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
	"strings"
	"testing"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testAccAuthenticationoauthaction_add = `

	resource "citrixadc_authenticationoauthaction" "tf_authenticationoauthaction" {
		name                  = "tf_authenticationoauthaction"
		authorizationendpoint = "https://example.com/"
		tokenendpoint         = "https://example.com/"
		clientid              = "id"
		clientsecret          = "secret"
		resourceuri           = "http://www.exampleadd.com"
		requestattribute   = "name=true@@@"
	}
`

const testAccAuthenticationoauthaction_update = `

	resource "citrixadc_authenticationoauthaction" "tf_authenticationoauthaction" {
		name                  = "tf_authenticationoauthaction"
		authorizationendpoint = "https://example.com/"
		tokenendpoint         = "https://example.com/"
		clientid              = "id"
		clientsecret          = "secret"
		resourceuri			  = "http://www.exampleupdate.com"
		requestattribute   = "name1=false@@@"
	}
`

func TestAccAuthenticationoauthaction_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationoauthactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationoauthaction_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationoauthactionExist("citrixadc_authenticationoauthaction.tf_authenticationoauthaction", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationoauthaction.tf_authenticationoauthaction", "name", "tf_authenticationoauthaction"),
					resource.TestCheckResourceAttr("citrixadc_authenticationoauthaction.tf_authenticationoauthaction", "resourceuri", "http://www.exampleadd.com"),
					resource.TestCheckResourceAttr("citrixadc_authenticationoauthaction.tf_authenticationoauthaction", "requestattribute", "name=true@@@"),
				),
			},
			{
				Config: testAccAuthenticationoauthaction_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationoauthactionExist("citrixadc_authenticationoauthaction.tf_authenticationoauthaction", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationoauthaction.tf_authenticationoauthaction", "name", "tf_authenticationoauthaction"),
					resource.TestCheckResourceAttr("citrixadc_authenticationoauthaction.tf_authenticationoauthaction", "resourceuri", "http://www.exampleupdate.com"),
					resource.TestCheckResourceAttr("citrixadc_authenticationoauthaction.tf_authenticationoauthaction", "requestattribute", "name1=false@@@"),
				),
			},
		},
	})
}

func TestAccAuthenticationoauthaction_import(t *testing.T) {
	const resAddr = "citrixadc_authenticationoauthaction.tf_authenticationoauthaction"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationoauthactionDestroy,
		Steps: []resource.TestStep{
			{Config: testAccAuthenticationoauthaction_add},
			{
				Config:            testAccAuthenticationoauthaction_add,
				ResourceName:      resAddr,
				ImportState:       true,
				ImportStateVerify: true,
				// clientsecret is a sensitive attribute NITRO never echoes back, and
				// clientsecret_wo_version is a client-side write-only version tracker not
				// returned by the API; neither can round-trip through import.
				ImportStateVerifyIgnore: []string{"clientsecret", "clientsecret_wo_version"},
			},
		},
	})
}

func testAccCheckAuthenticationoauthactionExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No authenticationoauthaction name is set")
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
		data, err := client.FindResource(service.Authenticationoauthaction.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("authenticationoauthaction %s not found", n)
		}

		return nil
	}
}

func testAccCheckAuthenticationoauthactionDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_authenticationoauthaction" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Authenticationoauthaction.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("authenticationoauthaction %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccAuthenticationoauthactionDataSource_basic = `

	resource "citrixadc_authenticationoauthaction" "tf_authenticationoauthaction_ds" {
		name                  = "tf_authenticationoauthaction_ds"
		authorizationendpoint = "https://example.com/auth"
		tokenendpoint         = "https://example.com/token"
		clientid              = "datasource_test_id"
		clientsecret          = "datasource_test_secret"
		resourceuri           = "http://www.datasourcetest.com"
	}

	data "citrixadc_authenticationoauthaction" "tf_authenticationoauthaction_ds" {
		name = citrixadc_authenticationoauthaction.tf_authenticationoauthaction_ds.name
	}
`

func TestAccAuthenticationoauthactionDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationoauthactionDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_authenticationoauthaction.tf_authenticationoauthaction_ds", "name", "tf_authenticationoauthaction_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationoauthaction.tf_authenticationoauthaction_ds", "authorizationendpoint", "https://example.com/auth"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationoauthaction.tf_authenticationoauthaction_ds", "tokenendpoint", "https://example.com/token"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationoauthaction.tf_authenticationoauthaction_ds", "clientid", "datasource_test_id"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationoauthaction.tf_authenticationoauthaction_ds", "resourceuri", "http://www.datasourcetest.com"),
				),
			},
		},
	})
}

