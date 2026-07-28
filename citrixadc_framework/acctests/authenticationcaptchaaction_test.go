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

const testAccAuthenticationcaptchaaction_add = `
	resource "citrixadc_authenticationcaptchaaction" "tf_captchaaction" {
		name                       = "tf_captchaaction"
		secretkey                  = "secret"
		sitekey                    = "key"
		serverurl                  = "http://www.example.com/"
		defaultauthenticationgroup = "old_group"
		scorethreshold			 = 3
	}
`
const testAccAuthenticationcaptchaaction_update = `
	resource "citrixadc_authenticationcaptchaaction" "tf_captchaaction" {
		name                       = "tf_captchaaction"
		secretkey                  = "new_secret"
		sitekey                    = "key"
		serverurl                  = "http://www.example.com/"
		defaultauthenticationgroup = "new_group"
		scorethreshold			 = 6
	}
`

func TestAccAuthenticationcaptchaaction_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationcaptchaactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationcaptchaaction_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationcaptchaactionExist("citrixadc_authenticationcaptchaaction.tf_captchaaction", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationcaptchaaction.tf_captchaaction", "name", "tf_captchaaction"),
					resource.TestCheckResourceAttr("citrixadc_authenticationcaptchaaction.tf_captchaaction", "secretkey", "secret"),
					resource.TestCheckResourceAttr("citrixadc_authenticationcaptchaaction.tf_captchaaction", "defaultauthenticationgroup", "old_group"),
					resource.TestCheckResourceAttr("citrixadc_authenticationcaptchaaction.tf_captchaaction", "scorethreshold", "3"),
				),
			},
			{
				Config: testAccAuthenticationcaptchaaction_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationcaptchaactionExist("citrixadc_authenticationcaptchaaction.tf_captchaaction", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationcaptchaaction.tf_captchaaction", "name", "tf_captchaaction"),
					resource.TestCheckResourceAttr("citrixadc_authenticationcaptchaaction.tf_captchaaction", "secretkey", "new_secret"),
					resource.TestCheckResourceAttr("citrixadc_authenticationcaptchaaction.tf_captchaaction", "defaultauthenticationgroup", "new_group"),
					resource.TestCheckResourceAttr("citrixadc_authenticationcaptchaaction.tf_captchaaction", "scorethreshold", "6"),
				),
			},
		},
	})
}

func TestAccAuthenticationcaptchaaction_import(t *testing.T) {
	const resAddr = "citrixadc_authenticationcaptchaaction.tf_captchaaction"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationcaptchaactionDestroy,
		Steps: []resource.TestStep{
			{Config: testAccAuthenticationcaptchaaction_add},
			{
				Config:            testAccAuthenticationcaptchaaction_add,
				ResourceName:      resAddr,
				ImportState:       true,
				ImportStateVerify: true,
				// secretkey/sitekey are sensitive and never echoed back by NITRO
				// (retained from config on Read), so they are null after import.
				// The *_wo_version trackers are Computed and not populated by the
				// GET response, so none of these can round-trip through import.
				ImportStateVerifyIgnore: []string{"secretkey", "secretkey_wo_version", "sitekey", "sitekey_wo_version"},
			},
		},
	})
}

func testAccCheckAuthenticationcaptchaactionExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No authenticationcaptchaaction name is set")
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
		data, err := client.FindResource("authenticationcaptchaaction", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("authenticationcaptchaaction %s not found", n)
		}

		return nil
	}
}

func testAccCheckAuthenticationcaptchaactionDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_authenticationcaptchaaction" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("authenticationcaptchaaction", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("authenticationcaptchaaction %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccAuthenticationcaptchaactionDataSource_basic = `
	resource "citrixadc_authenticationcaptchaaction" "tf_captchaaction" {
		name                       = "tf_captchaaction_ds"
		secretkey                  = "secret"
		sitekey                    = "key"
		serverurl                  = "http://www.example.com/"
		defaultauthenticationgroup = "test_group"
		scorethreshold             = 7
	}

	data "citrixadc_authenticationcaptchaaction" "tf_captchaaction" {
		name       = citrixadc_authenticationcaptchaaction.tf_captchaaction.name
		depends_on = [citrixadc_authenticationcaptchaaction.tf_captchaaction]
	}
`

func TestAccAuthenticationcaptchaactionDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationcaptchaactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationcaptchaactionDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_authenticationcaptchaaction.tf_captchaaction", "name", "tf_captchaaction_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationcaptchaaction.tf_captchaaction", "serverurl", "http://www.example.com/"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationcaptchaaction.tf_captchaaction", "defaultauthenticationgroup", "test_group"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationcaptchaaction.tf_captchaaction", "scorethreshold", "7"),
				),
			},
		},
	})
}

// ---- secretkey backward-compat tests ----

const testAccAuthenticationcaptchaaction_secretkey_step1 = `
	variable "authenticationcaptchaaction_secretkey" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_authenticationcaptchaaction" "tf_captchaaction" {
		name                       = "tf_captchaaction"
		secretkey                  = var.authenticationcaptchaaction_secretkey
		sitekey                    = "key"
		serverurl                  = "http://www.example.com/"
		defaultauthenticationgroup = "old_group"
		scorethreshold             = 3
	}
`

const testAccAuthenticationcaptchaaction_secretkey_step2 = `
	variable "authenticationcaptchaaction_secretkey_2" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_authenticationcaptchaaction" "tf_captchaaction" {
		name                       = "tf_captchaaction"
		secretkey                  = var.authenticationcaptchaaction_secretkey_2
		sitekey                    = "key"
		serverurl                  = "http://www.example.com/"
		defaultauthenticationgroup = "new_group"
		scorethreshold             = 6
	}
`

func TestAccAuthenticationcaptchaaction_secretkey_backward_compat(t *testing.T) {
	t.Setenv("TF_VAR_authenticationcaptchaaction_secretkey", "secret1")
	t.Setenv("TF_VAR_authenticationcaptchaaction_secretkey_2", "secret2")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationcaptchaactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationcaptchaaction_secretkey_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationcaptchaactionExist("citrixadc_authenticationcaptchaaction.tf_captchaaction", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationcaptchaaction.tf_captchaaction", "name", "tf_captchaaction"),
					resource.TestCheckResourceAttr("citrixadc_authenticationcaptchaaction.tf_captchaaction", "defaultauthenticationgroup", "old_group"),
					resource.TestCheckResourceAttr("citrixadc_authenticationcaptchaaction.tf_captchaaction", "scorethreshold", "3"),
				),
			},
			{
				Config: testAccAuthenticationcaptchaaction_secretkey_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationcaptchaactionExist("citrixadc_authenticationcaptchaaction.tf_captchaaction", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationcaptchaaction.tf_captchaaction", "name", "tf_captchaaction"),
					resource.TestCheckResourceAttr("citrixadc_authenticationcaptchaaction.tf_captchaaction", "defaultauthenticationgroup", "new_group"),
					resource.TestCheckResourceAttr("citrixadc_authenticationcaptchaaction.tf_captchaaction", "scorethreshold", "6"),
				),
			},
		},
	})
}

// ---- secretkey write-only (ephemeral) tests ----

const testAccAuthenticationcaptchaaction_secretkey_wo_step1 = `
	variable "authenticationcaptchaaction_secretkey_wo" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_authenticationcaptchaaction" "tf_captchaaction" {
		name                       = "tf_captchaaction"
		secretkey_wo               = var.authenticationcaptchaaction_secretkey_wo
		secretkey_wo_version       = 1
		sitekey                    = "key"
		serverurl                  = "http://www.example.com/"
		defaultauthenticationgroup = "old_group"
		scorethreshold             = 3
	}
`

const testAccAuthenticationcaptchaaction_secretkey_wo_step2 = `
	variable "authenticationcaptchaaction_secretkey_wo_2" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_authenticationcaptchaaction" "tf_captchaaction" {
		name                       = "tf_captchaaction"
		secretkey_wo               = var.authenticationcaptchaaction_secretkey_wo_2
		secretkey_wo_version       = 2
		sitekey                    = "key"
		serverurl                  = "http://www.example.com/"
		defaultauthenticationgroup = "new_group"
		scorethreshold             = 6
	}
`

