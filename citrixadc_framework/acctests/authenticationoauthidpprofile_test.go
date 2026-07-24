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

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testAccAuthenticationoauthidpprofile_add = `

	resource "citrixadc_authenticationoauthidpprofile" "tf_idpprofile" {
		name         = "tf_idpprofile"
		clientid     = "cliId"
		clientsecret = "secret"
		redirecturl  = "http://www.example.com/1/"
	}
`
const testAccAuthenticationoauthidpprofile_update = `

	resource "citrixadc_authenticationoauthidpprofile" "tf_idpprofile" {
		name          = "tf_idpprofile"
		clientid      = "cliId1"
		clientsecret  = "secret"
		redirecturl   = "http://www.example11.com/1/"
	}
`

func TestAccAuthenticationoauthidpprofile_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationoauthidpprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationoauthidpprofile_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationoauthidpprofileExist("citrixadc_authenticationoauthidpprofile.tf_idpprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationoauthidpprofile.tf_idpprofile", "name", "tf_idpprofile"),
					resource.TestCheckResourceAttr("citrixadc_authenticationoauthidpprofile.tf_idpprofile", "clientid", "cliId"),
				),
			},
			{
				Config: testAccAuthenticationoauthidpprofile_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationoauthidpprofileExist("citrixadc_authenticationoauthidpprofile.tf_idpprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationoauthidpprofile.tf_idpprofile", "name", "tf_idpprofile"),
					resource.TestCheckResourceAttr("citrixadc_authenticationoauthidpprofile.tf_idpprofile", "clientid", "cliId1"),
				),
			},
		},
	})
}

func TestAccAuthenticationoauthidpprofile_import(t *testing.T) {
	const resAddr = "citrixadc_authenticationoauthidpprofile.tf_idpprofile"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationoauthidpprofileDestroy,
		Steps: []resource.TestStep{
			{Config: testAccAuthenticationoauthidpprofile_add},
			{
				Config:            testAccAuthenticationoauthidpprofile_add,
				ResourceName:      resAddr,
				ImportState:       true,
				ImportStateVerify: true,
				// clientsecret is not returned by the NITRO API (secret) and is only
				// retained from config, so it cannot round-trip through import.
				ImportStateVerifyIgnore: []string{"clientsecret"},
			},
		},
	})
}

func testAccCheckAuthenticationoauthidpprofileExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No authenticationoauthidpprofile name is set")
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
		data, err := client.FindResource("authenticationoauthidpprofile", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("authenticationoauthidpprofile %s not found", n)
		}

		return nil
	}
}

func testAccCheckAuthenticationoauthidpprofileDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_authenticationoauthidpprofile" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("authenticationoauthidpprofile", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("authenticationoauthidpprofile %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccAuthenticationoauthidpprofileDataSource_basic = `

	resource "citrixadc_authenticationoauthidpprofile" "tf_idpprofile_ds" {
		name         = "tf_idpprofile_ds"
		clientid     = "cliId_datasource"
		clientsecret = "secret"
		redirecturl  = "http://www.example.com/datasource/"
	}

	data "citrixadc_authenticationoauthidpprofile" "tf_idpprofile_ds" {
		name = citrixadc_authenticationoauthidpprofile.tf_idpprofile_ds.name
	}
`

func TestAccAuthenticationoauthidpprofileDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationoauthidpprofileDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_authenticationoauthidpprofile.tf_idpprofile_ds", "name", "tf_idpprofile_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationoauthidpprofile.tf_idpprofile_ds", "clientid", "cliId_datasource"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationoauthidpprofile.tf_idpprofile_ds", "redirecturl", "http://www.example.com/datasource/"),
				),
			},
		},
	})
}

// Test backward-compatible path: using clientsecret (Sensitive attribute)
const testAccAuthenticationoauthidpprofile_clientsecret_step1 = `

	variable "authenticationoauthidpprofile_clientsecret" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_authenticationoauthidpprofile" "test" {
		name         = "tf_idpprofile_compat"
		clientid     = "cliId_compat"
		clientsecret = var.authenticationoauthidpprofile_clientsecret
		redirecturl  = "http://www.example.com/compat/"
	}
`

