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

const testAccAppfwprofile_cmdinjection_binding_basic = `
	resource "citrixadc_appfwprofile" "tf_appfwprofile" {
		name                     = "tf_appfwprofile"
		bufferoverflowaction     = ["none"]
		contenttypeaction        = ["none"]
		cookieconsistencyaction  = ["none"]
		creditcard               = ["none"]
		creditcardaction         = ["none"]
		crosssitescriptingaction = ["none"]
		csrftagaction            = ["none"]
		denyurlaction            = ["none"]
		dynamiclearning          = ["none"]
		fieldconsistencyaction   = ["none"]
		fieldformataction        = ["none"]
		fileuploadtypesaction    = ["none"]
		inspectcontenttypes      = ["none"]
		jsondosaction            = ["none"]
		jsonsqlinjectionaction   = ["none"]
		jsonxssaction            = ["none"]
		multipleheaderaction     = ["none"]
		sqlinjectionaction       = ["none"]
		starturlaction           = ["none"]
		type                     = ["HTML"]
		xmlattachmentaction      = ["none"]
		xmldosaction             = ["none"]
		xmlformataction          = ["none"]
		xmlsoapfaultaction       = ["none"]
		xmlsqlinjectionaction    = ["none"]
		xmlvalidationaction      = ["none"]
		xmlwsiaction             = ["none"]
		xmlxssaction             = ["none"]
	}
	resource "citrixadc_appfwprofile_cmdinjection_binding" "tf_binding" {
		name                 = citrixadc_appfwprofile.tf_appfwprofile.name
		cmdinjection         = "tf_cmdinjection"
		formactionurl_cmd    = "http://10.10.10.10/"
		as_scan_location_cmd = "HEADER"
		as_value_type_cmd    = "Keyword"
		as_value_expr_cmd    = "[a-z]+grep"
		alertonly            = "OFF"
		isvalueregex_cmd     = "REGEX"
		isautodeployed       = "NOTAUTODEPLOYED"
		comment              = "Testing"
	}
	resource "citrixadc_appfwprofile_cmdinjection_binding" "tf_binding2" {
		name                 = citrixadc_appfwprofile.tf_appfwprofile.name
		cmdinjection         = "tf_cmdinjection"
		formactionurl_cmd    = "http://10.10.10.10/"
		as_scan_location_cmd = "COOKIE"
		as_value_type_cmd    = "Keyword"
		as_value_expr_cmd    = "[a-z]+grep"
		alertonly            = "OFF"
		isvalueregex_cmd     = "REGEX"
		isautodeployed       = "NOTAUTODEPLOYED"
		comment              = "Testing"
	}
`

const testAccAppfwprofile_cmdinjection_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion
	resource "citrixadc_appfwprofile" "tf_appfwprofile" {
		name                     = "tf_appfwprofile"
		bufferoverflowaction     = ["none"]
		contenttypeaction        = ["none"]
		cookieconsistencyaction  = ["none"]
		creditcard               = ["none"]
		creditcardaction         = ["none"]
		crosssitescriptingaction = ["none"]
		csrftagaction            = ["none"]
		denyurlaction            = ["none"]
		dynamiclearning          = ["none"]
		fieldconsistencyaction   = ["none"]
		fieldformataction        = ["none"]
		fileuploadtypesaction    = ["none"]
		inspectcontenttypes      = ["none"]
		jsondosaction            = ["none"]
		jsonsqlinjectionaction   = ["none"]
		jsonxssaction            = ["none"]
		multipleheaderaction     = ["none"]
		sqlinjectionaction       = ["none"]
		starturlaction           = ["none"]
		type                     = ["HTML"]
		xmlattachmentaction      = ["none"]
		xmldosaction             = ["none"]
		xmlformataction          = ["none"]
		xmlsoapfaultaction       = ["none"]
		xmlsqlinjectionaction    = ["none"]
		xmlvalidationaction      = ["none"]
		xmlwsiaction             = ["none"]
		xmlxssaction             = ["none"]
	}
