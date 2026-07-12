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

const testAccAuthenticationvserver_cspolicy_binding_basic = `
	resource "citrixadc_authenticationvserver" "tf_authenticationvserver" {
		name           = "tf_authenticationvserver"
		servicetype    = "SSL"
		comment        = "new"
		authentication = "ON"
		state          = "DISABLED"
	}
	resource "citrixadc_lbvserver" "foo_lbvserver" {
		name        = "test_policy_lb"
		servicetype = "HTTP"
		ipv46       = "192.122.3.3"
		port        = 8000
		comment     = "hello"
	}
	resource "citrixadc_csaction" "tf_csaction" {
		name            = "tf_csaction"
		targetlbvserver = citrixadc_lbvserver.foo_lbvserver.name
	}
	resource "citrixadc_cspolicy" "foo_cspolicy" {
		policyname = "test_policy"
		rule       = "TRUE"
		action     = citrixadc_csaction.tf_csaction.name
	}
	resource "citrixadc_authenticationvserver_cspolicy_binding" "tf_bind" {
		name      = citrixadc_authenticationvserver.tf_authenticationvserver.name
		policy    = citrixadc_cspolicy.foo_cspolicy.policyname
		priority  = 90
		bindpoint = "REQUEST" #doesnot unbind for RESPONSE
	}
`

const testAccAuthenticationvserver_cspolicy_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion
	resource "citrixadc_authenticationvserver" "tf_authenticationvserver" {
		name           = "tf_authenticationvserver"
		servicetype    = "SSL"
		comment        = "new"
		authentication = "ON"
		state          = "DISABLED"
	}
	resource "citrixadc_lbvserver" "foo_lbvserver" {
		name        = "test_policy_lb"
		servicetype = "HTTP"
		ipv46       = "192.122.3.3"
		port        = 8000
		comment     = "hello"
	}
	resource "citrixadc_csaction" "tf_csaction" {
		name            = "tf_csaction"
		targetlbvserver = citrixadc_lbvserver.foo_lbvserver.name
	}
	resource "citrixadc_cspolicy" "foo_cspolicy" {
		policyname = "test_policy"
		rule       = "TRUE"
		action     = citrixadc_csaction.tf_csaction.name
	}
`

const testAccAuthenticationvserverCspolicyBindingDataSource_basic = `
	resource "citrixadc_authenticationvserver" "tf_authenticationvserver" {
		name           = "tf_authenticationvserver"
		servicetype    = "SSL"
		comment        = "new"
		authentication = "ON"
		state          = "DISABLED"
	}
	resource "citrixadc_lbvserver" "foo_lbvserver" {
		name        = "test_policy_lb"
		servicetype = "HTTP"
		ipv46       = "192.122.3.3"
		port        = 8000
		comment     = "hello"
	}
	resource "citrixadc_csaction" "tf_csaction" {
		name            = "tf_csaction"
		targetlbvserver = citrixadc_lbvserver.foo_lbvserver.name
	}
	resource "citrixadc_cspolicy" "foo_cspolicy" {
		policyname = "test_policy"
		rule       = "TRUE"
		action     = citrixadc_csaction.tf_csaction.name
	}
	resource "citrixadc_authenticationvserver_cspolicy_binding" "tf_bind" {
		name      = citrixadc_authenticationvserver.tf_authenticationvserver.name
		policy    = citrixadc_cspolicy.foo_cspolicy.policyname
		priority  = 90
	}

	data "citrixadc_authenticationvserver_cspolicy_binding" "tf_bind" {
		name      = citrixadc_authenticationvserver_cspolicy_binding.tf_bind.name
		policy    = citrixadc_authenticationvserver_cspolicy_binding.tf_bind.policy
		depends_on = [citrixadc_authenticationvserver_cspolicy_binding.tf_bind]
	}
