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
	"strings"
	"testing"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testAccCrvserver_appflowpolicy_binding_basic = `

	resource "citrixadc_crvserver" "crvserver" {
		name = "my_vserver"
		servicetype = "HTTP"
		arp = "OFF"
	}
	resource "citrixadc_appflowpolicy" "tf_appflowpolicy" {
		name      = "tf_appflowpolicy"
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

	resource "citrixadc_crvserver_appflowpolicy_binding" "crvserver_appflowpolicy_binding" {
		name = citrixadc_crvserver.crvserver.name
		policyname = citrixadc_appflowpolicy.tf_appflowpolicy.name
		gotopriorityexpression = "END"
		labelname = citrixadc_crvserver.crvserver.name
		invoke = true
		labeltype = "reqvserver"
		priority = 1
	}
`

const testAccCrvserver_appflowpolicy_binding_basic_step2 = `

	resource "citrixadc_crvserver" "crvserver" {
		name = "my_vserver"
		servicetype = "HTTP"
		arp = "OFF"
	}
	resource "citrixadc_appflowpolicy" "tf_appflowpolicy" {
		name      = "tf_appflowpolicy"
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

func TestAccCrvserver_appflowpolicy_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCrvserver_appflowpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCrvserver_appflowpolicy_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCrvserver_appflowpolicy_bindingExist("citrixadc_crvserver_appflowpolicy_binding.crvserver_appflowpolicy_binding", nil),
				),
			},
			{
				Config: testAccCrvserver_appflowpolicy_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCrvserver_appflowpolicy_bindingNotExist("citrixadc_crvserver_appflowpolicy_binding.crvserver_appflowpolicy_binding", "my_vserver,tf_appflowpolicy"),
				),
			},
		},
	})
}

func TestAccCrvserver_appflowpolicy_binding_import(t *testing.T) {
	const resAddr = "citrixadc_crvserver_appflowpolicy_binding.crvserver_appflowpolicy_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCrvserver_appflowpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccCrvserver_appflowpolicy_binding_basic},
			{Config: testAccCrvserver_appflowpolicy_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
		},
	})
}

func testAccCheckCrvserver_appflowpolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No crvserver_appflowpolicy_binding id is set")
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
			return fmt.Errorf("Error parsing ID %v: %v", bindingId, err)
		}
		name := idMap["name"]
		policyname := idMap["policyname"]

		findParams := service.FindParams{
			ResourceType:             "crvserver_appflowpolicy_binding",
			ResourceName:             name,
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
			return fmt.Errorf("crvserver_appflowpolicy_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckCrvserver_appflowpolicy_bindingNotExist(n string, id string) resource.TestCheckFunc {
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
			return fmt.Errorf("Error parsing ID %v: %v", id, err)
		}
		name := idMap["name"]
		policyname := idMap["policyname"]

		findParams := service.FindParams{
			ResourceType:             "crvserver_appflowpolicy_binding",
			ResourceName:             name,
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
			return fmt.Errorf("crvserver_appflowpolicy_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckCrvserver_appflowpolicy_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_crvserver_appflowpolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Crvserver_appflowpolicy_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("crvserver_appflowpolicy_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccCrvserver_appflowpolicy_bindingDataSource_basic = `

	resource "citrixadc_crvserver" "crvserver" {
		name = "my_vserver_ds"
		servicetype = "HTTP"
		arp = "OFF"
	}
	resource "citrixadc_appflowpolicy" "tf_appflowpolicy" {
		name      = "tf_appflowpolicy_ds"
		action    = citrixadc_appflowaction.tf_appflowaction.name
		rule      = "client.TCP.DSTPORT.EQ(22)"
	}
	resource "citrixadc_appflowaction" "tf_appflowaction" {
		name = "test_action_ds"
		collectors     = [citrixadc_appflowcollector.tf_appflowcollector.name]
		securityinsight = "ENABLED"
		botinsight      = "ENABLED"
		videoanalytics  = "ENABLED"
	}
	resource "citrixadc_appflowcollector" "tf_appflowcollector" {
		name      = "tf_collector_ds"
		ipaddress = "192.168.2.2"
		port      = 80
	}

	resource "citrixadc_crvserver_appflowpolicy_binding" "crvserver_appflowpolicy_binding" {
		name = citrixadc_crvserver.crvserver.name
		policyname = citrixadc_appflowpolicy.tf_appflowpolicy.name
		priority = 1
	}

	data "citrixadc_crvserver_appflowpolicy_binding" "crvserver_appflowpolicy_binding" {
		name       = citrixadc_crvserver_appflowpolicy_binding.crvserver_appflowpolicy_binding.name
		policyname = citrixadc_crvserver_appflowpolicy_binding.crvserver_appflowpolicy_binding.policyname
		depends_on = [citrixadc_crvserver_appflowpolicy_binding.crvserver_appflowpolicy_binding]
	}
