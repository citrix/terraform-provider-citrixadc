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

const testAccCsvserver_botpolicy_binding_basic = `
	resource "citrixadc_csvserver_botpolicy_binding" "tf_csvserver_botpolicy_binding" {
        name = citrixadc_csvserver.tf_csvserver.name
        policyname = citrixadc_botpolicy.tf_botpolicy.name
        priority = 5
	}

	resource "citrixadc_csvserver" "tf_csvserver" {
		name = "tf_csvserver"
		ipv46 = "10.202.11.11"
		port = 8080
		servicetype = "HTTP"
	}

	resource citrixadc_botpolicy tf_botpolicy {
		name = "tf_botpolicy"  
		profilename = "BOT_BYPASS"
		rule  = "true"
	}
`

const testAccCsvserver_botpolicy_binding_basic_step2 = `
	resource "citrixadc_csvserver" "tf_csvserver" {
		name = "tf_csvserver"
		ipv46 = "10.202.11.11"
		port = 8080
		servicetype = "HTTP"
	}

	resource citrixadc_botpolicy tf_botpolicy {
		name = "tf_botpolicy"  
		profilename = "BOT_BYPASS"
		rule  = "true"
	}
`

func TestAccCsvserver_botpolicy_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCsvserver_botpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCsvserver_botpolicy_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCsvserver_botpolicy_bindingExist("citrixadc_csvserver_botpolicy_binding.tf_csvserver_botpolicy_binding", nil),
				),
			},
			{
				Config: testAccCsvserver_botpolicy_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCsvserver_botpolicy_bindingNotExist("citrixadc_csvserver_botpolicy_binding.tf_csvserver_botpolicy_binding", "tf_csvserver,tf_botpolicy"),
				),
			},
		},
	})
}

func testAccCheckCsvserver_botpolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No csvserver_botpolicy_binding id is set")
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
			return fmt.Errorf("Error parsing ID %q: %v", bindingId, err)
		}
		name := idMap["name"]
		policyname := idMap["policyname"]

		findParams := service.FindParams{
			ResourceType:             "csvserver_botpolicy_binding",
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
			return fmt.Errorf("csvserver_botpolicy_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckCsvserver_botpolicy_bindingNotExist(n string, id string) resource.TestCheckFunc {
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
			return fmt.Errorf("Error parsing ID %q: %v", id, err)
		}
		csvserverName := idMap["name"]
		policyName := idMap["policyname"]

		findParams := service.FindParams{
			ResourceType:             "csvserver_botpolicy_binding",
			ResourceName:             csvserverName,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the right policy name
		found := false
		for _, v := range dataArr {
			if v["policyname"].(string) == policyName {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("csvserver_botpolicy_binding %s was found, but is should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckCsvserver_botpolicy_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_csvserver_botpolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("csvserver_botpolicy_binding", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("csvserver_botpolicy_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccCsvserver_botpolicy_binding_import(t *testing.T) {
	const resAddr = "citrixadc_csvserver_botpolicy_binding.tf_csvserver_botpolicy_binding"

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
		CheckDestroy:             testAccCheckCsvserver_botpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccCsvserver_botpolicy_binding_basic},
			{Config: testAccCsvserver_botpolicy_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
			{Config: testAccCsvserver_botpolicy_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateIdFunc: legacyID, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
		},
	})
}

const testAccCsvserver_botpolicy_bindingDataSource_basic = `
	resource "citrixadc_csvserver_botpolicy_binding" "tf_csvserver_botpolicy_binding" {
        name = citrixadc_csvserver.tf_csvserver.name
        policyname = citrixadc_botpolicy.tf_botpolicy.name
        priority = 5
	}

	resource "citrixadc_csvserver" "tf_csvserver" {
		name = "tf_csvserver"
		ipv46 = "10.202.11.11"
		port = 8080
		servicetype = "HTTP"
	}

	resource citrixadc_botpolicy tf_botpolicy {
		name = "tf_botpolicy"  
		profilename = "BOT_BYPASS"
		rule  = "true"
	}

	data "citrixadc_csvserver_botpolicy_binding" "tf_csvserver_botpolicy_binding" {
		name = "tf_csvserver"
		policyname = "tf_botpolicy"
		depends_on = [citrixadc_csvserver_botpolicy_binding.tf_csvserver_botpolicy_binding]
	}
`

func TestAccCsvserver_botpolicy_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCsvserver_botpolicy_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_csvserver_botpolicy_binding.tf_csvserver_botpolicy_binding", "name", "tf_csvserver"),
					resource.TestCheckResourceAttr("data.citrixadc_csvserver_botpolicy_binding.tf_csvserver_botpolicy_binding", "policyname", "tf_botpolicy"),
					resource.TestCheckResourceAttr("data.citrixadc_csvserver_botpolicy_binding.tf_csvserver_botpolicy_binding", "priority", "5"),
				),
			},
		},
	})
}

