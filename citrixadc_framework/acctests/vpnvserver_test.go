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

const testAccVpnvserver_add = `
	resource "citrixadc_ipset" "tf_ipset" {
		name = "tf_test_ipset"
	}
	resource "citrixadc_vpnvserver" "foo" {
		name                     = "tf.citrix.example.com"
		servicetype              = "SSL"
		ipv46                    = "3.3.3.3"
		port                     = 443
		ipset                    = citrixadc_ipset.tf_ipset.name
		dtls                     = "OFF"
		downstateflush           = "DISABLED"
		listenpolicy             = "NONE"
		tcpprofilename           = "nstcp_default_XA_XD_profile"
		secureprivateaccess		= "ENABLED"
		accessrestrictedpageredirect = "NS"
		deviceposture 		  = "DISABLED"
	}
`

const testAccVpnvserver_update = `
	resource "citrixadc_ipset" "tf_ipset" {
		name = "tf_test_ipset"
	}
	resource "citrixadc_vpnvserver" "foo" {
		name                     = "tf.citrix.example.com"
		servicetype              = "SSL"
		ipv46                    = "3.3.3.3"
		port                     = 443
		ipset                    = citrixadc_ipset.tf_ipset.name
		dtls                     = "OFF"
		downstateflush           = "ENABLED"
		listenpolicy             = "NONE"
		tcpprofilename           = "nstcp_default_XA_XD_profile"
		secureprivateaccess		= "DISABLED"
		deviceposture 		  = "ENABLED"
	}
`

func TestAccVpnvserver_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVpnvserverDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnvserver_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnvserverExist("citrixadc_vpnvserver.foo", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnvserver.foo", "name", "tf.citrix.example.com"),
					resource.TestCheckResourceAttr("citrixadc_vpnvserver.foo", "servicetype", "SSL"),
					resource.TestCheckResourceAttr("citrixadc_vpnvserver.foo", "ipv46", "3.3.3.3"),
					resource.TestCheckResourceAttr("citrixadc_vpnvserver.foo", "downstateflush", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_vpnvserver.foo", "secureprivateaccess", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_vpnvserver.foo", "deviceposture", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_vpnvserver.foo", "accessrestrictedpageredirect", "NS"),
				),
			},
			{
				Config: testAccVpnvserver_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnvserverExist("citrixadc_vpnvserver.foo", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnvserver.foo", "name", "tf.citrix.example.com"),
					resource.TestCheckResourceAttr("citrixadc_vpnvserver.foo", "servicetype", "SSL"),
					resource.TestCheckResourceAttr("citrixadc_vpnvserver.foo", "ipv46", "3.3.3.3"),
					resource.TestCheckResourceAttr("citrixadc_vpnvserver.foo", "downstateflush", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_vpnvserver.foo", "secureprivateaccess", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_vpnvserver.foo", "deviceposture", "ENABLED"),
				),
			},
		},
	})
}

func testAccCheckVpnvserverExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No vpnvserver name is set")
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
		data, err := client.FindResource(service.Vpnvserver.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("vpnvserver %s not found", n)
		}

		return nil
	}
}

func testAccCheckVpnvserverDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_vpnvserver" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Vpnvserver.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("vpnvserver %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccVpnvserverDataSource_basic = `
	resource "citrixadc_ipset" "tf_ipset" {
		name = "tf_test_ipset"
	}
	resource "citrixadc_vpnvserver" "foo" {
		name                     = "tf.citrix.example.com"
		servicetype              = "SSL"
		ipv46                    = "3.3.3.3"
		port                     = 443
		ipset                    = citrixadc_ipset.tf_ipset.name
		dtls                     = "OFF"
		downstateflush           = "DISABLED"
		listenpolicy             = "NONE"
		tcpprofilename           = "nstcp_default_XA_XD_profile"
		secureprivateaccess		= "ENABLED"
		accessrestrictedpageredirect = "NS"
		deviceposture 		  = "DISABLED"
	}

data "citrixadc_vpnvserver" "foo" {
	name = citrixadc_vpnvserver.foo.name
}
`

func TestAccVpnvserverDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnvserverDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_vpnvserver.foo", "name", "tf.citrix.example.com"),
					resource.TestCheckResourceAttr("data.citrixadc_vpnvserver.foo", "servicetype", "SSL"),
					resource.TestCheckResourceAttr("data.citrixadc_vpnvserver.foo", "ipv46", "3.3.3.3"),
					resource.TestCheckResourceAttr("data.citrixadc_vpnvserver.foo", "downstateflush", "DISABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_vpnvserver.foo", "secureprivateaccess", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_vpnvserver.foo", "deviceposture", "DISABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_vpnvserver.foo", "accessrestrictedpageredirect", "NS"),
				),
			},
		},
	})
}
