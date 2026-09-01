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

const testAccCsparameter_basic = `
	resource "citrixadc_csparameter" "tf_csparameter" {
		stateupdate = "ENABLED"
	}
`

const testAccCsparameter_basic_update = `
	resource "citrixadc_csparameter" "tf_csparameter" {
		stateupdate = "DISABLED"
	}
`

func TestAccCsparameter_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCsparameterDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCsparameter_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCsparameterExist("citrixadc_csparameter.tf_csparameter", nil),
					resource.TestCheckResourceAttr("citrixadc_csparameter.tf_csparameter", "stateupdate", "ENABLED"),
				),
			},
			{
				Config: testAccCsparameter_basic_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCsparameterExist("citrixadc_csparameter.tf_csparameter", nil),
					resource.TestCheckResourceAttr("citrixadc_csparameter.tf_csparameter", "stateupdate", "DISABLED"),
				),
			},
		},
	})
}

func TestAccCsparameter_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckCsparameterDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.0.0"},
				},
				Config: testAccCsparameter_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckCsparameterExist("citrixadc_csparameter.tf_csparameter", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccCsparameter_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckCsparameterExist("citrixadc_csparameter.tf_csparameter", nil)),
			},
		},
	})
}

func testAccCheckCsparameterExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No csparameter name is set")
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
		data, err := client.FindResource(service.Csparameter.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("csparameter %s not found", n)
		}

		return nil
	}
}

func testAccCheckCsparameterDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_csparameter" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Csparameter.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("csparameter %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

// csparameter is a singleton with a single mutable attribute, stateupdate
// (NITRO default DISABLED). Step 1 sets it to a non-default value; step 2
// removes it from config, which must unset it back to the NITRO default.
const testAccCsparameter_unset_step1 = `
	resource "citrixadc_csparameter" "tf_unset" {
		stateupdate = "ENABLED"
	}
`

const testAccCsparameter_unset_step2 = `
	resource "citrixadc_csparameter" "tf_unset" {
		# stateupdate removed from config -> provider must unset it (revert to
		# NITRO default, "DISABLED").
	}
`

func TestAccCsparameter_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCsparameterDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default value is applied and persisted.
				Config: testAccCsparameter_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCsparameterExist("citrixadc_csparameter.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_csparameter.tf_unset", "stateupdate", "ENABLED"),
				),
			},
			{
				// Removing the attribute must unset it: state (read back from the
				// appliance) reverts to the documented NITRO default, and the
				// implicit post-apply plan must be empty.
				Config: testAccCsparameter_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCsparameterExist("citrixadc_csparameter.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_csparameter.tf_unset", "stateupdate", "DISABLED"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckCsparameterADCValue("stateupdate", "DISABLED"),
				),
			},
		},
	})
}

// testAccCheckCsparameterADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted
// it.
func testAccCheckCsparameterADCValue(attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Csparameter.Type(), "")
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("csparameter not found on appliance")
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("csparameter: appliance attr %q = %q, want %q (unset did not revert it)", attr, got, want)
		}
		return nil
	}
}

func TestAccCsparameter_import(t *testing.T) {
	const resAddr = "citrixadc_csparameter.tf_csparameter"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCsparameterDestroy,
		Steps: []resource.TestStep{
			{Config: testAccCsparameter_basic},
			{
				Config:                  testAccCsparameter_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

const testAccCsparameterDataSource_basic = `

	resource "citrixadc_csparameter" "tf_csparameter" {
		stateupdate = "ENABLED"
	}

	data "citrixadc_csparameter" "tf_csparameter" {
		depends_on = [citrixadc_csparameter.tf_csparameter]
	}
`

func TestAccCsparameterDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccCsparameterDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_csparameter.tf_csparameter", "stateupdate", "ENABLED"),
					// Universal runtime-binding proof.
					resource.TestCheckResourceAttrSet("data.citrixadc_csparameter.tf_csparameter", "id"),
				),
			},
		},
	})
}
