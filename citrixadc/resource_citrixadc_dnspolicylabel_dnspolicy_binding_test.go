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
	"strings"
	"testing"
)

const testAccDnspolicylabel_dnspolicy_binding_basic = `

resource "citrixadc_dnspolicy" "dnspolicy" {
	name = "policy_A"
	rule = "CLIENT.IP.SRC.IN_SUBNET(1.1.1.1/24)"
	drop = "YES"
	}
  resource "citrixadc_dnspolicylabel" "dnspolicylabel" {
	labelname = "blue_label"
	transform = "dns_req"
  
	}
  resource "citrixadc_dnspolicylabel_dnspolicy_binding" "dnspolicylabel_dnspolicy_binding" {
	labelname  = citrixadc_dnspolicylabel.dnspolicylabel.labelname
	policyname = citrixadc_dnspolicy.dnspolicy.name
	priority   = 10
  
	}
  

`

const testAccDnspolicylabel_dnspolicy_binding_basic_step2 = `
resource "citrixadc_dnspolicy" "dnspolicy" {
	name = "policy_A"
	rule = "CLIENT.IP.SRC.IN_SUBNET(1.1.1.1/24)"
	drop = "YES"
	}
  resource "citrixadc_dnspolicylabel" "dnspolicylabel" {
	labelname = "blue_label"
	transform = "dns_req"
  
	}
`

func TestAccDnspolicylabel_dnspolicy_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckDnspolicylabel_dnspolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDnspolicylabel_dnspolicy_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDnspolicylabel_dnspolicy_bindingExist("citrixadc_dnspolicylabel_dnspolicy_binding.dnspolicylabel_dnspolicy_binding", nil),
				),
			},
			{
				Config: testAccDnspolicylabel_dnspolicy_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDnspolicylabel_dnspolicy_bindingNotExist("citrixadc_dnspolicylabel_dnspolicy_binding.dnspolicylabel_dnspolicy_binding", "blue_label,policy_A"),
				),
			},
		},
	})
}

func testAccCheckDnspolicylabel_dnspolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No dnspolicylabel_dnspolicy_binding id is set")
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

		bindingId := rs.Primary.ID

		idSlice := strings.SplitN(bindingId, ",", 2)

		labelname := idSlice[0]
		policyname := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "dnspolicylabel_dnspolicy_binding",
			ResourceName:             labelname,
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
			return fmt.Errorf("dnspolicylabel_dnspolicy_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckDnspolicylabel_dnspolicy_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}
		idSlice := strings.SplitN(id, ",", 2)

		labelname := idSlice[0]
		policyname := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "dnspolicylabel_dnspolicy_binding",
			ResourceName:             labelname,
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
			return fmt.Errorf("dnspolicylabel_dnspolicy_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckDnspolicylabel_dnspolicy_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_dnspolicylabel_dnspolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Dnspolicylabel_dnspolicy_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("dnspolicylabel_dnspolicy_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
