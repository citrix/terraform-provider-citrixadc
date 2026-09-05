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

const testAccHasecureheartbeats_basic = `

	resource "citrixadc_hasecureheartbeats" "tf_hasecureheartbeats" {
		state = "ENABLED"
		hapsk = "presharedkey123"
	}

`

// The update keeps state=ENABLED and rotates hapsk. NITRO rejects specifying
// hapsk together with state=DISABLED (errorcode 1092 "Arguments cannot both be
// specified [haPSK, state==DISABLED]"), so the two cannot be changed in one step.
const testAccHasecureheartbeats_update = `

	resource "citrixadc_hasecureheartbeats" "tf_hasecureheartbeats" {
		state = "ENABLED"
		hapsk = "presharedkey456"
	}

`

func TestAccHasecureheartbeats_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckHasecureheartbeatsDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccHasecureheartbeats_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHasecureheartbeatsExist("citrixadc_hasecureheartbeats.tf_hasecureheartbeats", nil),
					resource.TestCheckResourceAttr("citrixadc_hasecureheartbeats.tf_hasecureheartbeats", "state", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_hasecureheartbeats.tf_hasecureheartbeats", "hapsk", "presharedkey123"),
					// Independent appliance-level confirmation.
					testAccCheckHasecureheartbeatsADCValue("state", "ENABLED"),
				),
			},
			{
				Config: testAccHasecureheartbeats_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHasecureheartbeatsExist("citrixadc_hasecureheartbeats.tf_hasecureheartbeats", nil),
					resource.TestCheckResourceAttr("citrixadc_hasecureheartbeats.tf_hasecureheartbeats", "state", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_hasecureheartbeats.tf_hasecureheartbeats", "hapsk", "presharedkey456"),
					testAccCheckHasecureheartbeatsADCValue("state", "ENABLED"),
				),
			},
		},
	})
}

func testAccCheckHasecureheartbeatsExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No hasecureheartbeats id is set")
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
		data, err := client.FindResource(service.Hasecureheartbeats.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("hasecureheartbeats %s not found", n)
		}

		return nil
	}
}

// hasecureheartbeats is a global configuration singleton with no NITRO delete
// operation; there is nothing to assert on destroy.
func testAccCheckHasecureheartbeatsDestroy(s *terraform.State) error {
	return nil
}

// testAccCheckHasecureheartbeatsADCValue asserts an attribute's value directly on
// the appliance (not just in Terraform state). A missing key is treated as an
// empty value, which is how the appliance reports an unset attribute.
func testAccCheckHasecureheartbeatsADCValue(attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Hasecureheartbeats.Type(), "")
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("hasecureheartbeats not found on appliance")
		}
		got := ""
		if v, ok := data[attr]; ok && v != nil {
			got = strings.TrimSpace(fmt.Sprintf("%v", v))
		}
		if got != want {
			return fmt.Errorf("hasecureheartbeats: appliance attr %q = %q, want %q", attr, got, want)
		}
		return nil
	}
}

const testAccHasecureheartbeatsDataSource_basic = `

	resource "citrixadc_hasecureheartbeats" "tf_hasecureheartbeats" {
		state = "ENABLED"
		hapsk = "presharedkey123"
	}

	data "citrixadc_hasecureheartbeats" "tf_hasecureheartbeats" {
		depends_on = [citrixadc_hasecureheartbeats.tf_hasecureheartbeats]
	}
`

func TestAccHasecureheartbeatsDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckHasecureheartbeatsDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccHasecureheartbeatsDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_hasecureheartbeats.tf_hasecureheartbeats", "state", "ENABLED"),
				),
			},
		},
	})
}

func TestAccHasecureheartbeats_import(t *testing.T) {
	const resAddr = "citrixadc_hasecureheartbeats.tf_hasecureheartbeats"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckHasecureheartbeatsDestroy,
		Steps: []resource.TestStep{
			{Config: testAccHasecureheartbeats_basic},
			{
				Config:            testAccHasecureheartbeats_basic,
				ResourceName:      resAddr,
				ImportState:       true,
				ImportStateVerify: true,
				// hapsk is a secret retained from config (NITRO returns only an
				// encrypted form or omits it); hapsk_wo_version is a config-only
				// tracker not returned by GET. Neither round-trips on import.
				ImportStateVerifyIgnore: []string{"hapsk", "hapsk_wo_version"},
			},
		},
	})
}

