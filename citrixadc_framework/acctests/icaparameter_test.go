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

const testAccIcaparameter_basic = `


resource "citrixadc_icaparameter" "tf_icaparameter" {
	edtpmtuddf           = "ENABLED"
	edtpmtuddftimeout    = 200
	l7latencyfrequency   = 0
	enablesronhafailover = "YES"
	edtpmtudrediscovery = "DISABLED"
	edtlosstolerant = "DISABLED"
	dfpersistence = "DISABLED"
	hdxinsightnonnsap = "NO"
	}
  
`
const testAccIcaparameter_update = `

resource "citrixadc_icaparameter" "tf_icaparameter" {
	edtpmtuddf           = "ENABLED"
	edtpmtuddftimeout    = 100
	l7latencyfrequency   = 30
	enablesronhafailover = "NO"
	edtpmtudrediscovery = "ENABLED"
	edtlosstolerant = "ENABLED"
	dfpersistence = "ENABLED"
	hdxinsightnonnsap = "YES"
	}
  
`

func TestAccIcaparameter_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccIcaparameter_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIcaparameterExist("citrixadc_icaparameter.tf_icaparameter", nil),
					resource.TestCheckResourceAttr("citrixadc_icaparameter.tf_icaparameter", "edtpmtuddf", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_icaparameter.tf_icaparameter", "edtpmtuddftimeout", "200"),
					resource.TestCheckResourceAttr("citrixadc_icaparameter.tf_icaparameter", "l7latencyfrequency", "0"),
					resource.TestCheckResourceAttr("citrixadc_icaparameter.tf_icaparameter", "enablesronhafailover", "YES"),
					resource.TestCheckResourceAttr("citrixadc_icaparameter.tf_icaparameter", "edtpmtudrediscovery", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_icaparameter.tf_icaparameter", "edtlosstolerant", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_icaparameter.tf_icaparameter", "dfpersistence", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_icaparameter.tf_icaparameter", "hdxinsightnonnsap", "NO"),
				),
			},
			{
				Config: testAccIcaparameter_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIcaparameterExist("citrixadc_icaparameter.tf_icaparameter", nil),
					resource.TestCheckResourceAttr("citrixadc_icaparameter.tf_icaparameter", "edtpmtuddf", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_icaparameter.tf_icaparameter", "edtpmtuddftimeout", "100"),
					resource.TestCheckResourceAttr("citrixadc_icaparameter.tf_icaparameter", "l7latencyfrequency", "30"),
					resource.TestCheckResourceAttr("citrixadc_icaparameter.tf_icaparameter", "enablesronhafailover", "NO"),
					resource.TestCheckResourceAttr("citrixadc_icaparameter.tf_icaparameter", "edtpmtudrediscovery", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_icaparameter.tf_icaparameter", "edtlosstolerant", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_icaparameter.tf_icaparameter", "dfpersistence", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_icaparameter.tf_icaparameter", "hdxinsightnonnsap", "YES"),
				),
			},
		},
	})
}

func TestAccIcaparameter_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccIcaparameter_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckIcaparameterExist("citrixadc_icaparameter.tf_icaparameter", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccIcaparameter_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckIcaparameterExist("citrixadc_icaparameter.tf_icaparameter", nil)),
			},
		},
	})
}

func testAccCheckIcaparameterExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No icaparameter name is set")
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
		data, err := client.FindResource("icaparameter", "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("icaparameter %s not found", n)
		}

		return nil
	}
}

// icaparameter is a singleton. step1 sets every unset-eligible attribute to a
// non-default value; step2 removes them all so the provider must unset them,
// reverting each to its documented NITRO default.
const testAccIcaparameter_unset_step1 = `
resource "citrixadc_icaparameter" "tf_unset" {
	dfpersistence        = "ENABLED"
	edtpmtuddf           = "DISABLED"
	edtpmtuddftimeout    = 200
	edtpmtudrediscovery  = "ENABLED"
	enablesronhafailover = "YES"
	hdxinsightnonnsap    = "NO"
	l7latencyfrequency   = 30
}
`

