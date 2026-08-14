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

const testAccVideooptimizationdetectionpolicy_add = `
	resource "citrixadc_videooptimizationdetectionaction" "tf_detectionaction" {
		name = "tf_videooptimizationdetectionaction"
		type = "clear_text_abr"
	}
	
	resource "citrixadc_videooptimizationdetectionpolicy" "tf_detectionpolicy" {
		name   = "tf_videooptimizationdetectionpolicy"
		rule   = "true"
		action = citrixadc_videooptimizationdetectionaction.tf_detectionaction.name
	}
`

const testAccVideooptimizationdetectionpolicy_update = `
	resource "citrixadc_videooptimizationdetectionaction" "tf_detectionaction" {
		name = "tf_videooptimizationdetectionaction"
		type = "clear_text_abr"
	}
	
	resource "citrixadc_videooptimizationdetectionpolicy" "tf_detectionpolicy" {
		name   = "tf_videooptimizationdetectionpolicy"
		rule   = "false"
		action = citrixadc_videooptimizationdetectionaction.tf_detectionaction.name
	}
`

const testAccVideooptimizationdetectionpolicyDataSource_basic = `
	resource "citrixadc_videooptimizationdetectionaction" "tf_detectionaction" {
		name = "tf_videooptimizationdetectionaction"
		type = "clear_text_abr"
	}
	
	resource "citrixadc_videooptimizationdetectionpolicy" "tf_detectionpolicy" {
		name   = "tf_videooptimizationdetectionpolicy"
		rule   = "true"
		action = citrixadc_videooptimizationdetectionaction.tf_detectionaction.name
	}

	data "citrixadc_videooptimizationdetectionpolicy" "tf_detectionpolicy" {
		name = citrixadc_videooptimizationdetectionpolicy.tf_detectionpolicy.name
	}
`

func TestAccVideooptimizationdetectionpolicy_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVideooptimizationdetectionpolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVideooptimizationdetectionpolicy_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVideooptimizationdetectionpolicyExist("citrixadc_videooptimizationdetectionpolicy.tf_detectionpolicy", nil),
					resource.TestCheckResourceAttr("citrixadc_videooptimizationdetectionpolicy.tf_detectionpolicy", "name", "tf_videooptimizationdetectionpolicy"),
					resource.TestCheckResourceAttr("citrixadc_videooptimizationdetectionpolicy.tf_detectionpolicy", "rule", "true"),
					resource.TestCheckResourceAttr("citrixadc_videooptimizationdetectionpolicy.tf_detectionpolicy", "action", "tf_videooptimizationdetectionaction"),
				),
			},
			{
				Config: testAccVideooptimizationdetectionpolicy_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVideooptimizationdetectionpolicyExist("citrixadc_videooptimizationdetectionpolicy.tf_detectionpolicy", nil),
					resource.TestCheckResourceAttr("citrixadc_videooptimizationdetectionpolicy.tf_detectionpolicy", "name", "tf_videooptimizationdetectionpolicy"),
					resource.TestCheckResourceAttr("citrixadc_videooptimizationdetectionpolicy.tf_detectionpolicy", "rule", "false"),
					resource.TestCheckResourceAttr("citrixadc_videooptimizationdetectionpolicy.tf_detectionpolicy", "action", "tf_videooptimizationdetectionaction"),
				),
			},
		},
	})
}

func testAccCheckVideooptimizationdetectionpolicyExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No videooptimizationdetectionpolicy name is set")
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
		data, err := client.FindResource("videooptimizationdetectionpolicy", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("videooptimizationdetectionpolicy %s not found", n)
		}

		return nil
	}
}

func testAccCheckVideooptimizationdetectionpolicyDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_videooptimizationdetectionpolicy" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("videooptimizationdetectionpolicy", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("videooptimizationdetectionpolicy %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccVideooptimizationdetectionpolicy_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_videooptimizationdetectionpolicy.tf_detectionpolicy"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVideooptimizationdetectionpolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVideooptimizationdetectionpolicy_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckVideooptimizationdetectionpolicyExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Videooptimizationdetectionpolicy.Type(), "tf_videooptimizationdetectionpolicy"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccVideooptimizationdetectionpolicy_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckVideooptimizationdetectionpolicyExist(resAddr, nil)),
			},
		},
	})
}

func TestAccVideooptimizationdetectionpolicy_import(t *testing.T) {
	const resAddr = "citrixadc_videooptimizationdetectionpolicy.tf_detectionpolicy"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVideooptimizationdetectionpolicyDestroy,
		Steps: []resource.TestStep{
			{Config: testAccVideooptimizationdetectionpolicy_add},
			{
				Config:                  testAccVideooptimizationdetectionpolicy_add,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccVideooptimizationdetectionpolicy_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckVideooptimizationdetectionpolicyDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccVideooptimizationdetectionpolicy_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckVideooptimizationdetectionpolicyExist("citrixadc_videooptimizationdetectionpolicy.tf_detectionpolicy", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccVideooptimizationdetectionpolicy_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckVideooptimizationdetectionpolicyExist("citrixadc_videooptimizationdetectionpolicy.tf_detectionpolicy", nil)),
			},
		},
	})
}

