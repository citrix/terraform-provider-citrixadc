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

const testAccAppfwprofile_sqlinjection_binding_basic = `
	resource "citrixadc_appfwprofile_sqlinjection_binding" "appfw-szw-bi-test-sqlinject-relax-7" {
		name                 = citrixadc_appfwprofile.demo_appfw.name
		sqlinjection         = "data"
		isautodeployed       = "NOTAUTODEPLOYED"
		as_scan_location_sql = "FORMFIELD"
		formactionurl_sql    = "^https://citrix.csg.com/analytics/saw.dll$"
		as_value_type_sql    = "Keyword"
		isvalueregex_sql     = "REGEX"
		as_value_expr_sql    = ".*"
		state                = "ENABLED"
		depends_on           = [citrixadc_appfwprofile.demo_appfw]
	}

	resource "citrixadc_appfwprofile_sqlinjection_binding" "appfw-szw-bi-test-sqlinject-relax-8" {
		name                 = citrixadc_appfwprofile.demo_appfw.name
		sqlinjection         = "data"
		isautodeployed       = "NOTAUTODEPLOYED"
		as_scan_location_sql = "FORMFIELD"
		formactionurl_sql    = "^https://citrix.csg.com/dv/ui/api/v1/maps/suggest$"
		as_value_type_sql    = "Keyword"
		isvalueregex_sql     = "REGEX"
		as_value_expr_sql    = ".*"
		state                = "ENABLED"
		depends_on           = [citrixadc_appfwprofile.demo_appfw]
	}

	resource "citrixadc_appfwprofile_sqlinjection_binding" "appfw-szw-bi-test-sqlinject-relax-9" {
		name                 = citrixadc_appfwprofile.demo_appfw.name
		sqlinjection         = "data"
		isautodeployed       = "NOTAUTODEPLOYED"
		as_scan_location_sql = "FORMFIELD"
		formactionurl_sql    = "^https://citrix.csg.com/dv/ui/api/v1/maps/suggest$"
		as_value_type_sql    = "Keyword"
		isvalueregex_sql     = "REGEX"
		as_value_expr_sql    = ".*"
		state                = "ENABLED"
		ruletype             = "DENY"
		depends_on           = [citrixadc_appfwprofile.demo_appfw]
	}

	resource "citrixadc_appfwprofile" "demo_appfw" {
		name                     = "demo_appfwprofile"
		type                     = ["HTML"]
	}
`

func TestAccAppfwprofile_sqlinjection_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAppfwprofile_sqlinjection_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAppfwprofile_sqlinjection_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwprofile_sqlinjection_bindingExist("citrixadc_appfwprofile_sqlinjection_binding.appfw-szw-bi-test-sqlinject-relax-7", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_sqlinjection_binding.appfw-szw-bi-test-sqlinject-relax-7", "name", "demo_appfwprofile"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_sqlinjection_binding.appfw-szw-bi-test-sqlinject-relax-7", "sqlinjection", "data"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_sqlinjection_binding.appfw-szw-bi-test-sqlinject-relax-7", "formactionurl_sql", "^https://citrix.csg.com/analytics/saw.dll$"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_sqlinjection_binding.appfw-szw-bi-test-sqlinject-relax-7", "as_scan_location_sql", "FORMFIELD"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_sqlinjection_binding.appfw-szw-bi-test-sqlinject-relax-7", "isvalueregex_sql", "REGEX"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_sqlinjection_binding.appfw-szw-bi-test-sqlinject-relax-7", "state", "ENABLED"),
					testAccCheckAppfwprofile_sqlinjection_bindingExist("citrixadc_appfwprofile_sqlinjection_binding.appfw-szw-bi-test-sqlinject-relax-8", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_sqlinjection_binding.appfw-szw-bi-test-sqlinject-relax-8", "name", "demo_appfwprofile"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_sqlinjection_binding.appfw-szw-bi-test-sqlinject-relax-8", "sqlinjection", "data"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_sqlinjection_binding.appfw-szw-bi-test-sqlinject-relax-8", "formactionurl_sql", "^https://citrix.csg.com/dv/ui/api/v1/maps/suggest$"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_sqlinjection_binding.appfw-szw-bi-test-sqlinject-relax-8", "as_scan_location_sql", "FORMFIELD"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_sqlinjection_binding.appfw-szw-bi-test-sqlinject-relax-8", "isvalueregex_sql", "REGEX"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_sqlinjection_binding.appfw-szw-bi-test-sqlinject-relax-8", "state", "ENABLED"),
					testAccCheckAppfwprofile_sqlinjection_bindingExist("citrixadc_appfwprofile_sqlinjection_binding.appfw-szw-bi-test-sqlinject-relax-9", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_sqlinjection_binding.appfw-szw-bi-test-sqlinject-relax-9", "name", "demo_appfwprofile"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_sqlinjection_binding.appfw-szw-bi-test-sqlinject-relax-9", "sqlinjection", "data"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_sqlinjection_binding.appfw-szw-bi-test-sqlinject-relax-9", "formactionurl_sql", "^https://citrix.csg.com/dv/ui/api/v1/maps/suggest$"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_sqlinjection_binding.appfw-szw-bi-test-sqlinject-relax-9", "as_scan_location_sql", "FORMFIELD"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_sqlinjection_binding.appfw-szw-bi-test-sqlinject-relax-9", "isvalueregex_sql", "REGEX"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_sqlinjection_binding.appfw-szw-bi-test-sqlinject-relax-9", "state", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_sqlinjection_binding.appfw-szw-bi-test-sqlinject-relax-9", "ruletype", "DENY"),
				),
			},
		},
	})
}

