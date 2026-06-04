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

// Participating-entity config reused from:
//   - authenticationvserver_test.go            (parent: citrixadc_authenticationvserver)
//   - authenticationsmartaccesspolicy_test.go  (policy + its action profile)
//   - authenticationsmartaccessprofile_test.go (the action referenced by the policy)
//
// Dependency chain:
//   authenticationsmartaccessprofile (action) ->
//   authenticationsmartaccesspolicy  (policy, action = profile.name) ->
//   authenticationvserver            (parent vserver) ->
//   authenticationvserver_authenticationsmartaccesspolicy_binding (name = vserver.name, policy = policy.name)

const testAccAuthenticationvserver_authenticationsmartaccesspolicy_binding_basic_step1 = `
resource "citrixadc_authenticationsmartaccessprofile" "tf_authenticationsmartaccessprofile" {
  name    = "tf_authenticationsmartaccessprofile"
  tags    = "test_tag"
  comment = "test_comment"
}

resource "citrixadc_authenticationsmartaccesspolicy" "tf_authenticationsmartaccesspolicy" {
  name    = "tf_authenticationsmartaccesspolicy"
  action  = citrixadc_authenticationsmartaccessprofile.tf_authenticationsmartaccessprofile.name
  rule    = "TRUE"
  comment = "test_comment"
}

resource "citrixadc_authenticationvserver" "tf_authenticationvserver" {
  name           = "tf_authenticationvserver"
  servicetype    = "SSL"
  comment        = "Hello"
  authentication = "ON"
  state          = "ENABLED"
}

resource "citrixadc_authenticationvserver_authenticationsmartaccesspolicy_binding" "tf_binding" {
  name     = citrixadc_authenticationvserver.tf_authenticationvserver.name
  policy   = citrixadc_authenticationsmartaccesspolicy.tf_authenticationsmartaccesspolicy.name
  priority = 100
  depends_on = [
    citrixadc_authenticationvserver.tf_authenticationvserver,
    citrixadc_authenticationsmartaccesspolicy.tf_authenticationsmartaccesspolicy,
  ]
}
`

// Step 2 keeps the participating entities but drops the binding to verify proper unbind/delete.
const testAccAuthenticationvserver_authenticationsmartaccesspolicy_binding_basic_step2 = `
resource "citrixadc_authenticationsmartaccessprofile" "tf_authenticationsmartaccessprofile" {
  name    = "tf_authenticationsmartaccessprofile"
  tags    = "test_tag"
  comment = "test_comment"
}

resource "citrixadc_authenticationsmartaccesspolicy" "tf_authenticationsmartaccesspolicy" {
  name    = "tf_authenticationsmartaccesspolicy"
  action  = citrixadc_authenticationsmartaccessprofile.tf_authenticationsmartaccessprofile.name
  rule    = "TRUE"
  comment = "test_comment"
}

resource "citrixadc_authenticationvserver" "tf_authenticationvserver" {
  name           = "tf_authenticationvserver"
  servicetype    = "SSL"
  comment        = "Hello"
  authentication = "ON"
  state          = "ENABLED"
}
`

const testAccAuthenticationvserverAuthenticationsmartaccesspolicyBindingDataSource_basic = `
resource "citrixadc_authenticationsmartaccessprofile" "tf_authenticationsmartaccessprofile" {
  name    = "tf_authenticationsmartaccessprofile"
  tags    = "test_tag"
  comment = "test_comment"
}

resource "citrixadc_authenticationsmartaccesspolicy" "tf_authenticationsmartaccesspolicy" {
  name    = "tf_authenticationsmartaccesspolicy"
  action  = citrixadc_authenticationsmartaccessprofile.tf_authenticationsmartaccessprofile.name
  rule    = "TRUE"
  comment = "test_comment"
}

resource "citrixadc_authenticationvserver" "tf_authenticationvserver" {
  name           = "tf_authenticationvserver"
  servicetype    = "SSL"
  comment        = "Hello"
  authentication = "ON"
  state          = "ENABLED"
}

resource "citrixadc_authenticationvserver_authenticationsmartaccesspolicy_binding" "tf_binding" {
  name     = citrixadc_authenticationvserver.tf_authenticationvserver.name
  policy   = citrixadc_authenticationsmartaccesspolicy.tf_authenticationsmartaccesspolicy.name
  priority = 100
  depends_on = [
    citrixadc_authenticationvserver.tf_authenticationvserver,
    citrixadc_authenticationsmartaccesspolicy.tf_authenticationsmartaccesspolicy,
  ]
}

data "citrixadc_authenticationvserver_authenticationsmartaccesspolicy_binding" "tf_binding" {
  name       = citrixadc_authenticationvserver_authenticationsmartaccesspolicy_binding.tf_binding.name
  policy     = citrixadc_authenticationvserver_authenticationsmartaccesspolicy_binding.tf_binding.policy
  depends_on = [citrixadc_authenticationvserver_authenticationsmartaccesspolicy_binding.tf_binding]
}
`

