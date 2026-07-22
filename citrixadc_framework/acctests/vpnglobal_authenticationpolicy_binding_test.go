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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// vpnglobal is a singleton on the ADC, so there is no parent resource to create.
// The binding's participating entity is the advanced authenticationpolicy (and its
// authenticationldapaction prerequisite), reused from authenticationpolicy_test.go.

const testAccVpnglobalAuthenticationpolicyBinding_basic_step1 = `
	resource "citrixadc_authenticationldapaction" "tf_authenticationldapaction" {
		name          = "tf_vpnglobal_ldapaction"
		serverip      = "1.2.3.4"
		serverport    = 8080
		authtimeout   = 1
		ldaploginname = "username"
	}

	resource "citrixadc_authenticationpolicy" "tf_authenticationpolicy" {
		name    = "tf_vpnglobal_authpolicy"
		rule    = "true"
		action  = citrixadc_authenticationldapaction.tf_authenticationldapaction.name
		comment = "new_policy"
	}

	resource "citrixadc_vpnglobal_authenticationpolicy_binding" "tf_binding" {
		policyname = citrixadc_authenticationpolicy.tf_authenticationpolicy.name
		priority   = 90
		secondary  = false

		depends_on = [citrixadc_authenticationpolicy.tf_authenticationpolicy]
	}
`

const testAccVpnglobalAuthenticationpolicyBinding_basic_step2 = `
	# Keep the participating entities without the actual binding to confirm proper deletion.
	resource "citrixadc_authenticationldapaction" "tf_authenticationldapaction" {
		name          = "tf_vpnglobal_ldapaction"
		serverip      = "1.2.3.4"
		serverport    = 8080
		authtimeout   = 1
		ldaploginname = "username"
	}

	resource "citrixadc_authenticationpolicy" "tf_authenticationpolicy" {
		name    = "tf_vpnglobal_authpolicy"
		rule    = "true"
		action  = citrixadc_authenticationldapaction.tf_authenticationldapaction.name
		comment = "new_policy"
	}
`

func TestAccVpnglobalAuthenticationpolicyBinding_basic(t *testing.T) {
	t.Skip("TODO: Requires review")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVpnglobalAuthenticationpolicyBindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnglobalAuthenticationpolicyBinding_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnglobalAuthenticationpolicyBindingExist("citrixadc_vpnglobal_authenticationpolicy_binding.tf_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnglobal_authenticationpolicy_binding.tf_binding", "policyname", "tf_vpnglobal_authpolicy"),
					resource.TestCheckResourceAttr("citrixadc_vpnglobal_authenticationpolicy_binding.tf_binding", "priority", "90"),
				),
			},
			{
				Config: testAccVpnglobalAuthenticationpolicyBinding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnglobalAuthenticationpolicyBindingNotExist("citrixadc_vpnglobal_authenticationpolicy_binding.tf_binding", "tf_vpnglobal_authpolicy"),
				),
			},
		},
	})
}

func TestAccVpnglobalAuthenticationpolicyBinding_import(t *testing.T) {
	t.Skip("TODO: Requires review")
	const resAddr = "citrixadc_vpnglobal_authenticationpolicy_binding.tf_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVpnglobalAuthenticationpolicyBindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnglobalAuthenticationpolicyBinding_basic_step1,
			},
			{
				Config:                  testAccVpnglobalAuthenticationpolicyBinding_basic_step1,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccCheckVpnglobalAuthenticationpolicyBindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No vpnglobal_authenticationpolicy_binding id is set")
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

		// ID is a plain value (single unique attr: policyname)
		policyname := rs.Primary.ID

		findParams := service.FindParams{
			ResourceType:             service.Vpnglobal_authenticationpolicy_binding.Type(),
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
			if val, ok := v["policyname"].(string); ok && val == policyname {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("vpnglobal_authenticationpolicy_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckVpnglobalAuthenticationpolicyBindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		policyname := id

		findParams := service.FindParams{
			ResourceType:             service.Vpnglobal_authenticationpolicy_binding.Type(),
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
			if val, ok := v["policyname"].(string); ok && val == policyname {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("vpnglobal_authenticationpolicy_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckVpnglobalAuthenticationpolicyBindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_vpnglobal_authenticationpolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		policyname := rs.Primary.ID

		findParams := service.FindParams{
			ResourceType:             service.Vpnglobal_authenticationpolicy_binding.Type(),
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			// A missing-resource error means the binding is gone, which is what we want.
			continue
		}

		for _, v := range dataArr {
			if val, ok := v["policyname"].(string); ok && val == policyname {
				return fmt.Errorf("vpnglobal_authenticationpolicy_binding %s still exists", rs.Primary.ID)
			}
		}
	}

	return nil
}

const testAccVpnglobalAuthenticationpolicyBindingDataSource_basic = `
	resource "citrixadc_authenticationldapaction" "tf_authenticationldapaction" {
		name          = "tf_vpnglobal_ldapaction"
		serverip      = "1.2.3.4"
		serverport    = 8080
		authtimeout   = 1
		ldaploginname = "username"
	}

	resource "citrixadc_authenticationpolicy" "tf_authenticationpolicy" {
		name    = "tf_vpnglobal_authpolicy"
		rule    = "true"
		action  = citrixadc_authenticationldapaction.tf_authenticationldapaction.name
		comment = "new_policy"
	}

	resource "citrixadc_vpnglobal_authenticationpolicy_binding" "tf_binding" {
		policyname = citrixadc_authenticationpolicy.tf_authenticationpolicy.name
		priority   = 90
		secondary  = false

		depends_on = [citrixadc_authenticationpolicy.tf_authenticationpolicy]
	}

	data "citrixadc_vpnglobal_authenticationpolicy_binding" "tf_binding" {
		policyname = citrixadc_vpnglobal_authenticationpolicy_binding.tf_binding.policyname
		depends_on = [citrixadc_vpnglobal_authenticationpolicy_binding.tf_binding]
	}
`

func TestAccVpnglobalAuthenticationpolicyBindingDataSource_basic(t *testing.T) {
	t.Skip("TODO: Requires review")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnglobalAuthenticationpolicyBindingDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_vpnglobal_authenticationpolicy_binding.tf_binding", "policyname", "tf_vpnglobal_authpolicy"),
					resource.TestCheckResourceAttr("data.citrixadc_vpnglobal_authenticationpolicy_binding.tf_binding", "priority", "90"),
				),
			},
		},
	})
}
