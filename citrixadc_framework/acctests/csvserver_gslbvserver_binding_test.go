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

const testAccCsvserver_gslbvserver_binding_basic = `
	resource "citrixadc_csvserver_gslbvserver_binding" "tf_csvserver_gslbvserver_binding" {
        name = citrixadc_csvserver.tf_csvserver.name
        vserver = citrixadc_gslbvserver.tf_gslbvserver.name
	}

	resource "citrixadc_csvserver" "tf_csvserver" {
		name = "tf_csvserver"
		servicetype = "HTTP"
		targettype = "GSLB"
	}

	resource "citrixadc_gslbvserver" "tf_gslbvserver" {
		name = "tf_gslbvserver"
		servicetype = "HTTP"
	}
`

const testAccCsvserver_gslbvserver_binding_basic_step2 = `
	resource "citrixadc_csvserver" "tf_csvserver" {
		name = "tf_csvserver"
		servicetype = "HTTP"
		targettype = "GSLB"
	}

	resource "citrixadc_gslbvserver" "tf_gslbvserver" {
		name = "tf_gslbvserver"
		servicetype = "HTTP"
	}
`

func TestAccCsvserver_gslbvserver_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCsvserver_gslbvserver_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCsvserver_gslbvserver_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCsvserver_gslbvserver_bindingExist("citrixadc_csvserver_gslbvserver_binding.tf_csvserver_gslbvserver_binding", nil),
				),
			},
			{
				Config: testAccCsvserver_gslbvserver_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCsvserver_gslbvserver_bindingNotExist("citrixadc_csvserver_gslbvserver_binding.tf_csvserver_gslbvserver_binding", "tf_csvserver,tf_gslbvserver"),
				),
			},
		},
	})
}

func testAccCheckCsvserver_gslbvserver_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No csvserver_gslbvserver_binding id is set")
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

		idMap, _, err := utils.ParseIdString(bindingId, []string{"name", "vserver"}, nil)
		if err != nil {
			return err
		}
		name := idMap["name"]
		vserver := idMap["vserver"]

		findParams := service.FindParams{
			ResourceType:             "csvserver_gslbvserver_binding",
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
			if v["vserver"].(string) == vserver {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("csvserver_gslbvserver_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckCsvserver_gslbvserver_bindingNotExist(n string, id string) resource.TestCheckFunc {
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
		vserver := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "csvserver_gslbvserver_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the right gslbvserver name
		found := false
		for _, v := range dataArr {
			if v["vserver"].(string) == vserver {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("csvserver_gslbvserver_binding %s was found, but is should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckCsvserver_gslbvserver_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_csvserver_gslbvserver_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Csvserver_gslbvserver_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("csvserver_gslbvserver_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccCsvserver_gslbvserver_bindingDataSource_basic = `
	resource "citrixadc_csvserver" "tf_csvserver_ds" {
		name = "tf_csvserver_ds"
		servicetype = "HTTP"
		targettype = "GSLB"
	}

	resource "citrixadc_gslbvserver" "tf_gslbvserver_ds" {
		name = "tf_gslbvserver_ds"
		servicetype = "HTTP"
	}

	resource "citrixadc_csvserver_gslbvserver_binding" "tf_csvserver_gslbvserver_binding_ds" {
		name = citrixadc_csvserver.tf_csvserver_ds.name
		vserver = citrixadc_gslbvserver.tf_gslbvserver_ds.name
	}

	data "citrixadc_csvserver_gslbvserver_binding" "tf_csvserver_gslbvserver_binding_ds_read" {
		name = citrixadc_csvserver_gslbvserver_binding.tf_csvserver_gslbvserver_binding_ds.name
		vserver = citrixadc_csvserver_gslbvserver_binding.tf_csvserver_gslbvserver_binding_ds.vserver
		depends_on = [citrixadc_csvserver_gslbvserver_binding.tf_csvserver_gslbvserver_binding_ds]
	}
`

func TestAcccsvserver_gslbvserver_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccCsvserver_gslbvserver_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_csvserver_gslbvserver_binding.tf_csvserver_gslbvserver_binding_ds_read", "name", "tf_csvserver_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_csvserver_gslbvserver_binding.tf_csvserver_gslbvserver_binding_ds_read", "vserver", "tf_gslbvserver_ds"),
					resource.TestCheckResourceAttrSet("data.citrixadc_csvserver_gslbvserver_binding.tf_csvserver_gslbvserver_binding_ds_read", "id"),
				),
			},
		},
	})
}

// testAccCsvserver_gslbvserver_binding_upgrade_basic is the config used by the
// sdkv2 -> framework state-upgrade test. It reuses the same values and resource
// labels as testAccCsvserver_gslbvserver_binding_basic so it is valid under BOTH
// the SDK v2 2.2.0 schema and the current framework schema.
const testAccCsvserver_gslbvserver_binding_upgrade_basic = `
	resource "citrixadc_csvserver_gslbvserver_binding" "tf_csvserver_gslbvserver_binding" {
        name = citrixadc_csvserver.tf_csvserver.name
        vserver = citrixadc_gslbvserver.tf_gslbvserver.name
	}

	resource "citrixadc_csvserver" "tf_csvserver" {
		name = "tf_csvserver"
		servicetype = "HTTP"
		targettype = "GSLB"
	}

	resource "citrixadc_gslbvserver" "tf_gslbvserver" {
		name = "tf_gslbvserver"
		servicetype = "HTTP"
	}
`

// TestAccCsvserver_gslbvserver_binding_sdkv2StateUpgrade verifies that a binding
// created with the last SDK v2 release (2.2.0, legacy comma-separated ID) is
// correctly refreshed/planned/applied by the current framework provider.
func TestAccCsvserver_gslbvserver_binding_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckCsvserver_gslbvserver_bindingDestroy,
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
				Config: testAccCsvserver_gslbvserver_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCsvserver_gslbvserver_bindingExist("citrixadc_csvserver_gslbvserver_binding.tf_csvserver_gslbvserver_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_csvserver_gslbvserver_binding.tf_csvserver_gslbvserver_binding", "id", "tf_csvserver,tf_gslbvserver"),
				),
			},
			// Step 2: same config, current (framework) provider. Terraform
			// refreshes the legacy-id state through the framework Read
			// (exercising ParseIdString on the legacy id) then plans/applies.
			// The framework recomputes the id on read to the new key:value form.
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccCsvserver_gslbvserver_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCsvserver_gslbvserver_bindingExist("citrixadc_csvserver_gslbvserver_binding.tf_csvserver_gslbvserver_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_csvserver_gslbvserver_binding.tf_csvserver_gslbvserver_binding", "id", "name:tf_csvserver,vserver:tf_gslbvserver"),
				),
			},
		},
	})
}

func TestAccCsvserver_gslbvserver_binding_import(t *testing.T) {
	const resAddr = "citrixadc_csvserver_gslbvserver_binding.tf_csvserver_gslbvserver_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCsvserver_gslbvserver_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccCsvserver_gslbvserver_binding_basic},
			{Config: testAccCsvserver_gslbvserver_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
		},
	})
}
