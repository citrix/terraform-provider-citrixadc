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

const testAccNshmackey_add = `

	resource "citrixadc_nshmackey" "tf_nshmackey" {
		name     = "tf_nshmackey"
		digest   = "MD4"
		keyvalue = "AUTO"
		comment  = "Testing"
	}
`
const testAccNshmackey_update = `

	resource "citrixadc_nshmackey" "tf_nshmackey" {
		name     = "tf_nshmackey"
		digest   = "MD2"
		keyvalue = "AUTO"
		comment  = "Testing_sample"
	}
`

func TestAccNshmackey_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNshmackeyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNshmackey_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNshmackeyExist("citrixadc_nshmackey.tf_nshmackey", nil),
					resource.TestCheckResourceAttr("citrixadc_nshmackey.tf_nshmackey", "name", "tf_nshmackey"),
					resource.TestCheckResourceAttr("citrixadc_nshmackey.tf_nshmackey", "digest", "MD4"),
					resource.TestCheckResourceAttr("citrixadc_nshmackey.tf_nshmackey", "comment", "Testing"),
				),
			},
			{
				Config: testAccNshmackey_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNshmackeyExist("citrixadc_nshmackey.tf_nshmackey", nil),
					resource.TestCheckResourceAttr("citrixadc_nshmackey.tf_nshmackey", "name", "tf_nshmackey"),
					resource.TestCheckResourceAttr("citrixadc_nshmackey.tf_nshmackey", "digest", "MD2"),
					resource.TestCheckResourceAttr("citrixadc_nshmackey.tf_nshmackey", "comment", "Testing_sample"),
				),
			},
		},
	})
}

func testAccCheckNshmackeyExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No nshmackey name is set")
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
		data, err := client.FindResource("nshmackey", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("nshmackey %s not found", n)
		}

		return nil
	}
}

func testAccCheckNshmackeyDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_nshmackey" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("nshmackey", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("nshmackey %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccNshmackeyDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNshmackeyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNshmackeyDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_nshmackey.tf_hmackey_ds", "name", "tf_hmackey_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_nshmackey.tf_hmackey_ds", "digest", "SHA256"),
					resource.TestCheckResourceAttr("data.citrixadc_nshmackey.tf_hmackey_ds", "comment", "Test HMAC key for datasource"),
				),
			},
		},
	})
}

const testAccNshmackeyDataSource_basic = `

resource "citrixadc_nshmackey" "tf_hmackey_ds" {
	name     = "tf_hmackey_ds"
	digest   = "SHA256"
	keyvalue = "616263"
	comment  = "Test HMAC key for datasource"
}

data "citrixadc_nshmackey" "tf_hmackey_ds" {
	name = citrixadc_nshmackey.tf_hmackey_ds.name
	depends_on = [citrixadc_nshmackey.tf_hmackey_ds]
}
`

// --- Ephemeral / Write-Only Tests for keyvalue ---

const testAccNshmackey_keyvalue_step1 = `
	variable "nshmackey_keyvalue" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_nshmackey" "tf_nshmackey" {
		name     = "tf_nshmackey"
		digest   = "MD4"
		keyvalue = var.nshmackey_keyvalue
		comment  = "Backward compat test step1"
	}
`

const testAccNshmackey_keyvalue_step2 = `
	variable "nshmackey_keyvalue_2" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_nshmackey" "tf_nshmackey" {
		name     = "tf_nshmackey"
		digest   = "MD4"
		keyvalue = var.nshmackey_keyvalue_2
		comment  = "Backward compat test step2"
	}
`

func TestAccNshmackey_keyvalue_backward_compat(t *testing.T) {
	t.Setenv("TF_VAR_nshmackey_keyvalue", "0102030405060708090a0b0c0d0e0f10")
	t.Setenv("TF_VAR_nshmackey_keyvalue_2", "1112131415161718191a1b1c1d1e1f20")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNshmackeyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNshmackey_keyvalue_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNshmackeyExist("citrixadc_nshmackey.tf_nshmackey", nil),
					resource.TestCheckResourceAttr("citrixadc_nshmackey.tf_nshmackey", "name", "tf_nshmackey"),
					resource.TestCheckResourceAttr("citrixadc_nshmackey.tf_nshmackey", "digest", "MD4"),
				),
			},
			{
				Config: testAccNshmackey_keyvalue_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNshmackeyExist("citrixadc_nshmackey.tf_nshmackey", nil),
					resource.TestCheckResourceAttr("citrixadc_nshmackey.tf_nshmackey", "name", "tf_nshmackey"),
					resource.TestCheckResourceAttr("citrixadc_nshmackey.tf_nshmackey", "digest", "MD4"),
				),
			},
		},
	})
}

const testAccNshmackey_keyvalue_wo_step1 = `
	variable "nshmackey_keyvalue_wo" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_nshmackey" "tf_nshmackey" {
		name                = "tf_nshmackey"
		digest              = "MD4"
		keyvalue_wo         = var.nshmackey_keyvalue_wo
		keyvalue_wo_version = 1
		comment             = "Write-only ephemeral test step1"
	}
`

const testAccNshmackey_keyvalue_wo_step2 = `
	variable "nshmackey_keyvalue_wo_2" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_nshmackey" "tf_nshmackey" {
		name                = "tf_nshmackey"
		digest              = "MD4"
		keyvalue_wo         = var.nshmackey_keyvalue_wo_2
		keyvalue_wo_version = 2
		comment             = "Write-only ephemeral test step2"
	}
`

func TestAccNshmackey_keyvalue_wo_ephemeral(t *testing.T) {
	t.Setenv("TF_VAR_nshmackey_keyvalue_wo", "0102030405060708090a0b0c0d0e0f10")
	t.Setenv("TF_VAR_nshmackey_keyvalue_wo_2", "1112131415161718191a1b1c1d1e1f20")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNshmackeyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNshmackey_keyvalue_wo_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNshmackeyExist("citrixadc_nshmackey.tf_nshmackey", nil),
					resource.TestCheckResourceAttr("citrixadc_nshmackey.tf_nshmackey", "keyvalue_wo_version", "1"),
				),
			},
			{
				Config: testAccNshmackey_keyvalue_wo_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNshmackeyExist("citrixadc_nshmackey.tf_nshmackey", nil),
					resource.TestCheckResourceAttr("citrixadc_nshmackey.tf_nshmackey", "keyvalue_wo_version", "2"),
				),
			},
		},
	})
}

func TestAccNshmackey_import(t *testing.T) {
	const resAddr = "citrixadc_nshmackey.tf_nshmackey"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNshmackeyDestroy,
		Steps: []resource.TestStep{
			{Config: testAccNshmackey_add},
			{
				Config:            testAccNshmackey_add,
				ResourceName:      resAddr,
				ImportState:       true,
				ImportStateVerify: true,
				// keyvalue is a secret the NITRO API never echoes back, and
				// keyvalue_wo_version is a computed version tracker not populated on import;
				// neither can round-trip through import.
				ImportStateVerifyIgnore: []string{"keyvalue", "keyvalue_wo_version"},
			},
		},
	})
}

func TestAccNshmackey_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckNshmackeyDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccNshmackey_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNshmackeyExist("citrixadc_nshmackey.tf_nshmackey", nil),
				),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccNshmackey_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNshmackeyExist("citrixadc_nshmackey.tf_nshmackey", nil),
				),
			},
		},
	})
}
