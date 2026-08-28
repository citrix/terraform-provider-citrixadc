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
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const testAccVpnsecureprivateaccessprofile_basic = `


resource "citrixadc_vpnsecureprivateaccessprofile" "tf_vpnsecureprivateaccessprofile" {
	name                        = "my_spaprofile"
	url                         = "https://spa.example.com"
	forceclienttype             = "ON"
	chromeenterprisepremiummode = "OFF"
}

`
const testAccVpnsecureprivateaccessprofile_update = `


resource "citrixadc_vpnsecureprivateaccessprofile" "tf_vpnsecureprivateaccessprofile" {
	name                        = "my_spaprofile"
	url                         = "https://spa-updated.example.com"
	forceclienttype             = "OFF"
	chromeenterprisepremiummode = "WITHOUT_PARTNER_CONNECTOR"
}

`

func TestAccVpnsecureprivateaccessprofile_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVpnsecureprivateaccessprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnsecureprivateaccessprofile_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnsecureprivateaccessprofileExist("citrixadc_vpnsecureprivateaccessprofile.tf_vpnsecureprivateaccessprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnsecureprivateaccessprofile.tf_vpnsecureprivateaccessprofile", "name", "my_spaprofile"),
					resource.TestCheckResourceAttr("citrixadc_vpnsecureprivateaccessprofile.tf_vpnsecureprivateaccessprofile", "url", "https://spa.example.com"),
					resource.TestCheckResourceAttr("citrixadc_vpnsecureprivateaccessprofile.tf_vpnsecureprivateaccessprofile", "forceclienttype", "ON"),
					resource.TestCheckResourceAttr("citrixadc_vpnsecureprivateaccessprofile.tf_vpnsecureprivateaccessprofile", "chromeenterprisepremiummode", "OFF"),
				),
			},
			{
				Config: testAccVpnsecureprivateaccessprofile_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnsecureprivateaccessprofileExist("citrixadc_vpnsecureprivateaccessprofile.tf_vpnsecureprivateaccessprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnsecureprivateaccessprofile.tf_vpnsecureprivateaccessprofile", "name", "my_spaprofile"),
					resource.TestCheckResourceAttr("citrixadc_vpnsecureprivateaccessprofile.tf_vpnsecureprivateaccessprofile", "url", "https://spa-updated.example.com"),
					resource.TestCheckResourceAttr("citrixadc_vpnsecureprivateaccessprofile.tf_vpnsecureprivateaccessprofile", "forceclienttype", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_vpnsecureprivateaccessprofile.tf_vpnsecureprivateaccessprofile", "chromeenterprisepremiummode", "WITHOUT_PARTNER_CONNECTOR"),
				),
			},
		},
	})
}

func TestAccVpnsecureprivateaccessprofile_import(t *testing.T) {
	const resAddr = "citrixadc_vpnsecureprivateaccessprofile.tf_vpnsecureprivateaccessprofile"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVpnsecureprivateaccessprofileDestroy,
		Steps: []resource.TestStep{
			{Config: testAccVpnsecureprivateaccessprofile_basic},
			{
				Config:            testAccVpnsecureprivateaccessprofile_basic,
				ResourceName:      resAddr,
				ImportState:       true,
				ImportStateVerify: true,
				// sharedsecret_wo_version is a write-only version tracker (defaulted to 1) that
				// is not returned by NITRO, so it cannot round-trip through import.
				ImportStateVerifyIgnore: []string{"sharedsecret_wo_version"},
			},
		},
	})
}

func testAccCheckVpnsecureprivateaccessprofileExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No vpnsecureprivateaccessprofile name is set")
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
		data, err := client.FindResource("vpnsecureprivateaccessprofile", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("vpnsecureprivateaccessprofile %s not found", n)
		}

		return nil
	}
}

func testAccCheckVpnsecureprivateaccessprofileDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_vpnsecureprivateaccessprofile" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("vpnsecureprivateaccessprofile", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("vpnsecureprivateaccessprofile %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccVpnsecureprivateaccessprofileDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnsecureprivateaccessprofileDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_vpnsecureprivateaccessprofile.test", "name", "tf_ds_spaprofile"),
					resource.TestCheckResourceAttrSet("data.citrixadc_vpnsecureprivateaccessprofile.test", "id"),
					resource.TestCheckResourceAttr("data.citrixadc_vpnsecureprivateaccessprofile.test", "forceclienttype", "ON"),
				),
			},
		},
	})
}

