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

// systemcpuparam is a set/get singleton. pemode accepts DEFAULT/CPUBOUND.
// SAFETY: tests only ever set pemode = "DEFAULT". CPUBOUND rebinds packet
// engines to CPUs and can be disruptive (possible reboot) on a live appliance,
// so it is intentionally never exercised here.

const testAccSystemcpuparam_basic = `
	resource "citrixadc_systemcpuparam" "tf_systemcpuparam" {
		pemode = "DEFAULT"
	}
`

func TestAccSystemcpuparam_basic(t *testing.T) {
	t.Skip("TODO: Requires review")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// Singleton resource - cannot be deleted from the ADC.
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				Config: testAccSystemcpuparam_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSystemcpuparamExist("citrixadc_systemcpuparam.tf_systemcpuparam", nil),
					resource.TestCheckResourceAttr("citrixadc_systemcpuparam.tf_systemcpuparam", "pemode", "DEFAULT"),
				),
			},
		},
	})
}

func testAccCheckSystemcpuparamExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No systemcpuparam name is set")
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
		data, err := client.FindResource(service.Systemcpuparam.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("systemcpuparam %s not found", n)
		}

		return nil
	}
}

const testAccSystemcpuparamDataSource_basic = `

	resource "citrixadc_systemcpuparam" "tf_systemcpuparam" {
		pemode = "DEFAULT"
	}

	data "citrixadc_systemcpuparam" "tf_systemcpuparam" {
		depends_on = [citrixadc_systemcpuparam.tf_systemcpuparam]
	}
`

func TestAccSystemcpuparamDataSource_basic(t *testing.T) {
	t.Skip("TODO: Requires review")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccSystemcpuparamDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_systemcpuparam.tf_systemcpuparam", "pemode", "DEFAULT"),
				),
			},
		},
	})
}
