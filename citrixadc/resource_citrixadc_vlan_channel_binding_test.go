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
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"strings"
	"testing"
)

const testAccVlan_channel_binding_basic = `

	resource "citrixadc_vlan_channel_binding" "tf_vlan_channel_binding" {
		vlanid = 2
		ifnum  = "LA/2"
		tagged = false
	}
	
`

const testAccVlan_channel_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion
`

func TestAccVlan_channel_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckVlan_channel_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVlan_channel_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlan_channel_bindingExist("citrixadc_vlan_channel_binding.tf_vlan_channel_binding", nil),
				),
			},
			{
				Config: testAccVlan_channel_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlan_channel_bindingNotExist("citrixadc_vlan_channel_binding.tf_vlan_channel_binding", "2,LA/2"),
				),
			},
		},
	})
}

func testAccCheckVlan_channel_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No vlan_channel_binding id is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		client := testAccProvider.Meta().(*NetScalerNitroClient).client

		bindingId := rs.Primary.ID

		idSlice := strings.SplitN(bindingId, ",", 2)

		vlanid := idSlice[0]
		ifnum := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "vlan_channel_binding",
			ResourceName:             vlanid,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the matching ifnum
		found := false
		for _, v := range dataArr {
			if v["ifnum"].(string) == ifnum {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("vlan_channel_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckVlan_channel_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*NetScalerNitroClient).client

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}
		idSlice := strings.SplitN(id, ",", 2)

		vlanid := idSlice[0]
		ifnum := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "vlan_channel_binding",
			ResourceName:             vlanid,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to hopefully not find the one with the matching ifnum
		found := false
		for _, v := range dataArr {
			if v["ifnum"].(string) == ifnum {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("vlan_channel_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckVlan_channel_bindingDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_vlan_channel_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Vlan_channel_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("vlan_channel_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
