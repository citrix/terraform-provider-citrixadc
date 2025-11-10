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

const testAccAuthenticationsamlpolicy_add = `
	resource "citrixadc_authenticationsamlaction" "tf_samlaction" {
		name                    = "tf_samlaction"
		metadataurl             = "http://www.example.com"
		samltwofactor           = "OFF"
		requestedauthncontext   = "minimum"
		digestmethod            = "SHA1"
		signaturealg            = "RSA-SHA256"
		metadatarefreshinterval = 1
	}
	resource "citrixadc_authenticationsamlpolicy" "tf_samlpolicy" {
		name      = "tf_samlpolicy"
		rule      = "NS_TRUE"
		reqaction = citrixadc_authenticationsamlaction.tf_samlaction.name
	}
`
const testAccAuthenticationsamlpolicy_update = `
	resource "citrixadc_authenticationsamlaction" "tf_samlaction" {
		name                    = "tf_samlaction"
		metadataurl             = "http://www.example.com"
		samltwofactor           = "OFF"
		requestedauthncontext   = "minimum"
		digestmethod            = "SHA1"
		signaturealg            = "RSA-SHA256"
		metadatarefreshinterval = 1
	}
	resource "citrixadc_authenticationsamlpolicy" "tf_samlpolicy" {
		name      = "tf_samlpolicy"
		rule      = "NS_FALSE"
		reqaction = citrixadc_authenticationsamlaction.tf_samlaction.name
	}
`

func TestAccAuthenticationsamlpolicy_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckAuthenticationsamlpolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationsamlpolicy_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationsamlpolicyExist("citrixadc_authenticationsamlpolicy.tf_samlpolicy", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationsamlpolicy.tf_samlpolicy", "name", "tf_samlpolicy"),
					resource.TestCheckResourceAttr("citrixadc_authenticationsamlpolicy.tf_samlpolicy", "rule", "NS_TRUE"),
				),
			},
			{
				Config: testAccAuthenticationsamlpolicy_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationsamlpolicyExist("citrixadc_authenticationsamlpolicy.tf_samlpolicy", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationsamlpolicy.tf_samlpolicy", "name", "tf_samlpolicy"),
					resource.TestCheckResourceAttr("citrixadc_authenticationsamlpolicy.tf_samlpolicy", "rule", "NS_FALSE"),
				),
			},
		},
	})
}

func testAccCheckAuthenticationsamlpolicyExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No authenticationsamlpolicy name is set")
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
		data, err := client.FindResource(service.Authenticationsamlpolicy.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("authenticationsamlpolicy %s not found", n)
		}

		return nil
	}
}

func testAccCheckAuthenticationsamlpolicyDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_authenticationsamlpolicy" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Authenticationsamlpolicy.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("authenticationsamlpolicy %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
