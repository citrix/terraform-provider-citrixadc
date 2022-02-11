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

const testAccVpnurlaction_add = `

	resource "citrixadc_vpnurlaction" "foo" {
		
		name 				= "tf_vpnurlaction"
		actualurl 			= "www.citrix.com"
		linkname 			= "new_link"
		applicationtype 	= "CVPN"
		clientlessaccess 	= "ON"
		comment 			= "Testing"
		ssotype 			= "unifiedgateway"
		vservername 		= "vserver1"
	}
`
const testAccVpnurlaction_update = `

	resource "citrixadc_vpnurlaction" "foo" {
		
		name 				= "tf_vpnurlaction"
		actualurl 			= "https://www.citrix.com/products/citrix-adc/"
		linkname 			= "new_link"
		applicationtype 	= "CVPN"
		clientlessaccess 	= "OFF"
		comment 			= "Testing"
		ssotype 			= "unifiedgateway"
		vservername 		= "vserver1"
	}
`

func TestAccVpnurlaction_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckVpnurlactionDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccVpnurlaction_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnurlactionExist("citrixadc_vpnurlaction.foo", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnurlaction.foo", "name", "tf_vpnurlaction"),
					resource.TestCheckResourceAttr("citrixadc_vpnurlaction.foo", "actualurl", "www.citrix.com"),
					resource.TestCheckResourceAttr("citrixadc_vpnurlaction.foo", "clientlessaccess", "ON"),
				),
			},
			resource.TestStep{
				Config: testAccVpnurlaction_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnurlactionExist("citrixadc_vpnurlaction.foo", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnurlaction.foo", "name", "tf_vpnurlaction"),
					resource.TestCheckResourceAttr("citrixadc_vpnurlaction.foo", "actualurl", "https://www.citrix.com/products/citrix-adc/"),
					resource.TestCheckResourceAttr("citrixadc_vpnurlaction.foo", "clientlessaccess", "OFF"),
				),
			},
		},
	})
}

func testAccCheckVpnurlactionExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No vpnurlaction name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource("vpnurlaction", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("vpnurlaction %s not found", n)
		}

		return nil
	}
}

func testAccCheckVpnurlactionDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_vpnurlaction" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource("vpnurlaction", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("vpnurlaction %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
