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

const testAccNsencryptionkey_add = `
	resource "citrixadc_nsencryptionkey" "tf_encryptionkey" {
		name     = "tf_encryptionkey"
		method   = "AES256"
		keyvalue = "26ea5537b7e0746089476e5658f9327c0b10c3b4778c673a5b38cee182874711"
		padding  = "ON"
		iv       = "c2bf0b2e15c15004d6b14bcdc7e5e365"
		comment  = "Testing"
	}
`

const testAccNsencryptionkey_update = `
	resource "citrixadc_nsencryptionkey" "tf_encryptionkey" {
		name     = "tf_encryptionkey"
		method   = "AES256"
		keyvalue = "26ea5537b7e0746089476e5658f9327c0b10c3b4778c673a5b38cee182874711"
		padding  = "DEFAULT"
		iv       = "c2bf0b2e15c15004d6b14bcdc7e5e123"
		comment  = "Testing_sample"
	}
`

func TestAccNsencryptionkey_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNsencryptionkeyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNsencryptionkey_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsencryptionkeyExist("citrixadc_nsencryptionkey.tf_encryptionkey", nil),
					resource.TestCheckResourceAttr("citrixadc_nsencryptionkey.tf_encryptionkey", "name", "tf_encryptionkey"),
					resource.TestCheckResourceAttr("citrixadc_nsencryptionkey.tf_encryptionkey", "padding", "ON"),
					resource.TestCheckResourceAttr("citrixadc_nsencryptionkey.tf_encryptionkey", "iv", "c2bf0b2e15c15004d6b14bcdc7e5e365"),
					resource.TestCheckResourceAttr("citrixadc_nsencryptionkey.tf_encryptionkey", "comment", "Testing"),
				),
			},
			{
				Config: testAccNsencryptionkey_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsencryptionkeyExist("citrixadc_nsencryptionkey.tf_encryptionkey", nil),
					resource.TestCheckResourceAttr("citrixadc_nsencryptionkey.tf_encryptionkey", "name", "tf_encryptionkey"),
					resource.TestCheckResourceAttr("citrixadc_nsencryptionkey.tf_encryptionkey", "padding", "DEFAULT"),
					resource.TestCheckResourceAttr("citrixadc_nsencryptionkey.tf_encryptionkey", "iv", "c2bf0b2e15c15004d6b14bcdc7e5e123"),
					resource.TestCheckResourceAttr("citrixadc_nsencryptionkey.tf_encryptionkey", "comment", "Testing_sample"),
				),
			},
		},
	})
}

func TestAccNsencryptionkey_import(t *testing.T) {
	const resAddr = "citrixadc_nsencryptionkey.tf_encryptionkey"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNsencryptionkeyDestroy,
		Steps: []resource.TestStep{
			{Config: testAccNsencryptionkey_add},
			{
				Config:            testAccNsencryptionkey_add,
				ResourceName:      resAddr,
				ImportState:       true,
				ImportStateVerify: true,
				// keyvalue is a Sensitive attribute that NITRO never echoes back, and
				// keyvalue_wo_version is a state-only version tracker (also not returned
				// by NITRO) - neither can round-trip through import.
				ImportStateVerifyIgnore: []string{"keyvalue", "keyvalue_wo_version"},
			},
		},
	})
}

func testAccCheckNsencryptionkeyExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No nsencryptionkey name is set")
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
		data, err := client.FindResource("nsencryptionkey", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("nsencryptionkey %s not found", n)
		}

		return nil
	}
}

func testAccCheckNsencryptionkeyDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_nsencryptionkey" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("nsencryptionkey", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("nsencryptionkey %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccNsencryptionkeyDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccNsencryptionkeyDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_nsencryptionkey.tf_encryptionkey_ds", "name", "tf_encryptionkey_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_nsencryptionkey.tf_encryptionkey_ds", "method", "AES256"),
					resource.TestCheckResourceAttr("data.citrixadc_nsencryptionkey.tf_encryptionkey_ds", "padding", "ON"),
					resource.TestCheckResourceAttr("data.citrixadc_nsencryptionkey.tf_encryptionkey_ds", "iv", "c2bf0b2e15c15004d6b14bcdc7e5e365"),
					resource.TestCheckResourceAttr("data.citrixadc_nsencryptionkey.tf_encryptionkey_ds", "comment", "DataSource test for nsencryptionkey"),
				),
			},
		},
	})
}

const testAccNsencryptionkeyDataSource_basic = `

	resource "citrixadc_nsencryptionkey" "tf_encryptionkey_ds" {
		name     = "tf_encryptionkey_ds"
		method   = "AES256"
		keyvalue = "26ea5537b7e0746089476e5658f9327c0b10c3b4778c673a5b38cee182874711"
		padding  = "ON"
		iv       = "c2bf0b2e15c15004d6b14bcdc7e5e365"
		comment  = "DataSource test for nsencryptionkey"
	}

	data "citrixadc_nsencryptionkey" "tf_encryptionkey_ds" {
		name       = citrixadc_nsencryptionkey.tf_encryptionkey_ds.name
		depends_on = [citrixadc_nsencryptionkey.tf_encryptionkey_ds]
	}
`

