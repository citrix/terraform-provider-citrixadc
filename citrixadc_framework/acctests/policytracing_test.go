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

// NOTE on the policytracing resource:
//   - Models the NITRO POST /policytracing?action=clear endpoint, which clears
//     the captured policy-tracing records. The clear body is empty
//     ({"policytracing":{}}); the clear action takes no arguments, so the
//     RESOURCE schema exposes ONLY the synthetic, constant id "policytracing"
//     and has NO writable attributes.
//   - This is an ACTION-ONLY (one-shot side-effect) resource: Create performs the
//     clear action, Read is a no-op (preserves state), Update is a no-op, and
//     Delete is a state-only removal. There is NO add/update/delete endpoint, so
//     the resource CANNOT be verified by reading it back from the ADC. (NITRO does
//     expose get(all)/count, which backs the read-only datasource below, but
//     there is no GET-by-id to re-resolve "this clear".)
//   - The clear side-effect is NOT NITRO-verifiable: clearing the trace records
//     produces no queryable managed object keyed by this action.
//   - The Exist check below only verifies that the resource landed in Terraform
//     state with its synthetic constant id ("policytracing"); it does NOT (and
//     cannot) verify the clear side-effect via NITRO.
//   - There is NO CheckDestroy: the clear action has no inverse on NITRO, and
//     there is no GET-by-id to confirm absence; Delete is a state-only removal.
//
// This mirrors the action-only test precedent in aaasession_test.go,
// clusterpropstatus_test.go and clusterfiles_test.go (single apply step,
// state-only Exist check, no CheckDestroy, TestCheckResourceAttrSet on "id"),
// adapted for the clear action.

// Single apply step: the resource takes no attributes (clear has no arguments),
// so there is no in-place update to exercise.
const testAccPolicytracing_basic = `
resource "citrixadc_policytracing" "tf_policytracing" {
}

`

func TestAccPolicytracing_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// No CheckDestroy: the clear action has no inverse on NITRO and there is
		// no GET-by-id to confirm absence; Delete is a state-only removal.
		Steps: []resource.TestStep{
			{
				Config: testAccPolicytracing_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPolicytracingExist("citrixadc_policytracing.tf_policytracing", nil),
					// "id" is the synthetic constant state handle "policytracing".
					resource.TestCheckResourceAttrSet("citrixadc_policytracing.tf_policytracing", "id"),
				),
			},
		},
	})
}

// testAccCheckPolicytracingExist is a state-only existence check.
//
// policytracing is an action-only resource: Read is a no-op and there is no
// GET-by-id endpoint, so we CANNOT verify the clear via NITRO. We only assert
// that Terraform recorded the resource in state with a non-empty ID (which
// equals the synthetic constant "policytracing" after a successful
// POST ?action=clear). This mirrors testAccCheckClusterpropstatusExist /
// testAccCheckAaasessionExist.
func testAccCheckPolicytracingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No policytracing ID is set")
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

// TestAccPolicytracingDataSource_basic is SKIPPED by default.
//
// The policytracing datasource is backed by the NITRO get(all) endpoint, which
// returns the currently captured policy-tracing records. Its Read explicitly
// errors when the result set is empty ("policytracing returned empty array")
// and when no record matches the supplied filters ("no policytracing record
// matched the provided filters").
//
// On a clean testbed there are typically NO captured policy-tracing records
// (and the policytracing clear resource above actively wipes them), so the
// get(all) call returns an empty array and the datasource Read errors out.
// There is also no Terraform resource that *creates* trace records to set up a
// non-empty result deterministically — those records are produced as a runtime
// side-effect of live traffic flowing through policies. Running this test
// against an empty testbed would therefore fail spuriously.
//
// Following the repo precedent for show-only diagnostic datasources whose
// results are likely empty without out-of-band setup, the datasource test is
// gated behind t.Skip. To exercise it, generate policy-trace records on the
// target ADC first (e.g. drive traffic through traced policies) and then run
// this test with the skip removed.
func TestAccPolicytracingDataSource_basic(t *testing.T) {
	t.Skip("policytracing datasource get(all) errors on empty results; policy-trace records cannot be created via Terraform and are typically absent on a clean testbed. Generate trace records out-of-band and remove this skip to run.")

	const testAccPolicytracingDataSource_basic = `
data "citrixadc_policytracing" "tf_policytracing" {
}
`

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccPolicytracingDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.citrixadc_policytracing.tf_policytracing", "id"),
				),
			},
		},
	})
}
