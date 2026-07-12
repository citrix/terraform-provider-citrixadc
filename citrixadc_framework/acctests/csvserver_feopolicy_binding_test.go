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

const testAccCsvserver_feopolicy_binding_basic = `
	# Since the feopolicy resource is not yet available on Terraform,
	# the tf_feopolicy policy must be created by hand in order for the script to run correctly.
	# You can do that by using the following Citrix ADC cli command:
	# add feo policy tf_feopolicy TRUE BASIC

	resource "citrixadc_feopolicy" "tf_feopolicy" {
		name   = "tf_feopolicy"
		action = "BASIC"
		rule   = "true"
	}

	resource "citrixadc_csvserver_feopolicy_binding" "tf_csvserver_feopolicy_binding" {
        name 					= citrixadc_csvserver.tf_csvserver.name
        policyname 				= citrixadc_feopolicy.tf_feopolicy.name
        bindpoint 				= "REQUEST"
        gotopriorityexpression 	= "END"
        priority 				= 1  
	}

	resource "citrixadc_csvserver" "tf_csvserver" {
		name 		= "tf_csvserver"
		ipv46 		= "10.202.11.11"
		port 		= 8080
		servicetype = "HTTP"
	}
`

const testAccCsvserver_feopolicy_binding_basic_step2 = `

	resource "citrixadc_feopolicy" "tf_feopolicy" {
		name   = "tf_feopolicy"
		action = "BASIC"
		rule   = "true"
	}
	resource "citrixadc_csvserver" "tf_csvserver" {
		name 		= "tf_csvserver"
		ipv46 		= "10.202.11.11"
		port 		= 8080
		servicetype = "HTTP"
	}
`

func TestAccCsvserver_feopolicy_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCsvserver_feopolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCsvserver_feopolicy_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCsvserver_feopolicy_bindingExist("citrixadc_csvserver_feopolicy_binding.tf_csvserver_feopolicy_binding", nil),
				),
			},
			{
				Config: testAccCsvserver_feopolicy_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCsvserver_feopolicy_bindingNotExist("citrixadc_csvserver_feopolicy_binding.tf_csvserver_feopolicy_binding", "tf_csvserver,tf_feopolicy"),
				),
			},
		},
	})
}

func testAccCheckCsvserver_feopolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No csvserver_feopolicy_binding id is set")
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
			return fmt.Errorf("Error parsing ID %s: %v", bindingId, err)
		}
		name := idMap["name"]
		policyname := idMap["policyname"]

		findParams := service.FindParams{
			ResourceType:             "csvserver_feopolicy_binding",
			ResourceName:             name,
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
			if v["policyname"].(string) == policyname {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("csvserver_feopolicy_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckCsvserver_feopolicy_bindingNotExist(n string, id string) resource.TestCheckFunc {
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

		csvserverName := idSlice[0]
		policyName := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "csvserver_feopolicy_binding",
			ResourceName:             csvserverName,
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
			return fmt.Errorf("csvserver_feopolicy_binding %s was found, but is should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckCsvserver_feopolicy_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_csvserver_feopolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Csvserver_feopolicy_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("csvserver_feopolicy_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccCsvserver_feopolicy_bindingDataSource_basic = `

	resource "citrixadc_feopolicy" "tf_feopolicy" {
		name   = "my_feopolicy_ds"
		action = "BASIC"
		rule   = "true"
	}
	resource "citrixadc_csvserver" "tf_csvserver" {
		name 		= "my_csvserver_ds"
		ipv46 		= "10.202.11.12"
		port 		= 8081
		servicetype = "HTTP"
	}
	resource "citrixadc_csvserver_feopolicy_binding" "tf_csvserver_feopolicy_binding" {
		name 					= citrixadc_csvserver.tf_csvserver.name
		policyname 				= citrixadc_feopolicy.tf_feopolicy.name
		bindpoint 				= "REQUEST"
		gotopriorityexpression 	= "END"
		priority 				= 10  
	}

	data "citrixadc_csvserver_feopolicy_binding" "tf_csvserver_feopolicy_binding_ds" {
		name       = citrixadc_csvserver_feopolicy_binding.tf_csvserver_feopolicy_binding.name
		policyname = citrixadc_csvserver_feopolicy_binding.tf_csvserver_feopolicy_binding.policyname
		depends_on = [citrixadc_csvserver_feopolicy_binding.tf_csvserver_feopolicy_binding]
	}
`

func TestAcccsvserver_feopolicy_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccCsvserver_feopolicy_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_csvserver_feopolicy_binding.tf_csvserver_feopolicy_binding_ds", "name", "my_csvserver_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_csvserver_feopolicy_binding.tf_csvserver_feopolicy_binding_ds", "policyname", "my_feopolicy_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_csvserver_feopolicy_binding.tf_csvserver_feopolicy_binding_ds", "priority", "10"),
					resource.TestCheckResourceAttr("data.citrixadc_csvserver_feopolicy_binding.tf_csvserver_feopolicy_binding_ds", "gotopriorityexpression", "END"),
				),
			},
		},
	})
}

