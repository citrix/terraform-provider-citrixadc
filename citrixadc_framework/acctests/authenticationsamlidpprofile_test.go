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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"testing"
)

const testAccAuthenticationsamlidpprofile_add = `
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
`
const testAccAuthenticationsamlidpprofile_update = `
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
		rejectunsignedrequests      = "OFF"
		signaturealg                = "RSA-SHA1"
		digestmethod                = "SHA256"
		nameidformat                = "Unspecified"
	}
`

func TestAccAuthenticationsamlidpprofile_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doSslPrecheckforsamlidpprofile(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationsamlidpprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationsamlidpprofile_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationsamlidpprofileExist("citrixadc_authenticationsamlidpprofile.tf_samlidpprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationsamlidpprofile.tf_samlidpprofile", "name", "tf_samlidpprofile"),
					resource.TestCheckResourceAttr("citrixadc_authenticationsamlidpprofile.tf_samlidpprofile", "digestmethod", "SHA1"),
				),
			},
			{
				Config: testAccAuthenticationsamlidpprofile_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationsamlidpprofileExist("citrixadc_authenticationsamlidpprofile.tf_samlidpprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationsamlidpprofile.tf_samlidpprofile", "name", "tf_samlidpprofile"),
					resource.TestCheckResourceAttr("citrixadc_authenticationsamlidpprofile.tf_samlidpprofile", "digestmethod", "SHA256"),
				),
			},
		},
	})
}

func testAccCheckAuthenticationsamlidpprofileExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No authenticationsamlidpprofile name is set")
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
		data, err := client.FindResource(service.Authenticationsamlidpprofile.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("authenticationsamlidpprofile %s not found", n)
		}

		return nil
	}
}
func doSslPrecheckforsamlidpprofile(t *testing.T) {
	testAccPreCheck(t)

	uploads := []string{
		"certificate1.crt",
		"key1.pem",
	}

	c, err := testHelperInstantiateClient("", "", "", false)
	if err != nil {
		t.Fatalf("Failed to instantiate client. %v\n", err)
	}

	//c := testAccProvider.Meta().(*NetScalerNitroClient)
	for _, filename := range uploads {
		err := uploadTestdataFile(c, t, filename, "/var/tmp")
		if err != nil {
			t.Errorf("%v", err)
		}
	}
}

func testAccCheckAuthenticationsamlidpprofileDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_authenticationsamlidpprofile" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Authenticationsamlidpprofile.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("authenticationsamlidpprofile %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
