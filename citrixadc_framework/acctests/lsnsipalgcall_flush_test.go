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

// NOTE on the lsnsipalgcall_flush resource:
//   - Models the NITRO POST /lsnsipalgcall?action=flush endpoint, invoked via
//     ActOnResource(service.Lsnsipalgcall.Type(), &payload, "flush").
//   - This is an ACTION-ONLY (one-shot side-effect) resource: Create performs
//     the flush action, Read is a no-op (preserves state), Update is a no-op,
//     and Delete is a state-only removal. There is no add/update/delete endpoint.
//   - `callid` is the only Required attribute; it identifies the SIP ALG call to
//     flush and is the only field placed in the flush payload. It is
//     RequiresReplace (a different callid forces a new resource).
//   - `nodeid` is an Optional GET-only cluster filter and is intentionally
//     excluded from the flush payload (Pattern 15). It is left unset here so the
//     config is self-contained for a standalone testbed.
//   - The flush of a non-existent call is a harmless no-op on the ADC, so the
//     test does not require a pre-existing SIP ALG call on the testbed.
//   - The Exist check below only verifies that the resource landed in Terraform
//     state with its synthetic ID ("lsnsipalgcall_flush"); it does NOT (and
//     cannot) verify the flush side-effect via NITRO, because Read is a no-op and
//     there is no stable GET-by-id to re-resolve "this flushed call".
//   - There is no CheckDestroy: the flush action has no inverse on NITRO, and
//     Delete is a state-only removal.
//
// This mirrors the action-only test precedent in protocolhttpband_clear_test.go
// and endpointinfo_clear_test.go (single apply step, state-only Exist check, no
// CheckDestroy, TestCheckResourceAttrSet on "id"), adapted for the flush action
// and its single Required callid attribute.

// Single apply step: callid is RequiresReplace, so there is no in-place update
// to exercise. The flush action is self-contained via a plausible callid that
// need not exist on the testbed.
const testAccLsnsipalgcallFlush_basic = `
resource "citrixadc_lsnsipalgcall_flush" "tf_lsnsipalgcall_flush" {
  callid = "12345-abcde"
}

`

func TestAccLsnsipalgcallFlush_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// No CheckDestroy: the flush action has no inverse on NITRO and there is
		// no GET-by-id to confirm absence; Delete is a state-only removal.
		Steps: []resource.TestStep{
			{
				Config: testAccLsnsipalgcallFlush_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsnsipalgcallFlushExist("citrixadc_lsnsipalgcall_flush.tf_lsnsipalgcall_flush", nil),
					// "id" is the synthetic state handle "lsnsipalgcall_flush".
					resource.TestCheckResourceAttrSet("citrixadc_lsnsipalgcall_flush.tf_lsnsipalgcall_flush", "id"),
					// Assert only the attribute actually set in HCL.
					resource.TestCheckResourceAttr("citrixadc_lsnsipalgcall_flush.tf_lsnsipalgcall_flush", "callid", "12345-abcde"),
				),
			},
		},
	})
}

// testAccCheckLsnsipalgcallFlushExist is a state-only existence check.
//
// lsnsipalgcall_flush is an action-only resource: Read is a no-op and there is no
// stable GET-by-id endpoint to re-resolve a flushed call, so we CANNOT verify the
// flush via NITRO. We only assert that Terraform recorded the resource in state
// with a non-empty ID (which equals the synthetic "lsnsipalgcall_flush" after a
// successful POST ?action=flush). This mirrors testAccCheckProtocolhttpbandClearExist.
func testAccCheckLsnsipalgcallFlushExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lsnsipalgcall_flush ID is set")
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
