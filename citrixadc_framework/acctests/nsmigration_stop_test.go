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
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// nsmigration_stop wraps POST /nsmigration?action=stop (aborts/rolls back an
// in-progress migration). Action-only: empty payload, no per-action GET, Delete is
// a state-only removal. See nsmigration_start_test.go for the multi-action notes.
//
// SKIPPED by default: stop requires a migration to be in progress. On a standalone
// appliance (and with no migration running) NITRO returns errorcode 257 "Operation
// not permitted [Migration is not in progress]". Remove the t.Skip only on a
// migration-capable testbed with an active migration.

const testAccNsmigrationStop_basic = `
resource "citrixadc_nsmigration_stop" "tf_stop" {
}
`

func TestAccNsmigrationStop_basic(t *testing.T) {
	t.Skip("Skipping nsmigration_stop test: stop requires an in-progress migration (NITRO errorcode 257 \"Migration is not in progress\"); not available on standalone appliances")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNsmigrationStop_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsmigrationActionExist("citrixadc_nsmigration_stop.tf_stop"),
					resource.TestCheckResourceAttrSet("citrixadc_nsmigration_stop.tf_stop", "id"),
				),
			},
		},
	})
}
