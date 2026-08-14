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

const testAccNetbridge_add = `
	resource "citrixadc_vxlanvlanmap" "tf_vxlanvlanmp" {
		name = "tf_vxlanvlanmp"
	}
	resource "citrixadc_vxlanvlanmap" "tf_vxlanvlanmp1" {
		name = "tf_vxlanvlanmpsample"
	}
	resource "citrixadc_netbridge" "tf_netbridge" {
		name         = "tf_netbridge"
		vxlanvlanmap = citrixadc_vxlanvlanmap.tf_vxlanvlanmp.name
	}
`
const testAccNetbridge_update = `
	resource "citrixadc_vxlanvlanmap" "tf_vxlanvlanmp" {
		name = "tf_vxlanvlanmp"
	}
	resource "citrixadc_vxlanvlanmap" "tf_vxlanvlanmp1" {
		name = "tf_vxlanvlanmpsample"
	}
	resource "citrixadc_netbridge" "tf_netbridge" {
		name         = "tf_netbridge"
		vxlanvlanmap = citrixadc_vxlanvlanmap.tf_vxlanvlanmp1.name
	}
`

const testAccNetbridgeDataSource_basic = `
	resource "citrixadc_vxlanvlanmap" "tf_vxlanvlanmp" {
		name = "tf_vxlanvlanmp"
	}
	resource "citrixadc_netbridge" "tf_netbridge" {
		name         = "tf_netbridge_ds"
		vxlanvlanmap = citrixadc_vxlanvlanmap.tf_vxlanvlanmp.name
	}

	data "citrixadc_netbridge" "tf_netbridge_ds" {
		name = citrixadc_netbridge.tf_netbridge.name
	}
`

func TestAccNetbridge_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNetbridgeDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNetbridge_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetbridgeExist("citrixadc_netbridge.tf_netbridge", nil),
					resource.TestCheckResourceAttr("citrixadc_netbridge.tf_netbridge", "name", "tf_netbridge"),
					resource.TestCheckResourceAttr("citrixadc_netbridge.tf_netbridge", "vxlanvlanmap", "tf_vxlanvlanmp"),
				),
			},
			{
				Config: testAccNetbridge_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetbridgeExist("citrixadc_netbridge.tf_netbridge", nil),
					resource.TestCheckResourceAttr("citrixadc_netbridge.tf_netbridge", "name", "tf_netbridge"),
					resource.TestCheckResourceAttr("citrixadc_netbridge.tf_netbridge", "vxlanvlanmap", "tf_vxlanvlanmpsample"),
				),
			},
		},
	})
}

func testAccCheckNetbridgeExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No netbridge name is set")
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
		data, err := client.FindResource(service.Netbridge.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("netbridge %s not found", n)
		}

		return nil
	}
}

func testAccCheckNetbridgeDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_netbridge" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Netbridge.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("netbridge %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccNetbridge_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_netbridge.tf_netbridge"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNetbridgeDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNetbridge_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckNetbridgeExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Netbridge.Type(), "tf_netbridge"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccNetbridge_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckNetbridgeExist(resAddr, nil)),
			},
		},
	})
}

func TestAccNetbridge_import(t *testing.T) {
	const resAddr = "citrixadc_netbridge.tf_netbridge"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNetbridgeDestroy,
		Steps: []resource.TestStep{
			{Config: testAccNetbridge_add},
			{
				Config:                  testAccNetbridge_add,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccNetbridge_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckNetbridgeDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccNetbridge_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckNetbridgeExist("citrixadc_netbridge.tf_netbridge", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccNetbridge_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckNetbridgeExist("citrixadc_netbridge.tf_netbridge", nil)),
			},
		},
	})
}

// The netbridge unset test covers vxlanvlanmap, the only unset-eligible
// attribute. step1 sets it to a real vxlanvlanmap; step2 removes it from
// config, requiring the provider to unset it (revert to the NITRO default,
// which is empty).
const testAccNetbridge_unset_step1 = `
	resource "citrixadc_vxlanvlanmap" "tf_vxlanvlanmp" {
		name = "tf_netbridge_unset_map"
	}
	resource "citrixadc_netbridge" "tf_unset" {
		name         = "tf_netbridge_unset"
		vxlanvlanmap = citrixadc_vxlanvlanmap.tf_vxlanvlanmp.name
	}
`

const testAccNetbridge_unset_step2 = `
	resource "citrixadc_vxlanvlanmap" "tf_vxlanvlanmp" {
		name = "tf_netbridge_unset_map"
	}
	resource "citrixadc_netbridge" "tf_unset" {
		name = "tf_netbridge_unset"
		# vxlanvlanmap removed from config -> the provider must unset it.
	}
`

func TestAccNetbridge_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNetbridgeDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default value is applied and persisted.
				Config: testAccNetbridge_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetbridgeExist("citrixadc_netbridge.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_netbridge.tf_unset", "vxlanvlanmap", "tf_netbridge_unset_map"),
				),
			},
			{
				// Removing the attribute must unset it: state reverts to the NITRO
				// default (empty) and the implicit post-apply plan must be empty.
				Config: testAccNetbridge_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetbridgeExist("citrixadc_netbridge.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_netbridge.tf_unset", "vxlanvlanmap", ""),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckNetbridgeADCValue("tf_netbridge_unset", "vxlanvlanmap", ""),
				),
			},
		},
	})
}

// testAccCheckNetbridgeADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted
// it.
func testAccCheckNetbridgeADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Netbridge.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("netbridge %s not found on appliance", name)
		}
		got := ""
		if v, ok := data[attr]; ok && v != nil {
			got = strings.TrimSpace(fmt.Sprintf("%v", v))
		}
		if got != want {
			return fmt.Errorf("netbridge %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func TestAccNetbridgeDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNetbridgeDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNetbridgeDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_netbridge.tf_netbridge_ds", "name", "tf_netbridge_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_netbridge.tf_netbridge_ds", "vxlanvlanmap", "tf_vxlanvlanmp"),
				),
			},
		},
	})
}
