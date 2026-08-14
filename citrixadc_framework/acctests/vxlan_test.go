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

const testAccVxlan_add = `
	resource "citrixadc_vlan" "tf_vlan" {
		vlanid    = 40
		aliasname = "Management VLAN"
	}
	resource "citrixadc_vxlan" "tf_vxlan" {
		vxlanid            = 123
		vlan               = citrixadc_vlan.tf_vlan.vlanid
		port               = 33
		dynamicrouting     = "DISABLED"
		ipv6dynamicrouting = "DISABLED"
		innervlantagging   = "ENABLED"
	}
`
const testAccVxlan_update = `
	resource "citrixadc_vlan" "tf_vlan" {
		vlanid    = 40
		aliasname = "Management VLAN"
	}
	resource "citrixadc_vxlan" "tf_vxlan" {
		vxlanid            = 123
		vlan               = citrixadc_vlan.tf_vlan.vlanid
		port               = 8080
		dynamicrouting     = "DISABLED"
		ipv6dynamicrouting = "DISABLED"
		innervlantagging   = "ENABLED"
	}
`

func TestAccVxlan_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVxlanDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVxlan_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVxlanExist("citrixadc_vxlan.tf_vxlan", nil),
					resource.TestCheckResourceAttr("citrixadc_vxlan.tf_vxlan", "vxlanid", "123"),
					resource.TestCheckResourceAttr("citrixadc_vxlan.tf_vxlan", "port", "33"),
				),
			},
			{
				Config: testAccVxlan_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVxlanExist("citrixadc_vxlan.tf_vxlan", nil),
					resource.TestCheckResourceAttr("citrixadc_vxlan.tf_vxlan", "vxlanid", "123"),
					resource.TestCheckResourceAttr("citrixadc_vxlan.tf_vxlan", "port", "8080"),
				),
			},
		},
	})
}

func TestAccVxlan_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_vxlan.tf_vxlan"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVxlanDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVxlan_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckVxlanExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Vxlan.Type(), "123"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccVxlan_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckVxlanExist(resAddr, nil)),
			},
		},
	})
}

func testAccCheckVxlanExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No vxlan name is set")
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
		data, err := client.FindResource(service.Vxlan.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("vxlan %s not found", n)
		}

		return nil
	}
}

func testAccCheckVxlanDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_vxlan" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Vxlan.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("vxlan %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccVxlan_import(t *testing.T) {
	const resAddr = "citrixadc_vxlan.tf_vxlan"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVxlanDestroy,
		Steps: []resource.TestStep{
			{Config: testAccVxlan_add},
			{
				Config:                  testAccVxlan_add,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

const testAccVxlanDataSource_basic = `
	resource "citrixadc_vlan" "tf_vlan" {
		vlanid    = 40
		aliasname = "Management VLAN"
	}
	resource "citrixadc_vxlan" "tf_vxlan" {
		vxlanid            = 123
		vlan               = citrixadc_vlan.tf_vlan.vlanid
		port               = 33
		dynamicrouting     = "DISABLED"
		ipv6dynamicrouting = "DISABLED"
		innervlantagging   = "ENABLED"
	}

data "citrixadc_vxlan" "tf_vxlan" {
	vxlanid = citrixadc_vxlan.tf_vxlan.vxlanid
}
`

func TestAccVxlan_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckVxlanDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccVxlan_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckVxlanExist("citrixadc_vxlan.tf_vxlan", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccVxlan_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckVxlanExist("citrixadc_vxlan.tf_vxlan", nil)),
			},
		},
	})
}

// testAccVxlan_unset_step1 sets the unset-eligible attributes to non-default
// values; step2 removes them so the provider must unset them back to the
// documented NITRO defaults (port=4789, dynamicrouting/ipv6dynamicrouting/
// innervlantagging=DISABLED).
const testAccVxlan_unset_step1 = `
	resource "citrixadc_vxlan" "tf_unset" {
		vxlanid            = 456
		port               = 8080
		dynamicrouting     = "ENABLED"
		ipv6dynamicrouting = "ENABLED"
		innervlantagging   = "ENABLED"
	}
`

const testAccVxlan_unset_step2 = `
	resource "citrixadc_vxlan" "tf_unset" {
		vxlanid = 456
		# All unset-eligible attributes removed from config -> the provider must
		# unset them (revert to NITRO defaults).
	}
`

func TestAccVxlan_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVxlanDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccVxlan_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVxlanExist("citrixadc_vxlan.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_vxlan.tf_unset", "port", "8080"),
					resource.TestCheckResourceAttr("citrixadc_vxlan.tf_unset", "dynamicrouting", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_vxlan.tf_unset", "ipv6dynamicrouting", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_vxlan.tf_unset", "innervlantagging", "ENABLED"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from
				// the appliance) reverts to the documented NITRO defaults, and the
				// implicit post-apply plan must be empty.
				Config: testAccVxlan_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVxlanExist("citrixadc_vxlan.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_vxlan.tf_unset", "port", "4789"),
					resource.TestCheckResourceAttr("citrixadc_vxlan.tf_unset", "dynamicrouting", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_vxlan.tf_unset", "ipv6dynamicrouting", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_vxlan.tf_unset", "innervlantagging", "DISABLED"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckVxlanADCValue("456", "port", "4789"),
					testAccCheckVxlanADCValue("456", "dynamicrouting", "DISABLED"),
					testAccCheckVxlanADCValue("456", "innervlantagging", "DISABLED"),
				),
			},
		},
	})
}

// testAccCheckVxlanADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted
// it.
func testAccCheckVxlanADCValue(id, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Vxlan.Type(), id)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("vxlan %s not found on appliance", id)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("vxlan %s: appliance attr %q = %q, want %q (unset did not revert it)", id, attr, got, want)
		}
		return nil
	}
}

func TestAccVxlanDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccVxlanDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_vxlan.tf_vxlan", "vxlanid", "123"),
					resource.TestCheckResourceAttr("data.citrixadc_vxlan.tf_vxlan", "vlan", "40"),
					resource.TestCheckResourceAttr("data.citrixadc_vxlan.tf_vxlan", "port", "33"),
					resource.TestCheckResourceAttr("data.citrixadc_vxlan.tf_vxlan", "dynamicrouting", "DISABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_vxlan.tf_vxlan", "ipv6dynamicrouting", "DISABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_vxlan.tf_vxlan", "innervlantagging", "ENABLED"),
				),
			},
		},
	})
}
