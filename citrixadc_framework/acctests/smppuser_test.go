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

const testAccSmppuser_basic = `


resource "citrixadc_smppuser" "tf_smppuser" {
	username = "user1"
	password = "abc"
	}
`
const testAccSmppuser_update = `


resource "citrixadc_smppuser" "tf_smppuser" {
	username = "user1"
	password = "abcd"
	}
`

func TestAccSmppuser_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSmppuserDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSmppuser_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSmppuserExist("citrixadc_smppuser.tf_smppuser", nil),
					resource.TestCheckResourceAttr("citrixadc_smppuser.tf_smppuser", "username", "user1"),
					resource.TestCheckResourceAttr("citrixadc_smppuser.tf_smppuser", "password", "abc"),
				),
			},
			{
				Config: testAccSmppuser_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSmppuserExist("citrixadc_smppuser.tf_smppuser", nil),
					resource.TestCheckResourceAttr("citrixadc_smppuser.tf_smppuser", "username", "user1"),
					resource.TestCheckResourceAttr("citrixadc_smppuser.tf_smppuser", "password", "abcd"),
				),
			},
		},
	})
}

func testAccCheckSmppuserExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No smppuser name is set")
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
		data, err := client.FindResource("smppuser", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("smppuser %s not found", n)
		}

		return nil
	}
}

func testAccCheckSmppuserDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_smppuser" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("smppuser", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("smppuser %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccSmppuserDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSmppuserDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSmppuserDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_smppuser.tf_smppuser_ds", "username", "tf_smppuser_ds"),
				),
			},
		},
	})
}

const testAccSmppuserDataSource_basic = `

resource "citrixadc_smppuser" "tf_smppuser_ds" {
	username = "tf_smppuser_ds"
	password = "datasourcepass"
}

data "citrixadc_smppuser" "tf_smppuser_ds" {
	username = citrixadc_smppuser.tf_smppuser_ds.username
	depends_on = [citrixadc_smppuser.tf_smppuser_ds]
}
`

// Test backward-compatible path: using password (Sensitive attribute)
const testAccSmppuser_password_step1 = `
	variable "smppuser_password" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_smppuser" "tf_smppuser_ephem" {
		username = "tf_smppuser_ephem"
		password = var.smppuser_password
	}
`

const testAccSmppuser_password_step2 = `
	variable "smppuser_password_2" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_smppuser" "tf_smppuser_ephem" {
		username = "tf_smppuser_ephem"
		password = var.smppuser_password_2
	}
`

func TestAccSmppuser_password_backward_compat(t *testing.T) {
	t.Setenv("TF_VAR_smppuser_password", "smpppass1")
	t.Setenv("TF_VAR_smppuser_password_2", "smpppass2")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSmppuserDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSmppuser_password_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSmppuserExist("citrixadc_smppuser.tf_smppuser_ephem", nil),
					resource.TestCheckResourceAttr("citrixadc_smppuser.tf_smppuser_ephem", "username", "tf_smppuser_ephem"),
				),
			},
			{
				Config: testAccSmppuser_password_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSmppuserExist("citrixadc_smppuser.tf_smppuser_ephem", nil),
					resource.TestCheckResourceAttr("citrixadc_smppuser.tf_smppuser_ephem", "username", "tf_smppuser_ephem"),
				),
			},
		},
	})
}

func TestAccSmppuser_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckSmppuserDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccSmppuser_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSmppuserExist("citrixadc_smppuser.tf_smppuser", nil),
				),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccSmppuser_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSmppuserExist("citrixadc_smppuser.tf_smppuser", nil),
				),
			},
		},
	})
}

// Test ephemeral path: using password_wo (WriteOnly attribute) with version tracker
const testAccSmppuser_password_wo_step1 = `
	variable "smppuser_password_wo" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_smppuser" "tf_smppuser_ephem" {
		username          = "tf_smppuser_ephem"
		password_wo       = var.smppuser_password_wo
		password_wo_version = 1
	}
`

const testAccSmppuser_password_wo_step2 = `
	variable "smppuser_password_wo_2" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_smppuser" "tf_smppuser_ephem" {
		username          = "tf_smppuser_ephem"
		password_wo       = var.smppuser_password_wo_2
		password_wo_version = 2
	}
`

func TestAccSmppuser_password_wo_ephemeral(t *testing.T) {
	t.Setenv("TF_VAR_smppuser_password_wo", "ephem_smpppass1")
	t.Setenv("TF_VAR_smppuser_password_wo_2", "ephem_smpppass2")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSmppuserDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSmppuser_password_wo_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSmppuserExist("citrixadc_smppuser.tf_smppuser_ephem", nil),
					resource.TestCheckResourceAttr("citrixadc_smppuser.tf_smppuser_ephem", "username", "tf_smppuser_ephem"),
					resource.TestCheckResourceAttr("citrixadc_smppuser.tf_smppuser_ephem", "password_wo_version", "1"),
				),
			},
			{
				Config: testAccSmppuser_password_wo_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSmppuserExist("citrixadc_smppuser.tf_smppuser_ephem", nil),
					resource.TestCheckResourceAttr("citrixadc_smppuser.tf_smppuser_ephem", "username", "tf_smppuser_ephem"),
					resource.TestCheckResourceAttr("citrixadc_smppuser.tf_smppuser_ephem", "password_wo_version", "2"),
				),
			},
		},
	})
}
