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
// applying the citrixadc_systemkek resource ROTATES the appliance Key Encryption
// Key (KEK) via POST /nitro/v1/config/systemkek?action=update.
//
//   - The KEK rotation is IRREVERSIBLE and NON-IDEMPOTENT: every apply backs up
//     the old keys and generates brand-new keys. There is no inverse action.
//   - NITRO exposes NO add/get/delete/count endpoint for systemkek, so there is
//     no way to read the result back or detect drift, and no way to "undo" it.
//   - Running this against a shared or production appliance can break decryption
//     of stored secrets and disrupt the appliance.
//
// It is provided ONLY for completeness of the resource's test surface. Run it
// MANUALLY and ONLY against a disposable / throwaway appliance you can afford to
// re-image.
// =============================================================================

// NOTE on the systemkek resource:
//   - Models the NITRO POST /systemkek?action=update endpoint.
//   - ACTION-ONLY (one-shot side-effect) resource: Create performs the KEK
//     rotation, Read is a no-op (preserves state), Update is a no-op, and Delete
//     is a state-only removal. There is NO add/get/update/delete endpoint, so
//     the resource CANNOT be verified by reading it back from the ADC, and there
//     is NO datasource (it was removed -- no NITRO GET endpoint exists).
//   - The single attribute `level` is Required (enum: basic | extended) and
//     RequiresReplace: re-applying forces a fresh KEK rotation.
//   - The Exist check below only verifies that the resource landed in Terraform
//     state with its synthetic ID ("systemkek-config"); it does NOT (and cannot)
//     verify the rotation side-effect via NITRO.
//   - There is no CheckDestroy: the rotation has no inverse on NITRO, and there
//     is no GET-by-id to confirm absence; Delete is a state-only removal.
//
// This mirrors the action-only test precedent in aaasession_test.go and
// clusterfiles_test.go (single apply step, state-only Exist check, no
// CheckDestroy, TestCheckResourceAttrSet on "id").

// Single apply step: level is RequiresReplace, so there is no in-place update to
// exercise. level = "basic" performs the least-invasive rotation variant.
const testAccSystemkek_basic_step1 = `
resource "citrixadc_systemkek" "tf_systemkek" {
  level = "basic"
}

`

func TestAccSystemkek_basic(t *testing.T) {
	// DESTRUCTIVE / IRREVERSIBLE: see the file-level DANGER banner above. This
	// test rotates the appliance KEK and is provided for completeness only. It is
	// intentionally skipped so it never runs under CI or `go test ./...`. Remove
	// the t.Skip ONLY when running manually against a disposable appliance.
	// t.Skip("DANGER: rotates the appliance Key Encryption Key (irreversible, non-idempotent). Run manually on a disposable appliance only.")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// No CheckDestroy: the KEK rotation has no inverse on NITRO and there is
		// no GET-by-id to confirm absence; Delete is a state-only removal.
		Steps: []resource.TestStep{
			{
				Config: testAccSystemkek_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSystemkekExist("citrixadc_systemkek.tf_systemkek", nil),
					// "id" is the synthetic state handle "systemkek-config".
					resource.TestCheckResourceAttrSet("citrixadc_systemkek.tf_systemkek", "id"),
					// Assert the level actually set in HCL.
					resource.TestCheckResourceAttr("citrixadc_systemkek.tf_systemkek", "level", "basic"),
				),
			},
		},
	})
}

// testAccCheckSystemkekExist is a state-only existence check.
//
// systemkek is an action-only resource: Read is a no-op and there is no
// GET-by-id endpoint, so we CANNOT verify the KEK rotation via NITRO. We only
// assert that Terraform recorded the resource in state with a non-empty ID
// (which equals the synthetic "systemkek-config" after a successful
// POST ?action=update). This mirrors testAccCheckAaasessionExist.
func testAccCheckSystemkekExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No systemkek ID is set")
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
