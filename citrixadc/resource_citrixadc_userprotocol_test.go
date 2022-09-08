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
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"testing"
)

const testAccUserprotocol_basic = `

	resource "citrixadc_userprotocol" "tf_userprotocol" {
		name      = "my_userprotocol"
		transport = "TCP"
		extension = "my_extension"
		comment   = "my_comment"
	} 
`
const testAccUserprotocol_update = `

	resource "citrixadc_userprotocol" "tf_userprotocol" {
		name      = "my_userprotocol"
		transport = "SSL"
		extension = "my_extension_mqtt"
		comment   = "my_new_comment"
	} 
`
func TestAccUserprotocol_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckUserprotocolDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccUserprotocol_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUserprotocolExist("citrixadc_userprotocol.tf_userprotocol", nil),
					resource.TestCheckResourceAttr("citrixadc_userprotocol.tf_userprotocol", "name", "my_userprotocol"),
					resource.TestCheckResourceAttr("citrixadc_userprotocol.tf_userprotocol", "transport", "TCP"),
					resource.TestCheckResourceAttr("citrixadc_userprotocol.tf_userprotocol", "extension", "my_extension"),
					resource.TestCheckResourceAttr("citrixadc_userprotocol.tf_userprotocol", "comment", "my_comment"),
				),
			},
			resource.TestStep{
				Config: testAccUserprotocol_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUserprotocolExist("citrixadc_userprotocol.tf_userprotocol", nil),
					resource.TestCheckResourceAttr("citrixadc_userprotocol.tf_userprotocol", "name", "my_userprotocol"),
					resource.TestCheckResourceAttr("citrixadc_userprotocol.tf_userprotocol", "transport", "SSL"),
					resource.TestCheckResourceAttr("citrixadc_userprotocol.tf_userprotocol", "extension", "my_extension_mqtt"),
					resource.TestCheckResourceAttr("citrixadc_userprotocol.tf_userprotocol", "comment", "my_new_comment"),
				),
			},
		},
	})
}

func testAccCheckUserprotocolExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No userprotocol name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource("userprotocol", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("userprotocol %s not found", n)
		}

		return nil
	}
}

func testAccCheckUserprotocolDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_userprotocol" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource("userprotocol", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("userprotocol %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
