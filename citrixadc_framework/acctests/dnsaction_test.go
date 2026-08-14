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

const testAccDnsaction_add = `

	resource "citrixadc_dnsprofile" "dnsprofile" {
		dnsprofilename         = "tf_profile1"
		dnsquerylogging        = "DISABLED"
		dnsanswerseclogging    = "DISABLED"
		dnsextendedlogging     = "DISABLED"
		dnserrorlogging        = "DISABLED"
		cacherecords           = "ENABLED"
		cachenegativeresponses = "ENABLED"
		dropmultiqueryrequest  = "DISABLED"
		cacheecsresponses      = "DISABLED"
	}

	resource "citrixadc_dnsaction" "dnsaction" {
		actionname       = "tf_action1"
		actiontype       = "Rewrite_Response"
		ipaddress        = ["192.0.2.20","192.0.2.56","198.51.130.10"]
		dnsprofilename   = citrixadc_dnsprofile.dnsprofile.dnsprofilename
	}
`

const testAccDnsactionDataSource_basic = `

	resource "citrixadc_dnsprofile" "dnsprofile" {
		dnsprofilename         = "tf_profile1"
		dnsquerylogging        = "DISABLED"
		dnsanswerseclogging    = "DISABLED"
		dnsextendedlogging     = "DISABLED"
		dnserrorlogging        = "DISABLED"
		cacherecords           = "ENABLED"
		cachenegativeresponses = "ENABLED"
		dropmultiqueryrequest  = "DISABLED"
		cacheecsresponses      = "DISABLED"
	}

	resource "citrixadc_dnsaction" "dnsaction" {
		actionname       = "tf_action1"
		actiontype       = "Rewrite_Response"
		ipaddress        = ["192.0.2.20","192.0.2.56","198.51.130.10"]
		dnsprofilename   = citrixadc_dnsprofile.dnsprofile.dnsprofilename
	}

	data "citrixadc_dnsaction" "dnsaction_data" {
		actionname = citrixadc_dnsaction.dnsaction.actionname
	}
`

func TestAccDnsaction_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDnsactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDnsaction_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDnsactionExist("citrixadc_dnsaction.dnsaction", nil),
					resource.TestCheckResourceAttr("citrixadc_dnsaction.dnsaction", "actionname", "tf_action1"),
					resource.TestCheckResourceAttr("citrixadc_dnsaction.dnsaction", "actiontype", "Rewrite_Response"),
					resource.TestCheckResourceAttr("citrixadc_dnsaction.dnsaction", "dnsprofilename", "tf_profile1"),
				),
			},
		},
	})
}

func TestAccDnsactionDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDnsactionDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_dnsaction.dnsaction_data", "actionname", "tf_action1"),
					resource.TestCheckResourceAttr("data.citrixadc_dnsaction.dnsaction_data", "actiontype", "Rewrite_Response"),
					resource.TestCheckResourceAttr("data.citrixadc_dnsaction.dnsaction_data", "dnsprofilename", "tf_profile1"),
				),
			},
		},
	})
}

func TestAccDnsaction_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_dnsaction.dnsaction"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDnsactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDnsaction_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckDnsactionExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Dnsaction.Type(), "tf_action1"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccDnsaction_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckDnsactionExist(resAddr, nil)),
			},
		},
	})
}

func TestAccDnsaction_import(t *testing.T) {
	const resAddr = "citrixadc_dnsaction.dnsaction"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDnsactionDestroy,
		Steps: []resource.TestStep{
			{Config: testAccDnsaction_add},
			{
				Config:                  testAccDnsaction_add,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccDnsaction_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckDnsactionDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccDnsaction_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckDnsactionExist("citrixadc_dnsaction.dnsaction", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccDnsaction_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckDnsactionExist("citrixadc_dnsaction.dnsaction", nil)),
			},
		},
	})
}

// The dnsaction unset test covers ttl, the only spec-unsettable mutable
// attribute (it appears in the NITRO unset payload and has a documented
// default of 3600). Step1 sets a non-default ttl; step2 removes it, so the
// provider must unset it and the appliance reverts it to 3600.
const testAccDnsaction_unset_step1 = `
	resource "citrixadc_dnsaction" "tf_unset" {
		actionname = "tf_test_dnsaction_unset"
		actiontype = "Rewrite_Response"
		ipaddress  = ["192.0.2.20", "192.0.2.56"]
		ttl        = 1800
	}
`

const testAccDnsaction_unset_step2 = `
	resource "citrixadc_dnsaction" "tf_unset" {
		actionname = "tf_test_dnsaction_unset"
		actiontype = "Rewrite_Response"
		ipaddress  = ["192.0.2.20", "192.0.2.56"]
		# ttl removed from config -> the provider must unset it (revert to 3600).
	}
`

func TestAccDnsaction_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDnsactionDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default value is applied and persisted.
				Config: testAccDnsaction_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDnsactionExist("citrixadc_dnsaction.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_dnsaction.tf_unset", "ttl", "1800"),
				),
			},
			{
				// Removing ttl must unset it: state (read back from the appliance)
				// reverts to the documented NITRO default, and the implicit
				// post-apply plan must be empty.
				Config: testAccDnsaction_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDnsactionExist("citrixadc_dnsaction.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_dnsaction.tf_unset", "ttl", "3600"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckDnsactionADCValue("tf_test_dnsaction_unset", "ttl", "3600"),
				),
			},
		},
	})
}

// testAccCheckDnsactionADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted
// it.
func testAccCheckDnsactionADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Dnsaction.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("dnsaction %s not found on appliance", name)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("dnsaction %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func testAccCheckDnsactionExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No dnsaction name is set")
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
		data, err := client.FindResource(service.Dnsaction.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("dnsaction %s not found", n)
		}

		return nil
	}
}

func testAccCheckDnsactionDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_dnsaction" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Dnsaction.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("dnsaction %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
