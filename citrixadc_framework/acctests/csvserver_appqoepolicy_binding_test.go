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

const testAccCsvserver_appqoepolicy_binding_basic = `

	resource "citrixadc_appqoeaction" "tf_appqoeaction" {
		name        = "tf_appqoeaction"
		priority    = "LOW"
		respondwith = "NS"
		delay       = 40
	}
	resource "citrixadc_appqoepolicy" "tf_appqoepolicy" {
		name   = "tf_appqoepolicy"
		rule   = "true"
		action = citrixadc_appqoeaction.tf_appqoeaction.name
	}

	resource "citrixadc_csvserver_appqoepolicy_binding" "tf_csvserver_appqoepolicy_binding" {
		name 		= citrixadc_csvserver.tf_csvserver.name
		policyname 	= citrixadc_appqoepolicy.tf_appqoepolicy.name
		bindpoint 	= "REQUEST"
		priority 	= 5
	}

	resource "citrixadc_csvserver" "tf_csvserver" {
		name        = "tf_csvserver"
		ipv46       = "10.10.10.33"
		port        = 80
		servicetype = "HTTP"
	}
`

const testAccCsvserver_appqoepolicy_binding_basic_step2 = `

	resource "citrixadc_appqoeaction" "tf_appqoeaction" {
		name        = "tf_appqoeaction"
		priority    = "LOW"
		respondwith = "NS"
		delay       = 40
	}
	resource "citrixadc_appqoepolicy" "tf_appqoepolicy" {
		name   = "tf_appqoepolicy"
		rule   = "true"
		action = citrixadc_appqoeaction.tf_appqoeaction.name
	}
	resource "citrixadc_csvserver" "tf_csvserver" {
		name        = "tf_csvserver"
		ipv46       = "10.10.10.33"
		port        = 80
		servicetype = "HTTP"
	}
`

func TestAccCsvserver_appqoepolicy_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCsvserver_appqoepolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCsvserver_appqoepolicy_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCsvserver_appqoepolicy_bindingExist("citrixadc_csvserver_appqoepolicy_binding.tf_csvserver_appqoepolicy_binding", nil),
				),
			},
			{
				Config: testAccCsvserver_appqoepolicy_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCsvserver_appqoepolicy_bindingNotExist("citrixadc_csvserver_appqoepolicy_binding.tf_csvserver_appqoepolicy_binding", "tf_csvserver,tf_appqoepolicy"),
				),
			},
		},
	})
}

