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
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const testAccContentinspectionparameter_basic = `

resource "citrixadc_contentinspectionparameter" "tf_contentinspectionparameter" {
	undefaction = "RESET"
	}
`
const testAccContentinspectionparameter_update = `

resource "citrixadc_contentinspectionparameter" "tf_contentinspectionparameter" {
	undefaction = "DROP"
	}
`

func TestAccContentinspectionparameter_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccContentinspectionparameter_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckContentinspectionparameterExist("citrixadc_contentinspectionparameter.tf_contentinspectionparameter", nil),
					resource.TestCheckResourceAttr("citrixadc_contentinspectionparameter.tf_contentinspectionparameter", "undefaction", "RESET"),
				),
			},
			{
				Config: testAccContentinspectionparameter_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckContentinspectionparameterExist("citrixadc_contentinspectionparameter.tf_contentinspectionparameter", nil),
					resource.TestCheckResourceAttr("citrixadc_contentinspectionparameter.tf_contentinspectionparameter", "undefaction", "DROP"),
				),
			},
		},
	})
}

func testAccCheckContentinspectionparameterExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No contentinspectionparameter name is set")
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
		data, err := client.FindResource("contentinspectionparameter", "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("contentinspectionparameter %s not found", n)
		}

		return nil
	}
}

func TestAccContentinspectionparameter_import(t *testing.T) {
	const resAddr = "citrixadc_contentinspectionparameter.tf_contentinspectionparameter"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{Config: testAccContentinspectionparameter_basic},
			{
				Config:                  testAccContentinspectionparameter_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

const testAccContentinspectionparameterDataSource_basic = `

resource "citrixadc_contentinspectionparameter" "tf_contentinspectionparameter" {
	undefaction = "RESET"
}

data "citrixadc_contentinspectionparameter" "tf_contentinspectionparameter_datasource" {
	depends_on = [citrixadc_contentinspectionparameter.tf_contentinspectionparameter]
}
`

func TestAccContentinspectionparameter_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccContentinspectionparameter_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckContentinspectionparameterExist("citrixadc_contentinspectionparameter.tf_contentinspectionparameter", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccContentinspectionparameter_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckContentinspectionparameterExist("citrixadc_contentinspectionparameter.tf_contentinspectionparameter", nil)),
			},
		},
	})
}

const testAccContentinspectionparameter_unset_step1 = `

resource "citrixadc_contentinspectionparameter" "tf_contentinspectionparameter" {
	undefaction = "DROP"
	}
`

const testAccContentinspectionparameter_unset_step2 = `

resource "citrixadc_contentinspectionparameter" "tf_contentinspectionparameter" {
	# undefaction removed from config -> the provider must unset it
	# (revert to NITRO default, "NOINSPECTION").
	}
`

func TestAccContentinspectionparameter_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				// Non-default value is applied and persisted.
				Config: testAccContentinspectionparameter_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckContentinspectionparameterExist("citrixadc_contentinspectionparameter.tf_contentinspectionparameter", nil),
					resource.TestCheckResourceAttr("citrixadc_contentinspectionparameter.tf_contentinspectionparameter", "undefaction", "DROP"),
				),
			},
			{
				// Removing the attribute must unset it: state (read back from the
				// appliance) reverts to the documented NITRO default, and the
				// implicit post-apply plan must be empty.
				Config: testAccContentinspectionparameter_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckContentinspectionparameterExist("citrixadc_contentinspectionparameter.tf_contentinspectionparameter", nil),
					resource.TestCheckResourceAttr("citrixadc_contentinspectionparameter.tf_contentinspectionparameter", "undefaction", "NOINSPECTION"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckContentinspectionparameterADCValue("undefaction", "NOINSPECTION"),
				),
			},
		},
	})
}

// testAccCheckContentinspectionparameterADCValue asserts an attribute's value
// directly on the appliance (not just in Terraform state), proving the unset
// actually reverted it.
func testAccCheckContentinspectionparameterADCValue(attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource("contentinspectionparameter", "")
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("contentinspectionparameter not found on appliance")
		}
		got := fmt.Sprintf("%v", data[attr])
		if got != want {
			return fmt.Errorf("contentinspectionparameter: appliance attr %q = %q, want %q (unset did not revert it)", attr, got, want)
		}
		return nil
	}
}

func TestAccContentinspectionparameterDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccContentinspectionparameterDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_contentinspectionparameter.tf_contentinspectionparameter_datasource", "undefaction", "RESET"),
				),
			},
		},
	})
}
