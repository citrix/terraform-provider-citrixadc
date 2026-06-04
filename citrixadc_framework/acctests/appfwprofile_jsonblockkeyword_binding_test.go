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

// Step 1 creates the parent appfwprofile and binds a JSON block keyword to it.
// All attributes on this binding are RequiresReplace (no in-place update); the
// binding is keyed by name + jsonblockkeyword + keyname_json_blockkeyword +
// jsonblockkeywordurl. The participating parent entity config (citrixadc_appfwprofile
// with name + type) is reused from appfwprofile_test.go; type includes "JSON" so a
// JSON keyword bind is valid on the profile.
// Read-only/computed attributes (alertonly, resourceid) are never set or asserted.
const testAccAppfwprofileJsonblockkeywordBinding_basic_step1 = `

resource "citrixadc_appfwprofile" "tf_appfwprofile" {
  name = "tf_appfwprofile_jsonblockkeyword"
  type = ["HTML", "JSON"]
}

resource "citrixadc_appfwprofile_jsonblockkeyword_binding" "tf_appfwprofile_jsonblockkeyword_binding" {
  name                         = citrixadc_appfwprofile.tf_appfwprofile.name
  jsonblockkeyword             = "tf_jsonblockkeyword"
  keyname_json_blockkeyword    = "tf_keyname"
  jsonblockkeywordurl          = "http://www.example.com"
  iskeyregex_json_blockkeyword = "NOTREGEX"
  jsonblockkeywordtype         = "literal"
  state                        = "ENABLED"
  comment                      = "Testing"

  depends_on = [citrixadc_appfwprofile.tf_appfwprofile]
}
`

// Step 2 drops the binding to confirm it is deleted (only the parent remains).
const testAccAppfwprofileJsonblockkeywordBinding_basic_step2 = `

resource "citrixadc_appfwprofile" "tf_appfwprofile" {
  name = "tf_appfwprofile_jsonblockkeyword"
  type = ["HTML", "JSON"]
}
`

func TestAccAppfwprofileJsonblockkeywordBinding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAppfwprofileJsonblockkeywordBindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAppfwprofileJsonblockkeywordBinding_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwprofileJsonblockkeywordBindingExist("citrixadc_appfwprofile_jsonblockkeyword_binding.tf_appfwprofile_jsonblockkeyword_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_jsonblockkeyword_binding.tf_appfwprofile_jsonblockkeyword_binding", "name", "tf_appfwprofile_jsonblockkeyword"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_jsonblockkeyword_binding.tf_appfwprofile_jsonblockkeyword_binding", "jsonblockkeyword", "tf_jsonblockkeyword"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_jsonblockkeyword_binding.tf_appfwprofile_jsonblockkeyword_binding", "keyname_json_blockkeyword", "tf_keyname"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_jsonblockkeyword_binding.tf_appfwprofile_jsonblockkeyword_binding", "jsonblockkeywordurl", "http://www.example.com"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_jsonblockkeyword_binding.tf_appfwprofile_jsonblockkeyword_binding", "iskeyregex_json_blockkeyword", "NOTREGEX"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_jsonblockkeyword_binding.tf_appfwprofile_jsonblockkeyword_binding", "jsonblockkeywordtype", "literal"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_jsonblockkeyword_binding.tf_appfwprofile_jsonblockkeyword_binding", "state", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_jsonblockkeyword_binding.tf_appfwprofile_jsonblockkeyword_binding", "comment", "Testing"),
				),
			},
			{
				Config: testAccAppfwprofileJsonblockkeywordBinding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwprofileJsonblockkeywordBindingNotExist(
						"citrixadc_appfwprofile.tf_appfwprofile",
						"tf_appfwprofile_jsonblockkeyword",
						"tf_jsonblockkeyword",
						"tf_keyname",
						"http://www.example.com",
					),
				),
			},
		},
	})
}

func testAccCheckAppfwprofileJsonblockkeywordBindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No appfwprofile_jsonblockkeyword_binding ID is set")
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
		jsonblockkeyword := idMap["jsonblockkeyword"]
		keynameJsonBlockkeyword := idMap["keyname_json_blockkeyword"]
		jsonblockkeywordurl := idMap["jsonblockkeywordurl"]

		findParams := service.FindParams{
			ResourceType:             service.Appfwprofile_jsonblockkeyword_binding.Type(),
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			return err
		}

		found := false
		for _, v := range dataArr {
			if val, ok := v["jsonblockkeyword"].(string); !ok || val != jsonblockkeyword {
				continue
			}
			if keynameJsonBlockkeyword != "" {
				if val, ok := v["keyname_json_blockkeyword"].(string); !ok || val != keynameJsonBlockkeyword {
					continue
				}
			}
			if jsonblockkeywordurl != "" {
				if val, ok := v["jsonblockkeywordurl"].(string); !ok || val != jsonblockkeywordurl {
					continue
				}
			}
			found = true
			break
		}

		if !found {
			return fmt.Errorf("appfwprofile_jsonblockkeyword_binding %s not found", n)
		}

		return nil
	}
}

