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

const testAccMapbmr_bmrv4network_binding_basic = `
	resource "citrixadc_mapbmr" "tf_mapbmr" {
		name           = "tf_mapbmr"
		ruleipv6prefix = "2001:db8:abcd:12::/64"
		psidoffset     = 6
		eabitlength    = 16
		psidlength     = 8
	}
	resource "citrixadc_mapbmr_bmrv4network_binding" "tf_binding" {
		name    = citrixadc_mapbmr.tf_mapbmr.name
		network = "1.2.3.0"
		netmask = "255.255.255.0"
	}
`

const testAccMapbmr_bmrv4network_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion
	resource "citrixadc_mapbmr" "tf_mapbmr" {
		name           = "tf_mapbmr"
		ruleipv6prefix = "2001:db8:abcd:12::/64"
		psidoffset     = 6
		eabitlength    = 16
		psidlength     = 8
	}
`

func TestAccMapbmr_bmrv4network_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckMapbmr_bmrv4network_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccMapbmr_bmrv4network_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMapbmr_bmrv4network_bindingExist("citrixadc_mapbmr_bmrv4network_binding.tf_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_mapbmr_bmrv4network_binding.tf_binding", "name", "tf_mapbmr"),
					resource.TestCheckResourceAttr("citrixadc_mapbmr_bmrv4network_binding.tf_binding", "network", "1.2.3.0"),
					resource.TestCheckResourceAttr("citrixadc_mapbmr_bmrv4network_binding.tf_binding", "netmask", "255.255.255.0"),
				),
			},
			{
				Config: testAccMapbmr_bmrv4network_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMapbmr_bmrv4network_bindingNotExist("citrixadc_mapbmr_bmrv4network_binding.tf_binding", "tf_mapbmr,1.2.3.0"),
				),
			},
		},
	})
}

func TestAccMapbmr_bmrv4network_binding_import(t *testing.T) {
	const resAddr = "citrixadc_mapbmr_bmrv4network_binding.tf_binding"

	// Backward-compat: import via the LEGACY SDK v2 id. Rebuild the legacy positional id from
	// the current canonical key:value id (raw values, only the keys actually set, in legacy
	// order: name,network) so it matches exactly what SDK v2 wrote.
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
		ordr := []string{"name", "network"}
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
		CheckDestroy:             testAccCheckMapbmr_bmrv4network_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccMapbmr_bmrv4network_binding_basic},
			{Config: testAccMapbmr_bmrv4network_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
			{Config: testAccMapbmr_bmrv4network_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateIdFunc: legacyID, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
		},
	})
}

func testAccCheckMapbmr_bmrv4network_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No mapbmr_bmrv4network_binding id is set")
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

		idMap, _, err := utils.ParseIdString(bindingId, []string{"name", "network"}, nil)
		if err != nil {
			return err
		}
		name := idMap["name"]
		networkid := idMap["network"]

		findParams := service.FindParams{
			ResourceType:             "mapbmr_bmrv4network_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the matching secondIdComponent
		found := false
		for _, v := range dataArr {
			if v["network"].(string) == networkid {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("mapbmr_bmrv4network_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckMapbmr_bmrv4network_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}
		idMap, _, err := utils.ParseIdString(id, []string{"name", "network"}, nil)
		if err != nil {
			return err
		}
		name := idMap["name"]
		networkid := idMap["network"]

		findParams := service.FindParams{
			ResourceType:             "mapbmr_bmrv4network_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to hopefully not find the one with the matching secondIdComponent
		found := false
		for _, v := range dataArr {
			if v["network"].(string) == networkid {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("mapbmr_bmrv4network_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckMapbmr_bmrv4network_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_mapbmr_bmrv4network_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("mapbmr_bmrv4network_binding", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("mapbmr_bmrv4network_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccMapbmr_bmrv4network_bindingDataSource_basic = `
	resource "citrixadc_mapbmr" "tf_mapbmr" {
		name           = "tf_mapbmr"
		ruleipv6prefix = "2001:db8:abcd:12::/64"
		psidoffset     = 6
		eabitlength    = 16
		psidlength     = 8
	}
	resource "citrixadc_mapbmr_bmrv4network_binding" "tf_binding" {
		name    = citrixadc_mapbmr.tf_mapbmr.name
		network = "1.2.3.0"
		netmask = "255.255.255.0"
	}
	
	data "citrixadc_mapbmr_bmrv4network_binding" "tf_binding" {
		name    = citrixadc_mapbmr_bmrv4network_binding.tf_binding.name
		network = citrixadc_mapbmr_bmrv4network_binding.tf_binding.network
		depends_on = [citrixadc_mapbmr_bmrv4network_binding.tf_binding]
	}
`

func TestAccMapbmr_bmrv4network_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccMapbmr_bmrv4network_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_mapbmr_bmrv4network_binding.tf_binding", "name", "tf_mapbmr"),
					resource.TestCheckResourceAttr("data.citrixadc_mapbmr_bmrv4network_binding.tf_binding", "network", "1.2.3.0"),
					resource.TestCheckResourceAttr("data.citrixadc_mapbmr_bmrv4network_binding.tf_binding", "netmask", "255.255.255.0"),
				),
			},
		},
	})
}

const testAccMapbmr_bmrv4network_binding_upgrade_basic = `
	resource "citrixadc_mapbmr" "tf_mapbmr" {
		name           = "tf_mapbmr"
		ruleipv6prefix = "2001:db8:abcd:12::/64"
		psidoffset     = 6
		eabitlength    = 16
		psidlength     = 8
	}
	resource "citrixadc_mapbmr_bmrv4network_binding" "tf_binding" {
		name    = citrixadc_mapbmr.tf_mapbmr.name
		network = "1.2.3.0"
		netmask = "255.255.255.0"
	}
`

func TestAccMapbmr_bmrv4network_binding_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckMapbmr_bmrv4network_bindingDestroy,
		Steps: []resource.TestStep{
			// Step 1: create the binding with the last SDK v2 release (writes state with the legacy comma ID).
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.2.0",
					},
				},
				Config: testAccMapbmr_bmrv4network_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMapbmr_bmrv4network_bindingExist("citrixadc_mapbmr_bmrv4network_binding.tf_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_mapbmr_bmrv4network_binding.tf_binding", "id", "tf_mapbmr,1.2.3.0"),
				),
			},
			// Step 2: same config through the current (Framework) provider. Read parses the legacy id
			// via ParseIdString and recomputes it to the new key:value format.
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccMapbmr_bmrv4network_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMapbmr_bmrv4network_bindingExist("citrixadc_mapbmr_bmrv4network_binding.tf_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_mapbmr_bmrv4network_binding.tf_binding", "id", "name:tf_mapbmr,network:1.2.3.0"),
				),
			},
		},
	})
}

func TestAccMapbmr_bmrv4network_binding_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_mapbmr_bmrv4network_binding.tf_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckMapbmr_bmrv4network_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccMapbmr_bmrv4network_binding_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckMapbmr_bmrv4network_bindingExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResourceWithArgsMap(service.Mapbmr_bmrv4network_binding.Type(), "tf_mapbmr", map[string]string{"network": "1.2.3.0", "netmask": "255.255.255.0"}); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccMapbmr_bmrv4network_binding_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckMapbmr_bmrv4network_bindingExist(resAddr, nil)),
			},
		},
	})
}
