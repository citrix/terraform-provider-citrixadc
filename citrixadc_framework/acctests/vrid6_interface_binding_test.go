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
	"net/url"
	"strings"
	"testing"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

// NOTE: ifnum is testbed-specific and binding an interface to a VMAC6 can be
// DISRUPTIVE. Replace TODO_PLACEHOLDER with a free interface (e.g. "1/2") that is
// safe to use on the target ADC before running this test.

// step1: create the vrid6 parent, then bind an interface to it.
const testAccVrid6_interface_binding_basic_step1 = `
resource "citrixadc_vrid6" "tf_vrid6" {
	vrid6_id   = 100
	preemption = "DISABLED"
	sharing    = "DISABLED"
	tracking   = "NONE"
}

resource "citrixadc_vrid6_interface_binding" "tf_vrid6_interface_binding" {
	vrid_id = citrixadc_vrid6.tf_vrid6.vrid6_id
	ifnum   = "1/1" // free interface, e.g. "1/1" (testbed-specific, disruptive)

	depends_on = [citrixadc_vrid6.tf_vrid6]
}
`

// step2: drop the binding (keep the parent) to exercise delete of the binding.
const testAccVrid6_interface_binding_basic_step2 = `
resource "citrixadc_vrid6" "tf_vrid6" {
	vrid6_id   = 100
	preemption = "DISABLED"
	sharing    = "DISABLED"
	tracking   = "NONE"
}
`

func TestAccVrid6_interface_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVrid6_interface_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVrid6_interface_binding_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVrid6_interface_bindingExist("citrixadc_vrid6_interface_binding.tf_vrid6_interface_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_vrid6_interface_binding.tf_vrid6_interface_binding", "vrid_id", "100"),
					resource.TestCheckResourceAttr("citrixadc_vrid6_interface_binding.tf_vrid6_interface_binding", "ifnum", "1/1"),
				),
			},
			{
				Config: testAccVrid6_interface_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVrid6_interface_bindingNotExist("citrixadc_vrid6_interface_binding.tf_vrid6_interface_binding", "id:100,ifnum:1%2F1"),
				),
			},
		},
	})
}

// vrid6InterfaceRowMatchesForTest reports whether an aggregate
// vrid6_interface_binding row corresponds to the wanted ifnum. Verified live on NS
// VPX: the row is {"id","vlan","flags"} and does NOT echo "ifnum". When present it
// is matched, otherwise accepted by row presence (the parent vrid6 id already
// scopes the result). Mirrors the resource read fallback.
func vrid6InterfaceRowMatchesForTest(m map[string]interface{}, want string) bool {
	raw, ok := m["ifnum"]
	if !ok || raw == nil {
		return true
	}
	if s, ok := raw.(string); ok {
		return s == want
	}
	return false
}

// testAccCheckVrid6_interface_bindingExist mirrors the resource's aggregate-read
// helper: it queries the parent aggregate endpoint (vrid6_binding/<id>), flattens
// the nested "vrid6_interface_binding" arrays, and matches the row by ifnum.
func testAccCheckVrid6_interface_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No vrid6_interface_binding id is set")
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
			nested, ok := parent["vrid6_interface_binding"]
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
				if vrid6InterfaceRowMatchesForTest(m, ifnum) {
					found = true
					break
				}
			}
			if found {
				break
			}
		}

		if !found {
			return fmt.Errorf("vrid6_interface_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckVrid6_interface_bindingNotExist(n string, id string) resource.TestCheckFunc {
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
			return nil
		}

		found := false
		for _, parent := range parentArr {
			nested, ok := parent["vrid6_interface_binding"]
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
				if vrid6InterfaceRowMatchesForTest(m, ifnum) {
					found = true
					break
				}
			}
			if found {
				break
			}
		}

		if found {
			return fmt.Errorf("vrid6_interface_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckVrid6_interface_bindingDestroy(s *terraform.State) error {
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_vrid6_interface_binding" {
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
			continue
		}

		for _, parent := range parentArr {
			nested, ok := parent["vrid6_interface_binding"]
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
				if vrid6InterfaceRowMatchesForTest(m, ifnum) {
					return fmt.Errorf("vrid6_interface_binding %s still exists", rs.Primary.ID)
				}
			}
		}
	}

	return nil
}

