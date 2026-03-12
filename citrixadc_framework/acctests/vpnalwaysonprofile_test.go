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

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
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

const testAccVpnalwaysonprofileDataSource_basic = `

	resource "citrixadc_vpnalwaysonprofile" "tf_vpnalwaysonprofile" {
		name = "tf_vpnalwaysonprofile"
		clientcontrol = "DENY"
		locationbasedvpn = "Remote"
		networkaccessonvpnfailure = "onlyToGateway"
	}

	data "citrixadc_vpnalwaysonprofile" "tf_vpnalwaysonprofile" {
		name = citrixadc_vpnalwaysonprofile.tf_vpnalwaysonprofile.name
	}
`

func TestAccVpnalwaysonprofile_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVpnalwaysonprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnalwaysonprofile_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnalwaysonprofileExist("citrixadc_vpnalwaysonprofile.tf_vpnalwaysonprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnalwaysonprofile.tf_vpnalwaysonprofile", "name", "tf_vpnalwaysonprofile"),
					resource.TestCheckResourceAttr("citrixadc_vpnalwaysonprofile.tf_vpnalwaysonprofile", "clientcontrol", "DENY"),
					resource.TestCheckResourceAttr("citrixadc_vpnalwaysonprofile.tf_vpnalwaysonprofile", "locationbasedvpn", "Remote"),
					resource.TestCheckResourceAttr("citrixadc_vpnalwaysonprofile.tf_vpnalwaysonprofile", "networkaccessonvpnfailure", "onlyToGateway"),
				),
			},
			{
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

		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource("vpnalwaysonprofile", rs.Primary.ID)

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
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_vpnalwaysonprofile" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("vpnalwaysonprofile", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("vpnalwaysonprofile %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccVpnalwaysonprofileDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnalwaysonprofileDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_vpnalwaysonprofile.tf_vpnalwaysonprofile", "name", "tf_vpnalwaysonprofile"),
					resource.TestCheckResourceAttr("data.citrixadc_vpnalwaysonprofile.tf_vpnalwaysonprofile", "clientcontrol", "DENY"),
					resource.TestCheckResourceAttr("data.citrixadc_vpnalwaysonprofile.tf_vpnalwaysonprofile", "locationbasedvpn", "Remote"),
					resource.TestCheckResourceAttr("data.citrixadc_vpnalwaysonprofile.tf_vpnalwaysonprofile", "networkaccessonvpnfailure", "onlyToGateway"),
				),
			},
		},
	})
}
