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
	"strings"
	"testing"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
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

func TestAccNat64_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_nat64.tf_nat64"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNat64Destroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNat64_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckNat64Exist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Nat64.Type(), "tf_nat64"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccNat64_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckNat64Exist(resAddr, nil)),
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

func TestAccNat64_import(t *testing.T) {
	const resAddr = "citrixadc_nat64.tf_nat64"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNat64Destroy,
		Steps: []resource.TestStep{
			{Config: testAccNat64_add},
			{
				Config:                  testAccNat64_add,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccNat64_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckNat64Destroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccNat64_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckNat64Exist("citrixadc_nat64.tf_nat64", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccNat64_add,
				Check:                    resource.ComposeTestCheckFunc(testAccCheckNat64Exist("citrixadc_nat64.tf_nat64", nil)),
			},
		},
	})
}

// Unset test: netprofile is the only unset-eligible attribute (name is the key,
// acl6name is required). Step 1 sets netprofile to a non-default value; step 2
// removes it so the provider must issue a NITRO ?action=unset, reverting it on
// the appliance.
const testAccNat64_unset_step1 = `
	resource "citrixadc_nsacl6" "tf_nsacl6_unset" {
		acl6name   = "tf_nsacl6_unset"
		acl6action = "ALLOW"
		logstate   = "ENABLED"
		stateful   = "NO"
		ratelimit  = 120
		state      = "ENABLED"
		priority   = 20
		protocol   = "TCP"
	}
	resource "citrixadc_netprofile" "tf_netprofile_unset" {
		name                   = "tf_netprofile_unset"
		proxyprotocol          = "ENABLED"
		proxyprotocoltxversion = "V1"
	}
	resource "citrixadc_nat64" "tf_nat64_unset" {
		name       = "tf_nat64_unset"
		acl6name   = citrixadc_nsacl6.tf_nsacl6_unset.acl6name
		netprofile = citrixadc_netprofile.tf_netprofile_unset.name
	}
`

const testAccNat64_unset_step2 = `
	resource "citrixadc_nsacl6" "tf_nsacl6_unset" {
		acl6name   = "tf_nsacl6_unset"
		acl6action = "ALLOW"
		logstate   = "ENABLED"
		stateful   = "NO"
		ratelimit  = 120
		state      = "ENABLED"
		priority   = 20
		protocol   = "TCP"
	}
	resource "citrixadc_netprofile" "tf_netprofile_unset" {
		name                   = "tf_netprofile_unset"
		proxyprotocol          = "ENABLED"
		proxyprotocoltxversion = "V1"
	}
	resource "citrixadc_nat64" "tf_nat64_unset" {
		name     = "tf_nat64_unset"
		acl6name = citrixadc_nsacl6.tf_nsacl6_unset.acl6name
		# netprofile removed from config -> provider must unset it.
	}
`

func TestAccNat64_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNat64Destroy,
		Steps: []resource.TestStep{
			{
				// Non-default value is applied and persisted.
				Config: testAccNat64_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNat64Exist("citrixadc_nat64.tf_nat64_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_nat64.tf_nat64_unset", "netprofile", "tf_netprofile_unset"),
				),
			},
			{
				// Removing netprofile must unset it: state reverts to the NITRO
				// default (empty) and the implicit post-apply plan must be empty.
				Config: testAccNat64_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNat64Exist("citrixadc_nat64.tf_nat64_unset", nil),
					resource.TestCheckNoResourceAttr("citrixadc_nat64.tf_nat64_unset", "netprofile"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckNat64ADCValue("tf_nat64_unset", "netprofile", ""),
				),
			},
		},
	})
}

// testAccCheckNat64ADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually
// reverted it.
func testAccCheckNat64ADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Nat64.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("nat64 %s not found on appliance", name)
		}
		got := ""
		if v, ok := data[attr]; ok && v != nil {
			got = strings.TrimSpace(fmt.Sprintf("%v", v))
		}
		if got != want {
			return fmt.Errorf("nat64 %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

const testAccNat64DataSource_basic = `
	resource "citrixadc_nsacl6" "tf_nsacl6_ds" {
		acl6name   = "tf_nsacl6_ds"
		acl6action = "ALLOW"
		logstate   = "ENABLED"
		stateful   = "NO"
		ratelimit  = 120
		state      = "ENABLED"
		priority   = 20
		protocol   = "TCP"
	}
	resource "citrixadc_netprofile" "tf_netprofile_ds" {
		name                   = "tf_netprofile_ds"
		proxyprotocol          = "ENABLED"
		proxyprotocoltxversion = "V1"
	}
	resource "citrixadc_nat64" "tf_nat64_ds" {
		name       = "tf_nat64_ds"
		acl6name   = citrixadc_nsacl6.tf_nsacl6_ds.acl6name
		netprofile = citrixadc_netprofile.tf_netprofile_ds.name
	}
	data "citrixadc_nat64" "tf_nat64_ds_data" {
		name = citrixadc_nat64.tf_nat64_ds.name
	}
`

func TestAccNat64DataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNat64DataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_nat64.tf_nat64_ds_data", "name", "tf_nat64_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_nat64.tf_nat64_ds_data", "acl6name", "tf_nsacl6_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_nat64.tf_nat64_ds_data", "netprofile", "tf_netprofile_ds"),
				),
			},
		},
	})
}
