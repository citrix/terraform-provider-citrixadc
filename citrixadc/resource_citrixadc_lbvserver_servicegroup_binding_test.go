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

const testAccLbvserver_servicegroup_binding_basic = `
resource "citrixadc_lbvserver" "tf_lbvserver" {
  ipv46       = "10.10.10.33"
  name        = "tf_lbvserver"
  port        = 80
  servicetype = "HTTP"
}

resource "citrixadc_lbvserver" "tf_lbvserver2" {
  ipv46       = "10.10.10.34"
  name        = "tf_lbvserver2"
  port        = 80
  servicetype = "HTTP"
}

resource "citrixadc_servicegroup" "tf_servicegroup" {
    servicegroupname = "tf_servicegroup"
    servicetype  = "HTTP"
}

resource "citrixadc_servicegroup" "tf_servicegroup2" {
    servicegroupname = "tf_servicegroup2"
    servicetype  = "HTTP"
}

resource "citrixadc_lbvserver_servicegroup_binding" "tf_binding" {
  name = citrixadc_lbvserver.tf_lbvserver.name
  servicegroupname = citrixadc_servicegroup.tf_servicegroup.servicegroupname
}

resource "citrixadc_lbvserver_servicegroup_binding" "tf_binding2" {
  name = citrixadc_lbvserver.tf_lbvserver.name
  servicegroupname = citrixadc_servicegroup.tf_servicegroup2.servicegroupname
  order = 4
}

resource "citrixadc_lbvserver_servicegroup_binding" "tf_binding3" {
  name = citrixadc_lbvserver.tf_lbvserver2.name
  servicegroupname = citrixadc_servicegroup.tf_servicegroup.servicegroupname
}

resource "citrixadc_lbvserver_servicegroup_binding" "tf_binding4" {
  name = citrixadc_lbvserver.tf_lbvserver2.name
  servicegroupname = citrixadc_servicegroup.tf_servicegroup2.servicegroupname
  order = 4
}

`

func TestAccLbvserver_servicegroup_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckLbvserver_servicegroup_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLbvserver_servicegroup_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbvserver_servicegroup_bindingExist("citrixadc_lbvserver_servicegroup_binding.tf_binding", nil, map[string]interface{}{}),
					testAccCheckLbvserver_servicegroup_bindingExist("citrixadc_lbvserver_servicegroup_binding.tf_binding2", nil, map[string]interface{}{"order": 4}),
					testAccCheckLbvserver_servicegroup_bindingExist("citrixadc_lbvserver_servicegroup_binding.tf_binding3", nil, map[string]interface{}{}),
					testAccCheckLbvserver_servicegroup_bindingExist("citrixadc_lbvserver_servicegroup_binding.tf_binding4", nil, map[string]interface{}{"order": 4}),
				),
			},
		},
	})
}

func testAccCheckLbvserver_servicegroup_bindingExist(n string, id *string, expectedValues map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lbvserver_servicegroup_binding name is set")
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
		servicegroupname := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "lbvserver_servicegroup_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		if err != nil {
			return err
		}

		// Iterate through results to find the one with the right monitor name
		foundIndex := -1
		data := map[string]interface{}{}
		for i, v := range dataArr {
			if v["servicegroupname"].(string) == servicegroupname {
				foundIndex = i
				data = v
				break
			}
		}

		// Resource is missing
		if foundIndex == -1 {
			return fmt.Errorf("lbvserver_servicegroup_binding %s not found", n)
		}

		// Iterate through all expected values and validate them
		for key, expectedValue := range expectedValues {
			if actualValue, exists := data[key]; !exists {
				return fmt.Errorf("Expected key %q not found in retrieved data", key)
			} else if !compareValues(expectedValue, actualValue) {
				return fmt.Errorf("Expected value for %q differs. Expected: %v, Retrieved: %v",
					key, expectedValue, actualValue)
			}
		}

		return nil
	}
}

func testAccCheckLbvserver_servicegroup_bindingDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_lbvserver_servicegroup_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Lbvserver_servicegroup_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("lbvserver_servicegroup_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
