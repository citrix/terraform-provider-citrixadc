/*
Copyright 2024 Citrix Systems, Inc

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

const testAccLbvserver_lbpolicy_binding_basic = `
	resource "citrixadc_lbvserver" "tf_lbvserver" {
		ipv46       = "10.10.10.33"
		name        = "tf_lbvserver"
		port        = 80
		servicetype = "HTTP"
	}
	resource "citrixadc_lbpolicy" "tf_pol" {
		name   = "tf_pol"
		rule   = "true"
		action = "NOLBACTION"
	}
	resource "citrixadc_lbvserver_lbpolicy_binding" "tf_bind" {
		name       = citrixadc_lbvserver.tf_lbvserver.name
		policyname = citrixadc_lbpolicy.tf_pol.name
		priority   = 10
	}
`

const testAccLbvserver_lbpolicy_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion
	
	resource "citrixadc_lbvserver" "tf_lbvserver" {
		ipv46       = "10.10.10.33"
		name        = "tf_lbvserver"
		port        = 80
		servicetype = "HTTP"
	}
	resource "citrixadc_lbpolicy" "tf_pol" {
		name   = "tf_pol"
		rule   = "true"
		action = "NOLBACTION"
	}
`

func TestAccLbvserver_lbpolicy_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLbvserver_lbpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLbvserver_lbpolicy_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbvserver_lbpolicy_bindingExist("citrixadc_lbvserver_lbpolicy_binding.tf_bind", nil),
				),
			},
			{
				Config: testAccLbvserver_lbpolicy_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbvserver_lbpolicy_bindingNotExist("citrixadc_lbvserver_lbpolicy_binding.tf_bind", "tf_lbvserver,tf_pol"),
				),
			},
		},
	})
}

func testAccCheckLbvserver_lbpolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lbvserver_lbpolicy_binding id is set")
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

		bindingId := rs.Primary.ID

		idSlice := strings.SplitN(bindingId, ",", 2)

		name := idSlice[0]
		policyname := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "lbvserver_lbpolicy_binding",
			ResourceName:             name,
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
			return fmt.Errorf("lbvserver_lbpolicy_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckLbvserver_lbpolicy_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}
		idSlice := strings.SplitN(id, ",", 2)

		name := idSlice[0]
		policyname := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "lbvserver_lbpolicy_binding",
			ResourceName:             name,
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
			return fmt.Errorf("lbvserver_lbpolicy_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckLbvserver_lbpolicy_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_lbvserver_lbpolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("lbvserver_lbpolicy_binding", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("lbvserver_lbpolicy_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
