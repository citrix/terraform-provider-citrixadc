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

// Step 1 creates the parent appfwprofile and binds a bypass-list entry to it.
// All attributes on this binding are RequiresReplace (no in-place update); the
// binding is keyed by name + as_bypass_list + as_bypass_list_value_type +
// as_bypass_list_location. The participating parent entity config
// (citrixadc_appfwprofile with name + type) is reused from appfwprofile_test.go.
const testAccAppfwprofileBypasslistBinding_basic_step1 = `

resource "citrixadc_appfwprofile" "tf_appfwprofile" {
  name = "tf_appfwprofile_bypasslist"
  type = ["HTML"]
}

resource "citrixadc_appfwprofile_bypasslist_binding" "tf_appfwprofile_bypasslist_binding" {
  name                      = citrixadc_appfwprofile.tf_appfwprofile.name
  as_bypass_list            = "http://www.example.com/bypass"
  as_bypass_list_value_type = "literal"
  as_bypass_list_location   = "URL"
  as_bypass_list_action     = "log"
  state                     = "ENABLED"
  comment                   = "Testing"

  depends_on = [citrixadc_appfwprofile.tf_appfwprofile]
}
`

// Step 2 drops the binding to confirm it is deleted (only the parent remains).
const testAccAppfwprofileBypasslistBinding_basic_step2 = `

resource "citrixadc_appfwprofile" "tf_appfwprofile" {
  name = "tf_appfwprofile_bypasslist"
  type = ["HTML"]
}
`

func TestAccAppfwprofileBypasslistBinding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAppfwprofileBypasslistBindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAppfwprofileBypasslistBinding_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwprofileBypasslistBindingExist("citrixadc_appfwprofile_bypasslist_binding.tf_appfwprofile_bypasslist_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_bypasslist_binding.tf_appfwprofile_bypasslist_binding", "name", "tf_appfwprofile_bypasslist"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_bypasslist_binding.tf_appfwprofile_bypasslist_binding", "as_bypass_list", "http://www.example.com/bypass"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_bypasslist_binding.tf_appfwprofile_bypasslist_binding", "as_bypass_list_value_type", "literal"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_bypasslist_binding.tf_appfwprofile_bypasslist_binding", "as_bypass_list_location", "URL"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_bypasslist_binding.tf_appfwprofile_bypasslist_binding", "as_bypass_list_action", "log"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_bypasslist_binding.tf_appfwprofile_bypasslist_binding", "state", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_bypasslist_binding.tf_appfwprofile_bypasslist_binding", "comment", "Testing"),
				),
			},
			{
				Config: testAccAppfwprofileBypasslistBinding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwprofileBypasslistBindingNotExist(
						"citrixadc_appfwprofile.tf_appfwprofile",
						"tf_appfwprofile_bypasslist",
						"http://www.example.com/bypass",
						"literal",
						"URL",
					),
				),
			},
		},
	})
}

func TestAccAppfwprofileBypasslistBinding_import(t *testing.T) {
	const resAddr = "citrixadc_appfwprofile_bypasslist_binding.tf_appfwprofile_bypasslist_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAppfwprofileBypasslistBindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAppfwprofileBypasslistBinding_basic_step1,
			},
			{
				Config:            testAccAppfwprofileBypasslistBinding_basic_step1,
				ResourceName:      resAddr,
				ImportState:       true,
				ImportStateVerify: true,
				// Full round-trip: the identity attrs (name, as_bypass_list,
				// as_bypass_list_value_type, as_bypass_list_location) are restored from
				// the parsed composite ID, and the echoed config attrs
				// (as_bypass_list_action, comment, state) plus the Computed
				// resourceid/alertonly/isautodeployed are repopulated from the GET row.
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccCheckAppfwprofileBypasslistBindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No appfwprofile_bypasslist_binding ID is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		idMap, _, err := utils.ParseIdString(rs.Primary.ID, nil, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", rs.Primary.ID, err)
		}

		name := idMap["name"]
		asBypassList := idMap["as_bypass_list"]
		asBypassListValueType := idMap["as_bypass_list_value_type"]
		asBypassListLocation := idMap["as_bypass_list_location"]

		findParams := service.FindParams{
			ResourceType:             service.Appfwprofile_bypasslist_binding.Type(),
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			return err
		}

		found := false
		for _, v := range dataArr {
			if val, ok := v["as_bypass_list"].(string); !ok || val != asBypassList {
				continue
			}
			if asBypassListValueType != "" {
				if val, ok := v["as_bypass_list_value_type"].(string); !ok || val != asBypassListValueType {
					continue
				}
			}
			if asBypassListLocation != "" {
				if val, ok := v["as_bypass_list_location"].(string); !ok || val != asBypassListLocation {
					continue
				}
			}
			found = true
			break
		}

		if !found {
			return fmt.Errorf("appfwprofile_bypasslist_binding %s not found", n)
		}

		return nil
	}
}

