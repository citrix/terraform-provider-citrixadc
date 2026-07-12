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

const testAccVpnglobal_authenticationcertpolicy_binding_basic = `
	resource "citrixadc_authenticationcertaction" "tf_certaction" {
		name                       = "tf_certaction"
		twofactor                  = "ON"
		defaultauthenticationgroup = "new_group"
		usernamefield              = "Subject:CN"
		groupnamefield             = "subject:grp"
	}
	resource "citrixadc_authenticationcertpolicy" "tf_certpolicy" {
		name      = "tf_certpolicy"
		rule      = "ns_true"
		reqaction = citrixadc_authenticationcertaction.tf_certaction.name
	}
	resource "citrixadc_vpnglobal_authenticationcertpolicy_binding" "tf_bind" {
		policyname      = citrixadc_authenticationcertpolicy.tf_certpolicy.name
		priority        = 20
		groupextraction = false
		secondary       = false
	}
`

const testAccVpnglobal_authenticationcertpolicy_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion
	resource "citrixadc_authenticationcertaction" "tf_certaction" {
		name                       = "tf_certaction"
		twofactor                  = "ON"
		defaultauthenticationgroup = "new_group"
		usernamefield              = "Subject:CN"
		groupnamefield             = "subject:grp"
	}
	resource "citrixadc_authenticationcertpolicy" "tf_certpolicy" {
		name      = "tf_certpolicy"
		rule      = "ns_true"
		reqaction = citrixadc_authenticationcertaction.tf_certaction.name
	}
`

func TestAccVpnglobal_authenticationcertpolicy_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVpnglobal_authenticationcertpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnglobal_authenticationcertpolicy_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnglobal_authenticationcertpolicy_bindingExist("citrixadc_vpnglobal_authenticationcertpolicy_binding.tf_bind", nil),
				),
			},
			{
				Config: testAccVpnglobal_authenticationcertpolicy_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnglobal_authenticationcertpolicy_bindingNotExist("citrixadc_vpnglobal_authenticationcertpolicy_binding.tf_bind", "tf_certpolicy"),
				),
			},
		},
	})
}

func testAccCheckVpnglobal_authenticationcertpolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No vpnglobal_authenticationcertpolicy_binding id is set")
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

		policyname := rs.Primary.ID

		findParams := service.FindParams{
			ResourceType:             "vpnglobal_authenticationcertpolicy_binding",
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
			if v["policyname"].(string) == policyname {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("vpnglobal_authenticationcertpolicy_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckVpnglobal_authenticationcertpolicy_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		policyname := id

		findParams := service.FindParams{
			ResourceType:             "vpnglobal_authenticationcertpolicy_binding",
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
			if v["policyname"].(string) == policyname {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("vpnglobal_authenticationcertpolicy_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckVpnglobal_authenticationcertpolicy_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_vpnglobal_authenticationcertpolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Vpnglobal_authenticationcertpolicy_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("vpnglobal_authenticationcertpolicy_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

// testAccVpnglobal_authenticationcertpolicy_binding_upgrade_basic reuses the _basic
// config (binding + all prerequisite resources). It is valid under BOTH the SDK v2
// 2.2.0 schema and the current Framework schema because the migration restored the
// SDK v2 attribute names.
const testAccVpnglobal_authenticationcertpolicy_binding_upgrade_basic = `
	resource "citrixadc_authenticationcertaction" "tf_certaction" {
		name                       = "tf_certaction"
		twofactor                  = "ON"
		defaultauthenticationgroup = "new_group"
		usernamefield              = "Subject:CN"
		groupnamefield             = "subject:grp"
	}
	resource "citrixadc_authenticationcertpolicy" "tf_certpolicy" {
		name      = "tf_certpolicy"
		rule      = "ns_true"
		reqaction = citrixadc_authenticationcertaction.tf_certaction.name
	}
	resource "citrixadc_vpnglobal_authenticationcertpolicy_binding" "tf_bind" {
		policyname      = citrixadc_authenticationcertpolicy.tf_certpolicy.name
		priority        = 20
		groupextraction = false
		secondary       = false
	}
