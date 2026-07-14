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

// =============================================================================
// !!! DANGER -- DO NOT RUN THIS TEST IN CI / AUTOMATION !!!
// =============================================================================
//
// This test is intentionally NOT run by CI or any automation harness because
// applying the citrixadc_systemsession resource KILLS administrative sessions
// via POST /nitro/v1/config/systemsession?action=kill.
//
//   - `kill systemsession -all` (all = true) terminates ALL admin sessions,
//     INCLUDING the provider's own NITRO session and every other operator's
//     session. `all = true` MUST NEVER be used in automation -- it would sever
//     the very connection the provider is using and break the test mid-run.
//   - Even killing a SPECIFIC `sid` is destructive: it forcibly logs out that
//     session. A valid sid is also runtime-specific (it only exists while that
//     login is active), so it cannot be hardcoded for unattended runs.
//   - NITRO exposes no inverse "un-kill" action, so the effect cannot be undone.
//
// Run the resource test MANUALLY and ONLY with a KNOWN, DISPOSABLE sid (e.g. a
// throwaway login you created yourself) on an appliance you control. The benign
// read-only datasource test is separately skip-gated below.
// =============================================================================

// NOTE on the systemsession resource:
//   - Models the NITRO POST /systemsession?action=kill endpoint.
//   - ACTION-ONLY (one-shot side-effect) resource: Create performs the kill
//     action, Read is a no-op (preserves state), Update is a no-op, and Delete
//     is a state-only removal. There is NO add/update/delete endpoint, so the
//     killed session CANNOT be re-resolved/verified by reading it back.
//   - Exactly one of `sid` or `all` must be supplied (enforced by ValidateConfig);
//     both are RequiresReplace. The synthetic ID is the sid value, or "all".
//   - The Exist check below only verifies that the resource landed in Terraform
//     state with a non-empty synthetic ID; it does NOT (and cannot) verify the
//     kill side-effect via NITRO.
//   - There is no CheckDestroy: the kill action has no inverse on NITRO, and
//     there is no GET-by-id to confirm absence; Delete is a state-only removal.
//
// This mirrors the action-only test precedent in aaasession_test.go and
// clusterfiles_test.go (single apply step, state-only Exist check, no
// CheckDestroy, TestCheckResourceAttrSet on "id").

// Single apply step: sid is RequiresReplace, so there is no in-place update to
// exercise.
//
// IMPORTANT: a REAL, runtime-specific session id MUST be substituted for the
// TODO_PLACEHOLDER below before running this manually -- there is no generic
// valid sid. Use `all = true` instead is FORBIDDEN: it would kill the provider's
// own NITRO session (and everyone else's) and is never permitted in automation.
const testAccSystemsession_basic_step1 = `
resource "citrixadc_systemsession" "tf_systemsession" {
  sid = 13555
}

`

func TestAccSystemsession_basic(t *testing.T) {
	t.Skip("Requires a real, runtime-specific live session sid;")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// No CheckDestroy: the kill action has no inverse on NITRO and there is
		// no GET-by-id to confirm absence; Delete is a state-only removal.
		Steps: []resource.TestStep{
			{
				Config: testAccSystemsession_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSystemsessionExist("citrixadc_systemsession.tf_systemsession", nil),
					// "id" is the synthetic state handle (the sid value, or "all").
					resource.TestCheckResourceAttrSet("citrixadc_systemsession.tf_systemsession", "id"),
				),
			},
		},
	})
}

// testAccCheckSystemsessionExist is a state-only existence check.
//
// systemsession is an action-only resource: Read is a no-op and there is no
// GET-by-id endpoint that re-resolves a killed session, so we CANNOT verify the
// kill via NITRO. We only assert that Terraform recorded the resource in state
// with a non-empty ID (the synthetic sid value, or "all", after a successful
// POST ?action=kill). This mirrors testAccCheckAaasessionExist.
func testAccCheckSystemsessionExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No systemsession ID is set")
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

// Datasource: systemsession HAS a get-by-sid datasource that returns read-only
// session fields (username, logintime, clientipaddress, ...). Reading a session
// is BENIGN (no kill), but it requires a REAL, runtime-specific sid that exists
// only while that login is active -- it cannot be hardcoded for unattended runs.
// Following the runtime/session datasource convention (see
// TestAccAaasessionDataSource_basic), this test is skip-gated. Remove the t.Skip
// and substitute a known live sid for TODO_PLACEHOLDER to run it manually.
const testAccSystemsessionDataSource_basic = `

data "citrixadc_systemsession" "tf_systemsession" {
  sid = 13555
}
`

func TestAccSystemsessionDataSource_basic(t *testing.T) {
	t.Skip("Requires a real, runtime-specific live session sid;")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSystemsessionDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Keep assertions minimal: session output fields depend on the
					// live session that matched the supplied sid. The synthetic
					// "id" is the only stable value we can assert here.
					resource.TestCheckResourceAttrSet("data.citrixadc_systemsession.tf_systemsession", "id"),
				),
			},
		},
	})
}
