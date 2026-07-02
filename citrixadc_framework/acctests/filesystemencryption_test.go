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
// citrixadc_filesystemencryption is an ACTION-ONLY resource:
//   - Create fires  ?action=enable  (encrypts /flash and /var)
//   - Delete fires  ?action=disable (decrypts the file system)
//   - There is NO Update endpoint; every argument is RequiresReplace.
//
// These operations are REAL, POTENTIALLY DESTRUCTIVE platform operations that
// overwrite /flash and /var and toggle full file-system encryption on the ADC.
// They are ONLY permitted on a platform/appliance that supports File System
// Encryption (i.e. the nameless-singleton GET reports supportedstate == ENABLED).
// On a VPX / CPX / unsupported platform the enable action will fail.
//
// DO NOT run these tests against a production, shared, or otherwise important
// appliance. Run only on a disposable testbed that explicitly supports FS
// encryption. TODO_PLACEHOLDER: confirm the target testbed supports FS
// encryption (supportedstate == ENABLED) before enabling this test, and supply
// a real passphrase value in place of every "TODO_PLACEHOLDER_passphrase*".
//
// Because enable/disable is nameless (singleton) and action-only, there is no
// meaningful "resource no longer exists" state after disable, so CheckDestroy
// is intentionally omitted.
// ============================================================================

import (
	"fmt"
	"testing"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// ---------------------------------------------------------------------------
// Basic (enable) test
// ---------------------------------------------------------------------------
//
// No update step: all arguments are RequiresReplace and there is no NITRO
// set/update endpoint, so a single enable step is the only meaningful CRUD path.
const testAccFilesystemencryption_basic_step1 = `
resource "citrixadc_filesystemencryption" "tf_filesystemencryption" {
  ntimes0flash = 0
  ntimes0var   = 0
  passphrase   = "MySecret123"
}
`

func TestAccFilesystemencryption_basic(t *testing.T) {
	// !!! DESTRUCTIVE / PLATFORM-GATED !!!
	// This performs a REAL enable (and disable on teardown) of full file-system
	// encryption. Only run on a disposable testbed that supports FS encryption
	// (supportedstate == ENABLED). See the file-level warning banner above.
	// TODO_PLACEHOLDER: remove this Skip once you have confirmed the target
	// testbed supports FS encryption and you have supplied a real passphrase.
	// t.Skip("filesystemencryption performs a REAL, potentially destructive enable/disable and is platform-gated (supportedstate must be ENABLED). Remove this Skip only on a supported disposable testbed with a real passphrase set.")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// Action-only resource: disable is fired on teardown, but there is no
		// meaningful "not found" state to assert, so CheckDestroy is omitted.
		Steps: []resource.TestStep{
			{
				Config: testAccFilesystemencryption_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFilesystemencryptionExist("citrixadc_filesystemencryption.tf_filesystemencryption", nil),
					resource.TestCheckResourceAttr("citrixadc_filesystemencryption.tf_filesystemencryption", "ntimes0flash", "0"),
					resource.TestCheckResourceAttr("citrixadc_filesystemencryption.tf_filesystemencryption", "ntimes0var", "0"),
					// Computed read-only state populated from the nameless singleton GET.
					resource.TestCheckResourceAttrSet("citrixadc_filesystemencryption.tf_filesystemencryption", "supportedstate"),
					resource.TestCheckResourceAttrSet("citrixadc_filesystemencryption.tf_filesystemencryption", "effectivestate"),
				),
			},
		},
	})
}

// testAccCheckFilesystemencryptionExist verifies the resource is in state and
// that the nameless singleton GET is reachable. The GET may legitimately return
// nothing meaningful on some platforms (the resource is action-only), so a nil
// GET response is tolerated - we only hard-fail if the state ID is missing or
// the GET itself errors.
func testAccCheckFilesystemencryptionExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No filesystemencryption ID is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		// Nameless singleton GET (args=nodeid may filter, but "" targets the node).
		if _, err := client.FindResource(service.Filesystemencryption.Type(), ""); err != nil {
			return err
		}

		return nil
	}
}

// ---------------------------------------------------------------------------
// Write-only / ephemeral (passphrase_wo) test
// ---------------------------------------------------------------------------
//
// passphrase_wo is WriteOnly: its value is never persisted in state and cannot
// be asserted with TestCheckResourceAttr. passphrase_wo_version IS stored in
// state; bumping it forces a RequiresReplace (disable+enable) so the new
// passphrase is applied.
const testAccFilesystemencryption_wo_step1 = `
	variable "filesystemencryption_passphrase_wo" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_filesystemencryption" "tf_filesystemencryption" {
		ntimes0flash          = 0
		ntimes0var            = 0
		passphrase_wo         = var.filesystemencryption_passphrase_wo
		passphrase_wo_version = 1
	}
`

