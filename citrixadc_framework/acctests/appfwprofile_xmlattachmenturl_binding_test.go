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

const testAccAppfwprofile_xmlattachmenturl_binding_basic = `
	resource "citrixadc_appfwprofile" "tf_appfwprofile" {
		name                     = "tf_appfwprofile"
		type                     = ["HTML"]
	}
	resource "citrixadc_appfwprofile_xmlattachmenturl_binding" "tf_binding" {
		name                          = citrixadc_appfwprofile.tf_appfwprofile.name
		xmlattachmenturl              = ".*"
		xmlattachmentcontenttype      = "abc*"
		alertonly                     = "ON"
		state                         = "ENABLED"
		isautodeployed                = "AUTODEPLOYED"
		comment                       = "Testing"
		xmlattachmentcontenttypecheck = "ON"
		xmlmaxattachmentsize          = "1000"
		xmlmaxattachmentsizecheck     = "ON"
	}
`

const testAccAppfwprofile_xmlattachmenturl_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion
	resource "citrixadc_appfwprofile" "tf_appfwprofile" {
		name                     = "tf_appfwprofile"
		type                     = ["HTML"]
	}
`

func TestAccAppfwprofile_xmlattachmenturl_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAppfwprofile_xmlattachmenturl_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAppfwprofile_xmlattachmenturl_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwprofile_xmlattachmenturl_bindingExist("citrixadc_appfwprofile_xmlattachmenturl_binding.tf_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_xmlattachmenturl_binding.tf_binding", "name", "tf_appfwprofile"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_xmlattachmenturl_binding.tf_binding", "xmlattachmenturl", ".*"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_xmlattachmenturl_binding.tf_binding", "xmlattachmentcontenttype", "abc*"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_xmlattachmenturl_binding.tf_binding", "alertonly", "ON"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_xmlattachmenturl_binding.tf_binding", "state", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_xmlattachmenturl_binding.tf_binding", "isautodeployed", "AUTODEPLOYED"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_xmlattachmenturl_binding.tf_binding", "comment", "Testing"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_xmlattachmenturl_binding.tf_binding", "xmlattachmentcontenttypecheck", "ON"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_xmlattachmenturl_binding.tf_binding", "xmlmaxattachmentsize", "1000"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_xmlattachmenturl_binding.tf_binding", "xmlmaxattachmentsizecheck", "ON"),
				),
			},
			{
				Config: testAccAppfwprofile_xmlattachmenturl_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwprofile_xmlattachmenturl_bindingNotExist("citrixadc_appfwprofile_xmlattachmenturl_binding.tf_binding", "tf_appfwprofile,.*"),
				),
			},
		},
	})
}

func testAccCheckAppfwprofile_xmlattachmenturl_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No appfwprofile_xmlattachmenturl_binding id is set")
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

		idSlice := strings.SplitN(bindingId, ",", 2)

		name := idSlice[0]
		xmlattachmenturl := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "appfwprofile_xmlattachmenturl_binding",
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
			if v["xmlattachmenturl"].(string) == xmlattachmenturl {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("appfwprofile_xmlattachmenturl_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckAppfwprofile_xmlattachmenturl_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}
		idSlice := strings.SplitN(id, ",", 2)

		name := idSlice[0]
		xmlattachmenturl := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "appfwprofile_xmlattachmenturl_binding",
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
			if v["xmlattachmenturl"].(string) == xmlattachmenturl {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("appfwprofile_xmlattachmenturl_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckAppfwprofile_xmlattachmenturl_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_appfwprofile_xmlattachmenturl_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Appfwprofile_xmlattachmenturl_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("appfwprofile_xmlattachmenturl_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
