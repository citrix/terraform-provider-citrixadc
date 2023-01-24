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
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"testing"
)

const testAccAuthenticationvserver_add = `

	resource "citrixadc_authenticationvserver" "tf_authenticationvserver" {
		name  		   = "tf_authenticationvserver"
		servicetype    = "SSL"
		comment 	   = "Hello"
		authentication = "ON"
		state          = "ENABLED"
	}
`
const testAccAuthenticationvserver_update = `

	resource "citrixadc_authenticationvserver" "tf_authenticationvserver" {
		name  		   = "tf_authenticationvserver"
		servicetype    = "SSL"
		comment 	   = "New"
		authentication = "ON"
		state          = "DISABLED"
	}
`
func TestAccAuthenticationvserver_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAuthenticationvserverDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccAuthenticationvserver_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationvserverExist("citrixadc_authenticationvserver.tf_authenticationvserver", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationvserver.tf_authenticationvserver", "name", "tf_authenticationvserver"),
					resource.TestCheckResourceAttr("citrixadc_authenticationvserver.tf_authenticationvserver", "comment", "Hello"),
					resource.TestCheckResourceAttr("citrixadc_authenticationvserver.tf_authenticationvserver", "state", "ENABLED"),
				),
			},
			resource.TestStep{
				Config: testAccAuthenticationvserver_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationvserverExist("citrixadc_authenticationvserver.tf_authenticationvserver", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationvserver.tf_authenticationvserver", "name", "tf_authenticationvserver"),
					resource.TestCheckResourceAttr("citrixadc_authenticationvserver.tf_authenticationvserver", "comment", "New"),
					resource.TestCheckResourceAttr("citrixadc_authenticationvserver.tf_authenticationvserver", "state", "DISABLED"),
				),
			},
		},
	})
}

func testAccCheckAuthenticationvserverExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No authenticationvserver name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(service.Authenticationvserver.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("authenticationvserver %s not found", n)
		}

		return nil
	}
}

func testAccCheckAuthenticationvserverDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_authenticationvserver" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Authenticationvserver.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("authenticationvserver %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
