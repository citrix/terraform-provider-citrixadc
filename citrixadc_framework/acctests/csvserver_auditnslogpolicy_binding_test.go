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

const testAccCsvserver_auditnslogpolicy_binding_basic = `

	resource "citrixadc_auditnslogaction" "tf_auditnslogaction" {
		name     = "tf_auditnslogaction"
		serverip = "1.1.1.1"
		loglevel = ["ALERT", "CRITICAL"]
	}
	resource "citrixadc_auditnslogpolicy" "tf_auditnslogpolicy" {
		name   = "tf_auditnslogpolicy"
		rule   = "ns_true"
		action = citrixadc_auditnslogaction.tf_auditnslogaction.name
	}

	resource "citrixadc_csvserver_auditnslogpolicy_binding" "tf_csvserver_auditnslogpolicy_binding" {
        name 		= citrixadc_csvserver.tf_csvserver.name
        policyname 	= citrixadc_auditnslogpolicy.tf_auditnslogpolicy.name
        priority 	= 5
	}

	resource "citrixadc_csvserver" "tf_csvserver" {
		name 		= "tf_csvserver"
		ipv46 		= "10.202.11.11"
		port 		= 8080
		servicetype = "HTTP"
	}
`

const testAccCsvserver_auditnslogpolicy_binding_basic_step2 = `

	resource "citrixadc_auditnslogaction" "tf_auditnslogaction" {
		name     = "tf_auditnslogaction"
		serverip = "1.1.1.1"
		loglevel = ["ALERT", "CRITICAL"]
	}
	resource "citrixadc_auditnslogpolicy" "tf_auditnslogpolicy" {
		name   = "tf_auditnslogpolicy"
		rule   = "ns_true"
		action = citrixadc_auditnslogaction.tf_auditnslogaction.name
	}
	resource "citrixadc_csvserver" "tf_csvserver" {
		name 		= "tf_csvserver"
		ipv46 		= "10.202.11.11"
		port 		= 8080
		servicetype = "HTTP"
	}
`

const testAccCsvserver_auditnslogpolicy_bindingDataSource_basic = `

	resource "citrixadc_auditnslogaction" "tf_auditnslogaction" {
		name     = "tf_auditnslogaction"
		serverip = "1.1.1.1"
		loglevel = ["ALERT", "CRITICAL"]
	}
	resource "citrixadc_auditnslogpolicy" "tf_auditnslogpolicy" {
		name   = "tf_auditnslogpolicy"
		rule   = "ns_true"
		action = citrixadc_auditnslogaction.tf_auditnslogaction.name
	}

	resource "citrixadc_csvserver_auditnslogpolicy_binding" "tf_csvserver_auditnslogpolicy_binding" {
        name 		= citrixadc_csvserver.tf_csvserver.name
        policyname 	= citrixadc_auditnslogpolicy.tf_auditnslogpolicy.name
        priority 	= 5
	}

	resource "citrixadc_csvserver" "tf_csvserver" {
		name 		= "tf_csvserver"
		ipv46 		= "10.202.11.11"
		port 		= 8080
		servicetype = "HTTP"
	}

	data "citrixadc_csvserver_auditnslogpolicy_binding" "tf_csvserver_auditnslogpolicy_binding" {
		name       = citrixadc_csvserver_auditnslogpolicy_binding.tf_csvserver_auditnslogpolicy_binding.name
		policyname = citrixadc_csvserver_auditnslogpolicy_binding.tf_csvserver_auditnslogpolicy_binding.policyname
	}
`

func TestAccCsvserver_auditnslogpolicy_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCsvserver_auditnslogpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCsvserver_auditnslogpolicy_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCsvserver_auditnslogpolicy_bindingExist("citrixadc_csvserver_auditnslogpolicy_binding.tf_csvserver_auditnslogpolicy_binding", nil),
				),
			},
			{
				Config: testAccCsvserver_auditnslogpolicy_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCsvserver_auditnslogpolicy_bindingNotExist("citrixadc_csvserver_auditnslogpolicy_binding.tf_csvserver_auditnslogpolicy_binding", "tf_csvserver,tf_auditnslogpolicy"),
				),
			},
		},
	})
}

func TestAccCsvserver_auditnslogpolicy_binding_import(t *testing.T) {
	const resAddr = "citrixadc_csvserver_auditnslogpolicy_binding.tf_csvserver_auditnslogpolicy_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCsvserver_auditnslogpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccCsvserver_auditnslogpolicy_binding_basic},
			{Config: testAccCsvserver_auditnslogpolicy_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
		},
	})
}

