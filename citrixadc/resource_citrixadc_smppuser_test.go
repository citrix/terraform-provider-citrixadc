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

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const testAccSmppuser_basic = `
	resource "citrixadc_smppuser" "user1" {
		username = "demouser"
		password = "demouser123"
	}
`

func TestAccSmppuser_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSmppuserDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccSmppuser_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSmppuserExist("citrixadc_smppuser.user1", nil),
					resource.TestCheckResourceAttr("citrixadc_smppuser.user1", "username", "demouser"),
				),
			},
		},
	})
}

func testAccCheckSmppuserExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No smppuser name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource("smppuser", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("smppuser %s not found", n)
		}

		return nil
	}
}

func testAccCheckSmppuserDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_smppuser" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource("smppuser", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("smppuser %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
