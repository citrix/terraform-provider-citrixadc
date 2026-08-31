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

const testAccAuthenticationcertaction_add = `
	resource "citrixadc_authenticationcertaction" "tf_certaction" {
		name                       = "tf_certaction"
		twofactor                  = "ON"
		defaultauthenticationgroup = "old_group"
		usernamefield              = "Subject:CN"
		groupnamefield             = "subject:grp"
	}
`
const testAccAuthenticationcertaction_update = `
	resource "citrixadc_authenticationcertaction" "tf_certaction" {
		name                       = "tf_certaction"
		twofactor                  = "OFF"
		defaultauthenticationgroup = "new_group"
		usernamefield              = "Subject:CN"
		groupnamefield             = "subject:grp"
	}
`

const testAccAuthenticationcertactionDataSource_basic = `
	resource "citrixadc_authenticationcertaction" "tf_certaction_ds" {
		name                       = "tf_certaction_ds"
		twofactor                  = "ON"
		defaultauthenticationgroup = "test_group"
		usernamefield              = "Subject:CN"
		groupnamefield             = "subject:grp"
	}

	data "citrixadc_authenticationcertaction" "tf_certaction_ds" {
		name = citrixadc_authenticationcertaction.tf_certaction_ds.name
		depends_on = [citrixadc_authenticationcertaction.tf_certaction_ds]
	}
`

func TestAccAuthenticationcertaction_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationcertactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationcertaction_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationcertactionExist("citrixadc_authenticationcertaction.tf_certaction", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationcertaction.tf_certaction", "name", "tf_certaction"),
					resource.TestCheckResourceAttr("citrixadc_authenticationcertaction.tf_certaction", "twofactor", "ON"),
					resource.TestCheckResourceAttr("citrixadc_authenticationcertaction.tf_certaction", "defaultauthenticationgroup", "old_group"),
				),
			},
			{
				Config: testAccAuthenticationcertaction_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationcertactionExist("citrixadc_authenticationcertaction.tf_certaction", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationcertaction.tf_certaction", "name", "tf_certaction"),
					resource.TestCheckResourceAttr("citrixadc_authenticationcertaction.tf_certaction", "twofactor", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_authenticationcertaction.tf_certaction", "defaultauthenticationgroup", "new_group"),
				),
			},
		},
	})
}

func testAccCheckAuthenticationcertactionExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No authenticationcertaction name is set")
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
		data, err := client.FindResource(service.Authenticationcertaction.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("authenticationcertaction %s not found", n)
		}

		return nil
	}
}

func testAccCheckAuthenticationcertactionDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_authenticationcertaction" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Authenticationcertaction.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("authenticationcertaction %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccAuthenticationcertaction_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_authenticationcertaction.tf_certaction"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationcertactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationcertaction_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAuthenticationcertactionExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Authenticationcertaction.Type(), "tf_certaction"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccAuthenticationcertaction_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAuthenticationcertactionExist(resAddr, nil)),
			},
		},
	})
}

func TestAccAuthenticationcertaction_import(t *testing.T) {
	const resAddr = "citrixadc_authenticationcertaction.tf_certaction"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationcertactionDestroy,
		Steps: []resource.TestStep{
			{Config: testAccAuthenticationcertaction_add},
			{
				Config:                  testAccAuthenticationcertaction_add,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccAuthenticationcertaction_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckAuthenticationcertactionDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.0.0"},
				},
				Config: testAccAuthenticationcertaction_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAuthenticationcertactionExist("citrixadc_authenticationcertaction.tf_certaction", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccAuthenticationcertaction_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAuthenticationcertactionExist("citrixadc_authenticationcertaction.tf_certaction", nil)),
			},
		},
	})
}

// twofactor is the only spec-unsettable attribute with a documented NITRO
// default (OFF). Removing it from config must unset it back to that default.
const testAccAuthenticationcertaction_unset_step1 = `
	resource "citrixadc_authenticationcertaction" "tf_unset" {
		name      = "tf_test_certaction_unset"
		twofactor = "ON"
	}
`

const testAccAuthenticationcertaction_unset_step2 = `
	resource "citrixadc_authenticationcertaction" "tf_unset" {
		name = "tf_test_certaction_unset"
		# twofactor removed from config -> provider must unset it (revert to "OFF").
	}
`

func TestAccAuthenticationcertaction_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationcertactionDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default value is applied and persisted.
				Config: testAccAuthenticationcertaction_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationcertactionExist("citrixadc_authenticationcertaction.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationcertaction.tf_unset", "twofactor", "ON"),
				),
			},
			{
				// Removing the attribute must unset it: state (read back from the
				// appliance) reverts to the documented NITRO default, and the
				// implicit post-apply plan must be empty.
				Config: testAccAuthenticationcertaction_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationcertactionExist("citrixadc_authenticationcertaction.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationcertaction.tf_unset", "twofactor", "OFF"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckAuthenticationcertactionADCValue("tf_test_certaction_unset", "twofactor", "OFF"),
				),
			},
		},
	})
}

// testAccCheckAuthenticationcertactionADCValue asserts an attribute's value
// directly on the appliance (not just in Terraform state), proving the unset
// actually reverted it.
func testAccCheckAuthenticationcertactionADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Authenticationcertaction.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("authenticationcertaction %s not found on appliance", name)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("authenticationcertaction %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func TestAccAuthenticationcertactionDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationcertactionDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_authenticationcertaction.tf_certaction_ds", "name", "tf_certaction_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationcertaction.tf_certaction_ds", "twofactor", "ON"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationcertaction.tf_certaction_ds", "defaultauthenticationgroup", "test_group"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationcertaction.tf_certaction_ds", "usernamefield", "Subject:CN"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationcertaction.tf_certaction_ds", "groupnamefield", "subject:grp"),
				),
			},
		},
	})
}
