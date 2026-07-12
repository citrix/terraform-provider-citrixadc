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

const testAccLbvserver_contentinspectionpolicy_binding_basic = `

	resource "citrixadc_contentinspectionpolicy" "tf_contentinspectionpolicy" {
		name   = "tf_contentinspectionpolicy"
		rule   = "false"
		action = "DROP"
	}

	resource "citrixadc_lbvserver_contentinspectionpolicy_binding" "tf_lbvserver_contentinspectionpolicy_binding" {
		bindpoint 				= "REQUEST"
		gotopriorityexpression 	= "END"
		name 					= citrixadc_lbvserver.tf_lbvserver.name
		policyname 				= citrixadc_contentinspectionpolicy.tf_contentinspectionpolicy.name
		priority 				= 1    
	}

	resource "citrixadc_lbvserver" "tf_lbvserver" {
		name        = "tf_lbvserver"
		ipv46       = "10.10.10.33"
		port        = 80
		servicetype = "HTTP"
	}
`

const testAccLbvserver_contentispectionpolicy_binding_basic_step2 = `

	resource "citrixadc_contentinspectionpolicy" "tf_contentinspectionpolicy" {
		name   = "tf_contentinspectionpolicy"
		rule   = "false"
		action = "DROP"
	}
	resource "citrixadc_lbvserver" "tf_lbvserver" {
		name        = "tf_lbvserver"
		ipv46       = "10.10.10.33"
		port        = 80
		servicetype = "HTTP"
	}
`

func TestAccLbvserver_contentinspectionpolicy_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLbvserver_contentinspectionpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLbvserver_contentinspectionpolicy_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbvserver_contentinspectionpolicy_bindingExist("citrixadc_lbvserver_contentinspectionpolicy_binding.tf_lbvserver_contentinspectionpolicy_binding", nil),
				),
			},
			{
				Config: testAccLbvserver_contentispectionpolicy_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbvserver_contentispectionpolicy_bindingNotExist("citrixadc_lbvserver_contentinspectionpolicy_binding.tf_lbvserver_contentinspectionpolicy_binding", "tf_lbvserver,tf_contentinspectionpolicy"),
				),
			},
		},
	})
}

func TestAccLbvserver_contentinspectionpolicy_binding_import(t *testing.T) {
	const resAddr = "citrixadc_lbvserver_contentinspectionpolicy_binding.tf_lbvserver_contentinspectionpolicy_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLbvserver_contentinspectionpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccLbvserver_contentinspectionpolicy_binding_basic},
			{Config: testAccLbvserver_contentinspectionpolicy_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
		},
	})
}

