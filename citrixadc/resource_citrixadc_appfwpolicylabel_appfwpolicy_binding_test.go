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

const testAccAppfwpolicylabel_appfwpolicy_binding_basic = `
	resource "citrixadc_appfwpolicylabel" "tf_appfwpolicylabel" {
		labelname       = "tf_appfwpolicylabel"
		policylabeltype = "http_req"
	}
	resource "citrixadc_appfwprofile" "tf_appfwprofile" {
		name                     = "tf_appfwprofile"
		type                     = ["HTML"]
	}
	resource "citrixadc_appfwpolicy" "tf_appfwpolicy1" {
		name        = "tf_appfwpolicy1"
		profilename = citrixadc_appfwprofile.tf_appfwprofile.name
		rule        = "true"
	}
	resource "citrixadc_appfwpolicy" "tf_appfwpolicy2" {
		name        = "tf_appfwpolicy2"
		profilename = citrixadc_appfwprofile.tf_appfwprofile.name
		rule        = "true"
	}
	resource "citrixadc_appfwpolicylabel_appfwpolicy_binding" "tf_binding1" {
		labelname  = citrixadc_appfwpolicylabel.tf_appfwpolicylabel.labelname
		policyname = citrixadc_appfwpolicy.tf_appfwpolicy1.name
		priority   = 90
	}
	resource "citrixadc_appfwpolicylabel_appfwpolicy_binding" "tf_binding2" {
		labelname  = citrixadc_appfwpolicylabel.tf_appfwpolicylabel.labelname
		policyname = citrixadc_appfwpolicy.tf_appfwpolicy2.name
		priority   = 100
	}
`

const testAccAppfwpolicylabel_appfwpolicy_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion
	resource "citrixadc_appfwpolicylabel" "tf_appfwpolicylabel" {
		labelname       = "tf_appfwpolicylabel"
		policylabeltype = "http_req"
	}
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

func TestAccAppfwpolicylabel_appfwpolicy_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAppfwpolicylabel_appfwpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAppfwpolicylabel_appfwpolicy_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwpolicylabel_appfwpolicy_bindingExist("citrixadc_appfwpolicylabel_appfwpolicy_binding.tf_binding1", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwpolicylabel_appfwpolicy_binding.tf_binding1", "priority", "90"),
					testAccCheckAppfwpolicylabel_appfwpolicy_bindingExist("citrixadc_appfwpolicylabel_appfwpolicy_binding.tf_binding2", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwpolicylabel_appfwpolicy_binding.tf_binding2", "priority", "100"),
				),
			},
			{
				Config: testAccAppfwpolicylabel_appfwpolicy_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwpolicylabel_appfwpolicy_bindingNotExist("citrixadc_appfwpolicylabel_appfwpolicy_binding.tf_binding1", "tf_appfwpolicylabel,tf_appfwpolicy1"),
					testAccCheckAppfwpolicylabel_appfwpolicy_bindingNotExist("citrixadc_appfwpolicylabel_appfwpolicy_binding.tf_binding2", "tf_appfwpolicylabel,tf_appfwpolicy2"),
				),
			},
		},
	})
}

func testAccCheckAppfwpolicylabel_appfwpolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No appfwpolicylabel_appfwpolicy_binding id is set")
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

		labelname := idSlice[0]
		policyname := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "appfwpolicylabel_appfwpolicy_binding",
			ResourceName:             labelname,
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
			if v["policyname"].(string) == policyname {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("appfwpolicylabel_appfwpolicy_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckAppfwpolicylabel_appfwpolicy_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*NetScalerNitroClient).client

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}
		idSlice := strings.SplitN(id, ",", 2)

		labelname := idSlice[0]
		policyname := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "appfwpolicylabel_appfwpolicy_binding",
			ResourceName:             labelname,
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
			return fmt.Errorf("appfwpolicylabel_appfwpolicy_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckAppfwpolicylabel_appfwpolicy_bindingDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_appfwpolicylabel_appfwpolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Appfwpolicylabel_appfwpolicy_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("appfwpolicylabel_appfwpolicy_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
