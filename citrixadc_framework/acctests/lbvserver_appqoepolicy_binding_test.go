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

const testAccLbvserver_appqoepolicy_binding_basic_step1 = `

	resource "citrixadc_appqoeaction" "tf_appqoeaction" {
		name        = "tf_appqoeaction"
		priority    = "LOW"
		respondwith = "NS"
		delay       = 40
	}
	resource "citrixadc_appqoepolicy" "tf_appqoepolicy" {
		name   = "appqoe-pol-primd"
		rule   = "true"
		action = citrixadc_appqoeaction.tf_appqoeaction.name
	}
	resource "citrixadc_lbvserver" "tf_lbvserver" {
		name        = "tf_lbvserver"
		servicetype = "HTTP"
	}
	
	resource "citrixadc_lbvserver_appqoepolicy_binding" "foo" {
		name = citrixadc_lbvserver.tf_lbvserver.name
		policyname = citrixadc_appqoepolicy.tf_appqoepolicy.name
		bindpoint = "REQUEST"
		gotopriorityexpression = "END"
		priority = 56 
	}
`

const testAccLbvserver_appqoepolicy_binding_basic_step2 = `

	resource "citrixadc_appqoeaction" "tf_appqoeaction" {
		name        = "tf_appqoeaction"
		priority    = "LOW"
		respondwith = "NS"
		delay       = 40
	}
	resource "citrixadc_appqoepolicy" "tf_appqoepolicy" {
		name   = "appqoe-pol-primd"
		rule   = "true"
		action = citrixadc_appqoeaction.tf_appqoeaction.name
	}
	resource "citrixadc_lbvserver" "tf_lbvserver" {
		name        = "tf_lbvserver"
		servicetype = "HTTP"
	}
`

func TestAccLbvserver_appqoepolicy_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLbvserver_appqoepolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLbvserver_appqoepolicy_binding_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbvserver_appqoepolicy_bindingExist("citrixadc_lbvserver_appqoepolicy_binding.foo", nil),
				),
			},
			{
				Config: testAccLbvserver_appqoepolicy_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbvserver_appqoepolicy_bindingNotExist("citrixadc_lbvserver_appqoepolicy_binding.foo", "tf_lbvserver,appqoe-pol-primd"),
				),
			},
		},
	})
}

func TestAccLbvserver_appqoepolicy_binding_import(t *testing.T) {
	const resAddr = "citrixadc_lbvserver_appqoepolicy_binding.foo"

	// Backward-compat: import via the LEGACY SDK v2 id. Rebuild the legacy positional id from
	// the current canonical key:value id (raw values, only the keys actually set, in legacy
	// order: name,policyname) so it matches exactly what SDK v2 wrote.
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
		ordr := []string{"name", "policyname"}
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
		CheckDestroy:             testAccCheckLbvserver_appqoepolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccLbvserver_appqoepolicy_binding_basic_step1},
			{Config: testAccLbvserver_appqoepolicy_binding_basic_step1, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
			{Config: testAccLbvserver_appqoepolicy_binding_basic_step1, ResourceName: resAddr, ImportState: true, ImportStateIdFunc: legacyID, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
		},
	})
}

