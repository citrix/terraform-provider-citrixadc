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

// Step 1 creates the parent appfwprofile and binds a deny-list entry to it.
// All attributes on this binding are RequiresReplace (no in-place update); the
// binding is keyed by name + as_deny_list + as_deny_list_value_type +
// as_deny_list_location. The participating parent entity config
// (citrixadc_appfwprofile with name + type) is reused from appfwprofile_test.go
// (and matches appfwprofile_bypasslist_binding_test.go).
// Note: as_deny_list_action is a LIST attribute, so it is set as a list in HCL.
const testAccAppfwprofileDenylistBinding_basic_step1 = `

resource "citrixadc_appfwprofile" "tf_appfwprofile" {
  name = "tf_appfwprofile_denylist"
  type = ["HTML"]
}

resource "citrixadc_appfwprofile_denylist_binding" "tf_appfwprofile_denylist_binding" {
  name                    = citrixadc_appfwprofile.tf_appfwprofile.name
  as_deny_list            = "http://www.example.com/deny"
  as_deny_list_value_type = "literal"
  as_deny_list_location   = "URL"
  as_deny_list_action     = ["log"]
  state                   = "ENABLED"
  comment                 = "Testing"

  depends_on = [citrixadc_appfwprofile.tf_appfwprofile]
}
`

// Step 2 drops the binding to confirm it is deleted (only the parent remains).
const testAccAppfwprofileDenylistBinding_basic_step2 = `

resource "citrixadc_appfwprofile" "tf_appfwprofile" {
  name = "tf_appfwprofile_denylist"
  type = ["HTML"]
}
`

func TestAccAppfwprofileDenylistBinding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAppfwprofileDenylistBindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAppfwprofileDenylistBinding_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwprofileDenylistBindingExist("citrixadc_appfwprofile_denylist_binding.tf_appfwprofile_denylist_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_denylist_binding.tf_appfwprofile_denylist_binding", "name", "tf_appfwprofile_denylist"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_denylist_binding.tf_appfwprofile_denylist_binding", "as_deny_list", "http://www.example.com/deny"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_denylist_binding.tf_appfwprofile_denylist_binding", "as_deny_list_value_type", "literal"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_denylist_binding.tf_appfwprofile_denylist_binding", "as_deny_list_location", "URL"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_denylist_binding.tf_appfwprofile_denylist_binding", "as_deny_list_action.#", "1"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_denylist_binding.tf_appfwprofile_denylist_binding", "as_deny_list_action.0", "log"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_denylist_binding.tf_appfwprofile_denylist_binding", "state", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_denylist_binding.tf_appfwprofile_denylist_binding", "comment", "Testing"),
				),
			},
			{
				Config: testAccAppfwprofileDenylistBinding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwprofileDenylistBindingNotExist(
						"citrixadc_appfwprofile.tf_appfwprofile",
						"tf_appfwprofile_denylist",
						"http://www.example.com/deny",
						"literal",
						"URL",
					),
				),
			},
		},
	})
}

func TestAccAppfwprofileDenylistBinding_import(t *testing.T) {
	const resAddr = "citrixadc_appfwprofile_denylist_binding.tf_appfwprofile_denylist_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAppfwprofileDenylistBindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAppfwprofileDenylistBinding_basic_step1,
			},
			{
				Config:            testAccAppfwprofileDenylistBinding_basic_step1,
				ResourceName:      resAddr,
				ImportState:       true,
				ImportStateVerify: true,
				// Full round-trip: the identity attributes (name, as_deny_list,
				// as_deny_list_value_type, as_deny_list_location) are backfilled from
				// the parsed composite ID, and the remaining user-supplied attributes
				// (as_deny_list_action, comment, state) are echoed verbatim by the GET
				// response and read back in appfwprofile_denylist_bindingSetAttrFromGet.
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccCheckAppfwprofileDenylistBindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No appfwprofile_denylist_binding ID is set")
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
		asDenyList := idMap["as_deny_list"]
		asDenyListValueType := idMap["as_deny_list_value_type"]
		asDenyListLocation := idMap["as_deny_list_location"]

		findParams := service.FindParams{
			ResourceType:             service.Appfwprofile_denylist_binding.Type(),
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			return err
		}

		found := false
		for _, v := range dataArr {
			if val, ok := v["as_deny_list"].(string); !ok || val != asDenyList {
				continue
			}
			if asDenyListValueType != "" {
				if val, ok := v["as_deny_list_value_type"].(string); !ok || val != asDenyListValueType {
					continue
				}
			}
			if asDenyListLocation != "" {
				if val, ok := v["as_deny_list_location"].(string); !ok || val != asDenyListLocation {
					continue
				}
			}
			found = true
			break
		}

		if !found {
			return fmt.Errorf("appfwprofile_denylist_binding %s not found", n)
		}

		return nil
	}
}

