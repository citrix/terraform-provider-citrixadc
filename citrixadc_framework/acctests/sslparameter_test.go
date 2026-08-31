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

const testAccSslparameter_basic = `
	resource "citrixadc_sslparameter" "default" {
		denysslreneg   = "NONSECURE"
		defaultprofile = "ENABLED"
		operationqueuelimit = 4096
	}
`
const testAccSslparameter_basic_update = `
	resource "citrixadc_sslparameter" "default" {
		denysslreneg   = "ALL"
		defaultprofile = "ENABLED"
		operationqueuelimit = 4088
	}
`

const testAccSslparameterDataSource_basic = `
	resource "citrixadc_sslparameter" "default" {
		denysslreneg   = "NONSECURE"
		defaultprofile = "ENABLED"
		operationqueuelimit = 4096
	}

	data "citrixadc_sslparameter" "default" {
		depends_on = [citrixadc_sslparameter.default]
	}
`

func TestAccSslparameter_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// sslparameter resource do not have DELETE operation
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				Config: testAccSslparameter_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslparameterExist("citrixadc_sslparameter.default", nil),
					resource.TestCheckResourceAttr("citrixadc_sslparameter.default", "denysslreneg", "NONSECURE"),
					resource.TestCheckResourceAttr("citrixadc_sslparameter.default", "defaultprofile", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslparameter.default", "operationqueuelimit", "4096"),
				),
			},
			{
				Config: testAccSslparameter_basic_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslparameterExist("citrixadc_sslparameter.default", nil),
					resource.TestCheckResourceAttr("citrixadc_sslparameter.default", "denysslreneg", "ALL"),
					resource.TestCheckResourceAttr("citrixadc_sslparameter.default", "defaultprofile", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslparameter.default", "operationqueuelimit", "4088"),
				),
			},
		},
	})
}

func testAccCheckSslparameterExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lb vserver name is set")
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
		data, err := client.FindResource(service.Sslparameter.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("SSL Parameter %s not found", n)
		}

		return nil
	}
}

func TestAccSslparameter_import(t *testing.T) {
	const resAddr = "citrixadc_sslparameter.default"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{Config: testAccSslparameter_basic},
			{
				Config:                  testAccSslparameter_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccSslparameter_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.0.0"},
				},
				Config: testAccSslparameter_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckSslparameterExist("citrixadc_sslparameter.default", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccSslparameter_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckSslparameterExist("citrixadc_sslparameter.default", nil)),
			},
		},
	})
}

// sslparameter is a singleton config resource. Step 1 sets a set of
// unset-eligible attributes to valid NON-default values; step 2 removes them
// from config so the provider must unset them (revert to the documented NITRO
// defaults). Each attribute wired here has a documented server default and
// unsets cleanly on a standalone appliance.
const testAccSslparameter_unset_step1 = `
	resource "citrixadc_sslparameter" "unset" {
		crlmemorysizemb          = 500
		dropreqwithnohostheader  = "YES"
		encrypttriggerpktcount   = 40
		insertcertspace          = "NO"
		insertionencoding        = "UTF-8"
		ndcppcompliancecertcheck = "YES"
		ocspcachesize            = 50
		pushenctriggertimeout    = 100
		quantumsize              = "16384"
		sendclosenotify          = "NO"
		snihttphostmatch         = "STRICT"
		sslierrorcache           = "ENABLED"
		ssltriggertimeout        = 150
		strictcachecks           = "YES"
		undefactioncontrol       = "NOOP"
		undefactiondata          = "RESET"
	}
`

const testAccSslparameter_unset_step2 = `
	resource "citrixadc_sslparameter" "unset" {
		# All unset-eligible attributes removed from config -> the provider must
		# unset them (revert to NITRO defaults).
	}
`

