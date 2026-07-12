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

const testAccAuditnslogglobal_auditnslogpolicy_binding_basic = `

	resource "citrixadc_auditnslogaction" "tf_auditnslogaction" {
		name     = "my_auditnslogaction"
		serverip = "1.1.1.1"
		loglevel = ["ALERT", "CRITICAL"]
	}
	resource "citrixadc_auditnslogpolicy" "tf_auditnslogpolicy" {
		name   = "my_auditnslogpolicy"
		rule   = "true"
		action = citrixadc_auditnslogaction.tf_auditnslogaction.name
	}

	resource "citrixadc_auditnslogglobal_auditnslogpolicy_binding" "tf_auditnslogglobal_auditnslogpolicy_binding" {
		policyname 		= citrixadc_auditnslogpolicy.tf_auditnslogpolicy.name
		priority   		= 100
		globalbindtype 	= "SYSTEM_GLOBAL"
	}
`

const testAccAuditnslogglobal_auditnslogpolicy_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion

	resource "citrixadc_auditnslogaction" "tf_auditnslogaction" {
		name     = "my_auditnslogaction"
		serverip = "1.1.1.1"
		loglevel = ["ALERT", "CRITICAL"]
	}
	resource "citrixadc_auditnslogpolicy" "tf_auditnslogpolicy" {
		name   = "my_auditnslogpolicy"
		rule   = "true"
		action = citrixadc_auditnslogaction.tf_auditnslogaction.name
	}
`

func TestAccAuditnslogglobal_auditnslogpolicy_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuditnslogglobal_auditnslogpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuditnslogglobal_auditnslogpolicy_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuditnslogglobal_auditnslogpolicy_bindingExist("citrixadc_auditnslogglobal_auditnslogpolicy_binding.tf_auditnslogglobal_auditnslogpolicy_binding", nil),
				),
			},
			{
				Config: testAccAuditnslogglobal_auditnslogpolicy_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuditnslogglobal_auditnslogpolicy_bindingNotExist("citrixadc_auditnslogglobal_auditnslogpolicy_binding.tf_auditnslogglobal_auditnslogpolicy_binding", "my_auditnslogpolicy"),
				),
			},
		},
	})
}

func testAccCheckAuditnslogglobal_auditnslogpolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No auditnslogglobal_auditnslogpolicy_binding id is set")
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

		idMap, _, err := utils.ParseIdString(rs.Primary.ID, []string{"policyname"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID: %v", err)
		}
		policyname := idMap["policyname"]

		findParams := service.FindParams{
			ResourceType:             "auditnslogglobal_auditnslogpolicy_binding",
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
			return fmt.Errorf("auditnslogglobal_auditnslogpolicy_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckAuditnslogglobal_auditnslogpolicy_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		policyname := id

		findParams := service.FindParams{
			ResourceType:             "auditnslogglobal_auditnslogpolicy_binding",
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
			return fmt.Errorf("auditnslogglobal_auditnslogpolicy_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckAuditnslogglobal_auditnslogpolicy_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_auditnslogglobal_auditnslogpolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("auditnslogglobal_auditnslogpolicy_binding", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("auditnslogglobal_auditnslogpolicy_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccAuditnslogglobalAuditnslogpolicyBindingDataSource_basic = `

	resource "citrixadc_auditnslogaction" "tf_auditnslogaction" {
		name     = "my_auditnslogaction"
		serverip = "1.1.1.1"
		loglevel = ["ALERT", "CRITICAL"]
	}
	resource "citrixadc_auditnslogpolicy" "tf_auditnslogpolicy" {
		name   = "my_auditnslogpolicy"
		rule   = "true"
		action = citrixadc_auditnslogaction.tf_auditnslogaction.name
	}

	resource "citrixadc_auditnslogglobal_auditnslogpolicy_binding" "tf_auditnslogglobal_auditnslogpolicy_binding" {
		policyname 		= citrixadc_auditnslogpolicy.tf_auditnslogpolicy.name
		priority   		= 100
		globalbindtype 	= "SYSTEM_GLOBAL"
	}

	data "citrixadc_auditnslogglobal_auditnslogpolicy_binding" "tf_auditnslogglobal_auditnslogpolicy_binding" {
		policyname     = citrixadc_auditnslogglobal_auditnslogpolicy_binding.tf_auditnslogglobal_auditnslogpolicy_binding.policyname
		globalbindtype = citrixadc_auditnslogglobal_auditnslogpolicy_binding.tf_auditnslogglobal_auditnslogpolicy_binding.globalbindtype
		depends_on     = [citrixadc_auditnslogglobal_auditnslogpolicy_binding.tf_auditnslogglobal_auditnslogpolicy_binding]
	}
