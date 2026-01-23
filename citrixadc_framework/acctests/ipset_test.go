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

// TODO: add ipset_nsip6_binding testcase
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
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckIpsetDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccIpset_no_bindings,
				Check: resource.ComposeTestCheckFunc(testAccCheckIpsetExist("citrixadc_ipset.foo", nil),
					resource.TestCheckResourceAttr("citrixadc_ipset.foo", "name", "tf_test_ipset"),
				),
			},
			{
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
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckIpsetDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccIpset_no_bindings,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpsetExist("citrixadc_ipset.foo", nil),
				),
			},
			{
				Config: testAccIpset_single_binding,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpsetExist("citrixadc_ipset.foo", nil),
				),
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

		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Ipset.Type(), rs.Primary.ID)

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
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_ipset" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Ipset.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("LB vserver %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccIpset_swap_ipv4_step1 = `
resource "citrixadc_csvserver" "test_csvserver" {
  ipset = citrixadc_ipset.tf_ipset.name
  name        = "tf_csvserver"
  ipv46 = "10.78.60.22"
  port        = 80
  servicetype = "HTTP"

}

resource "citrixadc_ipset" "tf_ipset" {
  name = "tf_test_ipset"
  nsipbinding = [
	citrixadc_nsip.nsip2.ipaddress,
  ]

}

resource "citrixadc_nsip" "nsip1" {
  ipaddress = "10.1.1.1"
  type      = "VIP"
  netmask   = "255.255.255.0"
}
resource "citrixadc_nsip" "nsip2" {
  ipaddress = "10.2.2.2"
  type      = "VIP"
  netmask   = "255.255.255.0"
}
`

const testAccIpset_swap_ipv4_step2 = `
resource "citrixadc_csvserver" "test_csvserver" {
  ipset = citrixadc_ipset.tf_ipset.name
  name        = "tf_csvserver"
  ipv46 = "10.78.60.22"
  port        = 80
  servicetype = "HTTP"

}

resource "citrixadc_ipset" "tf_ipset" {
  name = "tf_test_ipset"
  nsipbinding = [
	citrixadc_nsip.nsip1.ipaddress,
  ]

}

resource "citrixadc_nsip" "nsip1" {
  ipaddress = "10.1.1.1"
  type      = "VIP"
  netmask   = "255.255.255.0"
}
resource "citrixadc_nsip" "nsip2" {
  ipaddress = "10.2.2.2"
  type      = "VIP"
  netmask   = "255.255.255.0"
}
`

func TestAccIpset_ipv4_swaps(t *testing.T) {
	if isCpxRun {
		t.Skip("No support in CPX")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckIpsetDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccIpset_swap_ipv4_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpsetExist("citrixadc_ipset.tf_ipset", nil),
				),
			},
			{
				Config: testAccIpset_swap_ipv4_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpsetExist("citrixadc_ipset.tf_ipset", nil),
				),
			},
		},
	})
}

const testAccIpset_swap_ipv6_step1 = `
resource "citrixadc_csvserver" "test_csvserver" {
  ipset = citrixadc_ipset.tf_ipset.name
  name        = "tf_csvserver"
  ipv46 = "10.78.60.22"
  port        = 80
  servicetype = "HTTP"

}

resource "citrixadc_ipset" "tf_ipset" {
  name = "tf_test_ipset"
  nsip6binding = [
	citrixadc_nsip6.tf_nsip1.ipv6address
  ]

}

resource "citrixadc_nsip6" "tf_nsip1" {
    ipv6address = "2001:db8:100::fb/64"
    type = "VIP"
    icmp = "DISABLED"
}

resource "citrixadc_nsip6" "tf_nsip2" {
    ipv6address = "2001:db8:100::fc/64"
    type = "VIP"
    icmp = "DISABLED" 
}
`

const testAccIpset_swap_ipv6_step2 = `
resource "citrixadc_csvserver" "test_csvserver" {
  ipset = citrixadc_ipset.tf_ipset.name
  name        = "tf_csvserver"
  ipv46 = "10.78.60.22"
  port        = 80
  servicetype = "HTTP"

}

resource "citrixadc_ipset" "tf_ipset" {
  name = "tf_test_ipset"
  nsip6binding = [
	citrixadc_nsip6.tf_nsip2.ipv6address
  ]

}

resource "citrixadc_nsip6" "tf_nsip1" {
    ipv6address = "2001:db8:100::fb/64"
    type = "VIP"
    icmp = "DISABLED"
}

resource "citrixadc_nsip6" "tf_nsip2" {
    ipv6address = "2001:db8:100::fc/64"
    type = "VIP"
    icmp = "DISABLED" 
}
`

func TestAccIpset_ipv6_swaps(t *testing.T) {
	if isCpxRun {
		t.Skip("No support in CPX")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckIpsetDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccIpset_swap_ipv6_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpsetExist("citrixadc_ipset.tf_ipset", nil),
				),
			},
			{
				Config: testAccIpset_swap_ipv6_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpsetExist("citrixadc_ipset.tf_ipset", nil),
				),
			},
		},
	})
}
