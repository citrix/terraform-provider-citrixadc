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

// vpnpcoipconnection_kill is an ACTION-ONLY resource: Create performs a POST
// ?action=kill to terminate active PCoIP connections. There is no GET-backed
// managed object, no update/set endpoint, and no delete endpoint (Delete is a
// state-only removal).
//
// Consequences for testing (Test-generation pitfall #4):
//   - The Exist check below does NOT call client.FindResource(...) — there is no
//     stable NITRO row to look up for a killed connection. It only verifies that
//     the synthetic ID "vpnpcoipconnection_kill" was set in Terraform state.
//   - No CheckDestroy is generated — there is no NITRO delete endpoint to verify.
//   - No datasource test is generated — the resource has no stable GET-backed
//     object.
//   - A single apply step is sufficient: every attribute is RequiresReplace, so
//     there is no in-place update path to exercise. On an idle ADC with no live
//     PCoIP connections, the kill is effectively a no-op; we only assert that the
//     apply succeeds and that the resource is recorded in state with its ID set.

const testAccVpnpcoipconnectionKill_basic_step1 = `
resource "citrixadc_vpnpcoipconnection_kill" "tf_vpnpcoipconnection_kill" {
  all = true
}

`

func TestAccVpnpcoipconnectionKill_basic(t *testing.T) {
	t.Skip("TODO: Requires review - the kill action terminates live PCoIP connections; disruptive on a shared testbed")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// No CheckDestroy: action-only resource has no NITRO delete endpoint.
		Steps: []resource.TestStep{
			{
				Config: testAccVpnpcoipconnectionKill_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnpcoipconnectionKillExist("citrixadc_vpnpcoipconnection_kill.tf_vpnpcoipconnection_kill", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnpcoipconnection_kill.tf_vpnpcoipconnection_kill", "all", "true"),
				),
			},
		},
	})
}

// testAccCheckVpnpcoipconnectionKillExist verifies the resource exists in Terraform state.
//
// IMPORTANT: This check intentionally does NOT call client.FindResource(...).
// vpnpcoipconnection_kill is action-only (kill); the killed connection is not a
// persistent, queryable managed object, so there is no NITRO GET that can
// re-resolve "this" record. We therefore only assert that the synthetic ID
// "vpnpcoipconnection_kill" was set in state, which confirms the kill action's
// Create path completed successfully.
func testAccCheckVpnpcoipconnectionKillExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No vpnpcoipconnection_kill ID is set")
		}

		if rs.Primary.ID != "vpnpcoipconnection_kill" {
			return fmt.Errorf("Unexpected synthetic ID: got %q, want %q", rs.Primary.ID, "vpnpcoipconnection_kill")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		// No NITRO call: action-only resource (kill) has no GET-backed object.
		return nil
	}
}
