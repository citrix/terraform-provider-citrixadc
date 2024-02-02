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
	"testing"
)

const testAccContentinspectionglobal_contentinspectionpolicy_binding_basic = `

	resource "citrixadc_contentinspectionpolicy" "tf_contentinspectionpolicy" {
		name   = "my_ci_policy"
		rule   = "false"
		action = "DROP"
	}

	resource "citrixadc_contentinspectionglobal_contentinspectionpolicy_binding" "tf_ci_binding" {
		policyname = citrixadc_contentinspectionpolicy.tf_contentinspectionpolicy.name
		priority   = 100
	}
`

const testAccContentinspectionglobal_contentinspectionpolicy_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion

	resource "citrixadc_contentinspectionpolicy" "tf_contentinspectionpolicy" {
		name   = "my_ci_policy"
		rule   = "false"
		action = "DROP"
	}
`

func TestAccContentinspectionglobal_contentinspectionpolicy_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckContentinspectionglobal_contentinspectionpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccContentinspectionglobal_contentinspectionpolicy_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckContentinspectionglobal_contentinspectionpolicy_bindingExist("citrixadc_contentinspectionglobal_contentinspectionpolicy_binding.tf_ci_binding", nil),
				),
			},
			{
				Config: testAccContentinspectionglobal_contentinspectionpolicy_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckContentinspectionglobal_contentinspectionpolicy_bindingNotExist("citrixadc_contentinspectionglobal_contentinspectionpolicy_binding.tf_ci_binding", "my_ci_policy", "REQ_DEFAULT"),
				),
			},
		},
	})
}

func testAccCheckContentinspectionglobal_contentinspectionpolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No contentinspectionglobal_contentinspectionpolicy_binding id is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		client := testAccProvider.Meta().(*NetScalerNitroClient).client

		policyname := rs.Primary.ID

		findParams := service.FindParams{
			ResourceType:             "contentinspectionglobal_contentinspectionpolicy_binding",
			ArgsMap:                  map[string]string{"type": "REQ_DEFAULT"},
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the matching policyname
		found := false
		for _, v := range dataArr {
			if v["policyname"].(string) == policyname {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("contentinspectionglobal_contentinspectionpolicy_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckContentinspectionglobal_contentinspectionpolicy_bindingNotExist(n string, id string, type_val string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*NetScalerNitroClient).client

		policyname := id

		findParams := service.FindParams{
			ResourceType:             "contentinspectionglobal_contentinspectionpolicy_binding",
			ArgsMap:                  map[string]string{"type": type_val},
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to hopefully not find the one with the matching policyname
		found := false
		for _, v := range dataArr {
			if v["policyname"].(string) == policyname {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("contentinspectionglobal_contentinspectionpolicy_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckContentinspectionglobal_contentinspectionpolicy_bindingDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_contentinspectionglobal_contentinspectionpolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource("contentinspectionglobal_contentinspectionpolicy_binding", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("contentinspectionglobal_contentinspectionpolicy_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