// Datasource read-only outputs (flags, vlan) are not asserted because their
// values are appliance-determined; only the renamed key (vrid_id) and the member
// (ifnum) are asserted.
const testAccVrid6_interface_bindingDataSource_basic = `
resource "citrixadc_vrid6" "tf_vrid6" {
	vrid6_id   = 100
	preemption = "DISABLED"
	sharing    = "DISABLED"
	tracking   = "NONE"
}

resource "citrixadc_vrid6_interface_binding" "tf_vrid6_interface_binding" {
	vrid_id = citrixadc_vrid6.tf_vrid6.vrid6_id
	ifnum   = "1/1" // free interface, e.g. "1/1" (testbed-specific, disruptive)

	depends_on = [citrixadc_vrid6.tf_vrid6]
}

data "citrixadc_vrid6_interface_binding" "tf_vrid6_interface_binding" {
	vrid_id = citrixadc_vrid6_interface_binding.tf_vrid6_interface_binding.vrid_id
	ifnum   = citrixadc_vrid6_interface_binding.tf_vrid6_interface_binding.ifnum

	depends_on = [citrixadc_vrid6_interface_binding.tf_vrid6_interface_binding]
}
`

func TestAccVrid6_interface_binding_import(t *testing.T) {
	const resAddr = "citrixadc_vrid6_interface_binding.tf_vrid6_interface_binding"

	// Backward-compat: import via the LEGACY SDK v2 id. Rebuild the legacy positional id from
	// the current canonical key:value id (raw values, only the keys actually set, in legacy
	// order: id,ifnum) so it matches exactly what SDK v2 wrote.
	legacyID := func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resAddr]
		if !ok {
			return "", fmt.Errorf("resource not found in state: %s", resAddr)
		}
		kv := map[string]string{}
		for _, p := range strings.Split(rs.Primary.ID, ",") {
			if i := strings.Index(p, ":"); i >= 0 {
				v, _ := url.QueryUnescape(p[i+1:])
				kv[p[:i]] = v
			}
		}
		ordr := []string{"id", "ifnum"}
		parts := make([]string, 0, len(ordr))
		for _, k := range ordr {
			if v, ok := kv[k]; ok {
				parts = append(parts, v)
			}
		}
		// Fallback: a positional (non key:value) id has no key:value parts to reorder; import it as-is.
		if len(parts) == 0 {
			return rs.Primary.ID, nil
		}
		return strings.Join(parts, ","), nil
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVrid6_interface_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVrid6_interface_binding_basic_step1,
			},
			{
				Config:            testAccVrid6_interface_binding_basic_step1,
				ResourceName:      resAddr,
				ImportState:       true,
				ImportStateVerify: true,
				// Full round-trip: vrid_id and ifnum are backfilled from the composite
				// ID during Read (the firmware does not echo "ifnum" in the aggregate
				// read, but it is always recoverable from the ID), so nothing needs to
				// be ignored on import.
				ImportStateVerifyIgnore: []string{},
			},
			{Config: testAccVrid6_interface_binding_basic_step1, ResourceName: resAddr, ImportState: true, ImportStateIdFunc: legacyID, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
		},
	})
}

func TestAccVrid6_interface_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVrid6_interface_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVrid6_interface_bindingDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_vrid6_interface_binding.tf_vrid6_interface_binding", "vrid_id", "100"),
					resource.TestCheckResourceAttr("data.citrixadc_vrid6_interface_binding.tf_vrid6_interface_binding", "ifnum", "1/1"),
					// Universal runtime-binding proof that the data source resolved.
					resource.TestCheckResourceAttrSet("data.citrixadc_vrid6_interface_binding.tf_vrid6_interface_binding", "id"),
				),
			},
		},
	})
}

func TestAccVrid6_interface_binding_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_vrid6_interface_binding.tf_vrid6_interface_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVrid6_interface_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVrid6_interface_binding_basic_step1,
				Check:  resource.ComposeTestCheckFunc(testAccCheckVrid6_interface_bindingExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResourceWithArgs(service.Vrid6_interface_binding.Type(), "100", []string{fmt.Sprintf("ifnum:%s", utils.UrlEncode("1/1"))}); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccVrid6_interface_binding_basic_step1,
				Check:  resource.ComposeTestCheckFunc(testAccCheckVrid6_interface_bindingExist(resAddr, nil)),
			},
		},
	})
}
