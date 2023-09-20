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

const testAccNetbridge_iptunnel_binding_basic = `
	resource "citrixadc_vxlanvlanmap" "tf_vxlanvlanmp" {
		name = "tf_vxlanvlanmp"
	}
	resource "citrixadc_nsip" "nsip" {
		ipaddress = "2.2.2.1"
		type      = "VIP"
		netmask   = "255.255.255.0"
	}
	resource "citrixadc_netbridge" "tf_netbridge" {
		name         = "tf_netbridge"
		vxlanvlanmap = citrixadc_vxlanvlanmap.tf_vxlanvlanmp.name
	}
	resource "citrixadc_iptunnel" "tf_iptunnel" {
		name             = "tf_iptunnel"
		remote           = "66.0.0.11"
		remotesubnetmask = "255.255.255.255"
		local            = citrixadc_nsip.nsip.ipaddress
		protocol         = "GRE"
	}
	resource "citrixadc_netbridge_iptunnel_binding" "tf_binding" {
		name   = citrixadc_netbridge.tf_netbridge.name
		tunnel = citrixadc_iptunnel.tf_iptunnel.name
	}
`

const testAccNetbridge_iptunnel_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion
	resource "citrixadc_vxlanvlanmap" "tf_vxlanvlanmp" {
		name = "tf_vxlanvlanmp"
	}
	resource "citrixadc_nsip" "nsip" {
		ipaddress = "2.2.2.1"
		type      = "VIP"
		netmask   = "255.255.255.0"
	}
	resource "citrixadc_netbridge" "tf_netbridge" {
		name         = "tf_netbridge"
		vxlanvlanmap = citrixadc_vxlanvlanmap.tf_vxlanvlanmp.name
	}
	resource "citrixadc_iptunnel" "tf_iptunnel" {
		name             = "tf_iptunnel"
		remote           = "66.0.0.11"
		remotesubnetmask = "255.255.255.255"
		local            = citrixadc_nsip.nsip.ipaddress
		protocol         = "GRE"
	}
`

func TestAccNetbridge_iptunnel_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNetbridge_iptunnel_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNetbridge_iptunnel_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetbridge_iptunnel_bindingExist("citrixadc_netbridge_iptunnel_binding.tf_binding", nil),
				),
			},
			{
				Config: testAccNetbridge_iptunnel_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetbridge_iptunnel_bindingNotExist("citrixadc_netbridge_iptunnel_binding.tf_binding", "tf_netbridge,tf_iptunnel"),
				),
			},
		},
	})
}

func testAccCheckNetbridge_iptunnel_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No netbridge_iptunnel_binding id is set")
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

		name := idSlice[0]
		tunnel := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "netbridge_iptunnel_binding",
			ResourceName:             name,
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
			if v["tunnel"].(string) == tunnel {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("netbridge_iptunnel_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckNetbridge_iptunnel_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*NetScalerNitroClient).client

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}
		idSlice := strings.SplitN(id, ",", 2)

		name := idSlice[0]
		tunnel := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "netbridge_iptunnel_binding",
			ResourceName:             name,
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
			if v["tunnel"].(string) == tunnel {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("netbridge_iptunnel_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckNetbridge_iptunnel_bindingDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_netbridge_iptunnel_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Netbridge_iptunnel_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("netbridge_iptunnel_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