const testAccIcaparameter_unset_step2 = `
resource "citrixadc_icaparameter" "tf_unset" {
	# All unset-eligible attributes removed from config -> the provider must
	# unset them (revert to NITRO defaults).
}
`

func TestAccIcaparameter_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccIcaparameter_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIcaparameterExist("citrixadc_icaparameter.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_icaparameter.tf_unset", "dfpersistence", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_icaparameter.tf_unset", "edtpmtuddf", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_icaparameter.tf_unset", "edtpmtuddftimeout", "200"),
					resource.TestCheckResourceAttr("citrixadc_icaparameter.tf_unset", "edtpmtudrediscovery", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_icaparameter.tf_unset", "enablesronhafailover", "YES"),
					resource.TestCheckResourceAttr("citrixadc_icaparameter.tf_unset", "hdxinsightnonnsap", "NO"),
					resource.TestCheckResourceAttr("citrixadc_icaparameter.tf_unset", "l7latencyfrequency", "30"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from the
				// appliance) reverts to the documented NITRO defaults, and the
				// implicit post-apply plan must be empty.
				Config: testAccIcaparameter_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIcaparameterExist("citrixadc_icaparameter.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_icaparameter.tf_unset", "dfpersistence", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_icaparameter.tf_unset", "edtpmtuddf", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_icaparameter.tf_unset", "edtpmtuddftimeout", "100"),
					resource.TestCheckResourceAttr("citrixadc_icaparameter.tf_unset", "edtpmtudrediscovery", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_icaparameter.tf_unset", "enablesronhafailover", "NO"),
					resource.TestCheckResourceAttr("citrixadc_icaparameter.tf_unset", "hdxinsightnonnsap", "YES"),
					resource.TestCheckResourceAttr("citrixadc_icaparameter.tf_unset", "l7latencyfrequency", "0"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckIcaparameterADCValue("dfpersistence", "DISABLED"),
					testAccCheckIcaparameterADCValue("edtpmtuddf", "ENABLED"),
					testAccCheckIcaparameterADCValue("enablesronhafailover", "NO"),
				),
			},
		},
	})
}

// testAccCheckIcaparameterADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted it.
func testAccCheckIcaparameterADCValue(attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Icaparameter.Type(), "")
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("icaparameter not found on appliance")
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("icaparameter: appliance attr %q = %q, want %q (unset did not revert it)", attr, got, want)
		}
		return nil
	}
}

func TestAccIcaparameter_import(t *testing.T) {
	const resAddr = "citrixadc_icaparameter.tf_icaparameter"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{Config: testAccIcaparameter_basic},
			{
				Config:                  testAccIcaparameter_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

const testAccIcaparameterDataSource_basic = `

resource "citrixadc_icaparameter" "tf_icaparameter" {
	edtpmtuddf           = "ENABLED"
	edtpmtuddftimeout    = 200
	l7latencyfrequency   = 0
	enablesronhafailover = "YES"
	edtpmtudrediscovery = "DISABLED"
	edtlosstolerant = "DISABLED"
	dfpersistence = "DISABLED"
	hdxinsightnonnsap = "NO"
}

data "citrixadc_icaparameter" "tf_icaparameter_ds" {
	depends_on = [citrixadc_icaparameter.tf_icaparameter]
}
`

func TestAccIcaparameterDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIcaparameterDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_icaparameter.tf_icaparameter_ds", "edtpmtuddf", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_icaparameter.tf_icaparameter_ds", "edtpmtuddftimeout", "200"),
					resource.TestCheckResourceAttr("data.citrixadc_icaparameter.tf_icaparameter_ds", "l7latencyfrequency", "0"),
					resource.TestCheckResourceAttr("data.citrixadc_icaparameter.tf_icaparameter_ds", "enablesronhafailover", "YES"),
					resource.TestCheckResourceAttr("data.citrixadc_icaparameter.tf_icaparameter_ds", "edtpmtudrediscovery", "DISABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_icaparameter.tf_icaparameter_ds", "edtlosstolerant", "DISABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_icaparameter.tf_icaparameter_ds", "dfpersistence", "DISABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_icaparameter.tf_icaparameter_ds", "hdxinsightnonnsap", "NO"),
				),
			},
		},
	})
}
