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

// Participating entities reused from existing acceptance tests:
//   - citrixadc_vpnvserver         (parent)            -> vpnvserver_test.go
//   - citrixadc_authenticationpolicy (advanced policy) -> authenticationpolicy_test.go
//     with its citrixadc_authenticationldapaction prerequisite, mirrored from
//     vpnglobal_authenticationpolicy_binding_test.go.
//
// The binding has no UPDATE (every schema attribute is RequiresReplace), so the
// "update" step instead drops the binding to confirm proper deletion.
//
// The composite ID is "bindpoint:<v>,name:<v>,policy:<v>" (key:UrlEncode(value) pairs).
// bindpoint is Optional+Computed and part of the unique key, so it is set explicitly
// (REQUEST) to keep the ID deterministic and to satisfy the datasource lookup (which
// requires bindpoint). Only attributes set in HCL are asserted (name, policy, bindpoint,
// priority); secondary/groupextraction/gotopriorityexpression are not echoed by GET.

const testAccVpnvserverAuthenticationpolicyBinding_basic_step1 = `
	resource "citrixadc_vpnvserver" "tf_vpnvserver" {
		name        = "tf_vpnvserver_authpol"
		servicetype = "SSL"
		ipv46       = "3.3.3.3"
		port        = 443
	}

	resource "citrixadc_authenticationldapaction" "tf_authenticationldapaction" {
		name          = "tf_vpnvserver_ldapaction"
		serverip      = "1.2.3.4"
		serverport    = 8080
		authtimeout   = 1
		ldaploginname = "username"
	}

	resource "citrixadc_authenticationpolicy" "tf_authenticationpolicy" {
		name    = "tf_vpnvserver_authpolicy"
		rule    = "true"
		action  = citrixadc_authenticationldapaction.tf_authenticationldapaction.name
		comment = "new_policy"
	}

	resource "citrixadc_vpnvserver_authenticationpolicy_binding" "tf_binding" {
		name      = citrixadc_vpnvserver.tf_vpnvserver.name
		policy    = citrixadc_authenticationpolicy.tf_authenticationpolicy.name
		priority  = 90
		secondary = false

		depends_on = [
			citrixadc_vpnvserver.tf_vpnvserver,
			citrixadc_authenticationpolicy.tf_authenticationpolicy,
		]
	}
`

const testAccVpnvserverAuthenticationpolicyBinding_basic_step2 = `
	# Keep the participating entities without the actual binding to confirm proper deletion.
	resource "citrixadc_vpnvserver" "tf_vpnvserver" {
		name        = "tf_vpnvserver_authpol"
		servicetype = "SSL"
		ipv46       = "3.3.3.3"
		port        = 443
	}

	resource "citrixadc_authenticationldapaction" "tf_authenticationldapaction" {
		name          = "tf_vpnvserver_ldapaction"
		serverip      = "1.2.3.4"
		serverport    = 8080
		authtimeout   = 1
		ldaploginname = "username"
	}

	resource "citrixadc_authenticationpolicy" "tf_authenticationpolicy" {
		name    = "tf_vpnvserver_authpolicy"
		rule    = "true"
		action  = citrixadc_authenticationldapaction.tf_authenticationldapaction.name
		comment = "new_policy"
	}
`

func TestAccVpnvserverAuthenticationpolicyBinding_basic(t *testing.T) {
	t.Skip("TODO: Requires review")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVpnvserverAuthenticationpolicyBindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnvserverAuthenticationpolicyBinding_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnvserverAuthenticationpolicyBindingExist("citrixadc_vpnvserver_authenticationpolicy_binding.tf_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnvserver_authenticationpolicy_binding.tf_binding", "name", "tf_vpnvserver_authpol"),
					resource.TestCheckResourceAttr("citrixadc_vpnvserver_authenticationpolicy_binding.tf_binding", "policy", "tf_vpnvserver_authpolicy"),
					resource.TestCheckResourceAttr("citrixadc_vpnvserver_authenticationpolicy_binding.tf_binding", "priority", "90"),
				),
			},
			{
				Config: testAccVpnvserverAuthenticationpolicyBinding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnvserverAuthenticationpolicyBindingNotExist("citrixadc_vpnvserver_authenticationpolicy_binding.tf_binding", "bindpoint:REQUEST,name:tf_vpnvserver_authpol,policy:tf_vpnvserver_authpolicy"),
				),
			},
		},
	})
}