// testAccCheckAppfwprofileBypasslistBindingNotExist verifies, in step 2, that
// the binding was removed while the parent appfwprofile still exists.
func testAccCheckAppfwprofileBypasslistBindingNotExist(parentResource, name, asBypassList, asBypassListValueType, asBypassListLocation string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		_, ok := s.RootModule().Resources[parentResource]
		if !ok {
			return fmt.Errorf("Parent not found: %s", parentResource)
		}

		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		findParams := service.FindParams{
			ResourceType:             service.Appfwprofile_bypasslist_binding.Type(),
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			// No bindings on the parent at all is an acceptable "not found" outcome.
			return nil
		}

		for _, v := range dataArr {
			if val, ok := v["as_bypass_list"].(string); !ok || val != asBypassList {
				continue
			}
			if asBypassListValueType != "" {
				if val, ok := v["as_bypass_list_value_type"].(string); !ok || val != asBypassListValueType {
					continue
				}
			}
			if asBypassListLocation != "" {
				if val, ok := v["as_bypass_list_location"].(string); !ok || val != asBypassListLocation {
					continue
				}
			}
			return fmt.Errorf("appfwprofile_bypasslist_binding for %s/%s still exists", name, asBypassList)
		}

		return nil
	}
}

func testAccCheckAppfwprofileBypasslistBindingDestroy(s *terraform.State) error {
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_appfwprofile_bypasslist_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No appfwprofile_bypasslist_binding ID is set")
		}

		idMap, _, err := utils.ParseIdString(rs.Primary.ID, nil, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", rs.Primary.ID, err)
		}

		name := idMap["name"]
		asBypassList := idMap["as_bypass_list"]
		asBypassListValueType := idMap["as_bypass_list_value_type"]
		asBypassListLocation := idMap["as_bypass_list_location"]

		findParams := service.FindParams{
			ResourceType:             service.Appfwprofile_bypasslist_binding.Type(),
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			// Parent gone or no bindings - binding is destroyed.
			continue
		}

		for _, v := range dataArr {
			if val, ok := v["as_bypass_list"].(string); !ok || val != asBypassList {
				continue
			}
			if asBypassListValueType != "" {
				if val, ok := v["as_bypass_list_value_type"].(string); !ok || val != asBypassListValueType {
					continue
				}
			}
			if asBypassListLocation != "" {
				if val, ok := v["as_bypass_list_location"].(string); !ok || val != asBypassListLocation {
					continue
				}
			}
			return fmt.Errorf("appfwprofile_bypasslist_binding %s still exists", rs.Primary.ID)
		}
	}

	return nil
}

// Datasource requires name, as_bypass_list, as_bypass_list_value_type and
// as_bypass_list_location (all Required in the datasource schema), so the data
// block references all four from the resource.
const testAccAppfwprofileBypasslistBindingDataSource_basic = `

resource "citrixadc_appfwprofile" "tf_appfwprofile" {
  name = "tf_appfwprofile_bypasslist"
  type = ["HTML"]
}

resource "citrixadc_appfwprofile_bypasslist_binding" "tf_appfwprofile_bypasslist_binding" {
  name                      = citrixadc_appfwprofile.tf_appfwprofile.name
  as_bypass_list            = "http://www.example.com/bypass"
  as_bypass_list_value_type = "literal"
  as_bypass_list_location   = "URL"
  as_bypass_list_action     = "log"
  state                     = "ENABLED"
  comment                   = "Testing"

  depends_on = [citrixadc_appfwprofile.tf_appfwprofile]
}

data "citrixadc_appfwprofile_bypasslist_binding" "tf_appfwprofile_bypasslist_binding" {
  name                      = citrixadc_appfwprofile_bypasslist_binding.tf_appfwprofile_bypasslist_binding.name
  as_bypass_list            = citrixadc_appfwprofile_bypasslist_binding.tf_appfwprofile_bypasslist_binding.as_bypass_list
  as_bypass_list_value_type = citrixadc_appfwprofile_bypasslist_binding.tf_appfwprofile_bypasslist_binding.as_bypass_list_value_type
  as_bypass_list_location   = citrixadc_appfwprofile_bypasslist_binding.tf_appfwprofile_bypasslist_binding.as_bypass_list_location

  depends_on = [citrixadc_appfwprofile_bypasslist_binding.tf_appfwprofile_bypasslist_binding]
}
`

func TestAccAppfwprofileBypasslistBindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAppfwprofileBypasslistBindingDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_bypasslist_binding.tf_appfwprofile_bypasslist_binding", "name", "tf_appfwprofile_bypasslist"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_bypasslist_binding.tf_appfwprofile_bypasslist_binding", "as_bypass_list", "http://www.example.com/bypass"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_bypasslist_binding.tf_appfwprofile_bypasslist_binding", "as_bypass_list_value_type", "literal"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_bypasslist_binding.tf_appfwprofile_bypasslist_binding", "as_bypass_list_location", "URL"),
				),
			},
		},
	})
}
