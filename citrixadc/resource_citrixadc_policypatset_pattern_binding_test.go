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
	"strings"
	"testing"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

const testAccPolicypatset_pattern_binding_basic_step1 = `

resource "citrixadc_policypatset" "tf_patset" {
    name = "tf_patset"
    comment = "some comment"
}

resource "citrixadc_policypatset_pattern_binding" "tf_bind" {
    name = citrixadc_policypatset.tf_patset.name
    string = "pattern1,/postfix"
}

`

const testAccPolicypatset_pattern_binding_basic_step2 = `

resource "citrixadc_policypatset" "tf_patset" {
    name = "tf_patset"
    comment = "some comment"
}

resource "citrixadc_policypatset_pattern_binding" "tf_bind" {
    name = citrixadc_policypatset.tf_patset.name
    string = "pattern2,/postfix"
}

`

func TestAccPolicypatset_pattern_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPolicypatset_pattern_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPolicypatset_pattern_binding_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPolicypatset_pattern_bindingExist("citrixadc_policypatset_pattern_binding.tf_bind", nil),
				),
			},
			{
				Config: testAccPolicypatset_pattern_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPolicypatset_pattern_bindingExist("citrixadc_policypatset_pattern_binding.tf_bind", nil),
				),
			},
		},
	})
}

func testAccCheckPolicypatset_pattern_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No policypatset_pattern_binding name is set")
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

		name := idSlice[0]
		stringText := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "policypatset_pattern_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 2823,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the right String
		foundIndex := -1
		for i, v := range dataArr {
			if v["String"].(string) == stringText {
				foundIndex = i
				break
			}
		}

		// Resource is missing
		if foundIndex == -1 {
			return fmt.Errorf("FindResourceArrayWithParams  could not find pattern_binding %v", bindingId)
		}

		return nil
	}
}

func testAccCheckPolicypatset_pattern_bindingDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_policypatset_pattern_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Policypatset_pattern_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("policypatset_pattern_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