func TestAccSslparameter_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccSslparameter_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslparameterExist("citrixadc_sslparameter.unset", nil),
					resource.TestCheckResourceAttr("citrixadc_sslparameter.unset", "crlmemorysizemb", "500"),
					resource.TestCheckResourceAttr("citrixadc_sslparameter.unset", "dropreqwithnohostheader", "YES"),
					resource.TestCheckResourceAttr("citrixadc_sslparameter.unset", "encrypttriggerpktcount", "40"),
					resource.TestCheckResourceAttr("citrixadc_sslparameter.unset", "insertcertspace", "NO"),
					resource.TestCheckResourceAttr("citrixadc_sslparameter.unset", "insertionencoding", "UTF-8"),
					resource.TestCheckResourceAttr("citrixadc_sslparameter.unset", "ndcppcompliancecertcheck", "YES"),
					resource.TestCheckResourceAttr("citrixadc_sslparameter.unset", "ocspcachesize", "50"),
					resource.TestCheckResourceAttr("citrixadc_sslparameter.unset", "pushenctriggertimeout", "100"),
					resource.TestCheckResourceAttr("citrixadc_sslparameter.unset", "quantumsize", "16384"),
					resource.TestCheckResourceAttr("citrixadc_sslparameter.unset", "sendclosenotify", "NO"),
					resource.TestCheckResourceAttr("citrixadc_sslparameter.unset", "snihttphostmatch", "STRICT"),
					resource.TestCheckResourceAttr("citrixadc_sslparameter.unset", "sslierrorcache", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslparameter.unset", "ssltriggertimeout", "150"),
					resource.TestCheckResourceAttr("citrixadc_sslparameter.unset", "strictcachecks", "YES"),
					resource.TestCheckResourceAttr("citrixadc_sslparameter.unset", "undefactioncontrol", "NOOP"),
					resource.TestCheckResourceAttr("citrixadc_sslparameter.unset", "undefactiondata", "RESET"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from
				// the appliance) reverts to the documented NITRO defaults, and the
				// implicit post-apply plan must be empty.
				Config: testAccSslparameter_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslparameterExist("citrixadc_sslparameter.unset", nil),
					resource.TestCheckResourceAttr("citrixadc_sslparameter.unset", "crlmemorysizemb", "256"),
					resource.TestCheckResourceAttr("citrixadc_sslparameter.unset", "dropreqwithnohostheader", "NO"),
					resource.TestCheckResourceAttr("citrixadc_sslparameter.unset", "encrypttriggerpktcount", "45"),
					resource.TestCheckResourceAttr("citrixadc_sslparameter.unset", "insertcertspace", "YES"),
					resource.TestCheckResourceAttr("citrixadc_sslparameter.unset", "insertionencoding", "Unicode"),
					resource.TestCheckResourceAttr("citrixadc_sslparameter.unset", "ndcppcompliancecertcheck", "NO"),
					resource.TestCheckResourceAttr("citrixadc_sslparameter.unset", "ocspcachesize", "10"),
					resource.TestCheckResourceAttr("citrixadc_sslparameter.unset", "pushenctriggertimeout", "1"),
					resource.TestCheckResourceAttr("citrixadc_sslparameter.unset", "quantumsize", "8192"),
					resource.TestCheckResourceAttr("citrixadc_sslparameter.unset", "sendclosenotify", "YES"),
					resource.TestCheckResourceAttr("citrixadc_sslparameter.unset", "snihttphostmatch", "CERT"),
					resource.TestCheckResourceAttr("citrixadc_sslparameter.unset", "sslierrorcache", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslparameter.unset", "ssltriggertimeout", "100"),
					resource.TestCheckResourceAttr("citrixadc_sslparameter.unset", "strictcachecks", "NO"),
					resource.TestCheckResourceAttr("citrixadc_sslparameter.unset", "undefactioncontrol", "CLIENTAUTH"),
					resource.TestCheckResourceAttr("citrixadc_sslparameter.unset", "undefactiondata", "NOOP"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckSslparameterADCValue("strictcachecks", "NO"),
					testAccCheckSslparameterADCValue("quantumsize", "8192"),
					testAccCheckSslparameterADCValue("sendclosenotify", "YES"),
				),
			},
		},
	})
}

// testAccCheckSslparameterADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted
// it.
func testAccCheckSslparameterADCValue(attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Sslparameter.Type(), "")
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("sslparameter not found on appliance")
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("sslparameter: appliance attr %q = %q, want %q (unset did not revert it)", attr, got, want)
		}
		return nil
	}
}

func TestAccSslparameterDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccSslparameterDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_sslparameter.default", "denysslreneg", "NONSECURE"),
					resource.TestCheckResourceAttr("data.citrixadc_sslparameter.default", "defaultprofile", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_sslparameter.default", "operationqueuelimit", "4096"),
				),
			},
		},
	})
}
