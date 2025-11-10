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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"strings"
	"testing"
)

const testAccAaagroup_auditsyslogpolicy_binding_basic = `

	resource "citrixadc_aaagroup_auditsyslogpolicy_binding" "tf_aaagroup_auditsyslogpolicy_binding" {
		groupname = citrixadc_aaagroup.tf_aaagroup.groupname
		policy    = citrixadc_auditsyslogpolicy.tf_auditsyslogpolicy.name
		priority  = 100
	}
	resource "citrixadc_aaagroup" "tf_aaagroup" {
		groupname = "my_group"
		weight    = 100
		loggedin  = false
	}
  
	resource "citrixadc_auditsyslogaction" "tf_syslogaction" {
		name       = "tf_syslogaction"
		serverip   = "10.78.60.33"
		serverport = 514
		loglevel = [
		"ERROR",
		"NOTICE",
		]
	}
	resource "citrixadc_auditsyslogpolicy" "tf_auditsyslogpolicy" {
		name   = "tf_auditsyslogpolicy"
		rule   = "ns_true"
		action = citrixadc_auditsyslogaction.tf_syslogaction.name
	}
`

const testAccAaagroup_auditsyslogpolicy_binding_basic_step2 = `
	resource "citrixadc_aaagroup" "tf_aaagroup" {
		groupname = "my_group"
		weight    = 100
		loggedin  = false
	}
	resource "citrixadc_auditsyslogaction" "tf_syslogaction" {
		name       = "tf_syslogaction"
		serverip   = "10.78.60.33"
		serverport = 514
		loglevel = [
		"ERROR",
		"NOTICE",
		]
	}
	resource "citrixadc_auditsyslogpolicy" "tf_auditsyslogpolicy" {
		name   = "tf_auditsyslogpolicy"
		rule   = "ns_true"
		action = citrixadc_auditsyslogaction.tf_syslogaction.name
	}
`

func TestAccAaagroup_auditsyslogpolicy_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckAaagroup_auditsyslogpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAaagroup_auditsyslogpolicy_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaagroup_auditsyslogpolicy_bindingExist("citrixadc_aaagroup_auditsyslogpolicy_binding.tf_aaagroup_auditsyslogpolicy_binding", nil),
				),
			},
			{
				Config: testAccAaagroup_auditsyslogpolicy_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaagroup_auditsyslogpolicy_bindingNotExist("citrixadc_aaagroup_auditsyslogpolicy_binding.tf_aaagroup_auditsyslogpolicy_binding", "my_group,tf_auditsyslogpolicy"),
				),
			},
		},
	})
}

func testAccCheckAaagroup_auditsyslogpolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No aaagroup_auditsyslogpolicy_binding id is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		// Use the shared utility function to get a configured client
		client, err := testAccGetClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		bindingId := rs.Primary.ID

		idSlice := strings.SplitN(bindingId, ",", 2)

		groupname := idSlice[0]
		policy := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "aaagroup_auditsyslogpolicy_binding",
			ResourceName:             groupname,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the matching policy
		found := false
		for _, v := range dataArr {
			if v["policy"].(string) == policy {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("aaagroup_auditsyslogpolicy_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckAaagroup_auditsyslogpolicy_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}
		idSlice := strings.SplitN(id, ",", 2)

		groupname := idSlice[0]
		policy := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "aaagroup_auditsyslogpolicy_binding",
			ResourceName:             groupname,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to hopefully not find the one with the matching policy
		found := false
		for _, v := range dataArr {
			if v["policy"].(string) == policy {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("aaagroup_auditsyslogpolicy_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckAaagroup_auditsyslogpolicy_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_aaagroup_auditsyslogpolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Aaagroup_auditsyslogpolicy_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("aaagroup_auditsyslogpolicy_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
