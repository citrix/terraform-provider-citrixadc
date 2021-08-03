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

const testAccLbgroup_lbvserver_binding_basic = `
	resource "citrixadc_lbgroup_lbvserver_binding" "tf_lbvserverbinding" {
		name = citrixadc_lbgroup.tf_lbgroup.name
		vservername = citrixadc_lbvserver.tf_lbvserver.name
	}

	resource "citrixadc_lbvserver" "tf_lbvserver" {
		name 		= "tf_lbvserver"
		ipv46       = "1.1.1.8"
		port        = "80"
		servicetype = "HTTP"
	}
	
	resource "citrixadc_lbgroup" "tf_lbgroup" {
		name = "tf_lbgroup"
	}
`

const testAccLbgroup_lbvserver_binding_basic_step2 = `
	resource "citrixadc_lbvserver" "tf_lbvserver" {
		name = "tf_lbvserver"
		ipv46       = "1.1.1.8"
		port        = "80"
		servicetype = "HTTP"
	}
	
	resource "citrixadc_lbgroup" "tf_lbgroup" {
		name = "tf_lbgroup"
	}
`

func TestAccLbgroup_lbvserver_binding_basic(t *testing.T) {
	if adcTestbed != "STANDALONE" {
		t.Skipf("ADC testbed is %s. Expected STANDALONE.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckLbgroup_lbvserver_bindingDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccLbgroup_lbvserver_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbgroup_lbvserver_bindingExist("citrixadc_lbgroup_lbvserver_binding.tf_lbvserverbinding", nil),
				),
			},
			resource.TestStep{
				Config: testAccLbgroup_lbvserver_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbgroup_lbvserver_bindingNotExist("citrixadc_lbgroup_lbvserver_binding.tf_lbvserverbinding", "tf_lbgroup,tf_lbvserver"),
				),
			},
		},
	})
}

func testAccCheckLbgroup_lbvserver_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lbgroup_lbvserver_binding name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		bindingId := rs.Primary.ID
		idSlice := strings.SplitN(bindingId, ",", 2)
		lbgroupName := idSlice[0]
		lbvserverName := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             service.Lbgroup_lbvserver_binding.Type(),
			ResourceName:             lbgroupName,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := nsClient.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the right policy name
		foundIndex := -1
		for i, v := range dataArr {
			if v["vservername"].(string) == lbvserverName {
				foundIndex = i
				break
			}
		}

		// Resource is missing
		if foundIndex == -1 {
			return fmt.Errorf("Cannot find lbgroup_lbvserver_binding ID %v", bindingId)
		}

		return nil
	}
}

func testAccCheckLbgroup_lbvserver_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*NetScalerNitroClient).client

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}
		idSlice := strings.SplitN(id, ",", 2)

		lbgroupName := idSlice[0]
		lbvserverName := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "lbgroup_lbvserver_binding",
			ResourceName:             lbgroupName,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the right policy name
		found := false
		for _, v := range dataArr {
			if v["vservername"].(string) == lbvserverName {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("Lbgroup_lbvserver_binding %s was found, but is should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckLbgroup_lbvserver_bindingDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_lbgroup_lbvserver_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Lbgroup_lbvserver_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("lbgroup_lbvserver_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
