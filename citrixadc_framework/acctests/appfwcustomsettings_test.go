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

// NOTE on the appfwcustomsettings resource (ACTION-ONLY export resource):
//   - Models the NITRO POST /appfwcustomsettings?action=export endpoint.
//   - This is a one-shot side-effect (action) resource:
//       Create = ActOnResource("export")
//       Read   = preserve-state no-op (NITRO has NO GET endpoint)
//       Update = no-op (all attributes are RequiresReplace)
//       Delete = state-only removal (export has no inverse on NITRO)
//   - Because there is no GET endpoint that reports export state, there is no
//     datasource for this resource (Bug Pattern 13 in
//     .claude/agents/FeatureDeveloper.md). The datasource files were removed by
//     the FeatureDeveloper agent, so NO datasource test is generated.
//   - The exist check below only verifies that the resource has a state ID;
//     it does NOT (and cannot) verify the export side-effect via NITRO.
//   - There is no destroy check (Delete is state-only; the export side-effect
//     cannot be undone via API).
//   - LIVE-RUN CAVEAT: the appfwcustomsettings export CLI verb is deprecated,
//     and the export source object (`name`) and writable `target` must already
//     exist / be valid on the ADC for the POST to succeed. A live run may
//     require a pre-existing custom-settings object on the testbed. Adjust the
//     `name`/`target` values below to match the testbed before running.

const testAccAppfwcustomsettings_basic_step1 = `
resource "citrixadc_appfwcustomsettings" "tf_appfwcustomsettings" {
  name   = "default"
  target = "local:exported_customsettings"
}

`

func TestAccAppfwcustomsettings_basic(t *testing.T) {
	t.Skip("TODO: Requires review")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// No CheckDestroy: the export action has no inverse on NITRO; Delete is a
		// state-only removal that does not touch the ADC.
		Steps: []resource.TestStep{
			{
				// Single-step apply: triggers the export action and confirms the
				// apply succeeds and the resource is recorded in state with an ID.
				// All attributes are RequiresReplace, so there is no in-place
				// update path worth a second step.
				Config: testAccAppfwcustomsettings_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwcustomsettingsExist("citrixadc_appfwcustomsettings.tf_appfwcustomsettings", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwcustomsettings.tf_appfwcustomsettings", "name", "default"),
					resource.TestCheckResourceAttr("citrixadc_appfwcustomsettings.tf_appfwcustomsettings", "target", "local:exported_customsettings"),
				),
			},
		},
	})
}

// Datasource: SKIPPED — no datasource files exist for appfwcustomsettings.
// NITRO has no GET endpoint for export state (Bug Pattern 13). The
// FeatureDeveloper agent deliberately removed the datasource for this
// action-only resource; we do not generate a datasource test.

// testAccCheckAppfwcustomsettingsExist is a state-only existence check.
// There is no NITRO GET endpoint for appfwcustomsettings (only ?action=export),
// so we cannot verify the export actually occurred on the ADC. We only assert
// that Terraform has recorded the resource in state with a non-empty ID (which
// equals `name` after a successful POST ?action=export). This is why this
// function intentionally does NOT call client.FindResource(...).
func testAccCheckAppfwcustomsettingsExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No appfwcustomsettings ID is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		// No NITRO GET to verify against; presence of the state ID is the only
		// confirmation we can make for this action-only export resource.
		return nil
	}
}
