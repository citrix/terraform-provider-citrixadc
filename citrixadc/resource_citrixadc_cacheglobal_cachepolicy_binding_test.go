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

const testAccCacheglobal_cachepolicy_binding_basic = `


	resource "citrixadc_cachepolicy" "tf_cachepolicy" {
		policyname  = "my_cachepolicy"
		rule        = "true"
		action      = "CACHE"
	}
	resource "citrixadc_cacheglobal_cachepolicy_binding" "tf_cacheglobal_cachepolicy_binding" {
		policy   = citrixadc_cachepolicy.tf_cachepolicy.policyname
		priority = 100
		type     = "REQ_DEFAULT"
	}
  
`

const testAccCacheglobal_cachepolicy_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion
	
	resource "citrixadc_cachepolicy" "tf_cachepolicy" {
		policyname  = "my_cachepolicy"
		rule        = "true"
		action      = "CACHE"
	}

`

func TestAccCacheglobal_cachepolicy_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckCacheglobal_cachepolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCacheglobal_cachepolicy_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCacheglobal_cachepolicy_bindingExist("citrixadc_cacheglobal_cachepolicy_binding.tf_cacheglobal_cachepolicy_binding", nil),
				),
			},
			{
				Config: testAccCacheglobal_cachepolicy_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCacheglobal_cachepolicy_bindingNotExist("citrixadc_cacheglobal_cachepolicy_binding.tf_cacheglobal_cachepolicy_binding", "my_cachepolicy", "REQ_DEFAULT"),
				),
			},
		},
	})
}

func testAccCheckCacheglobal_cachepolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No cacheglobal_cachepolicy_binding id is set")
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

		policy := rs.Primary.ID

		findParams := service.FindParams{
			ResourceType:             "cacheglobal_cachepolicy_binding",
			ArgsMap:                  map[string]string{"type": rs.Primary.Attributes["type"]},
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the matching policy
		found := false
		for _, v := range dataArr {
			if v["policy"].(string) == policy {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("cacheglobal_cachepolicy_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckCacheglobal_cachepolicy_bindingNotExist(n string, id string, type_val string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		policy := id

		findParams := service.FindParams{
			ResourceType:             "cacheglobal_cachepolicy_binding",
			ArgsMap:                  map[string]string{"type": type_val},
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to hopefully not find the one with the matching policy
		found := false
		for _, v := range dataArr {
			if v["policy"].(string) == policy {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("cacheglobal_cachepolicy_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckCacheglobal_cachepolicy_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_cacheglobal_cachepolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Cacheglobal_cachepolicy_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("cacheglobal_cachepolicy_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
