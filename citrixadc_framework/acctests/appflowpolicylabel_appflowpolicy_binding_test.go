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

const testAccAppflowpolicylabel_appflowpolicy_binding_basic = `

	resource "citrixadc_appflowpolicylabel_appflowpolicy_binding" "tf_appflowpolicylabel_appflowpolicy_binding" {
		labelname  = citrixadc_appflowpolicylabel.tf_appflowpolicylabel.labelname
		policyname = citrixadc_appflowpolicy.tf_appflowpolicy.name
		priority   = 30
	}

	resource "citrixadc_appflowpolicylabel" "tf_appflowpolicylabel" {
	  labelname       = "tf_policylabel"
	  policylabeltype = "OTHERTCP"
	}
	
	resource "citrixadc_appflowpolicy" "tf_appflowpolicy" {
	  name      = "test_policy"
	  action    = citrixadc_appflowaction.tf_appflowaction.name
	  rule      = "client.TCP.DSTPORT.EQ(22)"
	}
	resource "citrixadc_appflowaction" "tf_appflowaction" {
	  name = "test_action"
	  collectors     = [citrixadc_appflowcollector.tf_appflowcollector.name]
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

const testAccAppflowpolicylabel_appflowpolicy_binding_basic_step2 = `
	
	resource "citrixadc_appflowpolicylabel" "tf_appflowpolicylabel" {
	  labelname       = "tf_policylabel"
	  policylabeltype = "OTHERTCP"
	}
	
	resource "citrixadc_appflowpolicy" "tf_appflowpolicy" {
	  name      = "test_policy"
	  action    = citrixadc_appflowaction.tf_appflowaction.name
	  rule      = "client.TCP.DSTPORT.EQ(22)"
	}
	resource "citrixadc_appflowaction" "tf_appflowaction" {
	  name = "test_action"
	  collectors     = [citrixadc_appflowcollector.tf_appflowcollector.name]
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

func TestAccAppflowpolicylabel_appflowpolicy_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAppflowpolicylabel_appflowpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAppflowpolicylabel_appflowpolicy_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppflowpolicylabel_appflowpolicy_bindingExist("citrixadc_appflowpolicylabel_appflowpolicy_binding.tf_appflowpolicylabel_appflowpolicy_binding", nil),
				),
			},
			{
				Config: testAccAppflowpolicylabel_appflowpolicy_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppflowpolicylabel_appflowpolicy_bindingNotExist("citrixadc_appflowpolicylabel_appflowpolicy_binding.tf_appflowpolicylabel_appflowpolicy_binding", "tf_policylabel,test_policy"),
				),
			},
		},
	})
}

func testAccCheckAppflowpolicylabel_appflowpolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No appflowpolicylabel_appflowpolicy_binding id is set")
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

		idMap, _, err := utils.ParseIdString(bindingId, []string{"labelname", "policyname"}, nil)
		if err != nil {
			return err
		}
		labelname := idMap["labelname"]
		policyname := idMap["policyname"]

		findParams := service.FindParams{
			ResourceType:             "appflowpolicylabel_appflowpolicy_binding",
			ResourceName:             labelname,
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
			return fmt.Errorf("appflowpolicylabel_appflowpolicy_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckAppflowpolicylabel_appflowpolicy_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		idMap, _, err := utils.ParseIdString(id, []string{"labelname", "policyname"}, nil)
		if err != nil {
			return err
		}
		labelname := idMap["labelname"]
		policyname := idMap["policyname"]

		findParams := service.FindParams{
			ResourceType:             "appflowpolicylabel_appflowpolicy_binding",
			ResourceName:             labelname,
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
			return fmt.Errorf("appflowpolicylabel_appflowpolicy_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckAppflowpolicylabel_appflowpolicy_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_appflowpolicylabel_appflowpolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Appflowpolicylabel_appflowpolicy_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("appflowpolicylabel_appflowpolicy_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccAppflowpolicylabel_appflowpolicy_bindingDataSource_basic = `

	resource "citrixadc_appflowpolicylabel_appflowpolicy_binding" "tf_appflowpolicylabel_appflowpolicy_binding" {
		labelname  = citrixadc_appflowpolicylabel.tf_appflowpolicylabel.labelname
		policyname = citrixadc_appflowpolicy.tf_appflowpolicy.name
		priority   = 30
	}

	resource "citrixadc_appflowpolicylabel" "tf_appflowpolicylabel" {
	  labelname       = "tf_policylabel"
	  policylabeltype = "OTHERTCP"
	}
	
	resource "citrixadc_appflowpolicy" "tf_appflowpolicy" {
	  name      = "test_policy"
	  action    = citrixadc_appflowaction.tf_appflowaction.name
	  rule      = "client.TCP.DSTPORT.EQ(22)"
	}
	resource "citrixadc_appflowaction" "tf_appflowaction" {
	  name = "test_action"
	  collectors     = [citrixadc_appflowcollector.tf_appflowcollector.name]
	  securityinsight = "ENABLED"
	  botinsight      = "ENABLED"
	  videoanalytics  = "ENABLED"
	}
	resource "citrixadc_appflowcollector" "tf_appflowcollector" {
	  name      = "tf_collector"
	  ipaddress = "192.168.2.2"
	  port      = 80
	}

	data "citrixadc_appflowpolicylabel_appflowpolicy_binding" "tf_appflowpolicylabel_appflowpolicy_binding" {
		labelname  = citrixadc_appflowpolicylabel_appflowpolicy_binding.tf_appflowpolicylabel_appflowpolicy_binding.labelname
		policyname = citrixadc_appflowpolicylabel_appflowpolicy_binding.tf_appflowpolicylabel_appflowpolicy_binding.policyname
		depends_on = [citrixadc_appflowpolicylabel_appflowpolicy_binding.tf_appflowpolicylabel_appflowpolicy_binding]
	}
`

func TestAccAppflowpolicylabel_appflowpolicy_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAppflowpolicylabel_appflowpolicy_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_appflowpolicylabel_appflowpolicy_binding.tf_appflowpolicylabel_appflowpolicy_binding", "labelname", "tf_policylabel"),
					resource.TestCheckResourceAttr("data.citrixadc_appflowpolicylabel_appflowpolicy_binding.tf_appflowpolicylabel_appflowpolicy_binding", "policyname", "test_policy"),
					resource.TestCheckResourceAttr("data.citrixadc_appflowpolicylabel_appflowpolicy_binding.tf_appflowpolicylabel_appflowpolicy_binding", "priority", "30"),
				),
			},
		},
	})
}

