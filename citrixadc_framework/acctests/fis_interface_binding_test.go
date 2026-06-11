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

// IMPORTANT — fis_interface_binding has NO NITRO read path.
//
// NITRO exposes only add (PUT) and delete for fis_interface_binding; there is no
// GET endpoint, and the aggregate fis_binding/<name> endpoint does NOT surface
// bound interface members (verified live on NS14.1 — `show fis` lists the bound
// interface over the CLI, but no NITRO GET returns it). The resource Read is
// therefore a tolerant no-op and the datasource was removed.
//
// As a result, these tests CANNOT verify the binding through a NITRO read. We
// follow the no-GET / action-only precedent: the Exist check verifies the
// resource is present in Terraform state with a non-empty id (state-only), and
// CheckDestroy verifies it has been removed from state. No datasource test is
// generated because the datasource does not exist.

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// NOTE on ifnum value below:
//   fis_interface_binding binds a *physical* interface (slot/port notation, e.g.
//   "1/2") to an FIS. The interface MUST exist and be free on the appliance, and
//   valid interface numbers are TESTBED-SPECIFIC. Replace TODO_PLACEHOLDER with a
//   real, free physical interface on your ADC (e.g. "1/2"). The same value must
//   be used in the config and in the attribute checks.

// Step 1: create the FIS parent and bind a physical interface to it.
const testAccFis_interface_binding_basic_step1 = `

resource "citrixadc_fis" "tf_fis" {
	name = "tf_fis"
}

resource "citrixadc_fis_interface_binding" "tf_fis_interface_binding" {
	name  = citrixadc_fis.tf_fis.name
	ifnum = "1/2" // testbed-specific free physical interface, e.g. "1/2"

	depends_on = [citrixadc_fis.tf_fis]
}
`

// Step 2: drop the binding (keep the FIS parent) to exercise delete + CheckDestroy.
const testAccFis_interface_binding_basic_step2 = `

resource "citrixadc_fis" "tf_fis" {
	name = "tf_fis"
}
`

func TestAccFis_interface_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckFis_interface_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccFis_interface_binding_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFis_interface_bindingExist("citrixadc_fis_interface_binding.tf_fis_interface_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_fis_interface_binding.tf_fis_interface_binding", "name", "tf_fis"),
					resource.TestCheckResourceAttr("citrixadc_fis_interface_binding.tf_fis_interface_binding", "ifnum", "1/2"),
				),
			},
			{
				Config: testAccFis_interface_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					// Binding dropped from config; CheckDestroy (state-only) confirms removal.
					testAccCheckFis_interface_bindingGoneFromState("citrixadc_fis_interface_binding.tf_fis_interface_binding"),
				),
			},
		},
	})
}

// testAccCheckFis_interface_bindingExist is a STATE-ONLY check.
//
// fis_interface_binding has no NITRO GET path (members are only visible via the
// CLI `show fis`), so we cannot confirm the binding through the appliance. We
// instead verify the resource is present in Terraform state with a non-empty id,
// mirroring the action-only / no-GET resource test precedent.
func testAccCheckFis_interface_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No fis_interface_binding id is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		// No NITRO read is possible for fis_interface_binding; presence in state
		// with a non-empty id is the only verifiable signal.
		return nil
	}
}

// testAccCheckFis_interface_bindingGoneFromState confirms the binding is no
// longer in Terraform state after it is dropped from config (state-only — there
// is no NITRO read to confirm appliance-side removal).
func testAccCheckFis_interface_bindingGoneFromState(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if _, ok := s.RootModule().Resources[n]; ok {
			return fmt.Errorf("fis_interface_binding %s is still present in state, expected it to be removed", n)
		}
		return nil
	}
}

// testAccCheckFis_interface_bindingDestroy is a STATE-ONLY destroy check.
//
// There is no NITRO GET for fis_interface_binding, so appliance-side removal
// cannot be verified. We assert no fis_interface_binding resources remain in
// Terraform state, which is the verifiable contract for a no-GET binding.
func testAccCheckFis_interface_bindingDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_fis_interface_binding" {
			continue
		}
		return fmt.Errorf("fis_interface_binding %s still present in state, expected it to be destroyed", rs.Primary.ID)
	}
	return nil
}
