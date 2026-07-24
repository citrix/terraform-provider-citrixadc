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

func testAccCheckAuthenticationazurekeyvaultExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No authenticationazurekeyvault name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Authenticationazurekeyvault.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("authenticationazurekeyvault %s not found", n)
		}

		return nil
	}
}

func testAccCheckAuthenticationazurekeyvaultDestroy(s *terraform.State) error {
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_authenticationazurekeyvault" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Authenticationazurekeyvault.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("authenticationazurekeyvault %s still exists", rs.Primary.ID)
		}
	}

	return nil
}

const testAccAuthenticationazurekeyvaultDataSource_basic = `

variable "authenticationazurekeyvault_clientsecret_wo" {
	type      = string
	sensitive = true
}

resource "citrixadc_authenticationazurekeyvault" "tf_authenticationazurekeyvault" {
  name                    = "tf_authenticationazurekeyvault"
  clientid                = "<clientid>"
  clientsecret_wo         = var.authenticationazurekeyvault_clientsecret_wo
  clientsecret_wo_version = 1
  servicekeyname          = "TestKey"
  vaultname               = "https://tfadctest.vault.azure.net/"
  authentication          = "ENABLED"
  refreshinterval         = 50
}

data "citrixadc_authenticationazurekeyvault" "tf_authenticationazurekeyvault" {
  name       = citrixadc_authenticationazurekeyvault.tf_authenticationazurekeyvault.name
  depends_on = [citrixadc_authenticationazurekeyvault.tf_authenticationazurekeyvault]
}
`

func TestAccAuthenticationazurekeyvaultDataSource_basic(t *testing.T) {
	t.Skip("Requires valid Azure credentials")
	t.Setenv("TF_VAR_authenticationazurekeyvault_clientsecret_wo", "<clientsecret>")
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationazurekeyvaultDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_authenticationazurekeyvault.tf_authenticationazurekeyvault", "name", "tf_authenticationazurekeyvault"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationazurekeyvault.tf_authenticationazurekeyvault", "clientid", "<clientid>"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationazurekeyvault.tf_authenticationazurekeyvault", "servicekeyname", "TestKey"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationazurekeyvault.tf_authenticationazurekeyvault", "vaultname", "https://tfadctest.vault.azure.net/"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationazurekeyvault.tf_authenticationazurekeyvault", "authentication", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationazurekeyvault.tf_authenticationazurekeyvault", "refreshinterval", "50"),
				),
			},
		},
	})
}

// Test backward-compatible path: using clientsecret (Sensitive attribute)
const testAccAuthenticationazurekeyvault_clientsecret_step1 = `

	variable "authenticationazurekeyvault_clientsecret" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_authenticationazurekeyvault" "tf_authenticationazurekeyvault" {
		name           = "tf_authenticationazurekeyvault"
		clientid       = "<clientid>"
		clientsecret   = var.authenticationazurekeyvault_clientsecret
		servicekeyname = "TestKey"
		vaultname      = "https://tfadctest.vault.azure.net/"
	}
`

// Update backward-compatible path: change clientsecret value
const testAccAuthenticationazurekeyvault_clientsecret_step2 = `

	 variable "authenticationazurekeyvault_clientsecret_2" {
	  type      = string
	  sensitive = true
	}

	 resource "citrixadc_authenticationazurekeyvault" "tf_authenticationazurekeyvault" {
		name           = "tf_authenticationazurekeyvault"
		clientid       = "<clientid>"
		clientsecret   = var.authenticationazurekeyvault_clientsecret_2
		servicekeyname = "TestKey"
		vaultname      = "https://tfadcnew.vault.azure.net/"
	}
`

