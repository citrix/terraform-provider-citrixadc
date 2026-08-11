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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testAccSnmpalarm_basic = `
resource "citrixadc_snmpalarm" "tf_snmpalarm" {
	trapname = "CPU-USAGE"
	thresholdvalue = 15
	normalvalue    = 10
	state          = "DISABLED"
	severity       = "Minor"
	}
  
`
const testAccSnmpalarm_update = `
resource "citrixadc_snmpalarm" "tf_snmpalarm" {
	trapname = "CPU-USAGE"
	thresholdvalue = 20
	normalvalue    = 15
	state          = "ENABLED"
	severity       = "Major"
	}
  
`

func TestAccSnmpalarm_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccSnmpalarm_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSnmpalarmExist("citrixadc_snmpalarm.tf_snmpalarm", nil),
					resource.TestCheckResourceAttr("citrixadc_snmpalarm.tf_snmpalarm", "thresholdvalue", "15"),
					resource.TestCheckResourceAttr("citrixadc_snmpalarm.tf_snmpalarm", "normalvalue", "10"),
					resource.TestCheckResourceAttr("citrixadc_snmpalarm.tf_snmpalarm", "state", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_snmpalarm.tf_snmpalarm", "severity", "Minor"),
				),
			},
			{
				Config: testAccSnmpalarm_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSnmpalarmExist("citrixadc_snmpalarm.tf_snmpalarm", nil),
					resource.TestCheckResourceAttr("citrixadc_snmpalarm.tf_snmpalarm", "thresholdvalue", "20"),
					resource.TestCheckResourceAttr("citrixadc_snmpalarm.tf_snmpalarm", "normalvalue", "15"),
					resource.TestCheckResourceAttr("citrixadc_snmpalarm.tf_snmpalarm", "state", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_snmpalarm.tf_snmpalarm", "severity", "Major"),
				),
			},
		},
	})
}

func TestAccSnmpalarm_import(t *testing.T) {
	const resAddr = "citrixadc_snmpalarm.tf_snmpalarm"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{Config: testAccSnmpalarm_basic},
			{
				Config:                  testAccSnmpalarm_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccCheckSnmpalarmExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No snmpalarm name is set")
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
		data, err := client.FindResource(service.Snmpalarm.Type(), rs.Primary.Attributes["trapclass"])

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("snmpalarm %s not found", n)
		}

		return nil
	}
}

func TestAccSnmpalarm_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccSnmpalarm_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckSnmpalarmExist("citrixadc_snmpalarm.tf_snmpalarm", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccSnmpalarm_basic,
				Check:                    resource.ComposeTestCheckFunc(testAccCheckSnmpalarmExist("citrixadc_snmpalarm.tf_snmpalarm", nil)),
			},
		},
	})
}

func TestAccSnmpalarmDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccSnmpalarmDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_snmpalarm.tf_snmpalarm_ds", "trapname", "CPU-USAGE"),
					resource.TestCheckResourceAttr("data.citrixadc_snmpalarm.tf_snmpalarm_ds", "thresholdvalue", "25"),
					resource.TestCheckResourceAttr("data.citrixadc_snmpalarm.tf_snmpalarm_ds", "normalvalue", "20"),
				),
			},
		},
	})
}

// The snmpalarm unset test covers the alarm-independent unset-eligible
// attributes (logging, severity) on the MEMORY alarm. thresholdvalue/normalvalue
// have no documented server default, time is alarm-specific and not echoed on
// GET, and state is applied through the enable/disable actions -- so they are
// excluded from unset.
const testAccSnmpalarm_unset_step1 = `
resource "citrixadc_snmpalarm" "tf_unset" {
	trapname = "MEMORY"
	logging  = "DISABLED"
	severity = "Critical"
}
`

const testAccSnmpalarm_unset_step2 = `
resource "citrixadc_snmpalarm" "tf_unset" {
	trapname = "MEMORY"
	# logging and severity removed from config -> the provider must unset them
	# (revert to NITRO defaults: logging=ENABLED, severity=Unknown).
}
`

func TestAccSnmpalarm_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccSnmpalarm_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSnmpalarmExist("citrixadc_snmpalarm.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_snmpalarm.tf_unset", "logging", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_snmpalarm.tf_unset", "severity", "Critical"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from
				// the appliance) reverts to the documented NITRO defaults, and the
				// implicit post-apply plan must be empty.
				Config: testAccSnmpalarm_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSnmpalarmExist("citrixadc_snmpalarm.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_snmpalarm.tf_unset", "logging", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_snmpalarm.tf_unset", "severity", "Unknown"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckSnmpalarmADCValue("MEMORY", "logging", "ENABLED"),
					testAccCheckSnmpalarmADCValue("MEMORY", "severity", "Unknown"),
				),
			},
		},
	})
}

// testAccCheckSnmpalarmADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted
// it.
func testAccCheckSnmpalarmADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Snmpalarm.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("snmpalarm %s not found on appliance", name)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("snmpalarm %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

const testAccSnmpalarmDataSource_basic = `

resource "citrixadc_snmpalarm" "tf_snmpalarm_ds" {
	trapname       = "CPU-USAGE"
	thresholdvalue = 25
	normalvalue    = 20
	state          = "ENABLED"
	severity       = "Warning"
}

data "citrixadc_snmpalarm" "tf_snmpalarm_ds" {
	trapname = citrixadc_snmpalarm.tf_snmpalarm_ds.trapname
	depends_on = [citrixadc_snmpalarm.tf_snmpalarm_ds]
}
`
