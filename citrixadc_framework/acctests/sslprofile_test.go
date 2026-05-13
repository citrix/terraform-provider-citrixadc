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

const testAccSslprofile_add = `
	resource "citrixadc_sslprofile" "foo" {
		name = "tfAcc_sslprofile"
		ecccurvebindings = []
	}
`
const testAccSslprofile_update = `
	resource "citrixadc_sslprofile" "foo" {
		name = "tfAcc_sslprofile"
		hsts = "ENABLED"
		snienable = "ENABLED"
		ecccurvebindings = []
		sslclientlogs = "ENABLED"
		encryptedclienthello = "ENABLED"
		defaultsni = 60
		allowunknownsni = "ENABLED"
		allowextendedmastersecret = "YES"
	}
`

func TestAccSslprofile_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslprofile_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslprofileExist("citrixadc_sslprofile.foo", nil),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.foo", "name", "tfAcc_sslprofile"),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.foo", "sslclientlogs", "DISABLED"),
				),
			},
			{
				Config: testAccSslprofile_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslprofileExist("citrixadc_sslprofile.foo", nil),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.foo", "name", "tfAcc_sslprofile"),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.foo", "hsts", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.foo", "sslclientlogs", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.foo", "encryptedclienthello", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.foo", "defaultsni", "60"),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.foo", "allowunknownsni", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.foo", "allowextendedmastersecret", "YES"),
				),
			},
		},
	})
}

const testAccSslprofile_ecccurvebinding_bind = `
	resource "citrixadc_sslprofile" "foo" {
		name = "tfAcc_sslprofile"
		ecccurvebindings = ["P_256"]
	}
`
const testAccSslprofile_ecccurvebinding_unbind = `
	resource "citrixadc_sslprofile" "foo" {
		name = "tfAcc_sslprofile"
		ecccurvebindings = []
	}
`

func TestAccSslprofile_ecccurve_binding(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslprofile_ecccurvebinding_bind,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslprofileExist("citrixadc_sslprofile.foo", nil),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.foo", "name", "tfAcc_sslprofile"),
				),
			},
			{
				Config: testAccSslprofile_ecccurvebinding_unbind,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslprofileExist("citrixadc_sslprofile.foo", nil),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.foo", "name", "tfAcc_sslprofile"),
				),
			},
		},
	})
}

const testAccSslprofile_cipherbinding_bind = `
	resource "citrixadc_sslprofile" "foo" {
		name = "tfAcc_sslprofile"
		ecccurvebindings = []
		cipherbindings {
			ciphername     = "HIGH"
			cipherpriority = 10
	}
	}
`
const testAccSslprofile_cipherbinding_unbind = `
	resource "citrixadc_sslprofile" "foo" {
		name = "tfAcc_sslprofile"
		ecccurvebindings = []
	}
`

func TestAccSslprofile_cipher_binding(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslprofile_cipherbinding_bind,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslprofileExist("citrixadc_sslprofile.foo", nil),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.foo", "name", "tfAcc_sslprofile"),
				),
			},
			{
				Config: testAccSslprofile_cipherbinding_unbind,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslprofileExist("citrixadc_sslprofile.foo", nil),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.foo", "name", "tfAcc_sslprofile"),
				),
			},
		},
	})
}

func testAccCheckSslprofileExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SSL Profile name is set")
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
		data, err := client.FindResource(service.Sslprofile.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("SSL Profile %s not found", n)
		}

		return nil
	}
}

func testAccCheckSslprofileDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_sslprofile" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Sslprofile.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("SSL Profile %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccSslprofileDataSource_basic = `
	resource "citrixadc_sslprofile" "tf_sslprofile" {
		name = "tf_sslprofile_datasource"
		hsts = "ENABLED"
		snienable = "ENABLED"
		ecccurvebindings = []
		sslclientlogs = "ENABLED"
		encryptedclienthello = "ENABLED"
		defaultsni = "60"
		allowunknownsni = "ENABLED"
		allowextendedmastersecret = "YES"
	}

	data "citrixadc_sslprofile" "tf_sslprofile_datasource" {
		name = citrixadc_sslprofile.tf_sslprofile.name
	}
`

func TestAccSslprofileDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSslprofileDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_sslprofile.tf_sslprofile_datasource", "name", "tf_sslprofile_datasource"),
					resource.TestCheckResourceAttr("data.citrixadc_sslprofile.tf_sslprofile_datasource", "hsts", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_sslprofile.tf_sslprofile_datasource", "snienable", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_sslprofile.tf_sslprofile_datasource", "sslclientlogs", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_sslprofile.tf_sslprofile_datasource", "encryptedclienthello", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_sslprofile.tf_sslprofile_datasource", "defaultsni", "60"),
					resource.TestCheckResourceAttr("data.citrixadc_sslprofile.tf_sslprofile_datasource", "allowunknownsni", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_sslprofile.tf_sslprofile_datasource", "allowextendedmastersecret", "YES"),
				),
			},
		},
	})
}

// Test backward-compatible path: using sessionticketkeydata (Sensitive attribute)
const testAccSslprofile_sessionticketkeydata_step1 = `
	variable "sslprofile_sessionticketkeydata" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_sslprofile" "tf_sslprofile_ephem" {
		name                   = "tf_sslprofile_ephem"
		sessionticket          = "ENABLED"
		sessionticketkeydata   = var.sslprofile_sessionticketkeydata
	}
`

const testAccSslprofile_sessionticketkeydata_step2 = `
	variable "sslprofile_sessionticketkeydata_2" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_sslprofile" "tf_sslprofile_ephem" {
		name                   = "tf_sslprofile_ephem"
		sessionticket          = "ENABLED"
		sessionticketkeydata   = var.sslprofile_sessionticketkeydata_2
	}
`

func TestAccSslprofile_sessionticketkeydata_backward_compat(t *testing.T) {
	t.Setenv("TF_VAR_sslprofile_sessionticketkeydata", "0102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f20")
	t.Setenv("TF_VAR_sslprofile_sessionticketkeydata_2", "2122232425262728292a2b2c2d2e2f303132333435363738393a3b3c3d3e3f40")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslprofile_sessionticketkeydata_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslprofileExist("citrixadc_sslprofile.tf_sslprofile_ephem", nil),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.tf_sslprofile_ephem", "name", "tf_sslprofile_ephem"),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.tf_sslprofile_ephem", "sessionticket", "ENABLED"),
				),
			},
			{
				Config: testAccSslprofile_sessionticketkeydata_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslprofileExist("citrixadc_sslprofile.tf_sslprofile_ephem", nil),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.tf_sslprofile_ephem", "name", "tf_sslprofile_ephem"),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.tf_sslprofile_ephem", "sessionticket", "ENABLED"),
				),
			},
		},
	})
}

// Test ephemeral path: using sessionticketkeydata_wo (WriteOnly attribute) with version tracker
const testAccSslprofile_sessionticketkeydata_wo_step1 = `
	variable "sslprofile_sessionticketkeydata_wo" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_sslprofile" "tf_sslprofile_ephem" {
		name                              = "tf_sslprofile_ephem"
		sessionticket                     = "ENABLED"
		sessionticketkeydata_wo           = var.sslprofile_sessionticketkeydata_wo
		sessionticketkeydata_wo_version   = 1
	}
`

const testAccSslprofile_sessionticketkeydata_wo_step2 = `
	variable "sslprofile_sessionticketkeydata_wo_2" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_sslprofile" "tf_sslprofile_ephem" {
		name                              = "tf_sslprofile_ephem"
		sessionticket                     = "ENABLED"
		sessionticketkeydata_wo           = var.sslprofile_sessionticketkeydata_wo_2
		sessionticketkeydata_wo_version   = 2
	}
`

func TestAccSslprofile_sessionticketkeydata_wo_ephemeral(t *testing.T) {
	t.Setenv("TF_VAR_sslprofile_sessionticketkeydata_wo", "0102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f20")
	t.Setenv("TF_VAR_sslprofile_sessionticketkeydata_wo_2", "2122232425262728292a2b2c2d2e2f303132333435363738393a3b3c3d3e3f40")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslprofile_sessionticketkeydata_wo_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslprofileExist("citrixadc_sslprofile.tf_sslprofile_ephem", nil),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.tf_sslprofile_ephem", "name", "tf_sslprofile_ephem"),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.tf_sslprofile_ephem", "sessionticketkeydata_wo_version", "1"),
				),
			},
			{
				Config: testAccSslprofile_sessionticketkeydata_wo_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslprofileExist("citrixadc_sslprofile.tf_sslprofile_ephem", nil),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.tf_sslprofile_ephem", "name", "tf_sslprofile_ephem"),
					resource.TestCheckResourceAttr("citrixadc_sslprofile.tf_sslprofile_ephem", "sessionticketkeydata_wo_version", "2"),
				),
			},
		},
	})
}
