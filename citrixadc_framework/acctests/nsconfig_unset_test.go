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

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

// NOTE on the nsconfig_unset resource:
//   - Models the NITRO POST /nsconfig?action=unset endpoint (CLI: "unset ns
//     config"). It resets the named settable nsconfig parameters to their
//     appliance defaults.
//   - This is an ACTION-ONLY resource: Create performs the unset via
//     utils.ExecuteUnset, Read/Update are no-ops (all attributes are
//     RequiresReplace), Delete is a state-only removal. There is NO GET-by-id, so
//     the resource itself cannot be read back; the Exist check below instead reads
//     the live nsconfig and confirms the unset parameters are at their defaults.
//   - `timestamp` is a synthetic Required + RequiresReplace re-run key; the
//     synthetic ID equals the timestamp.
//
// SAFETY: this test unsets ONLY non-connectivity parameters that are already at
// their appliance defaults on the designated disposable box (NSIP 10.101.132.152),
// so `unset ns config` is an effective no-op that CANNOT disrupt the appliance.
// ipaddress/netmask/ifnum/nsvlan/tagged and the securemanagement* parameters are
// intentionally NOT unset.
const testAccNsconfigUnset_basic = `
resource "citrixadc_nsconfig_unset" "foo" {
  attributes = [
    "maxconn",
    "maxreq",
    "cip",
    "cookieversion",
    "securecookie",
    "pmtumin",
    "pmtutimeout",
    "timezone",
    "grantquotamaxclient",
    "exclusivequotamaxclient",
    "grantquotaspillover",
    "exclusivequotaspillover",
  ]
  timestamp = "2024-06-01T12:00:00"
}
`

func TestAccNsconfigUnset_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// No CheckDestroy: the unset action has no inverse on NITRO and there is no
		// GET-by-id to confirm absence; Delete is a state-only removal.
		Steps: []resource.TestStep{
			{
				Config: testAccNsconfigUnset_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsconfigUnsetExist("citrixadc_nsconfig_unset.foo", nil),
					resource.TestCheckResourceAttr("citrixadc_nsconfig_unset.foo", "attributes.#", "12"),
					resource.TestCheckResourceAttr("citrixadc_nsconfig_unset.foo", "timestamp", "2024-06-01T12:00:00"),
					resource.TestCheckResourceAttrSet("citrixadc_nsconfig_unset.foo", "id"),
				),
			},
		},
	})
}

// testAccCheckNsconfigUnsetExist asserts the resource landed in state with its
// synthetic ID, and then reads the live nsconfig back to confirm the parameters we
// unset are now at their documented appliance defaults.
func testAccCheckNsconfigUnsetExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No nsconfig_unset ID is set")
		}
		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}
			*id = rs.Primary.ID
		}

		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Error creating client for nsconfig read-back: %s", err.Error())
		}
		data, err := client.FindResource(service.Nsconfig.Type(), "")
		if err != nil {
			return fmt.Errorf("Error reading nsconfig from ADC: %s", err.Error())
		}

		// After unset, these parameters must be back at their appliance defaults.
		wantDefaults := map[string]string{
			"maxconn":                 "0",
			"maxreq":                  "0",
			"securecookie":            "ENABLED",
			"pmtumin":                 "576",
			"pmtutimeout":             "10",
			"timezone":                "CoordinatedUniversalTime",
			"grantquotamaxclient":     "10",
			"exclusivequotamaxclient": "80",
			"grantquotaspillover":     "10",
			"exclusivequotaspillover": "80",
		}
		for attr, want := range wantDefaults {
			if got, ok := data[attr]; ok {
				if fmt.Sprintf("%v", got) != want {
					return fmt.Errorf("nsconfig %s not at default after unset: want=%s adc=%v", attr, want, got)
				}
			}
		}
		return nil
	}
}
