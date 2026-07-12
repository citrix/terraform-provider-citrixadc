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

const testAccCsvserver_analyticsprofile_binding_basic = `
	resource "citrixadc_csvserver_analyticsprofile_binding" "tf_csvserver_analyticsprofile_binding" {
        name = citrixadc_csvserver.tf_csvserver.name
		analyticsprofile = "ns_analytics_global_profile"
	}

	resource "citrixadc_csvserver" "tf_csvserver" {
		name = "tf_csvserver"
		ipv46 = "1.1.1.2"
		port = 80
		servicetype = "HTTP"
	}
`

const testAccCsvserver_analyticsprofile_binding_basic_step2 = `
	resource "citrixadc_csvserver" "tf_csvserver" {
		name = "tf_csvserver"
		ipv46 = "1.1.1.2"
		port = 80
		servicetype = "HTTP"
	}
`

const testAccCsvserver_analyticsprofile_bindingDataSource_basic = `
	resource "citrixadc_csvserver_analyticsprofile_binding" "tf_csvserver_analyticsprofile_binding" {
        name = citrixadc_csvserver.tf_csvserver.name
		analyticsprofile = "ns_analytics_global_profile"
	}

	resource "citrixadc_csvserver" "tf_csvserver" {
		name = "tf_csvserver"
		ipv46 = "1.1.1.2"
		port = 80
		servicetype = "HTTP"
	}

	data "citrixadc_csvserver_analyticsprofile_binding" "tf_csvserver_analyticsprofile_binding" {
		name = citrixadc_csvserver_analyticsprofile_binding.tf_csvserver_analyticsprofile_binding.name
		analyticsprofile = citrixadc_csvserver_analyticsprofile_binding.tf_csvserver_analyticsprofile_binding.analyticsprofile
	}
`

func TestAccCsvserver_analyticsprofile_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCsvserver_analyticsprofile_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCsvserver_analyticsprofile_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCsvserver_analyticsprofile_bindingExist("citrixadc_csvserver_analyticsprofile_binding.tf_csvserver_analyticsprofile_binding", nil),
				),
			},
			{
				Config: testAccCsvserver_analyticsprofile_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCsvserver_analyticsprofile_bindingNotExist("citrixadc_csvserver_analyticsprofile_binding.tf_csvserver_analyticsprofile_binding", "tf_csvserver,ns_analytics_global_profile"),
				),
			},
		},
	})
}

func TestAccCsvserver_analyticsprofile_binding_import(t *testing.T) {
	const resAddr = "citrixadc_csvserver_analyticsprofile_binding.tf_csvserver_analyticsprofile_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCsvserver_analyticsprofile_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccCsvserver_analyticsprofile_binding_basic},
			{Config: testAccCsvserver_analyticsprofile_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
		},
	})
}

func testAccCheckCsvserver_analyticsprofile_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No csvserver_analyticsprofile_binding id is set")
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

		idMap, _, err := utils.ParseIdString(bindingId, []string{"name", "analyticsprofile"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID: %v", err)
		}
		name := idMap["name"]
		analyticsprofile := idMap["analyticsprofile"]

		findParams := service.FindParams{
			ResourceType:             "csvserver_analyticsprofile_binding",
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
			if v["analyticsprofile"].(string) == analyticsprofile {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("csvserver_analyticsprofile_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckCsvserver_analyticsprofile_bindingNotExist(n string, id string) resource.TestCheckFunc {
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

		name := idSlice[0]
		analyticsprofile := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "csvserver_analyticsprofile_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the right profile name
		found := false
		for _, v := range dataArr {
			if v["analyticsprofile"].(string) == analyticsprofile {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("csvserver_analyticsprofile_binding %s was found, but is should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckCsvserver_analyticsprofile_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_csvserver_analyticsprofile_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("csvserver_analyticsprofile_binding", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("csvserver_analyticsprofile_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccCsvserver_analyticsprofile_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCsvserver_analyticsprofile_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_csvserver_analyticsprofile_binding.tf_csvserver_analyticsprofile_binding", "name", "tf_csvserver"),
					resource.TestCheckResourceAttr("data.citrixadc_csvserver_analyticsprofile_binding.tf_csvserver_analyticsprofile_binding", "analyticsprofile", "ns_analytics_global_profile"),
				),
			},
		},
	})
}

// testAccCsvserver_analyticsprofile_binding_upgrade_basic reuses the _basic config
// (binding + its csvserver prerequisite). It is valid under BOTH the SDK v2 2.2.0
// schema and the current Framework schema because the migration restored the SDK v2
// attribute names.
const testAccCsvserver_analyticsprofile_binding_upgrade_basic = `
	resource "citrixadc_csvserver_analyticsprofile_binding" "tf_csvserver_analyticsprofile_binding" {
        name = citrixadc_csvserver.tf_csvserver.name
		analyticsprofile = "ns_analytics_global_profile"
	}

	resource "citrixadc_csvserver" "tf_csvserver" {
		name = "tf_csvserver"
		ipv46 = "1.1.1.2"
		port = 80
		servicetype = "HTTP"
	}
`

// TestAccCsvserver_analyticsprofile_binding_sdkv2StateUpgrade verifies that state
// written by the last SDK v2 release is correctly upgraded when the same config is
// subsequently managed by the current Framework provider. Step 1 creates the binding
// with citrix/citrixadc 2.2.0 (writes the legacy id "tf_csvserver,ns_analytics_global_profile" —
// the SDK v2 d.SetId(name,analyticsprofile)). Step 2 refreshes/plans/applies the same
// config through the Framework provider, exercising ParseIdString on the legacy id; the
// Framework recomputes the id on Read (SetAttrFromGet) into the canonical new format
// "analyticsprofile:ns_analytics_global_profile,name:tf_csvserver".
func TestAccCsvserver_analyticsprofile_binding_sdkv2StateUpgrade(t *testing.T) {
	resourceAddr := "citrixadc_csvserver_analyticsprofile_binding.tf_csvserver_analyticsprofile_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckCsvserver_analyticsprofile_bindingDestroy,
		Steps: []resource.TestStep{
			// Step 1: create with the last SDK v2 release -> state carries the legacy id.
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.2.0",
					},
				},
				Config: testAccCsvserver_analyticsprofile_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCsvserver_analyticsprofile_bindingExist(resourceAddr, nil),
					resource.TestCheckResourceAttr(resourceAddr, "id", "tf_csvserver,ns_analytics_global_profile"),
				),
			},
			// Step 2: refresh/plan/apply the SAME config through the current Framework
			// provider. The legacy-id state is read via ParseIdString and the id is
			// recomputed on Read into the new key:value format.
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccCsvserver_analyticsprofile_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCsvserver_analyticsprofile_bindingExist(resourceAddr, nil),
					resource.TestCheckResourceAttr(resourceAddr, "id", "analyticsprofile:ns_analytics_global_profile,name:tf_csvserver"),
				),
			},
		},
	})
}
