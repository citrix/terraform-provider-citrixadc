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

const testAccSslhsmkey_basic = `
resource "citrixadc_sslhsmkey" "tf_hsmkey1" {
	hsmkeyname = "hsmkey1"
	hsmtype = "Fillme"
	password = "Fillme"
	serialnum = "Fillme"
}
`

const testAccSslhsmkeyDataSource_basic = `
resource "citrixadc_sslhsmkey" "tf_hsmkey1" {
	hsmkeyname = "hsmkey1"
	hsmtype = "Fillme"
	password = "Fillme"
	serialnum = "Fillme"
}

data "citrixadc_sslhsmkey" "tf_hsmkey1" {
	hsmkeyname = citrixadc_sslhsmkey.tf_hsmkey1.hsmkeyname
}
`

func TestAccSslhsmkey_basic(t *testing.T) {
	if adcTestbed != "STANDALONE_HSM" {
		t.Skipf("ADC testbed is %s. Expected STANDALONE_HSM.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslhsmkeyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslhsmkey_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslhsmkeyExist("citrixadc_sslhsmkey.tf_hsmkey1", nil),
				),
			},
		},
	})
}

func testAccCheckSslhsmkeyExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No sslhsmkey name is set")
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
		data, err := client.FindResource(service.Sslhsmkey.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("sslhsmkey %s not found", n)
		}

		return nil
	}
}

func testAccCheckSslhsmkeyDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_sslhsmkey" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Sslhsmkey.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("sslhsmkey %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccSslhsmkeyDataSource_basic(t *testing.T) {
	if adcTestbed != "STANDALONE_HSM" {
		t.Skipf("ADC testbed is %s. Expected STANDALONE_HSM.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSslhsmkeyDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_sslhsmkey.tf_hsmkey1", "hsmkeyname", "hsmkey1"),
					resource.TestCheckResourceAttr("data.citrixadc_sslhsmkey.tf_hsmkey1", "hsmtype", "Fillme"),
					resource.TestCheckResourceAttr("data.citrixadc_sslhsmkey.tf_hsmkey1", "serialnum", "Fillme"),
				),
			},
		},
	})
}

// Test backward-compatible path: using password (Sensitive attribute)
const testAccSslhsmkey_password_step1 = `
	variable "sslhsmkey_password" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_sslhsmkey" "tf_hsmkey_ephem" {
		hsmkeyname = "hsmkey_ephem"
		hsmtype    = "Fillme"
		serialnum  = "Fillme"
		password   = var.sslhsmkey_password
	}
`

const testAccSslhsmkey_password_step2 = `
	variable "sslhsmkey_password_2" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_sslhsmkey" "tf_hsmkey_ephem" {
		hsmkeyname = "hsmkey_ephem"
		hsmtype    = "Fillme"
		serialnum  = "Fillme"
		password   = var.sslhsmkey_password_2
	}
`

func TestAccSslhsmkey_password_backward_compat(t *testing.T) {
	if adcTestbed != "STANDALONE_HSM" {
		t.Skipf("ADC testbed is %s. Expected STANDALONE_HSM.", adcTestbed)
	}
	t.Setenv("TF_VAR_sslhsmkey_password", "hsmpass1")
	t.Setenv("TF_VAR_sslhsmkey_password_2", "hsmpass2")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslhsmkeyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslhsmkey_password_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslhsmkeyExist("citrixadc_sslhsmkey.tf_hsmkey_ephem", nil),
					resource.TestCheckResourceAttr("citrixadc_sslhsmkey.tf_hsmkey_ephem", "hsmkeyname", "hsmkey_ephem"),
				),
			},
			{
				Config: testAccSslhsmkey_password_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslhsmkeyExist("citrixadc_sslhsmkey.tf_hsmkey_ephem", nil),
					resource.TestCheckResourceAttr("citrixadc_sslhsmkey.tf_hsmkey_ephem", "hsmkeyname", "hsmkey_ephem"),
				),
			},
		},
	})
}

func TestAccSslhsmkey_import(t *testing.T) {
	if adcTestbed != "STANDALONE_HSM" {
		t.Skipf("ADC testbed is %s. Expected STANDALONE_HSM.", adcTestbed)
	}
	const resAddr = "citrixadc_sslhsmkey.tf_hsmkey1"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslhsmkeyDestroy,
		Steps: []resource.TestStep{
			{Config: testAccSslhsmkey_basic},
			{
				Config:                  testAccSslhsmkey_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccSslhsmkey_sdkv2StateUpgrade(t *testing.T) {
	if adcTestbed != "STANDALONE_HSM" {
		t.Skipf("ADC testbed is %s. Expected STANDALONE_HSM.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckSslhsmkeyDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccSslhsmkey_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslhsmkeyExist("citrixadc_sslhsmkey.tf_hsmkey1", nil),
				),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccSslhsmkey_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslhsmkeyExist("citrixadc_sslhsmkey.tf_hsmkey1", nil),
				),
			},
		},
	})
}

// Test ephemeral path: using password_wo (WriteOnly attribute) with version tracker
const testAccSslhsmkey_password_wo_step1 = `
	variable "sslhsmkey_password_wo" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_sslhsmkey" "tf_hsmkey_ephem" {
		hsmkeyname          = "hsmkey_ephem"
		hsmtype             = "Fillme"
		serialnum           = "Fillme"
		password_wo         = var.sslhsmkey_password_wo
		password_wo_version = 1
	}
`

const testAccSslhsmkey_password_wo_step2 = `
	variable "sslhsmkey_password_wo_2" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_sslhsmkey" "tf_hsmkey_ephem" {
		hsmkeyname          = "hsmkey_ephem"
		hsmtype             = "Fillme"
		serialnum           = "Fillme"
		password_wo         = var.sslhsmkey_password_wo_2
		password_wo_version = 2
	}
`

func TestAccSslhsmkey_password_wo_ephemeral(t *testing.T) {
	if adcTestbed != "STANDALONE_HSM" {
		t.Skipf("ADC testbed is %s. Expected STANDALONE_HSM.", adcTestbed)
	}
	t.Setenv("TF_VAR_sslhsmkey_password_wo", "ephem_hsmpass1")
	t.Setenv("TF_VAR_sslhsmkey_password_wo_2", "ephem_hsmpass2")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslhsmkeyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslhsmkey_password_wo_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslhsmkeyExist("citrixadc_sslhsmkey.tf_hsmkey_ephem", nil),
					resource.TestCheckResourceAttr("citrixadc_sslhsmkey.tf_hsmkey_ephem", "hsmkeyname", "hsmkey_ephem"),
					resource.TestCheckResourceAttr("citrixadc_sslhsmkey.tf_hsmkey_ephem", "password_wo_version", "1"),
				),
			},
			{
				Config: testAccSslhsmkey_password_wo_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslhsmkeyExist("citrixadc_sslhsmkey.tf_hsmkey_ephem", nil),
					resource.TestCheckResourceAttr("citrixadc_sslhsmkey.tf_hsmkey_ephem", "hsmkeyname", "hsmkey_ephem"),
					resource.TestCheckResourceAttr("citrixadc_sslhsmkey.tf_hsmkey_ephem", "password_wo_version", "2"),
				),
			},
		},
	})
}
