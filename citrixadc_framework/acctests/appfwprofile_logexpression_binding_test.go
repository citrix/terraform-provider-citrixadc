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
	"testing"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
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
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAppfwprofile_logexpression_bindingDestroy,
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

		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		bindingId := rs.Primary.ID

		idMap, _, err := utils.ParseIdString(bindingId, []string{"name", "logexpression"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", bindingId, err)
		}
		name := idMap["name"]
		logexpression := idMap["logexpression"]

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
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		idMap, _, err := utils.ParseIdString(id, []string{"name", "logexpression"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", id, err)
		}
		name := idMap["name"]
		logexpression := idMap["logexpression"]

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
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_appfwprofile_logexpression_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("appfwprofile_logexpression_binding", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("appfwprofile_logexpression_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccAppfwprofile_logexpression_bindingDataSource_basic = `
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

	data "citrixadc_appfwprofile_logexpression_binding" "tf_binding1" {
		name          = citrixadc_appfwprofile_logexpression_binding.tf_binding1.name
		logexpression = citrixadc_appfwprofile_logexpression_binding.tf_binding1.logexpression
	}
`

func TestAccAppfwprofile_logexpression_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAppfwprofile_logexpression_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_logexpression_binding.tf_binding1", "name", "tf_appfwprofile"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_logexpression_binding.tf_binding1", "logexpression", "tf_logexp"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_logexpression_binding.tf_binding1", "as_logexpression", "HTTP.REQ.IS_VALID"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_logexpression_binding.tf_binding1", "alertonly", "OFF"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_logexpression_binding.tf_binding1", "isautodeployed", "NOTAUTODEPLOYED"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_logexpression_binding.tf_binding1", "comment", "Testing"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_logexpression_binding.tf_binding1", "state", "ENABLED"),
				),
			},
		},
	})
}

const testAccAppfwprofile_logexpression_binding_upgrade_basic = `
	resource "citrixadc_appfwprofile" "tf_appfwprofile" {
		name = "tf_appfwprofile"
		type = ["HTML"]
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
`

func TestAccAppfwprofile_logexpression_binding_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckAppfwprofile_logexpression_bindingDestroy,
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
				Config: testAccAppfwprofile_logexpression_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwprofile_logexpression_bindingExist("citrixadc_appfwprofile_logexpression_binding.tf_binding1", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_logexpression_binding.tf_binding1", "id", "tf_appfwprofile,tf_logexp"),
				),
			},
			{
				// Step 2: refresh/plan the legacy-id state through the current
				// framework provider. Read exercises ParseIdString on the legacy id
				// and SetAttrFromGet recomputes the id into the new key:value form.
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccAppfwprofile_logexpression_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwprofile_logexpression_bindingExist("citrixadc_appfwprofile_logexpression_binding.tf_binding1", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_logexpression_binding.tf_binding1", "id", "logexpression:tf_logexp,name:tf_appfwprofile"),
				),
			},
		},
	})
}

func TestAccAppfwprofile_logexpression_binding_import(t *testing.T) {
	const resAddr = "citrixadc_appfwprofile_logexpression_binding.tf_binding1"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAppfwprofile_logexpression_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccAppfwprofile_logexpression_binding_basic},
			{Config: testAccAppfwprofile_logexpression_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{"alertonly", "isautodeployed"}},
		},
	})
}
