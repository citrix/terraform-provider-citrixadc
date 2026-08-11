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

const testAccAuthenticationoauthidppolicy_add = `
	resource "citrixadc_authenticationoauthidpprofile" "tf_idpprofile" {
		name         = "tf_idpprofile"
		clientid     = "cliId"
		clientsecret = "secret"
		redirecturl  = "http://www.example.com/1/"
	}
	resource "citrixadc_authenticationoauthidppolicy" "tf_idppolicy" {
		name    = "tf_idppolicy"
		rule    = "true"
		action  = citrixadc_authenticationoauthidpprofile.tf_idpprofile.name
		comment = "add_policy"
	}
`
const testAccAuthenticationoauthidppolicy_update = `
	resource "citrixadc_authenticationoauthidpprofile" "tf_idpprofile" {
		name         = "tf_idpprofile"
		clientid     = "cliId"
		clientsecret = "secret"
		redirecturl  = "http://www.example.com/1/"
	}
	resource "citrixadc_authenticationoauthidppolicy" "tf_idppolicy" {
		name    = "tf_idppolicy"
		rule    = "false"
		action  = citrixadc_authenticationoauthidpprofile.tf_idpprofile.name
		comment = "update_policy"
	}
`

func TestAccAuthenticationoauthidppolicy_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationoauthidppolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationoauthidppolicy_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationoauthidppolicyExist("citrixadc_authenticationoauthidppolicy.tf_idppolicy", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationoauthidppolicy.tf_idppolicy", "name", "tf_idppolicy"),
					resource.TestCheckResourceAttr("citrixadc_authenticationoauthidppolicy.tf_idppolicy", "rule", "true"),
					resource.TestCheckResourceAttr("citrixadc_authenticationoauthidppolicy.tf_idppolicy", "comment", "add_policy"),
				),
			},
			{
				Config: testAccAuthenticationoauthidppolicy_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationoauthidppolicyExist("citrixadc_authenticationoauthidppolicy.tf_idppolicy", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationoauthidppolicy.tf_idppolicy", "name", "tf_idppolicy"),
					resource.TestCheckResourceAttr("citrixadc_authenticationoauthidppolicy.tf_idppolicy", "rule", "false"),
					resource.TestCheckResourceAttr("citrixadc_authenticationoauthidppolicy.tf_idppolicy", "comment", "update_policy"),
				),
			},
		},
	})
}

func testAccCheckAuthenticationoauthidppolicyExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No authenticationoauthidppolicy name is set")
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
		data, err := client.FindResource("authenticationoauthidppolicy", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("authenticationoauthidppolicy %s not found", n)
		}

		return nil
	}
}

func testAccCheckAuthenticationoauthidppolicyDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_authenticationoauthidppolicy" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("authenticationoauthidppolicy", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("authenticationoauthidppolicy %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccAuthenticationoauthidppolicyDataSource_basic = `
	resource "citrixadc_authenticationoauthidpprofile" "tf_idpprofile" {
		name         = "tf_idpprofile_ds"
		clientid     = "cliId"
		clientsecret = "secret"
		redirecturl  = "http://www.example.com/1/"
	}
	resource "citrixadc_authenticationoauthidppolicy" "tf_idppolicy" {
		name    = "tf_idppolicy_ds"
		rule    = "true"
		action  = citrixadc_authenticationoauthidpprofile.tf_idpprofile.name
		comment = "datasource_test"
	}
	data "citrixadc_authenticationoauthidppolicy" "tf_idppolicy_ds" {
		name = citrixadc_authenticationoauthidppolicy.tf_idppolicy.name
	}
`

func TestAccAuthenticationoauthidppolicy_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_authenticationoauthidppolicy.tf_idppolicy"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationoauthidppolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationoauthidppolicy_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAuthenticationoauthidppolicyExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Authenticationoauthidppolicy.Type(), "tf_idppolicy"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccAuthenticationoauthidppolicy_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAuthenticationoauthidppolicyExist(resAddr, nil)),
			},
		},
	})
}

func TestAccAuthenticationoauthidppolicy_import(t *testing.T) {
	const resAddr = "citrixadc_authenticationoauthidppolicy.tf_idppolicy"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationoauthidppolicyDestroy,
		Steps: []resource.TestStep{
			{Config: testAccAuthenticationoauthidppolicy_add},
			{
				Config:                  testAccAuthenticationoauthidppolicy_add,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccAuthenticationoauthidppolicy_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckAuthenticationoauthidppolicyDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccAuthenticationoauthidppolicy_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAuthenticationoauthidppolicyExist("citrixadc_authenticationoauthidppolicy.tf_idppolicy", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccAuthenticationoauthidppolicy_add,
				Check:                    resource.ComposeTestCheckFunc(testAccCheckAuthenticationoauthidppolicyExist("citrixadc_authenticationoauthidppolicy.tf_idppolicy", nil)),
			},
		},
	})
}

