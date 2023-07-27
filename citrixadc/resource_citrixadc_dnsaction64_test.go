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

const testAccDnsaction64_add = `


resource "citrixadc_dnsaction64" "dnsaction64" {
	actionname = "default_DNS64_action1"
    prefix = "64:ff9c::/96"
    mappedrule = "DNS.RR.RDATA.IP.IN_SUBNET(10.0.0.0/8)"
    excluderule = "DNS.RR.RDATA.IPV6.IN_SUBNET(::ffff:0:0/96)"
}

`

const testAccDnsaction64_update = `

resource "citrixadc_dnsaction64" "dnsaction64" {
	actionname = "default_DNS64_action1"
    prefix = "64:ff9c::/96"
    mappedrule = "DNS.RR.TYPE.EQ(A)"
    excluderule = "DNS.RR.TYPE.EQ(AAAA)"
}

`

func TestAccDnsaction64_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDnsaction64Destroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDnsaction64_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDnsaction64Exist("citrixadc_dnsaction64.dnsaction64", nil),
					resource.TestCheckResourceAttr("citrixadc_dnsaction64.dnsaction64", "actionname", "default_DNS64_action1"),
					resource.TestCheckResourceAttr("citrixadc_dnsaction64.dnsaction64", "prefix", "64:ff9c::/96"),
					resource.TestCheckResourceAttr("citrixadc_dnsaction64.dnsaction64", "mappedrule", "DNS.RR.RDATA.IP.IN_SUBNET(10.0.0.0/8)"),
					resource.TestCheckResourceAttr("citrixadc_dnsaction64.dnsaction64", "excluderule", "DNS.RR.RDATA.IPV6.IN_SUBNET(::ffff:0:0/96)"),
				),
			},
			{
				Config: testAccDnsaction64_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDnsaction64Exist("citrixadc_dnsaction64.dnsaction64", nil),
					resource.TestCheckResourceAttr("citrixadc_dnsaction64.dnsaction64", "actionname", "default_DNS64_action1"),
					resource.TestCheckResourceAttr("citrixadc_dnsaction64.dnsaction64", "prefix", "64:ff9c::/96"),
					resource.TestCheckResourceAttr("citrixadc_dnsaction64.dnsaction64", "mappedrule", "DNS.RR.TYPE.EQ(A)"),
					resource.TestCheckResourceAttr("citrixadc_dnsaction64.dnsaction64", "excluderule", "DNS.RR.TYPE.EQ(AAAA)"),
				),
			},
		},
	})
}

func testAccCheckDnsaction64Exist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No dnsaction64 name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(service.Dnsaction64.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("dnsaction64 %s not found", n)
		}

		return nil
	}
}

func testAccCheckDnsaction64Destroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_dnsaction64" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Dnsaction64.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("dnsaction64 %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
