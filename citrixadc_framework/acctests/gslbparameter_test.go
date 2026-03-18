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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testAccGslbparameter_basic = `

resource "citrixadc_gslbparameter" "tf_gslbparameter" {
	ldnsentrytimeout = 50
	rtttolerance     = 6
	ldnsmask         = "255.255.255.255"
	gslbsyncsaveconfigcommand = "DISABLED"
	}
`

const testAccGslbparameter_update = `

resource "citrixadc_gslbparameter" "tf_gslbparameter" {
	ldnsentrytimeout = 70
	rtttolerance     = 8
	ldnsmask         = "255.255.255.254"
	gslbsyncsaveconfigcommand = "ENABLED"
	}
`

func TestAccGslbparameter_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// gslb resource do not have DELETE operation
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				Config: testAccGslbparameter_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGslbparameterExist("citrixadc_gslbparameter.tf_gslbparameter", nil),
					resource.TestCheckResourceAttr("citrixadc_gslbparameter.tf_gslbparameter", "ldnsentrytimeout", "50"),
					resource.TestCheckResourceAttr("citrixadc_gslbparameter.tf_gslbparameter", "rtttolerance", "6"),
					resource.TestCheckResourceAttr("citrixadc_gslbparameter.tf_gslbparameter", "ldnsmask", "255.255.255.255"),
					resource.TestCheckResourceAttr("citrixadc_gslbparameter.tf_gslbparameter", "gslbsyncsaveconfigcommand", "DISABLED"),
				),
			},
			{
				Config: testAccGslbparameter_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGslbparameterExist("citrixadc_gslbparameter.tf_gslbparameter", nil),
					resource.TestCheckResourceAttr("citrixadc_gslbparameter.tf_gslbparameter", "ldnsentrytimeout", "70"),
					resource.TestCheckResourceAttr("citrixadc_gslbparameter.tf_gslbparameter", "rtttolerance", "8"),
					resource.TestCheckResourceAttr("citrixadc_gslbparameter.tf_gslbparameter", "ldnsmask", "255.255.255.254"),
					resource.TestCheckResourceAttr("citrixadc_gslbparameter.tf_gslbparameter", "gslbsyncsaveconfigcommand", "ENABLED"),
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

		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Gslbparameter.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("gslbparameter %s not found", n)
		}

		return nil
	}
}

const testAccGslbparameterDataSource_basic = `

	resource "citrixadc_gslbparameter" "tf_gslbparameter" {
		ldnsentrytimeout = 50
		rtttolerance     = 6
		ldnsmask         = "255.255.255.255"
		gslbsyncsaveconfigcommand = "DISABLED"
	}

	data "citrixadc_gslbparameter" "tf_gslbparameter" {
		depends_on = [citrixadc_gslbparameter.tf_gslbparameter]
	}
`

func TestAccGslbparameterDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccGslbparameterDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_gslbparameter.tf_gslbparameter", "ldnsentrytimeout", "50"),
					resource.TestCheckResourceAttr("data.citrixadc_gslbparameter.tf_gslbparameter", "rtttolerance", "6"),
					resource.TestCheckResourceAttr("data.citrixadc_gslbparameter.tf_gslbparameter", "ldnsmask", "255.255.255.255"),
					resource.TestCheckResourceAttr("data.citrixadc_gslbparameter.tf_gslbparameter", "gslbsyncsaveconfigcommand", "DISABLED"),
				),
			},
		},
	})
}
