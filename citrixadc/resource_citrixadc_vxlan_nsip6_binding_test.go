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

const testAccVxlan_nsip6_binding_basic = `
	resource "citrixadc_vxlan" "tf_vxlan" {
		vxlanid            = 123
		port               = 33
		dynamicrouting     = "DISABLED"
		ipv6dynamicrouting = "DISABLED"
		innervlantagging   = "ENABLED"
	}
	resource "citrixadc_nsip6" "test_nsip" {
		ipv6address = "2001:db8:100::fb/64"
		type        = "VIP"
		icmp        = "DISABLED"
	}
	resource "citrixadc_vxlan_nsip6_binding" "tf_binding" {
		vxlanid   = citrixadc_vxlan.tf_vxlan.vxlanid
		ipaddress = citrixadc_nsip6.test_nsip.ipv6address
		netmask   = "255.255.255.0"
	}
`

const testAccVxlan_nsip6_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion
	resource "citrixadc_vxlan" "tf_vxlan" {
		vxlanid            = 123
		port               = 33
		dynamicrouting     = "DISABLED"
		ipv6dynamicrouting = "DISABLED"
		innervlantagging   = "ENABLED"
	}
	resource "citrixadc_nsip6" "test_nsip" {
		ipv6address = "2001:db8:100::fb/64"
		type        = "VIP"
		icmp        = "DISABLED"
	}
`

func TestAccVxlan_nsip6_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckVxlan_nsip6_bindingDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccVxlan_nsip6_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVxlan_nsip6_bindingExist("citrixadc_vxlan_nsip6_binding.tf_binding", nil),
				),
			},
			resource.TestStep{
				Config: testAccVxlan_nsip6_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVxlan_nsip6_bindingNotExist("citrixadc_vxlan_nsip6_binding.tf_binding", "123,2001:db8:100::fb/64"),
				),
			},
		},
	})
}

func testAccCheckVxlan_nsip6_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No vxlan_nsip6_binding id is set")
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

		vxlanid := idSlice[0]
		ipaddress := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "vxlan_nsip6_binding",
			ResourceName:             vxlanid,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}
		// Resource is missing
		if len(dataArr) == 0 {
			return fmt.Errorf("Cannot find vlan_nsip_binding %s", bindingId)
		}

		// Iterate through results to find the one with the matching secondIdComponent
		foundIndex := -1
		for i, v := range dataArr {
			if v["ipaddress"].(string) == ipaddress {
				foundIndex = i
				break
			}
		}

		// Resource is missing
		if foundIndex == -1 {
			return fmt.Errorf("Resource missing vxlan_nsip6_binding %s", bindingId)
		}

		return nil
	}
}

func testAccCheckVxlan_nsip6_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*NetScalerNitroClient).client

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}
		idSlice := strings.SplitN(id, ",", 2)

		vxlanid := idSlice[0]
		ipaddress := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "vxlan_nsip6_binding",
			ResourceName:             vxlanid,
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

		// Iterate through results to find the one with the matching secondIdComponent
		foundIndex := -1
		for i, v := range dataArr {
			if v["ipaddress"].(string) == ipaddress {
				foundIndex = i
				break
			}
		}

		// Resource is missing
		if foundIndex == -1 {
			return fmt.Errorf("Resource missing vxlan_nsip6_binding %s", id)
		}

		return nil
	}
}

func testAccCheckVxlan_nsip6_bindingDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_vxlan_nsip6_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Vxlan_nsip6_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("vxlan_nsip6_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
