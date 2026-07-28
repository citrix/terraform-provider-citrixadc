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

// NOTE on the application_export resource:
//   - Models the NITRO application `?action=export` action. This is an
//     ACTION-ONLY resource (NOT standard CRUD):
//       * Create  = NITRO POST /application?action=export (exports an existing
//                   AppExpert application to a template file on the appliance).
//       * Read    = no-op: NITRO exposes NO get/get(all) endpoint for export
//                   state, so drift cannot be detected.
//       * Update  = no-op: all attributes are RequiresReplace and there is no
//                   NITRO set/update endpoint.
//       * Delete  = no-op: export is a one-shot side-effect with no inverse API
//                   ("un-export"); Delete only drops the resource from state.
//   - Because there is NO GET endpoint, there is intentionally NO datasource for
//     this resource and NO datasource test in this file (Pattern 13).
//   - The Exist check below is STATE-ONLY: Create assigns the synthetic ID
//     "application_export"; the check asserts Terraform recorded that ID. It does
//     NOT (and cannot) verify the export side-effect via a NITRO GET.
//   - There is NO CheckDestroy: Delete is a state-only no-op with no GET-by-id to
//     confirm absence. This mirrors the action-only precedent in gslbconfig and
//     clusterfiles tests (single apply step, state-only Exist, no CheckDestroy).
//
// nitro_rest reference (nitro_rest/app/application.html), export operation:
//   URL:    /nitro/v1/config/application?action=export   (HTTP POST)
//   Payload: {"application":{ "appname":<String>,
//                             "apptemplatefilename":<String>,
//                             "deploymentfilename":<String> }}
//   appname is mandatory; apptemplatefilename / deploymentfilename are optional.
//
// IMPORTANT LIVE PREREQUISITE (not available on the shared testbed):
//   ?action=export requires that an AppExpert application named `appname` ALREADY
//   EXISTS on the ADC (imported beforehand) and that `apptemplatefilename` names a
//   writable destination template file on the appliance filesystem. Neither the
//   source application nor a writable destination is provisioned on the shared
//   testbed, so this test is SKIP-GATED. Before running live you MUST:
//     1. Import/create an AppExpert application on the ADC and set `appname` to it, and
//     2. Set `apptemplatefilename` to a writable destination file on the appliance
//        (replace TODO_PLACEHOLDER), and update the assertion below to match.

// Single apply step. Create runs the ?action=export operation. All attributes are
// RequiresReplace, so there is no meaningful in-place update step.
//
// TODO_PLACEHOLDER: replace "TODO_PLACEHOLDER" for apptemplatefilename with a
// writable destination template file on the ADC appliance, and set `appname` to an
// AppExpert application that already exists on the appliance.
const testAccApplicationExport_basic = `
resource "citrixadc_application_export" "tf_application_export" {
  appname             = "tf_test_application"
  apptemplatefilename = "TODO_PLACEHOLDER"
  # deploymentfilename = "TODO_PLACEHOLDER" # optional: a valid deployment file on the appliance
}

`

func TestAccApplicationExport_basic(t *testing.T) {
	t.Skip("TODO: Requires review - export needs a writable destination / prerequisite not available on the shared testbed")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// No CheckDestroy: export has no inverse NITRO API and no GET-by-id to
		// confirm absence. Delete is a state-only no-op at Terraform teardown.
		Steps: []resource.TestStep{
			{
				Config: testAccApplicationExport_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckApplicationExportExist("citrixadc_application_export.tf_application_export", nil),
					// Create assigns the synthetic ID "application_export".
					resource.TestCheckResourceAttr("citrixadc_application_export.tf_application_export", "id", "application_export"),
					// Assert the attributes actually set in HCL.
					resource.TestCheckResourceAttr("citrixadc_application_export.tf_application_export", "appname", "tf_test_application"),
					// TODO_PLACEHOLDER: once apptemplatefilename above is set to a real
					// filename, update the expected value in this assertion to match.
					resource.TestCheckResourceAttr("citrixadc_application_export.tf_application_export", "apptemplatefilename", "TODO_PLACEHOLDER"),
				),
			},
		},
	})
}

// testAccCheckApplicationExportExist is a STATE-ONLY existence check.
//
// application_export is an action-only resource: ?action=export is a one-shot
// side-effect with NO GET endpoint and no inverse API, so we CANNOT verify the
// export via NITRO. Create assigns the synthetic ID "application_export"; this
// check only asserts Terraform recorded that ID in state. This mirrors
// testAccCheckGslbconfigExist and testAccCheckClusterfilesActionExist.
//
// TODO_PLACEHOLDER: a GET-based existence check is intentionally omitted because
// the NITRO application object exposes no get/get(all) endpoint for export state.
func testAccCheckApplicationExportExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No application_export ID is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		// No NITRO GET-by-id to verify against; the synthetic state ID is the only
		// confirmation we can make for this action-only resource.
		if rs.Primary.ID != "application_export" {
			return fmt.Errorf("application_export ID = %q, want the synthetic ID %q", rs.Primary.ID, "application_export")
		}

		return nil
	}
}
