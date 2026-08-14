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

const testAccOnlinkipv6prefix_basic = `


	resource "citrixadc_onlinkipv6prefix" "tf_onlinkipv6prefix" {
		ipv6prefix      = "8000::/64"
		onlinkprefix    = "YES"
		autonomusprefix = "NO"
	}
`
const testAccOnlinkipv6prefix_update = `


	resource "citrixadc_onlinkipv6prefix" "tf_onlinkipv6prefix" {
		ipv6prefix      = "8000::/64"
		onlinkprefix    = "NO"
		autonomusprefix = "YES"
	}
`

func TestAccOnlinkipv6prefix_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckOnlinkipv6prefixDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccOnlinkipv6prefix_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOnlinkipv6prefixExist("citrixadc_onlinkipv6prefix.tf_onlinkipv6prefix", nil),
					resource.TestCheckResourceAttr("citrixadc_onlinkipv6prefix.tf_onlinkipv6prefix", "onlinkprefix", "YES"),
					resource.TestCheckResourceAttr("citrixadc_onlinkipv6prefix.tf_onlinkipv6prefix", "autonomusprefix", "NO"),
				),
			},
			{
				Config: testAccOnlinkipv6prefix_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOnlinkipv6prefixExist("citrixadc_onlinkipv6prefix.tf_onlinkipv6prefix", nil),
					resource.TestCheckResourceAttr("citrixadc_onlinkipv6prefix.tf_onlinkipv6prefix", "onlinkprefix", "NO"),
					resource.TestCheckResourceAttr("citrixadc_onlinkipv6prefix.tf_onlinkipv6prefix", "autonomusprefix", "YES"),
				),
			},
		},
	})
}

func TestAccOnlinkipv6prefix_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_onlinkipv6prefix.tf_onlinkipv6prefix"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckOnlinkipv6prefixDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccOnlinkipv6prefix_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckOnlinkipv6prefixExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Onlinkipv6prefix.Type(), "8000::/64"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccOnlinkipv6prefix_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckOnlinkipv6prefixExist(resAddr, nil)),
			},
		},
	})
}

func TestAccOnlinkipv6prefix_import(t *testing.T) {
	const resAddr = "citrixadc_onlinkipv6prefix.tf_onlinkipv6prefix"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckOnlinkipv6prefixDestroy,
		Steps: []resource.TestStep{
			{Config: testAccOnlinkipv6prefix_basic},
			{
				Config:                  testAccOnlinkipv6prefix_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccCheckOnlinkipv6prefixExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No onlinkipv6prefix name is set")
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
		data, err := client.FindResource(service.Onlinkipv6prefix.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("onlinkipv6prefix %s not found", n)
		}

		return nil
	}
}

func testAccCheckOnlinkipv6prefixDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_onlinkipv6prefix" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Onlinkipv6prefix.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("onlinkipv6prefix %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccOnlinkipv6prefix_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckOnlinkipv6prefixDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccOnlinkipv6prefix_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckOnlinkipv6prefixExist("citrixadc_onlinkipv6prefix.tf_onlinkipv6prefix", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccOnlinkipv6prefix_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckOnlinkipv6prefixExist("citrixadc_onlinkipv6prefix.tf_onlinkipv6prefix", nil)),
			},
		},
	})
}

// The onlinkipv6prefix unset test exercises every spec-unsettable attribute:
// step1 sets each to a valid non-default value; step2 removes them from config so
// the provider unsets them and the appliance reverts to the documented NITRO
// defaults (onlinkprefix=YES, autonomusprefix=YES, depricateprefix=NO,
// decrementprefixlifetimes=NO, prefixpreferredlifetime=604800,
// prefixvalidelifetime=2592000).
const testAccOnlinkipv6prefix_unset_step1 = `
resource "citrixadc_onlinkipv6prefix" "tf_unset" {
  ipv6prefix               = "a000::/64"
  onlinkprefix             = "NO"
  autonomusprefix          = "NO"
  depricateprefix          = "YES"
  decrementprefixlifetimes = "YES"
  prefixpreferredlifetime  = 302400
  prefixvalidelifetime     = 1296000
}
`

