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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testAccAppfwprofile_jsoncmdurl_binding_basic = `
	resource "citrixadc_appfwprofile" "tf_appfwprofile" {
		name                     = "tf_appfwprofile"
		type                     = ["HTML"]
	}
	resource "citrixadc_appfwprofile_jsoncmdurl_binding" "tf_binding" {
		name           = citrixadc_appfwprofile.tf_appfwprofile.name
		jsoncmdurl     = "^https://sd2\\-zgw\\.test\\.ctxns\\.com/api/document/content$"
		alertonly      = "ON"
		isautodeployed = "AUTODEPLOYED"
		comment        = "Testing"
		state          = "DISABLED"
	}
	resource "citrixadc_appfwprofile_jsoncmdurl_binding" "tf_binding2" {
		name           = citrixadc_appfwprofile.tf_appfwprofile.name
		jsoncmdurl     = "^https://sd2\\-zgw\\.test\\.ctxns\\.com/api/v1/resource/temp$"
		keyname_json_cmd = "id"
		as_value_type_json_cmd = "SpecialString"
		as_value_expr_json_cmd = "$"
		alertonly      = "ON"
		isautodeployed = "AUTODEPLOYED"
		comment        = "Testing"
		state          = "DISABLED"
	}
`

const testAccAppfwprofile_jsoncmdurl_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion
	resource "citrixadc_appfwprofile" "tf_appfwprofile" {
		name                     = "tf_appfwprofile"
		type                     = ["HTML"]
	}
