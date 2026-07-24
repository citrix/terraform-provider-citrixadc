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

const testAccSslprofile_add = `
	resource "citrixadc_sslprofile" "foo" {
		name = "tfAcc_sslprofile"
		ecccurvebindings = []
	}
`
const testAccSslprofile_update = `
	resource "citrixadc_sslprofile" "foo" {
		name = "tfAcc_sslprofile"
		hsts = "ENABLED"
		snienable = "ENABLED"
		ecccurvebindings = []
		sslclientlogs = "ENABLED"
		encryptedclienthello = "ENABLED"
		defaultsni = 60
		allowunknownsni = "ENABLED"
		allowextendedmastersecret = "YES"
	}
`

func TestAccSslprofile_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslprofile_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslprofileExist("citrixadc_sslprofile.foo", nil),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.foo", "name", "tfAcc_sslprofile"),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.foo", "sslclientlogs", "DISABLED"),
				),
			},
			{
				Config: testAccSslprofile_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslprofileExist("citrixadc_sslprofile.foo", nil),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.foo", "name", "tfAcc_sslprofile"),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.foo", "hsts", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.foo", "sslclientlogs", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.foo", "encryptedclienthello", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.foo", "defaultsni", "60"),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.foo", "allowunknownsni", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.foo", "allowextendedmastersecret", "YES"),
				),
			},
		},
	})
}

const testAccSslprofile_ecccurvebinding_bind = `
	resource "citrixadc_sslprofile" "foo" {
		name = "tfAcc_sslprofile"
		ecccurvebindings = ["P_256"]
	}
`
const testAccSslprofile_ecccurvebinding_unbind = `
	resource "citrixadc_sslprofile" "foo" {
		name = "tfAcc_sslprofile"
		ecccurvebindings = []
	}
`

func TestAccSslprofile_ecccurve_binding(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslprofile_ecccurvebinding_bind,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslprofileExist("citrixadc_sslprofile.foo", nil),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.foo", "name", "tfAcc_sslprofile"),
				),
			},
			{
				Config: testAccSslprofile_ecccurvebinding_unbind,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslprofileExist("citrixadc_sslprofile.foo", nil),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.foo", "name", "tfAcc_sslprofile"),
				),
			},
		},
	})
}

const testAccSslprofile_cipherbinding_bind = `
	resource "citrixadc_sslprofile" "foo" {
		name = "tfAcc_sslprofile"
		ecccurvebindings = []
		cipherbindings {
			ciphername     = "HIGH"
			cipherpriority = 10
	}
	}
`
const testAccSslprofile_cipherbinding_unbind = `
	resource "citrixadc_sslprofile" "foo" {
		name = "tfAcc_sslprofile"
		ecccurvebindings = []
	}
`

func TestAccSslprofile_cipher_binding(t *testing.T) {
	if adcTestbed != "STANDALONE_DEFAULT_SSL_PROFILE" {
		t.Skipf("ADC testbed is %s. Expected STANDALONE_DEFAULT_SSL_PROFILE.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslprofile_cipherbinding_bind,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslprofileExist("citrixadc_sslprofile.foo", nil),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.foo", "name", "tfAcc_sslprofile"),
				),
			},
			{
				Config: testAccSslprofile_cipherbinding_unbind,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslprofileExist("citrixadc_sslprofile.foo", nil),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.foo", "name", "tfAcc_sslprofile"),
				),
			},
		},
	})
}

func TestAccSslprofile_import(t *testing.T) {
	const resAddr = "citrixadc_sslprofile.foo"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslprofileDestroy,
		Steps: []resource.TestStep{
			{Config: testAccSslprofile_add},
			{
				Config:            testAccSslprofile_add,
				ResourceName:      resAddr,
				ImportState:       true,
				ImportStateVerify: true,
				// sessionticketkeydata_wo_version is a write-only version tracker that
				// NITRO does not return, so it cannot round-trip through import.
				ImportStateVerifyIgnore: []string{"sessionticketkeydata_wo_version"},
			},
		},
	})
}

func testAccCheckSslprofileExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SSL Profile name is set")
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
		data, err := client.FindResource(service.Sslprofile.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("SSL Profile %s not found", n)
		}

		return nil
	}
}

func testAccCheckSslprofileDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_sslprofile" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Sslprofile.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("SSL Profile %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccSslprofileDataSource_basic = `
	resource "citrixadc_sslprofile" "tf_sslprofile" {
		name = "tf_sslprofile_datasource"
		hsts = "ENABLED"
		snienable = "ENABLED"
		ecccurvebindings = []
		sslclientlogs = "ENABLED"
		encryptedclienthello = "ENABLED"
		defaultsni = "60"
		allowunknownsni = "ENABLED"
		allowextendedmastersecret = "YES"
	}

	data "citrixadc_sslprofile" "tf_sslprofile_datasource" {
		name = citrixadc_sslprofile.tf_sslprofile.name
	}
`

func TestAccSslprofileDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSslprofileDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_sslprofile.tf_sslprofile_datasource", "name", "tf_sslprofile_datasource"),
					resource.TestCheckResourceAttr("data.citrixadc_sslprofile.tf_sslprofile_datasource", "hsts", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_sslprofile.tf_sslprofile_datasource", "snienable", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_sslprofile.tf_sslprofile_datasource", "sslclientlogs", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_sslprofile.tf_sslprofile_datasource", "encryptedclienthello", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_sslprofile.tf_sslprofile_datasource", "defaultsni", "60"),
					resource.TestCheckResourceAttr("data.citrixadc_sslprofile.tf_sslprofile_datasource", "allowunknownsni", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_sslprofile.tf_sslprofile_datasource", "allowextendedmastersecret", "YES"),
				),
			},
		},
	})
}

// Test backward-compatible path: using sessionticketkeydata (Sensitive attribute)
const testAccSslprofile_sessionticketkeydata_step1 = `
	variable "sslprofile_sessionticketkeydata" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_sslprofile" "tf_sslprofile_ephem" {
		name                   = "tf_sslprofile_ephem"
		sessionticket          = "ENABLED"
		sessionticketkeydata   = var.sslprofile_sessionticketkeydata
	}
`

const testAccSslprofile_sessionticketkeydata_step2 = `
	variable "sslprofile_sessionticketkeydata_2" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_sslprofile" "tf_sslprofile_ephem" {
		name                   = "tf_sslprofile_ephem"
		sessionticket          = "ENABLED"
		sessionticketkeydata   = var.sslprofile_sessionticketkeydata_2
	}
`

func TestAccSslprofile_sessionticketkeydata_backward_compat(t *testing.T) {
	t.Setenv("TF_VAR_sslprofile_sessionticketkeydata", "0102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f20")
	t.Setenv("TF_VAR_sslprofile_sessionticketkeydata_2", "2122232425262728292a2b2c2d2e2f303132333435363738393a3b3c3d3e3f40")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslprofile_sessionticketkeydata_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslprofileExist("citrixadc_sslprofile.tf_sslprofile_ephem", nil),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.tf_sslprofile_ephem", "name", "tf_sslprofile_ephem"),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.tf_sslprofile_ephem", "sessionticket", "ENABLED"),
				),
			},
			{
				Config: testAccSslprofile_sessionticketkeydata_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslprofileExist("citrixadc_sslprofile.tf_sslprofile_ephem", nil),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.tf_sslprofile_ephem", "name", "tf_sslprofile_ephem"),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.tf_sslprofile_ephem", "sessionticket", "ENABLED"),
				),
			},
		},
	})
}

func TestAccSslprofile_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckSslprofileDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccSslprofile_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslprofileExist("citrixadc_sslprofile.foo", nil),
				),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccSslprofile_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslprofileExist("citrixadc_sslprofile.foo", nil),
				),
			},
		},
	})
}

// Test ephemeral path: using sessionticketkeydata_wo (WriteOnly attribute) with version tracker
const testAccSslprofile_sessionticketkeydata_wo_step1 = `
	variable "sslprofile_sessionticketkeydata_wo" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_sslprofile" "tf_sslprofile_ephem" {
		name                              = "tf_sslprofile_ephem"
		sessionticket                     = "ENABLED"
		sessionticketkeydata_wo           = var.sslprofile_sessionticketkeydata_wo
		sessionticketkeydata_wo_version   = 1
	}
`

const testAccSslprofile_sessionticketkeydata_wo_step2 = `
	variable "sslprofile_sessionticketkeydata_wo_2" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_sslprofile" "tf_sslprofile_ephem" {
		name                              = "tf_sslprofile_ephem"
		sessionticket                     = "ENABLED"
		sessionticketkeydata_wo           = var.sslprofile_sessionticketkeydata_wo_2
		sessionticketkeydata_wo_version   = 2
	}
`

