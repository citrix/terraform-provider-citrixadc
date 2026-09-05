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

const testAccLsnclient_network_binding_basic = `

resource "citrixadc_lsnclient" "tf_lsnclient" {
	clientname = "my_lsnclient"
}

resource "citrixadc_lsnclient_network_binding" "tf_lsnclient_network_binding" {
	clientname = citrixadc_lsnclient.tf_lsnclient.clientname
	network    = "10.222.74.160"
	netmask    = "255.255.255.255"
	td         = 0
}
  
`

const testAccLsnclient_network_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion
`

func TestAccLsnclient_network_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLsnclient_network_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLsnclient_network_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsnclient_network_bindingExist("citrixadc_lsnclient_network_binding.tf_lsnclient_network_binding", nil),
				),
			},
			{
				Config: testAccLsnclient_network_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsnclient_network_bindingNotExist("citrixadc_lsnclient_network_binding.tf_lsnclient_network_binding", "my_lsnclient,10.222.74.160"),
				),
			},
		},
	})
}

func testAccCheckLsnclient_network_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lsnclient_network_binding id is set")
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

		idMap, _, err := utils.ParseIdString(bindingId, []string{"clientname", "network"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %v: %v", bindingId, err)
		}
		clientname := idMap["clientname"]
		network := idMap["network"]

		findParams := service.FindParams{
			ResourceType:             "lsnclient_network_binding",
			ResourceName:             clientname,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the matching network
		found := false
		for _, v := range dataArr {
			if v["network"].(string) == network {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("lsnclient_network_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckLsnclient_network_bindingNotExist(n string, id string) resource.TestCheckFunc {
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

		clientname := idSlice[0]
		network := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "lsnclient_network_binding",
			ResourceName:             clientname,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to hopefully not find the one with the matching network
		found := false
		for _, v := range dataArr {
			if v["network"].(string) == network {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("lsnclient_network_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckLsnclient_network_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_lsnclient_network_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("lsnclient_network_binding", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("lsnclient_network_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccLsnclient_network_bindingDataSource_basic = `

resource "citrixadc_lsnclient" "tf_lsnclient" {
	clientname = "my_lsnclient"
}

resource "citrixadc_lsnclient_network_binding" "tf_lsnclient_network_binding" {
	clientname = citrixadc_lsnclient.tf_lsnclient.clientname
	network    = "10.222.74.160"
	netmask    = "255.255.255.255"
	td         = 0
}

data "citrixadc_lsnclient_network_binding" "tf_lsnclient_network_binding" {
	clientname = citrixadc_lsnclient_network_binding.tf_lsnclient_network_binding.clientname
	network    = citrixadc_lsnclient_network_binding.tf_lsnclient_network_binding.network
}
`

func TestAccLsnclient_network_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccLsnclient_network_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_lsnclient_network_binding.tf_lsnclient_network_binding", "clientname", "my_lsnclient"),
					resource.TestCheckResourceAttr("data.citrixadc_lsnclient_network_binding.tf_lsnclient_network_binding", "network", "10.222.74.160"),
					resource.TestCheckResourceAttr("data.citrixadc_lsnclient_network_binding.tf_lsnclient_network_binding", "netmask", "255.255.255.255"),
				),
			},
		},
	})
}

const testAccLsnclient_network_binding_upgrade_basic = `

resource "citrixadc_lsnclient" "tf_lsnclient" {
	clientname = "my_lsnclient"
}

resource "citrixadc_lsnclient_network_binding" "tf_lsnclient_network_binding" {
	clientname = citrixadc_lsnclient.tf_lsnclient.clientname
	network    = "10.222.74.160"
	netmask    = "255.255.255.255"
	td         = 0
}

`

func TestAccLsnclient_network_binding_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckLsnclient_network_bindingDestroy,
		Steps: []resource.TestStep{
			// Step 1: Create the resource with the last SDK v2 release (writes state with the legacy comma ID).
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.0.0",
					},
				},
				Config: testAccLsnclient_network_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsnclient_network_bindingExist("citrixadc_lsnclient_network_binding.tf_lsnclient_network_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_lsnclient_network_binding.tf_lsnclient_network_binding", "id", "my_lsnclient,10.222.74.160"),
				),
			},
			// Step 2: Refresh the legacy-id state through the current (framework) provider.
			// Read exercises ParseIdString on the legacy id and recomputes the canonical new-format id.
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccLsnclient_network_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsnclient_network_bindingExist("citrixadc_lsnclient_network_binding.tf_lsnclient_network_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_lsnclient_network_binding.tf_lsnclient_network_binding", "id", "clientname:my_lsnclient,network:10.222.74.160"),
				),
			},
		},
	})
}

func TestAccLsnclient_network_binding_import(t *testing.T) {
	const resAddr = "citrixadc_lsnclient_network_binding.tf_lsnclient_network_binding"

	// Backward-compat: import via the LEGACY SDK v2 id. Rebuild the legacy positional id from
	// the current canonical key:value id (raw values, only the keys actually set, in legacy
	// order: clientname,network) so it matches exactly what SDK v2 wrote.
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
		ordr := []string{"clientname", "network"}
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
		CheckDestroy:             testAccCheckLsnclient_network_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccLsnclient_network_binding_basic},
			{Config: testAccLsnclient_network_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{"td"}},
			{Config: testAccLsnclient_network_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateIdFunc: legacyID, ImportStateVerify: true, ImportStateVerifyIgnore: []string{"td"}},
		},
	})
}

func TestAccLsnclient_network_binding_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_lsnclient_network_binding.tf_lsnclient_network_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLsnclient_network_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLsnclient_network_binding_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckLsnclient_network_bindingExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResourceWithArgs(service.Lsnclient_network_binding.Type(), "my_lsnclient", []string{"network:10.222.74.160"}); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccLsnclient_network_binding_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckLsnclient_network_bindingExist(resAddr, nil)),
			},
		},
	})
}
