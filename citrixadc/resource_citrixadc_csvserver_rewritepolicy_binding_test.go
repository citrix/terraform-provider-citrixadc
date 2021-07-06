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
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"fmt"
	"strings"
	"testing"
)

const testAccCsvserver_rewritepolicy_binding_basic_step1 = `
resource "citrixadc_csvserver" "tf_csvserver" {
  ipv46       = "10.10.10.33"
  name        = "tf_csvserver"
  port        = 80
  servicetype = "HTTP"
}

resource "citrixadc_csvserver" "tf_csvserver2" {
  ipv46       = "10.10.10.34"
  name        = "tf_csvserver2"
  port        = 80
  servicetype = "HTTP"
}

resource "citrixadc_rewritepolicy" "tf_rewrite_policy" {
  name   = "tf_test_rewrite_policy"
  action = "DROP"
  rule   = "HTTP.REQ.URL.PATH_AND_QUERY.CONTAINS(\"helloandby\")"
}

resource "citrixadc_csvserver_rewritepolicy_binding" "tf_bind" {
    name = citrixadc_csvserver.tf_csvserver.name
    policyname = citrixadc_rewritepolicy.tf_rewrite_policy.name
    priority = 100
    bindpoint = "REQUEST"
}

resource "citrixadc_csvserver_rewritepolicy_binding" "tf_bind2" {
    name = citrixadc_csvserver.tf_csvserver2.name
    policyname = citrixadc_rewritepolicy.tf_rewrite_policy.name
    priority = 100
    bindpoint = "REQUEST"
}
`

const testAccCsvserver_rewritepolicy_binding_basic_step2 = `
resource "citrixadc_csvserver" "tf_csvserver" {
  ipv46       = "10.10.10.33"
  name        = "tf_csvserver"
  port        = 80
  servicetype = "HTTP"
}

resource "citrixadc_csvserver" "tf_csvserver2" {
  ipv46       = "10.10.10.34"
  name        = "tf_csvserver2"
  port        = 80
  servicetype = "HTTP"
}

resource "citrixadc_rewritepolicy" "tf_rewrite_policy" {
  name   = "tf_test_rewrite_policy"
  action = "DROP"
  rule   = "HTTP.REQ.URL.PATH_AND_QUERY.CONTAINS(\"helloandby\")"
}

resource "citrixadc_csvserver_rewritepolicy_binding" "tf_bind" {
    name = citrixadc_csvserver.tf_csvserver.name
    policyname = citrixadc_rewritepolicy.tf_rewrite_policy.name
    priority = 110
    bindpoint = "REQUEST"
}

resource "citrixadc_csvserver_rewritepolicy_binding" "tf_bind2" {
    name = citrixadc_csvserver.tf_csvserver2.name
    policyname = citrixadc_rewritepolicy.tf_rewrite_policy.name
    priority = 120
    bindpoint = "REQUEST"
}
`

func TestAccCsvserver_rewritepolicy_binding_basic(t *testing.T) {
	if adcTestbed != "STANDALONE" {
		t.Skipf("ADC testbed is %s. Expected STANDALONE.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCsvserver_rewritepolicy_bindingDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCsvserver_rewritepolicy_binding_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCsvserver_rewritepolicy_bindingExist("citrixadc_csvserver_rewritepolicy_binding.tf_bind", nil),
					testAccCheckCsvserver_rewritepolicy_bindingExist("citrixadc_csvserver_rewritepolicy_binding.tf_bind2", nil),
				),
			},
			resource.TestStep{
				Config: testAccCsvserver_rewritepolicy_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCsvserver_rewritepolicy_bindingExist("citrixadc_csvserver_rewritepolicy_binding.tf_bind", nil),
					testAccCheckCsvserver_rewritepolicy_bindingExist("citrixadc_csvserver_rewritepolicy_binding.tf_bind2", nil),
				),
			},
		},
	})
}

func testAccCheckCsvserver_rewritepolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No csvserver_rewritepolicy_binding name is set")
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
		policyname := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "csvserver_rewritepolicy_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the right policy name
		foundIndex := -1
		for i, v := range dataArr {
			if v["policyname"].(string) == policyname {
				foundIndex = i
				break
			}
		}

		// Resource is missing
		if foundIndex == -1 {
			return fmt.Errorf("Cannot find csvserver_rewritepolicy_binding %s", bindingId)
		}

		return nil
	}
}

func testAccCheckCsvserver_rewritepolicy_bindingDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_csvserver_rewritepolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Csvserver_rewritepolicy_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("csvserver_rewritepolicy_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
