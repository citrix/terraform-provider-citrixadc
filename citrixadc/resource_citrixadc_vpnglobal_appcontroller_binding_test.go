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
	"log"
	"net/url"
	"testing"
)

const testAccVpnglobal_appcontroller_binding_basic = `

	resource "citrixadc_vpnglobal_appcontroller_binding" "tf_vpnglobal_appcontroller_binding" {
		appcontroller = "http://www.citrix.com"
	}

`

func TestAccVpnglobal_appcontroller_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckVpnglobal_appcontroller_bindingDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccVpnglobal_appcontroller_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnglobal_appcontroller_bindingExist("citrixadc_vpnglobal_appcontroller_binding.tf_vpnglobal_appcontroller_binding", nil),
				),
			},
		},
	})
}

func testAccCheckVpnglobal_appcontroller_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No vpnglobal_appcontroller_binding id is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		client := testAccProvider.Meta().(*NetScalerNitroClient).client

		appcontroller, _ := url.QueryUnescape(rs.Primary.ID)

		findParams := service.FindParams{
			ResourceType: "vpnglobal_appcontroller_binding",
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			log.Printf("[DEBUG] citrixadc-provider: Error during FindResourceArrayWithParams %s", err.Error())
			return err
		}

		// Resource is missing
		if len(dataArr) == 0 {
			log.Printf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams returned empty array")
			return nil
		}

		// Iterate through results to find the one with the right id
		foundIndex := -1
		for i, v := range dataArr {
			if v["appcontroller"].(string) == appcontroller {
				foundIndex = i
				break
			}
		}

		if foundIndex == -1 {
			return fmt.Errorf("vpnglobal_appcontroller_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckVpnglobal_appcontroller_bindingDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_vpnglobal_appcontroller_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Vpnglobal_appcontroller_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("vpnglobal_appcontroller_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