// --- Ephemeral / Write-Only Tests for keyvalue ---

const testAccNsencryptionkey_keyvalue_step1 = `
	variable "nsencryptionkey_keyvalue" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_nsencryptionkey" "tf_encryptionkey" {
		name     = "tf_encryptionkey"
		method   = "AES256"
		keyvalue = var.nsencryptionkey_keyvalue
		padding  = "ON"
		iv       = "c2bf0b2e15c15004d6b14bcdc7e5e365"
		comment  = "Backward compat test step1"
	}
`

const testAccNsencryptionkey_keyvalue_step2 = `
	variable "nsencryptionkey_keyvalue_2" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_nsencryptionkey" "tf_encryptionkey" {
		name     = "tf_encryptionkey"
		method   = "AES256"
		keyvalue = var.nsencryptionkey_keyvalue_2
		padding  = "ON"
		iv       = "c2bf0b2e15c15004d6b14bcdc7e5e365"
		comment  = "Backward compat test step2"
	}
`

func TestAccNsencryptionkey_keyvalue_backward_compat(t *testing.T) {
	t.Setenv("TF_VAR_nsencryptionkey_keyvalue", "26ea5537b7e0746089476e5658f9327c0b10c3b4778c673a5b38cee182874711")
	t.Setenv("TF_VAR_nsencryptionkey_keyvalue_2", "8b1d2e3f4a5b6c7d8e9f0a1b2c3d4e5f6a7b8c9d0e1f2a3b4c5d6e7f8a9b0c1d")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNsencryptionkeyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNsencryptionkey_keyvalue_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsencryptionkeyExist("citrixadc_nsencryptionkey.tf_encryptionkey", nil),
					resource.TestCheckResourceAttr("citrixadc_nsencryptionkey.tf_encryptionkey", "name", "tf_encryptionkey"),
					resource.TestCheckResourceAttr("citrixadc_nsencryptionkey.tf_encryptionkey", "method", "AES256"),
				),
			},
			{
				Config: testAccNsencryptionkey_keyvalue_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsencryptionkeyExist("citrixadc_nsencryptionkey.tf_encryptionkey", nil),
					resource.TestCheckResourceAttr("citrixadc_nsencryptionkey.tf_encryptionkey", "name", "tf_encryptionkey"),
					resource.TestCheckResourceAttr("citrixadc_nsencryptionkey.tf_encryptionkey", "method", "AES256"),
				),
			},
		},
	})
}

const testAccNsencryptionkey_keyvalue_wo_step1 = `
	variable "nsencryptionkey_keyvalue_wo" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_nsencryptionkey" "tf_encryptionkey" {
		name                = "tf_encryptionkey"
		method              = "AES256"
		keyvalue_wo         = var.nsencryptionkey_keyvalue_wo
		keyvalue_wo_version = 1
		padding             = "ON"
		iv                  = "c2bf0b2e15c15004d6b14bcdc7e5e365"
		comment             = "Write-only ephemeral test step1"
	}
`

const testAccNsencryptionkey_keyvalue_wo_step2 = `
	variable "nsencryptionkey_keyvalue_wo_2" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_nsencryptionkey" "tf_encryptionkey" {
		name                = "tf_encryptionkey"
		method              = "AES256"
		keyvalue_wo         = var.nsencryptionkey_keyvalue_wo_2
		keyvalue_wo_version = 2
		padding             = "ON"
		iv                  = "c2bf0b2e15c15004d6b14bcdc7e5e365"
		comment             = "Write-only ephemeral test step2"
	}
`

func TestAccNsencryptionkey_keyvalue_wo_ephemeral(t *testing.T) {
	t.Setenv("TF_VAR_nsencryptionkey_keyvalue_wo", "26ea5537b7e0746089476e5658f9327c0b10c3b4778c673a5b38cee182874711")
	t.Setenv("TF_VAR_nsencryptionkey_keyvalue_wo_2", "8b1d2e3f4a5b6c7d8e9f0a1b2c3d4e5f6a7b8c9d0e1f2a3b4c5d6e7f8a9b0c1d")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNsencryptionkeyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNsencryptionkey_keyvalue_wo_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsencryptionkeyExist("citrixadc_nsencryptionkey.tf_encryptionkey", nil),
					resource.TestCheckResourceAttr("citrixadc_nsencryptionkey.tf_encryptionkey", "keyvalue_wo_version", "1"),
				),
			},
			{
				Config: testAccNsencryptionkey_keyvalue_wo_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsencryptionkeyExist("citrixadc_nsencryptionkey.tf_encryptionkey", nil),
					resource.TestCheckResourceAttr("citrixadc_nsencryptionkey.tf_encryptionkey", "keyvalue_wo_version", "2"),
				),
			},
		},
	})
}

func TestAccNsencryptionkey_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckNsencryptionkeyDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccNsencryptionkey_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsencryptionkeyExist("citrixadc_nsencryptionkey.tf_encryptionkey", nil),
				),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccNsencryptionkey_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsencryptionkeyExist("citrixadc_nsencryptionkey.tf_encryptionkey", nil),
				),
			},
		},
	})
}
