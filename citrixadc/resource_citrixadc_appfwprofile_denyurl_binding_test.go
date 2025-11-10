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
	"net/url"
	"strings"
	"testing"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testAccAppfwprofile_denyurl_binding_basic = `
	resource citrixadc_appfwprofile demo_appfw {
		name = "tfAcc_appfwprofile"
		type = ["HTML"]
	}

	resource citrixadc_appfwprofile_denyurl_binding appfwprofile_denyurl1 {
		name = citrixadc_appfwprofile.demo_appfw.name
		denyurl = "debug[.][^/?]*(|[?].*)$"
		alertonly      = "OFF"
		isautodeployed = "NOTAUTODEPLOYED"
		state          = "ENABLED"
	}
	
	resource citrixadc_appfwprofile_denyurl_binding appfwprofile_denyurl2 {
		name = citrixadc_appfwprofile.demo_appfw.name
		denyurl = "warning[.][^/?]*(|[?].*)$"
		alertonly      = "OFF"
		isautodeployed = "NOTAUTODEPLOYED"
		state          = "ENABLED"
	}
`
const testAccAppfwprofile_denyurl_binding_basic_step2 = `
	resource citrixadc_appfwprofile demo_appfw {
		name = "tfAcc_appfwprofile"
		type = ["HTML"]
	}
`

func TestAccAppfwprofile_denyurl_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckAppfwprofile_denyurl_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAppfwprofile_denyurl_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwprofile_denyurl_bindingExist("citrixadc_appfwprofile_denyurl_binding.appfwprofile_denyurl1", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_denyurl_binding.appfwprofile_denyurl1", "name", "tfAcc_appfwprofile"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_denyurl_binding.appfwprofile_denyurl1", "denyurl", "debug[.][^/?]*(|[?].*)$"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_denyurl_binding.appfwprofile_denyurl1", "alertonly", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_denyurl_binding.appfwprofile_denyurl1", "isautodeployed", "NOTAUTODEPLOYED"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_denyurl_binding.appfwprofile_denyurl1", "state", "ENABLED"),

					testAccCheckAppfwprofile_denyurl_bindingExist("citrixadc_appfwprofile_denyurl_binding.appfwprofile_denyurl2", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_denyurl_binding.appfwprofile_denyurl2", "name", "tfAcc_appfwprofile"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_denyurl_binding.appfwprofile_denyurl2", "denyurl", "warning[.][^/?]*(|[?].*)$"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_denyurl_binding.appfwprofile_denyurl2", "alertonly", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_denyurl_binding.appfwprofile_denyurl2", "isautodeployed", "NOTAUTODEPLOYED"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_denyurl_binding.appfwprofile_denyurl2", "state", "ENABLED"),
				),
			},
			{
				Config: testAccAppfwprofile_denyurl_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwprofile_denyurl_bindingNotExist("citrixadc_appfwprofile_denyurl_binding.appfwprofile_denyurl1", "tfAcc_appfwprofile,debug[.][^/?]*(|[?].*)$"),
					testAccCheckAppfwprofile_denyurl_bindingNotExist("citrixadc_appfwprofile_denyurl_binding.appfwprofile_denyurl2", "tfAcc_appfwprofile,warning[.][^/?]*(|[?].*)$"),
				),
			},
		},
	})
}

func testAccCheckAppfwprofile_denyurl_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No appfwprofile_denyurl_binding name is set")
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

		bindingID := rs.Primary.ID
		idSlice := strings.SplitN(bindingID, ",", 2)

		if len(idSlice) < 2 {
			return fmt.Errorf("Cannot deduce appfwprofile and denyurl from ID string")
		}

		profileName := idSlice[0]
		denyURL := idSlice[1]

		findParams := service.FindParams{
			ResourceType: service.Appfwprofile_denyurl_binding.Type(),
			ResourceName: profileName,
		}
		findParams.FilterMap = make(map[string]string)
		findParams.FilterMap["denyurl"] = url.QueryEscape(denyURL)
		data, err := client.FindResourceArrayWithParams(findParams)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("appfwprofile_denyurl_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckAppfwprofile_denyurl_bindingNotExist(n string, id string) resource.TestCheckFunc {
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
		denyurl := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "appfwprofile_denyurl_binding",
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
			if v["denyurl"].(string) == denyurl {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("appfwprofile_denyurl_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckAppfwprofile_denyurl_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_appfwprofile_denyurl_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Appfwprofile_denyurl_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("appfwprofile_denyurl_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
