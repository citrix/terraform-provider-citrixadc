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

const testAccPolicystringmap_pattern_binding_basic_step1 = `

resource "citrixadc_policystringmap" "tf_policystringmap" {
    name = "tf_policystringmap"
    comment = "Some comment"
}

resource "citrixadc_policystringmap_pattern_binding" "tf_bind1" {
    name = citrixadc_policystringmap.tf_policystringmap.name
    key = "key1"
    value = "value1"
	comment = "key1-value1"
}

resource "citrixadc_policystringmap_pattern_binding" "tf_bind2" {
    name = citrixadc_policystringmap.tf_policystringmap.name
    key = "key2"
    value = "value2"
	comment = "key2-value2"
}
`

const testAccPolicystringmap_pattern_binding_basic_step2 = `

resource "citrixadc_policystringmap" "tf_policystringmap" {
    name = "tf_policystringmap"
    comment = "Some comment"
}

resource "citrixadc_policystringmap_pattern_binding" "tf_bind1" {
    name = citrixadc_policystringmap.tf_policystringmap.name
    key = "key1"
    value = "value1_new"
}

resource "citrixadc_policystringmap_pattern_binding" "tf_bind2" {
    name = citrixadc_policystringmap.tf_policystringmap.name
    key = "key2"
    value = "value2"
}
`

func TestAccPolicystringmap_pattern_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckPolicystringmap_pattern_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPolicystringmap_pattern_binding_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPolicystringmap_pattern_bindingExist("citrixadc_policystringmap_pattern_binding.tf_bind1", nil),
					resource.TestCheckResourceAttr("citrixadc_policystringmap_pattern_binding.tf_bind1", "key", "key1"),
					resource.TestCheckResourceAttr("citrixadc_policystringmap_pattern_binding.tf_bind1", "value", "value1"),
					resource.TestCheckResourceAttr("citrixadc_policystringmap_pattern_binding.tf_bind1", "comment", "key1-value1"),
					testAccCheckPolicystringmap_pattern_bindingExist("citrixadc_policystringmap_pattern_binding.tf_bind2", nil),
					resource.TestCheckResourceAttr("citrixadc_policystringmap_pattern_binding.tf_bind2", "key", "key2"),
					resource.TestCheckResourceAttr("citrixadc_policystringmap_pattern_binding.tf_bind2", "value", "value2"),
					resource.TestCheckResourceAttr("citrixadc_policystringmap_pattern_binding.tf_bind2", "comment", "key2-value2"),
				),
			},
			{
				Config: testAccPolicystringmap_pattern_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPolicystringmap_pattern_bindingExist("citrixadc_policystringmap_pattern_binding.tf_bind1", nil),
					resource.TestCheckResourceAttr("citrixadc_policystringmap_pattern_binding.tf_bind1", "value", "value1_new"),
					testAccCheckPolicystringmap_pattern_bindingExist("citrixadc_policystringmap_pattern_binding.tf_bind2", nil),
				),
			},
		},
	})
}

func testAccCheckPolicystringmap_pattern_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No policystringmap_pattern_binding name is set")
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
		idMap, _, err := utils.ParseIdString(bindingId, []string{"name", "key"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", bindingId, err)
		}
		name := idMap["name"]
		key := idMap["key"]

		findParams := service.FindParams{
			ResourceType:             "policystringmap_pattern_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		foundIndex := -1
		for i, v := range dataArr {
			if v["key"].(string) == key {
				foundIndex = i
				break
			}
		}
		if foundIndex == -1 {
			return fmt.Errorf("Could not find binding %s", bindingId)
		}

		return nil
	}
}

