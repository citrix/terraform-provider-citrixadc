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

// NOTE on the ping6 resource:
//   - Models the NITRO POST /ping6 action, which runs an IPv6 ping from the ADC
//     (CLI: "ping6 <host>"). It is a NON-DESTRUCTIVE, ACTION-ONLY (one-shot
//     side-effect) diagnostic resource: Create performs the ping, Read is a no-op
//     (preserves state), Update is a no-op (every attribute is RequiresReplace),
//     and Delete is a state-only removal.
//   - There is NO get/add/update/delete endpoint, so the resource CANNOT be
//     verified by reading it back from the ADC, and it has NO datasource.
//   - The Exist check below only verifies that the resource landed in Terraform
//     state with its synthetic ID ("ping6-config"); it does NOT (and cannot)
//     verify the ping side-effect via NITRO.
//   - There is no CheckDestroy: the ping action has no inverse on NITRO.
//
// This mirrors the action-only test precedent (single apply step, state-only
// Exist check, no CheckDestroy, TestCheckResourceAttrSet on "id").

// Single apply step: every attribute is RequiresReplace, so there is no in-place
// update to exercise. `hostname` is "::1" (the IPv6 loopback) so the ping is
// non-destructive; replace it with a reachable IPv6 target as needed.
const testAccPing6_basic = `
resource "citrixadc_ping6" "tf_ping6" {
  hostname = "::1"
  c        = 1
}
`

func TestAccPing6_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// No CheckDestroy: the ping6 action has no inverse on NITRO and there is
		// no GET-by-id to confirm absence; Delete is a state-only removal.
		Steps: []resource.TestStep{
			{
				Config: testAccPing6_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPing6Exist("citrixadc_ping6.tf_ping6", nil),
					resource.TestCheckResourceAttr("citrixadc_ping6.tf_ping6", "hostname", "::1"),
					// "id" is the synthetic state handle "ping6-config".
					resource.TestCheckResourceAttrSet("citrixadc_ping6.tf_ping6", "id"),
				),
			},
		},
	})
}

// testAccCheckPing6Exist is a state-only existence check.
//
// ping6 is an action-only resource: Read is a no-op and there is no GET-by-id
// endpoint, so we CANNOT verify the ping via NITRO. We only assert that
// Terraform recorded the resource in state with a non-empty ID (which equals the
// synthetic "ping6-config" after a successful POST).
func testAccCheckPing6Exist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ping6 ID is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		// No NITRO GET-by-id to verify against; presence of the synthetic state ID
		// is the only confirmation we can make for an action-only resource.
		return nil
	}
}
