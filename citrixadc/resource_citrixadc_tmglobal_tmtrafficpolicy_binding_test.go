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

const testAccTmglobal_tmtrafficpolicy_binding_basic = `

	resource "citrixadc_tmtrafficaction" "tf_tmtrafficaction" {
		name             = "my_trafficaction"
		apptimeout       = 5
		sso              = "OFF"
		persistentcookie = "ON"
	}
	resource "citrixadc_tmtrafficpolicy" "tf_tmtrafficpolicy" {
		name   = "my_tmtrafficpolicy"
		rule   = "true"
		action = citrixadc_tmtrafficaction.tf_tmtrafficaction.name
	}

	resource "citrixadc_tmglobal_tmtrafficpolicy_binding" "tf_tmglobal_tmtrafficpolicy_binding" {
		priority 	= "100"
		policyname = citrixadc_tmtrafficpolicy.tf_tmtrafficpolicy.name
	}
`

const testAccTmglobal_tmtrafficpolicy_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion

	resource "citrixadc_tmtrafficaction" "tf_tmtrafficaction" {
		name             = "my_trafficaction"
		apptimeout       = 5
		sso              = "OFF"
		persistentcookie = "ON"
	}
	resource "citrixadc_tmtrafficpolicy" "tf_tmtrafficpolicy" {
		name   = "my_tmtrafficpolicy"
		rule   = "true"
		action = citrixadc_tmtrafficaction.tf_tmtrafficaction.name
	}
`

func TestAccTmglobal_tmtrafficpolicy_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckTmglobal_tmtrafficpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTmglobal_tmtrafficpolicy_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTmglobal_tmtrafficpolicy_bindingExist("citrixadc_tmglobal_tmtrafficpolicy_binding.tf_tmglobal_tmtrafficpolicy_binding", nil),
				),
			},
			{
				Config: testAccTmglobal_tmtrafficpolicy_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTmglobal_tmtrafficpolicy_bindingNotExist("citrixadc_tmglobal_tmtrafficpolicy_binding.tf_tmglobal_tmtrafficpolicy_binding", "my_tmtrafficpolicy"),
				),
			},
		},
	})
}

func testAccCheckTmglobal_tmtrafficpolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No tmglobal_tmtrafficpolicy_binding id is set")
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

		policyname := rs.Primary.ID

		findParams := service.FindParams{
			ResourceType:             "tmglobal_tmtrafficpolicy_binding",
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
			return fmt.Errorf("tmglobal_tmtrafficpolicy_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckTmglobal_tmtrafficpolicy_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		policyname := id

		findParams := service.FindParams{
			ResourceType:             "tmglobal_tmtrafficpolicy_binding",
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
			return fmt.Errorf("tmglobal_tmtrafficpolicy_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckTmglobal_tmtrafficpolicy_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_tmglobal_tmtrafficpolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Tmglobal_tmtrafficpolicy_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("tmglobal_tmtrafficpolicy_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
