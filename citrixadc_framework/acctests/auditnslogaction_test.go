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

const testAccAuditnslogaction_basic = `


resource "citrixadc_auditnslogaction" "tf_auditnslogaction" {
	name     = "my_auditnslogaction"
	serverip = "10.222.74.180"
	loglevel = ["ALERT", "CRITICAL"]
	tcp      = "ALL"
	acl      = "ENABLED"
	protocolviolations = "NONE"
	}
  
`
const testAccAuditnslogaction_update = `


resource "citrixadc_auditnslogaction" "tf_auditnslogaction" {
	name     = "my_auditnslogaction"
	serverip = "10.222.74.180"
	loglevel = ["ALERT", "CRITICAL"]
	tcp      = "NONE"
	acl      = "DISABLED"
	protocolviolations = "ALL"
	}
  
`

func TestAccAuditnslogaction_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuditnslogactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuditnslogaction_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuditnslogactionExist("citrixadc_auditnslogaction.tf_auditnslogaction", nil),
					resource.TestCheckResourceAttr("citrixadc_auditnslogaction.tf_auditnslogaction", "name", "my_auditnslogaction"),
					resource.TestCheckResourceAttr("citrixadc_auditnslogaction.tf_auditnslogaction", "serverip", "10.222.74.180"),
					resource.TestCheckResourceAttr("citrixadc_auditnslogaction.tf_auditnslogaction", "tcp", "ALL"),
					resource.TestCheckResourceAttr("citrixadc_auditnslogaction.tf_auditnslogaction", "acl", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_auditnslogaction.tf_auditnslogaction", "protocolviolations", "NONE"),
				),
			},
			{
				Config: testAccAuditnslogaction_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuditnslogactionExist("citrixadc_auditnslogaction.tf_auditnslogaction", nil),
					resource.TestCheckResourceAttr("citrixadc_auditnslogaction.tf_auditnslogaction", "name", "my_auditnslogaction"),
					resource.TestCheckResourceAttr("citrixadc_auditnslogaction.tf_auditnslogaction", "serverip", "10.222.74.180"),
					resource.TestCheckResourceAttr("citrixadc_auditnslogaction.tf_auditnslogaction", "tcp", "NONE"),
					resource.TestCheckResourceAttr("citrixadc_auditnslogaction.tf_auditnslogaction", "acl", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_auditnslogaction.tf_auditnslogaction", "protocolviolations", "ALL"),
				),
			},
		},
	})
}

func TestAccAuditnslogaction_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_auditnslogaction.tf_auditnslogaction"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuditnslogactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuditnslogaction_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAuditnslogactionExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Auditnslogaction.Type(), "my_auditnslogaction"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccAuditnslogaction_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAuditnslogactionExist(resAddr, nil)),
			},
		},
	})
}

func TestAccAuditnslogaction_import(t *testing.T) {
	const resAddr = "citrixadc_auditnslogaction.tf_auditnslogaction"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuditnslogactionDestroy,
		Steps: []resource.TestStep{
			{Config: testAccAuditnslogaction_basic},
			{
				Config:                  testAccAuditnslogaction_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccCheckAuditnslogactionExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No auditnslogaction name is set")
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
		data, err := client.FindResource(service.Auditnslogaction.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("auditnslogaction %s not found", n)
		}

		return nil
	}
}

func testAccCheckAuditnslogactionDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_auditnslogaction" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Auditnslogaction.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("auditnslogaction %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccAuditnslogaction_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckAuditnslogactionDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.0.0"},
				},
				Config: testAccAuditnslogaction_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuditnslogactionExist("citrixadc_auditnslogaction.tf_auditnslogaction", nil),
				),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccAuditnslogaction_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuditnslogactionExist("citrixadc_auditnslogaction.tf_auditnslogaction", nil),
				),
			},
		},
	})
}

// The auditnslogaction unset test covers the mutable, spec-unsettable
// attributes that revert cleanly to their documented NITRO defaults. loglevel
// (mandatory-on-add) and urlfiltering (omitted from GET at default) are
// intentionally excluded.
const testAccAuditnslogaction_unset_step1 = `
resource "citrixadc_auditnslogaction" "tf_unset" {
  name                 = "tf_auditnslogaction_unset"
  serverip             = "10.222.74.180"
  loglevel             = ["ALERT", "CRITICAL"]
  serverport           = 9999
  dateformat           = "DDMMYYYY"
  logfacility          = "LOCAL3"
  tcp                  = "ALL"
  acl                  = "ENABLED"
  timezone             = "LOCAL_TIME"
  userdefinedauditlog  = "YES"
  appflowexport        = "ENABLED"
  lsn                  = "ENABLED"
  alg                  = "ENABLED"
  subscriberlog        = "ENABLED"
  sslinterception      = "ENABLED"
  contentinspectionlog = "ENABLED"
  protocolviolations   = "ALL"
}
`

