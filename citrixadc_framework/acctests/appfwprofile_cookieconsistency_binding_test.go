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

const testAccAppfwprofile_cookieconsistency_binding_basic_step1 = `
	resource citrixadc_appfwprofile_cookieconsistency_binding demo_binding1 {
		name              = citrixadc_appfwprofile.demo_appfw.name
		cookieconsistency = "^logon_[0-9A-Za-z]{2,15}$"
	}
	resource citrixadc_appfwprofile_cookieconsistency_binding demo_binding2 {
		name              = citrixadc_appfwprofile.demo_appfw.name
		cookieconsistency = "^logout_[0-9A-Za-z]{2,15}$"
	}

	resource citrixadc_appfwprofile demo_appfw {
		name                     = "demo_appfwprofile"
		type                     = ["HTML"]
	}
`

const testAccAppfwprofile_cookieconsistency_binding_basic_step2 = `
	resource citrixadc_appfwprofile demo_appfw {
		name                     = "demo_appfwprofile"
		type                     = ["HTML"]
	}
`

func TestAccAppfwprofile_cookieconsistency_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAppfwprofile_cookieconsistency_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAppfwprofile_cookieconsistency_binding_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwprofile_cookieconsistency_bindingExist("citrixadc_appfwprofile_cookieconsistency_binding.demo_binding1", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_cookieconsistency_binding.demo_binding1", "name", "demo_appfwprofile"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_cookieconsistency_binding.demo_binding1", "cookieconsistency", "^logon_[0-9A-Za-z]{2,15}$"),
					testAccCheckAppfwprofile_cookieconsistency_bindingExist("citrixadc_appfwprofile_cookieconsistency_binding.demo_binding2", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_cookieconsistency_binding.demo_binding2", "name", "demo_appfwprofile"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_cookieconsistency_binding.demo_binding2", "cookieconsistency", "^logout_[0-9A-Za-z]{2,15}$"),
				),
			},
			{
				Config: testAccAppfwprofile_cookieconsistency_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwprofile_cookieconsistency_bindingNotExist("citrixadc_appfwprofile_cookieconsistency_binding.demo_binding1", "demo_appfwprofile,^logon_[0-9A-Za-z]{2,15}$"),
					testAccCheckAppfwprofile_cookieconsistency_bindingNotExist("citrixadc_appfwprofile_cookieconsistency_binding.demo_binding2", "demo_appfwprofile,^logout_[0-9A-Za-z]{2,15}$"),
				),
			},
		},
	})
}

func testAccCheckAppfwprofile_cookieconsistency_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No appfwprofile_cookieconsistency_binding name is set")
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
		idMap, _, err := utils.ParseIdString(bindingId, []string{"name", "cookieconsistency"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %v: %v", bindingId, err)
		}
		appFwName := idMap["name"]
		cookieconsistency := idMap["cookieconsistency"]

		findParams := service.FindParams{
			ResourceType:             service.Appfwprofile_cookieconsistency_binding.Type(),
			ResourceName:             appFwName,
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
			if v["cookieconsistency"].(string) == cookieconsistency {
				foundIndex = i
				break
			}
		}

		// Resource is missing
		if foundIndex == -1 {
			return fmt.Errorf("Cannot find appfwprofile_cookieconsistency_binding ID %v", bindingId)
		}

		return nil
	}

}

func testAccCheckAppfwprofile_cookieconsistency_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}
		idSlice := strings.SplitN(id, ",", 3)

		name := idSlice[0]
		cookieconsistency := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "appfwprofile_cookieconsistency_binding",
			ResourceName:             name,
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
			if v["cookieconsistency"].(string) == cookieconsistency {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("appfwprofile_cookieconsistency_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckAppfwprofile_cookieconsistency_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_appfwprofile_cookieconsistency_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Appfwprofile_cookieconsistency_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("appfwprofile_cookieconsistency_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccAppfwprofile_cookieconsistency_bindingDataSource_basic = `
	resource citrixadc_appfwprofile demo_appfw {
		name                     = "demo_appfwprofile"
		type                     = ["HTML"]
	}
	resource citrixadc_appfwprofile_cookieconsistency_binding demo_binding1 {
		name              = citrixadc_appfwprofile.demo_appfw.name
		cookieconsistency = "^logon_[0-9A-Za-z]{2,15}$"
	}

	data "citrixadc_appfwprofile_cookieconsistency_binding" "demo_binding1" {
		name              = citrixadc_appfwprofile_cookieconsistency_binding.demo_binding1.name
		cookieconsistency = citrixadc_appfwprofile_cookieconsistency_binding.demo_binding1.cookieconsistency
	}
`

func TestAccAppfwprofile_cookieconsistency_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAppfwprofile_cookieconsistency_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_cookieconsistency_binding.demo_binding1", "name", "demo_appfwprofile"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_cookieconsistency_binding.demo_binding1", "cookieconsistency", "^logon_[0-9A-Za-z]{2,15}$"),
				),
			},
		},
	})
}