func testAccCheckCsvserver_auditnslogpolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No csvserver_auditnslogpolicy_binding id is set")
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
			return fmt.Errorf("Error parsing ID %s: %v", bindingId, err)
		}
		name := idMap["name"]
		policyname := idMap["policyname"]

		findParams := service.FindParams{
			ResourceType:             "csvserver_auditnslogpolicy_binding",
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
			return fmt.Errorf("csvserver_auditnslogpolicy_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckCsvserver_auditnslogpolicy_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		idMap, _, err := utils.ParseIdString(id, []string{"name", "policyname"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", id, err)
		}
		csvserverName := idMap["name"]
		policyName := idMap["policyname"]

		findParams := service.FindParams{
			ResourceType:             "csvserver_auditnslogpolicy_binding",
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
			return fmt.Errorf("csvserver_auditnslogpolicy_binding %s was found, but is should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckCsvserver_auditnslogpolicy_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_csvserver_auditnslogpolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Csvserver_auditnslogpolicy_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("csvserver_auditnslogpolicy_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccCsvserver_auditnslogpolicy_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccCsvserver_auditnslogpolicy_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_csvserver_auditnslogpolicy_binding.tf_csvserver_auditnslogpolicy_binding", "name", "tf_csvserver"),
					resource.TestCheckResourceAttr("data.citrixadc_csvserver_auditnslogpolicy_binding.tf_csvserver_auditnslogpolicy_binding", "policyname", "tf_auditnslogpolicy"),
					resource.TestCheckResourceAttr("data.citrixadc_csvserver_auditnslogpolicy_binding.tf_csvserver_auditnslogpolicy_binding", "priority", "5"),
				),
			},
		},
	})
}

// testAccCsvserver_auditnslogpolicy_binding_upgrade_basic reuses the _basic config
// (binding + all prerequisite resources). It is valid under BOTH the SDK v2 2.2.0
// schema and the current Framework schema because the migration restored the SDK v2
// attribute names.
const testAccCsvserver_auditnslogpolicy_binding_upgrade_basic = `

	resource "citrixadc_auditnslogaction" "tf_auditnslogaction" {
		name     = "tf_auditnslogaction"
		serverip = "1.1.1.1"
		loglevel = ["ALERT", "CRITICAL"]
	}
	resource "citrixadc_auditnslogpolicy" "tf_auditnslogpolicy" {
		name   = "tf_auditnslogpolicy"
		rule   = "ns_true"
		action = citrixadc_auditnslogaction.tf_auditnslogaction.name
	}

	resource "citrixadc_csvserver_auditnslogpolicy_binding" "tf_csvserver_auditnslogpolicy_binding" {
        name 		= citrixadc_csvserver.tf_csvserver.name
        policyname 	= citrixadc_auditnslogpolicy.tf_auditnslogpolicy.name
        priority 	= 5
	}

	resource "citrixadc_csvserver" "tf_csvserver" {
		name 		= "tf_csvserver"
		ipv46 		= "10.202.11.11"
		port 		= 8080
		servicetype = "HTTP"
	}
`

// TestAccCsvserver_auditnslogpolicy_binding_sdkv2StateUpgrade verifies that state
// written by the last SDK v2 release (legacy comma-separated ID) is correctly
// upgraded when the same config is subsequently managed by the current Framework
// provider. Step 1 creates the binding with citrix/citrixadc 2.2.0 (writes the
// legacy id "tf_csvserver,tf_auditnslogpolicy"). Step 2 refreshes/plans/applies the
// same config through the Framework provider, exercising ParseIdString on the legacy
// id; because the Framework recomputes the id on Read (SetAttrFromGet), the id
// upgrades to the new "key:value" form.
func TestAccCsvserver_auditnslogpolicy_binding_sdkv2StateUpgrade(t *testing.T) {
	resourceAddr := "citrixadc_csvserver_auditnslogpolicy_binding.tf_csvserver_auditnslogpolicy_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckCsvserver_auditnslogpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			// Step 1: create with the last SDK v2 release -> state carries the legacy id.
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.2.0",
					},
				},
				Config: testAccCsvserver_auditnslogpolicy_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCsvserver_auditnslogpolicy_bindingExist(resourceAddr, nil),
					resource.TestCheckResourceAttr(resourceAddr, "id", "tf_csvserver,tf_auditnslogpolicy"),
				),
			},
			// Step 2: refresh/plan/apply the SAME config through the current Framework
			// provider. The legacy-id state is read via ParseIdString and the id is
			// recomputed to the new key:value format.
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccCsvserver_auditnslogpolicy_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCsvserver_auditnslogpolicy_bindingExist(resourceAddr, nil),
					resource.TestCheckResourceAttr(resourceAddr, "id", "name:tf_csvserver,policyname:tf_auditnslogpolicy"),
				),
			},
		},
	})
}
