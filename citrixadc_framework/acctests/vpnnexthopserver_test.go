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

const testAccVpnnexthopserver_add = `

	resource "citrixadc_vpnnexthopserver" "tf_vpnnexthopserver" {
		name        = "tf_vpnnexthopserver"
		nexthopip   = "2.6.1.5"
		nexthopport = "200"
	}
`
const testAccVpnnexthopserver_update = `

	resource "citrixadc_vpnnexthopserver" "tf_vpnnexthopserver" {
		name        = "tf_vpnnexthopserver"
		nexthopip   = "2.6.1.5"
		nexthopport = "300"
	}
`

func TestAccVpnnexthopserver_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVpnnexthopserverDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnnexthopserver_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnnexthopserverExist("citrixadc_vpnnexthopserver.tf_vpnnexthopserver", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnnexthopserver.tf_vpnnexthopserver", "name", "tf_vpnnexthopserver"),
					resource.TestCheckResourceAttr("citrixadc_vpnnexthopserver.tf_vpnnexthopserver", "nexthopport", "200"),
				),
			},
			{
				Config: testAccVpnnexthopserver_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnnexthopserverExist("citrixadc_vpnnexthopserver.tf_vpnnexthopserver", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnnexthopserver.tf_vpnnexthopserver", "name", "tf_vpnnexthopserver"),
					resource.TestCheckResourceAttr("citrixadc_vpnnexthopserver.tf_vpnnexthopserver", "nexthopport", "300"),
				),
			},
		},
	})
}

func testAccCheckVpnnexthopserverExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No vpnnexthopserver name is set")
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
		data, err := client.FindResource(service.Vpnnexthopserver.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("vpnnexthopserver %s not found", n)
		}

		return nil
	}
}

func testAccCheckVpnnexthopserverDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_vpnnexthopserver" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Vpnnexthopserver.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("vpnnexthopserver %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccVpnnexthopserverDataSource_basic = `

	resource "citrixadc_vpnnexthopserver" "tf_vpnnexthopserver" {
		name        = "tf_vpnnexthopserver"
		nexthopip   = "2.6.1.5"
		nexthopport = "200"
	}

data "citrixadc_vpnnexthopserver" "tf_vpnnexthopserver" {
	name = citrixadc_vpnnexthopserver.tf_vpnnexthopserver.name
}
`

func TestAccVpnnexthopserverDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVpnnexthopserverDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnnexthopserverDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_vpnnexthopserver.tf_vpnnexthopserver", "name", "tf_vpnnexthopserver"),
					resource.TestCheckResourceAttr("data.citrixadc_vpnnexthopserver.tf_vpnnexthopserver", "nexthopip", "2.6.1.5"),
					resource.TestCheckResourceAttr("data.citrixadc_vpnnexthopserver.tf_vpnnexthopserver", "nexthopport", "200"),
				),
			},
		},
	})
}
