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

const testAccAaatacacsparams_basic = `


resource "citrixadc_aaatacacsparams" "tf_aaatacacsparams" {
	serverip      = "10.222.74.158"
	serverport    = 49
	authtimeout   = 5
	authorization = "OFF"
	}
  
`
const testAccAaatacacsparams_update = `


resource "citrixadc_aaatacacsparams" "tf_aaatacacsparams" {
	serverip      = "10.222.74.159"
	serverport    = 50
	authtimeout   = 6
	authorization = "ON"
	}
  
`

func TestAccAaatacacsparams_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccAaatacacsparams_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaatacacsparamsExist("citrixadc_aaatacacsparams.tf_aaatacacsparams", nil),
					resource.TestCheckResourceAttr("citrixadc_aaatacacsparams.tf_aaatacacsparams", "serverip", "10.222.74.158"),
					resource.TestCheckResourceAttr("citrixadc_aaatacacsparams.tf_aaatacacsparams", "serverport", "49"),
					resource.TestCheckResourceAttr("citrixadc_aaatacacsparams.tf_aaatacacsparams", "authtimeout", "5"),
					resource.TestCheckResourceAttr("citrixadc_aaatacacsparams.tf_aaatacacsparams", "authorization", "OFF"),
				),
			},
			{
				Config: testAccAaatacacsparams_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaatacacsparamsExist("citrixadc_aaatacacsparams.tf_aaatacacsparams", nil),
					resource.TestCheckResourceAttr("citrixadc_aaatacacsparams.tf_aaatacacsparams", "serverip", "10.222.74.159"),
					resource.TestCheckResourceAttr("citrixadc_aaatacacsparams.tf_aaatacacsparams", "serverport", "50"),
					resource.TestCheckResourceAttr("citrixadc_aaatacacsparams.tf_aaatacacsparams", "authtimeout", "6"),
					resource.TestCheckResourceAttr("citrixadc_aaatacacsparams.tf_aaatacacsparams", "authorization", "ON"),
				),
			},
		},
	})
}

func testAccCheckAaatacacsparamsExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No aaatacacsparams name is set")
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
		data, err := client.FindResource(service.Aaatacacsparams.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("aaatacacsparams %s not found", n)
		}

		return nil
	}
}

const testAccAaatacacsparamsDataSource_basic = `


resource "citrixadc_aaatacacsparams" "tf_aaatacacsparams" {
	serverip      = "10.222.74.158"
	serverport    = 49
	authtimeout   = 5
	authorization = "OFF"
	}

	data "citrixadc_aaatacacsparams" "tf_aaatacacsparams" {
		depends_on = [citrixadc_aaatacacsparams.tf_aaatacacsparams]
	}
  
`

func TestAccAaatacacsparamsDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccAaatacacsparamsDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_aaatacacsparams.tf_aaatacacsparams", "serverip", "10.222.74.158"),
					resource.TestCheckResourceAttr("data.citrixadc_aaatacacsparams.tf_aaatacacsparams", "serverport", "49"),
					resource.TestCheckResourceAttr("data.citrixadc_aaatacacsparams.tf_aaatacacsparams", "authtimeout", "5"),
					resource.TestCheckResourceAttr("data.citrixadc_aaatacacsparams.tf_aaatacacsparams", "authorization", "OFF"),
				),
			},
		},
	})
}

// Test backward-compatible path: using tacacssecret (Sensitive attribute)
const testAccAaatacacsparams_tacacssecret_step1 = `

	variable "aaatacacsparams_tacacssecret" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_aaatacacsparams" "tf_aaatacacsparams" {
		serverip      = "10.222.74.158"
		serverport    = 49
		authtimeout   = 5
		authorization = "OFF"
		tacacssecret  = var.aaatacacsparams_tacacssecret
	}
`

// Update backward-compatible path: change tacacssecret value
const testAccAaatacacsparams_tacacssecret_step2 = `

	variable "aaatacacsparams_tacacssecret_2" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_aaatacacsparams" "tf_aaatacacsparams" {
		serverip      = "10.222.74.158"
		serverport    = 49
		authtimeout   = 5
		authorization = "OFF"
		tacacssecret  = var.aaatacacsparams_tacacssecret_2
	}
`

