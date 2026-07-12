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
	"log"
	"testing"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testAccVlan_interface_binding_basic_step1 = `
resource "citrixadc_vlan" "tf_vlan" {
    vlanid = 50
    aliasname = "Management VLAN"
}

resource "citrixadc_vlan_interface_binding" "tf_bind" {
    vlanid = citrixadc_vlan.tf_vlan.vlanid
    ifnum = "1/1"
}

`

const testAccVlan_interface_binding_basic_step2 = `
resource "citrixadc_vlan" "tf_vlan" {
    vlanid = 50
    aliasname = "Management VLAN"
}
`

func TestAccVlan_interface_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVlan_interface_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVlan_interface_binding_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlan_interface_bindingExist("citrixadc_vlan_interface_binding.tf_bind", nil),
				),
			},
			{
				Config: testAccVlan_interface_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlan_interface_bindingNotExist("50,1/1"),
				),
			},
		},
	})
}

func testAccCheckVlan_interface_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No vlan_interface_binding name is set")
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

		idMap, _, err := utils.ParseIdString(rs.Primary.ID, []string{"vlanid", "ifnum"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", rs.Primary.ID, err)
		}
		vlanid := idMap["vlanid"]
		ifnum := idMap["ifnum"]

		log.Printf("[DEBUG] citrixadc-provider: Reading vlan_interface_binding state %s", rs.Primary.ID)
		findParams := service.FindParams{
			ResourceType:             "vlan_interface_binding",
			ResourceName:             vlanid,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Resource is missing
		if len(dataArr) == 0 {
			return fmt.Errorf("Could not retrieve any entries for vlan_interface_binding %s", vlanid)
		}

		// Iterate through results to find the one with the right monitor name
		foundIndex := -1
		for i, v := range dataArr {
			if v["ifnum"].(string) == ifnum {
				foundIndex = i
				break
			}
		}

		// Resource is missing
		if foundIndex == -1 {
			return fmt.Errorf("Did not find vlan_interface_binding with id %s", rs.Primary.ID)
		}

		return nil
	}
}

func testAccCheckVlan_interface_bindingNotExist(bindingId string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		idMap, _, err := utils.ParseIdString(bindingId, []string{"vlanid", "ifnum"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", bindingId, err)
		}
		vlanid := idMap["vlanid"]
		ifnum := idMap["ifnum"]

		log.Printf("[DEBUG] citrixadc-provider: Reading vlan_interface_binding state %s", bindingId)
		findParams := service.FindParams{
			ResourceType:             "vlan_interface_binding",
			ResourceName:             vlanid,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Resource is missing
		if len(dataArr) == 0 {
			return nil
		}

		// Iterate through results to find the one with the right monitor name
		foundIndex := -1
		for i, v := range dataArr {
			if v["ifnum"].(string) == ifnum {
				foundIndex = i
				break
			}
		}

		// Resource is missing
		if foundIndex == -1 {
			return nil
		}

		return fmt.Errorf("vlan_interface_binding %s still exists", bindingId)
	}
}

func TestAccVlan_interface_binding_import(t *testing.T) {
	const resAddr = "citrixadc_vlan_interface_binding.tf_bind"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVlan_interface_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccVlan_interface_binding_basic_step1},
			{Config: testAccVlan_interface_binding_basic_step1, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
		},
	})
}

func testAccCheckVlan_interface_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_vlan_interface_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Vlan_interface_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("vlan_interface_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccVlan_interface_bindingDataSource_basic = `
resource "citrixadc_vlan" "tf_vlan" {
    vlanid = 50
    aliasname = "Management VLAN"
}

resource "citrixadc_vlan_interface_binding" "tf_bind" {
    vlanid = citrixadc_vlan.tf_vlan.vlanid
    ifnum = "1/1"
}

data "citrixadc_vlan_interface_binding" "tf_bind" {
	vlanid = 50
	ifnum  = "1/1"
	depends_on = [citrixadc_vlan_interface_binding.tf_bind]
}
`

// testAccVlan_interface_binding_upgrade_basic mirrors the _basic_step1 config
// (a vlan + an interface bound to it). It is valid under BOTH the SDK v2 2.2.0
// schema and the current framework schema, so it can be applied with the old
// provider in step 1 and re-planned with the new provider in step 2 of the
// state-upgrade test below.
const testAccVlan_interface_binding_upgrade_basic = `
resource "citrixadc_vlan" "tf_vlan" {
    vlanid = 50
    aliasname = "Management VLAN"
}

resource "citrixadc_vlan_interface_binding" "tf_bind" {
    vlanid = citrixadc_vlan.tf_vlan.vlanid
    ifnum = "1/1"
}
`

// TestAccVlan_interface_binding_sdkv2StateUpgrade verifies that a binding created
// by the LAST SDK v2 release (2.2.0) — which writes the legacy comma-joined id
// "vlanid,ifnum" (e.g. "50,1/1") — is refreshed and re-applied correctly by the
// CURRENT framework provider. Step 2 exercises ParseIdString on the legacy id
// during the framework Read.
//
// On this branch the framework RECOMPUTES the id on Read (SetAttrFromGet calls
// vlanInterfaceBindingComposeId, setting data.Id to the canonical new format), so
// after the step-2 refresh the id becomes "vlanid:50,ifnum:1%2F1".
func TestAccVlan_interface_binding_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckVlan_interface_bindingDestroy,
		Steps: []resource.TestStep{
			// Step 1: create with the last SDK v2 release from the registry. This
			// writes state carrying the LEGACY comma-joined id.
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.2.0",
					},
				},
				Config: testAccVlan_interface_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlan_interface_bindingExist("citrixadc_vlan_interface_binding.tf_bind", nil),
					resource.TestCheckResourceAttr("citrixadc_vlan_interface_binding.tf_bind", "id", "50,1/1"),
				),
			},
			// Step 2: same config through the CURRENT framework provider. Terraform
			// refreshes the legacy-id state through the framework Read (exercising
			// ParseIdString on the legacy id) then plans/applies. The framework
			// recomputes the id to the canonical new format.
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccVlan_interface_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlan_interface_bindingExist("citrixadc_vlan_interface_binding.tf_bind", nil),
					resource.TestCheckResourceAttr("citrixadc_vlan_interface_binding.tf_bind", "id", "vlanid:50,ifnum:1%2F1"),
				),
			},
		},
	})
}

func TestAccVlan_interface_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccVlan_interface_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_vlan_interface_binding.tf_bind", "vlanid", "50"),
					resource.TestCheckResourceAttr("data.citrixadc_vlan_interface_binding.tf_bind", "ifnum", "1/1"),
					resource.TestCheckResourceAttr("data.citrixadc_vlan_interface_binding.tf_bind", "tagged", "false"),
				),
			},
		},
	})
}
