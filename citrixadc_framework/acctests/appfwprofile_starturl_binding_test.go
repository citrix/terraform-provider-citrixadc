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

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testAccAppfwprofile_starturl_binding_basic = `
	resource citrixadc_appfwprofile demo_appfw {
		name = "tfAcc_appfwprofile"
		type = ["HTML"]
	}

	resource citrixadc_appfwprofile_starturl_binding appfwprofile_starturl1 {
		name = citrixadc_appfwprofile.demo_appfw.name
		starturl = "^[^?]+[.](html?|shtml|js|gif|jpg|jpeg|png|swf|pif|pdf|css|csv)$"
		alertonly      = "OFF"
		isautodeployed = "NOTAUTODEPLOYED"
		state          = "ENABLED"
	}

	resource citrixadc_appfwprofile_starturl_binding appfwprofile_starturl2 {
		name = citrixadc_appfwprofile.demo_appfw.name
		starturl = "^[^?]+[.](html?|shtml|js|gif|jpg|jpeg|png|swf|pdf|css|csv)$"
		alertonly      = "OFF"
		isautodeployed = "NOTAUTODEPLOYED"
		state          = "ENABLED"
	}
`
const testAccAppfwprofile_starturl_binding_basic_step2 = `
	resource citrixadc_appfwprofile demo_appfw {
		name = "tfAcc_appfwprofile"
		type = ["HTML"]
	}
`

func TestAccAppfwprofile_starturl_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAppfwprofile_starturl_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAppfwprofile_starturl_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwprofile_starturl_bindingExist("citrixadc_appfwprofile_starturl_binding.appfwprofile_starturl1", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_starturl_binding.appfwprofile_starturl1", "name", "tfAcc_appfwprofile"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_starturl_binding.appfwprofile_starturl1", "starturl", "^[^?]+[.](html?|shtml|js|gif|jpg|jpeg|png|swf|pif|pdf|css|csv)$"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_starturl_binding.appfwprofile_starturl1", "alertonly", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_starturl_binding.appfwprofile_starturl1", "isautodeployed", "NOTAUTODEPLOYED"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_starturl_binding.appfwprofile_starturl1", "state", "ENABLED"),
					testAccCheckAppfwprofile_starturl_bindingExist("citrixadc_appfwprofile_starturl_binding.appfwprofile_starturl2", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_starturl_binding.appfwprofile_starturl2", "name", "tfAcc_appfwprofile"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_starturl_binding.appfwprofile_starturl2", "starturl", "^[^?]+[.](html?|shtml|js|gif|jpg|jpeg|png|swf|pdf|css|csv)$"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_starturl_binding.appfwprofile_starturl2", "alertonly", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_starturl_binding.appfwprofile_starturl2", "isautodeployed", "NOTAUTODEPLOYED"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_starturl_binding.appfwprofile_starturl2", "state", "ENABLED"),
				),
			},
			{
				Config: testAccAppfwprofile_starturl_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwprofile_starturl_bindingNotExist("citrixadc_appfwprofile_starturl_binding.appfwprofile_starturl1", "tfAcc_appfwprofile,^[^?]+[.](html?|shtml|js|gif|jpg|jpeg|png|swf|pif|pdf|css|csv)$"),
					testAccCheckAppfwprofile_starturl_bindingNotExist("citrixadc_appfwprofile_starturl_binding.appfwprofile_starturl2", "tfAcc_appfwprofile,^[^?]+[.](html?|shtml|js|gif|jpg|jpeg|png|swf|pdf|css|csv)$"),
				),
			},
		},
	})
}

