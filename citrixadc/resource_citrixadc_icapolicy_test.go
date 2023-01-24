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
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"testing"
)

const testAccIcapolicy_basic = `


resource "citrixadc_icapolicy" "tf_icapolicy" {
	name   = "my_ica_policy"
	rule   = true
	action = "my_ica_action"
  }
`

const testAccIcapolicy_update = `


resource "citrixadc_icapolicy" "tf_icapolicy" {
	name   = "my_ica_policy"
	rule   = false
	action = "my_ica_action2"
  }
`

func TestAccIcapolicy_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIcapolicyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccIcapolicy_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIcapolicyExist("citrixadc_icapolicy.tf_icapolicy", nil),
					resource.TestCheckResourceAttr("citrixadc_icapolicy.tf_icapolicy", "name", "my_ica_policy"),
					resource.TestCheckResourceAttr("citrixadc_icapolicy.tf_icapolicy", "rule", "true"),
					resource.TestCheckResourceAttr("citrixadc_icapolicy.tf_icapolicy", "action", "my_ica_action"),
				),
			},
			resource.TestStep{
				Config: testAccIcapolicy_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIcapolicyExist("citrixadc_icapolicy.tf_icapolicy", nil),
					resource.TestCheckResourceAttr("citrixadc_icapolicy.tf_icapolicy", "name", "my_ica_policy"),
					resource.TestCheckResourceAttr("citrixadc_icapolicy.tf_icapolicy", "rule", "false"),
					resource.TestCheckResourceAttr("citrixadc_icapolicy.tf_icapolicy", "action", "my_ica_action2"),
				),
			},
		},
	})
}

func testAccCheckIcapolicyExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No icapolicy name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource("icapolicy", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("icapolicy %s not found", n)
		}

		return nil
	}
}

func testAccCheckIcapolicyDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_icapolicy" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource("icapolicy", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("icapolicy %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
