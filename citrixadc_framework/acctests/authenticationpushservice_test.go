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

const testAccAuthenticationpushservice_add = `
	resource "citrixadc_authenticationpushservice" "tf_pushservice" {
		name            = "tf_pushservice"
		clientid        = "cliId"
		clientsecret    = "secret"
		customerid      = "cusID"
		refreshinterval = 50
	}
`
const testAccAuthenticationpushservice_update = `
	resource "citrixadc_authenticationpushservice" "tf_pushservice" {
		name            = "tf_pushservice"
		clientid        = "cliId"
		clientsecret    = "secret"
		customerid      = "cusID1"
		refreshinterval = 80
	}
`

func TestAccAuthenticationpushservice_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationpushserviceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationpushservice_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationpushserviceExist("citrixadc_authenticationpushservice.tf_pushservice", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationpushservice.tf_pushservice", "name", "tf_pushservice"),
					resource.TestCheckResourceAttr("citrixadc_authenticationpushservice.tf_pushservice", "refreshinterval", "50"),
				),
			},
			{
				Config: testAccAuthenticationpushservice_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationpushserviceExist("citrixadc_authenticationpushservice.tf_pushservice", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationpushservice.tf_pushservice", "name", "tf_pushservice"),
					resource.TestCheckResourceAttr("citrixadc_authenticationpushservice.tf_pushservice", "refreshinterval", "80"),
				),
			},
		},
	})
}

func TestAccAuthenticationpushservice_import(t *testing.T) {
	const resAddr = "citrixadc_authenticationpushservice.tf_pushservice"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationpushserviceDestroy,
		Steps: []resource.TestStep{
			{Config: testAccAuthenticationpushservice_add},
			{
				Config:            testAccAuthenticationpushservice_add,
				ResourceName:      resAddr,
				ImportState:       true,
				ImportStateVerify: true,
				// clientsecret is Sensitive and not returned by the NITRO API;
				// clientsecret_wo_version is a write-only version tracker that is
				// retained from config and not echoed back. Neither can round-trip
				// through import.
				ImportStateVerifyIgnore: []string{"clientsecret", "clientsecret_wo_version"},
			},
		},
	})
}

func testAccCheckAuthenticationpushserviceExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No authenticationpushservice name is set")
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
		data, err := client.FindResource("authenticationpushservice", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("authenticationpushservice %s not found", n)
		}

		return nil
	}
}

func testAccCheckAuthenticationpushserviceDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_authenticationpushservice" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("authenticationpushservice", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("authenticationpushservice %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccAuthenticationpushserviceDataSource_basic = `
	resource "citrixadc_authenticationpushservice" "tf_pushservice" {
		name            = "tf_pushservice_ds"
		clientid        = "cliId"
		clientsecret    = "secret"
		customerid      = "cusID"
		refreshinterval = 50
	}

	data "citrixadc_authenticationpushservice" "tf_pushservice_data" {
		name = citrixadc_authenticationpushservice.tf_pushservice.name
	}
`

func TestAccAuthenticationpushserviceDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationpushserviceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationpushserviceDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_authenticationpushservice.tf_pushservice_data", "name", "tf_pushservice_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationpushservice.tf_pushservice_data", "customerid", "cusID"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationpushservice.tf_pushservice_data", "refreshinterval", "50"),
				),
			},
		},
	})
}

// Test backward-compatible path: using clientsecret (Sensitive attribute)
const testAccAuthenticationpushservice_clientsecret_step1 = `

	variable "authenticationpushservice_clientsecret" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_authenticationpushservice" "test" {
		name            = "tf_pushservice_compat"
		clientid        = "cliId_compat"
		clientsecret    = var.authenticationpushservice_clientsecret
		customerid      = "cusID_compat"
		refreshinterval = 50
	}
`

// Update backward-compatible path: change clientsecret value
const testAccAuthenticationpushservice_clientsecret_step2 = `

	variable "authenticationpushservice_clientsecret_2" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_authenticationpushservice" "test" {
		name            = "tf_pushservice_compat"
		clientid        = "cliId_compat"
		clientsecret    = var.authenticationpushservice_clientsecret_2
		customerid      = "cusID_compat"
		refreshinterval = 80
	}
`

func TestAccAuthenticationpushservice_clientsecret_backward_compat(t *testing.T) {
	t.Setenv("TF_VAR_authenticationpushservice_clientsecret", "secret_value1")
	t.Setenv("TF_VAR_authenticationpushservice_clientsecret_2", "secret_value2")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationpushserviceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationpushservice_clientsecret_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationpushserviceExist("citrixadc_authenticationpushservice.test", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationpushservice.test", "name", "tf_pushservice_compat"),
					resource.TestCheckResourceAttr("citrixadc_authenticationpushservice.test", "clientid", "cliId_compat"),
				),
			},
			{
				Config: testAccAuthenticationpushservice_clientsecret_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationpushserviceExist("citrixadc_authenticationpushservice.test", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationpushservice.test", "name", "tf_pushservice_compat"),
					resource.TestCheckResourceAttr("citrixadc_authenticationpushservice.test", "clientid", "cliId_compat"),
				),
			},
		},
	})
}

