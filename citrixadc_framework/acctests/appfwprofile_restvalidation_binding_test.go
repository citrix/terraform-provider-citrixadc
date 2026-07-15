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

// Step 1 creates the parent appfwprofile and binds a REST validation relaxation
// rule to it. All attributes on this binding are RequiresReplace (no in-place
// update); the binding is keyed by name + rest_validation_action + restvalidation
// (the composite ID order is taken from Create / SetAttrFromGetForDatasource in
// resource_schema.go). rest_validation_action is part of the composite ID / delete
// key so it is always set. alertonly, resourceid and isautodeployed are read-only /
// server-assigned/derived and are intentionally NOT set or asserted. The
// participating parent entity config (citrixadc_appfwprofile with name + type) is
// reused from appfwprofile_test.go / appfwprofile_grpcvalidation_binding_test.go.
const testAccAppfwprofileRestvalidationBinding_basic_step1 = `

resource "citrixadc_appfwprofile" "tf_appfwprofile" {
  name = "tf_appfwprofile_restvalidation"
  type = ["HTML", "JSON"]
}

resource "citrixadc_appfwprofile_restvalidation_binding" "tf_appfwprofile_restvalidation_binding" {
  name                   = citrixadc_appfwprofile.tf_appfwprofile.name
  restvalidation         = "GET:/v1/bookstore/viewbooks"
  rest_validation_action = "log"
  state                  = "ENABLED"
  comment                = "Testing"

  depends_on = [citrixadc_appfwprofile.tf_appfwprofile]
}
`

// Step 2 drops the binding to confirm it is deleted (only the parent remains).
const testAccAppfwprofileRestvalidationBinding_basic_step2 = `

resource "citrixadc_appfwprofile" "tf_appfwprofile" {
  name = "tf_appfwprofile_restvalidation"
  type = ["HTML", "JSON"]
}
`

func TestAccAppfwprofileRestvalidationBinding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAppfwprofileRestvalidationBindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAppfwprofileRestvalidationBinding_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwprofileRestvalidationBindingExist("citrixadc_appfwprofile_restvalidation_binding.tf_appfwprofile_restvalidation_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_restvalidation_binding.tf_appfwprofile_restvalidation_binding", "name", "tf_appfwprofile_restvalidation"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_restvalidation_binding.tf_appfwprofile_restvalidation_binding", "restvalidation", "GET:/v1/bookstore/viewbooks"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_restvalidation_binding.tf_appfwprofile_restvalidation_binding", "rest_validation_action", "log"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_restvalidation_binding.tf_appfwprofile_restvalidation_binding", "state", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_restvalidation_binding.tf_appfwprofile_restvalidation_binding", "comment", "Testing"),
				),
			},
			{
				Config: testAccAppfwprofileRestvalidationBinding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwprofileRestvalidationBindingNotExist(
						"citrixadc_appfwprofile.tf_appfwprofile",
						"tf_appfwprofile_restvalidation",
						"GET:/v1/bookstore/viewbooks",
						"log",
					),
				),
			},
		},
	})
}

func testAccCheckAppfwprofileRestvalidationBindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No appfwprofile_restvalidation_binding ID is set")
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
		restvalidation := idMap["restvalidation"]
		restValidationAction := idMap["rest_validation_action"]

		findParams := service.FindParams{
			ResourceType:             service.Appfwprofile_restvalidation_binding.Type(),
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			return err
		}

		found := false
		for _, v := range dataArr {
			if val, ok := v["restvalidation"].(string); !ok || val != restvalidation {
				continue
			}
			if restValidationAction != "" {
				if val, ok := v["rest_validation_action"].(string); !ok || val != restValidationAction {
					continue
				}
			}
			found = true
			break
		}

		if !found {
			return fmt.Errorf("appfwprofile_restvalidation_binding %s not found", n)
		}

		return nil
	}
}

