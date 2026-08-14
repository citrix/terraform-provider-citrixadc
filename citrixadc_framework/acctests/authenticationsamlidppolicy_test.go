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
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
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
		PreCheck:                 func() { doSslcertkeyPreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationsamlidppolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationsamlidppolicy_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationsamlidppolicyExist("citrixadc_authenticationsamlidppolicy.tf_samlidppolicy", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationsamlidppolicy.tf_samlidppolicy", "name", "tf_samlidppolicy"),
					resource.TestCheckResourceAttr("citrixadc_authenticationsamlidppolicy.tf_samlidppolicy", "rule", "false"),
					resource.TestCheckResourceAttr("citrixadc_authenticationsamlidppolicy.tf_samlidppolicy", "comment", "aSimpleTesting"),
				),
			},
			{
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

		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Authenticationsamlidppolicy.Type(), rs.Primary.ID)

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
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_authenticationsamlidppolicy" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Authenticationsamlidppolicy.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("authenticationsamlidppolicy %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccAuthenticationsamlidppolicy_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_authenticationsamlidppolicy.tf_samlidppolicy"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doSslcertkeyPreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationsamlidppolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationsamlidppolicy_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAuthenticationsamlidppolicyExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Authenticationsamlidppolicy.Type(), "tf_samlidppolicy"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccAuthenticationsamlidppolicy_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAuthenticationsamlidppolicyExist(resAddr, nil)),
			},
		},
	})
}

func TestAccAuthenticationsamlidppolicy_import(t *testing.T) {
	const resAddr = "citrixadc_authenticationsamlidppolicy.tf_samlidppolicy"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doSslcertkeyPreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationsamlidppolicyDestroy,
		Steps: []resource.TestStep{
			{Config: testAccAuthenticationsamlidppolicy_add},
			{
				Config:                  testAccAuthenticationsamlidppolicy_add,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

const testAccAuthenticationsamlidppolicyDataSource_basic = `
	resource "citrixadc_sslcertkey" "tf_sslcertkey_ds" {
		certkey = "tf_sslcertkey_ds"
		cert    = "/var/tmp/certificate1.crt"
		key     = "/var/tmp/key1.pem"
	}
	resource "citrixadc_authenticationsamlidpprofile" "tf_samlidpprofile_ds" {
		name                        = "tf_samlidpprofile_ds"
		samlspcertname              = citrixadc_sslcertkey.tf_sslcertkey_ds.certkey
		assertionconsumerserviceurl = "http://www.example.com"
		sendpassword                = "OFF"
		samlissuername              = "new_user"
		rejectunsignedrequests      = "ON"
		signaturealg                = "RSA-SHA1"
		digestmethod                = "SHA1"
		nameidformat                = "Unspecified"
	}
	resource "citrixadc_authenticationsamlidppolicy" "tf_samlidppolicy_ds" {
		name    = "tf_samlidppolicy_ds"
		rule    = "false"
		action  = citrixadc_authenticationsamlidpprofile.tf_samlidpprofile_ds.name
		comment = "DataSource Test"
	}
	data "citrixadc_authenticationsamlidppolicy" "tf_samlidppolicy_ds" {
		name = citrixadc_authenticationsamlidppolicy.tf_samlidppolicy_ds.name
	}
`

func TestAccAuthenticationsamlidppolicy_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckAuthenticationsamlidppolicyDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccAuthenticationsamlidppolicy_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAuthenticationsamlidppolicyExist("citrixadc_authenticationsamlidppolicy.tf_samlidppolicy", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccAuthenticationsamlidppolicy_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAuthenticationsamlidppolicyExist("citrixadc_authenticationsamlidppolicy.tf_samlidppolicy", nil)),
			},
		},
	})
}

