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

// NOTE on the appfwlearningdata_reset resource:
//   - Models the NITRO POST /appfwlearningdata?action=reset endpoint. Create
//     performs the action via ActOnResource(service.Appfwlearningdata.Type(),
//     &payload, "reset"), which CLEARS all App-Firewall learned-data databases and
//     zeroes the transaction count. The reset payload is EMPTY
//     ({"appfwlearningdata":{}}) and takes no arguments (confirmed by the NetScaler
//     CLI `reset appfw learningdata`), so the resource has NO configurable
//     attributes.
//   - Read/Update are no-ops (NITRO has no GET endpoint reporting reset-state) and
//     Delete is a state-only removal (the reset action has no inverse on NITRO).
//     The synthetic id is "appfwlearningdata_reset".
//   - This is an action-only resource: it CANNOT be verified by reading it back
//     from the ADC.
//
// The resource acceptance test is SKIPPED by default because applying it RESETS
// (destroys) whatever App-Firewall learning data has accumulated on the target ADC
// — a disruptive, non-reversible side effect that would corrupt any real learning
// run sharing the testbed. Remove the t.Skip only on a throwaway appliance.
//
// Mirrors the action-only test precedent (single apply step, state-only Exist
// check, no CheckDestroy, TestCheckResourceAttrSet on "id").

const testAccAppfwlearningdataReset_basic = `
resource "citrixadc_appfwlearningdata_reset" "tf_appfwlearningdata_reset" {
}

`

func TestAccAppfwlearningdataReset_basic(t *testing.T) {
	// The reset action clears App-Firewall learned data on the target ADC, which is
	// disruptive and non-reversible. Skip unless running against a throwaway box.
	t.Skip("Skipping appfwlearningdata resource test: the reset action clears (destroys) App-Firewall learning data on the ADC and is disruptive/non-reversible")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// No CheckDestroy: the reset action has no inverse on NITRO and there is no
		// GET-by-id to confirm absence; Delete is a state-only removal.
		Steps: []resource.TestStep{
			{
				Config: testAccAppfwlearningdataReset_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwlearningdataResetExist("citrixadc_appfwlearningdata_reset.tf_appfwlearningdata_reset", nil),
					// "id" is the synthetic state handle "appfwlearningdata_reset".
					resource.TestCheckResourceAttrSet("citrixadc_appfwlearningdata_reset.tf_appfwlearningdata_reset", "id"),
				),
			},
		},
	})
}

// testAccCheckAppfwlearningdataResetExist is a state-only existence check.
//
// appfwlearningdata_reset is an action-only resource: Read is a no-op and there is
// no GET-by-id endpoint, so we CANNOT verify the reset via NITRO. We only assert
// that Terraform recorded the resource in state with a non-empty ID (which equals
// the synthetic "appfwlearningdata_reset" after a successful POST ?action=reset).
func testAccCheckAppfwlearningdataResetExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No appfwlearningdata_reset ID is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		// No NITRO GET-by-id to verify against; presence of the synthetic state ID
		// is the only confirmation we can make for an action-only resource.
		return nil
	}
}

// The datasource test reads the App-Firewall learned-data table (get(all)) scoped
// to a profile/security check. The citrixadc_appfwlearningdata DATASOURCE is
// UNCHANGED by the reset/export action split, so this sub-test is preserved
// verbatim here (still targeting the citrixadc_appfwlearningdata data source type).
// It is SKIPPED by default because it requires an ADC with a configured
// App-Firewall profile that has actually accumulated learned data for the given
// security check; on a fresh testbed the table is empty and the read returns no
// rows.
const testAccAppfwlearningdataDataSource_basic = `
resource "citrixadc_appfwprofile" "tf_appfwprofile" {
  name = "tf_appfwlearningdata_ds_profile"
  type = ["HTML"]
}

data "citrixadc_appfwlearningdata" "tf_appfwlearningdata" {
  profilename   = citrixadc_appfwprofile.tf_appfwprofile.name
  securitycheck = "startURL"
  depends_on    = [citrixadc_appfwprofile.tf_appfwprofile]
}
`

func TestAccAppfwlearningdataDataSource_basic(t *testing.T) {
	t.Skip("Skipping appfwlearningdata datasource test: get(all) returns rows only when the ADC has accumulated App-Firewall learned data for the profile/security check")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAppfwlearningdataDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.citrixadc_appfwlearningdata.tf_appfwlearningdata", "id"),
				),
			},
		},
	})
}
