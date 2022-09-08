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

const testAccTmsessionaction_basic = `


	resource "citrixadc_tmsessionaction" "tf_tmsessionaction" {
		name                       = "my_tmsession_action"
		sesstimeout                = 10
		defaultauthorizationaction = "ALLOW"
		sso                        = "OFF"
	}
`


const testAccTmsessionaction_update = `


	resource "citrixadc_tmsessionaction" "tf_tmsessionaction" {
		name                       = "my_tmsession_action"
		sesstimeout                = 20
		defaultauthorizationaction = "DENY"
		sso                        = "OFF"
	}
`

func TestAccTmsessionaction_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckTmsessionactionDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccTmsessionaction_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTmsessionactionExist("citrixadc_tmsessionaction.tf_tmsessionaction", nil),
					resource.TestCheckResourceAttr("citrixadc_tmsessionaction.tf_tmsessionaction", "name", "my_tmsession_action"),
					resource.TestCheckResourceAttr("citrixadc_tmsessionaction.tf_tmsessionaction", "sesstimeout", "10"),
					resource.TestCheckResourceAttr("citrixadc_tmsessionaction.tf_tmsessionaction", "defaultauthorizationaction", "ALLOW"),
					resource.TestCheckResourceAttr("citrixadc_tmsessionaction.tf_tmsessionaction", "sso", "OFF"),
				),
			},
			resource.TestStep{
				Config: testAccTmsessionaction_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTmsessionactionExist("citrixadc_tmsessionaction.tf_tmsessionaction", nil),
					resource.TestCheckResourceAttr("citrixadc_tmsessionaction.tf_tmsessionaction", "name", "my_tmsession_action"),
					resource.TestCheckResourceAttr("citrixadc_tmsessionaction.tf_tmsessionaction", "sesstimeout", "20"),
					resource.TestCheckResourceAttr("citrixadc_tmsessionaction.tf_tmsessionaction", "defaultauthorizationaction", "DENY"),
					resource.TestCheckResourceAttr("citrixadc_tmsessionaction.tf_tmsessionaction", "sso", "OFF"),
				),
			},
		},
	})
}

func testAccCheckTmsessionactionExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No tmsessionaction name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(service.Tmsessionaction.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("tmsessionaction %s not found", n)
		}

		return nil
	}
}

func testAccCheckTmsessionactionDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_tmsessionaction" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Tmsessionaction.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("tmsessionaction %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
