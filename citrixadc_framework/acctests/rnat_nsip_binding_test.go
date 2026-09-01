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

const testAccRnat_nsip_binding_basic = `

	resource "citrixadc_rnat" "tfrnat" {
		name             = "my_rnat"
		network          = "10.2.2.0"
		netmask          = "255.255.255.255"
		useproxyport     = "ENABLED"
		srcippersistency = "DISABLED"
		connfailover     = "DISABLED"
	}
	resource "citrixadc_nsip" "tf_nsip" {
		ipaddress = "10.222.74.200"
		type      = "VIP"
		netmask   = "255.255.255.0"
		icmp      = "ENABLED"
	}
	resource "citrixadc_rnat_nsip_binding" "tf_rnat_nsip_binding" {
		name  = citrixadc_rnat.tfrnat.name
		natip = citrixadc_nsip.tf_nsip.ipaddress
	}
  
`

const testAccRnat_nsip_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion

	resource "citrixadc_rnat" "tfrnat" {
		name             = "my_rnat"
		network          = "10.2.2.0"
		netmask          = "255.255.255.255"
		useproxyport     = "ENABLED"
		srcippersistency = "DISABLED"
		connfailover     = "DISABLED"
	}
	resource "citrixadc_nsip" "tf_nsip" {
		ipaddress = "10.222.74.200"
		type      = "VIP"
		netmask   = "255.255.255.0"
		icmp      = "ENABLED"
	}
`

func TestAccRnat_nsip_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRnat_nsip_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccRnat_nsip_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRnat_nsip_bindingExist("citrixadc_rnat_nsip_binding.tf_rnat_nsip_binding", nil),
				),
			},
			{
				Config: testAccRnat_nsip_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRnat_nsip_bindingNotExist("citrixadc_rnat_nsip_binding.tf_rnat_nsip_binding", "my_rnat,10.222.74.200"),
				),
			},
		},
	})
}

func testAccCheckRnat_nsip_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No rnat_nsip_binding id is set")
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

		idMap, _, err := utils.ParseIdString(bindingId, []string{"name", "natip"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", bindingId, err)
		}
		name := idMap["name"]
		natip := idMap["natip"]

		findParams := service.FindParams{
			ResourceType:             "rnat_nsip_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the matching natip
		found := false
		for _, v := range dataArr {
			if v["natip"].(string) == natip {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("rnat_nsip_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckRnat_nsip_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}
		idSlice := strings.SplitN(id, ",", 2)

		name := idSlice[0]
		natip := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "rnat_nsip_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to hopefully not find the one with the matching natip
		found := false
		for _, v := range dataArr {
			if v["natip"].(string) == natip {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("rnat_nsip_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckRnat_nsip_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_rnat_nsip_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("rnat_nsip_binding", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("rnat_nsip_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccRnat_nsip_bindingDataSource_basic = `

	resource "citrixadc_rnat" "tfrnat" {
		name             = "my_rnat"
		network          = "10.2.2.0"
		netmask          = "255.255.255.255"
		useproxyport     = "ENABLED"
		srcippersistency = "DISABLED"
		connfailover     = "DISABLED"
	}
	resource "citrixadc_nsip" "tf_nsip" {
		ipaddress = "10.222.74.200"
		type      = "VIP"
		netmask   = "255.255.255.0"
		icmp      = "ENABLED"
	}
	resource "citrixadc_rnat_nsip_binding" "tf_rnat_nsip_binding" {
		name  = citrixadc_rnat.tfrnat.name
		natip = citrixadc_nsip.tf_nsip.ipaddress
	}

	data "citrixadc_rnat_nsip_binding" "tf_rnat_nsip_binding" {
		name  = citrixadc_rnat_nsip_binding.tf_rnat_nsip_binding.name
		natip = citrixadc_rnat_nsip_binding.tf_rnat_nsip_binding.natip
		depends_on = [citrixadc_rnat_nsip_binding.tf_rnat_nsip_binding]
	}