func TestAccAuthenticationazurekeyvault_import(t *testing.T) {
	t.Skip("TODO: Requires review")
	const resAddr = "citrixadc_authenticationazurekeyvault.tf_authenticationazurekeyvault"
	t.Setenv("TF_VAR_authenticationazurekeyvault_clientsecret", "<clientsecret>")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationazurekeyvaultDestroy,
		Steps: []resource.TestStep{
			{Config: testAccAuthenticationazurekeyvault_clientsecret_step1},
			{
				Config:                  testAccAuthenticationazurekeyvault_clientsecret_step1,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccAuthenticationazurekeyvault_clientsecret_backward_compat(t *testing.T) {
	t.Setenv("TF_VAR_authenticationazurekeyvault_clientsecret", "<clientsecret>")
	t.Setenv("TF_VAR_authenticationazurekeyvault_clientsecret_2", "<clientsecret_2>")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationazurekeyvaultDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationazurekeyvault_clientsecret_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationazurekeyvaultExist("citrixadc_authenticationazurekeyvault.tf_authenticationazurekeyvault", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationazurekeyvault.tf_authenticationazurekeyvault", "name", "tf_authenticationazurekeyvault"),
					resource.TestCheckResourceAttr("citrixadc_authenticationazurekeyvault.tf_authenticationazurekeyvault", "clientid", "<clientid>"),
				),
			},
			{
				Config: testAccAuthenticationazurekeyvault_clientsecret_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationazurekeyvaultExist("citrixadc_authenticationazurekeyvault.tf_authenticationazurekeyvault", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationazurekeyvault.tf_authenticationazurekeyvault", "name", "tf_authenticationazurekeyvault"),
					resource.TestCheckResourceAttr("citrixadc_authenticationazurekeyvault.tf_authenticationazurekeyvault", "clientid", "<clientid>"),
				),
			},
		},
	})
}

// Test ephemeral path: using clientsecret_wo (WriteOnly attribute) with version tracker
const testAccAuthenticationazurekeyvault_clientsecret_wo_step1 = `

	variable "authenticationazurekeyvault_clientsecret_wo" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_authenticationazurekeyvault" "tf_authenticationazurekeyvault" {
		name                    = "tf_authenticationazurekeyvault"
		clientid                = "<clientid>"
		clientsecret_wo         = var.authenticationazurekeyvault_clientsecret_wo
		clientsecret_wo_version = 1
		servicekeyname          = "TestKey"
		vaultname               = "https://tfadctest.vault.azure.net/"
	}
`

// Update ephemeral path: bump version to trigger update with new clientsecret
const testAccAuthenticationazurekeyvault_clientsecret_wo_step2 = `

	 variable "authenticationazurekeyvault_clientsecret_wo_2" {
	  type      = string
	  sensitive = true
	}

	 resource "citrixadc_authenticationazurekeyvault" "tf_authenticationazurekeyvault" {
		name                    = "tf_authenticationazurekeyvault"
		clientid                = "<clientid>"
		clientsecret_wo         = var.authenticationazurekeyvault_clientsecret_wo_2
		clientsecret_wo_version = 2
		servicekeyname          = "TestKey"
		vaultname               = "https://tfadctest.vault.azure.net/"
	}
`

func TestAccAuthenticationazurekeyvault_clientsecret_wo_ephemeral(t *testing.T) {
	t.Skip("Requires valid Azure credentials")
	t.Setenv("TF_VAR_authenticationazurekeyvault_clientsecret_wo", "<clientsecret_wo>")
	t.Setenv("TF_VAR_authenticationazurekeyvault_clientsecret_wo_2", "<clientsecret_wo_2>")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationazurekeyvaultDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationazurekeyvault_clientsecret_wo_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationazurekeyvaultExist("citrixadc_authenticationazurekeyvault.tf_authenticationazurekeyvault", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationazurekeyvault.tf_authenticationazurekeyvault", "clientsecret_wo_version", "1"),
					resource.TestCheckResourceAttr("citrixadc_authenticationazurekeyvault.tf_authenticationazurekeyvault", "name", "tf_authenticationazurekeyvault"),
					resource.TestCheckResourceAttr("citrixadc_authenticationazurekeyvault.tf_authenticationazurekeyvault", "clientid", "<clientid>"),
				),
			},
			{
				Config: testAccAuthenticationazurekeyvault_clientsecret_wo_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationazurekeyvaultExist("citrixadc_authenticationazurekeyvault.tf_authenticationazurekeyvault", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationazurekeyvault.tf_authenticationazurekeyvault", "clientsecret_wo_version", "2"),
					resource.TestCheckResourceAttr("citrixadc_authenticationazurekeyvault.tf_authenticationazurekeyvault", "name", "tf_authenticationazurekeyvault"),
					resource.TestCheckResourceAttr("citrixadc_authenticationazurekeyvault.tf_authenticationazurekeyvault", "clientid", "<clientid>"),
				),
			},
		},
	})
}