func TestAccAuthenticationcaptchaaction_secretkey_wo_ephemeral(t *testing.T) {
	t.Setenv("TF_VAR_authenticationcaptchaaction_secretkey_wo", "ephemeral_secret1")
	t.Setenv("TF_VAR_authenticationcaptchaaction_secretkey_wo_2", "ephemeral_secret2")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationcaptchaactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationcaptchaaction_secretkey_wo_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationcaptchaactionExist("citrixadc_authenticationcaptchaaction.tf_captchaaction", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationcaptchaaction.tf_captchaaction", "secretkey_wo_version", "1"),
				),
			},
			{
				Config: testAccAuthenticationcaptchaaction_secretkey_wo_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationcaptchaactionExist("citrixadc_authenticationcaptchaaction.tf_captchaaction", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationcaptchaaction.tf_captchaaction", "secretkey_wo_version", "2"),
				),
			},
		},
	})
}

// ---- sitekey backward-compat tests ----

const testAccAuthenticationcaptchaaction_sitekey_step1 = `
	variable "authenticationcaptchaaction_sitekey" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_authenticationcaptchaaction" "tf_captchaaction" {
		name                       = "tf_captchaaction"
		secretkey                  = "secret"
		sitekey                    = var.authenticationcaptchaaction_sitekey
		serverurl                  = "http://www.example.com/"
		defaultauthenticationgroup = "old_group"
		scorethreshold             = 3
	}
`

const testAccAuthenticationcaptchaaction_sitekey_step2 = `
	variable "authenticationcaptchaaction_sitekey_2" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_authenticationcaptchaaction" "tf_captchaaction" {
		name                       = "tf_captchaaction"
		secretkey                  = "secret"
		sitekey                    = var.authenticationcaptchaaction_sitekey_2
		serverurl                  = "http://www.example.com/"
		defaultauthenticationgroup = "new_group"
		scorethreshold             = 6
	}
`

func TestAccAuthenticationcaptchaaction_sitekey_backward_compat(t *testing.T) {
	t.Setenv("TF_VAR_authenticationcaptchaaction_sitekey", "sitekey1")
	t.Setenv("TF_VAR_authenticationcaptchaaction_sitekey_2", "sitekey2")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationcaptchaactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationcaptchaaction_sitekey_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationcaptchaactionExist("citrixadc_authenticationcaptchaaction.tf_captchaaction", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationcaptchaaction.tf_captchaaction", "name", "tf_captchaaction"),
					resource.TestCheckResourceAttr("citrixadc_authenticationcaptchaaction.tf_captchaaction", "defaultauthenticationgroup", "old_group"),
					resource.TestCheckResourceAttr("citrixadc_authenticationcaptchaaction.tf_captchaaction", "scorethreshold", "3"),
				),
			},
			{
				Config: testAccAuthenticationcaptchaaction_sitekey_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationcaptchaactionExist("citrixadc_authenticationcaptchaaction.tf_captchaaction", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationcaptchaaction.tf_captchaaction", "name", "tf_captchaaction"),
					resource.TestCheckResourceAttr("citrixadc_authenticationcaptchaaction.tf_captchaaction", "defaultauthenticationgroup", "new_group"),
					resource.TestCheckResourceAttr("citrixadc_authenticationcaptchaaction.tf_captchaaction", "scorethreshold", "6"),
				),
			},
		},
	})
}

// ---- sitekey write-only (ephemeral) tests ----

const testAccAuthenticationcaptchaaction_sitekey_wo_step1 = `
	variable "authenticationcaptchaaction_sitekey_wo" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_authenticationcaptchaaction" "tf_captchaaction" {
		name                       = "tf_captchaaction"
		secretkey                  = "secret"
		sitekey_wo                 = var.authenticationcaptchaaction_sitekey_wo
		sitekey_wo_version         = 1
		serverurl                  = "http://www.example.com/"
		defaultauthenticationgroup = "old_group"
		scorethreshold             = 3
	}
`

