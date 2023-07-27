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

const testAccAuthenticationoauthidppolicy_add = `
	resource "citrixadc_authenticationoauthidpprofile" "tf_idpprofile" {
		name         = "tf_idpprofile"
		clientid     = "cliId"
		clientsecret = "secret"
		redirecturl  = "http://www.example.com/1/"
	}
	resource "citrixadc_authenticationoauthidppolicy" "tf_idppolicy" {
		name    = "tf_idppolicy"
		rule    = "true"
		action  = citrixadc_authenticationoauthidpprofile.tf_idpprofile.name
		comment = "add_policy"
	}
`
const testAccAuthenticationoauthidppolicy_update = `
	resource "citrixadc_authenticationoauthidpprofile" "tf_idpprofile" {
		name         = "tf_idpprofile"
		clientid     = "cliId"
		clientsecret = "secret"
		redirecturl  = "http://www.example.com/1/"
	}
	resource "citrixadc_authenticationoauthidppolicy" "tf_idppolicy" {
		name    = "tf_idppolicy"
		rule    = "false"
		action  = citrixadc_authenticationoauthidpprofile.tf_idpprofile.name
		comment = "update_policy"
	}
`

func TestAccAuthenticationoauthidppolicy_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAuthenticationoauthidppolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationoauthidppolicy_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationoauthidppolicyExist("citrixadc_authenticationoauthidppolicy.tf_idppolicy", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationoauthidppolicy.tf_idppolicy", "name", "tf_idppolicy"),
					resource.TestCheckResourceAttr("citrixadc_authenticationoauthidppolicy.tf_idppolicy", "rule", "true"),
					resource.TestCheckResourceAttr("citrixadc_authenticationoauthidppolicy.tf_idppolicy", "comment", "add_policy"),
				),
			},
			{
				Config: testAccAuthenticationoauthidppolicy_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationoauthidppolicyExist("citrixadc_authenticationoauthidppolicy.tf_idppolicy", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationoauthidppolicy.tf_idppolicy", "name", "tf_idppolicy"),
					resource.TestCheckResourceAttr("citrixadc_authenticationoauthidppolicy.tf_idppolicy", "rule", "false"),
					resource.TestCheckResourceAttr("citrixadc_authenticationoauthidppolicy.tf_idppolicy", "comment", "update_policy"),
				),
			},
		},
	})
}

func testAccCheckAuthenticationoauthidppolicyExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No authenticationoauthidppolicy name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource("authenticationoauthidppolicy", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("authenticationoauthidppolicy %s not found", n)
		}

		return nil
	}
}

func testAccCheckAuthenticationoauthidppolicyDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_authenticationoauthidppolicy" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource("authenticationoauthidppolicy", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("authenticationoauthidppolicy %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
