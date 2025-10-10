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

const testAccAppfwprofile_logexpression_binding_basic = `
	resource "citrixadc_appfwprofile" "tf_appfwprofile" {
		name                     = "tf_appfwprofile"
		type                     = ["HTML"]
	}
	resource "citrixadc_appfwprofile_logexpression_binding" "tf_binding1" {
		name             = citrixadc_appfwprofile.tf_appfwprofile.name
		logexpression    = "tf_logexp"
		as_logexpression = "HTTP.REQ.IS_VALID"
		alertonly        = "ON"
		isautodeployed   = "AUTODEPLOYED"
		comment          = "Testing"
		state            = "ENABLED"
	}
	resource "citrixadc_appfwprofile_logexpression_binding" "tf_binding2" {
		name             = citrixadc_appfwprofile.tf_appfwprofile.name
		logexpression    = "new_tf_logexp"
		as_logexpression = "HTTP.REQ.IS_VALID"
		alertonly        = "ON"
		isautodeployed   = "AUTODEPLOYED"
		comment          = "Testing"
		state            = "ENABLED"
	}
`

const testAccAppfwprofile_logexpression_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion
	resource "citrixadc_appfwprofile" "tf_appfwprofile" {
		name                     = "tf_appfwprofile"
		type                     = ["HTML"]
	}
`

func TestAccAppfwprofile_logexpression_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAppfwprofile_logexpression_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAppfwprofile_logexpression_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwprofile_logexpression_bindingExist("citrixadc_appfwprofile_logexpression_binding.tf_binding1", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_logexpression_binding.tf_binding1", "name", "tf_appfwprofile"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_logexpression_binding.tf_binding1", "logexpression", "tf_logexp"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_logexpression_binding.tf_binding1", "as_logexpression", "HTTP.REQ.IS_VALID"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_logexpression_binding.tf_binding1", "alertonly", "ON"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_logexpression_binding.tf_binding1", "isautodeployed", "AUTODEPLOYED"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_logexpression_binding.tf_binding1", "comment", "Testing"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_logexpression_binding.tf_binding1", "state", "ENABLED"),
					testAccCheckAppfwprofile_logexpression_bindingExist("citrixadc_appfwprofile_logexpression_binding.tf_binding2", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_logexpression_binding.tf_binding2", "name", "tf_appfwprofile"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_logexpression_binding.tf_binding2", "logexpression", "new_tf_logexp"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_logexpression_binding.tf_binding2", "as_logexpression", "HTTP.REQ.IS_VALID"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_logexpression_binding.tf_binding2", "alertonly", "ON"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_logexpression_binding.tf_binding2", "isautodeployed", "AUTODEPLOYED"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_logexpression_binding.tf_binding2", "comment", "Testing"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_logexpression_binding.tf_binding2", "state", "ENABLED"),
				),
			},
			{
				Config: testAccAppfwprofile_logexpression_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwprofile_logexpression_bindingNotExist("citrixadc_appfwprofile_logexpression_binding.tf_binding1", "tf_appfwprofile,tf_logexp"),
					testAccCheckAppfwprofile_logexpression_bindingNotExist("citrixadc_appfwprofile_logexpression_binding.tf_binding2", "tf_appfwprofile,new_tf_logexp"),
				),
			},
		},
	})
}

func testAccCheckAppfwprofile_logexpression_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No appfwprofile_logexpression_binding id is set")
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
		logexpression := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "appfwprofile_logexpression_binding",
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
			if v["logexpression"].(string) == logexpression {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("appfwprofile_logexpression_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckAppfwprofile_logexpression_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*NetScalerNitroClient).client

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}
		idSlice := strings.SplitN(id, ",", 2)

		name := idSlice[0]
		logexpression := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "appfwprofile_logexpression_binding",
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
			if v["logexpression"].(string) == logexpression {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("appfwprofile_logexpression_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckAppfwprofile_logexpression_bindingDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_appfwprofile_logexpression_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource("appfwprofile_logexpression_binding", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("appfwprofile_logexpression_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