`

func TestAccAuditnslogglobalAuditnslogpolicyBindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAuditnslogglobalAuditnslogpolicyBindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_auditnslogglobal_auditnslogpolicy_binding.tf_auditnslogglobal_auditnslogpolicy_binding", "policyname", "my_auditnslogpolicy"),
					resource.TestCheckResourceAttr("data.citrixadc_auditnslogglobal_auditnslogpolicy_binding.tf_auditnslogglobal_auditnslogpolicy_binding", "priority", "100"),
					resource.TestCheckResourceAttr("data.citrixadc_auditnslogglobal_auditnslogpolicy_binding.tf_auditnslogglobal_auditnslogpolicy_binding", "globalbindtype", "SYSTEM_GLOBAL"),
				),
			},
		},
	})
}

// testAccAuditnslogglobal_auditnslogpolicy_binding_upgrade_basic reuses the _basic
// config (binding + all prerequisite resources). It is valid under BOTH the SDK v2
// 2.2.0 schema and the current Framework schema because the migration restored the
// SDK v2 attribute names.
const testAccAuditnslogglobal_auditnslogpolicy_binding_upgrade_basic = `

	resource "citrixadc_auditnslogaction" "tf_auditnslogaction" {
		name     = "my_auditnslogaction"
		serverip = "1.1.1.1"
		loglevel = ["ALERT", "CRITICAL"]
	}
	resource "citrixadc_auditnslogpolicy" "tf_auditnslogpolicy" {
		name   = "my_auditnslogpolicy"
		rule   = "true"
		action = citrixadc_auditnslogaction.tf_auditnslogaction.name
	}

	resource "citrixadc_auditnslogglobal_auditnslogpolicy_binding" "tf_auditnslogglobal_auditnslogpolicy_binding" {
		policyname 		= citrixadc_auditnslogpolicy.tf_auditnslogpolicy.name
		priority   		= 100
		globalbindtype 	= "SYSTEM_GLOBAL"
	}
`

// TestAccAuditnslogglobal_auditnslogpolicy_binding_sdkv2StateUpgrade verifies that
// state written by the last SDK v2 release (legacy id) is correctly upgraded when the
// same config is subsequently managed by the current Framework provider. Step 1
// creates the binding with citrix/citrixadc 2.2.0 (writes the legacy id
// "my_auditnslogpolicy" from d.SetId(policyname)). Step 2 refreshes/plans/applies the
// same config through the Framework provider, exercising ParseIdString on the legacy
// id; because the Framework recomputes the id on Read (SetAttrFromGet), the id upgrades
// to the new "key:value" form.
func TestAccAuditnslogglobal_auditnslogpolicy_binding_sdkv2StateUpgrade(t *testing.T) {
	resourceAddr := "citrixadc_auditnslogglobal_auditnslogpolicy_binding.tf_auditnslogglobal_auditnslogpolicy_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckAuditnslogglobal_auditnslogpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			// Step 1: create with the last SDK v2 release -> state carries the legacy id.
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.2.0",
					},
				},
				Config: testAccAuditnslogglobal_auditnslogpolicy_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuditnslogglobal_auditnslogpolicy_bindingExist(resourceAddr, nil),
					resource.TestCheckResourceAttr(resourceAddr, "id", "my_auditnslogpolicy"),
				),
			},
			// Step 2: refresh/plan/apply the SAME config through the current Framework
			// provider. The legacy-id state is read via ParseIdString and the id is
			// recomputed to the new key:value format.
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccAuditnslogglobal_auditnslogpolicy_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuditnslogglobal_auditnslogpolicy_bindingExist(resourceAddr, nil),
					resource.TestCheckResourceAttr(resourceAddr, "id", "globalbindtype:SYSTEM_GLOBAL,policyname:my_auditnslogpolicy"),
				),
			},
		},
	})
}

func TestAccAuditnslogglobal_auditnslogpolicy_binding_import(t *testing.T) {
	const resAddr = "citrixadc_auditnslogglobal_auditnslogpolicy_binding.tf_auditnslogglobal_auditnslogpolicy_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuditnslogglobal_auditnslogpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccAuditnslogglobal_auditnslogpolicy_binding_basic},
			{Config: testAccAuditnslogglobal_auditnslogpolicy_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
		},
	})
}
