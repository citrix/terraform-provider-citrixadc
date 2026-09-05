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

const testAccAaaparameter_basic = `

	resource "citrixadc_aaaparameter" "tf_aaaparameter" {
		enablestaticpagecaching    = "NO"
		enableenhancedauthfeedback = "YES"
		defaultauthtype            = "LOCAL"
		// maxaaausers is intentionally omitted: this appliance's license caps it at
		// "Unlimited" (NITRO reports 4294967295 = 2^32-1). Any value below the
		// license limit is silently ignored by the ADC ("MaxAAAUsers value less
		// than allowed by license, ignored"), so setting it produces a perpetual diff.
		maxloginattempts           = 5
		failedlogintimeout         = 15
	}

`
const testAccAaaparameter_update = `

	resource "citrixadc_aaaparameter" "tf_aaaparameter" {
		enablestaticpagecaching    = "YES"
		enableenhancedauthfeedback = "NO"
		defaultauthtype            = "LOCAL"
		// maxaaausers intentionally omitted (license-capped, see basic step above).
		maxloginattempts           = 10
		failedlogintimeout         = 20
	}
  
`

func TestAccAaaparameter_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccAaaparameter_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaaparameterExist("citrixadc_aaaparameter.tf_aaaparameter", nil),
					resource.TestCheckResourceAttr("citrixadc_aaaparameter.tf_aaaparameter", "enablestaticpagecaching", "NO"),
					resource.TestCheckResourceAttr("citrixadc_aaaparameter.tf_aaaparameter", "enableenhancedauthfeedback", "YES"),
					resource.TestCheckResourceAttr("citrixadc_aaaparameter.tf_aaaparameter", "defaultauthtype", "LOCAL"),
					// maxaaausers is license-capped ("Unlimited") on this appliance; not asserted.
					resource.TestCheckResourceAttr("citrixadc_aaaparameter.tf_aaaparameter", "maxloginattempts", "5"),
					resource.TestCheckResourceAttr("citrixadc_aaaparameter.tf_aaaparameter", "failedlogintimeout", "15"),
				),
			},
			{
				Config: testAccAaaparameter_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaaparameterExist("citrixadc_aaaparameter.tf_aaaparameter", nil),
					resource.TestCheckResourceAttr("citrixadc_aaaparameter.tf_aaaparameter", "enablestaticpagecaching", "YES"),
					resource.TestCheckResourceAttr("citrixadc_aaaparameter.tf_aaaparameter", "enableenhancedauthfeedback", "NO"),
					resource.TestCheckResourceAttr("citrixadc_aaaparameter.tf_aaaparameter", "defaultauthtype", "LOCAL"),
					// maxaaausers is license-capped ("Unlimited") on this appliance; not asserted.
					resource.TestCheckResourceAttr("citrixadc_aaaparameter.tf_aaaparameter", "maxloginattempts", "10"),
					resource.TestCheckResourceAttr("citrixadc_aaaparameter.tf_aaaparameter", "failedlogintimeout", "20"),
				),
			},
		},
	})
}

func testAccCheckAaaparameterExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No aaaparameter name is set")
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
		data, err := client.FindResource(service.Aaaparameter.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("aaaparameter %s not found", n)
		}

		return nil
	}
}

const testAccAaaparameterDataSource_basic = `

	resource "citrixadc_aaaparameter" "tf_aaaparameter" {
		enablestaticpagecaching    = "NO"
		enableenhancedauthfeedback = "YES"
		defaultauthtype            = "LOCAL"
		maxloginattempts           = 5
		failedlogintimeout         = 15
	}
	
	data "citrixadc_aaaparameter" "tf_aaaparameter" {
		depends_on = [citrixadc_aaaparameter.tf_aaaparameter]
	}
`

func TestAccAaaparameterDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccAaaparameterDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_aaaparameter.tf_aaaparameter", "enablestaticpagecaching", "NO"),
					resource.TestCheckResourceAttr("data.citrixadc_aaaparameter.tf_aaaparameter", "enableenhancedauthfeedback", "YES"),
					resource.TestCheckResourceAttr("data.citrixadc_aaaparameter.tf_aaaparameter", "defaultauthtype", "LOCAL"),
					resource.TestCheckResourceAttr("data.citrixadc_aaaparameter.tf_aaaparameter", "maxloginattempts", "5"),
					resource.TestCheckResourceAttr("data.citrixadc_aaaparameter.tf_aaaparameter", "failedlogintimeout", "15"),
					// Universal runtime-binding proof for the data source.
					resource.TestCheckResourceAttrSet("data.citrixadc_aaaparameter.tf_aaaparameter", "id"),
				),
			},
		},
	})
}

