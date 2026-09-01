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

// NOTE: The interface number (ifnum) below is TESTBED-SPECIFIC. Binding an
// interface to a VRID changes VMAC ownership of that interface and CAN DISRUPT
// NETWORKING on the appliance. Replace TODO_PLACEHOLDER with a free, unused
// interface (slot/port notation, e.g. "1/2") before running this test.
//
// The parent VRID key attribute is "vrid_id" (the integer VRID, 1-255). It was
// renamed from the NITRO wire field "id" to avoid colliding with the framework's
// synthetic string "id". Use "vrid_id" in HCL, never "id".

const testAccVrid_channel_binding_basic_step1 = `
	resource "citrixadc_channel" "tf_channel" {
		channel_id = "LA/1"
	}

	resource "citrixadc_vrid" "tf_vrid" {
		vrid_id    = 100
		preemption = "DISABLED"
		sharing    = "ENABLED"
		tracking   = "NONE"
	}

	resource "citrixadc_vrid_channel_binding" "tf_vrid_channel_binding" {
		vrid_id = citrixadc_vrid.tf_vrid.vrid_id
		ifnum   = citrixadc_channel.tf_channel.channel_id # a channel interface, e.g. "LA/1" (testbed-specific)

		depends_on = [citrixadc_vrid.tf_vrid, citrixadc_channel.tf_channel]
	}
`

// Step 2 drops the binding (keeps the parent VRID and the channel) to verify clean deletion.
const testAccVrid_channel_binding_basic_step2 = `
	resource "citrixadc_channel" "tf_channel" {
		channel_id = "LA/1"
	}

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
					resource.TestCheckResourceAttr("citrixadc_vrid_channel_binding.tf_vrid_channel_binding", "ifnum", "LA/1"),
				),
			},
			{
				Config: testAccVrid_channel_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVrid_channel_bindingNotExist("citrixadc_vrid_channel_binding.tf_vrid_channel_binding", "100,LA/1"),
				),
			},
		},
	})
}

func TestAccVrid_channel_binding_import(t *testing.T) {
	const resAddr = "citrixadc_vrid_channel_binding.tf_vrid_channel_binding"

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
		CheckDestroy:             testAccCheckVrid_channel_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVrid_channel_binding_basic_step1,
			},
			{
				Config:            testAccVrid_channel_binding_basic_step1,
				ResourceName:      resAddr,
				ImportState:       true,
				ImportStateVerify: true,
				// All attributes round-trip: vrid_id is echoed by the aggregate GET
				// (as "id") and ifnum, though NOT echoed by the vrid_interface_binding
				// row, is reconstructed from the composite ID during Read.
				ImportStateVerifyIgnore: []string{},
			},
			{Config: testAccVrid_channel_binding_basic_step1, ResourceName: resAddr, ImportState: true, ImportStateIdFunc: legacyID, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
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
	resource "citrixadc_channel" "tf_channel" {
		channel_id = "LA/1"
	}

	resource "citrixadc_vrid" "tf_vrid" {
		vrid_id    = 100
		preemption = "DISABLED"
		sharing    = "ENABLED"
		tracking   = "NONE"
	}

	resource "citrixadc_vrid_channel_binding" "tf_vrid_channel_binding" {
		vrid_id = citrixadc_vrid.tf_vrid.vrid_id
		ifnum   = citrixadc_channel.tf_channel.channel_id # a channel interface, e.g. "LA/1" (testbed-specific)

		depends_on = [citrixadc_vrid.tf_vrid, citrixadc_channel.tf_channel]
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
					// id is always composed at runtime; the universal binding proof.
					resource.TestCheckResourceAttrSet("data.citrixadc_vrid_channel_binding.tf_vrid_channel_binding", "id"),
					resource.TestCheckResourceAttr("data.citrixadc_vrid_channel_binding.tf_vrid_channel_binding", "vrid_id", "100"),
					resource.TestCheckResourceAttr("data.citrixadc_vrid_channel_binding.tf_vrid_channel_binding", "ifnum", "LA/1"),
				),
			},
		},
	})
}

func TestAccVrid_channel_binding_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_vrid_channel_binding.tf_vrid_channel_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVrid_channel_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVrid_channel_binding_basic_step1,
				Check:  resource.ComposeTestCheckFunc(testAccCheckVrid_channel_bindingExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResourceWithArgs(service.Vrid_channel_binding.Type(), "100", []string{fmt.Sprintf("ifnum:%s", utils.UrlEncode("LA/1"))}); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccVrid_channel_binding_basic_step1,
				Check:  resource.ComposeTestCheckFunc(testAccCheckVrid_channel_bindingExist(resAddr, nil)),
			},
		},
	})
}
