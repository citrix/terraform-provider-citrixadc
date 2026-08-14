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

const testAccAuthenticationcitrixauthaction_add = `

	resource "citrixadc_authenticationcitrixauthaction" "tf_citrixauthaction" {
		name               = "tf_citrixauthaction"
		authenticationtype = "CITRIXCONNECTOR"
		authentication     = "DISABLED"
	}
`
const testAccAuthenticationcitrixauthaction_update = `

	resource "citrixadc_authenticationcitrixauthaction" "tf_citrixauthaction" {
		name               = "tf_citrixauthaction"
		authenticationtype = "ATHENA"
		authentication     = "ENABLED"
	}
`

func TestAccAuthenticationcitrixauthaction_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationcitrixauthactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationcitrixauthaction_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationcitrixauthactionExist("citrixadc_authenticationcitrixauthaction.tf_citrixauthaction", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationcitrixauthaction.tf_citrixauthaction", "name", "tf_citrixauthaction"),
					resource.TestCheckResourceAttr("citrixadc_authenticationcitrixauthaction.tf_citrixauthaction", "authenticationtype", "CITRIXCONNECTOR"),
					resource.TestCheckResourceAttr("citrixadc_authenticationcitrixauthaction.tf_citrixauthaction", "authentication", "DISABLED"),
				),
			},
			{
				Config: testAccAuthenticationcitrixauthaction_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationcitrixauthactionExist("citrixadc_authenticationcitrixauthaction.tf_citrixauthaction", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationcitrixauthaction.tf_citrixauthaction", "name", "tf_citrixauthaction"),
					resource.TestCheckResourceAttr("citrixadc_authenticationcitrixauthaction.tf_citrixauthaction", "authenticationtype", "ATHENA"),
					resource.TestCheckResourceAttr("citrixadc_authenticationcitrixauthaction.tf_citrixauthaction", "authentication", "ENABLED"),
				),
			},
		},
	})
}

func TestAccAuthenticationcitrixauthaction_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckAuthenticationcitrixauthactionDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccAuthenticationcitrixauthaction_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAuthenticationcitrixauthactionExist("citrixadc_authenticationcitrixauthaction.tf_citrixauthaction", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccAuthenticationcitrixauthaction_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAuthenticationcitrixauthactionExist("citrixadc_authenticationcitrixauthaction.tf_citrixauthaction", nil)),
			},
		},
	})
}

const testAccAuthenticationcitrixauthaction_unset_step1 = `
resource "citrixadc_authenticationcitrixauthaction" "tf_unset" {
  name               = "tf_test_citrixauthaction_unset"
  authenticationtype = "ATHENA"
  authentication     = "DISABLED"
}
`

const testAccAuthenticationcitrixauthaction_unset_step2 = `
resource "citrixadc_authenticationcitrixauthaction" "tf_unset" {
  name = "tf_test_citrixauthaction_unset"
  # All unset-eligible attributes removed from config -> the provider must
  # unset them (revert to NITRO defaults).
}
`

func TestAccAuthenticationcitrixauthaction_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationcitrixauthactionDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccAuthenticationcitrixauthaction_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationcitrixauthactionExist("citrixadc_authenticationcitrixauthaction.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationcitrixauthaction.tf_unset", "authenticationtype", "ATHENA"),
					resource.TestCheckResourceAttr("citrixadc_authenticationcitrixauthaction.tf_unset", "authentication", "DISABLED"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from
				// the appliance) reverts to the documented NITRO defaults, and the
				// implicit post-apply plan must be empty.
				Config: testAccAuthenticationcitrixauthaction_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationcitrixauthactionExist("citrixadc_authenticationcitrixauthaction.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationcitrixauthaction.tf_unset", "authenticationtype", "CITRIXCONNECTOR"),
					resource.TestCheckResourceAttr("citrixadc_authenticationcitrixauthaction.tf_unset", "authentication", "ENABLED"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckAuthenticationcitrixauthactionADCValue("tf_test_citrixauthaction_unset", "authenticationtype", "CITRIXCONNECTOR"),
					testAccCheckAuthenticationcitrixauthactionADCValue("tf_test_citrixauthaction_unset", "authentication", "ENABLED"),
				),
			},
		},
	})
}

// testAccCheckAuthenticationcitrixauthactionADCValue asserts an attribute's value
// directly on the appliance (not just in Terraform state), proving the unset
// actually reverted it.
func testAccCheckAuthenticationcitrixauthactionADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Authenticationcitrixauthaction.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("authenticationcitrixauthaction %s not found on appliance", name)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("authenticationcitrixauthaction %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func testAccCheckAuthenticationcitrixauthactionExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No authenticationcitrixauthaction name is set")
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
		data, err := client.FindResource("authenticationcitrixauthaction", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("authenticationcitrixauthaction %s not found", n)
		}

		return nil
	}
}

func testAccCheckAuthenticationcitrixauthactionDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_authenticationcitrixauthaction" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("authenticationcitrixauthaction", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("authenticationcitrixauthaction %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccAuthenticationcitrixauthaction_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_authenticationcitrixauthaction.tf_citrixauthaction"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationcitrixauthactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationcitrixauthaction_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAuthenticationcitrixauthactionExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Authenticationcitrixauthaction.Type(), "tf_citrixauthaction"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccAuthenticationcitrixauthaction_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAuthenticationcitrixauthactionExist(resAddr, nil)),
			},
		},
	})
}

func TestAccAuthenticationcitrixauthaction_import(t *testing.T) {
	const resAddr = "citrixadc_authenticationcitrixauthaction.tf_citrixauthaction"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationcitrixauthactionDestroy,
		Steps: []resource.TestStep{
			{Config: testAccAuthenticationcitrixauthaction_add},
			{
				Config:                  testAccAuthenticationcitrixauthaction_add,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

const testAccAuthenticationcitrixauthactionDataSource_basic = `

	resource "citrixadc_authenticationcitrixauthaction" "tf_citrixauthaction" {
		name               = "tf_citrixauthaction_ds"
		authenticationtype = "CITRIXCONNECTOR"
		authentication     = "DISABLED"
	}

	data "citrixadc_authenticationcitrixauthaction" "tf_citrixauthaction_ds" {
		name = citrixadc_authenticationcitrixauthaction.tf_citrixauthaction.name
	}
`

func TestAccAuthenticationcitrixauthactionDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationcitrixauthactionDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_authenticationcitrixauthaction.tf_citrixauthaction_ds", "name", "tf_citrixauthaction_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationcitrixauthaction.tf_citrixauthaction_ds", "authenticationtype", "CITRIXCONNECTOR"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationcitrixauthaction.tf_citrixauthaction_ds", "authentication", "DISABLED"),
				),
			},
		},
	})
}
