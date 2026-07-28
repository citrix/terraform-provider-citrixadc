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
// FIPS HARDWARE REQUIRED -- TESTS ARE SKIP-GATED
// ============================================================================
// The sslfips resource performs Hardware Security Module (HSM) initialization
// via the `inithsm` argument. This requires a dedicated FIPS/HSM card. On a
// standard VPX appliance (no FIPS hardware) the NITRO call fails with errors
// such as "FIPS card not present" / "operation not supported on this platform".
//
// !!! DANGER -- DESTRUCTIVE !!!
// Initializing the HSM ERASES ALL existing FIPS key and certificate material on
// the appliance. NEVER run this against a production or shared FIPS appliance.
//
// Every test in this file is therefore t.Skip-gated. To run on a real FIPS
// appliance, remove the t.Skip line, supply real secret values via the
// TF_VAR_* environment variables below, and replace the TODO_PLACEHOLDER
// values, fully understanding that the HSM will be re-initialized.
// ============================================================================

package citrixadc

import (
	"fmt"
	"testing"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// Basic test uses the write-only (_wo) secret triples for sopassword,
// oldsopassword and userpassword. Values are supplied via TF_VAR_* below.
const testAccSslfips_basic_step1 = `
variable "sslfips_sopassword_wo" {
  type      = string
  sensitive = true
}
variable "sslfips_oldsopassword_wo" {
  type      = string
  sensitive = true
}
variable "sslfips_userpassword_wo" {
  type      = string
  sensitive = true
}

resource "citrixadc_sslfips" "tf_sslfips" {
  inithsm = "Level-2"

  sopassword_wo            = var.sslfips_sopassword_wo
  sopassword_wo_version    = 1
  oldsopassword_wo         = var.sslfips_oldsopassword_wo
  oldsopassword_wo_version = 1
  userpassword_wo          = var.sslfips_userpassword_wo
  userpassword_wo_version  = 1

  hsmlabel = "test_hsm_label"
}

`

const testAccSslfips_basic_step2 = `
variable "sslfips_sopassword_wo_2" {
  type      = string
  sensitive = true
}
variable "sslfips_oldsopassword_wo_2" {
  type      = string
  sensitive = true
}
variable "sslfips_userpassword_wo_2" {
  type      = string
  sensitive = true
}

resource "citrixadc_sslfips" "tf_sslfips" {
  inithsm = "Level-2"

  sopassword_wo            = var.sslfips_sopassword_wo_2
  sopassword_wo_version    = 2
  oldsopassword_wo         = var.sslfips_oldsopassword_wo_2
  oldsopassword_wo_version = 2
  userpassword_wo          = var.sslfips_userpassword_wo_2
  userpassword_wo_version  = 2

  hsmlabel = "test_hsm_label_updated"
}

`

func TestAccSslfips_basic(t *testing.T) {
	t.Skip("TODO: Requires review")
	// !!! DANGER -- DESTRUCTIVE & FIPS-HARDWARE-ONLY !!!
	// HSM initialization erases ALL FIPS key/cert data and requires a FIPS card.
	// The standalone VPX testbed has no FIPS hardware, so this test cannot run
	// here and must never be run against a shared/production FIPS appliance.

	// Replace these with real secret values before running on a FIPS appliance.
	t.Setenv("TF_VAR_sslfips_sopassword_wo", "TODO_PLACEHOLDER")
	t.Setenv("TF_VAR_sslfips_oldsopassword_wo", "TODO_PLACEHOLDER")
	t.Setenv("TF_VAR_sslfips_userpassword_wo", "TODO_PLACEHOLDER")
	t.Setenv("TF_VAR_sslfips_sopassword_wo_2", "TODO_PLACEHOLDER")
	t.Setenv("TF_VAR_sslfips_oldsopassword_wo_2", "TODO_PLACEHOLDER")
	t.Setenv("TF_VAR_sslfips_userpassword_wo_2", "TODO_PLACEHOLDER")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// Singleton resource: no CheckDestroy (sslfips always exists on ADC and
		// cannot be deleted).
		Steps: []resource.TestStep{
			{
				Config: testAccSslfips_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslfipsExist("citrixadc_sslfips.tf_sslfips", nil),
					resource.TestCheckResourceAttr("citrixadc_sslfips.tf_sslfips", "inithsm", "Level-2"),
					resource.TestCheckResourceAttr("citrixadc_sslfips.tf_sslfips", "sopassword_wo_version", "1"),
					resource.TestCheckResourceAttr("citrixadc_sslfips.tf_sslfips", "oldsopassword_wo_version", "1"),
					resource.TestCheckResourceAttr("citrixadc_sslfips.tf_sslfips", "userpassword_wo_version", "1"),
				),
			},
			{
				Config: testAccSslfips_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslfipsExist("citrixadc_sslfips.tf_sslfips", nil),
					resource.TestCheckResourceAttr("citrixadc_sslfips.tf_sslfips", "inithsm", "Level-2"),
					resource.TestCheckResourceAttr("citrixadc_sslfips.tf_sslfips", "sopassword_wo_version", "2"),
					resource.TestCheckResourceAttr("citrixadc_sslfips.tf_sslfips", "oldsopassword_wo_version", "2"),
					resource.TestCheckResourceAttr("citrixadc_sslfips.tf_sslfips", "userpassword_wo_version", "2"),
				),
			},
		},
	})
}

// Import test: sslfips is a singleton whose Terraform id is the synthetic
// constant "sslfips-config" (set in Create). ImportState uses passthrough, so
// the stored id is reused for import -- no ImportStateIdFunc is required.
// Skip-gated for the same FIPS-hardware/destructive reasons as the other tests
// in this file.
func TestAccSslfips_import(t *testing.T) {
	t.Skip("TODO: Requires review")
	// !!! DANGER -- DESTRUCTIVE & FIPS-HARDWARE-ONLY !!!
	// The create step performs HSM initialization, which erases ALL FIPS
	// key/cert data and requires a FIPS card not present on the VPX testbed.

	const resAddr = "citrixadc_sslfips.tf_sslfips"

	// Replace these with real secret values before running on a FIPS appliance.
	t.Setenv("TF_VAR_sslfips_sopassword_wo", "TODO_PLACEHOLDER")
	t.Setenv("TF_VAR_sslfips_oldsopassword_wo", "TODO_PLACEHOLDER")
	t.Setenv("TF_VAR_sslfips_userpassword_wo", "TODO_PLACEHOLDER")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// Singleton resource: no CheckDestroy (sslfips always exists on ADC and
		// cannot be deleted).
		Steps: []resource.TestStep{
			{
				Config: testAccSslfips_basic_step1,
			},
			{
				Config:                  testAccSslfips_basic_step1,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccCheckSslfipsExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No sslfips name is set")
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
		// Singleton resource: read without an ID.
		data, err := client.FindResource(service.Sslfips.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("sslfips %s not found", n)
		}

		return nil
	}
}
