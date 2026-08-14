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
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const testAccDbuser_basic = `
	resource "citrixadc_dbuser" "tf_dbuser" {
		username = "user1"
	}
`
const testAccDbuser_update = `
	resource "citrixadc_dbuser" "tf_dbuser" {
		username = "user1"
		password = "1234"
	}
`

const testAccDbuserDataSource_basic = `
	resource "citrixadc_dbuser" "tf_dbuser" {
		username = "user1"
		password = "1234"
	}
	
	data "citrixadc_dbuser" "tf_dbuser_ds" {
		username = citrixadc_dbuser.tf_dbuser.username
	}
`

func TestAccDbuserDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDbuserDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_dbuser.tf_dbuser_ds", "username", "user1"),
					resource.TestCheckResourceAttrSet("data.citrixadc_dbuser.tf_dbuser_ds", "id"),
				),
			},
		},
	})
}

func TestAccDbuser_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDbuserDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDbuser_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDbuserExist("citrixadc_dbuser.tf_dbuser", nil),
					resource.TestCheckResourceAttr("citrixadc_dbuser.tf_dbuser", "username", "user1"),
				),
			},
			{
				Config: testAccDbuser_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDbuserExist("citrixadc_dbuser.tf_dbuser", nil),
					resource.TestCheckResourceAttr("citrixadc_dbuser.tf_dbuser", "username", "user1"),
					resource.TestCheckResourceAttr("citrixadc_dbuser.tf_dbuser", "password", "1234"),
				),
			},
		},
	})
}

func TestAccDbuser_import(t *testing.T) {
	const resAddr = "citrixadc_dbuser.tf_dbuser"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDbuserDestroy,
		Steps: []resource.TestStep{
			{Config: testAccDbuser_basic},
			{
				Config:            testAccDbuser_basic,
				ResourceName:      resAddr,
				ImportState:       true,
				ImportStateVerify: true,
				// password_wo_version is a write-only version tracker that NITRO
				// does not return; on import there is no config to retain it from,
				// so it cannot round-trip.
				ImportStateVerifyIgnore: []string{"password_wo_version"},
			},
		},
	})
}

func testAccCheckDbuserExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No dbuser name is set")
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
		data, err := client.FindResource(service.Dbuser.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("dbuser %s not found", n)
		}

		return nil
	}
}

func testAccCheckDbuserDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_dbuser" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Dbuser.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("dbuser %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

// ============================================================
// Ephemeral / Write-Only tests for secret attribute: password
// ============================================================

const testAccDbuser_password_step1 = `

variable "dbuser_password" {
  type      = string
  sensitive = true
}

resource "citrixadc_dbuser" "tf_dbuser_password" {
  username = "tf_test_dbuser_password"
  password = var.dbuser_password
}
`

const testAccDbuser_password_step2 = `

variable "dbuser_password_2" {
  type      = string
  sensitive = true
}

resource "citrixadc_dbuser" "tf_dbuser_password" {
  username = "tf_test_dbuser_password"
  password = var.dbuser_password_2
}
`

func TestAccDbuser_password_backward_compat(t *testing.T) {
	t.Setenv("TF_VAR_dbuser_password", "Password1!")
	t.Setenv("TF_VAR_dbuser_password_2", "Password2!")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDbuserDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDbuser_password_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDbuserExist("citrixadc_dbuser.tf_dbuser_password", nil),
					resource.TestCheckResourceAttr("citrixadc_dbuser.tf_dbuser_password", "username", "tf_test_dbuser_password"),
				),
			},
			{
				Config: testAccDbuser_password_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDbuserExist("citrixadc_dbuser.tf_dbuser_password", nil),
					resource.TestCheckResourceAttr("citrixadc_dbuser.tf_dbuser_password", "username", "tf_test_dbuser_password"),
				),
			},
		},
	})
}

const testAccDbuser_password_wo_step1 = `

variable "dbuser_password_wo" {
  type      = string
  sensitive = true
}

resource "citrixadc_dbuser" "tf_dbuser_password_wo" {
  username            = "tf_test_dbuser_password_wo"
  password_wo         = var.dbuser_password_wo
  password_wo_version = 1
}
`

const testAccDbuser_password_wo_step2 = `

variable "dbuser_password_wo_2" {
  type      = string
  sensitive = true
}

resource "citrixadc_dbuser" "tf_dbuser_password_wo" {
  username            = "tf_test_dbuser_password_wo"
  password_wo         = var.dbuser_password_wo_2
  password_wo_version = 2
}
`

func TestAccDbuser_password_wo_ephemeral(t *testing.T) {
	t.Setenv("TF_VAR_dbuser_password_wo", "Password1!")
	t.Setenv("TF_VAR_dbuser_password_wo_2", "Password2!")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDbuserDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDbuser_password_wo_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDbuserExist("citrixadc_dbuser.tf_dbuser_password_wo", nil),
					resource.TestCheckResourceAttr("citrixadc_dbuser.tf_dbuser_password_wo", "username", "tf_test_dbuser_password_wo"),
					resource.TestCheckResourceAttr("citrixadc_dbuser.tf_dbuser_password_wo", "password_wo_version", "1"),
				),
			},
			{
				Config: testAccDbuser_password_wo_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDbuserExist("citrixadc_dbuser.tf_dbuser_password_wo", nil),
					resource.TestCheckResourceAttr("citrixadc_dbuser.tf_dbuser_password_wo", "username", "tf_test_dbuser_password_wo"),
					resource.TestCheckResourceAttr("citrixadc_dbuser.tf_dbuser_password_wo", "password_wo_version", "2"),
				),
			},
		},
	})
}

// testAccDbuser_upgrade_basic is the fixture for the SDK v2 -> Framework upgrade test.
// It sets a password because password_wo_version is Optional+Computed with a default of 1:
// upgrading from 2.2.0 state (which has no password_wo_version) resolves it null->1, which
// triggers an in-place dbuser Update. The ADC requires a password on a dbuser update
// (errorcode 1095 otherwise), so the upgrade fixture must carry one. (testAccDbuser_basic is
// intentionally password-less and is used by the create/import tests.)
const testAccDbuser_upgrade_basic = `
	resource "citrixadc_dbuser" "tf_dbuser" {
		username = "user1"
		password = "1234"
	}
`

func TestAccDbuser_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckDbuserDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccDbuser_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDbuserExist("citrixadc_dbuser.tf_dbuser", nil),
				),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccDbuser_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDbuserExist("citrixadc_dbuser.tf_dbuser", nil),
				),
			},
		},
	})
}

func TestAccDbuser_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_dbuser.tf_dbuser"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDbuserDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDbuser_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckDbuserExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Dbuser.Type(), "user1"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccDbuser_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckDbuserExist(resAddr, nil)),
			},
		},
	})
}