`

func TestAccAppfwprofile_cmdinjection_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAppfwprofile_cmdinjection_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAppfwprofile_cmdinjection_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwprofile_cmdinjection_bindingExist("citrixadc_appfwprofile_cmdinjection_binding.tf_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_cmdinjection_binding.tf_binding", "name", "tf_appfwprofile"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_cmdinjection_binding.tf_binding", "cmdinjection", "tf_cmdinjection"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_cmdinjection_binding.tf_binding", "formactionurl_cmd", "http://10.10.10.10/"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_cmdinjection_binding.tf_binding", "as_scan_location_cmd", "HEADER"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_cmdinjection_binding.tf_binding", "as_value_type_cmd", "Keyword"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_cmdinjection_binding.tf_binding", "as_value_expr_cmd", "[a-z]+grep"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_cmdinjection_binding.tf_binding", "alertonly", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_cmdinjection_binding.tf_binding", "isvalueregex_cmd", "REGEX"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_cmdinjection_binding.tf_binding", "comment", "Testing"),
					testAccCheckAppfwprofile_cmdinjection_bindingExist("citrixadc_appfwprofile_cmdinjection_binding.tf_binding2", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_cmdinjection_binding.tf_binding2", "as_scan_location_cmd", "COOKIE"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_cmdinjection_binding.tf_binding2", "as_value_type_cmd", "Keyword"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_cmdinjection_binding.tf_binding2", "as_value_expr_cmd", "[a-z]+grep"),
				),
			},
			{
				Config: testAccAppfwprofile_cmdinjection_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwprofile_cmdinjection_bindingNotExist("citrixadc_appfwprofile_cmdinjection_binding.tf_binding", "tf_appfwprofile,tf_cmdinjection,http://10.10.10.10/,HEADER,Keyword,[a-z]+grep"),
				),
			},
		},
	})
}

func testAccCheckAppfwprofile_cmdinjection_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No appfwprofile_cmdinjection_binding id is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		client := testAccProvider.Meta().(*NetScalerNitroClient).client

		bindingId := rs.Primary.ID
		idSlice := strings.Split(bindingId, ",")
		appFwName := idSlice[0]
		cmdinjection := idSlice[1]
		formactionurl_cmd := idSlice[2]
		as_scan_location_cmd := idSlice[3]
		as_value_type_cmd := ""
		as_value_expr_cmd := ""
		if len(idSlice) > 4 {
			as_value_type_cmd = idSlice[4]
			as_value_expr_cmd = idSlice[5]
		}

		findParams := service.FindParams{
			ResourceType:             "appfwprofile_cmdinjection_binding",
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
			if v["cmdinjection"].(string) == cmdinjection && v["formactionurl_cmd"].(string) == formactionurl_cmd && v["as_scan_location_cmd"].(string) == as_scan_location_cmd {
				if as_value_type_cmd != "" && as_value_expr_cmd != "" {
					if v["as_value_type_cmd"] != nil && v["as_value_expr_cmd"] != nil && v["as_value_type_cmd"].(string) == as_value_type_cmd && v["as_value_expr_cmd"].(string) == as_value_expr_cmd {
						foundIndex = i
						break
					}
				} else if v["as_value_type_cmd"] == nil && v["as_value_expr_cmd"] == nil {
					foundIndex = i
					break
				}
			}
		}

		// Resource is missing
		if foundIndex == -1 {
			return fmt.Errorf("Cannot find appfwprofile_cmdinjection_binding ID %v", bindingId)
		}

		return nil
	}
}

func testAccCheckAppfwprofile_cmdinjection_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*NetScalerNitroClient).client

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}
		idSlice := strings.Split(id, ",")

		appFwName := idSlice[0]
		cmdinjection := idSlice[1]
		formactionurl_cmd := idSlice[2]
		as_scan_location_cmd := idSlice[3]
		as_value_type_cmd := ""
		as_value_expr_cmd := ""
		if len(idSlice) > 4 {
			as_value_type_cmd = idSlice[4]
			as_value_expr_cmd = idSlice[5]
		}

		findParams := service.FindParams{
			ResourceType:             "appfwprofile_cmdinjection_binding",
			ResourceName:             appFwName,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to hopefully not find the one with the matching binding
		foundIndex := -1
		for i, v := range dataArr {
			if v["cmdinjection"].(string) == cmdinjection && v["formactionurl_cmd"].(string) == formactionurl_cmd && v["as_scan_location_cmd"].(string) == as_scan_location_cmd {
				if as_value_type_cmd != "" && as_value_expr_cmd != "" {
					if v["as_value_type_cmd"] != nil && v["as_value_expr_cmd"] != nil && v["as_value_type_cmd"].(string) == as_value_type_cmd && v["as_value_expr_cmd"].(string) == as_value_expr_cmd {
						foundIndex = i
						break
					}
				} else if v["as_value_type_cmd"] == nil && v["as_value_expr_cmd"] == nil {
					foundIndex = i
					break
				}
			}
		}

		if foundIndex != -1 {
			return fmt.Errorf("appfwprofile_cmdinjection_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckAppfwprofile_cmdinjection_bindingDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_appfwprofile_cmdinjection_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource("appfwprofile_cmdinjection_binding", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("appfwprofile_cmdinjection_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