`

func TestAcccrvserver_appflowpolicy_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccCrvserver_appflowpolicy_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_crvserver_appflowpolicy_binding.crvserver_appflowpolicy_binding", "name", "my_vserver_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_crvserver_appflowpolicy_binding.crvserver_appflowpolicy_binding", "policyname", "tf_appflowpolicy_ds"),
				),
			},
		},
	})
}

// testAcccrvserver_appflowpolicy_binding_upgrade_basic mirrors the _basic config
// (same resource labels + prerequisites) and must be valid under BOTH the SDK v2
// 2.2.0 schema and the current Framework schema.
const testAcccrvserver_appflowpolicy_binding_upgrade_basic = `

	resource "citrixadc_crvserver" "crvserver" {
		name = "my_vserver"
		servicetype = "HTTP"
		arp = "OFF"
	}
	resource "citrixadc_appflowpolicy" "tf_appflowpolicy" {
		name      = "tf_appflowpolicy"
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

	resource "citrixadc_crvserver_appflowpolicy_binding" "crvserver_appflowpolicy_binding" {
		name = citrixadc_crvserver.crvserver.name
		policyname = citrixadc_appflowpolicy.tf_appflowpolicy.name
		gotopriorityexpression = "END"
		labelname = citrixadc_crvserver.crvserver.name
		invoke = true
		labeltype = "reqvserver"
		priority = 1
	}
`

// TestAccCrvserver_appflowpolicy_binding_sdkv2StateUpgrade verifies that state written
// by the last SDK v2 release (legacy comma-separated id) is upgraded in place by the
// current Framework provider. Step 1 creates the binding with citrix/citrixadc 2.2.0
// (legacy id "name,policyname"); step 2 refreshes/plans/applies the SAME config through
// the Framework provider, whose Read re-derives the canonical new-format id
// ("name:...,policyname:...").
func TestAccCrvserver_appflowpolicy_binding_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckCrvserver_appflowpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			// Step 1: create with the last SDK v2 release; state carries the legacy id.
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.2.0",
					},
				},
				Config: testAcccrvserver_appflowpolicy_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCrvserver_appflowpolicy_bindingExist("citrixadc_crvserver_appflowpolicy_binding.crvserver_appflowpolicy_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_crvserver_appflowpolicy_binding.crvserver_appflowpolicy_binding", "id", "my_vserver,tf_appflowpolicy"),
				),
			},
			// Step 2: refresh/plan/apply the same config through the Framework provider,
			// which upgrades the legacy id to the new key:value format on Read.
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAcccrvserver_appflowpolicy_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCrvserver_appflowpolicy_bindingExist("citrixadc_crvserver_appflowpolicy_binding.crvserver_appflowpolicy_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_crvserver_appflowpolicy_binding.crvserver_appflowpolicy_binding", "id", "name:my_vserver,policyname:tf_appflowpolicy"),
				),
			},
		},
	})
}
