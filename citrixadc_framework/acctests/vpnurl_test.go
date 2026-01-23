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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"testing"
)

const testAccVpnurl_add = `

resource "citrixadc_vpnurl" "foo" {
	actualurl        = "http://www.citrix.com"
	appjson          = "xyz"
	applicationtype  = "CVPN"
	clientlessaccess = "OFF"
	comment          = "Testing"
	linkname         = "Description"
	ssotype          = "unifiedgateway"
	urlname          = "Firsturl"
	vservername      = "server1"
	}
`
const testAccVpnurl_update = `

resource "citrixadc_vpnurl" "foo" {
	actualurl        = "http://www.citrix1.com"
	appjson          = "xyz"
	applicationtype  = "CVPN"
	clientlessaccess = "OFF"
	comment          = "Testing"
	linkname         = "Description"
	ssotype          = "unifiedgateway"
	urlname          = "Firsturl"
	vservername      = "server1"
	}
`

func TestAccVpnurl_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVpnurlDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnurl_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnurlExist("citrixadc_vpnurl.foo", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnurl.foo", "urlname", "Firsturl"),
					resource.TestCheckResourceAttr("citrixadc_vpnurl.foo", "actualurl", "http://www.citrix.com"),
					resource.TestCheckResourceAttr("citrixadc_vpnurl.foo", "linkname", "Description"),
					resource.TestCheckResourceAttr("citrixadc_vpnurl.foo", "comment", "Testing"),
				),
			},
			{
				Config: testAccVpnurl_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnurlExist("citrixadc_vpnurl.foo", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnurl.foo", "urlname", "Firsturl"),
					resource.TestCheckResourceAttr("citrixadc_vpnurl.foo", "actualurl", "http://www.citrix1.com"),
					resource.TestCheckResourceAttr("citrixadc_vpnurl.foo", "ssotype", "unifiedgateway"),
					resource.TestCheckResourceAttr("citrixadc_vpnurl.foo", "vservername", "server1"),
				),
			},
		},
	})
}

func testAccCheckVpnurlExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No vpnurl name is set")
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
		data, err := client.FindResource(service.Vpnurl.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("vpnurl %s not found", n)
		}

		return nil
	}
}

func testAccCheckVpnurlDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_vpnurl" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Vpnurl.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("vpnurl %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
