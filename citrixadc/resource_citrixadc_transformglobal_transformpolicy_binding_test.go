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

const testAccTransformglobal_transformpolicy_binding_basic = `

resource "citrixadc_transformpolicy" "tf_trans_policy" {
	name        = "tf_trans_policy"
	profilename = "tf_trans_profile"
	rule        = "http.REQ.URL.CONTAINS(\"test_url\")"
  }
  resource "citrixadc_transformglobal_transformpolicy_binding" "transformglobal_transformpolicy_binding" {
	policyname = citrixadc_transformpolicy.tf_trans_policy.name
	priority   = 2
	type       = "REQ_DEFAULT"
  }
`

const testAccTransformglobal_transformpolicy_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion
	resource "citrixadc_transformpolicy" "tf_trans_policy" {
		name        = "tf_trans_policy"
		profilename = "tf_trans_profile"
		rule        = "http.REQ.URL.CONTAINS(\"test_url\")"
	  }
`

func TestAccTransformglobal_transformpolicy_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckTransformglobal_transformpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccTransformglobal_transformpolicy_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTransformglobal_transformpolicy_bindingExist("citrixadc_transformglobal_transformpolicy_binding.transformglobal_transformpolicy_binding", nil),
				),
			},
			resource.TestStep{
				Config: testAccTransformglobal_transformpolicy_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTransformglobal_transformpolicy_bindingNotExist("citrixadc_transformglobal_transformpolicy_binding.transformglobal_transformpolicy_binding","tf_trans_policy","REQ_DEFAULT"),
				),
			},
		},
	})
}

func testAccCheckTransformglobal_transformpolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No transformglobal_transformpolicy_binding id is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		client := testAccProvider.Meta().(*NetScalerNitroClient).client

		policyname := rs.Primary.ID
		typename := rs.Primary.Attributes["type"]
		findParams := service.FindParams{
			ResourceType:             "transformglobal_transformpolicy_binding",
			ArgsMap: 				  map[string]string{ "type":typename },
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
			if v["policyname"].(string) == policyname {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("transformglobal_transformpolicy_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckTransformglobal_transformpolicy_bindingNotExist(n string, id string,typename string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*NetScalerNitroClient).client
		policyname := id
		findParams := service.FindParams{
			ResourceType:             "transformglobal_transformpolicy_binding",
			ArgsMap: 				   map[string]string{ "type":typename },
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
			if v["policyname"].(string) == policyname {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("transformglobal_transformpolicy_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckTransformglobal_transformpolicy_bindingDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_transformglobal_transformpolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Transformglobal_transformpolicy_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("transformglobal_transformpolicy_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
