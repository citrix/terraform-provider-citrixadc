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

const testAccNsdiameter_add = `
	resource "citrixadc_nsdiameter" "tf_nsdiameter" {
		identity               = "citrixadc.com"
		realm                  = "com"
		serverclosepropagation = "OFF"
	}
`
const testAccNsdiameter_update = `
	resource "citrixadc_nsdiameter" "tf_nsdiameter" {
		identity               = "netscaler.com"
		realm                  = "com"
		serverclosepropagation = "ON"
	}
`

func TestAccNsdiameter_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				Config: testAccNsdiameter_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsdiameterExist("citrixadc_nsdiameter.tf_nsdiameter", nil),
					resource.TestCheckResourceAttr("citrixadc_nsdiameter.tf_nsdiameter", "identity", "citrixadc.com"),
					resource.TestCheckResourceAttr("citrixadc_nsdiameter.tf_nsdiameter", "realm", "com"),
					resource.TestCheckResourceAttr("citrixadc_nsdiameter.tf_nsdiameter", "serverclosepropagation", "OFF"),
				),
			},
			{
				Config: testAccNsdiameter_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsdiameterExist("citrixadc_nsdiameter.tf_nsdiameter", nil),
					resource.TestCheckResourceAttr("citrixadc_nsdiameter.tf_nsdiameter", "identity", "netscaler.com"),
					resource.TestCheckResourceAttr("citrixadc_nsdiameter.tf_nsdiameter", "realm", "com"),
					resource.TestCheckResourceAttr("citrixadc_nsdiameter.tf_nsdiameter", "serverclosepropagation", "ON"),
				),
			},
		},
	})
}

func testAccCheckNsdiameterExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No nsdiameter name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(service.Nsdiameter.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("nsdiameter %s not found", n)
		}

		return nil
	}
}

func testAccCheckNsdiameterDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_nsdiameter" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Nsdiameter.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("nsdiameter %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
