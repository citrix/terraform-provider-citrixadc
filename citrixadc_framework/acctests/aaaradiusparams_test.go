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

const testAccAaaradiusparams_basic = `


	resource "citrixadc_aaaradiusparams" "tf_aaaradiusparams" {
		radkey             = "sslvpn"
		radnasip           = "ENABLED"
		serverip           = "10.222.74.158"
		authtimeout        = 8
		messageauthenticator = "OFF"
	}
`
const testAccAaaradiusparams_update = `


	resource "citrixadc_aaaradiusparams" "tf_aaaradiusparams" {
		radkey             = "sslvpn2"
		radnasip           = "DISABLED"
		serverip           = "10.222.74.159"
		authtimeout        = 10
		messageauthenticator = "ON"
	}
`

func TestAccAaaradiusparams_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccAaaradiusparams_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaaradiusparamsExist("citrixadc_aaaradiusparams.tf_aaaradiusparams", nil),
					resource.TestCheckResourceAttr("citrixadc_aaaradiusparams.tf_aaaradiusparams", "radkey", "sslvpn"),
					resource.TestCheckResourceAttr("citrixadc_aaaradiusparams.tf_aaaradiusparams", "radnasip", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_aaaradiusparams.tf_aaaradiusparams", "serverip", "10.222.74.158"),
					resource.TestCheckResourceAttr("citrixadc_aaaradiusparams.tf_aaaradiusparams", "authtimeout", "8"),
					resource.TestCheckResourceAttr("citrixadc_aaaradiusparams.tf_aaaradiusparams", "messageauthenticator", "OFF"),
				),
			},
			{
				Config: testAccAaaradiusparams_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaaradiusparamsExist("citrixadc_aaaradiusparams.tf_aaaradiusparams", nil),
					resource.TestCheckResourceAttr("citrixadc_aaaradiusparams.tf_aaaradiusparams", "radkey", "sslvpn2"),
					resource.TestCheckResourceAttr("citrixadc_aaaradiusparams.tf_aaaradiusparams", "radnasip", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_aaaradiusparams.tf_aaaradiusparams", "serverip", "10.222.74.159"),
					resource.TestCheckResourceAttr("citrixadc_aaaradiusparams.tf_aaaradiusparams", "authtimeout", "10"),
					resource.TestCheckResourceAttr("citrixadc_aaaradiusparams.tf_aaaradiusparams", "messageauthenticator", "ON"),
				),
			},
		},
	})
}

func testAccCheckAaaradiusparamsExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No aaaradiusparams name is set")
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
		data, err := client.FindResource(service.Aaaradiusparams.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("aaaradiusparams %s not found", n)
		}

		return nil
	}
}

const testAccAaaradiusparamsDataSource_basic = `


	resource "citrixadc_aaaradiusparams" "tf_aaaradiusparams" {
		radkey             = "sslvpn"
		radnasip           = "ENABLED"
		serverip           = "10.222.74.158"
		authtimeout        = 8
		messageauthenticator = "OFF"
	}

	data "citrixadc_aaaradiusparams" "tf_aaaradiusparams" {
		depends_on = [citrixadc_aaaradiusparams.tf_aaaradiusparams]
	}
`

func TestAccAaaradiusparamsDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccAaaradiusparamsDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					// radkey is not checked as it's returned encrypted/hashed by the API
					resource.TestCheckResourceAttr("data.citrixadc_aaaradiusparams.tf_aaaradiusparams", "radnasip", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_aaaradiusparams.tf_aaaradiusparams", "serverip", "10.222.74.158"),
					resource.TestCheckResourceAttr("data.citrixadc_aaaradiusparams.tf_aaaradiusparams", "authtimeout", "8"),
					resource.TestCheckResourceAttr("data.citrixadc_aaaradiusparams.tf_aaaradiusparams", "messageauthenticator", "OFF"),
				),
			},
		},
	})
}

const testAccAaaradiusparams_radkey_step1 = `
	variable "aaaradiusparams_radkey" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_aaaradiusparams" "tf_aaaradiusparams" {
		radkey               = var.aaaradiusparams_radkey
		radnasip             = "ENABLED"
		serverip             = "10.222.74.158"
		authtimeout          = 8
		messageauthenticator = "OFF"
	}
`

const testAccAaaradiusparams_radkey_step2 = `
	variable "aaaradiusparams_radkey_2" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_aaaradiusparams" "tf_aaaradiusparams" {
		radkey               = var.aaaradiusparams_radkey_2
		radnasip             = "DISABLED"
		serverip             = "10.222.74.159"
		authtimeout          = 10
		messageauthenticator = "ON"
	}
`

