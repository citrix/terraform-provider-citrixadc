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

const testAccContentinspectionpolicy_basic = `

	resource "citrixadc_contentinspectionpolicy" "tf_contentinspectionpolicy" {
		name   = "my_ci_policy"
		rule   = "true"
		action = "RESET"
	}
  
`
const testAccContentinspectionpolicy_update = `

	resource "citrixadc_contentinspectionpolicy" "tf_contentinspectionpolicy" {
		name   = "my_ci_policy"
		rule   = "false"
		action = "DROP"
	}
  
`

const testAccContentinspectionpolicyDataSource_basic = `

	resource "citrixadc_contentinspectionpolicy" "tf_contentinspectionpolicy" {
		name   = "my_ci_policy_ds"
		rule   = "true"
		action = "RESET"
	}

	data "citrixadc_contentinspectionpolicy" "tf_contentinspectionpolicy_ds" {
		name = citrixadc_contentinspectionpolicy.tf_contentinspectionpolicy.name
	}
`

func TestAccContentinspectionpolicy_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckContentinspectionpolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccContentinspectionpolicy_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckContentinspectionpolicyExist("citrixadc_contentinspectionpolicy.tf_contentinspectionpolicy", nil),
					resource.TestCheckResourceAttr("citrixadc_contentinspectionpolicy.tf_contentinspectionpolicy", "name", "my_ci_policy"),
					resource.TestCheckResourceAttr("citrixadc_contentinspectionpolicy.tf_contentinspectionpolicy", "rule", "true"),
					resource.TestCheckResourceAttr("citrixadc_contentinspectionpolicy.tf_contentinspectionpolicy", "action", "RESET"),
				),
			},
			{
				Config: testAccContentinspectionpolicy_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckContentinspectionpolicyExist("citrixadc_contentinspectionpolicy.tf_contentinspectionpolicy", nil),
					resource.TestCheckResourceAttr("citrixadc_contentinspectionpolicy.tf_contentinspectionpolicy", "name", "my_ci_policy"),
					resource.TestCheckResourceAttr("citrixadc_contentinspectionpolicy.tf_contentinspectionpolicy", "rule", "false"),
					resource.TestCheckResourceAttr("citrixadc_contentinspectionpolicy.tf_contentinspectionpolicy", "action", "DROP"),
				),
			},
		},
	})
}

func testAccCheckContentinspectionpolicyExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No contentinspectionpolicy name is set")
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
		data, err := client.FindResource("contentinspectionpolicy", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("contentinspectionpolicy %s not found", n)
		}

		return nil
	}
}

func testAccCheckContentinspectionpolicyDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_contentinspectionpolicy" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("contentinspectionpolicy", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("contentinspectionpolicy %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccContentinspectionpolicy_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_contentinspectionpolicy.tf_contentinspectionpolicy"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckContentinspectionpolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccContentinspectionpolicy_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckContentinspectionpolicyExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Contentinspectionpolicy.Type(), "my_ci_policy"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccContentinspectionpolicy_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckContentinspectionpolicyExist(resAddr, nil)),
			},
		},
	})
}

func TestAccContentinspectionpolicy_import(t *testing.T) {
	const resAddr = "citrixadc_contentinspectionpolicy.tf_contentinspectionpolicy"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckContentinspectionpolicyDestroy,
		Steps: []resource.TestStep{
			{Config: testAccContentinspectionpolicy_basic},
			{
				Config:                  testAccContentinspectionpolicy_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccContentinspectionpolicy_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckContentinspectionpolicyDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccContentinspectionpolicy_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckContentinspectionpolicyExist("citrixadc_contentinspectionpolicy.tf_contentinspectionpolicy", nil),
				),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccContentinspectionpolicy_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckContentinspectionpolicyExist("citrixadc_contentinspectionpolicy.tf_contentinspectionpolicy", nil),
				),
			},
		},
	})
}

// The contentinspectionpolicy unset test covers the sole cleanly-unsettable
// mutable attribute: comment. undefaction and logaction are spec-unsettable too,
// but NITRO reverts them to the sentinel "Use Global" (echoed back on GET and
// rejected as an input value), so they cannot round-trip through an unset without
// causing a perpetual plan diff and are excluded here.
const testAccContentinspectionpolicy_unset_step1 = `
	resource "citrixadc_contentinspectionpolicy" "tf_unset" {
		name    = "tf_ci_unset"
		rule    = "true"
		action  = "RESET"
		comment = "tf test comment"
	}
`

const testAccContentinspectionpolicy_unset_step2 = `
	resource "citrixadc_contentinspectionpolicy" "tf_unset" {
		name   = "tf_ci_unset"
		rule   = "true"
		action = "RESET"
		# comment removed from config -> the provider must unset it (revert to the
		# NITRO default: no comment).
	}
`

func TestAccContentinspectionpolicy_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckContentinspectionpolicyDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default value is applied and persisted.
				Config: testAccContentinspectionpolicy_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckContentinspectionpolicyExist("citrixadc_contentinspectionpolicy.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_contentinspectionpolicy.tf_unset", "comment", "tf test comment"),
					testAccCheckContentinspectionpolicyADCValue("tf_ci_unset", "comment", "tf test comment"),
				),
			},
			{
				// Removing comment must unset it: state (read back from the appliance)
				// reverts to the NITRO default (absent), and the implicit post-apply
				// plan must be empty.
				Config: testAccContentinspectionpolicy_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckContentinspectionpolicyExist("citrixadc_contentinspectionpolicy.tf_unset", nil),
					resource.TestCheckNoResourceAttr("citrixadc_contentinspectionpolicy.tf_unset", "comment"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckContentinspectionpolicyADCValue("tf_ci_unset", "comment", ""),
				),
			},
		},
	})
}

// testAccCheckContentinspectionpolicyADCValue asserts an attribute's value directly
// on the appliance (not just in Terraform state), proving the unset actually
// reverted it. An absent attribute is treated as the empty string.
func testAccCheckContentinspectionpolicyADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Contentinspectionpolicy.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("contentinspectionpolicy %s not found on appliance", name)
		}
		got := ""
		if v, ok := data[attr]; ok && v != nil {
			got = strings.TrimSpace(fmt.Sprintf("%v", v))
		}
		if got != want {
			return fmt.Errorf("contentinspectionpolicy %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func TestAccContentinspectionpolicyDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccContentinspectionpolicyDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_contentinspectionpolicy.tf_contentinspectionpolicy_ds", "name", "my_ci_policy_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_contentinspectionpolicy.tf_contentinspectionpolicy_ds", "rule", "true"),
					resource.TestCheckResourceAttr("data.citrixadc_contentinspectionpolicy.tf_contentinspectionpolicy_ds", "action", "RESET"),
				),
			},
		},
	})
}
