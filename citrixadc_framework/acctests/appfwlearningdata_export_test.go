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

// NOTE on the appfwlearningdata_export resource:
//   - Models the NITRO POST /appfwlearningdata?action=export endpoint. Create
//     performs the action via ActOnResource(service.Appfwlearningdata.Type(),
//     &payload, "export"); Read is a no-op (preserves state), Update is a no-op
//     (all attributes are RequiresReplace), and Delete is a state-only removal.
//     There is NO get/add/update/delete-by-id endpoint for this action, so the
//     resource CANNOT be verified by reading it back from the ADC.
//   - The export payload carries exactly three attributes, derived from the NITRO
//     export payload / NetScaler CLI
//     `export appfw learningdata <profileName> <securityCheck> [-target <string>]`:
//     profilename (Required, RequiresReplace), securitycheck (Required,
//     RequiresReplace) and target (Optional, RequiresReplace). This test first
//     creates a parent citrixadc_appfwprofile, then exports its learned data.
//   - The Exist check below only verifies that the resource landed in Terraform
//     state with its synthetic ID ("appfwlearningdata_export"); it does NOT (and
//     cannot) verify the export side-effect via NITRO.
//   - There is no CheckDestroy: the export action has no inverse on NITRO and there
//     is no GET-by-id to confirm absence; Delete is a state-only removal. The
//     parent citrixadc_appfwprofile is still torn down by terraform destroy at test
//     end, leaving the appliance clean.
//
// PREREQUISITE / SKIP-GATE: export writes the learned data to a `target` file on
// the appliance and requires an App-Firewall profile that has actually accumulated
// learned data for the security check. Neither a writable destination nor the
// prerequisite learned data is guaranteed on the shared testbed, so this test is
// skip-gated until reviewed on a throwaway appliance.

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// Single apply step: profilename, securitycheck and target are all RequiresReplace,
// so there is no in-place update to exercise. The parent App-Firewall profile is
// created first and the export action is wired to it by reference + depends_on.
const testAccAppfwlearningdataExport_basic = `
resource "citrixadc_appfwprofile" "tf_appfwprofile" {
  name = "tf_appfwlearningdata_export_profile"
  type = ["HTML"]
}

resource "citrixadc_appfwlearningdata_export" "tf_appfwlearningdata_export" {
  profilename   = citrixadc_appfwprofile.tf_appfwprofile.name
  securitycheck = "startURL"
  target        = "/var/tmp/tf_appfwlearningdata_export.txt"
  depends_on    = [citrixadc_appfwprofile.tf_appfwprofile]
}

`

func TestAccAppfwlearningdataExport_basic(t *testing.T) {
	t.Skip("TODO: Requires review - export needs a writable destination / prerequisite not available on the shared testbed")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// No CheckDestroy: the export action has no inverse on NITRO and there is no
		// GET-by-id to confirm absence; Delete is a state-only removal. The parent
		// appfwprofile is torn down by terraform destroy.
		Steps: []resource.TestStep{
			{
				Config: testAccAppfwlearningdataExport_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwlearningdataExportExist("citrixadc_appfwlearningdata_export.tf_appfwlearningdata_export", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwlearningdata_export.tf_appfwlearningdata_export", "profilename", "tf_appfwlearningdata_export_profile"),
					resource.TestCheckResourceAttr("citrixadc_appfwlearningdata_export.tf_appfwlearningdata_export", "securitycheck", "startURL"),
					// "id" is the synthetic state handle "appfwlearningdata_export".
					resource.TestCheckResourceAttrSet("citrixadc_appfwlearningdata_export.tf_appfwlearningdata_export", "id"),
				),
			},
		},
	})
}

// testAccCheckAppfwlearningdataExportExist is a state-only existence check.
//
// appfwlearningdata_export is an action-only resource: Read is a no-op and there is
// no GET-by-id endpoint, so we CANNOT verify the export via NITRO. We only assert
// that Terraform recorded the resource in state with a non-empty ID (which equals
// the synthetic "appfwlearningdata_export" after a successful POST ?action=export).
func testAccCheckAppfwlearningdataExportExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No appfwlearningdata_export ID is set")
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
