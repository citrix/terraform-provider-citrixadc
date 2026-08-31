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

const testAccNsspparams_add = `

	resource "citrixadc_nsspparams" "tf_nsspparams" {
		basethreshold = 400
		throttle      = "Aggressive"
	}
`
const testAccNsspparams_update = `

	resource "citrixadc_nsspparams" "tf_nsspparams" {
		basethreshold = 200
		throttle      = "Normal"
	}
`

func TestAccNsspparams_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNsspparamsDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNsspparams_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsspparamsExist("citrixadc_nsspparams.tf_nsspparams", nil),
					resource.TestCheckResourceAttr("citrixadc_nsspparams.tf_nsspparams", "basethreshold", "400"),
					resource.TestCheckResourceAttr("citrixadc_nsspparams.tf_nsspparams", "throttle", "Aggressive"),
				),
			},
			{
				Config: testAccNsspparams_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsspparamsExist("citrixadc_nsspparams.tf_nsspparams", nil),
					resource.TestCheckResourceAttr("citrixadc_nsspparams.tf_nsspparams", "basethreshold", "200"),
					resource.TestCheckResourceAttr("citrixadc_nsspparams.tf_nsspparams", "throttle", "Normal"),
				),
			},
		},
	})
}

func testAccCheckNsspparamsExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No nsspparams name is set")
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
		data, err := client.FindResource(service.Nsspparams.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("nsspparams %s not found", n)
		}

		return nil
	}
}

func testAccCheckNsspparamsDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_nsspparams" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Nsspparams.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("nsspparams %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccNsspparams_import(t *testing.T) {
	const resAddr = "citrixadc_nsspparams.tf_nsspparams"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNsspparamsDestroy,
		Steps: []resource.TestStep{
			{Config: testAccNsspparams_add},
			{
				Config:                  testAccNsspparams_add,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccNsspparams_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckNsspparamsDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.0.0"},
				},
				Config: testAccNsspparams_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckNsspparamsExist("citrixadc_nsspparams.tf_nsspparams", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccNsspparams_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckNsspparamsExist("citrixadc_nsspparams.tf_nsspparams", nil)),
			},
		},
	})
}

// The nsspparams unset test exercises the two mutable read/write attributes
// (basethreshold, throttle). step1 sets them to non-default values; step2
// removes them from config so the provider unsets them, reverting to the
// documented NITRO defaults (basethreshold=200, throttle=Normal).
const testAccNsspparams_unset_step1 = `
	resource "citrixadc_nsspparams" "tf_unset" {
		basethreshold = 400
		throttle      = "Aggressive"
	}
`

const testAccNsspparams_unset_step2 = `
	resource "citrixadc_nsspparams" "tf_unset" {
		# All unset-eligible attributes removed from config -> the provider must
		# unset them (revert to NITRO defaults).
	}
`

func TestAccNsspparams_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNsspparamsDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccNsspparams_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsspparamsExist("citrixadc_nsspparams.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_nsspparams.tf_unset", "basethreshold", "400"),
					resource.TestCheckResourceAttr("citrixadc_nsspparams.tf_unset", "throttle", "Aggressive"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from
				// the appliance) reverts to the documented NITRO defaults, and the
				// implicit post-apply plan must be empty.
				Config: testAccNsspparams_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsspparamsExist("citrixadc_nsspparams.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_nsspparams.tf_unset", "basethreshold", "200"),
					resource.TestCheckResourceAttr("citrixadc_nsspparams.tf_unset", "throttle", "Normal"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckNsspparamsADCValue("basethreshold", "200"),
					testAccCheckNsspparamsADCValue("throttle", "Normal"),
				),
			},
		},
	})
}

// testAccCheckNsspparamsADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted it.
func testAccCheckNsspparamsADCValue(attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Nsspparams.Type(), "")
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("nsspparams not found on appliance")
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("nsspparams: appliance attr %q = %q, want %q (unset did not revert it)", attr, got, want)
		}
		return nil
	}
}

func TestAccNsspparamsDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccNsspparamsDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_nsspparams.tf_nsspparams_ds", "basethreshold", "350"),
					resource.TestCheckResourceAttr("data.citrixadc_nsspparams.tf_nsspparams_ds", "throttle", "Aggressive"),
				),
			},
		},
	})
}

const testAccNsspparamsDataSource_basic = `

	resource "citrixadc_nsspparams" "tf_nsspparams_ds" {
		basethreshold = 350
		throttle      = "Aggressive"
	}

	data "citrixadc_nsspparams" "tf_nsspparams_ds" {
		depends_on = [citrixadc_nsspparams.tf_nsspparams_ds]
	}
`