// Test ephemeral path: using clientsecret_wo (WriteOnly attribute) with version tracker
const testAccAuthenticationpushservice_clientsecret_wo_step1 = `

	variable "authenticationpushservice_clientsecret_wo" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_authenticationpushservice" "test" {
		name                    = "tf_pushservice_wo"
		clientid                = "cliId_wo"
		clientsecret_wo         = var.authenticationpushservice_clientsecret_wo
		clientsecret_wo_version = 1
		customerid              = "cusID_wo"
		refreshinterval         = 50
	}
`

// Update ephemeral path: bump version to trigger update with new secret
const testAccAuthenticationpushservice_clientsecret_wo_step2 = `

	variable "authenticationpushservice_clientsecret_wo_2" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_authenticationpushservice" "test" {
		name                    = "tf_pushservice_wo"
		clientid                = "cliId_wo"
		clientsecret_wo         = var.authenticationpushservice_clientsecret_wo_2
		clientsecret_wo_version = 2
		customerid              = "cusID_wo"
		refreshinterval         = 80
	}
`

func TestAccAuthenticationpushservice_clientsecret_wo_ephemeral(t *testing.T) {
	t.Setenv("TF_VAR_authenticationpushservice_clientsecret_wo", "ephemeral_secret1")
	t.Setenv("TF_VAR_authenticationpushservice_clientsecret_wo_2", "ephemeral_secret2")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationpushserviceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationpushservice_clientsecret_wo_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationpushserviceExist("citrixadc_authenticationpushservice.test", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationpushservice.test", "name", "tf_pushservice_wo"),
					resource.TestCheckResourceAttr("citrixadc_authenticationpushservice.test", "clientid", "cliId_wo"),
					resource.TestCheckResourceAttr("citrixadc_authenticationpushservice.test", "clientsecret_wo_version", "1"),
				),
			},
			{
				Config: testAccAuthenticationpushservice_clientsecret_wo_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationpushserviceExist("citrixadc_authenticationpushservice.test", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationpushservice.test", "name", "tf_pushservice_wo"),
					resource.TestCheckResourceAttr("citrixadc_authenticationpushservice.test", "clientid", "cliId_wo"),
					resource.TestCheckResourceAttr("citrixadc_authenticationpushservice.test", "clientsecret_wo_version", "2"),
				),
			},
		},
	})
}

func TestAccAuthenticationpushservice_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckAuthenticationpushserviceDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccAuthenticationpushservice_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationpushserviceExist("citrixadc_authenticationpushservice.tf_pushservice", nil),
				),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccAuthenticationpushservice_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationpushserviceExist("citrixadc_authenticationpushservice.tf_pushservice", nil),
				),
			},
		},
	})
}

// Step 1: refreshinterval set to a non-default value (default is 50).
const testAccAuthenticationpushservice_unset_step1 = `
	resource "citrixadc_authenticationpushservice" "tf_unset" {
		name            = "tf_test_authpushsvc_unset"
		clientid        = "cliId"
		clientsecret    = "secret"
		customerid      = "cusID"
		refreshinterval = 80
	}
`

// Step 2: refreshinterval removed from config -> provider must unset it so the
// appliance reverts it to its default (50).
const testAccAuthenticationpushservice_unset_step2 = `
	resource "citrixadc_authenticationpushservice" "tf_unset" {
		name            = "tf_test_authpushsvc_unset"
		clientid        = "cliId"
		clientsecret    = "secret"
		customerid      = "cusID"
	}
`

func TestAccAuthenticationpushservice_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationpushserviceDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default value applies and persists.
				Config: testAccAuthenticationpushservice_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationpushserviceExist("citrixadc_authenticationpushservice.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationpushservice.tf_unset", "refreshinterval", "80"),
				),
			},
			{
				// Removing it must unset -> state reverts to NITRO default (50),
				// and the implicit post-apply plan must be empty.
				Config: testAccAuthenticationpushservice_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationpushserviceExist("citrixadc_authenticationpushservice.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationpushservice.tf_unset", "refreshinterval", "50"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckAuthenticationpushserviceADCValue("tf_test_authpushsvc_unset", "refreshinterval", "50"),
				),
			},
		},
	})
}

// testAccCheckAuthenticationpushserviceADCValue asserts an attribute's value
// directly on the appliance (not just in Terraform state), proving the unset
// actually reverted it.
func testAccCheckAuthenticationpushserviceADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource("authenticationpushservice", name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("authenticationpushservice %s not found on appliance", name)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("authenticationpushservice %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}