// testAccAppfwprofile_cookieconsistency_binding_upgrade_basic reuses the _basic config
// (the binding + its appfwprofile prerequisite) with the same values. It is valid under
// BOTH the SDK v2 2.2.0 schema and the current Framework schema because the migration
// restored the SDK v2 attribute names (name, cookieconsistency).
const testAccAppfwprofile_cookieconsistency_binding_upgrade_basic = `
	resource citrixadc_appfwprofile_cookieconsistency_binding demo_binding1 {
		name              = citrixadc_appfwprofile.demo_appfw.name
		cookieconsistency = "^logon_[0-9A-Za-z]{2,15}$"
	}

	resource citrixadc_appfwprofile demo_appfw {
		name = "demo_appfwprofile"
		type = ["HTML"]
	}
`

// TestAccAppfwprofile_cookieconsistency_binding_sdkv2StateUpgrade verifies that state
// written by the last SDK v2 release is correctly upgraded when the same config is
// subsequently managed by the current Framework provider.
//
// Step 1 creates the binding with citrix/citrixadc 2.2.0, which writes the legacy
// comma-joined id (SDK v2 d.SetId(fmt.Sprintf("%s,%s", name, cookieconsistency)) =>
// "demo_appfwprofile,^logon_[0-9A-Za-z]{2,15}$").
// Step 2 refreshes/plans/applies the SAME config through the current Framework provider,
// exercising ParseIdString on the legacy id. The Framework recomputes the id on Read
// (appfwprofile_cookieconsistency_bindingSetAttrFromGet), so the id is upgraded to the
// canonical new key:value format
// "cookieconsistency:%5Elogon_%5B0-9A-Za-z%5D%7B2%2C15%7D%24,name:demo_appfwprofile".
func TestAccAppfwprofile_cookieconsistency_binding_sdkv2StateUpgrade(t *testing.T) {
	resourceAddr := "citrixadc_appfwprofile_cookieconsistency_binding.demo_binding1"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckAppfwprofile_cookieconsistency_bindingDestroy,
		Steps: []resource.TestStep{
			// Step 1: create with the last SDK v2 release -> state carries the legacy id.
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.2.0",
					},
				},
				Config: testAccAppfwprofile_cookieconsistency_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwprofile_cookieconsistency_bindingExist(resourceAddr, nil),
					resource.TestCheckResourceAttr(resourceAddr, "id", "demo_appfwprofile,^logon_[0-9A-Za-z]{2,15}$"),
				),
			},
			// Step 2: refresh/plan/apply the SAME config through the current Framework
			// provider. The legacy-id state is read via ParseIdString and the id is
			// recomputed on Read to the canonical new key:value format.
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccAppfwprofile_cookieconsistency_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwprofile_cookieconsistency_bindingExist(resourceAddr, nil),
					resource.TestCheckResourceAttr(resourceAddr, "id", "cookieconsistency:%5Elogon_%5B0-9A-Za-z%5D%7B2%2C15%7D%24,name:demo_appfwprofile"),
				),
			},
		},
	})
}

func TestAccAppfwprofile_cookieconsistency_binding_import(t *testing.T) {
	const resAddr = "citrixadc_appfwprofile_cookieconsistency_binding.demo_binding1"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAppfwprofile_cookieconsistency_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccAppfwprofile_cookieconsistency_binding_basic_step1},
			{Config: testAccAppfwprofile_cookieconsistency_binding_basic_step1, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
		},
	})
}
