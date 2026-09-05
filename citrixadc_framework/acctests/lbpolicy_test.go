/*
Copyright 2024 Citrix Systems, Inc

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

const testAccLbpolicy_basic = `

	resource "citrixadc_lbaction" "tf_act" {
		name  = "tf_act"
		type  = "SELECTIONORDER"
		value = [1]
	}
	
	resource "citrixadc_lbpolicy" "tf_pol" {
		name   = "tf_pol"
		rule   = "true"
		action = citrixadc_lbaction.tf_act.name
	}
  
`

func TestAccLbpolicy_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLbpolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLbpolicy_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbpolicyExist("citrixadc_lbpolicy.tf_pol", nil),
				),
			},
		},
	})
}

func testAccCheckLbpolicyExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lbpolicy name is set")
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
		data, err := client.FindResource("lbpolicy", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("lbpolicy %s not found", n)
		}

		return nil
	}
}

func testAccCheckLbpolicyDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_lbpolicy" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("lbpolicy", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("lbpolicy %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccLbpolicy_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_lbpolicy.tf_pol"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLbpolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLbpolicy_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckLbpolicyExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Lbpolicy.Type(), "tf_pol"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccLbpolicy_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckLbpolicyExist(resAddr, nil)),
			},
		},
	})
}

func TestAccLbpolicy_import(t *testing.T) {
	const resAddr = "citrixadc_lbpolicy.tf_pol"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLbpolicyDestroy,
		Steps: []resource.TestStep{
			{Config: testAccLbpolicy_basic},
			{
				Config:                  testAccLbpolicy_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

// The lbpolicy unset test covers the single cleanly-unsettable, config-removable
// attribute: comment. NITRO's other unset-listed attributes (undefaction,
// logaction) have display-only server defaults ("Use Global" / "None") that are
// NOT valid create/update payload values -- sending them back errors with 2818
// ("Invalid undef action or log action") -- so a schema Default cannot be added
// to force the config-removal plan diff without regressing create. They are
// therefore excluded from unset.
const testAccLbpolicy_unset_step1 = `

	resource "citrixadc_lbaction" "tf_act_unset" {
		name  = "tf_act_unset"
		type  = "SELECTIONORDER"
		value = [1]
	}

	resource "citrixadc_lbpolicy" "tf_pol_unset" {
		name    = "tf_pol_unset"
		rule    = "true"
		action  = citrixadc_lbaction.tf_act_unset.name
		comment = "unset me"
	}
`

const testAccLbpolicy_unset_step2 = `

	resource "citrixadc_lbaction" "tf_act_unset" {
		name  = "tf_act_unset"
		type  = "SELECTIONORDER"
		value = [1]
	}

	resource "citrixadc_lbpolicy" "tf_pol_unset" {
		name   = "tf_pol_unset"
		rule   = "true"
		action = citrixadc_lbaction.tf_act_unset.name
		# comment removed from config -> the provider must unset it (revert to the
		# NITRO default of an empty comment).
	}
`

func TestAccLbpolicy_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLbpolicyDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default value applied and persisted.
				Config: testAccLbpolicy_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbpolicyExist("citrixadc_lbpolicy.tf_pol_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_lbpolicy.tf_pol_unset", "comment", "unset me"),
				),
			},
			{
				// Removing comment must unset it. Under Option B there is no static
				// Default, so after the NITRO unset the appliance omits comment from GET
				// and it reads back as null/absent in state (not ""). The implicit
				// post-apply plan must be empty.
				Config: testAccLbpolicy_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbpolicyExist("citrixadc_lbpolicy.tf_pol_unset", nil),
					resource.TestCheckNoResourceAttr("citrixadc_lbpolicy.tf_pol_unset", "comment"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckLbpolicyADCValue("tf_pol_unset", "comment", ""),
				),
			},
		},
	})
}

// testAccCheckLbpolicyADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted it.
func testAccCheckLbpolicyADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Lbpolicy.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("lbpolicy %s not found on appliance", name)
		}
		got := ""
		if v, ok := data[attr]; ok && v != nil {
			got = strings.TrimSpace(fmt.Sprintf("%v", v))
		}
		if got != want {
			return fmt.Errorf("lbpolicy %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

const testAccLbpolicyDataSource_basic = `

	resource "citrixadc_lbaction" "tf_act_ds" {
		name  = "tf_act_ds"
		type  = "SELECTIONORDER"
		value = [1]
	}
	
	resource "citrixadc_lbpolicy" "tf_pol_ds" {
		name   = "tf_pol_ds"
		rule   = "true"
		action = citrixadc_lbaction.tf_act_ds.name
	}
  
	data "citrixadc_lbpolicy" "tf_pol_ds" {
		name       = citrixadc_lbpolicy.tf_pol_ds.name
		depends_on = [citrixadc_lbpolicy.tf_pol_ds]
	}
`

func TestAccLbpolicy_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckLbpolicyDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.0.0"},
				},
				Config: testAccLbpolicy_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckLbpolicyExist("citrixadc_lbpolicy.tf_pol", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccLbpolicy_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckLbpolicyExist("citrixadc_lbpolicy.tf_pol", nil)),
			},
		},
	})
}

func TestAccLbpolicyDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccLbpolicyDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					// id is the universal runtime-binding proof of a resolved data source.
					resource.TestCheckResourceAttrSet("data.citrixadc_lbpolicy.tf_pol_ds", "id"),
					resource.TestCheckResourceAttr("data.citrixadc_lbpolicy.tf_pol_ds", "name", "tf_pol_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_lbpolicy.tf_pol_ds", "rule", "true"),
					resource.TestCheckResourceAttr("data.citrixadc_lbpolicy.tf_pol_ds", "action", "tf_act_ds"),
				),
			},
		},
	})
}
