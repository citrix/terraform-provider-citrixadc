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

const testAccAaaldapparams_basic = `
	resource "citrixadc_aaaldapparams" "tf_aaaldapparams" {
		authtimeout = 4
		serverip    = "10.222.74.158"
		passwdchange = "DISABLED"
	}
  
`
const testAccAaaldapparams_update = `
	resource "citrixadc_aaaldapparams" "tf_aaaldapparams" {
		authtimeout = 5
		serverip    = "10.222.74.158"
		passwdchange = "ENABLED"
	}
  
`

func TestAccAaaldapparams_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccAaaldapparams_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaaldapparamsExist("citrixadc_aaaldapparams.tf_aaaldapparams", nil),
					resource.TestCheckResourceAttr("citrixadc_aaaldapparams.tf_aaaldapparams", "authtimeout", "4"),
					resource.TestCheckResourceAttr("citrixadc_aaaldapparams.tf_aaaldapparams", "serverip", "10.222.74.158"),
					resource.TestCheckResourceAttr("citrixadc_aaaldapparams.tf_aaaldapparams", "passwdchange", "DISABLED"),
				),
			},
			{
				Config: testAccAaaldapparams_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaaldapparamsExist("citrixadc_aaaldapparams.tf_aaaldapparams", nil),
					resource.TestCheckResourceAttr("citrixadc_aaaldapparams.tf_aaaldapparams", "authtimeout", "5"),
					resource.TestCheckResourceAttr("citrixadc_aaaldapparams.tf_aaaldapparams", "serverip", "10.222.74.158"),
					resource.TestCheckResourceAttr("citrixadc_aaaldapparams.tf_aaaldapparams", "passwdchange", "ENABLED"),
				),
			},
		},
	})
}

func TestAccAaaldapparams_import(t *testing.T) {
	const resAddr = "citrixadc_aaaldapparams.tf_aaaldapparams"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{Config: testAccAaaldapparams_basic},
			{
				Config:            testAccAaaldapparams_basic,
				ResourceName:      resAddr,
				ImportState:       true,
				ImportStateVerify: true,
				// ldapbinddnpassword_wo_version is a write-only version tracker that
				// NITRO does not return; it cannot round-trip through import.
				ImportStateVerifyIgnore: []string{"ldapbinddnpassword_wo_version"},
			},
		},
	})
}

func testAccCheckAaaldapparamsExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No aaaldapparams name is set")
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
		data, err := client.FindResource(service.Aaaldapparams.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("aaaldapparams %s not found", n)
		}

		return nil
	}
}

// Add this to the end of aaaldapparams_test.go

const testAccAaaldapparamsDataSource_basic = `

resource "citrixadc_aaaldapparams" "tf_aaaldapparams" {
	serverip = "1.2.3.4"
	serverport = 389
	authtimeout = 5
	ldapbase = "dc=aaa,dc=local"
	ldapbinddn = "cn=Manager,dc=aaa,dc=local"
	ldapbinddnpassword = "secret"
	ldaploginname = "samAccountName"
	searchfilter = "cn"
	groupattrname = "memberOf"
	subattributename = "cn"
	sectype = "PLAINTEXT"
	passwdchange = "ENABLED"
	nestedgroupextraction = "OFF"
	maxnestinglevel = 3
	groupnameidentifier = "samAccountName"
	groupsearchattribute = "memberOf"
	groupsearchsubattribute = "cn"
	groupsearchfilter = "memberOf"
	defaultauthenticationgroup = "default_group"
}

data "citrixadc_aaaldapparams" "tf_aaaldapparams" {
	depends_on = [citrixadc_aaaldapparams.tf_aaaldapparams]
}
`

// --- Ephemeral / Write-Only Tests for ldapbinddnpassword ---

const testAccAaaldapparams_ldapbinddnpassword_step1 = `
	variable "aaaldapparams_ldapbinddnpassword" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_aaaldapparams" "tf_aaaldapparams" {
		serverip            = "10.222.74.158"
		authtimeout         = 4
		passwdchange        = "DISABLED"
		ldapbinddnpassword  = var.aaaldapparams_ldapbinddnpassword
	}
`

const testAccAaaldapparams_ldapbinddnpassword_step2 = `
	variable "aaaldapparams_ldapbinddnpassword_2" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_aaaldapparams" "tf_aaaldapparams" {
		serverip            = "10.222.74.158"
		authtimeout         = 4
		passwdchange        = "DISABLED"
		ldapbinddnpassword  = var.aaaldapparams_ldapbinddnpassword_2
	}
`

