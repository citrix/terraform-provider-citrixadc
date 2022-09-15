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

const testAccReputationsettings_basic = `

resource "citrixadc_reputationsettings" "tf_reputationsettings" {
	proxyserver = "my_proxyserver"
	proxyport   = 3500
  }
  
`
const testAccReputationsettings_update = `

resource "citrixadc_reputationsettings" "tf_reputationsettings" {
	proxyserver = "my_proxyserver"
	proxyport   = 3600
  }
  
`

func TestAccReputationsettings_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccReputationsettings_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckReputationsettingsExist("citrixadc_reputationsettings.tf_reputationsettings", nil),
					resource.TestCheckResourceAttr("citrixadc_reputationsettings.tf_reputationsettings", "proxyserver", "my_proxyserver"),
					resource.TestCheckResourceAttr("citrixadc_reputationsettings.tf_reputationsettings", "proxyport", "3500"),
				),
			},
			resource.TestStep{
				Config: testAccReputationsettings_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckReputationsettingsExist("citrixadc_reputationsettings.tf_reputationsettings", nil),
					resource.TestCheckResourceAttr("citrixadc_reputationsettings.tf_reputationsettings", "proxyserver", "my_proxyserver"),
					resource.TestCheckResourceAttr("citrixadc_reputationsettings.tf_reputationsettings", "proxyport", "3600"),
				),
			},
		},
	})
}

func testAccCheckReputationsettingsExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No reputationsettings name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource("reputationsettings", "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("reputationsettings %s not found", n)
		}

		return nil
	}
}