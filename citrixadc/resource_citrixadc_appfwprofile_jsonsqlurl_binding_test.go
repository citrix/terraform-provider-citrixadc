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

const testAccAppfwprofile_jsonsqlurl_binding_basic = `
	resource "citrixadc_appfwprofile" "tf_appfwprofile" {
		name                     = "tf_appfwprofile"
		type                     = ["HTML"]
	}
	resource "citrixadc_appfwprofile_jsonsqlurl_binding" "tf_binding" {
		name           = citrixadc_appfwprofile.tf_appfwprofile.name
		jsonsqlurl     = "[abc][a-z]a*"
		isautodeployed = "AUTODEPLOYED"
		state          = "ENABLED"
		alertonly      = "ON"
		comment        = "Testing"
	}
	resource "citrixadc_appfwprofile_jsonsqlurl_binding" "tf_binding2" {
		name           = citrixadc_appfwprofile.tf_appfwprofile.name
		jsonsqlurl     = "[abc][a-z]a*"
		keyname_json_sql = "id"
		as_value_type_json_sql = "SpecialString"
		as_value_expr_json_sql = "p"
		isautodeployed = "AUTODEPLOYED"
		state          = "ENABLED"
		alertonly      = "ON"
		comment        = "Testing"
	}
`

const testAccAppfwprofile_jsonsqlurl_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion
	resource "citrixadc_appfwprofile" "tf_appfwprofile" {
		name                     = "tf_appfwprofile"
		type                     = ["HTML"]
	}
`

func TestAccAppfwprofile_jsonsqlurl_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAppfwprofile_jsonsqlurl_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAppfwprofile_jsonsqlurl_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwprofile_jsonsqlurl_bindingExist("citrixadc_appfwprofile_jsonsqlurl_binding.tf_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_jsonsqlurl_binding.tf_binding", "name", "tf_appfwprofile"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_jsonsqlurl_binding.tf_binding", "jsonsqlurl", "[abc][a-z]a*"),
					testAccCheckAppfwprofile_jsonsqlurl_bindingExist("citrixadc_appfwprofile_jsonsqlurl_binding.tf_binding2", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_jsonsqlurl_binding.tf_binding2", "name", "tf_appfwprofile"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_jsonsqlurl_binding.tf_binding2", "keyname_json_sql", "id"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_jsonsqlurl_binding.tf_binding2", "as_value_type_json_sql", "SpecialString"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_jsonsqlurl_binding.tf_binding2", "as_value_expr_json_sql", "p"),
				),
			},
			{
				Config: testAccAppfwprofile_jsonsqlurl_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwprofile_jsonsqlurl_bindingNotExist("citrixadc_appfwprofile_jsonsqlurl_binding.tf_binding", "tf_appfwprofile,[abc][a-z]a*"),
					testAccCheckAppfwprofile_jsonsqlurl_bindingNotExist("citrixadc_appfwprofile_jsonsqlurl_binding.tf_binding2", "tf_appfwprofile,[abc][a-z]a*,id,SpecialString,p"),
				),
			},
		},
	})
}

func testAccCheckAppfwprofile_jsonsqlurl_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No appfwprofile_jsonsqlurl_binding id is set")
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

		name := idSlice[0]
		jsonsqlurl := idSlice[1]
		keyname_json_sql := ""
		as_value_type_json_sql := ""
		as_value_expr_json_sql := ""
		if len(idSlice) > 2 {
			keyname_json_sql = idSlice[2]
		}
		if len(idSlice) > 4 {
			as_value_type_json_sql = idSlice[3]
			as_value_expr_json_sql = idSlice[4]
		}

		findParams := service.FindParams{
			ResourceType:             "appfwprofile_jsonsqlurl_binding",
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
			if v["jsonsqlurl"] != nil && v["jsonsqlurl"].(string) == jsonsqlurl {
				vKeyname := ""
				if v["keyname_json_sql"] != nil {
					vKeyname = v["keyname_json_sql"].(string)
				}
				if keyname_json_sql != "" {
					if vKeyname == keyname_json_sql {
						vType := ""
						vExpr := ""
						if v["as_value_type_json_sql"] != nil {
							vType = v["as_value_type_json_sql"].(string)
						}
						if v["as_value_expr_json_sql"] != nil {
							vExpr = v["as_value_expr_json_sql"].(string)
						}
						if as_value_type_json_sql != "" && as_value_expr_json_sql != "" {
							if vType == as_value_type_json_sql && vExpr == as_value_expr_json_sql {
								found = true
								break
							}
						} else if v["as_value_type_json_sql"] == nil && v["as_value_expr_json_sql"] == nil {
							found = true
							break
						}
					}
				} else if v["keyname_json_sql"] == nil {
					found = true
					break
				}
			}
		}

		if !found {
			return fmt.Errorf("appfwprofile_jsonsqlurl_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckAppfwprofile_jsonsqlurl_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*NetScalerNitroClient).client

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}
		idSlice := strings.Split(id, ",")

		name := idSlice[0]
		jsonsqlurl := idSlice[1]
		keyname_json_sql := ""
		as_value_type_json_sql := ""
		as_value_expr_json_sql := ""
		if len(idSlice) > 2 {
			keyname_json_sql = idSlice[2]
		}
		if len(idSlice) > 4 {
			as_value_type_json_sql = idSlice[3]
			as_value_expr_json_sql = idSlice[4]
		}

		findParams := service.FindParams{
			ResourceType:             "appfwprofile_jsonsqlurl_binding",
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
			if v["jsonsqlurl"] != nil && v["jsonsqlurl"].(string) == jsonsqlurl {
				vKeyname := ""
				if v["keyname_json_sql"] != nil {
					vKeyname = v["keyname_json_sql"].(string)
				}
				if keyname_json_sql != "" {
					if vKeyname == keyname_json_sql {
						vType := ""
						vExpr := ""
						if v["as_value_type_json_sql"] != nil {
							vType = v["as_value_type_json_sql"].(string)
						}
						if v["as_value_expr_json_sql"] != nil {
							vExpr = v["as_value_expr_json_sql"].(string)
						}
						if as_value_type_json_sql != "" && as_value_expr_json_sql != "" {
							if vType == as_value_type_json_sql && vExpr == as_value_expr_json_sql {
								found = true
								break
							}
						} else if v["as_value_type_json_sql"] == nil && v["as_value_expr_json_sql"] == nil {
							found = true
							break
						}
					}
				} else if v["keyname_json_sql"] == nil {
					found = true
					break
				}
			}
		}

		if found {
			return fmt.Errorf("appfwprofile_jsonsqlurl_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckAppfwprofile_jsonsqlurl_bindingDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_appfwprofile_jsonsqlurl_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource("appfwprofile_jsonsqlurl_binding", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("appfwprofile_jsonsqlurl_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
