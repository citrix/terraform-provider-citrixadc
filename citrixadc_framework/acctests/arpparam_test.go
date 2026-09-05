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

const testAccArpparam_add = `
	resource "citrixadc_arpparam" "tf_arpparam" {
		timeout         = 1000
		spoofvalidation = "ENABLED"
	}
`
const testAccArpparam_update = `
	resource "citrixadc_arpparam" "tf_arpparam" {
		timeout         = 1200
		spoofvalidation = "DISABLED"
	}
`

func TestAccArpparam_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccArpparam_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckArpparamExist("citrixadc_arpparam.tf_arpparam", nil),
					resource.TestCheckResourceAttr("citrixadc_arpparam.tf_arpparam", "timeout", "1000"),
					resource.TestCheckResourceAttr("citrixadc_arpparam.tf_arpparam", "spoofvalidation", "ENABLED"),
				),
			},
			{
				Config: testAccArpparam_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckArpparamExist("citrixadc_arpparam.tf_arpparam", nil),
					resource.TestCheckResourceAttr("citrixadc_arpparam.tf_arpparam", "timeout", "1200"),
					resource.TestCheckResourceAttr("citrixadc_arpparam.tf_arpparam", "spoofvalidation", "DISABLED"),
				),
			},
		},
	})
}

func testAccCheckArpparamExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No arpparam name is set")
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
		data, err := client.FindResource(service.Arpparam.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("arpparam %s not found", n)
		}

		return nil
	}
}

func TestAccArpparam_import(t *testing.T) {
	const resAddr = "citrixadc_arpparam.tf_arpparam"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{Config: testAccArpparam_add},
			{
				Config:                  testAccArpparam_add,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

const testAccArpparamDataSource_basic = `

	resource "citrixadc_arpparam" "tf_arpparam" {
		timeout         = 1000
		spoofvalidation = "ENABLED"
	}

	data "citrixadc_arpparam" "tf_arpparam" {
		depends_on = [citrixadc_arpparam.tf_arpparam]
	}
`

func TestAccArpparam_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.0.0"},
				},
				Config: testAccArpparam_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckArpparamExist("citrixadc_arpparam.tf_arpparam", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccArpparam_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckArpparamExist("citrixadc_arpparam.tf_arpparam", nil)),
			},
		},
	})
}

// arpparam is a singleton config resource. Its unset-eligible attributes
// (spoofvalidation, timeout) are set to non-default values in step1 and removed
// in step2; the provider must unset them so the appliance reverts to the
// documented NITRO defaults (DISABLED / 1200).
const testAccArpparam_unset_step1 = `
	resource "citrixadc_arpparam" "tf_unset" {
		spoofvalidation = "ENABLED"
		timeout         = 1000
	}
`

const testAccArpparam_unset_step2 = `
	resource "citrixadc_arpparam" "tf_unset" {
		# All unset-eligible attributes removed from config -> the provider must
		# unset them (revert to NITRO defaults).
	}
`

func TestAccArpparam_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccArpparam_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckArpparamExist("citrixadc_arpparam.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_arpparam.tf_unset", "spoofvalidation", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_arpparam.tf_unset", "timeout", "1000"),
				),
			},
			{
				// Removing the attributes must unset them: state reverts to the
				// documented NITRO defaults and the implicit post-apply plan is empty.
				Config: testAccArpparam_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckArpparamExist("citrixadc_arpparam.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_arpparam.tf_unset", "spoofvalidation", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_arpparam.tf_unset", "timeout", "1200"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckArpparamADCValue("spoofvalidation", "DISABLED"),
					testAccCheckArpparamADCValue("timeout", "1200"),
				),
			},
		},
	})
}

// testAccCheckArpparamADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted it.
func testAccCheckArpparamADCValue(attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Arpparam.Type(), "")
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("arpparam not found on appliance")
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("arpparam: appliance attr %q = %q, want %q (unset did not revert it)", attr, got, want)
		}
		return nil
	}
}

func TestAccArpparamDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccArpparamDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_arpparam.tf_arpparam", "timeout", "1000"),
					resource.TestCheckResourceAttr("data.citrixadc_arpparam.tf_arpparam", "spoofvalidation", "ENABLED"),
				),
			},
		},
	})
}