// testAccAuthenticationsamlidppolicy_unset_step1 sets the unset-eligible
// attribute (comment) to a valid non-default value.
const testAccAuthenticationsamlidppolicy_unset_step1 = `
	resource "citrixadc_sslcertkey" "tf_sslcertkey_unset" {
		certkey = "tf_sslcertkey_unset"
		cert    = "/var/tmp/certificate1.crt"
		key     = "/var/tmp/key1.pem"
	}
	resource "citrixadc_authenticationsamlidpprofile" "tf_samlidpprofile_unset" {
		name                        = "tf_samlidpprofile_unset"
		samlspcertname              = citrixadc_sslcertkey.tf_sslcertkey_unset.certkey
		assertionconsumerserviceurl = "http://www.example.com"
		sendpassword                = "OFF"
		samlissuername              = "new_user"
		rejectunsignedrequests      = "ON"
		signaturealg                = "RSA-SHA1"
		digestmethod                = "SHA1"
		nameidformat                = "Unspecified"
	}
	resource "citrixadc_authenticationsamlidppolicy" "tf_unset" {
		name    = "tf_test_samlidppolicy_unset"
		rule    = "true"
		action  = citrixadc_authenticationsamlidpprofile.tf_samlidpprofile_unset.name
		comment = "aSimpleTesting"
	}
`

// testAccAuthenticationsamlidppolicy_unset_step2 removes comment from the policy
// (keeping only the key + required attributes). The provider must unset it so the
// appliance reverts it to its default (empty).
const testAccAuthenticationsamlidppolicy_unset_step2 = `
	resource "citrixadc_sslcertkey" "tf_sslcertkey_unset" {
		certkey = "tf_sslcertkey_unset"
		cert    = "/var/tmp/certificate1.crt"
		key     = "/var/tmp/key1.pem"
	}
	resource "citrixadc_authenticationsamlidpprofile" "tf_samlidpprofile_unset" {
		name                        = "tf_samlidpprofile_unset"
		samlspcertname              = citrixadc_sslcertkey.tf_sslcertkey_unset.certkey
		assertionconsumerserviceurl = "http://www.example.com"
		sendpassword                = "OFF"
		samlissuername              = "new_user"
		rejectunsignedrequests      = "ON"
		signaturealg                = "RSA-SHA1"
		digestmethod                = "SHA1"
		nameidformat                = "Unspecified"
	}
	resource "citrixadc_authenticationsamlidppolicy" "tf_unset" {
		name   = "tf_test_samlidppolicy_unset"
		rule   = "true"
		action = citrixadc_authenticationsamlidpprofile.tf_samlidpprofile_unset.name
	}
`

func TestAccAuthenticationsamlidppolicy_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doSslcertkeyPreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationsamlidppolicyDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccAuthenticationsamlidppolicy_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationsamlidppolicyExist("citrixadc_authenticationsamlidppolicy.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationsamlidppolicy.tf_unset", "comment", "aSimpleTesting"),
					testAccCheckAuthenticationsamlidppolicyADCValue("tf_test_samlidppolicy_unset", "comment", "aSimpleTesting"),
				),
			},
			{
				// Removing the attributes must unset them: the appliance reverts
				// them to their defaults (empty), and the implicit post-apply plan
				// must be empty.
				Config: testAccAuthenticationsamlidppolicy_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationsamlidppolicyExist("citrixadc_authenticationsamlidppolicy.tf_unset", nil),
					testAccCheckAuthenticationsamlidppolicyADCValue("tf_test_samlidppolicy_unset", "comment", ""),
				),
			},
		},
	})
}

// testAccCheckAuthenticationsamlidppolicyADCValue asserts an attribute's value
// directly on the appliance (not just in Terraform state), proving the unset
// actually reverted it. An absent attribute is treated as the empty string.
func testAccCheckAuthenticationsamlidppolicyADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Authenticationsamlidppolicy.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("authenticationsamlidppolicy %s not found on appliance", name)
		}
		got := ""
		if v, ok := data[attr]; ok && v != nil {
			got = strings.TrimSpace(fmt.Sprintf("%v", v))
		}
		if got != want {
			return fmt.Errorf("authenticationsamlidppolicy %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func TestAccAuthenticationsamlidppolicyDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doSslcertkeyPreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationsamlidppolicyDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_authenticationsamlidppolicy.tf_samlidppolicy_ds", "name", "tf_samlidppolicy_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationsamlidppolicy.tf_samlidppolicy_ds", "rule", "false"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationsamlidppolicy.tf_samlidppolicy_ds", "comment", "DataSource Test"),
					resource.TestCheckResourceAttrSet("data.citrixadc_authenticationsamlidppolicy.tf_samlidppolicy_ds", "id"),
				),
			},
		},
	})
}
