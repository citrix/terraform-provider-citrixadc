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

// NOTE (referencing the "Netscaler AIGW Config for Citrix IT" runbook): tokenquota
// / quotarefreshfrequency / authtoken are BACKEND-only parameters — a frontend
// profile rejects them (errorcode 257 "AIGW Frontend Profile cannot take
// TokenQuota/QuotaRefresh/AuthToken parameters"). A backend profile REQUIRES both
// tokenquota and quotarefreshfrequency. endpointtype is mandatory on create and,
// like profiletype/authtoken, is create-only (a change forces replacement). On the
// current firmware the only valid endpointtype value is "azureopenai" (the
// runbook's "vertexai_anthropic" needs a newer build). Only tokenquota and
// quotarefreshfrequency are mutable in place.
const testAccNsaigwprofile_basic = `


resource "citrixadc_nsaigwprofile" "tf_nsaigwprofile" {
	name                  = "tf_nsaigwprofile"
	endpointtype          = "azureopenai"
	profiletype           = "backend"
	tokenquota            = 1000
	quotarefreshfrequency = 60
}

`
const testAccNsaigwprofile_update = `


resource "citrixadc_nsaigwprofile" "tf_nsaigwprofile" {
	name                  = "tf_nsaigwprofile"
	endpointtype          = "azureopenai"
	profiletype           = "backend"
	tokenquota            = 2000
	quotarefreshfrequency = 120
}

`

func TestAccNsaigwprofile_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNsaigwprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNsaigwprofile_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsaigwprofileExist("citrixadc_nsaigwprofile.tf_nsaigwprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_nsaigwprofile.tf_nsaigwprofile", "name", "tf_nsaigwprofile"),
					resource.TestCheckResourceAttr("citrixadc_nsaigwprofile.tf_nsaigwprofile", "endpointtype", "azureopenai"),
					resource.TestCheckResourceAttr("citrixadc_nsaigwprofile.tf_nsaigwprofile", "profiletype", "backend"),
					resource.TestCheckResourceAttr("citrixadc_nsaigwprofile.tf_nsaigwprofile", "tokenquota", "1000"),
					resource.TestCheckResourceAttr("citrixadc_nsaigwprofile.tf_nsaigwprofile", "quotarefreshfrequency", "60"),
				),
			},
			{
				Config: testAccNsaigwprofile_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsaigwprofileExist("citrixadc_nsaigwprofile.tf_nsaigwprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_nsaigwprofile.tf_nsaigwprofile", "name", "tf_nsaigwprofile"),
					resource.TestCheckResourceAttr("citrixadc_nsaigwprofile.tf_nsaigwprofile", "profiletype", "backend"),
					resource.TestCheckResourceAttr("citrixadc_nsaigwprofile.tf_nsaigwprofile", "tokenquota", "2000"),
					resource.TestCheckResourceAttr("citrixadc_nsaigwprofile.tf_nsaigwprofile", "quotarefreshfrequency", "120"),
				),
			},
		},
	})
}

func TestAccNsaigwprofile_import(t *testing.T) {
	const resAddr = "citrixadc_nsaigwprofile.tf_nsaigwprofile"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNsaigwprofileDestroy,
		Steps: []resource.TestStep{
			{Config: testAccNsaigwprofile_basic},
			{
				Config:            testAccNsaigwprofile_basic,
				ResourceName:      resAddr,
				ImportState:       true,
				ImportStateVerify: true,
				// authtoken is a secret not returned by NITRO, and authtoken_wo_version
				// is a write-only version tracker (defaulted to 1) that cannot round-trip
				// through import.
				ImportStateVerifyIgnore: []string{"authtoken", "authtoken_wo", "authtoken_wo_version"},
			},
		},
	})
}

func testAccCheckNsaigwprofileExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No nsaigwprofile name is set")
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
		data, err := client.FindResource("nsaigwprofile", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("nsaigwprofile %s not found", n)
		}

		return nil
	}
}

func testAccCheckNsaigwprofileDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_nsaigwprofile" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("nsaigwprofile", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("nsaigwprofile %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccNsaigwprofileDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNsaigwprofileDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_nsaigwprofile.test", "name", "tf_nsaigwprofile_ds"),
					resource.TestCheckResourceAttrSet("data.citrixadc_nsaigwprofile.test", "id"),
					resource.TestCheckResourceAttrSet("data.citrixadc_nsaigwprofile.test", "profiletype"),
					resource.TestCheckResourceAttrSet("data.citrixadc_nsaigwprofile.test", "tokenquota"),
				),
			},
		},
	})
}

const testAccNsaigwprofileDataSource_basic = `
resource "citrixadc_nsaigwprofile" "tf_nsaigwprofile_ds" {
	name                  = "tf_nsaigwprofile_ds"
	endpointtype          = "azureopenai"
	profiletype           = "backend"
	tokenquota            = 1000
	quotarefreshfrequency = 60
}

data "citrixadc_nsaigwprofile" "test" {
	name = citrixadc_nsaigwprofile.tf_nsaigwprofile_ds.name
}
`

// Test backward-compatible path: using authtoken (Sensitive attribute)
const testAccNsaigwprofile_authtoken_step1 = `
	variable "nsaigwprofile_authtoken" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_nsaigwprofile" "tf_nsaigwprofile_ephem" {
		name                  = "tf_nsaigwprofile_ephem"
		endpointtype          = "azureopenai"
		profiletype           = "backend"
		tokenquota            = 1000
		quotarefreshfrequency = 60
		authtoken             = var.nsaigwprofile_authtoken
	}
`

const testAccNsaigwprofile_authtoken_step2 = `
	variable "nsaigwprofile_authtoken_2" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_nsaigwprofile" "tf_nsaigwprofile_ephem" {
		name                  = "tf_nsaigwprofile_ephem"
		endpointtype          = "azureopenai"
		profiletype           = "backend"
		tokenquota            = 1000
		quotarefreshfrequency = 60
		authtoken             = var.nsaigwprofile_authtoken_2
	}
`

func TestAccNsaigwprofile_authtoken_backward_compat(t *testing.T) {
	t.Setenv("TF_VAR_nsaigwprofile_authtoken", "apikey1234567890")
	t.Setenv("TF_VAR_nsaigwprofile_authtoken_2", "apikey0987654321")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNsaigwprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNsaigwprofile_authtoken_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsaigwprofileExist("citrixadc_nsaigwprofile.tf_nsaigwprofile_ephem", nil),
					resource.TestCheckResourceAttr("citrixadc_nsaigwprofile.tf_nsaigwprofile_ephem", "name", "tf_nsaigwprofile_ephem"),
				),
			},
			{
				Config: testAccNsaigwprofile_authtoken_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsaigwprofileExist("citrixadc_nsaigwprofile.tf_nsaigwprofile_ephem", nil),
					resource.TestCheckResourceAttr("citrixadc_nsaigwprofile.tf_nsaigwprofile_ephem", "name", "tf_nsaigwprofile_ephem"),
				),
			},
		},
	})
}

// Test ephemeral path: using authtoken_wo (WriteOnly attribute) with version tracker
const testAccNsaigwprofile_authtoken_wo_step1 = `
	variable "nsaigwprofile_authtoken_wo" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_nsaigwprofile" "tf_nsaigwprofile_ephem" {
		name                  = "tf_nsaigwprofile_ephem"
		endpointtype          = "azureopenai"
		profiletype           = "backend"
		tokenquota            = 1000
		quotarefreshfrequency = 60
		authtoken_wo          = var.nsaigwprofile_authtoken_wo
		authtoken_wo_version  = 1
	}
`

const testAccNsaigwprofile_authtoken_wo_step2 = `
	variable "nsaigwprofile_authtoken_wo_2" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_nsaigwprofile" "tf_nsaigwprofile_ephem" {
		name                  = "tf_nsaigwprofile_ephem"
		endpointtype          = "azureopenai"
		profiletype           = "backend"
		tokenquota            = 1000
		quotarefreshfrequency = 60
		authtoken_wo          = var.nsaigwprofile_authtoken_wo_2
		authtoken_wo_version  = 2
	}
`

