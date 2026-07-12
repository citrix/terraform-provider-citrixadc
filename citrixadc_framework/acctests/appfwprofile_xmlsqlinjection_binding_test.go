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
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
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

		idMap, _, err := utils.ParseIdString(bindingId, []string{"name", "xmlsqlinjection", "as_scan_location_xmlsql"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", bindingId, err)
		}
		name := idMap["name"]
		xmlsqlinjection := idMap["xmlsqlinjection"]
		locationName := idMap["as_scan_location_xmlsql"]
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

		idMap, _, parseErr := utils.ParseIdString(bindingId, []string{"name", "xmlsqlinjection", "as_scan_location_xmlsql"}, nil)
		if parseErr != nil {
			return fmt.Errorf("Error parsing ID %s: %v", bindingId, parseErr)
		}
		name := idMap["name"]
		_, err := client.FindResource(service.Appfwprofile_xmlsqlinjection_binding.Type(), name)
		if err == nil {
			return fmt.Errorf("appfwprofile_xmlsqlinjection_binding %s still exists", name)
		}

	}

	return nil
}

const testAccAppfwprofileXmlsqlinjectionBindingDataSource_basic = `
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

	data "citrixadc_appfwprofile_xmlsqlinjection_binding" "tf_binding1_data" {
		name                     = citrixadc_appfwprofile_xmlsqlinjection_binding.tf_binding1.name
		xmlsqlinjection          = citrixadc_appfwprofile_xmlsqlinjection_binding.tf_binding1.xmlsqlinjection
		as_scan_location_xmlsql  = "ELEMENT"
		depends_on               = [citrixadc_appfwprofile_xmlsqlinjection_binding.tf_binding1]
	}

	data "citrixadc_appfwprofile_xmlsqlinjection_binding" "tf_binding2_data" {
		name                     = citrixadc_appfwprofile_xmlsqlinjection_binding.tf_binding2.name
		xmlsqlinjection          = citrixadc_appfwprofile_xmlsqlinjection_binding.tf_binding2.xmlsqlinjection
		as_scan_location_xmlsql  = "ELEMENT"
		depends_on               = [citrixadc_appfwprofile_xmlsqlinjection_binding.tf_binding2]
	}
`

func TestAccAppfwprofileXmlsqlinjectionBindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAppfwprofileXmlsqlinjectionBindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_xmlsqlinjection_binding.tf_binding1_data", "name", "tf_appfwprofile"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_xmlsqlinjection_binding.tf_binding1_data", "xmlsqlinjection", "hello"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_xmlsqlinjection_binding.tf_binding1_data", "alertonly", "OFF"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_xmlsqlinjection_binding.tf_binding1_data", "isautodeployed", "NOTAUTODEPLOYED"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_xmlsqlinjection_binding.tf_binding1_data", "state", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_xmlsqlinjection_binding.tf_binding1_data", "comment", "Testing"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_xmlsqlinjection_binding.tf_binding1_data", "as_scan_location_xmlsql", "ELEMENT"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_xmlsqlinjection_binding.tf_binding2_data", "name", "tf_appfwprofile"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_xmlsqlinjection_binding.tf_binding2_data", "xmlsqlinjection", "world"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_xmlsqlinjection_binding.tf_binding2_data", "alertonly", "OFF"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_xmlsqlinjection_binding.tf_binding2_data", "isautodeployed", "NOTAUTODEPLOYED"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_xmlsqlinjection_binding.tf_binding2_data", "state", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_xmlsqlinjection_binding.tf_binding2_data", "comment", "Testing"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_xmlsqlinjection_binding.tf_binding2_data", "as_scan_location_xmlsql", "ELEMENT"),
				),
			},
		},
	})
}

const testAccAppfwprofile_xmlsqlinjection_binding_upgrade_basic = `
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
`

func TestAccAppfwprofile_xmlsqlinjection_binding_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckAppfwprofile_xmlsqlinjection_bindingDestroy,
		Steps: []resource.TestStep{
			{
				// Step 1: create the binding with the last SDK v2 release (2.2.0),
				// which writes state using the legacy comma-joined id.
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.2.0",
					},
				},
				Config: testAccAppfwprofile_xmlsqlinjection_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwprofile_xmlsqlinjection_bindingExist("citrixadc_appfwprofile_xmlsqlinjection_binding.tf_binding1", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_xmlsqlinjection_binding.tf_binding1", "id", "tf_appfwprofile,hello,ELEMENT"),
				),
			},
			{
				// Step 2: refresh/plan the legacy-id state through the current
				// framework provider. Read exercises ParseIdString on the legacy id
				// and SetAttrFromGet recomputes the id into the new key:value form.
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccAppfwprofile_xmlsqlinjection_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwprofile_xmlsqlinjection_bindingExist("citrixadc_appfwprofile_xmlsqlinjection_binding.tf_binding1", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_xmlsqlinjection_binding.tf_binding1", "id", "as_scan_location_xmlsql:ELEMENT,name:tf_appfwprofile,xmlsqlinjection:hello"),
				),
			},
		},
	})
}

func TestAccAppfwprofile_xmlsqlinjection_binding_import(t *testing.T) {
	const resAddr = "citrixadc_appfwprofile_xmlsqlinjection_binding.tf_binding1"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAppfwprofile_xmlsqlinjection_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccAppfwprofile_xmlsqlinjection_binding_basic},
			{Config: testAccAppfwprofile_xmlsqlinjection_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{"alertonly", "isautodeployed"}},
		},
	})
}