func TestAccAaaradiusparams_radkey_backward_compat(t *testing.T) {
	t.Setenv("TF_VAR_aaaradiusparams_radkey", "sslvpn")
	t.Setenv("TF_VAR_aaaradiusparams_radkey_2", "sslvpn2")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccAaaradiusparams_radkey_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaaradiusparamsExist("citrixadc_aaaradiusparams.tf_aaaradiusparams", nil),
					resource.TestCheckResourceAttr("citrixadc_aaaradiusparams.tf_aaaradiusparams", "radnasip", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_aaaradiusparams.tf_aaaradiusparams", "serverip", "10.222.74.158"),
					resource.TestCheckResourceAttr("citrixadc_aaaradiusparams.tf_aaaradiusparams", "authtimeout", "8"),
					resource.TestCheckResourceAttr("citrixadc_aaaradiusparams.tf_aaaradiusparams", "messageauthenticator", "OFF"),
				),
			},
			{
				Config: testAccAaaradiusparams_radkey_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaaradiusparamsExist("citrixadc_aaaradiusparams.tf_aaaradiusparams", nil),
					resource.TestCheckResourceAttr("citrixadc_aaaradiusparams.tf_aaaradiusparams", "radnasip", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_aaaradiusparams.tf_aaaradiusparams", "serverip", "10.222.74.159"),
					resource.TestCheckResourceAttr("citrixadc_aaaradiusparams.tf_aaaradiusparams", "authtimeout", "10"),
					resource.TestCheckResourceAttr("citrixadc_aaaradiusparams.tf_aaaradiusparams", "messageauthenticator", "ON"),
				),
			},
		},
	})
}

const testAccAaaradiusparams_radkey_wo_step1 = `
	variable "aaaradiusparams_radkey_wo" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_aaaradiusparams" "tf_aaaradiusparams" {
		radkey_wo            = var.aaaradiusparams_radkey_wo
		radkey_wo_version    = 1
		radnasip             = "ENABLED"
		serverip             = "10.222.74.158"
		authtimeout          = 8
		messageauthenticator = "OFF"
	}
`

const testAccAaaradiusparams_radkey_wo_step2 = `
	variable "aaaradiusparams_radkey_wo_2" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_aaaradiusparams" "tf_aaaradiusparams" {
		radkey_wo            = var.aaaradiusparams_radkey_wo_2
		radkey_wo_version    = 2
		radnasip             = "DISABLED"
		serverip             = "10.222.74.159"
		authtimeout          = 10
		messageauthenticator = "ON"
	}
`

func TestAccAaaradiusparams_radkey_wo_ephemeral(t *testing.T) {
	t.Setenv("TF_VAR_aaaradiusparams_radkey_wo", "sslvpn")
	t.Setenv("TF_VAR_aaaradiusparams_radkey_wo_2", "sslvpn2")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccAaaradiusparams_radkey_wo_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaaradiusparamsExist("citrixadc_aaaradiusparams.tf_aaaradiusparams", nil),
					resource.TestCheckResourceAttr("citrixadc_aaaradiusparams.tf_aaaradiusparams", "radkey_wo_version", "1"),
					resource.TestCheckResourceAttr("citrixadc_aaaradiusparams.tf_aaaradiusparams", "radnasip", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_aaaradiusparams.tf_aaaradiusparams", "serverip", "10.222.74.158"),
					resource.TestCheckResourceAttr("citrixadc_aaaradiusparams.tf_aaaradiusparams", "authtimeout", "8"),
					resource.TestCheckResourceAttr("citrixadc_aaaradiusparams.tf_aaaradiusparams", "messageauthenticator", "OFF"),
				),
			},
			{
				Config: testAccAaaradiusparams_radkey_wo_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaaradiusparamsExist("citrixadc_aaaradiusparams.tf_aaaradiusparams", nil),
					resource.TestCheckResourceAttr("citrixadc_aaaradiusparams.tf_aaaradiusparams", "radkey_wo_version", "2"),
					resource.TestCheckResourceAttr("citrixadc_aaaradiusparams.tf_aaaradiusparams", "radnasip", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_aaaradiusparams.tf_aaaradiusparams", "serverip", "10.222.74.159"),
					resource.TestCheckResourceAttr("citrixadc_aaaradiusparams.tf_aaaradiusparams", "authtimeout", "10"),
					resource.TestCheckResourceAttr("citrixadc_aaaradiusparams.tf_aaaradiusparams", "messageauthenticator", "ON"),
				),
			},
		},
	})
}

func TestAccAaaradiusparams_import(t *testing.T) {
	const resAddr = "citrixadc_aaaradiusparams.tf_aaaradiusparams"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{Config: testAccAaaradiusparams_basic},
			{
				Config:            testAccAaaradiusparams_basic,
				ResourceName:      resAddr,
				ImportState:       true,
				ImportStateVerify: true,
				// radkey is a secret the API returns encrypted/hashed (cannot
				// round-trip); radkey_wo_version is a state-only version tracker
				// that the NITRO GET never echoes back.
				ImportStateVerifyIgnore: []string{"radkey", "radkey_wo_version"},
			},
		},
	})
}

