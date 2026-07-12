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

const testAccLbvserver_transformpolicy_binding_basic_step1 = `
resource "citrixadc_lbvserver" "tf_lbvserver" {
  ipv46       = "10.10.10.33"
  name        = "tf_lbvserver"
  port        = 80
  servicetype = "HTTP"
}

resource "citrixadc_transformprofile" "tf_trans_profile" {
  name = "tf_trans_profile"
  comment = "Some comment"
}

resource "citrixadc_transformpolicy" "tf_trans_policy" {
    name = "tf_trans_policy"
    profilename = citrixadc_transformprofile.tf_trans_profile.name
    rule = "http.REQ.URL.CONTAINS(\"test_url\")"
}

resource "citrixadc_lbvserver_transformpolicy_binding" "tf_binding" {
    name = citrixadc_lbvserver.tf_lbvserver.name
    policyname = citrixadc_transformpolicy.tf_trans_policy.name
    priority = 100
    bindpoint = "REQUEST"
    gotopriorityexpression = "END"
}
`

const testAccLbvserver_transformpolicy_binding_basic_step2 = `
resource "citrixadc_lbvserver" "tf_lbvserver" {
  ipv46       = "10.10.10.33"
  name        = "tf_lbvserver"
  port        = 80
  servicetype = "HTTP"
}

resource "citrixadc_transformprofile" "tf_trans_profile" {
  name = "tf_trans_profile"
  comment = "Some comment"
}

resource "citrixadc_transformpolicy" "tf_trans_policy" {
    name = "tf_trans_policy"
    profilename = citrixadc_transformprofile.tf_trans_profile.name
    rule = "http.REQ.URL.CONTAINS(\"test_url\")"
}

resource "citrixadc_lbvserver_transformpolicy_binding" "tf_binding" {
    name = citrixadc_lbvserver.tf_lbvserver.name
    policyname = citrixadc_transformpolicy.tf_trans_policy.name
    priority = 110
    bindpoint = "REQUEST"
    gotopriorityexpression = "NEXT"
}
`

func TestAccLbvserver_transformpolicy_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLbvserver_transformpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLbvserver_transformpolicy_binding_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbvserver_transformpolicy_bindingExist("citrixadc_lbvserver_transformpolicy_binding.tf_binding", nil),
				),
			},
			{
				Config: testAccLbvserver_transformpolicy_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbvserver_transformpolicy_bindingExist("citrixadc_lbvserver_transformpolicy_binding.tf_binding", nil),
				),
			},
		},
	})
}

func TestAccLbvserver_transformpolicy_binding_import(t *testing.T) {
	const resAddr = "citrixadc_lbvserver_transformpolicy_binding.tf_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLbvserver_transformpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccLbvserver_transformpolicy_binding_basic_step1},
			{Config: testAccLbvserver_transformpolicy_binding_basic_step1, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
		},
	})
}

func testAccCheckLbvserver_transformpolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lbvserver_transformpolicy_binding name is set")
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
			ResourceType:             "lbvserver_transformpolicy_binding",
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
			return fmt.Errorf("lbvserver_transformpolicy_binding %s not found", bindingId)
		}
		return nil
	}
}

func testAccCheckLbvserver_transformpolicy_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_lbvserver_transformpolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Lbvserver_transformpolicy_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("lbvserver_transformpolicy_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccLbvserver_transformpolicy_bindingDataSource_basic = `
resource "citrixadc_lbvserver" "tf_lbvserver" {
  ipv46       = "10.10.10.33"
  name        = "tf_lbvserver"
  port        = 80
  servicetype = "HTTP"
}

resource "citrixadc_transformprofile" "tf_trans_profile" {
  name = "tf_trans_profile"
  comment = "Some comment"
}

resource "citrixadc_transformpolicy" "tf_trans_policy" {
    name = "tf_trans_policy"
    profilename = citrixadc_transformprofile.tf_trans_profile.name
    rule = "http.REQ.URL.CONTAINS(\"test_url\")"
}

resource "citrixadc_lbvserver_transformpolicy_binding" "tf_binding" {
    name = citrixadc_lbvserver.tf_lbvserver.name
    policyname = citrixadc_transformpolicy.tf_trans_policy.name
    priority = 100
    bindpoint = "REQUEST"
    gotopriorityexpression = "END"
}

data "citrixadc_lbvserver_transformpolicy_binding" "tf_binding" {
    name = citrixadc_lbvserver_transformpolicy_binding.tf_binding.name
    policyname = citrixadc_lbvserver_transformpolicy_binding.tf_binding.policyname
    depends_on = [citrixadc_lbvserver_transformpolicy_binding.tf_binding]
}
`

func TestAccLbvserver_transformpolicy_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccLbvserver_transformpolicy_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_lbvserver_transformpolicy_binding.tf_binding", "name", "tf_lbvserver"),
					resource.TestCheckResourceAttr("data.citrixadc_lbvserver_transformpolicy_binding.tf_binding", "policyname", "tf_trans_policy"),
					resource.TestCheckResourceAttr("data.citrixadc_lbvserver_transformpolicy_binding.tf_binding", "priority", "100"),
					resource.TestCheckResourceAttr("data.citrixadc_lbvserver_transformpolicy_binding.tf_binding", "gotopriorityexpression", "END"),
				),
			},
		},
	})
}

const testAccLbvserver_transformpolicy_binding_upgrade_basic = `
resource "citrixadc_lbvserver" "tf_lbvserver" {
  ipv46       = "10.10.10.33"
  name        = "tf_lbvserver"
  port        = 80
  servicetype = "HTTP"
}

resource "citrixadc_transformprofile" "tf_trans_profile" {
  name = "tf_trans_profile"
  comment = "Some comment"
}

resource "citrixadc_transformpolicy" "tf_trans_policy" {
    name = "tf_trans_policy"
    profilename = citrixadc_transformprofile.tf_trans_profile.name
    rule = "http.REQ.URL.CONTAINS(\"test_url\")"
}

resource "citrixadc_lbvserver_transformpolicy_binding" "tf_binding" {
    name = citrixadc_lbvserver.tf_lbvserver.name
    policyname = citrixadc_transformpolicy.tf_trans_policy.name
    priority = 100
    bindpoint = "REQUEST"
    gotopriorityexpression = "END"
}
`

// TestAccLbvserver_transformpolicy_binding_sdkv2StateUpgrade verifies that a resource
// created with the last SDK v2 release (2.2.0, legacy comma-separated ID) is correctly
// adopted by the current Framework provider: the Framework Read parses the legacy ID via
// ParseIdString and recomputes it to the canonical new key:value format.
func TestAccLbvserver_transformpolicy_binding_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckLbvserver_transformpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			// Step 1: create with the last SDK v2 release -> legacy id "name,policyname".
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.2.0",
					},
				},
				Config: testAccLbvserver_transformpolicy_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbvserver_transformpolicy_bindingExist("citrixadc_lbvserver_transformpolicy_binding.tf_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_lbvserver_transformpolicy_binding.tf_binding", "id", "tf_lbvserver,tf_trans_policy"),
				),
			},
			// Step 2: refresh/plan/apply through the current Framework provider. Read
			// parses the legacy id and recomputes it to the new key:value format.
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccLbvserver_transformpolicy_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbvserver_transformpolicy_bindingExist("citrixadc_lbvserver_transformpolicy_binding.tf_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_lbvserver_transformpolicy_binding.tf_binding", "id", "name:tf_lbvserver,policyname:tf_trans_policy"),
				),
			},
		},
	})
}