func TestAccAaaldapparams_ldapbinddnpassword_backward_compat(t *testing.T) {
	t.Setenv("TF_VAR_aaaldapparams_ldapbinddnpassword", "value1")
	t.Setenv("TF_VAR_aaaldapparams_ldapbinddnpassword_2", "value2")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccAaaldapparams_ldapbinddnpassword_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaaldapparamsExist("citrixadc_aaaldapparams.tf_aaaldapparams", nil),
					resource.TestCheckResourceAttr("citrixadc_aaaldapparams.tf_aaaldapparams", "serverip", "10.222.74.158"),
				),
			},
			{
				Config: testAccAaaldapparams_ldapbinddnpassword_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaaldapparamsExist("citrixadc_aaaldapparams.tf_aaaldapparams", nil),
					resource.TestCheckResourceAttr("citrixadc_aaaldapparams.tf_aaaldapparams", "serverip", "10.222.74.158"),
				),
			},
		},
	})
}

const testAccAaaldapparams_ldapbinddnpassword_wo_step1 = `
	variable "aaaldapparams_ldapbinddnpassword_wo" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_aaaldapparams" "tf_aaaldapparams" {
		serverip                     = "10.222.74.158"
		authtimeout                  = 4
		passwdchange                 = "DISABLED"
		ldapbinddnpassword_wo        = var.aaaldapparams_ldapbinddnpassword_wo
		ldapbinddnpassword_wo_version = 1
	}
`

const testAccAaaldapparams_ldapbinddnpassword_wo_step2 = `
	variable "aaaldapparams_ldapbinddnpassword_wo_2" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_aaaldapparams" "tf_aaaldapparams" {
		serverip                     = "10.222.74.158"
		authtimeout                  = 4
		passwdchange                 = "DISABLED"
		ldapbinddnpassword_wo        = var.aaaldapparams_ldapbinddnpassword_wo_2
		ldapbinddnpassword_wo_version = 2
	}
`

func TestAccAaaldapparams_ldapbinddnpassword_wo_ephemeral(t *testing.T) {
	t.Setenv("TF_VAR_aaaldapparams_ldapbinddnpassword_wo", "ephemeral_value1")
	t.Setenv("TF_VAR_aaaldapparams_ldapbinddnpassword_wo_2", "ephemeral_value2")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccAaaldapparams_ldapbinddnpassword_wo_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaaldapparamsExist("citrixadc_aaaldapparams.tf_aaaldapparams", nil),
					resource.TestCheckResourceAttr("citrixadc_aaaldapparams.tf_aaaldapparams", "ldapbinddnpassword_wo_version", "1"),
				),
			},
			{
				Config: testAccAaaldapparams_ldapbinddnpassword_wo_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaaldapparamsExist("citrixadc_aaaldapparams.tf_aaaldapparams", nil),
					resource.TestCheckResourceAttr("citrixadc_aaaldapparams.tf_aaaldapparams", "ldapbinddnpassword_wo_version", "2"),
				),
			},
		},
	})
}

func TestAccAaaldapparams_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccAaaldapparams_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaaldapparamsExist("citrixadc_aaaldapparams.tf_aaaldapparams", nil),
				),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccAaaldapparams_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaaldapparamsExist("citrixadc_aaaldapparams.tf_aaaldapparams", nil),
				),
			},
		},
	})
}

func TestAccAaaldapparamsDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAaaldapparamsDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_aaaldapparams.tf_aaaldapparams", "serverip", "1.2.3.4"),
					resource.TestCheckResourceAttr("data.citrixadc_aaaldapparams.tf_aaaldapparams", "serverport", "389"),
					resource.TestCheckResourceAttr("data.citrixadc_aaaldapparams.tf_aaaldapparams", "authtimeout", "5"),
					resource.TestCheckResourceAttr("data.citrixadc_aaaldapparams.tf_aaaldapparams", "ldapbase", "dc=aaa,dc=local"),
					resource.TestCheckResourceAttr("data.citrixadc_aaaldapparams.tf_aaaldapparams", "ldapbinddn", "cn=Manager,dc=aaa,dc=local"),
					resource.TestCheckResourceAttr("data.citrixadc_aaaldapparams.tf_aaaldapparams", "ldaploginname", "samAccountName"),
					resource.TestCheckResourceAttr("data.citrixadc_aaaldapparams.tf_aaaldapparams", "searchfilter", "cn"),
					resource.TestCheckResourceAttr("data.citrixadc_aaaldapparams.tf_aaaldapparams", "groupattrname", "memberOf"),
					resource.TestCheckResourceAttr("data.citrixadc_aaaldapparams.tf_aaaldapparams", "subattributename", "cn"),
					resource.TestCheckResourceAttr("data.citrixadc_aaaldapparams.tf_aaaldapparams", "sectype", "PLAINTEXT"),
					resource.TestCheckResourceAttr("data.citrixadc_aaaldapparams.tf_aaaldapparams", "passwdchange", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_aaaldapparams.tf_aaaldapparams", "nestedgroupextraction", "OFF"),
					resource.TestCheckResourceAttr("data.citrixadc_aaaldapparams.tf_aaaldapparams", "maxnestinglevel", "3"),
					resource.TestCheckResourceAttr("data.citrixadc_aaaldapparams.tf_aaaldapparams", "groupnameidentifier", "samAccountName"),
					resource.TestCheckResourceAttr("data.citrixadc_aaaldapparams.tf_aaaldapparams", "groupsearchattribute", "memberOf"),
					resource.TestCheckResourceAttr("data.citrixadc_aaaldapparams.tf_aaaldapparams", "groupsearchsubattribute", "cn"),
					resource.TestCheckResourceAttr("data.citrixadc_aaaldapparams.tf_aaaldapparams", "groupsearchfilter", "memberOf"),
					resource.TestCheckResourceAttr("data.citrixadc_aaaldapparams.tf_aaaldapparams", "defaultauthenticationgroup", "default_group"),
				),
			},
		},
	})
}

