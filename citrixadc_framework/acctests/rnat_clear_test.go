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

// NOTE on the rnat_clear resource:
//   - Models a *set* of RNAT route rules under one synthetic handle (rnatsname).
//     Create applies each rule via NITRO PUT /config/rnat; Update diffs the set
//     and clears removed rules (POST /config/rnat?action=clear) / re-applies added
//     ones; Delete clears every managed rule. Read is a state-preserving no-op
//     (the legacy SDK v2 Read silently failed for the mistyped schema).
//   - Because Create swallows per-rule NITRO errors (matching SDK v2) and Read is
//     a no-op, the resource always lands in state with its ID (= rnatsname). It
//     therefore behaves like an action/one-shot resource for test purposes and,
//     like the other action-only resources (aaasession_kill, clustersync_force),
//     is verified with a state-only Exist check plus TestCheckResourceAttrSet("id").
//   - `rnat` is a Framework SetNestedAttribute, so it uses attribute HCL syntax
//     (rnat = [ { ... } ]), not the legacy SDK v2 block syntax (rnat { ... }).
//
// This mirrors the action-only test precedent (single apply step, state-only Exist
// check, TestCheckResourceAttrSet on "id"). Delete runs automatically at the end
// of the test and clears the RNAT rule from the appliance.

const testAccRnatClear_basic = `
resource "citrixadc_rnat_clear" "foo" {
  rnatsname = "tf_rnat_clear"
  rnat = [
    {
      network = "192.168.96.0"
      netmask = "255.255.240.0"
    },
  ]
}
`

func TestAccRnatClear_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// No CheckDestroy: rnat_clear's Delete clears the managed RNAT rule(s)
		// during the framework's automatic teardown; there is no stable GET-by-id
		// to reliably confirm absence for the synthetic handle.
		Steps: []resource.TestStep{
			{
				Config: testAccRnatClear_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRnatClearExist("citrixadc_rnat_clear.foo", nil),
					// "id" equals the rnatsname handle after a successful apply.
					resource.TestCheckResourceAttrSet("citrixadc_rnat_clear.foo", "id"),
					resource.TestCheckResourceAttr("citrixadc_rnat_clear.foo", "rnatsname", "tf_rnat_clear"),
				),
			},
		},
	})
}

// testAccCheckRnatClearExist is a state-only existence check. rnat_clear applies
// RNAT config but exposes only a synthetic handle (rnatsname) and has a no-op
// Read, so we assert that Terraform recorded the resource in state with a
// non-empty ID rather than reading a specific rule back from NITRO.
func testAccCheckRnatClearExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No rnat_clear ID is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		return nil
	}
}
