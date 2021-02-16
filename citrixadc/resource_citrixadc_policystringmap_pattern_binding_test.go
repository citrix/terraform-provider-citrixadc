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
	"github.com/chiradeep/go-nitro/netscaler"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"strings"
	"testing"
)

const testAccPolicystringmap_pattern_binding_basic_step1 = `

resource "citrixadc_policystringmap" "tf_policystringmap" {
    name = "tf_policystringmap"
    comment = "Some comment"
}

resource "citrixadc_policystringmap_pattern_binding" "tf_bind1" {
    name = citrixadc_policystringmap.tf_policystringmap.name
    key = "key1"
    value = "value1"
}

resource "citrixadc_policystringmap_pattern_binding" "tf_bind2" {
    name = citrixadc_policystringmap.tf_policystringmap.name
    key = "key2"
    value = "value2"
}
`

const testAccPolicystringmap_pattern_binding_basic_step2 = `

resource "citrixadc_policystringmap" "tf_policystringmap" {
    name = "tf_policystringmap"
    comment = "Some comment"
}

resource "citrixadc_policystringmap_pattern_binding" "tf_bind1" {
    name = citrixadc_policystringmap.tf_policystringmap.name
    key = "key1"
    value = "value1_new"
}

resource "citrixadc_policystringmap_pattern_binding" "tf_bind2" {
    name = citrixadc_policystringmap.tf_policystringmap.name
    key = "key2"
    value = "value2"
}
`

func TestAccPolicystringmap_pattern_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPolicystringmap_pattern_bindingDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccPolicystringmap_pattern_binding_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPolicystringmap_pattern_bindingExist("citrixadc_policystringmap_pattern_binding.tf_bind1", nil),
					testAccCheckPolicystringmap_pattern_bindingExist("citrixadc_policystringmap_pattern_binding.tf_bind2", nil),
				),
			},
			resource.TestStep{
				Config: testAccPolicystringmap_pattern_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPolicystringmap_pattern_bindingExist("citrixadc_policystringmap_pattern_binding.tf_bind1", nil),
					testAccCheckPolicystringmap_pattern_bindingExist("citrixadc_policystringmap_pattern_binding.tf_bind2", nil),
				),
			},
		},
	})
}

func testAccCheckPolicystringmap_pattern_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No policystringmap_pattern_binding name is set")
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
		key := idSlice[1]

		findParams := netscaler.FindParams{
			ResourceType:             "policystringmap_pattern_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		foundIndex := -1
		for i, v := range dataArr {
			if v["key"].(string) == key {
				foundIndex = i
				break
			}
		}
		if foundIndex == -1 {
			return fmt.Errorf("Could not find binding %s", bindingId)
		}

		return nil
	}
}

func testAccCheckPolicystringmap_pattern_bindingDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_policystringmap_pattern_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(netscaler.Policystringmap_pattern_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("policystringmap_pattern_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
