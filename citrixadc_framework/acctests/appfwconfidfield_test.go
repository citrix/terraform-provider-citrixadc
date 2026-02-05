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
	"log"
	"net/url"
	"testing"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testAccAppfwconfidfield_add = `
	resource "citrixadc_appfwconfidfield" "tf_confidfield1" {
		fieldname = "tf_confidfield"
		url       = "^https://sd2\\-zgw\\.test\\.ctxns\\.com/api/document/content$"
		isregex   = "REGEX"
		comment   = "Testing"
		state     = "DISABLED"
	}
	resource "citrixadc_appfwconfidfield" "tf_confidfield2" {
		fieldname = "tf_confidfield"
		url       = "^https://sd2\\-zgw\\.test\\.ctxns\\.com/api/v1/resource/temp$"
		isregex   = "REGEX"
		comment   = "Testing"
		state     = "DISABLED"
	}
`
const testAccAppfwconfidfield_update = `
	resource "citrixadc_appfwconfidfield" "tf_confidfield1" {
		fieldname = "tf_confidfield"
		url       = "^https://sd2\\-zgw\\.test\\.ctxns\\.com/api/document/content$"
		isregex   = "REGEX"
		comment   = "updated_Testing"
		state     = "DISABLED"
	}
`

func TestAccAppfwconfidfield_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAppfwconfidfieldDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAppfwconfidfield_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwconfidfieldExist("citrixadc_appfwconfidfield.tf_confidfield1", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwconfidfield.tf_confidfield1", "fieldname", "tf_confidfield"),
					resource.TestCheckResourceAttr("citrixadc_appfwconfidfield.tf_confidfield1", "comment", "Testing"),
					resource.TestCheckResourceAttr("citrixadc_appfwconfidfield.tf_confidfield1", "url", "^https://sd2\\-zgw\\.test\\.ctxns\\.com/api/document/content$"),
					resource.TestCheckResourceAttr("citrixadc_appfwconfidfield.tf_confidfield1", "isregex", "REGEX"),
					resource.TestCheckResourceAttr("citrixadc_appfwconfidfield.tf_confidfield1", "state", "DISABLED"),
					testAccCheckAppfwconfidfieldExist("citrixadc_appfwconfidfield.tf_confidfield2", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwconfidfield.tf_confidfield2", "fieldname", "tf_confidfield"),
					resource.TestCheckResourceAttr("citrixadc_appfwconfidfield.tf_confidfield2", "comment", "Testing"),
					resource.TestCheckResourceAttr("citrixadc_appfwconfidfield.tf_confidfield2", "url", "^https://sd2\\-zgw\\.test\\.ctxns\\.com/api/v1/resource/temp$"),
					resource.TestCheckResourceAttr("citrixadc_appfwconfidfield.tf_confidfield2", "isregex", "REGEX"),
					resource.TestCheckResourceAttr("citrixadc_appfwconfidfield.tf_confidfield2", "state", "DISABLED"),
				),
			},
			{
				Config: testAccAppfwconfidfield_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwconfidfieldExist("citrixadc_appfwconfidfield.tf_confidfield1", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwconfidfield.tf_confidfield1", "fieldname", "tf_confidfield"),
					resource.TestCheckResourceAttr("citrixadc_appfwconfidfield.tf_confidfield1", "comment", "updated_Testing"),
				),
			},
		},
	})
}

func testAccCheckAppfwconfidfieldExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No appfwconfidfield name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}
		appfwconfidfieldName := rs.Primary.ID
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		argsMap := make(map[string]string)
		argsMap["fieldname"] = url.QueryEscape(rs.Primary.Attributes["fieldname"])
		argsMap["url"] = url.QueryEscape(rs.Primary.Attributes["url"])
		findParams := service.FindParams{
			ResourceType: service.Appfwconfidfield.Type(),
			ArgsMap:      argsMap,
		}
		dataArray, err := client.FindResourceArrayWithParams(findParams)

		if err != nil {
			log.Printf("[WARN] citrix-provider: acceptance test: Clearing lb route state %s", appfwconfidfieldName)
			return nil
		}
		if len(dataArray) == 0 {
			log.Printf("[WARN] citrix-provider: acceptance test: Appfwconfidfield does not exist. Clearing state.")
			return nil
		}

		if len(dataArray) > 1 {
			return fmt.Errorf("[ERROR] citrix-provider: acceptance test: multiple entries found for Appfwconfidfield")
		}

		return nil
	}
}

func testAccCheckAppfwconfidfieldDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_appfwconfidfield" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}
		argsMap := make(map[string]string)
		argsMap["fieldname"] = url.QueryEscape(rs.Primary.Attributes["fieldname"])
		argsMap["url"] = url.QueryEscape(rs.Primary.Attributes["url"])
		findParams := service.FindParams{
			ResourceType: service.Appfwconfidfield.Type(),
			ArgsMap:      argsMap,
		}
		_, err := client.FindResourceArrayWithParams(findParams)

		if err == nil {
			return fmt.Errorf("appfwconfidfield %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccAppfwconfidfieldDataSource_basic = `
	resource "citrixadc_appfwconfidfield" "tf_confidfield1" {
		fieldname = "tf_confidfield"
		url       = "^https://sd2\\-zgw\\.test\\.ctxns\\.com/api/document/content$"
		isregex   = "REGEX"
		comment   = "Testing"
		state     = "DISABLED"
	}

	data "citrixadc_appfwconfidfield" "tf_confidfield1" {
		fieldname = citrixadc_appfwconfidfield.tf_confidfield1.fieldname
		url       = citrixadc_appfwconfidfield.tf_confidfield1.url
	}
`

func TestAccAppfwconfidfieldDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccAppfwconfidfieldDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_appfwconfidfield.tf_confidfield1", "fieldname", "tf_confidfield"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwconfidfield.tf_confidfield1", "url", "^https://sd2\\-zgw\\.test\\.ctxns\\.com/api/document/content$"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwconfidfield.tf_confidfield1", "isregex", "REGEX"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwconfidfield.tf_confidfield1", "comment", "Testing"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwconfidfield.tf_confidfield1", "state", "DISABLED"),
				),
			},
		},
	})
}
