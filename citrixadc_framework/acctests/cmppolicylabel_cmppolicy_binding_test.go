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

const testAccCmppolicylabel_cmppolicy_binding_basic = `

resource "citrixadc_cmppolicylabel_cmppolicy_binding" "tf_cmppolicylabel_cmppolicy_binding" {
	policyname = citrixadc_cmppolicy.tf_cmppolicy.name
	labelname  = citrixadc_cmppolicylabel.tf_cmppolicylabel.labelname
	priority   = 100
	}

  resource "citrixadc_cmppolicylabel" "tf_cmppolicylabel" {
	labelname = "my_cmppolicy_label"
	type      = "RES"
	}
  resource "citrixadc_cmppolicy" "tf_cmppolicy" {
	name      = "tf_cmppolicy"
	rule      = "HTTP.RES.HEADER(\"Content-Type\").CONTAINS(\"text\")"
	resaction = "COMPRESS"
	}
`

const testAccCmppolicylabel_cmppolicy_binding_basic_step2 = `

resource "citrixadc_cmppolicylabel" "tf_cmppolicylabel" {
	labelname = "my_cmppolicy_label"
	type      = "RES"
	}
resource "citrixadc_cmppolicy" "tf_cmppolicy" {
	name      = "tf_cmppolicy"
	rule      = "HTTP.RES.HEADER(\"Content-Type\").CONTAINS(\"text\")"
	resaction = "COMPRESS"
	}
`

func TestAccCmppolicylabel_cmppolicy_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCmppolicylabel_cmppolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCmppolicylabel_cmppolicy_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCmppolicylabel_cmppolicy_bindingExist("citrixadc_cmppolicylabel_cmppolicy_binding.tf_cmppolicylabel_cmppolicy_binding", nil),
				),
			},
			{
				Config: testAccCmppolicylabel_cmppolicy_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCmppolicylabel_cmppolicy_bindingNotExist("citrixadc_cmppolicylabel_cmppolicy_binding.tf_cmppolicylabel_cmppolicy_binding", "my_cmppolicy_label,tf_cmppolicy"),
				),
			},
		},
	})
}

// testAccCmppolicylabel_cmppolicy_binding_upgrade_basic is the config used by the
// sdkv2 -> framework state-upgrade test. It reuses the exact values from
// testAccCmppolicylabel_cmppolicy_binding_basic and only uses SDK v2 attribute names
// so it is valid under BOTH the last SDK v2 release (2.2.0) and the current framework schema.
const testAccCmppolicylabel_cmppolicy_binding_upgrade_basic = `

  resource "citrixadc_cmppolicylabel" "tf_cmppolicylabel" {
	labelname = "my_cmppolicy_label"
	type      = "RES"
	}
  resource "citrixadc_cmppolicy" "tf_cmppolicy" {
	name      = "tf_cmppolicy"
	rule      = "HTTP.RES.HEADER(\"Content-Type\").CONTAINS(\"text\")"
	resaction = "COMPRESS"
	}

resource "citrixadc_cmppolicylabel_cmppolicy_binding" "tf_cmppolicylabel_cmppolicy_binding" {
	policyname = citrixadc_cmppolicy.tf_cmppolicy.name
	labelname  = citrixadc_cmppolicylabel.tf_cmppolicylabel.labelname
	priority   = 100
	}
`

