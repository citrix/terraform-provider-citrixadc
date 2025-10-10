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

const testAccAppfwprofile_jsondosurl_binding_basic = `
	resource "citrixadc_appfwprofile" "tf_appfwprofile" {
		name                     = "tf_appfwprofile"
		type                     = ["HTML"]
	}
	resource "citrixadc_appfwprofile_jsondosurl_binding" "tf_binding1" {
		name                        = citrixadc_appfwprofile.tf_appfwprofile.name
		jsondosurl                  = ".*"
		state                       = "ENABLED"
		alertonly                   = "ON"
		isautodeployed              = "AUTODEPLOYED"
		jsonmaxarraylengthcheck     = "ON"
		jsonmaxdocumentlengthcheck  = "ON"
		jsonmaxcontainerdepth       = 5
		jsonmaxobjectkeylengthcheck = "OFF"
		jsonmaxarraylength          = 100000
		jsonmaxdocumentlength       = 200000
		jsonmaxobjectkeycountcheck  = "ON"
		jsonmaxobjectkeylength      = 128
		jsonmaxobjectkeycount       = 1000
		jsonmaxstringlengthcheck    = "ON"
		jsonmaxcontainerdepthcheck  = "ON"
		jsonmaxstringlength         = 1000
		comment                     = "Testing"
	}
`

const testAccAppfwprofile_jsondosurl_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion
	resource "citrixadc_appfwprofile" "tf_appfwprofile" {
		name                     = "tf_appfwprofile"
		type                     = ["HTML"]
	}
`

func TestAccAppfwprofile_jsondosurl_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAppfwprofile_jsondosurl_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAppfwprofile_jsondosurl_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwprofile_jsondosurl_bindingExist("citrixadc_appfwprofile_jsondosurl_binding.tf_binding1", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_jsondosurl_binding.tf_binding1", "name", "tf_appfwprofile"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_jsondosurl_binding.tf_binding1", "jsondosurl", ".*"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_jsondosurl_binding.tf_binding1", "state", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_jsondosurl_binding.tf_binding1", "alertonly", "ON"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_jsondosurl_binding.tf_binding1", "isautodeployed", "AUTODEPLOYED"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_jsondosurl_binding.tf_binding1", "jsonmaxarraylengthcheck", "ON"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_jsondosurl_binding.tf_binding1", "jsonmaxdocumentlengthcheck", "ON"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_jsondosurl_binding.tf_binding1", "jsonmaxcontainerdepth", "5"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_jsondosurl_binding.tf_binding1", "jsonmaxobjectkeylengthcheck", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_jsondosurl_binding.tf_binding1", "jsonmaxarraylength", "100000"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_jsondosurl_binding.tf_binding1", "jsonmaxdocumentlength", "200000"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_jsondosurl_binding.tf_binding1", "jsonmaxobjectkeycountcheck", "ON"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_jsondosurl_binding.tf_binding1", "jsonmaxobjectkeylength", "128"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_jsondosurl_binding.tf_binding1", "jsonmaxobjectkeycount", "1000"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_jsondosurl_binding.tf_binding1", "jsonmaxstringlengthcheck", "ON"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_jsondosurl_binding.tf_binding1", "jsonmaxcontainerdepthcheck", "ON"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_jsondosurl_binding.tf_binding1", "jsonmaxstringlength", "1000"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_jsondosurl_binding.tf_binding1", "comment", "Testing"),
				),
			},
			{
				Config: testAccAppfwprofile_jsondosurl_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwprofile_jsondosurl_bindingNotExist("citrixadc_appfwprofile_jsondosurl_binding.tf_binding1", "tf_appfwprofile,.*"),
				),
			},
		},
	})
}

func testAccCheckAppfwprofile_jsondosurl_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No appfwprofile_jsondosurl_binding id is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		client := testAccProvider.Meta().(*NetScalerNitroClient).client

		bindingId := rs.Primary.ID

		idSlice := strings.SplitN(bindingId, ",", 2)

		name := idSlice[0]
		jsondosurl := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "appfwprofile_jsondosurl_binding",
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
			if v["jsondosurl"].(string) == jsondosurl {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("appfwprofile_jsondosurl_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckAppfwprofile_jsondosurl_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*NetScalerNitroClient).client

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}
		idSlice := strings.SplitN(id, ",", 2)

		name := idSlice[0]
		jsondosurl := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "appfwprofile_jsondosurl_binding",
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
			if v["jsondosurl"].(string) == jsondosurl {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("appfwprofile_jsondosurl_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckAppfwprofile_jsondosurl_bindingDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_appfwprofile_jsondosurl_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource("appfwprofile_jsondosurl_binding", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("appfwprofile_jsondosurl_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
