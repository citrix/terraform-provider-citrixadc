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

// Step 1 creates the parent appfwprofile and binds a block keyword to it.
// All attributes on this binding are RequiresReplace (no in-place update); the
// binding is keyed by name + blockkeyword + fieldname + as_blockkeyword_formurl.
// The participating parent entity config (citrixadc_appfwprofile with name +
// type) is reused from appfwprofile_test.go.
const testAccAppfwprofileBlockkeywordBinding_basic_step1 = `

resource "citrixadc_appfwprofile" "tf_appfwprofile" {
  name = "tf_appfwprofile_blockkeyword"
  type = ["HTML"]
}

resource "citrixadc_appfwprofile_blockkeyword_binding" "tf_appfwprofile_blockkeyword_binding" {
  name                              = citrixadc_appfwprofile.tf_appfwprofile.name
  blockkeyword                      = "tf_blockkeyword"
  fieldname                         = "tf_fieldname"
  as_blockkeyword_formurl           = "http://www.example.com"
  as_fieldname_isregex_blockkeyword = "NOTREGEX"
  blockkeywordtype                  = "literal"
  state                             = "ENABLED"
  comment                           = "Testing"

  depends_on = [citrixadc_appfwprofile.tf_appfwprofile]
}
`

// Step 2 drops the binding to confirm it is deleted (only the parent remains).
const testAccAppfwprofileBlockkeywordBinding_basic_step2 = `

resource "citrixadc_appfwprofile" "tf_appfwprofile" {
  name = "tf_appfwprofile_blockkeyword"
  type = ["HTML"]
}
`

func TestAccAppfwprofileBlockkeywordBinding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAppfwprofileBlockkeywordBindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAppfwprofileBlockkeywordBinding_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwprofileBlockkeywordBindingExist("citrixadc_appfwprofile_blockkeyword_binding.tf_appfwprofile_blockkeyword_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_blockkeyword_binding.tf_appfwprofile_blockkeyword_binding", "name", "tf_appfwprofile_blockkeyword"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_blockkeyword_binding.tf_appfwprofile_blockkeyword_binding", "blockkeyword", "tf_blockkeyword"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_blockkeyword_binding.tf_appfwprofile_blockkeyword_binding", "fieldname", "tf_fieldname"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_blockkeyword_binding.tf_appfwprofile_blockkeyword_binding", "as_blockkeyword_formurl", "http://www.example.com"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_blockkeyword_binding.tf_appfwprofile_blockkeyword_binding", "as_fieldname_isregex_blockkeyword", "NOTREGEX"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_blockkeyword_binding.tf_appfwprofile_blockkeyword_binding", "blockkeywordtype", "literal"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_blockkeyword_binding.tf_appfwprofile_blockkeyword_binding", "state", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_blockkeyword_binding.tf_appfwprofile_blockkeyword_binding", "comment", "Testing"),
				),
			},
			{
				Config: testAccAppfwprofileBlockkeywordBinding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwprofileBlockkeywordBindingNotExist(
						"citrixadc_appfwprofile.tf_appfwprofile",
						"tf_appfwprofile_blockkeyword",
						"tf_blockkeyword",
						"tf_fieldname",
						"http://www.example.com",
					),
				),
			},
		},
	})
}

func TestAccAppfwprofileBlockkeywordBinding_import(t *testing.T) {
	const resAddr = "citrixadc_appfwprofile_blockkeyword_binding.tf_appfwprofile_blockkeyword_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAppfwprofileBlockkeywordBindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAppfwprofileBlockkeywordBinding_basic_step1,
			},
			{
				Config:            testAccAppfwprofileBlockkeywordBinding_basic_step1,
				ResourceName:      resAddr,
				ImportState:       true,
				ImportStateVerify: true,
				// Full round-trip: identity/ID-component attrs (name, blockkeyword,
				// fieldname, as_blockkeyword_formurl) are backfilled from the parsed ID
				// in the Read helper, and the echoed config attrs
				// (as_fieldname_isregex_blockkeyword, blockkeywordtype, state, comment)
				// are backfilled from the GET response in SetAttrFromGet on import. So
				// nothing needs to be ignored.
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccCheckAppfwprofileBlockkeywordBindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No appfwprofile_blockkeyword_binding ID is set")
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
		blockkeyword := idMap["blockkeyword"]
		fieldname := idMap["fieldname"]
		asBlockkeywordFormurl := idMap["as_blockkeyword_formurl"]

		findParams := service.FindParams{
			ResourceType:             service.Appfwprofile_blockkeyword_binding.Type(),
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			return err
		}

		found := false
		for _, v := range dataArr {
			if val, ok := v["blockkeyword"].(string); !ok || val != blockkeyword {
				continue
			}
			if fieldname != "" {
				if val, ok := v["fieldname"].(string); !ok || val != fieldname {
					continue
				}
			}
			if asBlockkeywordFormurl != "" {
				if val, ok := v["as_blockkeyword_formurl"].(string); !ok || val != asBlockkeywordFormurl {
					continue
				}
			}
			found = true
			break
		}

		if !found {
			return fmt.Errorf("appfwprofile_blockkeyword_binding %s not found", n)
		}

		return nil
	}
}

