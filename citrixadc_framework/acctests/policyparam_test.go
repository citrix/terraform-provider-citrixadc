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

const testAccPolicyparam_basic = `
	resource "citrixadc_policyparam" "tf_policyparam" {
		timeout = 5
	}
`

const testAccPolicyparam_basic_update = `
	resource "citrixadc_policyparam" "tf_policyparam" {
		timeout = 6
	}
`

func TestAccPolicyparam_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// policyparam resource does not have DELETE operation
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				Config: testAccPolicyparam_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPolicyparamExist("citrixadc_policyparam.tf_policyparam", nil),
					resource.TestCheckResourceAttr("citrixadc_policyparam.tf_policyparam", "timeout", "5"),
				),
			},
			{
				Config: testAccPolicyparam_basic_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPolicyparamExist("citrixadc_policyparam.tf_policyparam", nil),
					resource.TestCheckResourceAttr("citrixadc_policyparam.tf_policyparam", "timeout", "6"),
				),
			},
		},
	})
}

func TestAccPolicyparam_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.0.0"},
				},
				Config: testAccPolicyparam_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckPolicyparamExist("citrixadc_policyparam.tf_policyparam", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccPolicyparam_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckPolicyparamExist("citrixadc_policyparam.tf_policyparam", nil)),
			},
		},
	})
}

func testAccCheckPolicyparamExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No policyparam name is set")
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
		data, err := client.FindResource("policyparam", "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("policyparam %s not found", n)
		}

		return nil
	}
}

// policyparam is a singleton; timeout is its only unset-eligible attribute.
// NITRO default is 3900.
const testAccPolicyparam_unset_step1 = `
	resource "citrixadc_policyparam" "tf_unset" {
		timeout = 1000
	}
`

const testAccPolicyparam_unset_step2 = `
	resource "citrixadc_policyparam" "tf_unset" {
		# timeout removed from config -> the provider must unset it
		# (revert to the NITRO default, 3900).
	}
`

func TestAccPolicyparam_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// policyparam resource does not have DELETE operation
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				// Non-default value is applied and persisted.
				Config: testAccPolicyparam_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPolicyparamExist("citrixadc_policyparam.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_policyparam.tf_unset", "timeout", "1000"),
				),
			},
			{
				// Removing the attribute must unset it: state (read back from the
				// appliance) reverts to the documented NITRO default, and the
				// implicit post-apply plan must be empty.
				Config: testAccPolicyparam_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPolicyparamExist("citrixadc_policyparam.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_policyparam.tf_unset", "timeout", "3900"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckPolicyparamADCValue("timeout", "3900"),
				),
			},
		},
	})
}

// testAccCheckPolicyparamADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted
// it.
func testAccCheckPolicyparamADCValue(attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Policyparam.Type(), "")
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("policyparam not found on appliance")
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("policyparam: appliance attr %q = %q, want %q (unset did not revert it)", attr, got, want)
		}
		return nil
	}
}

func TestAccPolicyparam_import(t *testing.T) {
	const resAddr = "citrixadc_policyparam.tf_policyparam"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{Config: testAccPolicyparam_basic},
			{
				Config:                  testAccPolicyparam_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

const testAccPolicyparamDataSource_basic = `
	resource "citrixadc_policyparam" "tf_policyparam_ds" {
		timeout = 5000
	}

	data "citrixadc_policyparam" "tf_policyparam_ds" {
		depends_on = [citrixadc_policyparam.tf_policyparam_ds]
	}
`

func TestAccPolicyparamDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccPolicyparamDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.citrixadc_policyparam.tf_policyparam_ds", "id"),
					resource.TestCheckResourceAttr("data.citrixadc_policyparam.tf_policyparam_ds", "timeout", "5000"),
				),
			},
		},
	})
}