func TestAccSslprofile_sessionticketkeydata_wo_ephemeral(t *testing.T) {
	t.Setenv("TF_VAR_sslprofile_sessionticketkeydata_wo", "0102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f20")
	t.Setenv("TF_VAR_sslprofile_sessionticketkeydata_wo_2", "2122232425262728292a2b2c2d2e2f303132333435363738393a3b3c3d3e3f40")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslprofile_sessionticketkeydata_wo_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslprofileExist("citrixadc_sslprofile.tf_sslprofile_ephem", nil),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.tf_sslprofile_ephem", "name", "tf_sslprofile_ephem"),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.tf_sslprofile_ephem", "sessionticketkeydata_wo_version", "1"),
				),
			},
			{
				Config: testAccSslprofile_sessionticketkeydata_wo_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslprofileExist("citrixadc_sslprofile.tf_sslprofile_ephem", nil),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.tf_sslprofile_ephem", "name", "tf_sslprofile_ephem"),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.tf_sslprofile_ephem", "sessionticketkeydata_wo_version", "2"),
				),
			},
		},
	})
}

// ---------------------------------------------------------------------------
// Unset support test
//
// Step 1 sets every unset-eligible attribute to a valid non-default value; step 2
// removes them from the configuration, which must drive an ?action=unset so each
// attribute reverts to its NITRO/schema default with an empty post-apply plan.
//
// Three of the 40 unset-eligible attributes are intentionally omitted from the
// non-default set (their defaults are still exercised implicitly on every create,
// so their unset wiring is still safe; they simply cannot be driven to a
// non-default value from a self-contained front-end profile config):
//   - dh: its only non-default value (ENABLED) requires a companion dhFile file
//     ("Required argument missing [dhFile, dh==ENABLED]") that cannot be provided
//     from a self-contained Terraform config; dh=DISABLED (default) applies fine.
//   - serverauth: a backend-only attribute; setting it to ENABLED on the default
//     (front-end) profile is rejected ("Specified parameters are not applicable
//     for this type of SSL profile."); serverauth=DISABLED (default) applies fine.
//   - skipclientcertpolicycheck: its non-default value (ENABLED) requires both
//     clientAuth==ENABLED and clientCert==Mandatory, and clientCert==Mandatory in
//     turn requires clientAuth==ENABLED. That coupling cannot coexist with
//     unsetting clientauth in step 2, so it is omitted; its default (DISABLED)
//     applies fine.
// ---------------------------------------------------------------------------

const testAccSslprofile_unset_step1 = `
	resource "citrixadc_sslprofile" "tf_unset" {
		name                              = "tf_test_sslprofile_unset"

		allowextendedmastersecret         = "YES"
		allowunknownsni                   = "ENABLED"
		cipherredirect                    = "ENABLED"
		clientauth                        = "ENABLED"
		clientauthuseboundcachain         = "ENABLED"
		denysslreneg                      = "NO"
		dhekeyexchangewithpsk             = "YES"
		dhkeyexpsizelimit                 = "ENABLED"
		dropreqwithnohostheader           = "YES"
		encryptedclienthello              = "ENABLED"
		encrypttriggerpktcount            = 50
		hsts                              = "ENABLED"
		includesubdomains                 = "YES"
		maxage                            = 100
		maxrenegrate                      = 10
		ocspstapling                      = "ENABLED"
		preload                           = "YES"
		prevsessionkeylifetime            = 100
		pushenctriggertimeout             = 10
		quantumsize                       = "16384"
		redirectportrewrite               = "ENABLED"
		sendclosenotify                   = "NO"
		sessionticket                     = "ENABLED"
		sessreuse                         = "DISABLED"
		snienable                         = "ENABLED"
		snihttphostmatch                  = "STRICT"
		ssl3                              = "ENABLED"
		sslclientlogs                     = "ENABLED"
		sslimaxsessperserver              = 100
		sslinterception                   = "ENABLED"
		ssliocspcheck                     = "DISABLED"
		sslireneg                         = "DISABLED"
		sslredirect                       = "ENABLED"
		ssltriggertimeout                 = 200
		strictcachecks                    = "YES"
		tls13sessionticketsperauthcontext = 2
		zerorttearlydata                  = "ENABLED"
	}
`

const testAccSslprofile_unset_step2 = `
	resource "citrixadc_sslprofile" "tf_unset" {
		name = "tf_test_sslprofile_unset"
		# All unset-eligible attributes removed from config -> provider must unset them.
	}
`

