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
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccPolicydataset_value_binding_basic(t *testing.T) {
	if adcTestbed != "STANDALONE" {
		t.Skipf("ADC testbed is %s. Expected STANDALONE.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPolicydataset_value_bindingDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccPolicydataset_value_binding_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPolicydatasetExist("citrixadc_policydataset.tf_dataset", nil),
					testAccCheckPolicydatasetValue("citrixadc_policydataset_value_binding.tf_value1"),
					testAccCheckPolicydatasetValue("citrixadc_policydataset_value_binding.tf_value2"),
				),
			},
			resource.TestStep{
				Config: testAccPolicydataset_value_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPolicydatasetExist("citrixadc_policydataset.tf_dataset", nil),
					testAccCheckPolicydatasetValue("citrixadc_policydataset_value_binding.tf_value1"),
					testAccCheckPolicydatasetValue("citrixadc_policydataset_value_binding.tf_value3"),
				),
			},
			resource.TestStep{
				Config: testAccPolicydataset_value_binding_basic_step3,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPolicydatasetExist("citrixadc_policydataset.tf_dataset", nil),
					testAccCheckPolicydatasetValue("citrixadc_policydataset_value_binding.tf_value3"),
				),
			},
		},
	})
}

func testAccCheckPolicydatasetValue(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No binding id")
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		idSlice := strings.Split(rs.Primary.ID, ",")

		name := idSlice[0]
		value := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "policydataset_value_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 2823,
		}
		dataArr, err := nsClient.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return fmt.Errorf("Error during FindResourceArrayWithParams %s", err.Error())
		}

		// Resource is missing
		if len(dataArr) == 0 {
			return fmt.Errorf("FindResourceArrayWithParams returned empty array")
		}

		// Iterate through results to find the one with the right monitor name
		foundIndex := -1
		for i, v := range dataArr {
			if v["value"].(string) == value {
				foundIndex = i
				break
			}
		}

		// Resource is missing
		if foundIndex == -1 {
			return fmt.Errorf("FindResourceArrayWithParams monitor name not found in array")
		}

		return nil
	}
}

func testAccCheckPolicydataset_value_bindingDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_policydataset_value_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Policydataset_value_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("LB vserver %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccPolicydataset_value_binding_basic_step1 = `
resource "citrixadc_policydataset" "tf_dataset" {
  name    = "tf_dataset"
  type    = "number"
  comment = "hello"
}

resource "citrixadc_policydataset_value_binding" "tf_value1" {
  name = citrixadc_policydataset.tf_dataset.name

  value    = 100
  index    = 111
}

resource "citrixadc_policydataset_value_binding" "tf_value2" {
  name = citrixadc_policydataset.tf_dataset.name

  value    = 200
}
`

const testAccPolicydataset_value_binding_basic_step2 = `
resource "citrixadc_policydataset" "tf_dataset" {
  name    = "tf_dataset"
  type    = "number"
  comment = "hello"
}

resource "citrixadc_policydataset_value_binding" "tf_value1" {
  name = citrixadc_policydataset.tf_dataset.name

  value    = 100
  index    = 111
}

resource "citrixadc_policydataset_value_binding" "tf_value3" {
  name = citrixadc_policydataset.tf_dataset.name

  value    = 300
}
`

const testAccPolicydataset_value_binding_basic_step3 = `
resource "citrixadc_policydataset" "tf_dataset" {
  name    = "tf_dataset"
  type    = "number"
  comment = "hello"
}

resource "citrixadc_policydataset_value_binding" "tf_value3" {
  name = citrixadc_policydataset.tf_dataset.name

  value    = 300
  index  = 333
}
`
