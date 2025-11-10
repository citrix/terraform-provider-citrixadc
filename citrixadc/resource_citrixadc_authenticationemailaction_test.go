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

const testAccAuthenticationemailaction_add = `
	resource "citrixadc_authenticationemailaction" "tf_emailaction" {
		name      = "tf_emailaction"
		username  = "username@abc.com"
		password  = "secret"
		serverurl = "www.example.com"
		timeout   = 100
		type      = "SMTP"
	}
`
const testAccAuthenticationemailaction_update = `
	resource "citrixadc_authenticationemailaction" "tf_emailaction" {
		name      = "tf_emailaction"
		username  = "username1@abc.com"
		password  = "secret"
		serverurl = "www.example1.com"
		timeout   = 100
		type      = "SMTP"
	}
`

func TestAccAuthenticationemailaction_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckAuthenticationemailactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationemailaction_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationemailactionExist("citrixadc_authenticationemailaction.tf_emailaction", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationemailaction.tf_emailaction", "name", "tf_emailaction"),
					resource.TestCheckResourceAttr("citrixadc_authenticationemailaction.tf_emailaction", "username", "username@abc.com"),
					resource.TestCheckResourceAttr("citrixadc_authenticationemailaction.tf_emailaction", "serverurl", "www.example.com"),
				),
			},
			{
				Config: testAccAuthenticationemailaction_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationemailactionExist("citrixadc_authenticationemailaction.tf_emailaction", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationemailaction.tf_emailaction", "name", "tf_emailaction"),
					resource.TestCheckResourceAttr("citrixadc_authenticationemailaction.tf_emailaction", "username", "username1@abc.com"),
					resource.TestCheckResourceAttr("citrixadc_authenticationemailaction.tf_emailaction", "serverurl", "www.example1.com"),
				),
			},
		},
	})
}

func testAccCheckAuthenticationemailactionExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No authenticationemailaction name is set")
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
		data, err := client.FindResource("authenticationemailaction", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("authenticationemailaction %s not found", n)
		}

		return nil
	}
}

func testAccCheckAuthenticationemailactionDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_authenticationemailaction" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("authenticationemailaction", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("authenticationemailaction %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
