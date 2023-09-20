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
	//"strings"
	"testing"
)

const testAccBotglobal_botpolicy_binding_basic = `
	resource "citrixadc_botpolicy" "tf_botpolicy" {
		name        = "tf_botpolicy"
		profilename = "BOT_BYPASS"
		rule        = "true"
		comment     = "COMMENT FOR BOTPOLICY"
	}
	resource "citrixadc_botglobal_botpolicy_binding" "tf_binding" {
		policyname = citrixadc_botpolicy.tf_botpolicy.name
		priority   = 90
		type       = "REQ_OVERRIDE"
	}
`

const testAccBotglobal_botpolicy_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion
	resource "citrixadc_botpolicy" "tf_botpolicy" {
		name        = "tf_botpolicy"
		profilename = "BOT_BYPASS"
		rule        = "true"
		comment     = "COMMENT FOR BOTPOLICY"
	}
`

func TestAccBotglobal_botpolicy_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckBotglobal_botpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccBotglobal_botpolicy_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBotglobal_botpolicy_bindingExist("citrixadc_botglobal_botpolicy_binding.tf_binding", nil),
				),
			},
			{
				Config: testAccBotglobal_botpolicy_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBotglobal_botpolicy_bindingNotExist("citrixadc_botglobal_botpolicy_binding.tf_binding", "tf_botpolicy", "REQ_OVERRIDE"),
				),
			},
		},
	})
}

func testAccCheckBotglobal_botpolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No botglobal_botpolicy_binding id is set")
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
			ResourceType:             "botglobal_botpolicy_binding",
			ArgsMap:                  map[string]string{"type": typename},
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
			return fmt.Errorf("botglobal_botpolicy_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckBotglobal_botpolicy_bindingNotExist(n string, id string, typename string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*NetScalerNitroClient).client
		policyname := id

		findParams := service.FindParams{
			ResourceType:             "botglobal_botpolicy_binding",
			ArgsMap:                  map[string]string{"type": typename},
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
			return fmt.Errorf("botglobal_botpolicy_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckBotglobal_botpolicy_bindingDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_botglobal_botpolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource("botglobal_botpolicy_binding", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("botglobal_botpolicy_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
