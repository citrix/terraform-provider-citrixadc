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
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"testing"
)

const testAccAuthenticationradiusaction_add = `
	resource "citrixadc_authenticationradiusaction" "tf_radiusaction" {
		name         = "tf_radiusaction"
		radkey       = "secret"
		serverip     = "1.2.3.4"
		serverport   = 8080
		authtimeout  = 2
		radnasip     = "DISABLED"
		passencoding = "chap"
	}
`
const testAccAuthenticationradiusaction_update = `
	resource "citrixadc_authenticationradiusaction" "tf_radiusaction" {
		name         = "tf_radiusaction"
		radkey       = "secret"
		serverip     = "1.2.3.4"
		serverport   = 8080
		authtimeout  = 2
		radnasip     = "ENABLED"
		passencoding = "pap"
	}
`

func TestAccAuthenticationradiusaction_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAuthenticationradiusactionDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccAuthenticationradiusaction_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationradiusactionExist("citrixadc_authenticationradiusaction.tf_radiusaction", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationradiusaction.tf_radiusaction", "name", "tf_radiusaction"),
					resource.TestCheckResourceAttr("citrixadc_authenticationradiusaction.tf_radiusaction", "radnasip", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_authenticationradiusaction.tf_radiusaction", "passencoding", "chap"),
				),
			},
			resource.TestStep{
				Config: testAccAuthenticationradiusaction_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationradiusactionExist("citrixadc_authenticationradiusaction.tf_radiusaction", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationradiusaction.tf_radiusaction", "name", "tf_radiusaction"),
					resource.TestCheckResourceAttr("citrixadc_authenticationradiusaction.tf_radiusaction", "radnasip", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_authenticationradiusaction.tf_radiusaction", "passencoding", "pap"),
				),
			},
		},
	})
}

func testAccCheckAuthenticationradiusactionExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No authenticationradiusaction name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(service.Authenticationradiusaction.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("authenticationradiusaction %s not found", n)
		}

		return nil
	}
}

func testAccCheckAuthenticationradiusactionDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_authenticationradiusaction" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Authenticationradiusaction.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("authenticationradiusaction %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
