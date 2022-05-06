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
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"strings"
	"testing"
)

const testAccService_dospolicy_binding_basic = `
	# Since the dospolicy resource is not yet available on Terraform,
	# the tf_dospolicy policy must be created by hand in order for the script to run correctly.
	# You can do that by using the following Citrix ADC cli commands:
	# add dospolicy tf_dospolicy -qDepth 25

	resource "citrixadc_service" "tf_service" {
		servicetype         = "HTTP"
		name                = "tf_service"
		ipaddress           = "10.77.33.22"
		ip                  = "10.77.33.22"
		port                = "80"
		state               = "ENABLED"
		wait_until_disabled = true
	}
	resource "citrixadc_service_dospolicy_binding" "tf_binding" {
		name       = citrixadc_service.tf_service.name
		policyname = "tf_dospolicy"
	}
`

const testAccService_dospolicy_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion
	# Since the dospolicy resource is not yet available on Terraform,
	# the tf_dospolicy policy must be created by hand in order for the script to run correctly.
	# You can do that by using the following Citrix ADC cli commands:
	# add dospolicy tf_dospolicy -qDepth 25

	resource "citrixadc_service" "tf_service" {
		servicetype         = "HTTP"
		name                = "tf_service"
		ipaddress           = "10.77.33.22"
		ip                  = "10.77.33.22"
		port                = "80"
		state               = "ENABLED"
		wait_until_disabled = true
	}
`

func TestAccService_dospolicy_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckService_dospolicy_bindingDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccService_dospolicy_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckService_dospolicy_bindingExist("citrixadc_service_dospolicy_binding.tf_binding", nil),
				),
			},
			resource.TestStep{
				Config: testAccService_dospolicy_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckService_dospolicy_bindingNotExist("citrixadc_service_dospolicy_binding.tf_binding", "tf_service,tf_dospolicy"),
				),
			},
		},
	})
}

func testAccCheckService_dospolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No service_dospolicy_binding id is set")
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
			ResourceType:             "service_dospolicy_binding",
			ResourceName:             name,
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
			return fmt.Errorf("service_dospolicy_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckService_dospolicy_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*NetScalerNitroClient).client

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}
		idSlice := strings.SplitN(id, ",", 2)

		name := idSlice[0]
		policyname := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "service_dospolicy_binding",
			ResourceName:             name,
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
			return fmt.Errorf("service_dospolicy_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckService_dospolicy_bindingDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_service_dospolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Service_dospolicy_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("service_dospolicy_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
