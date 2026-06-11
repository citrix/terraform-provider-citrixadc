/*
Copyright 2016 Citrix Systems, Inc

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package citrixadc

// NOTE: policyurlset is an IMPORT-as-create resource (mirrors policypatsetfile).
// Create issues a NITRO Import action (POST ?action=Import); Read is the filtered
// list GET /policyurlset?args=imported:true matched by name (a plain GET-by-name
// returns errorcode 258 "No such resource"); Update is a no-op because EVERY
// write attribute is RequiresReplace; Delete is a plain DELETE /policyurlset/<name>
// that removes the urlset but reports a spurious errorcode 258 (suppressed).
//
// `url` is a "local:" file that MUST exist on the appliance at import time.
// doPolicyUrlSetPreChecks (PreCheck) uploads testdata/tftest.urlset to /var/tmp
// on the ADC so the "local:tftest.urlset" import resolves, mirroring the sibling
// policypatsetfile/appfwprotofile/apispecfile import tests.
//
// NOTE: `delimiter` and `rowseparator` are single-character CLI arguments
// (`-delimiter <character>` / `-rowSeparator <character>`). Their CLI "Default
// value: 44"/"10" are the internal codes for comma/newline, NOT literal values
// you may pass; passing "44"/"10" yields errorcode 1106 "String length exceeds
// maximum". The testdata file uses the appliance defaults (comma fields, newline
// rows), so both are omitted and the appliance defaults are used.

import (
	"fmt"
	"testing"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// All policyurlset configurable attributes are RequiresReplace, and Update is a
// no-op. The basic test therefore only creates and verifies (no in-place update
// step). `url` is supplied via the plain (backward-compatible) Sensitive attribute.
const testAccPolicyurlset_basic_step1 = `
resource "citrixadc_policyurlset" "tf_policyurlset" {
  name                = "tf_policyurlset"
  url                 = "local:tftest.urlset"
  interval            = 0
  matchedid           = 2
  privateset          = false
  subdomainexactmatch = false
  comment             = "tf policyurlset acc test"
}

`

func TestAccPolicyurlset_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doPolicyUrlSetPreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckPolicyurlsetDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPolicyurlset_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPolicyurlsetExist("citrixadc_policyurlset.tf_policyurlset", nil),
					resource.TestCheckResourceAttr("citrixadc_policyurlset.tf_policyurlset", "name", "tf_policyurlset"),
					resource.TestCheckResourceAttr("citrixadc_policyurlset.tf_policyurlset", "interval", "0"),
					resource.TestCheckResourceAttr("citrixadc_policyurlset.tf_policyurlset", "matchedid", "2"),
					resource.TestCheckResourceAttr("citrixadc_policyurlset.tf_policyurlset", "subdomainexactmatch", "false"),
				),
			},
		},
	})
}

func testAccCheckPolicyurlsetExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No policyurlset name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		// Imported urlsets are only listed via the filtered GET
		// /policyurlset?args=imported:true, not by a plain GET-by-name.
		data, err := findImportedPolicyurlsetByNameInTest(client, rs.Primary.ID)
		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("policyurlset %s not found", n)
		}

		return nil
	}
}

// findImportedPolicyurlsetByNameInTest lists imported urlsets
// (GET /policyurlset?args=imported:true) and returns the entry whose "name"
// matches, or (nil, nil) if not present.
func findImportedPolicyurlsetByNameInTest(client *service.NitroClient, name string) (map[string]interface{}, error) {
	findParams := service.FindParams{
		ResourceType: service.Policyurlset.Type(),
		ArgsMap:      map[string]string{"imported": "true"},
	}
	dataArr, err := client.FindResourceArrayWithParams(findParams)
	if err != nil {
		return nil, err
	}
	for _, item := range dataArr {
		if val, ok := item["name"]; ok && val != nil {
			if val.(string) == name {
				return item, nil
			}
		}
	}
	return nil, nil
}

func testAccCheckPolicyurlsetDestroy(s *terraform.State) error {
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_policyurlset" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		data, err := findImportedPolicyurlsetByNameInTest(client, rs.Primary.ID)
		if err != nil {
			continue
		}
		if data != nil {
			return fmt.Errorf("policyurlset %s still exists", rs.Primary.ID)
		}
	}
	return nil
}

// Datasource lookup key is `name`. The imported-urlset list
// (GET /policyurlset?args=imported:true) only returns name/imported/patterncount/
// url for each entry — it does NOT echo back interval/matchedid/etc., and `url`
// is a write-only secret — so only `name` is asserted here. doPolicyUrlSetPreChecks
// (PreCheck) uploads testdata/tftest.urlset so the "local:tftest.urlset" import
// resolves.
const testAccPolicyurlsetDataSource_basic = `

resource "citrixadc_policyurlset" "tf_policyurlset" {
  name                = "tf_policyurlset_ds"
  url                 = "local:tftest.urlset"
  interval            = 0
  matchedid           = 2
  subdomainexactmatch = false
}

data "citrixadc_policyurlset" "tf_policyurlset" {
  name       = citrixadc_policyurlset.tf_policyurlset.name
  depends_on = [citrixadc_policyurlset.tf_policyurlset]
}
`

func TestAccPolicyurlsetDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doPolicyUrlSetPreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccPolicyurlsetDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_policyurlset.tf_policyurlset", "name", "tf_policyurlset_ds"),
				),
			},
		},
	})
}

// Write-only ephemeral path: `url_wo` + `url_wo_version`. Both `url_wo` and
// `url_wo_version` are RequiresReplace, so each step recreates the resource; the
// version still confirms the write-only path is exercised. `url_wo` is never
// stored in state and must NOT be asserted with TestCheckResourceAttr.
//
// The import SOURCE is a "local:" file that must exist on the appliance, so the
// write-only value is the same local: source as the basic test (uploaded by the
// PreCheck). url_wo_version is bumped between steps to drive the recreate.
const testAccPolicyurlset_url_wo_step1 = `

	variable "policyurlset_url_wo" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_policyurlset" "tf_policyurlset" {
		name                = "tf_policyurlset_wo"
		url_wo              = var.policyurlset_url_wo
		url_wo_version      = 1
		interval            = 0
		matchedid           = 2
		subdomainexactmatch = false
	}
`

const testAccPolicyurlset_url_wo_step2 = `

	 variable "policyurlset_url_wo_2" {
	  type      = string
	  sensitive = true
	}

	 resource "citrixadc_policyurlset" "tf_policyurlset" {
		name                = "tf_policyurlset_wo"
		url_wo              = var.policyurlset_url_wo_2
		url_wo_version      = 2
		interval            = 0
		matchedid           = 2
		subdomainexactmatch = false
	}
`

func TestAccPolicyurlset_url_wo_ephemeral(t *testing.T) {
	// Both steps point url_wo at the same uploaded local: source; the
	// url_wo_version bump (1 -> 2) is what triggers the RequiresReplace recreate
	// and confirms the write-only path is exercised.
	t.Setenv("TF_VAR_policyurlset_url_wo", "local:tftest.urlset")
	t.Setenv("TF_VAR_policyurlset_url_wo_2", "local:tftest.urlset")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doPolicyUrlSetPreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckPolicyurlsetDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPolicyurlset_url_wo_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPolicyurlsetExist("citrixadc_policyurlset.tf_policyurlset", nil),
					resource.TestCheckResourceAttr("citrixadc_policyurlset.tf_policyurlset", "name", "tf_policyurlset_wo"),
					resource.TestCheckResourceAttr("citrixadc_policyurlset.tf_policyurlset", "url_wo_version", "1"),
				),
			},
			{
				Config: testAccPolicyurlset_url_wo_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPolicyurlsetExist("citrixadc_policyurlset.tf_policyurlset", nil),
					resource.TestCheckResourceAttr("citrixadc_policyurlset.tf_policyurlset", "name", "tf_policyurlset_wo"),
					resource.TestCheckResourceAttr("citrixadc_policyurlset.tf_policyurlset", "url_wo_version", "2"),
				),
			},
		},
	})
}
