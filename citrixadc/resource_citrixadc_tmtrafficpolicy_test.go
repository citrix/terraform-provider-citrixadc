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

const testAccTmtrafficpolicy_basic = `


resource "citrixadc_tmtrafficpolicy" "tf_tmtrafficpolicy" {
	name   = "my_tmtraffic_policy"
	rule   = "true"
	action = "my_tmtraffic_action"
  }
`
const testAccTmtrafficpolicy_update = `


resource "citrixadc_tmtrafficpolicy" "tf_tmtrafficpolicy" {
	name   = "my_tmtraffic_policy"
	rule   = "false"
	action = "my_tmtraffic_action2"
  }
`

func TestAccTmtrafficpolicy_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckTmtrafficpolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTmtrafficpolicy_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTmtrafficpolicyExist("citrixadc_tmtrafficpolicy.tf_tmtrafficpolicy", nil),
					resource.TestCheckResourceAttr("citrixadc_tmtrafficpolicy.tf_tmtrafficpolicy", "name", "my_tmtraffic_policy"),
					resource.TestCheckResourceAttr("citrixadc_tmtrafficpolicy.tf_tmtrafficpolicy", "action", "my_tmtraffic_action"),
					resource.TestCheckResourceAttr("citrixadc_tmtrafficpolicy.tf_tmtrafficpolicy", "rule", "true"),
				),
			},
			{
				Config: testAccTmtrafficpolicy_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTmtrafficpolicyExist("citrixadc_tmtrafficpolicy.tf_tmtrafficpolicy", nil),
					resource.TestCheckResourceAttr("citrixadc_tmtrafficpolicy.tf_tmtrafficpolicy", "name", "my_tmtraffic_policy"),
					resource.TestCheckResourceAttr("citrixadc_tmtrafficpolicy.tf_tmtrafficpolicy", "action", "my_tmtraffic_action2"),
					resource.TestCheckResourceAttr("citrixadc_tmtrafficpolicy.tf_tmtrafficpolicy", "rule", "false"),
				),
			},
		},
	})
}

func testAccCheckTmtrafficpolicyExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No tmtrafficpolicy name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(service.Tmtrafficpolicy.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("tmtrafficpolicy %s not found", n)
		}

		return nil
	}
}

func testAccCheckTmtrafficpolicyDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_tmtrafficpolicy" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Tmtrafficpolicy.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("tmtrafficpolicy %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
