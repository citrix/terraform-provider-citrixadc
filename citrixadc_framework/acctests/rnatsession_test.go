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

// NOTE on the rnatsession resource:
//   - Models the NITRO POST /rnatsession?action=flush endpoint, which flushes
//     RNAT (reverse NAT) sessions (CLI: "flush rnatSession").
//   - This is an ACTION-ONLY (one-shot side-effect) resource: Create performs the
//     flush action, Read is a no-op (preserves state), Update is a no-op (all
//     attributes are RequiresReplace), and Delete is a state-only removal. There
//     is NO get/add/update/delete endpoint, so the resource CANNOT be verified by
//     reading it back from the ADC, and it has NO datasource (removed per
//     Pattern 13 — no NITRO GET endpoint).
//   - All four attributes (network, netmask, natip, aclname) are OPTIONAL flush
//     filters; none is Required, and all are RequiresReplace. network/netmask
//     scope the flush to a subnet, natip filters by the RNAT NAT IP, and aclname
//     filters by a configured extended ACL. When all are omitted, the flush
//     clears ALL RNAT sessions. The basic test sets network + netmask, which the
//     ADC rejects with errorcode 775 ("ACL or NETWORK based RNAT rule does not
//     exist") UNLESS a matching NETWORK-based RNAT rule already exists for that
//     subnet. The test is therefore made self-contained: it first creates a
//     citrixadc_rnat NETWORK rule for the same subnet and the flush depends_on
//     it, so the scoped flush has a rule to act on.
//   - The Exist check below only verifies that the resource landed in Terraform
//     state with its synthetic ID ("rnatsession-config"); it does NOT (and
//     cannot) verify the flush side-effect via NITRO.
//   - There is no CheckDestroy: the flush action has no inverse on NITRO, and
//     there is no GET-by-id to confirm absence; Delete is a state-only removal.
//
// This mirrors the action-only test precedent in dnsproxyrecords_test.go /
// dnssubnetcache_test.go (single apply step, state-only Exist check, no
// CheckDestroy, TestCheckResourceAttrSet on "id"), adapted for the flush action.
//
// ENVIRONMENT: flush works on a standalone testbed; no ADC_TESTBED gate is
// needed.

// Single apply step: all attributes are RequiresReplace, so there is no in-place
// update to exercise. network + netmask flush RNAT sessions in the given subnet;
// the ADC requires a matching NETWORK-based RNAT rule to exist for that subnet,
// so we create a citrixadc_rnat rule first and have the flush depend on it. The
// rnat rule is a normal resource and is torn down automatically at end of test,
// leaving the appliance clean.
const testAccRnatsession_basic = `
resource "citrixadc_rnat" "tf_rnat_rnatsession" {
  name             = "tf_rnat_rnatsession"
  network          = "192.0.2.0"
  netmask          = "255.255.255.0"
  useproxyport     = "ENABLED"
  srcippersistency = "DISABLED"
  connfailover     = "DISABLED"
}

resource "citrixadc_rnatsession" "tf_rnatsession" {
  network = "192.0.2.0"
  netmask = "255.255.255.0"

  depends_on = [citrixadc_rnat.tf_rnat_rnatsession]
}

`

func TestAccRnatsession_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// No CheckDestroy: the flush action has no inverse on NITRO and there is
		// no GET-by-id to confirm absence; Delete is a state-only removal.
		Steps: []resource.TestStep{
			{
				Config: testAccRnatsession_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRnatsessionExist("citrixadc_rnatsession.tf_rnatsession", nil),
					resource.TestCheckResourceAttr("citrixadc_rnatsession.tf_rnatsession", "network", "192.0.2.0"),
					resource.TestCheckResourceAttr("citrixadc_rnatsession.tf_rnatsession", "netmask", "255.255.255.0"),
					// "id" is the synthetic state handle "rnatsession-config".
					resource.TestCheckResourceAttrSet("citrixadc_rnatsession.tf_rnatsession", "id"),
				),
			},
		},
	})
}

// testAccCheckRnatsessionExist is a state-only existence check.
//
// rnatsession is an action-only resource: Read is a no-op and there is no
// GET-by-id endpoint, so we CANNOT verify the flush via NITRO. We only assert
// that Terraform recorded the resource in state with a non-empty ID (which
// equals the synthetic "rnatsession-config" after a successful POST
// ?action=flush). This mirrors testAccCheckDnsproxyrecordsExist.
func testAccCheckRnatsessionExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No rnatsession ID is set")
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
