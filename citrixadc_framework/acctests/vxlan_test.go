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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"testing"
)

const testAccVxlan_add = `
	resource "citrixadc_vlan" "tf_vlan" {
		vlanid    = 40
		aliasname = "Management VLAN"
	}
	resource "citrixadc_vxlan" "tf_vxlan" {
		vxlanid            = 123
		vlan               = citrixadc_vlan.tf_vlan.vlanid
		port               = 33
		dynamicrouting     = "DISABLED"
		ipv6dynamicrouting = "DISABLED"
		innervlantagging   = "ENABLED"
	}
`
const testAccVxlan_update = `
	resource "citrixadc_vlan" "tf_vlan" {
		vlanid    = 40
		aliasname = "Management VLAN"
	}
	resource "citrixadc_vxlan" "tf_vxlan" {
		vxlanid            = 123
		vlan               = citrixadc_vlan.tf_vlan.vlanid
		port               = 8080
		dynamicrouting     = "DISABLED"
		ipv6dynamicrouting = "DISABLED"
		innervlantagging   = "ENABLED"
	}
`

func TestAccVxlan_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVxlanDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVxlan_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVxlanExist("citrixadc_vxlan.tf_vxlan", nil),
					resource.TestCheckResourceAttr("citrixadc_vxlan.tf_vxlan", "vxlanid", "123"),
					resource.TestCheckResourceAttr("citrixadc_vxlan.tf_vxlan", "port", "33"),
				),
			},
			{
				Config: testAccVxlan_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVxlanExist("citrixadc_vxlan.tf_vxlan", nil),
					resource.TestCheckResourceAttr("citrixadc_vxlan.tf_vxlan", "vxlanid", "123"),
					resource.TestCheckResourceAttr("citrixadc_vxlan.tf_vxlan", "port", "8080"),
				),
			},
		},
	})
}

func testAccCheckVxlanExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No vxlan name is set")
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
		data, err := client.FindResource(service.Vxlan.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("vxlan %s not found", n)
		}

		return nil
	}
}

func testAccCheckVxlanDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_vxlan" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Vxlan.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("vxlan %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
