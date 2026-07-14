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
// sslfipssimsource drives the FIPS Secure Information Management (SIM) key
// export flow on the SOURCE appliance (NITRO `enable`/`init` actions only --
// no get/add/delete). It requires a dedicated FIPS/HSM card. On a standard VPX
// appliance the NITRO call fails with errors such as "FIPS card not present" /
// "operation not supported on this platform".
//
// !!! DANGER -- FIPS SIM KEY EXPORT !!!
// This manipulates FIPS secret/key material exported between appliances. Never
// run against a production or shared FIPS appliance.
//
// This is an ACTION-ONLY resource:
//   - No NITRO GET endpoint, so the Exist check is state-only (no FindResource).
//   - No NITRO DELETE endpoint, so there is no CheckDestroy.
//   - Datasource was removed (no GET) -- there is intentionally no datasource test.
//
// The test below is t.Skip-gated. To run on a real FIPS appliance, remove the
// t.Skip line, supply real secret values via TF_VAR_*, and replace the
// TODO_PLACEHOLDER certfile path.
// ============================================================================

package citrixadc

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// Action-only resource: single apply (all attributes are RequiresReplace).
const testAccSslfipssimsource_basic = `
variable "sslfipssimsource_sourcesecret" {
  type      = string
  sensitive = true
}
variable "sslfipssimsource_targetsecret" {
  type      = string
  sensitive = true
}

resource "citrixadc_sslfipssimsource" "tf_sslfipssimsource" {
  certfile     = "TODO_PLACEHOLDER"
  sourcesecret = var.sslfipssimsource_sourcesecret
  targetsecret = var.sslfipssimsource_targetsecret
}
`

func TestAccSslfipssimsource_basic(t *testing.T) {
	t.Skip("TODO: Requires review")
	// !!! DANGER -- FIPS SIM KEY EXPORT & FIPS-HARDWARE-ONLY !!!
	// Exports FIPS secret/key material; requires a dedicated FIPS/HSM card not
	// present on the VPX testbed. Never run against a shared/production FIPS box.

	// Replace these with real secret values before running on a FIPS appliance.
	t.Setenv("TF_VAR_sslfipssimsource_sourcesecret", "TODO_PLACEHOLDER")
	t.Setenv("TF_VAR_sslfipssimsource_targetsecret", "TODO_PLACEHOLDER")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// Action-only resource: no CheckDestroy (no NITRO DELETE endpoint).
		Steps: []resource.TestStep{
			{
				Config: testAccSslfipssimsource_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslfipssimsourceExist("citrixadc_sslfipssimsource.tf_sslfipssimsource", nil),
				),
			},
		},
	})
}

// Action-only Exist check: there is NO NITRO GET endpoint for sslfipssimsource,
// so this only verifies the resource is present in Terraform state. No
// FindResource call is possible.
func testAccCheckSslfipssimsourceExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No sslfipssimsource ID is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		// No NITRO GET endpoint for this action-only resource; state presence is
		// the only thing we can assert.
		return nil
	}
}
