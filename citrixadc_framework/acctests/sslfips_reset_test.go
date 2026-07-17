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
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// ============================================================================
// FIPS HARDWARE REQUIRED & DESTRUCTIVE -- TEST IS SKIP-GATED
// ============================================================================
// NOTE on the sslfips_reset resource:
//   - Models the NITRO POST /sslfips?action=reset endpoint (CLI: "reset ssl
//     fips"), which resets the FIPS Hardware Security Module (HSM) on the
//     appliance.
//   - This is an ACTION-ONLY (one-shot side-effect) resource: Create performs
//     the reset via ActOnResource(service.Sslfips.Type(), &payload, "reset"),
//     Read/Update are no-ops, and Delete is a state-only removal. There is NO
//     get/add/update/delete endpoint, so the resource CANNOT be verified by
//     reading it back from the ADC, and it has NO datasource.
//   - The reset payload carries NO input attributes; the schema exposes only the
//     Computed synthetic id "sslfips_reset".
//
// !!! DANGER -- FIPS-HARDWARE-ONLY & DISRUPTIVE !!!
// The reset action requires a dedicated FIPS/HSM card. On a standard VPX
// appliance (no FIPS hardware) the NITRO call fails with errors such as "FIPS
// card not present" / "operation not supported on this platform". Resetting the
// FIPS card is disruptive to a shared appliance. This test is therefore
// t.Skip-gated, mirroring the parent sslfips_test.go. To run on a real FIPS
// appliance, remove the t.Skip line, fully understanding the consequences.
//
// The Exist check below only verifies that the resource landed in Terraform
// state with its synthetic ID ("sslfips_reset"); it does NOT (and cannot)
// verify the reset side-effect via NITRO. There is no CheckDestroy: the reset
// action has no inverse on NITRO and Delete is a state-only removal.
// ============================================================================

// Single apply step: the resource has no writable attributes, so there is no
// in-place update to exercise.
const testAccSslfipsReset_basic = `
resource "citrixadc_sslfips_reset" "tf_sslfips_reset" {
}

`

func TestAccSslfipsReset_basic(t *testing.T) {
	t.Skip("TODO: Requires review -- FIPS hardware required and disruptive; not runnable on the VPX testbed")
	// !!! DANGER -- FIPS-HARDWARE-ONLY & DISRUPTIVE !!!
	// Resetting the FIPS card requires a FIPS/HSM card not present on the VPX
	// testbed and is disruptive to a shared appliance.

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// No CheckDestroy: the reset action has no inverse on NITRO and there is
		// no GET-by-id to confirm absence; Delete is a state-only removal.
		Steps: []resource.TestStep{
			{
				Config: testAccSslfipsReset_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslfipsResetExist("citrixadc_sslfips_reset.tf_sslfips_reset", nil),
					// "id" is the synthetic state handle "sslfips_reset".
					resource.TestCheckResourceAttrSet("citrixadc_sslfips_reset.tf_sslfips_reset", "id"),
				),
			},
		},
	})
}

// testAccCheckSslfipsResetExist is a state-only existence check.
//
// sslfips_reset is an action-only resource: Read is a no-op and there is no
// GET-by-id endpoint, so we CANNOT verify the reset via NITRO. We only assert
// that Terraform recorded the resource in state with a non-empty ID (which
// equals the synthetic "sslfips_reset" after a successful POST ?action=reset).
func testAccCheckSslfipsResetExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No sslfips_reset ID is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		// No NITRO GET-by-id to verify against; presence of the synthetic state
		// ID is the only confirmation we can make for an action-only resource.
		return nil
	}
}