func TestAccAaaparameter_import(t *testing.T) {
	const resAddr = "citrixadc_aaaparameter.tf_aaaparameter"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{Config: testAccAaaparameter_basic},
			{
				Config:                  testAccAaaparameter_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

// aaaparameter is a singleton. Step 1 sets a set of unset-eligible toggle
// attributes to valid non-default values; step 2 removes them from config so the
// provider must unset them (revert to the documented NITRO defaults).
const testAccAaaparameter_unset_step1 = `

	resource "citrixadc_aaaparameter" "tf_aaaparameter" {
		aaadloglevel            = "DEBUG"
		apitokencache           = "ENABLED"
		defaultcspheader        = "DISABLED"
		enablesessionstickiness = "YES"
		enhancedepa             = "ENABLED"
		httponlycookie          = "DISABLED"
		loginencryption         = "ENABLED"
		maxkbquestions          = 4
		persistentloginattempts = "ENABLED"
		securityinsights        = "ENABLED"
	}
`

const testAccAaaparameter_unset_step2 = `

	resource "citrixadc_aaaparameter" "tf_aaaparameter" {
		# All unset-eligible attributes removed from config -> the provider must
		# unset them (revert to NITRO defaults).
	}
`

func TestAccAaaparameter_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccAaaparameter_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaaparameterExist("citrixadc_aaaparameter.tf_aaaparameter", nil),
					resource.TestCheckResourceAttr("citrixadc_aaaparameter.tf_aaaparameter", "aaadloglevel", "DEBUG"),
					resource.TestCheckResourceAttr("citrixadc_aaaparameter.tf_aaaparameter", "apitokencache", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_aaaparameter.tf_aaaparameter", "defaultcspheader", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_aaaparameter.tf_aaaparameter", "enablesessionstickiness", "YES"),
					resource.TestCheckResourceAttr("citrixadc_aaaparameter.tf_aaaparameter", "enhancedepa", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_aaaparameter.tf_aaaparameter", "httponlycookie", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_aaaparameter.tf_aaaparameter", "loginencryption", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_aaaparameter.tf_aaaparameter", "maxkbquestions", "4"),
					resource.TestCheckResourceAttr("citrixadc_aaaparameter.tf_aaaparameter", "persistentloginattempts", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_aaaparameter.tf_aaaparameter", "securityinsights", "ENABLED"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from
				// the appliance) reverts to the documented NITRO defaults, and the
				// implicit post-apply plan must be empty.
				Config: testAccAaaparameter_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaaparameterExist("citrixadc_aaaparameter.tf_aaaparameter", nil),
					resource.TestCheckResourceAttr("citrixadc_aaaparameter.tf_aaaparameter", "aaadloglevel", "INFORMATIONAL"),
					resource.TestCheckResourceAttr("citrixadc_aaaparameter.tf_aaaparameter", "apitokencache", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_aaaparameter.tf_aaaparameter", "defaultcspheader", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_aaaparameter.tf_aaaparameter", "enablesessionstickiness", "NO"),
					resource.TestCheckResourceAttr("citrixadc_aaaparameter.tf_aaaparameter", "enhancedepa", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_aaaparameter.tf_aaaparameter", "httponlycookie", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_aaaparameter.tf_aaaparameter", "loginencryption", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_aaaparameter.tf_aaaparameter", "maxkbquestions", "2"),
					resource.TestCheckResourceAttr("citrixadc_aaaparameter.tf_aaaparameter", "persistentloginattempts", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_aaaparameter.tf_aaaparameter", "securityinsights", "DISABLED"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckAaaparameterADCValue("defaultcspheader", "ENABLED"),
					testAccCheckAaaparameterADCValue("enablesessionstickiness", "NO"),
					testAccCheckAaaparameterADCValue("httponlycookie", "ENABLED"),
					testAccCheckAaaparameterADCValue("persistentloginattempts", "DISABLED"),
				),
			},
		},
	})
}

// testAccCheckAaaparameterADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted it.
func testAccCheckAaaparameterADCValue(attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Aaaparameter.Type(), "")
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("aaaparameter not found on appliance")
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("aaaparameter: appliance attr %q = %q, want %q (unset did not revert it)", attr, got, want)
		}
		return nil
	}
}

func TestAccAaaparameter_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.0.0"},
				},
				Config: testAccAaaparameter_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAaaparameterExist("citrixadc_aaaparameter.tf_aaaparameter", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccAaaparameter_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAaaparameterExist("citrixadc_aaaparameter.tf_aaaparameter", nil)),
			},
		},
	})
}