func TestAccNsaigwprofile_authtoken_wo_ephemeral(t *testing.T) {
	t.Setenv("TF_VAR_nsaigwprofile_authtoken_wo", "ephem_apikey1")
	t.Setenv("TF_VAR_nsaigwprofile_authtoken_wo_2", "ephem_apikey2")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNsaigwprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNsaigwprofile_authtoken_wo_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsaigwprofileExist("citrixadc_nsaigwprofile.tf_nsaigwprofile_ephem", nil),
					resource.TestCheckResourceAttr("citrixadc_nsaigwprofile.tf_nsaigwprofile_ephem", "name", "tf_nsaigwprofile_ephem"),
					resource.TestCheckResourceAttr("citrixadc_nsaigwprofile.tf_nsaigwprofile_ephem", "authtoken_wo_version", "1"),
				),
			},
			{
				Config: testAccNsaigwprofile_authtoken_wo_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsaigwprofileExist("citrixadc_nsaigwprofile.tf_nsaigwprofile_ephem", nil),
					resource.TestCheckResourceAttr("citrixadc_nsaigwprofile.tf_nsaigwprofile_ephem", "name", "tf_nsaigwprofile_ephem"),
					resource.TestCheckResourceAttr("citrixadc_nsaigwprofile.tf_nsaigwprofile_ephem", "authtoken_wo_version", "2"),
				),
			},
		},
	})
}

// Step 1: a backend profile with the unset-eligible attribute (quotarefreshfrequency)
// set to a non-default value. tokenquota is required for a backend profile and is
// NOT unsettable (NITRO leaves it unchanged on ?action=unset), so it is kept across
// both steps; endpointtype/profiletype are create-only and kept as well.
const testAccNsaigwprofile_unset_step1 = `
resource "citrixadc_nsaigwprofile" "tf_unset" {
	name                  = "tf_test_nsaigwprofile_unset"
	endpointtype          = "azureopenai"
	profiletype           = "backend"
	tokenquota            = 5000
	quotarefreshfrequency = 240
}
`

// Step 2: quotarefreshfrequency removed from config -> provider must unset it,
// reverting it on the appliance while the profile itself remains.
const testAccNsaigwprofile_unset_step2 = `
resource "citrixadc_nsaigwprofile" "tf_unset" {
	name         = "tf_test_nsaigwprofile_unset"
	endpointtype = "azureopenai"
	profiletype  = "backend"
	tokenquota   = 5000
}
`

func TestAccNsaigwprofile_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNsaigwprofileDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values apply and persist.
				Config: testAccNsaigwprofile_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsaigwprofileExist("citrixadc_nsaigwprofile.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_nsaigwprofile.tf_unset", "tokenquota", "5000"),
					resource.TestCheckResourceAttr("citrixadc_nsaigwprofile.tf_unset", "quotarefreshfrequency", "240"),
				),
			},
			{
				// Removing quotarefreshfrequency must unset it on the appliance while
				// tokenquota (kept in config) is unchanged. The implicit post-apply
				// non-refresh plan must be empty (proving the unset applied cleanly
				// with no churn).
				Config: testAccNsaigwprofile_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsaigwprofileExist("citrixadc_nsaigwprofile.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_nsaigwprofile.tf_unset", "tokenquota", "5000"),
					// Independent appliance-level confirmation tokenquota persisted.
					testAccCheckNsaigwprofileADCValue("tf_test_nsaigwprofile_unset", "tokenquota", "5000"),
				),
			},
		},
	})
}

// testAccCheckNsaigwprofileADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted it.
func testAccCheckNsaigwprofileADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Nsaigwprofile.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("nsaigwprofile %s not found on appliance", name)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("nsaigwprofile %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

// TestAccNsaigwprofile_selfHealing verifies the provider re-creates the profile when
// it is deleted out-of-band between apply steps (drift recovery).
func TestAccNsaigwprofile_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_nsaigwprofile.tf_nsaigwprofile"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNsaigwprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNsaigwprofile_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckNsaigwprofileExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Nsaigwprofile.Type(), "tf_nsaigwprofile"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccNsaigwprofile_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckNsaigwprofileExist(resAddr, nil)),
			},
		},
	})
}
