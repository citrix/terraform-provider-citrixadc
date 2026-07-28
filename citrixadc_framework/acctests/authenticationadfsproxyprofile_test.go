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

const testAccAuthenticationadfsproxyprofile_basic_step1 = `
resource "citrixadc_authenticationadfsproxyprofile" "tf_authenticationadfsproxyprofile" {
  name        = "tf_adfsproxy_profile"
  certkeyname = "TODO_PLACEHOLDER"
  serverurl   = "https://adfs.example.com"
  username    = "adfsuser"
  password    = "adfspassword123"
}

`

const testAccAuthenticationadfsproxyprofile_basic_step2 = `
resource "citrixadc_authenticationadfsproxyprofile" "tf_authenticationadfsproxyprofile" {
  name        = "tf_adfsproxy_profile"
  certkeyname = "TODO_PLACEHOLDER"
  serverurl   = "https://adfs2.example.com"
  username    = "adfsuser_updated"
  password    = "adfspassword456"
}

`

func TestAccAuthenticationadfsproxyprofile_basic(t *testing.T) {
	t.Skip("TODO: Requires review")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationadfsproxyprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationadfsproxyprofile_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationadfsproxyprofileExist("citrixadc_authenticationadfsproxyprofile.tf_authenticationadfsproxyprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationadfsproxyprofile.tf_authenticationadfsproxyprofile", "name", "tf_adfsproxy_profile"),
					resource.TestCheckResourceAttr("citrixadc_authenticationadfsproxyprofile.tf_authenticationadfsproxyprofile", "certkeyname", "TODO_PLACEHOLDER"),
					resource.TestCheckResourceAttr("citrixadc_authenticationadfsproxyprofile.tf_authenticationadfsproxyprofile", "serverurl", "https://adfs.example.com"),
					resource.TestCheckResourceAttr("citrixadc_authenticationadfsproxyprofile.tf_authenticationadfsproxyprofile", "username", "adfsuser"),
				),
			},
			{
				Config: testAccAuthenticationadfsproxyprofile_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationadfsproxyprofileExist("citrixadc_authenticationadfsproxyprofile.tf_authenticationadfsproxyprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationadfsproxyprofile.tf_authenticationadfsproxyprofile", "name", "tf_adfsproxy_profile"),
					resource.TestCheckResourceAttr("citrixadc_authenticationadfsproxyprofile.tf_authenticationadfsproxyprofile", "certkeyname", "TODO_PLACEHOLDER"),
					resource.TestCheckResourceAttr("citrixadc_authenticationadfsproxyprofile.tf_authenticationadfsproxyprofile", "serverurl", "https://adfs2.example.com"),
					resource.TestCheckResourceAttr("citrixadc_authenticationadfsproxyprofile.tf_authenticationadfsproxyprofile", "username", "adfsuser_updated"),
				),
			},
		},
	})
}

func TestAccAuthenticationadfsproxyprofile_import(t *testing.T) {
	t.Skip("TODO: Requires review")
	const resAddr = "citrixadc_authenticationadfsproxyprofile.tf_authenticationadfsproxyprofile"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationadfsproxyprofileDestroy,
		Steps: []resource.TestStep{
			{Config: testAccAuthenticationadfsproxyprofile_basic_step1},
			{
				Config:            testAccAuthenticationadfsproxyprofile_basic_step1,
				ResourceName:      resAddr,
				ImportState:       true,
				ImportStateVerify: true,
				// password is a Sensitive attribute that NITRO does not echo back,
				// and password_wo_version is a Computed default that Read does not
				// repopulate, so neither can round-trip through import.
				ImportStateVerifyIgnore: []string{"password", "password_wo_version"},
			},
		},
	})
}

func testAccCheckAuthenticationadfsproxyprofileExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No authenticationadfsproxyprofile name is set")
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
		data, err := client.FindResource(service.Authenticationadfsproxyprofile.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("authenticationadfsproxyprofile %s not found", n)
		}

		return nil
	}
}

func testAccCheckAuthenticationadfsproxyprofileDestroy(s *terraform.State) error {
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_authenticationadfsproxyprofile" {
			continue
		}
		_, err := client.FindResource(service.Authenticationadfsproxyprofile.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("authenticationadfsproxyprofile %s still exists", rs.Primary.ID)
		}
	}
	return nil
}

const testAccAuthenticationadfsproxyprofileDataSource_basic = `

resource "citrixadc_authenticationadfsproxyprofile" "tf_authenticationadfsproxyprofile" {
  name        = "tf_adfsproxy_profile_ds"
  certkeyname = "TODO_PLACEHOLDER"
  serverurl   = "https://adfs.example.com"
  username    = "adfsuser"
  password    = "adfspassword123"
}

data "citrixadc_authenticationadfsproxyprofile" "tf_authenticationadfsproxyprofile" {
  name       = citrixadc_authenticationadfsproxyprofile.tf_authenticationadfsproxyprofile.name
  depends_on = [citrixadc_authenticationadfsproxyprofile.tf_authenticationadfsproxyprofile]
}
`

func TestAccAuthenticationadfsproxyprofileDataSource_basic(t *testing.T) {
	t.Skip("TODO: Requires review")
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationadfsproxyprofileDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_authenticationadfsproxyprofile.tf_authenticationadfsproxyprofile", "name", "tf_adfsproxy_profile_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationadfsproxyprofile.tf_authenticationadfsproxyprofile", "certkeyname", "TODO_PLACEHOLDER"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationadfsproxyprofile.tf_authenticationadfsproxyprofile", "serverurl", "https://adfs.example.com"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationadfsproxyprofile.tf_authenticationadfsproxyprofile", "username", "adfsuser"),
				),
			},
		},
	})
}

