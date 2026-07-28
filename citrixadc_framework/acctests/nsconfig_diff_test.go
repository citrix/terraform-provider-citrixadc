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

// NOTE on the nsconfig_diff resource:
//   - Models the NITRO POST /nsconfig?action=diff endpoint (CLI: "diff ns config").
//   - This is an ACTION-ONLY (one-shot side-effect) resource: Create performs the
//     diff action via ActOnResource(service.Nsconfig.Type(), &payload, "diff"),
//     Read is a no-op (preserves state), Update is a no-op (all attributes are
//     RequiresReplace), and Delete is a state-only removal. There is NO
//     get/add/update/delete endpoint, so the resource CANNOT be verified by reading
//     it back from the ADC, and it has NO datasource (Pattern 13 — no NITRO GET).
//   - All action attributes (config1, config2, outtype, template,
//     ignoredevicespecific) are Optional. With none of config1/config2 supplied,
//     the diff compares the running config against the saved config, which is
//     always available on any standalone appliance and needs NO pre-existing
//     configuration on the testbed.
//   - `timestamp` is a synthetic Required + RequiresReplace re-run key; the
//     synthetic ID equals the timestamp.
//   - The Exist check below only verifies that the resource landed in Terraform
//     state with its synthetic ID; it does NOT (and cannot) verify the diff
//     side-effect via NITRO.
//   - There is no CheckDestroy: the diff action has no inverse on NITRO and there
//     is no GET-by-id to confirm absence; Delete is a state-only removal.
//
// This mirrors the action-only test precedent in protocolhttpband_clear_test.go /
// nsconfig_clear_test.go (single apply step, state-only Exist check, no
// CheckDestroy). The parent nsconfig action tests are not skip-gated, so no
// ADC_TESTBED gate is applied here.

// Single apply step: all attributes are RequiresReplace, so there is no in-place
// update to exercise. Only `timestamp` is set, which runs "diff ns config"
// comparing running vs saved config (simplest self-contained form, no
// pre-existing configuration required).
const testAccNsconfigDiff_basic = `
resource "citrixadc_nsconfig_diff" "tf_nsconfig_diff" {
  timestamp = "2024-06-01T12:00:00"
}

`

func TestAccNsconfigDiff_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// No CheckDestroy: the diff action has no inverse on NITRO and there is no
		// GET-by-id to confirm absence; Delete is a state-only removal.
		Steps: []resource.TestStep{
			{
				Config: testAccNsconfigDiff_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsconfigDiffExist("citrixadc_nsconfig_diff.tf_nsconfig_diff", nil),
					resource.TestCheckResourceAttr("citrixadc_nsconfig_diff.tf_nsconfig_diff", "timestamp", "2024-06-01T12:00:00"),
					// "id" is the synthetic state handle (equals the timestamp).
					resource.TestCheckResourceAttrSet("citrixadc_nsconfig_diff.tf_nsconfig_diff", "id"),
				),
			},
		},
	})
}

// testAccCheckNsconfigDiffExist is a state-only existence check.
//
// nsconfig_diff is an action-only resource: Read is a no-op and there is no
// GET-by-id endpoint, so we CANNOT verify the diff via NITRO. We only assert that
// Terraform recorded the resource in state with a non-empty ID (which equals the
// synthetic timestamp after a successful POST ?action=diff).
func testAccCheckNsconfigDiffExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No nsconfig_diff ID is set")
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
