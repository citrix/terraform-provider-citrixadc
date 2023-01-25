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

const testAccRnat6_nsip6_binding_basic = `

	resource "citrixadc_rnat6_nsip6_binding" "tf_rnat6_nsip6_binding" {
		name  = "my_rnat6"
		natip6 = "2001:db8:85a3::8a2e:370:7334"
	}
`

const testAccRnat6_nsip6_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion
`

func TestAccRnat6_nsip6_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckRnat6_nsip6_bindingDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccRnat6_nsip6_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRnat6_nsip6_bindingExist("citrixadc_rnat6_nsip6_binding.tf_rnat6_nsip6_binding", nil),
				),
			},
			resource.TestStep{
				Config: testAccRnat6_nsip6_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRnat6_nsip6_bindingNotExist("citrixadc_rnat6_nsip6_binding.tf_rnat6_nsip6_binding", "my_rnat6,2001:db8:85a3::8a2e:370:7334"),
				),
			},
		},
	})
}

func testAccCheckRnat6_nsip6_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No rnat6_nsip6_binding id is set")
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
		natip6 := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "rnat6_nsip6_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the matching natip6
		found := false
		for _, v := range dataArr {
			if v["natip6"].(string) == natip6 {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("rnat6_nsip6_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckRnat6_nsip6_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*NetScalerNitroClient).client

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}
		idSlice := strings.SplitN(id, ",", 2)

		name := idSlice[0]
		natip6 := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "rnat6_nsip6_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to hopefully not find the one with the matching natip6
		found := false
		for _, v := range dataArr {
			if v["natip6"].(string) == natip6 {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("rnat6_nsip6_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckRnat6_nsip6_bindingDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_rnat6_nsip6_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Rnat6_nsip6_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("rnat6_nsip6_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
