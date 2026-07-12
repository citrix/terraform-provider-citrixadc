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

const testAccLbvserver_rewritepolicy_binding_basic_step1 = `
resource "citrixadc_lbvserver" "tf_lbvserver" {
  ipv46       = "10.10.10.33"
  name        = "tf_lbvserver"
  port        = 80
  servicetype = "HTTP"
}

resource "citrixadc_lbvserver" "tf_lbvserver2" {
  ipv46       = "10.10.10.34"
  name        = "tf_lbvserver2"
  port        = 80
  servicetype = "HTTP"
}

resource "citrixadc_rewritepolicy" "tf_rewrite_policy" {
  name   = "tf_test_rewrite_policy"
  action = "DROP"
  rule   = "HTTP.REQ.URL.PATH_AND_QUERY.CONTAINS(\"helloandby\")"
}

resource "citrixadc_lbvserver_rewritepolicy_binding" "tf_bind" {
    name = citrixadc_lbvserver.tf_lbvserver.name
    policyname = citrixadc_rewritepolicy.tf_rewrite_policy.name
    priority = 100
    bindpoint = "REQUEST"
}

resource "citrixadc_lbvserver_rewritepolicy_binding" "tf_bind2" {
    name = citrixadc_lbvserver.tf_lbvserver2.name
    policyname = citrixadc_rewritepolicy.tf_rewrite_policy.name
    priority = 100
    bindpoint = "REQUEST"
}
`

const testAccLbvserver_rewritepolicy_binding_basic_step2 = `
resource "citrixadc_lbvserver" "tf_lbvserver" {
  ipv46       = "10.10.10.33"
  name        = "tf_lbvserver"
  port        = 80
  servicetype = "HTTP"
}

resource "citrixadc_lbvserver" "tf_lbvserver2" {
  ipv46       = "10.10.10.34"
  name        = "tf_lbvserver2"
  port        = 80
  servicetype = "HTTP"
}

resource "citrixadc_rewritepolicy" "tf_rewrite_policy" {
  name   = "tf_test_rewrite_policy"
  action = "DROP"
  rule   = "HTTP.REQ.URL.PATH_AND_QUERY.CONTAINS(\"helloandby\")"
}

resource "citrixadc_lbvserver_rewritepolicy_binding" "tf_bind" {
    name = citrixadc_lbvserver.tf_lbvserver.name
    policyname = citrixadc_rewritepolicy.tf_rewrite_policy.name
    priority = 110
    bindpoint = "REQUEST"
}

resource "citrixadc_lbvserver_rewritepolicy_binding" "tf_bind2" {
    name = citrixadc_lbvserver.tf_lbvserver2.name
    policyname = citrixadc_rewritepolicy.tf_rewrite_policy.name
    priority = 120
    bindpoint = "REQUEST"
}
`

const testAccLbvserver_rewritepolicy_bindingDataSource_basic = `
resource "citrixadc_lbvserver" "tf_lbvserver" {
  ipv46       = "10.10.10.33"
  name        = "tf_lbvserver"
  port        = 80
  servicetype = "HTTP"
}

resource "citrixadc_rewritepolicy" "tf_rewrite_policy" {
  name   = "tf_test_rewrite_policy"
  action = "DROP"
  rule   = "HTTP.REQ.URL.PATH_AND_QUERY.CONTAINS(\"helloandby\")"
}

resource "citrixadc_lbvserver_rewritepolicy_binding" "tf_bind" {
    name = citrixadc_lbvserver.tf_lbvserver.name
    policyname = citrixadc_rewritepolicy.tf_rewrite_policy.name
    priority = 100
    bindpoint = "REQUEST"
}

data "citrixadc_lbvserver_rewritepolicy_binding" "tf_bind" {
    name = citrixadc_lbvserver_rewritepolicy_binding.tf_bind.name
    policyname = citrixadc_lbvserver_rewritepolicy_binding.tf_bind.policyname
    bindpoint = citrixadc_lbvserver_rewritepolicy_binding.tf_bind.bindpoint
    depends_on = [citrixadc_lbvserver_rewritepolicy_binding.tf_bind]
}
`

func TestAccLbvserver_rewritepolicy_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLbvserver_rewritepolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLbvserver_rewritepolicy_binding_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbvserver_rewritepolicy_bindingExist("citrixadc_lbvserver_rewritepolicy_binding.tf_bind", nil),
					testAccCheckLbvserver_rewritepolicy_bindingExist("citrixadc_lbvserver_rewritepolicy_binding.tf_bind2", nil),
				),
			},
			{
				Config: testAccLbvserver_rewritepolicy_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbvserver_rewritepolicy_bindingExist("citrixadc_lbvserver_rewritepolicy_binding.tf_bind", nil),
					testAccCheckLbvserver_rewritepolicy_bindingExist("citrixadc_lbvserver_rewritepolicy_binding.tf_bind2", nil),
				),
			},
		},
	})
}

