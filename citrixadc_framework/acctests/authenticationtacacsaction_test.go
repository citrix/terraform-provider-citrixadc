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

const testAccAuthenticationtacacsaction_add = `
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
`
const testAccAuthenticationtacacsaction_update = `
	resource "citrixadc_authenticationtacacsaction" "tf_tacacsaction" {
		name            = "tf_tacacsaction"
		serverip        = "1.2.3.4"
		serverport      = 8080
		authtimeout     = 5
		authorization   = "OFF"
		accounting      = "OFF"
		auditfailedcmds = "ON"
		groupattrname   = "group"
	}
`

func TestAccAuthenticationtacacsaction_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationtacacsactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationtacacsaction_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationtacacsactionExist("citrixadc_authenticationtacacsaction.tf_tacacsaction", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationtacacsaction.tf_tacacsaction", "name", "tf_tacacsaction"),
					resource.TestCheckResourceAttr("citrixadc_authenticationtacacsaction.tf_tacacsaction", "authorization", "ON"),
					resource.TestCheckResourceAttr("citrixadc_authenticationtacacsaction.tf_tacacsaction", "accounting", "ON"),
				),
			},
			{
				Config: testAccAuthenticationtacacsaction_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationtacacsactionExist("citrixadc_authenticationtacacsaction.tf_tacacsaction", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationtacacsaction.tf_tacacsaction", "name", "tf_tacacsaction"),
					resource.TestCheckResourceAttr("citrixadc_authenticationtacacsaction.tf_tacacsaction", "authorization", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_authenticationtacacsaction.tf_tacacsaction", "accounting", "OFF"),
				),
			},
		},
	})
}

func testAccCheckAuthenticationtacacsactionExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No authenticationtacacsaction name is set")
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
		data, err := client.FindResource(service.Authenticationtacacsaction.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("authenticationtacacsaction %s not found", n)
		}

		return nil
	}
}

func testAccCheckAuthenticationtacacsactionDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_authenticationtacacsaction" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Authenticationtacacsaction.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("authenticationtacacsaction %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

// Test backward-compatible path: using tacacssecret (Sensitive attribute)
const testAccAuthenticationtacacsaction_tacacssecret_step1 = `

	variable "tacacsaction_tacacssecret" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_authenticationtacacsaction" "tf_tacacsaction" {
		name          = "tf_tacacsaction"
		serverip      = "1.2.3.4"
		serverport    = 8080
		authtimeout   = 5
		tacacssecret  = var.tacacsaction_tacacssecret
		authorization = "ON"
	}
`

const testAccAuthenticationtacacsaction_tacacssecret_step2 = `

	variable "tacacsaction_tacacssecret_2" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_authenticationtacacsaction" "tf_tacacsaction" {
		name          = "tf_tacacsaction"
		serverip      = "1.2.3.4"
		serverport    = 8080
		authtimeout   = 5
		tacacssecret  = var.tacacsaction_tacacssecret_2
		authorization = "ON"
	}
`

func TestAccAuthenticationtacacsaction_tacacssecret_backward_compat(t *testing.T) {
	t.Setenv("TF_VAR_tacacsaction_tacacssecret", "oldsecret123")
	t.Setenv("TF_VAR_tacacsaction_tacacssecret_2", "newsecret456")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationtacacsactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationtacacsaction_tacacssecret_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationtacacsactionExist("citrixadc_authenticationtacacsaction.tf_tacacsaction", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationtacacsaction.tf_tacacsaction", "name", "tf_tacacsaction"),
					resource.TestCheckResourceAttr("citrixadc_authenticationtacacsaction.tf_tacacsaction", "authorization", "ON"),
				),
			},
			{
				Config: testAccAuthenticationtacacsaction_tacacssecret_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationtacacsactionExist("citrixadc_authenticationtacacsaction.tf_tacacsaction", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationtacacsaction.tf_tacacsaction", "name", "tf_tacacsaction"),
					resource.TestCheckResourceAttr("citrixadc_authenticationtacacsaction.tf_tacacsaction", "authorization", "ON"),
				),
			},
		},
	})
}

// Test ephemeral path: using tacacssecret_wo (WriteOnly attribute) with version tracker
const testAccAuthenticationtacacsaction_tacacssecret_wo_step1 = `

	variable "tacacsaction_tacacssecret_wo" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_authenticationtacacsaction" "tf_tacacsaction" {
		name                   = "tf_tacacsaction"
		serverip               = "1.2.3.4"
		serverport             = 8080
		authtimeout            = 5
		tacacssecret_wo        = var.tacacsaction_tacacssecret_wo
		tacacssecret_wo_version = 1
		authorization          = "ON"
	}
`

const testAccAuthenticationtacacsaction_tacacssecret_wo_step2 = `

	variable "tacacsaction_tacacssecret_wo_2" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_authenticationtacacsaction" "tf_tacacsaction" {
		name                   = "tf_tacacsaction"
		serverip               = "1.2.3.4"
		serverport             = 8080
		authtimeout            = 5
		tacacssecret_wo        = var.tacacsaction_tacacssecret_wo_2
		tacacssecret_wo_version = 2
		authorization          = "ON"
	}
`

func TestAccAuthenticationtacacsaction_tacacssecret_wo_ephemeral(t *testing.T) {
	t.Setenv("TF_VAR_tacacsaction_tacacssecret_wo", "ephemeral_tacacs1")
	t.Setenv("TF_VAR_tacacsaction_tacacssecret_wo_2", "ephemeral_tacacs2")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationtacacsactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationtacacsaction_tacacssecret_wo_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationtacacsactionExist("citrixadc_authenticationtacacsaction.tf_tacacsaction", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationtacacsaction.tf_tacacsaction", "tacacssecret_wo_version", "1"),
				),
			},
			{
				Config: testAccAuthenticationtacacsaction_tacacssecret_wo_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationtacacsactionExist("citrixadc_authenticationtacacsaction.tf_tacacsaction", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationtacacsaction.tf_tacacsaction", "tacacssecret_wo_version", "2"),
				),
			},
		},
	})
}

const testAccAuthenticationtacacsactionDataSource_basic = `
	resource "citrixadc_authenticationtacacsaction" "tf_tacacsaction" {
		name            = "tf_tacacsaction_ds"
		serverip        = "1.2.3.5"
		serverport      = 49
		authtimeout     = 10
		authorization   = "ON"
		accounting      = "ON"
		auditfailedcmds = "ON"
		groupattrname   = "group"
	}

	data "citrixadc_authenticationtacacsaction" "tf_tacacsaction_ds" {
		name = citrixadc_authenticationtacacsaction.tf_tacacsaction.name
	}
`

func TestAccAuthenticationtacacsactionDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationtacacsactionDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_authenticationtacacsaction.tf_tacacsaction_ds", "name", "tf_tacacsaction_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationtacacsaction.tf_tacacsaction_ds", "serverip", "1.2.3.5"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationtacacsaction.tf_tacacsaction_ds", "serverport", "49"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationtacacsaction.tf_tacacsaction_ds", "authtimeout", "10"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationtacacsaction.tf_tacacsaction_ds", "authorization", "ON"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationtacacsaction.tf_tacacsaction_ds", "accounting", "ON"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationtacacsaction.tf_tacacsaction_ds", "auditfailedcmds", "ON"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationtacacsaction.tf_tacacsaction_ds", "groupattrname", "group"),
				),
			},
		},
	})
}
