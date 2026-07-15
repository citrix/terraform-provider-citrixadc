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

// NOTE: ifnum is testbed-specific and binding a channel interface to a VMAC6
// can be DISRUPTIVE. Replace TODO_PLACEHOLDER with a free interface (e.g. "1/1")
// that is safe to use on the target ADC before running this test.

// step1: create the vrid6 parent, then bind a channel interface to it.
const testAccVrid6_channel_binding_basic_step1 = `
resource "citrixadc_vrid6" "tf_vrid6" {
	vrid6_id   = 100
	preemption = "DISABLED"
	sharing    = "DISABLED"
	tracking   = "NONE"
}

resource "citrixadc_vrid6_channel_binding" "tf_vrid6_channel_binding" {
	vrid_id = citrixadc_vrid6.tf_vrid6.vrid6_id
	ifnum   = "1/1" // free interface, e.g. "1/1" (testbed-specific, disruptive)

	depends_on = [citrixadc_vrid6.tf_vrid6]
}
`

// step2: drop the binding (keep the parent) to exercise delete of the binding.
const testAccVrid6_channel_binding_basic_step2 = `
resource "citrixadc_vrid6" "tf_vrid6" {
	vrid6_id   = 100
	preemption = "DISABLED"
	sharing    = "DISABLED"
	tracking   = "NONE"
}
`

func TestAccVrid6_channel_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVrid6_channel_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVrid6_channel_binding_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVrid6_channel_bindingExist("citrixadc_vrid6_channel_binding.tf_vrid6_channel_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_vrid6_channel_binding.tf_vrid6_channel_binding", "vrid_id", "100"),
					resource.TestCheckResourceAttr("citrixadc_vrid6_channel_binding.tf_vrid6_channel_binding", "ifnum", "1/1"),
				),
			},
			{
				Config: testAccVrid6_channel_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVrid6_channel_bindingNotExist("citrixadc_vrid6_channel_binding.tf_vrid6_channel_binding", "id:100,ifnum:1%2F1"),
				),
			},
		},
	})
}

// vrid6ChannelRowMatchesForTest reports whether an aggregate vrid6_interface_binding
// row (where a physical interface bound via the channel endpoint lands) corresponds
// to the wanted ifnum. Verified live: the row does NOT echo "ifnum"; when present it
// is matched, otherwise accepted by row presence (the parent vrid6 id already scopes
// the result). Mirrors the resource read fallback.
func vrid6ChannelRowMatchesForTest(m map[string]interface{}, want string) bool {
	raw, ok := m["ifnum"]
	if !ok || raw == nil {
		return true
	}
	if s, ok := raw.(string); ok {
		return s == want
	}
	return false
}

// testAccCheckVrid6_channel_bindingExist mirrors the resource's aggregate-read
// helper: it queries the parent aggregate endpoint (vrid6_binding/<id>), flattens
// the nested "vrid6_interface_binding" arrays (where a physical interface bound via
// the channel endpoint actually lands), and matches the row by ifnum.
func testAccCheckVrid6_channel_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No vrid6_channel_binding id is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		idMap, _, err := utils.ParseIdString(rs.Primary.ID, []string{"id", "ifnum"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID: %v", err)
		}
		idValue := idMap["id"]
		ifnum := idMap["ifnum"]

		// Aggregate read of the parent endpoint and flatten the nested binding rows.
		findParams := service.FindParams{
			ResourceType:             "vrid6_binding",
			ResourceName:             idValue,
			ResourceMissingErrorCode: 258,
		}
		parentArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			return err
		}

		found := false
		for _, parent := range parentArr {
			nested, ok := parent["vrid6_interface_binding"] // verified live: physical interface lands as vrid6_interface_binding
			if !ok || nested == nil {
				continue
			}
			nestedArr, ok := nested.([]interface{})
			if !ok {
				continue
			}
			for _, item := range nestedArr {
				m, ok := item.(map[string]interface{})
				if !ok {
					continue
				}
				if vrid6ChannelRowMatchesForTest(m, ifnum) {
					found = true
					break
				}
			}
			if found {
				break
			}
		}

		if !found {
			return fmt.Errorf("vrid6_channel_binding %s not found", n)
		}

		return nil
	}
}

