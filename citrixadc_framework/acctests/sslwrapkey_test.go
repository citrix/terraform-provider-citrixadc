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

// ============================================================================
// FIPS / CRYPTO SUBSYSTEM REQUIRED -- TESTS ARE SKIP-GATED
// ============================================================================
// sslwrapkey creates a key-wrapping key via the NITRO `create` action and
// deletes it via DELETE. Wrap-key creation exercises the FIPS/crypto subsystem
// and is likely unsupported on a non-FIPS VPX appliance (the create action may
// fail with "operation not supported on this platform" / FIPS-related errors).
//
// password and salt are write-only secret triples. The tests below pass them
// via the _wo path and TF_VAR_* environment variables.
//
// Every test in this file is t.Skip-gated. To run on a real FIPS/crypto-capable
// appliance, remove the t.Skip line, supply real secret values via TF_VAR_*,
// and replace any TODO_PLACEHOLDER values.
// ============================================================================

package citrixadc

import (
	"fmt"
	"testing"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// Basic test. All attributes are RequiresReplace, so "step2" changes the key
// name (forces recreate). password/salt are supplied via the _wo path.
const testAccSslwrapkey_basic_step1 = `
variable "sslwrapkey_password_wo" {
  type      = string
  sensitive = true
}
variable "sslwrapkey_salt_wo" {
  type      = string
  sensitive = true
}

resource "citrixadc_sslwrapkey" "tf_sslwrapkey" {
  wrapkeyname         = "tf_wrapkey"
  password_wo         = var.sslwrapkey_password_wo
  password_wo_version = 1
  salt_wo             = var.sslwrapkey_salt_wo
  salt_wo_version     = 1
}

`

const testAccSslwrapkey_basic_step2 = `
variable "sslwrapkey_password_wo_2" {
  type      = string
  sensitive = true
}
variable "sslwrapkey_salt_wo_2" {
  type      = string
  sensitive = true
}

resource "citrixadc_sslwrapkey" "tf_sslwrapkey" {
  wrapkeyname         = "tf_wrapkey_updated"
  password_wo         = var.sslwrapkey_password_wo_2
  password_wo_version = 2
  salt_wo             = var.sslwrapkey_salt_wo_2
  salt_wo_version     = 2
}

`

func TestAccSslwrapkey_basic(t *testing.T) {
	// FIPS/crypto subsystem required: wrap-key creation is likely unsupported on
	// the non-FIPS VPX testbed.
	t.Skip("sslwrapkey creation needs the FIPS/crypto subsystem and is likely unsupported on the non-FIPS VPX testbed.")

	// Replace these with real secret values before running on a capable appliance.
	t.Setenv("TF_VAR_sslwrapkey_password_wo", "TODO_PLACEHOLDER")
	t.Setenv("TF_VAR_sslwrapkey_salt_wo", "TODO_PLACEHOLDER")
	t.Setenv("TF_VAR_sslwrapkey_password_wo_2", "TODO_PLACEHOLDER")
	t.Setenv("TF_VAR_sslwrapkey_salt_wo_2", "TODO_PLACEHOLDER")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslwrapkeyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslwrapkey_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslwrapkeyExist("citrixadc_sslwrapkey.tf_sslwrapkey", nil),
					resource.TestCheckResourceAttr("citrixadc_sslwrapkey.tf_sslwrapkey", "wrapkeyname", "tf_wrapkey"),
					resource.TestCheckResourceAttr("citrixadc_sslwrapkey.tf_sslwrapkey", "password_wo_version", "1"),
					resource.TestCheckResourceAttr("citrixadc_sslwrapkey.tf_sslwrapkey", "salt_wo_version", "1"),
				),
			},
			{
				Config: testAccSslwrapkey_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslwrapkeyExist("citrixadc_sslwrapkey.tf_sslwrapkey", nil),
					resource.TestCheckResourceAttr("citrixadc_sslwrapkey.tf_sslwrapkey", "wrapkeyname", "tf_wrapkey_updated"),
					resource.TestCheckResourceAttr("citrixadc_sslwrapkey.tf_sslwrapkey", "password_wo_version", "2"),
					resource.TestCheckResourceAttr("citrixadc_sslwrapkey.tf_sslwrapkey", "salt_wo_version", "2"),
				),
			},
		},
	})
}

func testAccCheckSslwrapkeyExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No sslwrapkey name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		// Named resource: read by wrapkeyname (held in the ID).
		data, err := client.FindResource(service.Sslwrapkey.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("sslwrapkey %s not found", n)
		}

		return nil
	}
}

func testAccCheckSslwrapkeyDestroy(s *terraform.State) error {
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_sslwrapkey" {
			continue
		}

		_, err := client.FindResource(service.Sslwrapkey.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("sslwrapkey %s still exists", rs.Primary.ID)
		}
	}
	return nil
}

// Datasource KEPT for sslwrapkey (datasource files present). Also skip-gated --
// wrap-key creation needs the FIPS/crypto subsystem.
const testAccSslwrapkeyDataSource_basic = `
variable "sslwrapkey_password_wo" {
  type      = string
  sensitive = true
}
variable "sslwrapkey_salt_wo" {
  type      = string
  sensitive = true
}

resource "citrixadc_sslwrapkey" "tf_sslwrapkey" {
  wrapkeyname         = "tf_wrapkey"
  password_wo         = var.sslwrapkey_password_wo
  password_wo_version = 1
  salt_wo             = var.sslwrapkey_salt_wo
  salt_wo_version     = 1
}

data "citrixadc_sslwrapkey" "tf_sslwrapkey" {
  wrapkeyname = citrixadc_sslwrapkey.tf_sslwrapkey.wrapkeyname
  depends_on  = [citrixadc_sslwrapkey.tf_sslwrapkey]
}
`

func TestAccSslwrapkeyDataSource_basic(t *testing.T) {
	// FIPS/crypto subsystem required: the datasource first creates the wrap key,
	// which is likely unsupported on the non-FIPS VPX testbed.
	t.Skip("sslwrapkey datasource depends on wrap-key creation, which needs the FIPS/crypto subsystem and is likely unsupported on the non-FIPS VPX testbed.")

	t.Setenv("TF_VAR_sslwrapkey_password_wo", "TODO_PLACEHOLDER")
	t.Setenv("TF_VAR_sslwrapkey_salt_wo", "TODO_PLACEHOLDER")

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSslwrapkeyDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_sslwrapkey.tf_sslwrapkey", "wrapkeyname", "tf_wrapkey"),
				),
			},
		},
	})
}
