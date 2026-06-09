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

// NOTE on the clusterfiles resource:
//   - Models the NITRO POST /clusterfiles?action=sync endpoint, which
//     synchronizes configuration files across the nodes of a cluster.
//   - This is an ACTION-ONLY (one-shot side-effect) resource: Create performs
//     the sync action, Read is a no-op (preserves state), Update is a no-op,
//     and Delete is a state-only removal. There is NO add/get/update/delete
//     endpoint, so the resource CANNOT be verified by reading it back from the
//     ADC, and the datasource was REMOVED (no NITRO GET endpoint exists).
//   - The single attribute `mode` is OPTIONAL and RequiresReplace; it is a list
//     of directory/file groups to sync. When omitted, the ADC defaults to "all".
//     Possible values = all, bookmarks, ssl, imports, misc, dns, krb, AAA,
//     app_catalog, all_plus_misc, all_minus_misc.
//   - The Exist check below only verifies that the resource landed in Terraform
//     state with its synthetic ID ("clusterfiles"); it does NOT (and cannot)
//     verify the sync side-effect via NITRO.
//   - There is no CheckDestroy: the sync action has no inverse on NITRO, and
//     there is no GET-by-id to confirm absence; Delete is a state-only removal.
//
// This mirrors the action-only test precedent in aaasession_test.go and
// appfwarchive_export_test.go (single apply step, state-only Exist check, no
// CheckDestroy, TestCheckResourceAttrSet on "id"), adapted for the sync action.
//
// ENVIRONMENT CAVEAT: clusterfiles sync only works on a node that is part of a
// cluster (CCO). On a standalone testbed the sync action errors out. Following
// the repo's cluster-test convention (see cluster_test.go and
// clusterfiles_syncer_test.go), this test is gated on ADC_TESTBED == "CLUSTER"
// and skipped otherwise.

// Single apply step: mode is RequiresReplace, so there is no in-place update to
// exercise. mode = ["all"] makes the sync explicit and self-contained.
const testAccClusterfiles_basic = `
resource "citrixadc_clusterfiles" "tf_clusterfiles" {
  mode = ["all"]
}

`

func TestAccClusterfiles_basic(t *testing.T) {
	if adcTestbed != "CLUSTER" {
		t.Skipf("ADC testbed is %s. Expected CLUSTER (clusterfiles sync requires a cluster node).", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// No CheckDestroy: the sync action has no inverse on NITRO and there is
		// no GET-by-id to confirm absence; Delete is a state-only removal.
		Steps: []resource.TestStep{
			{
				Config: testAccClusterfiles_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClusterfilesActionExist("citrixadc_clusterfiles.tf_clusterfiles", nil),
					// "id" is the synthetic state handle "clusterfiles".
					resource.TestCheckResourceAttrSet("citrixadc_clusterfiles.tf_clusterfiles", "id"),
					// Assert the mode list actually set in HCL.
					resource.TestCheckResourceAttr("citrixadc_clusterfiles.tf_clusterfiles", "mode.#", "1"),
					resource.TestCheckResourceAttr("citrixadc_clusterfiles.tf_clusterfiles", "mode.0", "all"),
				),
			},
		},
	})
}

// testAccCheckClusterfilesActionExist is a state-only existence check.
//
// clusterfiles is an action-only resource: Read is a no-op and there is no
// GET-by-id endpoint, so we CANNOT verify the sync via NITRO. We only assert
// that Terraform recorded the resource in state with a non-empty ID (which
// equals the synthetic "clusterfiles" after a successful POST ?action=sync).
// This mirrors testAccCheckAaasessionExist.
//
// NOTE: the name is testAccCheckClusterfilesActionExist (not ...Exist) to avoid
// colliding with testAccCheckClusterfilesExist already declared in
// clusterfiles_syncer_test.go within this same package.
func testAccCheckClusterfilesActionExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No clusterfiles ID is set")
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
