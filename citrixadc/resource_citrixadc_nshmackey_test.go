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

const testAccNshmackey_add = `

	resource "citrixadc_nshmackey" "tf_nshmackey" {
		name     = "tf_nshmackey"
		digest   = "MD4"
		keyvalue = "AUTO"
		comment  = "Testing"
	}
`
const testAccNshmackey_update = `

	resource "citrixadc_nshmackey" "tf_nshmackey" {
		name     = "tf_nshmackey"
		digest   = "MD2"
		keyvalue = "AUTO"
		comment  = "Testing_sample"
	}
`

func TestAccNshmackey_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNshmackeyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNshmackey_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNshmackeyExist("citrixadc_nshmackey.tf_nshmackey", nil),
					resource.TestCheckResourceAttr("citrixadc_nshmackey.tf_nshmackey", "name", "tf_nshmackey"),
					resource.TestCheckResourceAttr("citrixadc_nshmackey.tf_nshmackey", "digest", "MD4"),
					resource.TestCheckResourceAttr("citrixadc_nshmackey.tf_nshmackey", "comment", "Testing"),
				),
			},
			resource.TestStep{
				Config: testAccNshmackey_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNshmackeyExist("citrixadc_nshmackey.tf_nshmackey", nil),
					resource.TestCheckResourceAttr("citrixadc_nshmackey.tf_nshmackey", "name", "tf_nshmackey"),
					resource.TestCheckResourceAttr("citrixadc_nshmackey.tf_nshmackey", "digest", "MD2"),
					resource.TestCheckResourceAttr("citrixadc_nshmackey.tf_nshmackey", "comment", "Testing_sample"),
				),
			},
		},
	})
}

func testAccCheckNshmackeyExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No nshmackey name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource("nshmackey", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("nshmackey %s not found", n)
		}

		return nil
	}
}

func testAccCheckNshmackeyDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_nshmackey" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource("nshmackey", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("nshmackey %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
