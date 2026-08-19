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

const testAccNspartition_bridgegroup_binding_basic = `
	resource "citrixadc_nspartition" "tf_nspartition" {
		partitionname = "tf_nspartition"
		maxbandwidth  = 10240
		minbandwidth  = 512
		maxconn       = 512
		maxmemlimit   = 11
	}
	resource "citrixadc_bridgegroup" "tf_bridgegroup" {
		bridgegroup_id     = 2
		dynamicrouting     = "DISABLED"
		ipv6dynamicrouting = "DISABLED"
	}
	resource "citrixadc_nspartition_bridgegroup_binding" "tf_binding" {
		partitionname = citrixadc_nspartition.tf_nspartition.partitionname
		bridgegroup   = citrixadc_bridgegroup.tf_bridgegroup.bridgegroup_id
	}
`

const testAccNspartition_bridgegroup_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion
	resource "citrixadc_nspartition" "tf_nspartition" {
		partitionname = "tf_nspartition"
		maxbandwidth  = 10240
		minbandwidth  = 512
		maxconn       = 512
		maxmemlimit   = 11
	}
	resource "citrixadc_bridgegroup" "tf_bridgegroup" {
		bridgegroup_id     = 2
		dynamicrouting     = "DISABLED"
		ipv6dynamicrouting = "DISABLED"
	}
`

func TestAccNspartition_bridgegroup_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNspartition_bridgegroup_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNspartition_bridgegroup_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNspartition_bridgegroup_bindingExist("citrixadc_nspartition_bridgegroup_binding.tf_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_nspartition_bridgegroup_binding.tf_binding", "partitionname", "tf_nspartition"),
					resource.TestCheckResourceAttr("citrixadc_nspartition_bridgegroup_binding.tf_binding", "bridgegroup", "2"),
				),
			},
			{
				Config: testAccNspartition_bridgegroup_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNspartition_bridgegroup_bindingNotExist("citrixadc_nspartition_bridgegroup_binding.tf_binding", "tf_nspartition,2"),
				),
			},
		},
	})
}

func testAccCheckNspartition_bridgegroup_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No nspartition_bridgegroup_binding id is set")
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

		idMap, _, err := utils.ParseIdString(bindingId, []string{"partitionname", "bridgegroup"}, nil)
		if err != nil {
			return err
		}
		partitionname := idMap["partitionname"]
		bridgegroup := idMap["bridgegroup"]

		findParams := service.FindParams{
			ResourceType:             "nspartition_bridgegroup_binding",
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
			if v["bridgegroup"].(string) == bridgegroup {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("nspartition_bridgegroup_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckNspartition_bridgegroup_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}
		idMap, _, err := utils.ParseIdString(id, []string{"partitionname", "bridgegroup"}, nil)
		if err != nil {
			return err
		}
		partitionname := idMap["partitionname"]
		bridgegroup := idMap["bridgegroup"]

		findParams := service.FindParams{
			ResourceType:             "nspartition_bridgegroup_binding",
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
			if v["bridgegroup"].(string) == bridgegroup {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("nspartition_bridgegroup_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckNspartition_bridgegroup_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_nspartition_bridgegroup_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Nspartition_bridgegroup_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("nspartition_bridgegroup_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccNspartition_bridgegroup_bindingDataSource_basic = `
	resource "citrixadc_nspartition" "tf_nspartition" {
		partitionname = "tf_nspartition"
		maxbandwidth  = 10240
		minbandwidth  = 512
		maxconn       = 512
		maxmemlimit   = 11
	}
	resource "citrixadc_bridgegroup" "tf_bridgegroup" {
		bridgegroup_id     = 2
		dynamicrouting     = "DISABLED"
		ipv6dynamicrouting = "DISABLED"
	}
	resource "citrixadc_nspartition_bridgegroup_binding" "tf_binding" {
		partitionname = citrixadc_nspartition.tf_nspartition.partitionname
		bridgegroup   = citrixadc_bridgegroup.tf_bridgegroup.bridgegroup_id
	}

	data "citrixadc_nspartition_bridgegroup_binding" "tf_binding" {
		partitionname = citrixadc_nspartition_bridgegroup_binding.tf_binding.partitionname
		bridgegroup   = citrixadc_nspartition_bridgegroup_binding.tf_binding.bridgegroup
		depends_on    = [citrixadc_nspartition_bridgegroup_binding.tf_binding]
	}
