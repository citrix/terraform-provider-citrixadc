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

const testAccVideooptimizationpacingaction_add = `

	resource "citrixadc_videooptimizationpacingaction" "tf_pacingaction" {
		name 	= "tf_pacingaction"
		rate 	= 20
		comment = "Some Comment"
	}
`

const testAccVideooptimizationpacingaction_update = `

	resource "citrixadc_videooptimizationpacingaction" "tf_pacingaction" {
		name 	= "tf_pacingaction"
		rate 	= 10
		comment = "Some Comment"
	}
`

func TestAccVideooptimizationpacingaction_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVideooptimizationpacingactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVideooptimizationpacingaction_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVideooptimizationpacingactionExist("citrixadc_videooptimizationpacingaction.tf_pacingaction", nil),
					resource.TestCheckResourceAttr("citrixadc_videooptimizationpacingaction.tf_pacingaction", "name", "tf_pacingaction"),
					resource.TestCheckResourceAttr("citrixadc_videooptimizationpacingaction.tf_pacingaction", "rate", "20"),
				),
			},
			{
				Config: testAccVideooptimizationpacingaction_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVideooptimizationpacingactionExist("citrixadc_videooptimizationpacingaction.tf_pacingaction", nil),
					resource.TestCheckResourceAttr("citrixadc_videooptimizationpacingaction.tf_pacingaction", "name", "tf_pacingaction"),
					resource.TestCheckResourceAttr("citrixadc_videooptimizationpacingaction.tf_pacingaction", "rate", "10"),
				),
			},
		},
	})
}

func testAccCheckVideooptimizationpacingactionExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No videooptimizationpacingaction name is set")
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
		data, err := client.FindResource("videooptimizationpacingaction", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("videooptimizationpacingaction %s not found", n)
		}

		return nil
	}
}

func testAccCheckVideooptimizationpacingactionDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_videooptimizationpacingaction" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("videooptimizationpacingaction", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("videooptimizationpacingaction %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccVideooptimizationpacingaction_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_videooptimizationpacingaction.tf_pacingaction"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVideooptimizationpacingactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVideooptimizationpacingaction_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckVideooptimizationpacingactionExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Videooptimizationpacingaction.Type(), "tf_pacingaction"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccVideooptimizationpacingaction_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckVideooptimizationpacingactionExist(resAddr, nil)),
			},
		},
	})
}

const testAccVideooptimizationpacingactionDataSource_basic = `

	resource "citrixadc_videooptimizationpacingaction" "tf_pacingaction" {
		name 	= "tf_pacingaction"
		rate 	= 20
		comment = "Some Comment"
	}

	data "citrixadc_videooptimizationpacingaction" "tf_pacingaction" {
		name = citrixadc_videooptimizationpacingaction.tf_pacingaction.name
	}
`

func TestAccVideooptimizationpacingaction_import(t *testing.T) {
	const resAddr = "citrixadc_videooptimizationpacingaction.tf_pacingaction"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVideooptimizationpacingactionDestroy,
		Steps: []resource.TestStep{
			{Config: testAccVideooptimizationpacingaction_add},
			{
				Config:                  testAccVideooptimizationpacingaction_add,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccVideooptimizationpacingaction_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckVideooptimizationpacingactionDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.0.0"},
				},
				Config: testAccVideooptimizationpacingaction_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckVideooptimizationpacingactionExist("citrixadc_videooptimizationpacingaction.tf_pacingaction", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccVideooptimizationpacingaction_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckVideooptimizationpacingactionExist("citrixadc_videooptimizationpacingaction.tf_pacingaction", nil)),
			},
		},
	})
}

// The only unset-eligible attribute on videooptimizationpacingaction is comment
// (rate is Required, so always supplied; name/newname are the key/rename fields).
// Removing comment from config must unset it -> the appliance reverts it to empty.
const testAccVideooptimizationpacingaction_unset_step1 = `
	resource "citrixadc_videooptimizationpacingaction" "tf_unset" {
		name    = "tf_test_pacingaction_unset"
		rate    = 20
		comment = "Some Comment"
	}
`

const testAccVideooptimizationpacingaction_unset_step2 = `
	resource "citrixadc_videooptimizationpacingaction" "tf_unset" {
		name = "tf_test_pacingaction_unset"
		rate = 20
		# comment removed from config -> the provider must unset it (revert to empty).
	}
`

func TestAccVideooptimizationpacingaction_unset(t *testing.T) {
	const resAddr = "citrixadc_videooptimizationpacingaction.tf_unset"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVideooptimizationpacingactionDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default value applied and persisted.
				Config: testAccVideooptimizationpacingaction_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVideooptimizationpacingactionExist(resAddr, nil),
					resource.TestCheckResourceAttr(resAddr, "comment", "Some Comment"),
					testAccCheckVideooptimizationpacingactionADCValue("tf_test_pacingaction_unset", "comment", "Some Comment"),
				),
			},
			{
				// Removing comment must unset it: state (read back from the appliance)
				// reverts to empty and the implicit post-apply plan must be empty.
				Config: testAccVideooptimizationpacingaction_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVideooptimizationpacingactionExist(resAddr, nil),
					resource.TestCheckResourceAttr(resAddr, "comment", ""),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckVideooptimizationpacingactionADCValue("tf_test_pacingaction_unset", "comment", ""),
				),
			},
		},
	})
}

// testAccCheckVideooptimizationpacingactionADCValue asserts an attribute's value
// directly on the appliance (not just in Terraform state), proving the unset
// actually reverted it. A missing/nil attribute is treated as "".
func testAccCheckVideooptimizationpacingactionADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Videooptimizationpacingaction.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("videooptimizationpacingaction %s not found on appliance", name)
		}
		got := ""
		if v, ok := data[attr]; ok && v != nil {
			got = strings.TrimSpace(fmt.Sprintf("%v", v))
		}
		if got != want {
			return fmt.Errorf("videooptimizationpacingaction %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func TestAccVideooptimizationpacingactionDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccVideooptimizationpacingactionDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_videooptimizationpacingaction.tf_pacingaction", "name", "tf_pacingaction"),
					resource.TestCheckResourceAttr("data.citrixadc_videooptimizationpacingaction.tf_pacingaction", "rate", "20"),
					resource.TestCheckResourceAttr("data.citrixadc_videooptimizationpacingaction.tf_pacingaction", "comment", "Some Comment"),
				),
			},
		},
	})
}
