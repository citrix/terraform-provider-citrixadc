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

func TestAccGslbvserver_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckGslbvserverDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccGslbvserver_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGslbvserverExist("netscaler_gslbvserver.foo", nil),

					resource.TestCheckResourceAttr(
						"netscaler_gslbvserver.foo", "dnsrecordtype", "A"),
					resource.TestCheckResourceAttr(
						"netscaler_gslbvserver.foo", "name", "GSLB-East-Coast-Vserver"),
					resource.TestCheckResourceAttr(
						"netscaler_gslbvserver.foo", "servicetype", "HTTP"),
				),
			},
		},
	})
}

func testAccCheckGslbvserverExist(n string, id *string) resource.TestCheckFunc {
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
		data, err := nsClient.FindResource(netscaler.Gslbvserver.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("GSLB vserver %s not found", n)
		}

		return nil
	}
}

func testAccCheckGslbvserverDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "netscaler_gslbvserver" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(netscaler.Gslbvserver.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("GSLB vserver %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccGslbvserver_basic = `


resource "netscaler_gslbvserver" "foo" {
  
  dnsrecordtype = "A"
  name = "GSLB-East-Coast-Vserver"
  servicetype = "HTTP"
  domain {
	  domainname =  "www.fooco.co"
	  ttl = "60"
  }
  domain {
	  domainname = "www.barco.com"
	  ttl = "55"
  }
}
`
