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
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"testing"
)

const testAccVpnglobal_intranetip6_binding_basic = `
	resource "citrixadc_vpnglobal_intranetip6_binding" "tf_bind" {
		intranetip6 = "2.3.4.5"
		numaddr     = "45"
	}
`

const testAccVpnglobal_intranetip6_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion
`

func TestAccVpnglobal_intranetip6_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckVpnglobal_intranetip6_bindingDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccVpnglobal_intranetip6_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnglobal_intranetip6_bindingExist("citrixadc_vpnglobal_intranetip6_binding.tf_bind", nil),
				),
			},
			resource.TestStep{
				Config: testAccVpnglobal_intranetip6_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnglobal_intranetip6_bindingNotExist("citrixadc_vpnglobal_intranetip6_binding.tf_bind", "2.3.4.5"),
				),
			},
		},
	})
}

func testAccCheckVpnglobal_intranetip6_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No vpnglobal_intranetip6_binding id is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		client := testAccProvider.Meta().(*NetScalerNitroClient).client

		intranetip6 := rs.Primary.ID

		findParams := service.FindParams{
			ResourceType:             "vpnglobal_intranetip6_binding",
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
			if v["intranetip6"].(string) == intranetip6 {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("vpnglobal_intranetip6_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckVpnglobal_intranetip6_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*NetScalerNitroClient).client
		intranetip6 := id
		findParams := service.FindParams{
			ResourceType:             "vpnglobal_intranetip6_binding",
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
			if v["intranetip6"].(string) == intranetip6 {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("vpnglobal_intranetip6_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckVpnglobal_intranetip6_bindingDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_vpnglobal_intranetip6_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Vpnglobal_intranetip6_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("vpnglobal_intranetip6_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