func TestAccAaatacacsparams_tacacssecret_backward_compat(t *testing.T) {
	t.Setenv("TF_VAR_aaatacacsparams_tacacssecret", "value1")
	t.Setenv("TF_VAR_aaatacacsparams_tacacssecret_2", "value2")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccAaatacacsparams_tacacssecret_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaatacacsparamsExist("citrixadc_aaatacacsparams.tf_aaatacacsparams", nil),
					resource.TestCheckResourceAttr("citrixadc_aaatacacsparams.tf_aaatacacsparams", "serverip", "10.222.74.158"),
					resource.TestCheckResourceAttr("citrixadc_aaatacacsparams.tf_aaatacacsparams", "serverport", "49"),
					resource.TestCheckResourceAttr("citrixadc_aaatacacsparams.tf_aaatacacsparams", "authtimeout", "5"),
					resource.TestCheckResourceAttr("citrixadc_aaatacacsparams.tf_aaatacacsparams", "authorization", "OFF"),
				),
			},
			{
				Config: testAccAaatacacsparams_tacacssecret_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaatacacsparamsExist("citrixadc_aaatacacsparams.tf_aaatacacsparams", nil),
					resource.TestCheckResourceAttr("citrixadc_aaatacacsparams.tf_aaatacacsparams", "serverip", "10.222.74.158"),
					resource.TestCheckResourceAttr("citrixadc_aaatacacsparams.tf_aaatacacsparams", "serverport", "49"),
					resource.TestCheckResourceAttr("citrixadc_aaatacacsparams.tf_aaatacacsparams", "authtimeout", "5"),
					resource.TestCheckResourceAttr("citrixadc_aaatacacsparams.tf_aaatacacsparams", "authorization", "OFF"),
				),
			},
		},
	})
}

// Test ephemeral path: using tacacssecret_wo (WriteOnly attribute) with version tracker
const testAccAaatacacsparams_tacacssecret_wo_step1 = `

	variable "aaatacacsparams_tacacssecret_wo" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_aaatacacsparams" "tf_aaatacacsparams" {
		serverip                 = "10.222.74.158"
		serverport               = 49
		authtimeout              = 5
		authorization            = "OFF"
		tacacssecret_wo          = var.aaatacacsparams_tacacssecret_wo
		tacacssecret_wo_version  = 1
	}
`

// Update ephemeral path: bump version to trigger update with new secret
const testAccAaatacacsparams_tacacssecret_wo_step2 = `

	variable "aaatacacsparams_tacacssecret_wo_2" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_aaatacacsparams" "tf_aaatacacsparams" {
		serverip                 = "10.222.74.158"
		serverport               = 49
		authtimeout              = 5
		authorization            = "OFF"
		tacacssecret_wo          = var.aaatacacsparams_tacacssecret_wo_2
		tacacssecret_wo_version  = 2
	}
`

func TestAccAaatacacsparams_tacacssecret_wo_ephemeral(t *testing.T) {
	t.Setenv("TF_VAR_aaatacacsparams_tacacssecret_wo", "ephemeral_value1")
	t.Setenv("TF_VAR_aaatacacsparams_tacacssecret_wo_2", "ephemeral_value2")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccAaatacacsparams_tacacssecret_wo_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaatacacsparamsExist("citrixadc_aaatacacsparams.tf_aaatacacsparams", nil),
					resource.TestCheckResourceAttr("citrixadc_aaatacacsparams.tf_aaatacacsparams", "serverip", "10.222.74.158"),
					resource.TestCheckResourceAttr("citrixadc_aaatacacsparams.tf_aaatacacsparams", "serverport", "49"),
					resource.TestCheckResourceAttr("citrixadc_aaatacacsparams.tf_aaatacacsparams", "authtimeout", "5"),
					resource.TestCheckResourceAttr("citrixadc_aaatacacsparams.tf_aaatacacsparams", "authorization", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_aaatacacsparams.tf_aaatacacsparams", "tacacssecret_wo_version", "1"),
				),
			},
			{
				Config: testAccAaatacacsparams_tacacssecret_wo_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaatacacsparamsExist("citrixadc_aaatacacsparams.tf_aaatacacsparams", nil),
					resource.TestCheckResourceAttr("citrixadc_aaatacacsparams.tf_aaatacacsparams", "serverip", "10.222.74.158"),
					resource.TestCheckResourceAttr("citrixadc_aaatacacsparams.tf_aaatacacsparams", "serverport", "49"),
					resource.TestCheckResourceAttr("citrixadc_aaatacacsparams.tf_aaatacacsparams", "authtimeout", "5"),
					resource.TestCheckResourceAttr("citrixadc_aaatacacsparams.tf_aaatacacsparams", "authorization", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_aaatacacsparams.tf_aaatacacsparams", "tacacssecret_wo_version", "2"),
				),
			},
		},
	})
}

