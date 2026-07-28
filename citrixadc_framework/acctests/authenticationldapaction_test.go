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

const testAccAuthenticationldapaction_add = `
	resource "citrixadc_authenticationldapaction" "foo" {
		name   		  = "ldapaction"
		serverip 	  = "1.2.3.4"
		serverport 	  = 8080
		authtimeout   = 1
		ldaploginname = "username"
	}
`
const testAccAuthenticationldapaction_update = `
	resource "citrixadc_authenticationldapaction" "foo" {
		name   		  = "ldapaction"
		serverip	  = "1.2.4.5"
		serverport    = 8000
		authtimeout   = 2
		ldaploginname = "username"
	}
`

func TestAccAuthenticationldapaction_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationldapactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationldapaction_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationldapactionExist("citrixadc_authenticationldapaction.foo", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationldapaction.foo", "name", "ldapaction"),
					resource.TestCheckResourceAttr("citrixadc_authenticationldapaction.foo", "serverip", "1.2.3.4"),
					resource.TestCheckResourceAttr("citrixadc_authenticationldapaction.foo", "authtimeout", "1"),
				),
			},
			{
				Config: testAccAuthenticationldapaction_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationldapactionExist("citrixadc_authenticationldapaction.foo", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationldapaction.foo", "name", "ldapaction"),
					resource.TestCheckResourceAttr("citrixadc_authenticationldapaction.foo", "serverip", "1.2.4.5"),
					resource.TestCheckResourceAttr("citrixadc_authenticationldapaction.foo", "authtimeout", "2"),
				),
			},
		},
	})
}

func TestAccAuthenticationldapaction_import(t *testing.T) {
	const resAddr = "citrixadc_authenticationldapaction.foo"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationldapactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationldapaction_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationldapactionExist(resAddr, nil),
				),
			},
			{
				Config:            testAccAuthenticationldapaction_add,
				ResourceName:      resAddr,
				ImportState:       true,
				ImportStateVerify: true,
				// ldapbinddnpassword_wo_version is a write-only version tracker
				// that is not stored on the ADC and cannot round-trip through import.
				ImportStateVerifyIgnore: []string{"ldapbinddnpassword_wo_version"},
			},
		},
	})
}

func testAccCheckAuthenticationldapactionExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No authenticationldapaction name is set")
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
		data, err := client.FindResource(service.Authenticationldapaction.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("authenticationldapaction %s not found", n)
		}

		return nil
	}
}

func testAccCheckAuthenticationldapactionDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_authenticationldapaction" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Authenticationldapaction.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("authenticationldapaction %s still exists", rs.Primary.ID)
		}
	}

	return nil
}

// Test backward-compatible path: using ldapbinddnpassword (Sensitive attribute)
const testAccAuthenticationldapaction_ldapbinddnpassword_step1 = `

	variable "ldapaction_ldapbinddnpassword" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_authenticationldapaction" "foo" {
		name               = "ldapaction"
		serverip           = "1.2.3.4"
		serverport         = 8080
		authtimeout        = 1
		ldaploginname      = "username"
		ldapbinddnpassword = var.ldapaction_ldapbinddnpassword
	}
`

const testAccAuthenticationldapaction_ldapbinddnpassword_step2 = `

	variable "ldapaction_ldapbinddnpassword_2" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_authenticationldapaction" "foo" {
		name               = "ldapaction"
		serverip           = "1.2.3.4"
		serverport         = 8080
		authtimeout        = 1
		ldaploginname      = "username"
		ldapbinddnpassword = var.ldapaction_ldapbinddnpassword_2
	}
`

func TestAccAuthenticationldapaction_ldapbinddnpassword_backward_compat(t *testing.T) {
	t.Setenv("TF_VAR_ldapaction_ldapbinddnpassword", "oldldappass123")
	t.Setenv("TF_VAR_ldapaction_ldapbinddnpassword_2", "newldappass456")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationldapactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationldapaction_ldapbinddnpassword_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationldapactionExist("citrixadc_authenticationldapaction.foo", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationldapaction.foo", "name", "ldapaction"),
					resource.TestCheckResourceAttr("citrixadc_authenticationldapaction.foo", "serverip", "1.2.3.4"),
				),
			},
			{
				Config: testAccAuthenticationldapaction_ldapbinddnpassword_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationldapactionExist("citrixadc_authenticationldapaction.foo", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationldapaction.foo", "name", "ldapaction"),
					resource.TestCheckResourceAttr("citrixadc_authenticationldapaction.foo", "serverip", "1.2.3.4"),
				),
			},
		},
	})
}

