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
	"github.com/chiradeep/go-nitro/netscaler"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"testing"
)

func TestAccRnat_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckRnatDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccRnat_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRnatExist("citrixadc_rnat.foo", nil),
				),
			},
		},
	})
}

func testAccCheckRnatExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No rnat name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		d, err := nsClient.FindFilteredResourceArray(netscaler.Rnat.Type(), map[string]string{"network": "192.168.96.0", "netmask": "255.255.240.0", "natip": "*"})

		if err != nil {
			return err
		}

		if len(d) != 1 {
			return fmt.Errorf("Rnat rule %s not found", n)
		}

		return nil
	}
}

func testAccCheckRnatDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_rnat" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindFilteredResourceArray(netscaler.Rnat.Type(), map[string]string{"network": "192.168.96.0", "netmask": "255.255.240.0", "natip": "*"})
		if err == nil {
			return fmt.Errorf("Rnat rule %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccRnat_basic = `


resource "citrixadc_rnat" "foo" {
	rnatsname = "foo"
	rnat {
           network = "192.168.96.0"
           netmask = "255.255.240.0"
         }
}
`
