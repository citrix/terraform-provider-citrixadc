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

const testAccAuditnslogglobal_auditnslogpolicy_binding_basic = `

	resource "citrixadc_auditnslogaction" "tf_auditnslogaction" {
		name     = "my_auditnslogaction"
		serverip = "1.1.1.1"
		loglevel = ["ALERT", "CRITICAL"]
	}
	resource "citrixadc_auditnslogpolicy" "tf_auditnslogpolicy" {
		name   = "my_auditnslogpolicy"
		rule   = "true"
		action = citrixadc_auditnslogaction.tf_auditnslogaction.name
	}

	resource "citrixadc_auditnslogglobal_auditnslogpolicy_binding" "tf_auditnslogglobal_auditnslogpolicy_binding" {
		policyname 		= citrixadc_auditnslogpolicy.tf_auditnslogpolicy.name
		priority   		= 100
		globalbindtype 	= "SYSTEM_GLOBAL"
	}
`

const testAccAuditnslogglobal_auditnslogpolicy_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion

	resource "citrixadc_auditnslogaction" "tf_auditnslogaction" {
		name     = "my_auditnslogaction"
		serverip = "1.1.1.1"
		loglevel = ["ALERT", "CRITICAL"]
	}
	resource "citrixadc_auditnslogpolicy" "tf_auditnslogpolicy" {
		name   = "my_auditnslogpolicy"
		rule   = "true"
		action = citrixadc_auditnslogaction.tf_auditnslogaction.name
	}
`

func TestAccAuditnslogglobal_auditnslogpolicy_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAuditnslogglobal_auditnslogpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuditnslogglobal_auditnslogpolicy_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuditnslogglobal_auditnslogpolicy_bindingExist("citrixadc_auditnslogglobal_auditnslogpolicy_binding.tf_auditnslogglobal_auditnslogpolicy_binding", nil),
				),
			},
			{
				Config: testAccAuditnslogglobal_auditnslogpolicy_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuditnslogglobal_auditnslogpolicy_bindingNotExist("citrixadc_auditnslogglobal_auditnslogpolicy_binding.tf_auditnslogglobal_auditnslogpolicy_binding", "my_auditnslogpolicy"),
				),
			},
		},
	})
}

func testAccCheckAuditnslogglobal_auditnslogpolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No auditnslogglobal_auditnslogpolicy_binding id is set")
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
			ResourceType:             "auditnslogglobal_auditnslogpolicy_binding",
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
			return fmt.Errorf("auditnslogglobal_auditnslogpolicy_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckAuditnslogglobal_auditnslogpolicy_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*NetScalerNitroClient).client

		policyname := id

		findParams := service.FindParams{
			ResourceType:             "auditnslogglobal_auditnslogpolicy_binding",
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
			return fmt.Errorf("auditnslogglobal_auditnslogpolicy_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckAuditnslogglobal_auditnslogpolicy_bindingDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_auditnslogglobal_auditnslogpolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource("auditnslogglobal_auditnslogpolicy_binding", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("auditnslogglobal_auditnslogpolicy_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
