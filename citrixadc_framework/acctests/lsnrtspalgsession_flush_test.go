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

// NOTE on the lsnrtspalgsession_flush resource:
//   - Models the NITRO POST /lsnrtspalgsession?action=flush endpoint, which
//     flushes (clears) a transient runtime RTSP-ALG session for Large Scale NAT.
//   - This is an ACTION-ONLY (one-shot side-effect) resource: Create performs
//     the flush action via ActOnResource(service.Lsnrtspalgsession.Type(),
//     payload, "flush"), Read is a no-op (preserves state), Update is a no-op,
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
//     state with its synthetic ID ("lsnrtspalgsession_flush"); it does NOT (and
//     cannot) verify the flush side-effect via NITRO.
//   - There is no CheckDestroy: the flush action has no inverse on NITRO, the
//     runtime session is transient, and Delete is a state-only removal.
//
// This mirrors the action-only test precedent in protocolhttpband_clear_test.go /
// rnatsession_test.go (single apply step, state-only Exist check, no CheckDestroy,
// TestCheckResourceAttrSet on "id"), adapted for the flush action.

// Single apply step: sessionid is RequiresReplace, so there is no in-place
// update to exercise. Flushing a (likely absent) session is a harmless no-op.
const testAccLsnrtspalgsessionFlush_basic = `
resource "citrixadc_lsnrtspalgsession_flush" "tf_lsnrtspalgsession_flush" {
  sessionid = "10.102.43.13:6789"
}

`

func TestAccLsnrtspalgsessionFlush_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// No CheckDestroy: the flush action has no inverse on NITRO, the runtime
		// session is transient, and Delete is a state-only removal.
		Steps: []resource.TestStep{
			{
				Config: testAccLsnrtspalgsessionFlush_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsnrtspalgsessionFlushExist("citrixadc_lsnrtspalgsession_flush.tf_lsnrtspalgsession_flush", nil),
					// "id" is the synthetic state handle ("lsnrtspalgsession_flush").
					resource.TestCheckResourceAttrSet("citrixadc_lsnrtspalgsession_flush.tf_lsnrtspalgsession_flush", "id"),
					resource.TestCheckResourceAttr("citrixadc_lsnrtspalgsession_flush.tf_lsnrtspalgsession_flush", "id", "lsnrtspalgsession_flush"),
					// Assert only the attribute actually set in HCL.
					resource.TestCheckResourceAttr("citrixadc_lsnrtspalgsession_flush.tf_lsnrtspalgsession_flush", "sessionid", "10.102.43.13:6789"),
				),
			},
		},
	})
}

// testAccCheckLsnrtspalgsessionFlushExist is a state-only existence check.
//
// lsnrtspalgsession_flush is an action-only resource: Read is a no-op and there
// is no stable GET-by-id to re-resolve a flushed (transient) runtime session, so
// we CANNOT verify the flush via NITRO. We only assert that Terraform recorded
// the resource in state with a non-empty ID (which equals the synthetic
// "lsnrtspalgsession_flush" after a successful POST ?action=flush). This mirrors
// testAccCheckProtocolhttpbandClearExist.
func testAccCheckLsnrtspalgsessionFlushExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lsnrtspalgsession_flush ID is set")
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
