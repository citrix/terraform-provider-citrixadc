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

const testAccNspartition_vxlan_binding_basic = `
	resource "citrixadc_nspartition" "tf_nspartition" {
		partitionname = "tf_nspartition"
		maxbandwidth  = 10240
		minbandwidth  = 512
		maxconn       = 512
		maxmemlimit   = 11
	}
	resource "citrixadc_vxlan" "tf_vxlan" {
		vxlanid            = 123
		port               = 33
		dynamicrouting     = "DISABLED"
		ipv6dynamicrouting = "DISABLED"
		innervlantagging   = "ENABLED"
	}
	resource "citrixadc_nspartition_vxlan_binding" "tf_binding" {
		partitionname = citrixadc_nspartition.tf_nspartition.partitionname
		vxlan         = citrixadc_vxlan.tf_vxlan.vxlanid
	}
`

const testAccNspartition_vxlan_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion
	resource "citrixadc_nspartition" "tf_nspartition" {
		partitionname = "tf_nspartition"
		maxbandwidth  = 10240
		minbandwidth  = 512
		maxconn       = 512
		maxmemlimit   = 11
	}
	resource "citrixadc_vxlan" "tf_vxlan" {
		vxlanid            = 123
		port               = 33
		dynamicrouting     = "DISABLED"
		ipv6dynamicrouting = "DISABLED"
		innervlantagging   = "ENABLED"
	}
`

func TestAccNspartition_vxlan_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNspartition_vxlan_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNspartition_vxlan_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNspartition_vxlan_bindingExist("citrixadc_nspartition_vxlan_binding.tf_binding", nil),
				),
			},
			{
				Config: testAccNspartition_vxlan_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNspartition_vxlan_bindingNotExist("citrixadc_nspartition_vxlan_binding.tf_binding", "tf_nspartition,123"),
				),
			},
		},
	})
}

func TestAccNspartition_vxlan_binding_import(t *testing.T) {
	const resAddr = "citrixadc_nspartition_vxlan_binding.tf_binding"

	// Backward-compat: import via the LEGACY SDK v2 id. Rebuild the legacy positional id from
	// the current canonical key:value id (raw values, only the keys actually set, in legacy
	// order: partitionname,vxlan) so it matches exactly what SDK v2 wrote.
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
		ordr := []string{"partitionname", "vxlan"}
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
		CheckDestroy:             testAccCheckNspartition_vxlan_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccNspartition_vxlan_binding_basic},
			{Config: testAccNspartition_vxlan_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
			{Config: testAccNspartition_vxlan_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateIdFunc: legacyID, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
		},
	})
}

func testAccCheckNspartition_vxlan_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No nspartition_vxlan_binding id is set")
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

		idMap, _, err := utils.ParseIdString(bindingId, []string{"partitionname", "vxlan"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID: %v", err)
		}
		partitionname := idMap["partitionname"]
		vxlan := idMap["vxlan"]

		findParams := service.FindParams{
			ResourceType:             "nspartition_vxlan_binding",
			ResourceName:             partitionname,
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
			if v["vxlan"].(string) == vxlan {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("nspartition_vxlan_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckNspartition_vxlan_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		idMap, _, err := utils.ParseIdString(id, []string{"partitionname", "vxlan"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID: %v", err)
		}
		partitionname := idMap["partitionname"]
		vxlan := idMap["vxlan"]

		findParams := service.FindParams{
			ResourceType:             "nspartition_vxlan_binding",
			ResourceName:             partitionname,
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
			if v["vxlan"].(string) == vxlan {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("nspartition_vxlan_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckNspartition_vxlan_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_nspartition_vxlan_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("nspartition_vxlan_binding", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("nspartition_vxlan_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccNspartition_vxlan_binding_upgrade_basic = `
	resource "citrixadc_nspartition" "tf_nspartition" {
		partitionname = "tf_nspartition"
		maxbandwidth  = 10240
		minbandwidth  = 512
		maxconn       = 512
		maxmemlimit   = 11
	}
	resource "citrixadc_vxlan" "tf_vxlan" {
		vxlanid            = 123
		port               = 33
		dynamicrouting     = "DISABLED"
		ipv6dynamicrouting = "DISABLED"
		innervlantagging   = "ENABLED"
	}
	resource "citrixadc_nspartition_vxlan_binding" "tf_binding" {
		partitionname = citrixadc_nspartition.tf_nspartition.partitionname
		vxlan         = citrixadc_vxlan.tf_vxlan.vxlanid
	}
`

func TestAccNspartition_vxlan_binding_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckNspartition_vxlan_bindingDestroy,
		Steps: []resource.TestStep{
			// Step 1: Create the resource with the last SDK v2 release (writes state with the legacy comma ID).
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.2.0",
					},
				},
				Config: testAccNspartition_vxlan_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNspartition_vxlan_bindingExist("citrixadc_nspartition_vxlan_binding.tf_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_nspartition_vxlan_binding.tf_binding", "id", "tf_nspartition,123"),
				),
			},
			// Step 2: Refresh the legacy-id state through the current (framework) provider.
			// Read exercises ParseIdString on the legacy id and recomputes the canonical new-format id.
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccNspartition_vxlan_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNspartition_vxlan_bindingExist("citrixadc_nspartition_vxlan_binding.tf_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_nspartition_vxlan_binding.tf_binding", "id", "partitionname:tf_nspartition,vxlan:123"),
				),
			},
		},
	})
}

const testAccNspartition_vxlan_bindingDataSource_basic = `
	resource "citrixadc_nspartition" "tf_nspartition" {
		partitionname = "tf_nspartition"
		maxbandwidth  = 10240
		minbandwidth  = 512
		maxconn       = 512
		maxmemlimit   = 11
	}
	resource "citrixadc_vxlan" "tf_vxlan" {
		vxlanid            = 123
		port               = 33
		dynamicrouting     = "DISABLED"
		ipv6dynamicrouting = "DISABLED"
		innervlantagging   = "ENABLED"
	}
	resource "citrixadc_nspartition_vxlan_binding" "tf_binding" {
		partitionname = citrixadc_nspartition.tf_nspartition.partitionname
		vxlan         = citrixadc_vxlan.tf_vxlan.vxlanid
	}

	data "citrixadc_nspartition_vxlan_binding" "tf_binding" {
		partitionname = citrixadc_nspartition_vxlan_binding.tf_binding.partitionname
		vxlan         = citrixadc_nspartition_vxlan_binding.tf_binding.vxlan
	}
`

func TestAccNspartition_vxlan_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNspartition_vxlan_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_nspartition_vxlan_binding.tf_binding", "partitionname", "tf_nspartition"),
					resource.TestCheckResourceAttr("data.citrixadc_nspartition_vxlan_binding.tf_binding", "vxlan", "123"),
				),
			},
		},
	})
}

func TestAccNspartition_vxlan_binding_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_nspartition_vxlan_binding.tf_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNspartition_vxlan_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNspartition_vxlan_binding_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckNspartition_vxlan_bindingExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResourceWithArgsMap(service.Nspartition_vxlan_binding.Type(), "tf_nspartition", map[string]string{"vxlan": "123"}); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccNspartition_vxlan_binding_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckNspartition_vxlan_bindingExist(resAddr, nil)),
			},
		},
	})
}
