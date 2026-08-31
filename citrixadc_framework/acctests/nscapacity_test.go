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

func TestAccNscapacity_basic(t *testing.T) {
	t.Skip("Requires License Server Configuration.")
	// if isCpxRun {
	// 	t.Skip("Feature not supported in CPX")
	// }
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNscapacity_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNscapacityExist("citrixadc_nscapacity.tf_capacity", nil),
				),
			},
			// {
			// 	Config: testAccNscapacity_basic_step2,
			// 	Check: resource.ComposeTestCheckFunc(
			// 		testAccCheckNscapacityExist("citrixadc_nscapacity.tf_capacity", nil),
			// 	),
			// },
			// {
			// 	Config: testAccNscapacity_basic_step3,
			// 	Check: resource.ComposeTestCheckFunc(
			// 		testAccCheckNscapacityExist("citrixadc_nscapacity.tf_capacity", nil),
			// 	),
			// },
		},
	})
}

func testAccCheckNscapacityExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lb vserver name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		return nil
	}
}

const testAccNscapacity_basic_step1 = `
# Pooled license
resource "citrixadc_nscapacity" "tf_capacity" {
	bandwidth = 100
	unit = "Mbps"
	edition = "Platinum"
}
`

// const testAccNscapacity_basic_step2 = `
// # vCPU license
// resource "citrixadc_nscapacity" "tf_capacity" {
// 	vcpu = true
// 	edition = "Standard"
// }
// `

// const testAccNscapacity_basic_step3 = `
// # CICO license
// resource "citrixadc_nscapacity" "tf_capacity" {
// 	platform = "VP10000"
// }
// `

const testAccNscapacityDataSource_basic = `

	data "citrixadc_nscapacity" "tf_nscapacity" {
	}
`

func TestAccNscapacityDataSource_basic(t *testing.T) {
	t.Skip("Requires License Server Configuration.")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccNscapacityDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.citrixadc_nscapacity.tf_nscapacity", "id"),
				),
			},
		},
	})
}

// testAccNscapacity_unset_step1 applies a pooled-license capacity with a
// non-default bandwidth (spec-unsettable attribute).
const testAccNscapacity_unset_step1 = `
resource "citrixadc_nscapacity" "tf_unset" {
	bandwidth = 100
	unit      = "Mbps"
	edition   = "Platinum"
}
`

// testAccNscapacity_unset_step2 removes every optional attribute; the removed
// spec-unsettable "bandwidth" must be reverted to its appliance default via the
// NITRO ?action=unset operation.
const testAccNscapacity_unset_step2 = `
resource "citrixadc_nscapacity" "tf_unset" {
}
`

func TestAccNscapacity_unset(t *testing.T) {
	t.Skip("Requires License Server Configuration.")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				// Non-default value is applied and persisted.
				Config: testAccNscapacity_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNscapacityExist("citrixadc_nscapacity.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_nscapacity.tf_unset", "bandwidth", "100"),
				),
			},
			{
				// Removing the attribute must unset it: state (read back from the
				// appliance) reverts to the NITRO default, and the implicit
				// post-apply plan must be empty.
				Config: testAccNscapacity_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNscapacityExist("citrixadc_nscapacity.tf_unset", nil),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckNscapacityADCValue("bandwidth", "0"),
				),
			},
		},
	})
}

// testAccCheckNscapacityADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted
// it. nscapacity is a singleton, so it is fetched with an empty name.
func testAccCheckNscapacityADCValue(attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Nscapacity.Type(), "")
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("nscapacity not found on appliance")
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("nscapacity: appliance attr %q = %q, want %q (unset did not revert it)", attr, got, want)
		}
		return nil
	}
}

func TestAccNscapacity_import(t *testing.T) {
	t.Skip("Requires License Server Configuration.")
	const resAddr = "citrixadc_nscapacity.tf_capacity"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{Config: testAccNscapacity_basic_step1},
			{
				Config:                  testAccNscapacity_basic_step1,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccNscapacity_sdkv2StateUpgrade(t *testing.T) {
	t.Skip("Requires License Server Configuration.")
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.0.0"},
				},
				Config: testAccNscapacity_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNscapacityExist("citrixadc_nscapacity.tf_capacity", nil),
				),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccNscapacity_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNscapacityExist("citrixadc_nscapacity.tf_capacity", nil),
				),
			},
		},
	})
}
