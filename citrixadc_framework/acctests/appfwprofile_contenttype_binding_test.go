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

const testAccAppfwprofile_contenttype_binding_basic = `
	resource "citrixadc_appfwprofile" "tf_appfwprofile" {
		name                     = "tf_appfwprofile"
		type                     = ["HTML"]
	}
	resource "citrixadc_appfwprofile_contenttype_binding" "tf_binding1" {
		name           = citrixadc_appfwprofile.tf_appfwprofile.name
		contenttype    = "hello"
		state          = "ENABLED"
		alertonly      = "ON"
		isautodeployed = "NOTAUTODEPLOYED"
		comment        = "Testing"
	}
	resource "citrixadc_appfwprofile_contenttype_binding" "tf_binding2" {
		name           = citrixadc_appfwprofile.tf_appfwprofile.name
		contenttype    = "world"
		state          = "ENABLED"
		alertonly      = "ON"
		isautodeployed = "NOTAUTODEPLOYED"
		comment        = "Testing"
	}
`

const testAccAppfwprofile_contenttype_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion
	resource "citrixadc_appfwprofile" "tf_appfwprofile" {
		name                     = "tf_appfwprofile"
		type                     = ["HTML"]
	}
`

func TestAccAppfwprofile_contenttype_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAppfwprofile_contenttype_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAppfwprofile_contenttype_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwprofile_contenttype_bindingExist("citrixadc_appfwprofile_contenttype_binding.tf_binding1", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_contenttype_binding.tf_binding1", "contenttype", "hello"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_contenttype_binding.tf_binding1", "state", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_contenttype_binding.tf_binding1", "alertonly", "ON"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_contenttype_binding.tf_binding1", "isautodeployed", "NOTAUTODEPLOYED"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_contenttype_binding.tf_binding1", "comment", "Testing"),
					testAccCheckAppfwprofile_contenttype_bindingExist("citrixadc_appfwprofile_contenttype_binding.tf_binding2", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_contenttype_binding.tf_binding2", "contenttype", "world"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_contenttype_binding.tf_binding2", "state", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_contenttype_binding.tf_binding2", "alertonly", "ON"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_contenttype_binding.tf_binding2", "isautodeployed", "NOTAUTODEPLOYED"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_contenttype_binding.tf_binding2", "comment", "Testing"),
				),
			},
			{
				Config: testAccAppfwprofile_contenttype_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwprofile_contenttype_bindingNotExist("citrixadc_appfwprofile_contenttype_binding.tf_binding1", "tf_appfwprofile,hello"),
					testAccCheckAppfwprofile_contenttype_bindingNotExist("citrixadc_appfwprofile_contenttype_binding.tf_binding2", "tf_appfwprofile,world"),
				),
			},
		},
	})
}

func testAccCheckAppfwprofile_contenttype_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No appfwprofile_contenttype_binding id is set")
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

		idMap, _, err := utils.ParseIdString(bindingId, []string{"name", "contenttype"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", bindingId, err)
		}
		name := idMap["name"]
		contenttype := idMap["contenttype"]

		findParams := service.FindParams{
			ResourceType:             "appfwprofile_contenttype_binding",
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
			if v["contenttype"].(string) == contenttype {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("appfwprofile_contenttype_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckAppfwprofile_contenttype_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}
		idSlice := strings.SplitN(id, ",", 2)

		name := idSlice[0]
		contenttype := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "appfwprofile_contenttype_binding",
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
			if v["contenttype"].(string) == contenttype {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("appfwprofile_contenttype_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckAppfwprofile_contenttype_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_appfwprofile_contenttype_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Appfwprofile_contenttype_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("appfwprofile_contenttype_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccAppfwprofile_contenttype_bindingDataSource_basic = `
	resource "citrixadc_appfwprofile" "tf_appfwprofile" {
		name                     = "tf_appfwprofile"
		type                     = ["HTML"]
	}
	resource "citrixadc_appfwprofile_contenttype_binding" "tf_binding1" {
		name           = citrixadc_appfwprofile.tf_appfwprofile.name
		contenttype    = "hello"
		state          = "ENABLED"
		alertonly      = "ON"
		isautodeployed = "NOTAUTODEPLOYED"
		comment        = "Testing"
	}

	data "citrixadc_appfwprofile_contenttype_binding" "tf_binding1" {
		name        = citrixadc_appfwprofile_contenttype_binding.tf_binding1.name
		contenttype = citrixadc_appfwprofile_contenttype_binding.tf_binding1.contenttype
	}