// Update backward-compatible path: change clientsecret value
const testAccAuthenticationoauthidpprofile_clientsecret_step2 = `

	variable "authenticationoauthidpprofile_clientsecret_2" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_authenticationoauthidpprofile" "test" {
		name         = "tf_idpprofile_compat"
		clientid     = "cliId_compat"
		clientsecret = var.authenticationoauthidpprofile_clientsecret_2
		redirecturl  = "http://www.example.com/compat/"
	}
`

func TestAccAuthenticationoauthidpprofile_clientsecret_backward_compat(t *testing.T) {
	t.Setenv("TF_VAR_authenticationoauthidpprofile_clientsecret", "secret_value1")
	t.Setenv("TF_VAR_authenticationoauthidpprofile_clientsecret_2", "secret_value2")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationoauthidpprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationoauthidpprofile_clientsecret_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationoauthidpprofileExist("citrixadc_authenticationoauthidpprofile.test", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationoauthidpprofile.test", "name", "tf_idpprofile_compat"),
					resource.TestCheckResourceAttr("citrixadc_authenticationoauthidpprofile.test", "clientid", "cliId_compat"),
				),
			},
			{
				Config: testAccAuthenticationoauthidpprofile_clientsecret_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationoauthidpprofileExist("citrixadc_authenticationoauthidpprofile.test", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationoauthidpprofile.test", "name", "tf_idpprofile_compat"),
					resource.TestCheckResourceAttr("citrixadc_authenticationoauthidpprofile.test", "clientid", "cliId_compat"),
				),
			},
		},
	})
}

func TestAccAuthenticationoauthidpprofile_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckAuthenticationoauthidpprofileDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccAuthenticationoauthidpprofile_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationoauthidpprofileExist("citrixadc_authenticationoauthidpprofile.tf_idpprofile", nil),
				),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccAuthenticationoauthidpprofile_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationoauthidpprofileExist("citrixadc_authenticationoauthidpprofile.tf_idpprofile", nil),
				),
			},
		},
	})
}

// Test ephemeral path: using clientsecret_wo (WriteOnly attribute) with version tracker
const testAccAuthenticationoauthidpprofile_clientsecret_wo_step1 = `

	variable "authenticationoauthidpprofile_clientsecret_wo" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_authenticationoauthidpprofile" "test" {
		name                    = "tf_idpprofile_wo"
		clientid                = "cliId_wo"
		clientsecret_wo         = var.authenticationoauthidpprofile_clientsecret_wo
		clientsecret_wo_version = 1
		redirecturl             = "http://www.example.com/wo/"
	}
`

// Update ephemeral path: bump version to trigger update with new secret
const testAccAuthenticationoauthidpprofile_clientsecret_wo_step2 = `

	variable "authenticationoauthidpprofile_clientsecret_wo_2" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_authenticationoauthidpprofile" "test" {
		name                    = "tf_idpprofile_wo"
		clientid                = "cliId_wo"
		clientsecret_wo         = var.authenticationoauthidpprofile_clientsecret_wo_2
		clientsecret_wo_version = 2
		redirecturl             = "http://www.example.com/wo/"
	}
`

func TestAccAuthenticationoauthidpprofile_clientsecret_wo_ephemeral(t *testing.T) {
	t.Setenv("TF_VAR_authenticationoauthidpprofile_clientsecret_wo", "ephemeral_secret1")
	t.Setenv("TF_VAR_authenticationoauthidpprofile_clientsecret_wo_2", "ephemeral_secret2")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationoauthidpprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationoauthidpprofile_clientsecret_wo_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationoauthidpprofileExist("citrixadc_authenticationoauthidpprofile.test", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationoauthidpprofile.test", "name", "tf_idpprofile_wo"),
					resource.TestCheckResourceAttr("citrixadc_authenticationoauthidpprofile.test", "clientid", "cliId_wo"),
					resource.TestCheckResourceAttr("citrixadc_authenticationoauthidpprofile.test", "clientsecret_wo_version", "1"),
				),
			},
			{
				Config: testAccAuthenticationoauthidpprofile_clientsecret_wo_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationoauthidpprofileExist("citrixadc_authenticationoauthidpprofile.test", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationoauthidpprofile.test", "name", "tf_idpprofile_wo"),
					resource.TestCheckResourceAttr("citrixadc_authenticationoauthidpprofile.test", "clientid", "cliId_wo"),
					resource.TestCheckResourceAttr("citrixadc_authenticationoauthidpprofile.test", "clientsecret_wo_version", "2"),
				),
			},
		},
	})
}

