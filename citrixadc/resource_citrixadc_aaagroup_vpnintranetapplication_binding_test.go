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

const testAccAaagroup_vpnintranetapplication_binding_basic = `

	resource "citrixadc_aaagroup_vpnintranetapplication_binding" "tf_aaagroup_vpnintranetapplication_binding" {
		groupname           = citrixadc_aaagroup.tf_aaagroup.groupname
		intranetapplication = citrixadc_vpnintranetapplication.tf_vpnintranetapplication.intranetapplication
	}

	resource "citrixadc_aaagroup" "tf_aaagroup" {
		groupname = "my_group"
		weight    = 100
		loggedin  = false
	}
	resource "citrixadc_vpnintranetapplication" "tf_vpnintranetapplication" {
		intranetapplication = "tf_vpnintranetapplication"
		protocol            = "UDP"
		destip              = "2.3.6.5"
		interception        = "TRANSPARENT"
	}
  
`

const testAccAaagroup_vpnintranetapplication_binding_basic_step2 = `	
	resource "citrixadc_aaagroup" "tf_aaagroup" {
		groupname = "my_group"
		weight    = 100
		loggedin  = false
	}
	resource "citrixadc_vpnintranetapplication" "tf_vpnintranetapplication" {
		intranetapplication = "tf_vpnintranetapplication"
		protocol            = "UDP"
		destip              = "2.3.6.5"
		interception        = "TRANSPARENT"
	}
`

func TestAccAaagroup_vpnintranetapplication_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckAaagroup_vpnintranetapplication_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAaagroup_vpnintranetapplication_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaagroup_vpnintranetapplication_bindingExist("citrixadc_aaagroup_vpnintranetapplication_binding.tf_aaagroup_vpnintranetapplication_binding", nil),
				),
			},
			{
				Config: testAccAaagroup_vpnintranetapplication_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaagroup_vpnintranetapplication_bindingNotExist("citrixadc_aaagroup_vpnintranetapplication_binding.tf_aaagroup_vpnintranetapplication_binding", "my_group,tf_vpnintranetapplication"),
				),
			},
		},
	})
}

func testAccCheckAaagroup_vpnintranetapplication_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No aaagroup_vpnintranetapplication_binding id is set")
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
		intranetapplication := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "aaagroup_vpnintranetapplication_binding",
			ResourceName:             groupname,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the matching intranetapplication
		found := false
		for _, v := range dataArr {
			if v["intranetapplication"].(string) == intranetapplication {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("aaagroup_vpnintranetapplication_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckAaagroup_vpnintranetapplication_bindingNotExist(n string, id string) resource.TestCheckFunc {
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
		intranetapplication := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "aaagroup_vpnintranetapplication_binding",
			ResourceName:             groupname,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to hopefully not find the one with the matching intranetapplication
		found := false
		for _, v := range dataArr {
			if v["intranetapplication"].(string) == intranetapplication {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("aaagroup_vpnintranetapplication_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckAaagroup_vpnintranetapplication_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_aaagroup_vpnintranetapplication_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Aaagroup_vpnintranetapplication_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("aaagroup_vpnintranetapplication_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
