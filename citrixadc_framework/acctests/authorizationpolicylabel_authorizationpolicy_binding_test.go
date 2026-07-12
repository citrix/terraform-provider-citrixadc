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

const testAccAuthorizationpolicylabel_authorizationpolicy_binding_basic = `

resource "citrixadc_authorizationpolicy" "authorize" {
	name   = "tp-authorize-1"
	rule   = "true"
	action = "DENY"
	}
  resource "citrixadc_authorizationpolicylabel" "authorizationpolicylabel" {
	labelname = "trans_http_url"
	}
  resource "citrixadc_authorizationpolicylabel_authorizationpolicy_binding" "authorizationpolicylabel_authorizationpolicy_binding" {
	policyname = citrixadc_authorizationpolicy.authorize.name
	labelname = citrixadc_authorizationpolicylabel.authorizationpolicylabel.labelname
	priority = 2
	}
`

const testAccAuthorizationpolicylabel_authorizationpolicy_binding_basic_step2 = `
	resource "citrixadc_authorizationpolicy" "authorize" {
	name   = "tp-authorize-1"
	rule   = "true"
	action = "DENY"
	}
	resource "citrixadc_authorizationpolicylabel" "authorizationpolicylabel" {
	labelname = "trans_http_url"
	}
`

const testAccAuthorizationpolicylabel_authorizationpolicy_bindingDataSource_basic = `

resource "citrixadc_authorizationpolicy" "authorize" {
	name   = "tp-authorize-1"
	rule   = "true"
	action = "DENY"
	}
  resource "citrixadc_authorizationpolicylabel" "authorizationpolicylabel" {
	labelname = "trans_http_url"
	}
  resource "citrixadc_authorizationpolicylabel_authorizationpolicy_binding" "authorizationpolicylabel_authorizationpolicy_binding" {
	policyname = citrixadc_authorizationpolicy.authorize.name
	labelname = citrixadc_authorizationpolicylabel.authorizationpolicylabel.labelname
	priority = 2
	}

	data "citrixadc_authorizationpolicylabel_authorizationpolicy_binding" "authorizationpolicylabel_authorizationpolicy_binding" {
		labelname  = citrixadc_authorizationpolicylabel_authorizationpolicy_binding.authorizationpolicylabel_authorizationpolicy_binding.labelname
		policyname = citrixadc_authorizationpolicylabel_authorizationpolicy_binding.authorizationpolicylabel_authorizationpolicy_binding.policyname
		depends_on = [citrixadc_authorizationpolicylabel_authorizationpolicy_binding.authorizationpolicylabel_authorizationpolicy_binding]
	}
`

func TestAccAuthorizationpolicylabel_authorizationpolicy_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthorizationpolicylabel_authorizationpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthorizationpolicylabel_authorizationpolicy_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthorizationpolicylabel_authorizationpolicy_bindingExist("citrixadc_authorizationpolicylabel_authorizationpolicy_binding.authorizationpolicylabel_authorizationpolicy_binding", nil),
				),
			},
			{
				Config: testAccAuthorizationpolicylabel_authorizationpolicy_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthorizationpolicylabel_authorizationpolicy_bindingNotExist("citrixadc_authorizationpolicylabel_authorizationpolicy_binding.authorizationpolicylabel_authorizationpolicy_binding", "trans_http_url,tp-authorize-1"),
				),
			},
		},
	})
}

