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
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"testing"
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

func TestAccUservserver_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckUservserverDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
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
			resource.TestStep{
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

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource("uservserver", rs.Primary.ID)

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
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_uservserver" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource("uservserver", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("uservserver %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