func testAccCheckPolicystringmap_pattern_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_policystringmap_pattern_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Policystringmap_pattern_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("policystringmap_pattern_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccPolicystringmap_pattern_bindingDataSource_basic = `

resource "citrixadc_policystringmap" "tf_policystringmap" {
    name = "tf_policystringmap_datasource"
    comment = "Some comment"
}

resource "citrixadc_policystringmap_pattern_binding" "tf_bind1" {
    name = citrixadc_policystringmap.tf_policystringmap.name
    key = "key1"
    value = "value1"
	comment = "key1-value1"
}

data "citrixadc_policystringmap_pattern_binding" "tf_bind1" {
  name   = citrixadc_policystringmap_pattern_binding.tf_bind1.name
  key    = citrixadc_policystringmap_pattern_binding.tf_bind1.key
  depends_on = [citrixadc_policystringmap_pattern_binding.tf_bind1]
}
`

func TestAccPolicystringmap_pattern_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccPolicystringmap_pattern_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_policystringmap_pattern_binding.tf_bind1", "name", "tf_policystringmap_datasource"),
					resource.TestCheckResourceAttr("data.citrixadc_policystringmap_pattern_binding.tf_bind1", "key", "key1"),
					resource.TestCheckResourceAttr("data.citrixadc_policystringmap_pattern_binding.tf_bind1", "value", "value1"),
					resource.TestCheckResourceAttr("data.citrixadc_policystringmap_pattern_binding.tf_bind1", "comment", "key1-value1"),
				),
			},
		},
	})
}

const testAccPolicystringmap_pattern_binding_upgrade_basic = `

resource "citrixadc_policystringmap" "tf_policystringmap" {
    name = "tf_policystringmap"
    comment = "Some comment"
}

resource "citrixadc_policystringmap_pattern_binding" "tf_bind1" {
    name = citrixadc_policystringmap.tf_policystringmap.name
    key = "key1"
    value = "value1"
    comment = "key1-value1"
}
`

// TestAccPolicystringmap_pattern_binding_sdkv2StateUpgrade verifies that a binding
// created by the LAST SDK v2 release (2.2.0) — which writes the legacy comma-joined
// id "name,key" — is refreshed and re-applied correctly by the CURRENT framework
// provider. Step 2 exercises ParseIdString on the legacy id during the framework
// Read; the framework SetAttrFromGet then recomputes data.Id into the canonical new
// "key:value" format, so the id upgrades in place.
func TestAccPolicystringmap_pattern_binding_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckPolicystringmap_pattern_bindingDestroy,
		Steps: []resource.TestStep{
			// Step 1: create with the last SDK v2 release from the registry. This
			// writes state carrying the LEGACY comma-joined id "name,key".
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.2.0",
					},
				},
				Config: testAccPolicystringmap_pattern_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPolicystringmap_pattern_bindingExist("citrixadc_policystringmap_pattern_binding.tf_bind1", nil),
					resource.TestCheckResourceAttr("citrixadc_policystringmap_pattern_binding.tf_bind1", "id", "tf_policystringmap,key1"),
				),
			},
			// Step 2: same config through the CURRENT framework provider. Terraform
			// refreshes the legacy-id state through the framework Read (exercising
			// ParseIdString on the legacy id), then plans/applies. SetAttrFromGet
			// recomputes the id to the new "key:value" format.
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccPolicystringmap_pattern_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPolicystringmap_pattern_bindingExist("citrixadc_policystringmap_pattern_binding.tf_bind1", nil),
					resource.TestCheckResourceAttr("citrixadc_policystringmap_pattern_binding.tf_bind1", "id", "key:key1,name:tf_policystringmap"),
				),
			},
		},
	})
}

func TestAccPolicystringmap_pattern_binding_import(t *testing.T) {
	const resAddr = "citrixadc_policystringmap_pattern_binding.tf_bind1"

	// Backward-compat: import via the LEGACY SDK v2 id. Rebuild the legacy positional id from
	// the current canonical key:value id (raw values, only the keys actually set, in legacy
	// order: name,key) so it matches exactly what SDK v2 wrote.
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
		ordr := []string{"name", "key"}
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
		CheckDestroy:             testAccCheckPolicystringmap_pattern_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccPolicystringmap_pattern_binding_basic_step1},
			{Config: testAccPolicystringmap_pattern_binding_basic_step1, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
			{Config: testAccPolicystringmap_pattern_binding_basic_step1, ResourceName: resAddr, ImportState: true, ImportStateIdFunc: legacyID, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
		},
	})
}

// TestAccPolicystringmap_pattern_binding_selfHealing verifies the provider re-creates
// the binding when it is deleted out-of-band between apply steps (drift recovery).
func TestAccPolicystringmap_pattern_binding_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_policystringmap_pattern_binding.tf_bind1"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckPolicystringmap_pattern_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPolicystringmap_pattern_binding_basic_step1,
				Check:  resource.ComposeTestCheckFunc(testAccCheckPolicystringmap_pattern_bindingExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResourceWithArgs(service.Policystringmap_pattern_binding.Type(), "tf_policystringmap", []string{"key:key1"}); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccPolicystringmap_pattern_binding_basic_step1,
				Check:  resource.ComposeTestCheckFunc(testAccCheckPolicystringmap_pattern_bindingExist(resAddr, nil)),
			},
		},
	})
}
