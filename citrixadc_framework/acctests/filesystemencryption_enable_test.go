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

// ============================================================================
// !!! DESTRUCTIVE / PLATFORM-GATED ACCEPTANCE TESTS - READ BEFORE RUNNING !!!
// ============================================================================
//
// citrixadc_filesystemencryption_enable is an ACTION-ONLY resource:
//   - Create fires  ?action=enable  (encrypts /flash and /var)
//   - There is NO GET, Update or Delete endpoint; every argument is
//     RequiresReplace and Read/Update/Delete are no-ops.
//
// The enable action is a REAL, POTENTIALLY DESTRUCTIVE platform operation that
// overwrites /flash and /var, toggles full file-system encryption on the ADC,
// AND REQUIRES A REBOOT. It is only permitted on a platform/appliance that
// supports File System Encryption (the nameless-singleton GET reports
// supportedstate == ENABLED). On a VPX / CPX / unsupported platform it fails.
//
// DO NOT run these tests against a production, shared, or otherwise important
// appliance. Every test in this file is Skip-gated for that reason.
//
// Because enable is nameless (singleton) and action-only, there is no GET
// endpoint and no "resource no longer exists" state, so the Exist check is
// state-only and CheckDestroy is intentionally omitted.
// ============================================================================

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// ---------------------------------------------------------------------------
// Basic (enable) test
// ---------------------------------------------------------------------------
//
// No update step: all arguments are RequiresReplace and there is no NITRO
// set/update endpoint, so a single enable step is the only meaningful path.
// TODO_PLACEHOLDER: supply a real passphrase for the target testbed in place
// of "MySecret123".
const testAccFilesystemencryptionEnable_basic_step1 = `
resource "citrixadc_filesystemencryption_enable" "tf_filesystemencryption_enable" {
  ntimes0flash = 0
  ntimes0var   = 0
  passphrase   = "MySecret123"
}
`

func TestAccFilesystemencryptionEnable_basic(t *testing.T) {
	t.Skip("TODO: Requires review - encrypts/decrypts the appliance flash filesystem and requires a reboot; must not run on a shared testbed")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// Action-only resource: no GET/DELETE, so CheckDestroy is omitted.
		Steps: []resource.TestStep{
			{
				Config: testAccFilesystemencryptionEnable_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFilesystemencryptionEnableExist("citrixadc_filesystemencryption_enable.tf_filesystemencryption_enable", nil),
					resource.TestCheckResourceAttr("citrixadc_filesystemencryption_enable.tf_filesystemencryption_enable", "ntimes0flash", "0"),
					resource.TestCheckResourceAttr("citrixadc_filesystemencryption_enable.tf_filesystemencryption_enable", "ntimes0var", "0"),
					// Synthetic id set by Create for this action-only resource.
					resource.TestCheckResourceAttr("citrixadc_filesystemencryption_enable.tf_filesystemencryption_enable", "id", "filesystemencryption_enable"),
				),
			},
		},
	})
}

// testAccCheckFilesystemencryptionEnableExist is a STATE-ONLY check. enable is an
// action-only resource with no NITRO GET endpoint, so there is nothing to query
// on the appliance - we only verify the resource is present in state with the
// synthetic id set by Create.
func testAccCheckFilesystemencryptionEnableExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No filesystemencryption_enable ID is set")
		}

		// Action-only resource: Create sets a synthetic id. There is no GET
		// endpoint to consult, so this is a pure state assertion.
		if rs.Primary.ID != "filesystemencryption_enable" {
			return fmt.Errorf("Unexpected filesystemencryption_enable ID %q, expected synthetic id \"filesystemencryption_enable\"", rs.Primary.ID)
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		return nil
	}
}