`

func TestAccNspartition_bridgegroup_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNspartition_bridgegroup_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_nspartition_bridgegroup_binding.tf_binding", "partitionname", "tf_nspartition"),
					resource.TestCheckResourceAttr("data.citrixadc_nspartition_bridgegroup_binding.tf_binding", "bridgegroup", "2"),
				),
			},
		},
	})
}

const testAccNspartition_bridgegroup_binding_upgrade_basic = `
	resource "citrixadc_nspartition" "tf_nspartition" {
		partitionname = "tf_nspartition"
		maxbandwidth  = 10240
		minbandwidth  = 512
		maxconn       = 512
		maxmemlimit   = 11
	}
	resource "citrixadc_bridgegroup" "tf_bridgegroup" {
		bridgegroup_id     = 2
		dynamicrouting     = "DISABLED"
		ipv6dynamicrouting = "DISABLED"
	}
	resource "citrixadc_nspartition_bridgegroup_binding" "tf_binding" {
		partitionname = citrixadc_nspartition.tf_nspartition.partitionname
		bridgegroup   = citrixadc_bridgegroup.tf_bridgegroup.bridgegroup_id
	}
`

func TestAccNspartition_bridgegroup_binding_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckNspartition_bridgegroup_bindingDestroy,
		Steps: []resource.TestStep{
			// Step 1: Create the resource with the last SDK v2 release (writes state with the legacy comma ID).
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.2.0",
					},
				},
				Config: testAccNspartition_bridgegroup_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNspartition_bridgegroup_bindingExist("citrixadc_nspartition_bridgegroup_binding.tf_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_nspartition_bridgegroup_binding.tf_binding", "id", "tf_nspartition,2"),
				),
			},
			// Step 2: Refresh the legacy-id state through the current (framework) provider.
			// Read exercises ParseIdString on the legacy id and recomputes the canonical new-format id.
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccNspartition_bridgegroup_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNspartition_bridgegroup_bindingExist("citrixadc_nspartition_bridgegroup_binding.tf_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_nspartition_bridgegroup_binding.tf_binding", "id", "bridgegroup:2,partitionname:tf_nspartition"),
				),
			},
		},
	})
}

func TestAccNspartition_bridgegroup_binding_import(t *testing.T) {
	const resAddr = "citrixadc_nspartition_bridgegroup_binding.tf_binding"

	// Backward-compat: import via the LEGACY SDK v2 id. Rebuild the legacy positional id from
	// the current canonical key:value id (raw values, only the keys actually set, in legacy
	// order: partitionname,bridgegroup) so it matches exactly what SDK v2 wrote.
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
		ordr := []string{"partitionname", "bridgegroup"}
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
		CheckDestroy:             testAccCheckNspartition_bridgegroup_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccNspartition_bridgegroup_binding_basic},
			{Config: testAccNspartition_bridgegroup_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
			{Config: testAccNspartition_bridgegroup_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateIdFunc: legacyID, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
		},
	})
}

func TestAccNspartition_bridgegroup_binding_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_nspartition_bridgegroup_binding.tf_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNspartition_bridgegroup_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNspartition_bridgegroup_binding_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckNspartition_bridgegroup_bindingExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResourceWithArgsMap(service.Nspartition_bridgegroup_binding.Type(), "tf_nspartition", map[string]string{"bridgegroup": "2"}); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccNspartition_bridgegroup_binding_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckNspartition_bridgegroup_bindingExist(resAddr, nil)),
			},
		},
	})
}
