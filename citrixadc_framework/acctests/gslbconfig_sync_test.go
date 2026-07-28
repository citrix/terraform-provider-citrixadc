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

// NOTE on the gslbconfig_sync resource:
//   - Models the NITRO POST /gslbconfig?action=sync endpoint, which synchronizes
//     the GSLB configuration from the master site across all participating GSLB
//     sites.
//   - This is an ACTION-ONLY (one-shot side-effect) resource: Create performs the
//     sync action, Read is a no-op (preserves state), Update is a no-op, and
//     Delete is a state-only removal. There is NO add/get/update/delete endpoint,
//     so the resource CANNOT be verified by reading it back from the ADC, and the
//     datasource was REMOVED (no NITRO GET endpoint exists). gslbconfig_sync is the
//     canonical Pattern 13 (datasource-removed) reference.
//   - All attributes are OPTIONAL: preview, debug, forcesync, nowarn, saveconfig,
//     command. They form a mutually-exclusive group (command cannot be combined
//     with forcesync/preview; preview is mutually exclusive with saveconfig);
//     debug is a standalone verbosity flag. This test sets only the standalone
//     `debug = true` flag, which maps to "sync gslb config -debug" and is the
//     simplest self-contained, non-conflicting form. A bare empty block
//     (`resource "citrixadc_gslbconfig_sync" "tf_gslbconfig" {}`) running
//     "sync gslb config" with no args is equally valid.
//   - The Exist check below only verifies that the resource landed in Terraform
//     state with its synthetic ID ("gslbconfig_sync"); it does NOT (and cannot)
//     verify the sync side-effect via NITRO.
//   - There is no CheckDestroy: the sync action has no inverse on NITRO, and there
//     is no GET-by-id to confirm absence; Delete is a state-only removal.
//
// This mirrors the action-only test precedent in clusterfiles_test.go and
// hafiles_test.go (single apply step, state-only Exist check, no CheckDestroy,
// TestCheckResourceAttrSet on "id"), adapted for the GSLB sync action.
//
// ENVIRONMENT CAVEAT: "sync gslb config" is only meaningful when GSLB sites are
// configured. On a box with no GSLB sites the action may warn/no-op or error.
// This test does not over-gate on a specific ADC_TESTBED (there is no GSLB-sites
// testbed marker); run it manually against an ADC that has GSLB sites configured.

// Single apply step. All attributes are Optional; the action runs unconditionally
// on Create. `debug = true` makes the sync explicit and self-contained while
// avoiding the mutually-exclusive forcesync/command/preview/saveconfig options
// (which would also require real peer GSLB sites to be useful).
const testAccGslbconfigSync_basic = `
resource "citrixadc_gslbconfig_sync" "tf_gslbconfig" {
  debug = true
}

`

func TestAccGslbconfigSync_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// No CheckDestroy: the sync action has no inverse on NITRO and there is
		// no GET-by-id to confirm absence; Delete is a state-only removal.
		Steps: []resource.TestStep{
			{
				Config: testAccGslbconfigSync_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGslbconfigSyncExist("citrixadc_gslbconfig_sync.tf_gslbconfig", nil),
					// "id" is the synthetic state handle "gslbconfig_sync".
					resource.TestCheckResourceAttrSet("citrixadc_gslbconfig_sync.tf_gslbconfig", "id"),
					// Assert the flag actually set in HCL.
					resource.TestCheckResourceAttr("citrixadc_gslbconfig_sync.tf_gslbconfig", "debug", "true"),
				),
			},
		},
	})
}

// testAccCheckGslbconfigSyncExist is a state-only existence check.
//
// gslbconfig_sync is an action-only resource: Read is a no-op and there is no
// GET-by-id endpoint, so we CANNOT verify the sync via NITRO. We only assert that
// Terraform recorded the resource in state with a non-empty ID (which equals the
// synthetic "gslbconfig_sync" after a successful POST ?action=sync). This mirrors
// testAccCheckHafilesExist and testAccCheckClusterfilesActionExist.
func testAccCheckGslbconfigSyncExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No gslbconfig_sync ID is set")
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
