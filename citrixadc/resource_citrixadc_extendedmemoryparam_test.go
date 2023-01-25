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

const testAccExtendedmemoryparam_add = `
	resource "citrixadc_extendedmemoryparam" "tf_extendedmemoryparam" {
		memlimit = 512
	}
`
const testAccExtendedmemoryparam_update = `
	resource "citrixadc_extendedmemoryparam" "tf_extendedmemoryparam" {
		memlimit = 256
	}
`

func TestAccExtendedmemoryparam_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccExtendedmemoryparam_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckExtendedmemoryparamExist("citrixadc_extendedmemoryparam.tf_extendedmemoryparam", nil),
					resource.TestCheckResourceAttr("citrixadc_extendedmemoryparam.tf_extendedmemoryparam", "memlimit", "512"),
				),
			},
			resource.TestStep{
				Config: testAccExtendedmemoryparam_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckExtendedmemoryparamExist("citrixadc_extendedmemoryparam.tf_extendedmemoryparam", nil),
					resource.TestCheckResourceAttr("citrixadc_extendedmemoryparam.tf_extendedmemoryparam", "memlimit", "256"),
				),
			},
		},
	})
}

func testAccCheckExtendedmemoryparamExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No extendedmemoryparam name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(service.Extendedmemoryparam.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("extendedmemoryparam %s not found", n)
		}

		return nil
	}
}
