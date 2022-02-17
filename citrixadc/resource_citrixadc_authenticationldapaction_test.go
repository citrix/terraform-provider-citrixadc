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
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"testing"
)

const testAccAuthenticationldapaction_add = `
	resource "citrixadc_authenticationldapaction" "foo" {
		name   		  = "ldapaction"
		serverip 	  = "1.2.3.4"
		serverport 	  = 8080
		authtimeout   = 1
		ldaploginname = "username"
	}
`
const testAccAuthenticationldapaction_update = `
	resource "citrixadc_authenticationldapaction" "foo" {
		name   		  = "ldapaction"
		serverip	  = "1.2.4.5"
		serverport    = 8000
		authtimeout   = 2
		ldaploginname = "username"
	}
`

func TestAccAuthenticationldapaction_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAuthenticationldapactionDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccAuthenticationldapaction_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationldapactionExist("citrixadc_authenticationldapaction.foo", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationldapaction.foo", "name", "ldapaction"),
					resource.TestCheckResourceAttr("citrixadc_authenticationldapaction.foo", "serverip", "1.2.3.4"),
					resource.TestCheckResourceAttr("citrixadc_authenticationldapaction.foo", "authtimeout", "1"),
				),
			},
			resource.TestStep{
				Config: testAccAuthenticationldapaction_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationldapactionExist("citrixadc_authenticationldapaction.foo", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationldapaction.foo", "name", "ldapaction"),
					resource.TestCheckResourceAttr("citrixadc_authenticationldapaction.foo", "serverip", "1.2.4.5"),
					resource.TestCheckResourceAttr("citrixadc_authenticationldapaction.foo", "authtimeout", "2"),
				),
			},
		},
	})
}

func testAccCheckAuthenticationldapactionExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No authenticationldapaction name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(service.Authenticationldapaction.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("authenticationldapaction %s not found", n)
		}

		return nil
	}
}

func testAccCheckAuthenticationldapactionDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_authenticationldapaction" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Authenticationldapaction.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("authenticationldapaction %s still exists", rs.Primary.ID)
		}
	}

	return nil
}
