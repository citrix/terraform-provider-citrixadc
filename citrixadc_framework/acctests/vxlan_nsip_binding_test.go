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
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const testAccVxlan_nsip_binding_basic = `
	resource "citrixadc_vxlan" "tf_vxlan" {
		vxlanid            = 123
		port               = 33
		dynamicrouting     = "DISABLED"
		ipv6dynamicrouting = "DISABLED"
		innervlantagging   = "ENABLED"
	}
	resource "citrixadc_nsip" "tf_snip" {
		ipaddress = "10.222.74.146"
		type      = "SNIP"
		netmask   = "255.255.255.0"
		icmp      = "ENABLED"
		state     = "ENABLED"
	}
	resource "citrixadc_vxlan_nsip_binding" "tf_binding" {
		vxlanid   = citrixadc_vxlan.tf_vxlan.vxlanid
		ipaddress = citrixadc_nsip.tf_snip.ipaddress
		netmask   = citrixadc_nsip.tf_snip.netmask
	}
`

const testAccVxlan_nsip_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion
	resource "citrixadc_vxlan" "tf_vxlan" {
		vxlanid            = 123
		port               = 33
		dynamicrouting     = "DISABLED"
		ipv6dynamicrouting = "DISABLED"
		innervlantagging   = "ENABLED"
	}
	resource "citrixadc_nsip" "tf_snip" {
		ipaddress = "10.222.74.146"
		type      = "SNIP"
		netmask   = "255.255.255.0"
		icmp      = "ENABLED"
		state     = "ENABLED"
	}
`

func TestAccVxlan_nsip_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVxlan_nsip_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVxlan_nsip_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVxlan_nsip_bindingExist("citrixadc_vxlan_nsip_binding.tf_binding", nil),
				),
			},
			{
				Config: testAccVxlan_nsip_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVxlan_nsip_bindingNotExist("citrixadc_vxlan_nsip_binding.tf_binding", "123,10.222.74.146"),
				),
			},
		},
	})
}

func testAccCheckVxlan_nsip_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No vxlan_nsip_binding id is set")
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

		bindingId := rs.Primary.ID

		idMap, _, err := utils.ParseIdString(bindingId, []string{"vxlanid", "ipaddress"}, nil)
		if err != nil {
			return err
		}
		vxlanid := idMap["vxlanid"]
		ipaddress := idMap["ipaddress"]

		findParams := service.FindParams{
			ResourceType:             "vxlan_nsip_binding",
			ResourceName:             vxlanid,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}
		// Resource is missing
		if len(dataArr) == 0 {
			return fmt.Errorf("Cannot find vxlan_nsip_binding %s", bindingId)
		}

		// Iterate through results to find the one with the matching secondIdComponent
		foundIndex := -1
		for i, v := range dataArr {
			if v["ipaddress"].(string) == ipaddress {
				foundIndex = i
				break
			}
		}

		if foundIndex == -1 {
			return fmt.Errorf("Resource missing vxlan_nsip_binding %s", bindingId)
		}

		return nil
	}
}

func testAccCheckVxlan_nsip_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}
		idMap, _, err := utils.ParseIdString(id, []string{"vxlanid", "ipaddress"}, nil)
		if err != nil {
			return err
		}
		vxlanid := idMap["vxlanid"]
		ipaddress := idMap["ipaddress"]

		findParams := service.FindParams{
			ResourceType:             "vxlan_nsip_binding",
			ResourceName:             vxlanid,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Resource is missing
		if len(dataArr) == 0 {
			return nil
		}

		// Iterate through results to hopefully not find the one with the matching secondIdComponent
		foundIndex := -1
		for i, v := range dataArr {
			if v["ipaddress"].(string) == ipaddress {
				foundIndex = i
				break
			}
		}

		// Resource is missing
		if foundIndex == -1 {
			return nil
		}

		return fmt.Errorf("Resource still exists vxlan_nsip_binding %s", id)
	}
}

func testAccCheckVxlan_nsip_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_vxlan_nsip_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Vxlan_nsip_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("vxlan_nsip_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccVxlan_nsip_bindingDataSource_basic = `
	resource "citrixadc_vxlan" "tf_vxlan" {
		vxlanid            = 123
		port               = 33
		dynamicrouting     = "DISABLED"
		ipv6dynamicrouting = "DISABLED"
		innervlantagging   = "ENABLED"
	}
	resource "citrixadc_nsip" "tf_snip" {
		ipaddress = "10.222.74.146"
		type      = "SNIP"
		netmask   = "255.255.255.0"
		icmp      = "ENABLED"
		state     = "ENABLED"
	}
	resource "citrixadc_vxlan_nsip_binding" "tf_binding" {
		vxlanid   = citrixadc_vxlan.tf_vxlan.vxlanid
		ipaddress = citrixadc_nsip.tf_snip.ipaddress
		netmask   = citrixadc_nsip.tf_snip.netmask
	}

	data "citrixadc_vxlan_nsip_binding" "tf_binding" {
		vxlanid   = citrixadc_vxlan_nsip_binding.tf_binding.vxlanid
		ipaddress = citrixadc_vxlan_nsip_binding.tf_binding.ipaddress
		depends_on = [citrixadc_vxlan_nsip_binding.tf_binding]
	}
