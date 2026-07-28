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

// NOTE on the nsconfig_convert resource:
//   - Models the NITRO POST /nsconfig?action=convert endpoint (CLI:
//     "convert ns config <configFile>"), which converts a config file into the
//     nitro graph format.
//   - This is an ACTION-ONLY (one-shot side-effect) resource: Create performs the
//     convert action via ActOnResource(service.Nsconfig.Type(), &payload,
//     "convert"), Read is a no-op (preserves state), Update is a no-op (all
//     attributes are RequiresReplace), and Delete is a state-only removal. There
//     is NO get/add/update/delete endpoint, so the resource CANNOT be verified by
//     reading it back from the ADC, and it has NO datasource (Pattern 13 — no
//     NITRO GET).
//   - `configfile` is Required (CLI + NITRO doc both mark it mandatory; Pattern 8).
//     It must point to a config file that already exists on the appliance. The
//     saved running configuration "/nsconfig/ns.conf" is always present on any
//     standalone appliance, so it is the safest self-contained value and needs NO
//     pre-existing setup on the testbed. With `responsefile` omitted, the nitro
//     graph is returned in the API response (discarded by ActOnResource).
//   - `timestamp` is a synthetic Required + RequiresReplace re-run key; the
//     synthetic ID equals the timestamp.
//   - The Exist check below only verifies that the resource landed in Terraform
//     state with its synthetic ID; it does NOT (and cannot) verify the convert
//     side-effect via NITRO.
//   - There is no CheckDestroy: the convert action has no inverse on NITRO and
//     there is no GET-by-id to confirm absence; Delete is a state-only removal.
//
// This mirrors the action-only test precedent in protocolhttpband_clear_test.go /
// nsconfig_clear_test.go (single apply step, state-only Exist check, no
// CheckDestroy). The parent nsconfig action tests are not skip-gated, so no
// ADC_TESTBED gate is applied here.

// Single apply step: all attributes are RequiresReplace, so there is no in-place
// update to exercise. configfile points at the always-present saved config
// "/nsconfig/ns.conf" (no pre-existing configuration required).
const testAccNsconfigConvert_basic = `
resource "citrixadc_nsconfig_convert" "tf_nsconfig_convert" {
  configfile = "/nsconfig/ns.conf"
  timestamp  = "2024-06-01T12:00:00"
}

`

func TestAccNsconfigConvert_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// No CheckDestroy: the convert action has no inverse on NITRO and there is
		// no GET-by-id to confirm absence; Delete is a state-only removal.
		Steps: []resource.TestStep{
			{
				Config: testAccNsconfigConvert_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsconfigConvertExist("citrixadc_nsconfig_convert.tf_nsconfig_convert", nil),
					resource.TestCheckResourceAttr("citrixadc_nsconfig_convert.tf_nsconfig_convert", "configfile", "/nsconfig/ns.conf"),
					resource.TestCheckResourceAttr("citrixadc_nsconfig_convert.tf_nsconfig_convert", "timestamp", "2024-06-01T12:00:00"),
					// "id" is the synthetic state handle (equals the timestamp).
					resource.TestCheckResourceAttrSet("citrixadc_nsconfig_convert.tf_nsconfig_convert", "id"),
				),
			},
		},
	})
}

// testAccCheckNsconfigConvertExist is a state-only existence check.
//
// nsconfig_convert is an action-only resource: Read is a no-op and there is no
// GET-by-id endpoint, so we CANNOT verify the convert via NITRO. We only assert
// that Terraform recorded the resource in state with a non-empty ID (which equals
// the synthetic timestamp after a successful POST ?action=convert).
func testAccCheckNsconfigConvertExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No nsconfig_convert ID is set")
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