const testAccFilesystemencryption_wo_step2 = `
	variable "filesystemencryption_passphrase_wo_2" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_filesystemencryption" "tf_filesystemencryption" {
		ntimes0flash          = 0
		ntimes0var            = 0
		passphrase_wo         = var.filesystemencryption_passphrase_wo_2
		passphrase_wo_version = 2
	}
`

func TestAccFilesystemencryption_writeOnly(t *testing.T) {
	// !!! DESTRUCTIVE / PLATFORM-GATED !!!
	// This performs REAL enable/disable cycles of full file-system encryption
	// (each passphrase_wo_version bump triggers a RequiresReplace = disable+enable).
	// Only run on a disposable testbed that supports FS encryption
	// (supportedstate == ENABLED). See the file-level warning banner above.
	// TODO_PLACEHOLDER: remove this Skip once you have confirmed the target
	// testbed supports FS encryption and you have supplied real passphrase values
	// (via TF_VAR_* below or your own values).
	t.Skip("filesystemencryption performs a REAL, potentially destructive enable/disable and is platform-gated (supportedstate must be ENABLED). Remove this Skip only on a supported disposable testbed with real passphrase values set.")

	// TODO_PLACEHOLDER: replace these with real passphrase values for the testbed.
	t.Setenv("TF_VAR_filesystemencryption_passphrase_wo", "TODO_PLACEHOLDER_passphrase_wo_1")
	t.Setenv("TF_VAR_filesystemencryption_passphrase_wo_2", "TODO_PLACEHOLDER_passphrase_wo_2")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccFilesystemencryption_wo_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFilesystemencryptionExist("citrixadc_filesystemencryption.tf_filesystemencryption", nil),
					// Do NOT assert passphrase_wo - write-only values are not stored in state.
					resource.TestCheckResourceAttr("citrixadc_filesystemencryption.tf_filesystemencryption", "passphrase_wo_version", "1"),
					resource.TestCheckResourceAttr("citrixadc_filesystemencryption.tf_filesystemencryption", "ntimes0flash", "0"),
					resource.TestCheckResourceAttr("citrixadc_filesystemencryption.tf_filesystemencryption", "ntimes0var", "0"),
				),
			},
			{
				Config: testAccFilesystemencryption_wo_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFilesystemencryptionExist("citrixadc_filesystemencryption.tf_filesystemencryption", nil),
					// Version bump confirms the RequiresReplace update path was triggered.
					resource.TestCheckResourceAttr("citrixadc_filesystemencryption.tf_filesystemencryption", "passphrase_wo_version", "2"),
				),
			},
		},
	})
}

// ---------------------------------------------------------------------------
// Datasource test (nameless singleton GET exposing supportedstate/effectivestate)
// ---------------------------------------------------------------------------
//
// The datasource reads the nameless singleton GET. It requires the resource to
// have enabled FS encryption first, so it inherits the same destructive/platform
// gating as the resource tests.
const testAccFilesystemencryptionDataSource_basic = `
resource "citrixadc_filesystemencryption" "tf_filesystemencryption" {
  ntimes0flash = 0
  ntimes0var   = 0
  passphrase   = "TODO_PLACEHOLDER_passphrase"
}

data "citrixadc_filesystemencryption" "tf_filesystemencryption" {
  depends_on = [citrixadc_filesystemencryption.tf_filesystemencryption]
}
`

func TestAccFilesystemencryptionDataSource_basic(t *testing.T) {
	// !!! DESTRUCTIVE / PLATFORM-GATED !!!
	// This creates the resource (real enable) before reading the datasource.
	// Only run on a disposable testbed that supports FS encryption. See banner above.
	// TODO_PLACEHOLDER: remove this Skip and supply a real passphrase to run.
	t.Skip("filesystemencryption datasource test creates the resource via a REAL, potentially destructive enable and is platform-gated. Remove this Skip only on a supported disposable testbed with a real passphrase set.")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccFilesystemencryptionDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Read-only state exposed by the nameless singleton GET.
					resource.TestCheckResourceAttrSet("data.citrixadc_filesystemencryption.tf_filesystemencryption", "supportedstate"),
					resource.TestCheckResourceAttrSet("data.citrixadc_filesystemencryption.tf_filesystemencryption", "effectivestate"),
					// passphrase / passphrase_wo are secrets never returned by NITRO - not checked.
				),
			},
		},
	})
}
