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

const testAccAuthenticationnegotiateaction_add = `
	resource "citrixadc_authenticationnegotiateaction" "tf_negotiateaction" {
		name                       = "tf_negotiateaction"
		domain                     = "DomainName"
		domainuser                 = "username"
		domainuserpasswd           = "password"
		ntlmpath                   = "http://www.example.com/"
		defaultauthenticationgroup = "grpname"
	}
`
const testAccAuthenticationnegotiateaction_update = `
	resource "citrixadc_authenticationnegotiateaction" "tf_negotiateaction" {
		name                       = "tf_negotiateaction"
		domain                     = "DomainName"
		domainuser                 = "new_username"
		domainuserpasswd           = "password"
		ntlmpath                   = "http://www.example.com/"
		defaultauthenticationgroup = "new_grpname"
	}
`

func TestAccAuthenticationnegotiateaction_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationnegotiateactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationnegotiateaction_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationnegotiateactionExist("citrixadc_authenticationnegotiateaction.tf_negotiateaction", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationnegotiateaction.tf_negotiateaction", "name", "tf_negotiateaction"),
					resource.TestCheckResourceAttr("citrixadc_authenticationnegotiateaction.tf_negotiateaction", "domainuser", "username"),
					resource.TestCheckResourceAttr("citrixadc_authenticationnegotiateaction.tf_negotiateaction", "defaultauthenticationgroup", "grpname"),
				),
			},
			{
				Config: testAccAuthenticationnegotiateaction_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationnegotiateactionExist("citrixadc_authenticationnegotiateaction.tf_negotiateaction", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationnegotiateaction.tf_negotiateaction", "name", "tf_negotiateaction"),
					resource.TestCheckResourceAttr("citrixadc_authenticationnegotiateaction.tf_negotiateaction", "domainuser", "new_username"),
					resource.TestCheckResourceAttr("citrixadc_authenticationnegotiateaction.tf_negotiateaction", "defaultauthenticationgroup", "new_grpname"),
				),
			},
		},
	})
}

func testAccCheckAuthenticationnegotiateactionExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No authenticationnegotiateaction name is set")
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
		data, err := client.FindResource(service.Authenticationnegotiateaction.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("authenticationnegotiateaction %s not found", n)
		}

		return nil
	}
}

func testAccCheckAuthenticationnegotiateactionDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_authenticationnegotiateaction" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Authenticationnegotiateaction.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("authenticationnegotiateaction %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccAuthenticationnegotiateactionDataSource_basic = `
	resource "citrixadc_authenticationnegotiateaction" "tf_negotiateaction_ds" {
		name                       = "tf_negotiateaction_ds"
		domain                     = "DomainName"
		domainuser                 = "username"
		domainuserpasswd           = "password"
		ntlmpath                   = "http://www.example.com/"
		defaultauthenticationgroup = "grpname"
	}

	data "citrixadc_authenticationnegotiateaction" "tf_negotiateaction_ds" {
		name = citrixadc_authenticationnegotiateaction.tf_negotiateaction_ds.name
	}
`

const testAccAuthenticationnegotiateaction_domainuserpasswd_step1 = `
	variable "authenticationnegotiateaction_domainuserpasswd" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_authenticationnegotiateaction" "tf_negotiateaction_bc" {
		name                       = "tf_negotiateaction_bc"
		domain                     = "DomainName"
		domainuser                 = "username"
		domainuserpasswd           = var.authenticationnegotiateaction_domainuserpasswd
		ntlmpath                   = "http://www.example.com/"
		defaultauthenticationgroup = "grpname"
	}
`

const testAccAuthenticationnegotiateaction_domainuserpasswd_step2 = `
	variable "authenticationnegotiateaction_domainuserpasswd_2" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_authenticationnegotiateaction" "tf_negotiateaction_bc" {
		name                       = "tf_negotiateaction_bc"
		domain                     = "DomainName"
		domainuser                 = "username"
		domainuserpasswd           = var.authenticationnegotiateaction_domainuserpasswd_2
		ntlmpath                   = "http://www.example.com/"
		defaultauthenticationgroup = "grpname"
	}
`

