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

const testAccAaauser_tmsessionpolicy_binding_basic = `

	resource "citrixadc_aaauser" "tf_aaauser" {
		username = "user1"
		password = "my_pass"
	}
	resource "citrixadc_tmsessionaction" "tf_tmsessionaction" {
		name                       = "my_tmsession_action"
		sesstimeout                = 10
		defaultauthorizationaction = "ALLOW"
		sso                        = "OFF"
	}
	resource "citrixadc_tmsessionpolicy" "tf_tmsessionpolicy" {
		name   = "my_tmsession_policy"
		rule   = "true"
		action = citrixadc_tmsessionaction.tf_tmsessionaction.name
	}

	resource "citrixadc_aaauser_tmsessionpolicy_binding" "tf_aaauser_tmsessionpolicy_binding" {
		username = citrixadc_aaauser.tf_aaauser.username
		policy    = citrixadc_tmsessionpolicy.tf_tmsessionpolicy.name
		type     = "REQUEST"
		priority  = 100
	}
`

const testAccAaauser_tmsessionpolicy_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion

	resource "citrixadc_aaauser" "tf_aaauser" {
		username = "user1"
		password = "my_pass"
	}
	resource "citrixadc_tmsessionaction" "tf_tmsessionaction" {
		name                       = "my_tmsession_action"
		sesstimeout                = 10
		defaultauthorizationaction = "ALLOW"
		sso                        = "OFF"
	}
	resource "citrixadc_tmsessionpolicy" "tf_tmsessionpolicy" {
		name   = "my_tmsession_policy"
		rule   = "true"
		action = citrixadc_tmsessionaction.tf_tmsessionaction.name
	}
`

func TestAccAaauser_tmsessionpolicy_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAaauser_tmsessionpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAaauser_tmsessionpolicy_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaauser_tmsessionpolicy_bindingExist("citrixadc_aaauser_tmsessionpolicy_binding.tf_aaauser_tmsessionpolicy_binding", nil),
				),
			},
			{
				Config: testAccAaauser_tmsessionpolicy_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaauser_tmsessionpolicy_bindingNotExist("citrixadc_aaauser_tmsessionpolicy_binding.tf_aaauser_tmsessionpolicy_binding", "user1,tf_tmsesspolicy"),
				),
			},
		},
	})
}

func TestAccAaauser_tmsessionpolicy_binding_import(t *testing.T) {
	const resAddr = "citrixadc_aaauser_tmsessionpolicy_binding.tf_aaauser_tmsessionpolicy_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAaauser_tmsessionpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccAaauser_tmsessionpolicy_binding_basic},
			{Config: testAccAaauser_tmsessionpolicy_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{"type"}},
		},
	})
}

func testAccCheckAaauser_tmsessionpolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No aaauser_tmsessionpolicy_binding id is set")
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

		idMap, _, err := utils.ParseIdString(bindingId, []string{"username", "policy"}, nil)
		if err != nil {
			return err
		}
		username := idMap["username"]
		policy := idMap["policy"]

		findParams := service.FindParams{
			ResourceType:             "aaauser_tmsessionpolicy_binding",
			ResourceName:             username,
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
			return fmt.Errorf("aaauser_tmsessionpolicy_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckAaauser_tmsessionpolicy_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}
		idMap, _, err := utils.ParseIdString(id, []string{"username", "policy"}, nil)
		if err != nil {
			return err
		}
		username := idMap["username"]
		policy := idMap["policy"]

		findParams := service.FindParams{
			ResourceType:             "aaauser_tmsessionpolicy_binding",
			ResourceName:             username,
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
			return fmt.Errorf("aaauser_tmsessionpolicy_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckAaauser_tmsessionpolicy_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_aaauser_tmsessionpolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Aaauser_tmsessionpolicy_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("aaauser_tmsessionpolicy_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccAaauser_tmsessionpolicy_bindingDataSource_basic = `

	resource "citrixadc_aaauser" "tf_aaauser" {
		username = "user1"
		password = "my_pass"
	}
	resource "citrixadc_tmsessionaction" "tf_tmsessionaction" {
		name                       = "my_tmsession_action"
		sesstimeout                = 10
		defaultauthorizationaction = "ALLOW"
		sso                        = "OFF"
	}
	resource "citrixadc_tmsessionpolicy" "tf_tmsessionpolicy" {
		name   = "my_tmsession_policy"
		rule   = "true"
		action = citrixadc_tmsessionaction.tf_tmsessionaction.name
	}

	resource "citrixadc_aaauser_tmsessionpolicy_binding" "tf_aaauser_tmsessionpolicy_binding" {
		username = citrixadc_aaauser.tf_aaauser.username
		policy    = citrixadc_tmsessionpolicy.tf_tmsessionpolicy.name
		type     = "REQUEST"
		priority  = 100
	}

	data "citrixadc_aaauser_tmsessionpolicy_binding" "tf_aaauser_tmsessionpolicy_binding" {
		username = citrixadc_aaauser_tmsessionpolicy_binding.tf_aaauser_tmsessionpolicy_binding.username
		policy   = citrixadc_aaauser_tmsessionpolicy_binding.tf_aaauser_tmsessionpolicy_binding.policy
	}
