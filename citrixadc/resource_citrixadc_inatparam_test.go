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

const testAccInatparam_basic = `

	resource "citrixadc_inatparam" "tf_inatparam" {
		nat46ignoretos    = "NO"
		nat46zerochecksum = "ENABLED"
		nat46v6mtu        = "1400"
	}
`
const testAccInatparam_update = `

	resource "citrixadc_inatparam" "tf_inatparam" {
		nat46ignoretos    = "YES"
		nat46zerochecksum = "DISABLED"
		nat46v6mtu        = "1300"
	}
`

func TestAccInatparam_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccInatparam_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckInatparamExist("citrixadc_inatparam.tf_inatparam", nil),
					resource.TestCheckResourceAttr("citrixadc_inatparam.tf_inatparam", "nat46ignoretos", "NO"),
					resource.TestCheckResourceAttr("citrixadc_inatparam.tf_inatparam", "nat46zerochecksum", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_inatparam.tf_inatparam", "nat46v6mtu", "1400"),
				),
			},
			resource.TestStep{
				Config: testAccInatparam_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckInatparamExist("citrixadc_inatparam.tf_inatparam", nil),
					resource.TestCheckResourceAttr("citrixadc_inatparam.tf_inatparam", "nat46ignoretos", "YES"),
					resource.TestCheckResourceAttr("citrixadc_inatparam.tf_inatparam", "nat46zerochecksum", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_inatparam.tf_inatparam", "nat46v6mtu", "1300"),
				),
			},
		},
	})
}

func testAccCheckInatparamExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No inatparam name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(service.Inatparam.Type(), rs.Primary.Attributes["td"])

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("inatparam %s not found", n)
		}

		return nil
	}
}