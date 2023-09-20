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

const testAccTmtrafficaction_basic = `


	resource "citrixadc_tmtrafficaction" "tf_tmtrafficaction" {
		name             = "my_traffic_action"
		apptimeout       = 5
		sso              = "OFF"
		persistentcookie = "ON"
	}
`

const testAccTmtrafficaction_update = `


	resource "citrixadc_tmtrafficaction" "tf_tmtrafficaction" {
		name             = "my_traffic_action"
		apptimeout       = 10
		sso              = "ON"
		persistentcookie = "OFF"
	}
`

func TestAccTmtrafficaction_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckTmtrafficactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTmtrafficaction_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTmtrafficactionExist("citrixadc_tmtrafficaction.tf_tmtrafficaction", nil),
					resource.TestCheckResourceAttr("citrixadc_tmtrafficaction.tf_tmtrafficaction", "name", "my_traffic_action"),
					resource.TestCheckResourceAttr("citrixadc_tmtrafficaction.tf_tmtrafficaction", "apptimeout", "5"),
					resource.TestCheckResourceAttr("citrixadc_tmtrafficaction.tf_tmtrafficaction", "sso", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_tmtrafficaction.tf_tmtrafficaction", "persistentcookie", "ON"),
				),
			},
			{
				Config: testAccTmtrafficaction_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTmtrafficactionExist("citrixadc_tmtrafficaction.tf_tmtrafficaction", nil),
					resource.TestCheckResourceAttr("citrixadc_tmtrafficaction.tf_tmtrafficaction", "name", "my_traffic_action"),
					resource.TestCheckResourceAttr("citrixadc_tmtrafficaction.tf_tmtrafficaction", "apptimeout", "10"),
					resource.TestCheckResourceAttr("citrixadc_tmtrafficaction.tf_tmtrafficaction", "sso", "ON"),
					resource.TestCheckResourceAttr("citrixadc_tmtrafficaction.tf_tmtrafficaction", "persistentcookie", "OFF"),
				),
			},
		},
	})
}

func testAccCheckTmtrafficactionExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No tmtrafficaction name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(service.Tmtrafficaction.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("tmtrafficaction %s not found", n)
		}

		return nil
	}
}

func testAccCheckTmtrafficactionDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_tmtrafficaction" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Tmtrafficaction.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("tmtrafficaction %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
