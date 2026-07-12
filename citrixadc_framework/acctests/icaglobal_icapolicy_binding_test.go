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

const testAccIcaglobal_icapolicy_binding_basic = `

	resource "citrixadc_icaaction" "tf_icaaction" {
		name              = "tf_icaaction"
		accessprofilename = "default_ica_accessprofile"
	}
	resource "citrixadc_icapolicy" "tf_icapolicy" {
		name   = "tf_icapolicy"
		rule   = true
		action = citrixadc_icaaction.tf_icaaction.name
	}

	resource "citrixadc_icaglobal_icapolicy_binding" "tf_icaglobal_icapolicy_binding" {
		policyname = citrixadc_icapolicy.tf_icapolicy.name
		priority   = 100
		type       = "ICA_REQ_DEFAULT"
	}
`

const testAccIcaglobal_icapolicy_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion

	resource "citrixadc_icaaction" "tf_icaaction" {
		name              = "tf_icaaction"
		accessprofilename = "default_ica_accessprofile"
	}
	resource "citrixadc_icapolicy" "tf_icapolicy" {
		name   = "tf_icapolicy"
		rule   = true
		action = citrixadc_icaaction.tf_icaaction.name
	}
`

const testAccIcaglobal_icapolicy_bindingDataSource_basic = `

	resource "citrixadc_icaaction" "tf_icaaction" {
		name              = "tf_icaaction"
		accessprofilename = "default_ica_accessprofile"
	}
	resource "citrixadc_icapolicy" "tf_icapolicy" {
		name   = "tf_icapolicy"
		rule   = true
		action = citrixadc_icaaction.tf_icaaction.name
	}

	resource "citrixadc_icaglobal_icapolicy_binding" "tf_icaglobal_icapolicy_binding" {
		policyname = citrixadc_icapolicy.tf_icapolicy.name
		priority   = 100
		type       = "ICA_REQ_DEFAULT"
	}

	data "citrixadc_icaglobal_icapolicy_binding" "tf_icaglobal_icapolicy_binding" {
		policyname = citrixadc_icaglobal_icapolicy_binding.tf_icaglobal_icapolicy_binding.policyname
		type       = citrixadc_icaglobal_icapolicy_binding.tf_icaglobal_icapolicy_binding.type
		depends_on = [citrixadc_icaglobal_icapolicy_binding.tf_icaglobal_icapolicy_binding]
	}
`

func TestAccIcaglobal_icapolicy_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckIcaglobal_icapolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccIcaglobal_icapolicy_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIcaglobal_icapolicy_bindingExist("citrixadc_icaglobal_icapolicy_binding.tf_icaglobal_icapolicy_binding", nil),
				),
			},
			{
				Config: testAccIcaglobal_icapolicy_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIcaglobal_icapolicy_bindingNotExist("citrixadc_icaglobal_icapolicy_binding.tf_icaglobal_icapolicy_binding", "tf_icapolicy", "ICA_REQ_DEFAULT"),
				),
			},
		},
	})
}

func TestAccIcaglobal_icapolicy_binding_import(t *testing.T) {
	const resAddr = "citrixadc_icaglobal_icapolicy_binding.tf_icaglobal_icapolicy_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckIcaglobal_icapolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccIcaglobal_icapolicy_binding_basic},
			{Config: testAccIcaglobal_icapolicy_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
		},
	})
}

func testAccCheckIcaglobal_icapolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No icaglobal_icapolicy_binding id is set")
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
			ResourceType:             "icaglobal_icapolicy_binding",
			ArgsMap:                  map[string]string{"type": rs.Primary.Attributes["type"]},
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
			return fmt.Errorf("icaglobal_icapolicy_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckIcaglobal_icapolicy_bindingNotExist(n string, id string, typename string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		policyname := id

		findParams := service.FindParams{
			ResourceType:             "icaglobal_icapolicy_binding",
			ArgsMap:                  map[string]string{"type": typename},
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
			return fmt.Errorf("icaglobal_icapolicy_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckIcaglobal_icapolicy_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_icaglobal_icapolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("icaglobal_icapolicy_binding", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("icaglobal_icapolicy_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccIcaglobal_icapolicy_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIcaglobal_icapolicy_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_icaglobal_icapolicy_binding.tf_icaglobal_icapolicy_binding", "policyname", "tf_icapolicy"),
					resource.TestCheckResourceAttr("data.citrixadc_icaglobal_icapolicy_binding.tf_icaglobal_icapolicy_binding", "priority", "100"),
					resource.TestCheckResourceAttr("data.citrixadc_icaglobal_icapolicy_binding.tf_icaglobal_icapolicy_binding", "type", "ICA_REQ_DEFAULT"),
				),
			},
		},
	})
}

const testAccIcaglobal_icapolicy_binding_upgrade_basic = `

	resource "citrixadc_icaaction" "tf_icaaction" {
		name              = "tf_icaaction"
		accessprofilename = "default_ica_accessprofile"
	}
	resource "citrixadc_icapolicy" "tf_icapolicy" {
		name   = "tf_icapolicy"
		rule   = true
		action = citrixadc_icaaction.tf_icaaction.name
	}

	resource "citrixadc_icaglobal_icapolicy_binding" "tf_icaglobal_icapolicy_binding" {
		policyname = citrixadc_icapolicy.tf_icapolicy.name
		priority   = 100
		type       = "ICA_REQ_DEFAULT"
	}
`

// TestAccIcaglobal_icapolicy_binding_sdkv2StateUpgrade verifies that a binding created
// with the last SDK v2 release (legacy plain-policyname id) is refreshed and upgraded to
// the new key:value id format by the current Framework provider.
func TestAccIcaglobal_icapolicy_binding_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckIcaglobal_icapolicy_bindingDestroy,
		Steps: []resource.TestStep{
			// Step 1: create with the last SDK v2 release. State is written with the legacy id (policyname).
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.2.0",
					},
				},
				Config: testAccIcaglobal_icapolicy_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIcaglobal_icapolicy_bindingExist("citrixadc_icaglobal_icapolicy_binding.tf_icaglobal_icapolicy_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_icaglobal_icapolicy_binding.tf_icaglobal_icapolicy_binding", "id", "tf_icapolicy"),
				),
			},
			// Step 2: refresh the legacy-id state through the current (Framework) provider.
			// Read exercises ParseIdString on the legacy id and recomputes the id to the new format.
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccIcaglobal_icapolicy_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIcaglobal_icapolicy_bindingExist("citrixadc_icaglobal_icapolicy_binding.tf_icaglobal_icapolicy_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_icaglobal_icapolicy_binding.tf_icaglobal_icapolicy_binding", "id", "policyname:tf_icapolicy,type:ICA_REQ_DEFAULT"),
				),
			},
		},
	})
}