// Ephemeral / Write-Only tests for clientsecret

const testAccAuthenticationoauthaction_clientsecret_step1 = `
	variable "authenticationoauthaction_clientsecret" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_authenticationoauthaction" "test" {
		name                  = "tf_oauthaction_compat"
		authorizationendpoint = "https://example.com/"
		tokenendpoint         = "https://example.com/"
		clientid              = "id"
		clientsecret          = var.authenticationoauthaction_clientsecret
		resourceuri           = "http://www.example.com"
	}
`

const testAccAuthenticationoauthaction_clientsecret_step2 = `
	variable "authenticationoauthaction_clientsecret_2" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_authenticationoauthaction" "test" {
		name                  = "tf_oauthaction_compat"
		authorizationendpoint = "https://example.com/"
		tokenendpoint         = "https://example.com/"
		clientid              = "id"
		clientsecret          = var.authenticationoauthaction_clientsecret_2
		resourceuri           = "http://www.example.com"
	}
`

func TestAccAuthenticationoauthaction_clientsecret_backward_compat(t *testing.T) {
	t.Setenv("TF_VAR_authenticationoauthaction_clientsecret", "secret1")
	t.Setenv("TF_VAR_authenticationoauthaction_clientsecret_2", "secret2")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationoauthactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationoauthaction_clientsecret_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationoauthactionExist("citrixadc_authenticationoauthaction.test", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationoauthaction.test", "name", "tf_oauthaction_compat"),
				),
			},
			{
				Config: testAccAuthenticationoauthaction_clientsecret_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationoauthactionExist("citrixadc_authenticationoauthaction.test", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationoauthaction.test", "name", "tf_oauthaction_compat"),
				),
			},
		},
	})
}

const testAccAuthenticationoauthaction_wo_step1 = `
	variable "authenticationoauthaction_clientsecret_wo" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_authenticationoauthaction" "test" {
		name                      = "tf_oauthaction_wo"
		authorizationendpoint     = "https://example.com/"
		tokenendpoint             = "https://example.com/"
		clientid                  = "id"
		clientsecret_wo           = var.authenticationoauthaction_clientsecret_wo
		clientsecret_wo_version   = 1
		resourceuri               = "http://www.example.com"
	}
`

const testAccAuthenticationoauthaction_wo_step2 = `
	variable "authenticationoauthaction_clientsecret_wo_2" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_authenticationoauthaction" "test" {
		name                      = "tf_oauthaction_wo"
		authorizationendpoint     = "https://example.com/"
		tokenendpoint             = "https://example.com/"
		clientid                  = "id"
		clientsecret_wo           = var.authenticationoauthaction_clientsecret_wo_2
		clientsecret_wo_version   = 2
		resourceuri               = "http://www.example.com"
	}
`

func TestAccAuthenticationoauthaction_clientsecret_wo_ephemeral(t *testing.T) {
	t.Setenv("TF_VAR_authenticationoauthaction_clientsecret_wo", "ephemeral_secret1")
	t.Setenv("TF_VAR_authenticationoauthaction_clientsecret_wo_2", "ephemeral_secret2")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationoauthactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationoauthaction_wo_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationoauthactionExist("citrixadc_authenticationoauthaction.test", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationoauthaction.test", "clientsecret_wo_version", "1"),
				),
			},
			{
				Config: testAccAuthenticationoauthaction_wo_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationoauthactionExist("citrixadc_authenticationoauthaction.test", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationoauthaction.test", "clientsecret_wo_version", "2"),
				),
			},
		},
	})
}

