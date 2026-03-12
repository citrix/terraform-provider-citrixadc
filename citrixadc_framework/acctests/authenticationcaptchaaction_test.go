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

const testAccAuthenticationcaptchaaction_add = `
	resource "citrixadc_authenticationcaptchaaction" "tf_captchaaction" {
		name                       = "tf_captchaaction"
		secretkey                  = "secret"
		sitekey                    = "key"
		serverurl                  = "http://www.example.com/"
		defaultauthenticationgroup = "old_group"
		scorethreshold			 = 3
	}
`
const testAccAuthenticationcaptchaaction_update = `
	resource "citrixadc_authenticationcaptchaaction" "tf_captchaaction" {
		name                       = "tf_captchaaction"
		secretkey                  = "new_secret"
		sitekey                    = "key"
		serverurl                  = "http://www.example.com/"
		defaultauthenticationgroup = "new_group"
		scorethreshold			 = 6
	}
`

func TestAccAuthenticationcaptchaaction_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationcaptchaactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationcaptchaaction_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationcaptchaactionExist("citrixadc_authenticationcaptchaaction.tf_captchaaction", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationcaptchaaction.tf_captchaaction", "name", "tf_captchaaction"),
					resource.TestCheckResourceAttr("citrixadc_authenticationcaptchaaction.tf_captchaaction", "secretkey", "secret"),
					resource.TestCheckResourceAttr("citrixadc_authenticationcaptchaaction.tf_captchaaction", "defaultauthenticationgroup", "old_group"),
					resource.TestCheckResourceAttr("citrixadc_authenticationcaptchaaction.tf_captchaaction", "scorethreshold", "3"),
				),
			},
			{
				Config: testAccAuthenticationcaptchaaction_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationcaptchaactionExist("citrixadc_authenticationcaptchaaction.tf_captchaaction", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationcaptchaaction.tf_captchaaction", "name", "tf_captchaaction"),
					resource.TestCheckResourceAttr("citrixadc_authenticationcaptchaaction.tf_captchaaction", "secretkey", "new_secret"),
					resource.TestCheckResourceAttr("citrixadc_authenticationcaptchaaction.tf_captchaaction", "defaultauthenticationgroup", "new_group"),
					resource.TestCheckResourceAttr("citrixadc_authenticationcaptchaaction.tf_captchaaction", "scorethreshold", "6"),
				),
			},
		},
	})
}

func testAccCheckAuthenticationcaptchaactionExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No authenticationcaptchaaction name is set")
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
		data, err := client.FindResource("authenticationcaptchaaction", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("authenticationcaptchaaction %s not found", n)
		}

		return nil
	}
}

func testAccCheckAuthenticationcaptchaactionDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_authenticationcaptchaaction" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("authenticationcaptchaaction", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("authenticationcaptchaaction %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccAuthenticationcaptchaactionDataSource_basic = `
	resource "citrixadc_authenticationcaptchaaction" "tf_captchaaction" {
		name                       = "tf_captchaaction_ds"
		secretkey                  = "secret"
		sitekey                    = "key"
		serverurl                  = "http://www.example.com/"
		defaultauthenticationgroup = "test_group"
		scorethreshold             = 7
	}

	data "citrixadc_authenticationcaptchaaction" "tf_captchaaction" {
		name       = citrixadc_authenticationcaptchaaction.tf_captchaaction.name
		depends_on = [citrixadc_authenticationcaptchaaction.tf_captchaaction]
	}
`

func TestAccAuthenticationcaptchaactionDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationcaptchaactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationcaptchaactionDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_authenticationcaptchaaction.tf_captchaaction", "name", "tf_captchaaction_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationcaptchaaction.tf_captchaaction", "serverurl", "http://www.example.com/"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationcaptchaaction.tf_captchaaction", "defaultauthenticationgroup", "test_group"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationcaptchaaction.tf_captchaaction", "scorethreshold", "7"),
				),
			},
		},
	})
}
