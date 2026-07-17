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

// NOTE on the endpointinfo_clear resource:
//   - Models the NITRO POST /endpointinfo?action=clear endpoint, which clears
//     endpoint information for the given endpoint kind.
//   - This is an ACTION-ONLY (one-shot side-effect) resource: Create performs the
//     clear action via ActOnResource(service.Endpointinfo.Type(), &payload,
//     "clear"), Read is a no-op (preserves state), Update is a no-op (the single
//     attribute is RequiresReplace), and Delete is a state-only removal. There is
//     NO get/add/update/delete endpoint for the clear action, so the resource
//     CANNOT be verified by reading it back from the ADC, and it has no datasource.
//   - The single input attribute `endpointkind` is Optional and RequiresReplace;
//     its only allowed value is IP. Clearing IP endpoint info is self-contained and
//     needs no pre-existing configuration on the testbed (clearing empty endpoint
//     info is a no-op success).
//   - The Exist check below only verifies that the resource landed in Terraform
//     state with its synthetic ID ("endpointinfo_clear-IP"); it does NOT (and
//     cannot) verify the clear side-effect via NITRO.
//   - There is no CheckDestroy: the clear action has no inverse on NITRO, and there
//     is no GET-by-id to confirm absence; Delete is a state-only removal.
//
// This mirrors the action-only test precedent in protocolhttpband_clear_test.go
// (single apply step, state-only Exist check, no CheckDestroy).
//
// ENVIRONMENT: the parent endpointinfo test is not skip-gated (works on a
// standalone testbed), so no ADC_TESTBED gate is needed here either.

// Single apply step: the `endpointkind` attribute is RequiresReplace, so there is
// no in-place update to exercise. endpointkind = "IP" clears the IP endpoint info
// (the only supported kind; no pre-existing configuration required).
const testAccEndpointinfoClear_basic = `
resource "citrixadc_endpointinfo_clear" "tf_endpointinfo_clear" {
  endpointkind = "IP"
}

`

func TestAccEndpointinfoClear_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// No CheckDestroy: the clear action has no inverse on NITRO and there is
		// no GET-by-id to confirm absence; Delete is a state-only removal.
		Steps: []resource.TestStep{
			{
				Config: testAccEndpointinfoClear_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckEndpointinfoClearExist("citrixadc_endpointinfo_clear.tf_endpointinfo_clear", nil),
					resource.TestCheckResourceAttr("citrixadc_endpointinfo_clear.tf_endpointinfo_clear", "endpointkind", "IP"),
					// "id" is the synthetic state handle "endpointinfo_clear-IP".
					resource.TestCheckResourceAttrSet("citrixadc_endpointinfo_clear.tf_endpointinfo_clear", "id"),
				),
			},
		},
	})
}

// testAccCheckEndpointinfoClearExist is a state-only existence check.
//
// endpointinfo_clear is an action-only resource: Read is a no-op and there is no
// GET-by-id endpoint, so we CANNOT verify the clear via NITRO. We only assert that
// Terraform recorded the resource in state with a non-empty ID (which equals the
// synthetic "endpointinfo_clear-IP" after a successful POST ?action=clear).
func testAccCheckEndpointinfoClearExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No endpointinfo_clear ID is set")
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
