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

const testAccAppflowaction_basic = `	

	resource "citrixadc_appflowaction" "tf_appflowaction" {
		name            = "test_action"
		collectors      = [citrixadc_appflowcollector.tf_appflowcollector.name ]
		securityinsight = "ENABLED"
		botinsight      = "ENABLED"
		videoanalytics  = "ENABLED"
	}
	resource "citrixadc_appflowcollector" "tf_appflowcollector" {
		name      = "tf_collector"
		ipaddress = "192.168.2.2"
		port      = 80
	}

`

const testAccAppflowaction_update = `	

	resource "citrixadc_appflowaction" "tf_appflowaction" {
		name            = "test_action"
		collectors      = [citrixadc_appflowcollector.tf_appflowcollector.name]
		securityinsight = "DISABLED"
		botinsight      = "DISABLED"
		videoanalytics  = "DISABLED"
	}

	resource "citrixadc_appflowcollector" "tf_appflowcollector" {
		name      = "tf_collector"
		ipaddress = "192.168.2.2"
		port      = 80
	}

`

func TestAccAppflowaction_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAppflowactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAppflowaction_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppflowactionExist("citrixadc_appflowaction.tf_appflowaction", nil),
					resource.TestCheckResourceAttr("citrixadc_appflowaction.tf_appflowaction", "name", "test_action"),
					resource.TestCheckResourceAttr("citrixadc_appflowaction.tf_appflowaction", "securityinsight", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowaction.tf_appflowaction", "videoanalytics", "ENABLED"),
				),
			},
			{
				Config: testAccAppflowaction_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppflowactionExist("citrixadc_appflowaction.tf_appflowaction", nil),
					resource.TestCheckResourceAttr("citrixadc_appflowaction.tf_appflowaction", "name", "test_action"),
					resource.TestCheckResourceAttr("citrixadc_appflowaction.tf_appflowaction", "securityinsight", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowaction.tf_appflowaction", "videoanalytics", "DISABLED"),
				),
			},
		},
	})
}

func testAccCheckAppflowactionExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No appflowaction name is set")
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
		data, err := client.FindResource(service.Appflowaction.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("appflowaction %s not found", n)
		}

		return nil
	}
}

func testAccCheckAppflowactionDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_appflowaction" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Appflowaction.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("appflowaction %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccAppflowactionDataSource_basic = `

	resource "citrixadc_appflowaction" "tf_appflowaction" {
		name            = "test_action"
		collectors      = [citrixadc_appflowcollector.tf_appflowcollector.name ]
		securityinsight = "ENABLED"
		botinsight      = "ENABLED"
		videoanalytics  = "ENABLED"
	}
	resource "citrixadc_appflowcollector" "tf_appflowcollector" {
		name      = "tf_collector"
		ipaddress = "192.168.2.2"
		port      = 80
	}

	data "citrixadc_appflowaction" "tf_appflowaction" {
		name = citrixadc_appflowaction.tf_appflowaction.name
	}
`

func TestAccAppflowactionDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccAppflowactionDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_appflowaction.tf_appflowaction", "name", "test_action"),
					resource.TestCheckResourceAttr("data.citrixadc_appflowaction.tf_appflowaction", "securityinsight", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_appflowaction.tf_appflowaction", "botinsight", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_appflowaction.tf_appflowaction", "videoanalytics", "ENABLED"),
				),
			},
		},
	})
}

func TestAccAppflowaction_import(t *testing.T) {
	const resAddr = "citrixadc_appflowaction.tf_appflowaction"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAppflowactionDestroy,
		Steps: []resource.TestStep{
			{Config: testAccAppflowaction_basic},
			{
				Config:                  testAccAppflowaction_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccAppflowaction_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckAppflowactionDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.0.0"},
				},
				Config: testAccAppflowaction_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAppflowactionExist("citrixadc_appflowaction.tf_appflowaction", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccAppflowaction_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAppflowactionExist("citrixadc_appflowaction.tf_appflowaction", nil)),
			},
		},
	})
}

func TestAccAppflowaction_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_appflowaction.tf_appflowaction"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAppflowactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAppflowaction_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAppflowactionExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Appflowaction.Type(), "test_action"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccAppflowaction_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAppflowactionExist(resAddr, nil)),
			},
		},
	})
}

const testAccAppflowaction_unset_step1 = `
	resource "citrixadc_appflowaction" "tf_unset" {
		name                   = "tf_test_appflowaction_unset"
		botinsight             = "ENABLED"
		ciinsight              = "ENABLED"
		clientsidemeasurements = "ENABLED"
		distributionalgorithm  = "ENABLED"
		pagetracking           = "ENABLED"
		securityinsight        = "ENABLED"
		videoanalytics         = "ENABLED"
		webinsight             = "DISABLED"
	}
`

const testAccAppflowaction_unset_step2 = `
	resource "citrixadc_appflowaction" "tf_unset" {
		name = "tf_test_appflowaction_unset"
		# All unset-eligible attributes removed from config -> the provider must
		# unset them (revert to NITRO defaults).
	}
`

func TestAccAppflowaction_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAppflowactionDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccAppflowaction_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppflowactionExist("citrixadc_appflowaction.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_appflowaction.tf_unset", "botinsight", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowaction.tf_unset", "ciinsight", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowaction.tf_unset", "clientsidemeasurements", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowaction.tf_unset", "distributionalgorithm", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowaction.tf_unset", "pagetracking", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowaction.tf_unset", "securityinsight", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowaction.tf_unset", "videoanalytics", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowaction.tf_unset", "webinsight", "DISABLED"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from
				// the appliance) reverts to the documented NITRO defaults, and the
				// implicit post-apply plan must be empty.
				Config: testAccAppflowaction_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppflowactionExist("citrixadc_appflowaction.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_appflowaction.tf_unset", "botinsight", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowaction.tf_unset", "ciinsight", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowaction.tf_unset", "clientsidemeasurements", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowaction.tf_unset", "distributionalgorithm", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowaction.tf_unset", "pagetracking", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowaction.tf_unset", "securityinsight", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowaction.tf_unset", "videoanalytics", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowaction.tf_unset", "webinsight", "ENABLED"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckAppflowactionADCValue("tf_test_appflowaction_unset", "botinsight", "DISABLED"),
					testAccCheckAppflowactionADCValue("tf_test_appflowaction_unset", "securityinsight", "DISABLED"),
					testAccCheckAppflowactionADCValue("tf_test_appflowaction_unset", "webinsight", "ENABLED"),
				),
			},
		},
	})
}

// testAccCheckAppflowactionADCValue asserts an attribute's value directly on
// the appliance (not just in Terraform state), proving the unset actually
// reverted it.
func testAccCheckAppflowactionADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Appflowaction.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("appflowaction %s not found on appliance", name)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("appflowaction %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}
