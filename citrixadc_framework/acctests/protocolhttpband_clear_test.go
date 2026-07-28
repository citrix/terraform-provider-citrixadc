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

// NOTE on the protocolhttpband_clear resource:
//   - Models the NITRO POST /protocolhttpband?action=clear endpoint, which clears
//     HTTP band statistics (CLI: "clear protocol httpband -type REQUEST").
//   - This is an ACTION-ONLY (one-shot side-effect) resource: Create performs the
//     clear action via ActOnResource(service.Protocolhttpband.Type(), &payload,
//     "clear"), Read is a no-op (preserves state), Update is a no-op (the single
//     attribute is RequiresReplace), and Delete is a state-only removal. There is
//     NO get/add/update/delete endpoint, so the resource CANNOT be verified by
//     reading it back from the ADC, and it has NO datasource (removed per
//     Pattern 13 — no NITRO GET endpoint).
//   - The single attribute `type` is Required and RequiresReplace; its allowed
//     values are REQUEST | RESPONSE | MQTT_JUMBO_REQ. Clearing the REQUEST band
//     statistics is the simplest self-contained form and needs no pre-existing
//     configuration on the testbed (clearing empty statistics is a no-op success).
//   - The Exist check below only verifies that the resource landed in Terraform
//     state with its synthetic ID ("protocolhttpband_clear-REQUEST"); it does NOT
//     (and cannot) verify the clear side-effect via NITRO.
//   - There is no CheckDestroy: the clear action has no inverse on NITRO, and
//     there is no GET-by-id to confirm absence; Delete is a state-only removal.
//
// This mirrors the action-only test precedent in dnsproxyrecords_test.go /
// rnatsession_test.go (single apply step, state-only Exist check, no CheckDestroy,
// TestCheckResourceAttrSet on "id"), adapted for the clear action.
//
// ENVIRONMENT: clear works on a standalone testbed; no ADC_TESTBED gate is needed.

// Single apply step: the `type` attribute is RequiresReplace, so there is no
// in-place update to exercise. type = "REQUEST" clears the HTTP request band
// statistics (simplest self-contained form, no pre-existing configuration
// required).
const testAccProtocolhttpbandClear_basic = `
resource "citrixadc_protocolhttpband_clear" "tf_protocolhttpband_clear" {
  type = "REQUEST"
}

`

func TestAccProtocolhttpbandClear_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// No CheckDestroy: the clear action has no inverse on NITRO and there is
		// no GET-by-id to confirm absence; Delete is a state-only removal.
		Steps: []resource.TestStep{
			{
				Config: testAccProtocolhttpbandClear_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckProtocolhttpbandClearExist("citrixadc_protocolhttpband_clear.tf_protocolhttpband_clear", nil),
					resource.TestCheckResourceAttr("citrixadc_protocolhttpband_clear.tf_protocolhttpband_clear", "type", "REQUEST"),
					// "id" is the synthetic state handle "protocolhttpband_clear-REQUEST".
					resource.TestCheckResourceAttrSet("citrixadc_protocolhttpband_clear.tf_protocolhttpband_clear", "id"),
				),
			},
		},
	})
}

// testAccCheckProtocolhttpbandClearExist is a state-only existence check.
//
// protocolhttpband_clear is an action-only resource: Read is a no-op and there is
// no GET-by-id endpoint, so we CANNOT verify the clear via NITRO. We only assert
// that Terraform recorded the resource in state with a non-empty ID (which equals
// the synthetic "protocolhttpband_clear-REQUEST" after a successful POST
// ?action=clear). This mirrors testAccCheckRnatsessionFlushExist.
func testAccCheckProtocolhttpbandClearExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No protocolhttpband_clear ID is set")
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
