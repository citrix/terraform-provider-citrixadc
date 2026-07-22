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

// NOTE on the systemsignedexereport_disable resource:
//   - Models ONLY the NITRO "disable" action on the "systemsignedexereport"
//     object (POST /nitro/v1/config/systemsignedexereport?action=disable).
//   - This is a ZERO-ATTRIBUTE, ACTION-ONLY resource: Create performs the
//     "disable" action and Read/Update/Delete are no-ops. There is NO
//     get/add/set endpoint, so the resource CANNOT be verified by reading it
//     back and it has NO datasource.
//   - The Exist check below only verifies the resource landed in Terraform state
//     with its synthetic id ("systemsignedexereport_disable"); it cannot verify
//     the side effect via NITRO.
//   - Disabling this feature is non-destructive, so the test is NOT skipped.
//
// This mirrors the action-only test precedent (single apply step, state-only
// Exist check, no CheckDestroy, TestCheckResourceAttrSet on "id").

const testAccSystemsignedexereportDisable_basic = `
resource "citrixadc_systemsignedexereport_disable" "tf_systemsignedexereport_disable" {
}
`

func TestAccSystemsignedexereportDisable_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// No CheckDestroy: there is no GET-by-id to confirm absence; Delete is a
		// no-op (the disable action has no inverse NITRO API on this resource).
		Steps: []resource.TestStep{
			{
				Config: testAccSystemsignedexereportDisable_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSystemsignedexereportDisableExist("citrixadc_systemsignedexereport_disable.tf_systemsignedexereport_disable", nil),
					// "id" is the synthetic state handle "systemsignedexereport_disable".
					resource.TestCheckResourceAttrSet("citrixadc_systemsignedexereport_disable.tf_systemsignedexereport_disable", "id"),
				),
			},
		},
	})
}

// testAccCheckSystemsignedexereportDisableExist is a state-only existence check.
// systemsignedexereport_disable is an action-only resource with no GET-by-id
// endpoint, so we only assert Terraform recorded the resource in state with a
// non-empty ID.
func testAccCheckSystemsignedexereportDisableExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No systemsignedexereport_disable ID is set")
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
