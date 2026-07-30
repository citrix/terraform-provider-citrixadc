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
	"testing"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testAccAppflowglobal_appflowpolicy_binding_basic = `

	resource "citrixadc_appflowglobal_appflowpolicy_binding" "tf_appflowglobal_appflowpolicy_binding" {
		policyname     = citrixadc_appflowpolicy.tf_appflowpolicy.name
		globalbindtype = "SYSTEM_GLOBAL"
		type           = "REQ_OVERRIDE"
		priority       = 55
	}

	
	resource "citrixadc_appflowpolicy" "tf_appflowpolicy" {
	  name   = "test_policy"
	  action = citrixadc_appflowaction.tf_appflowaction.name
	  rule   = "client.TCP.DSTPORT.EQ(22)"
	}
	resource "citrixadc_appflowaction" "tf_appflowaction" {
	  name            = "test_action"
	  collectors      = [citrixadc_appflowcollector.tf_appflowcollector.name]
	  securityinsight = "ENABLED"
	  botinsight      = "ENABLED"
	  videoanalytics  = "ENABLED"
	}
	resource "citrixadc_appflowcollector" "tf_appflowcollector" {
	  name      = "tf_collector"
	  ipaddress = "192.168.2.2"
	  port      = 80
	}
`

const testAccAppflowglobal_appflowpolicy_binding_basic_step2 = `

	
	
	resource "citrixadc_appflowpolicy" "tf_appflowpolicy" {
	  name   = "test_policy"
	  action = citrixadc_appflowaction.tf_appflowaction.name
	  rule   = "client.TCP.DSTPORT.EQ(22)"
	}
	resource "citrixadc_appflowaction" "tf_appflowaction" {
	  name            = "test_action"
	  collectors      = [citrixadc_appflowcollector.tf_appflowcollector.name]
	  securityinsight = "ENABLED"
	  botinsight      = "ENABLED"
	  videoanalytics  = "ENABLED"
	}
	resource "citrixadc_appflowcollector" "tf_appflowcollector" {
	  name      = "tf_collector"
	  ipaddress = "192.168.2.2"
	  port      = 80
	}
`

func TestAccAppflowglobal_appflowpolicy_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAppflowglobal_appflowpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAppflowglobal_appflowpolicy_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppflowglobal_appflowpolicy_bindingExist("citrixadc_appflowglobal_appflowpolicy_binding.tf_appflowglobal_appflowpolicy_binding", nil),
				),
			},
			{
				Config: testAccAppflowglobal_appflowpolicy_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppflowglobal_appflowpolicy_bindingNotExist("citrixadc_appflowglobal_appflowpolicy_binding.tf_appflowglobal_appflowpolicy_binding", "test3_policy", "REQ_DEFAULT"),
				),
			},
		},
	})
}