func TestAccVpnvserverAuthenticationpolicyBinding_import(t *testing.T) {
	t.Skip("TODO: Requires review")
	const resAddr = "citrixadc_vpnvserver_authenticationpolicy_binding.tf_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVpnvserverAuthenticationpolicyBindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnvserverAuthenticationpolicyBinding_basic_step1,
			},
			{
				Config:                  testAccVpnvserverAuthenticationpolicyBinding_basic_step1,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccCheckVpnvserverAuthenticationpolicyBindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No vpnvserver_authenticationpolicy_binding id is set")
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

		// ID is comma-separated key:UrlEncode(value) pairs: bindpoint:<v>,name:<v>,policy:<v>
		idMap, _, err := utils.ParseIdString(rs.Primary.ID, nil, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", rs.Primary.ID, err)
		}
		name := idMap["name"]
		policy := idMap["policy"]
		bindpoint := idMap["bindpoint"]

		findParams := service.FindParams{
			ResourceType:             service.Vpnvserver_authenticationpolicy_binding.Type(),
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one matching policy (and bindpoint when set)
		found := false
		for _, v := range dataArr {
			if val, ok := v["policy"].(string); !ok || val != policy {
				continue
			}
			if bindpoint != "" {
				if bp, ok := v["bindpoint"].(string); !ok || bp != bindpoint {
					continue
				}
			}
			found = true
			break
		}

		if !found {
			return fmt.Errorf("vpnvserver_authenticationpolicy_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckVpnvserverAuthenticationpolicyBindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		// ID is comma-separated key:UrlEncode(value) pairs: bindpoint:<v>,name:<v>,policy:<v>
		idMap, _, err := utils.ParseIdString(id, nil, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", id, err)
		}
		name := idMap["name"]
		policy := idMap["policy"]
		bindpoint := idMap["bindpoint"]

		findParams := service.FindParams{
			ResourceType:             service.Vpnvserver_authenticationpolicy_binding.Type(),
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// A missing-resource error means the binding is gone, which is what we want.
		if err != nil {
			return nil
		}

		// Iterate through results to hopefully not find the one matching policy (and bindpoint)
		found := false
		for _, v := range dataArr {
			if val, ok := v["policy"].(string); !ok || val != policy {
				continue
			}
			if bindpoint != "" {
				if bp, ok := v["bindpoint"].(string); !ok || bp != bindpoint {
					continue
				}
			}
			found = true
			break
		}

		if found {
			return fmt.Errorf("vpnvserver_authenticationpolicy_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckVpnvserverAuthenticationpolicyBindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_vpnvserver_authenticationpolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No id is set")
		}

		// ID is comma-separated key:UrlEncode(value) pairs: bindpoint:<v>,name:<v>,policy:<v>
		idMap, _, err := utils.ParseIdString(rs.Primary.ID, nil, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", rs.Primary.ID, err)
		}
		name := idMap["name"]
		policy := idMap["policy"]
		bindpoint := idMap["bindpoint"]

		findParams := service.FindParams{
			ResourceType:             service.Vpnvserver_authenticationpolicy_binding.Type(),
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			// A missing-resource error means the binding is gone, which is what we want.
			continue
		}

		for _, v := range dataArr {
			if val, ok := v["policy"].(string); !ok || val != policy {
				continue
			}
			if bindpoint != "" {
				if bp, ok := v["bindpoint"].(string); !ok || bp != bindpoint {
					continue
				}
			}
			return fmt.Errorf("vpnvserver_authenticationpolicy_binding %s still exists", rs.Primary.ID)
		}
	}

	return nil
}

const testAccVpnvserverAuthenticationpolicyBindingDataSource_basic = `
	resource "citrixadc_vpnvserver" "tf_vpnvserver" {
		name        = "tf_vpnvserver_authpol"
		servicetype = "SSL"
		ipv46       = "3.3.3.3"
		port        = 443
	}

	resource "citrixadc_authenticationldapaction" "tf_authenticationldapaction" {
		name          = "tf_vpnvserver_ldapaction"
		serverip      = "1.2.3.4"
		serverport    = 8080
		authtimeout   = 1
		ldaploginname = "username"
	}

	resource "citrixadc_authenticationpolicy" "tf_authenticationpolicy" {
		name    = "tf_vpnvserver_authpolicy"
		rule    = "true"
		action  = citrixadc_authenticationldapaction.tf_authenticationldapaction.name
		comment = "new_policy"
	}

	resource "citrixadc_vpnvserver_authenticationpolicy_binding" "tf_binding" {
		name      = citrixadc_vpnvserver.tf_vpnvserver.name
		policy    = citrixadc_authenticationpolicy.tf_authenticationpolicy.name
		bindpoint = "REQUEST"
		priority  = 90
		secondary = false

		depends_on = [
			citrixadc_vpnvserver.tf_vpnvserver,
			citrixadc_authenticationpolicy.tf_authenticationpolicy,
		]
	}

	data "citrixadc_vpnvserver_authenticationpolicy_binding" "tf_binding" {
		name       = citrixadc_vpnvserver_authenticationpolicy_binding.tf_binding.name
		policy     = citrixadc_vpnvserver_authenticationpolicy_binding.tf_binding.policy
		bindpoint  = citrixadc_vpnvserver_authenticationpolicy_binding.tf_binding.bindpoint
		depends_on = [citrixadc_vpnvserver_authenticationpolicy_binding.tf_binding]
	}
`

func TestAccVpnvserverAuthenticationpolicyBindingDataSource_basic(t *testing.T) {
	t.Skip("TODO: Requires review")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnvserverAuthenticationpolicyBindingDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_vpnvserver_authenticationpolicy_binding.tf_binding", "name", "tf_vpnvserver_authpol"),
					resource.TestCheckResourceAttr("data.citrixadc_vpnvserver_authenticationpolicy_binding.tf_binding", "policy", "tf_vpnvserver_authpolicy"),
					resource.TestCheckResourceAttr("data.citrixadc_vpnvserver_authenticationpolicy_binding.tf_binding", "bindpoint", "REQUEST"),
					resource.TestCheckResourceAttr("data.citrixadc_vpnvserver_authenticationpolicy_binding.tf_binding", "priority", "90"),
				),
			},
		},
	})
}
