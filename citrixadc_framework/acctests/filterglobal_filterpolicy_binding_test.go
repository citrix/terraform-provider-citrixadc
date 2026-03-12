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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testAccFilterglobal_filterpolicy_binding_basic_step1 = `
resource "citrixadc_filterpolicy" "tf_filterpolicy" {
    name = "tf_filterpolicy"
    reqaction = "DROP"
    rule = "REQ.HTTP.URL CONTAINS http://abcd.com"
}

resource "citrixadc_filterglobal_filterpolicy_binding" "tf_filterglobal" {
    policyname = citrixadc_filterpolicy.tf_filterpolicy.name
    priority = 200
    state = "ENABLED"
}

`

const testAccFilterglobal_filterpolicy_binding_basic_step2 = `
resource "citrixadc_filterpolicy" "tf_filterpolicy" {
    name = "tf_filterpolicy"
    reqaction = "DROP"
    rule = "REQ.HTTP.URL CONTAINS http://abcd.com"
}

resource "citrixadc_filterglobal_filterpolicy_binding" "tf_filterglobal" {
    policyname = citrixadc_filterpolicy.tf_filterpolicy.name
    priority = 100
    state = "DISABLED"
}

`

const testAccFilterglobal_filterpolicy_binding_basic_step3 = `
resource "citrixadc_filterpolicy" "tf_filterpolicy" {
    name = "tf_filterpolicy"
    reqaction = "DROP"
    rule = "REQ.HTTP.URL CONTAINS http://abcd.com"
}

`

func TestAccFilterglobal_filterpolicy_binding_basic(t *testing.T) {
	t.Skipf("filterpolicy is not supported in 13.1")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckFilterglobal_filterpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccFilterglobal_filterpolicy_binding_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFilterglobal_filterpolicy_bindingExist("citrixadc_filterglobal_filterpolicy_binding.tf_filterglobal", nil),
				),
			},
			{
				Config: testAccFilterglobal_filterpolicy_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFilterglobal_filterpolicy_bindingExist("citrixadc_filterglobal_filterpolicy_binding.tf_filterglobal", nil),
				),
			},
			{
				Config: testAccFilterglobal_filterpolicy_binding_basic_step3,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFilterglobal_filterpolicy_bindingNotExist("citrixadc_filterglobal_filterpolicy_binding.tf_filterglobal", "tf_filterpolicy"),
				),
			},
		},
	})
}

func testAccCheckFilterglobal_filterpolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No filterglobal_filterpolicy_binding id is set")
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

		policyname := rs.Primary.ID

		findParams := service.FindParams{
			ResourceType: "filterglobal_filterpolicy_binding",
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
			return fmt.Errorf("filterglobal_filterpolicy_binding %s not found", policyname)
		}

		return nil
	}
}

func testAccCheckFilterglobal_filterpolicy_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		policyname := id

		findParams := service.FindParams{
			ResourceType: "filterglobal_filterpolicy_binding",
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
			return fmt.Errorf("filterglobal_filterpolicy_binding %s was found, but it should have been destroyed", policyname)
		}

		return nil
	}
}

func testAccCheckFilterglobal_filterpolicy_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_filterglobal_filterpolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		policyname := rs.Primary.ID

		findParams := service.FindParams{
			ResourceType: "filterglobal_filterpolicy_binding",
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
			return fmt.Errorf("filterglobal_filterpolicy_binding %s was found, but it should have been destroyed", policyname)
		}

	}

	return nil
}
