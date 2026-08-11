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
	"strings"
	"testing"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
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

func TestAccAuthenticationsamlidpprofile_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_authenticationsamlidpprofile.tf_samlidpprofile"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doSslPrecheckforsamlidpprofile(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationsamlidpprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationsamlidpprofile_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAuthenticationsamlidpprofileExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Authenticationsamlidpprofile.Type(), "tf_samlidpprofile"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccAuthenticationsamlidpprofile_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAuthenticationsamlidpprofileExist(resAddr, nil)),
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

func TestAccAuthenticationsamlidpprofile_import(t *testing.T) {
	const resAddr = "citrixadc_authenticationsamlidpprofile.tf_samlidpprofile"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doSslPrecheckforsamlidpprofile(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationsamlidpprofileDestroy,
		Steps: []resource.TestStep{
			{Config: testAccAuthenticationsamlidpprofile_add},
			{
				Config:                  testAccAuthenticationsamlidpprofile_add,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"sendpassword"},
			},
		},
	})
}

const testAccAuthenticationsamlidpprofileDataSource_basic = `
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

	data "citrixadc_authenticationsamlidpprofile" "tf_samlidpprofile" {
		name = citrixadc_authenticationsamlidpprofile.tf_samlidpprofile.name
		depends_on = [citrixadc_authenticationsamlidpprofile.tf_samlidpprofile]
	}
`

func TestAccAuthenticationsamlidpprofile_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { doSslPrecheckforsamlidpprofile(t) },
		CheckDestroy: testAccCheckAuthenticationsamlidpprofileDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccAuthenticationsamlidpprofile_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAuthenticationsamlidpprofileExist("citrixadc_authenticationsamlidpprofile.tf_samlidpprofile", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccAuthenticationsamlidpprofile_add,
				Check:                    resource.ComposeTestCheckFunc(testAccCheckAuthenticationsamlidpprofileExist("citrixadc_authenticationsamlidpprofile.tf_samlidpprofile", nil)),
			},
		},
	})
}

// The unset test exercises the mutable attributes that have a documented NITRO
// server default. Step 1 sets each to a valid non-default value; step 2 removes
// them from config, and the provider must unset them so the appliance reverts
// them to their documented defaults (with an implicit empty post-apply plan).
const testAccAuthenticationsamlidpprofile_unset_step1 = `
resource "citrixadc_authenticationsamlidpprofile" "tf_unset" {
  name                    = "tf_test_samlidpprofile_unset"
  rejectunsignedrequests  = "OFF"
  signaturealg            = "RSA-SHA1"
  digestmethod            = "SHA1"
  nameidformat            = "persistent"
  samlbinding             = "ARTIFACT"
  skewtime                = 10
  signassertion           = "BOTH"
  logoutbinding           = "REDIRECT"
}
`

const testAccAuthenticationsamlidpprofile_unset_step2 = `
resource "citrixadc_authenticationsamlidpprofile" "tf_unset" {
  name = "tf_test_samlidpprofile_unset"
  # All unset-eligible attributes removed from config -> the provider must
  # unset them (revert to NITRO defaults).
}
`

func TestAccAuthenticationsamlidpprofile_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationsamlidpprofileDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccAuthenticationsamlidpprofile_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationsamlidpprofileExist("citrixadc_authenticationsamlidpprofile.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationsamlidpprofile.tf_unset", "rejectunsignedrequests", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_authenticationsamlidpprofile.tf_unset", "signaturealg", "RSA-SHA1"),
					resource.TestCheckResourceAttr("citrixadc_authenticationsamlidpprofile.tf_unset", "digestmethod", "SHA1"),
					resource.TestCheckResourceAttr("citrixadc_authenticationsamlidpprofile.tf_unset", "nameidformat", "persistent"),
					resource.TestCheckResourceAttr("citrixadc_authenticationsamlidpprofile.tf_unset", "samlbinding", "ARTIFACT"),
					resource.TestCheckResourceAttr("citrixadc_authenticationsamlidpprofile.tf_unset", "skewtime", "10"),
					resource.TestCheckResourceAttr("citrixadc_authenticationsamlidpprofile.tf_unset", "signassertion", "BOTH"),
					resource.TestCheckResourceAttr("citrixadc_authenticationsamlidpprofile.tf_unset", "logoutbinding", "REDIRECT"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from
				// the appliance) reverts to the documented NITRO defaults, and the
				// implicit post-apply plan must be empty.
				Config: testAccAuthenticationsamlidpprofile_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationsamlidpprofileExist("citrixadc_authenticationsamlidpprofile.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationsamlidpprofile.tf_unset", "rejectunsignedrequests", "ON"),
					resource.TestCheckResourceAttr("citrixadc_authenticationsamlidpprofile.tf_unset", "signaturealg", "RSA-SHA256"),
					resource.TestCheckResourceAttr("citrixadc_authenticationsamlidpprofile.tf_unset", "digestmethod", "SHA256"),
					resource.TestCheckResourceAttr("citrixadc_authenticationsamlidpprofile.tf_unset", "nameidformat", "transient"),
					resource.TestCheckResourceAttr("citrixadc_authenticationsamlidpprofile.tf_unset", "samlbinding", "POST"),
					resource.TestCheckResourceAttr("citrixadc_authenticationsamlidpprofile.tf_unset", "skewtime", "5"),
					resource.TestCheckResourceAttr("citrixadc_authenticationsamlidpprofile.tf_unset", "signassertion", "ASSERTION"),
					resource.TestCheckResourceAttr("citrixadc_authenticationsamlidpprofile.tf_unset", "logoutbinding", "POST"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckAuthenticationsamlidpprofileADCValue("tf_test_samlidpprofile_unset", "rejectunsignedrequests", "ON"),
					testAccCheckAuthenticationsamlidpprofileADCValue("tf_test_samlidpprofile_unset", "digestmethod", "SHA256"),
					testAccCheckAuthenticationsamlidpprofileADCValue("tf_test_samlidpprofile_unset", "samlbinding", "POST"),
					testAccCheckAuthenticationsamlidpprofileADCValue("tf_test_samlidpprofile_unset", "signassertion", "ASSERTION"),
				),
			},
		},
	})
}

// testAccCheckAuthenticationsamlidpprofileADCValue asserts an attribute's value
// directly on the appliance (not just in Terraform state), proving the unset
// actually reverted it.
func testAccCheckAuthenticationsamlidpprofileADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Authenticationsamlidpprofile.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("authenticationsamlidpprofile %s not found on appliance", name)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("authenticationsamlidpprofile %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func TestAccAuthenticationsamlidpprofileDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doSslPrecheckforsamlidpprofile(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationsamlidpprofileDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_authenticationsamlidpprofile.tf_samlidpprofile", "name", "tf_samlidpprofile"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationsamlidpprofile.tf_samlidpprofile", "digestmethod", "SHA1"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationsamlidpprofile.tf_samlidpprofile", "rejectunsignedrequests", "ON"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationsamlidpprofile.tf_samlidpprofile", "signaturealg", "RSA-SHA1"),
				),
			},
		},
	})
}
