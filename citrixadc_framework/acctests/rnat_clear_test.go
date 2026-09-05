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
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// NOTE on the rnat_clear resource:
//   - Models a *set* of legacy (UNNAMED) RNAT rules under one synthetic handle
//     (rnatsname). `rnat` is a Framework SetNestedBlock, so it uses the SDK v2
//     BLOCK HCL syntax `rnat { ... }` (not `rnat = [ { ... } ]`).
//   - The unnamed `set rnat` path is UNSUPPORTED on NetScaler 14.1+ (errorcode 275
//     "Operation not supported by device"): RNAT moved to the named RNAT4 model,
//     which is the separate `citrixadc_rnat` resource. Create/Update now SURFACE
//     that failure instead of silently swallowing it (as SDK v2 did), so on 14.1.x
//     firmware the apply legitimately fails rather than reporting a phantom success.
//
// The test therefore (1) proves the block HCL syntax is accepted via a
// firmware-independent PlanOnly step, and (2) proves the failed `set rnat` is now
// surfaced via an ExpectError step (all current lab firmware is 14.1.x). On older
// firmware where unnamed RNAT is supported, step 2 would need to become a real
// apply + appliance check.

const testAccRnatClear_basic = `
resource "citrixadc_rnat_clear" "foo" {
  rnatsname = "tf_rnat_clear"
  rnat {
    network = "192.168.96.0"
    netmask = "255.255.240.0"
  }
}
`

func TestAccRnatClear_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				// (A) The `rnat { ... }` block HCL syntax is accepted and plans.
				// Plan-only, so this is firmware-independent (no NITRO apply).
				Config:             testAccRnatClear_basic,
				PlanOnly:           true,
				ExpectNonEmptyPlan: true,
			},
			{
				// (B) A failed `set rnat` is now surfaced (not swallowed). Unnamed
				// RNAT is unsupported on 14.1.x -> errorcode 275.
				Config:      testAccRnatClear_basic,
				ExpectError: regexp.MustCompile(`(?i)Unable to apply rnat rule`),
			},
		},
	})
}