const testAccAuthenticationcaptchaaction_sitekey_wo_step2 = `
	variable "authenticationcaptchaaction_sitekey_wo_2" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_authenticationcaptchaaction" "tf_captchaaction" {
		name                       = "tf_captchaaction"
		secretkey                  = "secret"
		sitekey_wo                 = var.authenticationcaptchaaction_sitekey_wo_2
		sitekey_wo_version         = 2
		serverurl                  = "http://www.example.com/"
		defaultauthenticationgroup = "new_group"
		scorethreshold             = 6
	}
`

func TestAccAuthenticationcaptchaaction_sitekey_wo_ephemeral(t *testing.T) {
	t.Setenv("TF_VAR_authenticationcaptchaaction_sitekey_wo", "ephemeral_sitekey1")
	t.Setenv("TF_VAR_authenticationcaptchaaction_sitekey_wo_2", "ephemeral_sitekey2")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationcaptchaactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationcaptchaaction_sitekey_wo_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationcaptchaactionExist("citrixadc_authenticationcaptchaaction.tf_captchaaction", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationcaptchaaction.tf_captchaaction", "sitekey_wo_version", "1"),
				),
			},
			{
				Config: testAccAuthenticationcaptchaaction_sitekey_wo_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationcaptchaactionExist("citrixadc_authenticationcaptchaaction.tf_captchaaction", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationcaptchaaction.tf_captchaaction", "sitekey_wo_version", "2"),
				),
			},
		},
	})
}

// ---- unset support test ----

// Step 1: scorethreshold set to a non-default value (default is 5).
const testAccAuthenticationcaptchaaction_unset_step1 = `
	resource "citrixadc_authenticationcaptchaaction" "tf_unset" {
		name           = "tf_test_authenticationcaptchaaction_unset"
		secretkey      = "secret"
		sitekey        = "key"
		serverurl      = "http://www.example.com/"
		scorethreshold = 3
	}
`

// Step 2: scorethreshold removed from config -> provider must issue ?action=unset
// so the appliance reverts it to the NITRO default (5).
const testAccAuthenticationcaptchaaction_unset_step2 = `
	resource "citrixadc_authenticationcaptchaaction" "tf_unset" {
		name           = "tf_test_authenticationcaptchaaction_unset"
		secretkey      = "secret"
		sitekey        = "key"
		serverurl      = "http://www.example.com/"
	}
`

func TestAccAuthenticationcaptchaaction_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationcaptchaactionDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default value applies and persists.
				Config: testAccAuthenticationcaptchaaction_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationcaptchaactionExist("citrixadc_authenticationcaptchaaction.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationcaptchaaction.tf_unset", "scorethreshold", "3"),
				),
			},
			{
				// Removing it must unset -> state reverts to the NITRO default,
				// and the implicit post-apply plan must be empty.
				Config: testAccAuthenticationcaptchaaction_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationcaptchaactionExist("citrixadc_authenticationcaptchaaction.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationcaptchaaction.tf_unset", "scorethreshold", "5"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckAuthenticationcaptchaactionADCValue("tf_test_authenticationcaptchaaction_unset", "scorethreshold", "5"),
				),
			},
		},
	})
}

// testAccCheckAuthenticationcaptchaactionADCValue asserts an attribute's value
// directly on the appliance (not just in Terraform state), proving the unset
// actually reverted it.
func testAccCheckAuthenticationcaptchaactionADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Authenticationcaptchaaction.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("authenticationcaptchaaction %s not found on appliance", name)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("authenticationcaptchaaction %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func TestAccAuthenticationcaptchaaction_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckAuthenticationcaptchaactionDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccAuthenticationcaptchaaction_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationcaptchaactionExist("citrixadc_authenticationcaptchaaction.tf_captchaaction", nil),
				),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccAuthenticationcaptchaaction_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationcaptchaactionExist("citrixadc_authenticationcaptchaaction.tf_captchaaction", nil),
				),
			},
		},
	})
}
