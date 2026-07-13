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

const testAccDnskey_add = `

resource "citrixadc_dnskey" "dnskey" {
	keyname            = "adckey_1"
	publickey           = "/nsconfig/dns/dnskey_test.key"
	privatekey          = "/nsconfig/dns/dnskey_test.private"
	expires            = 120
	units1             = "DAYS"
	notificationperiod = 7
	units2             = "DAYS"
	ttl                = 3600
	rollovermethod	 = "PrePublication"
	autorollover	 = "ENABLED"
	}
`
const testAccDnskey_update = `

resource "citrixadc_dnskey" "dnskey" {
	keyname            = "adckey_1"
	publickey           = "/nsconfig/dns/dnskey_test.key"
	privatekey          = "/nsconfig/dns/dnskey_test.private"
	expires            = 121
	units1             = "HOURS"
	notificationperiod = 12
	units2             = "HOURS"
	ttl                = 3601
	rollovermethod	 = "DoubleSignature"
	autorollover	 = "DISABLED"
	}
`

const testAccDnskeyDataSource_basic = `

resource "citrixadc_dnskey" "dnskey" {
	keyname            = "adckey_ds_test"
	publickey  = "/nsconfig/dns/dnskey_test.key"
	privatekey = "/nsconfig/dns/dnskey_test.private"
	expires            = 120
	units1             = "DAYS"
	notificationperiod = 7
	units2             = "DAYS"
	ttl                = 3600
}

data "citrixadc_dnskey" "dnskey" {
	keyname = citrixadc_dnskey.dnskey.keyname
}
`

func TestAccDnskey_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doDnskeyPreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDnskeyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDnskey_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDnskeyExist("citrixadc_dnskey.dnskey", nil),
					resource.TestCheckResourceAttr("citrixadc_dnskey.dnskey", "keyname", "adckey_1"),
					resource.TestCheckResourceAttr("citrixadc_dnskey.dnskey", "expires", "120"),
					resource.TestCheckResourceAttr("citrixadc_dnskey.dnskey", "units1", "DAYS"),
					resource.TestCheckResourceAttr("citrixadc_dnskey.dnskey", "notificationperiod", "7"),
					resource.TestCheckResourceAttr("citrixadc_dnskey.dnskey", "units2", "DAYS"),
					resource.TestCheckResourceAttr("citrixadc_dnskey.dnskey", "ttl", "3600"),
					resource.TestCheckResourceAttr("citrixadc_dnskey.dnskey", "rollovermethod", "PrePublication"),
					resource.TestCheckResourceAttr("citrixadc_dnskey.dnskey", "autorollover", "ENABLED"),
				),
			},
			{
				Config: testAccDnskey_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDnskeyExist("citrixadc_dnskey.dnskey", nil),
					resource.TestCheckResourceAttr("citrixadc_dnskey.dnskey", "keyname", "adckey_1"),
					resource.TestCheckResourceAttr("citrixadc_dnskey.dnskey", "expires", "121"),
					resource.TestCheckResourceAttr("citrixadc_dnskey.dnskey", "units1", "HOURS"),
					resource.TestCheckResourceAttr("citrixadc_dnskey.dnskey", "notificationperiod", "12"),
					resource.TestCheckResourceAttr("citrixadc_dnskey.dnskey", "units2", "HOURS"),
					resource.TestCheckResourceAttr("citrixadc_dnskey.dnskey", "ttl", "3601"),
					resource.TestCheckResourceAttr("citrixadc_dnskey.dnskey", "rollovermethod", "DoubleSignature"),
					resource.TestCheckResourceAttr("citrixadc_dnskey.dnskey", "autorollover", "DISABLED"),
				),
			},
		},
	})
}

func testAccCheckDnskeyExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No dnskey name is set")
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
		data, err := client.FindResource(service.Dnskey.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("dnskey %s not found", n)
		}

		return nil
	}
}

func testAccCheckDnskeyDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_dnskey" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Dnskey.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("dnskey %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

// Test backward-compatible path: using password (Sensitive attribute)
const testAccDnskey_password_step1 = `

	variable "dnskey_password" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_dnskey" "tf_dnskey" {
		keyname    = "tf_dnskey_test"
		publickey  = "/nsconfig/dns/dnskey_test.key"
		privatekey = "/nsconfig/dns/dnskey_test.private"
		password   = var.dnskey_password
	}
`

