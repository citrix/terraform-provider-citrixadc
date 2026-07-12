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

	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	//"strings"
	"testing"
)

const testAccBotglobal_botpolicy_binding_basic = `
	resource "citrixadc_botpolicy" "tf_botpolicy" {
		name        = "tf_botpolicy"
		profilename = "BOT_BYPASS"
		rule        = "true"
		comment     = "COMMENT FOR BOTPOLICY"
	}
	resource "citrixadc_botglobal_botpolicy_binding" "tf_binding" {
		policyname = citrixadc_botpolicy.tf_botpolicy.name
		priority   = 90
		type       = "REQ_OVERRIDE"
	}
`

const testAccBotglobal_botpolicy_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion
	resource "citrixadc_botpolicy" "tf_botpolicy" {
		name        = "tf_botpolicy"
		profilename = "BOT_BYPASS"
		rule        = "true"
		comment     = "COMMENT FOR BOTPOLICY"
	}
`

func TestAccBotglobal_botpolicy_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckBotglobal_botpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccBotglobal_botpolicy_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBotglobal_botpolicy_bindingExist("citrixadc_botglobal_botpolicy_binding.tf_binding", nil),
				),
			},
			{
				Config: testAccBotglobal_botpolicy_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBotglobal_botpolicy_bindingNotExist("citrixadc_botglobal_botpolicy_binding.tf_binding", "tf_botpolicy", "REQ_OVERRIDE"),
				),
			},
		},
	})
}

func TestAccBotglobal_botpolicy_binding_import(t *testing.T) {
	const resAddr = "citrixadc_botglobal_botpolicy_binding.tf_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckBotglobal_botpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccBotglobal_botpolicy_binding_basic},
			{Config: testAccBotglobal_botpolicy_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
		},
	})
}

func testAccCheckBotglobal_botpolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No botglobal_botpolicy_binding id is set")
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
		typename := rs.Primary.Attributes["type"]

		findParams := service.FindParams{
			ResourceType:             "botglobal_botpolicy_binding",
			ArgsMap:                  map[string]string{"type": typename},
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
			return fmt.Errorf("botglobal_botpolicy_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckBotglobal_botpolicy_bindingNotExist(n string, id string, typename string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		policyname := id

		findParams := service.FindParams{
			ResourceType:             "botglobal_botpolicy_binding",
			ArgsMap:                  map[string]string{"type": typename},
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
			return fmt.Errorf("botglobal_botpolicy_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckBotglobal_botpolicy_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_botglobal_botpolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("botglobal_botpolicy_binding", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("botglobal_botpolicy_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccBotglobalBotpolicyBindingDataSource_basic = `
	resource "citrixadc_botpolicy" "tf_botpolicy" {
		name        = "tf_botpolicy"
		profilename = "BOT_BYPASS"
		rule        = "true"
		comment     = "COMMENT FOR BOTPOLICY"
	}
	resource "citrixadc_botglobal_botpolicy_binding" "tf_binding" {
		policyname = citrixadc_botpolicy.tf_botpolicy.name
		priority   = 90
		type       = "REQ_OVERRIDE"
	}

	data "citrixadc_botglobal_botpolicy_binding" "tf_binding" {
		policyname = citrixadc_botglobal_botpolicy_binding.tf_binding.policyname
		type       = citrixadc_botglobal_botpolicy_binding.tf_binding.type
		depends_on = [citrixadc_botglobal_botpolicy_binding.tf_binding]
	}
`

func TestAccBotglobalBotpolicyBindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccBotglobalBotpolicyBindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_botglobal_botpolicy_binding.tf_binding", "policyname", "tf_botpolicy"),
					resource.TestCheckResourceAttr("data.citrixadc_botglobal_botpolicy_binding.tf_binding", "priority", "90"),
					resource.TestCheckResourceAttr("data.citrixadc_botglobal_botpolicy_binding.tf_binding", "type", "REQ_OVERRIDE"),
				),
			},
		},
	})
}

const testAccBotglobal_botpolicy_binding_upgrade_basic = `
	resource "citrixadc_botpolicy" "tf_botpolicy" {
		name        = "tf_botpolicy"
		profilename = "BOT_BYPASS"
		rule        = "true"
		comment     = "COMMENT FOR BOTPOLICY"
	}
	resource "citrixadc_botglobal_botpolicy_binding" "tf_binding" {
		policyname = citrixadc_botpolicy.tf_botpolicy.name
		priority   = 90
		type       = "REQ_OVERRIDE"
	}
`

func TestAccBotglobal_botpolicy_binding_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckBotglobal_botpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				// Step 1: create the binding with the last SDK v2 release (2.2.0),
				// which writes state using the legacy id (policyname).
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.2.0",
					},
				},
				Config: testAccBotglobal_botpolicy_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBotglobal_botpolicy_bindingExist("citrixadc_botglobal_botpolicy_binding.tf_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_botglobal_botpolicy_binding.tf_binding", "id", "tf_botpolicy"),
				),
			},
			{
				// Step 2: refresh/plan the legacy-id state through the current
				// framework provider. Read exercises ParseIdString on the legacy id
				// and SetAttrFromGet recomputes the id into the new key:value form.
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccBotglobal_botpolicy_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBotglobal_botpolicy_bindingExist("citrixadc_botglobal_botpolicy_binding.tf_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_botglobal_botpolicy_binding.tf_binding", "id", "policyname:tf_botpolicy,type:REQ_OVERRIDE"),
				),
			},
		},
	})
}