`

func TestAccVxlan_nsip_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccVxlan_nsip_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_vxlan_nsip_binding.tf_binding", "vxlanid", "123"),
					resource.TestCheckResourceAttr("data.citrixadc_vxlan_nsip_binding.tf_binding", "ipaddress", "10.222.74.146"),
					resource.TestCheckResourceAttr("data.citrixadc_vxlan_nsip_binding.tf_binding", "netmask", "255.255.255.0"),
				),
			},
		},
	})
}

const testAccVxlan_nsip_binding_upgrade_basic = `
	resource "citrixadc_vxlan" "tf_vxlan" {
		vxlanid            = 123
		port               = 33
		dynamicrouting     = "DISABLED"
		ipv6dynamicrouting = "DISABLED"
		innervlantagging   = "ENABLED"
	}
	resource "citrixadc_nsip" "tf_snip" {
		ipaddress = "10.222.74.146"
		type      = "SNIP"
		netmask   = "255.255.255.0"
		icmp      = "ENABLED"
		state     = "ENABLED"
	}
	resource "citrixadc_vxlan_nsip_binding" "tf_binding" {
		vxlanid   = citrixadc_vxlan.tf_vxlan.vxlanid
		ipaddress = citrixadc_nsip.tf_snip.ipaddress
		netmask   = citrixadc_nsip.tf_snip.netmask
	}
`

// TestAccVxlan_nsip_binding_sdkv2StateUpgrade verifies that state written by the last
// SDK v2 release (legacy comma-joined id "vxlanid,ipaddress") is transparently upgraded
// by the current Framework provider. Step 1 creates the binding with citrix/citrixadc
// 2.2.0; step 2 refreshes/plans the same config through the current Framework provider,
// whose Read parses the legacy id and recomputes it to the new
// "vxlanid:<v>,ipaddress:<v>" canonical format (SetAttrFromGet).
func TestAccVxlan_nsip_binding_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckVxlan_nsip_bindingDestroy,
		Steps: []resource.TestStep{
			{
				// Step 1: create with the last SDK v2 release, writing the legacy id.
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.2.0",
					},
				},
				Config: testAccVxlan_nsip_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVxlan_nsip_bindingExist("citrixadc_vxlan_nsip_binding.tf_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_vxlan_nsip_binding.tf_binding", "id", "123,10.222.74.146"),
				),
			},
			{
				// Step 2: refresh/apply the same config through the current Framework
				// provider. Read exercises ParseIdString on the legacy id, then
				// recomputes the id to the new key:value canonical format.
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccVxlan_nsip_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVxlan_nsip_bindingExist("citrixadc_vxlan_nsip_binding.tf_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_vxlan_nsip_binding.tf_binding", "id", "vxlanid:123,ipaddress:10.222.74.146"),
				),
			},
		},
	})
}

func TestAccVxlan_nsip_binding_import(t *testing.T) {
	const resAddr = "citrixadc_vxlan_nsip_binding.tf_binding"

	// Backward-compat: import via the LEGACY SDK v2 id. Rebuild the legacy positional id from
	// the current canonical key:value id (raw values, only the keys actually set, in legacy
	// order: vxlanid,ipaddress) so it matches exactly what SDK v2 wrote.
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
		ordr := []string{"vxlanid", "ipaddress"}
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
		CheckDestroy:             testAccCheckVxlan_nsip_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccVxlan_nsip_binding_basic},
			{Config: testAccVxlan_nsip_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
			{Config: testAccVxlan_nsip_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateIdFunc: legacyID, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
		},
	})
}

func TestAccVxlan_nsip_binding_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_vxlan_nsip_binding.tf_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVxlan_nsip_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVxlan_nsip_binding_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckVxlan_nsip_bindingExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResourceWithArgsMap(service.Vxlan_nsip_binding.Type(), "123", map[string]string{"ipaddress": "10.222.74.146", "netmask": "255.255.255.0"}); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccVxlan_nsip_binding_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckVxlan_nsip_bindingExist(resAddr, nil)),
			},
		},
	})
}
