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

const testAccIcaaction_basic = `

	resource "citrixadc_icaaction" "tf_icaaction" {
		name              = "my_ica_action"
		accessprofilename = "default_ica_accessprofile"
	}
  
`

const testAccIcaaction_update = `

	resource "citrixadc_icaaccessprofile" "tf_icaaccessprofile" {
		name                   = "my_profile"
		connectclientlptports  = "DEFAULT"
		localremotedatasharing = "DEFAULT"
	}

	resource "citrixadc_icaaction" "tf_icaaction" {
		name              = "my_ica_action"
		accessprofilename = citrixadc_icaaccessprofile.tf_icaaccessprofile.name
	}
  
`

const testAccIcaactionDataSource_basic = `
	resource "citrixadc_icaaction" "tf_icaaction" {
		name              = "my_ica_action"
		accessprofilename = "default_ica_accessprofile"
	}
	
	data "citrixadc_icaaction" "tf_icaaction_ds" {
		name = citrixadc_icaaction.tf_icaaction.name
	}
`

func TestAccIcaactionDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIcaactionDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_icaaction.tf_icaaction_ds", "name", "my_ica_action"),
					resource.TestCheckResourceAttr("data.citrixadc_icaaction.tf_icaaction_ds", "accessprofilename", "default_ica_accessprofile"),
					resource.TestCheckResourceAttrSet("data.citrixadc_icaaction.tf_icaaction_ds", "id"),
				),
			},
		},
	})
}

func TestAccIcaaction_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckIcaactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccIcaaction_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIcaactionExist("citrixadc_icaaction.tf_icaaction", nil),
					resource.TestCheckResourceAttr("citrixadc_icaaction.tf_icaaction", "name", "my_ica_action"),
					resource.TestCheckResourceAttr("citrixadc_icaaction.tf_icaaction", "accessprofilename", "default_ica_accessprofile"),
				),
			},
			{
				Config: testAccIcaaction_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIcaactionExist("citrixadc_icaaction.tf_icaaction", nil),
					resource.TestCheckResourceAttr("citrixadc_icaaction.tf_icaaction", "name", "my_ica_action"),
					resource.TestCheckResourceAttr("citrixadc_icaaction.tf_icaaction", "accessprofilename", "my_profile"),
				),
			},
		},
	})
}

func TestAccIcaaction_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_icaaction.tf_icaaction"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckIcaactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccIcaaction_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckIcaactionExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Icaaction.Type(), "my_ica_action"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccIcaaction_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckIcaactionExist(resAddr, nil)),
			},
		},
	})
}

func TestAccIcaaction_import(t *testing.T) {
	const resAddr = "citrixadc_icaaction.tf_icaaction"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckIcaactionDestroy,
		Steps: []resource.TestStep{
			{Config: testAccIcaaction_basic},
			{
				Config:                  testAccIcaaction_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccIcaaction_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckIcaactionDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.0.0"},
				},
				Config: testAccIcaaction_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckIcaactionExist("citrixadc_icaaction.tf_icaaction", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccIcaaction_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckIcaactionExist("citrixadc_icaaction.tf_icaaction", nil)),
			},
		},
	})
}

// The icaaction unset test covers the single spec-unsettable, mutable attribute
// accessprofilename. latencyprofilename is also listed in the NITRO unset payload
// but the appliance rejects its unset with "Profile does not exist" (errorcode
// 3248), so it is intentionally excluded.
// A latencyprofilename must remain configured on the action across both steps:
// NITRO only permits unsetting accessprofilename while a latencyprofile is bound
// (otherwise it rejects the unset with "Profile does not exist", errorcode 3248).
// latencyprofilename is therefore a prerequisite here, not itself under test.
const testAccIcaaction_unset_step1 = `
	resource "citrixadc_icalatencyprofile" "tf_lp" {
		name                = "tf_test_icaaction_unset_lp"
		l7latencymonitoring = "ENABLED"
	}

	resource "citrixadc_icaaction" "tf_unset" {
		name               = "tf_test_icaaction_unset"
		accessprofilename  = "default_ica_accessprofile"
		latencyprofilename = citrixadc_icalatencyprofile.tf_lp.name
	}
`

const testAccIcaaction_unset_step2 = `
	resource "citrixadc_icalatencyprofile" "tf_lp" {
		name                = "tf_test_icaaction_unset_lp"
		l7latencymonitoring = "ENABLED"
	}

	resource "citrixadc_icaaction" "tf_unset" {
		name               = "tf_test_icaaction_unset"
		latencyprofilename = citrixadc_icalatencyprofile.tf_lp.name
		# accessprofilename removed from config -> the provider must unset it
		# (revert to the NITRO default, which is empty/absent).
	}
`

func TestAccIcaaction_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckIcaactionDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default value is applied and persisted.
				Config: testAccIcaaction_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIcaactionExist("citrixadc_icaaction.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_icaaction.tf_unset", "accessprofilename", "default_ica_accessprofile"),
					resource.TestCheckResourceAttr("citrixadc_icaaction.tf_unset", "latencyprofilename", "tf_test_icaaction_unset_lp"),
				),
			},
			{
				// Removing the attribute must unset it: state (read back from the
				// appliance) reverts to the NITRO default (empty/absent), and the
				// implicit post-apply plan must be empty.
				Config: testAccIcaaction_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIcaactionExist("citrixadc_icaaction.tf_unset", nil),
					resource.TestCheckNoResourceAttr("citrixadc_icaaction.tf_unset", "accessprofilename"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckIcaactionADCValue("tf_test_icaaction_unset", "accessprofilename", ""),
				),
			},
		},
	})
}

// testAccCheckIcaactionADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted
// it. A nil/absent value on the appliance is treated as the empty string.
func testAccCheckIcaactionADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Icaaction.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("icaaction %s not found on appliance", name)
		}
		got := ""
		if v, ok := data[attr]; ok && v != nil {
			got = strings.TrimSpace(fmt.Sprintf("%v", v))
		}
		if got != want {
			return fmt.Errorf("icaaction %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func testAccCheckIcaactionExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No icaaction name is set")
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
		data, err := client.FindResource("icaaction", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("icaaction %s not found", n)
		}

		return nil
	}
}

func testAccCheckIcaactionDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_icaaction" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("icaaction", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("icaaction %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