// ---------------------------------------------------------------------------
// Unset acceptance test
//
// Eligible (unset-wired) attributes: authentication, refreshinterval, signaturealg.
// Step 1 sets them to non-default values that apply on the appliance; step 2 removes
// them from config so the provider issues ?action=unset and each reverts to its NITRO
// default (authentication=ENABLED, refreshinterval=50, signaturealg=RS256), with an
// empty post-apply plan.
//
// Note: signaturealg's only ADC-accepted value is "RS256", which equals its default,
// so it cannot be varied to a distinct non-default; it is still exercised through the
// removal path to confirm no perpetual diff.
// ---------------------------------------------------------------------------

const testAccAuthenticationazurekeyvault_unset_step1 = `
resource "citrixadc_authenticationazurekeyvault" "tf_unset" {
  name            = "tf_test_authenticationazurekeyvault_unset"
  clientid        = "00000000-0000-0000-0000-000000000000"
  clientsecret    = "tf_test_unset_secret"
  servicekeyname  = "TestKey"
  vaultname       = "https://tfadctest.vault.azure.net/"
  authentication  = "DISABLED"
  refreshinterval = 100
  signaturealg    = "RS256"
}
`

const testAccAuthenticationazurekeyvault_unset_step2 = `
resource "citrixadc_authenticationazurekeyvault" "tf_unset" {
  name           = "tf_test_authenticationazurekeyvault_unset"
  clientid       = "00000000-0000-0000-0000-000000000000"
  clientsecret   = "tf_test_unset_secret"
  servicekeyname = "TestKey"
  vaultname      = "https://tfadctest.vault.azure.net/"
  # authentication, refreshinterval, signaturealg removed -> provider must unset them
}
`

func TestAccAuthenticationazurekeyvault_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationazurekeyvaultDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values apply and persist.
				Config: testAccAuthenticationazurekeyvault_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationazurekeyvaultExist("citrixadc_authenticationazurekeyvault.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationazurekeyvault.tf_unset", "authentication", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_authenticationazurekeyvault.tf_unset", "refreshinterval", "100"),
					resource.TestCheckResourceAttr("citrixadc_authenticationazurekeyvault.tf_unset", "signaturealg", "RS256"),
				),
			},
			{
				// Removing them must unset -> state reverts to NITRO defaults,
				// and the implicit post-apply plan must be empty.
				Config: testAccAuthenticationazurekeyvault_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationazurekeyvaultExist("citrixadc_authenticationazurekeyvault.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationazurekeyvault.tf_unset", "authentication", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_authenticationazurekeyvault.tf_unset", "refreshinterval", "50"),
					resource.TestCheckResourceAttr("citrixadc_authenticationazurekeyvault.tf_unset", "signaturealg", "RS256"),
					// Independent appliance-level confirmation the unset took effect:
					testAccCheckAuthenticationazurekeyvaultADCValue("tf_test_authenticationazurekeyvault_unset", "authentication", "ENABLED"),
					testAccCheckAuthenticationazurekeyvaultADCValue("tf_test_authenticationazurekeyvault_unset", "refreshinterval", "50"),
					testAccCheckAuthenticationazurekeyvaultADCValue("tf_test_authenticationazurekeyvault_unset", "signaturealg", "RS256"),
				),
			},
		},
	})
}

// testAccCheckAuthenticationazurekeyvaultADCValue asserts an attribute's value directly
// on the appliance (not just in Terraform state), proving the unset actually reverted it.
func testAccCheckAuthenticationazurekeyvaultADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Authenticationazurekeyvault.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("authenticationazurekeyvault %s not found on appliance", name)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("authenticationazurekeyvault %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}
