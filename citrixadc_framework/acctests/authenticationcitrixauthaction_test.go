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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"testing"
)

const testAccAuthenticationcitrixauthaction_add = `

	resource "citrixadc_authenticationcitrixauthaction" "tf_citrixauthaction" {
		name               = "tf_citrixauthaction"
		authenticationtype = "CITRIXCONNECTOR"
		authentication     = "DISABLED"
	}
`
const testAccAuthenticationcitrixauthaction_update = `

	resource "citrixadc_authenticationcitrixauthaction" "tf_citrixauthaction" {
		name               = "tf_citrixauthaction"
		authenticationtype = "ATHENA"
		authentication     = "ENABLED"
	}
`

func TestAccAuthenticationcitrixauthaction_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationcitrixauthactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationcitrixauthaction_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationcitrixauthactionExist("citrixadc_authenticationcitrixauthaction.tf_citrixauthaction", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationcitrixauthaction.tf_citrixauthaction", "name", "tf_citrixauthaction"),
					resource.TestCheckResourceAttr("citrixadc_authenticationcitrixauthaction.tf_citrixauthaction", "authenticationtype", "CITRIXCONNECTOR"),
					resource.TestCheckResourceAttr("citrixadc_authenticationcitrixauthaction.tf_citrixauthaction", "authentication", "DISABLED"),
				),
			},
			{
				Config: testAccAuthenticationcitrixauthaction_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationcitrixauthactionExist("citrixadc_authenticationcitrixauthaction.tf_citrixauthaction", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationcitrixauthaction.tf_citrixauthaction", "name", "tf_citrixauthaction"),
					resource.TestCheckResourceAttr("citrixadc_authenticationcitrixauthaction.tf_citrixauthaction", "authenticationtype", "ATHENA"),
					resource.TestCheckResourceAttr("citrixadc_authenticationcitrixauthaction.tf_citrixauthaction", "authentication", "ENABLED"),
				),
			},
		},
	})
}

func testAccCheckAuthenticationcitrixauthactionExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No authenticationcitrixauthaction name is set")
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
		data, err := client.FindResource("authenticationcitrixauthaction", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("authenticationcitrixauthaction %s not found", n)
		}

		return nil
	}
}

func testAccCheckAuthenticationcitrixauthactionDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_authenticationcitrixauthaction" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("authenticationcitrixauthaction", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("authenticationcitrixauthaction %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
