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

const testAccAaapreauthenticationparameter_basic = `

	resource "citrixadc_aaapreauthenticationparameter" "tf_aaapreauthenticationparameter" {
		preauthenticationaction = "ALLOW"
		deletefiles    = "/var/tmp/*.files"
	}
`
const testAccAaapreauthenticationparameter_update = `

	resource "citrixadc_aaapreauthenticationparameter" "tf_aaapreauthenticationparameter" {
		preauthenticationaction = "DENY"
		deletefiles    = "/var/tmp/*.files"
	}
`

func TestAccAaapreauthenticationparameter_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccAaapreauthenticationparameter_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaapreauthenticationparameterExist("citrixadc_aaapreauthenticationparameter.tf_aaapreauthenticationparameter", nil),
					resource.TestCheckResourceAttr("citrixadc_aaapreauthenticationparameter.tf_aaapreauthenticationparameter", "preauthenticationaction", "ALLOW"),
					resource.TestCheckResourceAttr("citrixadc_aaapreauthenticationparameter.tf_aaapreauthenticationparameter", "deletefiles", "/var/tmp/*.files"),
				),
			},
			{
				Config: testAccAaapreauthenticationparameter_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaapreauthenticationparameterExist("citrixadc_aaapreauthenticationparameter.tf_aaapreauthenticationparameter", nil),
					resource.TestCheckResourceAttr("citrixadc_aaapreauthenticationparameter.tf_aaapreauthenticationparameter", "preauthenticationaction", "DENY"),
					resource.TestCheckResourceAttr("citrixadc_aaapreauthenticationparameter.tf_aaapreauthenticationparameter", "deletefiles", "/var/tmp/*.files"),
				),
			},
		},
	})
}

func TestAccAaapreauthenticationparameter_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccAaapreauthenticationparameter_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAaapreauthenticationparameterExist("citrixadc_aaapreauthenticationparameter.tf_aaapreauthenticationparameter", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccAaapreauthenticationparameter_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAaapreauthenticationparameterExist("citrixadc_aaapreauthenticationparameter.tf_aaapreauthenticationparameter", nil)),
			},
		},
	})
}

func testAccCheckAaapreauthenticationparameterExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No aaapreauthenticationparameter name is set")
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
		data, err := client.FindResource(service.Aaapreauthenticationparameter.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("aaapreauthenticationparameter %s not found", n)
		}

		return nil
	}
}

// The aaapreauthenticationparameter unset test covers the two attributes that
// have a documented NITRO server default (always echoed back by GET) and unset
// cleanly: preauthenticationaction (-> "ALLOW") and rule (-> "ns_true").
// killprocess and deletefiles are NOT covered: NITRO omits them from GET after
// unset (no server default), so they cannot round-trip a schema Default.
const testAccAaapreauthenticationparameter_unset_step1 = `
	resource "citrixadc_aaapreauthenticationparameter" "tf_unset" {
		preauthenticationaction = "DENY"
		rule                    = "ns_false"
	}
`

const testAccAaapreauthenticationparameter_unset_step2 = `
	resource "citrixadc_aaapreauthenticationparameter" "tf_unset" {
		# All unset-eligible attributes removed from config -> the provider must
		# unset them (revert to NITRO defaults).
	}
`

func TestAccAaapreauthenticationparameter_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccAaapreauthenticationparameter_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaapreauthenticationparameterExist("citrixadc_aaapreauthenticationparameter.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_aaapreauthenticationparameter.tf_unset", "preauthenticationaction", "DENY"),
					resource.TestCheckResourceAttr("citrixadc_aaapreauthenticationparameter.tf_unset", "rule", "ns_false"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from the
				// appliance) reverts to the documented NITRO defaults, and the
				// implicit post-apply plan must be empty.
				Config: testAccAaapreauthenticationparameter_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaapreauthenticationparameterExist("citrixadc_aaapreauthenticationparameter.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_aaapreauthenticationparameter.tf_unset", "preauthenticationaction", "ALLOW"),
					resource.TestCheckResourceAttr("citrixadc_aaapreauthenticationparameter.tf_unset", "rule", "ns_true"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckAaapreauthenticationparameterADCValue("preauthenticationaction", "ALLOW"),
					testAccCheckAaapreauthenticationparameterADCValue("rule", "ns_true"),
				),
			},
		},
	})
}

// testAccCheckAaapreauthenticationparameterADCValue asserts an attribute's value
// directly on the appliance (not just in Terraform state), proving the unset
// actually reverted it.
func testAccCheckAaapreauthenticationparameterADCValue(attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Aaapreauthenticationparameter.Type(), "")
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("aaapreauthenticationparameter not found on appliance")
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("aaapreauthenticationparameter: appliance attr %q = %q, want %q (unset did not revert it)", attr, got, want)
		}
		return nil
	}
}

func TestAccAaapreauthenticationparameter_import(t *testing.T) {
	const resAddr = "citrixadc_aaapreauthenticationparameter.tf_aaapreauthenticationparameter"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{Config: testAccAaapreauthenticationparameter_basic},
			{
				Config:                  testAccAaapreauthenticationparameter_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

const testAccAaapreauthenticationparameterDataSource_basic = `

	resource "citrixadc_aaapreauthenticationparameter" "tf_aaapreauthenticationparameter" {
		preauthenticationaction = "ALLOW"
		deletefiles    = "/var/tmp/*.files"
	}
	
	data "citrixadc_aaapreauthenticationparameter" "tf_aaapreauthenticationparameter" {
		depends_on = [citrixadc_aaapreauthenticationparameter.tf_aaapreauthenticationparameter]
	}
`

func TestAccAaapreauthenticationparameterDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccAaapreauthenticationparameterDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_aaapreauthenticationparameter.tf_aaapreauthenticationparameter", "preauthenticationaction", "ALLOW"),
					resource.TestCheckResourceAttr("data.citrixadc_aaapreauthenticationparameter.tf_aaapreauthenticationparameter", "deletefiles", "/var/tmp/*.files"),
				),
			},
		},
	})
}
