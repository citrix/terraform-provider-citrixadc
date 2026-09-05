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

const testAccCsvserver_cachepolicy_binding_basic = `
	
	resource "citrixadc_cachepolicy" "tf_cachepolicy" {
		policyname  = "tf_cachepolicy"
		rule        = "true"
		action      = "CACHE"
	}
	resource "citrixadc_csvserver_cachepolicy_binding" "tf_csvserver_cachepolicy_binding" {
        name 		= citrixadc_csvserver.tf_csvserver.name
        policyname 	= citrixadc_cachepolicy.tf_cachepolicy.policyname
        priority 	= 5       
		bindpoint 	= "REQUEST" 
	}

	resource "citrixadc_csvserver" "tf_csvserver" {
		name 		= "tf_csvserver"
		ipv46 		= "10.202.11.11"
		port 		= 8080
		servicetype = "HTTP"
	}
`

const testAccCsvserver_cachepolicy_binding_basic_step2 = `
	
	resource "citrixadc_cachepolicy" "tf_cachepolicy" {
		policyname  = "tf_cachepolicy"
		rule        = "true"
		action      = "CACHE"
	}
	resource "citrixadc_csvserver" "tf_csvserver" {
		name 		= "tf_csvserver"
		ipv46 		= "10.202.11.11"
		port 		= 8080
		servicetype = "HTTP"
	}
`

func TestAccCsvserver_cachepolicy_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCsvserver_cachepolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCsvserver_cachepolicy_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCsvserver_cachepolicy_bindingExist("citrixadc_csvserver_cachepolicy_binding.tf_csvserver_cachepolicy_binding", nil),
				),
			},
			{
				Config: testAccCsvserver_cachepolicy_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCsvserver_cachepolicy_bindingNotExist("citrixadc_csvserver_cachepolicy_binding.tf_csvserver_cachepolicy_binding", "tf_csvserver,tf_cachepolicy"),
				),
			},
		},
	})
}

func TestAccCsvserver_cachepolicy_binding_import(t *testing.T) {
	const resAddr = "citrixadc_csvserver_cachepolicy_binding.tf_csvserver_cachepolicy_binding"

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
		CheckDestroy:             testAccCheckCsvserver_cachepolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccCsvserver_cachepolicy_binding_basic},
			{Config: testAccCsvserver_cachepolicy_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
			{Config: testAccCsvserver_cachepolicy_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateIdFunc: legacyID, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
		},
	})
}

func testAccCheckCsvserver_cachepolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No csvserver_cachepolicy_binding id is set")
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
			ResourceType:             "csvserver_cachepolicy_binding",
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
			return fmt.Errorf("csvserver_cachepolicy_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckCsvserver_cachepolicy_bindingNotExist(n string, id string) resource.TestCheckFunc {
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

		csvserverName := idSlice[0]
		policyName := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "csvserver_cachepolicy_binding",
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
			return fmt.Errorf("csvserver_cachepolicy_binding %s was found, but is should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckCsvserver_cachepolicy_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_csvserver_cachepolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Csvserver_cachepolicy_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("csvserver_cachepolicy_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccCsvserver_cachepolicy_bindingDataSource_basic = `
	
	resource "citrixadc_cachepolicy" "tf_cachepolicy" {
		policyname  = "tf_cachepolicy"
		rule        = "true"
		action      = "CACHE"
	}
	resource "citrixadc_csvserver_cachepolicy_binding" "tf_csvserver_cachepolicy_binding" {
        name 		= citrixadc_csvserver.tf_csvserver.name
        policyname 	= citrixadc_cachepolicy.tf_cachepolicy.policyname
        priority 	= 5       
		bindpoint 	= "REQUEST" 
	}

	resource "citrixadc_csvserver" "tf_csvserver" {
		name 		= "tf_csvserver"
		ipv46 		= "10.202.11.11"
		port 		= 8080
		servicetype = "HTTP"
	}

	data "citrixadc_csvserver_cachepolicy_binding" "tf_csvserver_cachepolicy_binding" {
		name       = citrixadc_csvserver_cachepolicy_binding.tf_csvserver_cachepolicy_binding.name
		policyname = citrixadc_csvserver_cachepolicy_binding.tf_csvserver_cachepolicy_binding.policyname
		bindpoint  = citrixadc_csvserver_cachepolicy_binding.tf_csvserver_cachepolicy_binding.bindpoint
	}
`

func TestAccCsvserver_cachepolicy_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCsvserver_cachepolicy_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_csvserver_cachepolicy_binding.tf_csvserver_cachepolicy_binding", "name", "tf_csvserver"),
					resource.TestCheckResourceAttr("data.citrixadc_csvserver_cachepolicy_binding.tf_csvserver_cachepolicy_binding", "policyname", "tf_cachepolicy"),
					resource.TestCheckResourceAttr("data.citrixadc_csvserver_cachepolicy_binding.tf_csvserver_cachepolicy_binding", "priority", "5"),
					resource.TestCheckResourceAttr("data.citrixadc_csvserver_cachepolicy_binding.tf_csvserver_cachepolicy_binding", "bindpoint", "REQUEST"),
				),
			},
		},
	})
}

