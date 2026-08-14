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

const testAccGslbparameter_basic = `

resource "citrixadc_gslbparameter" "tf_gslbparameter" {
	ldnsentrytimeout = 50
	rtttolerance     = 6
	ldnsmask         = "255.255.255.255"
	gslbsyncsaveconfigcommand = "DISABLED"
	}
`

const testAccGslbparameter_update = `

resource "citrixadc_gslbparameter" "tf_gslbparameter" {
	ldnsentrytimeout = 70
	rtttolerance     = 8
	ldnsmask         = "255.255.255.254"
	gslbsyncsaveconfigcommand = "ENABLED"
	}
`

func TestAccGslbparameter_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// gslb resource do not have DELETE operation
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				Config: testAccGslbparameter_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGslbparameterExist("citrixadc_gslbparameter.tf_gslbparameter", nil),
					resource.TestCheckResourceAttr("citrixadc_gslbparameter.tf_gslbparameter", "ldnsentrytimeout", "50"),
					resource.TestCheckResourceAttr("citrixadc_gslbparameter.tf_gslbparameter", "rtttolerance", "6"),
					resource.TestCheckResourceAttr("citrixadc_gslbparameter.tf_gslbparameter", "ldnsmask", "255.255.255.255"),
					resource.TestCheckResourceAttr("citrixadc_gslbparameter.tf_gslbparameter", "gslbsyncsaveconfigcommand", "DISABLED"),
				),
			},
			{
				Config: testAccGslbparameter_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGslbparameterExist("citrixadc_gslbparameter.tf_gslbparameter", nil),
					resource.TestCheckResourceAttr("citrixadc_gslbparameter.tf_gslbparameter", "ldnsentrytimeout", "70"),
					resource.TestCheckResourceAttr("citrixadc_gslbparameter.tf_gslbparameter", "rtttolerance", "8"),
					resource.TestCheckResourceAttr("citrixadc_gslbparameter.tf_gslbparameter", "ldnsmask", "255.255.255.254"),
					resource.TestCheckResourceAttr("citrixadc_gslbparameter.tf_gslbparameter", "gslbsyncsaveconfigcommand", "ENABLED"),
				),
			},
		},
	})
}

func testAccCheckGslbparameterExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No gslbparameter name is set")
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
		data, err := client.FindResource(service.Gslbparameter.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("gslbparameter %s not found", n)
		}

		return nil
	}
}

func TestAccGslbparameter_import(t *testing.T) {
	const resAddr = "citrixadc_gslbparameter.tf_gslbparameter"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{Config: testAccGslbparameter_basic},
			{
				Config:                  testAccGslbparameter_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccGslbparameter_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccGslbparameter_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckGslbparameterExist("citrixadc_gslbparameter.tf_gslbparameter", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccGslbparameter_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckGslbparameterExist("citrixadc_gslbparameter.tf_gslbparameter", nil)),
			},
		},
	})
}

// The gslbparameter unset test covers the spec-unsettable, documented-default
// mutable attributes. step1 sets them to valid non-default values; step2 removes
// them from config so the provider must unset them (revert to NITRO defaults).
const testAccGslbparameter_unset_step1 = `
resource "citrixadc_gslbparameter" "tf_unset" {
	automaticconfigsync       = "ENABLED"
	dropldnsreq               = "ENABLED"
	gslbconfigsyncmonitor     = "ENABLED"
	gslbsyncinterval          = 20
	gslbsynclocfiles          = "DISABLED"
	gslbsyncmode              = "FullSync"
	gslbsyncsaveconfigcommand = "ENABLED"
	ldnsentrytimeout          = 70
	mepkeepalivetimeout       = 20
	rtttolerance              = 8
	undefaction               = "DROP"
	v6ldnsmasklen             = 64
}
`

const testAccGslbparameter_unset_step2 = `
resource "citrixadc_gslbparameter" "tf_unset" {
	# All unset-eligible attributes removed from config -> the provider must
	# unset them (revert to NITRO defaults).
}
`

