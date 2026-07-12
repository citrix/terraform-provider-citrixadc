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

const testAccSystemglobal_authenticationradiuspolicy_binding_basic = `

	resource "citrixadc_systemglobal_authenticationradiuspolicy_binding" "tf_systemglobal_authenticationradiuspolicy_binding" {
		policyname = citrixadc_authenticationradiuspolicy.tf_radiuspolicy.name
		priority   = 50
	}
	
	resource "citrixadc_authenticationradiusaction" "tf_radiusaction" {
		name         = "tf_radiusaction"
		radkey       = "secret"
		serverip     = "1.2.3.4"
		serverport   = 8080
		authtimeout  = 2
		radnasip     = "DISABLED"
		passencoding = "chap"
	}
	resource "citrixadc_authenticationradiuspolicy" "tf_radiuspolicy" {
		name      = "tf_radiuspolicy"
		rule      = "NS_TRUE"
		reqaction = citrixadc_authenticationradiusaction.tf_radiusaction.name
	}
`

const testAccSystemglobal_authenticationradiuspolicy_binding_basic_step2 = `
	
resource "citrixadc_authenticationradiusaction" "tf_radiusaction" {
	name         = "tf_radiusaction"
	radkey       = "secret"
	serverip     = "1.2.3.4"
	serverport   = 8080
	authtimeout  = 2
	radnasip     = "DISABLED"
	passencoding = "chap"
}
resource "citrixadc_authenticationradiuspolicy" "tf_radiuspolicy" {
	name      = "tf_radiuspolicy"
	rule      = "NS_TRUE"
	reqaction = citrixadc_authenticationradiusaction.tf_radiusaction.name
}
`

func TestAccSystemglobal_authenticationradiuspolicy_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSystemglobal_authenticationradiuspolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSystemglobal_authenticationradiuspolicy_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSystemglobal_authenticationradiuspolicy_bindingExist("citrixadc_systemglobal_authenticationradiuspolicy_binding.tf_systemglobal_authenticationradiuspolicy_binding", nil),
				),
			},
			{
				Config: testAccSystemglobal_authenticationradiuspolicy_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSystemglobal_authenticationradiuspolicy_bindingNotExist("citrixadc_systemglobal_authenticationradiuspolicy_binding.tf_systemglobal_authenticationradiuspolicy_binding", "tf_radiuspolicy"),
				),
			},
		},
	})
}

func testAccCheckSystemglobal_authenticationradiuspolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No systemglobal_authenticationradiuspolicy_binding id is set")
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
			ResourceType:             "systemglobal_authenticationradiuspolicy_binding",
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
			return fmt.Errorf("systemglobal_authenticationradiuspolicy_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckSystemglobal_authenticationradiuspolicy_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		policyname := id

		findParams := service.FindParams{
			ResourceType:             "systemglobal_authenticationradiuspolicy_binding",
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
			return fmt.Errorf("systemglobal_authenticationradiuspolicy_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckSystemglobal_authenticationradiuspolicy_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_systemglobal_authenticationradiuspolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Systemglobal_authenticationradiuspolicy_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("systemglobal_authenticationradiuspolicy_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccSystemglobal_authenticationradiuspolicy_bindingDataSource_basic = `
	resource "citrixadc_authenticationradiusaction" "tf_radiusaction" {
		name         = "tf_radiusaction"
		radkey       = "secret"
		serverip     = "1.2.3.4"
		serverport   = 8080
		authtimeout  = 2
		radnasip     = "DISABLED"
		passencoding = "chap"
	}
	resource "citrixadc_authenticationradiuspolicy" "tf_radiuspolicy" {
		name      = "tf_radiuspolicy"
		rule      = "NS_TRUE"
		reqaction = citrixadc_authenticationradiusaction.tf_radiusaction.name
	}
	resource "citrixadc_systemglobal_authenticationradiuspolicy_binding" "tf_systemglobal_authenticationradiuspolicy_binding" {
		policyname = citrixadc_authenticationradiuspolicy.tf_radiuspolicy.name
		priority   = 50
	}
	
	data "citrixadc_systemglobal_authenticationradiuspolicy_binding" "tf_systemglobal_authenticationradiuspolicy_binding" {
		policyname = citrixadc_authenticationradiuspolicy.tf_radiuspolicy.name
		depends_on = [citrixadc_systemglobal_authenticationradiuspolicy_binding.tf_systemglobal_authenticationradiuspolicy_binding]
	}
`

func TestAccSystemglobal_authenticationradiuspolicy_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSystemglobal_authenticationradiuspolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSystemglobal_authenticationradiuspolicy_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_systemglobal_authenticationradiuspolicy_binding.tf_systemglobal_authenticationradiuspolicy_binding", "policyname", "tf_radiuspolicy"),
					resource.TestCheckResourceAttr("data.citrixadc_systemglobal_authenticationradiuspolicy_binding.tf_systemglobal_authenticationradiuspolicy_binding", "priority", "50"),
				),
			},
		},
	})
}