// testAccCheckAppfwprofileRestvalidationBindingNotExist verifies, in step 2, that
// the binding was removed while the parent appfwprofile still exists.
func testAccCheckAppfwprofileRestvalidationBindingNotExist(parentResource, name, restvalidation, restValidationAction string) resource.TestCheckFunc {
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
			ResourceType:             service.Appfwprofile_restvalidation_binding.Type(),
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			// No bindings on the parent at all is an acceptable "not found" outcome.
			return nil
		}

		for _, v := range dataArr {
			if val, ok := v["restvalidation"].(string); !ok || val != restvalidation {
				continue
			}
			if restValidationAction != "" {
				if val, ok := v["rest_validation_action"].(string); !ok || val != restValidationAction {
					continue
				}
			}
			return fmt.Errorf("appfwprofile_restvalidation_binding for %s/%s still exists", name, restvalidation)
		}

		return nil
	}
}

func testAccCheckAppfwprofileRestvalidationBindingDestroy(s *terraform.State) error {
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_appfwprofile_restvalidation_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No appfwprofile_restvalidation_binding ID is set")
		}

		idMap, _, err := utils.ParseIdString(rs.Primary.ID, nil, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", rs.Primary.ID, err)
		}

		name := idMap["name"]
		restvalidation := idMap["restvalidation"]
		restValidationAction := idMap["rest_validation_action"]

		findParams := service.FindParams{
			ResourceType:             service.Appfwprofile_restvalidation_binding.Type(),
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			// Parent gone or no bindings - binding is destroyed.
			continue
		}

		for _, v := range dataArr {
			if val, ok := v["restvalidation"].(string); !ok || val != restvalidation {
				continue
			}
			if restValidationAction != "" {
				if val, ok := v["rest_validation_action"].(string); !ok || val != restValidationAction {
					continue
				}
			}
			return fmt.Errorf("appfwprofile_restvalidation_binding %s still exists", rs.Primary.ID)
		}
	}

	return nil
}

// Datasource requires name, restvalidation and rest_validation_action (all Required
// in the datasource schema), so the data block references all three from the
// resource.
const testAccAppfwprofileRestvalidationBindingDataSource_basic = `

resource "citrixadc_appfwprofile" "tf_appfwprofile" {
  name = "tf_appfwprofile_restvalidation"
  type = ["HTML", "JSON"]
}

resource "citrixadc_appfwprofile_restvalidation_binding" "tf_appfwprofile_restvalidation_binding" {
  name                   = citrixadc_appfwprofile.tf_appfwprofile.name
  restvalidation         = "GET:/v1/bookstore/viewbooks"
  rest_validation_action = "log"
  state                  = "ENABLED"
  comment                = "Testing"

  depends_on = [citrixadc_appfwprofile.tf_appfwprofile]
}

data "citrixadc_appfwprofile_restvalidation_binding" "tf_appfwprofile_restvalidation_binding" {
  name                   = citrixadc_appfwprofile_restvalidation_binding.tf_appfwprofile_restvalidation_binding.name
  restvalidation         = citrixadc_appfwprofile_restvalidation_binding.tf_appfwprofile_restvalidation_binding.restvalidation
  rest_validation_action = citrixadc_appfwprofile_restvalidation_binding.tf_appfwprofile_restvalidation_binding.rest_validation_action

  depends_on = [citrixadc_appfwprofile_restvalidation_binding.tf_appfwprofile_restvalidation_binding]
}
`

func TestAccAppfwprofileRestvalidationBinding_import(t *testing.T) {
	const resAddr = "citrixadc_appfwprofile_restvalidation_binding.tf_appfwprofile_restvalidation_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAppfwprofileRestvalidationBindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAppfwprofileRestvalidationBinding_basic_step1,
			},
			{
				Config:            testAccAppfwprofileRestvalidationBinding_basic_step1,
				ResourceName:      resAddr,
				ImportState:       true,
				ImportStateVerify: true,
				// Full round-trip: name/restvalidation/rest_validation_action are
				// backfilled from the parsed composite ID, and comment/state are read
				// back from the GET row, so no attribute needs to be ignored on import.
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccAppfwprofileRestvalidationBindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAppfwprofileRestvalidationBindingDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_restvalidation_binding.tf_appfwprofile_restvalidation_binding", "name", "tf_appfwprofile_restvalidation"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_restvalidation_binding.tf_appfwprofile_restvalidation_binding", "restvalidation", "GET:/v1/bookstore/viewbooks"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_restvalidation_binding.tf_appfwprofile_restvalidation_binding", "rest_validation_action", "log"),
				),
			},
		},
	})
}