// testAccCheckAppfwprofileDenylistBindingNotExist verifies, in step 2, that the
// binding was removed while the parent appfwprofile still exists.
func testAccCheckAppfwprofileDenylistBindingNotExist(parentResource, name, asDenyList, asDenyListValueType, asDenyListLocation string) resource.TestCheckFunc {
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
			ResourceType:             service.Appfwprofile_denylist_binding.Type(),
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			// No bindings on the parent at all is an acceptable "not found" outcome.
			return nil
		}

		for _, v := range dataArr {
			if val, ok := v["as_deny_list"].(string); !ok || val != asDenyList {
				continue
			}
			if asDenyListValueType != "" {
				if val, ok := v["as_deny_list_value_type"].(string); !ok || val != asDenyListValueType {
					continue
				}
			}
			if asDenyListLocation != "" {
				if val, ok := v["as_deny_list_location"].(string); !ok || val != asDenyListLocation {
					continue
				}
			}
			return fmt.Errorf("appfwprofile_denylist_binding for %s/%s still exists", name, asDenyList)
		}

		return nil
	}
}

func testAccCheckAppfwprofileDenylistBindingDestroy(s *terraform.State) error {
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_appfwprofile_denylist_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No appfwprofile_denylist_binding ID is set")
		}

		idMap, _, err := utils.ParseIdString(rs.Primary.ID, nil, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", rs.Primary.ID, err)
		}

		name := idMap["name"]
		asDenyList := idMap["as_deny_list"]
		asDenyListValueType := idMap["as_deny_list_value_type"]
		asDenyListLocation := idMap["as_deny_list_location"]

		findParams := service.FindParams{
			ResourceType:             service.Appfwprofile_denylist_binding.Type(),
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			// Parent gone or no bindings - binding is destroyed.
			continue
		}

		for _, v := range dataArr {
			if val, ok := v["as_deny_list"].(string); !ok || val != asDenyList {
				continue
			}
			if asDenyListValueType != "" {
				if val, ok := v["as_deny_list_value_type"].(string); !ok || val != asDenyListValueType {
					continue
				}
			}
			if asDenyListLocation != "" {
				if val, ok := v["as_deny_list_location"].(string); !ok || val != asDenyListLocation {
					continue
				}
			}
			return fmt.Errorf("appfwprofile_denylist_binding %s still exists", rs.Primary.ID)
		}
	}

	return nil
}

// Datasource requires name, as_deny_list, as_deny_list_value_type and
// as_deny_list_location (all Required in the datasource schema), so the data
// block references all four from the resource.
const testAccAppfwprofileDenylistBindingDataSource_basic = `

resource "citrixadc_appfwprofile" "tf_appfwprofile" {
  name = "tf_appfwprofile_denylist"
  type = ["HTML"]
}

resource "citrixadc_appfwprofile_denylist_binding" "tf_appfwprofile_denylist_binding" {
  name                    = citrixadc_appfwprofile.tf_appfwprofile.name
  as_deny_list            = "http://www.example.com/deny"
  as_deny_list_value_type = "literal"
  as_deny_list_location   = "URL"
  as_deny_list_action     = ["log"]
  state                   = "ENABLED"
  comment                 = "Testing"

  depends_on = [citrixadc_appfwprofile.tf_appfwprofile]
}

data "citrixadc_appfwprofile_denylist_binding" "tf_appfwprofile_denylist_binding" {
  name                    = citrixadc_appfwprofile_denylist_binding.tf_appfwprofile_denylist_binding.name
  as_deny_list            = citrixadc_appfwprofile_denylist_binding.tf_appfwprofile_denylist_binding.as_deny_list
  as_deny_list_value_type = citrixadc_appfwprofile_denylist_binding.tf_appfwprofile_denylist_binding.as_deny_list_value_type
  as_deny_list_location   = citrixadc_appfwprofile_denylist_binding.tf_appfwprofile_denylist_binding.as_deny_list_location

  depends_on = [citrixadc_appfwprofile_denylist_binding.tf_appfwprofile_denylist_binding]
}
`

func TestAccAppfwprofileDenylistBindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAppfwprofileDenylistBindingDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_denylist_binding.tf_appfwprofile_denylist_binding", "name", "tf_appfwprofile_denylist"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_denylist_binding.tf_appfwprofile_denylist_binding", "as_deny_list", "http://www.example.com/deny"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_denylist_binding.tf_appfwprofile_denylist_binding", "as_deny_list_value_type", "literal"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_denylist_binding.tf_appfwprofile_denylist_binding", "as_deny_list_location", "URL"),
				),
			},
		},
	})
}
