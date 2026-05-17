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

const testAccNsencryptionparams_basic = `
	resource "citrixadc_nsencryptionparams" "tf_nsencryptionparams" {
		method = "DES-CFB"
	}
`

const testAccNsencryptionparams_update = `
	resource "citrixadc_nsencryptionparams" "tf_nsencryptionparams" {
		method = "RC4"
	}
`

func TestAccNsencryptionparams_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccNsencryptionparams_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsencryptionparamsExist("citrixadc_nsencryptionparams.tf_nsencryptionparams", nil),
					resource.TestCheckResourceAttr("citrixadc_nsencryptionparams.tf_nsencryptionparams", "method", "DES-CFB"),
				),
			},
			{
				Config: testAccNsencryptionparams_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsencryptionparamsExist("citrixadc_nsencryptionparams.tf_nsencryptionparams", nil),
					resource.TestCheckResourceAttr("citrixadc_nsencryptionparams.tf_nsencryptionparams", "method", "RC4"),
				),
			},
		},
	})
}

func testAccCheckNsencryptionparamsExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No nsencryptionparams name is set")
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
		data, err := client.FindResource(service.Nsencryptionparams.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("nsencryptionparams %s not found", n)
		}

		return nil
	}
}

const testAccNsencryptionparamsDataSource_basic = `

resource "citrixadc_nsencryptionparams" "tf_nsencryptionparams" {
	method = "AES256"
}

data "citrixadc_nsencryptionparams" "tf_nsencryptionparams_ds" {
	depends_on = [citrixadc_nsencryptionparams.tf_nsencryptionparams]
}
`

func TestAccNsencryptionparamsDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNsencryptionparamsDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_nsencryptionparams.tf_nsencryptionparams_ds", "method", "AES256"),
					resource.TestCheckResourceAttrSet("data.citrixadc_nsencryptionparams.tf_nsencryptionparams_ds", "id"),
				),
			},
		},
	})
}

// Test backward-compatible path: using keyvalue (Sensitive attribute)
const testAccNsencryptionparams_keyvalue_step1 = `
	variable "nsencryptionparams_keyvalue" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_nsencryptionparams" "tf_nsencryptionparams" {
		method   = "AES256"
		keyvalue = var.nsencryptionparams_keyvalue
	}
`

const testAccNsencryptionparams_keyvalue_step2 = `
	variable "nsencryptionparams_keyvalue_2" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_nsencryptionparams" "tf_nsencryptionparams" {
		method   = "AES256"
		keyvalue = var.nsencryptionparams_keyvalue_2
	}
`

func TestAccNsencryptionparams_keyvalue_backward_compat(t *testing.T) {
	t.Setenv("TF_VAR_nsencryptionparams_keyvalue", "")
	t.Setenv("TF_VAR_nsencryptionparams_keyvalue_2", "")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccNsencryptionparams_keyvalue_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsencryptionparamsExist("citrixadc_nsencryptionparams.tf_nsencryptionparams", nil),
					resource.TestCheckResourceAttr("citrixadc_nsencryptionparams.tf_nsencryptionparams", "method", "AES256"),
				),
			},
			{
				Config: testAccNsencryptionparams_keyvalue_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsencryptionparamsExist("citrixadc_nsencryptionparams.tf_nsencryptionparams", nil),
					resource.TestCheckResourceAttr("citrixadc_nsencryptionparams.tf_nsencryptionparams", "method", "AES256"),
				),
			},
		},
	})
}

// Test ephemeral path: using keyvalue_wo (WriteOnly attribute) with version tracker
const testAccNsencryptionparams_wo_step1 = `
	variable "nsencryptionparams_keyvalue_wo" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_nsencryptionparams" "tf_nsencryptionparams" {
		method              = "AES256"
		keyvalue_wo         = var.nsencryptionparams_keyvalue_wo
		keyvalue_wo_version = 1
	}
`

const testAccNsencryptionparams_wo_step2 = `
	variable "nsencryptionparams_keyvalue_wo_2" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_nsencryptionparams" "tf_nsencryptionparams" {
		method              = "AES256"
		keyvalue_wo         = var.nsencryptionparams_keyvalue_wo_2
		keyvalue_wo_version = 2
	}
`

func TestAccNsencryptionparams_keyvalue_wo_ephemeral(t *testing.T) {
	t.Setenv("TF_VAR_nsencryptionparams_keyvalue_wo", "")
	t.Setenv("TF_VAR_nsencryptionparams_keyvalue_wo_2", "")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccNsencryptionparams_wo_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsencryptionparamsExist("citrixadc_nsencryptionparams.tf_nsencryptionparams", nil),
					resource.TestCheckResourceAttr("citrixadc_nsencryptionparams.tf_nsencryptionparams", "method", "AES256"),
					resource.TestCheckResourceAttr("citrixadc_nsencryptionparams.tf_nsencryptionparams", "keyvalue_wo_version", "1"),
				),
			},
			{
				Config: testAccNsencryptionparams_wo_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsencryptionparamsExist("citrixadc_nsencryptionparams.tf_nsencryptionparams", nil),
					resource.TestCheckResourceAttr("citrixadc_nsencryptionparams.tf_nsencryptionparams", "method", "AES256"),
					resource.TestCheckResourceAttr("citrixadc_nsencryptionparams.tf_nsencryptionparams", "keyvalue_wo_version", "2"),
				),
			},
		},
	})
}
