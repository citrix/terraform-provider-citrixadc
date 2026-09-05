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

// NOTE on the appfwarchive_export resource:
//   - Models the NITRO POST /appfwarchive?action=export endpoint.
//   - This is a one-shot side-effect (action) resource: Read is a no-op,
//     Update is a no-op, Delete is a no-op. There is no GET endpoint that
//     reports export state, so there is no datasource for this resource
//     (Bug Pattern 13 in .claude/agents/FeatureDeveloper.md).
//   - The exist check below only verifies that the resource has a state ID;
//     it does NOT (and cannot) verify the export side-effect via NITRO.
//   - There is no destroy check (Delete is a no-op; the side-effect cannot
//     be undone via API).
//   - The test is self-contained: doAppfwarchivePreChecks (in appfwarchive_test.go)
//     enables the AppFw feature and uploads a tar fixture to /var/tmp. Each config
//     below first imports that fixture via a citrixadc_appfwarchive prereq so the
//     archive `name` exists on the ADC, then exports it. `target` is a writable
//     local:/var/tmp path. No external archive host or pre-seeded archive needed.

const testAccAppfwarchiveExport_basic_step1 = `
resource "citrixadc_appfwarchive" "tf_prereq" {
  name = "tfappfwarch"
  src  = "local:tfappfwarchive.tar"
}

resource "citrixadc_appfwarchive_export" "tf_appfwarchive_export" {
  name       = citrixadc_appfwarchive.tf_prereq.name
  target     = "local:tfappfwarch_export"
  depends_on = [citrixadc_appfwarchive.tf_prereq]
}

`

const testAccAppfwarchiveExport_basic_step2 = `
resource "citrixadc_appfwarchive" "tf_prereq" {
  name = "tfappfwarch"
  src  = "local:tfappfwarchive.tar"
}

resource "citrixadc_appfwarchive_export" "tf_appfwarchive_export" {
  name       = citrixadc_appfwarchive.tf_prereq.name
  target     = "local:tfappfwarch_export_v2"
  depends_on = [citrixadc_appfwarchive.tf_prereq]
}

`

func TestAccAppfwarchiveExport_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doAppfwarchivePreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// The export action itself has no inverse on NITRO (Delete is a no-op),
		// but the citrixadc_appfwarchive prereq IS destroyable; verify it is
		// cleaned up so no archive is orphaned on the shared ADC.
		CheckDestroy: testAccCheckAppfwarchiveDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAppfwarchiveExport_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwarchiveExportExist("citrixadc_appfwarchive_export.tf_appfwarchive_export", nil),
					resource.TestCheckResourceAttrSet("citrixadc_appfwarchive_export.tf_appfwarchive_export", "name"),
					resource.TestCheckResourceAttrSet("citrixadc_appfwarchive_export.tf_appfwarchive_export", "target"),
				),
			},
			{
				// All attributes are RequiresReplace, so this exercises
				// destroy+recreate, not an in-place update.
				Config: testAccAppfwarchiveExport_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwarchiveExportExist("citrixadc_appfwarchive_export.tf_appfwarchive_export", nil),
					resource.TestCheckResourceAttrSet("citrixadc_appfwarchive_export.tf_appfwarchive_export", "name"),
					resource.TestCheckResourceAttrSet("citrixadc_appfwarchive_export.tf_appfwarchive_export", "target"),
				),
			},
		},
	})
}

// Datasource: SKIPPED — no datasource files exist for appfwarchive_export.
// NITRO has no GET endpoint for export state (Bug Pattern 13). The
// FeatureDeveloper agent deliberately omitted the datasource for this
// action-only resource; we do not generate a datasource test.

// testAccCheckAppfwarchiveExportExist is a state-only existence check.
// There is no NITRO GET endpoint for export state, so we cannot verify the
// export actually occurred on the ADC. We only assert that Terraform has
// recorded the resource in state with a non-empty ID (which equals `name`
// after a successful POST ?action=export).
func testAccCheckAppfwarchiveExportExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No appfwarchive_export ID is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		// No NITRO GET to verify against; presence of state ID is the only
		// confirmation we can make.
		return nil
	}
}
