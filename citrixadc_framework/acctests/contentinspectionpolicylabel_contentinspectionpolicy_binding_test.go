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

const testAccContentinspectionpolicylabel_contentinspectionpolicy_binding_basic = `

	resource "citrixadc_contentinspectionpolicylabel" "tf_contentinspectionpolicylabel" {
		labelname = "my_ci_label"
		type      = "RES"
	}
	resource "citrixadc_contentinspectionpolicy" "tf_contentinspectionpolicy" {
		name   = "my_ci_policy"
		rule   = "false"
		action = "DROP"
	}

	resource "citrixadc_contentinspectionpolicylabel_contentinspectionpolicy_binding" "tf_ci_binding" {
		labelname  = citrixadc_contentinspectionpolicylabel.tf_contentinspectionpolicylabel.labelname
		policyname = citrixadc_contentinspectionpolicy.tf_contentinspectionpolicy.name
		priority   = 100
	}
  
`

const testAccContentinspectionpolicylabel_contentinspectionpolicy_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion

	resource "citrixadc_contentinspectionpolicylabel" "tf_contentinspectionpolicylabel" {
		labelname = "my_ci_label"
		type      = "RES"
	}
	resource "citrixadc_contentinspectionpolicy" "tf_contentinspectionpolicy" {
		name   = "my_ci_policy"
		rule   = "false"
		action = "DROP"
	}
`

func TestAccContentinspectionpolicylabel_contentinspectionpolicy_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckContentinspectionpolicylabel_contentinspectionpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccContentinspectionpolicylabel_contentinspectionpolicy_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckContentinspectionpolicylabel_contentinspectionpolicy_bindingExist("citrixadc_contentinspectionpolicylabel_contentinspectionpolicy_binding.tf_ci_binding", nil),
				),
			},
			{
				Config: testAccContentinspectionpolicylabel_contentinspectionpolicy_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckContentinspectionpolicylabel_contentinspectionpolicy_bindingNotExist("citrixadc_contentinspectionpolicylabel_contentinspectionpolicy_binding.tf_ci_binding", "my_ci_label,my_ci_policy"),
				),
			},
		},
	})
}

func testAccCheckContentinspectionpolicylabel_contentinspectionpolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No contentinspectionpolicylabel_contentinspectionpolicy_binding id is set")
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

		// Parse both the new key:value ID format and the legacy comma format.
		idMap, _, err := utils.ParseIdString(bindingId, []string{"labelname", "policyname"}, nil)
		if err != nil {
			return err
		}
		labelname := idMap["labelname"]
		policyname := idMap["policyname"]

		findParams := service.FindParams{
			ResourceType:             "contentinspectionpolicylabel_contentinspectionpolicy_binding",
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
			return fmt.Errorf("contentinspectionpolicylabel_contentinspectionpolicy_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckContentinspectionpolicylabel_contentinspectionpolicy_bindingNotExist(n string, id string) resource.TestCheckFunc {
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

		labelname := idSlice[0]
		policyname := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "contentinspectionpolicylabel_contentinspectionpolicy_binding",
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
			return fmt.Errorf("contentinspectionpolicylabel_contentinspectionpolicy_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckContentinspectionpolicylabel_contentinspectionpolicy_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_contentinspectionpolicylabel_contentinspectionpolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("contentinspectionpolicylabel_contentinspectionpolicy_binding", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("contentinspectionpolicylabel_contentinspectionpolicy_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccContentinspectionpolicylabel_contentinspectionpolicy_bindingDataSource_basic = `

	resource "citrixadc_contentinspectionpolicylabel" "tf_contentinspectionpolicylabel" {
		labelname = "tf_contentinspectionpolicylabel_ds"
		type      = "RES"
	}

	resource "citrixadc_contentinspectionpolicy" "tf_contentinspectionpolicy" {
		name   = "tf_contentinspectionpolicy_ds"
		rule   = "false"
		action = "DROP"
	}

	resource "citrixadc_contentinspectionpolicylabel_contentinspectionpolicy_binding" "tf_contentinspectionpolicylabel_contentinspectionpolicy_binding" {
		labelname  = citrixadc_contentinspectionpolicylabel.tf_contentinspectionpolicylabel.labelname
		policyname = citrixadc_contentinspectionpolicy.tf_contentinspectionpolicy.name
		priority   = 100
	}

	data "citrixadc_contentinspectionpolicylabel_contentinspectionpolicy_binding" "tf_contentinspectionpolicylabel_contentinspectionpolicy_binding" {
		labelname  = citrixadc_contentinspectionpolicylabel_contentinspectionpolicy_binding.tf_contentinspectionpolicylabel_contentinspectionpolicy_binding.labelname
		policyname = citrixadc_contentinspectionpolicylabel_contentinspectionpolicy_binding.tf_contentinspectionpolicylabel_contentinspectionpolicy_binding.policyname
		depends_on = [citrixadc_contentinspectionpolicylabel_contentinspectionpolicy_binding.tf_contentinspectionpolicylabel_contentinspectionpolicy_binding]
	}