// TestAccHasecureheartbeats_hapsk_backward_compat exercises the legacy plain (non
// write-only) hapsk attribute: it is supplied via a sensitive variable and
// changed between steps, mirroring aaaproxyparam's proxypassword backward-compat
// test.
const testAccHasecureheartbeats_hapsk_step1 = `
	variable "hasecureheartbeats_hapsk" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_hasecureheartbeats" "tf_hasecureheartbeats" {
		state = "ENABLED"
		hapsk = var.hasecureheartbeats_hapsk
	}
`
const testAccHasecureheartbeats_hapsk_step2 = `
	variable "hasecureheartbeats_hapsk_2" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_hasecureheartbeats" "tf_hasecureheartbeats" {
		state = "ENABLED"
		hapsk = var.hasecureheartbeats_hapsk_2
	}
`

func TestAccHasecureheartbeats_hapsk_backward_compat(t *testing.T) {
	t.Setenv("TF_VAR_hasecureheartbeats_hapsk", "presharedkey123")
	t.Setenv("TF_VAR_hasecureheartbeats_hapsk_2", "presharedkey456")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckHasecureheartbeatsDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccHasecureheartbeats_hapsk_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHasecureheartbeatsExist("citrixadc_hasecureheartbeats.tf_hasecureheartbeats", nil),
					resource.TestCheckResourceAttr("citrixadc_hasecureheartbeats.tf_hasecureheartbeats", "state", "ENABLED"),
				),
			},
			{
				Config: testAccHasecureheartbeats_hapsk_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHasecureheartbeatsExist("citrixadc_hasecureheartbeats.tf_hasecureheartbeats", nil),
					resource.TestCheckResourceAttr("citrixadc_hasecureheartbeats.tf_hasecureheartbeats", "state", "ENABLED"),
				),
			},
		},
	})
}

// TestAccHasecureheartbeats_hapsk_wo_ephemeral exercises the write-only
// (ephemeral) hapsk_wo secret + hapsk_wo_version tracker, mirroring
// aaaproxyparam's proxypassword_wo ephemeral test. The secret is supplied via a
// sensitive variable and rotated by bumping the version; it must never be
// persisted to state.
const testAccHasecureheartbeats_hapsk_wo_step1 = `
	variable "hasecureheartbeats_hapsk_wo" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_hasecureheartbeats" "tf_hasecureheartbeats" {
		state            = "ENABLED"
		hapsk_wo         = var.hasecureheartbeats_hapsk_wo
		hapsk_wo_version = 1
	}
`
const testAccHasecureheartbeats_hapsk_wo_step2 = `
	variable "hasecureheartbeats_hapsk_wo_2" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_hasecureheartbeats" "tf_hasecureheartbeats" {
		state            = "ENABLED"
		hapsk_wo         = var.hasecureheartbeats_hapsk_wo_2
		hapsk_wo_version = 2
	}
`

func TestAccHasecureheartbeats_hapsk_wo_ephemeral(t *testing.T) {
	t.Setenv("TF_VAR_hasecureheartbeats_hapsk_wo", "ephemeral_key1")
	t.Setenv("TF_VAR_hasecureheartbeats_hapsk_wo_2", "ephemeral_key2")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckHasecureheartbeatsDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccHasecureheartbeats_hapsk_wo_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHasecureheartbeatsExist("citrixadc_hasecureheartbeats.tf_hasecureheartbeats", nil),
					resource.TestCheckResourceAttr("citrixadc_hasecureheartbeats.tf_hasecureheartbeats", "hapsk_wo_version", "1"),
					resource.TestCheckResourceAttr("citrixadc_hasecureheartbeats.tf_hasecureheartbeats", "state", "ENABLED"),
					// The write-only secret and the plain secret must never be
					// persisted to Terraform state.
					resource.TestCheckNoResourceAttr("citrixadc_hasecureheartbeats.tf_hasecureheartbeats", "hapsk_wo"),
					resource.TestCheckNoResourceAttr("citrixadc_hasecureheartbeats.tf_hasecureheartbeats", "hapsk"),
				),
			},
			{
				Config: testAccHasecureheartbeats_hapsk_wo_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHasecureheartbeatsExist("citrixadc_hasecureheartbeats.tf_hasecureheartbeats", nil),
					resource.TestCheckResourceAttr("citrixadc_hasecureheartbeats.tf_hasecureheartbeats", "hapsk_wo_version", "2"),
					resource.TestCheckResourceAttr("citrixadc_hasecureheartbeats.tf_hasecureheartbeats", "state", "ENABLED"),
					resource.TestCheckNoResourceAttr("citrixadc_hasecureheartbeats.tf_hasecureheartbeats", "hapsk_wo"),
					resource.TestCheckNoResourceAttr("citrixadc_hasecureheartbeats.tf_hasecureheartbeats", "hapsk"),
				),
			},
		},
	})
}
