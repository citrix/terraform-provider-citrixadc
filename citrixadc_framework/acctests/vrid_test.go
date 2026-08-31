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

const testAccVrid_add = `
	resource "citrixadc_vrid" "tf_vrid" {
		vrid_id              = 3
		priority             = 30
		preemption           = "DISABLED"
		sharing              = "ENABLED"
		tracking             = "NONE"
		trackifnumpriority   = 0
		preemptiondelaytimer = 0
	}
`
const testAccVrid_update = `
	resource "citrixadc_vrid" "tf_vrid" {
		vrid_id              = 3
		priority             = 20
		preemption           = "ENABLED"
		sharing              = "ENABLED"
		tracking             = "NONE"
		trackifnumpriority   = 0
		preemptiondelaytimer = 0
	}
`

func TestAccVrid_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVridDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVrid_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVridExist("citrixadc_vrid.tf_vrid", nil),
					resource.TestCheckResourceAttr("citrixadc_vrid.tf_vrid", "vrid_id", "3"),
					resource.TestCheckResourceAttr("citrixadc_vrid.tf_vrid", "priority", "30"),
					resource.TestCheckResourceAttr("citrixadc_vrid.tf_vrid", "preemption", "DISABLED"),
				),
			},
			{
				Config: testAccVrid_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVridExist("citrixadc_vrid.tf_vrid", nil),
					resource.TestCheckResourceAttr("citrixadc_vrid.tf_vrid", "vrid_id", "3"),
					resource.TestCheckResourceAttr("citrixadc_vrid.tf_vrid", "priority", "20"),
					resource.TestCheckResourceAttr("citrixadc_vrid.tf_vrid", "preemption", "ENABLED"),
				),
			},
		},
	})
}

func TestAccVrid_import(t *testing.T) {
	const resAddr = "citrixadc_vrid.tf_vrid"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVridDestroy,
		Steps: []resource.TestStep{
			{Config: testAccVrid_add},
			{
				Config:                  testAccVrid_add,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"preemptiondelaytimer", "trackifnumpriority"},
			},
		},
	})
}

func testAccCheckVridExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No vrid name is set")
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
		data, err := client.FindResource(service.Vrid.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("vrid %s not found", n)
		}

		return nil
	}
}

func testAccCheckVridDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_vrid" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Vrid.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("vrid %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccVrid_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_vrid.tf_vrid"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVridDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVrid_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckVridExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Vrid.Type(), "3"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccVrid_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckVridExist(resAddr, nil)),
			},
		},
	})
}

const testAccVridDataSource_basic = `
	resource "citrixadc_vrid" "tf_vrid" {
		vrid_id              = 3
		priority             = 30
		preemption           = "DISABLED"
		sharing              = "ENABLED"
		tracking             = "NONE"
	}

data "citrixadc_vrid" "tf_vrid" {
	vrid_id = citrixadc_vrid.tf_vrid.vrid_id
}
`

func TestAccVrid_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckVridDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.0.0"},
				},
				Config: testAccVrid_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckVridExist("citrixadc_vrid.tf_vrid", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccVrid_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckVridExist("citrixadc_vrid.tf_vrid", nil)),
			},
		},
	})
}

// testAccVrid_unset_step1 sets every unset-eligible attribute to a valid
// non-default value; step2 removes them so the provider must unset them and the
// appliance reverts to the documented NITRO defaults.
const testAccVrid_unset_step1 = `
	resource "citrixadc_vrid" "tf_unset" {
		vrid_id              = 5
		priority             = 100
		preemption           = "DISABLED"
		sharing              = "ENABLED"
		tracking             = "ALL"
		trackifnumpriority   = 50
		preemptiondelaytimer = 100
	}
`

const testAccVrid_unset_step2 = `
	resource "citrixadc_vrid" "tf_unset" {
		vrid_id = 5
		# All unset-eligible attributes removed from config -> the provider must
		# unset them (revert to NITRO defaults).
	}
`

func TestAccVrid_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVridDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccVrid_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVridExist("citrixadc_vrid.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_vrid.tf_unset", "priority", "100"),
					resource.TestCheckResourceAttr("citrixadc_vrid.tf_unset", "preemption", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_vrid.tf_unset", "sharing", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_vrid.tf_unset", "tracking", "ALL"),
					resource.TestCheckResourceAttr("citrixadc_vrid.tf_unset", "trackifnumpriority", "50"),
					resource.TestCheckResourceAttr("citrixadc_vrid.tf_unset", "preemptiondelaytimer", "100"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from
				// the appliance) reverts to the documented NITRO defaults, and the
				// implicit post-apply plan must be empty.
				Config: testAccVrid_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVridExist("citrixadc_vrid.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_vrid.tf_unset", "priority", "255"),
					resource.TestCheckResourceAttr("citrixadc_vrid.tf_unset", "preemption", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_vrid.tf_unset", "sharing", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_vrid.tf_unset", "tracking", "NONE"),
					resource.TestCheckResourceAttr("citrixadc_vrid.tf_unset", "trackifnumpriority", "0"),
					resource.TestCheckResourceAttr("citrixadc_vrid.tf_unset", "preemptiondelaytimer", "0"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckVridADCValue("5", "priority", "255"),
					testAccCheckVridADCValue("5", "preemption", "ENABLED"),
					testAccCheckVridADCValue("5", "sharing", "DISABLED"),
					testAccCheckVridADCValue("5", "tracking", "NONE"),
				),
			},
		},
	})
}

// testAccCheckVridADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted it.
func testAccCheckVridADCValue(id, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Vrid.Type(), id)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("vrid %s not found on appliance", id)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("vrid %s: appliance attr %q = %q, want %q (unset did not revert it)", id, attr, got, want)
		}
		return nil
	}
}

func TestAccVridDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccVridDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_vrid.tf_vrid", "vrid_id", "3"),
					resource.TestCheckResourceAttr("data.citrixadc_vrid.tf_vrid", "priority", "30"),
					resource.TestCheckResourceAttr("data.citrixadc_vrid.tf_vrid", "preemption", "DISABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_vrid.tf_vrid", "sharing", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_vrid.tf_vrid", "tracking", "NONE"),
				),
			},
		},
	})
}
