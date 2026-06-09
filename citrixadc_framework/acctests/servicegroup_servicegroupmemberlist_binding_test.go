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

// servicegroup_servicegroupmemberlist_binding is a write-only bulk member-list binding.
// NITRO exposes NO get/get(all)/count endpoint for this binding (Pattern 13):
//   - Create  = PUT (UpdateUnnamedResource) with the full member set.
//   - Read    = no-op (state is preserved; drift detection is impossible by design).
//   - Update  = no-op (all attributes are RequiresReplace; the resource is not updateable).
//   - Delete  = DeleteResource(servicegroupname) (path key, no args).
//
// Because there is no GET-by-id for the binding, this test follows the no-GET /
// action-only pattern (cf. clusterfiles / gslbconfig / appfwarchive_export):
//   - testAccCheckServicegroupServicegroupmemberlistBindingExist is STATE-ONLY: it
//     verifies rs.Primary.ID != "" and does NOT call client.FindResource on the
//     binding type (there is no GET endpoint to call).
//   - CheckDestroy is omitted from the TestCase: there is no GET to verify the
//     binding was removed from the ADC.
//
// The datasource was removed (Pattern 13, no GET), so there is no datasource test.
//
// The participating parent entity (servicegroup) MUST be autoscale-enabled. The bulk
// desired-state memberlist binding is a NITRO "Operation not permitted" (errorcode 257)
// on a plain (autoscale=DISABLED) servicegroup: individual members on a plain group must
// be bound one-at-a-time via servicegroup_servicegroupmember_binding. The bulk member-list
// PUT is reserved for autoscale servicegroups (verified live: autoscale=API returns
// errorcode 0; plain HTTP returns errorcode 257 for any member field set / method). We use
// autoscale = "API" (CLOUD/POLICY require platform support / memberport not available on the
// test appliance). The binding wires itself to the servicegroup by reference + depends_on.

const testAccServicegroupServicegroupmemberlistBinding_basic = `
resource "citrixadc_servicegroup" "tf_servicegroup" {
  servicegroupname = "tf_memberlist_svcgroup"
  servicetype      = "HTTP"
  autoscale        = "API"
}

resource "citrixadc_servicegroup_servicegroupmemberlist_binding" "tf_binding" {
  servicegroupname = citrixadc_servicegroup.tf_servicegroup.servicegroupname

  members = [
    {
      ip     = "10.10.10.10"
      port   = 80
      weight = 10
      state  = "ENABLED"
      order  = 1
    },
    {
      ip     = "10.10.10.11"
      port   = 80
      weight = 20
      state  = "ENABLED"
      order  = 2
    },
  ]

  depends_on = [citrixadc_servicegroup.tf_servicegroup]
}
`

func TestAccServicegroupServicegroupmemberlistBinding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// No CheckDestroy: NITRO has no GET endpoint for this binding, so the
		// removal cannot be verified against the ADC.
		Steps: []resource.TestStep{
			{
				Config: testAccServicegroupServicegroupmemberlistBinding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServicegroupServicegroupmemberlistBindingExist("citrixadc_servicegroup_servicegroupmemberlist_binding.tf_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_servicegroup_servicegroupmemberlist_binding.tf_binding", "servicegroupname", "tf_memberlist_svcgroup"),
					resource.TestCheckResourceAttr("citrixadc_servicegroup_servicegroupmemberlist_binding.tf_binding", "members.#", "2"),
					resource.TestCheckResourceAttr("citrixadc_servicegroup_servicegroupmemberlist_binding.tf_binding", "members.0.ip", "10.10.10.10"),
					resource.TestCheckResourceAttr("citrixadc_servicegroup_servicegroupmemberlist_binding.tf_binding", "members.0.port", "80"),
					resource.TestCheckResourceAttr("citrixadc_servicegroup_servicegroupmemberlist_binding.tf_binding", "members.0.weight", "10"),
					resource.TestCheckResourceAttr("citrixadc_servicegroup_servicegroupmemberlist_binding.tf_binding", "members.0.state", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_servicegroup_servicegroupmemberlist_binding.tf_binding", "members.0.order", "1"),
					resource.TestCheckResourceAttr("citrixadc_servicegroup_servicegroupmemberlist_binding.tf_binding", "members.1.ip", "10.10.10.11"),
					resource.TestCheckResourceAttr("citrixadc_servicegroup_servicegroupmemberlist_binding.tf_binding", "members.1.port", "80"),
					resource.TestCheckResourceAttr("citrixadc_servicegroup_servicegroupmemberlist_binding.tf_binding", "members.1.weight", "20"),
					resource.TestCheckResourceAttr("citrixadc_servicegroup_servicegroupmemberlist_binding.tf_binding", "members.1.state", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_servicegroup_servicegroupmemberlist_binding.tf_binding", "members.1.order", "2"),
				),
			},
		},
	})
}

func testAccCheckServicegroupServicegroupmemberlistBindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No servicegroup_servicegroupmemberlist_binding ID is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		// NOTE: This is a STATE-ONLY existence check. NITRO exposes NO get/get(all)/count
		// endpoint for servicegroup_servicegroupmemberlist_binding (write-only bulk binding,
		// Pattern 13), so there is no FindResource call to verify against the ADC. The
		// presence of a non-empty ID in Terraform state is the only observable signal.
		return nil
	}
}
