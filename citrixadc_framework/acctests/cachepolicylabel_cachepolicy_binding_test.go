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

const testAccCachepolicylabel_cachepolicy_binding_basic = `

	resource "citrixadc_cachepolicylabel" "tf_policylabel" {
		labelname = "my_cachepolicylabel"
		evaluates = "REQ"
	}
	resource "citrixadc_cachepolicy" "tf_cachepolicy" {
		policyname  = "my_cachepolicy"
		rule        = "true"
		action      = "CACHE"
	}

	resource "citrixadc_cachepolicylabel_cachepolicy_binding" "tf_cachepolicylabel_cachepolicy_binding" {
		labelname  = citrixadc_cachepolicylabel.tf_policylabel.labelname
		priority   = 100
		policyname = citrixadc_cachepolicy.tf_cachepolicy.policyname
	}
  
  
`

const testAccCachepolicylabel_cachepolicy_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion

	resource "citrixadc_cachepolicylabel" "tf_policylabel" {
		labelname = "my_cachepolicylabel"
		evaluates = "REQ"
	}
	resource "citrixadc_cachepolicy" "tf_cachepolicy" {
		policyname  = "my_cachepolicy"
		rule        = "true"
		action      = "CACHE"
	}
`

func TestAccCachepolicylabel_cachepolicy_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCachepolicylabel_cachepolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCachepolicylabel_cachepolicy_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCachepolicylabel_cachepolicy_bindingExist("citrixadc_cachepolicylabel_cachepolicy_binding.tf_cachepolicylabel_cachepolicy_binding", nil),
				),
			},
			{
				Config: testAccCachepolicylabel_cachepolicy_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCachepolicylabel_cachepolicy_bindingNotExist("citrixadc_cachepolicylabel_cachepolicy_binding.tf_cachepolicylabel_cachepolicy_binding", "my_cachepolicylabel,my_cachepolicy"),
				),
			},
		},
	})
}

func TestAccCachepolicylabel_cachepolicy_binding_import(t *testing.T) {
	const resAddr = "citrixadc_cachepolicylabel_cachepolicy_binding.tf_cachepolicylabel_cachepolicy_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCachepolicylabel_cachepolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccCachepolicylabel_cachepolicy_binding_basic},
			{Config: testAccCachepolicylabel_cachepolicy_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
		},
	})
}

// testAccCachepolicylabel_cachepolicy_binding_upgrade_basic is the config used by the
// sdkv2 -> framework state-upgrade test. It reuses the exact values from
// testAccCachepolicylabel_cachepolicy_binding_basic and only uses SDK v2 attribute names
// so it is valid under BOTH the last SDK v2 release (2.2.0) and the current framework schema.
const testAccCachepolicylabel_cachepolicy_binding_upgrade_basic = `

	resource "citrixadc_cachepolicylabel" "tf_policylabel" {
		labelname = "my_cachepolicylabel"
		evaluates = "REQ"
	}
	resource "citrixadc_cachepolicy" "tf_cachepolicy" {
		policyname  = "my_cachepolicy"
		rule        = "true"
		action      = "CACHE"
	}

	resource "citrixadc_cachepolicylabel_cachepolicy_binding" "tf_cachepolicylabel_cachepolicy_binding" {
		labelname  = citrixadc_cachepolicylabel.tf_policylabel.labelname
		priority   = 100
		policyname = citrixadc_cachepolicy.tf_cachepolicy.policyname
	}