// Test ephemeral path: using ldapbinddnpassword_wo (WriteOnly attribute) with version tracker
const testAccAuthenticationldapaction_ldapbinddnpassword_wo_step1 = `

	variable "ldapaction_ldapbinddnpassword_wo" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_authenticationldapaction" "foo" {
		name                          = "ldapaction"
		serverip                      = "1.2.3.4"
		serverport                    = 8080
		authtimeout                   = 1
		ldaploginname                 = "username"
		ldapbinddnpassword_wo         = var.ldapaction_ldapbinddnpassword_wo
		ldapbinddnpassword_wo_version = 1
	}
`

const testAccAuthenticationldapaction_ldapbinddnpassword_wo_step2 = `

	variable "ldapaction_ldapbinddnpassword_wo_2" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_authenticationldapaction" "foo" {
		name                          = "ldapaction"
		serverip                      = "1.2.3.4"
		serverport                    = 8080
		authtimeout                   = 1
		ldaploginname                 = "username"
		ldapbinddnpassword_wo         = var.ldapaction_ldapbinddnpassword_wo_2
		ldapbinddnpassword_wo_version = 2
	}
`

func TestAccAuthenticationldapaction_ldapbinddnpassword_wo_ephemeral(t *testing.T) {
	t.Setenv("TF_VAR_ldapaction_ldapbinddnpassword_wo", "ephemeral_ldap1")
	t.Setenv("TF_VAR_ldapaction_ldapbinddnpassword_wo_2", "ephemeral_ldap2")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationldapactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationldapaction_ldapbinddnpassword_wo_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationldapactionExist("citrixadc_authenticationldapaction.foo", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationldapaction.foo", "ldapbinddnpassword_wo_version", "1"),
				),
			},
			{
				Config: testAccAuthenticationldapaction_ldapbinddnpassword_wo_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationldapactionExist("citrixadc_authenticationldapaction.foo", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationldapaction.foo", "ldapbinddnpassword_wo_version", "2"),
				),
			},
		},
	})
}

func TestAccAuthenticationldapaction_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckAuthenticationldapactionDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccAuthenticationldapaction_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationldapactionExist("citrixadc_authenticationldapaction.foo", nil),
				),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccAuthenticationldapaction_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationldapactionExist("citrixadc_authenticationldapaction.foo", nil),
				),
			},
		},
	})
}

const testAccAuthenticationldapactionDataSource_basic = `
	resource "citrixadc_authenticationldapaction" "tf_authenticationldapaction" {
		name   		  = "tf_authenticationldapaction"
		serverip 	  = "1.2.3.4"
		serverport 	  = 8080
		authtimeout   = 1
		ldaploginname = "username"
	}

	data "citrixadc_authenticationldapaction" "tf_authenticationldapaction" {
		name = citrixadc_authenticationldapaction.tf_authenticationldapaction.name
	}
`

// Unset test: step1 sets every unset-eligible attribute to a non-default value;
// step2 removes them from config so the provider must issue ?action=unset and the
// appliance reverts each to its NITRO default. groupnameidentifier and
// groupsearchattribute are kept in both steps because nestedgroupextraction=ON
// (a step1 non-default) requires them as prerequisites; they are not unset-eligible.
const testAccAuthenticationldapaction_unset_step1 = `
	resource "citrixadc_authenticationldapaction" "tf_unset" {
		name                  = "tf_test_authenticationldapaction_unset"
		serverip              = "1.2.3.4"
		groupnameidentifier   = "samAccountName"
		groupsearchattribute  = "memberOf"
		authentication        = "DISABLED"
		authtimeout           = 10
		cloudattributes       = "ENABLED"
		email                 = "mail2"
		followreferrals       = "ON"
		nestedgroupextraction = "ON"
		passwdchange          = "ENABLED"
		referraldnslookup     = "SRV-REC"
		requireuser           = "NO"
		sectype               = "TLS"
		serverport            = 636
		validateservercert    = "YES"
	}
`

const testAccAuthenticationldapaction_unset_step2 = `
	resource "citrixadc_authenticationldapaction" "tf_unset" {
		name                  = "tf_test_authenticationldapaction_unset"
		serverip              = "1.2.3.4"
		groupnameidentifier   = "samAccountName"
		groupsearchattribute  = "memberOf"
		# unset-eligible attributes removed from config -> provider must unset them
	}
`