const testAccVpnsecureprivateaccessprofileDataSource_basic = `
resource "citrixadc_vpnsecureprivateaccessprofile" "tf_ds_spaprofile" {
	name            = "tf_ds_spaprofile"
	url             = "https://spa-ds.example.com"
	forceclienttype = "ON"
}

data "citrixadc_vpnsecureprivateaccessprofile" "test" {
	name = citrixadc_vpnsecureprivateaccessprofile.tf_ds_spaprofile.name
}
`

// Test backward-compatible path: using sharedsecret (Sensitive attribute).
// sharedsecret has a minimum length of 32.
const testAccVpnsecureprivateaccessprofile_sharedsecret_step1 = `
	variable "spa_sharedsecret" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_vpnsecureprivateaccessprofile" "tf_spaprofile_ephem" {
		name            = "tf_spaprofile_ephem"
		url             = "https://spa-ephem.example.com"
		forceclienttype = "ON"
		sharedsecret    = var.spa_sharedsecret
	}
`

const testAccVpnsecureprivateaccessprofile_sharedsecret_step2 = `
	variable "spa_sharedsecret_2" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_vpnsecureprivateaccessprofile" "tf_spaprofile_ephem" {
		name            = "tf_spaprofile_ephem"
		url             = "https://spa-ephem.example.com"
		forceclienttype = "ON"
		sharedsecret    = var.spa_sharedsecret_2
	}
`

func TestAccVpnsecureprivateaccessprofile_sharedsecret_backward_compat(t *testing.T) {
	t.Setenv("TF_VAR_spa_sharedsecret", "0123456789abcdef0123456789abcdef")
	t.Setenv("TF_VAR_spa_sharedsecret_2", "fedcba9876543210fedcba9876543210")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVpnsecureprivateaccessprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnsecureprivateaccessprofile_sharedsecret_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnsecureprivateaccessprofileExist("citrixadc_vpnsecureprivateaccessprofile.tf_spaprofile_ephem", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnsecureprivateaccessprofile.tf_spaprofile_ephem", "name", "tf_spaprofile_ephem"),
					resource.TestCheckResourceAttr("citrixadc_vpnsecureprivateaccessprofile.tf_spaprofile_ephem", "forceclienttype", "ON"),
				),
			},
			{
				Config: testAccVpnsecureprivateaccessprofile_sharedsecret_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnsecureprivateaccessprofileExist("citrixadc_vpnsecureprivateaccessprofile.tf_spaprofile_ephem", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnsecureprivateaccessprofile.tf_spaprofile_ephem", "name", "tf_spaprofile_ephem"),
					resource.TestCheckResourceAttr("citrixadc_vpnsecureprivateaccessprofile.tf_spaprofile_ephem", "forceclienttype", "ON"),
				),
			},
		},
	})
}

// Test ephemeral path: using sharedsecret_wo (WriteOnly attribute) with version tracker.
const testAccVpnsecureprivateaccessprofile_sharedsecret_wo_step1 = `
	variable "spa_sharedsecret_wo" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_vpnsecureprivateaccessprofile" "tf_spaprofile_ephem" {
		name                    = "tf_spaprofile_ephem"
		url                     = "https://spa-ephem.example.com"
		forceclienttype         = "ON"
		sharedsecret_wo         = var.spa_sharedsecret_wo
		sharedsecret_wo_version = 1
	}
`

const testAccVpnsecureprivateaccessprofile_sharedsecret_wo_step2 = `
	variable "spa_sharedsecret_wo_2" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_vpnsecureprivateaccessprofile" "tf_spaprofile_ephem" {
		name                    = "tf_spaprofile_ephem"
		url                     = "https://spa-ephem.example.com"
		forceclienttype         = "ON"
		sharedsecret_wo         = var.spa_sharedsecret_wo_2
		sharedsecret_wo_version = 2
	}
`

