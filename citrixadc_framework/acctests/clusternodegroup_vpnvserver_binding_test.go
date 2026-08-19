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

const testAccClusternodegroup_vpnvserver_binding_basic = `

	resource "citrixadc_clusternodegroup" "tf_clusternodegroup" {
		name   = "my_clusternode_ds"
		strict = "NO"
		sticky = "YES"
	}

	resource "citrixadc_clusternodegroup_clusternode_binding" "tf_clusternodegroup_clusternode_binding" {
		name = citrixadc_clusternodegroup.tf_clusternodegroup.name
		node = 0
	}

	resource "citrixadc_vpnvserver" "tf_vpnvserver" {
		name        = "my_vpn_vserver_ds"
		servicetype = "SSL"
	}

	resource "citrixadc_clusternodegroup_vpnvserver_binding" "tf_clusternodegroup_vpnvserver_binding" {
		name    = citrixadc_clusternodegroup.tf_clusternodegroup.name
		vserver = citrixadc_vpnvserver.tf_vpnvserver.name
		depends_on = [citrixadc_clusternodegroup_clusternode_binding.tf_clusternodegroup_clusternode_binding]
	}
	
`

const testAccClusternodegroup_vpnvserver_binding_basic_step2 = `

	resource "citrixadc_clusternodegroup" "tf_clusternodegroup" {
		name   = "my_clusternode_ds"
		strict = "NO"
		sticky = "YES"
	}

	resource "citrixadc_clusternodegroup_clusternode_binding" "tf_clusternodegroup_clusternode_binding" {
		name = citrixadc_clusternodegroup.tf_clusternodegroup.name
		node = 0
	}

	resource "citrixadc_vpnvserver" "tf_vpnvserver" {
		name        = "my_vpn_vserver_ds"
		servicetype = "SSL"
	}
	# Keep the above bound resources without the actual binding to check proper deletion
`

func TestAccClusternodegroup_vpnvserver_binding_basic(t *testing.T) {
	if adcTestbed != "CLUSTER" {
		t.Skipf("ADC testbed is %s. Expected CLUSTER.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckClusternodegroup_vpnvserver_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccClusternodegroup_vpnvserver_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClusternodegroup_vpnvserver_bindingExist("citrixadc_clusternodegroup_vpnvserver_binding.tf_clusternodegroup_vpnvserver_binding", nil),
				),
			},
			{
				Config: testAccClusternodegroup_vpnvserver_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClusternodegroup_vpnvserver_bindingNotExist("citrixadc_clusternodegroup_vpnvserver_binding.tf_clusternodegroup_vpnvserver_binding", "my_vpn_group,my_vpnvserver"),
				),
			},
		},
	})
}

