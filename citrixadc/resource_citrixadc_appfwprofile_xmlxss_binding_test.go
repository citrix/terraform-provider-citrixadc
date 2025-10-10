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

const testAccAppfwprofile_xmlxss_binding_basic = `
	resource "citrixadc_appfwprofile" "tf_appfwprofile" {
		name                     = "tf_appfwprofile"
		type                     = ["HTML"]
	}
	resource "citrixadc_appfwprofile_xmlxss_binding" "tf_binding1" {
		name                    = citrixadc_appfwprofile.tf_appfwprofile.name
		xmlxss                  = "tf_xmlxss"
		state                   = "ENABLED"
		alertonly               = "ON"
		isregex_xmlxss          = "NOTREGEX"
		isautodeployed          = "AUTODEPLOYED"
	}
	resource "citrixadc_appfwprofile_xmlxss_binding" "tf_binding2" {
		name                    = citrixadc_appfwprofile.tf_appfwprofile.name
		xmlxss                  = "new_tf_xmlxss"
		state                   = "ENABLED"
		alertonly               = "ON"
		isregex_xmlxss          = "NOTREGEX"
		isautodeployed          = "AUTODEPLOYED"
	}
`

const testAccAppfwprofile_xmlxss_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion
	resource "citrixadc_appfwprofile" "tf_appfwprofile" {
		name                     = "tf_appfwprofile"
		type                     = ["HTML"]
	}
`

func TestAccAppfwprofile_xmlxss_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAppfwprofile_xmlxss_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAppfwprofile_xmlxss_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwprofile_xmlxss_bindingExist("citrixadc_appfwprofile_xmlxss_binding.tf_binding1", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_xmlxss_binding.tf_binding1", "name", "tf_appfwprofile"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_xmlxss_binding.tf_binding1", "xmlxss", "tf_xmlxss"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_xmlxss_binding.tf_binding1", "state", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_xmlxss_binding.tf_binding1", "alertonly", "ON"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_xmlxss_binding.tf_binding1", "isautodeployed", "AUTODEPLOYED"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_xmlxss_binding.tf_binding1", "isregex_xmlxss", "NOTREGEX"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_xmlxss_binding.tf_binding1", "as_scan_location_xmlxss", "ELEMENT"),
					testAccCheckAppfwprofile_xmlxss_bindingExist("citrixadc_appfwprofile_xmlxss_binding.tf_binding2", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_xmlxss_binding.tf_binding2", "name", "tf_appfwprofile"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_xmlxss_binding.tf_binding2", "xmlxss", "new_tf_xmlxss"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_xmlxss_binding.tf_binding2", "state", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_xmlxss_binding.tf_binding2", "alertonly", "ON"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_xmlxss_binding.tf_binding2", "isautodeployed", "AUTODEPLOYED"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_xmlxss_binding.tf_binding2", "isregex_xmlxss", "NOTREGEX"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_xmlxss_binding.tf_binding2", "as_scan_location_xmlxss", "ELEMENT"),
				),
			},
			{
				Config: testAccAppfwprofile_xmlxss_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwprofile_xmlxss_bindingNotExist("citrixadc_appfwprofile_xmlxss_binding.tf_binding", "tf_appfwprofile,tf_xmlxss,ELEMENT"),
				),
			},
		},
	})
}

func testAccCheckAppfwprofile_xmlxss_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No appfwprofile_xmlxss_binding id is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		client := testAccProvider.Meta().(*NetScalerNitroClient).client

		bindingId := rs.Primary.ID

		idSlice := strings.SplitN(bindingId, ",", 3)

		name := idSlice[0]
		xmlxss := idSlice[1]
		as_scan_location_xmlxss := idSlice[2]

		findParams := service.FindParams{
			ResourceType:             "appfwprofile_xmlxss_binding",
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
			if v["xmlxss"].(string) == xmlxss {
				if v["as_scan_location_xmlxss"].(string) == as_scan_location_xmlxss {
					found = true
					break
				}
			}
		}

		if !found {
			return fmt.Errorf("appfwprofile_xmlxss_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckAppfwprofile_xmlxss_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*NetScalerNitroClient).client

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}
		idSlice := strings.SplitN(id, ",", 3)

		name := idSlice[0]
		xmlxss := idSlice[1]
		as_scan_location_xmlxss := idSlice[2]

		findParams := service.FindParams{
			ResourceType:             "appfwprofile_xmlxss_binding",
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
			if v["xmlxss"].(string) == xmlxss {
				if v["as_scan_location_xmlxss"].(string) == as_scan_location_xmlxss {
					found = true
					break
				}
			}
		}

		if found {
			return fmt.Errorf("appfwprofile_xmlxss_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckAppfwprofile_xmlxss_bindingDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_appfwprofile_xmlxss_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Appfwprofile_xmlxss_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("appfwprofile_xmlxss_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
