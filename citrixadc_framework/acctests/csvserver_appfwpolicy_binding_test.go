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

const testAccCsvserver_appfwpolicy_binding_basic = `
	resource citrixadc_csvserver_appfwpolicy_binding demo_binding {
		name = citrixadc_csvserver.demo_cs.name
		priority = 100
		policyname  = citrixadc_appfwpolicy.demo_appfwpolicy.name
		gotopriorityexpression = "END"
	}
	resource "citrixadc_csvserver" "demo_cs" {
		ipv46       = "10.10.10.33"
		name        = "demo_csvserver"
		port        = 80
		servicetype = "HTTP"
	}

	resource citrixadc_appfwprofile demo_appfwprofile {
		name = "demo_appfwprofile"
		bufferoverflowaction = ["none"]
		contenttypeaction = ["none"]
		cookieconsistencyaction = ["none"]
		creditcard = ["none"]
		creditcardaction = ["none"]
		crosssitescriptingaction = ["none"]
		csrftagaction = ["none"]
		denyurlaction = ["none"]
		dynamiclearning = ["none"]
		fieldconsistencyaction = ["none"]
		fieldformataction = ["none"]
		fileuploadtypesaction = ["none"]
		inspectcontenttypes = ["none"]
		jsondosaction = ["none"]
		jsonsqlinjectionaction = ["none"]
		jsonxssaction = ["none"]
		multipleheaderaction = ["none"]
		sqlinjectionaction = ["none"]
		starturlaction = ["none"]
		type = ["HTML"]
		xmlattachmentaction = ["none"]
		xmldosaction = ["none"]
		xmlformataction = ["none"]
		xmlsoapfaultaction = ["none"]
		xmlsqlinjectionaction = ["none"]
		xmlvalidationaction = ["none"]
		xmlwsiaction = ["none"]
		xmlxssaction = ["none"]
	}

	resource citrixadc_appfwpolicy demo_appfwpolicy {
		name = "demo_appfwpolicy"
		profilename = citrixadc_appfwprofile.demo_appfwprofile.name
		rule = "true"
	}
`

func TestAccCsvserver_appfwpolicy_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCsvserver_appfwpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCsvserver_appfwpolicy_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCsvserver_appfwpolicy_bindingExist("citrixadc_csvserver_appfwpolicy_binding.demo_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_csvserver_appfwpolicy_binding.demo_binding", "name", "demo_csvserver"),
					resource.TestCheckResourceAttr("citrixadc_csvserver_appfwpolicy_binding.demo_binding", "priority", "100"),
					resource.TestCheckResourceAttr("citrixadc_csvserver_appfwpolicy_binding.demo_binding", "policyname", "demo_appfwpolicy"),
					resource.TestCheckResourceAttr("citrixadc_csvserver_appfwpolicy_binding.demo_binding", "gotopriorityexpression", "END"),
				),
			},
		},
	})
}

func testAccCheckCsvserver_appfwpolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No csvserver_appfwpolicy_binding name is set")
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
		idMap, _, err := utils.ParseIdString(bindingId, []string{"name", "policyname"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", bindingId, err)
		}
		csvserverName := idMap["name"]
		appfwPolicyName := idMap["policyname"]

		findParams := service.FindParams{
			ResourceType:             service.Csvserver_appfwpolicy_binding.Type(),
			ResourceName:             csvserverName,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the right policy name
		foundIndex := -1
		for i, v := range dataArr {
			if v["policyname"].(string) == appfwPolicyName {
				foundIndex = i
				break
			}
		}

		// Resource is missing
		if foundIndex == -1 {
			return fmt.Errorf("Cannot find csvserver_appfwpolicy_binding ID %v", bindingId)
		}

		return nil
	}
}

