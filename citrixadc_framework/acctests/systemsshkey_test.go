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
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// NOTE: systemsshkey is an IMPORT-as-create resource. Create issues a NITRO
// Import action (POST ?action=Import); Read is a GET-all-and-match (no
// get-by-name); Update is a no-op because EVERY write attribute is
// RequiresReplace; Delete is DeleteResourceWithArgs(name, ["sshkeytype:..."]).
// There is therefore no in-place update path, so the basic test below has a
// single create+verify step.
//
// SAFETY: this test imports a DISPOSABLE PUBLIC key only (sshkeytype = PUBLIC).
// It MUST NOT import a PRIVATE host key, which would alter the appliance's SSH
// host identity.
//
// `src` is a "local:" file that MUST exist on the appliance at import time.
// doSystemSshKeyPreChecks (PreCheck) uploads testdata/tftest_sshkey.pub to
// /var/tmp on the ADC so the "local:tftest_sshkey.pub" import resolves,
// mirroring the sibling policypatsetfile/apispecfile import tests.
//
// `src` is NOT returned by GET, so it is never asserted below.
const testAccSystemsshkey_basic_step1 = `
resource "citrixadc_systemsshkey" "tf_systemsshkey" {
  name       = "tf_sshkey_test"
  src        = "local:tftest_sshkey.pub"
  sshkeytype = "PUBLIC"
}

`

func TestAccSystemsshkey_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doSystemSshKeyPreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSystemsshkeyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSystemsshkey_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSystemsshkeyExist("citrixadc_systemsshkey.tf_systemsshkey", nil),
					resource.TestCheckResourceAttr("citrixadc_systemsshkey.tf_systemsshkey", "name", "tf_sshkey_test"),
					resource.TestCheckResourceAttr("citrixadc_systemsshkey.tf_systemsshkey", "sshkeytype", "PUBLIC"),
				),
			},
		},
	})
}

// findSystemsshkeyInTest lists all ssh keys (GET /systemsshkey) and returns the
// entry whose name + sshkeytype match the composite ID, or (nil, nil) if not
// present. NITRO exposes only get(all) for this resource (no get-by-name).
func findSystemsshkeyInTest(client *service.NitroClient, id string) (map[string]interface{}, error) {
	idMap, _, err := utils.ParseIdString(id, []string{"name", "sshkeytype"}, nil)
	if err != nil {
		return nil, fmt.Errorf("Error parsing ID %q: %v", id, err)
	}

	findParams := service.FindParams{
		ResourceType:             service.Systemsshkey.Type(),
		ResourceMissingErrorCode: 258,
	}
	dataArr, err := client.FindResourceArrayWithParams(findParams)
	if err != nil {
		return nil, err
	}

	for _, item := range dataArr {
		nameMatch := false
		typeMatch := false
		if val, ok := item["name"].(string); ok && val == idMap["name"] {
			nameMatch = true
		}
		if val, ok := item["sshkeytype"].(string); ok && val == idMap["sshkeytype"] {
			typeMatch = true
		}
		if nameMatch && typeMatch {
			return item, nil
		}
	}
	return nil, nil
}

func testAccCheckSystemsshkeyExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No systemsshkey name is set")
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

		data, err := findSystemsshkeyInTest(client, rs.Primary.ID)
		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("systemsshkey %s not found", n)
		}

		return nil
	}
}

func testAccCheckSystemsshkeyDestroy(s *terraform.State) error {
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_systemsshkey" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		data, err := findSystemsshkeyInTest(client, rs.Primary.ID)
		if err != nil {
			continue
		}
		if data != nil {
			return fmt.Errorf("systemsshkey %s still exists", rs.Primary.ID)
		}
	}

	return nil
}

// Datasource test - the data source reads the imported ssh key by name +
// sshkeytype. `depends_on` ensures the resource is imported before the data
// source reads it. `src` is not returned by GET, so only name and sshkeytype
// are asserted.
const testAccSystemsshkeyDataSource_basic = `

resource "citrixadc_systemsshkey" "tf_systemsshkey" {
  name       = "tf_sshkey_test"
  src        = "local:tftest_sshkey.pub"
  sshkeytype = "PUBLIC"
}

data "citrixadc_systemsshkey" "tf_systemsshkey" {
  name       = citrixadc_systemsshkey.tf_systemsshkey.name
  sshkeytype = citrixadc_systemsshkey.tf_systemsshkey.sshkeytype
  depends_on = [citrixadc_systemsshkey.tf_systemsshkey]
}
`

func TestAccSystemsshkeyDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doSystemSshKeyPreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSystemsshkeyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSystemsshkeyDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_systemsshkey.tf_systemsshkey", "name", "tf_sshkey_test"),
					resource.TestCheckResourceAttr("data.citrixadc_systemsshkey.tf_systemsshkey", "sshkeytype", "PUBLIC"),
				),
			},
		},
	})
}