func testAccCheckClusternodegroup_vpnvserver_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No clusternodegroup_vpnvserver_binding id is set")
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

		idMap, _, err := utils.ParseIdString(bindingId, []string{"name", "vserver"}, nil)
		if err != nil {
			return err
		}
		name := idMap["name"]
		vserver := idMap["vserver"]

		findParams := service.FindParams{
			ResourceType:             "clusternodegroup_vpnvserver_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the matching vserver
		found := false
		for _, v := range dataArr {
			if v["vserver"].(string) == vserver {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("clusternodegroup_vpnvserver_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckClusternodegroup_vpnvserver_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		idMap, _, err := utils.ParseIdString(id, []string{"name", "vserver"}, nil)
		if err != nil {
			return err
		}
		name := idMap["name"]
		vserver := idMap["vserver"]

		findParams := service.FindParams{
			ResourceType:             "clusternodegroup_vpnvserver_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to hopefully not find the one with the matching vserver
		found := false
		for _, v := range dataArr {
			if v["vserver"].(string) == vserver {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("clusternodegroup_vpnvserver_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckClusternodegroup_vpnvserver_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_clusternodegroup_vpnvserver_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Clusternodegroup_vpnvserver_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("clusternodegroup_vpnvserver_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccClusternodegroup_vpnvserver_bindingDataSource_basic = `

	resource "citrixadc_clusternodegroup" "tf_clusternodegroup" {
		name   = "my_tf_group"
		strict = "NO"
		sticky = "YES"
	}

	resource "citrixadc_clusternodegroup_clusternode_binding" "tf_clusternodegroup_clusternode_binding" {
		name = citrixadc_clusternodegroup.tf_clusternodegroup.name
		node = 0
	}

	resource "citrixadc_vpnvserver" "tf_vpnvserver" {
		name        = "my_vpn_vserver_ds"
		servicetype = "SSL"
	}

	resource "citrixadc_clusternodegroup_vpnvserver_binding" "tf_clusternodegroup_vpnvserver_binding" {
		name    = citrixadc_clusternodegroup.tf_clusternodegroup.name
		vserver = citrixadc_vpnvserver.tf_vpnvserver.name
		depends_on = [citrixadc_clusternodegroup_clusternode_binding.tf_clusternodegroup_clusternode_binding]
	}

	data "citrixadc_clusternodegroup_vpnvserver_binding" "tf_clusternodegroup_vpnvserver_binding" {
		name    = citrixadc_clusternodegroup_vpnvserver_binding.tf_clusternodegroup_vpnvserver_binding.name
		vserver = citrixadc_clusternodegroup_vpnvserver_binding.tf_clusternodegroup_vpnvserver_binding.vserver
		depends_on = [citrixadc_clusternodegroup_vpnvserver_binding.tf_clusternodegroup_vpnvserver_binding]
	}
`

func TestAccclusternodegroup_vpnvserver_bindingDataSource_basic(t *testing.T) {
	if adcTestbed != "CLUSTER" {
		t.Skipf("ADC testbed is %s. Expected CLUSTER.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccClusternodegroup_vpnvserver_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_clusternodegroup_vpnvserver_binding.tf_clusternodegroup_vpnvserver_binding", "name", "my_tf_group"),
					resource.TestCheckResourceAttr("data.citrixadc_clusternodegroup_vpnvserver_binding.tf_clusternodegroup_vpnvserver_binding", "vserver", "my_vpn_vserver_ds"),
				),
			},
		},
	})
}

// Config used by the SDK v2 -> Framework state-upgrade test. It reuses the same
// values as testAccClusternodegroup_vpnvserver_binding_basic and keeps the exact
// resource labels so the Exist/Destroy helpers and addresses match. It is valid
// under both the SDK v2 2.2.0 schema and the current provider schema.
const testAccClusternodegroup_vpnvserver_binding_upgrade_basic = `

	resource "citrixadc_clusternodegroup" "tf_clusternodegroup" {
		name   = "my_clusternode_ds"
		strict = "NO"
		sticky = "YES"
	}

	resource "citrixadc_clusternodegroup_clusternode_binding" "tf_clusternodegroup_clusternode_binding" {
		name = citrixadc_clusternodegroup.tf_clusternodegroup.name
		node = 0
	}

	resource "citrixadc_vpnvserver" "tf_vpnvserver" {
		name        = "my_vpn_vserver_ds"
		servicetype = "SSL"
	}

	resource "citrixadc_clusternodegroup_vpnvserver_binding" "tf_clusternodegroup_vpnvserver_binding" {
		name    = citrixadc_clusternodegroup.tf_clusternodegroup.name
		vserver = citrixadc_vpnvserver.tf_vpnvserver.name
		depends_on = [citrixadc_clusternodegroup_clusternode_binding.tf_clusternodegroup_clusternode_binding]
	}

`

func TestAccClusternodegroup_vpnvserver_binding_sdkv2StateUpgrade(t *testing.T) {
	if adcTestbed != "CLUSTER" {
		t.Skipf("ADC testbed is %s. Expected CLUSTER.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckClusternodegroup_vpnvserver_bindingDestroy,
		Steps: []resource.TestStep{
			{
				// Step 1: create with the last SDK v2 release (v2.2.0) from the
				// Terraform registry. This writes state with the legacy comma id.
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.2.0",
					},
				},
				Config: testAccClusternodegroup_vpnvserver_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClusternodegroup_vpnvserver_bindingExist("citrixadc_clusternodegroup_vpnvserver_binding.tf_clusternodegroup_vpnvserver_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_clusternodegroup_vpnvserver_binding.tf_clusternodegroup_vpnvserver_binding", "id", "my_clusternode_ds,my_vpn_vserver_ds"),
				),
			},
			{
				// Step 2: refresh/plan/apply the same config through the current
				// provider. The current provider still serves this resource via
				// SDK v2 (the Framework resource is not registered), so the Read
				// does not recompute the id to the new key:value format; the id
				// stays the legacy value. Assert only that the binding survives.
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccClusternodegroup_vpnvserver_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClusternodegroup_vpnvserver_bindingExist("citrixadc_clusternodegroup_vpnvserver_binding.tf_clusternodegroup_vpnvserver_binding", nil),
				),
			},
		},
	})
}

func TestAccClusternodegroup_vpnvserver_binding_import(t *testing.T) {
	if adcTestbed != "CLUSTER" {
		t.Skipf("ADC testbed is %s. Expected CLUSTER.", adcTestbed)
	}
	const resAddr = "citrixadc_clusternodegroup_vpnvserver_binding.tf_clusternodegroup_vpnvserver_binding"

	// Backward-compat: import via the LEGACY SDK v2 id. Rebuild the legacy positional id from
	// the current canonical key:value id (raw values, only the keys actually set, in legacy
	// order: name,vserver) so it matches exactly what SDK v2 wrote.
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
		ordr := []string{"name", "vserver"}
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
		CheckDestroy:             testAccCheckClusternodegroup_vpnvserver_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccClusternodegroup_vpnvserver_binding_basic},
			{Config: testAccClusternodegroup_vpnvserver_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
			{Config: testAccClusternodegroup_vpnvserver_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateIdFunc: legacyID, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
		},
	})
}

func TestAccClusternodegroup_vpnvserver_binding_selfHealing(t *testing.T) {
	if adcTestbed != "CLUSTER" {
		t.Skipf("ADC testbed is %s. Expected CLUSTER.", adcTestbed)
	}
	const resAddr = "citrixadc_clusternodegroup_vpnvserver_binding.tf_clusternodegroup_vpnvserver_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckClusternodegroup_vpnvserver_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccClusternodegroup_vpnvserver_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClusternodegroup_vpnvserver_bindingExist(resAddr, nil),
				),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResourceWithArgsMap(service.Clusternodegroup_vpnvserver_binding.Type(), "my_clusternode_ds", map[string]string{"vserver": "my_vpn_vserver_ds"}); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccClusternodegroup_vpnvserver_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClusternodegroup_vpnvserver_bindingExist(resAddr, nil),
				),
			},
		},
	})
}
