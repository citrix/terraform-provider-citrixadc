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
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"testing"
)

const testAccVpnpcoipprofile_add = `

	resource "citrixadc_vpnpcoipprofile" "tf_vpnpcoipprofile" {
		name               = "tf_vpnpcoipprofile"
		conserverurl       = "http://www.example.com"
		sessionidletimeout = 80
	}
`
const testAccVpnpcoipprofile_update = `

	resource "citrixadc_vpnpcoipprofile" "tf_vpnpcoipprofile" {
		name               = "tf_vpnpcoipprofile"
		conserverurl       = "http://www.example.com"
		sessionidletimeout = 90
	}
`

func TestAccVpnpcoipprofile_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckVpnpcoipprofileDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccVpnpcoipprofile_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnpcoipprofileExist("citrixadc_vpnpcoipprofile.tf_vpnpcoipprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnpcoipprofile.tf_vpnpcoipprofile", "name", "tf_vpnpcoipprofile"),
					resource.TestCheckResourceAttr("citrixadc_vpnpcoipprofile.tf_vpnpcoipprofile", "sessionidletimeout", "80"),
				),
			},
			resource.TestStep{
				Config: testAccVpnpcoipprofile_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnpcoipprofileExist("citrixadc_vpnpcoipprofile.tf_vpnpcoipprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnpcoipprofile.tf_vpnpcoipprofile", "name", "tf_vpnpcoipprofile"),
					resource.TestCheckResourceAttr("citrixadc_vpnpcoipprofile.tf_vpnpcoipprofile", "sessionidletimeout", "90"),
				),
			},
		},
	})
}

func testAccCheckVpnpcoipprofileExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No vpnpcoipprofile name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource("vpnpcoipprofile", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("vpnpcoipprofile %s not found", n)
		}

		return nil
	}
}

func testAccCheckVpnpcoipprofileDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_vpnpcoipprofile" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource("vpnpcoipprofile", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("vpnpcoipprofile %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
