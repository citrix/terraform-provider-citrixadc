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

const testAccAuthenticationtacacspolicy_add = `
	resource "citrixadc_authenticationtacacsaction" "tf_tacacsaction" {
		name            = "tf_tacacsaction"
		serverip        = "1.2.3.4"
		serverport      = 8080
		authtimeout     = 5
		authorization   = "ON"
		accounting      = "ON"
		auditfailedcmds = "ON"
		groupattrname   = "group"
	}
	resource "citrixadc_authenticationtacacspolicy" "tf_tacacspolicy" {
		name= "tf_tacacspolicy"
		rule= "NS_TRUE"
		reqaction= citrixadc_authenticationtacacsaction.tf_tacacsaction.name
		
	}
`
const testAccAuthenticationtacacspolicy_update = `
	resource "citrixadc_authenticationtacacsaction" "tf_tacacsaction" {
		name            = "tf_tacacsaction"
		serverip        = "1.2.3.4"
		serverport      = 8080
		authtimeout     = 5
		authorization   = "ON"
		accounting      = "ON"
		auditfailedcmds = "ON"
		groupattrname   = "group"
	}
	resource "citrixadc_authenticationtacacspolicy" "tf_tacacspolicy" {
		name= "tf_tacacspolicy"
		rule= "NS_FALSE"
		reqaction= citrixadc_authenticationtacacsaction.tf_tacacsaction.name
		
	}
`

func TestAccAuthenticationtacacspolicy_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationtacacspolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationtacacspolicy_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationtacacspolicyExist("citrixadc_authenticationtacacspolicy.tf_tacacspolicy", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationtacacspolicy.tf_tacacspolicy", "name", "tf_tacacspolicy"),
					resource.TestCheckResourceAttr("citrixadc_authenticationtacacspolicy.tf_tacacspolicy", "rule", "NS_TRUE"),
				),
			},
			{
				Config: testAccAuthenticationtacacspolicy_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationtacacspolicyExist("citrixadc_authenticationtacacspolicy.tf_tacacspolicy", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationtacacspolicy.tf_tacacspolicy", "name", "tf_tacacspolicy"),
					resource.TestCheckResourceAttr("citrixadc_authenticationtacacspolicy.tf_tacacspolicy", "rule", "NS_FALSE"),
				),
			},
		},
	})
}

func testAccCheckAuthenticationtacacspolicyExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No authenticationtacacspolicy name is set")
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
		data, err := client.FindResource(service.Authenticationtacacspolicy.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("authenticationtacacspolicy %s not found", n)
		}

		return nil
	}
}

func testAccCheckAuthenticationtacacspolicyDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_authenticationtacacspolicy" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Authenticationtacacspolicy.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("authenticationtacacspolicy %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
