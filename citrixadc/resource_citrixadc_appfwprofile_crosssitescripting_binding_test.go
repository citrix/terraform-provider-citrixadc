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
	"github.com/chiradeep/go-nitro/netscaler"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"testing"
)

const testAccAppfwprofile_crosssitescripting_binding_basic = `

resource citrixadc_appfwprofile_crosssitescripting_binding demo_binding {
	name                 = citrixadc_appfwprofile.demo_appfw.name
	crosssitescripting   = "demoxss"
	formactionurl_xss    = "http://www.example.com"
	as_scan_location_xss = "HEADER"
	isregex_xss          = "NOTREGEX"
	comment              = "democomment"
	state                = "ENABLED"
  }

  resource citrixadc_appfwprofile demo_appfw {
	name                     = "demo_appfwprofile"
	bufferoverflowaction     = ["none"]
	contenttypeaction        = ["none"]
	cookieconsistencyaction  = ["none"]
	creditcard               = ["none"]
	creditcardaction         = ["none"]
	crosssitescriptingaction = ["none"]
	csrftagaction            = ["none"]
	denyurlaction            = ["none"]
	dynamiclearning          = ["none"]
	fieldconsistencyaction   = ["none"]
	fieldformataction        = ["none"]
	fileuploadtypesaction    = ["none"]
	inspectcontenttypes      = ["none"]
	jsondosaction            = ["none"]
	jsonsqlinjectionaction   = ["none"]
	jsonxssaction            = ["none"]
	multipleheaderaction     = ["none"]
	sqlinjectionaction       = ["none"]
	starturlaction           = ["none"]
	type                     = ["HTML"]
	xmlattachmentaction      = ["none"]
	xmldosaction             = ["none"]
	xmlformataction          = ["none"]
	xmlsoapfaultaction       = ["none"]
	xmlsqlinjectionaction    = ["none"]
	xmlvalidationaction      = ["none"]
	xmlwsiaction             = ["none"]
	xmlxssaction             = ["none"]
  }
`

func TestAccAppfwprofile_crosssitescripting_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAppfwprofile_crosssitescripting_bindingDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccAppfwprofile_crosssitescripting_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwprofile_crosssitescripting_bindingExist("citrixadc_appfwprofile_crosssitescripting_binding.demo_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_crosssitescripting_binding.demo_binding", "name", "demo_appfwprofile"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_crosssitescripting_binding.demo_binding", "crosssitescripting", "demoxss"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_crosssitescripting_binding.demo_binding", "formactionurl_xss", "http://www.example.com"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_crosssitescripting_binding.demo_binding", "as_scan_location_xss", "HEADER"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_crosssitescripting_binding.demo_binding", "isregex_xss", "NOTREGEX"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_crosssitescripting_binding.demo_binding", "comment", "democomment"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_crosssitescripting_binding.demo_binding", "state", "ENABLED"),
				),
			},
		},
	})
}

func testAccCheckAppfwprofile_crosssitescripting_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No appfwprofile_crosssitescripting_binding name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(netscaler.Appfwprofile_crosssitescripting_binding.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("appfwprofile_crosssitescripting_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckAppfwprofile_crosssitescripting_bindingDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_appfwprofile_crosssitescripting_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(netscaler.Appfwprofile_crosssitescripting_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("appfwprofile_crosssitescripting_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
