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

const testAccSslrsakey_basic = `

	resource "citrixadc_sslrsakey" "tf_sslrsakey" {
		keyfile  = "/nsconfig/ssl/demo.pem"
		bits     = 2048
		aes256   = true
		password = "SecretPassword"
	}
`

func TestAccSslrsakey_basic(t *testing.T) {
	t.Skip("TODO: Requires cleanup of keyfile at ADC!")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSslrsakey_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslrsakeyExist("citrixadc_sslrsakey.tf_sslrsakey", nil),
				),
			},
		},
	})
}

func testAccCheckSslrsakeyExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No sslrsakey name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}
		return nil
	}
}

// Test backward-compatible path: using password (Sensitive attribute)
const testAccSslrsakey_password_step1 = `
	variable "sslrsakey_password" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_sslrsakey" "tf_sslrsakey_ephem" {
		keyfile  = "/nsconfig/ssl/demorsaephem.pem"
		bits     = 2048
		aes256   = true
		password = var.sslrsakey_password
	}
`

func TestAccSslrsakey_password_backward_compat(t *testing.T) {
	t.Skip("TODO: Requires cleanup of keyfile at ADC!")
	t.Setenv("TF_VAR_sslrsakey_password", "RsaPass1!")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSslrsakey_password_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslrsakeyExist("citrixadc_sslrsakey.tf_sslrsakey_ephem", nil),
				),
			},
		},
	})
}

// Test ephemeral path: using password_wo (WriteOnly attribute) with version tracker
const testAccSslrsakey_password_wo_step1 = `
	variable "sslrsakey_password_wo" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_sslrsakey" "tf_sslrsakey_ephem" {
		keyfile             = "/nsconfig/ssl/demorsaephem_wo.pem"
		bits                = 2048
		aes256              = true
		password_wo         = var.sslrsakey_password_wo
		password_wo_version = 1
	}
`

func TestAccSslrsakey_password_wo_ephemeral(t *testing.T) {
	t.Skip("TODO: Requires cleanup of keyfile at ADC!")
	t.Setenv("TF_VAR_sslrsakey_password_wo", "EpheRsaPass1!")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSslrsakey_password_wo_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslrsakeyExist("citrixadc_sslrsakey.tf_sslrsakey_ephem", nil),
					resource.TestCheckResourceAttr("citrixadc_sslrsakey.tf_sslrsakey_ephem", "password_wo_version", "1"),
				),
			},
		},
	})
}

func TestAccSslrsakey_sdkv2StateUpgrade(t *testing.T) {
	t.Skip("TODO: Requires cleanup of keyfile at ADC!")
	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccSslrsakey_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslrsakeyExist("citrixadc_sslrsakey.tf_sslrsakey", nil),
				),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccSslrsakey_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslrsakeyExist("citrixadc_sslrsakey.tf_sslrsakey", nil),
				),
			},
		},
	})
}
