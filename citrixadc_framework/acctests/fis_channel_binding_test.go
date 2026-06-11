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

// NOTE on ifnum value below:
//   fis_channel_binding is intended to bind a *channel* interface (e.g. an
//   LA/link-aggregate channel) to an FIS. Verified live on NS VPX, however, the
//   endpoint also accepts a plain physical interface, and -- importantly -- the
//   binding has NO NITRO read path either way (no GET surfaces it; see the check
//   helpers below). Because the user provided a free physical interface ("1/2") and
//   not a pre-existing channel, this test binds "1/2" (the write succeeds) and uses
//   the no-GET / action-only (state-only) verification precedent. If a real channel
//   is available on the target appliance, substitute its id (e.g. "LA/1"); the test
//   structure is unchanged because there is no readback regardless.

// Step 1: create the FIS parent and bind an interface to it.
const testAccFis_channel_binding_basic_step1 = `

resource "citrixadc_fis" "tf_fis" {
	name = "tf_fis"
}

resource "citrixadc_fis_channel_binding" "tf_fis_channel_binding" {
	name  = citrixadc_fis.tf_fis.name
	ifnum = "1/2" // free physical interface; a real channel id (e.g. "LA/1") also works

	depends_on = [citrixadc_fis.tf_fis]
}
`

// Step 2: drop the binding (keep the FIS parent) to exercise delete + CheckDestroy.
const testAccFis_channel_binding_basic_step2 = `

resource "citrixadc_fis" "tf_fis" {
	name = "tf_fis"
}
`

func TestAccFis_channel_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckFis_channel_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccFis_channel_binding_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFis_channel_bindingExist("citrixadc_fis_channel_binding.tf_fis_channel_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_fis_channel_binding.tf_fis_channel_binding", "name", "tf_fis"),
					resource.TestCheckResourceAttr("citrixadc_fis_channel_binding.tf_fis_channel_binding", "ifnum", "1/2"),
				),
			},
			{
				Config: testAccFis_channel_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					// Binding dropped from config. No NITRO read exists, so verify removal
					// from Terraform state (state-only, no-GET precedent).
					testAccCheckFis_channel_bindingGoneFromState("citrixadc_fis_channel_binding.tf_fis_channel_binding"),
				),
			},
		},
	})
}

// testAccCheckFis_channel_bindingExist is a STATE-ONLY check.
//
// IMPORTANT (verified live on NS VPX): fis_channel_binding has no NITRO read path.
// The bind PUT succeeds -- for a channel ifnum (e.g. "LA/1") AND for a plain physical
// interface (e.g. "1/2", which is what this test uses) -- but the binding is not
// surfaced by any GET: the aggregate fis_binding/<name> response carries only
// {"name"} with no fis_channel_binding array, and the direct endpoint returns a
// keyless empty body. We therefore follow the no-GET / action-only precedent
// (fis_interface_binding): presence in Terraform state with a non-empty id is the
// only verifiable signal.
func testAccCheckFis_channel_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No fis_channel_binding id is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		// No NITRO read is possible for fis_channel_binding; presence in state with a
		// non-empty id is the only verifiable signal.
		return nil
	}
}

// testAccCheckFis_channel_bindingGoneFromState confirms the binding is no longer in
// Terraform state after it is dropped from config (state-only -- there is no NITRO
// read to confirm appliance-side removal).
func testAccCheckFis_channel_bindingGoneFromState(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if _, ok := s.RootModule().Resources[n]; ok {
			return fmt.Errorf("fis_channel_binding %s is still present in state, expected it to be removed", n)
		}
		return nil
	}
}

// testAccCheckFis_channel_bindingDestroy is a STATE-ONLY destroy check. There is no
// NITRO GET for fis_channel_binding, so appliance-side removal cannot be verified;
// assert no such resources remain in Terraform state.
func testAccCheckFis_channel_bindingDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_fis_channel_binding" {
			continue
		}
		return fmt.Errorf("fis_channel_binding %s still present in state, expected it to be destroyed", rs.Primary.ID)
	}
	return nil
}
