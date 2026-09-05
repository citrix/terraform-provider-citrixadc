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

// NOTE on the nsmigration_start resource:
//   - nsmigration is a NITRO object supporting multiple actions (start / stop /
//     complete) plus a get. Mirroring the systemscalablemgmtthreads package, each
//     action is its own action-only resource. This one wraps
//     POST /nsmigration?action=start (begins a NetScaler session migration).
//   - Create performs the start action (empty payload; dumpsession is a GET-only
//     field rejected in the action body with NITRO errorcode 278). Read/Update are
//     no-ops (no per-action GET) and Delete is a state-only removal (the inverse is
//     the citrixadc_nsmigration_stop resource). Live migration state is queryable
//     via the citrixadc_nsmigration data source.
//
// SKIPPED by default: session migration is only available in an HA/migration
// deployment. On a standalone appliance NITRO returns errorcode 257 "Operation not
// permitted [Migration is not supported in standalone boxes]". Remove the t.Skip
// only on a migration-capable testbed.

const testAccNsmigrationStart_basic = `
resource "citrixadc_nsmigration_start" "tf_start" {
}
`

func TestAccNsmigrationStart_basic(t *testing.T) {
	t.Skip("Skipping nsmigration_start test: session migration is not supported on standalone appliances (NITRO errorcode 257 \"Migration is not supported in standalone boxes\")")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNsmigrationStart_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsmigrationActionExist("citrixadc_nsmigration_start.tf_start"),
					resource.TestCheckResourceAttrSet("citrixadc_nsmigration_start.tf_start", "id"),
				),
			},
		},
	})
}

// testAccCheckNsmigrationActionExist is a shared state-only existence check for the
// action-only nsmigration_{start,stop,complete} resources (no per-action GET).
func testAccCheckNsmigrationActionExist(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No %s ID is set", n)
		}
		return nil
	}
}
