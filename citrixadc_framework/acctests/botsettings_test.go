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

const testAccBotsettings_basic = `
	resource "citrixadc_botsettings" "default" {
		sessiontimeout= "900"
		proxyport = "8080"
		sessioncookiename = "citrix_bot_id"
		dfprequestlimit = "1"
		signatureautoupdate = "ON"
		trapurlautogenerate = "OFF"
		trapurlinterval = "3600"
		trapurllength = "32"
		proxyusername = "testuser"
	}
`
const testAccBotsettings_basic_update = `
	resource "citrixadc_botsettings" "default" {
		sessiontimeout= "950"
		proxyport = "80"
		sessioncookiename = "citrixbotid"
		dfprequestlimit = "3"
		signatureautoupdate = "ON"
		trapurlautogenerate = "ON"
		trapurlinterval = "3800"
		trapurllength = "33"
		proxyusername = "testuser1"
	}
`

func TestAccBotsettings_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// botsettings resource do not have DELETE operation
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				Config: testAccBotsettings_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBotsettingsExist("citrixadc_botsettings.default", nil),
					resource.TestCheckResourceAttr("citrixadc_botsettings.default", "sessiontimeout", "900"),
					resource.TestCheckResourceAttr("citrixadc_botsettings.default", "proxyport", "8080"),
					resource.TestCheckResourceAttr("citrixadc_botsettings.default", "sessioncookiename", "citrix_bot_id"),
					resource.TestCheckResourceAttr("citrixadc_botsettings.default", "dfprequestlimit", "1"),
					resource.TestCheckResourceAttr("citrixadc_botsettings.default", "signatureautoupdate", "ON"),
					resource.TestCheckResourceAttr("citrixadc_botsettings.default", "trapurlautogenerate", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_botsettings.default", "trapurlinterval", "3600"),
					resource.TestCheckResourceAttr("citrixadc_botsettings.default", "trapurllength", "32"),
					resource.TestCheckResourceAttr("citrixadc_botsettings.default", "proxyusername", "testuser"),
				),
			},
			{
				Config: testAccBotsettings_basic_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBotsettingsExist("citrixadc_botsettings.default", nil),
					resource.TestCheckResourceAttr("citrixadc_botsettings.default", "sessiontimeout", "950"),
					resource.TestCheckResourceAttr("citrixadc_botsettings.default", "proxyport", "80"),
					resource.TestCheckResourceAttr("citrixadc_botsettings.default", "sessioncookiename", "citrixbotid"),
					resource.TestCheckResourceAttr("citrixadc_botsettings.default", "dfprequestlimit", "3"),
					resource.TestCheckResourceAttr("citrixadc_botsettings.default", "signatureautoupdate", "ON"),
					resource.TestCheckResourceAttr("citrixadc_botsettings.default", "trapurlautogenerate", "ON"),
					resource.TestCheckResourceAttr("citrixadc_botsettings.default", "trapurlinterval", "3800"),
					resource.TestCheckResourceAttr("citrixadc_botsettings.default", "trapurllength", "33"),
					resource.TestCheckResourceAttr("citrixadc_botsettings.default", "proxyusername", "testuser1"),
				),
			},
		},
	})
}

func TestAccBotsettings_import(t *testing.T) {
	const resAddr = "citrixadc_botsettings.default"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// botsettings resource do not have DELETE operation
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{Config: testAccBotsettings_basic},
			{
				Config:            testAccBotsettings_basic,
				ResourceName:      resAddr,
				ImportState:       true,
				ImportStateVerify: true,
				// proxypassword_wo_version is a write-only version tracker stored in
				// state but never returned by NITRO, so it cannot round-trip on import.
				ImportStateVerifyIgnore: []string{"proxypassword_wo_version"},
			},
		},
	})
}

func testAccCheckBotsettingsExist(n string, id *string) resource.TestCheckFunc {
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
		data, err := client.FindResource("botsettings", "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("BOT Settings  %s not found", n)
		}

		return nil
	}
}

const testAccBotsettingsDataSource_basic = `
	resource "citrixadc_botsettings" "default" {
		sessiontimeout= "900"
		proxyport = "8080"
		sessioncookiename = "citrix_bot_id"
		dfprequestlimit = "1"
		signatureautoupdate = "ON"
		trapurlautogenerate = "OFF"
		trapurlinterval = "3600"
		trapurllength = "32"
		proxyusername = "testuser"
	}

	data "citrixadc_botsettings" "botsettings" {
		depends_on = [citrixadc_botsettings.default]
	}
`

func TestAccBotsettingsDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccBotsettingsDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_botsettings.botsettings", "sessiontimeout", "900"),
					resource.TestCheckResourceAttr("data.citrixadc_botsettings.botsettings", "proxyport", "8080"),
					resource.TestCheckResourceAttr("data.citrixadc_botsettings.botsettings", "sessioncookiename", "citrix_bot_id"),
					resource.TestCheckResourceAttr("data.citrixadc_botsettings.botsettings", "dfprequestlimit", "1"),
					resource.TestCheckResourceAttr("data.citrixadc_botsettings.botsettings", "signatureautoupdate", "ON"),
					resource.TestCheckResourceAttr("data.citrixadc_botsettings.botsettings", "trapurlautogenerate", "OFF"),
					resource.TestCheckResourceAttr("data.citrixadc_botsettings.botsettings", "trapurlinterval", "3600"),
					resource.TestCheckResourceAttr("data.citrixadc_botsettings.botsettings", "trapurllength", "32"),
					resource.TestCheckResourceAttr("data.citrixadc_botsettings.botsettings", "proxyusername", "testuser"),
				),
			},
		},
	})
}

// Test backward-compatible path: using proxypassword (Sensitive attribute)
const testAccBotsettings_proxypassword_step1 = `

	variable "botsettings_proxypassword" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_botsettings" "default" {
		proxyusername = "testuser"
		proxypassword = var.botsettings_proxypassword
		proxyport     = 8080
	}
`

// Update backward-compatible path: change proxypassword value
const testAccBotsettings_proxypassword_step2 = `

	variable "botsettings_proxypassword_2" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_botsettings" "default" {
		proxyusername = "testuser"
		proxypassword = var.botsettings_proxypassword_2
		proxyport     = 8080
	}
`

func TestAccBotsettings_proxypassword_backward_compat(t *testing.T) {
	t.Setenv("TF_VAR_botsettings_proxypassword", "oldpassword123")
	t.Setenv("TF_VAR_botsettings_proxypassword_2", "newpassword456")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccBotsettings_proxypassword_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBotsettingsExist("citrixadc_botsettings.default", nil),
					resource.TestCheckResourceAttr("citrixadc_botsettings.default", "proxyport", "8080"),
				),
			},
			{
				Config: testAccBotsettings_proxypassword_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBotsettingsExist("citrixadc_botsettings.default", nil),
					resource.TestCheckResourceAttr("citrixadc_botsettings.default", "proxyport", "8080"),
				),
			},
		},
	})
}

// Test ephemeral path: using proxypassword_wo (WriteOnly attribute) with version tracker
const testAccBotsettings_proxypassword_wo_step1 = `

	variable "botsettings_proxypassword_wo" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_botsettings" "default" {
		proxyusername            = "testuser"
		proxypassword_wo         = var.botsettings_proxypassword_wo
		proxypassword_wo_version = 1
		proxyport                = 8080
	}
`

// Update ephemeral path: bump version to trigger update with new password
const testAccBotsettings_proxypassword_wo_step2 = `

	variable "botsettings_proxypassword_wo_2" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_botsettings" "default" {
		proxyusername            = "testuser"
		proxypassword_wo         = var.botsettings_proxypassword_wo_2
		proxypassword_wo_version = 2
		proxyport                = 8080
	}
`

func TestAccBotsettings_proxypassword_wo_ephemeral(t *testing.T) {
	t.Setenv("TF_VAR_botsettings_proxypassword_wo", "ephemeral_pass1")
	t.Setenv("TF_VAR_botsettings_proxypassword_wo_2", "ephemeral_pass2")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccBotsettings_proxypassword_wo_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBotsettingsExist("citrixadc_botsettings.default", nil),
					resource.TestCheckResourceAttr("citrixadc_botsettings.default", "proxypassword_wo_version", "1"),
					resource.TestCheckResourceAttr("citrixadc_botsettings.default", "proxyport", "8080"),
				),
			},
			{
				Config: testAccBotsettings_proxypassword_wo_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBotsettingsExist("citrixadc_botsettings.default", nil),
					resource.TestCheckResourceAttr("citrixadc_botsettings.default", "proxypassword_wo_version", "2"),
					resource.TestCheckResourceAttr("citrixadc_botsettings.default", "proxyport", "8080"),
				),
			},
		},
	})
}

