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

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

// NOTE on the systemscalablemgmtthreads_disable resource:
//   - systemscalablemgmtthreads is a NITRO object supporting multiple actions
//     (enable / disable) plus a get. Mirroring the appfwlearningdata package, each
//     action is its own action-only resource. This one wraps
//     POST /systemscalablemgmtthreads?action=disable.
//   - Create performs the disable action (empty payload; nodeid is a GET-only filter
//     and is rejected in the action body with NITRO errorcode 278). Read/Update are
//     no-ops (no per-action GET) and Delete is a state-only removal (the inverse is
//     the citrixadc_systemscalablemgmtthreads_enable resource). Live feature state
//     is queryable via the citrixadc_systemscalablemgmtthreads data source.
//   - The Exist check below only verifies the resource landed in Terraform state
//     with its synthetic id ("systemscalablemgmtthreads_disable").
//
// SKIPPED by default: the Scalable Management Threads feature is platform-gated.
// On standard VPX/lab appliances NITRO returns errorcode 1501 "Operation not
// supported on this platform" for enable/disable/get. Remove the t.Skip only on a
// platform that supports the feature.

const testAccSystemscalablemgmtthreadsDisable_basic = `
resource "citrixadc_systemscalablemgmtthreads_disable" "tf_disable" {
}
`

func TestAccSystemscalablemgmtthreadsDisable_basic(t *testing.T) {
	t.Skip("Skipping systemscalablemgmtthreads_disable test: the Scalable Management Threads feature is platform-gated (NITRO errorcode 1501 \"Operation not supported on this platform\" on standard appliances)")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// No CheckDestroy: disable's inverse is the separate _enable resource and
		// Delete is a state-only removal.
		Steps: []resource.TestStep{
			{
				Config: testAccSystemscalablemgmtthreadsDisable_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSystemscalablemgmtthreadsDisableExist("citrixadc_systemscalablemgmtthreads_disable.tf_disable", nil),
					resource.TestCheckResourceAttrSet("citrixadc_systemscalablemgmtthreads_disable.tf_disable", "id"),
				),
			},
		},
	})
}

// testAccCheckSystemscalablemgmtthreadsDisableExist is a state-only existence
// check. The disable action has no per-action GET endpoint, so it only verifies
// the resource landed in Terraform state with its synthetic ID.
func testAccCheckSystemscalablemgmtthreadsDisableExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No systemscalablemgmtthreads_disable ID is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		return nil
	}
}
