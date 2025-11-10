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

const testAccAppfwprofile_xmldosurl_binding_basic = `
	resource "citrixadc_appfwprofile" "tf_appfwprofile" {
		name                     = "tf_appfwprofile"
		type                     = ["HTML"]
	}
	resource "citrixadc_appfwprofile_xmldosurl_binding" "tf_binding" {
		name                           = citrixadc_appfwprofile.tf_appfwprofile.name
		xmldosurl                      = ".*"
		state                          = "ENABLED"
		xmlsoaparraycheck              = "ON"
		xmlmaxelementdepthcheck        = "ON"
		xmlmaxfilesize                 = 100000
		xmlmaxfilesizecheck            = "OFF"
		xmlmaxnamespaceurilength       = 200
		xmlmaxnamespaceurilengthcheck  = "ON"
		xmlmaxelementnamelength        = 300
		xmlmaxelementnamelengthcheck   = "ON"
		xmlmaxelements                 = 30
		xmlmaxelementscheck            = "ON"
		xmlmaxattributes               = 20
		xmlmaxattributescheck          = "ON"
		xmlmaxchardatalength           = 1000
		xmlmaxchardatalengthcheck      = "ON"
		xmlmaxnamespaces               = 30
		xmlmaxnamespacescheck          = "ON"
		xmlmaxattributenamelength      = 200
		xmlmaxattributenamelengthcheck = "ON"
	}
`

const testAccAppfwprofile_xmldosurl_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion
	resource "citrixadc_appfwprofile" "tf_appfwprofile" {
		name                     = "tf_appfwprofile"
		type                     = ["HTML"]
	}
`

func TestAccAppfwprofile_xmldosurl_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckAppfwprofile_xmldosurl_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAppfwprofile_xmldosurl_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwprofile_xmldosurl_bindingExist("citrixadc_appfwprofile_xmldosurl_binding.tf_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_xmldosurl_binding.tf_binding", "name", "tf_appfwprofile"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_xmldosurl_binding.tf_binding", "xmldosurl", ".*"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_xmldosurl_binding.tf_binding", "state", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_xmldosurl_binding.tf_binding", "xmlsoaparraycheck", "ON"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_xmldosurl_binding.tf_binding", "xmlmaxelementdepthcheck", "ON"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_xmldosurl_binding.tf_binding", "xmlmaxfilesize", "100000"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_xmldosurl_binding.tf_binding", "xmlmaxfilesizecheck", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_xmldosurl_binding.tf_binding", "xmlmaxnamespaceurilength", "200"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_xmldosurl_binding.tf_binding", "xmlmaxnamespaceurilengthcheck", "ON"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_xmldosurl_binding.tf_binding", "xmlmaxelementnamelength", "300"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_xmldosurl_binding.tf_binding", "xmlmaxelementnamelengthcheck", "ON"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_xmldosurl_binding.tf_binding", "xmlmaxelements", "30"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_xmldosurl_binding.tf_binding", "xmlmaxelementscheck", "ON"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_xmldosurl_binding.tf_binding", "xmlmaxattributes", "20"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_xmldosurl_binding.tf_binding", "xmlmaxattributescheck", "ON"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_xmldosurl_binding.tf_binding", "xmlmaxchardatalength", "1000"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_xmldosurl_binding.tf_binding", "xmlmaxchardatalengthcheck", "ON"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_xmldosurl_binding.tf_binding", "xmlmaxnamespaces", "30"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_xmldosurl_binding.tf_binding", "xmlmaxnamespacescheck", "ON"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_xmldosurl_binding.tf_binding", "xmlmaxattributenamelength", "200"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_xmldosurl_binding.tf_binding", "xmlmaxattributenamelengthcheck", "ON"),
				),
			},
			{
				Config: testAccAppfwprofile_xmldosurl_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwprofile_xmldosurl_bindingNotExist("citrixadc_appfwprofile_xmldosurl_binding.tf_binding", "tf_appfwprofile,.*"),
				),
			},
		},
	})
}

func testAccCheckAppfwprofile_xmldosurl_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No appfwprofile_xmldosurl_binding id is set")
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

		idSlice := strings.SplitN(bindingId, ",", 2)

		name := idSlice[0]
		xmldosurl := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "appfwprofile_xmldosurl_binding",
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
			if v["xmldosurl"].(string) == xmldosurl {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("appfwprofile_xmldosurl_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckAppfwprofile_xmldosurl_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}
		idSlice := strings.SplitN(id, ",", 2)

		name := idSlice[0]
		xmldosurl := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "appfwprofile_xmldosurl_binding",
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
			if v["xmldosurl"].(string) == xmldosurl {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("appfwprofile_xmldosurl_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckAppfwprofile_xmldosurl_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_appfwprofile_xmldosurl_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Appfwprofile_xmldosurl_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("appfwprofile_xmldosurl_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
