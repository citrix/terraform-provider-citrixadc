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
// FIPS HARDWARE + FIRMWARE FILE REQUIRED -- TEST IS SKIP-GATED
// ============================================================================
// NOTE on the sslfips_change resource:
//   - Models the NITRO sslfips `change` operation. The operation is named
//     `change` but its URL query parameter is literally ?action=update (CLI:
//     "update ssl fips -fipsFW <path>"). It replaces the FIPS firmware on the
//     appliance's HSM card.
//   - This is an ACTION-ONLY (one-shot side-effect) resource: Create performs
//     the change via ActOnResource(service.Sslfips.Type(), &payload, "update"),
//     Read/Update are no-ops (the single attribute is RequiresReplace), and
//     Delete is a state-only removal. There is NO get/add/update/delete
//     endpoint, so the resource CANNOT be verified by reading it back from the
//     ADC, and it has NO datasource.
//   - The single attribute `fipsfw` is Required and RequiresReplace; it is the
//     path to the FIPS firmware file (max length 63).
//
// !!! DANGER -- FIPS-HARDWARE-ONLY !!!
// The change action requires a dedicated FIPS/HSM card AND a valid FIPS firmware
// file already present on the appliance. On a standard VPX appliance (no FIPS
// hardware, no firmware file) the NITRO call fails. This test is therefore
// t.Skip-gated, mirroring the parent sslfips_test.go. To run on a real FIPS
// appliance, remove the t.Skip line and replace the TODO_PLACEHOLDER fipsfw path
// with a real firmware file path present on that appliance.
//
// The Exist check below only verifies that the resource landed in Terraform
// state with its synthetic ID ("sslfips_change-<fipsfw>"); it does NOT (and
// cannot) verify the change side-effect via NITRO. There is no CheckDestroy: the
// change action has no inverse on NITRO and Delete is a state-only removal.
// ============================================================================

// Single apply step: the `fipsfw` attribute is RequiresReplace, so there is no
// in-place update to exercise. TODO_PLACEHOLDER must be replaced with a real
// FIPS firmware file path present on the target FIPS appliance.
const testAccSslfipsChange_basic = `
resource "citrixadc_sslfips_change" "tf_sslfips_change" {
  fipsfw = "TODO_PLACEHOLDER"
}

`

func TestAccSslfipsChange_basic(t *testing.T) {
	t.Skip("TODO: Requires review -- FIPS hardware and a real firmware file required; not runnable on the VPX testbed")
	// !!! DANGER -- FIPS-HARDWARE-ONLY !!!
	// Changing the FIPS firmware requires a FIPS/HSM card and a valid firmware
	// file, neither present on the VPX testbed.

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// No CheckDestroy: the change action has no inverse on NITRO and there is
		// no GET-by-id to confirm absence; Delete is a state-only removal.
		Steps: []resource.TestStep{
			{
				Config: testAccSslfipsChange_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslfipsChangeExist("citrixadc_sslfips_change.tf_sslfips_change", nil),
					resource.TestCheckResourceAttr("citrixadc_sslfips_change.tf_sslfips_change", "fipsfw", "TODO_PLACEHOLDER"),
					// "id" is the synthetic state handle "sslfips_change-<fipsfw>".
					resource.TestCheckResourceAttrSet("citrixadc_sslfips_change.tf_sslfips_change", "id"),
				),
			},
		},
	})
}

// testAccCheckSslfipsChangeExist is a state-only existence check.
//
// sslfips_change is an action-only resource: Read is a no-op and there is no
// GET-by-id endpoint, so we CANNOT verify the change via NITRO. We only assert
// that Terraform recorded the resource in state with a non-empty ID (which
// equals the synthetic "sslfips_change-<fipsfw>" after a successful POST
// ?action=update).
func testAccCheckSslfipsChangeExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No sslfips_change ID is set")
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
