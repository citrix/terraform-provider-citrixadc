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

// NOTE on the traceroute resource:
//   - Models the NITRO POST /traceroute action, which runs a traceroute from the
//     ADC (CLI: "traceroute <host>"). It is a NON-DESTRUCTIVE, ACTION-ONLY
//     (one-shot side-effect) diagnostic resource: Create performs the traceroute,
//     Read is a no-op (preserves state), Update is a no-op (every attribute is
//     RequiresReplace), and Delete is a state-only removal.
//   - There is NO get/add/update/delete endpoint, so the resource CANNOT be
//     verified by reading it back from the ADC, and it has NO datasource.
//   - The Exist check below only verifies that the resource landed in Terraform
//     state with its synthetic ID ("traceroute-config"); it does NOT (and cannot)
//     verify the traceroute side-effect via NITRO.
//   - There is no CheckDestroy: the traceroute action has no inverse on NITRO.
//
// This mirrors the action-only test precedent (single apply step, state-only
// Exist check, no CheckDestroy, TestCheckResourceAttrSet on "id").

// Single apply step: every attribute is RequiresReplace, so there is no in-place
// update to exercise. `host` is "127.0.0.1" (the IPv4 loopback) so the traceroute
// is non-destructive; replace it with a reachable target as needed.
const testAccTraceroute_basic = `
resource "citrixadc_traceroute" "tf_traceroute" {
  host = "127.0.0.1"
  m    = 3
}
`

func TestAccTraceroute_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// No CheckDestroy: the traceroute action has no inverse on NITRO and there
		// is no GET-by-id to confirm absence; Delete is a state-only removal.
		Steps: []resource.TestStep{
			{
				Config: testAccTraceroute_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTracerouteExist("citrixadc_traceroute.tf_traceroute", nil),
					resource.TestCheckResourceAttr("citrixadc_traceroute.tf_traceroute", "host", "127.0.0.1"),
					// "id" is the synthetic state handle "traceroute-config".
					resource.TestCheckResourceAttrSet("citrixadc_traceroute.tf_traceroute", "id"),
				),
			},
		},
	})
}

// testAccCheckTracerouteExist is a state-only existence check.
//
// traceroute is an action-only resource: Read is a no-op and there is no
// GET-by-id endpoint, so we CANNOT verify the traceroute via NITRO. We only
// assert that Terraform recorded the resource in state with a non-empty ID
// (which equals the synthetic "traceroute-config" after a successful POST).
func testAccCheckTracerouteExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No traceroute ID is set")
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
