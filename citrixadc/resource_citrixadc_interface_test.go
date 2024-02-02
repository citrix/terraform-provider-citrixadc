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
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccInterface_basic(t *testing.T) {
	if isCpxRun {
		t.Skip("skipping test CPX has different interface numbering")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccInterface_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckInterfaceExist("citrixadc_interface.tf_interface", nil, "1/1"),
				),
			},
			{
				Config: testAccInterface_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckInterfaceExist("citrixadc_interface.tf_interface", nil, "1/1"),
				),
			},
		},
	})
}

func testAccCheckInterfaceExist(n string, id *string, interfaceId string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lb vserver name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

		array, _ := nsClient.FindAllResources(service.Interface.Type())

		// Iterate over the retrieved addresses to find the particular interface id
		foundInterface := false
		for _, item := range array {
			if item["id"] == interfaceId {
				foundInterface = true
				break
			}
		}
		if !foundInterface {
			return fmt.Errorf("Could not find interface %v", interfaceId)
		}

		return nil
	}
}

const testAccInterface_basic_step1 = `
resource "citrixadc_interface" "tf_interface" {
    interface_id = "1/1"
    hamonitor = "OFF"
    mtu = 2000
}
`

const testAccInterface_basic_step2 = `
resource "citrixadc_interface" "tf_interface" {
    interface_id = "1/1"
    hamonitor = "ON"
    mtu = 1500
}
`
