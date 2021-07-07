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
	"strings"
	"testing"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
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
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckVlan_interface_bindingDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccVlan_interface_binding_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlan_interface_bindingExist("citrixadc_vlan_interface_binding.tf_bind", nil),
				),
			},
			resource.TestStep{
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

		client := testAccProvider.Meta().(*NetScalerNitroClient).client

		idSlice := strings.SplitN(rs.Primary.ID, ",", 2)

		vlanid := idSlice[0]
		ifnum := idSlice[1]

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

		client := testAccProvider.Meta().(*NetScalerNitroClient).client

		idSlice := strings.SplitN(bindingId, ",", 2)

		vlanid := idSlice[0]
		ifnum := idSlice[1]

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
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_vlan_interface_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Vlan_interface_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("vlan_interface_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
