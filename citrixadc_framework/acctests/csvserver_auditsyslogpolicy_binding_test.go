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

const testAccCsvserver_auditsyslogpolicy_binding_basic = `
	resource "citrixadc_csvserver_auditsyslogpolicy_binding" "tf_csvserver_auditsyslogpolicy_binding" {
        name = citrixadc_csvserver.tf_csvserver.name
        policyname = citrixadc_auditsyslogpolicy.tf_auditsyslogpolicy.name
        priority = 5
	}

	resource "citrixadc_csvserver" "tf_csvserver" {
		name = "tf_csvserver"
		ipv46 = "10.202.11.11"
		port = 8080
		servicetype = "HTTP"
	}

	resource "citrixadc_auditsyslogpolicy" "tf_auditsyslogpolicy" {
		name = "tf_auditsyslogpolicy"
		rule = "ns_true"
		action = citrixadc_auditsyslogaction.tf_syslogaction.name
	}

	resource "citrixadc_auditsyslogaction" "tf_syslogaction" {
		name = "tf_syslogaction"
		serverip = "10.78.60.53"
		serverport = 514
		loglevel = [
			"ERROR",
			"NOTICE",
		]
	}
`

const testAccCsvserver_auditsyslogpolicy_binding_basic_step2 = `
	resource "citrixadc_csvserver" "tf_csvserver" {
		name = "tf_csvserver"
		ipv46 = "10.202.11.11"
		port = 8080
		servicetype = "HTTP"
	}

	resource "citrixadc_auditsyslogpolicy" "tf_auditsyslogpolicy" {
		name = "tf_auditsyslogpolicy"
		rule = "ns_true"
		action = citrixadc_auditsyslogaction.tf_syslogaction.name
	}

	resource "citrixadc_auditsyslogaction" "tf_syslogaction" {
		name = "tf_syslogaction"
		serverip = "10.78.60.53"
		serverport = 514
		loglevel = [
			"ERROR",
			"NOTICE",
		]
	}
`

func TestAccCsvserver_auditsyslogpolicy_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCsvserver_auditsyslogpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCsvserver_auditsyslogpolicy_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCsvserver_auditsyslogpolicy_bindingExist("citrixadc_csvserver_auditsyslogpolicy_binding.tf_csvserver_auditsyslogpolicy_binding", nil),
				),
			},
			{
				Config: testAccCsvserver_auditsyslogpolicy_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCsvserver_auditsyslogpolicy_bindingNotExist("citrixadc_csvserver_auditsyslogpolicy_binding.tf_csvserver_auditsyslogpolicy_binding", "tf_csvserver,tf_auditsyslogpolicy"),
				),
			},
		},
	})
}

func testAccCheckCsvserver_auditsyslogpolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No csvserver_auditsyslogpolicy_binding id is set")
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
			ResourceType:             "csvserver_auditsyslogpolicy_binding",
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
			return fmt.Errorf("csvserver_auditsyslogpolicy_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckCsvserver_auditsyslogpolicy_bindingNotExist(n string, id string) resource.TestCheckFunc {
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
		csvserverName := idMap["name"]
		policyName := idMap["policyname"]

		findParams := service.FindParams{
			ResourceType:             "csvserver_auditsyslogpolicy_binding",
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
			return fmt.Errorf("csvserver_auditsyslogpolicy_binding %s was found, but is should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckCsvserver_auditsyslogpolicy_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_csvserver_auditsyslogpolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Csvserver_auditsyslogpolicy_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("csvserver_auditsyslogpolicy_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccCsvserver_auditsyslogpolicy_bindingDataSource_basic = `
	resource "citrixadc_csvserver_auditsyslogpolicy_binding" "tf_csvserver_auditsyslogpolicy_binding" {
        name = citrixadc_csvserver.tf_csvserver.name
        policyname = citrixadc_auditsyslogpolicy.tf_auditsyslogpolicy.name
        priority = 5
	}

	resource "citrixadc_csvserver" "tf_csvserver" {
		name = "tf_csvserver"
		ipv46 = "10.202.11.11"
		port = 8080
		servicetype = "HTTP"
	}

	resource "citrixadc_auditsyslogpolicy" "tf_auditsyslogpolicy" {
		name = "tf_auditsyslogpolicy"
		rule = "ns_true"
		action = citrixadc_auditsyslogaction.tf_syslogaction.name
	}

	resource "citrixadc_auditsyslogaction" "tf_syslogaction" {
		name = "tf_syslogaction"
		serverip = "10.78.60.53"
		serverport = 514
		loglevel = [
			"ERROR",
			"NOTICE",
		]
	}

	data "citrixadc_csvserver_auditsyslogpolicy_binding" "tf_csvserver_auditsyslogpolicy_binding" {
		name = "tf_csvserver"
		policyname = "tf_auditsyslogpolicy"
		depends_on = [citrixadc_csvserver_auditsyslogpolicy_binding.tf_csvserver_auditsyslogpolicy_binding]
	}
`

func TestAccCsvserver_auditsyslogpolicy_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCsvserver_auditsyslogpolicy_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_csvserver_auditsyslogpolicy_binding.tf_csvserver_auditsyslogpolicy_binding", "name", "tf_csvserver"),
					resource.TestCheckResourceAttr("data.citrixadc_csvserver_auditsyslogpolicy_binding.tf_csvserver_auditsyslogpolicy_binding", "policyname", "tf_auditsyslogpolicy"),
					resource.TestCheckResourceAttr("data.citrixadc_csvserver_auditsyslogpolicy_binding.tf_csvserver_auditsyslogpolicy_binding", "priority", "5"),
				),
			},
		},
	})
}