`

// TestAccCachepolicylabel_cachepolicy_binding_sdkv2StateUpgrade verifies that a binding
// created by the last SDK v2 release (which writes the legacy comma-joined id
// "my_cachepolicylabel,my_cachepolicy") is correctly read/refreshed by the current
// framework provider, which recomputes the id to the new
// "labelname:...,policyname:..." format on Read (SetAttrFromGet).
func TestAccCachepolicylabel_cachepolicy_binding_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckCachepolicylabel_cachepolicy_bindingDestroy,
		Steps: []resource.TestStep{
			// Step 1: create with the LAST SDK v2 release from the registry.
			// State is written with the legacy id.
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.2.0",
					},
				},
				Config: testAccCachepolicylabel_cachepolicy_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCachepolicylabel_cachepolicy_bindingExist("citrixadc_cachepolicylabel_cachepolicy_binding.tf_cachepolicylabel_cachepolicy_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_cachepolicylabel_cachepolicy_binding.tf_cachepolicylabel_cachepolicy_binding", "id", "my_cachepolicylabel,my_cachepolicy"),
				),
			},
			// Step 2: same config through the CURRENT (framework) provider. Terraform
			// refreshes the legacy-id state (exercising ParseIdString on the legacy id),
			// and the framework Read recomputes the canonical new-format id.
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccCachepolicylabel_cachepolicy_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCachepolicylabel_cachepolicy_bindingExist("citrixadc_cachepolicylabel_cachepolicy_binding.tf_cachepolicylabel_cachepolicy_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_cachepolicylabel_cachepolicy_binding.tf_cachepolicylabel_cachepolicy_binding", "id", "labelname:my_cachepolicylabel,policyname:my_cachepolicy"),
				),
			},
		},
	})
}

func testAccCheckCachepolicylabel_cachepolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No cachepolicylabel_cachepolicy_binding id is set")
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
			return err
		}
		labelname := idMap["labelname"]
		policyname := idMap["policyname"]

		findParams := service.FindParams{
			ResourceType:             "cachepolicylabel_cachepolicy_binding",
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
			return fmt.Errorf("cachepolicylabel_cachepolicy_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckCachepolicylabel_cachepolicy_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}
		idMap, _, err := utils.ParseIdString(id, []string{"labelname", "policyname"}, nil)
		if err != nil {
			return err
		}
		labelname := idMap["labelname"]
		policyname := idMap["policyname"]

		findParams := service.FindParams{
			ResourceType:             "cachepolicylabel_cachepolicy_binding",
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
			return fmt.Errorf("cachepolicylabel_cachepolicy_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckCachepolicylabel_cachepolicy_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_cachepolicylabel_cachepolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Cachepolicylabel_cachepolicy_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("cachepolicylabel_cachepolicy_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccCachepolicylabel_cachepolicy_bindingDataSource_basic = `

	resource "citrixadc_cachepolicylabel" "tf_policylabel" {
		labelname = "my_cachepolicylabel"
		evaluates = "REQ"
	}
	resource "citrixadc_cachepolicy" "tf_cachepolicy" {
		policyname  = "my_cachepolicy"
		rule        = "true"
		action      = "CACHE"
	}

	resource "citrixadc_cachepolicylabel_cachepolicy_binding" "tf_cachepolicylabel_cachepolicy_binding" {
		labelname  = citrixadc_cachepolicylabel.tf_policylabel.labelname
		priority   = 100
		policyname = citrixadc_cachepolicy.tf_cachepolicy.policyname
	}

	data "citrixadc_cachepolicylabel_cachepolicy_binding" "tf_cachepolicylabel_cachepolicy_binding" {
		labelname  = citrixadc_cachepolicylabel_cachepolicy_binding.tf_cachepolicylabel_cachepolicy_binding.labelname
		policyname = citrixadc_cachepolicylabel_cachepolicy_binding.tf_cachepolicylabel_cachepolicy_binding.policyname
		depends_on = [citrixadc_cachepolicylabel_cachepolicy_binding.tf_cachepolicylabel_cachepolicy_binding]
	}
`

func TestAcccachepolicylabel_cachepolicy_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCachepolicylabel_cachepolicy_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_cachepolicylabel_cachepolicy_binding.tf_cachepolicylabel_cachepolicy_binding", "labelname", "my_cachepolicylabel"),
					resource.TestCheckResourceAttr("data.citrixadc_cachepolicylabel_cachepolicy_binding.tf_cachepolicylabel_cachepolicy_binding", "policyname", "my_cachepolicy"),
					resource.TestCheckResourceAttr("data.citrixadc_cachepolicylabel_cachepolicy_binding.tf_cachepolicylabel_cachepolicy_binding", "priority", "100"),
				),
			},
		},
	})
}
