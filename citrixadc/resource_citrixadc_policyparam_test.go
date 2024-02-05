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
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

const testAccPolicyparam_basic = `
	resource "citrixadc_policyparam" "tf_policyparam" {
		timeout = 5
	}
`

const testAccPolicyparam_basic_update = `
	resource "citrixadc_policyparam" "tf_policyparam" {
		timeout = 6
	}
`

func TestAccPolicyparam_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		// policyparam resource does not have DELETE operation
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				Config: testAccPolicyparam_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPolicyparamExist("citrixadc_policyparam.tf_policyparam", nil),
					resource.TestCheckResourceAttr("citrixadc_policyparam.tf_policyparam", "timeout", "5"),
				),
			},
			{
				Config: testAccPolicyparam_basic_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPolicyparamExist("citrixadc_policyparam.tf_policyparam", nil),
					resource.TestCheckResourceAttr("citrixadc_policyparam.tf_policyparam", "timeout", "6"),
				),
			},
		},
	})
}

func testAccCheckPolicyparamExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No policyparam name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource("policyparam", "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("policyparam %s not found", n)
		}

		return nil
	}
}
