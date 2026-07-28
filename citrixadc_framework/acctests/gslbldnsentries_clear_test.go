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

// NOTE on the gslbldnsentries_clear resource:
//   - Models the NITRO POST /gslbldnsentries?action=clear endpoint, which flushes
//     the GSLB LDNS (local DNS) RTT entries. The clear action takes no arguments;
//     the payload is always empty.
//   - This is an ACTION-ONLY (one-shot side-effect) resource: Create performs the
//     clear action, Read is a no-op (preserves state), Update is a no-op, and
//     Delete is a state-only removal. There is NO add/update/delete endpoint, so
//     the resource CANNOT be verified by reading it back from the ADC. (NITRO does
//     expose get(all)/count, which backs the read-only datasource
//     citrixadc_gslbldnsentries below, but there is no GET-by-id to re-resolve
//     "this clear".)
//   - The single attribute `nodeid` is OPTIONAL (a GET-only cluster filter,
//     intentionally excluded from the clear payload) and RequiresReplace. When
//     omitted the clear flushes ALL LDNS entries, which is the simplest
//     self-contained config (no participating entity required) and is what the
//     basic test uses.
//   - The Exist check below only verifies that the resource landed in Terraform
//     state with its synthetic ID ("gslbldnsentries_clear"); it does NOT (and
//     cannot) verify the clear side-effect via NITRO.
//   - There is no CheckDestroy: the clear action has no inverse on NITRO, and there
//     is no GET-by-id to confirm absence; Delete is a state-only removal.
//
// This mirrors the action-only test precedent in aaasession_test.go,
// lbpersistentsessions_test.go and clusterpropstatus_test.go (single apply step,
// state-only Exist check, no CheckDestroy, TestCheckResourceAttrSet on "id"),
// adapted for the clear action.
//
// ENVIRONMENT: clearing LDNS entries works on any box (it is a no-op if none
// exist); no ADC_TESTBED gate is needed.

// Single apply step: nodeid is RequiresReplace, so there is no in-place update to
// exercise. An empty config clears ALL LDNS entries (simplest self-contained
// form, no participating entity required).
const testAccGslbldnsentriesClear_basic = `
resource "citrixadc_gslbldnsentries_clear" "tf_gslbldnsentries_clear" {
}

`

func TestAccGslbldnsentriesClear_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// No CheckDestroy: the clear action has no inverse on NITRO and there is
		// no GET-by-id to confirm absence; Delete is a state-only removal.
		Steps: []resource.TestStep{
			{
				Config: testAccGslbldnsentriesClear_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGslbldnsentriesClearExist("citrixadc_gslbldnsentries_clear.tf_gslbldnsentries_clear", nil),
					// "id" is the synthetic state handle "gslbldnsentries_clear".
					resource.TestCheckResourceAttrSet("citrixadc_gslbldnsentries_clear.tf_gslbldnsentries_clear", "id"),
				),
			},
		},
	})
}

// testAccCheckGslbldnsentriesClearExist is a state-only existence check.
//
// gslbldnsentries_clear is an action-only resource: Read is a no-op and there is
// no GET-by-id endpoint, so we CANNOT verify the clear via NITRO. We only assert
// that Terraform recorded the resource in state with a non-empty ID (which equals
// the synthetic "gslbldnsentries_clear" after a successful POST ?action=clear).
// This mirrors testAccCheckAaasessionExist / testAccCheckLbpersistentsessionsExist
// / testAccCheckClusterpropstatusExist.
func testAccCheckGslbldnsentriesClearExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No gslbldnsentries_clear ID is set")
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