// Step 1: unset-eligible attributes set to non-default values.
const testAccAuthenticationoauthidpprofile_unset_step1 = `

	resource "citrixadc_authenticationoauthidpprofile" "tf_unset" {
		name            = "tf_test_authoauthidp_unset"
		clientid        = "cliId"
		clientsecret    = "secret"
		redirecturl     = "http://www.example.com/1/"
		encrypttoken    = "ON"
		refreshinterval = 100
		sendpassword    = "ON"
		signaturealg    = "RS512"
		skewtime        = 10
	}
`

// Step 2: unset-eligible attributes removed from config -> provider must unset
// them, reverting each to its NITRO default.
const testAccAuthenticationoauthidpprofile_unset_step2 = `

	resource "citrixadc_authenticationoauthidpprofile" "tf_unset" {
		name         = "tf_test_authoauthidp_unset"
		clientid     = "cliId"
		clientsecret = "secret"
		redirecturl  = "http://www.example.com/1/"
	}
`

func TestAccAuthenticationoauthidpprofile_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationoauthidpprofileDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values apply and persist.
				Config: testAccAuthenticationoauthidpprofile_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationoauthidpprofileExist("citrixadc_authenticationoauthidpprofile.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationoauthidpprofile.tf_unset", "encrypttoken", "ON"),
					resource.TestCheckResourceAttr("citrixadc_authenticationoauthidpprofile.tf_unset", "refreshinterval", "100"),
					resource.TestCheckResourceAttr("citrixadc_authenticationoauthidpprofile.tf_unset", "sendpassword", "ON"),
					resource.TestCheckResourceAttr("citrixadc_authenticationoauthidpprofile.tf_unset", "signaturealg", "RS512"),
					resource.TestCheckResourceAttr("citrixadc_authenticationoauthidpprofile.tf_unset", "skewtime", "10"),
				),
			},
			{
				// Removing them must unset -> state reverts to NITRO defaults,
				// and the implicit post-apply plan must be empty.
				Config: testAccAuthenticationoauthidpprofile_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationoauthidpprofileExist("citrixadc_authenticationoauthidpprofile.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationoauthidpprofile.tf_unset", "encrypttoken", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_authenticationoauthidpprofile.tf_unset", "refreshinterval", "50"),
					resource.TestCheckResourceAttr("citrixadc_authenticationoauthidpprofile.tf_unset", "sendpassword", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_authenticationoauthidpprofile.tf_unset", "signaturealg", "RS256"),
					resource.TestCheckResourceAttr("citrixadc_authenticationoauthidpprofile.tf_unset", "skewtime", "5"),
					// Independent appliance-level confirmation the unset took effect:
					testAccCheckAuthenticationoauthidpprofileADCValue("tf_test_authoauthidp_unset", "encrypttoken", "OFF"),
					testAccCheckAuthenticationoauthidpprofileADCValue("tf_test_authoauthidp_unset", "refreshinterval", "50"),
					testAccCheckAuthenticationoauthidpprofileADCValue("tf_test_authoauthidp_unset", "sendpassword", "OFF"),
					testAccCheckAuthenticationoauthidpprofileADCValue("tf_test_authoauthidp_unset", "signaturealg", "RS256"),
					testAccCheckAuthenticationoauthidpprofileADCValue("tf_test_authoauthidp_unset", "skewtime", "5"),
				),
			},
		},
	})
}

// testAccCheckAuthenticationoauthidpprofileADCValue asserts an attribute's value
// directly on the appliance (not just in Terraform state), proving the unset
// actually reverted it to the NITRO default.
func testAccCheckAuthenticationoauthidpprofileADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource("authenticationoauthidpprofile", name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("authenticationoauthidpprofile %s not found on appliance", name)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("authenticationoauthidpprofile %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}
