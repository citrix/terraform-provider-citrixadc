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

// NOTE on the lldpneighbors_clear resource:
//   - lldpneighbors is an ACTION-ONLY runtime object. NITRO exposes ONLY
//     get(all), get-by-name, count, and the POST action ?action=clear. There is
//     NO add, NO update/set, and NO per-neighbor delete. The neighbor table is
//     populated by the LLDP discovery engine (LLDP frames received from directly
//     connected peers), not by the config API.
//   - The lldpneighbors_clear RESOURCE fires ?action=clear on Create with an
//     EMPTY payload ({"lldpneighbors":{}}) and NO args, and treats
//     Read/Update/Delete as no-ops. It records a FIXED synthetic Terraform state
//     ID "lldpneighbors_clear". Because there is no GET-by-id endpoint, the
//     resource CANNOT be verified by reading it back; the Exist check below only
//     asserts the synthetic state ID is set (mirrors ipsecalgsession_test.go /
//     cacheobject_test.go).
//   - There is NO CheckDestroy: clear has no inverse on NITRO and there is no
//     GET-by-id to confirm absence; Delete is a state-only removal (no-op).
//   - The resource has NO Required attributes. ifnum/nodeid are GET filters
//     (used only by the datasource), not clear-action args; on the resource they
//     are plain Optional inputs.
//
// TODO_PLACEHOLDER (prereq): For LLDP neighbors to ever appear on the appliance,
//   LLDP must be enabled per-interface (LLDP mode set to send/receive) and the
//   ADC must be directly connected to an LLDP-capable peer that is transmitting
//   LLDP frames, e.g.:
//       set lldpparam -mode TRANSMITTER RECEIVER   (or global holdtimeTxMult)
//       set interface <ifnum> -lldpmode TRANSMITTER RECEIVER
//   A bare ?action=clear succeeds regardless of whether any neighbors are present,
//   so the resource test below does NOT require this prereq. It only matters for
//   populating the datasource with a non-empty neighbor table.

// clear-all: no attributes set. This is the safest testable path — a bare
// ?action=clear (empty payload, no args) succeeds without any prerequisite LLDP
// neighbors on the appliance.
const testAccLldpneighborsClear_basic_step1 = `
resource "citrixadc_lldpneighbors_clear" "tf_lldpneighbors_clear" {
}

`

func TestAccLldpneighborsClear_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// No CheckDestroy: lldpneighbors_clear has no delete/GET-by-id endpoint (action-only).
		Steps: []resource.TestStep{
			{
				Config: testAccLldpneighborsClear_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLldpneighborsClearExist("citrixadc_lldpneighbors_clear.tf_lldpneighbors_clear", nil),
					// Fixed synthetic ID recorded after a successful ?action=clear.
					resource.TestCheckResourceAttrSet("citrixadc_lldpneighbors_clear.tf_lldpneighbors_clear", "id"),
					resource.TestCheckResourceAttr("citrixadc_lldpneighbors_clear.tf_lldpneighbors_clear", "id", "lldpneighbors_clear"),
				),
			},
		},
	})
}

// testAccCheckLldpneighborsClearExist is a state-only existence check.
//
// lldpneighbors_clear is an action-only resource: Read is a no-op and there is
// no GET-by-id endpoint, so we CANNOT verify the fired clear via NITRO. We only
// assert that Terraform recorded the resource in state with a non-empty ID
// (the fixed synthetic "lldpneighbors_clear" after a successful POST
// ?action=clear). This mirrors testAccCheckIpsecalgsessionExist.
func testAccCheckLldpneighborsClearExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lldpneighbors_clear ID is set")
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

// Datasource test: reads the LLDP neighbor table via get(all) and filters
// locally on the optional ifnum attribute (nodeid is an optional cluster-node
// GET filter).
//
// The datasource type citrixadc_lldpneighbors is UNCHANGED by the clear-action
// refactor, so this sub-test is preserved verbatim from the original
// lldpneighbors_test.go.
//
// TODO_PLACEHOLDER (prereq): the LLDP neighbor table is typically EMPTY on a
//
//	test appliance — it is populated only when LLDP is enabled per-interface and
//	an LLDP-capable peer is directly connected and transmitting LLDP frames (see
//	the resource note above). The datasource treats an empty table (or no filter
//	match) as VALID: it does NOT hard-fail, it sets the synthetic id
//	"lldpneighbors" and returns. Therefore the only assertion made below is that
//	the datasource id is set (read succeeded). Add stricter attribute assertions
//	once a known neighbor is guaranteed present on the testbed, e.g.:
//	    resource.TestCheckResourceAttrSet("data.citrixadc_lldpneighbors.tf_lldpneighbors", "ifnum"),
//	    resource.TestCheckResourceAttr("data.citrixadc_lldpneighbors.tf_lldpneighbors", "ifnum", "<TODO_PLACEHOLDER: live interface, e.g. 1/1>"),
//	    // chassisid, portid, systemname, etc. are testbed-specific and non-
//	    // deterministic — assert only once a known LLDP peer is guaranteed.
const testAccLldpneighborsDataSource_basic = `
data "citrixadc_lldpneighbors" "tf_lldpneighbors" {
  // No filters: match the first neighbor in the get(all) list. An empty neighbor
  // table is valid (see TODO_PLACEHOLDER above) and yields id = "lldpneighbors".
}

`

func TestAccLldpneighborsDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccLldpneighborsDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Read succeeded and the synthetic ID was composed. On an empty
					// neighbor table this is "lldpneighbors".
					resource.TestCheckResourceAttrSet("data.citrixadc_lldpneighbors.tf_lldpneighbors", "id"),
					resource.TestCheckResourceAttr("data.citrixadc_lldpneighbors.tf_lldpneighbors", "id", "lldpneighbors"),
					// TODO_PLACEHOLDER: add attribute assertions once a known LLDP
					// neighbor is guaranteed on the testbed (see note above).
				),
			},
		},
	})
}
