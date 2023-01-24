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

const testAccSnmpoption_basic = `
	resource "citrixadc_opoption" "tf_opoption" {
		snmpset              = "ENABLED"
		snmptraplogging      = "ENABLED"
		partitionnameintrap  = "ENABLED"
		snmptraplogginglevel = "WARNING"
	}
  
`
const testAccSnmpoption_update = `
	
	resource "citrixadc_opoption" "tf_opoption" {
		snmpset              = "DISABLED"
		snmptraplogging      = "DISABLED"
		partitionnameintrap  = "DISABLED"
		snmptraplogginglevel = "ERROR"

	}
  
`

func TestAccSnmpoption_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccSnmpoption_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSnmpoptionExist("citrixadc_opoption.tf_opoption", nil),
					resource.TestCheckResourceAttr("citrixadc_opoption.tf_opoption","snmpset", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_opoption.tf_opoption","snmptraplogging", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_opoption.tf_opoption","partitionnameintrap", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_opoption.tf_opoption","snmptraplogginglevel", "WARNING"),
				),
			},
			resource.TestStep{
				Config: testAccSnmpoption_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSnmpoptionExist("citrixadc_opoption.tf_opoption", nil),
					resource.TestCheckResourceAttr("citrixadc_opoption.tf_opoption","snmpset", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_opoption.tf_opoption","snmptraplogging", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_opoption.tf_opoption","partitionnameintrap", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_opoption.tf_opoption","snmptraplogginglevel", "ERROR"),
				),
			},
		},
	})
}

func testAccCheckSnmpoptionExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No snmpoption name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(service.Snmpoption.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("snmpoption %s not found", n)
		}

		return nil
	}
}