func TestAccSslprofile_unset(t *testing.T) {
	// The resource's other CRUD/datasource tests run on the default standalone
	// testbed without a skip guard, so this test mirrors that (no guard).
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslprofileDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values apply and persist.
				Config: testAccSslprofile_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslprofileExist("citrixadc_sslprofile.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.tf_unset", "allowextendedmastersecret", "YES"),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.tf_unset", "allowunknownsni", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.tf_unset", "cipherredirect", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.tf_unset", "clientauth", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.tf_unset", "clientauthuseboundcachain", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.tf_unset", "denysslreneg", "NO"),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.tf_unset", "dhekeyexchangewithpsk", "YES"),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.tf_unset", "dhkeyexpsizelimit", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.tf_unset", "dropreqwithnohostheader", "YES"),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.tf_unset", "encryptedclienthello", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.tf_unset", "encrypttriggerpktcount", "50"),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.tf_unset", "hsts", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.tf_unset", "includesubdomains", "YES"),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.tf_unset", "maxage", "100"),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.tf_unset", "maxrenegrate", "10"),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.tf_unset", "ocspstapling", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.tf_unset", "preload", "YES"),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.tf_unset", "prevsessionkeylifetime", "100"),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.tf_unset", "pushenctriggertimeout", "10"),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.tf_unset", "quantumsize", "16384"),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.tf_unset", "redirectportrewrite", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.tf_unset", "sendclosenotify", "NO"),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.tf_unset", "sessionticket", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.tf_unset", "sessreuse", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.tf_unset", "snienable", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.tf_unset", "snihttphostmatch", "STRICT"),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.tf_unset", "ssl3", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.tf_unset", "sslclientlogs", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.tf_unset", "sslimaxsessperserver", "100"),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.tf_unset", "sslinterception", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.tf_unset", "ssliocspcheck", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.tf_unset", "sslireneg", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.tf_unset", "sslredirect", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.tf_unset", "ssltriggertimeout", "200"),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.tf_unset", "strictcachecks", "YES"),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.tf_unset", "tls13sessionticketsperauthcontext", "2"),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.tf_unset", "zerorttearlydata", "ENABLED"),
				),
			},
			{
				// Removing the attributes must unset -> state reverts to NITRO defaults,
				// and the implicit post-apply plan must be empty.
				Config: testAccSslprofile_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslprofileExist("citrixadc_sslprofile.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.tf_unset", "allowextendedmastersecret", "NO"),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.tf_unset", "allowunknownsni", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.tf_unset", "cipherredirect", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.tf_unset", "clientauth", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.tf_unset", "clientauthuseboundcachain", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.tf_unset", "denysslreneg", "ALL"),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.tf_unset", "dhekeyexchangewithpsk", "NO"),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.tf_unset", "dhkeyexpsizelimit", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.tf_unset", "dropreqwithnohostheader", "NO"),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.tf_unset", "encryptedclienthello", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.tf_unset", "encrypttriggerpktcount", "45"),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.tf_unset", "hsts", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.tf_unset", "includesubdomains", "NO"),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.tf_unset", "maxage", "0"),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.tf_unset", "maxrenegrate", "0"),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.tf_unset", "ocspstapling", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.tf_unset", "preload", "NO"),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.tf_unset", "prevsessionkeylifetime", "0"),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.tf_unset", "pushenctriggertimeout", "1"),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.tf_unset", "quantumsize", "8192"),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.tf_unset", "redirectportrewrite", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.tf_unset", "sendclosenotify", "YES"),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.tf_unset", "sessionticket", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.tf_unset", "sessreuse", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.tf_unset", "snienable", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.tf_unset", "snihttphostmatch", "CERT"),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.tf_unset", "ssl3", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.tf_unset", "sslclientlogs", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.tf_unset", "sslimaxsessperserver", "10"),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.tf_unset", "sslinterception", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.tf_unset", "ssliocspcheck", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.tf_unset", "sslireneg", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.tf_unset", "sslredirect", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.tf_unset", "ssltriggertimeout", "100"),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.tf_unset", "strictcachecks", "NO"),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.tf_unset", "tls13sessionticketsperauthcontext", "1"),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.tf_unset", "zerorttearlydata", "DISABLED"),
					// Independent appliance-level confirmation the unset actually reverted values.
					testAccCheckSslprofileADCValue("tf_test_sslprofile_unset", "allowextendedmastersecret", "NO"),
					testAccCheckSslprofileADCValue("tf_test_sslprofile_unset", "hsts", "DISABLED"),
					testAccCheckSslprofileADCValue("tf_test_sslprofile_unset", "denysslreneg", "ALL"),
					testAccCheckSslprofileADCValue("tf_test_sslprofile_unset", "encrypttriggerpktcount", "45"),
					testAccCheckSslprofileADCValue("tf_test_sslprofile_unset", "quantumsize", "8192"),
					testAccCheckSslprofileADCValue("tf_test_sslprofile_unset", "sendclosenotify", "YES"),
					testAccCheckSslprofileADCValue("tf_test_sslprofile_unset", "sslimaxsessperserver", "10"),
					testAccCheckSslprofileADCValue("tf_test_sslprofile_unset", "ssltriggertimeout", "100"),
				),
			},
		},
	})
}

// testAccCheckSslprofileADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted it.
func testAccCheckSslprofileADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Sslprofile.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("sslprofile %s not found on appliance", name)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("sslprofile %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}
