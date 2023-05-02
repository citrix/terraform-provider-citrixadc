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

const testAccAaapreauthenticationaction_basic = `

	resource "citrixadc_aaapreauthenticationaction" "tf_aaapreauthenticationaction" {
		name                    = "tf_aaapreauthenticationaction"
		preauthenticationaction = "ALLOW"
		deletefiles             = "/var/tmp/new/hello.txt"
	}
  
`
const testAccAaapreauthenticationaction_update = `

	resource "citrixadc_aaapreauthenticationaction" "tf_aaapreauthenticationaction" {
		name                    = "tf_aaapreauthenticationaction"
		preauthenticationaction = "DENY"
		deletefiles             = "/var/tmp/new/hello.txt"
	}
`

func TestAccAaapreauthenticationaction_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAaapreauthenticationactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAaapreauthenticationaction_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaapreauthenticationactionExist("citrixadc_aaapreauthenticationaction.tf_aaapreauthenticationaction", nil),
					resource.TestCheckResourceAttr("citrixadc_aaapreauthenticationaction.tf_aaapreauthenticationaction", "preauthenticationaction", "ALLOW"),
				),
			},
			{
				Config: testAccAaapreauthenticationaction_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaapreauthenticationactionExist("citrixadc_aaapreauthenticationaction.tf_aaapreauthenticationaction", nil),
					resource.TestCheckResourceAttr("citrixadc_aaapreauthenticationaction.tf_aaapreauthenticationaction", "preauthenticationaction", "DENY"),
				),
			},
		},
	})
}

func testAccCheckAaapreauthenticationactionExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No aaapreauthenticationaction name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(service.Aaapreauthenticationaction.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("aaapreauthenticationaction %s not found", n)
		}

		return nil
	}
}

func testAccCheckAaapreauthenticationactionDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_aaapreauthenticationaction" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Aaapreauthenticationaction.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("aaapreauthenticationaction %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
