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

const testAccResponderparam_basic = `

	resource "citrixadc_responderparam" "tf_responderparam" {
		timeout = 5
		undefaction = "RESET"
	}
`

const testAccResponderparam_basic_update = `

	resource "citrixadc_responderparam" "tf_responderparam" {
		timeout = 6
		undefaction = "DROP"
	}
`

func TestAccResponderparam_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckResponderparamDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccResponderparam_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckResponderparamExist("citrixadc_responderparam.tf_responderparam", nil),
					resource.TestCheckResourceAttr("citrixadc_responderparam.tf_responderparam", "timeout", "5"),
					resource.TestCheckResourceAttr("citrixadc_responderparam.tf_responderparam", "undefaction", "RESET"),
				),
			},
			{
				Config: testAccResponderparam_basic_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckResponderparamExist("citrixadc_responderparam.tf_responderparam", nil),
					resource.TestCheckResourceAttr("citrixadc_responderparam.tf_responderparam", "timeout", "6"),
					resource.TestCheckResourceAttr("citrixadc_responderparam.tf_responderparam", "undefaction", "DROP"),
				),
			},
		},
	})
}

func testAccCheckResponderparamExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No responderparam name is set")
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
		data, err := client.FindResource(service.Responderparam.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("responderparam %s not found", n)
		}

		return nil
	}
}

func testAccCheckResponderparamDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_responderparam" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Responderparam.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("responderparam %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccResponderparam_import(t *testing.T) {
	const resAddr = "citrixadc_responderparam.tf_responderparam"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckResponderparamDestroy,
		Steps: []resource.TestStep{
			{Config: testAccResponderparam_basic},
			{
				Config:                  testAccResponderparam_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccResponderparam_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckResponderparamDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccResponderparam_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckResponderparamExist("citrixadc_responderparam.tf_responderparam", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccResponderparam_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckResponderparamExist("citrixadc_responderparam.tf_responderparam", nil)),
			},
		},
	})
}

const testAccResponderparamDataSource_basic = `
resource "citrixadc_responderparam" "tf_responderparam_ds" {
    timeout = 7
    undefaction = "DROP"
}

data "citrixadc_responderparam" "tf_responderparam_ds" {
  depends_on = [citrixadc_responderparam.tf_responderparam_ds]
}
`

// step1 sets the unset-eligible attributes to valid NON-default values; step2
// removes them so the provider must unset them (revert to NITRO defaults:
// timeout=3900, undefaction=NOOP).
const testAccResponderparam_unset_step1 = `
	resource "citrixadc_responderparam" "tf_unset" {
		timeout     = 1200
		undefaction = "RESET"
	}
`

const testAccResponderparam_unset_step2 = `
	resource "citrixadc_responderparam" "tf_unset" {
		# All unset-eligible attributes removed from config -> the provider must
		# unset them (revert to NITRO defaults).
	}
`

func TestAccResponderparam_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckResponderparamDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccResponderparam_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckResponderparamExist("citrixadc_responderparam.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_responderparam.tf_unset", "timeout", "1200"),
					resource.TestCheckResourceAttr("citrixadc_responderparam.tf_unset", "undefaction", "RESET"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from
				// the appliance) reverts to the documented NITRO defaults, and the
				// implicit post-apply plan must be empty.
				Config: testAccResponderparam_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckResponderparamExist("citrixadc_responderparam.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_responderparam.tf_unset", "timeout", "3900"),
					resource.TestCheckResourceAttr("citrixadc_responderparam.tf_unset", "undefaction", "NOOP"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckResponderparamADCValue("timeout", "3900"),
					testAccCheckResponderparamADCValue("undefaction", "NOOP"),
				),
			},
		},
	})
}

// testAccCheckResponderparamADCValue asserts an attribute's value directly on
// the appliance (not just in Terraform state), proving the unset actually
// reverted it.
func testAccCheckResponderparamADCValue(attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Responderparam.Type(), "")
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("responderparam not found on appliance")
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("responderparam: appliance attr %q = %q, want %q (unset did not revert it)", attr, got, want)
		}
		return nil
	}
}

func TestAccResponderparamDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResponderparamDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_responderparam.tf_responderparam_ds", "timeout", "7"),
					resource.TestCheckResourceAttr("data.citrixadc_responderparam.tf_responderparam_ds", "undefaction", "DROP"),
				),
			},
		},
	})
}
