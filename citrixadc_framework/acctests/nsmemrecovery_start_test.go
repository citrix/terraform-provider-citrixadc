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

// NOTE on the nsmemrecovery_start resource:
//   - Models the NITRO POST /nsmemrecovery?action=start endpoint, which recovers
//     a configurable percentage of memory from the freepools. Because the object
//     is action-only, the TF resource is named with the action suffix
//     (citrixadc_nsmemrecovery_start), mirroring citrixadc_appfwarchive_export.
//   - This is an ACTION-ONLY resource: Create performs the start action, Read is
//     a no-op, Update re-runs the start action when "percentage" changes, and
//     Delete is a state-only removal. There is NO get/add/set/delete/unset
//     endpoint, so the resource CANNOT be verified by reading it back and it has
//     NO datasource.
//   - The Exist check below only verifies the resource landed in Terraform state
//     with its synthetic id ("nsmemrecovery-config"); it cannot verify the side
//     effect via NITRO.
//
// This mirrors the action-only test precedent (single apply step + update step,
// state-only Exist check, no CheckDestroy, TestCheckResourceAttrSet on "id").

const testAccNsmemrecoveryStart_basic = `
resource "citrixadc_nsmemrecovery_start" "tf_nsmemrecovery_start" {
	percentage = 10
}
`

const testAccNsmemrecoveryStart_update = `
resource "citrixadc_nsmemrecovery_start" "tf_nsmemrecovery_start" {
	percentage = 20
}
`

func TestAccNsmemrecoveryStart_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// No CheckDestroy: the start action has no inverse on NITRO and there is
		// no GET-by-id to confirm absence; Delete is a state-only removal.
		Steps: []resource.TestStep{
			{
				Config: testAccNsmemrecoveryStart_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsmemrecoveryStartExist("citrixadc_nsmemrecovery_start.tf_nsmemrecovery_start", nil),
					resource.TestCheckResourceAttr("citrixadc_nsmemrecovery_start.tf_nsmemrecovery_start", "percentage", "10"),
					resource.TestCheckResourceAttrSet("citrixadc_nsmemrecovery_start.tf_nsmemrecovery_start", "id"),
				),
			},
			{
				Config: testAccNsmemrecoveryStart_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsmemrecoveryStartExist("citrixadc_nsmemrecovery_start.tf_nsmemrecovery_start", nil),
					resource.TestCheckResourceAttr("citrixadc_nsmemrecovery_start.tf_nsmemrecovery_start", "percentage", "20"),
				),
			},
		},
	})
}

// testAccCheckNsmemrecoveryStartExist is a state-only existence check.
// nsmemrecovery_start is an action-only resource with no GET-by-id endpoint, so
// we only assert Terraform recorded the resource in state with a non-empty ID.
func testAccCheckNsmemrecoveryStartExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No nsmemrecovery_start ID is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		// No NITRO GET-by-id to verify against for an action-only resource.
		return nil
	}
}