const testAccCsvserver_cachepolicy_binding_upgrade_basic = `

	resource "citrixadc_cachepolicy" "tf_cachepolicy" {
		policyname  = "tf_cachepolicy"
		rule        = "true"
		action      = "CACHE"
	}
	resource "citrixadc_csvserver" "tf_csvserver" {
		name 		= "tf_csvserver"
		ipv46 		= "10.202.11.11"
		port 		= 8080
		servicetype = "HTTP"
	}
	resource "citrixadc_csvserver_cachepolicy_binding" "tf_csvserver_cachepolicy_binding" {
		name 		= citrixadc_csvserver.tf_csvserver.name
		policyname 	= citrixadc_cachepolicy.tf_cachepolicy.policyname
		priority 	= 5
		bindpoint 	= "REQUEST"
	}
`

func TestAccCsvserver_cachepolicy_binding_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckCsvserver_cachepolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				// Step 1: create the binding with the last SDK v2 release (2.2.0),
				// which writes state using the legacy comma-joined id.
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.0.0",
					},
				},
				Config: testAccCsvserver_cachepolicy_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCsvserver_cachepolicy_bindingExist("citrixadc_csvserver_cachepolicy_binding.tf_csvserver_cachepolicy_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_csvserver_cachepolicy_binding.tf_csvserver_cachepolicy_binding", "id", "tf_csvserver,tf_cachepolicy"),
				),
			},
			{
				// Step 2: refresh/plan the legacy-id state through the current
				// framework provider. Read exercises ParseIdString on the legacy id
				// and SetAttrFromGet recomputes the id into the new key:value form.
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccCsvserver_cachepolicy_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCsvserver_cachepolicy_bindingExist("citrixadc_csvserver_cachepolicy_binding.tf_csvserver_cachepolicy_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_csvserver_cachepolicy_binding.tf_csvserver_cachepolicy_binding", "id", "bindpoint:REQUEST,name:tf_csvserver,policyname:tf_cachepolicy"),
				),
			},
		},
	})
}

func TestAccCsvserver_cachepolicy_binding_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_csvserver_cachepolicy_binding.tf_csvserver_cachepolicy_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCsvserver_cachepolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCsvserver_cachepolicy_binding_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckCsvserver_cachepolicy_bindingExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResourceWithArgsMap(service.Csvserver_cachepolicy_binding.Type(), "tf_csvserver", map[string]string{"policyname": "tf_cachepolicy", "priority": "5"}); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccCsvserver_cachepolicy_binding_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckCsvserver_cachepolicy_bindingExist(resAddr, nil)),
			},
		},
	})
}
