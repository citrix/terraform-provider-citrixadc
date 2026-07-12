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

const testAccAaagroup_auditnslogpolicy_binding_basic = `

	resource "citrixadc_aaagroup" "tf_aaagroup" {
		groupname = "my_group"
		weight    = 100
	}
	resource "citrixadc_auditnslogaction" "tf_auditnslogaction" {
		name     = "my_auditnslogaction"
		serverip = "1.1.1.1"
		loglevel = ["ALERT", "CRITICAL"]
	}
	resource "citrixadc_auditnslogpolicy" "tf_auditnslogpolicy" {
		name   = "my_auditnslogpolicy"
		rule   = "ns_true"
		action = citrixadc_auditnslogaction.tf_auditnslogaction.name
	}

	resource "citrixadc_aaagroup_auditnslogpolicy_binding" "tf_aaagroup_auditnslogpolicy_binding" {
		groupname = citrixadc_aaagroup.tf_aaagroup.groupname
		policy    = citrixadc_auditnslogpolicy.tf_auditnslogpolicy.name
		type     = "REQUEST"
		priority  = 150
	}
`

const testAccAaagroup_auditnslogpolicy_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion
	resource "citrixadc_aaagroup" "tf_aaagroup" {
		groupname = "my_group"
		weight    = 100
	}
	resource "citrixadc_auditnslogaction" "tf_auditnslogaction" {
		name     = "my_auditnslogaction"
		serverip = "1.1.1.1"
		loglevel = ["ALERT", "CRITICAL"]
	}
	resource "citrixadc_auditnslogpolicy" "tf_auditnslogpolicy" {
		name   = "my_auditnslogpolicy"
		rule   = "ns_true"
		action = citrixadc_auditnslogaction.tf_auditnslogaction.name
	}
`

const testAccAaagroupAuditnslogpolicyBindingDataSource_basic = `

	resource "citrixadc_aaagroup" "tf_aaagroup" {
		groupname = "my_group"
		weight    = 100
	}
	resource "citrixadc_auditnslogaction" "tf_auditnslogaction" {
		name     = "my_auditnslogaction"
		serverip = "1.1.1.1"
		loglevel = ["ALERT", "CRITICAL"]
	}
	resource "citrixadc_auditnslogpolicy" "tf_auditnslogpolicy" {
		name   = "my_auditnslogpolicy"
		rule   = "ns_true"
		action = citrixadc_auditnslogaction.tf_auditnslogaction.name
	}

	resource "citrixadc_aaagroup_auditnslogpolicy_binding" "tf_aaagroup_auditnslogpolicy_binding" {
		groupname = citrixadc_aaagroup.tf_aaagroup.groupname
		policy    = citrixadc_auditnslogpolicy.tf_auditnslogpolicy.name
		type      = "REQUEST"
		priority  = 150
	}

	data "citrixadc_aaagroup_auditnslogpolicy_binding" "tf_aaagroup_auditnslogpolicy_binding" {
		groupname = citrixadc_aaagroup_auditnslogpolicy_binding.tf_aaagroup_auditnslogpolicy_binding.groupname
		policy    = citrixadc_aaagroup_auditnslogpolicy_binding.tf_aaagroup_auditnslogpolicy_binding.policy
		depends_on = [citrixadc_aaagroup_auditnslogpolicy_binding.tf_aaagroup_auditnslogpolicy_binding]
	}