const testAccCsvserver_botpolicy_binding_upgrade_basic = `
	resource "citrixadc_csvserver_botpolicy_binding" "tf_csvserver_botpolicy_binding" {
        name = citrixadc_csvserver.tf_csvserver.name
        policyname = citrixadc_botpolicy.tf_botpolicy.name
        priority = 5
	}

	resource "citrixadc_csvserver" "tf_csvserver" {
		name = "tf_csvserver"
		ipv46 = "10.202.11.11"
		port = 8080
		servicetype = "HTTP"
	}

	resource citrixadc_botpolicy tf_botpolicy {
		name = "tf_botpolicy"
		profilename = "BOT_BYPASS"
		rule  = "true"
	}
`

// TestAccCsvserver_botpolicy_binding_sdkv2StateUpgrade verifies that state written
// by the last SDK v2 release (legacy comma-separated ID) is correctly upgraded when
// the same config is subsequently managed by the current Framework provider. Step 1
// creates the binding with citrix/citrixadc 2.2.0 (writes the legacy id
// "tf_csvserver,tf_botpolicy"). Step 2 refreshes/plans/applies the same config through
// the Framework provider, exercising ParseIdString on the legacy id; because the
// Framework recomputes the id on Read (SetAttrFromGet re-derives data.Id), the id
// upgrades to the new "key:value" form.
func TestAccCsvserver_botpolicy_binding_sdkv2StateUpgrade(t *testing.T) {
	resourceAddr := "citrixadc_csvserver_botpolicy_binding.tf_csvserver_botpolicy_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckCsvserver_botpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			// Step 1: create with the last SDK v2 release -> state carries the legacy id.
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.0.0",
					},
				},
				Config: testAccCsvserver_botpolicy_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCsvserver_botpolicy_bindingExist(resourceAddr, nil),
					resource.TestCheckResourceAttr(resourceAddr, "id", "tf_csvserver,tf_botpolicy"),
				),
			},
			// Step 2: refresh/plan/apply the SAME config through the current Framework
			// provider. The legacy-id state is read via ParseIdString and the id is
			// recomputed to the new key:value format.
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccCsvserver_botpolicy_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCsvserver_botpolicy_bindingExist(resourceAddr, nil),
					resource.TestCheckResourceAttr(resourceAddr, "id", "name:tf_csvserver,policyname:tf_botpolicy"),
				),
			},
		},
	})
}

func TestAccCsvserver_botpolicy_binding_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_csvserver_botpolicy_binding.tf_csvserver_botpolicy_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCsvserver_botpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCsvserver_botpolicy_binding_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckCsvserver_botpolicy_bindingExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResourceWithArgsMap(service.Csvserver_botpolicy_binding.Type(), "tf_csvserver", map[string]string{"policyname": "tf_botpolicy", "priority": "5"}); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccCsvserver_botpolicy_binding_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckCsvserver_botpolicy_bindingExist(resAddr, nil)),
			},
		},
	})
}