// testAccCsvserver_feopolicy_binding_upgrade_basic reuses the _basic config (binding +
// all prerequisite resources) with identical resource labels so the Exist/Destroy
// helpers and addresses match. It is valid under BOTH the SDK v2 2.2.0 schema and the
// current Framework schema because the migration restored the SDK v2 attribute names.
const testAccCsvserver_feopolicy_binding_upgrade_basic = `
	resource "citrixadc_feopolicy" "tf_feopolicy" {
		name   = "tf_feopolicy"
		action = "BASIC"
		rule   = "true"
	}

	resource "citrixadc_csvserver_feopolicy_binding" "tf_csvserver_feopolicy_binding" {
        name 					= citrixadc_csvserver.tf_csvserver.name
        policyname 				= citrixadc_feopolicy.tf_feopolicy.name
        bindpoint 				= "REQUEST"
        gotopriorityexpression 	= "END"
        priority 				= 1
	}

	resource "citrixadc_csvserver" "tf_csvserver" {
		name 		= "tf_csvserver"
		ipv46 		= "10.202.11.11"
		port 		= 8080
		servicetype = "HTTP"
	}
`

// TestAccCsvserver_feopolicy_binding_sdkv2StateUpgrade verifies that state written by
// the last SDK v2 release (legacy comma-separated ID) is correctly upgraded when the
// same config is subsequently managed by the current Framework provider. Step 1 creates
// the binding with citrix/citrixadc 2.2.0 (writes the legacy id
// "tf_csvserver,tf_feopolicy"). Step 2 refreshes/plans/applies the same config through
// the Framework provider, exercising ParseIdString on the legacy id; because the
// Framework recomputes the id on Read (SetAttrFromGet), the id upgrades to the new
// "key:value" form.
func TestAccCsvserver_feopolicy_binding_sdkv2StateUpgrade(t *testing.T) {
	resourceAddr := "citrixadc_csvserver_feopolicy_binding.tf_csvserver_feopolicy_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckCsvserver_feopolicy_bindingDestroy,
		Steps: []resource.TestStep{
			// Step 1: create with the last SDK v2 release -> state carries the legacy id.
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.2.0",
					},
				},
				Config: testAccCsvserver_feopolicy_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCsvserver_feopolicy_bindingExist(resourceAddr, nil),
					resource.TestCheckResourceAttr(resourceAddr, "id", "tf_csvserver,tf_feopolicy"),
				),
			},
			// Step 2: refresh/plan/apply the SAME config through the current Framework
			// provider. The legacy-id state is read via ParseIdString and the id is
			// recomputed to the new key:value format.
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccCsvserver_feopolicy_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCsvserver_feopolicy_bindingExist(resourceAddr, nil),
					resource.TestCheckResourceAttr(resourceAddr, "id", "name:tf_csvserver,policyname:tf_feopolicy"),
				),
			},
		},
	})
}

func TestAccCsvserver_feopolicy_binding_import(t *testing.T) {
	const resAddr = "citrixadc_csvserver_feopolicy_binding.tf_csvserver_feopolicy_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCsvserver_feopolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccCsvserver_feopolicy_binding_basic},
			{Config: testAccCsvserver_feopolicy_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
		},
	})
}
