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

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testAccauditsyslogparams_basic = `

	resource "citrixadc_auditsyslogparams" "tf_auditsyslogparams" {
		dateformat = "DDMMYYYY"
		loglevel   = ["EMERGENCY"]
		tcp        = "ALL"
		protocolviolations = "NONE"
		streamanalytics = "DISABLED"
	}
`
const testAccauditsyslogparams_update = `

	resource "citrixadc_auditsyslogparams" "tf_auditsyslogparams" {
		dateformat = "MMDDYYYY"
		loglevel   = ["EMERGENCY"]
		tcp        = "NONE"
		protocolviolations = "ALL"
		streamanalytics = "ENABLED"
	}
`

const testAccAuditsyslogparamsDataSource_basic = `

	resource "citrixadc_auditsyslogparams" "tf_auditsyslogparams" {
		dateformat = "DDMMYYYY"
		loglevel   = ["EMERGENCY"]
		tcp        = "ALL"
		protocolviolations = "NONE"
		streamanalytics = "DISABLED"
	}

	data "citrixadc_auditsyslogparams" "tf_auditsyslogparams_ds" {
		depends_on = [citrixadc_auditsyslogparams.tf_auditsyslogparams]
	}
`

func TestAccauditsyslogparams_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccauditsyslogparams_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckauditsyslogparamsExist("citrixadc_auditsyslogparams.tf_auditsyslogparams", nil),
					resource.TestCheckResourceAttr("citrixadc_auditsyslogparams.tf_auditsyslogparams", "dateformat", "DDMMYYYY"),
					resource.TestCheckResourceAttr("citrixadc_auditsyslogparams.tf_auditsyslogparams", "tcp", "ALL"),
					resource.TestCheckResourceAttr("citrixadc_auditsyslogparams.tf_auditsyslogparams", "protocolviolations", "NONE"),
					resource.TestCheckResourceAttr("citrixadc_auditsyslogparams.tf_auditsyslogparams", "streamanalytics", "DISABLED"),
				),
			},
			{
				Config: testAccauditsyslogparams_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckauditsyslogparamsExist("citrixadc_auditsyslogparams.tf_auditsyslogparams", nil),
					resource.TestCheckResourceAttr("citrixadc_auditsyslogparams.tf_auditsyslogparams", "dateformat", "MMDDYYYY"),
					resource.TestCheckResourceAttr("citrixadc_auditsyslogparams.tf_auditsyslogparams", "tcp", "NONE"),
					resource.TestCheckResourceAttr("citrixadc_auditsyslogparams.tf_auditsyslogparams", "protocolviolations", "ALL"),
					resource.TestCheckResourceAttr("citrixadc_auditsyslogparams.tf_auditsyslogparams", "streamanalytics", "ENABLED"),
				),
			},
		},
	})
}

func TestAccAuditsyslogparams_import(t *testing.T) {
	const resAddr = "citrixadc_auditsyslogparams.tf_auditsyslogparams"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{Config: testAccauditsyslogparams_basic},
			{
				Config:                  testAccauditsyslogparams_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccCheckauditsyslogparamsExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No auditsyslogparams name is set")
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
		data, err := client.FindResource(service.Auditsyslogparams.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("auditsyslogparams %s not found", n)
		}

		return nil
	}
}

func TestAccAuditsyslogparams_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccauditsyslogparams_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckauditsyslogparamsExist("citrixadc_auditsyslogparams.tf_auditsyslogparams", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccauditsyslogparams_basic,
				Check:                    resource.ComposeTestCheckFunc(testAccCheckauditsyslogparamsExist("citrixadc_auditsyslogparams.tf_auditsyslogparams", nil)),
			},
		},
	})
}

// auditsyslogparams is a singleton (unnamed) resource. The unset test sets the
// unset-eligible attributes to non-default values, then removes them from
// config so the provider must unset them (revert to the documented NITRO
// defaults). Only attributes that are also present in the basic config are
// wired for unset, so their added schema Defaults never activate in the basic /
// sdkv2StateUpgrade tests and cannot regress them.
const testAccauditsyslogparams_unset_step1 = `
	resource "citrixadc_auditsyslogparams" "tf_unset" {
		dateformat         = "DDMMYYYY"
		tcp                = "ALL"
		protocolviolations = "ALL"
		streamanalytics    = "ENABLED"
	}
`

const testAccauditsyslogparams_unset_step2 = `
	resource "citrixadc_auditsyslogparams" "tf_unset" {
		# All unset-eligible attributes removed from config -> the provider must
		# unset them (revert to NITRO defaults).
	}
`

func TestAccAuditsyslogparams_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccauditsyslogparams_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckauditsyslogparamsExist("citrixadc_auditsyslogparams.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_auditsyslogparams.tf_unset", "dateformat", "DDMMYYYY"),
					resource.TestCheckResourceAttr("citrixadc_auditsyslogparams.tf_unset", "tcp", "ALL"),
					resource.TestCheckResourceAttr("citrixadc_auditsyslogparams.tf_unset", "protocolviolations", "ALL"),
					resource.TestCheckResourceAttr("citrixadc_auditsyslogparams.tf_unset", "streamanalytics", "ENABLED"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from
				// the appliance) reverts to the documented NITRO defaults, and the
				// implicit post-apply plan must be empty.
				Config: testAccauditsyslogparams_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckauditsyslogparamsExist("citrixadc_auditsyslogparams.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_auditsyslogparams.tf_unset", "dateformat", "MMDDYYYY"),
					resource.TestCheckResourceAttr("citrixadc_auditsyslogparams.tf_unset", "tcp", "NONE"),
					resource.TestCheckResourceAttr("citrixadc_auditsyslogparams.tf_unset", "protocolviolations", "NONE"),
					resource.TestCheckResourceAttr("citrixadc_auditsyslogparams.tf_unset", "streamanalytics", "DISABLED"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckAuditsyslogparamsADCValue("dateformat", "MMDDYYYY"),
					testAccCheckAuditsyslogparamsADCValue("tcp", "NONE"),
				),
			},
		},
	})
}

// testAccCheckAuditsyslogparamsADCValue asserts an attribute's value directly on
// the appliance (not just in Terraform state), proving the unset actually
// reverted it.
func testAccCheckAuditsyslogparamsADCValue(attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Auditsyslogparams.Type(), "")
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("auditsyslogparams not found on appliance")
		}
		got := fmt.Sprintf("%v", data[attr])
		if got != want {
			return fmt.Errorf("auditsyslogparams: appliance attr %q = %q, want %q (unset did not revert it)", attr, got, want)
		}
		return nil
	}
}

func TestAccAuditsyslogparamsDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccAuditsyslogparamsDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_auditsyslogparams.tf_auditsyslogparams_ds", "dateformat", "DDMMYYYY"),
					resource.TestCheckResourceAttr("data.citrixadc_auditsyslogparams.tf_auditsyslogparams_ds", "tcp", "ALL"),
					resource.TestCheckResourceAttr("data.citrixadc_auditsyslogparams.tf_auditsyslogparams_ds", "protocolviolations", "NONE"),
					resource.TestCheckResourceAttr("data.citrixadc_auditsyslogparams.tf_auditsyslogparams_ds", "streamanalytics", "DISABLED"),
				),
			},
		},
	})
}