func testAccCheckAppflowglobal_appflowpolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No appflowglobal_appflowpolicy_binding id is set")
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

		// ID is the composite policyname:<v>,type:<v>; parse the policyname out of it.
		idMap, _, err := utils.ParseIdString(rs.Primary.ID, []string{"policyname"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %q: %v", rs.Primary.ID, err)
		}
		policyname := idMap["policyname"]
		typename := rs.Primary.Attributes["type"]

		findParams := service.FindParams{
			ResourceType:             "appflowglobal_appflowpolicy_binding",
			ArgsMap:                  map[string]string{"type": typename},
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the matching policyname
		found := false
		for _, v := range dataArr {
			if v["policyname"].(string) == policyname {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("appflowglobal_appflowpolicy_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckAppflowglobal_appflowpolicy_bindingNotExist(n string, id string, typename string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		policyname := id
		findParams := service.FindParams{
			ResourceType:             "appflowglobal_appflowpolicy_binding",
			ArgsMap:                  map[string]string{"type": typename},
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to hopefully not find the one with the matching policyname
		found := false
		for _, v := range dataArr {
			if v["policyname"].(string) == policyname {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("appflowglobal_appflowpolicy_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckAppflowglobal_appflowpolicy_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_appflowglobal_appflowpolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Appflowglobal_appflowpolicy_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("appflowglobal_appflowpolicy_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccAppflowglobal_appflowpolicy_bindingDataSource_basic = `

	resource "citrixadc_appflowglobal_appflowpolicy_binding" "tf_appflowglobal_appflowpolicy_binding" {
		policyname     = citrixadc_appflowpolicy.tf_appflowpolicy.name
		globalbindtype = "SYSTEM_GLOBAL"
		type           = "REQ_OVERRIDE"
		priority       = 55
	}

	
	resource "citrixadc_appflowpolicy" "tf_appflowpolicy" {
	  name   = "test_policy"
	  action = citrixadc_appflowaction.tf_appflowaction.name
	  rule   = "client.TCP.DSTPORT.EQ(22)"
	}
	resource "citrixadc_appflowaction" "tf_appflowaction" {
	  name            = "test_action"
	  collectors      = [citrixadc_appflowcollector.tf_appflowcollector.name]
	  securityinsight = "ENABLED"
	  botinsight      = "ENABLED"
	  videoanalytics  = "ENABLED"
	}
	resource "citrixadc_appflowcollector" "tf_appflowcollector" {
	  name      = "tf_collector"
	  ipaddress = "192.168.2.2"
	  port      = 80
	}

	data "citrixadc_appflowglobal_appflowpolicy_binding" "tf_appflowglobal_appflowpolicy_binding" {
		policyname = citrixadc_appflowglobal_appflowpolicy_binding.tf_appflowglobal_appflowpolicy_binding.policyname
		type       = citrixadc_appflowglobal_appflowpolicy_binding.tf_appflowglobal_appflowpolicy_binding.type
		depends_on = [citrixadc_appflowglobal_appflowpolicy_binding.tf_appflowglobal_appflowpolicy_binding]
	}
`

func TestAccAppflowglobal_appflowpolicy_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAppflowglobal_appflowpolicy_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_appflowglobal_appflowpolicy_binding.tf_appflowglobal_appflowpolicy_binding", "policyname", "test_policy"),
					resource.TestCheckResourceAttr("data.citrixadc_appflowglobal_appflowpolicy_binding.tf_appflowglobal_appflowpolicy_binding", "globalbindtype", "SYSTEM_GLOBAL"),
					resource.TestCheckResourceAttr("data.citrixadc_appflowglobal_appflowpolicy_binding.tf_appflowglobal_appflowpolicy_binding", "type", "REQ_OVERRIDE"),
					resource.TestCheckResourceAttr("data.citrixadc_appflowglobal_appflowpolicy_binding.tf_appflowglobal_appflowpolicy_binding", "priority", "55"),
				),
			},
		},
	})
}

// testAccAppflowglobal_appflowpolicy_binding_upgrade_basic reuses the _basic config
// (binding + all prerequisite resources). It is valid under BOTH the SDK v2 2.2.0
// schema and the current Framework schema because the migration restored the SDK v2
// attribute names.
const testAccAppflowglobal_appflowpolicy_binding_upgrade_basic = `

	resource "citrixadc_appflowglobal_appflowpolicy_binding" "tf_appflowglobal_appflowpolicy_binding" {
		policyname     = citrixadc_appflowpolicy.tf_appflowpolicy.name
		globalbindtype = "SYSTEM_GLOBAL"
		type           = "REQ_OVERRIDE"
		priority       = 55
	}

	resource "citrixadc_appflowpolicy" "tf_appflowpolicy" {
	  name   = "test_policy"
	  action = citrixadc_appflowaction.tf_appflowaction.name
	  rule   = "client.TCP.DSTPORT.EQ(22)"
	}
	resource "citrixadc_appflowaction" "tf_appflowaction" {
	  name            = "test_action"
	  collectors      = [citrixadc_appflowcollector.tf_appflowcollector.name]
	  securityinsight = "ENABLED"
	  botinsight      = "ENABLED"
	  videoanalytics  = "ENABLED"
	}
	resource "citrixadc_appflowcollector" "tf_appflowcollector" {
	  name      = "tf_collector"
	  ipaddress = "192.168.2.2"
	  port      = 80
	}
`

// TestAccAppflowglobal_appflowpolicy_binding_sdkv2StateUpgrade verifies that state
// written by the last SDK v2 release is correctly upgraded when the same config is
// subsequently managed by the current Framework provider. Step 1 creates the binding
// with citrix/citrixadc 2.2.0 (writes the legacy id "test_policy" — the SDK v2
// d.SetId(policyname)). Step 2 refreshes/plans/applies the same config through the
// Framework provider, exercising ParseIdString on the legacy id; the Framework
// recomputes the id on Read (SetAttrFromGet) into the composite key:value form. A single
// appflowpolicy can be bound at multiple bind points (type), so the id is upgraded from the
// legacy plain "test_policy" to "policyname:test_policy,type:REQ_OVERRIDE".
func TestAccAppflowglobal_appflowpolicy_binding_sdkv2StateUpgrade(t *testing.T) {
	resourceAddr := "citrixadc_appflowglobal_appflowpolicy_binding.tf_appflowglobal_appflowpolicy_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckAppflowglobal_appflowpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			// Step 1: create with the last SDK v2 release -> state carries the legacy id.
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.2.0",
					},
				},
				Config: testAccAppflowglobal_appflowpolicy_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppflowglobal_appflowpolicy_bindingExist(resourceAddr, nil),
					resource.TestCheckResourceAttr(resourceAddr, "id", "test_policy"),
				),
			},
			// Step 2: refresh/plan/apply the SAME config through the current Framework
			// provider. The legacy-id state is read via ParseIdString and the id is
			// recomputed on Read into the composite form (policyname:...,type:...).
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccAppflowglobal_appflowpolicy_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppflowglobal_appflowpolicy_bindingExist(resourceAddr, nil),
					resource.TestCheckResourceAttr(resourceAddr, "id", "policyname:test_policy,type:REQ_OVERRIDE"),
				),
			},
		},
	})
}

func TestAccAppflowglobal_appflowpolicy_binding_import(t *testing.T) {
	const resAddr = "citrixadc_appflowglobal_appflowpolicy_binding.tf_appflowglobal_appflowpolicy_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAppflowglobal_appflowpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccAppflowglobal_appflowpolicy_binding_basic},
			{Config: testAccAppflowglobal_appflowpolicy_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
		},
	})
}

func TestAccAppflowglobal_appflowpolicy_binding_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_appflowglobal_appflowpolicy_binding.tf_appflowglobal_appflowpolicy_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAppflowglobal_appflowpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAppflowglobal_appflowpolicy_binding_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAppflowglobal_appflowpolicy_bindingExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResourceWithArgsMap(service.Appflowglobal_appflowpolicy_binding.Type(), "", map[string]string{"policyname": "test_policy", "type": "REQ_OVERRIDE", "priority": "55"}); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccAppflowglobal_appflowpolicy_binding_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAppflowglobal_appflowpolicy_bindingExist(resAddr, nil)),
			},
		},
	})
}