func TestAccVpnsecureprivateaccessprofile_sharedsecret_wo_ephemeral(t *testing.T) {
	t.Setenv("TF_VAR_spa_sharedsecret_wo", "0123456789abcdef0123456789abcdef")
	t.Setenv("TF_VAR_spa_sharedsecret_wo_2", "fedcba9876543210fedcba9876543210")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVpnsecureprivateaccessprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnsecureprivateaccessprofile_sharedsecret_wo_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnsecureprivateaccessprofileExist("citrixadc_vpnsecureprivateaccessprofile.tf_spaprofile_ephem", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnsecureprivateaccessprofile.tf_spaprofile_ephem", "name", "tf_spaprofile_ephem"),
					resource.TestCheckResourceAttr("citrixadc_vpnsecureprivateaccessprofile.tf_spaprofile_ephem", "sharedsecret_wo_version", "1"),
				),
			},
			{
				Config: testAccVpnsecureprivateaccessprofile_sharedsecret_wo_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnsecureprivateaccessprofileExist("citrixadc_vpnsecureprivateaccessprofile.tf_spaprofile_ephem", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnsecureprivateaccessprofile.tf_spaprofile_ephem", "name", "tf_spaprofile_ephem"),
					resource.TestCheckResourceAttr("citrixadc_vpnsecureprivateaccessprofile.tf_spaprofile_ephem", "sharedsecret_wo_version", "2"),
				),
			},
		},
	})
}

// Step 1: unset-eligible enum attributes set to a valid non-default value.
const testAccVpnsecureprivateaccessprofile_unset_step1 = `
resource "citrixadc_vpnsecureprivateaccessprofile" "tf_unset" {
	name                        = "tf_test_spaprofile_unset"
	url                         = "https://spa-unset.example.com"
	forceclienttype             = "OFF"
	chromeenterprisepremiummode = "WITHOUT_PARTNER_CONNECTOR"
}
`

// Step 2: eligible attributes removed from config -> provider must unset them,
// reverting each to its NITRO default (forceclienttype=ON, chromeenterprisepremiummode=OFF).
const testAccVpnsecureprivateaccessprofile_unset_step2 = `
resource "citrixadc_vpnsecureprivateaccessprofile" "tf_unset" {
	name = "tf_test_spaprofile_unset"
	url  = "https://spa-unset.example.com"
}
`

func TestAccVpnsecureprivateaccessprofile_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVpnsecureprivateaccessprofileDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values apply and persist.
				Config: testAccVpnsecureprivateaccessprofile_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnsecureprivateaccessprofileExist("citrixadc_vpnsecureprivateaccessprofile.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnsecureprivateaccessprofile.tf_unset", "forceclienttype", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_vpnsecureprivateaccessprofile.tf_unset", "chromeenterprisepremiummode", "WITHOUT_PARTNER_CONNECTOR"),
				),
			},
			{
				// Removing them must unset -> state reverts to NITRO defaults.
				Config: testAccVpnsecureprivateaccessprofile_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnsecureprivateaccessprofileExist("citrixadc_vpnsecureprivateaccessprofile.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnsecureprivateaccessprofile.tf_unset", "forceclienttype", "ON"),
					resource.TestCheckResourceAttr("citrixadc_vpnsecureprivateaccessprofile.tf_unset", "chromeenterprisepremiummode", "OFF"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckVpnsecureprivateaccessprofileADCValue("tf_test_spaprofile_unset", "forceclienttype", "ON"),
					testAccCheckVpnsecureprivateaccessprofileADCValue("tf_test_spaprofile_unset", "chromeenterprisepremiummode", "OFF"),
				),
			},
		},
	})
}

// testAccCheckVpnsecureprivateaccessprofileADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted it.
func testAccCheckVpnsecureprivateaccessprofileADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Vpnsecureprivateaccessprofile.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("vpnsecureprivateaccessprofile %s not found on appliance", name)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("vpnsecureprivateaccessprofile %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

// TestAccVpnsecureprivateaccessprofile_selfHealing verifies the provider re-creates the profile
// when it is deleted out-of-band between apply steps (drift recovery).
func TestAccVpnsecureprivateaccessprofile_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_vpnsecureprivateaccessprofile.tf_vpnsecureprivateaccessprofile"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVpnsecureprivateaccessprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnsecureprivateaccessprofile_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckVpnsecureprivateaccessprofileExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Vpnsecureprivateaccessprofile.Type(), "my_spaprofile"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccVpnsecureprivateaccessprofile_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckVpnsecureprivateaccessprofileExist(resAddr, nil)),
			},
		},
	})
}