`

func TestAcccontentinspectionpolicylabel_contentinspectionpolicy_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccContentinspectionpolicylabel_contentinspectionpolicy_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_contentinspectionpolicylabel_contentinspectionpolicy_binding.tf_contentinspectionpolicylabel_contentinspectionpolicy_binding", "labelname", "tf_contentinspectionpolicylabel_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_contentinspectionpolicylabel_contentinspectionpolicy_binding.tf_contentinspectionpolicylabel_contentinspectionpolicy_binding", "policyname", "tf_contentinspectionpolicy_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_contentinspectionpolicylabel_contentinspectionpolicy_binding.tf_contentinspectionpolicylabel_contentinspectionpolicy_binding", "priority", "100"),
				),
			},
		},
	})
}

// testAccContentinspectionpolicylabel_contentinspectionpolicy_binding_upgrade_basic is the config used by the
// sdkv2 -> framework state-upgrade test. It reuses the same values and resource
// labels as testAccContentinspectionpolicylabel_contentinspectionpolicy_binding_basic so it is valid under BOTH
// the SDK v2 2.2.0 schema and the current framework schema.
const testAccContentinspectionpolicylabel_contentinspectionpolicy_binding_upgrade_basic = `

	resource "citrixadc_contentinspectionpolicylabel" "tf_contentinspectionpolicylabel" {
		labelname = "my_ci_label"
		type      = "RES"
	}
	resource "citrixadc_contentinspectionpolicy" "tf_contentinspectionpolicy" {
		name   = "my_ci_policy"
		rule   = "false"
		action = "DROP"
	}

	resource "citrixadc_contentinspectionpolicylabel_contentinspectionpolicy_binding" "tf_ci_binding" {
		labelname  = citrixadc_contentinspectionpolicylabel.tf_contentinspectionpolicylabel.labelname
		policyname = citrixadc_contentinspectionpolicy.tf_contentinspectionpolicy.name
		priority   = 100
	}

`

// TestAccContentinspectionpolicylabel_contentinspectionpolicy_binding_sdkv2StateUpgrade verifies that a binding
// created with the last SDK v2 release (2.2.0, legacy comma-separated ID) is
// correctly refreshed/planned/applied by the current framework provider.
func TestAccContentinspectionpolicylabel_contentinspectionpolicy_binding_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckContentinspectionpolicylabel_contentinspectionpolicy_bindingDestroy,
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
				Config: testAccContentinspectionpolicylabel_contentinspectionpolicy_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckContentinspectionpolicylabel_contentinspectionpolicy_bindingExist("citrixadc_contentinspectionpolicylabel_contentinspectionpolicy_binding.tf_ci_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_contentinspectionpolicylabel_contentinspectionpolicy_binding.tf_ci_binding", "id", "my_ci_label,my_ci_policy"),
				),
			},
			// Step 2: same config, current (framework) provider. Terraform
			// refreshes the legacy-id state through the framework Read
			// (exercising ParseIdString on the legacy id) then plans/applies.
			// The framework recomputes the id on read to the new key:value form.
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccContentinspectionpolicylabel_contentinspectionpolicy_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckContentinspectionpolicylabel_contentinspectionpolicy_bindingExist("citrixadc_contentinspectionpolicylabel_contentinspectionpolicy_binding.tf_ci_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_contentinspectionpolicylabel_contentinspectionpolicy_binding.tf_ci_binding", "id", "labelname:my_ci_label,policyname:my_ci_policy"),
				),
			},
		},
	})
}

func TestAccContentinspectionpolicylabel_contentinspectionpolicy_binding_import(t *testing.T) {
	const resAddr = "citrixadc_contentinspectionpolicylabel_contentinspectionpolicy_binding.tf_ci_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckContentinspectionpolicylabel_contentinspectionpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccContentinspectionpolicylabel_contentinspectionpolicy_binding_basic},
			{Config: testAccContentinspectionpolicylabel_contentinspectionpolicy_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
		},
	})
}
