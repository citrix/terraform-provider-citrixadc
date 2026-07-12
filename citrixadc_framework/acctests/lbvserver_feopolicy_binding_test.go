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

const testAccLbvserver_feopolicy_binding_basic = `

	resource "citrixadc_feopolicy" "tf_feopolicy" {
		name   = "tf_feopolicy"
		action = "BASIC"
		rule   = "true"
	}

	resource "citrixadc_lbvserver_feopolicy_binding" "tf_lbvserver_feopolicy_binding" {
		bindpoint 				= "REQUEST"
        gotopriorityexpression 	= "END"
        name 					= citrixadc_lbvserver.tf_lbvserver.name
        policyname 				= citrixadc_feopolicy.tf_feopolicy.name
        priority 				= 1  
	}

	resource "citrixadc_lbvserver" "tf_lbvserver" {
		name        = "tf_lbvserver"
		ipv46       = "10.10.10.33"
		port        = 80
		servicetype = "HTTP"
	}
`

const testAccLbvserver_feopolicy_binding_basic_step2 = `

	resource "citrixadc_feopolicy" "tf_feopolicy" {
		name   = "tf_feopolicy"
		action = "BASIC"
		rule   = "true"
	}
	resource "citrixadc_lbvserver" "tf_lbvserver" {
		name        = "tf_lbvserver"
		ipv46       = "10.10.10.33"
		port        = 80
		servicetype = "HTTP"
	}
`

func TestAccLbvserver_feopolicy_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLbvserver_feopolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLbvserver_feopolicy_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbvserver_feopolicy_bindingExist("citrixadc_lbvserver_feopolicy_binding.tf_lbvserver_feopolicy_binding", nil),
				),
			},
			{
				Config: testAccLbvserver_feopolicy_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbvserver_feopolicy_bindingNotExist("citrixadc_lbvserver_feopolicy_binding.tf_lbvserver_feopolicy_binding", "tf_lbvserver,tf_feopolicy"),
				),
			},
		},
	})
}

func TestAccLbvserver_feopolicy_binding_import(t *testing.T) {
	const resAddr = "citrixadc_lbvserver_feopolicy_binding.tf_lbvserver_feopolicy_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLbvserver_feopolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccLbvserver_feopolicy_binding_basic},
			{Config: testAccLbvserver_feopolicy_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
		},
	})
}

func testAccCheckLbvserver_feopolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lbvserver_feopolicy_binding id is set")
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
			ResourceType:             "lbvserver_feopolicy_binding",
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
			return fmt.Errorf("lbvserver_feopolicy_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckLbvserver_feopolicy_bindingNotExist(n string, id string) resource.TestCheckFunc {
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
			ResourceType:             "lbvserver_feopolicy_binding",
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
			return fmt.Errorf("lbvserver_feopolicy_binding %s was found, but is should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckLbvserver_feopolicy_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_lbvserver_feopolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Lbvserver_feopolicy_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("lbvserver_feopolicy_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccLbvserver_feopolicy_bindingDataSource_basic = `
resource "citrixadc_lbvserver" "tf_lbvserver" {
	name        = "tf_lbvserver"
	ipv46       = "10.10.10.33"
	port        = 80
	servicetype = "HTTP"
}

resource "citrixadc_feopolicy" "tf_feopolicy" {
	name   = "tf_feopolicy"
	action = "BASIC"
	rule   = "true"
}

resource "citrixadc_lbvserver_feopolicy_binding" "tf_lbvserver_feopolicy_binding" {
	bindpoint              = "REQUEST"
	gotopriorityexpression = "END"
	name                   = citrixadc_lbvserver.tf_lbvserver.name
	policyname             = citrixadc_feopolicy.tf_feopolicy.name
	priority               = 1
}

data "citrixadc_lbvserver_feopolicy_binding" "tf_lbvserver_feopolicy_binding" {
	name       = citrixadc_lbvserver_feopolicy_binding.tf_lbvserver_feopolicy_binding.name
	policyname = citrixadc_lbvserver_feopolicy_binding.tf_lbvserver_feopolicy_binding.policyname
	depends_on = [citrixadc_lbvserver_feopolicy_binding.tf_lbvserver_feopolicy_binding]
}
`

func TestAccLbvserver_feopolicy_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccLbvserver_feopolicy_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_lbvserver_feopolicy_binding.tf_lbvserver_feopolicy_binding", "name", "tf_lbvserver"),
					resource.TestCheckResourceAttr("data.citrixadc_lbvserver_feopolicy_binding.tf_lbvserver_feopolicy_binding", "policyname", "tf_feopolicy"),
					resource.TestCheckResourceAttr("data.citrixadc_lbvserver_feopolicy_binding.tf_lbvserver_feopolicy_binding", "priority", "1"),
					resource.TestCheckResourceAttr("data.citrixadc_lbvserver_feopolicy_binding.tf_lbvserver_feopolicy_binding", "gotopriorityexpression", "END"),
				),
			},
		},
	})
}

// testAccLbvserver_feopolicy_binding_upgrade_basic is the config used by the
// sdkv2 -> framework state-upgrade test. It reuses the same values and resource
// labels as testAccLbvserver_feopolicy_binding_basic so it is valid under BOTH
// the SDK v2 2.2.0 schema and the current framework schema.
const testAccLbvserver_feopolicy_binding_upgrade_basic = `

	resource "citrixadc_feopolicy" "tf_feopolicy" {
		name   = "tf_feopolicy"
		action = "BASIC"
		rule   = "true"
	}

	resource "citrixadc_lbvserver" "tf_lbvserver" {
		name        = "tf_lbvserver"
		ipv46       = "10.10.10.33"
		port        = 80
		servicetype = "HTTP"
	}

	resource "citrixadc_lbvserver_feopolicy_binding" "tf_lbvserver_feopolicy_binding" {
		bindpoint              = "REQUEST"
		gotopriorityexpression = "END"
		name                   = citrixadc_lbvserver.tf_lbvserver.name
		policyname             = citrixadc_feopolicy.tf_feopolicy.name
		priority               = 1
	}
`

// TestAccLbvserver_feopolicy_binding_sdkv2StateUpgrade verifies that a binding
// created with the last SDK v2 release (2.2.0, legacy comma-separated ID) is
// correctly refreshed/planned/applied by the current framework provider.
func TestAccLbvserver_feopolicy_binding_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckLbvserver_feopolicy_bindingDestroy,
		Steps: []resource.TestStep{
			// Step 1: create the binding with the last SDK v2 release.
			// State is written with the LEGACY comma-separated id.
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.2.0",
					},
				},
				Config: testAccLbvserver_feopolicy_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbvserver_feopolicy_bindingExist("citrixadc_lbvserver_feopolicy_binding.tf_lbvserver_feopolicy_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_lbvserver_feopolicy_binding.tf_lbvserver_feopolicy_binding", "id", "tf_lbvserver,tf_feopolicy"),
				),
			},
			// Step 2: same config, current (framework) provider. Terraform
			// refreshes the legacy-id state through the framework Read
			// (exercising ParseIdString on the legacy id) then plans/applies.
			// The framework recomputes the id on read to the new key:value form.
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccLbvserver_feopolicy_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbvserver_feopolicy_bindingExist("citrixadc_lbvserver_feopolicy_binding.tf_lbvserver_feopolicy_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_lbvserver_feopolicy_binding.tf_lbvserver_feopolicy_binding", "id", "name:tf_lbvserver,policyname:tf_feopolicy"),
				),
			},
		},
	})
}
