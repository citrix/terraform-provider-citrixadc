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
	"strings"
	"testing"
)

const testAccAaagroup_vpnsessionpolicy_binding_basic = `

	resource "citrixadc_aaagroup_vpnsessionpolicy_binding" "tf_aaagroup_vpnsessionpolicy_binding" {
		groupname = citrixadc_aaagroup.tf_aaagroup.groupname
		policy    = citrixadc_vpnsessionpolicy.tf_vpnsessionpolicy.name
		priority  = 100
	}
		
	resource "citrixadc_aaagroup" "tf_aaagroup" {
		groupname = "my_group"
		weight    = 100
		loggedin  = false
	}
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
`

const testAccAaagroup_vpnsessionpolicy_binding_basic_step2 = `
	resource "citrixadc_aaagroup" "tf_aaagroup" {
		groupname = "my_group"
		weight    = 100
		loggedin  = false
	}
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
`

func TestAccAaagroup_vpnsessionpolicy_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAaagroup_vpnsessionpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAaagroup_vpnsessionpolicy_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaagroup_vpnsessionpolicy_bindingExist("citrixadc_aaagroup_vpnsessionpolicy_binding.tf_aaagroup_vpnsessionpolicy_binding", nil),
				),
			},
			{
				Config: testAccAaagroup_vpnsessionpolicy_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaagroup_vpnsessionpolicy_bindingNotExist("citrixadc_aaagroup_vpnsessionpolicy_binding.tf_aaagroup_vpnsessionpolicy_binding", "my_group,tf_vpnsessionpolicy"),
				),
			},
		},
	})
}

func testAccCheckAaagroup_vpnsessionpolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No aaagroup_vpnsessionpolicy_binding id is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		client := testAccProvider.Meta().(*NetScalerNitroClient).client

		bindingId := rs.Primary.ID

		idSlice := strings.SplitN(bindingId, ",", 2)

		groupname := idSlice[0]
		policy := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "aaagroup_vpnsessionpolicy_binding",
			ResourceName:             groupname,
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
			return fmt.Errorf("aaagroup_vpnsessionpolicy_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckAaagroup_vpnsessionpolicy_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*NetScalerNitroClient).client

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}
		idSlice := strings.SplitN(id, ",", 2)

		groupname := idSlice[0]
		policy := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "aaagroup_vpnsessionpolicy_binding",
			ResourceName:             groupname,
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
			return fmt.Errorf("aaagroup_vpnsessionpolicy_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckAaagroup_vpnsessionpolicy_bindingDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_aaagroup_vpnsessionpolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Aaagroup_vpnsessionpolicy_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("aaagroup_vpnsessionpolicy_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