func TestAccAppfwprofile_starturl_binding_import(t *testing.T) {
	const resAddr = "citrixadc_appfwprofile_starturl_binding.appfwprofile_starturl1"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAppfwprofile_starturl_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccAppfwprofile_starturl_binding_basic},
			{Config: testAccAppfwprofile_starturl_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
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
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		bindingID := rs.Primary.ID
		idMap, _, err := utils.ParseIdString(bindingID, []string{"name", "starturl"}, nil)
		if err != nil {
			return fmt.Errorf("Cannot deduce appfwprofile and starturl from ID string: %v", err)
		}
		profileName := idMap["name"]
		startURL := idMap["starturl"]

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

func testAccCheckAppfwprofile_starturl_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}
		idSlice := strings.SplitN(id, ",", 3)

		name := idSlice[0]
		starturl := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "appfwprofile_starturl_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		findParams.FilterMap = make(map[string]string)
		findParams.FilterMap["starturl"] = url.QueryEscape(starturl)
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to hopefully not find the one with the matching secondIdComponent
		found := false
		for _, v := range dataArr {
			if v["starturl"].(string) == starturl {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("appfwprofile_starturl_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckAppfwprofile_starturl_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
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

const testAccAppfwprofileStarturlBindingDataSource_basic = `
	resource citrixadc_appfwprofile demo_appfw {
		name = "tfAcc_appfwprofile"
		type = ["HTML"]
	}

	resource citrixadc_appfwprofile_starturl_binding appfwprofile_starturl1 {
		name = citrixadc_appfwprofile.demo_appfw.name
		starturl = "^[^?]+[.](html?|shtml|js|gif|jpg|jpeg|png|swf|pif|pdf|css|csv)$"
		alertonly      = "OFF"
		isautodeployed = "NOTAUTODEPLOYED"
		state          = "ENABLED"
	}

	resource citrixadc_appfwprofile_starturl_binding appfwprofile_starturl2 {
		name = citrixadc_appfwprofile.demo_appfw.name
		starturl = "^[^?]+[.](html?|shtml|js|gif|jpg|jpeg|png|swf|pdf|css|csv)$"
		alertonly      = "OFF"
		isautodeployed = "NOTAUTODEPLOYED"
		state          = "ENABLED"
	}

	data "citrixadc_appfwprofile_starturl_binding" "appfwprofile_starturl1_data" {
		name       = citrixadc_appfwprofile_starturl_binding.appfwprofile_starturl1.name
		starturl   = citrixadc_appfwprofile_starturl_binding.appfwprofile_starturl1.starturl
		depends_on = [citrixadc_appfwprofile_starturl_binding.appfwprofile_starturl1]
	}

	data "citrixadc_appfwprofile_starturl_binding" "appfwprofile_starturl2_data" {
		name       = citrixadc_appfwprofile_starturl_binding.appfwprofile_starturl2.name
		starturl   = citrixadc_appfwprofile_starturl_binding.appfwprofile_starturl2.starturl
		depends_on = [citrixadc_appfwprofile_starturl_binding.appfwprofile_starturl2]
	}
`

const testAccAppfwprofile_starturl_binding_upgrade_basic = `
	resource citrixadc_appfwprofile demo_appfw {
		name = "tfAcc_appfwprofile"
		type = ["HTML"]
	}

	resource citrixadc_appfwprofile_starturl_binding appfwprofile_starturl1 {
		name = citrixadc_appfwprofile.demo_appfw.name
		starturl = "^[^?]+[.](html?|shtml|js|gif|jpg|jpeg|png|swf|pif|pdf|css|csv)$"
		alertonly      = "OFF"
		isautodeployed = "NOTAUTODEPLOYED"
		state          = "ENABLED"
	}
`

func TestAccAppfwprofile_starturl_binding_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckAppfwprofile_starturl_bindingDestroy,
		Steps: []resource.TestStep{
			{
				// Step 1: create the binding with the last SDK v2 release (2.2.0),
				// which writes state using the legacy comma-joined id.
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.2.0",
					},
				},
				Config: testAccAppfwprofile_starturl_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwprofile_starturl_bindingExist("citrixadc_appfwprofile_starturl_binding.appfwprofile_starturl1", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_starturl_binding.appfwprofile_starturl1", "id", "tfAcc_appfwprofile,^[^?]+[.](html?|shtml|js|gif|jpg|jpeg|png|swf|pif|pdf|css|csv)$"),
				),
			},
			{
				// Step 2: refresh/plan the legacy-id state through the current
				// framework provider. Read exercises ParseIdString on the legacy id
				// and SetAttrFromGet recomputes the id into the new key:value form.
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccAppfwprofile_starturl_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwprofile_starturl_bindingExist("citrixadc_appfwprofile_starturl_binding.appfwprofile_starturl1", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_starturl_binding.appfwprofile_starturl1", "id", "name:tfAcc_appfwprofile,starturl:%5E%5B%5E%3F%5D%2B%5B.%5D%28html%3F%7Cshtml%7Cjs%7Cgif%7Cjpg%7Cjpeg%7Cpng%7Cswf%7Cpif%7Cpdf%7Ccss%7Ccsv%29%24"),
				),
			},
		},
	})
}

func TestAccAppfwprofileStarturlBindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAppfwprofileStarturlBindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_starturl_binding.appfwprofile_starturl1_data", "name", "tfAcc_appfwprofile"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_starturl_binding.appfwprofile_starturl1_data", "starturl", "^[^?]+[.](html?|shtml|js|gif|jpg|jpeg|png|swf|pif|pdf|css|csv)$"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_starturl_binding.appfwprofile_starturl1_data", "alertonly", "OFF"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_starturl_binding.appfwprofile_starturl1_data", "isautodeployed", "NOTAUTODEPLOYED"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_starturl_binding.appfwprofile_starturl1_data", "state", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_starturl_binding.appfwprofile_starturl2_data", "name", "tfAcc_appfwprofile"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_starturl_binding.appfwprofile_starturl2_data", "starturl", "^[^?]+[.](html?|shtml|js|gif|jpg|jpeg|png|swf|pdf|css|csv)$"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_starturl_binding.appfwprofile_starturl2_data", "alertonly", "OFF"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_starturl_binding.appfwprofile_starturl2_data", "isautodeployed", "NOTAUTODEPLOYED"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_starturl_binding.appfwprofile_starturl2_data", "state", "ENABLED"),
				),
			},
		},
	})
}
