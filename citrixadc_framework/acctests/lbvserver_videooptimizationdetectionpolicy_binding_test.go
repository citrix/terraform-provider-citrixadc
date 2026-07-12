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

const testAccLbvserver_videooptimizationdetectionpolicy_binding_basic = `

	resource "citrixadc_videooptimizationdetectionpolicy" "tf_detectionpolicy" {
		name   = "tf_vop"
		rule   = "true"
		action = "DETECT_ENCRYPTED_ABR"
	}

	resource "citrixadc_lbvserver_videooptimizationdetectionpolicy_binding" "tf_vopolicy_binding" {
        bindpoint = "REQUEST"
        gotopriorityexpression = "END"
        name = citrixadc_lbvserver.tf_lbvserver.name
        policyname = citrixadc_videooptimizationdetectionpolicy.tf_detectionpolicy.name
        priority = 1
	}

	resource "citrixadc_lbvserver" "tf_lbvserver" {
		name        = "tf_lbvserver"
		ipv46       = "10.10.10.33"
		port        = 80
		servicetype = "HTTP"
	}
`

const testAccLbvserver_videooptimizationdetectionpolicy_binding_basic_step2 = `

	resource "citrixadc_videooptimizationdetectionpolicy" "tf_detectionpolicy" {
		name   = "tf_vop"
		rule   = "true"
		action = "DETECT_ENCRYPTED_ABR"
	}
	resource "citrixadc_lbvserver" "tf_lbvserver" {
		name        = "tf_lbvserver"
		ipv46       = "10.10.10.33"
		port        = 80
		servicetype = "HTTP"
	}
`

func TestAccLbvserver_videooptimizationdetectionpolicy_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLbvserver_videooptimizationdetectionpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLbvserver_videooptimizationdetectionpolicy_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbvserver_videooptimizationdetectionpolicy_bindingExist("citrixadc_lbvserver_videooptimizationdetectionpolicy_binding.tf_vopolicy_binding", nil),
				),
			},
			{
				Config: testAccLbvserver_videooptimizationdetectionpolicy_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbvserver_videooptimizationdetectionpolicy_bindingNotExist("citrixadc_lbvserver_videooptimizationdetectionpolicy_binding.tf_lbvserver_videooptimizationdetectionpolicy_binding", "tf_lbvserver,tf_vop"),
				),
			},
		},
	})
}

func TestAccLbvserver_videooptimizationdetectionpolicy_binding_import(t *testing.T) {
	const resAddr = "citrixadc_lbvserver_videooptimizationdetectionpolicy_binding.tf_vopolicy_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLbvserver_videooptimizationdetectionpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccLbvserver_videooptimizationdetectionpolicy_binding_basic},
			{Config: testAccLbvserver_videooptimizationdetectionpolicy_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
		},
	})
}

