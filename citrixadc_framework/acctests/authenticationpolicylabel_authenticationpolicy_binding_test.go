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

const testAccAuthenticationpolicylabel_authenticationpolicy_binding_basic = `

	resource "citrixadc_authenticationpolicylabel" "tf_authenticationpolicylabel" {
		labelname = "tf_authenticationpolicylabel"
		type      = "AAATM_REQ"
		comment   = "Testing"
	}
	resource "citrixadc_authenticationldapaction" "tf_authenticationldapaction" {
		name          = "ldapaction"
		serverip      = "1.2.3.4"
		serverport    = 8080
		authtimeout   = 1
		ldaploginname = "username"
	}
	resource "citrixadc_authenticationpolicy" "tf_authenticationpolicy" {
		name   = "tf_authenticationpolicy"
		rule   = "true"
		action = citrixadc_authenticationldapaction.tf_authenticationldapaction.name
	}
	resource "citrixadc_authenticationpolicylabel_authenticationpolicy_binding" "tf_bind" {
		labelname  = citrixadc_authenticationpolicylabel.tf_authenticationpolicylabel.labelname
		policyname = citrixadc_authenticationpolicy.tf_authenticationpolicy.name
		priority   = 20
	}
`

const testAccAuthenticationpolicylabel_authenticationpolicy_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion
	resource "citrixadc_authenticationpolicylabel" "tf_authenticationpolicylabel" {
		labelname = "tf_authenticationpolicylabel"
		type      = "AAATM_REQ"
		comment   = "Testing"
	}
	resource "citrixadc_authenticationldapaction" "tf_authenticationldapaction" {
		name          = "ldapaction"
		serverip      = "1.2.3.4"
		serverport    = 8080
		authtimeout   = 1
		ldaploginname = "username"
	}
	resource "citrixadc_authenticationpolicy" "tf_authenticationpolicy" {
		name   = "tf_authenticationpolicy"
		rule   = "true"
		action = citrixadc_authenticationldapaction.tf_authenticationldapaction.name
	}
`

const testAccAuthenticationpolicylabel_authenticationpolicy_bindingDataSource_basic = `

	resource "citrixadc_authenticationpolicylabel" "tf_authenticationpolicylabel" {
		labelname = "tf_authenticationpolicylabel"
		type      = "AAATM_REQ"
		comment   = "Testing"
	}
	resource "citrixadc_authenticationldapaction" "tf_authenticationldapaction" {
		name          = "ldapaction"
		serverip      = "1.2.3.4"
		serverport    = 8080
		authtimeout   = 1
		ldaploginname = "username"
	}
	resource "citrixadc_authenticationpolicy" "tf_authenticationpolicy" {
		name   = "tf_authenticationpolicy"
		rule   = "true"
		action = citrixadc_authenticationldapaction.tf_authenticationldapaction.name
	}
	resource "citrixadc_authenticationpolicylabel_authenticationpolicy_binding" "tf_bind" {
		labelname  = citrixadc_authenticationpolicylabel.tf_authenticationpolicylabel.labelname
		policyname = citrixadc_authenticationpolicy.tf_authenticationpolicy.name
		priority   = 20
	}

	data "citrixadc_authenticationpolicylabel_authenticationpolicy_binding" "tf_bind" {
		labelname  = citrixadc_authenticationpolicylabel_authenticationpolicy_binding.tf_bind.labelname
		policyname = citrixadc_authenticationpolicylabel_authenticationpolicy_binding.tf_bind.policyname
		depends_on = [citrixadc_authenticationpolicylabel_authenticationpolicy_binding.tf_bind]
	}
