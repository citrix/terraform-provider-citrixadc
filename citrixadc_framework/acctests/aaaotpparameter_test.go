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

const testAccAaaotpparameter_basic = `
	resource "citrixadc_aaaotpparameter" "tf_aaaotpparameter" {
		encryption = "OFF"
		maxotpdevices = 3
	}
`
const testAccAaaotpparameter_update = `
	resource "citrixadc_aaaotpparameter" "tf_aaaotpparameter" {
		encryption = "ON"
		maxotpdevices = 5
	}
`

func TestAccAaaotpparameter_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccAaaotpparameter_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaaotpparameterExist("citrixadc_aaaotpparameter.tf_aaaotpparameter", nil),
					resource.TestCheckResourceAttr("citrixadc_aaaotpparameter.tf_aaaotpparameter", "encryption", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_aaaotpparameter.tf_aaaotpparameter", "maxotpdevices", "3"),
				),
			},
			{
				Config: testAccAaaotpparameter_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaaotpparameterExist("citrixadc_aaaotpparameter.tf_aaaotpparameter", nil),
					resource.TestCheckResourceAttr("citrixadc_aaaotpparameter.tf_aaaotpparameter", "encryption", "ON"),
					resource.TestCheckResourceAttr("citrixadc_aaaotpparameter.tf_aaaotpparameter", "maxotpdevices", "5"),
				),
			},
		},
	})
}

func testAccCheckAaaotpparameterExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No aaaotpparameter name is set")
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
		data, err := client.FindResource("aaaotpparameter", "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("aaaotpparameter %s not found", n)
		}

		return nil
	}
}

func TestAccAaaotpparameter_import(t *testing.T) {
	const resAddr = "citrixadc_aaaotpparameter.tf_aaaotpparameter"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{Config: testAccAaaotpparameter_basic},
			{
				Config:                  testAccAaaotpparameter_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccAaaotpparameter_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.0.0"},
				},
				Config: testAccAaaotpparameter_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAaaotpparameterExist("citrixadc_aaaotpparameter.tf_aaaotpparameter", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccAaaotpparameter_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAaaotpparameterExist("citrixadc_aaaotpparameter.tf_aaaotpparameter", nil)),
			},
		},
	})
}

const testAccAaaotpparameter_unset_step1 = `
	resource "citrixadc_aaaotpparameter" "tf_unset" {
		encryption    = "ON"
		maxotpdevices = 7
	}
`

const testAccAaaotpparameter_unset_step2 = `
	resource "citrixadc_aaaotpparameter" "tf_unset" {
		# All unset-eligible attributes removed from config -> the provider must
		# unset them (revert to NITRO defaults: encryption=OFF, maxotpdevices=4).
	}
`

func TestAccAaaotpparameter_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccAaaotpparameter_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaaotpparameterExist("citrixadc_aaaotpparameter.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_aaaotpparameter.tf_unset", "encryption", "ON"),
					resource.TestCheckResourceAttr("citrixadc_aaaotpparameter.tf_unset", "maxotpdevices", "7"),
				),
			},
			{
				// Removing the attributes must unset them: state reverts to the
				// documented NITRO defaults and the implicit post-apply plan is empty.
				Config: testAccAaaotpparameter_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaaotpparameterExist("citrixadc_aaaotpparameter.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_aaaotpparameter.tf_unset", "encryption", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_aaaotpparameter.tf_unset", "maxotpdevices", "4"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckAaaotpparameterADCValue("encryption", "OFF"),
					testAccCheckAaaotpparameterADCValue("maxotpdevices", "4"),
				),
			},
		},
	})
}

// testAccCheckAaaotpparameterADCValue asserts an attribute's value directly on
// the appliance (not just in Terraform state), proving the unset reverted it.
func testAccCheckAaaotpparameterADCValue(attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Aaaotpparameter.Type(), "")
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("aaaotpparameter not found on appliance")
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("aaaotpparameter: appliance attr %q = %q, want %q (unset did not revert it)", attr, got, want)
		}
		return nil
	}
}

const testAccAaaotpparameterDataSource_basic = `
	resource "citrixadc_aaaotpparameter" "tf_aaaotpparameter" {
		encryption = "OFF"
		maxotpdevices = 3
	}
	
	data "citrixadc_aaaotpparameter" "tf_aaaotpparameter" {
		depends_on = [citrixadc_aaaotpparameter.tf_aaaotpparameter]
	}
`

func TestAccAaaotpparameterDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccAaaotpparameterDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_aaaotpparameter.tf_aaaotpparameter", "encryption", "OFF"),
					resource.TestCheckResourceAttr("data.citrixadc_aaaotpparameter.tf_aaaotpparameter", "maxotpdevices", "3"),
					// Universal runtime-binding proof that the data source resolved.
					resource.TestCheckResourceAttrSet("data.citrixadc_aaaotpparameter.tf_aaaotpparameter", "id"),
				),
			},
		},
	})
}
