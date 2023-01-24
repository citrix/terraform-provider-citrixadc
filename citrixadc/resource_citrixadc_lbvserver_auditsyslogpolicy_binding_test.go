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

const testAccLbvserver_auditsyslogpolicy_binding_basic_step1 = `
resource "citrixadc_lbvserver" "tf_lbvserver3" {
name        = "tf_lbvserver3"
servicetype = "HTTP"
}

resource "citrixadc_auditsyslogaction" "tf_syslogaction2" {
	name = "tf_syslogaction2"
	serverip = "10.124.67.93"
	loglevel = [
		"ERROR",
		"NOTICE",
	]
}

resource "citrixadc_auditsyslogpolicy" "tf_syslogpolicy2" {
	name = "tf_syslogpolicy2"
	rule = "true"
	action = citrixadc_auditsyslogaction.tf_syslogaction2.name

}

resource "citrixadc_lbvserver_auditsyslogpolicy_binding" "demo" {
	name = citrixadc_lbvserver.tf_lbvserver3.name
	policyname = citrixadc_auditsyslogpolicy.tf_syslogpolicy2.name
	invoke = "false"
	priority = 56
}
`

const testAccLbvserver_auditsyslogpolicy_binding_basic_step2 = `
resource "citrixadc_lbvserver" "tf_lbvserver3" {
name        = "tf_lbvserver3"
servicetype = "HTTP"
}

resource "citrixadc_auditsyslogaction" "tf_syslogaction2" {
	name = "tf_syslogaction2"
	serverip = "10.124.67.93"
	loglevel = [
		"ERROR",
		"NOTICE",
	]
}

resource "citrixadc_auditsyslogpolicy" "tf_syslogpolicy2" {
	name = "tf_syslogpolicy2"
	rule = "true"
	action = citrixadc_auditsyslogaction.tf_syslogaction2.name

}
`

func TestAccLbvserver_auditsyslogpolicy_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckLbvserver_auditsyslogpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccLbvserver_auditsyslogpolicy_binding_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbvserver_auditsyslogpolicy_bindingExist("citrixadc_lbvserver_auditsyslogpolicy_binding.demo", nil),
				),
			},
			resource.TestStep{
				Config: testAccLbvserver_auditsyslogpolicy_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbvserver_auditsyslogpolicy_bindingNotExist("citrixadc_lbvserver_auditsyslogpolicy_binding.demo", "tf_lbvserver3,tf_syslogaction2"),
				),
			},
		},
	})
}

func testAccCheckLbvserver_auditsyslogpolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lbvserver_auditsyslogpolicy_binding id is set")
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
			ResourceType:             "lbvserver_auditsyslogpolicy_binding",
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
			return fmt.Errorf("lbvserver_auditsyslogpolicy_binding %s not found", n)
		}

		return nil
	}
}
func testAccCheckLbvserver_auditsyslogpolicy_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*NetScalerNitroClient).client

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}

		idSlice := strings.SplitN(id, ",", 2)

		name := idSlice[0]
		policyname := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "lbvserver_auditsyslogpolicy_binding",
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

		if found {
			return fmt.Errorf("lbvserver_auditsyslogpolicy_binding %s not deleted", n)
		}

		return nil
	}
}
func testAccCheckLbvserver_auditsyslogpolicy_bindingDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_lbvserver_auditsyslogpolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Lbvserver_auditsyslogpolicy_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("lbvserver_auditsyslogpolicy_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
