/*
Copyright 2024 Citrix Systems, Inc

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

const testAccGslbvserver_lbpolicy_binding_basic = `
	resource "citrixadc_gslbvserver" "tf_gslbvserver" {
		name        = "tf_gslbvserver"
		servicetype = "HTTP"
	}
	resource "citrixadc_lbpolicy" "tf_pol" {
		name   = "tf_pol"
		rule   = "true"
		action = "NOLBACTION"
	}
	
	resource "citrixadc_gslbvserver_lbpolicy_binding" "tf_bind" {
		policyname = citrixadc_lbpolicy.tf_pol.name
		name       = citrixadc_gslbvserver.tf_gslbvserver.name
		priority   = 10
	}
	
`

const testAccGslbvserver_lbpolicy_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion

	resource "citrixadc_gslbvserver" "tf_gslbvserver" {
		name        = "tf_gslbvserver"
		servicetype = "HTTP"
	}
	resource "citrixadc_lbpolicy" "tf_pol" {
		name   = "tf_pol"
		rule   = "true"
		action = "NOLBACTION"
	}
`

const testAccGslbvserver_lbpolicy_bindingDataSource_basic = `
	resource "citrixadc_gslbvserver" "tf_gslbvserver" {
		name        = "tf_gslbvserver"
		servicetype = "HTTP"
	}
	resource "citrixadc_lbpolicy" "tf_pol" {
		name   = "tf_pol"
		rule   = "true"
		action = "NOLBACTION"
	}
	
	resource "citrixadc_gslbvserver_lbpolicy_binding" "tf_bind" {
		policyname = citrixadc_lbpolicy.tf_pol.name
		name       = citrixadc_gslbvserver.tf_gslbvserver.name
		priority   = 10
	}

	data "citrixadc_gslbvserver_lbpolicy_binding" "tf_bind" {
		name       = citrixadc_gslbvserver_lbpolicy_binding.tf_bind.name
		policyname = citrixadc_gslbvserver_lbpolicy_binding.tf_bind.policyname
		depends_on = [citrixadc_gslbvserver_lbpolicy_binding.tf_bind]
	}
`

func TestAccGslbvserver_lbpolicy_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckGslbvserver_lbpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGslbvserver_lbpolicy_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGslbvserver_lbpolicy_bindingExist("citrixadc_gslbvserver_lbpolicy_binding.tf_bind", nil),
				),
			},
			{
				Config: testAccGslbvserver_lbpolicy_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGslbvserver_lbpolicy_bindingNotExist("citrixadc_gslbvserver_lbpolicy_binding.tf_bind", "tf_gslbvserver,tf_pol"),
				),
			},
		},
	})
}

func TestAccGslbvserver_lbpolicy_binding_import(t *testing.T) {
	const resAddr = "citrixadc_gslbvserver_lbpolicy_binding.tf_bind"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckGslbvserver_lbpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccGslbvserver_lbpolicy_binding_basic},
			{Config: testAccGslbvserver_lbpolicy_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
		},
	})
}

func testAccCheckGslbvserver_lbpolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No gslbvserver_lbpolicy_binding id is set")
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
			return err
		}
		name := idMap["name"]
		policyname := idMap["policyname"]

		findParams := service.FindParams{
			ResourceType:             "gslbvserver_lbpolicy_binding",
			ResourceName:             name,
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
			return fmt.Errorf("gslbvserver_lbpolicy_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckGslbvserver_lbpolicy_bindingNotExist(n string, id string) resource.TestCheckFunc {
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

		name := idSlice[0]
		policyname := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "gslbvserver_lbpolicy_binding",
			ResourceName:             name,
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
			return fmt.Errorf("gslbvserver_lbpolicy_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckGslbvserver_lbpolicy_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_gslbvserver_lbpolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("gslbvserver_lbpolicy_binding", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("gslbvserver_lbpolicy_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccGslbvserver_lbpolicy_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccGslbvserver_lbpolicy_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_gslbvserver_lbpolicy_binding.tf_bind", "name", "tf_gslbvserver"),
					resource.TestCheckResourceAttr("data.citrixadc_gslbvserver_lbpolicy_binding.tf_bind", "policyname", "tf_pol"),
					resource.TestCheckResourceAttr("data.citrixadc_gslbvserver_lbpolicy_binding.tf_bind", "priority", "10"),
				),
			},
		},
	})
}

// testAccGslbvserver_lbpolicy_binding_upgrade_basic reuses the _basic config
// (participating gslbvserver + lbpolicy plus the binding). It must be valid
// under both the SDK v2 2.2.0 schema and the current framework schema.
const testAccGslbvserver_lbpolicy_binding_upgrade_basic = `
	resource "citrixadc_gslbvserver" "tf_gslbvserver" {
		name        = "tf_gslbvserver"
		servicetype = "HTTP"
	}
	resource "citrixadc_lbpolicy" "tf_pol" {
		name   = "tf_pol"
		rule   = "true"
		action = "NOLBACTION"
	}

	resource "citrixadc_gslbvserver_lbpolicy_binding" "tf_bind" {
		policyname = citrixadc_lbpolicy.tf_pol.name
		name       = citrixadc_gslbvserver.tf_gslbvserver.name
		priority   = 10
	}
`

// TestAccGslbvserver_lbpolicy_binding_sdkv2StateUpgrade verifies that state
// written by the last SDK v2 release (which stores the legacy comma-separated
// id "name,policyname") is read/planned/applied cleanly by the current
// framework provider, and that Read recomputes the id into the new key:value
// form ("name:...,policyname:...").
func TestAccGslbvserver_lbpolicy_binding_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckGslbvserver_lbpolicy_bindingDestroy,
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
				Config: testAccGslbvserver_lbpolicy_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGslbvserver_lbpolicy_bindingExist("citrixadc_gslbvserver_lbpolicy_binding.tf_bind", nil),
					resource.TestCheckResourceAttr("citrixadc_gslbvserver_lbpolicy_binding.tf_bind", "id", "tf_gslbvserver,tf_pol"),
				),
			},
			{
				// Step 2: refresh/plan the legacy-id state through the current
				// framework provider. Read exercises ParseIdString on the legacy id
				// and SetAttrFromGet recomputes the id into the new key:value form.
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccGslbvserver_lbpolicy_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGslbvserver_lbpolicy_bindingExist("citrixadc_gslbvserver_lbpolicy_binding.tf_bind", nil),
					resource.TestCheckResourceAttr("citrixadc_gslbvserver_lbpolicy_binding.tf_bind", "id", "name:tf_gslbvserver,policyname:tf_pol"),
				),
			},
		},
	})
}
