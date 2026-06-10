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

// NOTE on the lsnsession resource:
//   - Models the NITRO POST /lsnsession?action=flush endpoint.
//   - This is an ACTION-ONLY (one-shot side-effect) flush filter resource:
//     Create performs ActOnResource("lsnsession", payload, "flush") with the
//     optional filter selectors that are set; Read is a no-op (preserves
//     state), Update is a no-op (all attributes are RequiresReplace), and
//     Delete is a state-only removal. There is no GET-by-id endpoint that can
//     re-resolve "the flushed sessions", so the resource CANNOT be verified by
//     reading it back from the ADC.
//   - ALL attributes (nattype, clientname, network, netmask, network6, td,
//     natip, natport2, nodeid) are OPTIONAL flush filter selectors; none is
//     Required. nattype defaults to "NAT44" and netmask defaults to
//     "255.255.255.255".
//   - The Create uses a CONSTANT synthetic ID ("lsnsession") because there is
//     no get-by-name key on NITRO for this action-only resource.
//   - `netmask` has a Computed default of 255.255.255.255, so it is ALWAYS
//     present in the flush payload. The ADC requires `network` and `netmask` to
//     be supplied TOGETHER (errorcode 1093 "Argument pre-requisite missing
//     [netmask, network]" when netmask is present without network). The man
//     page (man flush lsn session) describes netmask as "Subnet mask for the IP
//     address specified by the network parameter", confirming the pairing. So
//     the config supplies a valid network/netmask pair (100.64.0.0/10, the
//     carrier-grade NAT / LSN client range) to make the flush filter
//     well-formed. A flush matching zero active sessions is a harmless no-op
//     returning 200.
//   - The Exist check below only verifies that the resource landed in Terraform
//     state with its synthetic ID ("lsnsession"); it does NOT (and cannot)
//     verify the flush side-effect via NITRO.
//   - There is no CheckDestroy: the flush action has no inverse on NITRO, and
//     there is no GET-by-id to confirm absence; Delete is a state-only removal.
//   - There is NO datasource for lsnsession (it was removed: no NITRO GET
//     endpoint), so no datasource test is generated.
//
// This mirrors the action-only test precedent in aaasession_test.go /
// clusterfiles_test.go / clusterpropstatus_test.go (state-only Exist check, no
// CheckDestroy, TestCheckResourceAttrSet on "id").

// Single apply step: all attributes are RequiresReplace, so there is no
// in-place update to exercise. The flush action is self-contained via a
// well-formed network/netmask pair (netmask has a Computed default, so network
// must accompany it; see note above).
const testAccLsnsession_basic = `
resource "citrixadc_lsnsession" "tf_lsnsession" {
  nattype = "NAT44"
  network = "100.64.0.0"
  netmask = "255.192.0.0"
}

`

func TestAccLsnsession_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// No CheckDestroy: the flush action has no inverse on NITRO and there is
		// no GET-by-id to confirm absence; Delete is a state-only removal.
		Steps: []resource.TestStep{
			{
				Config: testAccLsnsession_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsnsessionExist("citrixadc_lsnsession.tf_lsnsession", nil),
					// "id" is the synthetic state handle "lsnsession".
					resource.TestCheckResourceAttrSet("citrixadc_lsnsession.tf_lsnsession", "id"),
					// Assert only the filters actually set in HCL.
					resource.TestCheckResourceAttr("citrixadc_lsnsession.tf_lsnsession", "nattype", "NAT44"),
					resource.TestCheckResourceAttr("citrixadc_lsnsession.tf_lsnsession", "network", "100.64.0.0"),
					resource.TestCheckResourceAttr("citrixadc_lsnsession.tf_lsnsession", "netmask", "255.192.0.0"),
				),
			},
		},
	})
}

// testAccCheckLsnsessionExist is a state-only existence check.
//
// lsnsession is an action-only flush resource: Read is a no-op and there is no
// GET-by-id endpoint, so we CANNOT verify the flush via NITRO. We only assert
// that Terraform recorded the resource in state with a non-empty ID (which
// equals the synthetic "lsnsession" after a successful POST ?action=flush).
// This mirrors testAccCheckAaasessionExist.
func testAccCheckLsnsessionExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lsnsession ID is set")
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
