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

const testAccNspbr_basic = `


	resource "citrixadc_nspbr" "tf_nspbr" {
		name = "my_nspbr"
		action = "DENY"
	}
`
const testAccNspbr_update = `


	resource "citrixadc_nspbr" "tf_nspbr" {
		name = "my_nspbr"
		action = "ALLOW"
		nexthop = "true"
		nexthopval = "10.222.74.128"
	}
`

func TestAccNspbr_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNspbrDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNspbr_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNspbrExist("citrixadc_nspbr.tf_nspbr", nil),
					resource.TestCheckResourceAttr("citrixadc_nspbr.tf_nspbr", "name", "my_nspbr"),
					resource.TestCheckResourceAttr("citrixadc_nspbr.tf_nspbr", "action", "DENY"),
				),
			},
			// Commenting out update test, because this requires valid ip in the same subnet as nexthop
			// {
			// 	Config: testAccNspbr_update,
			// 	Check: resource.ComposeTestCheckFunc(
			// 		testAccCheckNspbrExist("citrixadc_nspbr.tf_nspbr", nil),
			// 		resource.TestCheckResourceAttr("citrixadc_nspbr.tf_nspbr", "name", "my_nspbr"),
			// 		resource.TestCheckResourceAttr("citrixadc_nspbr.tf_nspbr", "action", "ALLOW"),
			// 		resource.TestCheckResourceAttr("citrixadc_nspbr.tf_nspbr", "nexthopval", "10.222.74.128"),
			// 	),
			// },
		},
	})
}

func testAccCheckNspbrExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No nspbr name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(service.Nspbr.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("nspbr %s not found", n)
		}

		return nil
	}
}

func testAccCheckNspbrDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_nspbr" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Nspbr.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("nspbr %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
