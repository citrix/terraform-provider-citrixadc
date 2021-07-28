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

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccBotprofile_basic(t *testing.T) {
	if adcTestbed != "STANDALONE" {
		t.Skipf("ADC testbed is %s. Expected STANDALONE.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckBotprofileDestroy,
		Steps: []resource.TestStep{
			// create botprofile
			resource.TestStep{
				Config: testAccBotprofile_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBotprofileExist("citrixadc_botprofile.foo", nil),
					testAccCheckUserAgent(),
				),
			},
			// update botprofile comment
			resource.TestStep{
				Config: testAccBotprofile_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBotprofileExist("citrixadc_botprofile.foo", nil),
					resource.TestCheckResourceAttr("citrixadc_botprofile.foo", "comment", "Botprofile comment 2"),
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

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource("botprofile", rs.Primary.ID)

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
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_Botprofile" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource("botprofile", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("Botprofile %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccBotprofile_basic = `
resource "citrixadc_botprofile" "foo" {
	name = "test_Botprofile"
	comment = "Botprofile comment 1"
	bot_enable_white_list = "ON"
	devicefingerprint = "ON"
}
`

const testAccBotprofile_update = `
resource "citrixadc_botprofile" "foo" {
	name = "test_Botprofile"
	comment = "Botprofile comment 2"
}
`
