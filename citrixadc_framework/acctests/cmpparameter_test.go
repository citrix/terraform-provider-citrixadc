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

const testAccCmpparameter_basic = `


resource "citrixadc_cmpparameter" "tf_cmpparameter" {
	cmplevel    = "optimal"
	quantumsize = 20
	servercmp   = "OFF"
	randomgzipfilenameminlength = "12"
	randomgzipfilenamemaxlength = "20"
	randomgzipfilename = "ENABLED"
	}
`
const testAccCmpparameter_update = `


	resource "citrixadc_cmpparameter" "tf_cmpparameter" {
		cmplevel    = "bestspeed"
		quantumsize = 30
		servercmp   = "ON"
		randomgzipfilenameminlength = "14"
		randomgzipfilenamemaxlength = "22"
		randomgzipfilename = "DISABLED"
	}
`

func TestAccCmpparameter_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccCmpparameter_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCmpparameterExist("citrixadc_cmpparameter.tf_cmpparameter", nil),
					resource.TestCheckResourceAttr("citrixadc_cmpparameter.tf_cmpparameter", "cmplevel", "optimal"),
					resource.TestCheckResourceAttr("citrixadc_cmpparameter.tf_cmpparameter", "quantumsize", "20"),
					resource.TestCheckResourceAttr("citrixadc_cmpparameter.tf_cmpparameter", "servercmp", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_cmpparameter.tf_cmpparameter", "randomgzipfilenameminlength", "12"),
					resource.TestCheckResourceAttr("citrixadc_cmpparameter.tf_cmpparameter", "randomgzipfilenamemaxlength", "20"),
					resource.TestCheckResourceAttr("citrixadc_cmpparameter.tf_cmpparameter", "randomgzipfilename", "ENABLED"),
				),
			},
			{
				Config: testAccCmpparameter_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCmpparameterExist("citrixadc_cmpparameter.tf_cmpparameter", nil),
					resource.TestCheckResourceAttr("citrixadc_cmpparameter.tf_cmpparameter", "cmplevel", "bestspeed"),
					resource.TestCheckResourceAttr("citrixadc_cmpparameter.tf_cmpparameter", "quantumsize", "30"),
					resource.TestCheckResourceAttr("citrixadc_cmpparameter.tf_cmpparameter", "servercmp", "ON"),
					resource.TestCheckResourceAttr("citrixadc_cmpparameter.tf_cmpparameter", "randomgzipfilenameminlength", "14"),
					resource.TestCheckResourceAttr("citrixadc_cmpparameter.tf_cmpparameter", "randomgzipfilenamemaxlength", "22"),
					resource.TestCheckResourceAttr("citrixadc_cmpparameter.tf_cmpparameter", "randomgzipfilename", "DISABLED"),
				),
			},
		},
	})
}

func TestAccCmpparameter_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccCmpparameter_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckCmpparameterExist("citrixadc_cmpparameter.tf_cmpparameter", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccCmpparameter_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckCmpparameterExist("citrixadc_cmpparameter.tf_cmpparameter", nil)),
			},
		},
	})
}

func testAccCheckCmpparameterExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No cmpparameter name is set")
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
		data, err := client.FindResource(service.Cmpparameter.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("cmpparameter %s not found", n)
		}

		return nil
	}
}

