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

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

// NOTE on the clustersync_force resource:
//   - Models the NITRO POST /clustersync?action=Force endpoint, which forces a
//     synchronization of the cluster configuration across nodes.
//   - This is a ZERO-ATTRIBUTE, ACTION-ONLY resource: Create performs the Force
//     action, Read/Update are no-ops, and Delete is a state-only removal. There is
//     NO get/add/set/delete endpoint, so the resource CANNOT be verified by reading
//     it back and it has NO datasource (Pattern 13).
//   - The Exist check below only verifies the resource landed in Terraform state
//     with its synthetic id ("clustersync_force"); it cannot verify the side
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

const testAccClustersyncForce_basic = `
resource "citrixadc_clustersync_force" "tf_clustersync" {
}
`

func TestAccClustersyncForce_basic(t *testing.T) {
	if adcTestbed != "CLUSTER" {
		t.Skipf("ADC testbed is %s. Expected CLUSTER.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t); skipUnlessSyncableClusterNode(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// No CheckDestroy: the Force action has no inverse on NITRO and there is no
		// GET-by-id to confirm absence; Delete is a state-only removal.
		Steps: []resource.TestStep{
			{
				Config: testAccClustersyncForce_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClustersyncForceExist("citrixadc_clustersync_force.tf_clustersync", nil),
					// "id" is the synthetic state handle "clustersync_force".
					resource.TestCheckResourceAttrSet("citrixadc_clustersync_force.tf_clustersync", "id"),
				),
			},
		},
	})
}

// testAccCheckClustersyncForceExist is a state-only existence check. clustersync_force is an
// action-only resource with no GET-by-id endpoint, so we only assert Terraform
// recorded the resource in state with a non-empty ID.
// skipUnlessSyncableClusterNode skips the test unless the ADC under test is a
// NON-coordinator cluster node. PREREQUISITE: 'force cluster sync' is only
// permitted on a non-coordinator node (one that syncs its config FROM the
// configuration coordinator); NITRO rejects it on the configuration coordinator
// and via the Cluster IP (CLIP) with errorcode 2478 "Operation not permitted".
// Detected by reading the local clusternode's isconfigurationcoordinator flag,
// so point NS_URL at a non-CCO cluster node (not the CLIP) to run this test.
func skipUnlessSyncableClusterNode(t *testing.T) {
	client, err := testAccGetFrameworkClient()
	if err != nil {
		t.Fatalf("Failed to get test client: %v", err)
	}
	nodes, err := client.FindAllResources(service.Clusternode.Type())
	if err != nil {
		t.Fatalf("Failed to read clusternode to check the force-sync prerequisite: %v", err)
	}
	for _, n := range nodes {
		if fmt.Sprintf("%v", n["islocalnode"]) != "true" {
			continue
		}
		if fmt.Sprintf("%v", n["isconfigurationcoordinator"]) == "true" {
			t.Skipf("Prerequisite not met: clustersync_force ('force cluster sync') is only permitted on a "+
				"non-coordinator cluster node; the target (%v) is the configuration coordinator / CLIP, where "+
				"NITRO rejects it with errorcode 2478. Point NS_URL at a non-CCO cluster node. Skipping.", n["ipaddress"])
		}
		return // local node is a non-coordinator cluster node -> prerequisite met
	}
	t.Skipf("Prerequisite not met: clustersync_force requires connecting directly to a cluster node, but no " +
		"local cluster node was found (NS_URL is likely the CLIP or a non-cluster box). Skipping.")
}

func testAccCheckClustersyncForceExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No clustersync_force ID is set")
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
