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

const testAccGslbparameter_basic = `

resource "citrixadc_gslbparameter" "tf_gslbparameter" {
	ldnsentrytimeout = 50
	rtttolerance     = 6
	ldnsmask         = "255.255.255.255"
  }
`

const testAccGslbparameter_update = `

resource "citrixadc_gslbparameter" "tf_gslbparameter" {
	ldnsentrytimeout = 70
	rtttolerance     = 8
	ldnsmask         = "255.255.255.254"
  }
`

func TestAccGslbparameter_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		// gslb resource do not have DELETE operation
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccGslbparameter_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGslbparameterExist("citrixadc_gslbparameter.tf_gslbparameter", nil),
					resource.TestCheckResourceAttr("citrixadc_gslbparameter.tf_gslbparameter", "ldnsentrytimeout" , "50"),
					resource.TestCheckResourceAttr("citrixadc_gslbparameter.tf_gslbparameter", "rtttolerance" , "6"),
					resource.TestCheckResourceAttr("citrixadc_gslbparameter.tf_gslbparameter", "ldnsmask" , "255.255.255.255"),
				),
			},
			resource.TestStep{
				Config: testAccGslbparameter_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslparameterExist("citrixadc_gslbparameter.tf_gslbparameter", nil),
					resource.TestCheckResourceAttr("citrixadc_gslbparameter.tf_gslbparameter", "ldnsentrytimeout" , "70"),
					resource.TestCheckResourceAttr("citrixadc_gslbparameter.tf_gslbparameter", "rtttolerance" , "8"),
					resource.TestCheckResourceAttr("citrixadc_gslbparameter.tf_gslbparameter", "ldnsmask" , "255.255.255.254"),
				),
			},
		},
		
	})
}

func testAccCheckGslbparameterExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No gslbparameter name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(service.Gslbparameter.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("gslbparameter %s not found", n)
		}

		return nil
	}
}

func testAccCheckGslbparameterDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_gslbparameter" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Gslbparameter.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("gslbparameter %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
