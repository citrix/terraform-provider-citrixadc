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

const testAccVpnglobal_vpnsessionpolicy_binding_basic = `

	resource "citrixadc_vpnsessionaction" "tf_vpnsessionaction" {
		name                       = "newsession"
		sesstimeout                = "10"
		defaultauthorizationaction = "ALLOW"
	}
	
	resource "citrixadc_vpnsessionpolicy" "tf_vpnsessionpolicy" {
		name   = "tf_vpnsessionpolicy"
		rule   = "HTTP.REQ.HEADER(\"User-Agent\").CONTAINS(\"CitrixReceiver\").NOT"
		action = citrixadc_vpnsessionaction.tf_vpnsessionaction.name
	}
	resource "citrixadc_vpnglobal_vpnsessionpolicy_binding" "tf_bind" {
		policyname = citrixadc_vpnsessionpolicy.tf_vpnsessionpolicy.name
		priority = 20
	}
`

const testAccVpnglobal_vpnsessionpolicy_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion
	resource "citrixadc_vpnsessionaction" "tf_vpnsessionaction" {
		name                       = "newsession"
		sesstimeout                = "10"
		defaultauthorizationaction = "ALLOW"
	}
	  
	  resource "citrixadc_vpnsessionpolicy" "tf_vpnsessionpolicy" {
		name   = "tf_vpnsessionpolicy"
		rule   = "HTTP.REQ.HEADER(\"User-Agent\").CONTAINS(\"CitrixReceiver\")"
		action = citrixadc_vpnsessionaction.tf_vpnsessionaction.name
	}
`

func TestAccVpnglobal_vpnsessionpolicy_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckVpnglobal_vpnsessionpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccVpnglobal_vpnsessionpolicy_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnglobal_vpnsessionpolicy_bindingExist("citrixadc_vpnglobal_vpnsessionpolicy_binding.tf_bind", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnglobal_vpnsessionpolicy_binding.tf_bind", "policyname", "tf_vpnsessionpolicy"),
				),
			},
			resource.TestStep{
				Config: testAccVpnglobal_vpnsessionpolicy_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnglobal_vpnsessionpolicy_bindingNotExist("citrixadc_vpnglobal_vpnsessionpolicy_binding.tf_bind", "policyname"),
				),
			},
		},
	})
}

func testAccCheckVpnglobal_vpnsessionpolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No vpnglobal_vpnsessionpolicy_binding id is set")
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
			ResourceType:             "vpnglobal_vpnsessionpolicy_binding",
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
			return fmt.Errorf("vpnglobal_vpnsessionpolicy_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckVpnglobal_vpnsessionpolicy_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*NetScalerNitroClient).client
		policyname := id

		findParams := service.FindParams{
			ResourceType:             "vpnglobal_vpnsessionpolicy_binding",
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
			return fmt.Errorf("vpnglobal_vpnsessionpolicy_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckVpnglobal_vpnsessionpolicy_bindingDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_vpnglobal_vpnsessionpolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Vpnglobal_vpnsessionpolicy_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("vpnglobal_vpnsessionpolicy_binding %s still exists", rs.Primary.ID)
		}

	}
	return nil
}
