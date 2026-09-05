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

const testAccMcpprofile_basic = `


resource "citrixadc_mcpprofile" "tf_mcpprofile" {
	name      = "my_mcpprofile"
	proxymode = "FORWARD"
	comment   = "acctest mcp profile"
}

`
const testAccMcpprofile_update = `


resource "citrixadc_mcpprofile" "tf_mcpprofile" {
	name      = "my_mcpprofile"
	proxymode = "REVERSE"
	comment   = "updated comment"
}

`

func TestAccMcpprofile_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckMcpprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccMcpprofile_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMcpprofileExist("citrixadc_mcpprofile.tf_mcpprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_mcpprofile.tf_mcpprofile", "name", "my_mcpprofile"),
					resource.TestCheckResourceAttr("citrixadc_mcpprofile.tf_mcpprofile", "proxymode", "FORWARD"),
					resource.TestCheckResourceAttr("citrixadc_mcpprofile.tf_mcpprofile", "comment", "acctest mcp profile"),
				),
			},
			{
				Config: testAccMcpprofile_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMcpprofileExist("citrixadc_mcpprofile.tf_mcpprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_mcpprofile.tf_mcpprofile", "name", "my_mcpprofile"),
					resource.TestCheckResourceAttr("citrixadc_mcpprofile.tf_mcpprofile", "proxymode", "REVERSE"),
					resource.TestCheckResourceAttr("citrixadc_mcpprofile.tf_mcpprofile", "comment", "updated comment"),
				),
			},
		},
	})
}

func TestAccMcpprofile_import(t *testing.T) {
	const resAddr = "citrixadc_mcpprofile.tf_mcpprofile"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckMcpprofileDestroy,
		Steps: []resource.TestStep{
			{Config: testAccMcpprofile_basic},
			{
				Config:            testAccMcpprofile_basic,
				ResourceName:      resAddr,
				ImportState:       true,
				ImportStateVerify: true,
				// tokenorapi is a secret not round-tripped through GET, and the
				// tokenorapi_wo/_wo_version write-only trackers cannot be imported.
				ImportStateVerifyIgnore: []string{"tokenorapi", "tokenorapi_wo", "tokenorapi_wo_version"},
			},
		},
	})
}

func testAccCheckMcpprofileExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No mcpprofile name is set")
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
		data, err := client.FindResource("mcpprofile", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("mcpprofile %s not found", n)
		}

		return nil
	}
}

func testAccCheckMcpprofileDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_mcpprofile" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("mcpprofile", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("mcpprofile %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccMcpprofileDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccMcpprofileDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_mcpprofile.test", "name", "tf_mcpprofile_ds"),
					resource.TestCheckResourceAttrSet("data.citrixadc_mcpprofile.test", "id"),
					resource.TestCheckResourceAttrSet("data.citrixadc_mcpprofile.test", "proxymode"),
					resource.TestCheckResourceAttrSet("data.citrixadc_mcpprofile.test", "profiletype"),
				),
			},
		},
	})
}

const testAccMcpprofileDataSource_basic = `
resource "citrixadc_mcpprofile" "tf_mcpprofile_ds" {
	name      = "tf_mcpprofile_ds"
	proxymode = "FORWARD"
	comment   = "datasource test"
}

data "citrixadc_mcpprofile" "test" {
	name = citrixadc_mcpprofile.tf_mcpprofile_ds.name
}
`

// Step 1: every unset-eligible attribute set to a valid non-default value.
// proxymode=REVERSE forces host/url replacement DISABLED on the appliance, so
// those two are intentionally left unset here to respect that constraint.
// tokenorapi is a secret (password_key) and is exercised via the dedicated
// write-only ephemeral test, not the unset flow.
const testAccMcpprofile_unset_step1 = `
resource "citrixadc_mcpprofile" "tf_unset" {
	name                        = "tf_test_mcpprofile_unset"
	proxymode                   = "REVERSE"
	protocolversion             = "2025-03-26"
	comment                     = "unset test comment"
	insertheaderinclientrequest = "ENABLED"
}
`

// Step 2: all eligible attributes removed from config -> provider must unset them,
// reverting each to its NITRO default.
const testAccMcpprofile_unset_step2 = `
resource "citrixadc_mcpprofile" "tf_unset" {
	name = "tf_test_mcpprofile_unset"
}
`

func TestAccMcpprofile_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckMcpprofileDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values apply and persist.
				Config: testAccMcpprofile_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMcpprofileExist("citrixadc_mcpprofile.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_mcpprofile.tf_unset", "proxymode", "REVERSE"),
					resource.TestCheckResourceAttr("citrixadc_mcpprofile.tf_unset", "comment", "unset test comment"),
					resource.TestCheckResourceAttr("citrixadc_mcpprofile.tf_unset", "insertheaderinclientrequest", "ENABLED"),
				),
			},
			{
				// Removing them must unset -> state reverts to NITRO defaults.
				Config: testAccMcpprofile_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMcpprofileExist("citrixadc_mcpprofile.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_mcpprofile.tf_unset", "proxymode", "FORWARD"),
					resource.TestCheckResourceAttr("citrixadc_mcpprofile.tf_unset", "insertheaderinclientrequest", "DISABLED"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckMcpprofileADCValue("tf_test_mcpprofile_unset", "proxymode", "FORWARD"),
					testAccCheckMcpprofileADCValue("tf_test_mcpprofile_unset", "insertheaderinclientrequest", "DISABLED"),
				),
			},
		},
	})
}

// testAccCheckMcpprofileADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted it.
func testAccCheckMcpprofileADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Mcpprofile.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("mcpprofile %s not found on appliance", name)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("mcpprofile %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

// TestAccMcpprofile_selfHealing verifies the provider re-creates the profile when
// it is deleted out-of-band between apply steps (drift recovery).
func TestAccMcpprofile_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_mcpprofile.tf_mcpprofile"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckMcpprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccMcpprofile_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckMcpprofileExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Mcpprofile.Type(), "my_mcpprofile"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccMcpprofile_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckMcpprofileExist(resAddr, nil)),
			},
		},
	})
}

// tokenorapi is a NITRO password_key (secret). It supports two configuration
// paths: the backward-compatible plaintext `tokenorapi` attribute and the
// write-only ephemeral `tokenorapi_wo` (+ `tokenorapi_wo_version`) pair. The two
// tests below exercise each path, including secret rotation.

// Backward-compatible path: plaintext tokenorapi (Sensitive), rotated in step 2.
const testAccMcpprofile_tokenorapi_step1 = `
	variable "mcp_tokenorapi" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_mcpprofile" "tf_mcp_secret" {
		name       = "tf_mcp_secret"
		proxymode  = "FORWARD"
		tokenorapi = var.mcp_tokenorapi
	}
`

const testAccMcpprofile_tokenorapi_step2 = `
	variable "mcp_tokenorapi_2" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_mcpprofile" "tf_mcp_secret" {
		name       = "tf_mcp_secret"
		proxymode  = "FORWARD"
		tokenorapi = var.mcp_tokenorapi_2
	}
`

func TestAccMcpprofile_tokenorapi_backward_compat(t *testing.T) {
	t.Setenv("TF_VAR_mcp_tokenorapi", "Authorization: Bearer token1")
	t.Setenv("TF_VAR_mcp_tokenorapi_2", "Authorization: Bearer token2")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckMcpprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccMcpprofile_tokenorapi_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMcpprofileExist("citrixadc_mcpprofile.tf_mcp_secret", nil),
					resource.TestCheckResourceAttr("citrixadc_mcpprofile.tf_mcp_secret", "name", "tf_mcp_secret"),
				),
			},
			{
				Config: testAccMcpprofile_tokenorapi_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMcpprofileExist("citrixadc_mcpprofile.tf_mcp_secret", nil),
					resource.TestCheckResourceAttr("citrixadc_mcpprofile.tf_mcp_secret", "name", "tf_mcp_secret"),
				),
			},
		},
	})
}

// Ephemeral path: write-only tokenorapi_wo + tokenorapi_wo_version, rotated in
// step 2 by bumping the version.
const testAccMcpprofile_tokenorapi_wo_step1 = `
	variable "mcp_tokenorapi_wo" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_mcpprofile" "tf_mcp_secret" {
		name                  = "tf_mcp_secret"
		proxymode             = "FORWARD"
		tokenorapi_wo         = var.mcp_tokenorapi_wo
		tokenorapi_wo_version = 1
	}
`

const testAccMcpprofile_tokenorapi_wo_step2 = `
	variable "mcp_tokenorapi_wo_2" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_mcpprofile" "tf_mcp_secret" {
		name                  = "tf_mcp_secret"
		proxymode             = "FORWARD"
		tokenorapi_wo         = var.mcp_tokenorapi_wo_2
		tokenorapi_wo_version = 2
	}
`

func TestAccMcpprofile_tokenorapi_wo_ephemeral(t *testing.T) {
	t.Setenv("TF_VAR_mcp_tokenorapi_wo", "Authorization: Bearer ephem1")
	t.Setenv("TF_VAR_mcp_tokenorapi_wo_2", "Authorization: Bearer ephem2")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckMcpprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccMcpprofile_tokenorapi_wo_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMcpprofileExist("citrixadc_mcpprofile.tf_mcp_secret", nil),
					resource.TestCheckResourceAttr("citrixadc_mcpprofile.tf_mcp_secret", "name", "tf_mcp_secret"),
					resource.TestCheckResourceAttr("citrixadc_mcpprofile.tf_mcp_secret", "tokenorapi_wo_version", "1"),
				),
			},
			{
				Config: testAccMcpprofile_tokenorapi_wo_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMcpprofileExist("citrixadc_mcpprofile.tf_mcp_secret", nil),
					resource.TestCheckResourceAttr("citrixadc_mcpprofile.tf_mcp_secret", "name", "tf_mcp_secret"),
					resource.TestCheckResourceAttr("citrixadc_mcpprofile.tf_mcp_secret", "tokenorapi_wo_version", "2"),
				),
			},
		},
	})
}