func testAccCheckAuthorizationpolicylabel_authorizationpolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No authorizationpolicylabel_authorizationpolicy_binding id is set")
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
			ResourceType:             "authorizationpolicylabel_authorizationpolicy_binding",
			ResourceName:             labelname,
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
			return fmt.Errorf("authorizationpolicylabel_authorizationpolicy_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckAuthorizationpolicylabel_authorizationpolicy_bindingNotExist(n string, id string) resource.TestCheckFunc {
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
			ResourceType:             "authorizationpolicylabel_authorizationpolicy_binding",
			ResourceName:             labelname,
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
			return fmt.Errorf("authorizationpolicylabel_authorizationpolicy_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckAuthorizationpolicylabel_authorizationpolicy_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_authorizationpolicylabel_authorizationpolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Authorizationpolicylabel_authorizationpolicy_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("authorizationpolicylabel_authorizationpolicy_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccAuthorizationpolicylabel_authorizationpolicy_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthorizationpolicylabel_authorizationpolicy_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_authorizationpolicylabel_authorizationpolicy_binding.authorizationpolicylabel_authorizationpolicy_binding", "labelname", "trans_http_url"),
					resource.TestCheckResourceAttr("data.citrixadc_authorizationpolicylabel_authorizationpolicy_binding.authorizationpolicylabel_authorizationpolicy_binding", "policyname", "tp-authorize-1"),
					resource.TestCheckResourceAttr("data.citrixadc_authorizationpolicylabel_authorizationpolicy_binding.authorizationpolicylabel_authorizationpolicy_binding", "priority", "2"),
				),
			},
		},
	})
}

const testAccAuthorizationpolicylabel_authorizationpolicy_binding_upgrade_basic = `

resource "citrixadc_authorizationpolicy" "authorize" {
	name   = "tp-authorize-1"
	rule   = "true"
	action = "DENY"
	}
  resource "citrixadc_authorizationpolicylabel" "authorizationpolicylabel" {
	labelname = "trans_http_url"
	}
  resource "citrixadc_authorizationpolicylabel_authorizationpolicy_binding" "authorizationpolicylabel_authorizationpolicy_binding" {
	policyname = citrixadc_authorizationpolicy.authorize.name
	labelname = citrixadc_authorizationpolicylabel.authorizationpolicylabel.labelname
	priority = 2
	}
`

// TestAccAuthorizationpolicylabel_authorizationpolicy_binding_sdkv2StateUpgrade verifies that
// a binding created with the last SDK v2 release (v2.2.0), which writes state using the legacy
// comma-joined ID (labelname,policyname), is correctly refreshed and re-planned by the current
// Plugin Framework provider. The framework Read recomputes the ID to the new
// key:UrlEncode(value) format, so after the upgrade step the ID becomes the canonical new format.
func TestAccAuthorizationpolicylabel_authorizationpolicy_binding_sdkv2StateUpgrade(t *testing.T) {
	resourceAddr := "citrixadc_authorizationpolicylabel_authorizationpolicy_binding.authorizationpolicylabel_authorizationpolicy_binding"
	legacyId := "trans_http_url,tp-authorize-1"
	newId := "labelname:trans_http_url,policyname:tp-authorize-1"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckAuthorizationpolicylabel_authorizationpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			// Step 1: Create with the last SDK v2 release (writes legacy-format ID to state).
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.2.0",
					},
				},
				Config: testAccAuthorizationpolicylabel_authorizationpolicy_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthorizationpolicylabel_authorizationpolicy_bindingExist(resourceAddr, nil),
					resource.TestCheckResourceAttr(resourceAddr, "id", legacyId),
				),
			},
			// Step 2: Refresh/plan/apply the legacy-ID state through the current framework provider.
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccAuthorizationpolicylabel_authorizationpolicy_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthorizationpolicylabel_authorizationpolicy_bindingExist(resourceAddr, nil),
					resource.TestCheckResourceAttr(resourceAddr, "id", newId),
				),
			},
		},
	})
}

func TestAccAuthorizationpolicylabel_authorizationpolicy_binding_import(t *testing.T) {
	const resAddr = "citrixadc_authorizationpolicylabel_authorizationpolicy_binding.authorizationpolicylabel_authorizationpolicy_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthorizationpolicylabel_authorizationpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccAuthorizationpolicylabel_authorizationpolicy_binding_basic},
			{Config: testAccAuthorizationpolicylabel_authorizationpolicy_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
		},
	})
}