`

// TestAccVpnglobal_authenticationcertpolicy_binding_sdkv2StateUpgrade verifies that
// state written by the last SDK v2 release is correctly upgraded when the same config
// is subsequently managed by the current Framework provider. Step 1 creates the
// binding with citrix/citrixadc 2.2.0 (writes the legacy id "tf_certpolicy"). Step 2
// refreshes/plans/applies the same config through the Framework provider; the
// Framework recomputes the id on Read (SetAttrFromGet). This resource keys the id on
// a single unique attribute (policyname), so the legacy id and the new canonical id
// are both the plain policyname value "tf_certpolicy".
func TestAccVpnglobal_authenticationcertpolicy_binding_sdkv2StateUpgrade(t *testing.T) {
	resourceAddr := "citrixadc_vpnglobal_authenticationcertpolicy_binding.tf_bind"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckVpnglobal_authenticationcertpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			// Step 1: create with the last SDK v2 release -> state carries the legacy id.
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.2.0",
					},
				},
				Config: testAccVpnglobal_authenticationcertpolicy_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnglobal_authenticationcertpolicy_bindingExist(resourceAddr, nil),
					resource.TestCheckResourceAttr(resourceAddr, "id", "tf_certpolicy"),
				),
			},
			// Step 2: refresh/plan/apply the SAME config through the current Framework
			// provider. The id is recomputed to the canonical new format, which for this
			// single-key binding is the plain policyname value.
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccVpnglobal_authenticationcertpolicy_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnglobal_authenticationcertpolicy_bindingExist(resourceAddr, nil),
					resource.TestCheckResourceAttr(resourceAddr, "id", "tf_certpolicy"),
				),
			},
		},
	})
}

const testAccVpnglobal_authenticationcertpolicy_bindingDataSource_basic = `
	resource "citrixadc_authenticationcertaction" "tf_certaction" {
		name                       = "tf_certaction"
		twofactor                  = "ON"
		defaultauthenticationgroup = "new_group"
		usernamefield              = "Subject:CN"
		groupnamefield             = "subject:grp"
	}
	resource "citrixadc_authenticationcertpolicy" "tf_certpolicy" {
		name      = "tf_certpolicy"
		rule      = "ns_true"
		reqaction = citrixadc_authenticationcertaction.tf_certaction.name
	}
	resource "citrixadc_vpnglobal_authenticationcertpolicy_binding" "tf_bind" {
		policyname      = citrixadc_authenticationcertpolicy.tf_certpolicy.name
		priority        = 20
		groupextraction = false
		secondary       = false
	}

	data "citrixadc_vpnglobal_authenticationcertpolicy_binding" "tf_bind" {
		policyname = citrixadc_vpnglobal_authenticationcertpolicy_binding.tf_bind.policyname
		depends_on = [citrixadc_vpnglobal_authenticationcertpolicy_binding.tf_bind]
	}
`

func TestAccVpnglobal_authenticationcertpolicy_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnglobal_authenticationcertpolicy_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_vpnglobal_authenticationcertpolicy_binding.tf_bind", "policyname", "tf_certpolicy"),
					resource.TestCheckResourceAttr("data.citrixadc_vpnglobal_authenticationcertpolicy_binding.tf_bind", "priority", "20"),
				),
			},
		},
	})
}

func TestAccVpnglobal_authenticationcertpolicy_binding_import(t *testing.T) {
	const resAddr = "citrixadc_vpnglobal_authenticationcertpolicy_binding.tf_bind"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVpnglobal_authenticationcertpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccVpnglobal_authenticationcertpolicy_binding_basic},
			{Config: testAccVpnglobal_authenticationcertpolicy_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{"groupextraction"}},
		},
	})
}
