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

// NOTE on the dnsproxyrecords resource:
//   - Models the NITRO POST /dnsproxyrecords?action=flush endpoint, which flushes
//     DNS proxy records from the cache (e.g. "flush dns proxyRecords -type A").
//   - This is an ACTION-ONLY (one-shot side-effect) resource: Create performs the
//     flush action, Read is a no-op (preserves state), Update is a no-op, and
//     Delete is a state-only removal. There is NO get/add/update/delete endpoint,
//     so the resource CANNOT be verified by reading it back from the ADC, and it
//     has NO datasource (removed per Pattern 13 — no NITRO GET endpoint).
//   - Both attributes (type, negrectype) are OPTIONAL flush filters; neither is
//     Required, and both are RequiresReplace. type filters by record type
//     (A, NS, CNAME, SOA, MX, AAAA, SRV, RRSIG, NSEC, DNSKEY, PTR, TXT, NAPTR,
//     CAA); negrectype filters negative records (NXDOMAIN, NODATA). When both are
//     omitted, the flush clears ALL proxy records. The basic test sets
//     type = "A" (the simplest self-contained form; needs no pre-existing DNS
//     records on the testbed — flushing an empty type is a no-op success).
//   - The Exist check below only verifies that the resource landed in Terraform
//     state with its synthetic ID ("dnsproxyrecords-config"); it does NOT (and
//     cannot) verify the flush side-effect via NITRO.
//   - There is no CheckDestroy: the flush action has no inverse on NITRO, and
//     there is no GET-by-id to confirm absence; Delete is a state-only removal.
//
// This mirrors the action-only test precedent in lbpersistentsessions_test.go
// (single apply step, state-only Exist check, no CheckDestroy,
// TestCheckResourceAttrSet on "id"), adapted for the flush action.
//
// ENVIRONMENT: flush works on a standalone testbed; no ADC_TESTBED gate is
// needed.

// Single apply step: both attributes are RequiresReplace, so there is no
// in-place update to exercise. type = "A" flushes only A records (simplest
// self-contained form, no pre-existing DNS proxy records required).
const testAccDnsproxyrecords_basic = `
resource "citrixadc_dnsproxyrecords" "tf_dnsproxyrecords" {
  type = "A"
}

`

func TestAccDnsproxyrecords_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// No CheckDestroy: the flush action has no inverse on NITRO and there is
		// no GET-by-id to confirm absence; Delete is a state-only removal.
		Steps: []resource.TestStep{
			{
				Config: testAccDnsproxyrecords_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDnsproxyrecordsExist("citrixadc_dnsproxyrecords.tf_dnsproxyrecords", nil),
					resource.TestCheckResourceAttr("citrixadc_dnsproxyrecords.tf_dnsproxyrecords", "type", "A"),
					// "id" is the synthetic state handle "dnsproxyrecords-config".
					resource.TestCheckResourceAttrSet("citrixadc_dnsproxyrecords.tf_dnsproxyrecords", "id"),
				),
			},
		},
	})
}

// testAccCheckDnsproxyrecordsExist is a state-only existence check.
//
// dnsproxyrecords is an action-only resource: Read is a no-op and there is no
// GET-by-id endpoint, so we CANNOT verify the flush via NITRO. We only assert
// that Terraform recorded the resource in state with a non-empty ID (which
// equals the synthetic "dnsproxyrecords-config" after a successful POST
// ?action=flush). This mirrors testAccCheckLbpersistentsessionsExist.
func testAccCheckDnsproxyrecordsExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No dnsproxyrecords ID is set")
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
