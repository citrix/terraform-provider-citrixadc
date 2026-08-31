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

const testAccAuthenticationpolicy_add = `

	resource "citrixadc_authenticationldapaction" "tf_authenticationldapaction" {
		name          = "ldapaction"
		serverip      = "1.2.3.4"
		serverport    = 8080
		authtimeout   = 1
		ldaploginname = "username"
	}
	resource "citrixadc_authenticationpolicy" "tf_authenticationpolicy" {
		name   = "tf_authenticationpolicy"
		rule   = "true"
		action = citrixadc_authenticationldapaction.tf_authenticationldapaction.name
		comment= "new_policy"
	}
`
const testAccAuthenticationpolicy_update = `
	resource "citrixadc_authenticationldapaction" "tf_authenticationldapaction" {
		name          = "ldapaction"
		serverip      = "1.2.3.4"
		serverport    = 8080
		authtimeout   = 1
		ldaploginname = "username"
	}
	resource "citrixadc_authenticationpolicy" "tf_authenticationpolicy" {
		name   = "tf_authenticationpolicy"
		rule   = "true"
		action = citrixadc_authenticationldapaction.tf_authenticationldapaction.name
		comment= "updated"
	}
`

func TestAccAuthenticationpolicy_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationpolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationpolicy_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationpolicyExist("citrixadc_authenticationpolicy.tf_authenticationpolicy", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationpolicy.tf_authenticationpolicy", "name", "tf_authenticationpolicy"),
					resource.TestCheckResourceAttr("citrixadc_authenticationpolicy.tf_authenticationpolicy", "comment", "new_policy"),
				),
			},
			{
				Config: testAccAuthenticationpolicy_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationpolicyExist("citrixadc_authenticationpolicy.tf_authenticationpolicy", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationpolicy.tf_authenticationpolicy", "name", "tf_authenticationpolicy"),
					resource.TestCheckResourceAttr("citrixadc_authenticationpolicy.tf_authenticationpolicy", "comment", "updated"),
				),
			},
		},
	})
}

func testAccCheckAuthenticationpolicyExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No authenticationpolicy name is set")
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
		data, err := client.FindResource(service.Authenticationpolicy.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("authenticationpolicy %s not found", n)
		}

		return nil
	}
}

func testAccCheckAuthenticationpolicyDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_authenticationpolicy" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Authenticationpolicy.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("authenticationpolicy %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccAuthenticationpolicy_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_authenticationpolicy.tf_authenticationpolicy"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationpolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationpolicy_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAuthenticationpolicyExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Authenticationpolicy.Type(), "tf_authenticationpolicy"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccAuthenticationpolicy_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAuthenticationpolicyExist(resAddr, nil)),
			},
		},
	})
}

func TestAccAuthenticationpolicy_import(t *testing.T) {
	const resAddr = "citrixadc_authenticationpolicy.tf_authenticationpolicy"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationpolicyDestroy,
		Steps: []resource.TestStep{
			{Config: testAccAuthenticationpolicy_add},
			{
				Config:                  testAccAuthenticationpolicy_add,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

const testAccAuthenticationpolicyDataSource_basic = `
	resource "citrixadc_authenticationldapaction" "tf_authenticationldapaction" {
		name          = "ldapaction_ds"
		serverip      = "1.2.3.4"
		serverport    = 8080
		authtimeout   = 1
		ldaploginname = "username"
	}
	resource "citrixadc_authenticationpolicy" "tf_authenticationpolicy" {
		name   = "tf_authenticationpolicy_ds"
		rule   = "true"
		action = citrixadc_authenticationldapaction.tf_authenticationldapaction.name
		comment= "datasource_test"
	}
	data "citrixadc_authenticationpolicy" "tf_authenticationpolicy_ds" {
		name = citrixadc_authenticationpolicy.tf_authenticationpolicy.name
	}
`

// authenticationpolicy exposes one cleanly-unsettable optional attribute:
// comment. undefaction is not echoed back by NITRO GET (write-only-ish) and
// logaction requires a valid audit messagelog action, so neither is a reliable
// unset candidate; only comment is exercised here.
const testAccAuthenticationpolicy_unset_step1 = `
	resource "citrixadc_authenticationldapaction" "tf_authenticationldapaction" {
		name          = "ldapaction_unset"
		serverip      = "1.2.3.4"
		serverport    = 8080
		authtimeout   = 1
		ldaploginname = "username"
	}
	resource "citrixadc_authenticationpolicy" "tf_unset" {
		name    = "tf_authenticationpolicy_unset"
		rule    = "true"
		action  = citrixadc_authenticationldapaction.tf_authenticationldapaction.name
		comment = "unset_me"
	}
`

const testAccAuthenticationpolicy_unset_step2 = `
	resource "citrixadc_authenticationldapaction" "tf_authenticationldapaction" {
		name          = "ldapaction_unset"
		serverip      = "1.2.3.4"
		serverport    = 8080
		authtimeout   = 1
		ldaploginname = "username"
	}
	resource "citrixadc_authenticationpolicy" "tf_unset" {
		name   = "tf_authenticationpolicy_unset"
		rule   = "true"
		action = citrixadc_authenticationldapaction.tf_authenticationldapaction.name
		# comment removed from config -> provider must unset it (revert to default "").
	}
`

func TestAccAuthenticationpolicy_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationpolicyDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default value is applied and persisted.
				Config: testAccAuthenticationpolicy_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationpolicyExist("citrixadc_authenticationpolicy.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationpolicy.tf_unset", "comment", "unset_me"),
				),
			},
			{
				// Removing comment must unset it: state reverts to the NITRO default
				// (empty) and the implicit post-apply plan must be empty.
				Config: testAccAuthenticationpolicy_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationpolicyExist("citrixadc_authenticationpolicy.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationpolicy.tf_unset", "comment", ""),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckAuthenticationpolicyADCValue("tf_authenticationpolicy_unset", "comment", ""),
				),
			},
		},
	})
}

// testAccCheckAuthenticationpolicyADCValue asserts an attribute's value directly
// on the appliance (not just in Terraform state), proving the unset actually
// reverted it. A missing key is treated as the empty default.
func testAccCheckAuthenticationpolicyADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Authenticationpolicy.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("authenticationpolicy %s not found on appliance", name)
		}
		got := ""
		if v, ok := data[attr]; ok && v != nil {
			got = strings.TrimSpace(fmt.Sprintf("%v", v))
		}
		if got != want {
			return fmt.Errorf("authenticationpolicy %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func TestAccAuthenticationpolicy_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckAuthenticationpolicyDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.0.0"},
				},
				Config: testAccAuthenticationpolicy_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAuthenticationpolicyExist("citrixadc_authenticationpolicy.tf_authenticationpolicy", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccAuthenticationpolicy_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAuthenticationpolicyExist("citrixadc_authenticationpolicy.tf_authenticationpolicy", nil)),
			},
		},
	})
}

func TestAccAuthenticationpolicyDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationpolicyDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_authenticationpolicy.tf_authenticationpolicy_ds", "name", "tf_authenticationpolicy_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationpolicy.tf_authenticationpolicy_ds", "rule", "true"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationpolicy.tf_authenticationpolicy_ds", "comment", "datasource_test"),
					resource.TestCheckResourceAttrPair("data.citrixadc_authenticationpolicy.tf_authenticationpolicy_ds", "action", "citrixadc_authenticationldapaction.tf_authenticationldapaction", "name"),
				),
			},
		},
	})
}
