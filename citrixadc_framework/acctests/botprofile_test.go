/*
Copyright 2021 Citrix Systems, Inc

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

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccBotprofile_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckBotprofileDestroy,
		Steps: []resource.TestStep{
			// create botprofile
			{
				Config: testAccBotprofile_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBotprofileExist("citrixadc_botprofile.tf_botprofile", nil),
					testAccCheckUserAgent(),
					resource.TestCheckResourceAttr("citrixadc_botprofile.tf_botprofile", "addcookieflags", "secure"),
					resource.TestCheckResourceAttr("citrixadc_botprofile.tf_botprofile", "dfprequestlimit", "25"),
					resource.TestCheckResourceAttr("citrixadc_botprofile.tf_botprofile", "headlessbrowserdetection", "ON"),
					resource.TestCheckResourceAttr("citrixadc_botprofile.tf_botprofile", "sessioncookiename", "testCookie"),
					resource.TestCheckResourceAttr("citrixadc_botprofile.tf_botprofile", "sessiontimeout", "1200"),
					resource.TestCheckResourceAttr("citrixadc_botprofile.tf_botprofile", "spoofedreqaction.0", "LOG"),
					resource.TestCheckResourceAttr("citrixadc_botprofile.tf_botprofile", "verboseloglevel", "HTTP_FULL_HEADER"),
				),
			},
			// update botprofile actions
			{
				Config: testAccBotprofile_update_actions,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBotprofileExist("citrixadc_botprofile.tf_botprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_botprofile.tf_botprofile", "devicefingerprintaction.0", "LOG"),
					resource.TestCheckResourceAttr("citrixadc_botprofile.tf_botprofile", "devicefingerprintaction.1", "DROP"),
					resource.TestCheckResourceAttr("citrixadc_botprofile.tf_botprofile", "trapaction.0", "LOG"),
					resource.TestCheckResourceAttr("citrixadc_botprofile.tf_botprofile", "trapaction.1", "DROP"),
					testAccCheckUserAgent(),
					resource.TestCheckResourceAttr("citrixadc_botprofile.tf_botprofile", "addcookieflags", "secure"),
					resource.TestCheckResourceAttr("citrixadc_botprofile.tf_botprofile", "dfprequestlimit", "50"),
					resource.TestCheckResourceAttr("citrixadc_botprofile.tf_botprofile", "headlessbrowserdetection", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_botprofile.tf_botprofile", "sessioncookiename", "testCookie1"),
					resource.TestCheckResourceAttr("citrixadc_botprofile.tf_botprofile", "sessiontimeout", "1800"),
					resource.TestCheckResourceAttr("citrixadc_botprofile.tf_botprofile", "spoofedreqaction.0", "DROP"),
					resource.TestCheckResourceAttr("citrixadc_botprofile.tf_botprofile", "verboseloglevel", "NONE"),
				),
			},
			// update botprofile properties
			{
				Config: testAccBotprofile_update_properties,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBotprofileExist("citrixadc_botprofile.tf_botprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_botprofile.tf_botprofile", "errorurl", "http://www.citrix.com/products/citrix-adc/"),
					resource.TestCheckResourceAttr("citrixadc_botprofile.tf_botprofile", "trapurl", "/http://www.citrix.com/products/citrix-adc/"),
					resource.TestCheckResourceAttr("citrixadc_botprofile.tf_botprofile", "comment", "tf_botprofile comment 1"),
					resource.TestCheckResourceAttr("citrixadc_botprofile.tf_botprofile", "bot_enable_white_list", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_botprofile.tf_botprofile", "bot_enable_black_list", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_botprofile.tf_botprofile", "bot_enable_rate_limit", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_botprofile.tf_botprofile", "devicefingerprint", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_botprofile.tf_botprofile", "bot_enable_ip_reputation", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_botprofile.tf_botprofile", "trap", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_botprofile.tf_botprofile", "bot_enable_tps", "OFF"),
					testAccCheckUserAgent(),
				),
			},
		},
	})
}

func testAccCheckBotprofileExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Botprofile name is set")
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
		data, err := client.FindResource("botprofile", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("Botprofile %s not found", n)
		}

		return nil
	}
}

func testAccCheckBotprofileDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_Botprofile" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("botprofile", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("Botprofile %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccBotprofile_basic = `
resource "citrixadc_botprofile" "tf_botprofile" {
	name = "tf_botprofile"
	errorurl = "http://www.citrix.com"
	trapurl = "/http://www.citrix.com"
	comment = "tf_botprofile comment"
	bot_enable_white_list = "ON"
	bot_enable_black_list = "ON"
	bot_enable_rate_limit = "ON"
	devicefingerprint = "ON"
	devicefingerprintaction = ["LOG", "RESET"]
	bot_enable_ip_reputation = "ON"
	trap = "ON"
	trapaction = ["LOG", "RESET"]
	bot_enable_tps = "ON"
	addcookieflags	= "secure"
	dfprequestlimit = "25"
	headlessbrowserdetection = "ON"
	sessioncookiename = "testCookie"
	sessiontimeout = "1200"
	spoofedreqaction = ["LOG"]
	verboseloglevel = "HTTP_FULL_HEADER"
}
`

const testAccBotprofile_update_actions = `
resource "citrixadc_botprofile" "tf_botprofile" {
	name = "tf_botprofile"
	devicefingerprintaction = ["LOG", "DROP"]
	trapaction = ["LOG", "DROP"]
	addcookieflags	= "secure"
	dfprequestlimit = "50"
	headlessbrowserdetection = "OFF"
	sessioncookiename = "testCookie1"
	sessiontimeout = "1800"
	spoofedreqaction = ["DROP"]
	verboseloglevel = "NONE"
}
`

const testAccBotprofile_update_properties = `
resource "citrixadc_botprofile" "tf_botprofile" {
	name = "tf_botprofile"
	errorurl = "http://www.citrix.com/products/citrix-adc/"
	trapurl = "/http://www.citrix.com/products/citrix-adc/"
	comment = "tf_botprofile comment 1"
	bot_enable_white_list = "OFF"
	bot_enable_black_list = "OFF"
	bot_enable_rate_limit = "OFF"
	devicefingerprint = "OFF"
	bot_enable_ip_reputation = "OFF"
	trap = "OFF"
	bot_enable_tps = "OFF"
}
`

func TestAccBotprofileDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccBotprofileDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_botprofile.tf_botprofile_ds", "name", "tf_botprofile_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_botprofile.tf_botprofile_ds", "addcookieflags", "secure"),
					resource.TestCheckResourceAttr("data.citrixadc_botprofile.tf_botprofile_ds", "dfprequestlimit", "25"),
					resource.TestCheckResourceAttr("data.citrixadc_botprofile.tf_botprofile_ds", "headlessbrowserdetection", "ON"),
					resource.TestCheckResourceAttr("data.citrixadc_botprofile.tf_botprofile_ds", "sessioncookiename", "dsCookie"),
					resource.TestCheckResourceAttr("data.citrixadc_botprofile.tf_botprofile_ds", "sessiontimeout", "1200"),
					resource.TestCheckResourceAttr("data.citrixadc_botprofile.tf_botprofile_ds", "verboseloglevel", "HTTP_FULL_HEADER"),
					resource.TestCheckResourceAttr("data.citrixadc_botprofile.tf_botprofile_ds", "comment", "DATASOURCE TEST COMMENT"),
				),
			},
		},
	})
}

const testAccBotprofileDataSource_basic = `

resource "citrixadc_botprofile" "tf_botprofile_ds" {
	name = "tf_botprofile_ds"
	errorurl = "http://www.citrix.com"
	trapurl = "/http://www.citrix.com"
	comment = "DATASOURCE TEST COMMENT"
	bot_enable_white_list = "ON"
	bot_enable_black_list = "ON"
	bot_enable_rate_limit = "ON"
	devicefingerprint = "ON"
	bot_enable_ip_reputation = "ON"
	trap = "ON"
	bot_enable_tps = "ON"
	addcookieflags = "secure"
	dfprequestlimit = "25"
	headlessbrowserdetection = "ON"
	sessioncookiename = "dsCookie"
	sessiontimeout = "1200"
	verboseloglevel = "HTTP_FULL_HEADER"
}

data "citrixadc_botprofile" "tf_botprofile_ds" {
	name = citrixadc_botprofile.tf_botprofile_ds.name
	depends_on = [citrixadc_botprofile.tf_botprofile_ds]
}

`