const testAccSystemglobal_authenticationradiuspolicy_binding_upgrade_basic = `

	resource "citrixadc_systemglobal_authenticationradiuspolicy_binding" "tf_systemglobal_authenticationradiuspolicy_binding" {
		policyname = citrixadc_authenticationradiuspolicy.tf_radiuspolicy.name
		priority   = 50
	}

	resource "citrixadc_authenticationradiusaction" "tf_radiusaction" {
		name         = "tf_radiusaction"
		radkey       = "secret"
		serverip     = "1.2.3.4"
		serverport   = 8080
		authtimeout  = 2
		radnasip     = "DISABLED"
		passencoding = "chap"
	}
	resource "citrixadc_authenticationradiuspolicy" "tf_radiuspolicy" {
		name      = "tf_radiuspolicy"
		rule      = "NS_TRUE"
		reqaction = citrixadc_authenticationradiusaction.tf_radiusaction.name
	}
`

// TestAccSystemglobal_authenticationradiuspolicy_binding_sdkv2StateUpgrade verifies that
// state written by the last SDK v2 release (legacy id) is correctly upgraded when the
// same config is subsequently managed by the current Framework provider. Step 1 creates
// the binding with citrix/citrixadc 2.2.0, which writes state using the legacy id
// (d.SetId(policyname) -> "tf_radiuspolicy"). Step 2 refreshes/plans/applies the SAME
// config through the Framework provider, exercising ParseIdString on the legacy id. This
// is a single-unique-attribute global binding, so the Framework recomputes the id on Read
// (SetAttrFromGet: data.Id = policyname) to the plain policyname value, which equals the
// legacy id "tf_radiuspolicy".
func TestAccSystemglobal_authenticationradiuspolicy_binding_sdkv2StateUpgrade(t *testing.T) {
	resourceAddr := "citrixadc_systemglobal_authenticationradiuspolicy_binding.tf_systemglobal_authenticationradiuspolicy_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckSystemglobal_authenticationradiuspolicy_bindingDestroy,
		Steps: []resource.TestStep{
			// Step 1: create with the last SDK v2 release -> state carries the legacy id.
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.2.0",
					},
				},
				Config: testAccSystemglobal_authenticationradiuspolicy_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSystemglobal_authenticationradiuspolicy_bindingExist(resourceAddr, nil),
					resource.TestCheckResourceAttr(resourceAddr, "id", "tf_radiuspolicy"),
				),
			},
			// Step 2: refresh/plan/apply the SAME config through the current Framework
			// provider. The legacy-id state is read via ParseIdString and the id is
			// recomputed on Read (SetAttrFromGet). Single unique attribute -> the new
			// canonical id is the plain policyname value "tf_radiuspolicy".
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccSystemglobal_authenticationradiuspolicy_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSystemglobal_authenticationradiuspolicy_bindingExist(resourceAddr, nil),
					resource.TestCheckResourceAttr(resourceAddr, "id", "tf_radiuspolicy"),
				),
			},
		},
	})
}

func TestAccSystemglobal_authenticationradiuspolicy_binding_import(t *testing.T) {
	const resAddr = "citrixadc_systemglobal_authenticationradiuspolicy_binding.tf_systemglobal_authenticationradiuspolicy_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSystemglobal_authenticationradiuspolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccSystemglobal_authenticationradiuspolicy_binding_basic},
			{Config: testAccSystemglobal_authenticationradiuspolicy_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
		},
	})
}