func testAccCheckCsvserver_appqoepolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No csvserver_appqoepolicy_binding id is set")
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
			return fmt.Errorf("Error parsing ID: %v", err)
		}
		name := idMap["name"]
		policyname := idMap["policyname"]

		findParams := service.FindParams{
			ResourceType:             "csvserver_appqoepolicy_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the right monitor name
		found := false
		for _, v := range dataArr {
			if v["policyname"].(string) == policyname {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("csvserver_appqoepolicy_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckCsvserver_appqoepolicy_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}
		idMap, _, err := utils.ParseIdString(id, []string{"name", "policyname"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID: %v", err)
		}
		name := idMap["name"]
		policyName := idMap["policyname"]

		findParams := service.FindParams{
			ResourceType:             "csvserver_appqoepolicy_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the right policy name
		found := false
		for _, v := range dataArr {
			if v["policyname"].(string) == policyName {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("Csvserver_appqoepolicy_binding %s was found, but is should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckCsvserver_appqoepolicy_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_csvserver_appqoepolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("csvserver_appqoepolicy_binding", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("csvserver_appqoepolicy_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccCsvserver_appqoepolicy_bindingDataSource_basic = `

	resource "citrixadc_appqoeaction" "tf_appqoeaction" {
		name        = "tf_appqoeaction"
		priority    = "LOW"
		respondwith = "NS"
		delay       = 40
	}
	resource "citrixadc_appqoepolicy" "tf_appqoepolicy" {
		name   = "tf_appqoepolicy"
		rule   = "true"
		action = citrixadc_appqoeaction.tf_appqoeaction.name
	}

	resource "citrixadc_csvserver_appqoepolicy_binding" "tf_csvserver_appqoepolicy_binding" {
		name 		= citrixadc_csvserver.tf_csvserver.name
		policyname 	= citrixadc_appqoepolicy.tf_appqoepolicy.name
		bindpoint 	= "REQUEST"
		priority 	= 5
	}

	resource "citrixadc_csvserver" "tf_csvserver" {
		name        = "tf_csvserver"
		ipv46       = "10.10.10.33"
		port        = 80
		servicetype = "HTTP"
	}

	data "citrixadc_csvserver_appqoepolicy_binding" "tf_csvserver_appqoepolicy_binding" {
		name 		= citrixadc_csvserver_appqoepolicy_binding.tf_csvserver_appqoepolicy_binding.name
		policyname 	= citrixadc_csvserver_appqoepolicy_binding.tf_csvserver_appqoepolicy_binding.policyname
	}
`

func TestAccCsvserver_appqoepolicy_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCsvserver_appqoepolicy_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_csvserver_appqoepolicy_binding.tf_csvserver_appqoepolicy_binding", "name", "tf_csvserver"),
					resource.TestCheckResourceAttr("data.citrixadc_csvserver_appqoepolicy_binding.tf_csvserver_appqoepolicy_binding", "policyname", "tf_appqoepolicy"),
					resource.TestCheckResourceAttr("data.citrixadc_csvserver_appqoepolicy_binding.tf_csvserver_appqoepolicy_binding", "priority", "5"),
				),
			},
		},
	})
}

// testAccCsvserver_appqoepolicy_binding_upgrade_basic reuses the _basic config
// (binding + all prerequisite resources). It is valid under BOTH the SDK v2 2.2.0
// schema and the current Framework schema because the migration restored the SDK v2
// attribute names.
const testAccCsvserver_appqoepolicy_binding_upgrade_basic = `

	resource "citrixadc_appqoeaction" "tf_appqoeaction" {
		name        = "tf_appqoeaction"
		priority    = "LOW"
		respondwith = "NS"
		delay       = 40
	}
	resource "citrixadc_appqoepolicy" "tf_appqoepolicy" {
		name   = "tf_appqoepolicy"
		rule   = "true"
		action = citrixadc_appqoeaction.tf_appqoeaction.name
	}

	resource "citrixadc_csvserver_appqoepolicy_binding" "tf_csvserver_appqoepolicy_binding" {
		name 		= citrixadc_csvserver.tf_csvserver.name
		policyname 	= citrixadc_appqoepolicy.tf_appqoepolicy.name
		bindpoint 	= "REQUEST"
		priority 	= 5
	}

	resource "citrixadc_csvserver" "tf_csvserver" {
		name        = "tf_csvserver"
		ipv46       = "10.10.10.33"
		port        = 80
		servicetype = "HTTP"
	}
`

// TestAccCsvserver_appqoepolicy_binding_sdkv2StateUpgrade verifies that state written
// by the last SDK v2 release is correctly upgraded when the same config is subsequently
// managed by the current Framework provider. Step 1 creates the binding with
// citrix/citrixadc 2.2.0 (writes the legacy comma id "tf_csvserver,tf_appqoepolicy" —
// the SDK v2 d.SetId(fmt.Sprintf("%s,%s", name, policyname))). Step 2 refreshes/plans/
// applies the same config through the Framework provider, exercising ParseIdString on
// the legacy id; the Framework recomputes the id on Read (SetAttrFromGet), so the
// canonical new-format id becomes "name:tf_csvserver,policyname:tf_appqoepolicy".
func TestAccCsvserver_appqoepolicy_binding_sdkv2StateUpgrade(t *testing.T) {
	resourceAddr := "citrixadc_csvserver_appqoepolicy_binding.tf_csvserver_appqoepolicy_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckCsvserver_appqoepolicy_bindingDestroy,
		Steps: []resource.TestStep{
			// Step 1: create with the last SDK v2 release -> state carries the legacy id.
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.2.0",
					},
				},
				Config: testAccCsvserver_appqoepolicy_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCsvserver_appqoepolicy_bindingExist(resourceAddr, nil),
					resource.TestCheckResourceAttr(resourceAddr, "id", "tf_csvserver,tf_appqoepolicy"),
				),
			},
			// Step 2: refresh/plan/apply the SAME config through the current Framework
			// provider. The legacy-id state is read via ParseIdString and the id is
			// recomputed on Read into the new key:value format.
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccCsvserver_appqoepolicy_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCsvserver_appqoepolicy_bindingExist(resourceAddr, nil),
					resource.TestCheckResourceAttr(resourceAddr, "id", "name:tf_csvserver,policyname:tf_appqoepolicy"),
				),
			},
		},
	})
}

func TestAccCsvserver_appqoepolicy_binding_import(t *testing.T) {
	const resAddr = "citrixadc_csvserver_appqoepolicy_binding.tf_csvserver_appqoepolicy_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCsvserver_appqoepolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccCsvserver_appqoepolicy_binding_basic},
			{Config: testAccCsvserver_appqoepolicy_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
		},
	})
}
