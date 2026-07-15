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

// tmglobal_tmsessionpolicy_binding is an immutable GLOBAL binding (no update / RequiresReplace).
// A tmsessionpolicy (which needs a tmsessionaction) must exist before it can be bound to tmglobal.
// The tmsessionaction + tmsessionpolicy config below is reused from tmsessionpolicy_test.go /
// tmsessionaction_test.go with distinct names to avoid collisions with those tests.

const testAccTmglobal_tmsessionpolicy_binding_basic = `

	resource "citrixadc_tmsessionaction" "tf_tmsessionaction_glb" {
		name                       = "tf_tmsessaction_glb"
		sesstimeout                = 10
		defaultauthorizationaction = "ALLOW"
		sso                        = "OFF"
	}
	resource "citrixadc_tmsessionpolicy" "tf_tmsessionpolicy_glb" {
		name   = "tf_tmsessionpolicy_glb"
		rule   = "true"
		action = citrixadc_tmsessionaction.tf_tmsessionaction_glb.name
	}

	resource "citrixadc_tmglobal_tmsessionpolicy_binding" "tf_tmglobal_tmsessionpolicy_binding" {
		policyname             = citrixadc_tmsessionpolicy.tf_tmsessionpolicy_glb.name
		priority               = 100
		gotopriorityexpression = "END"
	}
`

const testAccTmglobal_tmsessionpolicy_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion

	resource "citrixadc_tmsessionaction" "tf_tmsessionaction_glb" {
		name                       = "tf_tmsessaction_glb"
		sesstimeout                = 10
		defaultauthorizationaction = "ALLOW"
		sso                        = "OFF"
	}
	resource "citrixadc_tmsessionpolicy" "tf_tmsessionpolicy_glb" {
		name   = "tf_tmsessionpolicy_glb"
		rule   = "true"
		action = citrixadc_tmsessionaction.tf_tmsessionaction_glb.name
	}
`

const testAccTmglobalTmsessionpolicyBindingDataSource_basic = `

	resource "citrixadc_tmsessionaction" "tf_tmsessionaction_glb" {
		name                       = "tf_tmsessaction_glb"
		sesstimeout                = 10
		defaultauthorizationaction = "ALLOW"
		sso                        = "OFF"
	}
	resource "citrixadc_tmsessionpolicy" "tf_tmsessionpolicy_glb" {
		name   = "tf_tmsessionpolicy_glb"
		rule   = "true"
		action = citrixadc_tmsessionaction.tf_tmsessionaction_glb.name
	}

	resource "citrixadc_tmglobal_tmsessionpolicy_binding" "tf_tmglobal_tmsessionpolicy_binding" {
		policyname             = citrixadc_tmsessionpolicy.tf_tmsessionpolicy_glb.name
		priority               = 100
		gotopriorityexpression = "END"
	}

	data "citrixadc_tmglobal_tmsessionpolicy_binding" "tf_tmglobal_tmsessionpolicy_binding" {
		policyname = citrixadc_tmglobal_tmsessionpolicy_binding.tf_tmglobal_tmsessionpolicy_binding.policyname
		depends_on = [citrixadc_tmglobal_tmsessionpolicy_binding.tf_tmglobal_tmsessionpolicy_binding]
	}
`

func TestAccTmglobal_tmsessionpolicy_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckTmglobal_tmsessionpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTmglobal_tmsessionpolicy_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTmglobal_tmsessionpolicy_bindingExist("citrixadc_tmglobal_tmsessionpolicy_binding.tf_tmglobal_tmsessionpolicy_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_tmglobal_tmsessionpolicy_binding.tf_tmglobal_tmsessionpolicy_binding", "policyname", "tf_tmsessionpolicy_glb"),
					resource.TestCheckResourceAttr("citrixadc_tmglobal_tmsessionpolicy_binding.tf_tmglobal_tmsessionpolicy_binding", "priority", "100"),
					resource.TestCheckResourceAttr("citrixadc_tmglobal_tmsessionpolicy_binding.tf_tmglobal_tmsessionpolicy_binding", "gotopriorityexpression", "END"),
				),
			},
			{
				// Immutable binding (RequiresReplace, no update). Step 2 removes the binding
				// (keeping the participating entities) to verify proper deletion.
				Config: testAccTmglobal_tmsessionpolicy_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTmglobal_tmsessionpolicy_bindingNotExist("citrixadc_tmglobal_tmsessionpolicy_binding.tf_tmglobal_tmsessionpolicy_binding", "tf_tmsessionpolicy_glb"),
				),
			},
		},
	})
}

func TestAccTmglobal_tmsessionpolicy_binding_import(t *testing.T) {
	const resAddr = "citrixadc_tmglobal_tmsessionpolicy_binding.tf_tmglobal_tmsessionpolicy_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckTmglobal_tmsessionpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTmglobal_tmsessionpolicy_binding_basic,
			},
			{
				Config:                  testAccTmglobal_tmsessionpolicy_binding_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccCheckTmglobal_tmsessionpolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No tmglobal_tmsessionpolicy_binding id is set")
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

		// Global-binding read quirk: the typed tmglobal_tmsessionpolicy_binding and
		// umbrella tmglobal_binding endpoints return empty on this firmware; the bound
		// policies are only exposed by the base "tmglobal" endpoint (matched on policyname).
		findParams := service.FindParams{
			ResourceType:             "tmglobal",
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
			return fmt.Errorf("tmglobal_tmsessionpolicy_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckTmglobal_tmsessionpolicy_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		policyname := id

		// See note in Exist: read bound policies from the base "tmglobal" endpoint.
		findParams := service.FindParams{
			ResourceType:             "tmglobal",
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
			return fmt.Errorf("tmglobal_tmsessionpolicy_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckTmglobal_tmsessionpolicy_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_tmglobal_tmsessionpolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		// Read bound policies from the base "tmglobal" endpoint (the typed binding
		// endpoint returns empty on this firmware) and ensure ours is gone.
		findParams := service.FindParams{
			ResourceType:             "tmglobal",
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			return err
		}
		for _, v := range dataArr {
			if pn, ok := v["policyname"].(string); ok && pn == rs.Primary.ID {
				return fmt.Errorf("tmglobal_tmsessionpolicy_binding %s still exists", rs.Primary.ID)
			}
		}

	}

	return nil
}

func TestAccTmglobalTmsessionpolicyBindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccTmglobalTmsessionpolicyBindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_tmglobal_tmsessionpolicy_binding.tf_tmglobal_tmsessionpolicy_binding", "policyname", "tf_tmsessionpolicy_glb"),
					resource.TestCheckResourceAttr("data.citrixadc_tmglobal_tmsessionpolicy_binding.tf_tmglobal_tmsessionpolicy_binding", "priority", "100"),
					resource.TestCheckResourceAttr("data.citrixadc_tmglobal_tmsessionpolicy_binding.tf_tmglobal_tmsessionpolicy_binding", "gotopriorityexpression", "END"),
				),
			},
		},
	})
}