const testAccCsvserver_auditsyslogpolicy_binding_upgrade_basic = `
	resource "citrixadc_csvserver_auditsyslogpolicy_binding" "tf_csvserver_auditsyslogpolicy_binding" {
        name = citrixadc_csvserver.tf_csvserver.name
        policyname = citrixadc_auditsyslogpolicy.tf_auditsyslogpolicy.name
        priority = 5
	}

	resource "citrixadc_csvserver" "tf_csvserver" {
		name = "tf_csvserver"
		ipv46 = "10.202.11.11"
		port = 8080
		servicetype = "HTTP"
	}

	resource "citrixadc_auditsyslogpolicy" "tf_auditsyslogpolicy" {
		name = "tf_auditsyslogpolicy"
		rule = "ns_true"
		action = citrixadc_auditsyslogaction.tf_syslogaction.name
	}

	resource "citrixadc_auditsyslogaction" "tf_syslogaction" {
		name = "tf_syslogaction"
		serverip = "10.78.60.53"
		serverport = 514
		loglevel = [
			"ERROR",
			"NOTICE",
		]
	}
`

// TestAccCsvserver_auditsyslogpolicy_binding_sdkv2StateUpgrade verifies that a binding
// created with the last SDK v2 release (2.2.0), whose state carries the legacy
// comma-joined id ("name,policyname"), is transparently upgraded when refreshed/applied
// through the current Framework provider. The Framework Read re-derives the canonical
// new-format id ("name:<v>,policyname:<v>") via SetAttrFromGet.
func TestAccCsvserver_auditsyslogpolicy_binding_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckCsvserver_auditsyslogpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			// Step 1: create with the last SDK v2 release -> legacy id in state.
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.2.0",
					},
				},
				Config: testAccCsvserver_auditsyslogpolicy_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCsvserver_auditsyslogpolicy_bindingExist("citrixadc_csvserver_auditsyslogpolicy_binding.tf_csvserver_auditsyslogpolicy_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_csvserver_auditsyslogpolicy_binding.tf_csvserver_auditsyslogpolicy_binding", "id", "tf_csvserver,tf_auditsyslogpolicy"),
				),
			},
			// Step 2: same config through the current Framework provider -> id upgraded to new format.
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccCsvserver_auditsyslogpolicy_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCsvserver_auditsyslogpolicy_bindingExist("citrixadc_csvserver_auditsyslogpolicy_binding.tf_csvserver_auditsyslogpolicy_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_csvserver_auditsyslogpolicy_binding.tf_csvserver_auditsyslogpolicy_binding", "id", "name:tf_csvserver,policyname:tf_auditsyslogpolicy"),
				),
			},
		},
	})
}

func TestAccCsvserver_auditsyslogpolicy_binding_import(t *testing.T) {
	const resAddr = "citrixadc_csvserver_auditsyslogpolicy_binding.tf_csvserver_auditsyslogpolicy_binding"

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
		CheckDestroy:             testAccCheckCsvserver_auditsyslogpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccCsvserver_auditsyslogpolicy_binding_basic},
			{Config: testAccCsvserver_auditsyslogpolicy_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
			{Config: testAccCsvserver_auditsyslogpolicy_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateIdFunc: legacyID, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
		},
	})
}

func TestAccCsvserver_auditsyslogpolicy_binding_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_csvserver_auditsyslogpolicy_binding.tf_csvserver_auditsyslogpolicy_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCsvserver_auditsyslogpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCsvserver_auditsyslogpolicy_binding_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckCsvserver_auditsyslogpolicy_bindingExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResourceWithArgsMap(service.Csvserver_auditsyslogpolicy_binding.Type(), "tf_csvserver", map[string]string{"policyname": "tf_auditsyslogpolicy", "priority": "5"}); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccCsvserver_auditsyslogpolicy_binding_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckCsvserver_auditsyslogpolicy_bindingExist(resAddr, nil)),
			},
		},
	})
}
