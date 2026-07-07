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

// NOTE on the appfwlearningdata resource/datasource:
//   - appfwlearningdata is the Application-Firewall LEARNED-DATA table. It is NOT a
//     normal CRUD object: NITRO exposes get(all), count, delete, and the reset/export
//     actions only. This provider models it BEST-EFFORT as an ACTION resource:
//     Create performs the "reset" action (POST /appfwlearningdata?action=reset, which
//     CLEARS learned data for the given profile/security check), Read/Update are
//     no-ops, and Delete is a state-only removal (the NITRO delete endpoint carries
//     no ?args= key selector, so no delete is issued). The synthetic id is
//     "appfwlearningdata-config".
//   - The reset/delete semantics are best-effort and should be verified on a live ADC.
//
// The resource acceptance test is SKIPPED by default because applying it RESETS
// (destroys) whatever App-Firewall learning data has accumulated on the target ADC
// — a disruptive, non-reversible side effect that would corrupt any real learning
// run sharing the testbed. Remove the t.Skip only on a throwaway appliance.
//
// Mirrors the action-only test precedent in streamsession_test.go (single apply
// step, state-only Exist check, no CheckDestroy, TestCheckResourceAttrSet on "id").

const testAccAppfwlearningdata_basic = `
resource "citrixadc_appfwprofile" "tf_appfwprofile" {
  name = "tf_appfwlearningdata_profile"
  type = ["HTML"]
}

resource "citrixadc_appfwlearningdata" "tf_appfwlearningdata" {
  profilename   = citrixadc_appfwprofile.tf_appfwprofile.name
  securitycheck = "startURL"
  starturl      = "^https?://[^/]+/$"

  depends_on = [citrixadc_appfwprofile.tf_appfwprofile]
}

`

func TestAccAppfwlearningdata_basic(t *testing.T) {
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
				Config: testAccAppfwlearningdata_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwlearningdataExist("citrixadc_appfwlearningdata.tf_appfwlearningdata", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwlearningdata.tf_appfwlearningdata", "securitycheck", "startURL"),
					// "id" is the synthetic state handle "appfwlearningdata-config".
					resource.TestCheckResourceAttrSet("citrixadc_appfwlearningdata.tf_appfwlearningdata", "id"),
				),
			},
		},
	})
}

// testAccCheckAppfwlearningdataExist is a state-only existence check.
//
// appfwlearningdata is an action resource: Read is a no-op and there is no
// GET-by-id endpoint, so we CANNOT verify the reset via NITRO. We only assert that
// Terraform recorded the resource in state with a non-empty ID (which equals the
// synthetic "appfwlearningdata-config" after a successful POST ?action=reset).
func testAccCheckAppfwlearningdataExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No appfwlearningdata ID is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		// No NITRO GET-by-id to verify against; presence of the synthetic state ID
		// is the only confirmation we can make for an action resource.
		return nil
	}
}

// The datasource test reads the App-Firewall learned-data table (get(all)) scoped
// to a profile/security check. It is SKIPPED by default because it requires an ADC
// with a configured App-Firewall profile that has actually accumulated learned data
// for the given security check; on a fresh testbed the table is empty and the read
// returns no rows.
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
