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

func testAccCheckAzureapplicationExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No azureapplication name is set")
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
		data, err := client.FindResource(service.Azureapplication.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("azureapplication %s not found", n)
		}

		return nil
	}
}

func testAccCheckAzureapplicationDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_azureapplication" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Azureapplication.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("azureapplication %s still exists", rs.Primary.ID)
		}
	}

	return nil
}

// Test backward-compatible path: using clientsecret (Sensitive attribute)
const testAccAzureapplication_clientsecret_step1 = `

	variable "azureapplication_clientsecret" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_azureapplication" "tf_azureapplication" {
		name          = "tf_azureapplication"
		vaultresource = "vault.azure.net"
		clientid      = "<clientid>"
		tenantid      = "<tenantid>"
		clientsecret  = var.azureapplication_clientsecret
	}
`

// Update backward-compatible path: change clientsecret value (forces recreation due to RequiresReplace)
const testAccAzureapplication_clientsecret_step2 = `

	 variable "azureapplication_clientsecret_2" {
	  type      = string
	  sensitive = true
	}

	 resource "citrixadc_azureapplication" "tf_azureapplication" {
		name          = "tf_azureapplication"
		vaultresource = "vault.azure.net"
		clientid      = "<clientid>"
		tenantid      = "<tenantid>"
		clientsecret  = var.azureapplication_clientsecret_2
	}
`

func TestAccAzureapplication_clientsecret_backward_compat(t *testing.T) {
	t.Skip("Requires valid Azure credentials")
	t.Setenv("TF_VAR_azureapplication_clientsecret", "<clientsecret>")
	t.Setenv("TF_VAR_azureapplication_clientsecret_2", "<clientsecret_2>")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAzureapplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureapplication_clientsecret_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAzureapplicationExist("citrixadc_azureapplication.tf_azureapplication", nil),
					resource.TestCheckResourceAttr("citrixadc_azureapplication.tf_azureapplication", "name", "tf_azureapplication"),
					resource.TestCheckResourceAttr("citrixadc_azureapplication.tf_azureapplication", "vaultresource", "vault.azure.net"),
				),
			},
			{
				Config: testAccAzureapplication_clientsecret_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAzureapplicationExist("citrixadc_azureapplication.tf_azureapplication", nil),
					resource.TestCheckResourceAttr("citrixadc_azureapplication.tf_azureapplication", "name", "tf_azureapplication"),
					resource.TestCheckResourceAttr("citrixadc_azureapplication.tf_azureapplication", "vaultresource", "vault.azure.net"),
				),
			},
		},
	})
}

// Test ephemeral path: using clientsecret_wo (WriteOnly attribute) with version tracker
const testAccAzureapplication_clientsecret_wo_step1 = `

	variable "azureapplication_clientsecret_wo" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_azureapplication" "tf_azureapplication" {
		name                    = "tf_azureapplication"
		vaultresource           = "vault.azure.net"
		clientid                = "<clientid>"
		tenantid                = "<tenantid>"
		clientsecret_wo         = var.azureapplication_clientsecret_wo
		clientsecret_wo_version = 1
	}
`

// Update ephemeral path: bump version to trigger update with new secret (forces recreation due to RequiresReplace)
const testAccAzureapplication_clientsecret_wo_step2 = `

	 variable "azureapplication_clientsecret_wo_2" {
	  type      = string
	  sensitive = true
	}

	 resource "citrixadc_azureapplication" "tf_azureapplication" {
		name                    = "tf_azureapplication"
		vaultresource           = "vault.azure.net"
		clientid                = "<clientid>"
		tenantid                = "<tenantid>"
		clientsecret_wo         = var.azureapplication_clientsecret_wo_2
		clientsecret_wo_version = 2
	}
`

func TestAccAzureapplication_clientsecret_wo_ephemeral(t *testing.T) {
	t.Skip("Requires valid Azure credentials")
	t.Setenv("TF_VAR_azureapplication_clientsecret_wo", "<clientsecret_wo>")
	t.Setenv("TF_VAR_azureapplication_clientsecret_wo_2", "<clientsecret_wo_2>")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAzureapplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureapplication_clientsecret_wo_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAzureapplicationExist("citrixadc_azureapplication.tf_azureapplication", nil),
					resource.TestCheckResourceAttr("citrixadc_azureapplication.tf_azureapplication", "clientsecret_wo_version", "1"),
					resource.TestCheckResourceAttr("citrixadc_azureapplication.tf_azureapplication", "name", "tf_azureapplication"),
					resource.TestCheckResourceAttr("citrixadc_azureapplication.tf_azureapplication", "vaultresource", "vault.azure.net"),
				),
			},
			{
				Config: testAccAzureapplication_clientsecret_wo_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAzureapplicationExist("citrixadc_azureapplication.tf_azureapplication", nil),
					resource.TestCheckResourceAttr("citrixadc_azureapplication.tf_azureapplication", "clientsecret_wo_version", "2"),
					resource.TestCheckResourceAttr("citrixadc_azureapplication.tf_azureapplication", "name", "tf_azureapplication"),
					resource.TestCheckResourceAttr("citrixadc_azureapplication.tf_azureapplication", "vaultresource", "vault.azure.net"),
				),
			},
		},
	})
}

func TestAccAzureapplication_import(t *testing.T) {
	t.Skip("TODO: Requires review")
	const resAddr = "citrixadc_azureapplication.tf_azureapplication"
	t.Setenv("TF_VAR_azureapplication_clientsecret", "<clientsecret>")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAzureapplicationDestroy,
		Steps: []resource.TestStep{
			{Config: testAccAzureapplication_clientsecret_step1},
			{
				Config:            testAccAzureapplication_clientsecret_step1,
				ResourceName:      resAddr,
				ImportState:       true,
				ImportStateVerify: true,
				// clientsecret is Sensitive and never echoed back by NITRO, and
				// clientsecret_wo_version is a computed version tracker not returned
				// by the API, so neither can round-trip through import.
				ImportStateVerifyIgnore: []string{"clientsecret", "clientsecret_wo_version"},
			},
		},
	})
}

const testAccAzureapplicationDataSource_basic = `

variable "azureapplication_clientsecret" {
	type      = string
	sensitive = true
}

resource "citrixadc_azureapplication" "tf_azureapplication" {
  name          = "tf_azureapplication"
  vaultresource = "vault.azure.net"
  clientid      = "<clientid>"
  tenantid      = "<tenantid>"
  clientsecret  = var.azureapplication_clientsecret
}

data "citrixadc_azureapplication" "tf_azureapplication" {
  name       = citrixadc_azureapplication.tf_azureapplication.name
  depends_on = [citrixadc_azureapplication.tf_azureapplication]
}
`

func TestAccAzureapplicationDataSource_basic(t *testing.T) {
	t.Skip("Requires valid Azure credentials")
	t.Setenv("TF_VAR_azureapplication_clientsecret", "<clientsecret>")
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureapplicationDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_azureapplication.tf_azureapplication", "name", "tf_azureapplication"),
					resource.TestCheckResourceAttr("data.citrixadc_azureapplication.tf_azureapplication", "vaultresource", "vault.azure.net"),
					resource.TestCheckResourceAttr("data.citrixadc_azureapplication.tf_azureapplication", "clientid", "<clientid>"),
					resource.TestCheckResourceAttr("data.citrixadc_azureapplication.tf_azureapplication", "tenantid", "<tenantid>"),
				),
			},
		},
	})
}