func testAccCheckLbvserver_appqoepolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lbvserver_appqoepolicy_binding id is set")
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

		idMap, _, err := utils.ParseIdString(bindingId, []string{"name", "policyname"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID: %v", err)
		}
		name := idMap["name"]
		policyname := idMap["policyname"]

		findParams := service.FindParams{
			ResourceType:             "lbvserver_appqoepolicy_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the right monitor name
		found := false
		for _, v := range dataArr {
			if v["policyname"].(string) == policyname {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("lbvserver_appqoepolicy_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckLbvserver_appqoepolicy_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}

		idMap, _, err := utils.ParseIdString(id, []string{"name", "policyname"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID: %v", err)
		}
		name := idMap["name"]
		policyname := idMap["policyname"]

		findParams := service.FindParams{
			ResourceType:             "lbvserver_appqoepolicy_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the right monitor name
		found := false
		for _, v := range dataArr {
			if v["policyname"].(string) == policyname {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("lbvserver_appqoepolicy_binding %s not deleted", n)
		}

		return nil
	}
}

func testAccCheckLbvserver_appqoepolicy_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_lbvserver_appqoepolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Lbvserver_appqoepolicy_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("lbvserver_appqoepolicy_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccLbvserver_appqoepolicy_bindingDataSource_basic = `

	resource "citrixadc_appqoeaction" "tf_appqoeaction" {
		name        = "tf_appqoeaction"
		priority    = "LOW"
		respondwith = "NS"
		delay       = 40
	}
	resource "citrixadc_appqoepolicy" "tf_appqoepolicy" {
		name   = "appqoe-pol-primd"
		rule   = "true"
		action = citrixadc_appqoeaction.tf_appqoeaction.name
	}
	resource "citrixadc_lbvserver" "tf_lbvserver" {
		name        = "tf_lbvserver"
		servicetype = "HTTP"
	}
	
	resource "citrixadc_lbvserver_appqoepolicy_binding" "foo" {
		name = citrixadc_lbvserver.tf_lbvserver.name
		policyname = citrixadc_appqoepolicy.tf_appqoepolicy.name
		bindpoint = "REQUEST"
		gotopriorityexpression = "END"
		priority = 56 
	}

	data "citrixadc_lbvserver_appqoepolicy_binding" "foo" {
		name = citrixadc_lbvserver_appqoepolicy_binding.foo.name
		policyname = citrixadc_lbvserver_appqoepolicy_binding.foo.policyname
		depends_on = [citrixadc_lbvserver_appqoepolicy_binding.foo]
	}
`

func TestAccLbvserver_appqoepolicy_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccLbvserver_appqoepolicy_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_lbvserver_appqoepolicy_binding.foo", "name", "tf_lbvserver"),
					resource.TestCheckResourceAttr("data.citrixadc_lbvserver_appqoepolicy_binding.foo", "policyname", "appqoe-pol-primd"),
					resource.TestCheckResourceAttr("data.citrixadc_lbvserver_appqoepolicy_binding.foo", "priority", "56"),
					resource.TestCheckResourceAttr("data.citrixadc_lbvserver_appqoepolicy_binding.foo", "gotopriorityexpression", "END"),
				),
			},
		},
	})
}

// testAccLbvserver_appqoepolicy_binding_upgrade_basic is the config used by the
// sdkv2 -> framework state-upgrade test. It reuses the same values and resource
// labels as testAccLbvserver_appqoepolicy_binding_basic_step1 so it is valid
// under BOTH the SDK v2 2.2.0 schema and the current framework schema.
const testAccLbvserver_appqoepolicy_binding_upgrade_basic = `

	resource "citrixadc_appqoeaction" "tf_appqoeaction" {
		name        = "tf_appqoeaction"
		priority    = "LOW"
		respondwith = "NS"
		delay       = 40
	}
	resource "citrixadc_appqoepolicy" "tf_appqoepolicy" {
		name   = "appqoe-pol-primd"
		rule   = "true"
		action = citrixadc_appqoeaction.tf_appqoeaction.name
	}
	resource "citrixadc_lbvserver" "tf_lbvserver" {
		name        = "tf_lbvserver"
		servicetype = "HTTP"
	}

	resource "citrixadc_lbvserver_appqoepolicy_binding" "foo" {
		name = citrixadc_lbvserver.tf_lbvserver.name
		policyname = citrixadc_appqoepolicy.tf_appqoepolicy.name
		bindpoint = "REQUEST"
		gotopriorityexpression = "END"
		priority = 56
	}
`

// TestAccLbvserver_appqoepolicy_binding_sdkv2StateUpgrade verifies that a binding
// created with the last SDK v2 release (2.2.0, legacy comma-separated ID) is
// correctly refreshed/planned/applied by the current framework provider.
func TestAccLbvserver_appqoepolicy_binding_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckLbvserver_appqoepolicy_bindingDestroy,
		Steps: []resource.TestStep{
			// Step 1: create the binding with the last SDK v2 release.
			// State is written with the LEGACY comma-separated id.
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.2.0",
					},
				},
				Config: testAccLbvserver_appqoepolicy_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbvserver_appqoepolicy_bindingExist("citrixadc_lbvserver_appqoepolicy_binding.foo", nil),
					resource.TestCheckResourceAttr("citrixadc_lbvserver_appqoepolicy_binding.foo", "id", "tf_lbvserver,appqoe-pol-primd"),
				),
			},
			// Step 2: same config, current (framework) provider. Terraform
			// refreshes the legacy-id state through the framework Read
			// (exercising ParseIdString on the legacy id) then plans/applies.
			// The framework recomputes the id on read to the new key:value form.
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccLbvserver_appqoepolicy_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbvserver_appqoepolicy_bindingExist("citrixadc_lbvserver_appqoepolicy_binding.foo", nil),
					resource.TestCheckResourceAttr("citrixadc_lbvserver_appqoepolicy_binding.foo", "id", "name:tf_lbvserver,policyname:appqoe-pol-primd"),
				),
			},
		},
	})
}

func TestAccLbvserver_appqoepolicy_binding_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_lbvserver_appqoepolicy_binding.foo"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLbvserver_appqoepolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLbvserver_appqoepolicy_binding_basic_step1,
				Check:  resource.ComposeTestCheckFunc(testAccCheckLbvserver_appqoepolicy_bindingExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResourceWithArgsMap(service.Lbvserver_appqoepolicy_binding.Type(), "tf_lbvserver", map[string]string{"policyname": "appqoe-pol-primd"}); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccLbvserver_appqoepolicy_binding_basic_step1,
				Check:  resource.ComposeTestCheckFunc(testAccCheckLbvserver_appqoepolicy_bindingExist(resAddr, nil)),
			},
		},
	})
}
