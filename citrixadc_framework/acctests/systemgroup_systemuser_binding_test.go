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

const testAccSystemgroup_systemuser_binding_basic = `

	resource "citrixadc_systemgroup" "tf_systemgroup" {
		groupname    = "tf_systemgroup"
		timeout      = 999
		promptstring = "bye>"
	}

	resource "citrixadc_systemuser" "tf_user" {
		username = "tf_user"
		password = "tf_password"
		timeout  = 200
	}

	resource "citrixadc_systemgroup_systemuser_binding" "tf_bind" {
		groupname = citrixadc_systemgroup.tf_systemgroup.groupname
		username  = citrixadc_systemuser.tf_user.username
	}
`

const testAccSystemgroup_systemuser_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion

	resource "citrixadc_systemgroup" "tf_systemgroup" {
		groupname    = "tf_systemgroup"
		timeout      = 999
		promptstring = "bye>"
	}

	resource "citrixadc_systemuser" "tf_user" {
		username = "tf_user"
		password = "tf_password"
		timeout  = 200
	}

`

func TestAccSystemgroup_systemuser_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSystemgroup_systemuser_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSystemgroup_systemuser_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSystemgroup_systemuser_bindingExist("citrixadc_systemgroup_systemuser_binding.tf_bind", nil),
				),
			},
			{
				Config: testAccSystemgroup_systemuser_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSystemgroup_systemuser_bindingNotExist("citrixadc_systemgroup_systemuser_binding.tf_bind", "tf_systemgroup,tf_user"),
				),
			},
		},
	})
}

func TestAccSystemgroup_systemuser_binding_import(t *testing.T) {
	const resAddr = "citrixadc_systemgroup_systemuser_binding.tf_bind"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSystemgroup_systemuser_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccSystemgroup_systemuser_binding_basic},
			{Config: testAccSystemgroup_systemuser_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
		},
	})
}

func testAccCheckSystemgroup_systemuser_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No systemgroup_systemuser_binding id is set")
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

		idMap, _, err := utils.ParseIdString(bindingId, []string{"groupname", "username"}, nil)
		if err != nil {
			return err
		}
		groupname := idMap["groupname"]
		username := idMap["username"]

		findParams := service.FindParams{
			ResourceType:             "systemgroup_systemuser_binding",
			ResourceName:             groupname,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the matching secondIdComponent
		found := false
		for _, v := range dataArr {
			if v["username"].(string) == username {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("systemgroup_systemuser_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckSystemgroup_systemuser_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		idMap, _, err := utils.ParseIdString(id, []string{"groupname", "username"}, nil)
		if err != nil {
			return err
		}
		groupname := idMap["groupname"]
		username := idMap["username"]

		findParams := service.FindParams{
			ResourceType:             "systemgroup_systemuser_binding",
			ResourceName:             groupname,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to hopefully not find the one with the matching secondIdComponent
		found := false
		for _, v := range dataArr {
			if v["username"].(string) == username {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("systemgroup_systemuser_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckSystemgroup_systemuser_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_systemgroup_systemuser_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Systemgroup_systemuser_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("systemgroup_systemuser_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccSystemgroup_systemuser_bindingDataSource_basic = `

	resource "citrixadc_systemgroup" "tf_systemgroup" {
		groupname    = "tf_systemgroup"
		timeout      = 999
		promptstring = "bye>"
	}

	resource "citrixadc_systemuser" "tf_user" {
		username = "tf_user"
		password = "tf_password"
		timeout  = 200
	}

	resource "citrixadc_systemgroup_systemuser_binding" "tf_bind" {
		groupname = citrixadc_systemgroup.tf_systemgroup.groupname
		username  = citrixadc_systemuser.tf_user.username
	}

	data "citrixadc_systemgroup_systemuser_binding" "tf_bind" {
		groupname = citrixadc_systemgroup_systemuser_binding.tf_bind.groupname
		username  = citrixadc_systemgroup_systemuser_binding.tf_bind.username
		depends_on = [citrixadc_systemgroup_systemuser_binding.tf_bind]
	}
`

func TestAccSystemgroup_systemuser_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSystemgroup_systemuser_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSystemgroup_systemuser_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_systemgroup_systemuser_binding.tf_bind", "groupname", "tf_systemgroup"),
					resource.TestCheckResourceAttr("data.citrixadc_systemgroup_systemuser_binding.tf_bind", "username", "tf_user"),
				),
			},
		},
	})
}

// testAccSystemgroup_systemuser_binding_upgrade_basic reuses the _basic config
// (binding + all prerequisite resources). It is valid under BOTH the SDK v2 2.2.0
// schema and the current Framework schema because the migration restored the SDK v2
// attribute names.
const testAccSystemgroup_systemuser_binding_upgrade_basic = `

	resource "citrixadc_systemgroup" "tf_systemgroup" {
		groupname    = "tf_systemgroup"
		timeout      = 999
		promptstring = "bye>"
	}

	resource "citrixadc_systemuser" "tf_user" {
		username = "tf_user"
		password = "tf_password"
		timeout  = 200
	}

	resource "citrixadc_systemgroup_systemuser_binding" "tf_bind" {
		groupname = citrixadc_systemgroup.tf_systemgroup.groupname
		username  = citrixadc_systemuser.tf_user.username
	}
`

// TestAccSystemgroup_systemuser_binding_sdkv2StateUpgrade verifies that state
// written by the last SDK v2 release (legacy comma-separated ID) is correctly
// upgraded when the same config is subsequently managed by the current Framework
// provider. Step 1 creates the binding with citrix/citrixadc 2.2.0 (writes the
// legacy id "tf_systemgroup,tf_user"). Step 2 refreshes/plans/applies the same
// config through the Framework provider, exercising ParseIdString on the legacy id;
// because the Framework recomputes the id on Read (SetAttrFromGet), the id upgrades
// to the new "key:value" form.
func TestAccSystemgroup_systemuser_binding_sdkv2StateUpgrade(t *testing.T) {
	resourceAddr := "citrixadc_systemgroup_systemuser_binding.tf_bind"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckSystemgroup_systemuser_bindingDestroy,
		Steps: []resource.TestStep{
			// Step 1: create with the last SDK v2 release -> state carries the legacy id.
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.2.0",
					},
				},
				Config: testAccSystemgroup_systemuser_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSystemgroup_systemuser_bindingExist(resourceAddr, nil),
					resource.TestCheckResourceAttr(resourceAddr, "id", "tf_systemgroup,tf_user"),
				),
			},
			// Step 2: refresh/plan/apply the SAME config through the current Framework
			// provider. The legacy-id state is read via ParseIdString and the id is
			// recomputed to the new key:value format.
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccSystemgroup_systemuser_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSystemgroup_systemuser_bindingExist(resourceAddr, nil),
					resource.TestCheckResourceAttr(resourceAddr, "id", "groupname:tf_systemgroup,username:tf_user"),
				),
			},
		},
	})
}
