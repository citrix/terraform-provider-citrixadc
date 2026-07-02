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

// NOTE on the rdpconnections resource:
//   - rdpconnections is an ACTION-ONLY runtime object. NITRO exposes ONLY
//     get(all), count, and the POST action ?action=kill. There is NO add, NO
//     update/set, and NO delete. The connection table is populated by live RDP
//     proxy (Gateway/VPN RDP proxy) sessions, not by the config API.
//   - The RESOURCE therefore fires ?action=kill on Create (optionally scoped by
//     username, or all=true to kill every active connection) and treats
//     Read/Update/Delete as no-ops. There is NO GET-by-id endpoint, so the
//     resource CANNOT be verified by reading it back; the Exist check below only
//     asserts the synthetic Terraform state ID is set (mirrors
//     ipsecalgsession_test.go / lldpneighbors_test.go).
//   - There is NO CheckDestroy: kill has no inverse on NITRO and there is no
//     GET-by-id to confirm absence; Delete is a state-only removal (no-op).
//   - The resource has NO Required attributes. username/all are Optional and
//     RequiresReplace kill selectors.
//   - Synthetic ID: when username is set, id = "rdpconnections-kill-<username>";
//     otherwise (all=true or bare) id = "rdpconnections-kill-all".
//
// TODO_PLACEHOLDER (prereq): For RDP connections to ever appear on the appliance,
//   the RDP Proxy feature (NetScaler Gateway / VPN RDP proxy) must be configured
//   and live RDP proxy sessions must be in-use (rdpserverprofile / rdpclientprofile
//   bound to a VPN/Gateway vserver, with a user actually tunneling RDP traffic).
//   A kill with all=true succeeds regardless of whether any connections are
//   present, so the resource test below does NOT require this prereq. It only
//   matters for populating the datasource with a non-empty connection table.

// kill-all: all=true. This is the safest testable path — a kill-all succeeds
// without any prerequisite active RDP connections on the appliance.
const testAccRdpconnections_basic_step1 = `
resource "citrixadc_rdpconnections" "tf_rdpconnections" {
  all = true
}

`

func TestAccRdpconnections_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// No CheckDestroy: rdpconnections has no delete/GET-by-id endpoint (action-only).
		Steps: []resource.TestStep{
			{
				Config: testAccRdpconnections_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRdpconnectionsExist("citrixadc_rdpconnections.tf_rdpconnections", nil),
					// Synthetic ID recorded after a successful kill-all (?action=kill,
					// all=true, no username). Per the resource code this is
					// "rdpconnections-kill-all".
					resource.TestCheckResourceAttrSet("citrixadc_rdpconnections.tf_rdpconnections", "id"),
					resource.TestCheckResourceAttr("citrixadc_rdpconnections.tf_rdpconnections", "id", "rdpconnections-kill-all"),
					resource.TestCheckResourceAttr("citrixadc_rdpconnections.tf_rdpconnections", "all", "true"),
				),
			},
		},
	})
}

// testAccCheckRdpconnectionsExist is a state-only existence check.
//
// rdpconnections is an action-only resource: Read is a no-op and there is no
// GET-by-id endpoint, so we CANNOT verify the fired kill via NITRO. We only
// assert that Terraform recorded the resource in state with a non-empty ID
// (the synthetic "rdpconnections-kill-all" / "rdpconnections-kill-<username>"
// after a successful POST ?action=kill). This mirrors
// testAccCheckIpsecalgsessionExist / testAccCheckLldpneighborsExist.
func testAccCheckRdpconnectionsExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No rdpconnections ID is set")
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

// Datasource test: reads the RDP connection table via get(all) and filters
// locally on the optional username attribute.
//
// TODO_PLACEHOLDER (prereq): the RDP connection table is typically EMPTY on a
//
//	test appliance — it is populated only by live RDP proxy sessions (RDP Proxy
//	feature configured on a NetScaler Gateway/VPN vserver, with a user actively
//	tunneling RDP; see the resource note above). The datasource treats an empty
//	table (or no filter match) as VALID: it does NOT hard-fail, it sets the
//	synthetic id "rdpconnections" and returns. Therefore the only assertion made
//	below is that the datasource id is set (read succeeded). Add stricter
//	attribute assertions once a known connection is guaranteed present on the
//	testbed, e.g.:
//	    resource.TestCheckResourceAttrSet("data.citrixadc_rdpconnections.tf_rdpconnections", "endpointip"),
//	    resource.TestCheckResourceAttr("data.citrixadc_rdpconnections.tf_rdpconnections", "endpointip", "<TODO_PLACEHOLDER: live endpoint IP>"),
//	    resource.TestCheckResourceAttrSet("data.citrixadc_rdpconnections.tf_rdpconnections", "targetip"),
//	    resource.TestCheckResourceAttr("data.citrixadc_rdpconnections.tf_rdpconnections", "targetip", "<TODO_PLACEHOLDER: live target IP>"),
//	    resource.TestCheckResourceAttrSet("data.citrixadc_rdpconnections.tf_rdpconnections", "endpointport"),
//	    resource.TestCheckResourceAttrSet("data.citrixadc_rdpconnections.tf_rdpconnections", "targetport"),
//	    resource.TestCheckResourceAttrSet("data.citrixadc_rdpconnections.tf_rdpconnections", "peid"),
//	    // endpointip/endpointport/targetip/targetport/peid are testbed-specific and
//	    // non-deterministic — assert only once a known RDP connection is guaranteed.
const testAccRdpconnectionsDataSource_basic = `
data "citrixadc_rdpconnections" "tf_rdpconnections" {
  // No filters: match the first connection in the get(all) list. An empty
  // connection table is valid (see TODO_PLACEHOLDER above) and yields
  // id = "rdpconnections".
}

`

func TestAccRdpconnectionsDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccRdpconnectionsDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Read succeeded and the synthetic ID was composed. On an empty
					// connection table this is "rdpconnections".
					resource.TestCheckResourceAttrSet("data.citrixadc_rdpconnections.tf_rdpconnections", "id"),
					resource.TestCheckResourceAttr("data.citrixadc_rdpconnections.tf_rdpconnections", "id", "rdpconnections"),
					// TODO_PLACEHOLDER: add attribute assertions once a known RDP
					// connection is guaranteed on the testbed (see note above).
				),
			},
		},
	})
}
