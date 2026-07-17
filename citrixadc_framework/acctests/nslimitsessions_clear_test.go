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

// NOTE on the nslimitsessions_clear resource:
//   - Models the NITRO POST /nslimitsessions?action=clear endpoint, which clears
//     the rate-limit sessions for a given rate-limit identifier.
//   - This is an ACTION-ONLY (one-shot side-effect) resource: Create performs the
//     clear action via ActOnResource(service.Nslimitsessions.Type(), &payload,
//     "clear"), Read is a no-op (preserves state), Update is a no-op (the single
//     input attribute is RequiresReplace), and Delete is a state-only removal.
//     There is NO get/add/update/delete endpoint for the clear action, so the
//     resource CANNOT be verified by reading it back from the ADC, and it has no
//     datasource (the datasource remains the unchanged citrixadc_nslimitsessions).
//   - The single input attribute `limitidentifier` is Required and RequiresReplace.
//     To be MEANINGFUL, it must reference a REAL nslimitidentifier configured on the
//     appliance; the config below creates one so the test is self-contained.
//   - The Exist check below only verifies that the resource landed in Terraform
//     state with its synthetic ID ("nslimitsessions_clear"); it does NOT (and
//     cannot) verify the clear side-effect via NITRO.
//   - There is no CheckDestroy: the clear action has no inverse on NITRO, and there
//     is no GET-by-id to confirm absence; Delete is a state-only removal.
//
// This mirrors the action-only test precedent in protocolhttpband_clear_test.go /
// endpointinfo_clear_test.go (single apply step, state-only Exist check, no
// CheckDestroy), adapted for the clear action. The original nslimitsessions action
// test was not skip-gated, so this runnable single-apply state-only test is kept.
//
// ENVIRONMENT: clear works on a standalone testbed; no ADC_TESTBED gate is needed.

// Single apply step: the `limitidentifier` attribute is RequiresReplace, so there
// is no in-place update to exercise. A real nslimitidentifier is created first so
// the clear action targets an existing rate-limit identifier.
const testAccNslimitsessionsClear_basic = `
	resource "citrixadc_nslimitidentifier" "tf_nslimitidentifier" {
		limitidentifier  = "tf_nslimitidentifier"
		threshold        = 1
		timeslice        = 1000
		limittype        = "BURSTY"
		mode             = "REQUEST_RATE"
		maxbandwidth     = 0
		trapsintimeslice = 1
	}
	resource "citrixadc_nslimitsessions_clear" "tf_nslimitsessions_clear" {
		limitidentifier = citrixadc_nslimitidentifier.tf_nslimitidentifier.limitidentifier
	}

`

func TestAccNslimitsessionsClear_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// No CheckDestroy: the clear action has no inverse on NITRO and there is
		// no GET-by-id to confirm absence; Delete is a state-only removal.
		Steps: []resource.TestStep{
			{
				Config: testAccNslimitsessionsClear_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNslimitsessionsClearExist("citrixadc_nslimitsessions_clear.tf_nslimitsessions_clear", nil),
					resource.TestCheckResourceAttr("citrixadc_nslimitsessions_clear.tf_nslimitsessions_clear", "limitidentifier", "tf_nslimitidentifier"),
					// "id" is the synthetic state handle "nslimitsessions_clear".
					resource.TestCheckResourceAttrSet("citrixadc_nslimitsessions_clear.tf_nslimitsessions_clear", "id"),
				),
			},
		},
	})
}

// testAccCheckNslimitsessionsClearExist is a state-only existence check.
//
// nslimitsessions_clear is an action-only resource: Read is a no-op and there is
// no GET-by-id endpoint, so we CANNOT verify the clear via NITRO. We only assert
// that Terraform recorded the resource in state with a non-empty ID (which equals
// the synthetic "nslimitsessions_clear" after a successful POST ?action=clear).
func testAccCheckNslimitsessionsClearExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No nslimitsessions_clear ID is set")
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
