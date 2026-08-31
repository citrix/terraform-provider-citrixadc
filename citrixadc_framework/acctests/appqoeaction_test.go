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

const testAccAppqoeaction_basic = `

	resource "citrixadc_appqoeaction" "tf_appqoeaction" {
		name        = "my_appqoeaction"
		priority    = "LOW"
		respondwith = "NS"
		delay       = "30"
	}
`
const testAccAppqoeaction_update = `

	resource "citrixadc_appqoeaction" "tf_appqoeaction" {
		name        = "my_appqoeaction"
		priority    = "HIGH"
		respondwith = "NS"
		delay       = "10"
	}
`

func TestAccAppqoeaction_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAppqoeactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAppqoeaction_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppqoeactionExist("citrixadc_appqoeaction.tf_appqoeaction", nil),
					resource.TestCheckResourceAttr("citrixadc_appqoeaction.tf_appqoeaction", "name", "my_appqoeaction"),
					resource.TestCheckResourceAttr("citrixadc_appqoeaction.tf_appqoeaction", "priority", "LOW"),
					resource.TestCheckResourceAttr("citrixadc_appqoeaction.tf_appqoeaction", "respondwith", "NS"),
					resource.TestCheckResourceAttr("citrixadc_appqoeaction.tf_appqoeaction", "delay", "30"),
				),
			},
			{
				Config: testAccAppqoeaction_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppqoeactionExist("citrixadc_appqoeaction.tf_appqoeaction", nil),
					resource.TestCheckResourceAttr("citrixadc_appqoeaction.tf_appqoeaction", "name", "my_appqoeaction"),
					resource.TestCheckResourceAttr("citrixadc_appqoeaction.tf_appqoeaction", "priority", "HIGH"),
					resource.TestCheckResourceAttr("citrixadc_appqoeaction.tf_appqoeaction", "respondwith", "NS"),
					resource.TestCheckResourceAttr("citrixadc_appqoeaction.tf_appqoeaction", "delay", "10"),
				),
			},
		},
	})
}

func TestAccAppqoeaction_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckAppqoeactionDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.0.0"},
				},
				Config: testAccAppqoeaction_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppqoeactionExist("citrixadc_appqoeaction.tf_appqoeaction", nil),
				),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccAppqoeaction_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppqoeactionExist("citrixadc_appqoeaction.tf_appqoeaction", nil),
				),
			},
		},
	})
}

func testAccCheckAppqoeactionExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No appqoeaction name is set")
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
		data, err := client.FindResource(service.Appqoeaction.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("appqoeaction %s not found", n)
		}

		return nil
	}
}

func testAccCheckAppqoeactionDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_appqoeaction" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Appqoeaction.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("appqoeaction %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccAppqoeaction_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_appqoeaction.tf_appqoeaction"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAppqoeactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAppqoeaction_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAppqoeactionExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Appqoeaction.Type(), "my_appqoeaction"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccAppqoeaction_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAppqoeactionExist(resAddr, nil)),
			},
		},
	})
}

func TestAccAppqoeaction_import(t *testing.T) {
	const resAddr = "citrixadc_appqoeaction.tf_appqoeaction"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAppqoeactionDestroy,
		Steps: []resource.TestStep{
			{Config: testAccAppqoeaction_basic},
			{
				Config:                  testAccAppqoeaction_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

const testAccAppqoeactionDataSource_basic = `

	resource "citrixadc_appqoeaction" "tf_appqoeaction" {
		name        = "my_appqoeaction_ds"
		priority    = "LOW"
		respondwith = "NS"
		delay       = "30"
	}

	data "citrixadc_appqoeaction" "tf_appqoeaction" {
		name = citrixadc_appqoeaction.tf_appqoeaction.name
	}
`

// The appqoeaction unset test covers the unset-eligible attributes that have a
// documented NITRO server default: numretries (default 3) and retryonreset
// (default NO). Step 1 sets them to non-default values; step 2 removes them
// from config, so the provider must unset them (revert to the NITRO defaults).
const testAccAppqoeaction_unset_step1 = `
	resource "citrixadc_appqoeaction" "tf_unset" {
		name         = "tf_test_appqoeaction_unset"
		priority     = "LOW"
		respondwith  = "NS"
		delay        = "30"
		numretries   = "5"
		retryonreset = "YES"
	}
`

const testAccAppqoeaction_unset_step2 = `
	resource "citrixadc_appqoeaction" "tf_unset" {
		name        = "tf_test_appqoeaction_unset"
		priority    = "LOW"
		respondwith = "NS"
		delay       = "30"
		# numretries and retryonreset removed from config -> the provider must
		# unset them (revert to NITRO defaults: 3 and NO).
	}
`

func TestAccAppqoeaction_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAppqoeactionDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccAppqoeaction_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppqoeactionExist("citrixadc_appqoeaction.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_appqoeaction.tf_unset", "numretries", "5"),
					resource.TestCheckResourceAttr("citrixadc_appqoeaction.tf_unset", "retryonreset", "YES"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from
				// the appliance) reverts to the documented NITRO defaults, and the
				// implicit post-apply plan must be empty.
				Config: testAccAppqoeaction_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppqoeactionExist("citrixadc_appqoeaction.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_appqoeaction.tf_unset", "numretries", "3"),
					resource.TestCheckResourceAttr("citrixadc_appqoeaction.tf_unset", "retryonreset", "NO"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckAppqoeactionADCValue("tf_test_appqoeaction_unset", "numretries", "3"),
					testAccCheckAppqoeactionADCValue("tf_test_appqoeaction_unset", "retryonreset", "NO"),
				),
			},
		},
	})
}

// testAccCheckAppqoeactionADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted
// it.
func testAccCheckAppqoeactionADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Appqoeaction.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("appqoeaction %s not found on appliance", name)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("appqoeaction %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func TestAccAppqoeactionDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccAppqoeactionDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_appqoeaction.tf_appqoeaction", "name", "my_appqoeaction_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_appqoeaction.tf_appqoeaction", "priority", "LOW"),
					resource.TestCheckResourceAttr("data.citrixadc_appqoeaction.tf_appqoeaction", "respondwith", "NS"),
					resource.TestCheckResourceAttr("data.citrixadc_appqoeaction.tf_appqoeaction", "delay", "30"),
				),
			},
		},
	})
}