// TestAccCmppolicylabel_cmppolicy_binding_sdkv2StateUpgrade verifies that a binding
// created by the last SDK v2 release (which writes the legacy comma-joined id
// "my_cmppolicy_label,tf_cmppolicy") is correctly read/refreshed by the current
// framework provider, which recomputes the id to the new
// "labelname:...,policyname:..." format on Read (SetAttrFromGet).
func TestAccCmppolicylabel_cmppolicy_binding_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckCmppolicylabel_cmppolicy_bindingDestroy,
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
				Config: testAccCmppolicylabel_cmppolicy_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCmppolicylabel_cmppolicy_bindingExist("citrixadc_cmppolicylabel_cmppolicy_binding.tf_cmppolicylabel_cmppolicy_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_cmppolicylabel_cmppolicy_binding.tf_cmppolicylabel_cmppolicy_binding", "id", "my_cmppolicy_label,tf_cmppolicy"),
				),
			},
			// Step 2: same config through the CURRENT (framework) provider. Terraform
			// refreshes the legacy-id state (exercising ParseIdString on the legacy id),
			// and the framework Read recomputes the canonical new-format id.
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccCmppolicylabel_cmppolicy_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCmppolicylabel_cmppolicy_bindingExist("citrixadc_cmppolicylabel_cmppolicy_binding.tf_cmppolicylabel_cmppolicy_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_cmppolicylabel_cmppolicy_binding.tf_cmppolicylabel_cmppolicy_binding", "id", "labelname:my_cmppolicy_label,policyname:tf_cmppolicy"),
				),
			},
		},
	})
}

func testAccCheckCmppolicylabel_cmppolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No cmppolicylabel_cmppolicy_binding id is set")
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
			ResourceType:             "cmppolicylabel_cmppolicy_binding",
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
			return fmt.Errorf("cmppolicylabel_cmppolicy_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckCmppolicylabel_cmppolicy_bindingNotExist(n string, id string) resource.TestCheckFunc {
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
			ResourceType:             "cmppolicylabel_cmppolicy_binding",
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
			return fmt.Errorf("cmppolicylabel_cmppolicy_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckCmppolicylabel_cmppolicy_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_cmppolicylabel_cmppolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Cmppolicylabel_cmppolicy_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("cmppolicylabel_cmppolicy_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccCmppolicylabel_cmppolicy_bindingDataSource_basic = `

	resource "citrixadc_cmppolicylabel" "tf_cmppolicylabel" {
		labelname = "tf_cmppolicylabel_ds"
		type      = "RES"
	}

	resource "citrixadc_cmppolicy" "tf_cmppolicy" {
		name      = "tf_cmppolicy_ds"
		rule      = "HTTP.RES.HEADER(\"Content-Type\").CONTAINS(\"text\")"
		resaction = "COMPRESS"
	}

	resource "citrixadc_cmppolicylabel_cmppolicy_binding" "tf_cmppolicylabel_cmppolicy_binding" {
		labelname  = citrixadc_cmppolicylabel.tf_cmppolicylabel.labelname
		policyname = citrixadc_cmppolicy.tf_cmppolicy.name
		priority   = 100
	}

	data "citrixadc_cmppolicylabel_cmppolicy_binding" "tf_cmppolicylabel_cmppolicy_binding" {
		labelname  = citrixadc_cmppolicylabel_cmppolicy_binding.tf_cmppolicylabel_cmppolicy_binding.labelname
		policyname = citrixadc_cmppolicylabel_cmppolicy_binding.tf_cmppolicylabel_cmppolicy_binding.policyname
		depends_on = [citrixadc_cmppolicylabel_cmppolicy_binding.tf_cmppolicylabel_cmppolicy_binding]
	}
`

func TestAcccmppolicylabel_cmppolicy_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccCmppolicylabel_cmppolicy_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_cmppolicylabel_cmppolicy_binding.tf_cmppolicylabel_cmppolicy_binding", "labelname", "tf_cmppolicylabel_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_cmppolicylabel_cmppolicy_binding.tf_cmppolicylabel_cmppolicy_binding", "policyname", "tf_cmppolicy_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_cmppolicylabel_cmppolicy_binding.tf_cmppolicylabel_cmppolicy_binding", "priority", "100"),
				),
			},
		},
	})
}

func TestAccCmppolicylabel_cmppolicy_binding_import(t *testing.T) {
	const resAddr = "citrixadc_cmppolicylabel_cmppolicy_binding.tf_cmppolicylabel_cmppolicy_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCmppolicylabel_cmppolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccCmppolicylabel_cmppolicy_binding_basic},
			{Config: testAccCmppolicylabel_cmppolicy_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
		},
	})
}
