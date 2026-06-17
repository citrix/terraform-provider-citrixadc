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

// !!! DANGER -- DO NOT RUN IN AUTOMATION !!!
//
// nsaptlicense is an ACTION-ONLY resource whose Create performs the NITRO
// `change` action (POST ?action=update). This action is DISRUPTIVE and
// NON-IDEMPOTENT: it ALLOCATES/CONSUMES pooled CADS license counts from a
// license server. Re-running it re-allocates licenses.
//
// Running this test will draw down real pooled licenses and requires a live
// license-server session (a valid `id` + `sessionid` obtained from that
// session). For that reason every test below is t.Skip-gated. Only un-skip and
// run it manually, against a disposable license pool, with real values
// substituted for the TODO_PLACEHOLDER fields.

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// Required: id, sessionid, bindtype, countavailable. serialno is the GET key
// used to read the allocated record back. All values must come from a real
// license-server session.
const testAccNsaptlicense_basic_step1 = `
resource "citrixadc_nsaptlicense" "tf_nsaptlicense" {
  id             = "TODO_PLACEHOLDER"
  sessionid      = "TODO_PLACEHOLDER"
  bindtype       = "TODO_PLACEHOLDER"
  countavailable = "TODO_PLACEHOLDER"
  serialno       = "TODO_PLACEHOLDER"
}

`

func TestAccNsaptlicense_basic(t *testing.T) {
	// DANGER: consumes pooled CADS licenses; must NOT run in automation.
	t.Skip("DANGER: nsaptlicense ALLOCATES pooled CADS licenses (DISRUPTIVE, non-idempotent) and needs a real id/sessionid from a license-server session. Run manually only.")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// Action-only resource: no CheckDestroy (NITRO exposes no delete endpoint;
		// allocated licenses remain on the appliance).
		Steps: []resource.TestStep{
			{
				Config: testAccNsaptlicense_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsaptlicenseExist("citrixadc_nsaptlicense.tf_nsaptlicense", nil),
					resource.TestCheckResourceAttrSet("citrixadc_nsaptlicense.tf_nsaptlicense", "id"),
					resource.TestCheckResourceAttr("citrixadc_nsaptlicense.tf_nsaptlicense", "id", "TODO_PLACEHOLDER"),
					resource.TestCheckResourceAttr("citrixadc_nsaptlicense.tf_nsaptlicense", "bindtype", "TODO_PLACEHOLDER"),
					resource.TestCheckResourceAttr("citrixadc_nsaptlicense.tf_nsaptlicense", "countavailable", "TODO_PLACEHOLDER"),
				),
			},
		},
	})
}

// testAccCheckNsaptlicenseExist is a STATE-ONLY check. Although nsaptlicense
// has a GET (array, filtered by serialno), the test is fully skip-gated and
// the allocation is disruptive, so we only confirm a non-empty synthetic ID
// here rather than issuing additional NITRO calls.
func testAccCheckNsaptlicenseExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No nsaptlicense ID is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		return nil
	}
}

// Datasource test for nsaptlicense. Gated with t.Skip because it requires a
// real allocated license record (filtered by serialno) to exist on the
// appliance, which in turn requires a disruptive allocation.
const testAccNsaptlicenseDataSource_basic = `

data "citrixadc_nsaptlicense" "tf_nsaptlicense" {
  serialno = "TODO_PLACEHOLDER"
}
`

func TestAccNsaptlicenseDataSource_basic(t *testing.T) {
	t.Skip("DANGER: nsaptlicense datasource needs a real allocated license record (serialno) on the appliance, which requires a disruptive pooled-license allocation. Run manually only.")
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNsaptlicenseDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_nsaptlicense.tf_nsaptlicense", "serialno", "TODO_PLACEHOLDER"),
				),
			},
		},
	})
}
