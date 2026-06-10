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

// NOTE on the gslbldnsentry resource:
//   - gslbldnsentry is an UNUSUAL "delete-only" NITRO resource: the NITRO API
//     exposes ONLY the `delete` verb (no add/get/update/count/clear). The only
//     thing you can do with it is REMOVE a single runtime-learned LDNS entry by
//     its IP address (CLI: `rm gslb ldnsentry <ipaddress>`).
//   - The provider models it as "delete-as-create": APPLYING this resource
//     performs the NITRO HTTP DELETE that removes the learned LDNS entry with the
//     given `ipaddress`. Read is a no-op (no GET endpoint), Update is a no-op
//     (`ipaddress` is RequiresReplace), and Delete (Terraform destroy) is a
//     state-only removal (the entry was already removed at create time).
//   - Because there is NO GET-by-id endpoint, the resource CANNOT be verified by
//     reading it back from the ADC, and the datasource was REMOVED (Pattern 13:
//     no NITRO GET endpoint exists), so there is NO datasource test below.
//   - This mirrors the action-only test precedent in gslbconfig_test.go and
//     clusterfiles_test.go: a single apply step, a state-only Exist check that
//     does NOT call the ADC, no CheckDestroy, and TestCheckResourceAttrSet on the
//     synthetic "id". The synthetic id equals the ipaddress acted upon.
//
// ENVIRONMENT: this works on any box. Applying it removes a runtime-learned LDNS
// entry keyed by ipaddress; if that IP is not currently learned the NITRO DELETE
// is treated as an idempotent no-op (nitro-go treats a 404 delete as
// already-deleted), so an arbitrary test IP (192.0.2.1, TEST-NET-1) applies
// cleanly. No special ADC_TESTBED gate is required.

// Single apply step. The only attribute is the required `ipaddress`
// (RequiresReplace). Applying performs the delete-as-create action.
const testAccGslbldnsentry_basic = `
resource "citrixadc_gslbldnsentry" "tf_gslbldnsentry" {
  ipaddress = "192.0.2.1"
}

`

func TestAccGslbldnsentry_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// No CheckDestroy: the delete-as-create action has no inverse on NITRO and
		// there is no GET-by-id to confirm absence; Terraform destroy is a
		// state-only removal (the entry was already removed at create time).
		Steps: []resource.TestStep{
			{
				Config: testAccGslbldnsentry_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGslbldnsentryExist("citrixadc_gslbldnsentry.tf_gslbldnsentry", nil),
					// "id" is the synthetic state handle (equals the ipaddress).
					resource.TestCheckResourceAttrSet("citrixadc_gslbldnsentry.tf_gslbldnsentry", "id"),
					// Assert the attribute set in HCL.
					resource.TestCheckResourceAttr("citrixadc_gslbldnsentry.tf_gslbldnsentry", "ipaddress", "192.0.2.1"),
					// id equals the ipaddress acted upon.
					resource.TestCheckResourceAttr("citrixadc_gslbldnsentry.tf_gslbldnsentry", "id", "192.0.2.1"),
				),
			},
		},
	})
}

// testAccCheckGslbldnsentryExist is a state-only existence check.
//
// gslbldnsentry is a delete-as-create / action-only resource: Read is a no-op
// and there is no GET-by-id endpoint, so we CANNOT verify the removal via NITRO.
// We only assert that Terraform recorded the resource in state with a non-empty
// ID (which equals the ipaddress after a successful delete-as-create). This
// mirrors testAccCheckGslbconfigExist and testAccCheckClusterfilesActionExist:
// it deliberately makes NO call to the ADC.
func testAccCheckGslbldnsentryExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No gslbldnsentry ID is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		// No NITRO GET-by-id to verify against; presence of the synthetic state
		// ID is the only confirmation we can make for a delete-as-create
		// action-only resource.
		return nil
	}
}
