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

const testAccAppfwprofile_fieldconsistency_binding_basic = `

	resource "citrixadc_appfwprofile" "tf_appfwprofile" {
		name                     = "tf_appfwprofile"
		type                     = ["HTML"]
	}
	resource "citrixadc_appfwprofile_fieldconsistency_binding" "tf_binding" {
		name              = citrixadc_appfwprofile.tf_appfwprofile.name
		fieldconsistency  = "tf_field"
		formactionurl_ffc = "^https://sd2\\-zgw\\.test\\.ctxns\\.com/api/document/content$"
		isautodeployed    = "NOTAUTODEPLOYED"
		state             = "DISABLED"
		alertonly         = "OFF"
		isregex_ffc       = "REGEX"
		comment           = "Testing"
	}
	resource "citrixadc_appfwprofile_fieldconsistency_binding" "tf_binding2" {
		name              = citrixadc_appfwprofile.tf_appfwprofile.name
		fieldconsistency  = "tf_field"
		formactionurl_ffc = "^https://sd2\\-zgw\\.test\\.ctxns\\.com/api/v1/resource/temp$"
		isautodeployed    = "NOTAUTODEPLOYED"
		state             = "DISABLED"
		alertonly         = "OFF"
		isregex_ffc       = "REGEX"
		comment           = "Testing"
	}
`

const testAccAppfwprofile_fieldconsistency_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion

	resource "citrixadc_appfwprofile" "tf_appfwprofile" {
		name                     = "tf_appfwprofile"
		type                     = ["HTML"]
	}
`

func TestAccAppfwprofile_fieldconsistency_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAppfwprofile_fieldconsistency_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAppfwprofile_fieldconsistency_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwprofile_fieldconsistency_bindingExist("citrixadc_appfwprofile_fieldconsistency_binding.tf_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_fieldconsistency_binding.tf_binding", "name", "tf_appfwprofile"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_fieldconsistency_binding.tf_binding", "fieldconsistency", "tf_field"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_fieldconsistency_binding.tf_binding", "formactionurl_ffc", "^https://sd2\\-zgw\\.test\\.ctxns\\.com/api/document/content$"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_fieldconsistency_binding.tf_binding", "isautodeployed", "NOTAUTODEPLOYED"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_fieldconsistency_binding.tf_binding", "state", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_fieldconsistency_binding.tf_binding", "alertonly", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_fieldconsistency_binding.tf_binding", "isregex_ffc", "REGEX"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_fieldconsistency_binding.tf_binding", "comment", "Testing"),
					testAccCheckAppfwprofile_fieldconsistency_bindingExist("citrixadc_appfwprofile_fieldconsistency_binding.tf_binding2", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_fieldconsistency_binding.tf_binding2", "name", "tf_appfwprofile"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_fieldconsistency_binding.tf_binding2", "fieldconsistency", "tf_field"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_fieldconsistency_binding.tf_binding2", "formactionurl_ffc", "^https://sd2\\-zgw\\.test\\.ctxns\\.com/api/v1/resource/temp$"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_fieldconsistency_binding.tf_binding2", "isautodeployed", "NOTAUTODEPLOYED"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_fieldconsistency_binding.tf_binding2", "state", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_fieldconsistency_binding.tf_binding2", "alertonly", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_fieldconsistency_binding.tf_binding2", "isregex_ffc", "REGEX"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_fieldconsistency_binding.tf_binding2", "comment", "Testing"),
				),
			},
			{
				Config: testAccAppfwprofile_fieldconsistency_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwprofile_fieldconsistency_bindingNotExist("citrixadc_appfwprofile_fieldconsistency_binding.tf_binding", "tf_appfwprofile,tf_field,^https://sd2\\-zgw\\.test\\.ctxns\\.com/api/document/content$"),
					testAccCheckAppfwprofile_fieldconsistency_bindingNotExist("citrixadc_appfwprofile_fieldconsistency_binding.tf_binding", "tf_appfwprofile,tf_field,^https://sd2\\-zgw\\.test\\.ctxns\\.com/api/v1/resource/temp$"),
				),
			},
		},
	})
}

func testAccCheckAppfwprofile_fieldconsistency_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No appfwprofile_fieldconsistency_binding id is set")
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

		idMap, _, err := utils.ParseIdString(bindingId, []string{"name", "fieldconsistency", "formactionurl_ffc"}, nil)
		if err != nil {
			return err
		}
		name := idMap["name"]
		fieldconsistency := idMap["fieldconsistency"]
		formactionurl_ffc := idMap["formactionurl_ffc"]

		findParams := service.FindParams{
			ResourceType:             "appfwprofile_fieldconsistency_binding",
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
			if v["fieldconsistency"].(string) == fieldconsistency {
				if v["formactionurl_ffc"].(string) == formactionurl_ffc {
					found = true
					break
				}
			}
		}

		if !found {
			return fmt.Errorf("appfwprofile_fieldconsistency_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckAppfwprofile_fieldconsistency_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}
		idMap, _, err := utils.ParseIdString(id, []string{"name", "fieldconsistency", "formactionurl_ffc"}, nil)
		if err != nil {
			return err
		}
		name := idMap["name"]
		fieldconsistency := idMap["fieldconsistency"]
		formactionurl_ffc := idMap["formactionurl_ffc"]

		findParams := service.FindParams{
			ResourceType:             "appfwprofile_fieldconsistency_binding",
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
			if v["fieldconsistency"].(string) == fieldconsistency {
				if v["formactionurl_ffc"].(string) == formactionurl_ffc {
					found = true
					break
				}
			}
		}

		if found {
			return fmt.Errorf("appfwprofile_fieldconsistency_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckAppfwprofile_fieldconsistency_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_appfwprofile_fieldconsistency_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Appfwprofile_fieldconsistency_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("appfwprofile_fieldconsistency_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccAppfwprofile_fieldconsistency_bindingDataSource_basic = `
	resource "citrixadc_appfwprofile" "tf_appfwprofile" {
		name                     = "tf_appfwprofile"
		type                     = ["HTML"]
	}
	resource "citrixadc_appfwprofile_fieldconsistency_binding" "tf_binding" {
		name              = citrixadc_appfwprofile.tf_appfwprofile.name
		fieldconsistency  = "tf_field"
		formactionurl_ffc = "^https://sd2\\-zgw\\.test\\.ctxns\\.com/api/document/content$"
		isautodeployed    = "NOTAUTODEPLOYED"
		state             = "DISABLED"
		alertonly         = "OFF"
		isregex_ffc       = "REGEX"
		comment           = "Testing"
	}

	data "citrixadc_appfwprofile_fieldconsistency_binding" "tf_binding_datasource" {
		name              = citrixadc_appfwprofile_fieldconsistency_binding.tf_binding.name
		fieldconsistency  = citrixadc_appfwprofile_fieldconsistency_binding.tf_binding.fieldconsistency
		formactionurl_ffc = citrixadc_appfwprofile_fieldconsistency_binding.tf_binding.formactionurl_ffc
	}
