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

const testAccSystemgroup_systemcmdpolicy_binding_basic = `

	resource "citrixadc_systemgroup" "tf_systemgroup" {
		groupname    = "tf_systemgroup"
		timeout      = 999
		promptstring = "bye>"
	}

	resource "citrixadc_systemcmdpolicy" "tf_policy" {
		policyname = "tf_policy"
		action     = "DENY"
		cmdspec    = "add.*"
	}

	resource "citrixadc_systemgroup_systemcmdpolicy_binding" "tf_bind" {
		groupname  = citrixadc_systemgroup.tf_systemgroup.groupname
		policyname = citrixadc_systemcmdpolicy.tf_policy.policyname
		priority   = 100
	}
`

const testAccSystemgroup_systemcmdpolicy_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion
	resource "citrixadc_systemgroup" "tf_systemgroup" {
		groupname    = "tf_systemgroup"
		timeout      = 999
		promptstring = "bye>"
	}

	resource "citrixadc_systemcmdpolicy" "tf_policy" {
		policyname = "tf_policy"
		action     = "DENY"
		cmdspec    = "add.*"
	}
`

func TestAccSystemgroup_systemcmdpolicy_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSystemgroup_systemcmdpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSystemgroup_systemcmdpolicy_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSystemgroup_systemcmdpolicy_bindingExist("citrixadc_systemgroup_systemcmdpolicy_binding.tf_bind", nil),
				),
			},
			{
				Config: testAccSystemgroup_systemcmdpolicy_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSystemgroup_systemcmdpolicy_bindingNotExist("citrixadc_systemgroup_systemcmdpolicy_binding.tf_bind", "tf_systemgroup,tf_policy"),
				),
			},
		},
	})
}

func testAccCheckSystemgroup_systemcmdpolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No systemgroup_systemcmdpolicy_binding id is set")
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
		policyname := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "systemgroup_systemcmdpolicy_binding",
			ResourceName:             groupname,
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
			return fmt.Errorf("systemgroup_systemcmdpolicy_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckSystemgroup_systemcmdpolicy_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*NetScalerNitroClient).client

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}
		idSlice := strings.SplitN(id, ",", 2)

		groupname := idSlice[0]
		policyname := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "systemgroup_systemcmdpolicy_binding",
			ResourceName:             groupname,
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
			return fmt.Errorf("systemgroup_systemcmdpolicy_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckSystemgroup_systemcmdpolicy_bindingDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_systemgroup_systemcmdpolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Systemgroup_systemcmdpolicy_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("systemgroup_systemcmdpolicy_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