func TestAccCmpparameter_import(t *testing.T) {
	const resAddr = "citrixadc_cmpparameter.tf_cmpparameter"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{Config: testAccCmpparameter_basic},
			{
				Config:                  testAccCmpparameter_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

const testAccCmpparameterDataSource_basic = `


	resource "citrixadc_cmpparameter" "tf_cmpparameter" {
		cmplevel    = "optimal"
		quantumsize = 20
		servercmp   = "OFF"
		randomgzipfilenameminlength = "12"
		randomgzipfilenamemaxlength = "20"
		randomgzipfilename = "ENABLED"
	}

	data "citrixadc_cmpparameter" "tf_cmpparameter" {
		depends_on = [citrixadc_cmpparameter.tf_cmpparameter]
	}
`

// cmpparameter is a singleton config resource. Step 1 sets every unset-eligible
// attribute to a valid non-default value; step 2 removes them all so the provider
// must unset them (revert to the documented NITRO defaults).
const testAccCmpparameter_unset_step1 = `
resource "citrixadc_cmpparameter" "tf_unset" {
	addvaryheader               = "ENABLED"
	cmpbypasspct                = 50
	cmplevel                    = "bestspeed"
	cmponpush                   = "ENABLED"
	externalcache               = "YES"
	heurexpiry                  = "ON"
	heurexpiryhistwt            = 25
	heurexpirythres             = 200
	quantumsize                 = 20000
	randomgzipfilename          = "ENABLED"
	randomgzipfilenameminlength = 12
	randomgzipfilenamemaxlength = 20
	servercmp                   = "OFF"
}
`

const testAccCmpparameter_unset_step2 = `
resource "citrixadc_cmpparameter" "tf_unset" {
	# All unset-eligible attributes removed from config -> the provider must
	# unset them (revert to the documented NITRO defaults).
}
`

func TestAccCmpparameter_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccCmpparameter_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCmpparameterExist("citrixadc_cmpparameter.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_cmpparameter.tf_unset", "addvaryheader", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_cmpparameter.tf_unset", "cmpbypasspct", "50"),
					resource.TestCheckResourceAttr("citrixadc_cmpparameter.tf_unset", "cmplevel", "bestspeed"),
					resource.TestCheckResourceAttr("citrixadc_cmpparameter.tf_unset", "cmponpush", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_cmpparameter.tf_unset", "externalcache", "YES"),
					resource.TestCheckResourceAttr("citrixadc_cmpparameter.tf_unset", "heurexpiry", "ON"),
					resource.TestCheckResourceAttr("citrixadc_cmpparameter.tf_unset", "heurexpiryhistwt", "25"),
					resource.TestCheckResourceAttr("citrixadc_cmpparameter.tf_unset", "heurexpirythres", "200"),
					resource.TestCheckResourceAttr("citrixadc_cmpparameter.tf_unset", "quantumsize", "20000"),
					resource.TestCheckResourceAttr("citrixadc_cmpparameter.tf_unset", "randomgzipfilename", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_cmpparameter.tf_unset", "randomgzipfilenameminlength", "12"),
					resource.TestCheckResourceAttr("citrixadc_cmpparameter.tf_unset", "randomgzipfilenamemaxlength", "20"),
					resource.TestCheckResourceAttr("citrixadc_cmpparameter.tf_unset", "servercmp", "OFF"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from
				// the appliance) reverts to the documented NITRO defaults, and the
				// implicit post-apply plan must be empty.
				Config: testAccCmpparameter_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCmpparameterExist("citrixadc_cmpparameter.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_cmpparameter.tf_unset", "addvaryheader", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_cmpparameter.tf_unset", "cmpbypasspct", "100"),
					resource.TestCheckResourceAttr("citrixadc_cmpparameter.tf_unset", "cmplevel", "optimal"),
					resource.TestCheckResourceAttr("citrixadc_cmpparameter.tf_unset", "cmponpush", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_cmpparameter.tf_unset", "externalcache", "NO"),
					resource.TestCheckNoResourceAttr("citrixadc_cmpparameter.tf_unset", "heurexpiry"),
					resource.TestCheckNoResourceAttr("citrixadc_cmpparameter.tf_unset", "heurexpiryhistwt"),
					resource.TestCheckNoResourceAttr("citrixadc_cmpparameter.tf_unset", "heurexpirythres"),
					resource.TestCheckResourceAttr("citrixadc_cmpparameter.tf_unset", "quantumsize", "57344"),
					resource.TestCheckResourceAttr("citrixadc_cmpparameter.tf_unset", "randomgzipfilename", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_cmpparameter.tf_unset", "randomgzipfilenameminlength", "8"),
					resource.TestCheckResourceAttr("citrixadc_cmpparameter.tf_unset", "randomgzipfilenamemaxlength", "63"),
					resource.TestCheckResourceAttr("citrixadc_cmpparameter.tf_unset", "servercmp", "ON"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckCmpparameterADCValue("cmplevel", "optimal"),
					testAccCheckCmpparameterADCValue("servercmp", "ON"),
					testAccCheckCmpparameterADCValue("cmpbypasspct", "100"),
				),
			},
		},
	})
}

// testAccCheckCmpparameterADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted it.
func testAccCheckCmpparameterADCValue(attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Cmpparameter.Type(), "")
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("cmpparameter not found on appliance")
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("cmpparameter: appliance attr %q = %q, want %q (unset did not revert it)", attr, got, want)
		}
		return nil
	}
}

func TestAccCmpparameterDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccCmpparameterDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_cmpparameter.tf_cmpparameter", "cmplevel", "optimal"),
					resource.TestCheckResourceAttr("data.citrixadc_cmpparameter.tf_cmpparameter", "quantumsize", "20"),
					resource.TestCheckResourceAttr("data.citrixadc_cmpparameter.tf_cmpparameter", "servercmp", "OFF"),
					resource.TestCheckResourceAttr("data.citrixadc_cmpparameter.tf_cmpparameter", "randomgzipfilenameminlength", "12"),
					resource.TestCheckResourceAttr("data.citrixadc_cmpparameter.tf_cmpparameter", "randomgzipfilenamemaxlength", "20"),
					resource.TestCheckResourceAttr("data.citrixadc_cmpparameter.tf_cmpparameter", "randomgzipfilename", "ENABLED"),
				),
			},
		},
	})
}
