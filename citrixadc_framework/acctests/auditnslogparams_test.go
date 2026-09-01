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

const testAccAuditnslogparams_basic = `

	resource "citrixadc_auditnslogparams" "tf_auditnslogparams" {
		dateformat = "DDMMYYYY"
		loglevel   = ["EMERGENCY"]
		tcp        = "ALL"
		protocolviolations = "NONE"
	}
`
const testAccAuditnslogparams_update = `

	resource "citrixadc_auditnslogparams" "tf_auditnslogparams" {
		dateformat = "MMDDYYYY"
		loglevel   = ["EMERGENCY"]
		tcp        = "NONE"
		protocolviolations = "ALL"
	}
`

func TestAccAuditnslogparams_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccAuditnslogparams_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuditnslogparamsExist("citrixadc_auditnslogparams.tf_auditnslogparams", nil),
					resource.TestCheckResourceAttr("citrixadc_auditnslogparams.tf_auditnslogparams", "dateformat", "DDMMYYYY"),
					resource.TestCheckResourceAttr("citrixadc_auditnslogparams.tf_auditnslogparams", "tcp", "ALL"),
					resource.TestCheckResourceAttr("citrixadc_auditnslogparams.tf_auditnslogparams", "protocolviolations", "NONE"),
				),
			},
			{
				Config: testAccAuditnslogparams_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuditnslogparamsExist("citrixadc_auditnslogparams.tf_auditnslogparams", nil),
					resource.TestCheckResourceAttr("citrixadc_auditnslogparams.tf_auditnslogparams", "dateformat", "MMDDYYYY"),
					resource.TestCheckResourceAttr("citrixadc_auditnslogparams.tf_auditnslogparams", "tcp", "NONE"),
					resource.TestCheckResourceAttr("citrixadc_auditnslogparams.tf_auditnslogparams", "protocolviolations", "ALL"),
				),
			},
		},
	})
}

func testAccCheckAuditnslogparamsExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No auditnslogparams name is set")
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
		data, err := client.FindResource(service.Auditnslogparams.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("auditnslogparams %s not found", n)
		}

		return nil
	}
}

func TestAccAuditnslogparamsDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccAuditnslogparamsDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.citrixadc_auditnslogparams.tf_auditnslogparams", "id"),
					resource.TestCheckResourceAttr("data.citrixadc_auditnslogparams.tf_auditnslogparams", "dateformat", "DDMMYYYY"),
					resource.TestCheckResourceAttr("data.citrixadc_auditnslogparams.tf_auditnslogparams", "tcp", "ALL"),
					resource.TestCheckResourceAttr("data.citrixadc_auditnslogparams.tf_auditnslogparams", "protocolviolations", "NONE"),
				),
			},
		},
	})
}

func TestAccAuditnslogparams_import(t *testing.T) {
	const resAddr = "citrixadc_auditnslogparams.tf_auditnslogparams"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{Config: testAccAuditnslogparams_basic},
			{
				Config:                  testAccAuditnslogparams_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccAuditnslogparams_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.0.0"},
				},
				Config: testAccAuditnslogparams_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAuditnslogparamsExist("citrixadc_auditnslogparams.tf_auditnslogparams", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccAuditnslogparams_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAuditnslogparamsExist("citrixadc_auditnslogparams.tf_auditnslogparams", nil)),
			},
		},
	})
}

// auditnslogparams is a singleton config resource. The unset test sets every
// unset-eligible attribute to a valid non-default value, then removes them all
// from config: the provider must issue a NITRO ?action=unset so the appliance
// reverts each to its documented default.
const testAccAuditnslogparams_unset_step1 = `
	resource "citrixadc_auditnslogparams" "tf_unset" {
		acl                  = "ENABLED"
		alg                  = "ENABLED"
		appflowexport        = "ENABLED"
		contentinspectionlog = "ENABLED"
		dateformat           = "DDMMYYYY"
		logfacility          = "LOCAL1"
		lsn                  = "ENABLED"
		protocolviolations   = "ALL"
		sslinterception      = "ENABLED"
		subscriberlog        = "ENABLED"
		tcp                  = "ALL"
		timezone             = "LOCAL_TIME"
		userdefinedauditlog  = "YES"
	}
`