// testAccCheckVrid6_channel_bindingNotExist verifies the binding row is no longer
// present in the parent aggregate response.
func testAccCheckVrid6_channel_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		idMap, _, err := utils.ParseIdString(id, []string{"id", "ifnum"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID: %v", err)
		}
		idValue := idMap["id"]
		ifnum := idMap["ifnum"]

		findParams := service.FindParams{
			ResourceType:             "vrid6_binding",
			ResourceName:             idValue,
			ResourceMissingErrorCode: 258,
		}
		parentArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			// Parent gone or no bindings -> binding is destroyed.
			return nil
		}

		found := false
		for _, parent := range parentArr {
			nested, ok := parent["vrid6_interface_binding"] // verified live: physical interface lands as vrid6_interface_binding
			if !ok || nested == nil {
				continue
			}
			nestedArr, ok := nested.([]interface{})
			if !ok {
				continue
			}
			for _, item := range nestedArr {
				m, ok := item.(map[string]interface{})
				if !ok {
					continue
				}
				if vrid6ChannelRowMatchesForTest(m, ifnum) {
					found = true
					break
				}
			}
			if found {
				break
			}
		}

		if found {
			return fmt.Errorf("vrid6_channel_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckVrid6_channel_bindingDestroy(s *terraform.State) error {
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_vrid6_channel_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		idMap, _, err := utils.ParseIdString(rs.Primary.ID, []string{"id", "ifnum"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID: %v", err)
		}
		idValue := idMap["id"]
		ifnum := idMap["ifnum"]

		findParams := service.FindParams{
			ResourceType:             "vrid6_binding",
			ResourceName:             idValue,
			ResourceMissingErrorCode: 258,
		}
		parentArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			// Parent gone -> binding is destroyed.
			continue
		}

		for _, parent := range parentArr {
			nested, ok := parent["vrid6_interface_binding"] // verified live: physical interface lands as vrid6_interface_binding
			if !ok || nested == nil {
				continue
			}
			nestedArr, ok := nested.([]interface{})
			if !ok {
				continue
			}
			for _, item := range nestedArr {
				m, ok := item.(map[string]interface{})
				if !ok {
					continue
				}
				if vrid6ChannelRowMatchesForTest(m, ifnum) {
					return fmt.Errorf("vrid6_channel_binding %s still exists", rs.Primary.ID)
				}
			}
		}
	}

	return nil
}

// Datasource read-only outputs (flags, vlan) are not asserted because their
// values are appliance-determined; only the renamed key (vrid_id) and the member
// (ifnum) are asserted.
const testAccVrid6_channel_bindingDataSource_basic = `
resource "citrixadc_vrid6" "tf_vrid6" {
	vrid6_id   = 100
	preemption = "DISABLED"
	sharing    = "DISABLED"
	tracking   = "NONE"
}

resource "citrixadc_vrid6_channel_binding" "tf_vrid6_channel_binding" {
	vrid_id = citrixadc_vrid6.tf_vrid6.vrid6_id
	ifnum   = "1/1" // free interface, e.g. "1/1" (testbed-specific, disruptive)

	depends_on = [citrixadc_vrid6.tf_vrid6]
}

data "citrixadc_vrid6_channel_binding" "tf_vrid6_channel_binding" {
	vrid_id = citrixadc_vrid6_channel_binding.tf_vrid6_channel_binding.vrid_id
	ifnum   = citrixadc_vrid6_channel_binding.tf_vrid6_channel_binding.ifnum

	depends_on = [citrixadc_vrid6_channel_binding.tf_vrid6_channel_binding]
}
`

func TestAccVrid6_channel_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVrid6_channel_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVrid6_channel_bindingDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_vrid6_channel_binding.tf_vrid6_channel_binding", "vrid_id", "100"),
					resource.TestCheckResourceAttr("data.citrixadc_vrid6_channel_binding.tf_vrid6_channel_binding", "ifnum", "1/1"),
				),
			},
		},
	})
}

func TestAccVrid6_channel_binding_import(t *testing.T) {
	const resAddr = "citrixadc_vrid6_channel_binding.tf_vrid6_channel_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVrid6_channel_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVrid6_channel_binding_basic_step1,
			},
			{
				Config:            testAccVrid6_channel_binding_basic_step1,
				ResourceName:      resAddr,
				ImportState:       true,
				ImportStateVerify: true,
				// vrid_id and ifnum are both composite-ID components and are
				// reconstructed from the parsed ID in readVrid6ChannelBindingFromApi,
				// so all attributes round-trip on import.
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}
