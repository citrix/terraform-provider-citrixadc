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
// nsaptlicense_change is an ACTION-ONLY resource whose Create performs the
// NITRO `update` action (POST ?action=update). This action is DISRUPTIVE and
// NON-IDEMPOTENT: it ALLOCATES/CONSUMES pooled CADS license counts from a
// license server and mutates APT license bindings. Re-running it re-allocates
// licenses.
//
// Running this test will draw down real pooled licenses and requires a live
// license-server session (a valid `id` + `sessionid` obtained from that
// session) plus a real APT license file/serial. For that reason the action
// test below is t.Skip-gated. Only un-skip and run it manually, against a
// disposable license pool, with real values substituted for the
// TODO_PLACEHOLDER fields.

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// Required: id, sessionid, bindtype, countavailable. serialno is the GET key
// used to read the allocated record back. All values must come from a real
// license-server session.
const testAccNsaptlicenseChange_basic_step1 = `
resource "citrixadc_nsaptlicense_change" "tf_nsaptlicense_change" {
  id             = "TODO_PLACEHOLDER"
  sessionid      = "TODO_PLACEHOLDER"
  bindtype       = "TODO_PLACEHOLDER"
  countavailable = "TODO_PLACEHOLDER"
  serialno       = "TODO_PLACEHOLDER"
}

`

func TestAccNsaptlicenseChange_basic(t *testing.T) {
	t.Skip("TODO: Requires review - changes APT license bindings; requires a real APT license file/serial and mutates licensing")
	// DANGER: consumes pooled CADS licenses; must NOT run in automation.
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// Action-only resource: no CheckDestroy (NITRO exposes no delete endpoint;
		// allocated licenses remain on the appliance).
		Steps: []resource.TestStep{
			{
				Config: testAccNsaptlicenseChange_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsaptlicenseChangeExist("citrixadc_nsaptlicense_change.tf_nsaptlicense_change", nil),
					resource.TestCheckResourceAttrSet("citrixadc_nsaptlicense_change.tf_nsaptlicense_change", "id"),
					resource.TestCheckResourceAttr("citrixadc_nsaptlicense_change.tf_nsaptlicense_change", "id", "TODO_PLACEHOLDER"),
					resource.TestCheckResourceAttr("citrixadc_nsaptlicense_change.tf_nsaptlicense_change", "bindtype", "TODO_PLACEHOLDER"),
					resource.TestCheckResourceAttr("citrixadc_nsaptlicense_change.tf_nsaptlicense_change", "countavailable", "TODO_PLACEHOLDER"),
				),
			},
		},
	})
}

// testAccCheckNsaptlicenseChangeExist is a STATE-ONLY check. The action-only
// resource carries a synthetic ID ("nsaptlicense_change") and performs a
// disruptive allocation, so we only confirm a non-empty ID in state here rather
// than issuing additional NITRO calls.
func testAccCheckNsaptlicenseChangeExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No nsaptlicense_change ID is set")
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
