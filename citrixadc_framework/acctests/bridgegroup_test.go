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

const testAccBridgegroup_add = `
	resource "citrixadc_bridgegroup" "tf_bridgegroup" {
		bridgegroup_id     = 2
		dynamicrouting     = "DISABLED"
		ipv6dynamicrouting = "DISABLED"
	}
`
const testAccBridgegroup_update = `
	resource "citrixadc_bridgegroup" "tf_bridgegroup" {
		bridgegroup_id     = 2
		dynamicrouting     = "ENABLED"
		ipv6dynamicrouting = "ENABLED"
	}
`

const testAccBridgegroupDataSource_basic = `
	resource "citrixadc_bridgegroup" "tf_bridgegroup" {
		bridgegroup_id     = 2
		dynamicrouting     = "DISABLED"
		ipv6dynamicrouting = "DISABLED"
	}

	data "citrixadc_bridgegroup" "tf_bridgegroup_ds" {
		depends_on        = [citrixadc_bridgegroup.tf_bridgegroup]
		bridgegroup_id    = citrixadc_bridgegroup.tf_bridgegroup.bridgegroup_id
	}
`

func TestAccBridgegroup_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckBridgegroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccBridgegroup_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBridgegroupExist("citrixadc_bridgegroup.tf_bridgegroup", nil),
					resource.TestCheckResourceAttr("citrixadc_bridgegroup.tf_bridgegroup", "bridgegroup_id", "2"),
					resource.TestCheckResourceAttr("citrixadc_bridgegroup.tf_bridgegroup", "dynamicrouting", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_bridgegroup.tf_bridgegroup", "ipv6dynamicrouting", "DISABLED"),
				),
			},
			{
				Config: testAccBridgegroup_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBridgegroupExist("citrixadc_bridgegroup.tf_bridgegroup", nil),
					resource.TestCheckResourceAttr("citrixadc_bridgegroup.tf_bridgegroup", "bridgegroup_id", "2"),
					resource.TestCheckResourceAttr("citrixadc_bridgegroup.tf_bridgegroup", "dynamicrouting", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_bridgegroup.tf_bridgegroup", "ipv6dynamicrouting", "ENABLED"),
				),
			},
		},
	})
}

func testAccCheckBridgegroupExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No bridgegroup name is set")
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
		data, err := client.FindResource(service.Bridgegroup.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("bridgegroup %s not found", n)
		}

		return nil
	}
}

func testAccCheckBridgegroupDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_bridgegroup" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Bridgegroup.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("bridgegroup %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccBridgegroup_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_bridgegroup.tf_bridgegroup"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckBridgegroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccBridgegroup_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckBridgegroupExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Bridgegroup.Type(), "2"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccBridgegroup_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckBridgegroupExist(resAddr, nil)),
			},
		},
	})
}

func TestAccBridgegroup_import(t *testing.T) {
	const resAddr = "citrixadc_bridgegroup.tf_bridgegroup"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckBridgegroupDestroy,
		Steps: []resource.TestStep{
			{Config: testAccBridgegroup_add},
			{
				Config:                  testAccBridgegroup_add,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccBridgegroup_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckBridgegroupDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.0.0"},
				},
				Config: testAccBridgegroup_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckBridgegroupExist("citrixadc_bridgegroup.tf_bridgegroup", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccBridgegroup_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckBridgegroupExist("citrixadc_bridgegroup.tf_bridgegroup", nil)),
			},
		},
	})
}

const testAccBridgegroup_unset_step1 = `
	resource "citrixadc_bridgegroup" "tf_unset" {
		bridgegroup_id     = 3
		dynamicrouting     = "ENABLED"
		ipv6dynamicrouting = "ENABLED"
	}
`

const testAccBridgegroup_unset_step2 = `
	resource "citrixadc_bridgegroup" "tf_unset" {
		bridgegroup_id = 3
		# unset-eligible attributes removed from config -> provider must unset
		# them (revert to NITRO defaults, "DISABLED").
	}
`

func TestAccBridgegroup_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckBridgegroupDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccBridgegroup_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBridgegroupExist("citrixadc_bridgegroup.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_bridgegroup.tf_unset", "dynamicrouting", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_bridgegroup.tf_unset", "ipv6dynamicrouting", "ENABLED"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from
				// the appliance) reverts to the documented NITRO defaults, and the
				// implicit post-apply plan must be empty.
				Config: testAccBridgegroup_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBridgegroupExist("citrixadc_bridgegroup.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_bridgegroup.tf_unset", "dynamicrouting", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_bridgegroup.tf_unset", "ipv6dynamicrouting", "DISABLED"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckBridgegroupADCValue("3", "dynamicrouting", "DISABLED"),
					testAccCheckBridgegroupADCValue("3", "ipv6dynamicrouting", "DISABLED"),
				),
			},
		},
	})
}

// testAccCheckBridgegroupADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted it.
func testAccCheckBridgegroupADCValue(id, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Bridgegroup.Type(), id)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("bridgegroup %s not found on appliance", id)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("bridgegroup %s: appliance attr %q = %q, want %q (unset did not revert it)", id, attr, got, want)
		}
		return nil
	}
}

func TestAccBridgegroupDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckBridgegroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccBridgegroupDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.citrixadc_bridgegroup.tf_bridgegroup_ds", "id"),
					resource.TestCheckResourceAttrPair("data.citrixadc_bridgegroup.tf_bridgegroup_ds", "bridgegroup_id", "citrixadc_bridgegroup.tf_bridgegroup", "bridgegroup_id"),
					resource.TestCheckResourceAttrPair("data.citrixadc_bridgegroup.tf_bridgegroup_ds", "dynamicrouting", "citrixadc_bridgegroup.tf_bridgegroup", "dynamicrouting"),
					resource.TestCheckResourceAttrPair("data.citrixadc_bridgegroup.tf_bridgegroup_ds", "ipv6dynamicrouting", "citrixadc_bridgegroup.tf_bridgegroup", "ipv6dynamicrouting"),
				),
			},
		},
	})
}
