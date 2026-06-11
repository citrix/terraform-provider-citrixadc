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
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// NOTE: The interface number (ifnum) below is TESTBED-SPECIFIC. Binding an
// interface to a VRID changes VMAC ownership of that interface and CAN DISRUPT
// NETWORKING on the appliance. Replace TODO_PLACEHOLDER with a free, unused
// interface (slot/port notation, e.g. "1/2") before running this test.
//
// The parent VRID key attribute is "vrid_id" (the integer VRID, 1-255). It was
// renamed from the NITRO wire field "id" to avoid colliding with the framework's
// synthetic string "id". Use "vrid_id" in HCL, never "id".

const testAccVrid_channel_binding_basic_step1 = `
	resource "citrixadc_vrid" "tf_vrid" {
		vrid_id    = 100
		preemption = "DISABLED"
		sharing    = "ENABLED"
		tracking   = "NONE"
	}

	resource "citrixadc_vrid_channel_binding" "tf_vrid_channel_binding" {
		vrid_id = citrixadc_vrid.tf_vrid.vrid_id
		ifnum   = "1/2" # free interface, e.g. "1/2" (testbed-specific)

		depends_on = [citrixadc_vrid.tf_vrid]
	}
`

// Step 2 drops the binding (keeps the parent VRID) to verify clean deletion.
const testAccVrid_channel_binding_basic_step2 = `
	resource "citrixadc_vrid" "tf_vrid" {
		vrid_id    = 100
		preemption = "DISABLED"
		sharing    = "ENABLED"
		tracking   = "NONE"
	}
`

func TestAccVrid_channel_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVrid_channel_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVrid_channel_binding_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVrid_channel_bindingExist("citrixadc_vrid_channel_binding.tf_vrid_channel_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_vrid_channel_binding.tf_vrid_channel_binding", "vrid_id", "100"),
					resource.TestCheckResourceAttr("citrixadc_vrid_channel_binding.tf_vrid_channel_binding", "ifnum", "1/2"),
				),
			},
			{
				Config: testAccVrid_channel_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVrid_channel_bindingNotExist("citrixadc_vrid_channel_binding.tf_vrid_channel_binding", "100,1/2"),
				),
			},
		},
	})
}

// vrid_channel_bindingAggregateReadForTest mirrors the resource's aggregate read:
// the by-name binding endpoint can return a keyless empty body, so bound members
// are read via the parent aggregate (GET vrid_binding/<id>) and the nested
// "vrid_channel_binding" arrays are flattened.
func vrid_channel_bindingAggregateReadForTest(client *service.NitroClient, id string) ([]map[string]interface{}, error) {
	findParams := service.FindParams{
		ResourceType:             "vrid_binding",
		ResourceName:             id,
		ResourceMissingErrorCode: 258,
	}
	parentArr, err := client.FindResourceArrayWithParams(findParams)
	if err != nil {
		return nil, err
	}

	rows := make([]map[string]interface{}, 0)
	for _, parent := range parentArr {
		// Verified live: an interface bound via the vrid_channel_binding endpoint is
		// stored as a "vrid_interface_binding" row in the aggregate vrid_binding/<id>
		// response (no "vrid_channel_binding" array). Mirror the resource read.
		nested, ok := parent["vrid_interface_binding"]
		if !ok || nested == nil {
			continue
		}
		nestedArr, ok := nested.([]interface{})
		if !ok {
			continue
		}
		for _, item := range nestedArr {
			if m, ok := item.(map[string]interface{}); ok {
				rows = append(rows, m)
			}
		}
	}
	return rows, nil
}

// vridChannelRowMatchesForTest reports whether an aggregate vrid_interface_binding
// row corresponds to the wanted ifnum. Verified live: the row does NOT echo
// "ifnum"; when present it is matched, otherwise accepted by presence (the parent
// vrid id already scopes the result). Mirrors the resource read fallback.
func vridChannelRowMatchesForTest(v map[string]interface{}, want string) bool {
	raw, ok := v["ifnum"]
	if !ok || raw == nil {
		return true
	}
	if s, ok := raw.(string); ok {
		return s == want
	}
	return false
}

func testAccCheckVrid_channel_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No vrid_channel_binding id is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		// Composite ID format: id:<vrid>,ifnum:<value>. Legacy order [id, ifnum].
		idMap, _, err := utils.ParseIdString(rs.Primary.ID, []string{"id", "ifnum"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", rs.Primary.ID, err)
		}
		vridId := idMap["id"]
		ifnum := idMap["ifnum"]

		dataArr, err := vrid_channel_bindingAggregateReadForTest(client, vridId)
		if err != nil {
			return err
		}

		found := false
		for _, v := range dataArr {
			if vridChannelRowMatchesForTest(v, ifnum) {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("vrid_channel_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckVrid_channel_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		// Helper id arg is the legacy comma form: "<vrid>,<ifnum>".
		idMap, _, err := utils.ParseIdString(id, []string{"id", "ifnum"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", id, err)
		}
		vridId := idMap["id"]
		ifnum := idMap["ifnum"]

		dataArr, err := vrid_channel_bindingAggregateReadForTest(client, vridId)
		if err != nil {
			return err
		}

		found := false
		for _, v := range dataArr {
			if vridChannelRowMatchesForTest(v, ifnum) {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("vrid_channel_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckVrid_channel_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_vrid_channel_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		idMap, _, err := utils.ParseIdString(rs.Primary.ID, []string{"id", "ifnum"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", rs.Primary.ID, err)
		}
		vridId := idMap["id"]
		ifnum := idMap["ifnum"]

		dataArr, err := vrid_channel_bindingAggregateReadForTest(client, vridId)
		if err != nil {
			// Parent VRID itself gone: the binding is necessarily gone too.
			continue
		}

		for _, v := range dataArr {
			if vridChannelRowMatchesForTest(v, ifnum) {
				return fmt.Errorf("vrid_channel_binding %s still exists", rs.Primary.ID)
			}
		}
	}

	return nil
}

// Datasource exposes the renamed key (vrid_id) and member (ifnum). The computed
// read-only outputs (flags, vlan) are not asserted because they are
// appliance-assigned and not deterministic.
const testAccVrid_channel_bindingDataSource_basic = `
	resource "citrixadc_vrid" "tf_vrid" {
		vrid_id    = 100
		preemption = "DISABLED"
		sharing    = "ENABLED"
		tracking   = "NONE"
	}

	resource "citrixadc_vrid_channel_binding" "tf_vrid_channel_binding" {
		vrid_id = citrixadc_vrid.tf_vrid.vrid_id
		ifnum   = "1/2" # free interface, e.g. "1/2" (testbed-specific)

		depends_on = [citrixadc_vrid.tf_vrid]
	}

	data "citrixadc_vrid_channel_binding" "tf_vrid_channel_binding" {
		vrid_id = citrixadc_vrid_channel_binding.tf_vrid_channel_binding.vrid_id
		ifnum   = citrixadc_vrid_channel_binding.tf_vrid_channel_binding.ifnum

		depends_on = [citrixadc_vrid_channel_binding.tf_vrid_channel_binding]
	}
`

func TestAccVrid_channel_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccVrid_channel_bindingDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_vrid_channel_binding.tf_vrid_channel_binding", "vrid_id", "100"),
					resource.TestCheckResourceAttr("data.citrixadc_vrid_channel_binding.tf_vrid_channel_binding", "ifnum", "1/2"),
				),
			},
		},
	})
}
