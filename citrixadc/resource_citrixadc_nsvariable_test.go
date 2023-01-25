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
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"testing"
)

const testAccNsvariable_add = `
	resource "citrixadc_nsvariable" "tf_nsvariable" {
		name          = "tf_nsvariable"
		type          = "text(20)"
		scope         = "global"
		iffull        = "undef"
		ifvaluetoobig = "undef"
		ifnovalue     = "undef"
		comment       = "Testing"
	}
`
const testAccNsvariable_update = `
	resource "citrixadc_nsvariable" "tf_nsvariable" {
		name          = "tf_nsvariable"
		type          = "text(20)"
		scope         = "global"
		iffull        = "lru"
		ifvaluetoobig = "truncate"
		ifnovalue     = "init"
		comment       = "Testing"
	}
`

func TestAccNsvariable_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNsvariableDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNsvariable_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsvariableExist("citrixadc_nsvariable.tf_nsvariable", nil),
					resource.TestCheckResourceAttr("citrixadc_nsvariable.tf_nsvariable", "name", "tf_nsvariable"),
					resource.TestCheckResourceAttr("citrixadc_nsvariable.tf_nsvariable", "iffull", "undef"),
					resource.TestCheckResourceAttr("citrixadc_nsvariable.tf_nsvariable", "ifvaluetoobig", "undef"),
					resource.TestCheckResourceAttr("citrixadc_nsvariable.tf_nsvariable", "ifnovalue", "undef"),
				),
			},
			resource.TestStep{
				Config: testAccNsvariable_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsvariableExist("citrixadc_nsvariable.tf_nsvariable", nil),
					resource.TestCheckResourceAttr("citrixadc_nsvariable.tf_nsvariable", "name", "tf_nsvariable"),
					resource.TestCheckResourceAttr("citrixadc_nsvariable.tf_nsvariable", "iffull", "lru"),
					resource.TestCheckResourceAttr("citrixadc_nsvariable.tf_nsvariable", "ifvaluetoobig", "truncate"),
					resource.TestCheckResourceAttr("citrixadc_nsvariable.tf_nsvariable", "ifnovalue", "init"),
				),
			},
		},
	})
}

func testAccCheckNsvariableExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No nsvariable name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(service.Nsvariable.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("nsvariable %s not found", n)
		}

		return nil
	}
}

func testAccCheckNsvariableDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_nsvariable" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Nsvariable.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("nsvariable %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
