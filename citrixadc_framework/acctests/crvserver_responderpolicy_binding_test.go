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
	"strings"
	"testing"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testAccCrvserver_responderpolicy_binding_basic = `

resource "citrixadc_crvserver" "crvserver" {
    name = "my_vserver"
    servicetype = "HTTP"
    arp = "OFF"
}
resource "citrixadc_responderpolicy" "tf_responderpolicy" {
	name    = "tf_responderpolicy1"
	action = "NOOP"
	rule = "HTTP.REQ.URL.PATH_AND_QUERY.CONTAINS(\"nosuchthing\")"
}
resource "citrixadc_crvserver_responderpolicy_binding" "crvserver_responderpolicy_binding" {
    name = citrixadc_crvserver.crvserver.name
    policyname = citrixadc_responderpolicy.tf_responderpolicy.name
    priority = 10
  
}
`

const testAccCrvserver_responderpolicy_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion
	resource "citrixadc_crvserver" "crvserver" {
		name = "my_vserver"
		servicetype = "HTTP"
		arp = "OFF"
	}
	resource "citrixadc_responderpolicy" "tf_responderpolicy" {
		name    = "tf_responderpolicy1"
		action = "NOOP"
		rule = "HTTP.REQ.URL.PATH_AND_QUERY.CONTAINS(\"nosuchthing\")"
	}
`

func TestAccCrvserver_responderpolicy_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCrvserver_responderpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCrvserver_responderpolicy_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCrvserver_responderpolicy_bindingExist("citrixadc_crvserver_responderpolicy_binding.crvserver_responderpolicy_binding", nil),
				),
			},
			{
				Config: testAccCrvserver_responderpolicy_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCrvserver_responderpolicy_bindingNotExist("citrixadc_crvserver_responderpolicy_binding.crvserver_responderpolicy_binding", "my_vserver,tf_responderpolicy1"),
				),
			},
		},
	})
}

func testAccCheckCrvserver_responderpolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No crvserver_responderpolicy_binding id is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		bindingId := rs.Primary.ID

		idMap, _, err := utils.ParseIdString(bindingId, []string{"name", "policyname"}, nil)
		if err != nil {
			return err
		}
		name := idMap["name"]
		policyname := idMap["policyname"]

		findParams := service.FindParams{
			ResourceType:             "crvserver_responderpolicy_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the matching policyname
		found := false
		for _, v := range dataArr {
			if v["policyname"].(string) == policyname {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("crvserver_responderpolicy_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckCrvserver_responderpolicy_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}
		idMap, _, err := utils.ParseIdString(id, []string{"name", "policyname"}, nil)
		if err != nil {
			return err
		}
		name := idMap["name"]
		policyname := idMap["policyname"]

		findParams := service.FindParams{
			ResourceType:             "crvserver_responderpolicy_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to hopefully not find the one with the matching policyname
		found := false
		for _, v := range dataArr {
			if v["policyname"].(string) == policyname {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("crvserver_responderpolicy_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckCrvserver_responderpolicy_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_crvserver_responderpolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Crvserver_responderpolicy_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("crvserver_responderpolicy_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccCrvserver_responderpolicy_bindingDataSource_basic = `

	resource "citrixadc_responderpolicy" "tf_responderpolicy" {
		name   = "tf_responderpolicy_ds"
		action = "NOOP"
		rule   = "HTTP.REQ.URL.PATH_AND_QUERY.CONTAINS(\"test_url\")"
	}
	resource "citrixadc_crvserver" "crvserver" {
		name        = "my_vserver_ds"
		servicetype = "HTTP"
		arp         = "OFF"
	}
	resource "citrixadc_crvserver_responderpolicy_binding" "crvserver_responderpolicy_binding" {
		name       = citrixadc_crvserver.crvserver.name
		policyname = citrixadc_responderpolicy.tf_responderpolicy.name
		priority   = 10
	}

	data "citrixadc_crvserver_responderpolicy_binding" "crvserver_responderpolicy_binding" {
		name       = citrixadc_crvserver_responderpolicy_binding.crvserver_responderpolicy_binding.name
		policyname = citrixadc_crvserver_responderpolicy_binding.crvserver_responderpolicy_binding.policyname
		depends_on = [citrixadc_crvserver_responderpolicy_binding.crvserver_responderpolicy_binding]
	}
`

func TestAcccrvserver_responderpolicy_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccCrvserver_responderpolicy_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_crvserver_responderpolicy_binding.crvserver_responderpolicy_binding", "name", "my_vserver_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_crvserver_responderpolicy_binding.crvserver_responderpolicy_binding", "policyname", "tf_responderpolicy_ds"),
				),
			},
		},
	})
}

// Config for the SDK v2 -> Framework state-upgrade test. It is identical to the
// _basic config (same terraform resource labels, same values) so it is valid under
// BOTH the last SDK v2 release (2.2.0) schema in step 1 and the current framework
// schema in step 2.
const testAccCrvserver_responderpolicy_binding_upgrade_basic = `

resource "citrixadc_crvserver" "crvserver" {
    name = "my_vserver"
    servicetype = "HTTP"
    arp = "OFF"
}
resource "citrixadc_responderpolicy" "tf_responderpolicy" {
	name    = "tf_responderpolicy1"
	action = "NOOP"
	rule = "HTTP.REQ.URL.PATH_AND_QUERY.CONTAINS(\"nosuchthing\")"
}
resource "citrixadc_crvserver_responderpolicy_binding" "crvserver_responderpolicy_binding" {
    name = citrixadc_crvserver.crvserver.name
    policyname = citrixadc_responderpolicy.tf_responderpolicy.name
    priority = 10

}
`

// TestAccCrvserver_responderpolicy_binding_sdkv2StateUpgrade verifies that state
// written by the last SDK v2 provider release (2.2.0), which stores the legacy
// comma-joined ID "name,policyname", is refreshed cleanly through the current
// Framework provider. On Read the Framework recomputes the ID to the canonical
// new "key:value" format, so after the step-2 apply the ID becomes
// "name:<name>,policyname:<policyname>".
func TestAccCrvserver_responderpolicy_binding_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckCrvserver_responderpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			// Step 1: create with the last SDK v2 release; state gets the legacy comma ID.
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.2.0",
					},
				},
				Config: testAccCrvserver_responderpolicy_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCrvserver_responderpolicy_bindingExist("citrixadc_crvserver_responderpolicy_binding.crvserver_responderpolicy_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_crvserver_responderpolicy_binding.crvserver_responderpolicy_binding", "id", "my_vserver,tf_responderpolicy1"),
				),
			},
			// Step 2: refresh/apply the legacy-ID state through the current Framework
			// provider. Read exercises ParseIdString on the legacy ID and recomputes
			// the canonical new-format ID.
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccCrvserver_responderpolicy_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCrvserver_responderpolicy_bindingExist("citrixadc_crvserver_responderpolicy_binding.crvserver_responderpolicy_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_crvserver_responderpolicy_binding.crvserver_responderpolicy_binding", "id", "name:my_vserver,policyname:tf_responderpolicy1"),
				),
			},
		},
	})
}

func TestAccCrvserver_responderpolicy_binding_import(t *testing.T) {
	const resAddr = "citrixadc_crvserver_responderpolicy_binding.crvserver_responderpolicy_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCrvserver_responderpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccCrvserver_responderpolicy_binding_basic},
			{Config: testAccCrvserver_responderpolicy_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
		},
	})
}
