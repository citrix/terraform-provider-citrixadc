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

// NOTE on the nsacls6_clear resource:
//   - Models the NITRO POST /nsacls6?action=clear endpoint, which clears the
//     configured extended IPv6 ACLs of the selected type.
//   - This is an ACTION-ONLY (one-shot side-effect) resource: Create performs the
//     "clear" action, Read is a no-op (preserves state), Update is a no-op (all
//     attributes are RequiresReplace), and Delete is a state-only removal. There
//     is NO add/get/delete endpoint, so the resource CANNOT be verified by reading
//     it back from the ADC.
//   - The single attribute `type` is OPTIONAL (enum CLASSIC | DFD, default CLASSIC)
//     and RequiresReplace. It selects the ACL type to clear.
//   - The Exist check below only verifies that the resource landed in Terraform
//     state with its synthetic ID ("nsacls6_clear"); it does NOT (and cannot)
//     verify the clear side-effect via NITRO.
//   - There is no CheckDestroy: the clear action has no inverse on NITRO, and
//     there is no GET-by-id to confirm absence; Delete is a state-only removal.
//
// This mirrors the action-only test precedent in clusterpropstatus_test.go and
// aaasession_test.go (single apply step, state-only Exist check, no CheckDestroy,
// TestCheckResourceAttrSet on "id"), adapted for the clear action.

// Single apply step: type is RequiresReplace, so there is no in-place update to
// exercise. CLASSIC is the default ACL type and the simplest valid value.
const testAccNsacls6Clear_basic = `
resource "citrixadc_nsacls6_clear" "tf_nsacls6_clear" {
  type = "CLASSIC"
}

`

func TestAccNsacls6Clear_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// No CheckDestroy: the clear action has no inverse on NITRO and there is
		// no GET-by-id to confirm absence; Delete is a state-only removal.
		Steps: []resource.TestStep{
			{
				Config: testAccNsacls6Clear_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsacls6ClearExist("citrixadc_nsacls6_clear.tf_nsacls6_clear", nil),
					resource.TestCheckResourceAttr("citrixadc_nsacls6_clear.tf_nsacls6_clear", "type", "CLASSIC"),
					// "id" is the synthetic state handle "nsacls6_clear".
					resource.TestCheckResourceAttrSet("citrixadc_nsacls6_clear.tf_nsacls6_clear", "id"),
				),
			},
		},
	})
}

// testAccCheckNsacls6ClearExist is a state-only existence check.
//
// nsacls6_clear is an action-only resource: Read is a no-op and there is no
// GET-by-id endpoint, so we CANNOT verify the clear via NITRO. We only assert
// that Terraform recorded the resource in state with a non-empty ID (which
// equals the synthetic "nsacls6_clear" after a successful POST ?action=clear).
func testAccCheckNsacls6ClearExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No nsacls6_clear ID is set")
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
