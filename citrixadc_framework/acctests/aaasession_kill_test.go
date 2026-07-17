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

// NOTE on the aaasession_kill resource:
//   - Models the NITRO POST /aaasession?action=kill endpoint.
//   - This is an ACTION-ONLY (one-shot side-effect) resource: Create performs
//     the kill action, Read is a no-op (preserves state), Update is a no-op,
//     and Delete is a state-only removal. There is no GET-by-id endpoint that
//     can re-resolve "this killed session", so the resource CANNOT be verified
//     by reading it back from the ADC.
//   - All attributes (username, groupname, iip, netmask, sessionkey, all,
//     nodeid) are OPTIONAL kill filters; none is Required. nodeid is a GET-only
//     cluster filter and is intentionally excluded from the kill payload.
//   - `all = true` kills all active AAA-TM/VPN sessions and is the simplest
//     self-contained config: it needs no pre-existing session on the testbed.
//   - The Exist check below only verifies that the resource landed in Terraform
//     state with its synthetic ID ("aaasession_kill"); it does NOT (and cannot)
//     verify the kill side-effect via NITRO.
//   - There is no CheckDestroy: the kill action has no inverse on NITRO, and
//     there is no GET-by-id to confirm absence; Delete is a state-only removal.
//
// This mirrors the action-only test precedent in appfwarchive_export_test.go
// (state-only Exist check, no CheckDestroy, TestCheckResourceAttrSet on "id"),
// adapted for the kill action and its optional filter attributes.

// Single apply step: all attributes are RequiresReplace, so there is no
// in-place update to exercise. The kill action is self-contained via all=true.
const testAccAaasessionKill_basic = `
resource "citrixadc_aaasession_kill" "tf_aaasession_kill" {
  all = true
}

`

func TestAccAaasessionKill_basic(t *testing.T) {
	t.Skip("TODO: Requires review - the kill action terminates live sessions and can disconnect the test's own management/NITRO session; unsafe on a shared testbed")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// No CheckDestroy: the kill action has no inverse on NITRO and there is
		// no GET-by-id to confirm absence; Delete is a state-only removal.
		Steps: []resource.TestStep{
			{
				Config: testAccAaasessionKill_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaasessionKillExist("citrixadc_aaasession_kill.tf_aaasession_kill", nil),
					// "id" is the synthetic state handle "aaasession_kill".
					resource.TestCheckResourceAttrSet("citrixadc_aaasession_kill.tf_aaasession_kill", "id"),
					// Assert only the filter actually set in HCL.
					resource.TestCheckResourceAttr("citrixadc_aaasession_kill.tf_aaasession_kill", "all", "true"),
				),
			},
		},
	})
}

// testAccCheckAaasessionKillExist is a state-only existence check.
//
// aaasession_kill is an action-only resource: Read is a no-op and there is no
// GET-by-id endpoint, so we CANNOT verify the kill via NITRO. We only assert
// that Terraform recorded the resource in state with a non-empty ID (which
// equals the synthetic "aaasession_kill" after a successful POST ?action=kill).
// This mirrors testAccCheckAppfwarchiveExportExist.
func testAccCheckAaasessionKillExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No aaasession_kill ID is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		// No NITRO GET-by-id to verify against; presence of the synthetic state
		// ID is the only confirmation we can make for an action-only resource.
		return nil
	}
}
