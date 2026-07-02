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

// NOTE on the application resource:
//   - Models the AppExpert application-template Import object. This is an
//     ACTION-ONLY resource (NOT standard CRUD):
//       * Create  = NITRO POST /application?action=Import (imports an app template file).
//       * Read    = no-op: the NITRO application object exposes NO get/get(all)
//                   endpoint, so drift cannot be detected and the object cannot
//                   be read back from the ADC.
//       * Update  = no-op: all attributes are RequiresReplace and there is no
//                   NITRO set/update endpoint.
//       * Delete  = DELETE /application?args=appname:<name> (runs at test teardown).
//   - Because there is NO GET endpoint, the datasource was REMOVED (Pattern 13),
//     so there is intentionally NO datasource test in this file.
//   - The Exist check below is STATE-ONLY: it verifies the resource landed in
//     Terraform state with its ID (== appname). It does NOT (and cannot) verify
//     the Import side-effect via a NITRO GET. A GET-based existence check is not
//     possible for this object (see TODO_PLACEHOLDER note below).
//   - There is no CheckDestroy: Delete has no GET-by-id to confirm absence.
//     Delete still runs during the normal test-teardown apply (Terraform destroy),
//     exercising the DELETE args=appname:<name> path; its success/failure surfaces
//     as a teardown error, not via a CheckDestroy read-back.
//
// This mirrors the action-only test precedent in gslbconfig_test.go and
// clusterfiles_test.go (single apply step, state-only Exist check, no CheckDestroy,
// TestCheckResourceAttrSet on "id").
//
// IMPORTANT LIVE PREREQUISITE (external file on the appliance):
//   The ?action=Import operation requires an AppExpert application-template file
//   to ALREADY EXIST on the ADC appliance filesystem (typically under
//   /var/tmp/... or the configured template location) BEFORE this test runs.
//   That file cannot be created by the Terraform config alone. Before running
//   this test live, you MUST:
//     1. Upload/place a valid application template file on the ADC appliance, and
//     2. Replace the TODO_PLACEHOLDER value of `apptemplatefilename` below with
//        that file's name.
//   Optionally set `deploymentfilename` to a valid deployment file present on the
//   appliance (left unset here). Until the prerequisite file exists and the
//   placeholder is filled in, the live Import will fail.

// Single apply step. Create runs the ?action=Import operation. All attributes are
// RequiresReplace, so there is no meaningful in-place update step.
//
// TODO_PLACEHOLDER: replace "TODO_PLACEHOLDER" for apptemplatefilename with the
// name of a valid AppExpert application template file that already exists on the
// ADC appliance filesystem (see the IMPORTANT LIVE PREREQUISITE note above).
const testAccApplication_basic = `
resource "citrixadc_application" "tf_application" {
  appname             = "tf_test_application"
  apptemplatefilename = "TODO_PLACEHOLDER"
  # deploymentfilename = "TODO_PLACEHOLDER" # optional: a valid deployment file on the appliance
}

`

func TestAccApplication_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// No CheckDestroy: the application object has no GET-by-id endpoint to
		// confirm absence after Delete. The DELETE args=appname path still runs
		// during Terraform destroy at test teardown.
		Steps: []resource.TestStep{
			{
				Config: testAccApplication_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckApplicationExist("citrixadc_application.tf_application", nil),
					// "id" is set to appname after a successful ?action=Import.
					resource.TestCheckResourceAttrSet("citrixadc_application.tf_application", "id"),
					// Assert the attributes actually set in HCL.
					resource.TestCheckResourceAttr("citrixadc_application.tf_application", "appname", "tf_test_application"),
					// TODO_PLACEHOLDER: once apptemplatefilename above is set to a real
					// filename, update the expected value in this assertion to match.
					resource.TestCheckResourceAttr("citrixadc_application.tf_application", "apptemplatefilename", "TODO_PLACEHOLDER"),
				),
			},
		},
	})
}

// testAccCheckApplicationExist is a STATE-ONLY existence check.
//
// application is an action-only resource: Read is a no-op and there is NO
// GET-by-id endpoint on NITRO, so we CANNOT verify the Import via NITRO. We only
// assert that Terraform recorded the resource in state with a non-empty ID (which
// equals the configured appname after a successful POST ?action=Import). This
// mirrors testAccCheckGslbconfigExist and testAccCheckClusterfilesActionExist.
//
// TODO_PLACEHOLDER: a GET-based existence check is intentionally omitted because
// the NITRO application object exposes no get/get(all) endpoint. If a future NITRO
// version adds a GET endpoint for this object, replace this state-only check with
// a client.FindResource(service.Application.Type(), rs.Primary.ID) verification.
func testAccCheckApplicationExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No application ID is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		// No NITRO GET-by-id to verify against; presence of the state ID is the
		// only confirmation we can make for an action-only resource.
		return nil
	}
}
