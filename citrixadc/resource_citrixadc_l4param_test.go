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

const testAccL4param_add = `
	resource "citrixadc_l4param" "tf_l4param" {
		l2connmethod = "Channel"
		l4switch     = "ENABLED"
	}
`
const testAccL4param_update = `
	resource "citrixadc_l4param" "tf_l4param" {
		l2connmethod = "MacVlanChannel"
		l4switch     = "DISABLED"
	}
`

func TestAccL4param_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckL4paramDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccL4param_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckL4paramExist("citrixadc_l4param.tf_l4param", nil),
					resource.TestCheckResourceAttr("citrixadc_l4param.tf_l4param", "l2connmethod", "Channel"),
					resource.TestCheckResourceAttr("citrixadc_l4param.tf_l4param", "l4switch", "ENABLED"),
				),
			},
			resource.TestStep{
				Config: testAccL4param_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckL4paramExist("citrixadc_l4param.tf_l4param", nil),
					resource.TestCheckResourceAttr("citrixadc_l4param.tf_l4param", "l2connmethod", "MacVlanChannel"),
					resource.TestCheckResourceAttr("citrixadc_l4param.tf_l4param", "l4switch", "DISABLED"),
				),
			},
		},
	})
}

func testAccCheckL4paramExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No l4param name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(service.L4param.Type(),"")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("l4param %s not found", n)
		}

		return nil
	}
}

func testAccCheckL4paramDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_l4param" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.L4param.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("l4param %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