func TestAccBotsettings_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccBotsettings_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBotsettingsExist("citrixadc_botsettings.default", nil),
				),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccBotsettings_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBotsettingsExist("citrixadc_botsettings.default", nil),
				),
			},
		},
	})
}

// Step 1: all unset-eligible attributes set to non-default values.
const testAccBotsettings_unset_step1 = `
	resource "citrixadc_botsettings" "tf_unset" {
		defaultnonintrusiveprofile = "BOT_BYPASS"
		proxyport                  = "3128"
		signatureautoupdate        = "ON"
		signatureurl               = "https://example.com/BotSignatureMapping.json"
		trapurlautogenerate        = "ON"
		trapurlinterval            = "7200"
		trapurllength              = "64"
	}
`

// Step 2: eligible attributes removed from config -> provider must unset them,
// so the appliance reverts each to its NITRO default.
const testAccBotsettings_unset_step2 = `
	resource "citrixadc_botsettings" "tf_unset" {
		# unset-eligible attributes removed from config -> provider issues ?action=unset
	}
`

func TestAccBotsettings_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// botsettings resource has no DELETE operation
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				// Non-default values apply and persist.
				Config: testAccBotsettings_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBotsettingsExist("citrixadc_botsettings.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_botsettings.tf_unset", "defaultnonintrusiveprofile", "BOT_BYPASS"),
					resource.TestCheckResourceAttr("citrixadc_botsettings.tf_unset", "proxyport", "3128"),
					resource.TestCheckResourceAttr("citrixadc_botsettings.tf_unset", "signatureautoupdate", "ON"),
					resource.TestCheckResourceAttr("citrixadc_botsettings.tf_unset", "signatureurl", "https://example.com/BotSignatureMapping.json"),
					resource.TestCheckResourceAttr("citrixadc_botsettings.tf_unset", "trapurlautogenerate", "ON"),
					resource.TestCheckResourceAttr("citrixadc_botsettings.tf_unset", "trapurlinterval", "7200"),
					resource.TestCheckResourceAttr("citrixadc_botsettings.tf_unset", "trapurllength", "64"),
				),
			},
			{
				// Removing them must unset -> state reverts to NITRO defaults,
				// and the implicit post-apply plan must be empty.
				Config: testAccBotsettings_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBotsettingsExist("citrixadc_botsettings.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_botsettings.tf_unset", "defaultnonintrusiveprofile", "BOT_STATS"),
					resource.TestCheckResourceAttr("citrixadc_botsettings.tf_unset", "proxyport", "8080"),
					resource.TestCheckResourceAttr("citrixadc_botsettings.tf_unset", "signatureautoupdate", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_botsettings.tf_unset", "signatureurl", "https://nsbotsignatures.s3.amazonaws.com/BotSignatureMapping.json"),
					resource.TestCheckResourceAttr("citrixadc_botsettings.tf_unset", "trapurlautogenerate", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_botsettings.tf_unset", "trapurlinterval", "3600"),
					resource.TestCheckResourceAttr("citrixadc_botsettings.tf_unset", "trapurllength", "32"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckBotsettingsADCValue("defaultnonintrusiveprofile", "BOT_STATS"),
					testAccCheckBotsettingsADCValue("proxyport", "8080"),
					testAccCheckBotsettingsADCValue("signatureautoupdate", "OFF"),
					testAccCheckBotsettingsADCValue("signatureurl", "https://nsbotsignatures.s3.amazonaws.com/BotSignatureMapping.json"),
					testAccCheckBotsettingsADCValue("trapurlautogenerate", "OFF"),
					testAccCheckBotsettingsADCValue("trapurlinterval", "3600"),
					testAccCheckBotsettingsADCValue("trapurllength", "32"),
				),
			},
		},
	})
}

// testAccCheckBotsettingsADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted it.
// botsettings is a singleton, so it is fetched with an empty resource name.
func testAccCheckBotsettingsADCValue(attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource("botsettings", "")
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("botsettings not found on appliance")
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("botsettings: appliance attr %q = %q, want %q (unset did not revert it)", attr, got, want)
		}
		return nil
	}
}
