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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// videooptimizationparameter is a SINGLETON resource:
//   - Static ID "videooptimizationparameter-config"
//   - Create/Update both via PUT (UpdateUnnamedResource); Delete is state-only (no NITRO delete)
//   - No CheckDestroy in the TestCase (the resource always exists on the ADC)
//
// Note: quicpacingrate is intentionally omitted from the config and assertions because it
// requires the VideoOptimization feature to be enabled on the ADC. On a testbed without that
// feature enabled, the appliance does not store/return quicpacingrate, so the GET response omits
// it and Terraform reports "inconsistent result after apply" (config value vs. null read-back).
// Re-add quicpacingrate to the config/assertions when running against an appliance with the
// VideoOptimization feature enabled. randomsamplingpercentage is the CLI-confirmed, safe primary
// attribute and is exercised across both steps to test the update/PUT path.

const testAccVideooptimizationparameter_basic_step1 = `

	resource "citrixadc_videooptimizationparameter" "tf_videooptimizationparameter" {
		randomsamplingpercentage = 10
	}

`

const testAccVideooptimizationparameter_basic_step2 = `

	resource "citrixadc_videooptimizationparameter" "tf_videooptimizationparameter" {
		randomsamplingpercentage = 25
	}

`

func TestAccVideooptimizationparameter_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccVideooptimizationparameter_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVideooptimizationparameterExist("citrixadc_videooptimizationparameter.tf_videooptimizationparameter", nil),
					resource.TestCheckResourceAttr("citrixadc_videooptimizationparameter.tf_videooptimizationparameter", "randomsamplingpercentage", "10"),
				),
			},
			{
				Config: testAccVideooptimizationparameter_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVideooptimizationparameterExist("citrixadc_videooptimizationparameter.tf_videooptimizationparameter", nil),
					resource.TestCheckResourceAttr("citrixadc_videooptimizationparameter.tf_videooptimizationparameter", "randomsamplingpercentage", "25"),
				),
			},
		},
	})
}

func TestAccVideooptimizationparameter_import(t *testing.T) {
	const resAddr = "citrixadc_videooptimizationparameter.tf_videooptimizationparameter"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{Config: testAccVideooptimizationparameter_basic_step1},
			{
				Config:                  testAccVideooptimizationparameter_basic_step1,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccCheckVideooptimizationparameterExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No videooptimizationparameter name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		// Get a configured client from the test helper
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		// Singleton resource: empty name passed to FindResource
		data, err := client.FindResource(service.Videooptimizationparameter.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("videooptimizationparameter %s not found", n)
		}

		return nil
	}
}

// --- Unset support test ---------------------------------------------------
//
// videooptimizationparameter has unset support wired for exactly one attribute:
//   randomsamplingpercentage (Float64, NITRO default 0).
// Step 1 sets it to a non-default value; step 2 removes it from config so the
// provider issues ?action=unset and the appliance reverts it to its default (0).
// This is a singleton resource, so there is no CheckDestroy and the resource
// name passed to FindResource is empty (mirrors testAccCheckVideooptimizationparameterExist).

const testAccVideooptimizationparameter_unset_step1 = `

	resource "citrixadc_videooptimizationparameter" "tf_unset" {
		randomsamplingpercentage = 25
	}

`

const testAccVideooptimizationparameter_unset_step2 = `

	resource "citrixadc_videooptimizationparameter" "tf_unset" {
		# randomsamplingpercentage removed from config -> provider must unset it
	}

`

func TestAccVideooptimizationparameter_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				// Non-default value applies and persists.
				Config: testAccVideooptimizationparameter_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVideooptimizationparameterExist("citrixadc_videooptimizationparameter.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_videooptimizationparameter.tf_unset", "randomsamplingpercentage", "25"),
				),
			},
			{
				// Removing it must unset -> state reverts to the NITRO default,
				// and the implicit post-apply plan must be empty.
				Config: testAccVideooptimizationparameter_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVideooptimizationparameterExist("citrixadc_videooptimizationparameter.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_videooptimizationparameter.tf_unset", "randomsamplingpercentage", "0"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckVideooptimizationparameterADCValue("randomsamplingpercentage", "0"),
				),
			},
		},
	})
}

// testAccCheckVideooptimizationparameterADCValue asserts an attribute's value
// directly on the appliance (not just in Terraform state), proving the unset
// actually reverted it. Singleton resource: FindResource takes an empty name.
func testAccCheckVideooptimizationparameterADCValue(attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Videooptimizationparameter.Type(), "")
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("videooptimizationparameter not found on appliance")
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("videooptimizationparameter: appliance attr %q = %q, want %q (unset did not revert it)", attr, got, want)
		}
		return nil
	}
}

const testAccVideooptimizationparameterDataSource_basic = `

	resource "citrixadc_videooptimizationparameter" "tf_videooptimizationparameter" {
		randomsamplingpercentage = 10
	}

	data "citrixadc_videooptimizationparameter" "tf_videooptimizationparameter" {
		depends_on = [citrixadc_videooptimizationparameter.tf_videooptimizationparameter]
	}
`

func TestAccVideooptimizationparameterDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccVideooptimizationparameterDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_videooptimizationparameter.tf_videooptimizationparameter", "randomsamplingpercentage", "10"),
				),
			},
		},
	})
}