// Update backward-compatible path: change password value
const testAccDnskey_password_step2 = `

	variable "dnskey_password_2" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_dnskey" "tf_dnskey" {
		keyname    = "tf_dnskey_test"
		publickey  = "/nsconfig/dns/dnskey_test.key"
		privatekey = "/nsconfig/dns/dnskey_test.private"
		password   = var.dnskey_password_2
	}
`

func TestAccDnskey_password_backward_compat(t *testing.T) {
	t.Setenv("TF_VAR_dnskey_password", "oldpassword123")
	t.Setenv("TF_VAR_dnskey_password_2", "newpassword456")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doDnskeyPreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDnskeyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDnskey_password_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDnskeyExist("citrixadc_dnskey.tf_dnskey", nil),
					resource.TestCheckResourceAttr("citrixadc_dnskey.tf_dnskey", "keyname", "tf_dnskey_test"),
				),
			},
			{
				Config: testAccDnskey_password_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDnskeyExist("citrixadc_dnskey.tf_dnskey", nil),
					resource.TestCheckResourceAttr("citrixadc_dnskey.tf_dnskey", "keyname", "tf_dnskey_test"),
				),
			},
		},
	})
}

// Test ephemeral path: using password_wo (WriteOnly attribute) with version tracker
const testAccDnskey_password_wo_step1 = `

	variable "dnskey_password_wo" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_dnskey" "tf_dnskey" {
		keyname             = "tf_dnskey_test"
		publickey           = "/nsconfig/dns/dnskey_test.key"
		privatekey          = "/nsconfig/dns/dnskey_test.private"
		password_wo         = var.dnskey_password_wo
		password_wo_version = 1
	}
`

// Update ephemeral path: bump version to trigger update with new password
const testAccDnskey_password_wo_step2 = `

	variable "dnskey_password_wo_2" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_dnskey" "tf_dnskey" {
		keyname             = "tf_dnskey_test"
		publickey           = "/nsconfig/dns/dnskey_test.key"
		privatekey          = "/nsconfig/dns/dnskey_test.private"
		password_wo         = var.dnskey_password_wo_2
		password_wo_version = 2
	}
`

func TestAccDnskey_password_wo_ephemeral(t *testing.T) {
	t.Setenv("TF_VAR_dnskey_password_wo", "ephemeral_pass1")
	t.Setenv("TF_VAR_dnskey_password_wo_2", "ephemeral_pass2")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doDnskeyPreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDnskeyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDnskey_password_wo_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDnskeyExist("citrixadc_dnskey.tf_dnskey", nil),
					resource.TestCheckResourceAttr("citrixadc_dnskey.tf_dnskey", "keyname", "tf_dnskey_test"),
					resource.TestCheckResourceAttr("citrixadc_dnskey.tf_dnskey", "password_wo_version", "1"),
				),
			},
			{
				Config: testAccDnskey_password_wo_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDnskeyExist("citrixadc_dnskey.tf_dnskey", nil),
					resource.TestCheckResourceAttr("citrixadc_dnskey.tf_dnskey", "keyname", "tf_dnskey_test"),
					resource.TestCheckResourceAttr("citrixadc_dnskey.tf_dnskey", "password_wo_version", "2"),
				),
			},
		},
	})
}

func TestAccDnskeyDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doDnskeyPreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDnskeyDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_dnskey.dnskey", "keyname", "adckey_ds_test"),
					resource.TestCheckResourceAttr("data.citrixadc_dnskey.dnskey", "expires", "120"),
					resource.TestCheckResourceAttr("data.citrixadc_dnskey.dnskey", "units1", "DAYS"),
					resource.TestCheckResourceAttr("data.citrixadc_dnskey.dnskey", "notificationperiod", "7"),
					resource.TestCheckResourceAttr("data.citrixadc_dnskey.dnskey", "units2", "DAYS"),
					resource.TestCheckResourceAttr("data.citrixadc_dnskey.dnskey", "ttl", "3600"),
				),
			},
		},
	})
}

func TestAccDnskey_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckDnskeyDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccDnskey_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDnskeyExist("citrixadc_dnskey.dnskey", nil),
				),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccDnskey_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDnskeyExist("citrixadc_dnskey.dnskey", nil),
				),
			},
		},
	})
}
