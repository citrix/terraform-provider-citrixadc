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

const testAccVpnparameter_add = `

	resource "citrixadc_vpnparameter" "tf_vpnparameter" {
		splitdns              = "LOCAL"
		sesstimeout           = 30
		clientsecuritylog     = "OFF"
		smartgroup            = 10
		splittunnel           = "ON"
		locallanaccess        = "ON"
		winsip                = "4.45.5.4"
		samesite              = "None"
		backendcertvalidation = "DISABLED"
		backendserversni      = "DISABLED"
		icasessiontimeout     = "OFF"
		iconwithreceiver      = "OFF"
		linuxpluginupgrade    = "Always"
		uitheme               = "DEFAULT"
		httpport              = [80]
		secureprivateaccess	= "ENABLED"
		maxiipperuser         = 5
		httptrackconnproxy	= "OFF"
		deviceposture = "DISABLED"
		backenddtls12 = "DISABLED"
		accessrestrictedpageredirect = "NS"
	}
`
const testAccVpnparameter_update = `

	resource "citrixadc_vpnparameter" "tf_vpnparameter" {
		splitdns              = "LOCAL"
		sesstimeout           = 30
		clientsecuritylog     = "OFF"
		smartgroup            = 10
		splittunnel           = "OFF"
		locallanaccess        = "OFF"
		winsip                = "4.45.5.4"
		samesite              = "None"
		backendcertvalidation = "DISABLED"
		backendserversni      = "DISABLED"
		icasessiontimeout     = "OFF"
		iconwithreceiver      = "OFF"
		linuxpluginupgrade    = "Always"
		uitheme               = "DEFAULT"
		httpport              = [80]
		secureprivateaccess	= "DISABLED"
		maxiipperuser         = 10
		httptrackconnproxy	= "ON"
		deviceposture = "ENABLED"
		backenddtls12 = "ENABLED"
	}
`

func TestAccVpnparameter_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// vpnparameter resource do not have DELETE operation
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnparameter_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnparameterExist("citrixadc_vpnparameter.tf_vpnparameter", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnparameter.tf_vpnparameter", "splittunnel", "ON"),
					resource.TestCheckResourceAttr("citrixadc_vpnparameter.tf_vpnparameter", "locallanaccess", "ON"),
					resource.TestCheckResourceAttr("citrixadc_vpnparameter.tf_vpnparameter", "secureprivateaccess", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_vpnparameter.tf_vpnparameter", "httptrackconnproxy", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_vpnparameter.tf_vpnparameter", "deviceposture", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_vpnparameter.tf_vpnparameter", "backenddtls12", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_vpnparameter.tf_vpnparameter", "accessrestrictedpageredirect", "NS"),
				),
			},
			{
				Config: testAccVpnparameter_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnparameterExist("citrixadc_vpnparameter.tf_vpnparameter", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnparameter.tf_vpnparameter", "splittunnel", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_vpnparameter.tf_vpnparameter", "locallanaccess", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_vpnparameter.tf_vpnparameter", "secureprivateaccess", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_vpnparameter.tf_vpnparameter", "httptrackconnproxy", "ON"),
					resource.TestCheckResourceAttr("citrixadc_vpnparameter.tf_vpnparameter", "deviceposture", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_vpnparameter.tf_vpnparameter", "backenddtls12", "ENABLED"),
				),
			},
		},
	})
}

func testAccCheckVpnparameterExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No vpnparameter name is set")
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
		data, err := client.FindResource(service.Vpnparameter.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("vpnparameter %s not found", n)
		}

		return nil
	}
}
