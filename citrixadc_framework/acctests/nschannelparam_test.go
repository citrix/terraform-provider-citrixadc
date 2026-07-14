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

const testAccNschannelparam_basic_step1 = `
	resource "citrixadc_nschannelparam" "tf_nschannelparam" {
		vfautorecover = "ENABLE"
	}
`

const testAccNschannelparam_basic_step2 = `
	resource "citrixadc_nschannelparam" "tf_nschannelparam" {
		vfautorecover = "DISABLE"
	}
`

func TestAccNschannelparam_basic(t *testing.T) {
	t.Skip("TODO: Requires review")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// Singleton resource: never truly deleted from the ADC (state-only removal).
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				Config: testAccNschannelparam_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNschannelparamExist("citrixadc_nschannelparam.tf_nschannelparam", nil),
					resource.TestCheckResourceAttr("citrixadc_nschannelparam.tf_nschannelparam", "vfautorecover", "ENABLE"),
				),
			},
			{
				Config: testAccNschannelparam_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNschannelparamExist("citrixadc_nschannelparam.tf_nschannelparam", nil),
					resource.TestCheckResourceAttr("citrixadc_nschannelparam.tf_nschannelparam", "vfautorecover", "DISABLE"),
				),
			},
		},
	})
}

func testAccCheckNschannelparamExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No nschannelparam name is set")
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
		// Singleton: get-all, no name argument.
		data, err := client.FindResource(service.Nschannelparam.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("nschannelparam %s not found", n)
		}

		return nil
	}
}

const testAccNschannelparamDataSource_basic = `

	resource "citrixadc_nschannelparam" "tf_nschannelparam" {
		vfautorecover = "ENABLE"
	}

	data "citrixadc_nschannelparam" "tf_nschannelparam" {
		depends_on = [citrixadc_nschannelparam.tf_nschannelparam]
	}
`

func TestAccNschannelparamDataSource_basic(t *testing.T) {
	t.Skip("TODO: Requires review")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccNschannelparamDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_nschannelparam.tf_nschannelparam", "vfautorecover", "ENABLE"),
				),
			},
		},
	})
}
