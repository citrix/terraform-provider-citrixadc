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

const testAccAuthenticationnegotiatepolicy_add = `

	resource "citrixadc_authenticationnegotiateaction" "tf_negotiateaction" {
		name                       = "tf_negotiateaction"
		domain                     = "DomainName"
		domainuser                 = "usersame"
		domainuserpasswd           = "password"
		ntlmpath                   = "http://www.example.com/"
		defaultauthenticationgroup = "new_grpname"
	}
	resource "citrixadc_authenticationnegotiatepolicy" "tf_negotiatepolicy" {
		name = "tf_negotiatepolicy"
		rule = "ns_true"
		reqaction = citrixadc_authenticationnegotiateaction.tf_negotiateaction.name
	}
`
const testAccAuthenticationnegotiatepolicy_update = `

	resource "citrixadc_authenticationnegotiateaction" "tf_negotiateaction" {
		name                       = "tf_negotiateaction"
		domain                     = "DomainName"
		domainuser                 = "usersame"
		domainuserpasswd           = "password"
		ntlmpath                   = "http://www.example.com/"
		defaultauthenticationgroup = "new_grpname"
	}
	resource "citrixadc_authenticationnegotiatepolicy" "tf_negotiatepolicy" {
		name = "tf_negotiatepolicy"
		rule = "ns_false"
		reqaction = citrixadc_authenticationnegotiateaction.tf_negotiateaction.name
	}
`

func TestAccAuthenticationnegotiatepolicy_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAuthenticationnegotiatepolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationnegotiatepolicy_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationnegotiatepolicyExist("citrixadc_authenticationnegotiatepolicy.tf_negotiatepolicy", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationnegotiatepolicy.tf_negotiatepolicy", "name", "tf_negotiatepolicy"),
					resource.TestCheckResourceAttr("citrixadc_authenticationnegotiatepolicy.tf_negotiatepolicy", "rule", "ns_true"),
				),
			},
			{
				Config: testAccAuthenticationnegotiatepolicy_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationnegotiatepolicyExist("citrixadc_authenticationnegotiatepolicy.tf_negotiatepolicy", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationnegotiatepolicy.tf_negotiatepolicy", "name", "tf_negotiatepolicy"),
					resource.TestCheckResourceAttr("citrixadc_authenticationnegotiatepolicy.tf_negotiatepolicy", "rule", "ns_false"),
				),
			},
		},
	})
}

func testAccCheckAuthenticationnegotiatepolicyExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No authenticationnegotiatepolicy name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(service.Authenticationnegotiatepolicy.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("authenticationnegotiatepolicy %s not found", n)
		}

		return nil
	}
}

func testAccCheckAuthenticationnegotiatepolicyDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_authenticationnegotiatepolicy" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Authenticationnegotiatepolicy.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("authenticationnegotiatepolicy %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