`

func TestAccAaauser_tmsessionpolicy_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAaauser_tmsessionpolicy_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_aaauser_tmsessionpolicy_binding.tf_aaauser_tmsessionpolicy_binding", "username", "user1"),
					resource.TestCheckResourceAttr("data.citrixadc_aaauser_tmsessionpolicy_binding.tf_aaauser_tmsessionpolicy_binding", "policy", "my_tmsession_policy"),
					resource.TestCheckResourceAttr("data.citrixadc_aaauser_tmsessionpolicy_binding.tf_aaauser_tmsessionpolicy_binding", "priority", "100"),
				),
			},
		},
	})
}

const testAccAaauser_tmsessionpolicy_binding_upgrade_basic = `

	resource "citrixadc_aaauser" "tf_aaauser" {
		username = "user1"
		password = "my_pass"
	}
	resource "citrixadc_tmsessionaction" "tf_tmsessionaction" {
		name                       = "my_tmsession_action"
		sesstimeout                = 10
		defaultauthorizationaction = "ALLOW"
		sso                        = "OFF"
	}
	resource "citrixadc_tmsessionpolicy" "tf_tmsessionpolicy" {
		name   = "my_tmsession_policy"
		rule   = "true"
		action = citrixadc_tmsessionaction.tf_tmsessionaction.name
	}

	resource "citrixadc_aaauser_tmsessionpolicy_binding" "tf_aaauser_tmsessionpolicy_binding" {
		username = citrixadc_aaauser.tf_aaauser.username
		policy    = citrixadc_tmsessionpolicy.tf_tmsessionpolicy.name
		type     = "REQUEST"
		priority  = 100
	}
`

func TestAccAaauser_tmsessionpolicy_binding_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckAaauser_tmsessionpolicy_bindingDestroy,
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
				Config: testAccAaauser_tmsessionpolicy_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaauser_tmsessionpolicy_bindingExist("citrixadc_aaauser_tmsessionpolicy_binding.tf_aaauser_tmsessionpolicy_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_aaauser_tmsessionpolicy_binding.tf_aaauser_tmsessionpolicy_binding", "id", "user1,my_tmsession_policy"),
				),
			},
			{
				// Step 2: refresh/plan the legacy-id state through the current
				// framework provider. Read exercises ParseIdString on the legacy id
				// and SetAttrFromGet recomputes the id into the new key:value form.
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccAaauser_tmsessionpolicy_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaauser_tmsessionpolicy_bindingExist("citrixadc_aaauser_tmsessionpolicy_binding.tf_aaauser_tmsessionpolicy_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_aaauser_tmsessionpolicy_binding.tf_aaauser_tmsessionpolicy_binding", "id", "policy:my_tmsession_policy,username:user1"),
				),
			},
		},
	})
}
