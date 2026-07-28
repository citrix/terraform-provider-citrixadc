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

// NOTE on the lbpersistentsessions_clear resource:
//   - Models the NITRO POST /lbpersistentsessions?action=clear endpoint, which
//     flushes load-balancing persistence sessions. This is the per-action
//     refactor of the former citrixadc_lbpersistentsessions resource (verb
//     "clear").
//   - This is an ACTION-ONLY (one-shot side-effect) resource: Create performs
//     the clear action, Read is a no-op (preserves state), Update is a no-op,
//     and Delete is a state-only removal. There is NO add/update/delete
//     endpoint, so the resource CANNOT be verified by reading it back from the
//     ADC. (NITRO does expose get(all)/count, which backs the read-only
//     datasource below, but there is no GET-by-id to re-resolve "this clear".)
//   - All attributes (vserver, persistenceparameter, nodeid) are OPTIONAL clear
//     filters; none is Required, and all are RequiresReplace. nodeid is a
//     GET-only cluster filter and is intentionally excluded from the clear
//     payload. When all are omitted, the clear flushes ALL persistence sessions,
//     which is the simplest self-contained config (it needs no pre-existing
//     vserver on the testbed) and is what the basic test uses.
//   - The Exist check below only verifies that the resource landed in Terraform
//     state with its synthetic ID ("lbpersistentsessions_clear"); it does NOT
//     (and cannot) verify the clear side-effect via NITRO.
//   - There is no CheckDestroy: the clear action has no inverse on NITRO, and
//     there is no GET-by-id to confirm absence; Delete is a state-only removal.
//
// This mirrors the action-only test precedent in aaasession_test.go and
// clusterpropstatus_test.go (single apply step, state-only Exist check, no
// CheckDestroy, TestCheckResourceAttrSet on "id"), adapted for the clear action.
//
// ENVIRONMENT: clear works on a standalone testbed; no ADC_TESTBED gate is
// needed (the schema/precedent indicate none).

// Single apply step: all attributes are RequiresReplace, so there is no
// in-place update to exercise. An empty config flushes all persistence sessions
// (simplest self-contained form, no pre-existing vserver required).
const testAccLbpersistentsessionsClear_basic = `
resource "citrixadc_lbpersistentsessions_clear" "tf_lbpersistentsessions_clear" {
}

`

func TestAccLbpersistentsessionsClear_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// No CheckDestroy: the clear action has no inverse on NITRO and there is
		// no GET-by-id to confirm absence; Delete is a state-only removal.
		Steps: []resource.TestStep{
			{
				Config: testAccLbpersistentsessionsClear_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbpersistentsessionsClearExist("citrixadc_lbpersistentsessions_clear.tf_lbpersistentsessions_clear", nil),
					// "id" is the synthetic state handle "lbpersistentsessions_clear".
					resource.TestCheckResourceAttrSet("citrixadc_lbpersistentsessions_clear.tf_lbpersistentsessions_clear", "id"),
				),
			},
		},
	})
}

// testAccCheckLbpersistentsessionsClearExist is a state-only existence check.
//
// lbpersistentsessions_clear is an action-only resource: Read is a no-op and
// there is no GET-by-id endpoint, so we CANNOT verify the clear via NITRO. We
// only assert that Terraform recorded the resource in state with a non-empty ID
// (which equals the synthetic "lbpersistentsessions_clear" after a successful
// POST ?action=clear). This mirrors testAccCheckAaasessionExist /
// testAccCheckClusterpropstatusExist.
func testAccCheckLbpersistentsessionsClearExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lbpersistentsessions_clear ID is set")
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
