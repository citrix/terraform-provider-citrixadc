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

// NOTE on the clustersync resource:
//   - Models the NITRO POST /clustersync?action=Force endpoint, which forces a
//     synchronization of the cluster configuration across nodes.
//   - This is a ZERO-ATTRIBUTE, ACTION-ONLY resource: Create performs the Force
//     action, Read/Update are no-ops, and Delete is a state-only removal. There is
//     NO get/add/set/delete endpoint, so the resource CANNOT be verified by reading
//     it back and it has NO datasource (Pattern 13).
//   - The Exist check below only verifies the resource landed in Terraform state
//     with its synthetic id ("clustersync-config"); it cannot verify the side
//     effect via NITRO.
//
// !!! DESTRUCTIVE !!!
// Forcing a cluster sync overwrites the running configuration on cluster nodes
// with the configuration coordinator's config. On a standalone/non-cluster testbed
// it fails, and on a real cluster it can disrupt node state. The test is therefore
// SKIPPED by default. Remove the t.Skip line only when running intentionally
// against a disposable cluster.
//
// This mirrors the action-only test precedent (single apply step, state-only Exist
// check, no CheckDestroy, TestCheckResourceAttrSet on "id").

const testAccClustersync_basic = `
resource "citrixadc_clustersync" "tf_clustersync" {
}
`

func TestAccClustersync_basic(t *testing.T) {
	if adcTestbed != "CLUSTER" {
		t.Skipf("ADC testbed is %s. Expected CLUSTER.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// No CheckDestroy: the Force action has no inverse on NITRO and there is no
		// GET-by-id to confirm absence; Delete is a state-only removal.
		Steps: []resource.TestStep{
			{
				Config: testAccClustersync_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClustersyncExist("citrixadc_clustersync.tf_clustersync", nil),
					// "id" is the synthetic state handle "clustersync-config".
					resource.TestCheckResourceAttrSet("citrixadc_clustersync.tf_clustersync", "id"),
				),
			},
		},
	})
}

// testAccCheckClustersyncExist is a state-only existence check. clustersync is an
// action-only resource with no GET-by-id endpoint, so we only assert Terraform
// recorded the resource in state with a non-empty ID.
func testAccCheckClustersyncExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No clustersync ID is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		// No NITRO GET-by-id to verify against for an action-only resource.
		return nil
	}
}
