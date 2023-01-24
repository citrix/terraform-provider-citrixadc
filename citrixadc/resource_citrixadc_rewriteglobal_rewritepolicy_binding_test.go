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

const testAccRewriteglobal_rewritepolicy_binding_basic = `

	resource "citrixadc_rewriteglobal_rewritepolicy_binding" "tf_rewriteglobal_rewritepolicy_binding" {
        policyname = citrixadc_rewritepolicy.tf_rewrite_policy.name
        priority = 5
        type = "REQ_DEFAULT"
		globalbindtype = "SYSTEM_GLOBAL"
        gotopriorityexpression = "END"
        invoke = "true"
        labelname = citrixadc_rewritepolicylabel.tf_rewritepolicylabel.labelname
        labeltype = "policylabel"
	}

	resource "citrixadc_rewritepolicy" "tf_rewrite_policy" {
		name = "tf_rewrite_policy"
		action = "DROP"
		rule = "HTTP.REQ.URL.PATH_AND_QUERY.CONTAINS(\"helloandby\")"
	}

	resource "citrixadc_rewritepolicylabel" "tf_rewritepolicylabel" {
		labelname = "tf_rewritepolicylabel"
		transform = "http_req"
	}
`

const testAccRewriteglobal_rewritepolicy_binding_basic_step2 = `

	resource "citrixadc_rewritepolicy" "tf_rewrite_policy" {
		name = "tf_rewrite_policy"
		action = "DROP"
		rule = "HTTP.REQ.URL.PATH_AND_QUERY.CONTAINS(\"helloandby\")"
	}	

	resource "citrixadc_rewritepolicylabel" "tf_rewritepolicylabel" {
		labelname = "tf_rewritepolicylabel"
		transform = "http_req"
	}
`

func TestAccRewriteglobal_rewritepolicy_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckRewriteglobal_rewritepolicy_bindingDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccRewriteglobal_rewritepolicy_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRewriteglobal_rewritepolicy_bindingExist("citrixadc_rewriteglobal_rewritepolicy_binding.tf_rewriteglobal_rewritepolicy_binding", nil),
				),
			},
			resource.TestStep{
				Config: testAccRewriteglobal_rewritepolicy_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRewriteglobal_rewritepolicy_bindingNotExist("citrixadc_rewriteglobal_rewritepolicy_binding.tf_rewriteglobal_rewritepolicy_binding", "tf_rewrite_policy,5,REQ_DEFAULT"),
				),
			},
		},
	})
}

func testAccCheckRewriteglobal_rewritepolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No rewriteglobal_rewritepolicy_binding id is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		client := testAccProvider.Meta().(*NetScalerNitroClient).client

		bindingId := rs.Primary.ID

		idSlice := strings.SplitN(bindingId, ",", 3)

		policyname := idSlice[0]
		priority := idSlice[1]
		type_bindpoint := idSlice[2]

		argsMap := make(map[string]string)
		argsMap["type"] = type_bindpoint

		findParams := service.FindParams{
			ResourceType: "rewriteglobal_rewritepolicy_binding",
			ArgsMap:      argsMap,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the matching components
		foundIndex := -1
		for i, rewriteglobal_rewritepolicy_binding := range dataArr {
			if rewriteglobal_rewritepolicy_binding["policyname"] != policyname {
				continue
			} else if rewriteglobal_rewritepolicy_binding["priority"] != priority {
				continue
			} else if rewriteglobal_rewritepolicy_binding["type"] != type_bindpoint {
				continue
			}
			foundIndex = i
			break
		}

		if foundIndex == -1 {
			return fmt.Errorf("rewriteglobal_rewritepolicy_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckRewriteglobal_rewritepolicy_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*NetScalerNitroClient).client

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}
		idSlice := strings.SplitN(id, ",", 3)

		policyname := idSlice[0]
		priority := idSlice[1]
		type_bindpoint := idSlice[2]

		argsMap := make(map[string]string)
		argsMap["type"] = type_bindpoint

		findParams := service.FindParams{
			ResourceType: "rewriteglobal_rewritepolicy_binding",
			ArgsMap:      argsMap,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to hopefully not find the one with the matching components
		foundIndex := -1
		for i, rewriteglobal_rewritepolicy_binding := range dataArr {
			if rewriteglobal_rewritepolicy_binding["policyname"] != policyname {
				continue
			} else if rewriteglobal_rewritepolicy_binding["priority"] != priority {
				continue
			} else if rewriteglobal_rewritepolicy_binding["type"] != type_bindpoint {
				continue
			}
			foundIndex = i
			break
		}

		if foundIndex != -1 {
			return fmt.Errorf("rewriteglobal_rewritepolicy_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckRewriteglobal_rewritepolicy_bindingDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_rewriteglobal_rewritepolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Rewriteglobal_rewritepolicy_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("rewriteglobal_rewritepolicy_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
