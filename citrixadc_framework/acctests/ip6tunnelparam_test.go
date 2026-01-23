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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"testing"
)

const testAccIp6tunnelparam_add = `
	resource "citrixadc_nsip6" "name" {
		ipv6address = "2001:db8:100::fa/64"
		type        = "VIP"
		icmp        = "DISABLED"
	}
	resource "citrixadc_ip6tunnelparam" "tf_ip6tunnelparam" {
		srcip                = split("/",citrixadc_nsip6.name.ipv6address)[0]
		dropfrag             = "YES"
		dropfragcputhreshold = 3
		srciproundrobin      = "YES"
		useclientsourceipv6  = "NO"
	}
`
const testAccIp6tunnelparam_update = `
	resource "citrixadc_nsip6" "name" {
		ipv6address = "2001:db8:100::fa/64"
		type        = "VIP"
		icmp        = "DISABLED"
	}
	resource "citrixadc_ip6tunnelparam" "tf_ip6tunnelparam" {
		srcip                = split("/",citrixadc_nsip6.name.ipv6address)[0]
		dropfrag             = "NO"
		dropfragcputhreshold = 1
		srciproundrobin      = "NO"
		useclientsourceipv6  = "NO"
	}
`

func TestAccIp6tunnelparam_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccIp6tunnelparam_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIp6tunnelparamExist("citrixadc_ip6tunnelparam.tf_ip6tunnelparam", nil),
					resource.TestCheckResourceAttr("citrixadc_ip6tunnelparam.tf_ip6tunnelparam", "srcip", "2001:db8:100::fa"),
					resource.TestCheckResourceAttr("citrixadc_ip6tunnelparam.tf_ip6tunnelparam", "dropfrag", "YES"),
					resource.TestCheckResourceAttr("citrixadc_ip6tunnelparam.tf_ip6tunnelparam", "dropfragcputhreshold", "3"),
					resource.TestCheckResourceAttr("citrixadc_ip6tunnelparam.tf_ip6tunnelparam", "srciproundrobin", "YES"),
				),
			},
			{
				Config: testAccIp6tunnelparam_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIp6tunnelparamExist("citrixadc_ip6tunnelparam.tf_ip6tunnelparam", nil),
					resource.TestCheckResourceAttr("citrixadc_ip6tunnelparam.tf_ip6tunnelparam", "srcip", "2001:db8:100::fa"),
					resource.TestCheckResourceAttr("citrixadc_ip6tunnelparam.tf_ip6tunnelparam", "dropfrag", "NO"),
					resource.TestCheckResourceAttr("citrixadc_ip6tunnelparam.tf_ip6tunnelparam", "dropfragcputhreshold", "1"),
					resource.TestCheckResourceAttr("citrixadc_ip6tunnelparam.tf_ip6tunnelparam", "srciproundrobin", "NO"),
				),
			},
		},
	})
}

func testAccCheckIp6tunnelparamExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ip6tunnelparam name is set")
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
		data, err := client.FindResource(service.Ip6tunnelparam.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("ip6tunnelparam %s not found", n)
		}

		return nil
	}
}