`

func TestAccAppfwprofile_contenttype_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAppfwprofile_contenttype_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_contenttype_binding.tf_binding1", "name", "tf_appfwprofile"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_contenttype_binding.tf_binding1", "contenttype", "hello"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_contenttype_binding.tf_binding1", "state", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_contenttype_binding.tf_binding1", "alertonly", "ON"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_contenttype_binding.tf_binding1", "isautodeployed", "NOTAUTODEPLOYED"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_contenttype_binding.tf_binding1", "comment", "Testing"),
				),
			},
		},
	})
}

// testAccAppfwprofile_contenttype_binding_upgrade_basic reuses the _basic config
// (binding + its prerequisite appfwprofile). It is valid under BOTH the SDK v2 2.2.0
// schema and the current Framework schema because the migration restored the SDK v2
// attribute names.
const testAccAppfwprofile_contenttype_binding_upgrade_basic = `
	resource "citrixadc_appfwprofile" "tf_appfwprofile" {
		name = "tf_appfwprofile"
		type = ["HTML"]
	}
	resource "citrixadc_appfwprofile_contenttype_binding" "tf_binding1" {
		name           = citrixadc_appfwprofile.tf_appfwprofile.name
		contenttype    = "hello"
		state          = "ENABLED"
		alertonly      = "ON"
		isautodeployed = "NOTAUTODEPLOYED"
		comment        = "Testing"
	}
`

// TestAccAppfwprofile_contenttype_binding_sdkv2StateUpgrade verifies that state
// written by the last SDK v2 release (legacy comma-separated ID) is correctly
// upgraded when the same config is subsequently managed by the current Framework
// provider. Step 1 creates the binding with citrix/citrixadc 2.2.0 (writes the
// legacy id "tf_appfwprofile,hello"). Step 2 refreshes/plans/applies the same config
// through the Framework provider, exercising ParseIdString on the legacy id; because
// the Framework recomputes the id on Read (SetAttrFromGet), the id upgrades to the
// new "key:value" form.
func TestAccAppfwprofile_contenttype_binding_sdkv2StateUpgrade(t *testing.T) {
	resourceAddr := "citrixadc_appfwprofile_contenttype_binding.tf_binding1"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckAppfwprofile_contenttype_bindingDestroy,
		Steps: []resource.TestStep{
			// Step 1: create with the last SDK v2 release -> state carries the legacy id.
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.2.0",
					},
				},
				Config: testAccAppfwprofile_contenttype_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwprofile_contenttype_bindingExist(resourceAddr, nil),
					resource.TestCheckResourceAttr(resourceAddr, "id", "tf_appfwprofile,hello"),
				),
			},
			// Step 2: refresh/plan/apply the SAME config through the current Framework
			// provider. The legacy-id state is read via ParseIdString and the id is
			// recomputed to the new key:value format.
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccAppfwprofile_contenttype_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwprofile_contenttype_bindingExist(resourceAddr, nil),
					resource.TestCheckResourceAttr(resourceAddr, "id", "contenttype:hello,name:tf_appfwprofile"),
				),
			},
		},
	})
}

func TestAccAppfwprofile_contenttype_binding_import(t *testing.T) {
	const resAddr = "citrixadc_appfwprofile_contenttype_binding.tf_binding1"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAppfwprofile_contenttype_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccAppfwprofile_contenttype_binding_basic},
			{Config: testAccAppfwprofile_contenttype_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
		},
	})
}
