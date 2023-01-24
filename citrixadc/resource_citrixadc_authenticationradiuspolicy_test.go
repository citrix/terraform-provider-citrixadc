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

const testAccAuthenticationradiuspolicy_add = `
	resource "citrixadc_authenticationradiusaction" "tf_radiusaction" {
		name         = "tf_radiusaction"
		radkey       = "secret"
		serverip     = "1.2.3.4"
		serverport   = 8080
		authtimeout  = 2
		radnasip     = "DISABLED"
		passencoding = "chap"
	}
	resource "citrixadc_authenticationradiuspolicy" "tf_radiuspolicy" {
		name      = "tf_radiuspolicy"
		rule      = "NS_TRUE"
		reqaction = citrixadc_authenticationradiusaction.tf_radiusaction.name
	}
`
const testAccAuthenticationradiuspolicy_update = `
	resource "citrixadc_authenticationradiusaction" "tf_radiusaction" {
		name         = "tf_radiusaction"
		radkey       = "secret"
		serverip     = "1.2.3.4"
		serverport   = 8080
		authtimeout  = 2
		radnasip     = "DISABLED"
		passencoding = "chap"
	}
	resource "citrixadc_authenticationradiuspolicy" "tf_radiuspolicy" {
		name      = "tf_radiuspolicy"
		rule      = "NS_FALSE"
		reqaction = citrixadc_authenticationradiusaction.tf_radiusaction.name
	}
`

func TestAccAuthenticationradiuspolicy_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAuthenticationradiuspolicyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccAuthenticationradiuspolicy_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationradiuspolicyExist("citrixadc_authenticationradiuspolicy.tf_radiuspolicy", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationradiuspolicy.tf_radiuspolicy", "name", "tf_radiuspolicy"),
					resource.TestCheckResourceAttr("citrixadc_authenticationradiuspolicy.tf_radiuspolicy", "rule", "NS_TRUE"),
				),
			},
			resource.TestStep{
				Config: testAccAuthenticationradiuspolicy_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationradiuspolicyExist("citrixadc_authenticationradiuspolicy.tf_radiuspolicy", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationradiuspolicy.tf_radiuspolicy", "name", "tf_radiuspolicy"),
					resource.TestCheckResourceAttr("citrixadc_authenticationradiuspolicy.tf_radiuspolicy", "rule", "NS_FALSE"),
				),
			},
		},
	})
}

func testAccCheckAuthenticationradiuspolicyExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No authenticationradiuspolicy name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(service.Authenticationradiuspolicy.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("authenticationradiuspolicy %s not found", n)
		}

		return nil
	}
}

func testAccCheckAuthenticationradiuspolicyDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_authenticationradiuspolicy" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Authenticationradiuspolicy.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("authenticationradiuspolicy %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
