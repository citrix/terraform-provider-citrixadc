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

const testAccSslcrl_basic = `
	resource "citrixadc_sslcrl" "tf_sslcrl" {
		crlname = "tf_sslcrl"
		crlpath = "/var/netscaler/ssl/crl_config_clnt_rsa1_1cert.pem"
		cacert  = "rootrsa_cert1"
	}
`

const testAccSslcrlDataSource_basic = `
	resource "citrixadc_sslcrl" "tf_sslcrl" {
		crlname = "tf_sslcrl"
		crlpath = "/var/netscaler/ssl/crl_config_clnt_rsa1_1cert.pem"
		cacert  = "rootrsa_cert1"
	}

	data "citrixadc_sslcrl" "tf_sslcrl" {
		crlname = citrixadc_sslcrl.tf_sslcrl.crlname
	}
`

func TestAccSslcrl_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslcrlDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslcrl_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslcrlExist("citrixadc_sslcrl.tf_sslcrl", nil),
				),
			},
		},
	})
}

func testAccCheckSslcrlExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No sslcrl name is set")
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
		data, err := client.FindResource(service.Sslcrl.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("sslcrl %s not found", n)
		}

		return nil
	}
}

func testAccCheckSslcrlDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_sslcrl" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Sslcrl.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("sslcrl %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccSslcrlDataSource_basic(t *testing.T) {
	t.Skipf("Find  a way to upload a CRL file to the ADC instance before running this test")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSslcrlDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_sslcrl.tf_sslcrl", "crlname", "tf_sslcrl"),
					resource.TestCheckResourceAttr("data.citrixadc_sslcrl.tf_sslcrl", "crlpath", "/var/netscaler/ssl/crl_config_clnt_rsa1_1cert.pem"),
					resource.TestCheckResourceAttr("data.citrixadc_sslcrl.tf_sslcrl", "cacert", "rootrsa_cert1"),
				),
			},
		},
	})
}

// Test backward-compatible path: using password (Sensitive attribute)
// password is the LDAP password used when refreshing CRL from an LDAP server
const testAccSslcrl_password_step1 = `
	variable "sslcrl_password" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_sslcrl" "tf_sslcrl_ephem" {
		crlname  = "tf_sslcrl_ephem"
		crlpath  = "/var/netscaler/ssl/crl_config_clnt_rsa1_1cert.pem"
		cacert   = "rootrsa_cert1"
		password = var.sslcrl_password
	}
`

const testAccSslcrl_password_step2 = `
	variable "sslcrl_password_2" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_sslcrl" "tf_sslcrl_ephem" {
		crlname  = "tf_sslcrl_ephem"
		crlpath  = "/var/netscaler/ssl/crl_config_clnt_rsa1_1cert.pem"
		cacert   = "rootrsa_cert1"
		password = var.sslcrl_password_2
	}
`

func TestAccSslcrl_password_backward_compat(t *testing.T) {
	t.Skipf("Need a valid CRL file on the ADC instance before running this test")
	t.Setenv("TF_VAR_sslcrl_password", "crlldappass1")
	t.Setenv("TF_VAR_sslcrl_password_2", "crlldappass2")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslcrlDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslcrl_password_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslcrlExist("citrixadc_sslcrl.tf_sslcrl_ephem", nil),
					resource.TestCheckResourceAttr("citrixadc_sslcrl.tf_sslcrl_ephem", "crlname", "tf_sslcrl_ephem"),
				),
			},
			{
				Config: testAccSslcrl_password_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslcrlExist("citrixadc_sslcrl.tf_sslcrl_ephem", nil),
					resource.TestCheckResourceAttr("citrixadc_sslcrl.tf_sslcrl_ephem", "crlname", "tf_sslcrl_ephem"),
				),
			},
		},
	})
}

// Test ephemeral path: using password_wo (WriteOnly attribute) with version tracker
const testAccSslcrl_password_wo_step1 = `
	variable "sslcrl_password_wo" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_sslcrl" "tf_sslcrl_ephem" {
		crlname             = "tf_sslcrl_ephem"
		crlpath             = "/var/netscaler/ssl/crl_config_clnt_rsa1_1cert.pem"
		cacert              = "rootrsa_cert1"
		password_wo         = var.sslcrl_password_wo
		password_wo_version = 1
	}
`

const testAccSslcrl_password_wo_step2 = `
	variable "sslcrl_password_wo_2" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_sslcrl" "tf_sslcrl_ephem" {
		crlname             = "tf_sslcrl_ephem"
		crlpath             = "/var/netscaler/ssl/crl_config_clnt_rsa1_1cert.pem"
		cacert              = "rootrsa_cert1"
		password_wo         = var.sslcrl_password_wo_2
		password_wo_version = 2
	}
`

func TestAccSslcrl_password_wo_ephemeral(t *testing.T) {
	t.Skipf("Need a valid CRL file on the ADC instance before running this test")
	t.Setenv("TF_VAR_sslcrl_password_wo", "ephem_crlpass1")
	t.Setenv("TF_VAR_sslcrl_password_wo_2", "ephem_crlpass2")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslcrlDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslcrl_password_wo_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslcrlExist("citrixadc_sslcrl.tf_sslcrl_ephem", nil),
					resource.TestCheckResourceAttr("citrixadc_sslcrl.tf_sslcrl_ephem", "crlname", "tf_sslcrl_ephem"),
					resource.TestCheckResourceAttr("citrixadc_sslcrl.tf_sslcrl_ephem", "password_wo_version", "1"),
				),
			},
			{
				Config: testAccSslcrl_password_wo_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslcrlExist("citrixadc_sslcrl.tf_sslcrl_ephem", nil),
					resource.TestCheckResourceAttr("citrixadc_sslcrl.tf_sslcrl_ephem", "crlname", "tf_sslcrl_ephem"),
					resource.TestCheckResourceAttr("citrixadc_sslcrl.tf_sslcrl_ephem", "password_wo_version", "2"),
				),
			},
		},
	})
}
