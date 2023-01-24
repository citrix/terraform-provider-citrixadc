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
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"testing"
)

const testAccAuthenticationldappolicy_add = `
	resource "citrixadc_authenticationldapaction" "tf_authenticationldapaction" {
		name          = "ldapaction"
		serverip      = "1.2.3.4"
		serverport    = 8080
		authtimeout   = 1
		ldaploginname = "username"
	}
	resource "citrixadc_authenticationldappolicy" "tf_authenticationldappolicy" {
		name = "tf_authenticationldappolicy"
		rule = "NS_TRUE"
		reqaction = citrixadc_authenticationldapaction.tf_authenticationldapaction.name
	}
`
const testAccAuthenticationldappolicy_update = `
	resource "citrixadc_authenticationldapaction" "tf_authenticationldapaction" {
		name          = "ldapaction"
		serverip      = "1.2.3.4"
		serverport    = 8080
		authtimeout   = 1
		ldaploginname = "username"
	}
	resource "citrixadc_authenticationldappolicy" "tf_authenticationldappolicy" {
		name   	  = "tf_authenticationldappolicy"
		rule	  = "NS_FALSE"
		reqaction = citrixadc_authenticationldapaction.tf_authenticationldapaction.name
	}
`

func TestAccAuthenticationldappolicy_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAuthenticationldappolicyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccAuthenticationldappolicy_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationldappolicyExist("citrixadc_authenticationldappolicy.tf_authenticationldappolicy", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationldappolicy.tf_authenticationldappolicy", "name", "tf_authenticationldappolicy"),
					resource.TestCheckResourceAttr("citrixadc_authenticationldappolicy.tf_authenticationldappolicy", "rule", "NS_TRUE"),
				),
			},
			resource.TestStep{
				Config: testAccAuthenticationldappolicy_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationldappolicyExist("citrixadc_authenticationldappolicy.tf_authenticationldappolicy", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationldappolicy.tf_authenticationldappolicy", "name", "tf_authenticationldappolicy"),
					resource.TestCheckResourceAttr("citrixadc_authenticationldappolicy.tf_authenticationldappolicy", "rule", "NS_FALSE"),
				),
			},
		},
	})
}

func testAccCheckAuthenticationldappolicyExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No authenticationldappolicy name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(service.Authenticationldappolicy.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("authenticationldappolicy %s not found", n)
		}

		return nil
	}
}

func testAccCheckAuthenticationldappolicyDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_authenticationldappolicy" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Authenticationldappolicy.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("authenticationldappolicy %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
