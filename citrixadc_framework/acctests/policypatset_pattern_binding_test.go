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

const testAccPolicypatset_pattern_binding_basic_step1 = `

resource "citrixadc_policypatset" "tf_patset" {
    name = "tf_patset"
    comment = "some comment"
}

resource "citrixadc_policypatset_pattern_binding" "tf_bind" {
    name = citrixadc_policypatset.tf_patset.name
    string = "pattern1,/postfix"
}

`

const testAccPolicypatset_pattern_binding_basic_step2 = `

resource "citrixadc_policypatset" "tf_patset" {
    name = "tf_patset"
    comment = "some comment"
}

resource "citrixadc_policypatset_pattern_binding" "tf_bind" {
    name = citrixadc_policypatset.tf_patset.name
    string = "pattern2,/postfix"
}

`

func TestAccPolicypatset_pattern_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckPolicypatset_pattern_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPolicypatset_pattern_binding_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPolicypatset_pattern_bindingExist("citrixadc_policypatset_pattern_binding.tf_bind", nil),
				),
			},
			{
				Config: testAccPolicypatset_pattern_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPolicypatset_pattern_bindingExist("citrixadc_policypatset_pattern_binding.tf_bind", nil),
				),
			},
		},
	})
}

func testAccCheckPolicypatset_pattern_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No policypatset_pattern_binding name is set")
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

		idMap, _, err := utils.ParseIdString(bindingId, []string{"name", "string"}, nil)
		if err != nil {
			return err
		}
		name := idMap["name"]
		stringText := idMap["string"]

		findParams := service.FindParams{
			ResourceType:             "policypatset_pattern_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 2823,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the right String
		foundIndex := -1
		for i, v := range dataArr {
			if v["String"].(string) == stringText {
				foundIndex = i
				break
			}
		}

		// Resource is missing
		if foundIndex == -1 {
			return fmt.Errorf("FindResourceArrayWithParams  could not find pattern_binding %v", bindingId)
		}

		return nil
	}
}

func testAccCheckPolicypatset_pattern_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_policypatset_pattern_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Policypatset_pattern_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("policypatset_pattern_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccPolicypatset_pattern_bindingDataSource_basic = `

resource "citrixadc_policypatset" "tf_patset" {
    name = "tf_patset_datasource"
    comment = "some comment"
}

resource "citrixadc_policypatset_pattern_binding" "tf_bind" {
    name = citrixadc_policypatset.tf_patset.name
    string = "pattern1,/postfix"
}

data "citrixadc_policypatset_pattern_binding" "tf_bind" {
  name   = citrixadc_policypatset_pattern_binding.tf_bind.name
  string = citrixadc_policypatset_pattern_binding.tf_bind.string
  depends_on = [citrixadc_policypatset_pattern_binding.tf_bind]
}
`

func TestAccPolicypatset_pattern_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccPolicypatset_pattern_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_policypatset_pattern_binding.tf_bind", "name", "tf_patset_datasource"),
					resource.TestCheckResourceAttr("data.citrixadc_policypatset_pattern_binding.tf_bind", "string", "pattern1,/postfix"),
				),
			},
		},
	})
}

const testAccPolicypatset_pattern_binding_upgrade_basic = `

resource "citrixadc_policypatset" "tf_patset" {
    name = "tf_patset"
    comment = "some comment"
}

resource "citrixadc_policypatset_pattern_binding" "tf_bind" {
    name = citrixadc_policypatset.tf_patset.name
    string = "pattern1,/postfix"
}

`

func TestAccPolicypatset_pattern_binding_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckPolicypatset_pattern_bindingDestroy,
		Steps: []resource.TestStep{
			// Step 1: create the resource with the last SDK v2 release (writes state with the legacy id).
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.2.0",
					},
				},
				Config: testAccPolicypatset_pattern_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPolicypatset_pattern_bindingExist("citrixadc_policypatset_pattern_binding.tf_bind", nil),
					resource.TestCheckResourceAttr("citrixadc_policypatset_pattern_binding.tf_bind", "id", "tf_patset,pattern1,/postfix"),
				),
			},
			// Step 2: refresh/plan/apply the legacy-id state through the current framework provider.
			// The framework Read recomputes the id into the new key:value format.
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccPolicypatset_pattern_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPolicypatset_pattern_bindingExist("citrixadc_policypatset_pattern_binding.tf_bind", nil),
					resource.TestCheckResourceAttr("citrixadc_policypatset_pattern_binding.tf_bind", "id", "name:tf_patset,string:pattern1%2C%2Fpostfix"),
				),
			},
		},
	})
}

func TestAccPolicypatset_pattern_binding_import(t *testing.T) {
	const resAddr = "citrixadc_policypatset_pattern_binding.tf_bind"

	// Backward-compat: import via the LEGACY SDK v2 id. Rebuild the legacy positional id from
	// the current canonical key:value id (raw values, only the keys actually set, in legacy
	// order: name,string) so it matches exactly what SDK v2 wrote.
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
		ordr := []string{"name", "string"}
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
		CheckDestroy:             testAccCheckPolicypatset_pattern_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccPolicypatset_pattern_binding_basic_step1},
			{Config: testAccPolicypatset_pattern_binding_basic_step1, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
			{Config: testAccPolicypatset_pattern_binding_basic_step1, ResourceName: resAddr, ImportState: true, ImportStateIdFunc: legacyID, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
		},
	})
}

func TestAccPolicypatset_pattern_binding_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_policypatset_pattern_binding.tf_bind"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckPolicypatset_pattern_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPolicypatset_pattern_binding_basic_step1,
				Check:  resource.ComposeTestCheckFunc(testAccCheckPolicypatset_pattern_bindingExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					// Mirrors the resource Delete: the pattern string is url.QueryEscape'd
					// ("pattern1,/postfix" -> "pattern1%2C%2Fpostfix") and passed as the "String" arg.
					if err := client.DeleteResourceWithArgsMap(service.Policypatset_pattern_binding.Type(), "tf_patset", map[string]string{"String": "pattern1%2C%2Fpostfix"}); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccPolicypatset_pattern_binding_basic_step1,
				Check:  resource.ComposeTestCheckFunc(testAccCheckPolicypatset_pattern_bindingExist(resAddr, nil)),
			},
		},
	})
}
