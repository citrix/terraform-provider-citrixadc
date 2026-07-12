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

const testAccVpnglobal_authenticationldappolicy_binding_basic = `

	resource "citrixadc_authenticationldapaction" "tf_authenticationldapaction" {
		name          = "tf_ldapaction"
		serverip      = "5.5.5.5"
		serverport    = 8080
		authtimeout   = 1
		ldaploginname = "username"
	}
	resource "citrixadc_authenticationldappolicy" "tf_authenticationldappolicy" {
		name      = "tf_ldappolicy"
		rule      = "NS_TRUE"
		reqaction = citrixadc_authenticationldapaction.tf_authenticationldapaction.name
	}
	resource "citrixadc_vpnglobal_authenticationldappolicy_binding" "tf_bind" {
		policyname = citrixadc_authenticationldappolicy.tf_authenticationldappolicy.name
		priority = 20
	}
 
`
const testAccVpnglobal_authenticationldappolicy_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion

	resource "citrixadc_authenticationldapaction" "tf_authenticationldapaction" {
		name          = "tf_ldapaction"
		serverip      = "5.5.5.5"
		serverport    = 8080
		authtimeout   = 1
		ldaploginname = "username"
	}
	resource "citrixadc_authenticationldappolicy" "tf_authenticationldappolicy" {
		name      = "tf_ldappolicy"
		rule      = "NS_TRUE"
		reqaction = citrixadc_authenticationldapaction.tf_authenticationldapaction.name
	}
`

func TestAccVpnglobal_authenticationldappolicy_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVpnglobal_authenticationldappolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnglobal_authenticationldappolicy_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnglobal_authenticationldappolicy_bindingExist("citrixadc_vpnglobal_authenticationldappolicy_binding.tf_bind", nil),
				),
			},
			{
				Config: testAccVpnglobal_authenticationldappolicy_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnglobal_authenticationldappolicy_bindingNotExist("citrixadc_vpnglobal_authenticationldappolicy_binding.tf_bind", "tf_ldappolicy"),
				),
			},
		},
	})
}

func TestAccVpnglobal_authenticationldappolicy_binding_import(t *testing.T) {
	const resAddr = "citrixadc_vpnglobal_authenticationldappolicy_binding.tf_bind"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVpnglobal_authenticationldappolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccVpnglobal_authenticationldappolicy_binding_basic},
			{Config: testAccVpnglobal_authenticationldappolicy_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
		},
	})
}

func testAccCheckVpnglobal_authenticationldappolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No vpnglobal_authenticationldappolicy_binding id is set")
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
			ResourceType:             "vpnglobal_authenticationldappolicy_binding",
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
			return fmt.Errorf("vpnglobal_authenticationldappolicy_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckVpnglobal_authenticationldappolicy_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		policyname := id

		findParams := service.FindParams{
			ResourceType:             "vpnglobal_authenticationldappolicy_binding",
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
			return fmt.Errorf("vpnglobal_authenticationldappolicy_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckVpnglobal_authenticationldappolicy_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_vpnglobal_authenticationldappolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Vpnglobal_authenticationldappolicy_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("vpnglobal_authenticationldappolicy_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccVpnglobal_authenticationldappolicy_bindingDataSource_basic = `
	resource "citrixadc_authenticationldapaction" "tf_authenticationldapaction" {
		name          = "tf_ldapaction"
		serverip      = "5.5.5.5"
		serverport    = 8080
		authtimeout   = 1
		ldaploginname = "username"
	}
	resource "citrixadc_authenticationldappolicy" "tf_authenticationldappolicy" {
		name      = "tf_ldappolicy"
		rule      = "NS_TRUE"
		reqaction = citrixadc_authenticationldapaction.tf_authenticationldapaction.name
	}
	resource "citrixadc_vpnglobal_authenticationldappolicy_binding" "tf_bind" {
		policyname = citrixadc_authenticationldappolicy.tf_authenticationldappolicy.name
		priority = 20
	}

	data "citrixadc_vpnglobal_authenticationldappolicy_binding" "tf_bind" {
		policyname = citrixadc_vpnglobal_authenticationldappolicy_binding.tf_bind.policyname
		depends_on = [citrixadc_vpnglobal_authenticationldappolicy_binding.tf_bind]
	}
`

func TestAccVpnglobal_authenticationldappolicy_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnglobal_authenticationldappolicy_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_vpnglobal_authenticationldappolicy_binding.tf_bind", "policyname", "tf_ldappolicy"),
					resource.TestCheckResourceAttr("data.citrixadc_vpnglobal_authenticationldappolicy_binding.tf_bind", "priority", "20"),
				),
			},
		},
	})
}

// Config for the SDK v2 -> Framework state-upgrade test. Reuses the _basic
// values and is valid under BOTH the last SDK v2 release (2.2.0) schema and the
// current Framework schema (uses only SDK v2 attribute names).
const testAccVpnglobal_authenticationldappolicy_binding_upgrade_basic = `

	resource "citrixadc_authenticationldapaction" "tf_authenticationldapaction" {
		name          = "tf_ldapaction"
		serverip      = "5.5.5.5"
		serverport    = 8080
		authtimeout   = 1
		ldaploginname = "username"
	}
	resource "citrixadc_authenticationldappolicy" "tf_authenticationldappolicy" {
		name      = "tf_ldappolicy"
		rule      = "NS_TRUE"
		reqaction = citrixadc_authenticationldapaction.tf_authenticationldapaction.name
	}
	resource "citrixadc_vpnglobal_authenticationldappolicy_binding" "tf_bind" {
		policyname = citrixadc_authenticationldappolicy.tf_authenticationldappolicy.name
		priority = 20
	}
`

// TestAccVpnglobal_authenticationldappolicy_binding_sdkv2StateUpgrade verifies that
// state written by the last SDK v2 release is transparently upgraded by the current
// Framework provider. This binding uses a single-key id (policyname), so the legacy
// id and the recomputed Framework canonical id are identical ("tf_ldappolicy").
// Step 1 creates the binding with citrix/citrixadc 2.2.0; step 2 refreshes/plans the
// same config through the current Framework provider, whose Read parses the id and
// recomputes it via SetAttrFromGet.
func TestAccVpnglobal_authenticationldappolicy_binding_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckVpnglobal_authenticationldappolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				// Step 1: create with the last SDK v2 release, writing the legacy id.
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.2.0",
					},
				},
				Config: testAccVpnglobal_authenticationldappolicy_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnglobal_authenticationldappolicy_bindingExist("citrixadc_vpnglobal_authenticationldappolicy_binding.tf_bind", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnglobal_authenticationldappolicy_binding.tf_bind", "id", "tf_ldappolicy"),
				),
			},
			{
				// Step 2: refresh/apply the same config through the current Framework
				// provider. Read exercises ParseIdString on the id, then recomputes it
				// via SetAttrFromGet. Single-key id => value is unchanged.
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccVpnglobal_authenticationldappolicy_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnglobal_authenticationldappolicy_bindingExist("citrixadc_vpnglobal_authenticationldappolicy_binding.tf_bind", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnglobal_authenticationldappolicy_binding.tf_bind", "id", "tf_ldappolicy"),
				),
			},
		},
	})
}
