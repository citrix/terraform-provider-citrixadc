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

const testAccNsspparams_add = `

	resource "citrixadc_nsspparams" "tf_nsspparams" {
		basethreshold = 400
		throttle      = "Aggressive"
	}
`
const testAccNsspparams_update = `

	resource "citrixadc_nsspparams" "tf_nsspparams" {
		basethreshold = 200
		throttle      = "Normal"
	}
`

func TestAccNsspparams_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNsspparamsDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNsspparams_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsspparamsExist("citrixadc_nsspparams.tf_nsspparams", nil),
					resource.TestCheckResourceAttr("citrixadc_nsspparams.tf_nsspparams", "basethreshold", "400"),
					resource.TestCheckResourceAttr("citrixadc_nsspparams.tf_nsspparams", "throttle", "Aggressive"),
				),
			},
			resource.TestStep{
				Config: testAccNsspparams_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsspparamsExist("citrixadc_nsspparams.tf_nsspparams", nil),
					resource.TestCheckResourceAttr("citrixadc_nsspparams.tf_nsspparams", "basethreshold", "200"),
					resource.TestCheckResourceAttr("citrixadc_nsspparams.tf_nsspparams", "throttle", "Normal"),
				),
			},
		},
	})
}

func testAccCheckNsspparamsExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No nsspparams name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(service.Nsspparams.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("nsspparams %s not found", n)
		}

		return nil
	}
}

func testAccCheckNsspparamsDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_nsspparams" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Nsspparams.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("nsspparams %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
