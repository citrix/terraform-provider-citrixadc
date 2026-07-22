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

// NOTE on the rdpconnections_kill resource:
//   - rdpconnections_kill models the NITRO rdpconnections `?action=kill` action
//     (POST). kill is a one-shot side-effect action: NITRO exposes ONLY get(all),
//     count, and the POST action ?action=kill for the underlying rdpconnections
//     object. There is NO add, NO update/set, and NO delete. The connection table
//     is populated by live RDP proxy (Gateway/VPN RDP proxy) sessions, not by the
//     config API.
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
//   - Synthetic ID: the resource records a static state ID of "rdpconnections_kill".
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
const testAccRdpconnectionsKill_basic_step1 = `
resource "citrixadc_rdpconnections_kill" "tf_rdpconnections_kill" {
  all = true
}

`

func TestAccRdpconnectionsKill_basic(t *testing.T) {
	t.Skip("TODO: Requires review - the kill action terminates live RDP connections; unsafe/disruptive on a shared testbed")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// No CheckDestroy: rdpconnections_kill has no delete/GET-by-id endpoint (action-only).
		Steps: []resource.TestStep{
			{
				Config: testAccRdpconnectionsKill_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRdpconnectionsKillExist("citrixadc_rdpconnections_kill.tf_rdpconnections_kill", nil),
					// Synthetic ID recorded after a successful kill-all (?action=kill,
					// all=true, no username). Per the resource code this is the static
					// "rdpconnections_kill".
					resource.TestCheckResourceAttrSet("citrixadc_rdpconnections_kill.tf_rdpconnections_kill", "id"),
					resource.TestCheckResourceAttr("citrixadc_rdpconnections_kill.tf_rdpconnections_kill", "id", "rdpconnections_kill"),
					resource.TestCheckResourceAttr("citrixadc_rdpconnections_kill.tf_rdpconnections_kill", "all", "true"),
				),
			},
		},
	})
}

// testAccCheckRdpconnectionsKillExist is a state-only existence check.
//
// rdpconnections_kill is an action-only resource: Read is a no-op and there is
// no GET-by-id endpoint, so we CANNOT verify the fired kill via NITRO. We only
// assert that Terraform recorded the resource in state with a non-empty ID
// (the synthetic "rdpconnections_kill" after a successful POST ?action=kill).
// This mirrors testAccCheckIpsecalgsessionExist / testAccCheckLldpneighborsExist.
func testAccCheckRdpconnectionsKillExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No rdpconnections_kill ID is set")
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
