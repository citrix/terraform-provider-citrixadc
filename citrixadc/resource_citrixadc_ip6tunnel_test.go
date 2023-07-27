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

const testAccIp6tunnel_add = `
	resource "citrixadc_nsip6" "test_nsip" {
		ipv6address = "23::30:20:23:34/64"
		type        = "VIP"
		icmp        = "DISABLED"
	}
	resource "citrixadc_ip6tunnel" "tf_ip6tunnel" {
		name   = "tf_ip6tunnel"
		remote = "2001:db8:0:b::/64"
		local  = trimsuffix(citrixadc_nsip6.test_nsip.ipv6address, "/64")
	}
`

func TestAccIp6tunnel_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIp6tunnelDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccIp6tunnel_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIp6tunnelExist("citrixadc_ip6tunnel.tf_ip6tunnel", nil),
					resource.TestCheckResourceAttr("citrixadc_ip6tunnel.tf_ip6tunnel", "name", "tf_ip6tunnel"),
					resource.TestCheckResourceAttr("citrixadc_ip6tunnel.tf_ip6tunnel", "remote", "2001:db8:0:b::/64"),
					resource.TestCheckResourceAttr("citrixadc_ip6tunnel.tf_ip6tunnel", "local", "23::30:20:23:34"),
				),
			},
		},
	})
}

func testAccCheckIp6tunnelExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ip6tunnel name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(service.Ip6tunnel.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("ip6tunnel %s not found", n)
		}

		return nil
	}
}

func testAccCheckIp6tunnelDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_ip6tunnel" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Ip6tunnel.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("ip6tunnel %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
