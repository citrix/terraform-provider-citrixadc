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

func TestAccService_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckServiceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccService_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServiceExist("netscaler_service.foo", nil),
					resource.TestCheckResourceAttr(
						"netscaler_service.foo", "lbvserver", "foo_lb"), resource.TestCheckResourceAttr(
						"netscaler_service.foo", "name", "foo_svc"), resource.TestCheckResourceAttr(
						"netscaler_service.foo", "port", "80"), resource.TestCheckResourceAttr(
						"netscaler_service.foo", "servername", "10.202.22.12"), resource.TestCheckResourceAttr(
						"netscaler_service.foo", "servicetype", "HTTP"),
				),
			},
		},
	})
}

func testAccCheckServiceExist(n string, id *string) resource.TestCheckFunc {
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
		data, err := nsClient.FindResource(netscaler.Service.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("LB vserver %s not found", n)
		}

		return nil
	}
}

func testAccCheckServiceDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "netscaler_service" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(netscaler.Service.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("LB vserver %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccService_basic = `

resource "netscaler_lbvserver" "foo" {
  
  ipv46 = "10.202.11.11"
  name = "foo_lb"
  port = 80
  servicetype = "HTTP"
}


resource "netscaler_service" "foo" {
  
  lbvserver = "foo_lb"
  name = "foo_svc"
  port = 80
  ip = "10.202.22.12"
  servicetype = "HTTP"

  depends_on = ["netscaler_lbvserver.foo"]

}
`
