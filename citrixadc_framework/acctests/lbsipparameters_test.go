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

const testAccLbsipparameters_basic = `
	resource "citrixadc_lbsipparameters" "tf_lbsipparameters" {
		addrportvip = "ENABLED"
		retrydur = 100
		rnatdstport = 80
		rnatsecuredstport = 81
		rnatsecuresrcport = 82
		rnatsrcport = 83
		sip503ratethreshold = 15
	}
`

const testAccLbsipparameters_basic_update = `
	resource "citrixadc_lbsipparameters" "tf_lbsipparameters" {
		addrportvip = "DISABLED"
		retrydur = 120
		rnatdstport = 1
		rnatsecuredstport = 2
		rnatsecuresrcport = 3
		rnatsrcport = 4
		sip503ratethreshold = 100
	}
`

func TestAccLbsipparameters_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// lbsipparameters resource do not have DELETE operation
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				Config: testAccLbsipparameters_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbsipparametersExist("citrixadc_lbsipparameters.tf_lbsipparameters", nil),
					resource.TestCheckResourceAttr("citrixadc_lbsipparameters.tf_lbsipparameters", "addrportvip", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_lbsipparameters.tf_lbsipparameters", "retrydur", "100"),
					resource.TestCheckResourceAttr("citrixadc_lbsipparameters.tf_lbsipparameters", "rnatdstport", "80"),
					resource.TestCheckResourceAttr("citrixadc_lbsipparameters.tf_lbsipparameters", "rnatsecuredstport", "81"),
					resource.TestCheckResourceAttr("citrixadc_lbsipparameters.tf_lbsipparameters", "rnatsecuresrcport", "82"),
					resource.TestCheckResourceAttr("citrixadc_lbsipparameters.tf_lbsipparameters", "rnatsrcport", "83"),
					resource.TestCheckResourceAttr("citrixadc_lbsipparameters.tf_lbsipparameters", "sip503ratethreshold", "15"),
				),
			},
			{
				Config: testAccLbsipparameters_basic_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbsipparametersExist("citrixadc_lbsipparameters.tf_lbsipparameters", nil),
					resource.TestCheckResourceAttr("citrixadc_lbsipparameters.tf_lbsipparameters", "addrportvip", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_lbsipparameters.tf_lbsipparameters", "retrydur", "120"),
					resource.TestCheckResourceAttr("citrixadc_lbsipparameters.tf_lbsipparameters", "rnatdstport", "1"),
					resource.TestCheckResourceAttr("citrixadc_lbsipparameters.tf_lbsipparameters", "rnatsecuredstport", "2"),
					resource.TestCheckResourceAttr("citrixadc_lbsipparameters.tf_lbsipparameters", "rnatsecuresrcport", "3"),
					resource.TestCheckResourceAttr("citrixadc_lbsipparameters.tf_lbsipparameters", "rnatsrcport", "4"),
					resource.TestCheckResourceAttr("citrixadc_lbsipparameters.tf_lbsipparameters", "sip503ratethreshold", "100"),
				),
			},
		},
	})
}

func testAccCheckLbsipparametersExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lbsipparameters name is set")
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
		data, err := client.FindResource(service.Lbsipparameters.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("lbsipparameters %s not found", n)
		}

		return nil
	}
}

func TestAccLbsipparameters_import(t *testing.T) {
	const resAddr = "citrixadc_lbsipparameters.tf_lbsipparameters"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{Config: testAccLbsipparameters_basic},
			{
				Config:                  testAccLbsipparameters_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

const testAccLbsipparametersDataSource_basic = `
	resource "citrixadc_lbsipparameters" "tf_lbsipparameters" {
		addrportvip = "ENABLED"
		retrydur = 100
		rnatdstport = 80
		rnatsecuredstport = 81
		rnatsecuresrcport = 82
		rnatsrcport = 83
		sip503ratethreshold = 15
	}

	data "citrixadc_lbsipparameters" "tf_lbsipparameters" {
		depends_on = [citrixadc_lbsipparameters.tf_lbsipparameters]
	}
`

func TestAccLbsipparameters_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccLbsipparameters_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbsipparametersExist("citrixadc_lbsipparameters.tf_lbsipparameters", nil),
				),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccLbsipparameters_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbsipparametersExist("citrixadc_lbsipparameters.tf_lbsipparameters", nil),
				),
			},
		},
	})
}

