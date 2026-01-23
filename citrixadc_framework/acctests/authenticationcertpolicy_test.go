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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"testing"
)

const testAccAuthenticationcertpolicy_add = `
	resource "citrixadc_authenticationcertaction" "tf_certaction" {
		name                       = "tf_certaction"
		twofactor                  = "ON"
		defaultauthenticationgroup = "new_group"
		usernamefield              = "Subject:CN"
		groupnamefield             = "subject:grp"
	}
	resource "citrixadc_authenticationcertpolicy" "tf_certpolicy" {
		name      = "tf_certpolicy"
		rule      = "ns_true"
		reqaction = citrixadc_authenticationcertaction.tf_certaction.name
	}
`
const testAccAuthenticationcertpolicy_update = `
	resource "citrixadc_authenticationcertaction" "tf_certaction" {
		name                       = "tf_certaction"
		twofactor                  = "ON"
		defaultauthenticationgroup = "new_group"
		usernamefield              = "Subject:CN"
		groupnamefield             = "subject:grp"
	}
	resource "citrixadc_authenticationcertpolicy" "tf_certpolicy" {
		name      = "tf_certpolicy"
		rule      = "ns_false"
		reqaction = citrixadc_authenticationcertaction.tf_certaction.name
	}
`

func TestAccAuthenticationcertpolicy_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationcertpolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationcertpolicy_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationcertpolicyExist("citrixadc_authenticationcertpolicy.tf_certpolicy", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationcertpolicy.tf_certpolicy", "name", "tf_certpolicy"),
					resource.TestCheckResourceAttr("citrixadc_authenticationcertpolicy.tf_certpolicy", "rule", "ns_true"),
				),
			},
			{
				Config: testAccAuthenticationcertpolicy_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationcertpolicyExist("citrixadc_authenticationcertpolicy.tf_certpolicy", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationcertpolicy.tf_certpolicy", "name", "tf_certpolicy"),
					resource.TestCheckResourceAttr("citrixadc_authenticationcertpolicy.tf_certpolicy", "rule", "ns_false"),
				),
			},
		},
	})
}

func testAccCheckAuthenticationcertpolicyExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No authenticationcertpolicy name is set")
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
		data, err := client.FindResource(service.Authenticationcertpolicy.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("authenticationcertpolicy %s not found", n)
		}

		return nil
	}
}

func testAccCheckAuthenticationcertpolicyDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_authenticationcertpolicy" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Authenticationcertpolicy.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("authenticationcertpolicy %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
