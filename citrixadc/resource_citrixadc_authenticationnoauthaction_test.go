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

const testAccAuthenticationnoauthaction_add = `
	resource "citrixadc_authenticationnoauthaction" "tf_noauthaction" {
		name                       = "tf_noauthaction"
		defaultauthenticationgroup = "old_group"
	}
`
const testAccAuthenticationnoauthaction_update = `
	resource "citrixadc_authenticationnoauthaction" "tf_noauthaction" {
		name                       = "tf_noauthaction"
		defaultauthenticationgroup = "new_group"
	}
`

func TestAccAuthenticationnoauthaction_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckAuthenticationnoauthactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationnoauthaction_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationnoauthactionExist("citrixadc_authenticationnoauthaction.tf_noauthaction", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationnoauthaction.tf_noauthaction", "name", "tf_noauthaction"),
					resource.TestCheckResourceAttr("citrixadc_authenticationnoauthaction.tf_noauthaction", "defaultauthenticationgroup", "old_group"),
				),
			},
			{
				Config: testAccAuthenticationnoauthaction_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationnoauthactionExist("citrixadc_authenticationnoauthaction.tf_noauthaction", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationnoauthaction.tf_noauthaction", "name", "tf_noauthaction"),
					resource.TestCheckResourceAttr("citrixadc_authenticationnoauthaction.tf_noauthaction", "defaultauthenticationgroup", "new_group"),
				),
			},
		},
	})
}

func testAccCheckAuthenticationnoauthactionExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No authenticationnoauthaction name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		// Use the shared utility function to get a configured client
		client, err := testAccGetClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource("authenticationnoauthaction", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("authenticationnoauthaction %s not found", n)
		}

		return nil
	}
}

func testAccCheckAuthenticationnoauthactionDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_authenticationnoauthaction" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("authenticationnoauthaction", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("authenticationnoauthaction %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
