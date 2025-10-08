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

const testAccAppfwprofile_starturl_binding_basic = `
	resource citrixadc_appfwprofile demo_appfw {
		name = "tfAcc_appfwprofile"
		bufferoverflowaction = ["none"]
		contenttypeaction = ["none"]
		cookieconsistencyaction = ["none"]
		creditcard = ["none"]
		creditcardaction = ["none"]
		crosssitescriptingaction = ["none"]
		csrftagaction = ["none"]
		denyurlaction = ["none"]
		dynamiclearning = ["none"]
		fieldconsistencyaction = ["none"]
		fieldformataction = ["none"]
		fileuploadtypesaction = ["none"]
		inspectcontenttypes = ["none"]
		jsondosaction = ["none"]
		jsonsqlinjectionaction = ["none"]
		jsonxssaction = ["none"]
		multipleheaderaction = ["none"]
		sqlinjectionaction = ["none"]
		starturlaction = ["none"]
		type = ["HTML"]
		xmlattachmentaction = ["none"]
		xmldosaction = ["none"]
		xmlformataction = ["none"]
		xmlsoapfaultaction = ["none"]
		xmlsqlinjectionaction = ["none"]
		xmlvalidationaction = ["none"]
		xmlwsiaction = ["none"]
		xmlxssaction = ["none"]
	}

	resource citrixadc_appfwprofile_starturl_binding appfwprofile_starturl1 {
		name = citrixadc_appfwprofile.demo_appfw.name
		starturl = "^[^?]+[.](html?|shtml|js|gif|jpg|jpeg|png|swf|pif|pdf|css|csv)$"
		alertonly      = "OFF"
		isautodeployed = "NOTAUTODEPLOYED"
		state          = "ENABLED"
	}
`

func TestAccAppfwprofile_starturl_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckAppfwprofile_starturl_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAppfwprofile_starturl_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwprofile_starturl_bindingExist("citrixadc_appfwprofile_starturl_binding.appfwprofile_starturl1", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_starturl_binding.appfwprofile_starturl1", "name", "tfAcc_appfwprofile"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_starturl_binding.appfwprofile_starturl1", "starturl", "^[^?]+[.](html?|shtml|js|gif|jpg|jpeg|png|swf|pif|pdf|css|csv)$"),
				),
			},
		},
	})
}

func testAccCheckAppfwprofile_starturl_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No appfwprofile_starturl_binding name is set")
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
			return fmt.Errorf("Cannot deduce appfwprofile and starturl from ID string")
		}

		profileName := idSlice[0]
		startURL := idSlice[1]

		findParams := service.FindParams{
			ResourceType: service.Appfwprofile_starturl_binding.Type(),
			ResourceName: profileName,
		}
		findParams.FilterMap = make(map[string]string)
		findParams.FilterMap["starturl"] = url.QueryEscape(startURL)
		data, err := client.FindResourceArrayWithParams(findParams)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("appfwprofile_starturl_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckAppfwprofile_starturl_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_appfwprofile_starturl_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Appfwprofile_starturl_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("appfwprofile_starturl_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
