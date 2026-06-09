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

// NOTE on the hasync resource:
//   - Models the NITRO POST /hasync?action=Force endpoint (capital F,
//     case-sensitive), which forces an HA configuration synchronization
//     between the nodes of an HA pair.
//   - This is an ACTION-ONLY (one-shot side-effect) resource: Create performs
//     the Force action, Read is a no-op (preserves state), Update is a no-op,
//     and Delete is a state-only removal. There is NO add/get/update/delete
//     endpoint, so the resource CANNOT be verified by reading it back from the
//     ADC, and the datasource was REMOVED (no NITRO GET endpoint exists).
//   - Attributes `force` (Optional, Bool) and `save` (Optional, String enum
//     YES|NO) are both RequiresReplace and both omittable. They ARE stored in
//     Terraform state (they are config attributes), so we can assert them even
//     though they are not Computed.
//   - The Exist check below only verifies that the resource landed in Terraform
//     state with its synthetic ID ("hasync"); it does NOT (and cannot) verify
//     the sync side-effect via NITRO.
//   - There is no CheckDestroy: the Force action has no inverse on NITRO, and
//     there is no GET-by-id to confirm absence; Delete is a state-only removal.
//
// This mirrors the action-only test precedent in hafiles_test.go,
// clusterfiles_test.go and aaasession_test.go (single apply step, state-only
// Exist check, no CheckDestroy, TestCheckResourceAttrSet on "id"), adapted for
// the Force action.
//
// ENVIRONMENT CAVEAT: forcing an HA sync only works on a node that is part of
// an HA pair. On a standalone testbed the Force action errors out. Following the
// repo's HA-test convention (see hafiles_test.go, hafailover_test.go and
// hanode_test.go), this test is gated on ADC_TESTBED == "HA_PAIR" and skipped
// otherwise.

// Single apply step: force/save are RequiresReplace, so there is no in-place
// update to exercise. Setting force = true and save = "YES" makes the action
// explicit and self-contained.
const testAccHasync_basic = `
resource "citrixadc_hasync" "tf_hasync" {
  force = true
  save  = "YES"
}

`

func TestAccHasync_basic(t *testing.T) {
	if adcTestbed != "HA_PAIR" {
		t.Skipf("ADC testbed is %s. Expected HA_PAIR (hasync Force requires an HA pair).", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// No CheckDestroy: the Force action has no inverse on NITRO and there is
		// no GET-by-id to confirm absence; Delete is a state-only removal.
		Steps: []resource.TestStep{
			{
				Config: testAccHasync_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHasyncExist("citrixadc_hasync.tf_hasync", nil),
					// "id" is the synthetic state handle "hasync".
					resource.TestCheckResourceAttrSet("citrixadc_hasync.tf_hasync", "id"),
					// Assert the config attributes actually set in HCL; both are
					// stored in state even though they are not Computed.
					resource.TestCheckResourceAttr("citrixadc_hasync.tf_hasync", "force", "true"),
					resource.TestCheckResourceAttr("citrixadc_hasync.tf_hasync", "save", "YES"),
				),
			},
		},
	})
}

// testAccCheckHasyncExist is a state-only existence check.
//
// hasync is an action-only resource: Read is a no-op and there is no GET-by-id
// endpoint, so we CANNOT verify the Force sync via NITRO. We only assert that
// Terraform recorded the resource in state with a non-empty ID (which equals
// the synthetic "hasync" after a successful POST ?action=Force). This mirrors
// testAccCheckHafilesExist, testAccCheckClusterfilesActionExist and
// testAccCheckAaasessionExist.
func testAccCheckHasyncExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No hasync ID is set")
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