// The videooptimizationdetectionpolicy unset test covers the spec-unsettable
// attributes (comment, logaction, undefaction). Step 1 sets them to non-default
// values; step 2 removes them from config so the provider issues the NITRO
// ?action=unset, reverting them to their appliance defaults (empty).
const testAccVideooptimizationdetectionpolicy_unset_step1 = `
	resource "citrixadc_videooptimizationdetectionaction" "tf_unset_action" {
		name = "tf_vodp_unset_action"
		type = "clear_text_abr"
	}

	resource "citrixadc_auditmessageaction" "tf_unset_msgaction" {
		name              = "tf_vodp_unset_msg"
		loglevel          = "NOTICE"
		stringbuilderexpr = "\"hello\""
		logtonewnslog     = "YES"
	}

	resource "citrixadc_videooptimizationdetectionpolicy" "tf_unset" {
		name        = "tf_vodp_unset"
		rule        = "true"
		action      = citrixadc_videooptimizationdetectionaction.tf_unset_action.name
		comment     = "managed by terraform"
		logaction   = citrixadc_auditmessageaction.tf_unset_msgaction.name
		undefaction = "RESET"
	}
`

const testAccVideooptimizationdetectionpolicy_unset_step2 = `
	resource "citrixadc_videooptimizationdetectionaction" "tf_unset_action" {
		name = "tf_vodp_unset_action"
		type = "clear_text_abr"
	}

	resource "citrixadc_auditmessageaction" "tf_unset_msgaction" {
		name              = "tf_vodp_unset_msg"
		loglevel          = "NOTICE"
		stringbuilderexpr = "\"hello\""
		logtonewnslog     = "YES"
	}

	resource "citrixadc_videooptimizationdetectionpolicy" "tf_unset" {
		name   = "tf_vodp_unset"
		rule   = "true"
		action = citrixadc_videooptimizationdetectionaction.tf_unset_action.name
		# comment, logaction and undefaction removed -> provider must unset them.
	}
`

func TestAccVideooptimizationdetectionpolicy_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVideooptimizationdetectionpolicyDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccVideooptimizationdetectionpolicy_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVideooptimizationdetectionpolicyExist("citrixadc_videooptimizationdetectionpolicy.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_videooptimizationdetectionpolicy.tf_unset", "comment", "managed by terraform"),
					resource.TestCheckResourceAttr("citrixadc_videooptimizationdetectionpolicy.tf_unset", "logaction", "tf_vodp_unset_msg"),
					resource.TestCheckResourceAttr("citrixadc_videooptimizationdetectionpolicy.tf_unset", "undefaction", "RESET"),
				),
			},
			{
				// Removing the attributes must unset them: the appliance reverts
				// them to their defaults (empty) and the implicit post-apply plan
				// must be empty.
				Config: testAccVideooptimizationdetectionpolicy_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVideooptimizationdetectionpolicyExist("citrixadc_videooptimizationdetectionpolicy.tf_unset", nil),
					testAccCheckVideooptimizationdetectionpolicyADCValue("tf_vodp_unset", "comment", ""),
					testAccCheckVideooptimizationdetectionpolicyADCValue("tf_vodp_unset", "logaction", ""),
					testAccCheckVideooptimizationdetectionpolicyADCValue("tf_vodp_unset", "undefaction", ""),
				),
			},
		},
	})
}

// testAccCheckVideooptimizationdetectionpolicyADCValue asserts an attribute's
// value directly on the appliance (not just in Terraform state), proving the
// unset reverted it.
func testAccCheckVideooptimizationdetectionpolicyADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Videooptimizationdetectionpolicy.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("videooptimizationdetectionpolicy %s not found on appliance", name)
		}
		got := ""
		if v, ok := data[attr]; ok && v != nil {
			got = strings.TrimSpace(fmt.Sprintf("%v", v))
		}
		if got != want {
			return fmt.Errorf("videooptimizationdetectionpolicy %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func TestAccVideooptimizationdetectionpolicyDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccVideooptimizationdetectionpolicyDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_videooptimizationdetectionpolicy.tf_detectionpolicy", "name", "tf_videooptimizationdetectionpolicy"),
					resource.TestCheckResourceAttr("data.citrixadc_videooptimizationdetectionpolicy.tf_detectionpolicy", "rule", "true"),
					resource.TestCheckResourceAttr("data.citrixadc_videooptimizationdetectionpolicy.tf_detectionpolicy", "action", "tf_videooptimizationdetectionaction"),
				),
			},
		},
	})
}