const testAccAuditnslogparams_unset_step2 = `
	resource "citrixadc_auditnslogparams" "tf_unset" {
		# All unset-eligible attributes removed from config -> the provider must
		# unset them (revert to NITRO defaults).
	}
`

func TestAccAuditnslogparams_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccAuditnslogparams_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuditnslogparamsExist("citrixadc_auditnslogparams.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_auditnslogparams.tf_unset", "acl", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_auditnslogparams.tf_unset", "alg", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_auditnslogparams.tf_unset", "appflowexport", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_auditnslogparams.tf_unset", "contentinspectionlog", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_auditnslogparams.tf_unset", "dateformat", "DDMMYYYY"),
					resource.TestCheckResourceAttr("citrixadc_auditnslogparams.tf_unset", "logfacility", "LOCAL1"),
					resource.TestCheckResourceAttr("citrixadc_auditnslogparams.tf_unset", "lsn", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_auditnslogparams.tf_unset", "protocolviolations", "ALL"),
					resource.TestCheckResourceAttr("citrixadc_auditnslogparams.tf_unset", "sslinterception", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_auditnslogparams.tf_unset", "subscriberlog", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_auditnslogparams.tf_unset", "tcp", "ALL"),
					resource.TestCheckResourceAttr("citrixadc_auditnslogparams.tf_unset", "timezone", "LOCAL_TIME"),
					resource.TestCheckResourceAttr("citrixadc_auditnslogparams.tf_unset", "userdefinedauditlog", "YES"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from
				// the appliance) reverts to the documented NITRO defaults, and the
				// implicit post-apply plan must be empty.
				Config: testAccAuditnslogparams_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuditnslogparamsExist("citrixadc_auditnslogparams.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_auditnslogparams.tf_unset", "acl", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_auditnslogparams.tf_unset", "alg", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_auditnslogparams.tf_unset", "appflowexport", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_auditnslogparams.tf_unset", "contentinspectionlog", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_auditnslogparams.tf_unset", "dateformat", "MMDDYYYY"),
					resource.TestCheckResourceAttr("citrixadc_auditnslogparams.tf_unset", "logfacility", "LOCAL0"),
					resource.TestCheckResourceAttr("citrixadc_auditnslogparams.tf_unset", "lsn", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_auditnslogparams.tf_unset", "protocolviolations", "NONE"),
					resource.TestCheckResourceAttr("citrixadc_auditnslogparams.tf_unset", "sslinterception", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_auditnslogparams.tf_unset", "subscriberlog", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_auditnslogparams.tf_unset", "tcp", "NONE"),
					resource.TestCheckResourceAttr("citrixadc_auditnslogparams.tf_unset", "timezone", "GMT_TIME"),
					resource.TestCheckResourceAttr("citrixadc_auditnslogparams.tf_unset", "userdefinedauditlog", "NO"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckAuditnslogparamsADCValue("acl", "DISABLED"),
					testAccCheckAuditnslogparamsADCValue("tcp", "NONE"),
					testAccCheckAuditnslogparamsADCValue("userdefinedauditlog", "NO"),
				),
			},
		},
	})
}

// testAccCheckAuditnslogparamsADCValue asserts an attribute's value directly on
// the appliance (not just in Terraform state), proving the unset actually
// reverted it.
func testAccCheckAuditnslogparamsADCValue(attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Auditnslogparams.Type(), "")
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("auditnslogparams not found on appliance")
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("auditnslogparams: appliance attr %q = %q, want %q (unset did not revert it)", attr, got, want)
		}
		return nil
	}
}

const testAccAuditnslogparamsDataSource_basic = `

resource "citrixadc_auditnslogparams" "tf_auditnslogparams" {
	dateformat = "DDMMYYYY"
	loglevel   = ["EMERGENCY"]
	tcp        = "ALL"
	protocolviolations = "NONE"
}

data "citrixadc_auditnslogparams" "tf_auditnslogparams" {
	depends_on = [citrixadc_auditnslogparams.tf_auditnslogparams]
}
`
