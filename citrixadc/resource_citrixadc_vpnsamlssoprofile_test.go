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

const testAccVpnsamlssoprofile_add = `

	resource "citrixadc_vpnsamlssoprofile" "tf_vpnsamlssoprofile" {
		name                        = "tf_vpnsamlssoprofile"
		assertionconsumerserviceurl = "http://www.example.com"
		relaystaterule              = "true"
		sendpassword				= "OFF"
		samlissuername              = "issuername"
		signaturealg                = "RSA-SHA1"
		digestmethod                = "SHA256"
		nameidformat                = "Unspecified"
	}  
`
const testAccVpnsamlssoprofile_update = `

	resource "citrixadc_vpnsamlssoprofile" "tf_vpnsamlssoprofile" {
		name                        = "tf_vpnsamlssoprofile"
		assertionconsumerserviceurl = "http://www.example.com"
		sendpassword				= "ON"
		relaystaterule              = "true"
		samlissuername              = "issuernewname"
		signaturealg                = "RSA-SHA1"
		digestmethod                = "SHA1"
		nameidformat                = "Unspecified"
	}
`

func TestAccVpnsamlssoprofile_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckVpnsamlssoprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnsamlssoprofile_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnsamlssoprofileExist("citrixadc_vpnsamlssoprofile.tf_vpnsamlssoprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnsamlssoprofile.tf_vpnsamlssoprofile", "name", "tf_vpnsamlssoprofile"),
					resource.TestCheckResourceAttr("citrixadc_vpnsamlssoprofile.tf_vpnsamlssoprofile", "digestmethod", "SHA256"),
					resource.TestCheckResourceAttr("citrixadc_vpnsamlssoprofile.tf_vpnsamlssoprofile", "samlissuername", "issuername"),
				),
			},
			{
				Config: testAccVpnsamlssoprofile_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnsamlssoprofileExist("citrixadc_vpnsamlssoprofile.tf_vpnsamlssoprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnsamlssoprofile.tf_vpnsamlssoprofile", "name", "tf_vpnsamlssoprofile"),
					resource.TestCheckResourceAttr("citrixadc_vpnsamlssoprofile.tf_vpnsamlssoprofile", "digestmethod", "SHA1"),
					resource.TestCheckResourceAttr("citrixadc_vpnsamlssoprofile.tf_vpnsamlssoprofile", "samlissuername", "issuernewname"),
				),
			},
		},
	})
}

func testAccCheckVpnsamlssoprofileExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No vpnsamlssoprofile name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(service.Vpnsamlssoprofile.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("vpnsamlssoprofile %s not found", n)
		}

		return nil
	}
}

func testAccCheckVpnsamlssoprofileDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_vpnsamlssoprofile" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Vpnsamlssoprofile.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("vpnsamlssoprofile %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