func testAccCheckCsvserver_appfwpolicy_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_csvserver_appfwpolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Csvserver_appfwpolicy_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("csvserver_appfwpolicy_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccCsvserver_appfwpolicy_bindingDataSource_basic = `
	resource citrixadc_csvserver_appfwpolicy_binding demo_binding {
		name = citrixadc_csvserver.demo_cs.name
		priority = 100
		policyname  = citrixadc_appfwpolicy.demo_appfwpolicy.name
		gotopriorityexpression = "END"
	}
	resource "citrixadc_csvserver" "demo_cs" {
		ipv46       = "10.10.10.33"
		name        = "demo_csvserver"
		port        = 80
		servicetype = "HTTP"
	}

	resource citrixadc_appfwprofile demo_appfwprofile {
		name = "demo_appfwprofile"
		bufferoverflowaction = ["none"]
		contenttypeaction = ["none"]
		cookieconsistencyaction = ["none"]
		creditcard = ["none"]
		creditcardaction = ["none"]
		crosssitescriptingaction = ["none"]
		csrftagaction = ["none"]
		denyurlaction = ["none"]
		dynamiclearning = ["none"]
		fieldconsistencyaction = ["none"]
		fieldformataction = ["none"]
		fileuploadtypesaction = ["none"]
		inspectcontenttypes = ["none"]
		jsondosaction = ["none"]
		jsonsqlinjectionaction = ["none"]
		jsonxssaction = ["none"]
		multipleheaderaction = ["none"]
		sqlinjectionaction = ["none"]
		starturlaction = ["none"]
		type = ["HTML"]
		xmlattachmentaction = ["none"]
		xmldosaction = ["none"]
		xmlformataction = ["none"]
		xmlsoapfaultaction = ["none"]
		xmlsqlinjectionaction = ["none"]
		xmlvalidationaction = ["none"]
		xmlwsiaction = ["none"]
		xmlxssaction = ["none"]
	}

	resource citrixadc_appfwpolicy demo_appfwpolicy {
		name = "demo_appfwpolicy"
		profilename = citrixadc_appfwprofile.demo_appfwprofile.name
		rule = "true"
	}

	data "citrixadc_csvserver_appfwpolicy_binding" "demo_binding" {
		name       = citrixadc_csvserver.demo_cs.name
		policyname = citrixadc_appfwpolicy.demo_appfwpolicy.name
		depends_on = [citrixadc_csvserver_appfwpolicy_binding.demo_binding]
	}
`

func TestAccCsvserver_appfwpolicy_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccCsvserver_appfwpolicy_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_csvserver_appfwpolicy_binding.demo_binding", "name", "demo_csvserver"),
					resource.TestCheckResourceAttr("data.citrixadc_csvserver_appfwpolicy_binding.demo_binding", "policyname", "demo_appfwpolicy"),
					resource.TestCheckResourceAttr("data.citrixadc_csvserver_appfwpolicy_binding.demo_binding", "priority", "100"),
					resource.TestCheckResourceAttr("data.citrixadc_csvserver_appfwpolicy_binding.demo_binding", "gotopriorityexpression", "END"),
				),
			},
		},
	})
}

// testAcccsvserver_appfwpolicy_binding_upgrade_basic is the config used by the
// sdkv2 -> framework state-upgrade test. It reuses the same values and resource
// label as testAccCsvserver_appfwpolicy_binding_basic so it is valid under BOTH
// the SDK v2 2.2.0 schema and the current framework schema.
const testAcccsvserver_appfwpolicy_binding_upgrade_basic = `
	resource citrixadc_csvserver_appfwpolicy_binding demo_binding {
		name = citrixadc_csvserver.demo_cs.name
		priority = 100
		policyname  = citrixadc_appfwpolicy.demo_appfwpolicy.name
		gotopriorityexpression = "END"
	}
	resource "citrixadc_csvserver" "demo_cs" {
		ipv46       = "10.10.10.33"
		name        = "demo_csvserver"
		port        = 80
		servicetype = "HTTP"
	}

	resource citrixadc_appfwprofile demo_appfwprofile {
		name = "demo_appfwprofile"
		bufferoverflowaction = ["none"]
		contenttypeaction = ["none"]
		cookieconsistencyaction = ["none"]
		creditcard = ["none"]
		creditcardaction = ["none"]
		crosssitescriptingaction = ["none"]
		csrftagaction = ["none"]
		denyurlaction = ["none"]
		dynamiclearning = ["none"]
		fieldconsistencyaction = ["none"]
		fieldformataction = ["none"]
		fileuploadtypesaction = ["none"]
		inspectcontenttypes = ["none"]
		jsondosaction = ["none"]
		jsonsqlinjectionaction = ["none"]
		jsonxssaction = ["none"]
		multipleheaderaction = ["none"]
		sqlinjectionaction = ["none"]
		starturlaction = ["none"]
		type = ["HTML"]
		xmlattachmentaction = ["none"]
		xmldosaction = ["none"]
		xmlformataction = ["none"]
		xmlsoapfaultaction = ["none"]
		xmlsqlinjectionaction = ["none"]
		xmlvalidationaction = ["none"]
		xmlwsiaction = ["none"]
		xmlxssaction = ["none"]
	}

	resource citrixadc_appfwpolicy demo_appfwpolicy {
		name = "demo_appfwpolicy"
		profilename = citrixadc_appfwprofile.demo_appfwprofile.name
		rule = "true"
	}
`

// TestAccCsvserver_appfwpolicy_binding_sdkv2StateUpgrade verifies that a binding
// created with the last SDK v2 release (2.2.0, legacy comma-separated ID) is
// correctly refreshed/planned/applied by the current framework provider.
func TestAccCsvserver_appfwpolicy_binding_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckCsvserver_appfwpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			// Step 1: create the binding with the last SDK v2 release.
			// State is written with the LEGACY comma-separated id.
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.2.0",
					},
				},
				Config: testAcccsvserver_appfwpolicy_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCsvserver_appfwpolicy_bindingExist("citrixadc_csvserver_appfwpolicy_binding.demo_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_csvserver_appfwpolicy_binding.demo_binding", "id", "demo_csvserver,demo_appfwpolicy"),
				),
			},
			// Step 2: same config, current (framework) provider. Terraform
			// refreshes the legacy-id state through the framework Read
			// (exercising ParseIdString on the legacy id) then plans/applies.
			// The framework recomputes the id on read to the new key:value form.
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAcccsvserver_appfwpolicy_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCsvserver_appfwpolicy_bindingExist("citrixadc_csvserver_appfwpolicy_binding.demo_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_csvserver_appfwpolicy_binding.demo_binding", "id", "name:demo_csvserver,policyname:demo_appfwpolicy"),
				),
			},
		},
	})
}

func TestAccCsvserver_appfwpolicy_binding_import(t *testing.T) {
	const resAddr = "citrixadc_csvserver_appfwpolicy_binding.demo_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCsvserver_appfwpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccCsvserver_appfwpolicy_binding_basic},
			{Config: testAccCsvserver_appfwpolicy_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
		},
	})
}
