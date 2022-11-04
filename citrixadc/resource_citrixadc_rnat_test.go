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
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"testing"
)

const testAccRnat_add = `

	resource "citrixadc_rnat" "tfrnat" {
		name             = "tfrnat"
		network          = "10.2.2.0"
		netmask          = "255.255.255.255"
		useproxyport     = "ENABLED"
		srcippersistency = "DISABLED"
		connfailover     = "DISABLED"
	}
`
const testAccRnat_update = `

	resource "citrixadc_rnat" "tfrnat" {
		name             = "tfrnat"
		network          = "10.2.2.0"
		netmask          = "255.255.255.255"
		useproxyport     = "DISABLED"
		srcippersistency = "DISABLED"
		connfailover     = "DISABLED"
	}
`

func TestAccRnat_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckRnatDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccRnat_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRnatExist("citrixadc_rnat.tfrnat", nil),
					resource.TestCheckResourceAttr("citrixadc_rnat.tfrnat","name", "tfrnat"),
					resource.TestCheckResourceAttr("citrixadc_rnat.tfrnat","useproxyport", "ENABLED"),
				),
			},
			resource.TestStep{
				Config: testAccRnat_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRnatExist("citrixadc_rnat.tfrnat", nil),
					resource.TestCheckResourceAttr("citrixadc_rnat.tfrnat","name", "tfrnat"),
					resource.TestCheckResourceAttr("citrixadc_rnat.tfrnat","useproxyport", "DISABLED"),
				),
			},
		},
	})
}

func testAccCheckRnatExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No rnat name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(service.Rnat.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("rnat %s not found", n)
		}

		return nil
	}
}

func testAccCheckRnatDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_rnat" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Rnat.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("rnat %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
