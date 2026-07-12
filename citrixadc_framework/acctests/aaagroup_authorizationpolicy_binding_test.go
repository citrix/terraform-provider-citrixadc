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

const testAccAaagroup_authorizationpolicy_binding_basic = `

	resource "citrixadc_aaagroup_authorizationpolicy_binding" "tf_aaagroup_authorizationpolicy_binding" {
		groupname = citrixadc_aaagroup.tf_aaagroup.groupname
		policy   = citrixadc_authorizationpolicy.tf_authorize.name
		type     = "REQUEST"
		priority = 100
	}
	resource "citrixadc_aaagroup" "tf_aaagroup" {
		groupname = "my_group"
		weight    = 100
	}
	resource "citrixadc_authorizationpolicy" "tf_authorize" {
		name   = "tp-authorize-1"
		rule   = "true"
		action = "ALLOW"
	}
`

const testAccAaagroup_authorizationpolicy_binding_basic_step2 = `
	resource "citrixadc_aaagroup" "tf_aaagroup" {
		groupname = "my_group"
		weight    = 100
	}
	resource "citrixadc_authorizationpolicy" "tf_authorize" {
		name   = "tp-authorize-1"
		rule   = "true"
		action = "ALLOW"
	}
`

const testAccAaagroupAuthorizationpolicyBindingDataSource_basic = `

	resource "citrixadc_aaagroup_authorizationpolicy_binding" "tf_aaagroup_authorizationpolicy_binding" {
		groupname = citrixadc_aaagroup.tf_aaagroup.groupname
		policy   = citrixadc_authorizationpolicy.tf_authorize.name
		type     = "REQUEST"
		priority = 100
	}
	resource "citrixadc_aaagroup" "tf_aaagroup" {
		groupname = "my_group"
		weight    = 100
	}
	resource "citrixadc_authorizationpolicy" "tf_authorize" {
		name   = "tp-authorize-1"
		rule   = "true"
		action = "ALLOW"
	}

	data "citrixadc_aaagroup_authorizationpolicy_binding" "tf_aaagroup_authorizationpolicy_binding" {
		groupname = citrixadc_aaagroup_authorizationpolicy_binding.tf_aaagroup_authorizationpolicy_binding.groupname
		policy    = citrixadc_aaagroup_authorizationpolicy_binding.tf_aaagroup_authorizationpolicy_binding.policy
		depends_on = [citrixadc_aaagroup_authorizationpolicy_binding.tf_aaagroup_authorizationpolicy_binding]
	}
`

func TestAccAaagroup_authorizationpolicy_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAaagroup_authorizationpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAaagroup_authorizationpolicy_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaagroup_authorizationpolicy_bindingExist("citrixadc_aaagroup_authorizationpolicy_binding.tf_aaagroup_authorizationpolicy_binding", nil),
				),
			},
			{
				Config: testAccAaagroup_authorizationpolicy_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaagroup_authorizationpolicy_bindingNotExist("citrixadc_aaagroup_authorizationpolicy_binding.tf_aaagroup_authorizationpolicy_binding", "my_group,tp-authorize-1"),
				),
			},
		},
	})
}

func TestAccAaagroupAuthorizationpolicyBindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccAaagroupAuthorizationpolicyBindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_aaagroup_authorizationpolicy_binding.tf_aaagroup_authorizationpolicy_binding", "groupname", "my_group"),
					resource.TestCheckResourceAttr("data.citrixadc_aaagroup_authorizationpolicy_binding.tf_aaagroup_authorizationpolicy_binding", "policy", "tp-authorize-1"),
					resource.TestCheckResourceAttr("data.citrixadc_aaagroup_authorizationpolicy_binding.tf_aaagroup_authorizationpolicy_binding", "priority", "100"),
				),
			},
		},
	})
}

func TestAccAaagroup_authorizationpolicy_binding_import(t *testing.T) {
	const resAddr = "citrixadc_aaagroup_authorizationpolicy_binding.tf_aaagroup_authorizationpolicy_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAaagroup_authorizationpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccAaagroup_authorizationpolicy_binding_basic},
			{Config: testAccAaagroup_authorizationpolicy_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{"type"}},
		},
	})
}

func testAccCheckAaagroup_authorizationpolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No aaagroup_authorizationpolicy_binding id is set")
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
			return fmt.Errorf("Error parsing ID %s: %v", bindingId, err)
		}
		groupname := idMap["groupname"]
		policy := idMap["policy"]

		findParams := service.FindParams{
			ResourceType:             "aaagroup_authorizationpolicy_binding",
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
			return fmt.Errorf("aaagroup_authorizationpolicy_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckAaagroup_authorizationpolicy_bindingNotExist(n string, id string) resource.TestCheckFunc {
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
			ResourceType:             "aaagroup_authorizationpolicy_binding",
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
			return fmt.Errorf("aaagroup_authorizationpolicy_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckAaagroup_authorizationpolicy_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_aaagroup_authorizationpolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Aaagroup_authorizationpolicy_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("aaagroup_authorizationpolicy_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccAaagroup_authorizationpolicy_binding_upgrade_basic = `
	resource "citrixadc_aaagroup" "tf_aaagroup" {
		groupname = "my_group"
		weight    = 100
	}
	resource "citrixadc_authorizationpolicy" "tf_authorize" {
		name   = "tp-authorize-1"
		rule   = "true"
		action = "ALLOW"
	}

	resource "citrixadc_aaagroup_authorizationpolicy_binding" "tf_aaagroup_authorizationpolicy_binding" {
		groupname = citrixadc_aaagroup.tf_aaagroup.groupname
		policy   = citrixadc_authorizationpolicy.tf_authorize.name
		type     = "REQUEST"
		priority = 100
	}
`

func TestAccAaagroup_authorizationpolicy_binding_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckAaagroup_authorizationpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				// Step 1: create the binding with the last SDK v2 release (2.2.0),
				// which writes state using the legacy comma-joined id.
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.2.0",
					},
				},
				Config: testAccAaagroup_authorizationpolicy_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaagroup_authorizationpolicy_bindingExist("citrixadc_aaagroup_authorizationpolicy_binding.tf_aaagroup_authorizationpolicy_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_aaagroup_authorizationpolicy_binding.tf_aaagroup_authorizationpolicy_binding", "id", "my_group,tp-authorize-1"),
				),
			},
			{
				// Step 2: refresh/plan the legacy-id state through the current
				// framework provider. Read exercises ParseIdString on the legacy id
				// and SetAttrFromGet recomputes the id into the new key:value form.
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccAaagroup_authorizationpolicy_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaagroup_authorizationpolicy_bindingExist("citrixadc_aaagroup_authorizationpolicy_binding.tf_aaagroup_authorizationpolicy_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_aaagroup_authorizationpolicy_binding.tf_aaagroup_authorizationpolicy_binding", "id", "groupname:my_group,policy:tp-authorize-1"),
				),
			},
		},
	})
}
