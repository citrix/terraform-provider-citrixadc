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

const testAccAuthenticationcertaction_add = `
	resource "citrixadc_authenticationcertaction" "tf_certaction" {
		name                       = "tf_certaction"
		twofactor                  = "ON"
		defaultauthenticationgroup = "old_group"
		usernamefield              = "Subject:CN"
		groupnamefield             = "subject:grp"
	}
`
const testAccAuthenticationcertaction_update = `
	resource "citrixadc_authenticationcertaction" "tf_certaction" {
		name                       = "tf_certaction"
		twofactor                  = "OFF"
		defaultauthenticationgroup = "new_group"
		usernamefield              = "Subject:CN"
		groupnamefield             = "subject:grp"
	}
`

func TestAccAuthenticationcertaction_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAuthenticationcertactionDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccAuthenticationcertaction_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationcertactionExist("citrixadc_authenticationcertaction.tf_certaction", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationcertaction.tf_certaction", "name", "tf_certaction"),
					resource.TestCheckResourceAttr("citrixadc_authenticationcertaction.tf_certaction", "twofactor", "ON"),
					resource.TestCheckResourceAttr("citrixadc_authenticationcertaction.tf_certaction", "defaultauthenticationgroup", "old_group"),
				),
			},
			resource.TestStep{
				Config: testAccAuthenticationcertaction_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationcertactionExist("citrixadc_authenticationcertaction.tf_certaction", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationcertaction.tf_certaction", "name", "tf_certaction"),
					resource.TestCheckResourceAttr("citrixadc_authenticationcertaction.tf_certaction", "twofactor", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_authenticationcertaction.tf_certaction", "defaultauthenticationgroup", "new_group"),
				),
			},
		},
	})
}

func testAccCheckAuthenticationcertactionExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No authenticationcertaction name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(service.Authenticationcertaction.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("authenticationcertaction %s not found", n)
		}

		return nil
	}
}

func testAccCheckAuthenticationcertactionDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_authenticationcertaction" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Authenticationcertaction.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("authenticationcertaction %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
