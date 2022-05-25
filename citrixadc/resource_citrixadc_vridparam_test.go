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

const testAccVridparam_add = `

	resource "citrixadc_vridparam" "tf_vridparam" {
		sendtomaster  = "ENABLED"
		hellointerval = 400
		deadinterval  = 4
	}
`
const testAccVridparam_update = `

	resource "citrixadc_vridparam" "tf_vridparam" {
		sendtomaster  = "DISABLED"
		hellointerval = 1000
		deadinterval  = 3
	}
`

func TestAccVridparam_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckVridparamDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccVridparam_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVridparamExist("citrixadc_vridparam.tf_vridparam", nil),
					resource.TestCheckResourceAttr("citrixadc_vridparam.tf_vridparam", "sendtomaster", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_vridparam.tf_vridparam", "hellointerval", "400"),
					resource.TestCheckResourceAttr("citrixadc_vridparam.tf_vridparam", "deadinterval", "4"),
				),
			},
			resource.TestStep{
				Config: testAccVridparam_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVridparamExist("citrixadc_vridparam.tf_vridparam", nil),
					resource.TestCheckResourceAttr("citrixadc_vridparam.tf_vridparam", "sendtomaster", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_vridparam.tf_vridparam", "hellointerval", "1000"),
					resource.TestCheckResourceAttr("citrixadc_vridparam.tf_vridparam", "deadinterval", "3"),
				),
			},
		},
	})
}

func testAccCheckVridparamExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No vridparam name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(service.Vridparam.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("vridparam %s not found", n)
		}

		return nil
	}
}

func testAccCheckVridparamDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_vridparam" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Vridparam.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("vridparam %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
