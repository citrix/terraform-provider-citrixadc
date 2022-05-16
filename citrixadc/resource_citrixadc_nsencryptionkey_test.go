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

const testAccNsencryptionkey_add = `
	resource "citrixadc_nsencryptionkey" "tf_encryptionkey" {
		name     = "tf_encryptionkey"
		method   = "AES256"
		keyvalue = "26ea5537b7e0746089476e5658f9327c0b10c3b4778c673a5b38cee182874711"
		padding  = "ON"
		iv       = "c2bf0b2e15c15004d6b14bcdc7e5e365"
		comment  = "Testing"
	}
`

const testAccNsencryptionkey_update = `
	resource "citrixadc_nsencryptionkey" "tf_encryptionkey" {
		name     = "tf_encryptionkey"
		method   = "AES256"
		keyvalue = "26ea5537b7e0746089476e5658f9327c0b10c3b4778c673a5b38cee182874711"
		padding  = "DEFAULT"
		iv       = "c2bf0b2e15c15004d6b14bcdc7e5e123"
		comment  = "Testing_sample"
	}
`

func TestAccNsencryptionkey_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNsencryptionkeyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNsencryptionkey_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsencryptionkeyExist("citrixadc_nsencryptionkey.tf_encryptionkey", nil),
					resource.TestCheckResourceAttr("citrixadc_nsencryptionkey.tf_encryptionkey", "name", "tf_encryptionkey"),
					resource.TestCheckResourceAttr("citrixadc_nsencryptionkey.tf_encryptionkey", "padding", "ON"),
					resource.TestCheckResourceAttr("citrixadc_nsencryptionkey.tf_encryptionkey", "iv", "c2bf0b2e15c15004d6b14bcdc7e5e365"),
					resource.TestCheckResourceAttr("citrixadc_nsencryptionkey.tf_encryptionkey", "comment", "Testing"),
				),
			},
			resource.TestStep{
				Config: testAccNsencryptionkey_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsencryptionkeyExist("citrixadc_nsencryptionkey.tf_encryptionkey", nil),
					resource.TestCheckResourceAttr("citrixadc_nsencryptionkey.tf_encryptionkey", "name", "tf_encryptionkey"),
					resource.TestCheckResourceAttr("citrixadc_nsencryptionkey.tf_encryptionkey", "padding", "DEFAULT"),
					resource.TestCheckResourceAttr("citrixadc_nsencryptionkey.tf_encryptionkey", "iv", "c2bf0b2e15c15004d6b14bcdc7e5e123"),
					resource.TestCheckResourceAttr("citrixadc_nsencryptionkey.tf_encryptionkey", "comment", "Testing_sample"),
				),
			},
		},
	})
}

func testAccCheckNsencryptionkeyExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No nsencryptionkey name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource("nsencryptionkey", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("nsencryptionkey %s not found", n)
		}

		return nil
	}
}

func testAccCheckNsencryptionkeyDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_nsencryptionkey" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource("nsencryptionkey", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("nsencryptionkey %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
