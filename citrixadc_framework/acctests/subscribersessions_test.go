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

// NOTE on the subscribersessions resource:
//   - subscribersessions is an ACTION-ONLY runtime object. NITRO exposes ONLY
//     get(all), count, and the POST action ?action=clear. There is NO add, NO
//     update/set, and NO delete. The subscriber session table is populated by
//     the live Subscriber/Gx/PCRF (Telco) engine, not by the config API.
//   - The RESOURCE therefore fires ?action=clear on Create (optionally scoped by
//     ip and/or vlan; a BARE clear with no scope flushes the ENTIRE subscriber
//     session DB and is always valid) and treats Read/Update/Delete as no-ops.
//     There is NO GET-by-id endpoint, so the resource CANNOT be verified by
//     reading it back; the Exist check below only asserts the synthetic Terraform
//     state ID is set (mirrors ipsecalgsession_test.go / rdpconnections_test.go).
//   - There is NO CheckDestroy: clear has no inverse on NITRO and there is no
//     GET-by-id to confirm absence; Delete is a state-only removal (no-op).
//   - The resource has NO Required attributes. ip (String) and vlan (Int64) are
//     Optional RequiresReplace clear selectors.
//   - Synthetic ID (from the resource code Create):
//       ip+vlan set  -> "subscribersessions-clear-<ip>-<vlan>"
//       ip only      -> "subscribersessions-clear-<ip>"
//       vlan only    -> "subscribersessions-clear-vlan-<vlan>"
//       bare clear   -> "subscribersessions-clear-all"
//
// !!! CAUTION — DESTRUCTIVE ACTION !!!
// TODO_PLACEHOLDER (destructive-action warning): The basic test below applies a
//   BARE clear (no ip/vlan), which flushes the ENTIRE subscriber session DB on
//   the target ADC. This is SAFE on a fresh/test appliance whose subscriber DB
//   is empty, but DO NOT run this test against any appliance carrying live
//   subscriber sessions you care about — every session will be cleared. To scope
//   the clear to a single subscriber instead, set ip (and optionally vlan) to a
//   live subscriber's values (see the scoped variant TODO below).
//
// TODO_PLACEHOLDER (prereq): The Subscriber / Gx / PCRF (Telco) feature must be
//   LICENSED and ENABLED on the target ADC for subscriber sessions to ever exist.
//   A bare clear is expected to succeed regardless of session-table contents, so
//   the resource test below does NOT strictly require this prereq; it only
//   matters for populating the datasource with a non-empty session table.
//
// TODO_PLACEHOLDER (variant not covered): a SCOPED clear (ip / vlan) targeting a
//   live session requires an actual subscriber session to exist on the appliance
//   (created by real Gx/PCRF subscriber traffic). Those ip/vlan values are
//   testbed-specific and non-deterministic, so the scoped-clear variant is
//   intentionally omitted. To exercise it manually, drive subscriber traffic,
//   read an active session's ip from the datasource, then apply:
//       resource "citrixadc_subscribersessions" "clear_one" {
//         ip   = <TODO_PLACEHOLDER: live subscriber IP, e.g. "10.0.0.5">
//         vlan = <TODO_PLACEHOLDER: live subscriber vlan, e.g. 10>
//       }

// clear-all: no scope attributes set. This is the safest testable path on a
// fresh appliance — a bare clear succeeds without any prerequisite subscriber
// sessions. See the DESTRUCTIVE ACTION caution above before running.
const testAccSubscribersessions_basic_step1 = `
resource "citrixadc_subscribersessions" "tf_subscribersessions" {
}

`

func TestAccSubscribersessions_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// No CheckDestroy: subscribersessions has no delete/GET-by-id endpoint (action-only).
		Steps: []resource.TestStep{
			{
				Config: testAccSubscribersessions_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSubscribersessionsExist("citrixadc_subscribersessions.tf_subscribersessions", nil),
					// Synthetic ID recorded after a successful bare clear (?action=clear,
					// no ip/vlan). Per the resource code this is "subscribersessions-clear-all".
					resource.TestCheckResourceAttrSet("citrixadc_subscribersessions.tf_subscribersessions", "id"),
					resource.TestCheckResourceAttr("citrixadc_subscribersessions.tf_subscribersessions", "id", "subscribersessions-clear-all"),
				),
			},
		},
	})
}

// testAccCheckSubscribersessionsExist is a state-only existence check.
//
// subscribersessions is an action-only resource: Read is a no-op and there is no
// GET-by-id endpoint, so we CANNOT verify the fired clear via NITRO. We only
// assert that Terraform recorded the resource in state with a non-empty ID (the
// synthetic "subscribersessions-clear-all" / "subscribersessions-clear-<scope>"
// after a successful POST ?action=clear). This mirrors
// testAccCheckIpsecalgsessionExist / testAccCheckRdpconnectionsExist.
func testAccCheckSubscribersessionsExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No subscribersessions ID is set")
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

// Datasource test: reads the subscriber session table via get(all) and filters
// locally on the optional ip/vlan/nodeid attributes.
//
// TODO_PLACEHOLDER (prereq): the subscriber session table is typically EMPTY on
//
//	a test appliance — it is populated only by live Gx/PCRF subscriber traffic
//	with the Subscriber/Gx/PCRF (Telco) feature licensed and enabled (see the
//	resource note above). The datasource treats an empty table (or no filter
//	match) as VALID: it does NOT hard-fail, it sets the synthetic id
//	"subscribersessions" and returns. Therefore the only assertion made below is
//	that the datasource id is set (read succeeded). Add stricter attribute
//	assertions once a known session is guaranteed present on the testbed, e.g.:
//	    resource.TestCheckResourceAttrSet("data.citrixadc_subscribersessions.tf_subscribersessions", "subscriptionidvalue"),
//	    resource.TestCheckResourceAttr("data.citrixadc_subscribersessions.tf_subscribersessions", "subscriptionidvalue", "<TODO_PLACEHOLDER: live subscription-id value>"),
//	    resource.TestCheckResourceAttrSet("data.citrixadc_subscribersessions.tf_subscribersessions", "subscriptionidtype"),
//	    resource.TestCheckResourceAttrSet("data.citrixadc_subscribersessions.tf_subscribersessions", "servicepath"),
//	    resource.TestCheckResourceAttrSet("data.citrixadc_subscribersessions.tf_subscribersessions", "ip"),
//	    // subscriptionidtype/subscriptionidvalue/subscriberrules/flags/ttl/idlettl/
//	    // avpdisplaybuffer/servicepath are testbed-specific and non-deterministic —
//	    // assert only once a known subscriber session is guaranteed present.
const testAccSubscribersessionsDataSource_basic = `
data "citrixadc_subscribersessions" "tf_subscribersessions" {
  // No filters: match the first session in the get(all) list. An empty session
  // table is valid (see TODO_PLACEHOLDER above) and yields id = "subscribersessions".
}

`

func TestAccSubscribersessionsDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSubscribersessionsDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Read succeeded and the synthetic ID was composed. On an empty
					// session table this is "subscribersessions".
					resource.TestCheckResourceAttrSet("data.citrixadc_subscribersessions.tf_subscribersessions", "id"),
					resource.TestCheckResourceAttr("data.citrixadc_subscribersessions.tf_subscribersessions", "id", "subscribersessions"),
					// TODO_PLACEHOLDER: add attribute assertions once a known subscriber
					// session is guaranteed present on the testbed (see note above).
				),
			},
		},
	})
}
