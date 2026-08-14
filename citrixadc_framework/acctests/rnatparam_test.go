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

const testAccRnatparam_add = `
	resource "citrixadc_rnatparam" "tf_rnatparam" {
		tcpproxy         = "DISABLED"
		srcippersistency = "ENABLED"
	}
`
const testAccRnatparam_update = `
	resource "citrixadc_rnatparam" "tf_rnatparam" {
		tcpproxy         = "ENABLED"
		srcippersistency = "DISABLED"
	}
`

func TestAccRnatparam_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccRnatparam_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRnatparamExist("citrixadc_rnatparam.tf_rnatparam", nil),
					resource.TestCheckResourceAttr("citrixadc_rnatparam.tf_rnatparam", "tcpproxy", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_rnatparam.tf_rnatparam", "srcippersistency", "ENABLED"),
				),
			},
			{
				Config: testAccRnatparam_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRnatparamExist("citrixadc_rnatparam.tf_rnatparam", nil),
					resource.TestCheckResourceAttr("citrixadc_rnatparam.tf_rnatparam", "tcpproxy", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_rnatparam.tf_rnatparam", "srcippersistency", "DISABLED"),
				),
			},
		},
	})
}

func TestAccRnatparam_import(t *testing.T) {
	const resAddr = "citrixadc_rnatparam.tf_rnatparam"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{Config: testAccRnatparam_add},
			{
				Config:                  testAccRnatparam_add,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccCheckRnatparamExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No rnatparam name is set")
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
		data, err := client.FindResource(service.Rnatparam.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("rnatparam %s not found", n)
		}

		return nil
	}
}

const testAccRnatparamDataSource_basic = `

	resource "citrixadc_rnatparam" "tf_rnatparam" {
		srcippersistency = "ENABLED"
		tcpproxy         = "ENABLED"
	}

	data "citrixadc_rnatparam" "tf_rnatparam" {
		depends_on = [citrixadc_rnatparam.tf_rnatparam]
	}
`

func TestAccRnatparam_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccRnatparam_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckRnatparamExist("citrixadc_rnatparam.tf_rnatparam", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccRnatparam_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckRnatparamExist("citrixadc_rnatparam.tf_rnatparam", nil)),
			},
		},
	})
}

// step1 sets both unsettable attrs to non-default values; step2 removes them
// so the provider must unset them, reverting to NITRO spec defaults
// (tcpproxy=ENABLED, srcippersistency=DISABLED).
const testAccRnatparam_unset_step1 = `
	resource "citrixadc_rnatparam" "tf_unset" {
		tcpproxy         = "DISABLED"
		srcippersistency = "ENABLED"
	}
`

const testAccRnatparam_unset_step2 = `
	resource "citrixadc_rnatparam" "tf_unset" {
		# All unset-eligible attributes removed from config -> the provider must
		# unset them (revert to NITRO defaults).
	}
`

func TestAccRnatparam_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccRnatparam_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRnatparamExist("citrixadc_rnatparam.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_rnatparam.tf_unset", "tcpproxy", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_rnatparam.tf_unset", "srcippersistency", "ENABLED"),
				),
			},
			{
				// Removing the attributes must unset them: state reverts to the
				// documented NITRO defaults, and the implicit post-apply plan
				// must be empty.
				Config: testAccRnatparam_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRnatparamExist("citrixadc_rnatparam.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_rnatparam.tf_unset", "tcpproxy", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_rnatparam.tf_unset", "srcippersistency", "DISABLED"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckRnatparamADCValue("tcpproxy", "ENABLED"),
					testAccCheckRnatparamADCValue("srcippersistency", "DISABLED"),
				),
			},
		},
	})
}

// testAccCheckRnatparamADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted it.
func testAccCheckRnatparamADCValue(attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Rnatparam.Type(), "")
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("rnatparam not found on appliance")
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("rnatparam: appliance attr %q = %q, want %q (unset did not revert it)", attr, got, want)
		}
		return nil
	}
}

func TestAccRnatparamDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccRnatparamDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_rnatparam.tf_rnatparam", "srcippersistency", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_rnatparam.tf_rnatparam", "tcpproxy", "ENABLED"),
				),
			},
		},
	})
}
