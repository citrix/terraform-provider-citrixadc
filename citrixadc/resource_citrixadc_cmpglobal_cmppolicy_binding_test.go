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
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"net/url"
	"testing"
)

const testAccCmpglobal_cmppolicy_binding_basic = `

resource "citrixadc_cmpglobal_cmppolicy_binding" "tf_cmpglobal_cmppolicy_binding" {
	globalbindtype = "SYSTEM_GLOBAL"
	priority   = 50
	policyname =citrixadc_cmppolicy.tf_cmppolicy.name
  }
  
  resource "citrixadc_cmppolicy" "tf_cmppolicy" {
	  name = "tf_cmppolicy"
	  rule = "HTTP.RES.HEADER(\"Content-Type\").CONTAINS(\"text\")"
	  resaction = "COMPRESS"
  }
`


const testAccCmpglobal_cmppolicy_binding_basic_step2 = `
	  
resource "citrixadc_cmppolicy" "tf_cmppolicy" {
	name = "tf_cmppolicy"
	rule = "HTTP.RES.HEADER(\"Content-Type\").CONTAINS(\"text\")"
	resaction = "COMPRESS"
}
`

func TestAccCmpglobal_cmppolicy_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCmpglobal_cmppolicy_bindingDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCmpglobal_cmppolicy_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCmpglobal_cmppolicy_bindingExist("citrixadc_cmpglobal_cmppolicy_binding.tf_cmpglobal_cmppolicy_binding", nil),
				),
			},
			resource.TestStep{
				Config: testAccCmpglobal_cmppolicy_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCmpglobal_cmppolicy_bindingNotExist("citrixadc_cmpglobal_cmppolicy_binding.tf_cmpglobal_cmppolicy_binding", "tf_cmppolicy", "RES_DEFAULT"),
				),
			},
		},
	})
}

func testAccCheckCmpglobal_cmppolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No cmpglobal_cmppolicy_binding id is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		client := testAccProvider.Meta().(*NetScalerNitroClient).client

		policyname := rs.Primary.ID
		argsMap := make(map[string]string)
		argsMap["type"] = url.QueryEscape(rs.Primary.Attributes["type"])


		findParams := service.FindParams{
			ResourceType:             "cmpglobal_cmppolicy_binding",
			ArgsMap:                   argsMap,
			ResourceMissingErrorCode:  258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the matching policyname
		found := false
		for _, v := range dataArr {
			if v["policyname"].(string) == policyname {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("cmpglobal_cmppolicy_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckCmpglobal_cmppolicy_bindingNotExist(n string, id string, typename string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*NetScalerNitroClient).client

		policyname := id
	
		findParams := service.FindParams{
			ResourceType:             "cmpglobal_cmppolicy_binding",
			ArgsMap:                  map[string]string{ "type":typename},
			ResourceMissingErrorCode:  258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to hopefully not find the one with the matching policyname
		found := false
		for _, v := range dataArr {
			if v["policyname"].(string) == policyname {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("cmpglobal_cmppolicy_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckCmpglobal_cmppolicy_bindingDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_cmpglobal_cmppolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Cmpglobal_cmppolicy_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("cmpglobal_cmppolicy_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