func TestAccAuthenticationnegotiateaction_domainuserpasswd_backward_compat(t *testing.T) {
	t.Setenv("TF_VAR_authenticationnegotiateaction_domainuserpasswd", "password1")
	t.Setenv("TF_VAR_authenticationnegotiateaction_domainuserpasswd_2", "password2")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationnegotiateactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationnegotiateaction_domainuserpasswd_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationnegotiateactionExist("citrixadc_authenticationnegotiateaction.tf_negotiateaction_bc", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationnegotiateaction.tf_negotiateaction_bc", "name", "tf_negotiateaction_bc"),
					resource.TestCheckResourceAttr("citrixadc_authenticationnegotiateaction.tf_negotiateaction_bc", "domainuser", "username"),
				),
			},
			{
				Config: testAccAuthenticationnegotiateaction_domainuserpasswd_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationnegotiateactionExist("citrixadc_authenticationnegotiateaction.tf_negotiateaction_bc", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationnegotiateaction.tf_negotiateaction_bc", "name", "tf_negotiateaction_bc"),
					resource.TestCheckResourceAttr("citrixadc_authenticationnegotiateaction.tf_negotiateaction_bc", "domainuser", "username"),
				),
			},
		},
	})
}

const testAccAuthenticationnegotiateaction_domainuserpasswd_wo_step1 = `
	variable "authenticationnegotiateaction_domainuserpasswd_wo" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_authenticationnegotiateaction" "tf_negotiateaction_wo" {
		name                       = "tf_negotiateaction_wo"
		domain                     = "DomainName"
		domainuser                 = "username"
		domainuserpasswd_wo        = var.authenticationnegotiateaction_domainuserpasswd_wo
		domainuserpasswd_wo_version = 1
		ntlmpath                   = "http://www.example.com/"
		defaultauthenticationgroup = "grpname"
	}
`

const testAccAuthenticationnegotiateaction_domainuserpasswd_wo_step2 = `
	variable "authenticationnegotiateaction_domainuserpasswd_wo_2" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_authenticationnegotiateaction" "tf_negotiateaction_wo" {
		name                       = "tf_negotiateaction_wo"
		domain                     = "DomainName"
		domainuser                 = "username"
		domainuserpasswd_wo        = var.authenticationnegotiateaction_domainuserpasswd_wo_2
		domainuserpasswd_wo_version = 2
		ntlmpath                   = "http://www.example.com/"
		defaultauthenticationgroup = "grpname"
	}
`

func TestAccAuthenticationnegotiateaction_domainuserpasswd_wo_ephemeral(t *testing.T) {
	t.Setenv("TF_VAR_authenticationnegotiateaction_domainuserpasswd_wo", "ephemeral_value1")
	t.Setenv("TF_VAR_authenticationnegotiateaction_domainuserpasswd_wo_2", "ephemeral_value2")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationnegotiateactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationnegotiateaction_domainuserpasswd_wo_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationnegotiateactionExist("citrixadc_authenticationnegotiateaction.tf_negotiateaction_wo", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationnegotiateaction.tf_negotiateaction_wo", "domainuserpasswd_wo_version", "1"),
				),
			},
			{
				Config: testAccAuthenticationnegotiateaction_domainuserpasswd_wo_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationnegotiateactionExist("citrixadc_authenticationnegotiateaction.tf_negotiateaction_wo", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationnegotiateaction.tf_negotiateaction_wo", "domainuserpasswd_wo_version", "2"),
				),
			},
		},
	})
}

func TestAccAuthenticationnegotiateaction_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckAuthenticationnegotiateactionDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccAuthenticationnegotiateaction_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationnegotiateactionExist("citrixadc_authenticationnegotiateaction.tf_negotiateaction", nil),
				),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccAuthenticationnegotiateaction_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationnegotiateactionExist("citrixadc_authenticationnegotiateaction.tf_negotiateaction", nil),
				),
			},
		},
	})
}

func TestAccAuthenticationnegotiateactionDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationnegotiateactionDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_authenticationnegotiateaction.tf_negotiateaction_ds", "name", "tf_negotiateaction_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationnegotiateaction.tf_negotiateaction_ds", "domain", "DomainName"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationnegotiateaction.tf_negotiateaction_ds", "domainuser", "username"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationnegotiateaction.tf_negotiateaction_ds", "ntlmpath", "http://www.example.com/"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationnegotiateaction.tf_negotiateaction_ds", "defaultauthenticationgroup", "grpname"),
				),
			},
		},
	})
}
