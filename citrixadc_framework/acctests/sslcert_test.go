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

const testAccSslcert_basic = `

	resource "citrixadc_sslcert" "tf_sslcert_ephem" {
		certfile = "/nsconfig/ssl/rootcert21.cert"
		reqfile  = "/nsconfig/ssl/rootcert2.req"
		certtype = "ROOT_CERT"
		keyfile  = "/nsconfig/ssl/rootcert2.key"
		keyform = "PEM"
	}
`

func TestAccSslcert_basic(t *testing.T) {
	t.Skip("TODO: Requires cleanup of certfile at ADC!")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doSslcertkeyPreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSslcert_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslcertExist("citrixadc_sslcert.tf_sslcert_ephem", nil),
				),
			},
		},
	})
}

func testAccCheckSslcertExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No sslcert name is set")
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

// Test backward-compatible path: using pempassphrase (Sensitive attribute)
const testAccSslcert_pempassphrase_step1 = `
	variable "sslcert_pempassphrase" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_sslcert" "tf_sslcert_ephem" {
		certfile = "/nsconfig/ssl/rootcert31.cert"
		reqfile  = "/nsconfig/ssl/rootcert3.req"
		certtype = "ROOT_CERT"
		keyfile  = "/nsconfig/ssl/rootcert3.key"
		keyform = "PEM"
		pempassphrase = var.sslcert_pempassphrase
	}
`

func TestAccSslcert_pempassphrase_backward_compat(t *testing.T) {
	t.Skip("TODO: Requires cleanup of certfile at ADC!")
	t.Setenv("TF_VAR_sslcert_pempassphrase", "1234567")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doSslcertkeyPreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSslcert_pempassphrase_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslcertExist("citrixadc_sslcert.tf_sslcert_ephem", nil),
				),
			},
		},
	})
}

// Test ephemeral path: using pempassphrase_wo (WriteOnly attribute) with version tracker
const testAccSslcert_pempassphrase_wo_step1 = `
	variable "sslcert_pempassphrase_wo" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_sslcert" "tf_sslcert_ephem" {
		certfile = "/nsconfig/ssl/rootcert311.cert"
		reqfile  = "/nsconfig/ssl/rootcert3.req"
		certtype = "ROOT_CERT"
		keyfile  = "/nsconfig/ssl/rootcert3.key"
		keyform = "PEM"
		pempassphrase_wo       = var.sslcert_pempassphrase_wo
		pempassphrase_wo_version = 1
	}
`

func TestAccSslcert_pempassphrase_wo_ephemeral(t *testing.T) {
	t.Skip("TODO: Requires cleanup of certfile at ADC!")
	t.Setenv("TF_VAR_sslcert_pempassphrase_wo", "1234567")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doSslcertkeyPreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSslcert_pempassphrase_wo_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslcertExist("citrixadc_sslcert.tf_sslcert_ephem", nil),
					resource.TestCheckResourceAttr("citrixadc_sslcert.tf_sslcert_ephem", "pempassphrase_wo_version", "1"),
				),
			},
		},
	})
}

func TestAccSslcert_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() { doSslcertkeyPreChecks(t) },
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccSslcert_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslcertExist("citrixadc_sslcert.tf_sslcert_ephem", nil),
				),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccSslcert_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslcertExist("citrixadc_sslcert.tf_sslcert_ephem", nil),
				),
			},
		},
	})
}