func TestAccAaaradiusparams_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccAaaradiusparams_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaaradiusparamsExist("citrixadc_aaaradiusparams.tf_aaaradiusparams", nil),
				),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccAaaradiusparams_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaaradiusparamsExist("citrixadc_aaaradiusparams.tf_aaaradiusparams", nil),
				),
			},
		},
	})
}

// Step 1: all unset-eligible attributes set to non-default values.
const testAccAaaradiusparams_unset_step1 = `
	resource "citrixadc_aaaradiusparams" "tf_unset" {
		radkey                 = "sslvpn"
		authentication         = "OFF"
		authservretry          = 4
		authtimeout            = 8
		callingstationid       = "ENABLED"
		messageauthenticator   = "OFF"
		passencoding           = "pap"
		serverport             = 1813
		tunnelendpointclientip = "ENABLED"
	}
`

// Step 2: the unset-eligible attributes are removed from config so the provider
// must issue ?action=unset, reverting each to its NITRO default.
const testAccAaaradiusparams_unset_step2 = `
	resource "citrixadc_aaaradiusparams" "tf_unset" {
		radkey = "sslvpn"
	}
`

func TestAccAaaradiusparams_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil, // singleton resource - never truly deleted
		Steps: []resource.TestStep{
			{
				// Non-default values apply and persist.
				Config: testAccAaaradiusparams_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaaradiusparamsExist("citrixadc_aaaradiusparams.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_aaaradiusparams.tf_unset", "authentication", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_aaaradiusparams.tf_unset", "authservretry", "4"),
					resource.TestCheckResourceAttr("citrixadc_aaaradiusparams.tf_unset", "authtimeout", "8"),
					resource.TestCheckResourceAttr("citrixadc_aaaradiusparams.tf_unset", "callingstationid", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_aaaradiusparams.tf_unset", "messageauthenticator", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_aaaradiusparams.tf_unset", "passencoding", "pap"),
					resource.TestCheckResourceAttr("citrixadc_aaaradiusparams.tf_unset", "serverport", "1813"),
					resource.TestCheckResourceAttr("citrixadc_aaaradiusparams.tf_unset", "tunnelendpointclientip", "ENABLED"),
				),
			},
			{
				// Removing them must unset -> state reverts to NITRO defaults,
				// and the implicit post-apply plan must be empty.
				Config: testAccAaaradiusparams_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaaradiusparamsExist("citrixadc_aaaradiusparams.tf_unset", nil),
					// State reverted to defaults.
					resource.TestCheckResourceAttr("citrixadc_aaaradiusparams.tf_unset", "authentication", "ON"),
					resource.TestCheckResourceAttr("citrixadc_aaaradiusparams.tf_unset", "authservretry", "3"),
					resource.TestCheckResourceAttr("citrixadc_aaaradiusparams.tf_unset", "authtimeout", "3"),
					resource.TestCheckResourceAttr("citrixadc_aaaradiusparams.tf_unset", "callingstationid", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_aaaradiusparams.tf_unset", "messageauthenticator", "ON"),
					resource.TestCheckResourceAttr("citrixadc_aaaradiusparams.tf_unset", "passencoding", "mschapv2"),
					resource.TestCheckResourceAttr("citrixadc_aaaradiusparams.tf_unset", "serverport", "1812"),
					resource.TestCheckResourceAttr("citrixadc_aaaradiusparams.tf_unset", "tunnelendpointclientip", "DISABLED"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckAaaradiusparamsADCValue("authentication", "ON"),
					testAccCheckAaaradiusparamsADCValue("authservretry", "3"),
					testAccCheckAaaradiusparamsADCValue("authtimeout", "3"),
					testAccCheckAaaradiusparamsADCValue("callingstationid", "DISABLED"),
					testAccCheckAaaradiusparamsADCValue("messageauthenticator", "ON"),
					testAccCheckAaaradiusparamsADCValue("passencoding", "mschapv2"),
					testAccCheckAaaradiusparamsADCValue("serverport", "1812"),
					testAccCheckAaaradiusparamsADCValue("tunnelendpointclientip", "DISABLED"),
				),
			},
		},
	})
}

// testAccCheckAaaradiusparamsADCValue asserts an attribute's value directly on
// the appliance (not just in Terraform state), proving the unset actually
// reverted it. aaaradiusparams is a singleton, so the resource name is empty.
func testAccCheckAaaradiusparamsADCValue(attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Aaaradiusparams.Type(), "")
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("aaaradiusparams not found on appliance")
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("aaaradiusparams: appliance attr %q = %q, want %q (unset did not revert it)", attr, got, want)
		}
		return nil
	}
}
