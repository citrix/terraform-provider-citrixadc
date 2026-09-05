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

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const testAccAdmparameter_basic = `
	resource "citrixadc_admparameter" "tf_admparameter" {
		admserviceconnect = "DISABLED"
	}
`
const testAccAdmparameter_update = `
	resource "citrixadc_admparameter" "tf_admparameter" {
		admserviceconnect = "ENABLED"
	}
`

func TestAccAdmparameter_basic(t *testing.T) {
	t.Skip("Autoconnect cannot be disabled for Citrix Internal NetScalers")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccAdmparameter_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmparameterExist("citrixadc_admparameter.tf_admparameter", nil),
					resource.TestCheckResourceAttr("citrixadc_admparameter.tf_admparameter", "admserviceconnect", "DISABLED"),
				),
			},
			{
				Config: testAccAdmparameter_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmparameterExist("citrixadc_admparameter.tf_admparameter", nil),
					resource.TestCheckResourceAttr("citrixadc_admparameter.tf_admparameter", "admserviceconnect", "ENABLED"),
				),
			},
		},
	})
}

func TestAccAdmparameter_import(t *testing.T) {
	t.Skip("Autoconnect cannot be disabled for Citrix Internal NetScalers")
	const resAddr = "citrixadc_admparameter.tf_admparameter"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{Config: testAccAdmparameter_basic},
			{
				Config:                  testAccAdmparameter_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccAdmparameter_sdkv2StateUpgrade(t *testing.T) {
	t.Skip("Autoconnect cannot be disabled for Citrix Internal NetScalers")
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.0.0"},
				},
				Config: testAccAdmparameter_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAdmparameterExist("citrixadc_admparameter.tf_admparameter", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccAdmparameter_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAdmparameterExist("citrixadc_admparameter.tf_admparameter", nil)),
			},
		},
	})
}

const testAccAdmparameter_unset_step1 = `
	resource "citrixadc_admparameter" "tf_unset" {
		admserviceconnect = "DISABLED"
	}
`

const testAccAdmparameter_unset_step2 = `
	resource "citrixadc_admparameter" "tf_unset" {
		# admserviceconnect removed from config -> the provider must unset it
		# (revert to NITRO default, "ENABLED").
	}
`

func TestAccAdmparameter_unset(t *testing.T) {
	t.Skip("Autoconnect cannot be disabled for Citrix Internal NetScalers")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				// Non-default value is applied and persisted.
				Config: testAccAdmparameter_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmparameterExist("citrixadc_admparameter.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_admparameter.tf_unset", "admserviceconnect", "DISABLED"),
				),
			},
			{
				// Removing the attribute must unset it: state (read back from the
				// appliance) reverts to the documented NITRO default, and the
				// implicit post-apply plan must be empty.
				Config: testAccAdmparameter_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAdmparameterExist("citrixadc_admparameter.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_admparameter.tf_unset", "admserviceconnect", "ENABLED"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckAdmparameterADCValue("admserviceconnect", "ENABLED"),
				),
			},
		},
	})
}

// testAccCheckAdmparameterADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted
// it. admparameter is an unnamed singleton, so it is read with an empty name.
func testAccCheckAdmparameterADCValue(attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource("admparameter", "")
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("admparameter not found on appliance")
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("admparameter: appliance attr %q = %q, want %q (unset did not revert it)", attr, got, want)
		}
		return nil
	}
}

func testAccCheckAdmparameterExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No admparameter name is set")
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
		data, err := client.FindResource("admparameter", "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("admparameter %s not found", n)
		}

		return nil
	}
}
