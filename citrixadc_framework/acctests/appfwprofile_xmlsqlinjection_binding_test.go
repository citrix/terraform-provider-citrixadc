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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testAccAppfwprofile_xmlsqlinjection_binding_basic = `
	resource "citrixadc_appfwprofile" "tf_appfwprofile" {
		name                     = "tf_appfwprofile"
		type                     = ["HTML"]
	}
	resource "citrixadc_appfwprofile_xmlsqlinjection_binding" "tf_binding1" {
		name                    = citrixadc_appfwprofile.tf_appfwprofile.name
		xmlsqlinjection         = "hello"
		alertonly               = "ON"
		isautodeployed          = "AUTODEPLOYED"
		state                   = "ENABLED"
		comment                 = "Testing"
	}
	resource "citrixadc_appfwprofile_xmlsqlinjection_binding" "tf_binding2" {
		name                    = citrixadc_appfwprofile.tf_appfwprofile.name
		xmlsqlinjection         = "world"
		alertonly               = "ON"
		isautodeployed          = "AUTODEPLOYED"
		state                   = "ENABLED"
		comment                 = "Testing"
	}
`

const testAccAppfwprofile_xmlsqlinjection_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion
	resource "citrixadc_appfwprofile" "tf_appfwprofile" {
		name                     = "tf_appfwprofile"
		type                     = ["HTML"]
	}
`

func TestAccAppfwprofile_xmlsqlinjection_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAppfwprofile_xmlsqlinjection_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAppfwprofile_xmlsqlinjection_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwprofile_xmlsqlinjection_bindingExist("citrixadc_appfwprofile_xmlsqlinjection_binding.tf_binding1", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_xmlsqlinjection_binding.tf_binding1", "name", "tf_appfwprofile"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_xmlsqlinjection_binding.tf_binding1", "xmlsqlinjection", "hello"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_xmlsqlinjection_binding.tf_binding1", "alertonly", "ON"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_xmlsqlinjection_binding.tf_binding1", "isautodeployed", "AUTODEPLOYED"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_xmlsqlinjection_binding.tf_binding1", "state", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_xmlsqlinjection_binding.tf_binding1", "comment", "Testing"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_xmlsqlinjection_binding.tf_binding1", "as_scan_location_xmlsql", "ELEMENT"),
					testAccCheckAppfwprofile_xmlsqlinjection_bindingExist("citrixadc_appfwprofile_xmlsqlinjection_binding.tf_binding2", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_xmlsqlinjection_binding.tf_binding2", "name", "tf_appfwprofile"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_xmlsqlinjection_binding.tf_binding2", "xmlsqlinjection", "world"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_xmlsqlinjection_binding.tf_binding2", "alertonly", "ON"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_xmlsqlinjection_binding.tf_binding2", "isautodeployed", "AUTODEPLOYED"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_xmlsqlinjection_binding.tf_binding2", "state", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_xmlsqlinjection_binding.tf_binding2", "comment", "Testing"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_xmlsqlinjection_binding.tf_binding2", "as_scan_location_xmlsql", "ELEMENT"),
				),
			},
			{
				Config: testAccAppfwprofile_xmlsqlinjection_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwprofile_xmlsqlinjection_bindingNotExist("citrixadc_appfwprofile_xmlsqlinjection_binding.tf_binding1", "tf_appfwprofile,hello", "ELEMENT"),
					testAccCheckAppfwprofile_xmlsqlinjection_bindingNotExist("citrixadc_appfwprofile_xmlsqlinjection_binding.tf_binding2", "tf_appfwprofile,world", "ELEMENT"),
				),
			},
		},
	})
}

func testAccCheckAppfwprofile_xmlsqlinjection_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No appfwprofile_xmlsqlinjection_binding id is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		bindingId := rs.Primary.ID

		idSlice := strings.Split(bindingId, ",")

		name := idSlice[0]
		xmlsqlinjection := idSlice[1]
		locationName := idSlice[2]
		findParams := service.FindParams{
			ResourceType:             "appfwprofile_xmlsqlinjection_binding",
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
			if v["xmlsqlinjection"].(string) == xmlsqlinjection {
				if v["as_scan_location_xmlsql"].(string) == locationName {
					found = true
					break
				}
			}
		}

		if !found {
			return fmt.Errorf("appfwprofile_xmlsqlinjection_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckAppfwprofile_xmlsqlinjection_bindingNotExist(n string, id string, locationName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}
		idSlice := strings.Split(id, ",")

		name := idSlice[0]
		xmlsqlinjection := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "appfwprofile_xmlsqlinjection_binding",
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
			if v["xmlsqlinjection"].(string) == xmlsqlinjection {
				if v["as_scan_location_xmlsql"].(string) == locationName {
					found = true
					break
				}
			}
		}

		if found {
			return fmt.Errorf("appfwprofile_xmlsqlinjection_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckAppfwprofile_xmlsqlinjection_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_appfwprofile_xmlsqlinjection_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		bindingId := rs.Primary.ID

		idSlice := strings.Split(bindingId, ",")

		name := idSlice[0]
		_, err := client.FindResource(service.Appfwprofile_xmlsqlinjection_binding.Type(), name)
		if err == nil {
			return fmt.Errorf("appfwprofile_xmlsqlinjection_binding %s still exists", name)
		}

	}

	return nil
}
