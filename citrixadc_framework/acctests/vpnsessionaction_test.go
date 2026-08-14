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

const testAccVpnsessionaction_add = `

	resource "citrixadc_vpnsessionaction" "foo" {
		name 					   = "newsession"
		sesstimeout                = "10"
  		defaultauthorizationaction = "ALLOW"
  		transparentinterception    = "ON"
  		clientidletimeout          = "10"
  		sso                        = "ON"
  		icaproxy                   = "ON"
  		wihome                     = "https://citrix.lab.com"
  		clientlessvpnmode          = "DISABLED"
  		
	}
`
const testAccVpnsessionaction_update = `

	resource "citrixadc_vpnsessionaction" "foo" {
		name 					   = "newsession"
		sesstimeout                = "20"
	 	defaultauthorizationaction = "DENY"
	  	transparentinterception    = "ON"
		clientidletimeout          = "20"
		sso                        = "ON"
		icaproxy                   = "OFF"
		wihome                     = "https://citrix.lab.com"
		clientlessvpnmode          = "DISABLED"
		httpport                   = [8080, 8000, 808]	
	}
`

const testAccVpnsessionactionDataSource_basic = `

	resource "citrixadc_vpnsessionaction" "foo" {
		name 					   = "newsession"
		sesstimeout                = "10"
  		defaultauthorizationaction = "ALLOW"
  		transparentinterception    = "ON"
  		clientidletimeout          = "10"
  		sso                        = "ON"
  		icaproxy                   = "ON"
  		wihome                     = "https://citrix.lab.com"
  		clientlessvpnmode          = "DISABLED"
	}

	data "citrixadc_vpnsessionaction" "foo" {
		name = citrixadc_vpnsessionaction.foo.name
	}
`

func TestAccVpnsessionaction_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVpnsessionactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnsessionaction_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnsessionactionExist("citrixadc_vpnsessionaction.foo", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnsessionaction.foo", "name", "newsession"),
					resource.TestCheckResourceAttr("citrixadc_vpnsessionaction.foo", "sesstimeout", "10"),
					resource.TestCheckResourceAttr("citrixadc_vpnsessionaction.foo", "defaultauthorizationaction", "ALLOW"),
				),
			},
			{
				Config: testAccVpnsessionaction_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnsessionactionExist("citrixadc_vpnsessionaction.foo", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnsessionaction.foo", "name", "newsession"),
					resource.TestCheckResourceAttr("citrixadc_vpnsessionaction.foo", "sesstimeout", "20"),
					resource.TestCheckResourceAttr("citrixadc_vpnsessionaction.foo", "clientidletimeout", "20"),
				),
			},
		},
	})
}

func testAccCheckVpnsessionactionExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No vpnsessionaction name is set")
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
		data, err := client.FindResource(service.Vpnsessionaction.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("vpnsessionaction %s not found", n)
		}

		return nil
	}
}

func testAccCheckVpnsessionactionDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_vpnsessionaction" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Vpnsessionaction.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("vpnsessionaction %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccVpnsessionaction_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_vpnsessionaction.foo"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVpnsessionactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnsessionaction_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckVpnsessionactionExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Vpnsessionaction.Type(), "newsession"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccVpnsessionaction_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckVpnsessionactionExist(resAddr, nil)),
			},
		},
	})
}

func TestAccVpnsessionaction_import(t *testing.T) {
	const resAddr = "citrixadc_vpnsessionaction.foo"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVpnsessionactionDestroy,
		Steps: []resource.TestStep{
			{Config: testAccVpnsessionaction_add},
			{
				Config:                  testAccVpnsessionaction_add,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccVpnsessionaction_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckVpnsessionactionDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccVpnsessionaction_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckVpnsessionactionExist("citrixadc_vpnsessionaction.foo", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccVpnsessionaction_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckVpnsessionactionExist("citrixadc_vpnsessionaction.foo", nil)),
			},
		},
	})
}

// The vpnsessionaction unset test covers the only attribute with a documented
// NITRO server default (advancedclientlessvpnmode, default DISABLED). Step 1
// sets it to a non-default value; step 2 removes it from config, and the
// provider must unset it (revert to DISABLED).
const testAccVpnsessionaction_unset_step1 = `
	resource "citrixadc_vpnsessionaction" "tf_unset" {
		name                      = "tf_test_vpnsessionaction_unset"
		advancedclientlessvpnmode = "ENABLED"
	}
`

const testAccVpnsessionaction_unset_step2 = `
	resource "citrixadc_vpnsessionaction" "tf_unset" {
		name = "tf_test_vpnsessionaction_unset"
		# advancedclientlessvpnmode removed from config -> provider must unset it
		# (revert to NITRO default, "DISABLED").
	}
`

func TestAccVpnsessionaction_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVpnsessionactionDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default value is applied and persisted.
				Config: testAccVpnsessionaction_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnsessionactionExist("citrixadc_vpnsessionaction.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnsessionaction.tf_unset", "advancedclientlessvpnmode", "ENABLED"),
				),
			},
			{
				// Removing the attribute must unset it: state (read back from the
				// appliance) reverts to the NITRO default, and the implicit
				// post-apply plan must be empty.
				Config: testAccVpnsessionaction_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnsessionactionExist("citrixadc_vpnsessionaction.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnsessionaction.tf_unset", "advancedclientlessvpnmode", "DISABLED"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckVpnsessionactionADCValue("tf_test_vpnsessionaction_unset", "advancedclientlessvpnmode", "DISABLED"),
				),
			},
		},
	})
}

// testAccCheckVpnsessionactionADCValue asserts an attribute's value directly on
// the appliance (not just in Terraform state), proving the unset actually
// reverted it.
func testAccCheckVpnsessionactionADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Vpnsessionaction.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("vpnsessionaction %s not found on appliance", name)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("vpnsessionaction %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func TestAccVpnsessionactionDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnsessionactionDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_vpnsessionaction.foo", "name", "newsession"),
					resource.TestCheckResourceAttr("data.citrixadc_vpnsessionaction.foo", "sesstimeout", "10"),
					resource.TestCheckResourceAttr("data.citrixadc_vpnsessionaction.foo", "defaultauthorizationaction", "ALLOW"),
				),
			},
		},
	})
}