// testAccAppflowpolicylabel_appflowpolicy_binding_upgrade_basic reuses the _basic config
// (the binding + all prerequisite resources). It is valid under BOTH the SDK v2 2.2.0
// schema and the current Framework schema because the migration restored the SDK v2
// attribute names.
const testAccAppflowpolicylabel_appflowpolicy_binding_upgrade_basic = `

	resource "citrixadc_appflowpolicylabel_appflowpolicy_binding" "tf_appflowpolicylabel_appflowpolicy_binding" {
		labelname  = citrixadc_appflowpolicylabel.tf_appflowpolicylabel.labelname
		policyname = citrixadc_appflowpolicy.tf_appflowpolicy.name
		priority   = 30
	}

	resource "citrixadc_appflowpolicylabel" "tf_appflowpolicylabel" {
	  labelname       = "tf_policylabel"
	  policylabeltype = "OTHERTCP"
	}

	resource "citrixadc_appflowpolicy" "tf_appflowpolicy" {
	  name      = "test_policy"
	  action    = citrixadc_appflowaction.tf_appflowaction.name
	  rule      = "client.TCP.DSTPORT.EQ(22)"
	}
	resource "citrixadc_appflowaction" "tf_appflowaction" {
	  name = "test_action"
	  collectors     = [citrixadc_appflowcollector.tf_appflowcollector.name]
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

// TestAccAppflowpolicylabel_appflowpolicy_binding_sdkv2StateUpgrade verifies that state
// written by the last SDK v2 release is correctly upgraded when the same config is
// subsequently managed by the current Framework provider.
//
// Step 1 creates the binding with citrix/citrixadc 2.2.0, which writes the legacy
// comma-joined id (SDK v2 d.SetId("labelname,policyname") => "tf_policylabel,test_policy").
// Step 2 refreshes/plans/applies the SAME config through the current Framework provider,
// exercising ParseIdString on the legacy id. The Framework recomputes the id on Read
// (appflowpolicylabel_appflowpolicy_bindingSetAttrFromGet), so the id is upgraded to the
// canonical new key:value format "labelname:tf_policylabel,policyname:test_policy".
func TestAccAppflowpolicylabel_appflowpolicy_binding_sdkv2StateUpgrade(t *testing.T) {
	resourceAddr := "citrixadc_appflowpolicylabel_appflowpolicy_binding.tf_appflowpolicylabel_appflowpolicy_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckAppflowpolicylabel_appflowpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			// Step 1: create with the last SDK v2 release -> state carries the legacy id.
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.2.0",
					},
				},
				Config: testAccAppflowpolicylabel_appflowpolicy_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppflowpolicylabel_appflowpolicy_bindingExist(resourceAddr, nil),
					resource.TestCheckResourceAttr(resourceAddr, "id", "tf_policylabel,test_policy"),
				),
			},
			// Step 2: refresh/plan/apply the SAME config through the current Framework
			// provider. The legacy-id state is read via ParseIdString and the id is
			// recomputed on Read to the canonical new key:value format.
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccAppflowpolicylabel_appflowpolicy_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppflowpolicylabel_appflowpolicy_bindingExist(resourceAddr, nil),
					resource.TestCheckResourceAttr(resourceAddr, "id", "labelname:tf_policylabel,policyname:test_policy"),
				),
			},
		},
	})
}

func TestAccAppflowpolicylabel_appflowpolicy_binding_import(t *testing.T) {
	const resAddr = "citrixadc_appflowpolicylabel_appflowpolicy_binding.tf_appflowpolicylabel_appflowpolicy_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAppflowpolicylabel_appflowpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccAppflowpolicylabel_appflowpolicy_binding_basic},
			{Config: testAccAppflowpolicylabel_appflowpolicy_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
		},
	})
}