`

const testAccAppfwprofile_fieldconsistency_binding_upgrade_basic = `

	resource "citrixadc_appfwprofile" "tf_appfwprofile" {
		name                     = "tf_appfwprofile"
		type                     = ["HTML"]
	}
	resource "citrixadc_appfwprofile_fieldconsistency_binding" "tf_binding" {
		name              = citrixadc_appfwprofile.tf_appfwprofile.name
		fieldconsistency  = "tf_field"
		formactionurl_ffc = "^https://sd2\\-zgw\\.test\\.ctxns\\.com/api/document/content$"
		isautodeployed    = "NOTAUTODEPLOYED"
		state             = "DISABLED"
		alertonly         = "OFF"
		isregex_ffc       = "REGEX"
		comment           = "Testing"
	}
`

// TestAccAppfwprofile_fieldconsistency_binding_sdkv2StateUpgrade verifies that a
// binding created with the last SDK v2 release (legacy comma-separated ID) is read
// and upgraded correctly by the current Framework provider. Step 1 provisions the
// binding with the v2.2.0 registry provider (writing the legacy ID). Step 2 refreshes
// that state through the current Framework provider, which recomputes the ID into the
// new key:UrlEncode(value) format on Read.
func TestAccAppfwprofile_fieldconsistency_binding_sdkv2StateUpgrade(t *testing.T) {
	// Legacy SDK v2 id: name,fieldconsistency,formactionurl_ffc (see
	// citrixadc/resource_citrixadc_appfwprofile_fieldconsistency_binding.go d.SetId).
	legacyId := "tf_appfwprofile,tf_field,^https://sd2\\-zgw\\.test\\.ctxns\\.com/api/document/content$"
	// New Framework id: fieldconsistency:<enc>,formactionurl_ffc:<enc>,name:<enc>
	// derived exactly as the resource Create/SetAttrFromGet idParts do.
	newId := fmt.Sprintf(
		"fieldconsistency:%s,formactionurl_ffc:%s,name:%s",
		utils.UrlEncode("tf_field"),
		utils.UrlEncode("^https://sd2\\-zgw\\.test\\.ctxns\\.com/api/document/content$"),
		utils.UrlEncode("tf_appfwprofile"),
	)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckAppfwprofile_fieldconsistency_bindingDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.2.0",
					},
				},
				Config: testAccAppfwprofile_fieldconsistency_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwprofile_fieldconsistency_bindingExist("citrixadc_appfwprofile_fieldconsistency_binding.tf_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_fieldconsistency_binding.tf_binding", "id", legacyId),
				),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccAppfwprofile_fieldconsistency_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwprofile_fieldconsistency_bindingExist("citrixadc_appfwprofile_fieldconsistency_binding.tf_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_fieldconsistency_binding.tf_binding", "id", newId),
				),
			},
		},
	})
}

func TestAccAppfwprofile_fieldconsistency_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccAppfwprofile_fieldconsistency_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_fieldconsistency_binding.tf_binding_datasource", "name", "tf_appfwprofile"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_fieldconsistency_binding.tf_binding_datasource", "fieldconsistency", "tf_field"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_fieldconsistency_binding.tf_binding_datasource", "formactionurl_ffc", "^https://sd2\\-zgw\\.test\\.ctxns\\.com/api/document/content$"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_fieldconsistency_binding.tf_binding_datasource", "isautodeployed", "NOTAUTODEPLOYED"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_fieldconsistency_binding.tf_binding_datasource", "state", "DISABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_fieldconsistency_binding.tf_binding_datasource", "alertonly", "OFF"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_fieldconsistency_binding.tf_binding_datasource", "isregex_ffc", "REGEX"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_fieldconsistency_binding.tf_binding_datasource", "comment", "Testing"),
				),
			},
		},
	})
}

func TestAccAppfwprofile_fieldconsistency_binding_import(t *testing.T) {
	const resAddr = "citrixadc_appfwprofile_fieldconsistency_binding.tf_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAppfwprofile_fieldconsistency_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccAppfwprofile_fieldconsistency_binding_basic},
			{Config: testAccAppfwprofile_fieldconsistency_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
		},
	})
}
