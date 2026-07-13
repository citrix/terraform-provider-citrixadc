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

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testAccAuthenticationemailaction_add = `
	resource "citrixadc_authenticationemailaction" "tf_emailaction" {
		name      = "tf_emailaction"
		username  = "username@abc.com"
		password  = "secret"
		serverurl = "www.example.com"
		timeout   = 100
		type      = "SMTP"
	}
`
const testAccAuthenticationemailaction_update = `
	resource "citrixadc_authenticationemailaction" "tf_emailaction" {
		name      = "tf_emailaction"
		username  = "username1@abc.com"
		password  = "secret"
		serverurl = "www.example1.com"
		timeout   = 100
		type      = "SMTP"
	}
`

func TestAccAuthenticationemailaction_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationemailactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationemailaction_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationemailactionExist("citrixadc_authenticationemailaction.tf_emailaction", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationemailaction.tf_emailaction", "name", "tf_emailaction"),
					resource.TestCheckResourceAttr("citrixadc_authenticationemailaction.tf_emailaction", "username", "username@abc.com"),
					resource.TestCheckResourceAttr("citrixadc_authenticationemailaction.tf_emailaction", "serverurl", "www.example.com"),
				),
			},
			{
				Config: testAccAuthenticationemailaction_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationemailactionExist("citrixadc_authenticationemailaction.tf_emailaction", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationemailaction.tf_emailaction", "name", "tf_emailaction"),
					resource.TestCheckResourceAttr("citrixadc_authenticationemailaction.tf_emailaction", "username", "username1@abc.com"),
					resource.TestCheckResourceAttr("citrixadc_authenticationemailaction.tf_emailaction", "serverurl", "www.example1.com"),
				),
			},
		},
	})
}

func testAccCheckAuthenticationemailactionExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No authenticationemailaction name is set")
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
		data, err := client.FindResource("authenticationemailaction", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("authenticationemailaction %s not found", n)
		}

		return nil
	}
}

func testAccCheckAuthenticationemailactionDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_authenticationemailaction" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("authenticationemailaction", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("authenticationemailaction %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccAuthenticationemailactionDataSource_basic = `
	resource "citrixadc_authenticationemailaction" "tf_emailaction" {
		name      = "tf_emailaction_ds"
		username  = "username@abc.com"
		password  = "secret"
		serverurl = "www.example.com"
		timeout   = 100
		type      = "SMTP"
	}

	data "citrixadc_authenticationemailaction" "tf_emailaction" {
		name       = citrixadc_authenticationemailaction.tf_emailaction.name
		depends_on = [citrixadc_authenticationemailaction.tf_emailaction]
	}
`

func TestAccAuthenticationemailactionDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationemailactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationemailactionDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_authenticationemailaction.tf_emailaction", "name", "tf_emailaction_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationemailaction.tf_emailaction", "username", "username@abc.com"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationemailaction.tf_emailaction", "serverurl", "www.example.com"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationemailaction.tf_emailaction", "type", "SMTP"),
				),
			},
		},
	})
}

// Backward-compatible test: uses the sensitive `password` attribute path
const testAccAuthenticationemailaction_password_step1 = `
	variable "authenticationemailaction_password" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_authenticationemailaction" "tf_emailaction" {
		name      = "tf_emailaction"
		username  = "username@abc.com"
		password  = var.authenticationemailaction_password
		serverurl = "www.example.com"
		timeout   = 100
		type      = "SMTP"
	}
`

const testAccAuthenticationemailaction_password_step2 = `
	variable "authenticationemailaction_password_2" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_authenticationemailaction" "tf_emailaction" {
		name      = "tf_emailaction"
		username  = "username@abc.com"
		password  = var.authenticationemailaction_password_2
		serverurl = "www.example.com"
		timeout   = 100
		type      = "SMTP"
	}
`

func TestAccAuthenticationemailaction_password_backward_compat(t *testing.T) {
	t.Setenv("TF_VAR_authenticationemailaction_password", "secret1")
	t.Setenv("TF_VAR_authenticationemailaction_password_2", "secret2")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationemailactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationemailaction_password_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationemailactionExist("citrixadc_authenticationemailaction.tf_emailaction", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationemailaction.tf_emailaction", "name", "tf_emailaction"),
					resource.TestCheckResourceAttr("citrixadc_authenticationemailaction.tf_emailaction", "username", "username@abc.com"),
					resource.TestCheckResourceAttr("citrixadc_authenticationemailaction.tf_emailaction", "serverurl", "www.example.com"),
				),
			},
			{
				Config: testAccAuthenticationemailaction_password_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationemailactionExist("citrixadc_authenticationemailaction.tf_emailaction", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationemailaction.tf_emailaction", "name", "tf_emailaction"),
					resource.TestCheckResourceAttr("citrixadc_authenticationemailaction.tf_emailaction", "username", "username@abc.com"),
					resource.TestCheckResourceAttr("citrixadc_authenticationemailaction.tf_emailaction", "serverurl", "www.example.com"),
				),
			},
		},
	})
}

// Ephemeral write-only test: uses the `password_wo` + `password_wo_version` path
const testAccAuthenticationemailaction_wo_step1 = `
	variable "authenticationemailaction_password_wo" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_authenticationemailaction" "tf_emailaction" {
		name               = "tf_emailaction"
		username           = "username@abc.com"
		password_wo        = var.authenticationemailaction_password_wo
		password_wo_version = 1
		serverurl          = "www.example.com"
		timeout            = 100
		type               = "SMTP"
	}
`

const testAccAuthenticationemailaction_wo_step2 = `
	variable "authenticationemailaction_password_wo_2" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_authenticationemailaction" "tf_emailaction" {
		name               = "tf_emailaction"
		username           = "username@abc.com"
		password_wo        = var.authenticationemailaction_password_wo_2
		password_wo_version = 2
		serverurl          = "www.example.com"
		timeout            = 100
		type               = "SMTP"
	}
`

func TestAccAuthenticationemailaction_password_wo_ephemeral(t *testing.T) {
	t.Setenv("TF_VAR_authenticationemailaction_password_wo", "ephemeral_secret1")
	t.Setenv("TF_VAR_authenticationemailaction_password_wo_2", "ephemeral_secret2")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationemailactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationemailaction_wo_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationemailactionExist("citrixadc_authenticationemailaction.tf_emailaction", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationemailaction.tf_emailaction", "name", "tf_emailaction"),
					resource.TestCheckResourceAttr("citrixadc_authenticationemailaction.tf_emailaction", "password_wo_version", "1"),
				),
			},
			{
				Config: testAccAuthenticationemailaction_wo_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationemailactionExist("citrixadc_authenticationemailaction.tf_emailaction", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationemailaction.tf_emailaction", "name", "tf_emailaction"),
					resource.TestCheckResourceAttr("citrixadc_authenticationemailaction.tf_emailaction", "password_wo_version", "2"),
				),
			},
		},
	})
}

func TestAccAuthenticationemailaction_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckAuthenticationemailactionDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccAuthenticationemailaction_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationemailactionExist("citrixadc_authenticationemailaction.tf_emailaction", nil),
				),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccAuthenticationemailaction_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationemailactionExist("citrixadc_authenticationemailaction.tf_emailaction", nil),
				),
			},
		},
	})
}
