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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testAccVpnglobal_staserver_binding_basic = `
	resource "citrixadc_vpnglobal_staserver_binding" "tf_bind" {
		staserver      = "http://www.example.com/"
		staaddresstype = "IPV4"
	}
`

const testAccVpnglobal_staserver_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion
`

const testAccVpnglobal_staserver_bindingDataSource_basic = `
	resource "citrixadc_vpnglobal_staserver_binding" "tf_bind" {
		staserver      = "http://www.example.com/"
		staaddresstype = "IPV4"
	}

	data "citrixadc_vpnglobal_staserver_binding" "tf_bind" {
		staserver  = citrixadc_vpnglobal_staserver_binding.tf_bind.staserver
		depends_on = [citrixadc_vpnglobal_staserver_binding.tf_bind]
	}
`

func TestAccVpnglobal_staserver_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVpnglobal_staserver_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnglobal_staserver_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnglobal_staserver_bindingExist("citrixadc_vpnglobal_staserver_binding.tf_bind", nil),
				),
			},
			{
				Config: testAccVpnglobal_staserver_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnglobal_staserver_bindingNotExist("citrixadc_vpnglobal_staserver_binding.tf_bind", "http://www.example.com/"),
				),
			},
		},
	})
}

func testAccCheckVpnglobal_staserver_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No vpnglobal_staserver_binding id is set")
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

		staserver := rs.Primary.ID

		findParams := service.FindParams{
			ResourceType:             "vpnglobal_staserver_binding",
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the matching secondIdComponent
		found := false
		for _, v := range dataArr {
			if v["staserver"].(string) == staserver {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("vpnglobal_staserver_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckVpnglobal_staserver_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		staserver := id
		findParams := service.FindParams{
			ResourceType:             "vpnglobal_staserver_binding",
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
			if v["staserver"].(string) == staserver {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("vpnglobal_staserver_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckVpnglobal_staserver_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_vpnglobal_staserver_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Vpnglobal_staserver_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("vpnglobal_staserver_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccVpnglobal_staserver_binding_upgrade_basic = `
	resource "citrixadc_vpnglobal_staserver_binding" "tf_bind" {
		staserver      = "http://www.example.com/"
		staaddresstype = "IPV4"
	}
`

func TestAccVpnglobal_staserver_binding_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckVpnglobal_staserver_bindingDestroy,
		Steps: []resource.TestStep{
			{
				// Step 1: create the binding with the last SDK v2 release (2.2.0),
				// which writes state using the legacy id (plain staserver value).
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.2.0",
					},
				},
				Config: testAccVpnglobal_staserver_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnglobal_staserver_bindingExist("citrixadc_vpnglobal_staserver_binding.tf_bind", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnglobal_staserver_binding.tf_bind", "id", "http://www.example.com/"),
				),
			},
			{
				// Step 2: refresh/plan the legacy-id state through the current
				// framework provider. Read re-derives the canonical id in
				// SetAttrFromGet. This is a single-key binding, so the canonical
				// id is the plain staserver value (same as the legacy id).
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccVpnglobal_staserver_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnglobal_staserver_bindingExist("citrixadc_vpnglobal_staserver_binding.tf_bind", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnglobal_staserver_binding.tf_bind", "id", "http://www.example.com/"),
				),
			},
		},
	})
}

func TestAccVpnglobal_staserver_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnglobal_staserver_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_vpnglobal_staserver_binding.tf_bind", "staserver", "http://www.example.com/"),
					resource.TestCheckResourceAttr("data.citrixadc_vpnglobal_staserver_binding.tf_bind", "staaddresstype", "IPV4"),
				),
			},
		},
	})
}

func TestAccVpnglobal_staserver_binding_import(t *testing.T) {
	const resAddr = "citrixadc_vpnglobal_staserver_binding.tf_bind"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVpnglobal_staserver_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccVpnglobal_staserver_binding_basic},
			{Config: testAccVpnglobal_staserver_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
		},
	})
}
