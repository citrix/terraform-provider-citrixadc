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

// NOTE on the ipsecalgsession resource:
//   - ipsecalgsession is an ACTION-ONLY runtime object. NITRO exposes ONLY
//     get(all), count, and the POST action ?action=flush. There is NO add, NO
//     update/set, and NO delete. The session table is populated by the IPSec ALG
//     traffic engine, not by the config API.
//   - The RESOURCE therefore fires ?action=flush on Create (optionally scoped by
//     sourceip/natip/destip; a BARE flush with no scope flushes ALL sessions and
//     is always valid) and treats Read/Update/Delete as no-ops. There is NO
//     GET-by-id endpoint, so the resource CANNOT be verified by reading it back;
//     the Exist check below only asserts the synthetic Terraform state ID is set
//     (mirrors cacheobject_test.go / gslbconfig_test.go).
//   - There is NO CheckDestroy: flush has no inverse on NITRO and there is no
//     GET-by-id to confirm absence; Delete is a state-only removal.
//   - The resource has NO Required attributes. All scope attributes are Optional
//     and RequiresReplace.
//
// TODO_PLACEHOLDER (prereq): The IPSec ALG feature may need to be licensed and/or
//   enabled/in-use on the target ADC. A bare flush-all is expected to succeed
//   regardless of session-table contents, but if the feature is entirely absent
//   the flush action may error. This cannot be asserted from the test; ensure the
//   IPSec ALG feature is available on the testbed out-of-band if the flush errors.
//
// TODO_PLACEHOLDER (variant not covered): a SCOPED flush (sourceip/natip/destip)
//   targeting a live session requires an actual IPSec ALG session to exist on the
//   appliance (created by real IPSec ALG traffic). Those IP values are
//   testbed-specific and non-deterministic, so the scoped-flush variant is
//   intentionally omitted. To exercise it manually, drive IPSec ALG traffic, read
//   an active session's sourceip from the datasource, then apply:
//       resource "citrixadc_ipsecalgsession" "flush_one" {
//         sourceip = <TODO_PLACEHOLDER: live source IP, e.g. "10.0.0.5">
//       }

// flush-all: no scope attributes set. This is the safest testable path — a bare
// flush succeeds without any prerequisite IPSec ALG sessions on the appliance.
const testAccIpsecalgsession_basic_step1 = `
resource "citrixadc_ipsecalgsession" "tf_ipsecalgsession" {
}

`

func TestAccIpsecalgsession_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// No CheckDestroy: ipsecalgsession has no delete/GET-by-id endpoint (action-only).
		Steps: []resource.TestStep{
			{
				Config: testAccIpsecalgsession_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpsecalgsessionExist("citrixadc_ipsecalgsession.tf_ipsecalgsession", nil),
					// Synthetic ID for a bare flush-all is "flush-all".
					resource.TestCheckResourceAttrSet("citrixadc_ipsecalgsession.tf_ipsecalgsession", "id"),
					resource.TestCheckResourceAttr("citrixadc_ipsecalgsession.tf_ipsecalgsession", "id", "flush-all"),
				),
			},
		},
	})
}

// testAccCheckIpsecalgsessionExist is a state-only existence check.
//
// ipsecalgsession is an action-only resource: Read is a no-op and there is no
// GET-by-id endpoint, so we CANNOT verify the fired flush via NITRO. We only
// assert that Terraform recorded the resource in state with a non-empty ID
// (which equals the synthetic "flush-all" / "flush:<scope>:<ip>" after a
// successful POST ?action=flush). This mirrors testAccCheckCacheobjectExist.
func testAccCheckIpsecalgsessionExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ipsecalgsession ID is set")
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

// Datasource test: reads the IPSec ALG session table via get(all) and filters
// locally on the optional sourceip/natip/destip attributes.
//
// TODO_PLACEHOLDER (prereq): the IPSec ALG session table is typically EMPTY on a
//
//	test appliance — it is populated only by live IPSec ALG traffic. The
//	datasource treats an empty table (or no filter match) as VALID: it does NOT
//	hard-fail, it sets the synthetic id "ipsecalgsession" and returns. Therefore
//	the only assertion made below is that the datasource id is set (read
//	succeeded). Add stricter attribute assertions once a known session is
//	guaranteed present on the testbed, e.g.:
//	    resource.TestCheckResourceAttrSet("data.citrixadc_ipsecalgsession.tf_ipsecalgsession", "sourceip"),
//	    resource.TestCheckResourceAttr("data.citrixadc_ipsecalgsession.tf_ipsecalgsession", "sourceip", "<TODO_PLACEHOLDER: live source IP>"),
const testAccIpsecalgsessionDataSource_basic = `
data "citrixadc_ipsecalgsession" "tf_ipsecalgsession" {
  // No filters: match the first session in the get(all) list. An empty session
  // table is valid (see TODO_PLACEHOLDER above) and yields id = "ipsecalgsession".
}

`

func TestAccIpsecalgsessionDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIpsecalgsessionDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Read succeeded and the synthetic ID was composed. On an empty
					// session table this is "ipsecalgsession".
					resource.TestCheckResourceAttrSet("data.citrixadc_ipsecalgsession.tf_ipsecalgsession", "id"),
					resource.TestCheckResourceAttr("data.citrixadc_ipsecalgsession.tf_ipsecalgsession", "id", "ipsecalgsession"),
					// TODO_PLACEHOLDER: add attribute assertions once a known IPSec
					// ALG session is guaranteed on the testbed (see note above).
				),
			},
		},
	})
}