// --- Unset Test ---
//
// Proves aaaldapparams unset support end to end for the 6 unset-eligible
// attributes (authtimeout, maxnestinglevel, nestedgroupextraction,
// passwdchange, sectype, serverport): step 1 sets them to non-default values;
// step 2 removes them from config so the provider issues ?action=unset and each
// reverts to its NITRO default, both in Terraform state and on the appliance.
//
// nestedgroupextraction=ON requires groupnameidentifier and groupsearchattribute
// (NITRO inter-field prerequisite), so those non-eligible fields are supplied in
// both steps; they are not unset.

const testAccAaaldapparams_unset_step1 = `
resource "citrixadc_aaaldapparams" "tf_unset" {
	serverip              = "10.222.74.158"
	groupnameidentifier   = "samAccountName"
	groupsearchattribute  = "memberOf"
	authtimeout           = 10
	maxnestinglevel       = 3
	nestedgroupextraction = "ON"
	passwdchange          = "ENABLED"
	sectype               = "PLAINTEXT"
	serverport            = 636
}
`

const testAccAaaldapparams_unset_step2 = `
resource "citrixadc_aaaldapparams" "tf_unset" {
	serverip             = "10.222.74.158"
	groupnameidentifier  = "samAccountName"
	groupsearchattribute = "memberOf"
	# authtimeout, maxnestinglevel, nestedgroupextraction, passwdchange,
	# sectype and serverport removed from config -> provider must unset them.
}
`

func TestAccAaaldapparams_unset(t *testing.T) {
	// The resource's other tests run on the default standalone testbed with no
	// skip guards, so none is added here.
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil, // singleton resource - never truly deleted
		Steps: []resource.TestStep{
			{
				// Non-default values apply and persist.
				Config: testAccAaaldapparams_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaaldapparamsExist("citrixadc_aaaldapparams.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_aaaldapparams.tf_unset", "authtimeout", "10"),
					resource.TestCheckResourceAttr("citrixadc_aaaldapparams.tf_unset", "maxnestinglevel", "3"),
					resource.TestCheckResourceAttr("citrixadc_aaaldapparams.tf_unset", "nestedgroupextraction", "ON"),
					resource.TestCheckResourceAttr("citrixadc_aaaldapparams.tf_unset", "passwdchange", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_aaaldapparams.tf_unset", "sectype", "PLAINTEXT"),
					resource.TestCheckResourceAttr("citrixadc_aaaldapparams.tf_unset", "serverport", "636"),
				),
			},
			{
				// Removing them must unset -> state reverts to NITRO defaults,
				// and the implicit post-apply plan must be empty.
				Config: testAccAaaldapparams_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaaldapparamsExist("citrixadc_aaaldapparams.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_aaaldapparams.tf_unset", "authtimeout", "3"),
					resource.TestCheckResourceAttr("citrixadc_aaaldapparams.tf_unset", "maxnestinglevel", "2"),
					resource.TestCheckResourceAttr("citrixadc_aaaldapparams.tf_unset", "nestedgroupextraction", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_aaaldapparams.tf_unset", "passwdchange", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_aaaldapparams.tf_unset", "sectype", "TLS"),
					resource.TestCheckResourceAttr("citrixadc_aaaldapparams.tf_unset", "serverport", "389"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckAaaldapparamsADCValue("authtimeout", "3"),
					testAccCheckAaaldapparamsADCValue("maxnestinglevel", "2"),
					testAccCheckAaaldapparamsADCValue("nestedgroupextraction", "OFF"),
					testAccCheckAaaldapparamsADCValue("passwdchange", "DISABLED"),
					testAccCheckAaaldapparamsADCValue("sectype", "TLS"),
					testAccCheckAaaldapparamsADCValue("serverport", "389"),
				),
			},
		},
	})
}

// testAccCheckAaaldapparamsADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted
// it. aaaldapparams is a singleton, so the resource is fetched with an empty name.
func testAccCheckAaaldapparamsADCValue(attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Aaaldapparams.Type(), "")
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("aaaldapparams not found on appliance")
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("aaaldapparams: appliance attr %q = %q, want %q (unset did not revert it)", attr, got, want)
		}
		return nil
	}
}
