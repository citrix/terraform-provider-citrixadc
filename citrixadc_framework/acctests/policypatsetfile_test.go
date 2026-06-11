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

import (
	"fmt"
	"testing"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// NOTE: policypatsetfile is an IMPORT-as-create resource. Create issues a
// NITRO Import action (POST ?action=Import); Read is a GET-by-name;
// Update is a no-op because EVERY write attribute is RequiresReplace; Delete
// is a plain DELETE /policypatsetfile/<name>. There is therefore no in-place
// update path, so the basic test below has a single create+verify step.
//
// `src` is a "local:" file that MUST exist on the appliance at import time.
// doPolicyPatSetFilePreChecks (PreCheck) uploads testdata/tftest.patset to
// /var/tmp on the ADC so the "local:tftest.patset" import resolves, mirroring
// the sibling appfwprotofile/apispecfile import tests.
// NOTE: `delimiter` is a single-character argument (CLI: `-delimiter
// <character>`). Its CLI "Default value: 10" is the internal code for the
// newline character, NOT a literal value you may pass. Passing "10" yields
// errorcode 1106 "String length exceeds maximum [10]". The testdata file is
// newline-delimited, so we omit `delimiter` and let the appliance use its
// newline default.
const testAccPolicypatsetfile_basic_step1 = `
resource "citrixadc_policypatsetfile" "tf_policypatsetfile" {
  name      = "tf_policypatsetfile"
  src       = "local:tftest.patset"
  charset   = "ASCII"
  comment   = "test_comment"
  overwrite = true
}

`

func TestAccPolicypatsetfile_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doPolicyPatSetFilePreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckPolicypatsetfileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPolicypatsetfile_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPolicypatsetfileExist("citrixadc_policypatsetfile.tf_policypatsetfile", nil),
					resource.TestCheckResourceAttr("citrixadc_policypatsetfile.tf_policypatsetfile", "name", "tf_policypatsetfile"),
					resource.TestCheckResourceAttr("citrixadc_policypatsetfile.tf_policypatsetfile", "src", "local:tftest.patset"),
					resource.TestCheckResourceAttr("citrixadc_policypatsetfile.tf_policypatsetfile", "charset", "ASCII"),
					resource.TestCheckResourceAttr("citrixadc_policypatsetfile.tf_policypatsetfile", "comment", "test_comment"),
				),
			},
		},
	})
}

func testAccCheckPolicypatsetfileExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No policypatsetfile name is set")
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
		// Imported patset files are only listed via the filtered GET
		// /policypatsetfile?args=imported:true, not by a plain GET-by-name.
		data, err := findImportedPatsetfileByNameInTest(client, rs.Primary.ID)
		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("policypatsetfile %s not found", n)
		}

		return nil
	}
}

// findImportedPatsetfileByNameInTest lists imported patset files
// (GET /policypatsetfile?args=imported:true) and returns the entry whose "name"
// matches, or (nil, nil) if not present.
func findImportedPatsetfileByNameInTest(client *service.NitroClient, name string) (map[string]interface{}, error) {
	findParams := service.FindParams{
		ResourceType: service.Policypatsetfile.Type(),
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

func testAccCheckPolicypatsetfileDestroy(s *terraform.State) error {
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_policypatsetfile" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		data, err := findImportedPatsetfileByNameInTest(client, rs.Primary.ID)
		if err != nil {
			continue
		}
		if data != nil {
			return fmt.Errorf("policypatsetfile %s still exists", rs.Primary.ID)
		}
	}

	return nil
}

// Datasource test - the data source reads the imported patset file by name.
// `depends_on` ensures the resource is imported before the data source reads it.
// doPolicyPatSetFilePreChecks (PreCheck) uploads testdata/tftest.patset so the
// "local:tftest.patset" import resolves.
const testAccPolicypatsetfileDataSource_basic = `

resource "citrixadc_policypatsetfile" "tf_policypatsetfile" {
  name      = "tf_policypatsetfile"
  src       = "local:tftest.patset"
  charset   = "ASCII"
  comment   = "test_comment"
  overwrite = true
}

data "citrixadc_policypatsetfile" "tf_policypatsetfile" {
  name       = citrixadc_policypatsetfile.tf_policypatsetfile.name
  depends_on = [citrixadc_policypatsetfile.tf_policypatsetfile]
}
`

func TestAccPolicypatsetfileDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doPolicyPatSetFilePreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccPolicypatsetfileDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_policypatsetfile.tf_policypatsetfile", "name", "tf_policypatsetfile"),
				),
			},
		},
	})
}