// testAccCheckAppfwprofileBlockkeywordBindingNotExist verifies, in step 2, that
// the binding was removed while the parent appfwprofile still exists.
func testAccCheckAppfwprofileBlockkeywordBindingNotExist(parentResource, name, blockkeyword, fieldname, asBlockkeywordFormurl string) resource.TestCheckFunc {
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
			ResourceType:             service.Appfwprofile_blockkeyword_binding.Type(),
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			// No bindings on the parent at all is an acceptable "not found" outcome.
			return nil
		}

		for _, v := range dataArr {
			if val, ok := v["blockkeyword"].(string); !ok || val != blockkeyword {
				continue
			}
			if fieldname != "" {
				if val, ok := v["fieldname"].(string); !ok || val != fieldname {
					continue
				}
			}
			if asBlockkeywordFormurl != "" {
				if val, ok := v["as_blockkeyword_formurl"].(string); !ok || val != asBlockkeywordFormurl {
					continue
				}
			}
			return fmt.Errorf("appfwprofile_blockkeyword_binding for %s/%s still exists", name, blockkeyword)
		}

		return nil
	}
}

func testAccCheckAppfwprofileBlockkeywordBindingDestroy(s *terraform.State) error {
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_appfwprofile_blockkeyword_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No appfwprofile_blockkeyword_binding ID is set")
		}

		idMap, _, err := utils.ParseIdString(rs.Primary.ID, nil, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", rs.Primary.ID, err)
		}

		name := idMap["name"]
		blockkeyword := idMap["blockkeyword"]
		fieldname := idMap["fieldname"]
		asBlockkeywordFormurl := idMap["as_blockkeyword_formurl"]

		findParams := service.FindParams{
			ResourceType:             service.Appfwprofile_blockkeyword_binding.Type(),
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			// Parent gone or no bindings - binding is destroyed.
			continue
		}

		for _, v := range dataArr {
			if val, ok := v["blockkeyword"].(string); !ok || val != blockkeyword {
				continue
			}
			if fieldname != "" {
				if val, ok := v["fieldname"].(string); !ok || val != fieldname {
					continue
				}
			}
			if asBlockkeywordFormurl != "" {
				if val, ok := v["as_blockkeyword_formurl"].(string); !ok || val != asBlockkeywordFormurl {
					continue
				}
			}
			return fmt.Errorf("appfwprofile_blockkeyword_binding %s still exists", rs.Primary.ID)
		}
	}

	return nil
}

// Datasource requires name, blockkeyword, fieldname and as_blockkeyword_formurl
// (all Required in the datasource schema), so the data block references all four
// from the resource.
const testAccAppfwprofileBlockkeywordBindingDataSource_basic = `

resource "citrixadc_appfwprofile" "tf_appfwprofile" {
  name = "tf_appfwprofile_blockkeyword"
  type = ["HTML"]
}

resource "citrixadc_appfwprofile_blockkeyword_binding" "tf_appfwprofile_blockkeyword_binding" {
  name                              = citrixadc_appfwprofile.tf_appfwprofile.name
  blockkeyword                      = "tf_blockkeyword"
  fieldname                         = "tf_fieldname"
  as_blockkeyword_formurl           = "http://www.example.com"
  as_fieldname_isregex_blockkeyword = "NOTREGEX"
  blockkeywordtype                  = "literal"
  state                             = "ENABLED"
  comment                           = "Testing"

  depends_on = [citrixadc_appfwprofile.tf_appfwprofile]
}

data "citrixadc_appfwprofile_blockkeyword_binding" "tf_appfwprofile_blockkeyword_binding" {
  name                    = citrixadc_appfwprofile_blockkeyword_binding.tf_appfwprofile_blockkeyword_binding.name
  blockkeyword            = citrixadc_appfwprofile_blockkeyword_binding.tf_appfwprofile_blockkeyword_binding.blockkeyword
  fieldname               = citrixadc_appfwprofile_blockkeyword_binding.tf_appfwprofile_blockkeyword_binding.fieldname
  as_blockkeyword_formurl = citrixadc_appfwprofile_blockkeyword_binding.tf_appfwprofile_blockkeyword_binding.as_blockkeyword_formurl

  depends_on = [citrixadc_appfwprofile_blockkeyword_binding.tf_appfwprofile_blockkeyword_binding]
}
`

func TestAccAppfwprofileBlockkeywordBindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAppfwprofileBlockkeywordBindingDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_blockkeyword_binding.tf_appfwprofile_blockkeyword_binding", "name", "tf_appfwprofile_blockkeyword"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_blockkeyword_binding.tf_appfwprofile_blockkeyword_binding", "blockkeyword", "tf_blockkeyword"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_blockkeyword_binding.tf_appfwprofile_blockkeyword_binding", "fieldname", "tf_fieldname"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_blockkeyword_binding.tf_appfwprofile_blockkeyword_binding", "as_blockkeyword_formurl", "http://www.example.com"),
				),
			},
		},
	})
}
