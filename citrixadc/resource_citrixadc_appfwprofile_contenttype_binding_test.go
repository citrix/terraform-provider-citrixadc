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

const testAccAppfwprofile_contenttype_binding_basic = `
	resource "citrixadc_appfwprofile" "tf_appfwprofile" {
		name                     = "tf_appfwprofile"
		type                     = ["HTML"]
	}
	resource "citrixadc_appfwprofile_contenttype_binding" "tf_binding1" {
		name           = citrixadc_appfwprofile.tf_appfwprofile.name
		contenttype    = "hello"
		state          = "ENABLED"
		alertonly      = "ON"
		isautodeployed = "NOTAUTODEPLOYED"
		comment        = "Testing"
	}
	resource "citrixadc_appfwprofile_contenttype_binding" "tf_binding2" {
		name           = citrixadc_appfwprofile.tf_appfwprofile.name
		contenttype    = "world"
		state          = "ENABLED"
		alertonly      = "ON"
		isautodeployed = "NOTAUTODEPLOYED"
		comment        = "Testing"
	}
`

const testAccAppfwprofile_contenttype_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion
	resource "citrixadc_appfwprofile" "tf_appfwprofile" {
		name                     = "tf_appfwprofile"
		type                     = ["HTML"]
	}
`

func TestAccAppfwprofile_contenttype_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckAppfwprofile_contenttype_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAppfwprofile_contenttype_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwprofile_contenttype_bindingExist("citrixadc_appfwprofile_contenttype_binding.tf_binding1", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_contenttype_binding.tf_binding1", "contenttype", "hello"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_contenttype_binding.tf_binding1", "state", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_contenttype_binding.tf_binding1", "alertonly", "ON"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_contenttype_binding.tf_binding1", "isautodeployed", "NOTAUTODEPLOYED"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_contenttype_binding.tf_binding1", "comment", "Testing"),
					testAccCheckAppfwprofile_contenttype_bindingExist("citrixadc_appfwprofile_contenttype_binding.tf_binding2", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_contenttype_binding.tf_binding2", "contenttype", "world"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_contenttype_binding.tf_binding2", "state", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_contenttype_binding.tf_binding2", "alertonly", "ON"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_contenttype_binding.tf_binding2", "isautodeployed", "NOTAUTODEPLOYED"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_contenttype_binding.tf_binding2", "comment", "Testing"),
				),
			},
			{
				Config: testAccAppfwprofile_contenttype_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwprofile_contenttype_bindingNotExist("citrixadc_appfwprofile_contenttype_binding.tf_binding1", "tf_appfwprofile,hello"),
					testAccCheckAppfwprofile_contenttype_bindingNotExist("citrixadc_appfwprofile_contenttype_binding.tf_binding2", "tf_appfwprofile,world"),
				),
			},
		},
	})
}

func testAccCheckAppfwprofile_contenttype_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No appfwprofile_contenttype_binding id is set")
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
		contenttype := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "appfwprofile_contenttype_binding",
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
			if v["contenttype"].(string) == contenttype {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("appfwprofile_contenttype_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckAppfwprofile_contenttype_bindingNotExist(n string, id string) resource.TestCheckFunc {
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
		contenttype := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "appfwprofile_contenttype_binding",
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
			if v["contenttype"].(string) == contenttype {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("appfwprofile_contenttype_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckAppfwprofile_contenttype_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_appfwprofile_contenttype_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Appfwprofile_contenttype_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("appfwprofile_contenttype_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
