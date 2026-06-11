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

// NOTE: The tracked interface number (trackifnum) below is TESTBED-SPECIFIC.
// Binding a tracked interface to a VRID affects VRRP priority tracking and CAN
// DISRUPT NETWORKING on the appliance. Replace TODO_PLACEHOLDER with a free,
// unused interface (slot/port notation, e.g. "1/2") before running this test.
//
// The parent VRID key attribute is "vrid_id" (the integer VRID, 1-255). It was
// renamed from the NITRO wire field "id" to avoid colliding with the framework's
// synthetic string "id". Use "vrid_id" in HCL, never "id".

const testAccVrid_trackinterface_binding_basic_step1 = `
	resource "citrixadc_vrid" "tf_vrid" {
		vrid_id    = 100
		priority   = 30
		preemption = "DISABLED"
		sharing    = "ENABLED"
		tracking   = "NONE"
	}

	resource "citrixadc_vrid_trackinterface_binding" "tf_vrid_trackinterface_binding" {
		vrid_id    = citrixadc_vrid.tf_vrid.vrid_id
		trackifnum = "1/2" # free interface, e.g. "1/2" (testbed-specific)

		depends_on = [citrixadc_vrid.tf_vrid]
	}
`

// Step 2 drops the binding (keeps the parent VRID) to verify clean deletion.
const testAccVrid_trackinterface_binding_basic_step2 = `
	resource "citrixadc_vrid" "tf_vrid" {
		vrid_id    = 100
		priority   = 30
		preemption = "DISABLED"
		sharing    = "ENABLED"
		tracking   = "NONE"
	}
`

func TestAccVrid_trackinterface_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVrid_trackinterface_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVrid_trackinterface_binding_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVrid_trackinterface_bindingExist("citrixadc_vrid_trackinterface_binding.tf_vrid_trackinterface_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_vrid_trackinterface_binding.tf_vrid_trackinterface_binding", "vrid_id", "100"),
					resource.TestCheckResourceAttr("citrixadc_vrid_trackinterface_binding.tf_vrid_trackinterface_binding", "trackifnum", "1/2"),
				),
			},
			{
				Config: testAccVrid_trackinterface_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVrid_trackinterface_bindingNotExist("citrixadc_vrid_trackinterface_binding.tf_vrid_trackinterface_binding", "100,1/2"),
				),
			},
		},
	})
}

// IMPORTANT - vrid_trackinterface_binding has NO NITRO read path (verified live on
// NS VPX). The bind PUT succeeds, but the binding is not surfaced by any GET: the
// aggregate vrid_binding/<id> response carries only {"id"} with no
// vrid_trackinterface_binding array, and the direct endpoint returns a keyless empty
// body. These tests therefore follow the no-GET / action-only precedent
// (fis_interface_binding): the Exist check verifies the resource is present in
// Terraform state with a non-empty id (state-only), step 2 verifies it is gone from
// state, and CheckDestroy verifies no such resources remain in state. No datasource
// test is generated because the datasource cannot read the binding on this firmware.

func testAccCheckVrid_trackinterface_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No vrid_trackinterface_binding id is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		// No NITRO read is possible for vrid_trackinterface_binding; presence in state
		// with a non-empty id is the only verifiable signal.
		return nil
	}
}

// testAccCheckVrid_trackinterface_bindingNotExist confirms the binding is no longer
// in Terraform state after it is dropped from config (state-only -- there is no
// NITRO read to confirm appliance-side removal).
func testAccCheckVrid_trackinterface_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if _, ok := s.RootModule().Resources[n]; ok {
			return fmt.Errorf("vrid_trackinterface_binding %s is still present in state, expected it to be removed", n)
		}
		return nil
	}
}

// testAccCheckVrid_trackinterface_bindingDestroy is a STATE-ONLY destroy check.
// There is no NITRO GET for vrid_trackinterface_binding, so appliance-side removal
// cannot be verified; assert no such resources remain in Terraform state.
func testAccCheckVrid_trackinterface_bindingDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_vrid_trackinterface_binding" {
			continue
		}
		return fmt.Errorf("vrid_trackinterface_binding %s still present in state, expected it to be destroyed", rs.Primary.ID)
	}
	return nil
}
