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
// sslfipssimtarget_enable models the NITRO sslfipssimtarget `?action=enable`
// action on the TARGET appliance. It requires a dedicated FIPS/HSM card. On a
// standard VPX appliance the NITRO call fails with errors such as "FIPS card
// not present" / "operation not supported on this platform".
//
// !!! DANGER -- FIPS SIM KEY IMPORT !!!
// This manipulates FIPS secret/key material exchanged between appliances. Never
// run against a production or shared FIPS appliance.
//
// This is an ACTION-ONLY resource:
//   - No NITRO GET endpoint, so the Exist check is state-only (no FindResource).
//     It only asserts the synthetic ID "sslfipssimtarget_enable".
//   - No NITRO DELETE endpoint, so there is no CheckDestroy.
//   - No datasource (no GET) -- there is intentionally no datasource test.
//
// The enable payload carries only keyvector and sourcesecret (both mandatory
// per the NITRO doc and CLI). The test below is t.Skip-gated. To run on a real
// FIPS appliance, remove the t.Skip line, supply the real secret value via
// TF_VAR_*, and replace the TODO_PLACEHOLDER keyvector path.
// ============================================================================

package citrixadc

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// Action-only resource: single apply (all attributes are RequiresReplace).
const testAccSslfipssimtargetEnable_basic = `
variable "sslfipssimtarget_enable_sourcesecret" {
  type      = string
  sensitive = true
}

resource "citrixadc_sslfipssimtarget_enable" "tf_sslfipssimtarget_enable" {
  keyvector    = "TODO_PLACEHOLDER"
  sourcesecret = var.sslfipssimtarget_enable_sourcesecret
}
`

func TestAccSslfipssimtargetEnable_basic(t *testing.T) {
	t.Skip("TODO: Requires review - requires FIPS/HSM hardware")
	// !!! DANGER -- FIPS SIM KEY IMPORT & FIPS-HARDWARE-ONLY !!!
	// Imports FIPS secret/key material; requires a dedicated FIPS/HSM card not
	// present on the VPX testbed. Never run against a shared/production FIPS box.

	// Replace this with a real secret value before running on a FIPS appliance.
	t.Setenv("TF_VAR_sslfipssimtarget_enable_sourcesecret", "TODO_PLACEHOLDER")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// Action-only resource: no CheckDestroy (no NITRO DELETE endpoint).
		Steps: []resource.TestStep{
			{
				Config: testAccSslfipssimtargetEnable_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslfipssimtargetEnableExist("citrixadc_sslfipssimtarget_enable.tf_sslfipssimtarget_enable", nil),
					resource.TestCheckResourceAttr("citrixadc_sslfipssimtarget_enable.tf_sslfipssimtarget_enable", "id", "sslfipssimtarget_enable"),
					resource.TestCheckResourceAttr("citrixadc_sslfipssimtarget_enable.tf_sslfipssimtarget_enable", "keyvector", "TODO_PLACEHOLDER"),
				),
			},
		},
	})
}

// Action-only, state-only Exist check: there is NO NITRO GET endpoint for
// sslfipssimtarget_enable, so this only verifies the resource is present in
// Terraform state and carries the synthetic ID "sslfipssimtarget_enable". No
// FindResource call is possible.
func testAccCheckSslfipssimtargetEnableExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No sslfipssimtarget_enable ID is set")
		}

		// Action-only resource: the ID is a synthetic constant assigned in Create.
		if rs.Primary.ID != "sslfipssimtarget_enable" {
			return fmt.Errorf("Unexpected sslfipssimtarget_enable ID %q, expected %q", rs.Primary.ID, "sslfipssimtarget_enable")
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
