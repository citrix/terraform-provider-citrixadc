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

const testAccNat64param_add = `
	resource "citrixadc_nat64param" "tf_nat64param" {
		nat64ignoretos    = "YES"
		nat64zerochecksum = "DISABLED"
		nat64v6mtu        = 1282
		nat64fragheader   = "DISABLED"
	}
`
const testAccNat64param_update = `
	resource "citrixadc_nat64param" "tf_nat64param" {
		nat64ignoretos    = "NO"
		nat64zerochecksum = "ENABLED"
		nat64v6mtu        = 1280
		nat64fragheader   = "ENABLED"
	}
`

func TestAccNat64param_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccNat64param_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNat64paramExist("citrixadc_nat64param.tf_nat64param", nil),
					resource.TestCheckResourceAttr("citrixadc_nat64param.tf_nat64param", "nat64ignoretos", "YES"),
					resource.TestCheckResourceAttr("citrixadc_nat64param.tf_nat64param", "nat64zerochecksum", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_nat64param.tf_nat64param", "nat64v6mtu", "1282"),
					resource.TestCheckResourceAttr("citrixadc_nat64param.tf_nat64param", "nat64fragheader", "DISABLED"),
				),
			},
			{
				Config: testAccNat64param_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNat64paramExist("citrixadc_nat64param.tf_nat64param", nil),
					resource.TestCheckResourceAttr("citrixadc_nat64param.tf_nat64param", "nat64ignoretos", "NO"),
					resource.TestCheckResourceAttr("citrixadc_nat64param.tf_nat64param", "nat64zerochecksum", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_nat64param.tf_nat64param", "nat64v6mtu", "1280"),
					resource.TestCheckResourceAttr("citrixadc_nat64param.tf_nat64param", "nat64fragheader", "ENABLED"),
				),
			},
		},
	})
}

func testAccCheckNat64paramExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No nat64param name is set")
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
		data, err := client.FindResource("nat64param", "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("nat64param %s not found", n)
		}

		return nil
	}
}

func TestAccNat64param_import(t *testing.T) {
	const resAddr = "citrixadc_nat64param.tf_nat64param"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{Config: testAccNat64param_add},
			{
				Config:                  testAccNat64param_add,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccNat64param_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccNat64param_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckNat64paramExist("citrixadc_nat64param.tf_nat64param", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccNat64param_add,
				Check:                    resource.ComposeTestCheckFunc(testAccCheckNat64paramExist("citrixadc_nat64param.tf_nat64param", nil)),
			},
		},
	})
}

// nat64param unset test: step1 sets all mutable attributes to non-default
// values; step2 removes them so the provider must unset them (revert to the
// documented NITRO defaults).
const testAccNat64param_unset_step1 = `
	resource "citrixadc_nat64param" "tf_unset" {
		nat64ignoretos    = "YES"
		nat64zerochecksum = "DISABLED"
		nat64v6mtu        = 1282
		nat64fragheader   = "DISABLED"
	}
`

const testAccNat64param_unset_step2 = `
	resource "citrixadc_nat64param" "tf_unset" {
		# All unset-eligible attributes removed from config -> the provider must
		# unset them (revert to NITRO defaults).
	}
`

func TestAccNat64param_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccNat64param_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNat64paramExist("citrixadc_nat64param.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_nat64param.tf_unset", "nat64ignoretos", "YES"),
					resource.TestCheckResourceAttr("citrixadc_nat64param.tf_unset", "nat64zerochecksum", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_nat64param.tf_unset", "nat64v6mtu", "1282"),
					resource.TestCheckResourceAttr("citrixadc_nat64param.tf_unset", "nat64fragheader", "DISABLED"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from
				// the appliance) reverts to the documented NITRO defaults, and the
				// implicit post-apply plan must be empty.
				Config: testAccNat64param_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNat64paramExist("citrixadc_nat64param.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_nat64param.tf_unset", "nat64ignoretos", "NO"),
					resource.TestCheckResourceAttr("citrixadc_nat64param.tf_unset", "nat64zerochecksum", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_nat64param.tf_unset", "nat64v6mtu", "1280"),
					resource.TestCheckResourceAttr("citrixadc_nat64param.tf_unset", "nat64fragheader", "ENABLED"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckNat64paramADCValue("nat64ignoretos", "NO"),
					testAccCheckNat64paramADCValue("nat64zerochecksum", "ENABLED"),
					testAccCheckNat64paramADCValue("nat64v6mtu", "1280"),
					testAccCheckNat64paramADCValue("nat64fragheader", "ENABLED"),
				),
			},
		},
	})
}

// testAccCheckNat64paramADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted it.
func testAccCheckNat64paramADCValue(attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Nat64param.Type(), "")
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("nat64param not found on appliance")
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("nat64param: appliance attr %q = %q, want %q (unset did not revert it)", attr, got, want)
		}
		return nil
	}
}

const testAccNat64paramDataSource_basic = `
	resource "citrixadc_nat64param" "tf_nat64param" {
		nat64ignoretos    = "YES"
		nat64zerochecksum = "DISABLED"
		nat64v6mtu        = 1282
		nat64fragheader   = "DISABLED"
	}

	data "citrixadc_nat64param" "tf_nat64param_ds" {
		td = citrixadc_nat64param.tf_nat64param.td
	}
`

func TestAccNat64paramDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccNat64paramDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_nat64param.tf_nat64param_ds", "nat64ignoretos", "YES"),
					resource.TestCheckResourceAttr("data.citrixadc_nat64param.tf_nat64param_ds", "nat64zerochecksum", "DISABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_nat64param.tf_nat64param_ds", "nat64v6mtu", "1282"),
					resource.TestCheckResourceAttr("data.citrixadc_nat64param.tf_nat64param_ds", "nat64fragheader", "DISABLED"),
				),
			},
		},
	})
}
