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

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

const testAccBotpolicy_add = `
	resource citrixadc_botpolicy tfAcc_botpolicy1 {
		name = "tfAcc_botpolicy1"  
		profilename = "BOT_BYPASS"
		rule  = "true"
		comment = "COMMENT FOR BOTPOLICY"
    }
`
const testAccBotpolicy_update = `
	resource citrixadc_botpolicy tfAcc_botpolicy1 {
		name = "tfAcc_botpolicy1"
		profilename = "BOT_BYPASS"
		rule = "true"
		comment = "CHANGED COMMENT"
	}
`

func TestAccBotpolicy_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckBotpolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccBotpolicy_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBotpolicyExist("citrixadc_botpolicy.tfAcc_botpolicy1", nil),
					resource.TestCheckResourceAttr("citrixadc_botpolicy.tfAcc_botpolicy1", "name", "tfAcc_botpolicy1"),
					resource.TestCheckResourceAttr("citrixadc_botpolicy.tfAcc_botpolicy1", "profilename", "BOT_BYPASS"),
					resource.TestCheckResourceAttr("citrixadc_botpolicy.tfAcc_botpolicy1", "rule", "true"),
					resource.TestCheckResourceAttr("citrixadc_botpolicy.tfAcc_botpolicy1", "comment", "COMMENT FOR BOTPOLICY"),
				),
			},
			{
				Config: testAccBotpolicy_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBotpolicyExist("citrixadc_botpolicy.tfAcc_botpolicy1", nil),
					resource.TestCheckResourceAttr("citrixadc_botpolicy.tfAcc_botpolicy1", "name", "tfAcc_botpolicy1"),
					resource.TestCheckResourceAttr("citrixadc_botpolicy.tfAcc_botpolicy1", "profilename", "BOT_BYPASS"),
					resource.TestCheckResourceAttr("citrixadc_botpolicy.tfAcc_botpolicy1", "rule", "true"),
					resource.TestCheckResourceAttr("citrixadc_botpolicy.tfAcc_botpolicy1", "comment", "CHANGED COMMENT"),
				),
			},
		},
	})
}

func testAccCheckBotpolicyExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No botpolicy name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource("botpolicy", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("BOT policy %s not found", n)
		}

		return nil
	}
}

func testAccCheckBotpolicyDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_botpolicy" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource("botpolicy", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("BOT policy %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
