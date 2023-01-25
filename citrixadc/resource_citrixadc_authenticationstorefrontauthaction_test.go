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

const testAccAuthenticationstorefrontauthaction_add = `
	resource "citrixadc_authenticationstorefrontauthaction" "tf_storefront" {
		name                       = "tf_storefront"
		serverurl                  = "http://www.example.com/"
		domain                     = "domainname"
		defaultauthenticationgroup = "group_name"
	}
`
const testAccAuthenticationstorefrontauthaction_update = `
	resource "citrixadc_authenticationstorefrontauthaction" "tf_storefront" {
		name                       = "tf_storefront"
		serverurl                  = "http://www.example.com/"
		domain                     = "new_domainname"
		defaultauthenticationgroup = "new_groupname"
	}
`

func TestAccAuthenticationstorefrontauthaction_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAuthenticationstorefrontauthactionDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccAuthenticationstorefrontauthaction_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationstorefrontauthactionExist("citrixadc_authenticationstorefrontauthaction.tf_storefront", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationstorefrontauthaction.tf_storefront", "name", "tf_storefront"),
					resource.TestCheckResourceAttr("citrixadc_authenticationstorefrontauthaction.tf_storefront", "domain", "domainname"),
				),
			},
			resource.TestStep{
				Config: testAccAuthenticationstorefrontauthaction_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationstorefrontauthactionExist("citrixadc_authenticationstorefrontauthaction.tf_storefront", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationstorefrontauthaction.tf_storefront", "name", "tf_storefront"),
					resource.TestCheckResourceAttr("citrixadc_authenticationstorefrontauthaction.tf_storefront", "domain", "new_domainname"),
				),
			},
		},
	})
}

func testAccCheckAuthenticationstorefrontauthactionExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No authenticationstorefrontauthaction name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource("authenticationstorefrontauthaction", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("authenticationstorefrontauthaction %s not found", n)
		}

		return nil
	}
}

func testAccCheckAuthenticationstorefrontauthactionDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_authenticationstorefrontauthaction" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource("authenticationstorefrontauthaction", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("authenticationstorefrontauthaction %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
