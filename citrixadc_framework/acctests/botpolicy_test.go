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

const testAccBotpolicy_add = `
	resource citrixadc_botpolicy tfAcc_botpolicy1 {
		name = "tfAcc_botpolicy1"  
		profilename = "BOT_BYPASS"
		rule  = "true"
		comment = "COMMENT FOR BOTPOLICY"
	}
`
const testAccBotpolicy_update = `
	resource citrixadc_botpolicy tfAcc_botpolicy1 {
		name = "tfAcc_botpolicy1"
		profilename = "BOT_BYPASS"
		rule = "true"
		comment = "CHANGED COMMENT"
	}
`

func TestAccBotpolicy_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckBotpolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccBotpolicy_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBotpolicyExist("citrixadc_botpolicy.tfAcc_botpolicy1", nil),
					resource.TestCheckResourceAttr("citrixadc_botpolicy.tfAcc_botpolicy1", "name", "tfAcc_botpolicy1"),
					resource.TestCheckResourceAttr("citrixadc_botpolicy.tfAcc_botpolicy1", "profilename", "BOT_BYPASS"),
					resource.TestCheckResourceAttr("citrixadc_botpolicy.tfAcc_botpolicy1", "rule", "true"),
					resource.TestCheckResourceAttr("citrixadc_botpolicy.tfAcc_botpolicy1", "comment", "COMMENT FOR BOTPOLICY"),
				),
			},
			{
				Config: testAccBotpolicy_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBotpolicyExist("citrixadc_botpolicy.tfAcc_botpolicy1", nil),
					resource.TestCheckResourceAttr("citrixadc_botpolicy.tfAcc_botpolicy1", "name", "tfAcc_botpolicy1"),
					resource.TestCheckResourceAttr("citrixadc_botpolicy.tfAcc_botpolicy1", "profilename", "BOT_BYPASS"),
					resource.TestCheckResourceAttr("citrixadc_botpolicy.tfAcc_botpolicy1", "rule", "true"),
					resource.TestCheckResourceAttr("citrixadc_botpolicy.tfAcc_botpolicy1", "comment", "CHANGED COMMENT"),
				),
			},
		},
	})
}

func TestAccBotpolicy_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_botpolicy.tfAcc_botpolicy1"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckBotpolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccBotpolicy_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckBotpolicyExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Botpolicy.Type(), "tfAcc_botpolicy1"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccBotpolicy_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckBotpolicyExist(resAddr, nil)),
			},
		},
	})
}

func TestAccBotpolicy_import(t *testing.T) {
	const resAddr = "citrixadc_botpolicy.tfAcc_botpolicy1"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckBotpolicyDestroy,
		Steps: []resource.TestStep{
			{Config: testAccBotpolicy_add},
			{
				Config:                  testAccBotpolicy_add,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccBotpolicy_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckBotpolicyDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.0.0"},
				},
				Config: testAccBotpolicy_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckBotpolicyExist("citrixadc_botpolicy.tfAcc_botpolicy1", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccBotpolicy_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckBotpolicyExist("citrixadc_botpolicy.tfAcc_botpolicy1", nil)),
			},
		},
	})
}

func testAccCheckBotpolicyExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No botpolicy name is set")
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
		data, err := client.FindResource("botpolicy", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("BOT policy %s not found", n)
		}

		return nil
	}
}

func testAccCheckBotpolicyDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_botpolicy" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("botpolicy", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("BOT policy %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

// The botpolicy unset test covers the mutable, spec-unsettable attributes
// comment and undefaction. logaction is also wired for unset but requires a
// pre-existing bot messagelog action to set to a non-default value, so it is not
// exercised here. After unset, undefaction reverts to the NITRO default "None"
// and comment reverts to absent (no value).
const testAccBotpolicy_unset_step1 = `
	resource citrixadc_botpolicy tfAcc_botpolicy_unset {
		name        = "tfAcc_botpolicy_unset"
		profilename = "BOT_BYPASS"
		rule        = "true"
		comment     = "UNSET PROBE COMMENT"
		undefaction = "RESET"
	}
`

const testAccBotpolicy_unset_step2 = `
	resource citrixadc_botpolicy tfAcc_botpolicy_unset {
		name        = "tfAcc_botpolicy_unset"
		profilename = "BOT_BYPASS"
		rule        = "true"
		# comment and undefaction removed from config -> provider must unset them
		# (revert to NITRO defaults).
	}
`

func TestAccBotpolicy_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckBotpolicyDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccBotpolicy_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBotpolicyExist("citrixadc_botpolicy.tfAcc_botpolicy_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_botpolicy.tfAcc_botpolicy_unset", "comment", "UNSET PROBE COMMENT"),
					resource.TestCheckResourceAttr("citrixadc_botpolicy.tfAcc_botpolicy_unset", "undefaction", "RESET"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from the
				// appliance) reverts to NITRO defaults and the implicit post-apply
				// plan must be empty.
				Config: testAccBotpolicy_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBotpolicyExist("citrixadc_botpolicy.tfAcc_botpolicy_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_botpolicy.tfAcc_botpolicy_unset", "undefaction", "None"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckBotpolicyADCValue("tfAcc_botpolicy_unset", "undefaction", "None"),
					testAccCheckBotpolicyADCValue("tfAcc_botpolicy_unset", "comment", ""),
				),
			},
		},
	})
}

// testAccCheckBotpolicyADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted
// it. An empty want matches an absent attribute.
func testAccCheckBotpolicyADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Botpolicy.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("botpolicy %s not found on appliance", name)
		}
		var got string
		if v, ok := data[attr]; ok && v != nil {
			got = strings.TrimSpace(fmt.Sprintf("%v", v))
		}
		if got != want {
			return fmt.Errorf("botpolicy %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func TestAccBotpolicyDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccBotpolicyDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.citrixadc_botpolicy.tf_botpolicy_ds", "id"),
					resource.TestCheckResourceAttr("data.citrixadc_botpolicy.tf_botpolicy_ds", "name", "tf_botpolicy_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_botpolicy.tf_botpolicy_ds", "profilename", "BOT_BYPASS"),
					resource.TestCheckResourceAttr("data.citrixadc_botpolicy.tf_botpolicy_ds", "rule", "true"),
					resource.TestCheckResourceAttr("data.citrixadc_botpolicy.tf_botpolicy_ds", "comment", "DATASOURCE TEST COMMENT"),
				),
			},
		},
	})
}

const testAccBotpolicyDataSource_basic = `

resource "citrixadc_botpolicy" "tf_botpolicy_ds" {
    name        = "tf_botpolicy_ds"
    profilename = "BOT_BYPASS"
    rule        = "true"
    comment     = "DATASOURCE TEST COMMENT"
}

data "citrixadc_botpolicy" "tf_botpolicy_ds" {
    name = citrixadc_botpolicy.tf_botpolicy_ds.name
    depends_on = [citrixadc_botpolicy.tf_botpolicy_ds]
}

`
