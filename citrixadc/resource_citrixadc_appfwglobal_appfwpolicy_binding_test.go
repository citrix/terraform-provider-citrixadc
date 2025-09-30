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

const testAccAppfwglobal_appfwpolicy_binding_basic = `
	resource "citrixadc_appfwprofile" "tf_appfwprofile" {
		name                     = "tf_appfwprofile"
		type                     = ["HTML"]
	}
	resource "citrixadc_appfwpolicy" "tf_appfwpolicy" {
		name        = "tf_appfwpolicy"
		profilename = citrixadc_appfwprofile.tf_appfwprofile.name
		rule        = "true"
	}
	resource "citrixadc_appfwglobal_appfwpolicy_binding" "tf_binding" {
		policyname = citrixadc_appfwpolicy.tf_appfwpolicy.name
		priority   = 30
		state      = "ENABLED"
	}
	resource "citrixadc_appfwglobal_appfwpolicy_binding" "tf_binding2" {
		policyname = citrixadc_appfwpolicy.tf_appfwpolicy.name
		priority   = 30
		state      = "ENABLED"
		type  = "REQ_OVERRIDE"
	}
`

const testAccAppfwglobal_appfwpolicy_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion
	resource "citrixadc_appfwprofile" "tf_appfwprofile" {
		name                     = "tf_appfwprofile"
		type                     = ["HTML"]
	}
	resource "citrixadc_appfwpolicy" "tf_appfwpolicy" {
		name        = "tf_appfwpolicy"
		profilename = citrixadc_appfwprofile.tf_appfwprofile.name
		rule        = "true"
	}
`

func TestAccAppfwglobal_appfwpolicy_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAppfwglobal_appfwpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAppfwglobal_appfwpolicy_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwglobal_appfwpolicy_bindingExist("citrixadc_appfwglobal_appfwpolicy_binding.tf_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwglobal_appfwpolicy_binding.tf_binding", "policyname", "tf_appfwpolicy"),
					resource.TestCheckResourceAttr("citrixadc_appfwglobal_appfwpolicy_binding.tf_binding", "priority", "30"),
					resource.TestCheckResourceAttr("citrixadc_appfwglobal_appfwpolicy_binding.tf_binding", "state", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_appfwglobal_appfwpolicy_binding.tf_binding", "type", "REQ_DEFAULT"),
					resource.TestCheckResourceAttr("citrixadc_appfwglobal_appfwpolicy_binding.tf_binding", "globalbindtype", "SYSTEM_GLOBAL"),
					testAccCheckAppfwglobal_appfwpolicy_bindingExist("citrixadc_appfwglobal_appfwpolicy_binding.tf_binding2", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwglobal_appfwpolicy_binding.tf_binding2", "policyname", "tf_appfwpolicy"),
					resource.TestCheckResourceAttr("citrixadc_appfwglobal_appfwpolicy_binding.tf_binding2", "priority", "30"),
					resource.TestCheckResourceAttr("citrixadc_appfwglobal_appfwpolicy_binding.tf_binding2", "state", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_appfwglobal_appfwpolicy_binding.tf_binding2", "type", "REQ_OVERRIDE"),
					resource.TestCheckResourceAttr("citrixadc_appfwglobal_appfwpolicy_binding.tf_binding2", "globalbindtype", "SYSTEM_GLOBAL"),
				),
			},
			{
				Config: testAccAppfwglobal_appfwpolicy_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwglobal_appfwpolicy_bindingNotExist("citrixadc_appfwglobal_appfwpolicy_binding.tf_binding", "tf_appfwpolicy", "REQ_DEFAULT"),
					testAccCheckAppfwglobal_appfwpolicy_bindingNotExist("citrixadc_appfwglobal_appfwpolicy_binding.tf_binding2", "tf_appfwpolicy", "REQ_OVERRIDE"),
				),
			},
		},
	})
}

func testAccCheckAppfwglobal_appfwpolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No appfwglobal_appfwpolicy_binding id is set")
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
		policyname := idSlice[0]
		bindpoint_type := idSlice[1]
		globalbindtype := idSlice[2]

		findParams := service.FindParams{
			ResourceType:             "appfwglobal_appfwpolicy_binding",
			ArgsMap:                  map[string]string{"type": bindpoint_type},
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
			if v["policyname"].(string) == policyname && v["globalbindtype"].(string) == globalbindtype {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("appfwglobal_appfwpolicy_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckAppfwglobal_appfwpolicy_bindingNotExist(n string, id string, typename string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*NetScalerNitroClient).client
		policyname := id
		findParams := service.FindParams{
			ResourceType:             "appfwglobal_appfwpolicy_binding",
			ArgsMap:                  map[string]string{"type": typename},
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
			if v["policyname"].(string) == policyname {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("appfwglobal_appfwpolicy_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckAppfwglobal_appfwpolicy_bindingDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_appfwglobal_appfwpolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Appfwglobal_appfwpolicy_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("appfwglobal_appfwpolicy_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
