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

// nscentralmanagementserver is a create+delete (no-update) resource. All schema
// attributes are RequiresReplace, so the basic test only creates and verifies.
// type is the unique key (enum: CLOUD or ONPREM). ValidateConfig enforces that
// exactly one of ipaddress / servername is set; this test uses ipaddress.
//
// NOTE: registering a central management server points the ADC at an ADM. A live
// run likely requires a reachable ADM at the configured ipaddress, so this test
// may not pass against an isolated VPX. It is generated correctly regardless.

func testAccCheckNscentralmanagementserverExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No nscentralmanagementserver name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Nscentralmanagementserver.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("nscentralmanagementserver %s not found", n)
		}

		return nil
	}
}

func testAccCheckNscentralmanagementserverDestroy(s *terraform.State) error {
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_nscentralmanagementserver" {
			continue
		}
		_, err := client.FindResource(service.Nscentralmanagementserver.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("nscentralmanagementserver %s still exists", rs.Primary.ID)
		}
	}
	return nil
}

const testAccNscentralmanagementserverDataSource_basic = `

	variable "nscentralmanagementserver_password" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_nscentralmanagementserver" "tf_nscentralmanagementserver" {
		type         = "ONPREM"
		ipaddress    = "10.101.132.128"
		username     = "nsroot"
		validatecert = "NO"
		password     = var.nscentralmanagementserver_password
	}

data "citrixadc_nscentralmanagementserver" "tf_nscentralmanagementserver" {
  type       = citrixadc_nscentralmanagementserver.tf_nscentralmanagementserver.type
  depends_on = [citrixadc_nscentralmanagementserver.tf_nscentralmanagementserver]
}
`

func TestAccNscentralmanagementserverDataSource_basic(t *testing.T) {
	t.Skip("Requires valid NetScaler Console Credentials.")
	t.Setenv("TF_VAR_nscentralmanagementserver_password", "admpassword123")
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNscentralmanagementserverDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_nscentralmanagementserver.tf_nscentralmanagementserver", "type", "ONPREM"),
					resource.TestCheckResourceAttr("data.citrixadc_nscentralmanagementserver.tf_nscentralmanagementserver", "ipaddress", "10.101.132.128"),
					resource.TestCheckResourceAttr("data.citrixadc_nscentralmanagementserver.tf_nscentralmanagementserver", "username", "nsroot"),
				),
			},
		},
	})
}

// ---------------------------------------------------------------------------
// Secret: password
// ---------------------------------------------------------------------------

// Test backward-compatible path: using password (Sensitive attribute)
const testAccNscentralmanagementserver_password_step1 = `

	variable "nscentralmanagementserver_password" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_nscentralmanagementserver" "tf_nscentralmanagementserver" {
		type         = "ONPREM"
		ipaddress    = "10.101.132.128"
		username     = "nsroot"
		validatecert = "NO"
		password     = var.nscentralmanagementserver_password
	}
`

func TestAccNscentralmanagementserver_password_backward_compat(t *testing.T) {
	t.Skip("Requires valid NetScaler Console Credentials.")
	t.Setenv("TF_VAR_nscentralmanagementserver_password", "admpassword123")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNscentralmanagementserverDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNscentralmanagementserver_password_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNscentralmanagementserverExist("citrixadc_nscentralmanagementserver.tf_nscentralmanagementserver", nil),
					resource.TestCheckResourceAttr("citrixadc_nscentralmanagementserver.tf_nscentralmanagementserver", "type", "ONPREM"),
					resource.TestCheckResourceAttr("citrixadc_nscentralmanagementserver.tf_nscentralmanagementserver", "ipaddress", "10.101.132.128"),
				),
			},
		},
	})
}

// Test ephemeral path: using password_wo (WriteOnly attribute) with version tracker
const testAccNscentralmanagementserver_password_wo_step1 = `

	variable "nscentralmanagementserver_password_wo" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_nscentralmanagementserver" "tf_nscentralmanagementserver" {
		type                = "ONPREM"
		ipaddress           = "10.101.132.128"
		username            = "nsroot"
		validatecert        = "NO"
		password_wo         = var.nscentralmanagementserver_password_wo
		password_wo_version = 1
	}
`

