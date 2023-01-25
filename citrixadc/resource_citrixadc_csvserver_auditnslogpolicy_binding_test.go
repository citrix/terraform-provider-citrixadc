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

const testAccCsvserver_auditnslogpolicy_binding_basic = `
	# Since the auditnslogpolicy resource is not yet available on Terraform,
	# the tf_auditnslogpolicy policy must be created by hand in order for the script to run correctly.
	# You can do that by using the following Citrix ADC cli commands:
	# add audit nslogAction tf_auditnslogaction 1.1.1.1 -loglevel NONE
	# add audit nslogPolicy tf_auditnslogpolicy ns_true tf_auditnslogaction

	resource "citrixadc_csvserver_auditnslogpolicy_binding" "tf_csvserver_auditnslogpolicy_binding" {
        name = citrixadc_csvserver.tf_csvserver.name
        policyname = "tf_auditnslogpolicy"
        priority = 5
	}

	resource "citrixadc_csvserver" "tf_csvserver" {
		name = "tf_csvserver"
		ipv46 = "10.202.11.11"
		port = 8080
		servicetype = "HTTP"
	}
`

const testAccCsvserver_auditnslogpolicy_binding_basic_step2 = `
	resource "citrixadc_csvserver" "tf_csvserver" {
		name = "tf_csvserver"
		ipv46 = "10.202.11.11"
		port = 8080
		servicetype = "HTTP"
	}
`

func TestAccCsvserver_auditnslogpolicy_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCsvserver_auditnslogpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCsvserver_auditnslogpolicy_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCsvserver_auditnslogpolicy_bindingExist("citrixadc_csvserver_auditnslogpolicy_binding.tf_csvserver_auditnslogpolicy_binding", nil),
				),
			},
			resource.TestStep{
				Config: testAccCsvserver_auditnslogpolicy_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCsvserver_auditnslogpolicy_bindingNotExist("citrixadc_csvserver_auditnslogpolicy_binding.tf_csvserver_auditnslogpolicy_binding", "tf_csvserver,tf_auditnslogpolicy"),
				),
			},
		},
	})
}

func testAccCheckCsvserver_auditnslogpolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No csvserver_auditnslogpolicy_binding id is set")
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
			ResourceType:             "csvserver_auditnslogpolicy_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the right monitor name
		found := false
		for _, v := range dataArr {
			if v["policyname"].(string) == policyname {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("csvserver_auditnslogpolicy_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckCsvserver_auditnslogpolicy_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*NetScalerNitroClient).client

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}
		idSlice := strings.SplitN(id, ",", 2)

		csvserverName := idSlice[0]
		policyName := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "csvserver_auditnslogpolicy_binding",
			ResourceName:             csvserverName,
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
			if v["policyname"].(string) == policyName {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("csvserver_auditnslogpolicy_binding %s was found, but is should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckCsvserver_auditnslogpolicy_bindingDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_csvserver_auditnslogpolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Csvserver_auditnslogpolicy_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("csvserver_auditnslogpolicy_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
