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

// NOTE on the nsextension_change resource:
//   - Models the NITRO nsextension `change` action, which reloads / recompiles an
//     extension object from its stored source file. Although the NITRO doc anchor
//     is named "change", the op is invoked at POST ?action=update (the literal
//     "change" verb is rejected by NITRO with errorcode 1240), so the provider
//     passes the "update" verb to ActOnResource(service.Nsextension.Type(), ...).
//   - This is an ACTION-ONLY (one-shot side-effect) resource: Create performs the
//     change/reload action, Read is a no-op (preserves state), Update is a no-op
//     (the single attribute `name` is RequiresReplace), and Delete is a state-only
//     removal. There is NO get/add/update/delete endpoint for the change action,
//     so the resource CANNOT be verified by reading it back from the ADC, and it
//     has NO datasource.
//   - The single attribute `name` is Required and RequiresReplace. The change
//     action reloads an EXISTING extension object from its stored source file, so
//     the named nsextension MUST already exist on the appliance before the change
//     action is applied. The config below therefore creates a citrixadc_nsextension
//     first (reusing the parent's basic config) and wires the change to it by
//     reference + depends_on.
//   - The Exist check below only verifies that the resource landed in Terraform
//     state with its synthetic ID ("nsextension_change-<name>"); it does NOT (and
//     cannot) verify the change side-effect via NITRO.
//   - There is no CheckDestroy on the change resource: the change action has no
//     inverse on NITRO, and there is no GET-by-id to confirm absence; Delete is a
//     state-only removal. (The prerequisite citrixadc_nsextension IS destroyed by
//     Terraform at end of test, but the change action itself has no CheckDestroy.)
//
// This mirrors the action-only test precedent in protocolhttpband_clear_test.go
// (single apply step, state-only Exist check, no CheckDestroy).
//
// PREREQUISITE / SKIP: creating the prerequisite citrixadc_nsextension requires
// importing a real reachable .lua extension file via `src` (see nsextension_test.go,
// where `src` is a TODO_PLACEHOLDER "local:tftest_extension.lua" and every parent
// test is skip-gated with t.Skip("TODO: Requires review")). Until a real .lua
// source is staged on the testbed, the prerequisite extension cannot be created
// and therefore the change action cannot run. This test is skip-gated the same way
// as the parent nsextension tests.

// Single apply step: `name` is RequiresReplace, so there is no in-place update to
// exercise. The prerequisite extension is created first and the change action is
// wired to it by reference + depends_on.
//
// TODO_PLACEHOLDER: replace `src` with a real reachable .lua extension (an uploaded
// local file or an accessible URL). "local:tftest_extension.lua" is a placeholder
// and will fail Import until a real source is supplied (see nsextension_test.go).
const testAccNsextensionChange_basic = `
resource "citrixadc_nsextension" "tf_nsextension" {
  name    = "tf_nsextension_change"
  src     = "local:tftest_extension.lua"
  comment = "created by nsextension_change acceptance test"
}

resource "citrixadc_nsextension_change" "tf_nsextension_change" {
  name       = citrixadc_nsextension.tf_nsextension.name
  depends_on = [citrixadc_nsextension.tf_nsextension]
}

`

func TestAccNsextensionChange_basic(t *testing.T) {
	// Skip-gated to match the parent nsextension tests: the prerequisite
	// citrixadc_nsextension cannot be created until a real .lua extension source
	// is staged on the testbed (src is a TODO_PLACEHOLDER).
	t.Skip("TODO: Requires review")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// No CheckDestroy on the change resource: the change action has no inverse
		// on NITRO and there is no GET-by-id to confirm absence; Delete is a
		// state-only removal.
		Steps: []resource.TestStep{
			{
				Config: testAccNsextensionChange_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsextensionChangeExist("citrixadc_nsextension_change.tf_nsextension_change", nil),
					resource.TestCheckResourceAttr("citrixadc_nsextension_change.tf_nsextension_change", "name", "tf_nsextension_change"),
					// "id" is the synthetic state handle "nsextension_change-<name>".
					resource.TestCheckResourceAttrSet("citrixadc_nsextension_change.tf_nsextension_change", "id"),
				),
			},
		},
	})
}

// testAccCheckNsextensionChangeExist is a state-only existence check.
//
// nsextension_change is an action-only resource: Read is a no-op and there is no
// GET-by-id endpoint for the change action, so we CANNOT verify the change via
// NITRO. We only assert that Terraform recorded the resource in state with a
// non-empty ID (which equals the synthetic "nsextension_change-<name>" after a
// successful POST ?action=update). This mirrors testAccCheckProtocolhttpbandClearExist.
func testAccCheckNsextensionChangeExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No nsextension_change ID is set")
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
