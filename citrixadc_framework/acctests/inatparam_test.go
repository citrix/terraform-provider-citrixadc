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

const testAccInatparam_basic = `

	resource "citrixadc_inatparam" "tf_inatparam" {
		nat46ignoretos    = "NO"
		nat46zerochecksum = "ENABLED"
		nat46v6mtu        = "1400"
	}
`
const testAccInatparam_update = `

	resource "citrixadc_inatparam" "tf_inatparam" {
		nat46ignoretos    = "YES"
		nat46zerochecksum = "DISABLED"
		nat46v6mtu        = "1300"
	}
`

func TestAccInatparam_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccInatparam_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckInatparamExist("citrixadc_inatparam.tf_inatparam", nil),
					resource.TestCheckResourceAttr("citrixadc_inatparam.tf_inatparam", "nat46ignoretos", "NO"),
					resource.TestCheckResourceAttr("citrixadc_inatparam.tf_inatparam", "nat46zerochecksum", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_inatparam.tf_inatparam", "nat46v6mtu", "1400"),
				),
			},
			{
				Config: testAccInatparam_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckInatparamExist("citrixadc_inatparam.tf_inatparam", nil),
					resource.TestCheckResourceAttr("citrixadc_inatparam.tf_inatparam", "nat46ignoretos", "YES"),
					resource.TestCheckResourceAttr("citrixadc_inatparam.tf_inatparam", "nat46zerochecksum", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_inatparam.tf_inatparam", "nat46v6mtu", "1300"),
				),
			},
		},
	})
}

func testAccCheckInatparamExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No inatparam name is set")
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
		data, err := client.FindResource(service.Inatparam.Type(), rs.Primary.Attributes["td"])

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("inatparam %s not found", n)
		}

		return nil
	}
}

func TestAccInatparam_import(t *testing.T) {
	const resAddr = "citrixadc_inatparam.tf_inatparam"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{Config: testAccInatparam_basic},
			{
				Config:                  testAccInatparam_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccInatparam_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccInatparam_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckInatparamExist("citrixadc_inatparam.tf_inatparam", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccInatparam_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckInatparamExist("citrixadc_inatparam.tf_inatparam", nil)),
			},
		},
	})
}

// The inatparam unset test covers the only mutable, spec-unsettable attribute
// the appliance accepts for a NITRO ?action=unset: nat46v6prefix. The nat46*
// toggle/mtu attributes are rejected with "Invalid argument [..]" on unset, so
// they are not wired. Step 1 sets a non-default prefix; step 2 removes it from
// config so the provider unsets it (the prefix reverts to no value / absent).
const testAccInatparam_unset_step1 = `

	resource "citrixadc_inatparam" "tf_unset" {
		nat46v6prefix = "2001:db8::/96"
	}
`

const testAccInatparam_unset_step2 = `

	resource "citrixadc_inatparam" "tf_unset" {
		# nat46v6prefix removed from config -> the provider must unset it.
	}
`

func TestAccInatparam_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				// Non-default value is applied and persisted.
				Config: testAccInatparam_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckInatparamExist("citrixadc_inatparam.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_inatparam.tf_unset", "nat46v6prefix", "2001:db8::/96"),
					testAccCheckInatparamADCValue("0", "nat46v6prefix", "2001:db8::/96"),
				),
			},
			{
				// Removing the attribute must unset it: state (read back from the
				// appliance) reverts to no value, and the implicit post-apply plan
				// must be empty.
				Config: testAccInatparam_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckInatparamExist("citrixadc_inatparam.tf_unset", nil),
					// Independent appliance-level confirmation the unset took effect
					// (the prefix is absent from GET after unset).
					testAccCheckInatparamADCValue("0", "nat46v6prefix", ""),
				),
			},
		},
	})
}

// testAccCheckInatparamADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted
// it. The inatparam resource is keyed on the traffic domain (td).
func testAccCheckInatparamADCValue(td, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Inatparam.Type(), td)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("inatparam td=%s not found on appliance", td)
		}
		got := ""
		if v, ok := data[attr]; ok && v != nil {
			got = strings.TrimSpace(fmt.Sprintf("%v", v))
		}
		if got != want {
			return fmt.Errorf("inatparam td=%s: appliance attr %q = %q, want %q (unset did not revert it)", td, attr, got, want)
		}
		return nil
	}
}

func TestAccInatparamDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccInatparamDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_inatparam.tf_inatparam_ds", "nat46ignoretos", "NO"),
					resource.TestCheckResourceAttr("data.citrixadc_inatparam.tf_inatparam_ds", "nat46zerochecksum", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_inatparam.tf_inatparam_ds", "nat46v6mtu", "1400"),
					resource.TestCheckResourceAttr("data.citrixadc_inatparam.tf_inatparam_ds", "td", "0"),
				),
			},
		},
	})
}

const testAccInatparamDataSource_basic = `

resource "citrixadc_inatparam" "tf_inatparam_ds" {
	nat46ignoretos    = "NO"
	nat46zerochecksum = "ENABLED"
	nat46v6mtu        = 1400
}

data "citrixadc_inatparam" "tf_inatparam_ds" {
	td = 0
	depends_on = [citrixadc_inatparam.tf_inatparam_ds]
}

`
