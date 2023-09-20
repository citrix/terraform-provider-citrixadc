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
	"testing"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

const testAccAppfwsettings_basic_step1 = `
resource "citrixadc_appfwsettings" "tf_appfwsettings" {
  defaultprofile           = "APPFW_DROP"
  undefaction              = "APPFW_DROP"
  sessiontimeout           = 800
  learnratelimit           = 300
  sessionlifetime          = 1000
  sessioncookiename        = "test_ns_id"
  importsizelimit          = 134217700
  signatureautoupdate      = "ON"
  signatureurl             = "https://example.com"
  cookiepostencryptprefix  = "ENCRYPTED"
  geolocationlogging       = "ON"
  ceflogging               = "ON"
  entitydecoding           = "ON"
  useconfigurablesecretkey = "ON"
  sessionlimit             = 0
  malformedreqaction = [
	"none",
  ]
  centralizedlearning = "ON"
  proxyport           = 9090
}

`

const testAccAppfwsettings_basic_step2 = `
resource "citrixadc_appfwsettings" "tf_appfwsettings" {
  defaultprofile           = "APPFW_BYPASS"
  undefaction              = "APPFW_BLOCK"
  sessiontimeout           = 900
  learnratelimit           = 400
  sessionlifetime          = 0
  sessioncookiename        = "citrix_ns_id"
  importsizelimit          = 134217728
  signatureautoupdate      = "OFF"
  signatureurl             = "https://s3.amazonaws.com/NSAppFwSignatures/SignaturesMapping.xml"
  cookiepostencryptprefix  = "ENC"
  geolocationlogging       = "OFF"
  ceflogging               = "OFF"
  entitydecoding           = "OFF"
  useconfigurablesecretkey = "OFF"
  sessionlimit             = 100000
  malformedreqaction = [
    "block",
    "log",
    "stats"
  ]
  centralizedlearning = "OFF"
  proxyport           = 8080
}

`

func TestAccAppfwsettings_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccAppfwsettings_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwsettingsExist("citrixadc_appfwsettings.tf_appfwsettings", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_appfwsettings", "defaultprofile", "APPFW_DROP"),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_appfwsettings", "undefaction", "APPFW_DROP"),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_appfwsettings", "sessiontimeout", "800"),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_appfwsettings", "learnratelimit", "300"),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_appfwsettings", "sessionlifetime", "1000"),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_appfwsettings", "sessioncookiename", "test_ns_id"),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_appfwsettings", "importsizelimit", "134217700"),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_appfwsettings", "signatureautoupdate", "ON"),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_appfwsettings", "signatureurl", "https://example.com"),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_appfwsettings", "cookiepostencryptprefix", "ENCRYPTED"),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_appfwsettings", "geolocationlogging", "ON"),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_appfwsettings", "ceflogging", "ON"),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_appfwsettings", "entitydecoding", "ON"),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_appfwsettings", "useconfigurablesecretkey", "ON"),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_appfwsettings", "sessionlimit", "0"),
					// Attribute not present for check
					//resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_appfwsettings", "malformedreqaction", "[\"none\"]"),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_appfwsettings", "centralizedlearning", "ON"),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_appfwsettings", "proxyport", "9090"),
				),
			},
			{
				Config: testAccAppfwsettings_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwsettingsExist("citrixadc_appfwsettings.tf_appfwsettings", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_appfwsettings", "defaultprofile", "APPFW_BYPASS"),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_appfwsettings", "undefaction", "APPFW_BLOCK"),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_appfwsettings", "sessiontimeout", "900"),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_appfwsettings", "learnratelimit", "400"),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_appfwsettings", "sessionlifetime", "0"),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_appfwsettings", "sessioncookiename", "citrix_ns_id"),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_appfwsettings", "importsizelimit", "134217728"),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_appfwsettings", "signatureautoupdate", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_appfwsettings", "signatureurl", "https://s3.amazonaws.com/NSAppFwSignatures/SignaturesMapping.xml"),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_appfwsettings", "cookiepostencryptprefix", "ENC"),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_appfwsettings", "geolocationlogging", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_appfwsettings", "ceflogging", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_appfwsettings", "entitydecoding", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_appfwsettings", "useconfigurablesecretkey", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_appfwsettings", "sessionlimit", "100000"),
					// Attribute not present for check
					//resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_appfwsettings", "malformedreqaction", "[\"block\", \"log\", \"stats\"]"),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_appfwsettings", "centralizedlearning", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_appfwsettings", "proxyport", "8080"),
				),
			},
		},
	})
}

func testAccCheckAppfwsettingsExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No appfwsettings name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(service.Appfwsettings.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("appfwsettings %s not found", n)
		}

		return nil
	}
}
