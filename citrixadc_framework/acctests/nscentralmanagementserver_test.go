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
const testAccNscentralmanagementserver_basic_step1 = `
resource "citrixadc_nscentralmanagementserver" "tf_nscentralmanagementserver" {
  type         = "ONPREM"
  ipaddress    = "192.0.2.100"
  username     = "admin_user"
  validatecert = "NO"
}

`

func TestAccNscentralmanagementserver_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNscentralmanagementserverDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNscentralmanagementserver_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNscentralmanagementserverExist("citrixadc_nscentralmanagementserver.tf_nscentralmanagementserver", nil),
					resource.TestCheckResourceAttr("citrixadc_nscentralmanagementserver.tf_nscentralmanagementserver", "type", "ONPREM"),
					resource.TestCheckResourceAttr("citrixadc_nscentralmanagementserver.tf_nscentralmanagementserver", "ipaddress", "192.0.2.100"),
					resource.TestCheckResourceAttr("citrixadc_nscentralmanagementserver.tf_nscentralmanagementserver", "username", "admin_user"),
					resource.TestCheckResourceAttr("citrixadc_nscentralmanagementserver.tf_nscentralmanagementserver", "validatecert", "NO"),
				),
			},
		},
	})
}

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

resource "citrixadc_nscentralmanagementserver" "tf_nscentralmanagementserver" {
  type         = "ONPREM"
  ipaddress    = "192.0.2.100"
  username     = "admin_user"
  validatecert = "NO"
}

data "citrixadc_nscentralmanagementserver" "tf_nscentralmanagementserver" {
  type       = citrixadc_nscentralmanagementserver.tf_nscentralmanagementserver.type
  depends_on = [citrixadc_nscentralmanagementserver.tf_nscentralmanagementserver]
}
`

func TestAccNscentralmanagementserverDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNscentralmanagementserverDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_nscentralmanagementserver.tf_nscentralmanagementserver", "type", "ONPREM"),
					resource.TestCheckResourceAttr("data.citrixadc_nscentralmanagementserver.tf_nscentralmanagementserver", "ipaddress", "192.0.2.100"),
					resource.TestCheckResourceAttr("data.citrixadc_nscentralmanagementserver.tf_nscentralmanagementserver", "username", "admin_user"),
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
		ipaddress    = "192.0.2.100"
		username     = "admin_user"
		validatecert = "NO"
		password     = var.nscentralmanagementserver_password
	}
`

// Update backward-compatible path: change password value (RequiresReplace -> recreate)
const testAccNscentralmanagementserver_password_step2 = `

	 variable "nscentralmanagementserver_password_2" {
	  type      = string
	  sensitive = true
	}

	 resource "citrixadc_nscentralmanagementserver" "tf_nscentralmanagementserver" {
		type         = "ONPREM"
		ipaddress    = "192.0.2.100"
		username     = "admin_user"
		validatecert = "NO"
		password     = var.nscentralmanagementserver_password_2
	}
`

func TestAccNscentralmanagementserver_password_backward_compat(t *testing.T) {
	t.Setenv("TF_VAR_nscentralmanagementserver_password", "oldpassword123")
	t.Setenv("TF_VAR_nscentralmanagementserver_password_2", "newpassword456")
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
					resource.TestCheckResourceAttr("citrixadc_nscentralmanagementserver.tf_nscentralmanagementserver", "ipaddress", "192.0.2.100"),
				),
			},
			{
				Config: testAccNscentralmanagementserver_password_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNscentralmanagementserverExist("citrixadc_nscentralmanagementserver.tf_nscentralmanagementserver", nil),
					resource.TestCheckResourceAttr("citrixadc_nscentralmanagementserver.tf_nscentralmanagementserver", "type", "ONPREM"),
					resource.TestCheckResourceAttr("citrixadc_nscentralmanagementserver.tf_nscentralmanagementserver", "ipaddress", "192.0.2.100"),
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
		ipaddress           = "192.0.2.100"
		username            = "admin_user"
		validatecert        = "NO"
		password_wo         = var.nscentralmanagementserver_password_wo
		password_wo_version = 1
	}
