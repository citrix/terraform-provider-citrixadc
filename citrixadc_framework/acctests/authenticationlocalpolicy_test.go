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

const testAccAuthenticationlocalpolicy_add = `
	resource "citrixadc_authenticationlocalpolicy" "tf_authenticationlocalpolicy" {
		name   = "tf_authenticationlocalpolicy"
		rule   = "ns_true"
	}
`
const testAccAuthenticationlocalpolicy_update = `
	resource "citrixadc_authenticationlocalpolicy" "tf_authenticationlocalpolicy" {
		name   = "tf_authenticationlocalpolicy"
		rule   = "ns_false"
	}
`

func TestAccAuthenticationlocalpolicy_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationlocalpolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationlocalpolicy_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationlocalpolicyExist("citrixadc_authenticationlocalpolicy.tf_authenticationlocalpolicy", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationlocalpolicy.tf_authenticationlocalpolicy", "name", "tf_authenticationlocalpolicy"),
					resource.TestCheckResourceAttr("citrixadc_authenticationlocalpolicy.tf_authenticationlocalpolicy", "rule", "ns_true"),
				),
			},
			{
				Config: testAccAuthenticationlocalpolicy_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationlocalpolicyExist("citrixadc_authenticationlocalpolicy.tf_authenticationlocalpolicy", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationlocalpolicy.tf_authenticationlocalpolicy", "name", "tf_authenticationlocalpolicy"),
					resource.TestCheckResourceAttr("citrixadc_authenticationlocalpolicy.tf_authenticationlocalpolicy", "rule", "ns_false"),
				),
			},
		},
	})
}

func testAccCheckAuthenticationlocalpolicyExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No authenticationlocalpolicy name is set")
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
		data, err := client.FindResource(service.Authenticationlocalpolicy.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("authenticationlocalpolicy %s not found", n)
		}

		return nil
	}
}

func testAccCheckAuthenticationlocalpolicyDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_authenticationlocalpolicy" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Authenticationlocalpolicy.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("authenticationlocalpolicy %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