const testAccOnlinkipv6prefix_unset_step2 = `
resource "citrixadc_onlinkipv6prefix" "tf_unset" {
  ipv6prefix = "a000::/64"
  # All unset-eligible attributes removed from config -> the provider must
  # unset them (revert to NITRO defaults).
}
`

func TestAccOnlinkipv6prefix_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckOnlinkipv6prefixDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccOnlinkipv6prefix_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOnlinkipv6prefixExist("citrixadc_onlinkipv6prefix.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_onlinkipv6prefix.tf_unset", "onlinkprefix", "NO"),
					resource.TestCheckResourceAttr("citrixadc_onlinkipv6prefix.tf_unset", "autonomusprefix", "NO"),
					resource.TestCheckResourceAttr("citrixadc_onlinkipv6prefix.tf_unset", "depricateprefix", "YES"),
					resource.TestCheckResourceAttr("citrixadc_onlinkipv6prefix.tf_unset", "decrementprefixlifetimes", "YES"),
					resource.TestCheckResourceAttr("citrixadc_onlinkipv6prefix.tf_unset", "prefixpreferredlifetime", "302400"),
					resource.TestCheckResourceAttr("citrixadc_onlinkipv6prefix.tf_unset", "prefixvalidelifetime", "1296000"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from the
				// appliance) reverts to the documented NITRO defaults, and the
				// implicit post-apply plan must be empty.
				Config: testAccOnlinkipv6prefix_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOnlinkipv6prefixExist("citrixadc_onlinkipv6prefix.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_onlinkipv6prefix.tf_unset", "onlinkprefix", "YES"),
					resource.TestCheckResourceAttr("citrixadc_onlinkipv6prefix.tf_unset", "autonomusprefix", "YES"),
					resource.TestCheckResourceAttr("citrixadc_onlinkipv6prefix.tf_unset", "depricateprefix", "NO"),
					resource.TestCheckResourceAttr("citrixadc_onlinkipv6prefix.tf_unset", "decrementprefixlifetimes", "NO"),
					resource.TestCheckResourceAttr("citrixadc_onlinkipv6prefix.tf_unset", "prefixpreferredlifetime", "604800"),
					resource.TestCheckResourceAttr("citrixadc_onlinkipv6prefix.tf_unset", "prefixvalidelifetime", "2592000"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckOnlinkipv6prefixADCValue("a000::/64", "onlinkprefix", "YES"),
					testAccCheckOnlinkipv6prefixADCValue("a000::/64", "autonomusprefix", "YES"),
					testAccCheckOnlinkipv6prefixADCValue("a000::/64", "depricateprefix", "NO"),
				),
			},
		},
	})
}

// testAccCheckOnlinkipv6prefixADCValue asserts an attribute's value directly on
// the appliance (not just in Terraform state), proving the unset reverted it.
func testAccCheckOnlinkipv6prefixADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Onlinkipv6prefix.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("onlinkipv6prefix %s not found on appliance", name)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("onlinkipv6prefix %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func TestAccOnlinkipv6prefixDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckOnlinkipv6prefixDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccOnlinkipv6prefixDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_onlinkipv6prefix.tf_onlinkipv6prefix_ds", "ipv6prefix", "9000::/64"),
					resource.TestCheckResourceAttr("data.citrixadc_onlinkipv6prefix.tf_onlinkipv6prefix_ds", "onlinkprefix", "YES"),
					resource.TestCheckResourceAttr("data.citrixadc_onlinkipv6prefix.tf_onlinkipv6prefix_ds", "autonomusprefix", "NO"),
				),
			},
		},
	})
}

const testAccOnlinkipv6prefixDataSource_basic = `

resource "citrixadc_onlinkipv6prefix" "tf_onlinkipv6prefix_ds" {
	ipv6prefix      = "9000::/64"
	onlinkprefix    = "YES"
	autonomusprefix = "NO"
}

data "citrixadc_onlinkipv6prefix" "tf_onlinkipv6prefix_ds" {
	ipv6prefix = citrixadc_onlinkipv6prefix.tf_onlinkipv6prefix_ds.ipv6prefix
	depends_on = [citrixadc_onlinkipv6prefix.tf_onlinkipv6prefix_ds]
}
`
