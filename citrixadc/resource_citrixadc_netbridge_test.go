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

const testAccNetbridge_add = `
	resource "citrixadc_vxlanvlanmap" "tf_vxlanvlanmp" {
		name = "tf_vxlanvlanmp"
	}
	resource "citrixadc_vxlanvlanmap" "tf_vxlanvlanmp1" {
		name = "tf_vxlanvlanmpsample"
	}
	resource "citrixadc_netbridge" "tf_netbridge" {
		name         = "tf_netbridge"
		vxlanvlanmap = citrixadc_vxlanvlanmap.tf_vxlanvlanmp.name
	}
`
const testAccNetbridge_update = `
	resource "citrixadc_vxlanvlanmap" "tf_vxlanvlanmp" {
		name = "tf_vxlanvlanmp"
	}
	resource "citrixadc_vxlanvlanmap" "tf_vxlanvlanmp1" {
		name = "tf_vxlanvlanmpsample"
	}
	resource "citrixadc_netbridge" "tf_netbridge" {
		name         = "tf_netbridge"
		vxlanvlanmap = citrixadc_vxlanvlanmap.tf_vxlanvlanmp1.name
	}
`

func TestAccNetbridge_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckNetbridgeDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNetbridge_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetbridgeExist("citrixadc_netbridge.tf_netbridge", nil),
					resource.TestCheckResourceAttr("citrixadc_netbridge.tf_netbridge", "name", "tf_netbridge"),
					resource.TestCheckResourceAttr("citrixadc_netbridge.tf_netbridge", "vxlanvlanmap", "tf_vxlanvlanmp"),
				),
			},
			{
				Config: testAccNetbridge_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetbridgeExist("citrixadc_netbridge.tf_netbridge", nil),
					resource.TestCheckResourceAttr("citrixadc_netbridge.tf_netbridge", "name", "tf_netbridge"),
					resource.TestCheckResourceAttr("citrixadc_netbridge.tf_netbridge", "vxlanvlanmap", "tf_vxlanvlanmpsample"),
				),
			},
		},
	})
}

func testAccCheckNetbridgeExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No netbridge name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		// Use the shared utility function to get a configured client
		client, err := testAccGetClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Netbridge.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("netbridge %s not found", n)
		}

		return nil
	}
}

func testAccCheckNetbridgeDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_netbridge" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Netbridge.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("netbridge %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