func testAccCheckLbvserver_rewritepolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lbvserver_rewritepolicy_binding name is set")
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
			ResourceType:             "lbvserver_rewritepolicy_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the right policy name
		foundIndex := -1
		for i, v := range dataArr {
			if v["policyname"].(string) == policyname {
				foundIndex = i
				break
			}
		}

		// Resource is missing
		if foundIndex == -1 {
			return fmt.Errorf("Cannot find lbvserver_rewritepolicy_binding %s", bindingId)
		}

		return nil
	}
}

func testAccCheckLbvserver_rewritepolicy_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_lbvserver_rewritepolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Lbvserver_rewritepolicy_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("lbvserver_rewritepolicy_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccLbvserver_rewritepolicy_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccLbvserver_rewritepolicy_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_lbvserver_rewritepolicy_binding.tf_bind", "name", "tf_lbvserver"),
					resource.TestCheckResourceAttr("data.citrixadc_lbvserver_rewritepolicy_binding.tf_bind", "policyname", "tf_test_rewrite_policy"),
					resource.TestCheckResourceAttr("data.citrixadc_lbvserver_rewritepolicy_binding.tf_bind", "priority", "100"),
					resource.TestCheckResourceAttr("data.citrixadc_lbvserver_rewritepolicy_binding.tf_bind", "bindpoint", "REQUEST"),
				),
			},
		},
	})
}

// testAcclbvserver_rewritepolicy_binding_upgrade_basic is the config used by the
// sdkv2 -> framework state-upgrade test. It reuses the same values and resource
// label (tf_bind) as testAccLbvserver_rewritepolicy_binding_basic_step1 so it is
// valid under BOTH the SDK v2 2.2.0 schema and the current framework schema.
const testAcclbvserver_rewritepolicy_binding_upgrade_basic = `
resource "citrixadc_lbvserver" "tf_lbvserver" {
  ipv46       = "10.10.10.33"
  name        = "tf_lbvserver"
  port        = 80
  servicetype = "HTTP"
}

resource "citrixadc_rewritepolicy" "tf_rewrite_policy" {
  name   = "tf_test_rewrite_policy"
  action = "DROP"
  rule   = "HTTP.REQ.URL.PATH_AND_QUERY.CONTAINS(\"helloandby\")"
}

resource "citrixadc_lbvserver_rewritepolicy_binding" "tf_bind" {
    name = citrixadc_lbvserver.tf_lbvserver.name
    policyname = citrixadc_rewritepolicy.tf_rewrite_policy.name
    priority = 100
    bindpoint = "REQUEST"
}
`

// TestAccLbvserver_rewritepolicy_binding_sdkv2StateUpgrade verifies that a binding
// created with the last SDK v2 release (2.2.0, legacy comma-separated ID) is
// correctly refreshed/planned/applied by the current framework provider.
func TestAccLbvserver_rewritepolicy_binding_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckLbvserver_rewritepolicy_bindingDestroy,
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
				Config: testAcclbvserver_rewritepolicy_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbvserver_rewritepolicy_bindingExist("citrixadc_lbvserver_rewritepolicy_binding.tf_bind", nil),
					resource.TestCheckResourceAttr("citrixadc_lbvserver_rewritepolicy_binding.tf_bind", "id", "tf_lbvserver,tf_test_rewrite_policy"),
				),
			},
			// Step 2: same config, current (framework) provider. Terraform
			// refreshes the legacy-id state through the framework Read
			// (exercising ParseIdString on the legacy id) then plans/applies.
			// The framework recomputes the id on read to the new key:value form.
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAcclbvserver_rewritepolicy_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbvserver_rewritepolicy_bindingExist("citrixadc_lbvserver_rewritepolicy_binding.tf_bind", nil),
					resource.TestCheckResourceAttr("citrixadc_lbvserver_rewritepolicy_binding.tf_bind", "id", "name:tf_lbvserver,policyname:tf_test_rewrite_policy"),
				),
			},
		},
	})
}

func TestAccLbvserver_rewritepolicy_binding_import(t *testing.T) {
	const resAddr = "citrixadc_lbvserver_rewritepolicy_binding.tf_bind"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLbvserver_rewritepolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccLbvserver_rewritepolicy_binding_basic_step1},
			{Config: testAccLbvserver_rewritepolicy_binding_basic_step1, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{"bindpoint"}},
		},
	})
}
