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

const testAccVrid6_add = `
	resource "citrixadc_vrid6" "tf_vrid6" {
		vrid6_id             = 3
		priority             = 30
		preemption           = "DISABLED"
		sharing              = "DISABLED"
		tracking             = "NONE"
		trackifnumpriority   = 0
		preemptiondelaytimer = 0
	}
`
const testAccVrid6_update = `
	resource "citrixadc_vrid6" "tf_vrid6" {
		vrid6_id             = 3
		priority             = 50
		preemption           = "ENABLED"
		sharing              = "DISABLED"
		tracking             = "NONE"
		trackifnumpriority   = 0
		preemptiondelaytimer = 0
	}
`

func TestAccVrid6_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVrid6Destroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVrid6_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVrid6Exist("citrixadc_vrid6.tf_vrid6", nil),
					resource.TestCheckResourceAttr("citrixadc_vrid6.tf_vrid6", "vrid6_id", "3"),
					resource.TestCheckResourceAttr("citrixadc_vrid6.tf_vrid6", "priority", "30"),
					resource.TestCheckResourceAttr("citrixadc_vrid6.tf_vrid6", "preemption", "DISABLED"),
				),
			},
			{
				Config: testAccVrid6_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVrid6Exist("citrixadc_vrid6.tf_vrid6", nil),
					resource.TestCheckResourceAttr("citrixadc_vrid6.tf_vrid6", "vrid6_id", "3"),
					resource.TestCheckResourceAttr("citrixadc_vrid6.tf_vrid6", "priority", "50"),
					resource.TestCheckResourceAttr("citrixadc_vrid6.tf_vrid6", "preemption", "ENABLED"),
				),
			},
		},
	})
}

func TestAccVrid6_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_vrid6.tf_vrid6"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVrid6Destroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVrid6_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckVrid6Exist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Vrid6.Type(), "3"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccVrid6_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckVrid6Exist(resAddr, nil)),
			},
		},
	})
}

func TestAccVrid6_import(t *testing.T) {
	const resAddr = "citrixadc_vrid6.tf_vrid6"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVrid6Destroy,
		Steps: []resource.TestStep{
			{Config: testAccVrid6_add},
			{
				Config:                  testAccVrid6_add,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"preemptiondelaytimer", "trackifnumpriority"},
			},
		},
	})
}

func testAccCheckVrid6Exist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No vrid6 name is set")
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
		data, err := client.FindResource(service.Vrid6.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("vrid6 %s not found", n)
		}

		return nil
	}
}

func testAccCheckVrid6Destroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_vrid6" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Vrid6.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("vrid6 %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccVrid6_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckVrid6Destroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccVrid6_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckVrid6Exist("citrixadc_vrid6.tf_vrid6", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccVrid6_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckVrid6Exist("citrixadc_vrid6.tf_vrid6", nil)),
			},
		},
	})
}

// Unset test: step1 sets the unset-eligible attributes to non-default values,
// step2 removes them so the provider must unset them (revert to NITRO defaults).
const testAccVrid6_unset_step1 = `
	resource "citrixadc_vrid6" "tf_unset" {
		vrid6_id             = 7
		priority             = 100
		preemption           = "DISABLED"
		sharing              = "ENABLED"
		tracking             = "ALL"
		preemptiondelaytimer = 10
		trackifnumpriority   = 5
	}
`

const testAccVrid6_unset_step2 = `
	resource "citrixadc_vrid6" "tf_unset" {
		vrid6_id = 7
		# All unset-eligible attributes removed from config -> the provider must
		# unset them (revert to NITRO defaults).
	}
`

func TestAccVrid6_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVrid6Destroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccVrid6_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVrid6Exist("citrixadc_vrid6.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_vrid6.tf_unset", "priority", "100"),
					resource.TestCheckResourceAttr("citrixadc_vrid6.tf_unset", "preemption", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_vrid6.tf_unset", "sharing", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_vrid6.tf_unset", "tracking", "ALL"),
					resource.TestCheckResourceAttr("citrixadc_vrid6.tf_unset", "preemptiondelaytimer", "10"),
					resource.TestCheckResourceAttr("citrixadc_vrid6.tf_unset", "trackifnumpriority", "5"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from the
				// appliance) reverts to the documented NITRO defaults, and the
				// implicit post-apply plan must be empty.
				Config: testAccVrid6_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVrid6Exist("citrixadc_vrid6.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_vrid6.tf_unset", "priority", "255"),
					resource.TestCheckResourceAttr("citrixadc_vrid6.tf_unset", "preemption", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_vrid6.tf_unset", "sharing", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_vrid6.tf_unset", "tracking", "NONE"),
					resource.TestCheckResourceAttr("citrixadc_vrid6.tf_unset", "preemptiondelaytimer", "0"),
					resource.TestCheckResourceAttr("citrixadc_vrid6.tf_unset", "trackifnumpriority", "0"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckVrid6ADCValue("7", "priority", "255"),
					testAccCheckVrid6ADCValue("7", "preemption", "ENABLED"),
					testAccCheckVrid6ADCValue("7", "tracking", "NONE"),
				),
			},
		},
	})
}

// testAccCheckVrid6ADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted
// it.
func testAccCheckVrid6ADCValue(id, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Vrid6.Type(), id)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("vrid6 %s not found on appliance", id)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("vrid6 %s: appliance attr %q = %q, want %q (unset did not revert it)", id, attr, got, want)
		}
		return nil
	}
}

const testAccVrid6DataSource_basic = `
	resource "citrixadc_vrid6" "tf_vrid6" {
		vrid6_id             = 3
		priority             = 30
		preemption           = "DISABLED"
		sharing              = "DISABLED"
		tracking             = "NONE"
	}

data "citrixadc_vrid6" "tf_vrid6" {
	vrid6_id = citrixadc_vrid6.tf_vrid6.vrid6_id
}
`

func TestAccVrid6DataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccVrid6DataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_vrid6.tf_vrid6", "vrid6_id", "3"),
					resource.TestCheckResourceAttr("data.citrixadc_vrid6.tf_vrid6", "priority", "30"),
					resource.TestCheckResourceAttr("data.citrixadc_vrid6.tf_vrid6", "preemption", "DISABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_vrid6.tf_vrid6", "sharing", "DISABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_vrid6.tf_vrid6", "tracking", "NONE"),
				),
			},
		},
	})
}
