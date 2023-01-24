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
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"testing"
)

const testAccVpnalwaysonprofile_basic = `

	resource "citrixadc_vpnalwaysonprofile" "tf_vpnalwaysonprofile" {
		name = "tf_vpnalwaysonprofile"
		clientcontrol = "DENY"
		locationbasedvpn = "Remote"
		networkaccessonvpnfailure = "onlyToGateway"
	}
`

const testAccVpnalwaysonprofile_basic_update = `

	resource "citrixadc_vpnalwaysonprofile" "tf_vpnalwaysonprofile" {
		name = "tf_vpnalwaysonprofile"
		clientcontrol = "ALLOW"
		locationbasedvpn = "Everywhere"
		networkaccessonvpnfailure = "fullAccess"
	}
`

func TestAccVpnalwaysonprofile_basic(t *testing.T) {
	if adcTestbed != "STANDALONE" {
		t.Skipf("ADC testbed is %s. Expected STANDALONE.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckVpnalwaysonprofileDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccVpnalwaysonprofile_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnalwaysonprofileExist("citrixadc_vpnalwaysonprofile.tf_vpnalwaysonprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnalwaysonprofile.tf_vpnalwaysonprofile", "name", "tf_vpnalwaysonprofile"),
					resource.TestCheckResourceAttr("citrixadc_vpnalwaysonprofile.tf_vpnalwaysonprofile", "clientcontrol", "DENY"),
					resource.TestCheckResourceAttr("citrixadc_vpnalwaysonprofile.tf_vpnalwaysonprofile", "locationbasedvpn", "Remote"),
					resource.TestCheckResourceAttr("citrixadc_vpnalwaysonprofile.tf_vpnalwaysonprofile", "networkaccessonvpnfailure", "onlyToGateway"),
				),
			},
			resource.TestStep{
				Config: testAccVpnalwaysonprofile_basic_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnalwaysonprofileExist("citrixadc_vpnalwaysonprofile.tf_vpnalwaysonprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnalwaysonprofile.tf_vpnalwaysonprofile", "clientcontrol", "ALLOW"),
					resource.TestCheckResourceAttr("citrixadc_vpnalwaysonprofile.tf_vpnalwaysonprofile", "locationbasedvpn", "Everywhere"),
					resource.TestCheckResourceAttr("citrixadc_vpnalwaysonprofile.tf_vpnalwaysonprofile", "networkaccessonvpnfailure", "fullAccess"),
				),
			},
		},
	})
}

func testAccCheckVpnalwaysonprofileExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No vpnalwaysonprofile name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource("vpnalwaysonprofile", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("vpnalwaysonprofile %s not found", n)
		}

		return nil
	}
}

func testAccCheckVpnalwaysonprofileDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_vpnalwaysonprofile" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource("vpnalwaysonprofile", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("vpnalwaysonprofile %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