`

func TestAccRnat_nsip_bindingDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccRnat_nsip_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_rnat_nsip_binding.tf_rnat_nsip_binding", "name", "my_rnat"),
					resource.TestCheckResourceAttr("data.citrixadc_rnat_nsip_binding.tf_rnat_nsip_binding", "natip", "10.222.74.200"),
					// Universal runtime-binding proof; read-only GET-only fields
					// (ownergroup/td) may be omitted for a basic standalone binding.
					resource.TestCheckResourceAttrSet("data.citrixadc_rnat_nsip_binding.tf_rnat_nsip_binding", "id"),
				),
			},
		},
	})
}

const testAccRnat_nsip_binding_upgrade_basic = `

	resource "citrixadc_rnat" "tfrnat" {
		name             = "my_rnat"
		network          = "10.2.2.0"
		netmask          = "255.255.255.255"
		useproxyport     = "ENABLED"
		srcippersistency = "DISABLED"
		connfailover     = "DISABLED"
	}
	resource "citrixadc_nsip" "tf_nsip" {
		ipaddress = "10.222.74.200"
		type      = "VIP"
		netmask   = "255.255.255.0"
		icmp      = "ENABLED"
	}
	resource "citrixadc_rnat_nsip_binding" "tf_rnat_nsip_binding" {
		name  = citrixadc_rnat.tfrnat.name
		natip = citrixadc_nsip.tf_nsip.ipaddress
	}

`

// TestAccRnat_nsip_binding_sdkv2StateUpgrade verifies that state written by the
// last SDK v2 release (legacy comma-joined id) is transparently upgraded by the
// current Framework provider. Step 1 creates the binding with citrix/citrixadc
// 2.2.0 (legacy id "name,natip"); step 2 refreshes/plans the same config through
// the current Framework provider, whose Read parses the legacy id and recomputes
// it to the new "name:<v>,natip:<v>" format (SetAttrFromGet).
func TestAccRnat_nsip_binding_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckRnat_nsip_bindingDestroy,
		Steps: []resource.TestStep{
			{
				// Step 1: create with the last SDK v2 release, writing the legacy id.
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.0.0",
					},
				},
				Config: testAccRnat_nsip_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRnat_nsip_bindingExist("citrixadc_rnat_nsip_binding.tf_rnat_nsip_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_rnat_nsip_binding.tf_rnat_nsip_binding", "id", "my_rnat,10.222.74.200"),
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
				Config: testAccRnat_nsip_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRnat_nsip_bindingExist("citrixadc_rnat_nsip_binding.tf_rnat_nsip_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_rnat_nsip_binding.tf_rnat_nsip_binding", "id", "name:my_rnat,natip:10.222.74.200"),
				),
			},
		},
	})
}

func TestAccRnat_nsip_binding_import(t *testing.T) {
	const resAddr = "citrixadc_rnat_nsip_binding.tf_rnat_nsip_binding"

	// Backward-compat: import via the LEGACY SDK v2 id. Rebuild the legacy positional id from
	// the current canonical key:value id (raw values, only the keys actually set, in legacy
	// order: name,natip) so it matches exactly what SDK v2 wrote.
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
		ordr := []string{"name", "natip"}
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
		CheckDestroy:             testAccCheckRnat_nsip_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccRnat_nsip_binding_basic},
			{Config: testAccRnat_nsip_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
			{Config: testAccRnat_nsip_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateIdFunc: legacyID, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
		},
	})
}

// TestAccRnat_nsip_binding_selfHealing verifies the provider re-creates the binding when
// it is deleted out-of-band between apply steps (drift recovery). The natip delete arg
// matches the resource's own Delete.
func TestAccRnat_nsip_binding_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_rnat_nsip_binding.tf_rnat_nsip_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRnat_nsip_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccRnat_nsip_binding_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckRnat_nsip_bindingExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResourceWithArgsMap(service.Rnat_nsip_binding.Type(), "my_rnat", map[string]string{"natip": "10.222.74.200"}); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccRnat_nsip_binding_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckRnat_nsip_bindingExist(resAddr, nil)),
			},
		},
	})
}