func TestAccAaatacacsparams_import(t *testing.T) {
	const resAddr = "citrixadc_aaatacacsparams.tf_aaatacacsparams"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccAaatacacsparams_basic,
			},
			{
				Config:            testAccAaatacacsparams_basic,
				ResourceName:      resAddr,
				ImportState:       true,
				ImportStateVerify: true,
				// tacacssecret_wo_version is a client-side write-only version tracker
				// that NITRO does not store or return, so it cannot round-trip on import.
				ImportStateVerifyIgnore: []string{"tacacssecret_wo_version"},
			},
		},
	})
}

// Unset test: step1 sets the unset-eligible attributes (serverport, authtimeout)
// to non-default values; step2 removes them from config so the provider must issue
// ?action=unset and the appliance reverts each to its NITRO default (serverport=49,
// authtimeout=3). The implicit post-apply plan on step2 must be empty (no perpetual diff).
const testAccAaatacacsparams_unset_step1 = `
resource "citrixadc_aaatacacsparams" "tf_unset" {
	serverip    = "10.222.74.200"
	serverport  = 88
	authtimeout = 10
}
`

const testAccAaatacacsparams_unset_step2 = `
resource "citrixadc_aaatacacsparams" "tf_unset" {
	serverip = "10.222.74.200"
	# serverport and authtimeout removed from config -> provider must unset them
}
`

func TestAccAaatacacsparams_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				// Non-default values apply and persist.
				Config: testAccAaatacacsparams_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaatacacsparamsExist("citrixadc_aaatacacsparams.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_aaatacacsparams.tf_unset", "serverport", "88"),
					resource.TestCheckResourceAttr("citrixadc_aaatacacsparams.tf_unset", "authtimeout", "10"),
				),
			},
			{
				// Removing them must unset -> state reverts to NITRO defaults,
				// and the implicit post-apply plan must be empty.
				Config: testAccAaatacacsparams_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaatacacsparamsExist("citrixadc_aaatacacsparams.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_aaatacacsparams.tf_unset", "serverport", "49"),
					resource.TestCheckResourceAttr("citrixadc_aaatacacsparams.tf_unset", "authtimeout", "3"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckAaatacacsparamsADCValue("serverport", "49"),
					testAccCheckAaatacacsparamsADCValue("authtimeout", "3"),
				),
			},
		},
	})
}

// testAccCheckAaatacacsparamsADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted it.
// aaatacacsparams is a singleton (params) resource, so it is fetched with an empty name.
func testAccCheckAaatacacsparamsADCValue(attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Aaatacacsparams.Type(), "")
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("aaatacacsparams not found on appliance")
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("aaatacacsparams: appliance attr %q = %q, want %q (unset did not revert it)", attr, got, want)
		}
		return nil
	}
}

func TestAccAaatacacsparams_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccAaatacacsparams_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaatacacsparamsExist("citrixadc_aaatacacsparams.tf_aaatacacsparams", nil),
				),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccAaatacacsparams_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaatacacsparamsExist("citrixadc_aaatacacsparams.tf_aaatacacsparams", nil),
				),
			},
		},
	})
}
