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
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"strings"
	"testing"
)

const testAccAppfwprofile_xmlxss_binding_basic = `
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
	resource "citrixadc_appfwprofile_xmlxss_binding" "tf_binding" {
		name                    = citrixadc_appfwprofile.tf_appfwprofile.name
		xmlxss                  = "tf_xmlxss"
		as_scan_location_xmlxss = "ELEMENT"
		state                   = "ENABLED"
		alertonly               = "ON"
		isregex_xmlxss          = "NOTREGEX"
		isautodeployed          = "AUTODEPLOYED"
	}
`

const testAccAppfwprofile_xmlxss_binding_basic_step2 = `
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

func TestAccAppfwprofile_xmlxss_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckAppfwprofile_xmlxss_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAppfwprofile_xmlxss_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwprofile_xmlxss_bindingExist("citrixadc_appfwprofile_xmlxss_binding.tf_binding", nil),
				),
			},
			{
				Config: testAccAppfwprofile_xmlxss_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwprofile_xmlxss_bindingNotExist("citrixadc_appfwprofile_xmlxss_binding.tf_binding", "tf_appfwprofile,tf_xmlxss,ELEMENT"),
				),
			},
		},
	})
}

func testAccCheckAppfwprofile_xmlxss_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No appfwprofile_xmlxss_binding id is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		// Use the shared utility function to get a configured client
		client, err := testAccGetClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		bindingId := rs.Primary.ID

		idSlice := strings.SplitN(bindingId, ",", 3)

		name := idSlice[0]
		xmlxss := idSlice[1]
		as_scan_location_xmlxss := idSlice[2]

		findParams := service.FindParams{
			ResourceType:             "appfwprofile_xmlxss_binding",
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
			if v["xmlxss"].(string) == xmlxss {
				if v["as_scan_location_xmlxss"].(string) == as_scan_location_xmlxss {
					found = true
					break
				}
			}
		}

		if !found {
			return fmt.Errorf("appfwprofile_xmlxss_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckAppfwprofile_xmlxss_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}
		idSlice := strings.SplitN(id, ",", 3)

		name := idSlice[0]
		xmlxss := idSlice[1]
		as_scan_location_xmlxss := idSlice[2]

		findParams := service.FindParams{
			ResourceType:             "appfwprofile_xmlxss_binding",
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
			if v["xmlxss"].(string) == xmlxss {
				if v["as_scan_location_xmlxss"].(string) == as_scan_location_xmlxss {
					found = true
					break
				}
			}
		}

		if found {
			return fmt.Errorf("appfwprofile_xmlxss_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckAppfwprofile_xmlxss_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_appfwprofile_xmlxss_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Appfwprofile_xmlxss_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("appfwprofile_xmlxss_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