func TestAccAuthenticationvserver_authenticationsmartaccesspolicy_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationvserver_authenticationsmartaccesspolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationvserver_authenticationsmartaccesspolicy_binding_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationvserver_authenticationsmartaccesspolicy_bindingExist("citrixadc_authenticationvserver_authenticationsmartaccesspolicy_binding.tf_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationvserver_authenticationsmartaccesspolicy_binding.tf_binding", "name", "tf_authenticationvserver"),
					resource.TestCheckResourceAttr("citrixadc_authenticationvserver_authenticationsmartaccesspolicy_binding.tf_binding", "policy", "tf_authenticationsmartaccesspolicy"),
					resource.TestCheckResourceAttr("citrixadc_authenticationvserver_authenticationsmartaccesspolicy_binding.tf_binding", "priority", "100"),
				),
			},
			{
				Config: testAccAuthenticationvserver_authenticationsmartaccesspolicy_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationvserver_authenticationsmartaccesspolicy_bindingNotExist("citrixadc_authenticationvserver_authenticationsmartaccesspolicy_binding.tf_binding", "tf_authenticationvserver", "tf_authenticationsmartaccesspolicy"),
				),
			},
		},
	})
}

func testAccCheckAuthenticationvserver_authenticationsmartaccesspolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No authenticationvserver_authenticationsmartaccesspolicy_binding id is set")
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

		// ID = name:<urlenc>,policy:<urlenc>
		idMap, _, err := utils.ParseIdString(rs.Primary.ID, []string{"name", "policy"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", rs.Primary.ID, err)
		}

		name := idMap["name"]
		policy := idMap["policy"]

		findParams := service.FindParams{
			ResourceType:             service.Authenticationvserver_authenticationsmartaccesspolicy_binding.Type(),
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the matching policy
		found := false
		for _, v := range dataArr {
			if val, ok := v["policy"].(string); ok && val == policy {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("authenticationvserver_authenticationsmartaccesspolicy_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckAuthenticationvserver_authenticationsmartaccesspolicy_bindingNotExist(n string, name string, policy string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		findParams := service.FindParams{
			ResourceType:             service.Authenticationvserver_authenticationsmartaccesspolicy_binding.Type(),
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			return err
		}

		// Iterate through results to hopefully not find the one with the matching policy
		found := false
		for _, v := range dataArr {
			if val, ok := v["policy"].(string); ok && val == policy {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("authenticationvserver_authenticationsmartaccesspolicy_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckAuthenticationvserver_authenticationsmartaccesspolicy_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_authenticationvserver_authenticationsmartaccesspolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		// ID = name:<urlenc>,policy:<urlenc>
		idMap, _, err := utils.ParseIdString(rs.Primary.ID, []string{"name", "policy"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", rs.Primary.ID, err)
		}

		name := idMap["name"]
		policy := idMap["policy"]

		findParams := service.FindParams{
			ResourceType:             service.Authenticationvserver_authenticationsmartaccesspolicy_binding.Type(),
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			// Parent vserver is gone -> binding is gone too.
			continue
		}

		for _, v := range dataArr {
			if val, ok := v["policy"].(string); ok && val == policy {
				return fmt.Errorf("authenticationvserver_authenticationsmartaccesspolicy_binding %s still exists", rs.Primary.ID)
			}
		}
	}

	return nil
}

func TestAccAuthenticationvserverAuthenticationsmartaccesspolicyBindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationvserverAuthenticationsmartaccesspolicyBindingDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_authenticationvserver_authenticationsmartaccesspolicy_binding.tf_binding", "name", "tf_authenticationvserver"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationvserver_authenticationsmartaccesspolicy_binding.tf_binding", "policy", "tf_authenticationsmartaccesspolicy"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationvserver_authenticationsmartaccesspolicy_binding.tf_binding", "priority", "100"),
				),
			},
		},
	})
}
