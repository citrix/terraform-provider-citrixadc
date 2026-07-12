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

const testAccLbvserver_cachepolicy_binding_basic = `
	resource "citrixadc_cachepolicy" "tf_cachepolicy" {
		policyname  = "tf_cachepolicy"
		rule        = "true"
		action      = "CACHE"
	}
	resource "citrixadc_lbvserver_cachepolicy_binding" "tf_citrixadc_lbvserver_cachepolicy_binding" {
		name 		= citrixadc_lbvserver.tf_lbvserver.name
		policyname 	= citrixadc_cachepolicy.tf_cachepolicy.policyname
		priority 	= 1
		bindpoint 	= "REQUEST"
	}

	resource "citrixadc_lbvserver" "tf_lbvserver" {
		name        = "tf_lbvserver"
		ipv46       = "10.10.10.33"
		port        = 80
		servicetype = "HTTP"
	}
`

const testAccLbvserver_cachepolicy_binding_basic_step2 = `

	resource "citrixadc_cachepolicy" "tf_cachepolicy" {
		policyname  = "tf_cachepolicy"
		rule        = "true"
		action      = "CACHE"
	}
	resource "citrixadc_lbvserver" "tf_lbvserver" {
		name        = "tf_lbvserver"
		ipv46       = "10.10.10.33"
		port        = 80
		servicetype = "HTTP"
	}
`

func TestAccLbvserver_cachepolicy_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLbvserver_cachepolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLbvserver_cachepolicy_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbvserver_cachepolicy_bindingExist("citrixadc_lbvserver_cachepolicy_binding.tf_citrixadc_lbvserver_cachepolicy_binding", nil),
				),
			},
			{
				Config: testAccLbvserver_cachepolicy_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbvserver_cachepolicy_bindingNotExist("citrixadc_lbvserver_cachepolicy_binding.tf_lbvserver_cachepolicy_binding", "tf_lbvserver,tf_cachepolicy"),
				),
			},
		},
	})
}

func testAccCheckLbvserver_cachepolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lbvserver_cachepolicy_binding id is set")
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
			ResourceType:             "lbvserver_cachepolicy_binding",
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
			return fmt.Errorf("lbvserver_cachepolicy_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckLbvserver_cachepolicy_bindingNotExist(n string, id string) resource.TestCheckFunc {
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

		lbvserverName := idSlice[0]
		policyName := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "lbvserver_cachepolicy_binding",
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
			return fmt.Errorf("lbvserver_cachepolicy_binding %s was found, but is should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckLbvserver_cachepolicy_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_lbvserver_cachepolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Lbvserver_cachepolicy_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("lbvserver_cachepolicy_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

// testAccLbvserver_cachepolicy_binding_upgrade_basic reuses the _basic config
// (binding + all prerequisite resources). It is valid under BOTH the SDK v2 2.2.0
// schema and the current Framework schema because the migration restored the SDK v2
// attribute names. The terraform resource label matches the one in _basic so the
// Exist/Destroy helpers and addresses line up.
const testAccLbvserver_cachepolicy_binding_upgrade_basic = `
	resource "citrixadc_cachepolicy" "tf_cachepolicy" {
		policyname  = "tf_cachepolicy"
		rule        = "true"
		action      = "CACHE"
	}
	resource "citrixadc_lbvserver_cachepolicy_binding" "tf_citrixadc_lbvserver_cachepolicy_binding" {
		name 		= citrixadc_lbvserver.tf_lbvserver.name
		policyname 	= citrixadc_cachepolicy.tf_cachepolicy.policyname
		priority 	= 1
		bindpoint 	= "REQUEST"
	}

	resource "citrixadc_lbvserver" "tf_lbvserver" {
		name        = "tf_lbvserver"
		ipv46       = "10.10.10.33"
		port        = 80
		servicetype = "HTTP"
	}
`

// TestAccLbvserver_cachepolicy_binding_sdkv2StateUpgrade verifies that state written
// by the last SDK v2 release (legacy comma-separated ID "tf_lbvserver,tf_cachepolicy")
// is correctly upgraded when the same config is subsequently managed by the current
// Framework provider. Step 1 creates the binding with citrix/citrixadc 2.2.0. Step 2
// refreshes/plans/applies the same config through the Framework provider, exercising
// ParseIdString on the legacy id; because the Framework recomputes the id on Read
// (SetAttrFromGet), the id upgrades to the new "key:value" form.
func TestAccLbvserver_cachepolicy_binding_sdkv2StateUpgrade(t *testing.T) {
	resourceAddr := "citrixadc_lbvserver_cachepolicy_binding.tf_citrixadc_lbvserver_cachepolicy_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckLbvserver_cachepolicy_bindingDestroy,
		Steps: []resource.TestStep{
			// Step 1: create with the last SDK v2 release -> state carries the legacy id.
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.2.0",
					},
				},
				Config: testAccLbvserver_cachepolicy_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbvserver_cachepolicy_bindingExist(resourceAddr, nil),
					resource.TestCheckResourceAttr(resourceAddr, "id", "tf_lbvserver,tf_cachepolicy"),
				),
			},
			// Step 2: refresh/plan/apply the SAME config through the current Framework
			// provider. The legacy-id state is read via ParseIdString and the id is
			// recomputed to the new key:value format.
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccLbvserver_cachepolicy_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbvserver_cachepolicy_bindingExist(resourceAddr, nil),
					resource.TestCheckResourceAttr(resourceAddr, "id", "bindpoint:REQUEST,name:tf_lbvserver,policyname:tf_cachepolicy"),
				),
			},
		},
	})
}

