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
package netscaler

import (
	"fmt"
	"testing"

	"github.com/chiradeep/go-nitro/netscaler"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccGslbservice_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckGslbserviceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccGslbservice_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGslbserviceExist("netscaler_gslbservice.foo", nil),

					resource.TestCheckResourceAttr(
						"netscaler_gslbservice.foo", "ipaddress", "172.16.1.101"),
					resource.TestCheckResourceAttr(
						"netscaler_gslbservice.foo", "port", "80"),
					resource.TestCheckResourceAttr(
						"netscaler_gslbservice.foo", "servicename", "gslb1vservice"),
					resource.TestCheckResourceAttr(
						"netscaler_gslbservice.foo", "servicetype", "HTTP"),
					resource.TestCheckResourceAttr(
						"netscaler_gslbservice.foo", "sitename", "Site-GSLB-East-Coast"),
				),
			},
		},
	})
}

func testAccCheckGslbserviceExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lb vserver name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(netscaler.Gslbservice.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("LB vserver %s not found", n)
		}

		return nil
	}
}

func testAccCheckGslbserviceDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "netscaler_gslbservice" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(netscaler.Gslbservice.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("LB vserver %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccGslbservice_basic = `
resource "netscaler_gslbsite" "foo" {
  
	siteipaddress = "172.31.11.20"
	sitename = "Site-GSLB-East-Coast"
  
  }

resource "netscaler_gslbservice" "foo" {
  
  ip = "172.16.1.101"
  port = "80"
  servicename = "gslb1vservice"
  servicetype = "HTTP"
  sitename = "${netscaler_gslbsite.foo.sitename}"

}
`
