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

const testAccLldpparam_basic = `
	resource "citrixadc_lldpparam" "tf_lldpparam" {
		holdtimetxmult = 3
		mode           = "TRANSMITTER"
		timer          = 40
	}
`
const testAccLldpparam_update = `
	resource "citrixadc_lldpparam" "tf_lldpparam" {
		holdtimetxmult = 10
		mode           = "RECEIVER"
		timer          = 60
	}
`

func TestAccLldpparam_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccLldpparam_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLldpparamExist("citrixadc_lldpparam.tf_lldpparam", nil),
					resource.TestCheckResourceAttr("citrixadc_lldpparam.tf_lldpparam", "holdtimetxmult", "3"),
					resource.TestCheckResourceAttr("citrixadc_lldpparam.tf_lldpparam", "mode", "TRANSMITTER"),
					resource.TestCheckResourceAttr("citrixadc_lldpparam.tf_lldpparam", "timer", "40"),
				),
			},
			{
				Config: testAccLldpparam_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLldpparamExist("citrixadc_lldpparam.tf_lldpparam", nil),
					resource.TestCheckResourceAttr("citrixadc_lldpparam.tf_lldpparam", "holdtimetxmult", "10"),
					resource.TestCheckResourceAttr("citrixadc_lldpparam.tf_lldpparam", "mode", "RECEIVER"),
					resource.TestCheckResourceAttr("citrixadc_lldpparam.tf_lldpparam", "timer", "60"),
				),
			},
		},
	})
}

func testAccCheckLldpparamExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lldpparam name is set")
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
		data, err := client.FindResource("lldpparam", "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("lldpparam %s not found", n)
		}

		return nil
	}
}

func TestAccLldpparam_import(t *testing.T) {
	const resAddr = "citrixadc_lldpparam.tf_lldpparam"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{Config: testAccLldpparam_basic},
			{
				Config:                  testAccLldpparam_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

const testAccLldpparamDataSource_basic = `
	resource "citrixadc_lldpparam" "tf_lldpparam" {
		holdtimetxmult = 3
		mode           = "TRANSMITTER"
		timer          = 40
	}

	data "citrixadc_lldpparam" "tf_lldpparam" {
		depends_on = [citrixadc_lldpparam.tf_lldpparam]
	}
`

func TestAccLldpparam_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.0.0"},
				},
				Config: testAccLldpparam_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckLldpparamExist("citrixadc_lldpparam.tf_lldpparam", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccLldpparam_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckLldpparamExist("citrixadc_lldpparam.tf_lldpparam", nil)),
			},
		},
	})
}

const testAccLldpparam_unset_step1 = `
	resource "citrixadc_lldpparam" "tf_unset" {
		holdtimetxmult = 10
		timer          = 60
	}
`

const testAccLldpparam_unset_step2 = `
	resource "citrixadc_lldpparam" "tf_unset" {
		# Unset-eligible attributes removed from config -> provider must unset
		# them (revert to NITRO defaults: holdtimetxmult=4, timer=30).
	}
`

func TestAccLldpparam_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccLldpparam_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLldpparamExist("citrixadc_lldpparam.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_lldpparam.tf_unset", "holdtimetxmult", "10"),
					resource.TestCheckResourceAttr("citrixadc_lldpparam.tf_unset", "timer", "60"),
				),
			},
			{
				// Removing the attributes must unset them: state reverts to the
				// documented NITRO defaults and the implicit post-apply plan is empty.
				Config: testAccLldpparam_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLldpparamExist("citrixadc_lldpparam.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_lldpparam.tf_unset", "holdtimetxmult", "4"),
					resource.TestCheckResourceAttr("citrixadc_lldpparam.tf_unset", "timer", "30"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckLldpparamADCValue("holdtimetxmult", "4"),
					testAccCheckLldpparamADCValue("timer", "30"),
				),
			},
		},
	})
}

// testAccCheckLldpparamADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted it.
func testAccCheckLldpparamADCValue(attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Lldpparam.Type(), "")
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("lldpparam not found on appliance")
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("lldpparam: appliance attr %q = %q, want %q (unset did not revert it)", attr, got, want)
		}
		return nil
	}
}

func TestAccLldpparamDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccLldpparamDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_lldpparam.tf_lldpparam", "holdtimetxmult", "3"),
					resource.TestCheckResourceAttr("data.citrixadc_lldpparam.tf_lldpparam", "mode", "TRANSMITTER"),
					resource.TestCheckResourceAttr("data.citrixadc_lldpparam.tf_lldpparam", "timer", "40"),
				),
			},
		},
	})
}
