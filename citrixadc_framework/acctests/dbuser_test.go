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
		loggedin = false
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
					resource.TestCheckResourceAttr("data.citrixadc_dbuser.tf_dbuser_ds", "loggedin", "false"),
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

func TestAccDbuser_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckDbuserDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccDbuser_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDbuserExist("citrixadc_dbuser.tf_dbuser", nil),
				),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccDbuser_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDbuserExist("citrixadc_dbuser.tf_dbuser", nil),
				),
			},
		},
	})
}