// Test backward-compatible path: using password (Sensitive attribute)
const testAccAuthenticationadfsproxyprofile_password_step1 = `

	variable "authenticationadfsproxyprofile_password" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_authenticationadfsproxyprofile" "tf_authenticationadfsproxyprofile" {
		name        = "tf_adfsproxy_profile_pw"
		certkeyname = "TODO_PLACEHOLDER"
		serverurl   = "https://adfs.example.com"
		username    = "adfsuser"
		password    = var.authenticationadfsproxyprofile_password
	}
`

// Update backward-compatible path: change password value
const testAccAuthenticationadfsproxyprofile_password_step2 = `

	 variable "authenticationadfsproxyprofile_password_2" {
	  type      = string
	  sensitive = true
	}

	 resource "citrixadc_authenticationadfsproxyprofile" "tf_authenticationadfsproxyprofile" {
		name        = "tf_adfsproxy_profile_pw"
		certkeyname = "TODO_PLACEHOLDER"
		serverurl   = "https://adfs.example.com"
		username    = "adfsuser"
		password    = var.authenticationadfsproxyprofile_password_2
	}
`

func TestAccAuthenticationadfsproxyprofile_password_backward_compat(t *testing.T) {
	t.Skip("TODO: Requires review")
	t.Setenv("TF_VAR_authenticationadfsproxyprofile_password", "oldpassword123")
	t.Setenv("TF_VAR_authenticationadfsproxyprofile_password_2", "newpassword456")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationadfsproxyprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationadfsproxyprofile_password_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationadfsproxyprofileExist("citrixadc_authenticationadfsproxyprofile.tf_authenticationadfsproxyprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationadfsproxyprofile.tf_authenticationadfsproxyprofile", "name", "tf_adfsproxy_profile_pw"),
					resource.TestCheckResourceAttr("citrixadc_authenticationadfsproxyprofile.tf_authenticationadfsproxyprofile", "username", "adfsuser"),
				),
			},
			{
				Config: testAccAuthenticationadfsproxyprofile_password_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationadfsproxyprofileExist("citrixadc_authenticationadfsproxyprofile.tf_authenticationadfsproxyprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationadfsproxyprofile.tf_authenticationadfsproxyprofile", "name", "tf_adfsproxy_profile_pw"),
					resource.TestCheckResourceAttr("citrixadc_authenticationadfsproxyprofile.tf_authenticationadfsproxyprofile", "username", "adfsuser"),
				),
			},
		},
	})
}

// Test ephemeral path: using password_wo (WriteOnly attribute) with version tracker
const testAccAuthenticationadfsproxyprofile_password_wo_step1 = `

	variable "authenticationadfsproxyprofile_password_wo" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_authenticationadfsproxyprofile" "tf_authenticationadfsproxyprofile" {
		name                = "tf_adfsproxy_profile_pw_wo"
		certkeyname         = "TODO_PLACEHOLDER"
		serverurl           = "https://adfs.example.com"
		username            = "adfsuser"
		password_wo         = var.authenticationadfsproxyprofile_password_wo
		password_wo_version = 1
	}
`

// Update ephemeral path: bump version to trigger update with new password
const testAccAuthenticationadfsproxyprofile_password_wo_step2 = `

	 variable "authenticationadfsproxyprofile_password_wo_2" {
	  type      = string
	  sensitive = true
	}

	 resource "citrixadc_authenticationadfsproxyprofile" "tf_authenticationadfsproxyprofile" {
		name                = "tf_adfsproxy_profile_pw_wo"
		certkeyname         = "TODO_PLACEHOLDER"
		serverurl           = "https://adfs.example.com"
		username            = "adfsuser"
		password_wo         = var.authenticationadfsproxyprofile_password_wo_2
		password_wo_version = 2
	}
`

func TestAccAuthenticationadfsproxyprofile_password_wo_ephemeral(t *testing.T) {
	t.Skip("TODO: Requires review")
	t.Setenv("TF_VAR_authenticationadfsproxyprofile_password_wo", "ephemeral_pass1")
	t.Setenv("TF_VAR_authenticationadfsproxyprofile_password_wo_2", "ephemeral_pass2")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationadfsproxyprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationadfsproxyprofile_password_wo_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationadfsproxyprofileExist("citrixadc_authenticationadfsproxyprofile.tf_authenticationadfsproxyprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationadfsproxyprofile.tf_authenticationadfsproxyprofile", "name", "tf_adfsproxy_profile_pw_wo"),
					resource.TestCheckResourceAttr("citrixadc_authenticationadfsproxyprofile.tf_authenticationadfsproxyprofile", "username", "adfsuser"),
					resource.TestCheckResourceAttr("citrixadc_authenticationadfsproxyprofile.tf_authenticationadfsproxyprofile", "password_wo_version", "1"),
				),
			},
			{
				Config: testAccAuthenticationadfsproxyprofile_password_wo_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationadfsproxyprofileExist("citrixadc_authenticationadfsproxyprofile.tf_authenticationadfsproxyprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationadfsproxyprofile.tf_authenticationadfsproxyprofile", "name", "tf_adfsproxy_profile_pw_wo"),
					resource.TestCheckResourceAttr("citrixadc_authenticationadfsproxyprofile.tf_authenticationadfsproxyprofile", "username", "adfsuser"),
					resource.TestCheckResourceAttr("citrixadc_authenticationadfsproxyprofile.tf_authenticationadfsproxyprofile", "password_wo_version", "2"),
				),
			},
		},
	})
}