`

func TestAccAuthenticationpolicylabel_authenticationpolicy_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationpolicylabel_authenticationpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationpolicylabel_authenticationpolicy_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationpolicylabel_authenticationpolicy_bindingExist("citrixadc_authenticationpolicylabel_authenticationpolicy_binding.tf_bind", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationpolicylabel_authenticationpolicy_binding.tf_bind", "labelname", "tf_authenticationpolicylabel"),
					resource.TestCheckResourceAttr("citrixadc_authenticationpolicylabel_authenticationpolicy_binding.tf_bind", "policyname", "tf_authenticationpolicy"),
					resource.TestCheckResourceAttr("citrixadc_authenticationpolicylabel_authenticationpolicy_binding.tf_bind", "priority", "20"),
				),
			},
			{
				Config: testAccAuthenticationpolicylabel_authenticationpolicy_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationpolicylabel_authenticationpolicy_bindingNotExist("citrixadc_authenticationpolicylabel_authenticationpolicy_binding.tf_bind", "tf_authenticationpolicylabel,tf_authenticationpolicy"),
				),
			},
		},
	})
}

func testAccCheckAuthenticationpolicylabel_authenticationpolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No authenticationpolicylabel_authenticationpolicy_binding id is set")
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

		idMap, _, err := utils.ParseIdString(bindingId, []string{"labelname", "policyname"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", bindingId, err)
		}
		labelname := idMap["labelname"]
		policyname := idMap["policyname"]

		findParams := service.FindParams{
			ResourceType:             "authenticationpolicylabel_authenticationpolicy_binding",
			ResourceName:             labelname,
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
			return fmt.Errorf("authenticationpolicylabel_authenticationpolicy_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckAuthenticationpolicylabel_authenticationpolicy_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		idMap, _, err := utils.ParseIdString(id, []string{"labelname", "policyname"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", id, err)
		}
		labelname := idMap["labelname"]
		policyname := idMap["policyname"]

		findParams := service.FindParams{
			ResourceType:             "authenticationpolicylabel_authenticationpolicy_binding",
			ResourceName:             labelname,
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
			return fmt.Errorf("authenticationpolicylabel_authenticationpolicy_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckAuthenticationpolicylabel_authenticationpolicy_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_authenticationpolicylabel_authenticationpolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Authenticationpolicylabel_authenticationpolicy_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("authenticationpolicylabel_authenticationpolicy_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccAuthenticationpolicylabel_authenticationpolicy_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationpolicylabel_authenticationpolicy_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_authenticationpolicylabel_authenticationpolicy_binding.tf_bind", "labelname", "tf_authenticationpolicylabel"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationpolicylabel_authenticationpolicy_binding.tf_bind", "policyname", "tf_authenticationpolicy"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationpolicylabel_authenticationpolicy_binding.tf_bind", "priority", "20"),
				),
			},
		},
	})
}

const testAccAuthenticationpolicylabel_authenticationpolicy_binding_upgrade_basic = `

	resource "citrixadc_authenticationpolicylabel" "tf_authenticationpolicylabel" {
		labelname = "tf_authenticationpolicylabel"
		type      = "AAATM_REQ"
		comment   = "Testing"
	}
	resource "citrixadc_authenticationldapaction" "tf_authenticationldapaction" {
		name          = "ldapaction"
		serverip      = "1.2.3.4"
		serverport    = 8080
		authtimeout   = 1
		ldaploginname = "username"
	}
	resource "citrixadc_authenticationpolicy" "tf_authenticationpolicy" {
		name   = "tf_authenticationpolicy"
		rule   = "true"
		action = citrixadc_authenticationldapaction.tf_authenticationldapaction.name
	}
	resource "citrixadc_authenticationpolicylabel_authenticationpolicy_binding" "tf_bind" {
		labelname  = citrixadc_authenticationpolicylabel.tf_authenticationpolicylabel.labelname
		policyname = citrixadc_authenticationpolicy.tf_authenticationpolicy.name
		priority   = 20
	}
`

func TestAccAuthenticationpolicylabel_authenticationpolicy_binding_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckAuthenticationpolicylabel_authenticationpolicy_bindingDestroy,
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
				Config: testAccAuthenticationpolicylabel_authenticationpolicy_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationpolicylabel_authenticationpolicy_bindingExist("citrixadc_authenticationpolicylabel_authenticationpolicy_binding.tf_bind", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationpolicylabel_authenticationpolicy_binding.tf_bind", "id", "tf_authenticationpolicylabel,tf_authenticationpolicy"),
				),
			},
			{
				// Step 2: refresh/plan the legacy-id state through the current
				// framework provider. Read exercises ParseIdString on the legacy id
				// and SetAttrFromGet recomputes the id into the new key:value form.
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccAuthenticationpolicylabel_authenticationpolicy_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationpolicylabel_authenticationpolicy_bindingExist("citrixadc_authenticationpolicylabel_authenticationpolicy_binding.tf_bind", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationpolicylabel_authenticationpolicy_binding.tf_bind", "id", "labelname:tf_authenticationpolicylabel,policyname:tf_authenticationpolicy"),
				),
			},
		},
	})
}

func TestAccAuthenticationpolicylabel_authenticationpolicy_binding_import(t *testing.T) {
	const resAddr = "citrixadc_authenticationpolicylabel_authenticationpolicy_binding.tf_bind"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationpolicylabel_authenticationpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccAuthenticationpolicylabel_authenticationpolicy_binding_basic},
			{Config: testAccAuthenticationpolicylabel_authenticationpolicy_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
		},
	})
}
