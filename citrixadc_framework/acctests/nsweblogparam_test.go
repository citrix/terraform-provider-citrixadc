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

const testAccNsweblogparam_add = `
	resource "citrixadc_nsweblogparam" "tf_nsweblofparam" {
		buffersizemb  = 32
		customreqhdrs = ["req1", "req2"]
		customrsphdrs = ["res1", "res2"]
	}
`
const testAccNsweblogparam_update = `
	resource "citrixadc_nsweblogparam" "tf_nsweblofparam" {
		buffersizemb  = 16
		customreqhdrs = ["req1", "req2"]
		customrsphdrs = ["res1", "res2"]
	}
`

func TestAccNsweblogparam_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccNsweblogparam_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsweblogparamExist("citrixadc_nsweblogparam.tf_nsweblofparam", nil),
					resource.TestCheckResourceAttr("citrixadc_nsweblogparam.tf_nsweblofparam", "buffersizemb", "32"),
				),
			},
			{
				Config: testAccNsweblogparam_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsweblogparamExist("citrixadc_nsweblogparam.tf_nsweblofparam", nil),
					resource.TestCheckResourceAttr("citrixadc_nsweblogparam.tf_nsweblofparam", "buffersizemb", "16"),
				),
			},
		},
	})
}

func testAccCheckNsweblogparamExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No nsweblogparam name is set")
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
		data, err := client.FindResource(service.Nsweblogparam.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("nsweblogparam %s not found", n)
		}

		return nil
	}
}
func TestAccNsweblogparamDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNsweblogparamDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.citrixadc_nsweblogparam.test", "id"),
					resource.TestCheckResourceAttrSet("data.citrixadc_nsweblogparam.test", "buffersizemb"),
				),
			},
		},
	})
}

func TestAccNsweblogparam_import(t *testing.T) {
	const resAddr = "citrixadc_nsweblogparam.tf_nsweblofparam"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{Config: testAccNsweblogparam_add},
			{
				Config:                  testAccNsweblogparam_add,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

const testAccNsweblogparamDataSource_basic = `
data "citrixadc_nsweblogparam" "test" {
}
`

// buffersizemb is the only spec-unsettable attribute (documented NITRO default
// value 16); the customreqhdrs/customrsphdrs lists have no documented server
// default and are therefore not unset-eligible.
const testAccNsweblogparam_unset_step1 = `
	resource "citrixadc_nsweblogparam" "tf_unset" {
		buffersizemb = 32
	}
`

const testAccNsweblogparam_unset_step2 = `
	resource "citrixadc_nsweblogparam" "tf_unset" {
		# buffersizemb removed from config -> the provider must unset it
		# (revert to the NITRO default, 16).
	}
`

func TestAccNsweblogparam_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				// Non-default value is applied and persisted.
				Config: testAccNsweblogparam_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsweblogparamExist("citrixadc_nsweblogparam.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_nsweblogparam.tf_unset", "buffersizemb", "32"),
				),
			},
			{
				// Removing the attribute must unset it: state (read back from the
				// appliance) reverts to the NITRO default and the implicit
				// post-apply plan must be empty.
				Config: testAccNsweblogparam_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsweblogparamExist("citrixadc_nsweblogparam.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_nsweblogparam.tf_unset", "buffersizemb", "16"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckNsweblogparamADCValue("buffersizemb", "16"),
				),
			},
		},
	})
}

// testAccCheckNsweblogparamADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted it.
func testAccCheckNsweblogparamADCValue(attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Nsweblogparam.Type(), "")
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("nsweblogparam not found on appliance")
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("nsweblogparam: appliance attr %q = %q, want %q (unset did not revert it)", attr, got, want)
		}
		return nil
	}
}

func TestAccNsweblogparam_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccNsweblogparam_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckNsweblogparamExist("citrixadc_nsweblogparam.tf_nsweblofparam", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccNsweblogparam_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckNsweblogparamExist("citrixadc_nsweblogparam.tf_nsweblofparam", nil)),
			},
		},
	})
}