func TestAccNscentralmanagementserver_password_wo_ephemeral(t *testing.T) {
	t.Skip("Requires valid NetScaler Console Credentials.")
	t.Setenv("TF_VAR_nscentralmanagementserver_password_wo", "admpassword123")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNscentralmanagementserverDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNscentralmanagementserver_password_wo_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNscentralmanagementserverExist("citrixadc_nscentralmanagementserver.tf_nscentralmanagementserver", nil),
					resource.TestCheckResourceAttr("citrixadc_nscentralmanagementserver.tf_nscentralmanagementserver", "type", "ONPREM"),
					resource.TestCheckResourceAttr("citrixadc_nscentralmanagementserver.tf_nscentralmanagementserver", "password_wo_version", "1"),
				),
			},
		},
	})
}

// ---------------------------------------------------------------------------
// Secret: adcpassword
// ---------------------------------------------------------------------------

// Test backward-compatible path: using adcpassword (Sensitive attribute)
const testAccNscentralmanagementserver_adcpassword_step1 = `
	variable "nscentralmanagementserver_password" {
	  type      = string
	  sensitive = true
	}

	variable "nscentralmanagementserver_adcpassword" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_nscentralmanagementserver" "tf_nscentralmanagementserver" {
		type         = "ONPREM"
		ipaddress    = "10.101.132.128"
		username     = "nsroot"
		validatecert = "NO"
		password     = var.nscentralmanagementserver_password
		adcusername = "nsroot"
		adcpassword  = var.nscentralmanagementserver_adcpassword
	}
`

func TestAccNscentralmanagementserver_adcpassword_backward_compat(t *testing.T) {
	t.Skip("Requires valid NetScaler Console Credentials.")
	t.Setenv("TF_VAR_nscentralmanagementserver_password", "admpassword123")
	t.Setenv("TF_VAR_nscentralmanagementserver_adcpassword", "oldadcpassword123")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNscentralmanagementserverDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNscentralmanagementserver_adcpassword_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNscentralmanagementserverExist("citrixadc_nscentralmanagementserver.tf_nscentralmanagementserver", nil),
					resource.TestCheckResourceAttr("citrixadc_nscentralmanagementserver.tf_nscentralmanagementserver", "type", "ONPREM"),
					resource.TestCheckResourceAttr("citrixadc_nscentralmanagementserver.tf_nscentralmanagementserver", "ipaddress", "192.0.2.100"),
				),
			},
		},
	})
}

// Test ephemeral path: using adcpassword_wo (WriteOnly attribute) with version tracker
const testAccNscentralmanagementserver_adcpassword_wo_step1 = `
	variable "nscentralmanagementserver_password" {
	  type      = string
	  sensitive = true
	}

	variable "nscentralmanagementserver_adcpassword_wo" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_nscentralmanagementserver" "tf_nscentralmanagementserver" {
		type                   = "ONPREM"
		ipaddress              = "10.101.132.128"
		username               = "nsroot"
		validatecert           = "NO"
		password     = var.nscentralmanagementserver_password
		adcusername = "nsroot"
		adcpassword_wo         = var.nscentralmanagementserver_adcpassword_wo
		adcpassword_wo_version = 1
	}
`

func TestAccNscentralmanagementserver_adcpassword_wo_ephemeral(t *testing.T) {
	t.Skip("Requires valid NetScaler Console Credentials.")
	t.Setenv("TF_VAR_nscentralmanagementserver_password", "admpassword123")
	t.Setenv("TF_VAR_nscentralmanagementserver_adcpassword_wo", "oldadcpassword123")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNscentralmanagementserverDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNscentralmanagementserver_adcpassword_wo_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNscentralmanagementserverExist("citrixadc_nscentralmanagementserver.tf_nscentralmanagementserver", nil),
					resource.TestCheckResourceAttr("citrixadc_nscentralmanagementserver.tf_nscentralmanagementserver", "type", "ONPREM"),
					resource.TestCheckResourceAttr("citrixadc_nscentralmanagementserver.tf_nscentralmanagementserver", "adcpassword_wo_version", "1"),
				),
			},
		},
	})
}