const testAccLbvserver_cachepolicy_bindingDataSource_basic = `
	resource "citrixadc_cachepolicy" "tf_cachepolicy" {
		policyname  = "tf_cachepolicy"
		rule        = "true"
		action      = "CACHE"
	}
	resource "citrixadc_lbvserver" "tf_lbvserver" {
		name        = "tf_lbvserver"
		ipv46       = "10.10.10.33"
		port        = 80
		servicetype = "HTTP"
	}
	resource "citrixadc_lbvserver_cachepolicy_binding" "tf_citrixadc_lbvserver_cachepolicy_binding" {
		name 		= citrixadc_lbvserver.tf_lbvserver.name
		policyname 	= citrixadc_cachepolicy.tf_cachepolicy.policyname
		priority 	= 1
		bindpoint 	= "REQUEST"
	}

	data "citrixadc_lbvserver_cachepolicy_binding" "tf_citrixadc_lbvserver_cachepolicy_binding" {
		name = citrixadc_lbvserver_cachepolicy_binding.tf_citrixadc_lbvserver_cachepolicy_binding.name
		policyname = citrixadc_lbvserver_cachepolicy_binding.tf_citrixadc_lbvserver_cachepolicy_binding.policyname
		bindpoint = citrixadc_lbvserver_cachepolicy_binding.tf_citrixadc_lbvserver_cachepolicy_binding.bindpoint
		depends_on = [citrixadc_lbvserver_cachepolicy_binding.tf_citrixadc_lbvserver_cachepolicy_binding]
	}
`

func TestAccLbvserver_cachepolicy_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccLbvserver_cachepolicy_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_lbvserver_cachepolicy_binding.tf_citrixadc_lbvserver_cachepolicy_binding", "name", "tf_lbvserver"),
					resource.TestCheckResourceAttr("data.citrixadc_lbvserver_cachepolicy_binding.tf_citrixadc_lbvserver_cachepolicy_binding", "policyname", "tf_cachepolicy"),
					resource.TestCheckResourceAttr("data.citrixadc_lbvserver_cachepolicy_binding.tf_citrixadc_lbvserver_cachepolicy_binding", "bindpoint", "REQUEST"),
					resource.TestCheckResourceAttr("data.citrixadc_lbvserver_cachepolicy_binding.tf_citrixadc_lbvserver_cachepolicy_binding", "priority", "1"),
				),
			},
		},
	})
}

func TestAccLbvserver_cachepolicy_binding_import(t *testing.T) {
	const resAddr = "citrixadc_lbvserver_cachepolicy_binding.tf_citrixadc_lbvserver_cachepolicy_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLbvserver_cachepolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccLbvserver_cachepolicy_binding_basic},
			{Config: testAccLbvserver_cachepolicy_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
		},
	})
}
