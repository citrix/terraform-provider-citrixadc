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

// NOTE on the dnssubnetcache resource:
//   - Models the NITRO POST /dnssubnetcache?action=flush endpoint, which flushes
//     ECS (EDNS Client Subnet) subnets from the runtime DNS cache
//     (e.g. "flush dns subnetCache -all").
//   - This is an ACTION-ONLY (one-shot side-effect) resource (Pattern 13):
//     Create performs the flush action, Read is a no-op (preserves state),
//     Update is a no-op, and Delete is a state-only removal. There is NO
//     get/add/update/delete endpoint, so the resource CANNOT be verified by
//     reading it back from the ADC, and it has NO datasource (removed per
//     Pattern 13 — no NITRO GET endpoint).
//   - The two write inputs (all, ecssubnet) are OPTIONAL flush filters; neither
//     is Required, and both are RequiresReplace. The CLI wants exactly one of
//     all | ecssubnet: "all = true" flushes every ECS subnet from the cache;
//     "ecssubnet" flushes a single subnet. (nodeid is a GET-only cluster filter
//     — Pattern 15 — and is intentionally not modelled in the schema, so it is
//     not exercised here.)
//   - The basic test sets all = true (the simplest self-contained form; needs no
//     pre-existing ECS subnet entries on the testbed — flushing an empty cache is
//     a no-op success).
//   - The Exist check below only verifies that the resource landed in Terraform
//     state with its synthetic ID ("dnssubnetcache-flush"); it does NOT (and
//     cannot) verify the flush side-effect via NITRO.
//   - There is no CheckDestroy: the flush action has no inverse on NITRO, and
//     there is no GET-by-id to confirm absence; Delete is a state-only removal.
//
// This mirrors the action-only test precedent in dnsproxyrecords_test.go /
// lbpersistentsessions_test.go (single apply step, state-only Exist check,
// no CheckDestroy, TestCheckResourceAttrSet on "id"), adapted for the flush
// action.
//
// ENVIRONMENT: flush works on a standalone testbed; no ADC_TESTBED gate is
// needed.

// Single apply step: both attributes are RequiresReplace, so there is no
// in-place update to exercise. all = true flushes every ECS subnet from the
// DNS cache (simplest self-contained form, no pre-existing ECS subnet entries
// required).
const testAccDnssubnetcache_basic = `
resource "citrixadc_dnssubnetcache" "tf_dnssubnetcache" {
  all = true
}

`

func TestAccDnssubnetcache_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// No CheckDestroy: the flush action has no inverse on NITRO and there is
		// no GET-by-id to confirm absence; Delete is a state-only removal.
		Steps: []resource.TestStep{
			{
				Config: testAccDnssubnetcache_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDnssubnetcacheExist("citrixadc_dnssubnetcache.tf_dnssubnetcache", nil),
					resource.TestCheckResourceAttr("citrixadc_dnssubnetcache.tf_dnssubnetcache", "all", "true"),
					// "id" is the synthetic state handle "dnssubnetcache-flush".
					resource.TestCheckResourceAttrSet("citrixadc_dnssubnetcache.tf_dnssubnetcache", "id"),
				),
			},
		},
	})
}

// testAccCheckDnssubnetcacheExist is a state-only existence check.
//
// dnssubnetcache is an action-only resource: Read is a no-op and there is no
// GET-by-id endpoint, so we CANNOT verify the flush via NITRO. We only assert
// that Terraform recorded the resource in state with a non-empty ID (which
// equals the synthetic "dnssubnetcache-flush" after a successful POST
// ?action=flush). This mirrors testAccCheckDnsproxyrecordsExist.
func testAccCheckDnssubnetcacheExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No dnssubnetcache ID is set")
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
