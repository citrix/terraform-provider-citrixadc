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

	"github.com/chiradeep/go-nitro/netscaler"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const testAccIpset_no_bindings = `
	resource "citrixadc_ipset" "foo" {
		name = "tf_test_ipset"
	}
`

const testAccIpset_single_binding = `
	resource "citrixadc_ipset" "foo" {
		name = "tf_test_ipset"
		nsipbinding = [
			citrixadc_nsip.nsip1.ipaddress,
		]
	}

	resource "citrixadc_nsip" "nsip1" {
		ipaddress = "1.1.1.1"
		type      = "VIP"
		netmask   = "255.255.255.0"
	}
`

const testAccIpset_2nd_binding_added = `
	resource "citrixadc_ipset" "foo" {
		name = "tf_test_ipset"
		nsipbinding = [
			citrixadc_nsip.nsip1.ipaddress,
			citrixadc_nsip.nsip2.ipaddress,
		]
	}

	resource "citrixadc_nsip" "nsip1" {
		ipaddress = "1.1.1.1"
		type      = "VIP"
		netmask   = "255.255.255.0"
	}
	resource "citrixadc_nsip" "nsip2" {
		ipaddress = "2.2.2.2"
		type      = "SNIP"
		netmask   = "255.255.255.0"
	}
`

const testAccIpset_name_changed = `
	resource "citrixadc_ipset" "foo" {
		name = "tf_test_ipset_newname"
	}
`

func TestAccIpset_no_bindings(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIpsetDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccIpset_no_bindings,
				Check: resource.ComposeTestCheckFunc(testAccCheckIpsetExist("citrixadc_ipset.foo", nil),
					resource.TestCheckResourceAttr("citrixadc_ipset.foo", "name", "tf_test_ipset"),
				),
			},
			resource.TestStep{
				Config: testAccIpset_name_changed,
				Check: resource.ComposeTestCheckFunc(testAccCheckIpsetExist("citrixadc_ipset.foo", nil),
					resource.TestCheckResourceAttr("citrixadc_ipset.foo", "name", "tf_test_ipset_newname"),
				),
			},
		},
	})
}

func TestAccIpset_with_bindings(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIpsetDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccIpset_no_bindings,
				Check: resource.ComposeTestCheckFunc(testAccCheckIpsetExist("citrixadc_ipset.foo", nil),
					resource.TestCheckResourceAttr("citrixadc_ipset.foo", "name", "tf_test_ipset"),
				),
			},
			resource.TestStep{
				Config: testAccIpset_single_binding,
				Check: resource.ComposeTestCheckFunc(testAccCheckIpsetExist("citrixadc_ipset.foo", nil),
					resource.TestCheckResourceAttr("citrixadc_ipset.foo", "name", "tf_test_ipset"),
					// TODO: add TechCheckResourceAttr() for checking bindings
					// resource.TestCheckResourceAttr("citrixadc_ipset.foo", "ipaddress", "1.1.1.1"),
				),
			},
			resource.TestStep{
				Config: testAccIpset_2nd_binding_added,
				Check:  resource.ComposeTestCheckFunc(testAccCheckIpsetExist("citrixadc_ipset.foo", nil)),// resource.TestCheckResourceAttr("citrixadc_ipset.foo", "name", "tf_test_ipset"),
				// resource.TestCheckResourceAttr("citrixadc_ipset.foo", "ipaddress", "[\"1.1.1.1\", \"2.2.2.2\"]"), // list order may vary

			},
		},
	})
}

func testAccCheckIpsetExist(n string, id *string) resource.TestCheckFunc {
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
		data, err := nsClient.FindResource(netscaler.Ipset.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("LB vserver %s not found", n)
		}

		return nil
	}
}

func testAccCheckIpsetDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_ipset" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(netscaler.Ipset.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("LB vserver %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