func TestAccGslbparameter_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccGslbparameter_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGslbparameterExist("citrixadc_gslbparameter.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_gslbparameter.tf_unset", "automaticconfigsync", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_gslbparameter.tf_unset", "dropldnsreq", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_gslbparameter.tf_unset", "gslbconfigsyncmonitor", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_gslbparameter.tf_unset", "gslbsyncinterval", "20"),
					resource.TestCheckResourceAttr("citrixadc_gslbparameter.tf_unset", "gslbsynclocfiles", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_gslbparameter.tf_unset", "gslbsyncmode", "FullSync"),
					resource.TestCheckResourceAttr("citrixadc_gslbparameter.tf_unset", "gslbsyncsaveconfigcommand", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_gslbparameter.tf_unset", "ldnsentrytimeout", "70"),
					resource.TestCheckResourceAttr("citrixadc_gslbparameter.tf_unset", "mepkeepalivetimeout", "20"),
					resource.TestCheckResourceAttr("citrixadc_gslbparameter.tf_unset", "rtttolerance", "8"),
					resource.TestCheckResourceAttr("citrixadc_gslbparameter.tf_unset", "undefaction", "DROP"),
					resource.TestCheckResourceAttr("citrixadc_gslbparameter.tf_unset", "v6ldnsmasklen", "64"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from the
				// appliance) reverts to the documented NITRO defaults and the implicit
				// post-apply plan must be empty.
				Config: testAccGslbparameter_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGslbparameterExist("citrixadc_gslbparameter.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_gslbparameter.tf_unset", "automaticconfigsync", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_gslbparameter.tf_unset", "dropldnsreq", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_gslbparameter.tf_unset", "gslbconfigsyncmonitor", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_gslbparameter.tf_unset", "gslbsyncinterval", "10"),
					resource.TestCheckResourceAttr("citrixadc_gslbparameter.tf_unset", "gslbsynclocfiles", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_gslbparameter.tf_unset", "gslbsyncmode", "IncrementalSync"),
					resource.TestCheckResourceAttr("citrixadc_gslbparameter.tf_unset", "gslbsyncsaveconfigcommand", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_gslbparameter.tf_unset", "ldnsentrytimeout", "180"),
					resource.TestCheckResourceAttr("citrixadc_gslbparameter.tf_unset", "mepkeepalivetimeout", "10"),
					resource.TestCheckResourceAttr("citrixadc_gslbparameter.tf_unset", "rtttolerance", "5"),
					resource.TestCheckResourceAttr("citrixadc_gslbparameter.tf_unset", "undefaction", "NOLBACTION"),
					resource.TestCheckResourceAttr("citrixadc_gslbparameter.tf_unset", "v6ldnsmasklen", "128"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckGslbparameterADCValue("gslbsyncmode", "IncrementalSync"),
					testAccCheckGslbparameterADCValue("ldnsentrytimeout", "180"),
					testAccCheckGslbparameterADCValue("undefaction", "NOLBACTION"),
				),
			},
		},
	})
}

// testAccCheckGslbparameterADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted it.
func testAccCheckGslbparameterADCValue(attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Gslbparameter.Type(), "")
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("gslbparameter not found on appliance")
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("gslbparameter: appliance attr %q = %q, want %q (unset did not revert it)", attr, got, want)
		}
		return nil
	}
}

const testAccGslbparameterDataSource_basic = `

	resource "citrixadc_gslbparameter" "tf_gslbparameter" {
		ldnsentrytimeout = 50
		rtttolerance     = 6
		ldnsmask         = "255.255.255.255"
		gslbsyncsaveconfigcommand = "DISABLED"
	}

	data "citrixadc_gslbparameter" "tf_gslbparameter" {
		depends_on = [citrixadc_gslbparameter.tf_gslbparameter]
	}
`

func TestAccGslbparameterDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccGslbparameterDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_gslbparameter.tf_gslbparameter", "ldnsentrytimeout", "50"),
					resource.TestCheckResourceAttr("data.citrixadc_gslbparameter.tf_gslbparameter", "rtttolerance", "6"),
					resource.TestCheckResourceAttr("data.citrixadc_gslbparameter.tf_gslbparameter", "ldnsmask", "255.255.255.255"),
					resource.TestCheckResourceAttr("data.citrixadc_gslbparameter.tf_gslbparameter", "gslbsyncsaveconfigcommand", "DISABLED"),
				),
			},
		},
	})
}
