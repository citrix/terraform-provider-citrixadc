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

const testAccInterfacepair_basic = `
	resource "citrixadc_interfacepair" "tf_interfacepair" {
		interface_id = 1
		ifnum        = ["LA/2", "LA/3"]
	}
  
`
const testAccInterfacepair_update = `
	resource "citrixadc_interfacepair" "tf_interfacepair" {
		interface_id = 1
		ifnum        = ["LA/4", "LA/5"]
	}
  
`

const testAccInterfacepairDataSource_basic = `
	resource "citrixadc_interfacepair" "tf_interfacepair" {
		interface_id = 1
		ifnum        = ["LA/2", "LA/3"]
	}
	
	data "citrixadc_interfacepair" "tf_interfacepair_ds" {
		interface_id = citrixadc_interfacepair.tf_interfacepair.interface_id
	}
`

func TestAccInterfacepairDataSource_basic(t *testing.T) {
	t.Skip("TODO: Need to find a way to test this datasource - requires specific hardware setup!")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccInterfacepairDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_interfacepair.tf_interfacepair_ds", "interface_id", "1"),
					resource.TestCheckResourceAttrSet("data.citrixadc_interfacepair.tf_interfacepair_ds", "id"),
				),
			},
		},
	})
}

func TestAccInterfacepair_basic(t *testing.T) {
	t.Skip("TODO: Need to find a way to test this resource!")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckInterfacepairDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccInterfacepair_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckInterfacepairExist("citrixadc_interfacepair.tf_interfacepair", nil),
				),
			},
			{
				Config: testAccInterfacepair_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckInterfacepairExist("citrixadc_interfacepair.tf_interfacepair", nil),
				),
			},
		},
	})
}

func testAccCheckInterfacepairExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No interfacepair name is set")
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
		data, err := client.FindResource(service.Interfacepair.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("interfacepair %s not found", n)
		}

		return nil
	}
}

func testAccCheckInterfacepairDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_interfacepair" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Interfacepair.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("interfacepair %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