// testAccCheckAppfwprofileJsonblockkeywordBindingNotExist verifies, in step 2, that
// the binding was removed while the parent appfwprofile still exists.
func testAccCheckAppfwprofileJsonblockkeywordBindingNotExist(parentResource, name, jsonblockkeyword, keynameJsonBlockkeyword, jsonblockkeywordurl string) resource.TestCheckFunc {
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
			ResourceType:             service.Appfwprofile_jsonblockkeyword_binding.Type(),
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			// No bindings on the parent at all is an acceptable "not found" outcome.
			return nil
		}

		for _, v := range dataArr {
			if val, ok := v["jsonblockkeyword"].(string); !ok || val != jsonblockkeyword {
				continue
			}
			if keynameJsonBlockkeyword != "" {
				if val, ok := v["keyname_json_blockkeyword"].(string); !ok || val != keynameJsonBlockkeyword {
					continue
				}
			}
			if jsonblockkeywordurl != "" {
				if val, ok := v["jsonblockkeywordurl"].(string); !ok || val != jsonblockkeywordurl {
					continue
				}
			}
			return fmt.Errorf("appfwprofile_jsonblockkeyword_binding for %s/%s still exists", name, jsonblockkeyword)
		}

		return nil
	}
}

func testAccCheckAppfwprofileJsonblockkeywordBindingDestroy(s *terraform.State) error {
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_appfwprofile_jsonblockkeyword_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No appfwprofile_jsonblockkeyword_binding ID is set")
		}

		idMap, _, err := utils.ParseIdString(rs.Primary.ID, nil, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", rs.Primary.ID, err)
		}

		name := idMap["name"]
		jsonblockkeyword := idMap["jsonblockkeyword"]
		keynameJsonBlockkeyword := idMap["keyname_json_blockkeyword"]
		jsonblockkeywordurl := idMap["jsonblockkeywordurl"]

		findParams := service.FindParams{
			ResourceType:             service.Appfwprofile_jsonblockkeyword_binding.Type(),
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			// Parent gone or no bindings - binding is destroyed.
			continue
		}

		for _, v := range dataArr {
			if val, ok := v["jsonblockkeyword"].(string); !ok || val != jsonblockkeyword {
				continue
			}
			if keynameJsonBlockkeyword != "" {
				if val, ok := v["keyname_json_blockkeyword"].(string); !ok || val != keynameJsonBlockkeyword {
					continue
				}
			}
			if jsonblockkeywordurl != "" {
				if val, ok := v["jsonblockkeywordurl"].(string); !ok || val != jsonblockkeywordurl {
					continue
				}
			}
			return fmt.Errorf("appfwprofile_jsonblockkeyword_binding %s still exists", rs.Primary.ID)
		}
	}

	return nil
}

// Datasource requires name, jsonblockkeyword, keyname_json_blockkeyword and
// jsonblockkeywordurl (all Required in the datasource schema), so the data block
// references all four from the resource.
const testAccAppfwprofileJsonblockkeywordBindingDataSource_basic = `

resource "citrixadc_appfwprofile" "tf_appfwprofile" {
  name = "tf_appfwprofile_jsonblockkeyword"
  type = ["HTML", "JSON"]
}

resource "citrixadc_appfwprofile_jsonblockkeyword_binding" "tf_appfwprofile_jsonblockkeyword_binding" {
  name                         = citrixadc_appfwprofile.tf_appfwprofile.name
  jsonblockkeyword             = "tf_jsonblockkeyword"
  keyname_json_blockkeyword    = "tf_keyname"
  jsonblockkeywordurl          = "http://www.example.com"
  iskeyregex_json_blockkeyword = "NOTREGEX"
  jsonblockkeywordtype         = "literal"
  state                        = "ENABLED"
  comment                      = "Testing"

  depends_on = [citrixadc_appfwprofile.tf_appfwprofile]
}

data "citrixadc_appfwprofile_jsonblockkeyword_binding" "tf_appfwprofile_jsonblockkeyword_binding" {
  name                      = citrixadc_appfwprofile_jsonblockkeyword_binding.tf_appfwprofile_jsonblockkeyword_binding.name
  jsonblockkeyword          = citrixadc_appfwprofile_jsonblockkeyword_binding.tf_appfwprofile_jsonblockkeyword_binding.jsonblockkeyword
  keyname_json_blockkeyword = citrixadc_appfwprofile_jsonblockkeyword_binding.tf_appfwprofile_jsonblockkeyword_binding.keyname_json_blockkeyword
  jsonblockkeywordurl       = citrixadc_appfwprofile_jsonblockkeyword_binding.tf_appfwprofile_jsonblockkeyword_binding.jsonblockkeywordurl

  depends_on = [citrixadc_appfwprofile_jsonblockkeyword_binding.tf_appfwprofile_jsonblockkeyword_binding]
}
`

func TestAccAppfwprofileJsonblockkeywordBindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAppfwprofileJsonblockkeywordBindingDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_jsonblockkeyword_binding.tf_appfwprofile_jsonblockkeyword_binding", "name", "tf_appfwprofile_jsonblockkeyword"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_jsonblockkeyword_binding.tf_appfwprofile_jsonblockkeyword_binding", "jsonblockkeyword", "tf_jsonblockkeyword"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_jsonblockkeyword_binding.tf_appfwprofile_jsonblockkeyword_binding", "keyname_json_blockkeyword", "tf_keyname"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_jsonblockkeyword_binding.tf_appfwprofile_jsonblockkeyword_binding", "jsonblockkeywordurl", "http://www.example.com"),
				),
			},
		},
	})
}
