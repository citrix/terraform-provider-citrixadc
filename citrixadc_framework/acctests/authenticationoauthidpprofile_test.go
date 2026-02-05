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

const testAccAuthenticationoauthidpprofile_add = `

	resource "citrixadc_authenticationoauthidpprofile" "tf_idpprofile" {
		name         = "tf_idpprofile"
		clientid     = "cliId"
		clientsecret = "secret"
		redirecturl  = "http://www.example.com/1/"
	}
`
const testAccAuthenticationoauthidpprofile_update = `

	resource "citrixadc_authenticationoauthidpprofile" "tf_idpprofile" {
		name          = "tf_idpprofile"
		clientid      = "cliId1"
		clientsecret  = "secret"
		redirecturl   = "http://www.example11.com/1/"
	}
`

func TestAccAuthenticationoauthidpprofile_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationoauthidpprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationoauthidpprofile_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationoauthidpprofileExist("citrixadc_authenticationoauthidpprofile.tf_idpprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationoauthidpprofile.tf_idpprofile", "name", "tf_idpprofile"),
					resource.TestCheckResourceAttr("citrixadc_authenticationoauthidpprofile.tf_idpprofile", "clientid", "cliId"),
				),
			},
			{
				Config: testAccAuthenticationoauthidpprofile_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationoauthidpprofileExist("citrixadc_authenticationoauthidpprofile.tf_idpprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationoauthidpprofile.tf_idpprofile", "name", "tf_idpprofile"),
					resource.TestCheckResourceAttr("citrixadc_authenticationoauthidpprofile.tf_idpprofile", "clientid", "cliId1"),
				),
			},
		},
	})
}

func testAccCheckAuthenticationoauthidpprofileExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No authenticationoauthidpprofile name is set")
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
		data, err := client.FindResource("authenticationoauthidpprofile", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("authenticationoauthidpprofile %s not found", n)
		}

		return nil
	}
}

func testAccCheckAuthenticationoauthidpprofileDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_authenticationoauthidpprofile" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("authenticationoauthidpprofile", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("authenticationoauthidpprofile %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccAuthenticationoauthidpprofileDataSource_basic = `

	resource "citrixadc_authenticationoauthidpprofile" "tf_idpprofile_ds" {
		name         = "tf_idpprofile_ds"
		clientid     = "cliId_datasource"
		clientsecret = "secret"
		redirecturl  = "http://www.example.com/datasource/"
	}

	data "citrixadc_authenticationoauthidpprofile" "tf_idpprofile_ds" {
		name = citrixadc_authenticationoauthidpprofile.tf_idpprofile_ds.name
	}
`

func TestAccAuthenticationoauthidpprofileDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationoauthidpprofileDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_authenticationoauthidpprofile.tf_idpprofile_ds", "name", "tf_idpprofile_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationoauthidpprofile.tf_idpprofile_ds", "clientid", "cliId_datasource"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationoauthidpprofile.tf_idpprofile_ds", "redirecturl", "http://www.example.com/datasource/"),
				),
			},
		},
	})
}
