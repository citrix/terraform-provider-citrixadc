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

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

const testAccResponderparam_basic = `

	resource "citrixadc_responderparam" "tf_responderparam" {
		timeout = 5
		undefaction = "RESET"
	}
`

const testAccResponderparam_basic_update = `

	resource "citrixadc_responderparam" "tf_responderparam" {
		timeout = 6
		undefaction = "DROP"
	}
`

func TestAccResponderparam_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckResponderparamDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccResponderparam_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckResponderparamExist("citrixadc_responderparam.tf_responderparam", nil),
					resource.TestCheckResourceAttr("citrixadc_responderparam.tf_responderparam", "timeout", "5"),
					resource.TestCheckResourceAttr("citrixadc_responderparam.tf_responderparam", "undefaction", "RESET"),
				),
			},
			{
				Config: testAccResponderparam_basic_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckResponderparamExist("citrixadc_responderparam.tf_responderparam", nil),
					resource.TestCheckResourceAttr("citrixadc_responderparam.tf_responderparam", "timeout", "6"),
					resource.TestCheckResourceAttr("citrixadc_responderparam.tf_responderparam", "undefaction", "DROP"),
				),
			},
		},
	})
}

func testAccCheckResponderparamExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No responderparam name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(service.Responderparam.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("responderparam %s not found", n)
		}

		return nil
	}
}

func testAccCheckResponderparamDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_responderparam" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Responderparam.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("responderparam %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
