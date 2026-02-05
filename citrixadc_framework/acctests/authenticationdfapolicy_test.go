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

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testAccAuthenticationdfapolicy_add = `
	resource "citrixadc_authenticationdfaaction" "tf_dfaaction" {
		name       = "tf_dfaaction"
		serverurl  = "https://example.com/"
		clientid   = "cliId"
		passphrase = "secret"
	}
	resource "citrixadc_authenticationdfapolicy" "td_dfapolicy" {
		name   = "td_dfapolicy"
		rule   = "NS_TRUE"
		action = citrixadc_authenticationdfaaction.tf_dfaaction.name
	}
`
const testAccAuthenticationdfapolicy_update = `
	resource "citrixadc_authenticationdfaaction" "tf_dfaaction" {
		name       = "tf_dfaaction"
		serverurl  = "https://example.com/"
		clientid   = "cliId"
		passphrase = "secret"
	}
	resource "citrixadc_authenticationdfapolicy" "td_dfapolicy" {
		name   = "td_dfapolicy"
		rule   = "NS_FALSE"
		action = citrixadc_authenticationdfaaction.tf_dfaaction.name
	}
`

func TestAccAuthenticationdfapolicy_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationdfapolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationdfapolicy_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationdfapolicyExist("citrixadc_authenticationdfapolicy.td_dfapolicy", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationdfapolicy.td_dfapolicy", "name", "td_dfapolicy"),
					resource.TestCheckResourceAttr("citrixadc_authenticationdfapolicy.td_dfapolicy", "rule", "NS_TRUE"),
				),
			},
			{
				Config: testAccAuthenticationdfapolicy_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationdfapolicyExist("citrixadc_authenticationdfapolicy.td_dfapolicy", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationdfapolicy.td_dfapolicy", "name", "td_dfapolicy"),
					resource.TestCheckResourceAttr("citrixadc_authenticationdfapolicy.td_dfapolicy", "rule", "NS_FALSE"),
				),
			},
		},
	})
}

func testAccCheckAuthenticationdfapolicyExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No authenticationdfapolicy name is set")
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
		data, err := client.FindResource(service.Authenticationdfapolicy.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("authenticationdfapolicy %s not found", n)
		}

		return nil
	}
}

func testAccCheckAuthenticationdfapolicyDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_authenticationdfapolicy" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Authenticationdfapolicy.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("authenticationdfapolicy %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccAuthenticationdfapolicyDataSource_basic = `
	resource "citrixadc_authenticationdfaaction" "tf_dfaaction" {
		name       = "tf_dfaaction"
		serverurl  = "https://example.com/"
		clientid   = "cliId"
		passphrase = "secret"
	}
	resource "citrixadc_authenticationdfapolicy" "td_dfapolicy" {
		name   = "td_dfapolicy"
		rule   = "NS_TRUE"
		action = citrixadc_authenticationdfaaction.tf_dfaaction.name
	}

	data "citrixadc_authenticationdfapolicy" "td_dfapolicy" {
		name = citrixadc_authenticationdfapolicy.td_dfapolicy.name
		depends_on = [citrixadc_authenticationdfapolicy.td_dfapolicy]
	}
`

func TestAccAuthenticationdfapolicyDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationdfapolicyDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_authenticationdfapolicy.td_dfapolicy", "name", "td_dfapolicy"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationdfapolicy.td_dfapolicy", "rule", "NS_TRUE"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationdfapolicy.td_dfapolicy", "action", "tf_dfaaction"),
				),
			},
		},
	})
}