`

func TestAccAppfwprofile_jsoncmdurl_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckAppfwprofile_jsoncmdurl_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAppfwprofile_jsoncmdurl_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwprofile_jsoncmdurl_bindingExist("citrixadc_appfwprofile_jsoncmdurl_binding.tf_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_jsoncmdurl_binding.tf_binding", "name", "tf_appfwprofile"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_jsoncmdurl_binding.tf_binding", "jsoncmdurl", "^https://sd2\\-zgw\\.test\\.ctxns\\.com/api/document/content$"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_jsoncmdurl_binding.tf_binding", "alertonly", "ON"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_jsoncmdurl_binding.tf_binding", "isautodeployed", "AUTODEPLOYED"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_jsoncmdurl_binding.tf_binding", "comment", "Testing"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_jsoncmdurl_binding.tf_binding", "state", "DISABLED"),
					testAccCheckAppfwprofile_jsoncmdurl_bindingExist("citrixadc_appfwprofile_jsoncmdurl_binding.tf_binding2", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_jsoncmdurl_binding.tf_binding2", "name", "tf_appfwprofile"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_jsoncmdurl_binding.tf_binding2", "jsoncmdurl", "^https://sd2\\-zgw\\.test\\.ctxns\\.com/api/v1/resource/temp$"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_jsoncmdurl_binding.tf_binding2", "keyname_json_cmd", "id"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_jsoncmdurl_binding.tf_binding2", "as_value_type_json_cmd", "SpecialString"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_jsoncmdurl_binding.tf_binding2", "as_value_expr_json_cmd", "$"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_jsoncmdurl_binding.tf_binding2", "alertonly", "ON"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_jsoncmdurl_binding.tf_binding2", "isautodeployed", "AUTODEPLOYED"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_jsoncmdurl_binding.tf_binding2", "comment", "Testing"),
				),
			},
			{
				Config: testAccAppfwprofile_jsoncmdurl_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwprofile_jsoncmdurl_bindingNotExist("citrixadc_appfwprofile_jsoncmdurl_binding.tf_binding", "tf_appfwprofile,^https://sd2\\-zgw\\.test\\.ctxns\\.com/api/document/content$"),
					testAccCheckAppfwprofile_jsoncmdurl_bindingNotExist("citrixadc_appfwprofile_jsoncmdurl_binding.tf_binding2", "tf_appfwprofile,^https://sd2\\-zgw\\.test\\.ctxns\\.com/api/v1/resource/temp$,id,SpecialString,$"),
				),
			},
		},
	})
}

func testAccCheckAppfwprofile_jsoncmdurl_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No appfwprofile_jsoncmdurl_binding id is set")
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
		idSlice := strings.Split(bindingId, ",")

		name := idSlice[0]
		jsoncmdurl := idSlice[1]
		keyname_json_cmd := ""
		as_value_type_json_cmd := ""
		as_value_expr_json_cmd := ""
		if len(idSlice) > 2 {
			keyname_json_cmd = idSlice[2]
		}
		if len(idSlice) > 4 {
			as_value_type_json_cmd = idSlice[3]
			as_value_expr_json_cmd = idSlice[4]
		}

		findParams := service.FindParams{
			ResourceType:             "appfwprofile_jsoncmdurl_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the matching components
		found := false
		for _, v := range dataArr {
			if v["jsoncmdurl"] != nil && v["jsoncmdurl"].(string) == jsoncmdurl {
				vKeyname := ""
				if v["keyname_json_cmd"] != nil {
					vKeyname = v["keyname_json_cmd"].(string)
				}
				if keyname_json_cmd != "" {
					if vKeyname == keyname_json_cmd {
						vType := ""
						vExpr := ""
						if v["as_value_type_json_cmd"] != nil {
							vType = v["as_value_type_json_cmd"].(string)
						}
						if v["as_value_expr_json_cmd"] != nil {
							vExpr = v["as_value_expr_json_cmd"].(string)
						}
						if as_value_type_json_cmd != "" && as_value_expr_json_cmd != "" {
							if vType == as_value_type_json_cmd && vExpr == as_value_expr_json_cmd {
								found = true
								break
							}
						} else if v["as_value_type_json_cmd"] == nil && v["as_value_expr_json_cmd"] == nil {
							found = true
							break
						}
					}
				} else if v["keyname_json_cmd"] == nil {
					found = true
					break
				}
			}
		}

		if !found {
			return fmt.Errorf("appfwprofile_jsoncmdurl_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckAppfwprofile_jsoncmdurl_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}
		idSlice := strings.Split(id, ",")

		name := idSlice[0]
		jsoncmdurl := idSlice[1]
		keyname_json_cmd := ""
		as_value_type_json_cmd := ""
		as_value_expr_json_cmd := ""
		if len(idSlice) > 2 {
			keyname_json_cmd = idSlice[2]
		}
		if len(idSlice) > 4 {
			as_value_type_json_cmd = idSlice[3]
			as_value_expr_json_cmd = idSlice[4]
		}

		findParams := service.FindParams{
			ResourceType:             "appfwprofile_jsoncmdurl_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to hopefully not find the one with the matching components
		found := false
		for _, v := range dataArr {
			if v["jsoncmdurl"] != nil && v["jsoncmdurl"].(string) == jsoncmdurl {
				vKeyname := ""
				if v["keyname_json_cmd"] != nil {
					vKeyname = v["keyname_json_cmd"].(string)
				}
				if keyname_json_cmd != "" {
					if vKeyname == keyname_json_cmd {
						vType := ""
						vExpr := ""
						if v["as_value_type_json_cmd"] != nil {
							vType = v["as_value_type_json_cmd"].(string)
						}
						if v["as_value_expr_json_cmd"] != nil {
							vExpr = v["as_value_expr_json_cmd"].(string)
						}
						if as_value_type_json_cmd != "" && as_value_expr_json_cmd != "" {
							if vType == as_value_type_json_cmd && vExpr == as_value_expr_json_cmd {
								found = true
								break
							}
						} else if v["as_value_type_json_cmd"] == nil && v["as_value_expr_json_cmd"] == nil {
							found = true
							break
						}
					}
				} else if v["keyname_json_cmd"] == nil {
					found = true
					break
				}
			}
		}

		if found {
			return fmt.Errorf("appfwprofile_jsoncmdurl_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckAppfwprofile_jsoncmdurl_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_appfwprofile_jsoncmdurl_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("appfwprofile_jsoncmdurl_binding", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("appfwprofile_jsoncmdurl_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
