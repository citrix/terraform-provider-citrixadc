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

// Step 1 creates the parent appfwprofile and binds a confidential field to it.
// All attributes on this binding are RequiresReplace (no in-place update); the
// binding is keyed by name + confidfield + cffield_url.
const testAccAppfwprofileAppfwconfidfieldBinding_basic_step1 = `

resource "citrixadc_appfwprofile" "tf_appfwprofile" {
  name = "tf_appfwprofile_confidfield"
  type = ["HTML"]
}

resource "citrixadc_appfwprofile_appfwconfidfield_binding" "tf_appfwprofile_appfwconfidfield_binding" {
  name            = citrixadc_appfwprofile.tf_appfwprofile.name
  confidfield     = "tf_confidfield"
  cffield_url     = "http://www.example.com"
  isregex_cffield = "NOTREGEX"
  state           = "ENABLED"
  comment         = "Testing"

  depends_on = [citrixadc_appfwprofile.tf_appfwprofile]
}
`

// Step 2 drops the binding to confirm it is deleted (only the parent remains).
const testAccAppfwprofileAppfwconfidfieldBinding_basic_step2 = `

resource "citrixadc_appfwprofile" "tf_appfwprofile" {
  name = "tf_appfwprofile_confidfield"
  type = ["HTML"]
}
`

func TestAccAppfwprofileAppfwconfidfieldBinding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAppfwprofileAppfwconfidfieldBindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAppfwprofileAppfwconfidfieldBinding_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwprofileAppfwconfidfieldBindingExist("citrixadc_appfwprofile_appfwconfidfield_binding.tf_appfwprofile_appfwconfidfield_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_appfwconfidfield_binding.tf_appfwprofile_appfwconfidfield_binding", "name", "tf_appfwprofile_confidfield"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_appfwconfidfield_binding.tf_appfwprofile_appfwconfidfield_binding", "confidfield", "tf_confidfield"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_appfwconfidfield_binding.tf_appfwprofile_appfwconfidfield_binding", "cffield_url", "http://www.example.com"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_appfwconfidfield_binding.tf_appfwprofile_appfwconfidfield_binding", "isregex_cffield", "NOTREGEX"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_appfwconfidfield_binding.tf_appfwprofile_appfwconfidfield_binding", "state", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_appfwconfidfield_binding.tf_appfwprofile_appfwconfidfield_binding", "comment", "Testing"),
				),
			},
			{
				Config: testAccAppfwprofileAppfwconfidfieldBinding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwprofileAppfwconfidfieldBindingNotExist(
						"citrixadc_appfwprofile.tf_appfwprofile",
						"tf_appfwprofile_confidfield",
						"tf_confidfield",
						"http://www.example.com",
					),
				),
			},
		},
	})
}

func testAccCheckAppfwprofileAppfwconfidfieldBindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No appfwprofile_appfwconfidfield_binding ID is set")
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
		confidfield := idMap["confidfield"]
		cffieldUrl := idMap["cffield_url"]

		findParams := service.FindParams{
			ResourceType:             service.Appfwprofile_appfwconfidfield_binding.Type(),
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			return err
		}

		found := false
		for _, v := range dataArr {
			if val, ok := v["confidfield"].(string); !ok || val != confidfield {
				continue
			}
			if cffieldUrl != "" {
				if val, ok := v["cffield_url"].(string); !ok || val != cffieldUrl {
					continue
				}
			}
			found = true
			break
		}

		if !found {
			return fmt.Errorf("appfwprofile_appfwconfidfield_binding %s not found", n)
		}

		return nil
	}
}

// testAccCheckAppfwprofileAppfwconfidfieldBindingNotExist verifies, in step 2,
// that the binding was removed while the parent appfwprofile still exists.
func testAccCheckAppfwprofileAppfwconfidfieldBindingNotExist(parentResource, name, confidfield, cffieldUrl string) resource.TestCheckFunc {
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
			ResourceType:             service.Appfwprofile_appfwconfidfield_binding.Type(),
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			// No bindings on the parent at all is an acceptable "not found" outcome.
			return nil
		}

		for _, v := range dataArr {
			if val, ok := v["confidfield"].(string); !ok || val != confidfield {
				continue
			}
			if cffieldUrl != "" {
				if val, ok := v["cffield_url"].(string); !ok || val != cffieldUrl {
					continue
				}
			}
			return fmt.Errorf("appfwprofile_appfwconfidfield_binding for %s/%s still exists", name, confidfield)
		}

		return nil
	}
}

func testAccCheckAppfwprofileAppfwconfidfieldBindingDestroy(s *terraform.State) error {
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_appfwprofile_appfwconfidfield_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No appfwprofile_appfwconfidfield_binding ID is set")
		}

		idMap, _, err := utils.ParseIdString(rs.Primary.ID, nil, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", rs.Primary.ID, err)
		}

		name := idMap["name"]
		confidfield := idMap["confidfield"]
		cffieldUrl := idMap["cffield_url"]

		findParams := service.FindParams{
			ResourceType:             service.Appfwprofile_appfwconfidfield_binding.Type(),
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			// Parent gone or no bindings - binding is destroyed.
			continue
		}

		for _, v := range dataArr {
			if val, ok := v["confidfield"].(string); !ok || val != confidfield {
				continue
			}
			if cffieldUrl != "" {
				if val, ok := v["cffield_url"].(string); !ok || val != cffieldUrl {
					continue
				}
			}
			return fmt.Errorf("appfwprofile_appfwconfidfield_binding %s still exists", rs.Primary.ID)
		}
	}

	return nil
}

// Datasource requires name, confidfield and cffield_url (all Required in the
// datasource schema), so the data block references all three from the resource.
const testAccAppfwprofileAppfwconfidfieldBindingDataSource_basic = `

resource "citrixadc_appfwprofile" "tf_appfwprofile" {
  name = "tf_appfwprofile_confidfield"
  type = ["HTML"]
}

resource "citrixadc_appfwprofile_appfwconfidfield_binding" "tf_appfwprofile_appfwconfidfield_binding" {
  name            = citrixadc_appfwprofile.tf_appfwprofile.name
  confidfield     = "tf_confidfield"
  cffield_url     = "http://www.example.com"
  isregex_cffield = "NOTREGEX"
  state           = "ENABLED"
  comment         = "Testing"

  depends_on = [citrixadc_appfwprofile.tf_appfwprofile]
}

data "citrixadc_appfwprofile_appfwconfidfield_binding" "tf_appfwprofile_appfwconfidfield_binding" {
  name        = citrixadc_appfwprofile_appfwconfidfield_binding.tf_appfwprofile_appfwconfidfield_binding.name
  confidfield = citrixadc_appfwprofile_appfwconfidfield_binding.tf_appfwprofile_appfwconfidfield_binding.confidfield
  cffield_url = citrixadc_appfwprofile_appfwconfidfield_binding.tf_appfwprofile_appfwconfidfield_binding.cffield_url

  depends_on = [citrixadc_appfwprofile_appfwconfidfield_binding.tf_appfwprofile_appfwconfidfield_binding]
}
`

func TestAccAppfwprofileAppfwconfidfieldBindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAppfwprofileAppfwconfidfieldBindingDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_appfwconfidfield_binding.tf_appfwprofile_appfwconfidfield_binding", "name", "tf_appfwprofile_confidfield"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_appfwconfidfield_binding.tf_appfwprofile_appfwconfidfield_binding", "confidfield", "tf_confidfield"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_appfwconfidfield_binding.tf_appfwprofile_appfwconfidfield_binding", "cffield_url", "http://www.example.com"),
				),
			},
		},
	})
}