`

func TestAccAaagroup_auditnslogpolicy_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAaagroup_auditnslogpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAaagroup_auditnslogpolicy_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaagroup_auditnslogpolicy_bindingExist("citrixadc_aaagroup_auditnslogpolicy_binding.tf_aaagroup_auditnslogpolicy_binding", nil),
				),
			},
			{
				Config: testAccAaagroup_auditnslogpolicy_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaagroup_auditnslogpolicy_bindingNotExist("citrixadc_aaagroup_auditnslogpolicy_binding.tf_aaagroup_auditnslogpolicy_binding", "my_group,tf_auditnslogpolicy"),
				),
			},
		},
	})
}

func TestAccAaagroupAuditnslogpolicyBindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccAaagroupAuditnslogpolicyBindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_aaagroup_auditnslogpolicy_binding.tf_aaagroup_auditnslogpolicy_binding", "groupname", "my_group"),
					resource.TestCheckResourceAttr("data.citrixadc_aaagroup_auditnslogpolicy_binding.tf_aaagroup_auditnslogpolicy_binding", "policy", "my_auditnslogpolicy"),
					resource.TestCheckResourceAttr("data.citrixadc_aaagroup_auditnslogpolicy_binding.tf_aaagroup_auditnslogpolicy_binding", "priority", "150"),
				),
			},
		},
	})
}

func TestAccAaagroup_auditnslogpolicy_binding_import(t *testing.T) {
	const resAddr = "citrixadc_aaagroup_auditnslogpolicy_binding.tf_aaagroup_auditnslogpolicy_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAaagroup_auditnslogpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccAaagroup_auditnslogpolicy_binding_basic},
			{Config: testAccAaagroup_auditnslogpolicy_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{"type"}},
		},
	})
}

func testAccCheckAaagroup_auditnslogpolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No aaagroup_auditnslogpolicy_binding id is set")
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

		idMap, _, err := utils.ParseIdString(bindingId, []string{"groupname", "policy"}, nil)
		if err != nil {
			return err
		}
		groupname := idMap["groupname"]
		policy := idMap["policy"]

		findParams := service.FindParams{
			ResourceType:             "aaagroup_auditnslogpolicy_binding",
			ResourceName:             groupname,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the matching policy
		found := false
		for _, v := range dataArr {
			if v["policy"].(string) == policy {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("aaagroup_auditnslogpolicy_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckAaagroup_auditnslogpolicy_bindingNotExist(n string, id string) resource.TestCheckFunc {
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

		groupname := idSlice[0]
		policy := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "aaagroup_auditnslogpolicy_binding",
			ResourceName:             groupname,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to hopefully not find the one with the matching policy
		found := false
		for _, v := range dataArr {
			if v["policy"].(string) == policy {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("aaagroup_auditnslogpolicy_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckAaagroup_auditnslogpolicy_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_aaagroup_auditnslogpolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Aaagroup_auditnslogpolicy_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("aaagroup_auditnslogpolicy_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

// testAccAaagroup_auditnslogpolicy_binding_upgrade_basic reuses the _basic config
// (binding + all prerequisite resources). It is valid under BOTH the SDK v2 2.2.0
// schema and the current Framework schema because the migration restored the SDK v2
// attribute names.
const testAccAaagroup_auditnslogpolicy_binding_upgrade_basic = `

	resource "citrixadc_aaagroup" "tf_aaagroup" {
		groupname = "my_group"
		weight    = 100
	}
	resource "citrixadc_auditnslogaction" "tf_auditnslogaction" {
		name     = "my_auditnslogaction"
		serverip = "1.1.1.1"
		loglevel = ["ALERT", "CRITICAL"]
	}
	resource "citrixadc_auditnslogpolicy" "tf_auditnslogpolicy" {
		name   = "my_auditnslogpolicy"
		rule   = "ns_true"
		action = citrixadc_auditnslogaction.tf_auditnslogaction.name
	}

	resource "citrixadc_aaagroup_auditnslogpolicy_binding" "tf_aaagroup_auditnslogpolicy_binding" {
		groupname = citrixadc_aaagroup.tf_aaagroup.groupname
		policy    = citrixadc_auditnslogpolicy.tf_auditnslogpolicy.name
		type      = "REQUEST"
		priority  = 150
	}
`

// TestAccAaagroup_auditnslogpolicy_binding_sdkv2StateUpgrade verifies that state
// written by the last SDK v2 release (legacy comma-separated ID) is correctly
// upgraded when the same config is subsequently managed by the current Framework
// provider. Step 1 creates the binding with citrix/citrixadc 2.2.0 (writes the
// legacy id "my_group,my_auditnslogpolicy"). Step 2 refreshes/plans/applies the same
// config through the Framework provider, exercising ParseIdString on the legacy id;
// because the Framework recomputes the id on Read (SetAttrFromGet), the id upgrades
// to the new "key:value" form.
func TestAccAaagroup_auditnslogpolicy_binding_sdkv2StateUpgrade(t *testing.T) {
	resourceAddr := "citrixadc_aaagroup_auditnslogpolicy_binding.tf_aaagroup_auditnslogpolicy_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckAaagroup_auditnslogpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			// Step 1: create with the last SDK v2 release -> state carries the legacy id.
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.2.0",
					},
				},
				Config: testAccAaagroup_auditnslogpolicy_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaagroup_auditnslogpolicy_bindingExist(resourceAddr, nil),
					resource.TestCheckResourceAttr(resourceAddr, "id", "my_group,my_auditnslogpolicy"),
				),
			},
			// Step 2: refresh/plan/apply the SAME config through the current Framework
			// provider. The legacy-id state is read via ParseIdString and the id is
			// recomputed to the new key:value format.
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccAaagroup_auditnslogpolicy_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaagroup_auditnslogpolicy_bindingExist(resourceAddr, nil),
					resource.TestCheckResourceAttr(resourceAddr, "id", "groupname:my_group,policy:my_auditnslogpolicy"),
				),
			},
		},
	})
}