func testAccCheckAppfwprofile_sqlinjection_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No appfwprofile_sqlinjection_binding name is set")
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
		idMap, _, err := utils.ParseIdString(bindingId, []string{"name", "sqlinjection", "formactionurl_sql", "as_scan_location_sql", "as_value_type_sql", "as_value_expr_sql", "ruletype"}, []string{"as_value_type_sql", "as_value_expr_sql", "ruletype"})
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", bindingId, err)
		}
		appFwName := idMap["name"]
		sqlinjection := idMap["sqlinjection"]
		formactionurl_sql := idMap["formactionurl_sql"]
		as_scan_location_sql := idMap["as_scan_location_sql"]
		as_value_type_sql := idMap["as_value_type_sql"]
		as_value_expr_sql := idMap["as_value_expr_sql"]

		findParams := service.FindParams{
			ResourceType:             service.Appfwprofile_sqlinjection_binding.Type(),
			ResourceName:             appFwName,
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
			if v["sqlinjection"].(string) == sqlinjection {
				if v["formactionurl_sql"] != nil && v["as_scan_location_sql"] != nil && v["as_value_type_sql"] != nil && v["as_value_expr_sql"] != nil && v["as_value_type_sql"].(string) == as_value_type_sql && v["as_value_expr_sql"].(string) == as_value_expr_sql && v["as_scan_location_sql"].(string) == as_scan_location_sql && v["formactionurl_sql"].(string) == formactionurl_sql {
					foundIndex = i
					break
				}
			}
		}

		// Resource is missing
		if foundIndex == -1 {
			return fmt.Errorf("Cannot find appfwprofile_sqlinjection_binding ID %v", bindingId)
		}

		return nil
	}
}

func testAccCheckAppfwprofile_sqlinjection_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_appfwprofile_sqlinjection_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		bindingId := rs.Primary.ID
		idMap, _, err := utils.ParseIdString(bindingId, []string{"name", "sqlinjection", "formactionurl_sql", "as_scan_location_sql", "as_value_type_sql", "as_value_expr_sql", "ruletype"}, []string{"as_value_type_sql", "as_value_expr_sql", "ruletype"})
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", bindingId, err)
		}
		appFwName := idMap["name"]

		_, err = client.FindResource(service.Appfwprofile_sqlinjection_binding.Type(), appFwName)
		if err == nil {
			return fmt.Errorf("appfwprofile_sqlinjection_binding %s still exists", appFwName)
		}

	}

	return nil
}