const testAccAuditnslogaction_unset_step2 = `
resource "citrixadc_auditnslogaction" "tf_unset" {
  name     = "tf_auditnslogaction_unset"
  serverip = "10.222.74.180"
  loglevel = ["ALERT", "CRITICAL"]
  # All unset-eligible attributes removed from config -> the provider must
  # unset them (revert to NITRO defaults).
}
`

func TestAccAuditnslogaction_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuditnslogactionDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccAuditnslogaction_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuditnslogactionExist("citrixadc_auditnslogaction.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_auditnslogaction.tf_unset", "serverport", "9999"),
					resource.TestCheckResourceAttr("citrixadc_auditnslogaction.tf_unset", "dateformat", "DDMMYYYY"),
					resource.TestCheckResourceAttr("citrixadc_auditnslogaction.tf_unset", "logfacility", "LOCAL3"),
					resource.TestCheckResourceAttr("citrixadc_auditnslogaction.tf_unset", "tcp", "ALL"),
					resource.TestCheckResourceAttr("citrixadc_auditnslogaction.tf_unset", "acl", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_auditnslogaction.tf_unset", "timezone", "LOCAL_TIME"),
					resource.TestCheckResourceAttr("citrixadc_auditnslogaction.tf_unset", "userdefinedauditlog", "YES"),
					resource.TestCheckResourceAttr("citrixadc_auditnslogaction.tf_unset", "appflowexport", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_auditnslogaction.tf_unset", "lsn", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_auditnslogaction.tf_unset", "alg", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_auditnslogaction.tf_unset", "subscriberlog", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_auditnslogaction.tf_unset", "sslinterception", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_auditnslogaction.tf_unset", "contentinspectionlog", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_auditnslogaction.tf_unset", "protocolviolations", "ALL"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from
				// the appliance) reverts to the documented NITRO defaults, and the
				// implicit post-apply plan must be empty.
				Config: testAccAuditnslogaction_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuditnslogactionExist("citrixadc_auditnslogaction.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_auditnslogaction.tf_unset", "serverport", "3023"),
					resource.TestCheckResourceAttr("citrixadc_auditnslogaction.tf_unset", "dateformat", "MMDDYYYY"),
					resource.TestCheckResourceAttr("citrixadc_auditnslogaction.tf_unset", "logfacility", "LOCAL0"),
					resource.TestCheckResourceAttr("citrixadc_auditnslogaction.tf_unset", "tcp", "NONE"),
					resource.TestCheckResourceAttr("citrixadc_auditnslogaction.tf_unset", "acl", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_auditnslogaction.tf_unset", "timezone", "GMT_TIME"),
					resource.TestCheckResourceAttr("citrixadc_auditnslogaction.tf_unset", "userdefinedauditlog", "NO"),
					resource.TestCheckResourceAttr("citrixadc_auditnslogaction.tf_unset", "appflowexport", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_auditnslogaction.tf_unset", "lsn", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_auditnslogaction.tf_unset", "alg", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_auditnslogaction.tf_unset", "subscriberlog", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_auditnslogaction.tf_unset", "sslinterception", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_auditnslogaction.tf_unset", "contentinspectionlog", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_auditnslogaction.tf_unset", "protocolviolations", "NONE"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckAuditnslogactionADCValue("tf_auditnslogaction_unset", "tcp", "NONE"),
					testAccCheckAuditnslogactionADCValue("tf_auditnslogaction_unset", "acl", "DISABLED"),
					testAccCheckAuditnslogactionADCValue("tf_auditnslogaction_unset", "protocolviolations", "NONE"),
				),
			},
		},
	})
}

// testAccCheckAuditnslogactionADCValue asserts an attribute's value directly on
// the appliance (not just in Terraform state), proving the unset actually
// reverted it.
func testAccCheckAuditnslogactionADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Auditnslogaction.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("auditnslogaction %s not found on appliance", name)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("auditnslogaction %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func TestAccAuditnslogactionDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccAuditnslogactionDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_auditnslogaction.tf_auditnslogaction_ds", "name", "tf_auditnslogaction_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_auditnslogaction.tf_auditnslogaction_ds", "serverip", "10.222.74.180"),
					resource.TestCheckResourceAttr("data.citrixadc_auditnslogaction.tf_auditnslogaction_ds", "tcp", "ALL"),
					resource.TestCheckResourceAttr("data.citrixadc_auditnslogaction.tf_auditnslogaction_ds", "acl", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_auditnslogaction.tf_auditnslogaction_ds", "protocolviolations", "NONE"),
				),
			},
		},
	})
}

const testAccAuditnslogactionDataSource_basic = `

resource "citrixadc_auditnslogaction" "tf_auditnslogaction_ds" {
    name     = "tf_auditnslogaction_ds"
    serverip = "10.222.74.180"
    loglevel = ["ALERT", "CRITICAL"]
    tcp      = "ALL"
    acl      = "ENABLED"
    protocolviolations = "NONE"
}

data "citrixadc_auditnslogaction" "tf_auditnslogaction_ds" {
    name = citrixadc_auditnslogaction.tf_auditnslogaction_ds.name
    depends_on = [citrixadc_auditnslogaction.tf_auditnslogaction_ds]
}

`
