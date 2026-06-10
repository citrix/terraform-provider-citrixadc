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

// NOTE on the lsnrtspalgsession resource:
//   - Models the NITRO POST /lsnrtspalgsession?action=flush endpoint, which
//     flushes (clears) a transient runtime RTSP-ALG session for Large Scale NAT.
//   - This is an ACTION-ONLY (one-shot side-effect) resource: Create performs
//     the flush action, Read is a no-op (preserves state), Update is a no-op,
//     and Delete is a state-only removal. There is NO add/update/delete
//     endpoint, so the resource CANNOT be verified by reading it back from the
//     ADC. (NITRO does expose get(all)/get(by sessionid)/count, which backs the
//     read-only datasource below, but there is no stable GET-by-id to re-resolve
//     "this flush" — the runtime session is transient and may already be gone.)
//   - `sessionid` is the single REQUIRED attribute (RequiresReplace); it
//     identifies the RTSP ALG session to flush. `nodeid` is OPTIONAL and is a
//     GET-only cluster filter (excluded from the flush payload), so the basic
//     test omits it (standalone testbed).
//   - Flushing a non-existent session is a harmless no-op on the ADC, so the
//     basic test uses a plausible sessionid that need not pre-exist.
//   - The Exist check below only verifies that the resource landed in Terraform
//     state with its synthetic ID (the sessionid value); it does NOT (and cannot)
//     verify the flush side-effect via NITRO.
//   - There is no CheckDestroy: the flush action has no inverse on NITRO, the
//     runtime session is transient, and Delete is a state-only removal.
//
// This mirrors the action-only test precedent in aaasession_test.go,
// clusterfiles_test.go and clusterpropstatus_test.go (single apply step,
// state-only Exist check, no CheckDestroy, TestCheckResourceAttrSet on "id"),
// adapted for the flush action.

// Single apply step: sessionid is RequiresReplace, so there is no in-place
// update to exercise. Flushing a (likely absent) session is a harmless no-op.
const testAccLsnrtspalgsession_basic = `
resource "citrixadc_lsnrtspalgsession" "tf_lsnrtspalgsession" {
  sessionid = "10.102.43.13:6789"
}

`

func TestAccLsnrtspalgsession_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// No CheckDestroy: the flush action has no inverse on NITRO, the runtime
		// session is transient, and Delete is a state-only removal.
		Steps: []resource.TestStep{
			{
				Config: testAccLsnrtspalgsession_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsnrtspalgsessionExist("citrixadc_lsnrtspalgsession.tf_lsnrtspalgsession", nil),
					// "id" is the synthetic state handle (the sessionid value).
					resource.TestCheckResourceAttrSet("citrixadc_lsnrtspalgsession.tf_lsnrtspalgsession", "id"),
					// Assert only the attribute actually set in HCL.
					resource.TestCheckResourceAttr("citrixadc_lsnrtspalgsession.tf_lsnrtspalgsession", "sessionid", "10.102.43.13:6789"),
				),
			},
		},
	})
}

// testAccCheckLsnrtspalgsessionExist is a state-only existence check.
//
// lsnrtspalgsession is an action-only resource: Read is a no-op and there is no
// stable GET-by-id to re-resolve a flushed (transient) runtime session, so we
// CANNOT verify the flush via NITRO. We only assert that Terraform recorded the
// resource in state with a non-empty ID (which equals the synthetic sessionid
// value after a successful POST ?action=flush). This mirrors
// testAccCheckAaasessionExist / testAccCheckClusterpropstatusExist.
func testAccCheckLsnrtspalgsessionExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lsnrtspalgsession ID is set")
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

// Datasource: lsnrtspalgsession HAS a get(by sessionid)-backed datasource. Its
// Read calls FindResource(..., sessionid), which errors out if no RTSP-ALG
// session with that sessionid currently exists on the ADC. RTSP-ALG sessions
// are transient runtime objects (and the resource above flushes them), so on a
// standalone testbed there is almost never a live session to look up. This test
// is therefore skipped by default; remove the t.Skip and set sessionid to a
// known live RTSP-ALG session before running. This mirrors the gated datasource
// precedent in aaasession_test.go.
const testAccLsnrtspalgsessionDataSource_basic = `

data "citrixadc_lsnrtspalgsession" "tf_lsnrtspalgsession" {
  sessionid = "10.102.43.13:6789"
}
`

func TestAccLsnrtspalgsessionDataSource_basic(t *testing.T) {
	t.Skip("Requires a live RTSP-ALG session with the given sessionid on the ADC; the get(by sessionid) datasource errors when the session is absent (RTSP-ALG sessions are transient).")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccLsnrtspalgsessionDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Keep assertions minimal: the live session's returned fields
					// depend on runtime state. The synthetic "id" (= sessionid) and
					// the looked-up sessionid are the stable values we can assert.
					resource.TestCheckResourceAttrSet("data.citrixadc_lsnrtspalgsession.tf_lsnrtspalgsession", "id"),
					resource.TestCheckResourceAttr("data.citrixadc_lsnrtspalgsession.tf_lsnrtspalgsession", "sessionid", "10.102.43.13:6789"),
				),
			},
		},
	})
}