func testAccCheckLbvserver_contentinspectionpolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lbvserver_contentinspectionpolicy_binding id is set")
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
			ResourceType:             "lbvserver_contentinspectionpolicy_binding",
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
			return fmt.Errorf("lbvserver_contentinspectionpolicy_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckLbvserver_contentispectionpolicy_bindingNotExist(n string, id string) resource.TestCheckFunc {
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
			ResourceType:             "lbvserver_contentinspectionpolicy_binding",
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
			return fmt.Errorf("lbvserver_contentinspectionpolicy_binding %s was found, but is should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckLbvserver_contentinspectionpolicy_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_lbvserver_contentinspectionpolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("lbvserver_contentinspectionpolicy_binding", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("lbvserver_contentinspectionpolicy_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccLbvserver_contentinspectionpolicy_bindingDataSource_basic = `

	resource "citrixadc_contentinspectionpolicy" "tf_contentinspectionpolicy" {
		name   = "tf_contentinspectionpolicy"
		rule   = "false"
		action = "DROP"
	}

	resource "citrixadc_lbvserver_contentinspectionpolicy_binding" "tf_lbvserver_contentinspectionpolicy_binding" {
		bindpoint 				= "REQUEST"
		gotopriorityexpression 	= "END"
		name 					= citrixadc_lbvserver.tf_lbvserver.name
		policyname 				= citrixadc_contentinspectionpolicy.tf_contentinspectionpolicy.name
		priority 				= 1    
	}

	resource "citrixadc_lbvserver" "tf_lbvserver" {
		name        = "tf_lbvserver"
		ipv46       = "10.10.10.33"
		port        = 80
		servicetype = "HTTP"
	}

	data "citrixadc_lbvserver_contentinspectionpolicy_binding" "tf_lbvserver_contentinspectionpolicy_binding" {
		name 		= citrixadc_lbvserver_contentinspectionpolicy_binding.tf_lbvserver_contentinspectionpolicy_binding.name
		bindpoint 	= citrixadc_lbvserver_contentinspectionpolicy_binding.tf_lbvserver_contentinspectionpolicy_binding.bindpoint
		policyname 	= citrixadc_lbvserver_contentinspectionpolicy_binding.tf_lbvserver_contentinspectionpolicy_binding.policyname
	}
`

func TestAccLbvserver_contentinspectionpolicy_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccLbvserver_contentinspectionpolicy_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_lbvserver_contentinspectionpolicy_binding.tf_lbvserver_contentinspectionpolicy_binding", "name", "tf_lbvserver"),
					resource.TestCheckResourceAttr("data.citrixadc_lbvserver_contentinspectionpolicy_binding.tf_lbvserver_contentinspectionpolicy_binding", "bindpoint", "REQUEST"),
					resource.TestCheckResourceAttr("data.citrixadc_lbvserver_contentinspectionpolicy_binding.tf_lbvserver_contentinspectionpolicy_binding", "policyname", "tf_contentinspectionpolicy"),
					resource.TestCheckResourceAttr("data.citrixadc_lbvserver_contentinspectionpolicy_binding.tf_lbvserver_contentinspectionpolicy_binding", "priority", "1"),
					resource.TestCheckResourceAttr("data.citrixadc_lbvserver_contentinspectionpolicy_binding.tf_lbvserver_contentinspectionpolicy_binding", "gotopriorityexpression", "END"),
				),
			},
		},
	})
}

// testAccLbvserver_contentinspectionpolicy_binding_upgrade_basic is valid under BOTH the SDK v2
// 2.2.0 schema and the current Framework schema (it uses only the SDK v2 attribute names, which
// the migration restored). It reuses the values from testAccLbvserver_contentinspectionpolicy_binding_basic.
const testAccLbvserver_contentinspectionpolicy_binding_upgrade_basic = `

	resource "citrixadc_contentinspectionpolicy" "tf_contentinspectionpolicy" {
		name   = "tf_contentinspectionpolicy"
		rule   = "false"
		action = "DROP"
	}

	resource "citrixadc_lbvserver" "tf_lbvserver" {
		name        = "tf_lbvserver"
		ipv46       = "10.10.10.33"
		port        = 80
		servicetype = "HTTP"
	}

	resource "citrixadc_lbvserver_contentinspectionpolicy_binding" "tf_lbvserver_contentinspectionpolicy_binding" {
		bindpoint              = "REQUEST"
		gotopriorityexpression = "END"
		name                   = citrixadc_lbvserver.tf_lbvserver.name
		policyname             = citrixadc_contentinspectionpolicy.tf_contentinspectionpolicy.name
		priority               = 1
	}
`

func TestAccLbvserver_contentinspectionpolicy_binding_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckLbvserver_contentinspectionpolicy_bindingDestroy,
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
				Config: testAccLbvserver_contentinspectionpolicy_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbvserver_contentinspectionpolicy_bindingExist("citrixadc_lbvserver_contentinspectionpolicy_binding.tf_lbvserver_contentinspectionpolicy_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_lbvserver_contentinspectionpolicy_binding.tf_lbvserver_contentinspectionpolicy_binding", "id", "tf_lbvserver,tf_contentinspectionpolicy"),
				),
			},
			{
				// Step 2: refresh/plan the legacy-id state through the current
				// framework provider. Read exercises ParseIdString on the legacy id
				// and SetAttrFromGet recomputes the id into the new key:value form.
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccLbvserver_contentinspectionpolicy_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbvserver_contentinspectionpolicy_bindingExist("citrixadc_lbvserver_contentinspectionpolicy_binding.tf_lbvserver_contentinspectionpolicy_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_lbvserver_contentinspectionpolicy_binding.tf_lbvserver_contentinspectionpolicy_binding", "id", "bindpoint:REQUEST,name:tf_lbvserver,policyname:tf_contentinspectionpolicy"),
				),
			},
		},
	})
}
