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

const testAccDnsglobal_dnspolicy_binding_basic = `

resource "citrixadc_dnspolicy" "dnspolicy" {
	name = "policy_A"
	rule = "CLIENT.IP.SRC.IN_SUBNET(1.1.1.1/24)"
	drop = "YES"
	}
  resource "citrixadc_dnsglobal_dnspolicy_binding" "dnsglobal_dnspolicy_binding" {
	policyname = citrixadc_dnspolicy.dnspolicy.name
	priority   = 30
	type       = "REQ_DEFAULT"
	}
`

const testAccDnsglobal_dnspolicy_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion
	resource "citrixadc_dnspolicy" "dnspolicy" {
		name = "policy_A"
		rule = "CLIENT.IP.SRC.IN_SUBNET(1.1.1.1/24)"
		drop = "YES"
	}
`

func TestAccDnsglobal_dnspolicy_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckDnsglobal_dnspolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDnsglobal_dnspolicy_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDnsglobal_dnspolicy_bindingExist("citrixadc_dnsglobal_dnspolicy_binding.dnsglobal_dnspolicy_binding", nil),
				),
			},
			{
				Config: testAccDnsglobal_dnspolicy_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDnsglobal_dnspolicy_bindingNotExist("citrixadc_dnsglobal_dnspolicy_binding.tf_binding", "policy_A", "REQ_DEFAULT"),
				),
			},
		},
	})
}

func testAccCheckDnsglobal_dnspolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No dnsglobal_dnspolicy_binding id is set")
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
		typename := rs.Primary.Attributes["type"]
		findParams := service.FindParams{
			ResourceType:             "dnsglobal_dnspolicy_binding",
			ArgsMap:                  map[string]string{"type": typename},
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
			return fmt.Errorf("dnsglobal_dnspolicy_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckDnsglobal_dnspolicy_bindingNotExist(n string, id string, typename string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		policyname := id

		findParams := service.FindParams{
			ResourceType:             "dnsglobal_dnspolicy_binding",
			ArgsMap:                  map[string]string{"type": typename},
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
			return fmt.Errorf("dnsglobal_dnspolicy_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckDnsglobal_dnspolicy_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_dnsglobal_dnspolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Dnsglobal_dnspolicy_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("dnsglobal_dnspolicy_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
