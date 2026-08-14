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

const testAccRnat_add = `

	resource "citrixadc_rnat" "tfrnat" {
		name             = "tfrnat"
		network          = "10.2.2.0"
		netmask          = "255.255.255.255"
		useproxyport     = "ENABLED"
		srcippersistency = "DISABLED"
		connfailover     = "DISABLED"
	}
`
const testAccRnat_update = `

	resource "citrixadc_rnat" "tfrnat" {
		name             = "tfrnat"
		network          = "10.2.2.0"
		netmask          = "255.255.255.255"
		useproxyport     = "DISABLED"
		srcippersistency = "DISABLED"
		connfailover     = "DISABLED"
	}
`

func TestAccRnat_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRnatDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccRnat_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRnatExist("citrixadc_rnat.tfrnat", nil),
					resource.TestCheckResourceAttr("citrixadc_rnat.tfrnat", "name", "tfrnat"),
					resource.TestCheckResourceAttr("citrixadc_rnat.tfrnat", "useproxyport", "ENABLED"),
				),
			},
			{
				Config: testAccRnat_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRnatExist("citrixadc_rnat.tfrnat", nil),
					resource.TestCheckResourceAttr("citrixadc_rnat.tfrnat", "name", "tfrnat"),
					resource.TestCheckResourceAttr("citrixadc_rnat.tfrnat", "useproxyport", "DISABLED"),
				),
			},
		},
	})
}

func testAccCheckRnatExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No rnat name is set")
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
		data, err := client.FindResource(service.Rnat.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("rnat %s not found", n)
		}

		return nil
	}
}

func testAccCheckRnatDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_rnat" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Rnat.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("rnat %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccRnatDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccRnatDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_rnat.test", "name", "tf_rnat_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_rnat.test", "network", "10.3.3.0"),
					resource.TestCheckResourceAttr("data.citrixadc_rnat.test", "netmask", "255.255.255.255"),
					resource.TestCheckResourceAttrSet("data.citrixadc_rnat.test", "id"),
				),
			},
		},
	})
}

func TestAccRnat_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_rnat.tfrnat"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRnatDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccRnat_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckRnatExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Rnat.Type(), "tfrnat"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccRnat_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckRnatExist(resAddr, nil)),
			},
		},
	})
}

func TestAccRnat_import(t *testing.T) {
	const resAddr = "citrixadc_rnat.tfrnat"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRnatDestroy,
		Steps: []resource.TestStep{
			{Config: testAccRnat_add},
			{
				Config:                  testAccRnat_add,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccRnat_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckRnatDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccRnat_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckRnatExist("citrixadc_rnat.tfrnat", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccRnat_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckRnatExist("citrixadc_rnat.tfrnat", nil)),
			},
		},
	})
}

// The rnat unset test covers the mutable, spec-unsettable toggle attributes
// that have documented NITRO defaults and no prerequisites: connfailover
// (default DISABLED), srcippersistency (default DISABLED) and useproxyport
// (default ENABLED). ownergroup (cluster-only) and td (requires a
// pre-existing traffic domain) have prerequisites that cannot be satisfied on
// a standalone appliance and are excluded.
const testAccRnat_unset_step1 = `
resource "citrixadc_rnat" "tf_unset" {
	name             = "tf_test_rnat_unset"
	network          = "10.5.5.0"
	netmask          = "255.255.255.255"
	connfailover     = "ENABLED"
	srcippersistency = "ENABLED"
	useproxyport     = "DISABLED"
}
`

const testAccRnat_unset_step2 = `
resource "citrixadc_rnat" "tf_unset" {
	name    = "tf_test_rnat_unset"
	network = "10.5.5.0"
	netmask = "255.255.255.255"
	# All unset-eligible attributes removed from config -> the provider must
	# unset them (revert to NITRO defaults).
}
`

func TestAccRnat_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRnatDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccRnat_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRnatExist("citrixadc_rnat.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_rnat.tf_unset", "connfailover", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_rnat.tf_unset", "srcippersistency", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_rnat.tf_unset", "useproxyport", "DISABLED"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from
				// the appliance) reverts to the documented NITRO defaults, and the
				// implicit post-apply plan must be empty.
				Config: testAccRnat_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRnatExist("citrixadc_rnat.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_rnat.tf_unset", "connfailover", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_rnat.tf_unset", "srcippersistency", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_rnat.tf_unset", "useproxyport", "ENABLED"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckRnatADCValue("tf_test_rnat_unset", "connfailover", "DISABLED"),
					testAccCheckRnatADCValue("tf_test_rnat_unset", "srcippersistency", "DISABLED"),
					testAccCheckRnatADCValue("tf_test_rnat_unset", "useproxyport", "ENABLED"),
				),
			},
		},
	})
}

// testAccCheckRnatADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted
// it.
func testAccCheckRnatADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Rnat.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("rnat %s not found on appliance", name)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("rnat %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

const testAccRnatDataSource_basic = `
resource "citrixadc_rnat" "test" {
	name    = "tf_rnat_ds"
	network = "10.3.3.0"
	netmask = "255.255.255.255"
}

data "citrixadc_rnat" "test" {
	name = citrixadc_rnat.test.name
}
`
