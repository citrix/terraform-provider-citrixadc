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

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testAccBotsettings_basic = `
	resource "citrixadc_botsettings" "default" {
		sessiontimeout= "900"
		proxyport = "8080"
		sessioncookiename = "citrix_bot_id"
		dfprequestlimit = "1"
		signatureautoupdate = "ON"
		trapurlautogenerate = "OFF"
		trapurlinterval = "3600"
		trapurllength = "32"
		proxyusername = "testuser"
	}
`
const testAccBotsettings_basic_update = `
	resource "citrixadc_botsettings" "default" {
		sessiontimeout= "950"
		proxyport = "80"
		sessioncookiename = "citrixbotid"
		dfprequestlimit = "3"
		signatureautoupdate = "ON"
		trapurlautogenerate = "ON"
		trapurlinterval = "3800"
		trapurllength = "33"
		proxyusername = "testuser1"
	}
`

func TestAccBotsettings_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		// botsettings resource do not have DELETE operation
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				Config: testAccBotsettings_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBotsettingsExist("citrixadc_botsettings.default", nil),
					resource.TestCheckResourceAttr("citrixadc_botsettings.default", "sessiontimeout", "900"),
					resource.TestCheckResourceAttr("citrixadc_botsettings.default", "proxyport", "8080"),
					resource.TestCheckResourceAttr("citrixadc_botsettings.default", "sessioncookiename", "citrix_bot_id"),
					resource.TestCheckResourceAttr("citrixadc_botsettings.default", "dfprequestlimit", "1"),
					resource.TestCheckResourceAttr("citrixadc_botsettings.default", "signatureautoupdate", "ON"),
					resource.TestCheckResourceAttr("citrixadc_botsettings.default", "trapurlautogenerate", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_botsettings.default", "trapurlinterval", "3600"),
					resource.TestCheckResourceAttr("citrixadc_botsettings.default", "trapurllength", "32"),
					resource.TestCheckResourceAttr("citrixadc_botsettings.default", "proxyusername", "testuser"),
				),
			},
			{
				Config: testAccBotsettings_basic_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBotsettingsExist("citrixadc_botsettings.default", nil),
					resource.TestCheckResourceAttr("citrixadc_botsettings.default", "sessiontimeout", "950"),
					resource.TestCheckResourceAttr("citrixadc_botsettings.default", "proxyport", "80"),
					resource.TestCheckResourceAttr("citrixadc_botsettings.default", "sessioncookiename", "citrixbotid"),
					resource.TestCheckResourceAttr("citrixadc_botsettings.default", "dfprequestlimit", "3"),
					resource.TestCheckResourceAttr("citrixadc_botsettings.default", "signatureautoupdate", "ON"),
					resource.TestCheckResourceAttr("citrixadc_botsettings.default", "trapurlautogenerate", "ON"),
					resource.TestCheckResourceAttr("citrixadc_botsettings.default", "trapurlinterval", "3800"),
					resource.TestCheckResourceAttr("citrixadc_botsettings.default", "trapurllength", "33"),
					resource.TestCheckResourceAttr("citrixadc_botsettings.default", "proxyusername", "testuser1"),
				),
			},
		},
	})
}

func testAccCheckBotsettingsExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lb vserver name is set")
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
		data, err := client.FindResource("botsettings", "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("BOT Settings  %s not found", n)
		}

		return nil
	}
}