func TestAccAuthenticationldapaction_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationldapactionDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values apply and persist.
				Config: testAccAuthenticationldapaction_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationldapactionExist("citrixadc_authenticationldapaction.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationldapaction.tf_unset", "authentication", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_authenticationldapaction.tf_unset", "authtimeout", "10"),
					resource.TestCheckResourceAttr("citrixadc_authenticationldapaction.tf_unset", "cloudattributes", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_authenticationldapaction.tf_unset", "email", "mail2"),
					resource.TestCheckResourceAttr("citrixadc_authenticationldapaction.tf_unset", "followreferrals", "ON"),
					resource.TestCheckResourceAttr("citrixadc_authenticationldapaction.tf_unset", "nestedgroupextraction", "ON"),
					resource.TestCheckResourceAttr("citrixadc_authenticationldapaction.tf_unset", "passwdchange", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_authenticationldapaction.tf_unset", "referraldnslookup", "SRV-REC"),
					resource.TestCheckResourceAttr("citrixadc_authenticationldapaction.tf_unset", "requireuser", "NO"),
					resource.TestCheckResourceAttr("citrixadc_authenticationldapaction.tf_unset", "sectype", "TLS"),
					resource.TestCheckResourceAttr("citrixadc_authenticationldapaction.tf_unset", "serverport", "636"),
					resource.TestCheckResourceAttr("citrixadc_authenticationldapaction.tf_unset", "validateservercert", "YES"),
				),
			},
			{
				// Removing them must unset -> state reverts to NITRO defaults,
				// and the implicit post-apply plan must be empty.
				Config: testAccAuthenticationldapaction_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationldapactionExist("citrixadc_authenticationldapaction.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationldapaction.tf_unset", "authentication", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_authenticationldapaction.tf_unset", "authtimeout", "3"),
					resource.TestCheckResourceAttr("citrixadc_authenticationldapaction.tf_unset", "cloudattributes", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_authenticationldapaction.tf_unset", "email", "mail"),
					resource.TestCheckResourceAttr("citrixadc_authenticationldapaction.tf_unset", "followreferrals", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_authenticationldapaction.tf_unset", "nestedgroupextraction", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_authenticationldapaction.tf_unset", "passwdchange", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_authenticationldapaction.tf_unset", "referraldnslookup", "A-REC"),
					resource.TestCheckResourceAttr("citrixadc_authenticationldapaction.tf_unset", "requireuser", "YES"),
					resource.TestCheckResourceAttr("citrixadc_authenticationldapaction.tf_unset", "sectype", "PLAINTEXT"),
					resource.TestCheckResourceAttr("citrixadc_authenticationldapaction.tf_unset", "serverport", "389"),
					resource.TestCheckResourceAttr("citrixadc_authenticationldapaction.tf_unset", "validateservercert", "NO"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckAuthenticationldapactionADCValue("tf_test_authenticationldapaction_unset", "authentication", "ENABLED"),
					testAccCheckAuthenticationldapactionADCValue("tf_test_authenticationldapaction_unset", "authtimeout", "3"),
					testAccCheckAuthenticationldapactionADCValue("tf_test_authenticationldapaction_unset", "cloudattributes", "DISABLED"),
					testAccCheckAuthenticationldapactionADCValue("tf_test_authenticationldapaction_unset", "email", "mail"),
					testAccCheckAuthenticationldapactionADCValue("tf_test_authenticationldapaction_unset", "followreferrals", "OFF"),
					testAccCheckAuthenticationldapactionADCValue("tf_test_authenticationldapaction_unset", "nestedgroupextraction", "OFF"),
					testAccCheckAuthenticationldapactionADCValue("tf_test_authenticationldapaction_unset", "passwdchange", "DISABLED"),
					testAccCheckAuthenticationldapactionADCValue("tf_test_authenticationldapaction_unset", "referraldnslookup", "A-REC"),
					testAccCheckAuthenticationldapactionADCValue("tf_test_authenticationldapaction_unset", "requireuser", "YES"),
					testAccCheckAuthenticationldapactionADCValue("tf_test_authenticationldapaction_unset", "sectype", "PLAINTEXT"),
					testAccCheckAuthenticationldapactionADCValue("tf_test_authenticationldapaction_unset", "serverport", "389"),
					testAccCheckAuthenticationldapactionADCValue("tf_test_authenticationldapaction_unset", "validateservercert", "NO"),
				),
			},
		},
	})
}

// testAccCheckAuthenticationldapactionADCValue asserts an attribute's value
// directly on the appliance (not just in Terraform state), proving the unset
// actually reverted it.
func testAccCheckAuthenticationldapactionADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Authenticationldapaction.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("authenticationldapaction %s not found on appliance", name)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("authenticationldapaction %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func TestAccAuthenticationldapactionDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationldapactionDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_authenticationldapaction.tf_authenticationldapaction", "name", "tf_authenticationldapaction"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationldapaction.tf_authenticationldapaction", "serverip", "1.2.3.4"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationldapaction.tf_authenticationldapaction", "serverport", "8080"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationldapaction.tf_authenticationldapaction", "authtimeout", "1"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationldapaction.tf_authenticationldapaction", "ldaploginname", "username"),
				),
			},
		},
	})
}