`

func TestAccAuthenticationvserver_cspolicy_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationvserver_cspolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationvserver_cspolicy_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationvserver_cspolicy_bindingExist("citrixadc_authenticationvserver_cspolicy_binding.tf_bind", nil),
				),
			},
			{
				Config: testAccAuthenticationvserver_cspolicy_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationvserver_cspolicy_bindingNotExist("citrixadc_authenticationvserver_cspolicy_binding.tf_bind", "tf_authenticationvserver,test_policy"),
				),
			},
		},
	})
}

func TestAccAuthenticationvserver_cspolicy_binding_import(t *testing.T) {
	const resAddr = "citrixadc_authenticationvserver_cspolicy_binding.tf_bind"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationvserver_cspolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccAuthenticationvserver_cspolicy_binding_basic},
			{Config: testAccAuthenticationvserver_cspolicy_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{"bindpoint"}},
		},
	})
}

func testAccCheckAuthenticationvserver_cspolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No authenticationvserver_cspolicy_binding id is set")
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

		idMap, _, err := utils.ParseIdString(bindingId, []string{"name", "policy"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID: %v", err)
		}
		name := idMap["name"]
		policy := idMap["policy"]

		findParams := service.FindParams{
			ResourceType:             "authenticationvserver_cspolicy_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the matching secondIdComponent
		found := false
		for _, v := range dataArr {
			if v["policy"].(string) == policy {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("authenticationvserver_cspolicy_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckAuthenticationvserver_cspolicy_bindingNotExist(n string, id string) resource.TestCheckFunc {
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

		name := idSlice[0]
		policy := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "authenticationvserver_cspolicy_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to hopefully not find the one with the matching secondIdComponent
		found := false
		for _, v := range dataArr {
			if v["policy"].(string) == policy {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("authenticationvserver_cspolicy_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckAuthenticationvserver_cspolicy_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_authenticationvserver_cspolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Authenticationvserver_cspolicy_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("authenticationvserver_cspolicy_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccAuthenticationvserverCspolicyBindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationvserverCspolicyBindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_authenticationvserver_cspolicy_binding.tf_bind", "name", "tf_authenticationvserver"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationvserver_cspolicy_binding.tf_bind", "policy", "test_policy"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationvserver_cspolicy_binding.tf_bind", "priority", "90"),
				),
			},
		},
	})
}

// testAccauthenticationvserver_cspolicy_binding_upgrade_basic is the config used by
// the sdkv2 -> framework state-upgrade test. It reuses the same values and resource
// labels as testAccAuthenticationvserver_cspolicy_binding_basic so it is valid under
// BOTH the SDK v2 2.2.0 schema and the current framework schema.
const testAccauthenticationvserver_cspolicy_binding_upgrade_basic = `
	resource "citrixadc_authenticationvserver" "tf_authenticationvserver" {
		name           = "tf_authenticationvserver"
		servicetype    = "SSL"
		comment        = "new"
		authentication = "ON"
		state          = "DISABLED"
	}
	resource "citrixadc_lbvserver" "foo_lbvserver" {
		name        = "test_policy_lb"
		servicetype = "HTTP"
		ipv46       = "192.122.3.3"
		port        = 8000
		comment     = "hello"
	}
	resource "citrixadc_csaction" "tf_csaction" {
		name            = "tf_csaction"
		targetlbvserver = citrixadc_lbvserver.foo_lbvserver.name
	}
	resource "citrixadc_cspolicy" "foo_cspolicy" {
		policyname = "test_policy"
		rule       = "TRUE"
		action     = citrixadc_csaction.tf_csaction.name
	}
	resource "citrixadc_authenticationvserver_cspolicy_binding" "tf_bind" {
		name      = citrixadc_authenticationvserver.tf_authenticationvserver.name
		policy    = citrixadc_cspolicy.foo_cspolicy.policyname
		priority  = 90
		bindpoint = "REQUEST" #doesnot unbind for RESPONSE
	}
`

// TestAccAuthenticationvserver_cspolicy_binding_sdkv2StateUpgrade verifies that a
// binding created with the last SDK v2 release (2.2.0, legacy comma-separated ID) is
// correctly refreshed/planned/applied by the current framework provider. The
// framework recomputes the id on Read to the new key:value form.
func TestAccAuthenticationvserver_cspolicy_binding_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckAuthenticationvserver_cspolicy_bindingDestroy,
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
				Config: testAccauthenticationvserver_cspolicy_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationvserver_cspolicy_bindingExist("citrixadc_authenticationvserver_cspolicy_binding.tf_bind", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationvserver_cspolicy_binding.tf_bind", "id", "tf_authenticationvserver,test_policy"),
				),
			},
			// Step 2: same config, current (framework) provider. Terraform
			// refreshes the legacy-id state through the framework Read
			// (exercising ParseIdString on the legacy id) then plans/applies.
			// The framework recomputes the id on read to the new key:value form.
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccauthenticationvserver_cspolicy_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationvserver_cspolicy_bindingExist("citrixadc_authenticationvserver_cspolicy_binding.tf_bind", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationvserver_cspolicy_binding.tf_bind", "id", "name:tf_authenticationvserver,policy:test_policy"),
				),
			},
		},
	})
}
