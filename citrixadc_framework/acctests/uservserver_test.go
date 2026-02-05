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

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testAccUservserver_basic = `


	resource "citrixadc_uservserver" "tf_uservserver" {
		name         = "my_user_vserver"
		userprotocol = "MQTT"
		ipaddress    = "10.222.74.180"
		port         = 3200
		defaultlb    = "mysv"
	}
`
const testAccUservserver_update = `


	resource "citrixadc_uservserver" "tf_uservserver" {
		name         = "my_user_vserver"
		userprotocol = "my_user_protocol"
		ipaddress    = "10.222.74.200"
		port         = 3500
		defaultlb    = "mysv"
	}
`

const testAccUservserverDataSource_basic = `

	resource "citrixadc_userprotocol" "tf_userprotocol" {
		name      = "MQTT"
		transport = "TCP"
		extension = "mqtt_code"
		comment   = "my_comment"
	}

	resource "citrixadc_lbvserver" "tf_defaultlb" {
		name        = "tf_defaultlb"
		servicetype = "USER_TCP"
	}

	resource "citrixadc_uservserver" "tf_uservserver" {
		name         = "my_user_vserver"
		userprotocol = "MQTT"
		ipaddress    = "10.222.74.180"
		port         = 80
		defaultlb    = citrixadc_lbvserver.tf_defaultlb.name
		depends_on   = [citrixadc_userprotocol.tf_userprotocol, citrixadc_lbvserver.tf_defaultlb]
	}

	data "citrixadc_uservserver" "tf_uservserver" {
		name = citrixadc_uservserver.tf_uservserver.name
		depends_on = [citrixadc_uservserver.tf_uservserver]
	}
`

func TestAccUservserver_basic(t *testing.T) {
	t.Skip("TODO: Need to find a way to test this resource!")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckUservserverDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccUservserver_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUservserverExist("citrixadc_uservserver.tf_uservserver", nil),
					resource.TestCheckResourceAttr("citrixadc_uservserver.tf_uservserver", "name", "my_user_vserver"),
					resource.TestCheckResourceAttr("citrixadc_uservserver.tf_uservserver", "userprotocol", "MQTT"),
					resource.TestCheckResourceAttr("citrixadc_uservserver.tf_uservserver", "ipaddress", "10.222.74.180"),
					resource.TestCheckResourceAttr("citrixadc_uservserver.tf_uservserver", "port", "3200"),
					resource.TestCheckResourceAttr("citrixadc_uservserver.tf_uservserver", "defaultlb", "mysv"),
				),
			},
			{
				Config: testAccUservserver_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUservserverExist("citrixadc_uservserver.tf_uservserver", nil),
					resource.TestCheckResourceAttr("citrixadc_uservserver.tf_uservserver", "name", "my_user_vserver"),
					resource.TestCheckResourceAttr("citrixadc_uservserver.tf_uservserver", "userprotocol", "my_user_protocol"),
					resource.TestCheckResourceAttr("citrixadc_uservserver.tf_uservserver", "ipaddress", "10.222.74.200"),
					resource.TestCheckResourceAttr("citrixadc_uservserver.tf_uservserver", "port", "3500"),
					resource.TestCheckResourceAttr("citrixadc_uservserver.tf_uservserver", "defaultlb", "mysv"),
				),
			},
		},
	})
}

func testAccCheckUservserverExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No uservserver name is set")
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
		data, err := client.FindResource("uservserver", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("uservserver %s not found", n)
		}

		return nil
	}
}

func testAccCheckUservserverDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_uservserver" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("uservserver", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("uservserver %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccUservserverDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUservserverDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_uservserver.tf_uservserver", "name", "my_user_vserver"),
					resource.TestCheckResourceAttr("data.citrixadc_uservserver.tf_uservserver", "userprotocol", "MQTT"),
					resource.TestCheckResourceAttr("data.citrixadc_uservserver.tf_uservserver", "ipaddress", "10.222.74.180"),
					resource.TestCheckResourceAttr("data.citrixadc_uservserver.tf_uservserver", "port", "80"),
				),
			},
		},
	})
}
