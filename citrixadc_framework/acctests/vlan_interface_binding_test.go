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
