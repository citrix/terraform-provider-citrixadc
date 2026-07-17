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

// NOTE on the policyurlset_export resource:
//   - Models the NITRO POST /policyurlset?action=export endpoint. Create performs
//     the action via ActOnResource(service.Policyurlset.Type(), &payload,
//     "export"); Read is a no-op (preserves state), Update is a no-op (all
//     attributes are RequiresReplace), and Delete is a state-only removal. There
//     is NO get/add/update/delete-by-id endpoint for this action, so the resource
//     CANNOT be verified by reading it back from the ADC, and it has NO
//     datasource (no NITRO GET endpoint).
//   - Attributes `name` (the url set to export, must already exist) and `url` (the
//     export DESTINATION) are both Required and RequiresReplace. `url` must be a
//     WRITABLE remote destination reachable from the appliance — only HTTP, HTTPS
//     and FTP protocols are supported (local: is NOT accepted for export). This
//     test therefore first creates a parent citrixadc_policyurlset (reusing the
//     parent test's working HCL and the doPolicyUrlSetPreChecks PreCheck), then
//     exports it to `url`.
//   - The Exist check below only verifies that the resource landed in Terraform
//     state with its synthetic ID ("policyurlset_export-<name>"); it does NOT
//     (and cannot) verify the export side-effect via NITRO.
//   - There is no CheckDestroy: the export action has no inverse on NITRO and
//     there is no GET-by-id to confirm absence; Delete is a state-only removal.
//     The parent citrixadc_policyurlset is still torn down by terraform destroy at
//     test end, leaving the appliance clean.
//   - The parent test (policyurlset_test.go) is NOT skip-gated, so neither is
//     this one.
//
// PREREQUISITE: `url` below is TODO_PLACEHOLDER because export requires a
// reachable, WRITABLE FTP/HTTP(S) destination server. Substitute a real
// destination reachable from the testbed appliance (e.g.
// "ftp://user:pass@host/path/urlset.csv") before this test can run green.

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// Single apply step: both `name` and `url` are RequiresReplace, so there is no
// in-place update to exercise. The parent url set is created first (from the
// parent test's working HCL) and the export action is wired to it by reference +
// depends_on. `url` is the export DESTINATION and must be a writable HTTP/HTTPS/
// FTP endpoint reachable from the appliance.
const testAccPolicyurlsetExport_basic = `
resource "citrixadc_policyurlset" "tf_policyurlset" {
  name                = "tf_policyurlset_export"
  url                 = "local:tftest.urlset"
  interval            = 0
  matchedid           = 2
  subdomainexactmatch = false
}

resource "citrixadc_policyurlset_export" "tf_policyurlset_export" {
  name       = citrixadc_policyurlset.tf_policyurlset.name
  url        = "TODO_PLACEHOLDER"
  depends_on = [citrixadc_policyurlset.tf_policyurlset]
}

`

func TestAccPolicyurlsetExport_basic(t *testing.T) {
	// export writes the CSV to a WRITABLE remote HTTP/HTTPS/FTP destination
	// (local: is not accepted). No such reachable destination exists for the
	// shared testbed, so this test is skipped until a real `url` destination is
	// substituted for the TODO_PLACEHOLDER in the config above.
	t.Skip("TODO: Requires review - export needs a writable remote FTP/HTTP(S) destination reachable from the appliance")
	resource.Test(t, resource.TestCase{
		// Reuse the parent PreCheck: it uploads testdata/tftest.urlset so the
		// parent policyurlset import resolves before the export action runs.
		PreCheck:                 func() { doPolicyUrlSetPreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// No CheckDestroy: the export action has no inverse on NITRO and there is
		// no GET-by-id to confirm absence; Delete is a state-only removal. The
		// parent policyurlset is torn down by terraform destroy.
		Steps: []resource.TestStep{
			{
				Config: testAccPolicyurlsetExport_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPolicyurlsetExportExist("citrixadc_policyurlset_export.tf_policyurlset_export", nil),
					resource.TestCheckResourceAttr("citrixadc_policyurlset_export.tf_policyurlset_export", "name", "tf_policyurlset_export"),
					// "id" is the synthetic state handle "policyurlset_export-<name>".
					resource.TestCheckResourceAttrSet("citrixadc_policyurlset_export.tf_policyurlset_export", "id"),
				),
			},
		},
	})
}

// testAccCheckPolicyurlsetExportExist is a state-only existence check.
//
// policyurlset_export is an action-only resource: Read is a no-op and there is no
// GET-by-id endpoint, so we CANNOT verify the export via NITRO. We only assert
// that Terraform recorded the resource in state with a non-empty ID (which equals
// the synthetic "policyurlset_export-<name>" after a successful POST
// ?action=export).
func testAccCheckPolicyurlsetExportExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No policyurlset_export ID is set")
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
