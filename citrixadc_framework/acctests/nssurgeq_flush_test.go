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

// NOTE on the nssurgeq_flush resource:
//   - Models the NITRO POST /nssurgeq?action=flush endpoint, which flushes the
//     surge queue (SurgeQ). With no arguments it performs a system-wide flush;
//     name/servername/port narrow the flush to a specific entity.
//   - This is an ACTION-ONLY (one-shot side-effect) resource: Create performs the
//     "flush" action, Read is a no-op (preserves state), Update is a no-op (all
//     attributes are RequiresReplace), and Delete is a state-only removal. There
//     is NO add/get/delete endpoint, so the resource CANNOT be verified by reading
//     it back from the ADC.
//   - All attributes are OPTIONAL: `name` (vserver/service/servicegroup),
//     `servername` (service-group member; requires name), and `port` (requires
//     servername). The simplest self-contained config is an empty resource block,
//     which performs a system-wide flush and avoids depending on any participating
//     entity existing on the testbed.
//   - The Exist check below only verifies that the resource landed in Terraform
//     state with its synthetic ID ("nssurgeq_flush"); it does NOT (and cannot)
//     verify the flush side-effect via NITRO.
//   - There is no CheckDestroy: the flush action has no inverse on NITRO, and
//     there is no GET-by-id to confirm absence; Delete is a state-only removal.
//
// This mirrors the action-only test precedent in clusterpropstatus_test.go and
// aaasession_test.go (single apply step, state-only Exist check, no CheckDestroy,
// TestCheckResourceAttrSet on "id"), adapted for the flush action.

// Single apply step: an empty resource block performs a system-wide flush, the
// simplest valid form with no dependency on a participating entity.
const testAccNssurgeqFlush_basic = `
resource "citrixadc_nssurgeq_flush" "tf_nssurgeq" {
}

`

func TestAccNssurgeqFlush_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// No CheckDestroy: the flush action has no inverse on NITRO and there is
		// no GET-by-id to confirm absence; Delete is a state-only removal.
		Steps: []resource.TestStep{
			{
				Config: testAccNssurgeqFlush_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNssurgeqFlushExist("citrixadc_nssurgeq_flush.tf_nssurgeq", nil),
					// "id" is the synthetic state handle "nssurgeq_flush".
					resource.TestCheckResourceAttrSet("citrixadc_nssurgeq_flush.tf_nssurgeq", "id"),
				),
			},
		},
	})
}

// testAccCheckNssurgeqFlushExist is a state-only existence check.
//
// nssurgeq_flush is an action-only resource: Read is a no-op and there is no
// GET-by-id endpoint, so we CANNOT verify the flush via NITRO. We only assert
// that Terraform recorded the resource in state with a non-empty ID (which
// equals the synthetic "nssurgeq_flush" after a successful POST ?action=flush).
func testAccCheckNssurgeqFlushExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No nssurgeq_flush ID is set")
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