// The lbsipparameters unset test sets every unset-eligible attribute to a
// non-default value, then removes them all from config. The provider must
// unset each one so the appliance reverts to the documented NITRO defaults.
const testAccLbsipparameters_unset_step1 = `
	resource "citrixadc_lbsipparameters" "tf_unset" {
		addrportvip         = "DISABLED"
		retrydur            = 100
		rnatdstport         = 80
		rnatsecuredstport   = 81
		rnatsecuresrcport   = 82
		rnatsrcport         = 83
		sip503ratethreshold = 15
	}
`

const testAccLbsipparameters_unset_step2 = `
	resource "citrixadc_lbsipparameters" "tf_unset" {
		# All unset-eligible attributes removed from config -> the provider must
		# unset them (revert to NITRO defaults).
	}
`

func TestAccLbsipparameters_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccLbsipparameters_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbsipparametersExist("citrixadc_lbsipparameters.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_lbsipparameters.tf_unset", "addrportvip", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_lbsipparameters.tf_unset", "retrydur", "100"),
					resource.TestCheckResourceAttr("citrixadc_lbsipparameters.tf_unset", "rnatdstport", "80"),
					resource.TestCheckResourceAttr("citrixadc_lbsipparameters.tf_unset", "rnatsecuredstport", "81"),
					resource.TestCheckResourceAttr("citrixadc_lbsipparameters.tf_unset", "rnatsecuresrcport", "82"),
					resource.TestCheckResourceAttr("citrixadc_lbsipparameters.tf_unset", "rnatsrcport", "83"),
					resource.TestCheckResourceAttr("citrixadc_lbsipparameters.tf_unset", "sip503ratethreshold", "15"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from
				// the appliance) reverts to the documented NITRO defaults, and the
				// implicit post-apply plan must be empty.
				Config: testAccLbsipparameters_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbsipparametersExist("citrixadc_lbsipparameters.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_lbsipparameters.tf_unset", "addrportvip", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_lbsipparameters.tf_unset", "retrydur", "120"),
					resource.TestCheckResourceAttr("citrixadc_lbsipparameters.tf_unset", "rnatdstport", "0"),
					resource.TestCheckResourceAttr("citrixadc_lbsipparameters.tf_unset", "rnatsecuredstport", "0"),
					resource.TestCheckResourceAttr("citrixadc_lbsipparameters.tf_unset", "rnatsecuresrcport", "0"),
					resource.TestCheckResourceAttr("citrixadc_lbsipparameters.tf_unset", "rnatsrcport", "0"),
					resource.TestCheckResourceAttr("citrixadc_lbsipparameters.tf_unset", "sip503ratethreshold", "100"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckLbsipparametersADCValue("addrportvip", "ENABLED"),
					testAccCheckLbsipparametersADCValue("retrydur", "120"),
				),
			},
		},
	})
}

// testAccCheckLbsipparametersADCValue asserts an attribute's value directly on
// the appliance (not just in Terraform state), proving the unset actually
// reverted it.
func testAccCheckLbsipparametersADCValue(attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Lbsipparameters.Type(), "")
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("lbsipparameters not found on appliance")
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("lbsipparameters: appliance attr %q = %q, want %q (unset did not revert it)", attr, got, want)
		}
		return nil
	}
}

func TestAccLbsipparametersDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccLbsipparametersDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_lbsipparameters.tf_lbsipparameters", "addrportvip", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_lbsipparameters.tf_lbsipparameters", "retrydur", "100"),
					resource.TestCheckResourceAttr("data.citrixadc_lbsipparameters.tf_lbsipparameters", "rnatdstport", "80"),
					resource.TestCheckResourceAttr("data.citrixadc_lbsipparameters.tf_lbsipparameters", "rnatsecuredstport", "81"),
					resource.TestCheckResourceAttr("data.citrixadc_lbsipparameters.tf_lbsipparameters", "rnatsecuresrcport", "82"),
					resource.TestCheckResourceAttr("data.citrixadc_lbsipparameters.tf_lbsipparameters", "rnatsrcport", "83"),
					resource.TestCheckResourceAttr("data.citrixadc_lbsipparameters.tf_lbsipparameters", "sip503ratethreshold", "15"),
				),
			},
		},
	})
}
