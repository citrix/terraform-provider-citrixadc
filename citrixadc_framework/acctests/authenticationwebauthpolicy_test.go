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

const testAccAuthenticationwebauthpolicy_add = `
	resource "citrixadc_authenticationwebauthaction" "tf_webauthaction" {
		name                       = "tf_webauthaction"
		serverip                   = "1.2.3.4"
		serverport                 = 8080
		fullreqexpr                = "TRUE"
		scheme                     = "http"
		successrule                = "http.RES.STATUS.EQ(200)"
		defaultauthenticationgroup = "new_group"
	}
	resource "citrixadc_authenticationwebauthpolicy" "tf_webauthpolicy" {
		name   = "tf_webauthpolicy"
		rule   = "NS_TRUE"
		action = citrixadc_authenticationwebauthaction.tf_webauthaction.name
	}
`
const testAccAuthenticationwebauthpolicy_update = `
	resource "citrixadc_authenticationwebauthaction" "tf_webauthaction" {
		name                       = "tf_webauthaction"
		serverip                   = "1.2.3.4"
		serverport                 = 8080
		fullreqexpr                = "TRUE"
		scheme                     = "http"
		successrule                = "http.RES.STATUS.EQ(200)"
		defaultauthenticationgroup = "new_group"
	}
	resource "citrixadc_authenticationwebauthpolicy" "tf_webauthpolicy" {
		name   = "tf_webauthpolicy"
		rule   = "NS_FALSE"
		action = citrixadc_authenticationwebauthaction.tf_webauthaction.name
	}
`

func TestAccAuthenticationwebauthpolicy_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationwebauthpolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationwebauthpolicy_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationwebauthpolicyExist("citrixadc_authenticationwebauthpolicy.tf_webauthpolicy", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationwebauthpolicy.tf_webauthpolicy", "name", "tf_webauthpolicy"),
					resource.TestCheckResourceAttr("citrixadc_authenticationwebauthpolicy.tf_webauthpolicy", "rule", "NS_TRUE"),
				),
			},
			{
				Config: testAccAuthenticationwebauthpolicy_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationwebauthpolicyExist("citrixadc_authenticationwebauthpolicy.tf_webauthpolicy", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationwebauthpolicy.tf_webauthpolicy", "name", "tf_webauthpolicy"),
					resource.TestCheckResourceAttr("citrixadc_authenticationwebauthpolicy.tf_webauthpolicy", "rule", "NS_FALSE"),
				),
			},
		},
	})
}

func testAccCheckAuthenticationwebauthpolicyExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No authenticationwebauthpolicy name is set")
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
		data, err := client.FindResource(service.Authenticationwebauthpolicy.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("authenticationwebauthpolicy %s not found", n)
		}

		return nil
	}
}

func testAccCheckAuthenticationwebauthpolicyDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_authenticationwebauthpolicy" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Authenticationwebauthpolicy.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("authenticationwebauthpolicy %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccAuthenticationwebauthpolicyDataSource_basic = `
	resource "citrixadc_authenticationwebauthaction" "tf_webauthaction" {
		name                       = "tf_webauthaction"
		serverip                   = "1.2.3.4"
		serverport                 = 8080
		fullreqexpr                = "TRUE"
		scheme                     = "http"
		successrule                = "http.RES.STATUS.EQ(200)"
		defaultauthenticationgroup = "new_group"
	}
	resource "citrixadc_authenticationwebauthpolicy" "tf_webauthpolicy" {
		name   = "tf_webauthpolicy_ds"
		rule   = "NS_TRUE"
		action = citrixadc_authenticationwebauthaction.tf_webauthaction.name
	}

	data "citrixadc_authenticationwebauthpolicy" "tf_webauthpolicy_ds" {
		name = citrixadc_authenticationwebauthpolicy.tf_webauthpolicy.name
	}
`

func TestAccAuthenticationwebauthpolicyDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationwebauthpolicyDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_authenticationwebauthpolicy.tf_webauthpolicy_ds", "name", "tf_webauthpolicy_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationwebauthpolicy.tf_webauthpolicy_ds", "rule", "NS_TRUE"),
					resource.TestCheckResourceAttrPair("data.citrixadc_authenticationwebauthpolicy.tf_webauthpolicy_ds", "action", "citrixadc_authenticationwebauthaction.tf_webauthaction", "name"),
				),
			},
		},
	})
}