func testAccCheckLbvserver_videooptimizationdetectionpolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lbvserver_videooptimizationdetectionpolicy_binding id is set")
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

		idMap, _, err := utils.ParseIdString(bindingId, []string{"name", "policyname"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID: %v", err)
		}
		lbvserverName := idMap["name"]
		policyName := idMap["policyname"]

		findParams := service.FindParams{
			ResourceType:             "lbvserver_videooptimizationdetectionpolicy_binding",
			ResourceName:             lbvserverName,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the right monitor name
		found := false
		for _, v := range dataArr {
			if v["policyname"].(string) == policyName {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("lbvserver_videooptimizationdetectionpolicy_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckLbvserver_videooptimizationdetectionpolicy_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		idMap, _, err := utils.ParseIdString(id, []string{"name", "policyname"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID: %v", err)
		}
		lbvserverName := idMap["name"]
		policyName := idMap["policyname"]

		findParams := service.FindParams{
			ResourceType:             "lbvserver_videooptimizationdetectionpolicy_binding",
			ResourceName:             lbvserverName,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the right policy name
		found := false
		for _, v := range dataArr {
			if v["policyname"].(string) == policyName {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("lbvserver_videooptimizationdetectionpolicy_binding %s was found, but is should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckLbvserver_videooptimizationdetectionpolicy_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_lbvserver_videooptimizationdetectionpolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("lbvserver_videooptimizationdetectionpolicy_binding", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("lbvserver_videooptimizationdetectionpolicy_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccLbvserver_videooptimizationdetectionpolicy_bindingDataSource_basic = `

	resource "citrixadc_videooptimizationdetectionpolicy" "tf_detectionpolicy" {
		name   = "tf_vop"
		rule   = "true"
		action = "DETECT_ENCRYPTED_ABR"
	}

	resource "citrixadc_lbvserver_videooptimizationdetectionpolicy_binding" "tf_vopolicy_binding" {
        bindpoint = "REQUEST"
        gotopriorityexpression = "END"
        name = citrixadc_lbvserver.tf_lbvserver.name
        policyname = citrixadc_videooptimizationdetectionpolicy.tf_detectionpolicy.name
        priority = 1
	}

	resource "citrixadc_lbvserver" "tf_lbvserver" {
		name        = "tf_lbvserver"
		ipv46       = "10.10.10.33"
		port        = 80
		servicetype = "HTTP"
	}

	data "citrixadc_lbvserver_videooptimizationdetectionpolicy_binding" "tf_vopolicy_binding" {
		name = citrixadc_lbvserver_videooptimizationdetectionpolicy_binding.tf_vopolicy_binding.name
		policyname = citrixadc_lbvserver_videooptimizationdetectionpolicy_binding.tf_vopolicy_binding.policyname
		bindpoint = citrixadc_lbvserver_videooptimizationdetectionpolicy_binding.tf_vopolicy_binding.bindpoint
		depends_on = [citrixadc_lbvserver_videooptimizationdetectionpolicy_binding.tf_vopolicy_binding]
	}
`

func TestAccLbvserver_videooptimizationdetectionpolicy_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccLbvserver_videooptimizationdetectionpolicy_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_lbvserver_videooptimizationdetectionpolicy_binding.tf_vopolicy_binding", "name", "tf_lbvserver"),
					resource.TestCheckResourceAttr("data.citrixadc_lbvserver_videooptimizationdetectionpolicy_binding.tf_vopolicy_binding", "policyname", "tf_vop"),
					resource.TestCheckResourceAttr("data.citrixadc_lbvserver_videooptimizationdetectionpolicy_binding.tf_vopolicy_binding", "priority", "1"),
					resource.TestCheckResourceAttr("data.citrixadc_lbvserver_videooptimizationdetectionpolicy_binding.tf_vopolicy_binding", "gotopriorityexpression", "END"),
					resource.TestCheckResourceAttr("data.citrixadc_lbvserver_videooptimizationdetectionpolicy_binding.tf_vopolicy_binding", "bindpoint", "REQUEST"),
				),
			},
		},
	})
}

// testAccLbvserver_videooptimizationdetectionpolicy_binding_upgrade_basic is used by the
// sdkv2 -> Framework state-upgrade test. It reuses the same config values as the _basic
// test and MUST be valid under BOTH the SDK v2 2.2.0 schema (step 1) and the current
// Framework schema (step 2). The resource label (tf_vopolicy_binding) is kept identical so
// the Exist/Destroy helpers and resource addresses match.
const testAccLbvserver_videooptimizationdetectionpolicy_binding_upgrade_basic = `

	resource "citrixadc_videooptimizationdetectionpolicy" "tf_detectionpolicy" {
		name   = "tf_vop"
		rule   = "true"
		action = "DETECT_ENCRYPTED_ABR"
	}

	resource "citrixadc_lbvserver_videooptimizationdetectionpolicy_binding" "tf_vopolicy_binding" {
        bindpoint = "REQUEST"
        gotopriorityexpression = "END"
        name = citrixadc_lbvserver.tf_lbvserver.name
        policyname = citrixadc_videooptimizationdetectionpolicy.tf_detectionpolicy.name
        priority = 1
	}

	resource "citrixadc_lbvserver" "tf_lbvserver" {
		name        = "tf_lbvserver"
		ipv46       = "10.10.10.33"
		port        = 80
		servicetype = "HTTP"
	}
`

// TestAccLbvserver_videooptimizationdetectionpolicy_binding_sdkv2StateUpgrade verifies that a
// resource created with the last SDK v2 release (2.2.0), which writes the legacy comma-joined
// ID (name,policyname), upgrades cleanly when refreshed/planned through the current Framework
// provider. On the Framework Read the ID is recomputed into the new key:value format (see
// lbvserver_videooptimizationdetectionpolicy_bindingSetAttrFromGet in resource_schema.go).
func TestAccLbvserver_videooptimizationdetectionpolicy_binding_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckLbvserver_videooptimizationdetectionpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			// Step 1: create with the last SDK v2 release; state holds the legacy ID.
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.2.0",
					},
				},
				Config: testAccLbvserver_videooptimizationdetectionpolicy_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbvserver_videooptimizationdetectionpolicy_bindingExist("citrixadc_lbvserver_videooptimizationdetectionpolicy_binding.tf_vopolicy_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_lbvserver_videooptimizationdetectionpolicy_binding.tf_vopolicy_binding", "id", "tf_lbvserver,tf_vop"),
				),
			},
			// Step 2: same config through the current Framework provider; Read recomputes
			// the legacy ID into the new key:value format.
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccLbvserver_videooptimizationdetectionpolicy_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbvserver_videooptimizationdetectionpolicy_bindingExist("citrixadc_lbvserver_videooptimizationdetectionpolicy_binding.tf_vopolicy_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_lbvserver_videooptimizationdetectionpolicy_binding.tf_vopolicy_binding", "id", "bindpoint:REQUEST,name:tf_lbvserver,policyname:tf_vop"),
				),
			},
		},
	})
}
