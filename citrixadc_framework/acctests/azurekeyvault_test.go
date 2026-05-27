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

// No attributes are updateable (all have RequiresReplace), so step1 and step2
// use the same config to test idempotency.
const testAccAzurekeyvault_basic_step1 = `

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

resource "citrixadc_azurekeyvault" "tf_azurekeyvault" {
  name             = "tf_azurekeyvault"
  azureapplication = citrixadc_azureapplication.tf_azureapplication.name
  azurevaultname   = "tfadctest.vault.azure.net"
}

`

const testAccAzurekeyvault_basic_step2 = `

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

resource "citrixadc_azurekeyvault" "tf_azurekeyvault" {
  name             = "tf_azurekeyvault"
  azureapplication = citrixadc_azureapplication.tf_azureapplication.name
  azurevaultname   = "tfadcnew.vault.azure.net"
}

`

func TestAccAzurekeyvault_basic(t *testing.T) {
	t.Skip("Requires valid Azure credentials")
	t.Setenv("TF_VAR_azureapplication_clientsecret", "<clientsecret>")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAzurekeyvaultDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzurekeyvault_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAzurekeyvaultExist("citrixadc_azurekeyvault.tf_azurekeyvault", nil),
					resource.TestCheckResourceAttr("citrixadc_azurekeyvault.tf_azurekeyvault", "name", "tf_azurekeyvault"),
					resource.TestCheckResourceAttr("citrixadc_azurekeyvault.tf_azurekeyvault", "azureapplication", "tf_azureapplication"),
					resource.TestCheckResourceAttr("citrixadc_azurekeyvault.tf_azurekeyvault", "azurevaultname", "tfadctest.vault.azure.net"),
				),
			},
			{
				Config: testAccAzurekeyvault_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAzurekeyvaultExist("citrixadc_azurekeyvault.tf_azurekeyvault", nil),
					resource.TestCheckResourceAttr("citrixadc_azurekeyvault.tf_azurekeyvault", "name", "tf_azurekeyvault"),
					resource.TestCheckResourceAttr("citrixadc_azurekeyvault.tf_azurekeyvault", "azureapplication", "tf_azureapplication"),
					resource.TestCheckResourceAttr("citrixadc_azurekeyvault.tf_azurekeyvault", "azurevaultname", "tfadcnew.vault.azure.net"),
				),
			},
		},
	})
}

func testAccCheckAzurekeyvaultExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No azurekeyvault name is set")
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
		data, err := client.FindResource(service.Azurekeyvault.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("azurekeyvault %s not found", n)
		}

		return nil
	}
}

func testAccCheckAzurekeyvaultDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_azurekeyvault" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Azurekeyvault.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("azurekeyvault %s still exists", rs.Primary.ID)
		}
	}

	return nil
}

const testAccAzurekeyvaultDataSource_basic = `

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

resource "citrixadc_azurekeyvault" "tf_azurekeyvault" {
  name             = "tf_azurekeyvault"
  azureapplication = citrixadc_azureapplication.tf_azureapplication.name
  azurevaultname   = "tfadcnew.vault.azure.net"
}

data "citrixadc_azurekeyvault" "tf_azurekeyvault" {
  name       = citrixadc_azurekeyvault.tf_azurekeyvault.name
  depends_on = [citrixadc_azurekeyvault.tf_azurekeyvault]
}
`

func TestAccAzurekeyvaultDataSource_basic(t *testing.T) {
	t.Skip("Requires valid Azure credentials")
	t.Setenv("TF_VAR_azureapplication_clientsecret", "<clientsecret>")
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAzurekeyvaultDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_azurekeyvault.tf_azurekeyvault", "name", "tf_azurekeyvault"),
					resource.TestCheckResourceAttr("data.citrixadc_azurekeyvault.tf_azurekeyvault", "azureapplication", "tf_azureapplication"),
					resource.TestCheckResourceAttr("data.citrixadc_azurekeyvault.tf_azurekeyvault", "azurevaultname", "tfadcnew.vault.azure.net"),
				),
			},
		},
	})
}