const testAccAppfwprofileSqlinjectionBindingDataSource_basic = `
	resource "citrixadc_appfwprofile_sqlinjection_binding" "appfw-szw-bi-test-sqlinject-relax-7" {
		name                 = citrixadc_appfwprofile.demo_appfw.name
		sqlinjection         = "data"
		isautodeployed       = "NOTAUTODEPLOYED"
		as_scan_location_sql = "FORMFIELD"
		formactionurl_sql    = "^https://citrix.csg.com/analytics/saw.dll$"
		as_value_type_sql    = "Keyword"
		isvalueregex_sql     = "REGEX"
		as_value_expr_sql    = ".*"
		state                = "ENABLED"
		depends_on           = [citrixadc_appfwprofile.demo_appfw]
	}

	resource "citrixadc_appfwprofile_sqlinjection_binding" "appfw-szw-bi-test-sqlinject-relax-8" {
		name                 = citrixadc_appfwprofile.demo_appfw.name
		sqlinjection         = "data"
		isautodeployed       = "NOTAUTODEPLOYED"
		as_scan_location_sql = "FORMFIELD"
		formactionurl_sql    = "^https://citrix.csg.com/dv/ui/api/v1/maps/suggest$"
		as_value_type_sql    = "Keyword"
		isvalueregex_sql     = "REGEX"
		as_value_expr_sql    = ".*"
		state                = "ENABLED"
		depends_on           = [citrixadc_appfwprofile.demo_appfw]
	}

	resource "citrixadc_appfwprofile" "demo_appfw" {
		name                     = "demo_appfwprofile"
		type                     = ["HTML"]
	}

	data "citrixadc_appfwprofile_sqlinjection_binding" "tf_binding_data" {
		name                 = citrixadc_appfwprofile_sqlinjection_binding.appfw-szw-bi-test-sqlinject-relax-7.name
		sqlinjection         = citrixadc_appfwprofile_sqlinjection_binding.appfw-szw-bi-test-sqlinject-relax-7.sqlinjection
		as_scan_location_sql = citrixadc_appfwprofile_sqlinjection_binding.appfw-szw-bi-test-sqlinject-relax-7.as_scan_location_sql
		formactionurl_sql    = citrixadc_appfwprofile_sqlinjection_binding.appfw-szw-bi-test-sqlinject-relax-7.formactionurl_sql
		as_value_type_sql    = citrixadc_appfwprofile_sqlinjection_binding.appfw-szw-bi-test-sqlinject-relax-7.as_value_type_sql
		as_value_expr_sql    = citrixadc_appfwprofile_sqlinjection_binding.appfw-szw-bi-test-sqlinject-relax-7.as_value_expr_sql
		depends_on           = [citrixadc_appfwprofile_sqlinjection_binding.appfw-szw-bi-test-sqlinject-relax-7]
	}

	data "citrixadc_appfwprofile_sqlinjection_binding" "tf_binding2_data" {
		name                 = citrixadc_appfwprofile_sqlinjection_binding.appfw-szw-bi-test-sqlinject-relax-8.name
		sqlinjection         = citrixadc_appfwprofile_sqlinjection_binding.appfw-szw-bi-test-sqlinject-relax-8.sqlinjection
		as_scan_location_sql = citrixadc_appfwprofile_sqlinjection_binding.appfw-szw-bi-test-sqlinject-relax-8.as_scan_location_sql
		formactionurl_sql    = citrixadc_appfwprofile_sqlinjection_binding.appfw-szw-bi-test-sqlinject-relax-8.formactionurl_sql
		as_value_type_sql    = citrixadc_appfwprofile_sqlinjection_binding.appfw-szw-bi-test-sqlinject-relax-8.as_value_type_sql
		as_value_expr_sql    = citrixadc_appfwprofile_sqlinjection_binding.appfw-szw-bi-test-sqlinject-relax-8.as_value_expr_sql
		depends_on           = [citrixadc_appfwprofile_sqlinjection_binding.appfw-szw-bi-test-sqlinject-relax-8]
	}
`

func TestAccAppfwprofileSqlinjectionBindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAppfwprofileSqlinjectionBindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_sqlinjection_binding.tf_binding_data", "name", "demo_appfwprofile"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_sqlinjection_binding.tf_binding_data", "sqlinjection", "data"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_sqlinjection_binding.tf_binding_data", "formactionurl_sql", "^https://citrix.csg.com/analytics/saw.dll$"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_sqlinjection_binding.tf_binding_data", "as_scan_location_sql", "FORMFIELD"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_sqlinjection_binding.tf_binding_data", "as_value_type_sql", "Keyword"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_sqlinjection_binding.tf_binding_data", "as_value_expr_sql", ".*"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_sqlinjection_binding.tf_binding_data", "isvalueregex_sql", "REGEX"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_sqlinjection_binding.tf_binding_data", "state", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_sqlinjection_binding.tf_binding_data", "isautodeployed", "NOTAUTODEPLOYED"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_sqlinjection_binding.tf_binding2_data", "name", "demo_appfwprofile"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_sqlinjection_binding.tf_binding2_data", "sqlinjection", "data"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_sqlinjection_binding.tf_binding2_data", "formactionurl_sql", "^https://citrix.csg.com/dv/ui/api/v1/maps/suggest$"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_sqlinjection_binding.tf_binding2_data", "as_scan_location_sql", "FORMFIELD"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_sqlinjection_binding.tf_binding2_data", "as_value_type_sql", "Keyword"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_sqlinjection_binding.tf_binding2_data", "as_value_expr_sql", ".*"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_sqlinjection_binding.tf_binding2_data", "isvalueregex_sql", "REGEX"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_sqlinjection_binding.tf_binding2_data", "state", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_sqlinjection_binding.tf_binding2_data", "isautodeployed", "NOTAUTODEPLOYED"),
				),
			},
		},
	})
}

const testAccAppfwprofile_sqlinjection_binding_upgrade_basic = `
	resource "citrixadc_appfwprofile_sqlinjection_binding" "appfw-szw-bi-test-sqlinject-relax-7" {
		name                 = citrixadc_appfwprofile.demo_appfw.name
		sqlinjection         = "data"
		isautodeployed       = "NOTAUTODEPLOYED"
		as_scan_location_sql = "FORMFIELD"
		formactionurl_sql    = "^https://citrix.csg.com/analytics/saw.dll$"
		as_value_type_sql    = "Keyword"
		isvalueregex_sql     = "REGEX"
		as_value_expr_sql    = ".*"
		state                = "ENABLED"
		depends_on           = [citrixadc_appfwprofile.demo_appfw]
	}

	resource "citrixadc_appfwprofile" "demo_appfw" {
		name = "demo_appfwprofile"
		type = ["HTML"]
	}
`

// TestAccAppfwprofile_sqlinjection_binding_sdkv2StateUpgrade verifies that a
// resource created with the last SDK v2 release (2.2.0, legacy comma-joined ID)
// upgrades cleanly when refreshed through the current Framework provider, and
// that the Framework Read re-derives the canonical new-format (key:value) ID.
func TestAccAppfwprofile_sqlinjection_binding_import(t *testing.T) {
	const resAddr = "citrixadc_appfwprofile_sqlinjection_binding.appfw-szw-bi-test-sqlinject-relax-7"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAppfwprofile_sqlinjection_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccAppfwprofile_sqlinjection_binding_basic},
			{Config: testAccAppfwprofile_sqlinjection_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
		},
	})
}

func TestAccAppfwprofile_sqlinjection_binding_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckAppfwprofile_sqlinjection_bindingDestroy,
		Steps: []resource.TestStep{
			// Step 1: create with the last SDK v2 release -> legacy positional ID.
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.2.0",
					},
				},
				Config: testAccAppfwprofile_sqlinjection_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwprofile_sqlinjection_bindingExist("citrixadc_appfwprofile_sqlinjection_binding.appfw-szw-bi-test-sqlinject-relax-7", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_sqlinjection_binding.appfw-szw-bi-test-sqlinject-relax-7", "id", "demo_appfwprofile,data,^https://citrix.csg.com/analytics/saw.dll$,FORMFIELD,Keyword,.*,ALLOW"),
				),
			},
			// Step 2: refresh/apply the SAME config through the current Framework
			// provider. Read parses the legacy ID and re-derives the new-format ID.
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccAppfwprofile_sqlinjection_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwprofile_sqlinjection_bindingExist("citrixadc_appfwprofile_sqlinjection_binding.appfw-szw-bi-test-sqlinject-relax-7", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_sqlinjection_binding.appfw-szw-bi-test-sqlinject-relax-7", "id", "as_scan_location_sql:FORMFIELD,as_value_expr_sql:.%2A,as_value_type_sql:Keyword,formactionurl_sql:%5Ehttps%3A%2F%2Fcitrix.csg.com%2Fanalytics%2Fsaw.dll%24,name:demo_appfwprofile,sqlinjection:data"),
				),
			},
		},
	})
}
