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

// NOTE on the clusterpropstatus_clear resource:
//   - Models the NITRO POST /clusterpropstatus?action=clear endpoint, which
//     clears the cluster property-propagation status counters. This works only
//     on a cluster IP (CLIP) / cluster node.
//   - This is an ACTION-ONLY (one-shot side-effect) resource: Create performs
//     the clear action, Read is a no-op (preserves state), Update is a no-op,
//     and Delete is a state-only removal. There is NO add/update/delete
//     endpoint, so the resource CANNOT be verified by reading it back from the
//     ADC.
//   - The single attribute `nodeid` is OPTIONAL (range 0-31) and RequiresReplace;
//     it identifies the cluster node whose propagation status to clear. When
//     omitted the clear applies to all nodes, which is the simplest self-contained
//     config and is what the basic test uses.
//   - The Exist check below only verifies that the resource landed in Terraform
//     state with its synthetic ID ("clusterpropstatus_clear"); it does NOT (and
//     cannot) verify the clear side-effect via NITRO.
//   - There is no CheckDestroy: the clear action has no inverse on NITRO, and
//     there is no GET-by-id to confirm absence; Delete is a state-only removal.
//
// This mirrors the action-only test precedent in aaasession_test.go and
// clusterfiles_test.go (single apply step, state-only Exist check, no
// CheckDestroy, TestCheckResourceAttrSet on "id"), adapted for the clear action.
//
// ENVIRONMENT CAVEAT: clusterpropstatus clear only works on a cluster (CLIP).
// On a standalone testbed the clear action errors out. Following the repo's
// cluster-test convention (see clusterfiles_test.go and clusternodegroup_test.go),
// this test is gated on ADC_TESTBED == "CLUSTER" and skipped otherwise.

// Single apply step: nodeid is RequiresReplace, so there is no in-place update to
// exercise. Omitting nodeid clears the status for all nodes (simplest valid form).
const testAccClusterpropstatusClear_basic = `
resource "citrixadc_clusterpropstatus_clear" "tf_clusterpropstatus_clear" {
}

`

func TestAccClusterpropstatusClear_basic(t *testing.T) {
	if adcTestbed != "CLUSTER" {
		t.Skipf("ADC testbed is %s. Expected CLUSTER (clusterpropstatus clear requires a cluster / CLIP).", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// No CheckDestroy: the clear action has no inverse on NITRO and there is
		// no GET-by-id to confirm absence; Delete is a state-only removal.
		Steps: []resource.TestStep{
			{
				Config: testAccClusterpropstatusClear_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClusterpropstatusClearExist("citrixadc_clusterpropstatus_clear.tf_clusterpropstatus_clear", nil),
					// "id" is the synthetic state handle "clusterpropstatus_clear".
					resource.TestCheckResourceAttrSet("citrixadc_clusterpropstatus_clear.tf_clusterpropstatus_clear", "id"),
				),
			},
		},
	})
}

// testAccCheckClusterpropstatusClearExist is a state-only existence check.
//
// clusterpropstatus_clear is an action-only resource: Read is a no-op and there
// is no GET-by-id endpoint, so we CANNOT verify the clear via NITRO. We only
// assert that Terraform recorded the resource in state with a non-empty ID
// (which equals the synthetic "clusterpropstatus_clear" after a successful POST
// ?action=clear). This mirrors testAccCheckAaasessionExist /
// testAccCheckClusterfilesActionExist.
func testAccCheckClusterpropstatusClearExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No clusterpropstatus_clear ID is set")
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
