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

const testAccVpnpcoipvserverprofile_add = `
	resource "citrixadc_vpnpcoipvserverprofile" "tf_vpnpcoipvserverprofile" {
		name        = "tf_vpnpcoipvserverprofile"
		logindomain = "domainname"
		udpport     = "802"
	}
`
const testAccVpnpcoipvserverprofile_update = `
	resource "citrixadc_vpnpcoipvserverprofile" "tf_vpnpcoipvserverprofile" {
		name        = "tf_vpnpcoipvserverprofile"
		logindomain = "domainname"
		udpport     = "200"
	}
`

const testAccVpnpcoipvserverprofileDataSource_basic = `
	resource "citrixadc_vpnpcoipvserverprofile" "tf_vpnpcoipvserverprofile" {
		name        = "tf_vpnpcoipvserverprofile"
		logindomain = "domainname"
		udpport     = "802"
	}

	data "citrixadc_vpnpcoipvserverprofile" "tf_vpnpcoipvserverprofile" {
		name = citrixadc_vpnpcoipvserverprofile.tf_vpnpcoipvserverprofile.name
	}
`

func TestAccVpnpcoipvserverprofile_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVpnpcoipvserverprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnpcoipvserverprofile_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnpcoipvserverprofileExist("citrixadc_vpnpcoipvserverprofile.tf_vpnpcoipvserverprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnpcoipvserverprofile.tf_vpnpcoipvserverprofile", "name", "tf_vpnpcoipvserverprofile"),
					resource.TestCheckResourceAttr("citrixadc_vpnpcoipvserverprofile.tf_vpnpcoipvserverprofile", "udpport", "802"),
				),
			},
			{
				Config: testAccVpnpcoipvserverprofile_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnpcoipvserverprofileExist("citrixadc_vpnpcoipvserverprofile.tf_vpnpcoipvserverprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnpcoipvserverprofile.tf_vpnpcoipvserverprofile", "name", "tf_vpnpcoipvserverprofile"),
					resource.TestCheckResourceAttr("citrixadc_vpnpcoipvserverprofile.tf_vpnpcoipvserverprofile", "udpport", "200"),
				),
			},
		},
	})
}

func testAccCheckVpnpcoipvserverprofileExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No vpnpcoipvserverprofile name is set")
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
		data, err := client.FindResource("vpnpcoipvserverprofile", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("vpnpcoipvserverprofile %s not found", n)
		}

		return nil
	}
}

func testAccCheckVpnpcoipvserverprofileDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_vpnpcoipvserverprofile" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("vpnpcoipvserverprofile", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("vpnpcoipvserverprofile %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccVpnpcoipvserverprofileDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnpcoipvserverprofileDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_vpnpcoipvserverprofile.tf_vpnpcoipvserverprofile", "name", "tf_vpnpcoipvserverprofile"),
					resource.TestCheckResourceAttr("data.citrixadc_vpnpcoipvserverprofile.tf_vpnpcoipvserverprofile", "logindomain", "domainname"),
					resource.TestCheckResourceAttr("data.citrixadc_vpnpcoipvserverprofile.tf_vpnpcoipvserverprofile", "udpport", "802"),
				),
			},
		},
	})
}
