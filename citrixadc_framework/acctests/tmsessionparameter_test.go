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

const testAccTmsessionparameter_basic = `


resource "citrixadc_tmsessionparameter" "tf_tmsessionparameter" {
	sesstimeout                = 40
	defaultauthorizationaction = "ALLOW"
	sso                        = "ON"
	ssodomain                  = 3
	}
  
`
const testAccTmsessionparameter_update = `


resource "citrixadc_tmsessionparameter" "tf_tmsessionparameter" {
	sesstimeout                = 50
	defaultauthorizationaction = "DENY"
	sso                        = "OFF"
	}
  
`

func TestAccTmsessionparameter_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccTmsessionparameter_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTmsessionparameterExist("citrixadc_tmsessionparameter.tf_tmsessionparameter", nil),
					resource.TestCheckResourceAttr("citrixadc_tmsessionparameter.tf_tmsessionparameter", "sesstimeout", "40"),
					resource.TestCheckResourceAttr("citrixadc_tmsessionparameter.tf_tmsessionparameter", "defaultauthorizationaction", "ALLOW"),
					resource.TestCheckResourceAttr("citrixadc_tmsessionparameter.tf_tmsessionparameter", "sso", "ON"),
					resource.TestCheckResourceAttr("citrixadc_tmsessionparameter.tf_tmsessionparameter", "ssodomain", "3"),
				),
			},
			{
				Config: testAccTmsessionparameter_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTmsessionparameterExist("citrixadc_tmsessionparameter.tf_tmsessionparameter", nil),
					resource.TestCheckResourceAttr("citrixadc_tmsessionparameter.tf_tmsessionparameter", "sesstimeout", "50"),
					resource.TestCheckResourceAttr("citrixadc_tmsessionparameter.tf_tmsessionparameter", "defaultauthorizationaction", "DENY"),
					resource.TestCheckResourceAttr("citrixadc_tmsessionparameter.tf_tmsessionparameter", "sso", "OFF"),
				),
			},
		},
	})
}

func testAccCheckTmsessionparameterExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No tmsessionparameter name is set")
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
		data, err := client.FindResource(service.Tmsessionparameter.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("tmsessionparameter %s not found", n)
		}

		return nil
	}
}

func TestAccTmsessionparameter_import(t *testing.T) {
	const resAddr = "citrixadc_tmsessionparameter.tf_tmsessionparameter"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{Config: testAccTmsessionparameter_basic},
			{
				Config:                  testAccTmsessionparameter_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

const testAccTmsessionparameterDataSource_basic = `


resource "citrixadc_tmsessionparameter" "tf_tmsessionparameter" {
	sesstimeout                = 40
	defaultauthorizationaction = "ALLOW"
	sso                        = "ON"
	ssodomain                  = 3
	}

data "citrixadc_tmsessionparameter" "tf_tmsessionparameter" {
    depends_on = [citrixadc_tmsessionparameter.tf_tmsessionparameter]
}
`

func TestAccTmsessionparameter_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.0.0"},
				},
				Config: testAccTmsessionparameter_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTmsessionparameterExist("citrixadc_tmsessionparameter.tf_tmsessionparameter", nil),
				),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccTmsessionparameter_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTmsessionparameterExist("citrixadc_tmsessionparameter.tf_tmsessionparameter", nil),
				),
			},
		},
	})
}

// tmsessionparameter is a singleton config resource. Step1 sets the
// unset-eligible attributes (each with a documented NITRO default) to
// non-default values; step2 removes them so the provider must unset them,
// reverting the appliance to the documented defaults.
const testAccTmsessionparameter_unset_step1 = `
resource "citrixadc_tmsessionparameter" "tf_unset" {
	defaultauthorizationaction = "ALLOW"
	homepage                   = "http://example.com"
	httponlycookie             = "NO"
	sesstimeout                = 40
	sso                        = "ON"
	ssocredential              = "SECONDARY"
}
`

const testAccTmsessionparameter_unset_step2 = `
resource "citrixadc_tmsessionparameter" "tf_unset" {
	# All unset-eligible attributes removed from config -> the provider must
	# unset them (revert to the documented NITRO defaults).
}
`

func TestAccTmsessionparameter_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccTmsessionparameter_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTmsessionparameterExist("citrixadc_tmsessionparameter.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_tmsessionparameter.tf_unset", "defaultauthorizationaction", "ALLOW"),
					resource.TestCheckResourceAttr("citrixadc_tmsessionparameter.tf_unset", "homepage", "http://example.com"),
					resource.TestCheckResourceAttr("citrixadc_tmsessionparameter.tf_unset", "httponlycookie", "NO"),
					resource.TestCheckResourceAttr("citrixadc_tmsessionparameter.tf_unset", "sesstimeout", "40"),
					resource.TestCheckResourceAttr("citrixadc_tmsessionparameter.tf_unset", "sso", "ON"),
					resource.TestCheckResourceAttr("citrixadc_tmsessionparameter.tf_unset", "ssocredential", "SECONDARY"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from
				// the appliance) reverts to the documented NITRO defaults, and the
				// implicit post-apply plan must be empty.
				Config: testAccTmsessionparameter_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTmsessionparameterExist("citrixadc_tmsessionparameter.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_tmsessionparameter.tf_unset", "defaultauthorizationaction", "DENY"),
					resource.TestCheckResourceAttr("citrixadc_tmsessionparameter.tf_unset", "homepage", "None"),
					resource.TestCheckResourceAttr("citrixadc_tmsessionparameter.tf_unset", "httponlycookie", "YES"),
					resource.TestCheckResourceAttr("citrixadc_tmsessionparameter.tf_unset", "sesstimeout", "30"),
					resource.TestCheckResourceAttr("citrixadc_tmsessionparameter.tf_unset", "sso", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_tmsessionparameter.tf_unset", "ssocredential", "PRIMARY"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckTmsessionparameterADCValue("defaultauthorizationaction", "DENY"),
					testAccCheckTmsessionparameterADCValue("sso", "OFF"),
					testAccCheckTmsessionparameterADCValue("httponlycookie", "YES"),
				),
			},
		},
	})
}

// testAccCheckTmsessionparameterADCValue asserts an attribute's value directly
// on the appliance (not just in Terraform state), proving the unset actually
// reverted it.
func testAccCheckTmsessionparameterADCValue(attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Tmsessionparameter.Type(), "")
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("tmsessionparameter not found on appliance")
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("tmsessionparameter: appliance attr %q = %q, want %q (unset did not revert it)", attr, got, want)
		}
		return nil
	}
}

func TestAccTmsessionparameterDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccTmsessionparameterDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					// "id" is the universal runtime-binding proof (singleton static ID).
					resource.TestCheckResourceAttrSet("data.citrixadc_tmsessionparameter.tf_tmsessionparameter", "id"),
					resource.TestCheckResourceAttr("data.citrixadc_tmsessionparameter.tf_tmsessionparameter", "sesstimeout", "40"),
					resource.TestCheckResourceAttr("data.citrixadc_tmsessionparameter.tf_tmsessionparameter", "defaultauthorizationaction", "ALLOW"),
					resource.TestCheckResourceAttr("data.citrixadc_tmsessionparameter.tf_tmsessionparameter", "sso", "ON"),
					resource.TestCheckResourceAttr("data.citrixadc_tmsessionparameter.tf_tmsessionparameter", "ssodomain", "3"),
				),
			},
		},
	})
}
