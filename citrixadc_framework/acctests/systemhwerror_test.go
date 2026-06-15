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

// NOTE on the systemhwerror resource:
//   - Models the NITRO POST /systemhwerror?action=check endpoint, which runs a
//     hardware/disk error check on the appliance (CLI: "check systemhwerror
//     -diskCheck ...") and returns read-only `response` / `hwerrorCount` fields.
//   - This is an ACTION-ONLY (one-shot side-effect) resource: Create performs the
//     check action, Read is a no-op (preserves state), Update is a no-op (the only
//     attribute is RequiresReplace), and Delete is a state-only removal. There is
//     NO add/get/update/delete endpoint, so the resource CANNOT be verified by
//     reading it back from the ADC, and the datasource was REMOVED (no NITRO GET
//     endpoint exists). This is the Pattern 13 (datasource-removed) shape.
//   - The single attribute `diskcheck` (bool, Required, RequiresReplace) maps to
//     the mandatory CLI `-diskCheck` argument ("Perform only disk error checking").
//     The check action is a non-destructive diagnostic and safe to run.
//   - The Exist check below only verifies that the resource landed in Terraform
//     state with its synthetic ID ("systemhwerror-config"); it does NOT (and
//     cannot) verify the check side-effect via NITRO.
//   - There is no CheckDestroy: the check action has no inverse on NITRO, and
//     there is no GET-by-id to confirm absence; Delete is a state-only removal.
//
// This mirrors the action-only test precedent in gslbconfig_test.go and
// clusterfiles_test.go (single apply step, state-only Exist check, no
// CheckDestroy, TestCheckResourceAttrSet on "id").

// Single apply step. `diskcheck` is the only (Required) attribute; the check
// action runs unconditionally on Create.
const testAccSystemhwerror_basic = `
resource "citrixadc_systemhwerror" "tf_systemhwerror" {
  diskcheck = true
}

`

func TestAccSystemhwerror_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// No CheckDestroy: the check action has no inverse on NITRO and there is
		// no GET-by-id to confirm absence; Delete is a state-only removal.
		Steps: []resource.TestStep{
			{
				Config: testAccSystemhwerror_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSystemhwerrorExist("citrixadc_systemhwerror.tf_systemhwerror", nil),
					// "id" is the synthetic state handle "systemhwerror-config".
					resource.TestCheckResourceAttrSet("citrixadc_systemhwerror.tf_systemhwerror", "id"),
					// Assert the flag actually set in HCL.
					resource.TestCheckResourceAttr("citrixadc_systemhwerror.tf_systemhwerror", "diskcheck", "true"),
				),
			},
		},
	})
}

// testAccCheckSystemhwerrorExist is a state-only existence check.
//
// systemhwerror is an action-only resource: Read is a no-op and there is no
// GET-by-id endpoint, so we CANNOT verify the check via NITRO. We only assert
// that Terraform recorded the resource in state with a non-empty ID (which
// equals the synthetic "systemhwerror-config" after a successful POST
// ?action=check). This mirrors testAccCheckGslbconfigExist.
func testAccCheckSystemhwerrorExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No systemhwerror ID is set")
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
