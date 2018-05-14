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
	"github.com/chiradeep/go-nitro/netscaler"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"testing"
)

func TestAccNsacl_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNsaclDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNsacl_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsaclExist("netscaler_nsacl.foo", nil),

					resource.TestCheckResourceAttr(
						"netscaler_nsacl.foo", "aclaction", "DENY"),
					resource.TestCheckResourceAttr(
						"netscaler_nsacl.foo", "aclname", "test_acl"),
					resource.TestCheckResourceAttr(
						"netscaler_nsacl.foo", "destip", "192.168.1.1"),
					resource.TestCheckResourceAttr(
						"netscaler_nsacl.foo", "protocol", "TCP"),
					resource.TestCheckResourceAttr(
						"netscaler_nsacl.foo", "srcport", "45-1024"),
				),
			},
		},
	})
}

func testAccCheckNsaclExist(n string, id *string) resource.TestCheckFunc {
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
		data, err := nsClient.FindResource(netscaler.Nsacl.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("LB vserver %s not found", n)
		}

		return nil
	}
}

func testAccCheckNsaclDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "netscaler_nsacl" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(netscaler.Nsacl.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("LB vserver %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccNsacl_basic = `


resource "netscaler_nsacl" "foo" {
  
  aclaction = "DENY"
  aclname = "test_acl"
  destipval = "192.168.1.33"
  protocol = "TCP"
  srcportval = "45-1024"
  priority = "100"

}
`
