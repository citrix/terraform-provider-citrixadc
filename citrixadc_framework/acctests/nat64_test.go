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

const testAccNat64_add = `
	resource "citrixadc_nsacl6" "tf_nsacl6" {
		acl6name   = "tf_nsacl6"
		acl6action = "ALLOW"
		logstate   = "ENABLED"
		stateful   = "NO"
		ratelimit  = 120
		state      = "ENABLED"
		priority   = 20
		protocol   = "TCP"
	}
	resource "citrixadc_netprofile" "tf_netprofile" {
		name                   = "tf_netprofile"
		proxyprotocol          = "ENABLED"
		proxyprotocoltxversion = "V1"
	}
	resource "citrixadc_netprofile" "tf_netprofile1" {
		name                   = "tf_netprofile1"
		proxyprotocol          = "ENABLED"
		proxyprotocoltxversion = "V1"
	}
	resource "citrixadc_nat64" "tf_nat64" {
		name       = "tf_nat64"
		acl6name   = citrixadc_nsacl6.tf_nsacl6.acl6name
		netprofile = citrixadc_netprofile.tf_netprofile.name
	}
`
const testAccNat64_update = `
	resource "citrixadc_nsacl6" "tf_nsacl6" {
		acl6name   = "tf_nsacl6"
		acl6action = "ALLOW"
		logstate   = "ENABLED"
		stateful   = "NO"
		ratelimit  = 120
		state      = "ENABLED"
		priority   = 20
		protocol   = "TCP"
	}
	resource "citrixadc_netprofile" "tf_netprofile" {
		name                   = "tf_netprofile"
		proxyprotocol          = "ENABLED"
		proxyprotocoltxversion = "V1"
	}
	resource "citrixadc_netprofile" "tf_netprofile1" {
		name                   = "tf_netprofile1"
		proxyprotocol          = "ENABLED"
		proxyprotocoltxversion = "V1"
	}
	resource "citrixadc_nat64" "tf_nat64" {
		name       = "tf_nat64"
		acl6name   = citrixadc_nsacl6.tf_nsacl6.acl6name
		netprofile = citrixadc_netprofile.tf_netprofile1.name
	}
`

func TestAccNat64_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNat64Destroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNat64_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNat64Exist("citrixadc_nat64.tf_nat64", nil),
					resource.TestCheckResourceAttr("citrixadc_nat64.tf_nat64", "name", "tf_nat64"),
					resource.TestCheckResourceAttr("citrixadc_nat64.tf_nat64", "netprofile", "tf_netprofile"),
				),
			},
			{
				Config: testAccNat64_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNat64Exist("citrixadc_nat64.tf_nat64", nil),
					resource.TestCheckResourceAttr("citrixadc_nat64.tf_nat64", "name", "tf_nat64"),
					resource.TestCheckResourceAttr("citrixadc_nat64.tf_nat64", "netprofile", "tf_netprofile1"),
				),
			},
		},
	})
}

func testAccCheckNat64Exist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No nat64 name is set")
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
		data, err := client.FindResource(service.Nat64.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("nat64 %s not found", n)
		}

		return nil
	}
}

func testAccCheckNat64Destroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_nat64" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Nat64.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("nat64 %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
