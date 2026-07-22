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

// NOTE on the shutdown resource:
//   - Models the NITRO POST /shutdown endpoint (CLI: "shutdown"), which shuts
//     down the Citrix ADC appliance.
//   - This is a ZERO-ATTRIBUTE, ACTION-ONLY resource: Create performs the
//     shutdown action, Read/Update are no-ops, and Delete is a state-only
//     removal. There is NO get/add/set/delete endpoint, so the resource CANNOT
//     be verified by reading it back and it has NO datasource (Pattern 13).
//   - The Exist check below only verifies the resource landed in Terraform state
//     with its synthetic id ("shutdown-config"); it cannot verify the side
//     effect via NITRO.
//
// !!! DESTRUCTIVE !!!
// Applying this resource SHUTS DOWN the appliance, which will make the testbed
// unreachable and fail the rest of the acceptance run. The test is therefore
// SKIPPED by default. Remove the t.Skip line below only when you intend to
// shut down a disposable appliance on purpose.
//
// This mirrors the action-only test precedent (single apply step, state-only
// Exist check, no CheckDestroy, TestCheckResourceAttrSet on "id").

const testAccShutdown_basic = `
resource "citrixadc_shutdown" "tf_shutdown" {
}
`

func TestAccShutdown_basic(t *testing.T) {
	t.Skip("Will shutdown target ADC testbed.")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// No CheckDestroy: the shutdown action has no inverse on NITRO and there
		// is no GET-by-id to confirm absence; Delete is a state-only removal.
		Steps: []resource.TestStep{
			{
				Config: testAccShutdown_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckShutdownExist("citrixadc_shutdown.tf_shutdown", nil),
					// "id" is the synthetic state handle "shutdown-config".
					resource.TestCheckResourceAttrSet("citrixadc_shutdown.tf_shutdown", "id"),
				),
			},
		},
	})
}

// testAccCheckShutdownExist is a state-only existence check. shutdown is an
// action-only resource with no GET-by-id endpoint, so we only assert Terraform
// recorded the resource in state with a non-empty ID.
func testAccCheckShutdownExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No shutdown ID is set")
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
