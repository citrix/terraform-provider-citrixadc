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

const testAccAppfwprofile_fieldformat_binding_basic = `
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
	resource "citrixadc_appfwprofile_fieldformat_binding" "tf_binding" {
		name                 = citrixadc_appfwprofile.tf_appfwprofile.name
		fieldformat          = "tf_field"
		formactionurl_ff     = "http://www.example.com"
		comment              = "Testing"
		state                = "ENABLED"
		fieldformatmaxlength = 20
		isregexff            = "NOTREGEX"
		fieldtype            = "alpha"
		alertonly            = "OFF"
	}
`

const testAccAppfwprofile_fieldformat_binding_basic_step2 = `
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

func TestAccAppfwprofile_fieldformat_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAppfwprofile_fieldformat_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAppfwprofile_fieldformat_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwprofile_fieldformat_bindingExist("citrixadc_appfwprofile_fieldformat_binding.tf_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_fieldformat_binding.tf_binding", "name", "tf_appfwprofile"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_fieldformat_binding.tf_binding", "fieldformat", "tf_field"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_fieldformat_binding.tf_binding", "formactionurl_ff", "http://www.example.com"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_fieldformat_binding.tf_binding", "comment", "Testing"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_fieldformat_binding.tf_binding", "state", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_fieldformat_binding.tf_binding", "fieldformatmaxlength", "20"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_fieldformat_binding.tf_binding", "isregexff", "NOTREGEX"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_fieldformat_binding.tf_binding", "fieldtype", "alpha"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_fieldformat_binding.tf_binding", "alertonly", "OFF"),
				),
			},
			{
				Config: testAccAppfwprofile_fieldformat_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwprofile_fieldformat_bindingNotExist("citrixadc_appfwprofile_fieldformat_binding.tf_binding", "tf_appfwprofile,tf_field,http://www.example.com"),
				),
			},
		},
	})
}

func testAccCheckAppfwprofile_fieldformat_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No appfwprofile_fieldformat_binding id is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		client := testAccProvider.Meta().(*NetScalerNitroClient).client

		bindingId := rs.Primary.ID

		idSlice := strings.SplitN(bindingId, ",", 3)

		name := idSlice[0]
		fieldformat := idSlice[1]
		formactionurl_ff := idSlice[2]

		findParams := service.FindParams{
			ResourceType:             "appfwprofile_fieldformat_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the matching secondIdComponent
		found := false
		for _, v := range dataArr {
			if v["fieldformat"].(string) == fieldformat {
				if v["formactionurl_ff"].(string) == formactionurl_ff {
					found = true
					break
				}
			}
		}

		if !found {
			return fmt.Errorf("appfwprofile_fieldformat_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckAppfwprofile_fieldformat_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*NetScalerNitroClient).client

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}
		idSlice := strings.SplitN(id, ",", 3)

		name := idSlice[0]
		fieldformat := idSlice[1]
		formactionurl_ff := idSlice[2]

		findParams := service.FindParams{
			ResourceType:             "appfwprofile_fieldformat_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to hopefully not find the one with the matching secondIdComponent
		found := false
		for _, v := range dataArr {
			if v["fieldformat"].(string) == fieldformat {
				if v["formactionurl_ff"].(string) == formactionurl_ff {
					found = true
					break
				}
			}
		}

		if found {
			return fmt.Errorf("appfwprofile_fieldformat_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckAppfwprofile_fieldformat_bindingDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_appfwprofile_fieldformat_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Appfwprofile_fieldformat_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("appfwprofile_fieldformat_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