// testAccAuthenticationoauthidppolicy_unset_step1 sets every unset-eligible
// attribute (comment, logaction, undefaction) to a valid non-default value.
// logaction requires an existing messagelog (audit message) action, created here.
const testAccAuthenticationoauthidppolicy_unset_step1 = `
	resource "citrixadc_auditmessageaction" "tf_unset_msgact" {
		name              = "tf_unset_msgact"
		loglevel          = "INFORMATIONAL"
		stringbuilderexpr = "\"oauthidppolicy unset test\""
	}
	resource "citrixadc_authenticationoauthidpprofile" "tf_unset_idpprofile" {
		name         = "tf_unset_idpprofile"
		clientid     = "cliId"
		clientsecret = "secret"
		redirecturl  = "http://www.example.com/1/"
	}
	resource "citrixadc_authenticationoauthidppolicy" "tf_unset" {
		name        = "tf_unset_idppolicy"
		rule        = "true"
		action      = citrixadc_authenticationoauthidpprofile.tf_unset_idpprofile.name
		comment     = "unset_me"
		logaction   = citrixadc_auditmessageaction.tf_unset_msgact.name
		undefaction = "RESET"
	}
`

// testAccAuthenticationoauthidppolicy_unset_step2 removes all unset-eligible
// attributes (only key + required attrs remain), so the provider must unset them
// on the appliance (revert to NITRO defaults, i.e. empty/absent).
const testAccAuthenticationoauthidppolicy_unset_step2 = `
	resource "citrixadc_auditmessageaction" "tf_unset_msgact" {
		name              = "tf_unset_msgact"
		loglevel          = "INFORMATIONAL"
		stringbuilderexpr = "\"oauthidppolicy unset test\""
	}
	resource "citrixadc_authenticationoauthidpprofile" "tf_unset_idpprofile" {
		name         = "tf_unset_idpprofile"
		clientid     = "cliId"
		clientsecret = "secret"
		redirecturl  = "http://www.example.com/1/"
	}
	resource "citrixadc_authenticationoauthidppolicy" "tf_unset" {
		name   = "tf_unset_idppolicy"
		rule   = "true"
		action = citrixadc_authenticationoauthidpprofile.tf_unset_idpprofile.name
	}
`

func TestAccAuthenticationoauthidppolicy_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationoauthidppolicyDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccAuthenticationoauthidppolicy_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationoauthidppolicyExist("citrixadc_authenticationoauthidppolicy.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationoauthidppolicy.tf_unset", "comment", "unset_me"),
					resource.TestCheckResourceAttr("citrixadc_authenticationoauthidppolicy.tf_unset", "logaction", "tf_unset_msgact"),
					resource.TestCheckResourceAttr("citrixadc_authenticationoauthidppolicy.tf_unset", "undefaction", "RESET"),
					testAccCheckAuthenticationoauthidppolicyADCValue("tf_unset_idppolicy", "comment", "unset_me"),
					testAccCheckAuthenticationoauthidppolicyADCValue("tf_unset_idppolicy", "logaction", "tf_unset_msgact"),
					testAccCheckAuthenticationoauthidppolicyADCValue("tf_unset_idppolicy", "undefaction", "RESET"),
				),
			},
			{
				// Removing the attributes must unset them: the appliance reverts
				// them to their defaults (empty/absent), and the implicit post-apply
				// plan must be empty.
				Config: testAccAuthenticationoauthidppolicy_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationoauthidppolicyExist("citrixadc_authenticationoauthidppolicy.tf_unset", nil),
					testAccCheckAuthenticationoauthidppolicyADCValue("tf_unset_idppolicy", "comment", ""),
					testAccCheckAuthenticationoauthidppolicyADCValue("tf_unset_idppolicy", "logaction", ""),
					testAccCheckAuthenticationoauthidppolicyADCValue("tf_unset_idppolicy", "undefaction", ""),
				),
			},
		},
	})
}

// testAccCheckAuthenticationoauthidppolicyADCValue asserts an attribute's value
// directly on the appliance (not just in Terraform state), proving the unset
// actually reverted it. An absent attribute is treated as the empty string.
func testAccCheckAuthenticationoauthidppolicyADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Authenticationoauthidppolicy.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("authenticationoauthidppolicy %s not found on appliance", name)
		}
		got := ""
		if v, ok := data[attr]; ok && v != nil {
			got = strings.TrimSpace(fmt.Sprintf("%v", v))
		}
		if got != want {
			return fmt.Errorf("authenticationoauthidppolicy %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func TestAccAuthenticationoauthidppolicyDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationoauthidppolicyDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_authenticationoauthidppolicy.tf_idppolicy_ds", "name", "tf_idppolicy_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationoauthidppolicy.tf_idppolicy_ds", "rule", "true"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationoauthidppolicy.tf_idppolicy_ds", "comment", "datasource_test"),
					resource.TestCheckResourceAttrPair("data.citrixadc_authenticationoauthidppolicy.tf_idppolicy_ds", "action", "citrixadc_authenticationoauthidppolicy.tf_idppolicy", "action"),
				),
			},
		},
	})
}
