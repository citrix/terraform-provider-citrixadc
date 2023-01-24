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
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"testing"
)

const testAccAuthenticationsamlidppolicy_add = `
	resource "citrixadc_sslcertkey" "tf_sslcertkey" {
		certkey = "tf_sslcertkey"
		cert    = "/var/tmp/certificate1.crt"
		key     = "/var/tmp/key1.pem"
	}
	resource "citrixadc_authenticationsamlidpprofile" "tf_samlidpprofile" {
		name                        = "tf_samlidpprofile"
		samlspcertname              = citrixadc_sslcertkey.tf_sslcertkey.certkey
		assertionconsumerserviceurl = "http://www.example.com"
		sendpassword                = "OFF"
		samlissuername              = "new_user"
		rejectunsignedrequests      = "ON"
		signaturealg                = "RSA-SHA1"
		digestmethod                = "SHA1"
		nameidformat                = "Unspecified"
	}
	resource "citrixadc_authenticationsamlidppolicy" "tf_samlidppolicy" {
		name    = "tf_samlidppolicy"
		rule    = "false"
		action  = citrixadc_authenticationsamlidpprofile.tf_samlidpprofile.name
		comment = "aSimpleTesting"
	}
`
const testAccAuthenticationsamlidppolicy_update = `
	resource "citrixadc_sslcertkey" "tf_sslcertkey" {
		certkey = "tf_sslcertkey"
		cert    = "/var/tmp/certificate1.crt"
		key     = "/var/tmp/key1.pem"
	}
	resource "citrixadc_authenticationsamlidpprofile" "tf_samlidpprofile" {
		name                        = "tf_samlidpprofile"
		samlspcertname              = citrixadc_sslcertkey.tf_sslcertkey.certkey
		assertionconsumerserviceurl = "http://www.example.com"
		sendpassword                = "OFF"
		samlissuername              = "new_user"
		rejectunsignedrequests      = "ON"
		signaturealg                = "RSA-SHA1"
		digestmethod                = "SHA1"
		nameidformat                = "Unspecified"
	}
	resource "citrixadc_authenticationsamlidppolicy" "tf_samlidppolicy" {
		name    = "tf_samlidppolicy"
		rule    = "true"
		action  = citrixadc_authenticationsamlidpprofile.tf_samlidpprofile.name
		comment = "aSimpleTesting1"
	}
`

func TestAccAuthenticationsamlidppolicy_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAuthenticationsamlidppolicyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccAuthenticationsamlidppolicy_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationsamlidppolicyExist("citrixadc_authenticationsamlidppolicy.tf_samlidppolicy", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationsamlidppolicy.tf_samlidppolicy", "name", "tf_samlidppolicy"),
					resource.TestCheckResourceAttr("citrixadc_authenticationsamlidppolicy.tf_samlidppolicy", "rule", "false"),
					resource.TestCheckResourceAttr("citrixadc_authenticationsamlidppolicy.tf_samlidppolicy", "comment", "aSimpleTesting"),
				),
			},
			resource.TestStep{
				Config: testAccAuthenticationsamlidppolicy_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationsamlidppolicyExist("citrixadc_authenticationsamlidppolicy.tf_samlidppolicy", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationsamlidppolicy.tf_samlidppolicy", "name", "tf_samlidppolicy"),
					resource.TestCheckResourceAttr("citrixadc_authenticationsamlidppolicy.tf_samlidppolicy", "rule", "true"),
					resource.TestCheckResourceAttr("citrixadc_authenticationsamlidppolicy.tf_samlidppolicy", "comment", "aSimpleTesting1"),
				),
			},
		},
	})
}

func testAccCheckAuthenticationsamlidppolicyExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No authenticationsamlidppolicy name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(service.Authenticationsamlidppolicy.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("authenticationsamlidppolicy %s not found", n)
		}

		return nil
	}
}

func testAccCheckAuthenticationsamlidppolicyDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_authenticationsamlidppolicy" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Authenticationsamlidppolicy.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("authenticationsamlidppolicy %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
