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

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
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
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAppfwprofile_sqlinjection_bindingDestroy,
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

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

		bindingId := rs.Primary.ID
		idSlice := strings.Split(bindingId, ",")
		appFwName := idSlice[0]
		sqlinjection := idSlice[1]
		formactionurl_sql := idSlice[2]
		as_scan_location_sql := idSlice[3]
		as_value_type_sql := idSlice[4]
		as_value_expr_sql := idSlice[5]

		findParams := service.FindParams{
			ResourceType:             service.Appfwprofile_sqlinjection_binding.Type(),
			ResourceName:             appFwName,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := nsClient.FindResourceArrayWithParams(findParams)

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
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_appfwprofile_sqlinjection_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		bindingId := rs.Primary.ID
		idSlice := strings.Split(bindingId, ",")
		appFwName := idSlice[0]

		_, err := nsClient.FindResource(service.Appfwprofile_sqlinjection_binding.Type(), appFwName)
		if err == nil {
			return fmt.Errorf("appfwprofile_sqlinjection_binding %s still exists", appFwName)
		}

	}

	return nil
}
