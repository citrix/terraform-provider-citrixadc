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

// NOTE on the streamsession resource:
//   - Models the NITRO POST /streamsession?action=clear endpoint, which clears the
//     stream session for a given stream identifier (CLI: "clear stream session <name>").
//   - This is an ACTION-ONLY (one-shot side-effect) resource: Create performs the
//     clear action, Read is a no-op (preserves state), Update is a no-op (the single
//     attribute `name` is RequiresReplace), and Delete is a state-only removal.
//     There is NO get/add/update/delete endpoint, so the resource CANNOT be verified
//     by reading it back from the ADC, and it has NO datasource (removed per
//     Pattern 13 — no NITRO GET endpoint).
//   - The single Required attribute `name` is the stream identifier name; it is
//     RequiresReplace. `clear stream session <name>` targets an existing
//     streamidentifier, so the test first creates a streamselector + streamidentifier
//     (config reused from streamidentifier_analyticsprofile_binding_test.go) and wires
//     name = the streamidentifier's name.
//   - The Exist check below only verifies that the resource landed in Terraform
//     state with its synthetic ID ("streamsession-config"); it does NOT (and cannot)
//     verify the clear side-effect via NITRO.
//   - There is no CheckDestroy: the clear action has no inverse on NITRO, and there
//     is no GET-by-id to confirm absence; Delete is a state-only removal.
//
// This mirrors the action-only test precedent in rnatsession_test.go /
// dnsproxyrecords_test.go (single apply step, state-only Exist check, no
// CheckDestroy, TestCheckResourceAttrSet on "id"), adapted for the clear action.
//
// ENVIRONMENT: clear works on a standalone testbed; no ADC_TESTBED gate is needed.

// Single apply step: `name` is RequiresReplace, so there is no in-place update to
// exercise. The streamselector -> streamidentifier prerequisites are created first,
// and the streamsession clears that identifier's session by name.
const testAccStreamsession_basic = `
resource "citrixadc_streamselector" "tf_streamselector" {
  name = "my_streamselector"
  rule = ["HTTP.REQ.URL", "CLIENT.IP.SRC"]
}

resource "citrixadc_streamidentifier" "tf_streamidentifier" {
  name         = "my_streamidentifier"
  selectorname = citrixadc_streamselector.tf_streamselector.name
  samplecount  = 10
  sort         = "CONNECTIONS"
  snmptrap     = "ENABLED"
  loglimit     = 500
  loginterval  = 60
  log          = "NONE"
}

resource "citrixadc_streamsession" "tf_streamsession" {
  name = citrixadc_streamidentifier.tf_streamidentifier.name

  depends_on = [citrixadc_streamidentifier.tf_streamidentifier]
}

`

func TestAccStreamsession_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// No CheckDestroy: the clear action has no inverse on NITRO and there is no
		// GET-by-id to confirm absence; Delete is a state-only removal.
		Steps: []resource.TestStep{
			{
				Config: testAccStreamsession_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckStreamsessionExist("citrixadc_streamsession.tf_streamsession", nil),
					resource.TestCheckResourceAttr("citrixadc_streamsession.tf_streamsession", "name", "my_streamidentifier"),
					// "id" is the synthetic state handle "streamsession-config".
					resource.TestCheckResourceAttrSet("citrixadc_streamsession.tf_streamsession", "id"),
				),
			},
		},
	})
}

// testAccCheckStreamsessionExist is a state-only existence check.
//
// streamsession is an action-only resource: Read is a no-op and there is no
// GET-by-id endpoint, so we CANNOT verify the clear via NITRO. We only assert that
// Terraform recorded the resource in state with a non-empty ID (which equals the
// synthetic "streamsession-config" after a successful POST ?action=clear). This
// mirrors testAccCheckRnatsessionExist.
func testAccCheckStreamsessionExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No streamsession ID is set")
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