`

// Update ephemeral path: bump version to trigger recreate with new password
const testAccNscentralmanagementserver_password_wo_step2 = `

	 variable "nscentralmanagementserver_password_wo_2" {
	  type      = string
	  sensitive = true
	}

	 resource "citrixadc_nscentralmanagementserver" "tf_nscentralmanagementserver" {
		type                = "ONPREM"
		ipaddress           = "192.0.2.100"
		username            = "admin_user"
		validatecert        = "NO"
		password_wo         = var.nscentralmanagementserver_password_wo_2
		password_wo_version = 2
	}
`

func TestAccNscentralmanagementserver_password_wo_ephemeral(t *testing.T) {
	t.Setenv("TF_VAR_nscentralmanagementserver_password_wo", "ephemeral_pass1")
	t.Setenv("TF_VAR_nscentralmanagementserver_password_wo_2", "ephemeral_pass2")
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
			{
				Config: testAccNscentralmanagementserver_password_wo_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNscentralmanagementserverExist("citrixadc_nscentralmanagementserver.tf_nscentralmanagementserver", nil),
					resource.TestCheckResourceAttr("citrixadc_nscentralmanagementserver.tf_nscentralmanagementserver", "type", "ONPREM"),
					resource.TestCheckResourceAttr("citrixadc_nscentralmanagementserver.tf_nscentralmanagementserver", "password_wo_version", "2"),
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

	variable "nscentralmanagementserver_adcpassword" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_nscentralmanagementserver" "tf_nscentralmanagementserver" {
		type         = "ONPREM"
		ipaddress    = "192.0.2.100"
		username     = "admin_user"
		validatecert = "NO"
		adcpassword  = var.nscentralmanagementserver_adcpassword
	}
`

// Update backward-compatible path: change adcpassword value (RequiresReplace -> recreate)
const testAccNscentralmanagementserver_adcpassword_step2 = `

	 variable "nscentralmanagementserver_adcpassword_2" {
	  type      = string
	  sensitive = true
	}

	 resource "citrixadc_nscentralmanagementserver" "tf_nscentralmanagementserver" {
		type         = "ONPREM"
		ipaddress    = "192.0.2.100"
		username     = "admin_user"
		validatecert = "NO"
		adcpassword  = var.nscentralmanagementserver_adcpassword_2
	}
`

func TestAccNscentralmanagementserver_adcpassword_backward_compat(t *testing.T) {
	t.Setenv("TF_VAR_nscentralmanagementserver_adcpassword", "oldadcpassword123")
	t.Setenv("TF_VAR_nscentralmanagementserver_adcpassword_2", "newadcpassword456")
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
			{
				Config: testAccNscentralmanagementserver_adcpassword_step2,
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

	variable "nscentralmanagementserver_adcpassword_wo" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_nscentralmanagementserver" "tf_nscentralmanagementserver" {
		type                   = "ONPREM"
		ipaddress              = "192.0.2.100"
		username               = "admin_user"
		validatecert           = "NO"
		adcpassword_wo         = var.nscentralmanagementserver_adcpassword_wo
		adcpassword_wo_version = 1
	}
`

// Update ephemeral path: bump version to trigger recreate with new adcpassword
const testAccNscentralmanagementserver_adcpassword_wo_step2 = `

	 variable "nscentralmanagementserver_adcpassword_wo_2" {
	  type      = string
	  sensitive = true
	}

	 resource "citrixadc_nscentralmanagementserver" "tf_nscentralmanagementserver" {
		type                   = "ONPREM"
		ipaddress              = "192.0.2.100"
		username               = "admin_user"
		validatecert           = "NO"
		adcpassword_wo         = var.nscentralmanagementserver_adcpassword_wo_2
		adcpassword_wo_version = 2
	}
`

func TestAccNscentralmanagementserver_adcpassword_wo_ephemeral(t *testing.T) {
	t.Setenv("TF_VAR_nscentralmanagementserver_adcpassword_wo", "ephemeral_adcpass1")
	t.Setenv("TF_VAR_nscentralmanagementserver_adcpassword_wo_2", "ephemeral_adcpass2")
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
			{
				Config: testAccNscentralmanagementserver_adcpassword_wo_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNscentralmanagementserverExist("citrixadc_nscentralmanagementserver.tf_nscentralmanagementserver", nil),
					resource.TestCheckResourceAttr("citrixadc_nscentralmanagementserver.tf_nscentralmanagementserver", "type", "ONPREM"),
					resource.TestCheckResourceAttr("citrixadc_nscentralmanagementserver.tf_nscentralmanagementserver", "adcpassword_wo_version", "2"),
				),
			},
		},
	})
}
