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

const testAccSslcertreq_basic = `

	resource "citrixadc_sslcertreq" "tf_sslcertreq" {
		reqfile          = "/nsconfig/ssl/rootcert21.req"
		keyfile          = "/nsconfig/ssl/rootcert2.key"
		countryname      = "in"
		statename        = "kar"
		organizationname = "xyz"
	}
`

func TestAccSslcertreq_basic(t *testing.T) {
	t.Skip("TODO: Requires cleanup of reqfile at ADC!")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doSslcertkeyPreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSslcertreq_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslcertreqExist("citrixadc_sslcertreq.tf_sslcertreq", nil),
				),
			},
		},
	})
}

func testAccCheckSslcertreqExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No sslcertreq name is set")
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

// Test backward-compatible path: using challengepassword (Sensitive attribute)
const testAccSslcertreq_challengepassword_step1 = `
	variable "sslcertreq_challengepassword" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_sslcertreq" "tf_sslcertreq_cp" {
		reqfile          = "/nsconfig/ssl/rootcert211.req"
		keyfile          = "/nsconfig/ssl/rootcert2.key"
		countryname      = "in"
		statename        = "kar"
		organizationname = "xyz"
		challengepassword = var.sslcertreq_challengepassword
	}
`

func TestAccSslcertreq_challengepassword_backward_compat(t *testing.T) {
	t.Skip("TODO: Requires cleanup of reqfile at ADC!")
	t.Setenv("TF_VAR_sslcertreq_challengepassword", "chalpass1")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSslcertreq_challengepassword_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslcertreqExist("citrixadc_sslcertreq.tf_sslcertreq_cp", nil),
				),
			},
		},
	})
}

// Test ephemeral path: using challengepassword_wo (WriteOnly attribute) with version tracker
const testAccSslcertreq_challengepassword_wo_step1 = `
	variable "sslcertreq_challengepassword_wo" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_sslcertreq" "tf_sslcertreq_cpwo" {
		reqfile          = "/nsconfig/ssl/rootcert2111.req"
		keyfile          = "/nsconfig/ssl/rootcert2.key"
		countryname      = "in"
		statename        = "kar"
		organizationname = "xyz"
		challengepassword_wo       = var.sslcertreq_challengepassword_wo
		challengepassword_wo_version = 1
	}
`

func TestAccSslcertreq_challengepassword_wo_ephemeral(t *testing.T) {
	t.Skip("TODO: Requires cleanup of reqfile at ADC!")
	t.Setenv("TF_VAR_sslcertreq_challengepassword_wo", "ephem_chalpass1")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSslcertreq_challengepassword_wo_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslcertreqExist("citrixadc_sslcertreq.tf_sslcertreq_cpwo", nil),
					resource.TestCheckResourceAttr("citrixadc_sslcertreq.tf_sslcertreq_cpwo", "challengepassword_wo_version", "1"),
				),
			},
		},
	})
}

// Test backward-compatible path: using pempassphrase (Sensitive attribute)
const testAccSslcertreq_pempassphrase_step1 = `
	variable "sslcertreq_pempassphrase" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_sslcertreq" "tf_sslcertreq_pp" {
		reqfile          = "/nsconfig/ssl/rootcert31.req"
		keyfile          = "/nsconfig/ssl/rootcert3.key"
		countryname      = "in"
		statename        = "kar"
		organizationname = "xyz"
		pempassphrase = var.sslcertreq_pempassphrase
		keyform = "PEM"
	}
`

func TestAccSslcertreq_pempassphrase_backward_compat(t *testing.T) {
	t.Skip("TODO: Requires cleanup of reqfile at ADC!")
	t.Setenv("TF_VAR_sslcertreq_pempassphrase", "1234567")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSslcertreq_pempassphrase_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslcertreqExist("citrixadc_sslcertreq.tf_sslcertreq_pp", nil),
				),
			},
		},
	})
}

// Test ephemeral path: using pempassphrase_wo (WriteOnly attribute) with version tracker
const testAccSslcertreq_pempassphrase_wo_step1 = `
	variable "sslcertreq_pempassphrase_wo" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_sslcertreq" "tf_sslcertreq_ppwo" {
		reqfile          = "/nsconfig/ssl/rootcert311.req"
		keyfile          = "/nsconfig/ssl/rootcert3.key"
		countryname      = "in"
		statename        = "kar"
		organizationname = "xyz"
		pempassphrase_wo       = var.sslcertreq_pempassphrase_wo
		pempassphrase_wo_version = 1
		keyform = "PEM"
	}
`

func TestAccSslcertreq_pempassphrase_wo_ephemeral(t *testing.T) {
	t.Skip("TODO: Requires cleanup of reqfile at ADC!")
	t.Setenv("TF_VAR_sslcertreq_pempassphrase_wo", "1234567")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSslcertreq_pempassphrase_wo_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslcertreqExist("citrixadc_sslcertreq.tf_sslcertreq_ppwo", nil),
					resource.TestCheckResourceAttr("citrixadc_sslcertreq.tf_sslcertreq_ppwo", "pempassphrase_wo_version", "1"),
				),
			},
		},
	})
}

func TestAccSslcertreq_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccSslcertreq_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslcertreqExist("citrixadc_sslcertreq.tf_sslcertreq", nil),
				),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccSslcertreq_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslcertreqExist("citrixadc_sslcertreq.tf_sslcertreq", nil),
				),
			},
		},
	})
}