// Unset acceptance test: proves that unset-eligible attributes take effect when
// set (step 1) and revert to their NITRO defaults when removed from config (step 2),
// with an empty post-apply plan and direct appliance confirmation.
//
// tenantid is not unset-eligible; it is kept constant in both steps because it is a
// mandatory prerequisite for oauthtype=INTUNE (the only non-default oauthtype value).
const testAccAuthenticationoauthaction_unset_step1 = `

	resource "citrixadc_authenticationoauthaction" "tf_unset" {
		name                    = "tf_test_authenticationoauthaction_unset"
		authorizationendpoint   = "https://example.com/"
		tokenendpoint           = "https://example.com/"
		clientid                = "id"
		clientsecret            = "secret"
		tenantid                = "my-tenant-id"
		authentication          = "DISABLED"
		oauthtype               = "INTUNE"
		pkce                    = "DISABLED"
		refreshinterval         = 100
		skewtime                = 10
		tokenendpointauthmethod = "client_secret_jwt"
	}
`

const testAccAuthenticationoauthaction_unset_step2 = `

	resource "citrixadc_authenticationoauthaction" "tf_unset" {
		name                    = "tf_test_authenticationoauthaction_unset"
		authorizationendpoint   = "https://example.com/"
		tokenendpoint           = "https://example.com/"
		clientid                = "id"
		clientsecret            = "secret"
		# eligible attributes removed from config -> provider must unset them.
		# tenantid (a non-eligible prerequisite of oauthtype=INTUNE) is also removed
		# because it is only needed to support the INTUNE value in step 1.
	}
`

func TestAccAuthenticationoauthaction_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationoauthactionDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values apply and persist.
				Config: testAccAuthenticationoauthaction_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationoauthactionExist("citrixadc_authenticationoauthaction.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationoauthaction.tf_unset", "authentication", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_authenticationoauthaction.tf_unset", "oauthtype", "INTUNE"),
					resource.TestCheckResourceAttr("citrixadc_authenticationoauthaction.tf_unset", "pkce", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_authenticationoauthaction.tf_unset", "refreshinterval", "100"),
					resource.TestCheckResourceAttr("citrixadc_authenticationoauthaction.tf_unset", "skewtime", "10"),
					resource.TestCheckResourceAttr("citrixadc_authenticationoauthaction.tf_unset", "tokenendpointauthmethod", "client_secret_jwt"),
				),
			},
			{
				// Removing them must unset -> state reverts to NITRO defaults,
				// and the implicit post-apply plan must be empty.
				Config: testAccAuthenticationoauthaction_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationoauthactionExist("citrixadc_authenticationoauthaction.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationoauthaction.tf_unset", "authentication", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_authenticationoauthaction.tf_unset", "oauthtype", "GENERIC"),
					resource.TestCheckResourceAttr("citrixadc_authenticationoauthaction.tf_unset", "pkce", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_authenticationoauthaction.tf_unset", "refreshinterval", "1440"),
					resource.TestCheckResourceAttr("citrixadc_authenticationoauthaction.tf_unset", "skewtime", "5"),
					resource.TestCheckResourceAttr("citrixadc_authenticationoauthaction.tf_unset", "tokenendpointauthmethod", "client_secret_post"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckAuthenticationoauthactionADCValue("tf_test_authenticationoauthaction_unset", "authentication", "ENABLED"),
					testAccCheckAuthenticationoauthactionADCValue("tf_test_authenticationoauthaction_unset", "oauthtype", "GENERIC"),
					testAccCheckAuthenticationoauthactionADCValue("tf_test_authenticationoauthaction_unset", "pkce", "ENABLED"),
					testAccCheckAuthenticationoauthactionADCValue("tf_test_authenticationoauthaction_unset", "refreshinterval", "1440"),
					testAccCheckAuthenticationoauthactionADCValue("tf_test_authenticationoauthaction_unset", "skewtime", "5"),
					testAccCheckAuthenticationoauthactionADCValue("tf_test_authenticationoauthaction_unset", "tokenendpointauthmethod", "client_secret_post"),
				),
			},
		},
	})
}

// testAccCheckAuthenticationoauthactionADCValue asserts an attribute's value directly
// on the appliance (not just in Terraform state), proving the unset actually reverted it.
func testAccCheckAuthenticationoauthactionADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Authenticationoauthaction.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("authenticationoauthaction %s not found on appliance", name)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("authenticationoauthaction %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func TestAccAuthenticationoauthaction_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckAuthenticationoauthactionDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccAuthenticationoauthaction_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationoauthactionExist("citrixadc_authenticationoauthaction.tf_authenticationoauthaction", nil),
				),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccAuthenticationoauthaction_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationoauthactionExist("citrixadc_authenticationoauthaction.tf_authenticationoauthaction", nil),
				),
			},
		},
	})
}
