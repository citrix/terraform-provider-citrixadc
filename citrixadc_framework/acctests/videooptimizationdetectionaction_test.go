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
	"testing"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const testAccVideooptimizationdetectionaction_add = `

	resource "citrixadc_videooptimizationdetectionaction" "tf_detectionaction" {
		name = "tf_videooptimizationdetectionaction"
		type = "clear_text_pd"
	}
  
`
const testAccVideooptimizationdetectionaction_update = `

	resource "citrixadc_videooptimizationdetectionaction" "tf_detectionaction" {
		name = "tf_videooptimizationdetectionaction"
		type = "clear_text_abr"
	}
  
`

const testAccVideooptimizationdetectionactionDataSource_basic = `

	resource "citrixadc_videooptimizationdetectionaction" "tf_detectionaction" {
		name = "tf_videooptimizationdetectionaction"
		type = "clear_text_pd"
	}

	data "citrixadc_videooptimizationdetectionaction" "tf_detectionaction" {
		name = citrixadc_videooptimizationdetectionaction.tf_detectionaction.name
	}
  
`

func TestAccVideooptimizationdetectionaction_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVideooptimizationdetectionactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVideooptimizationdetectionaction_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVideooptimizationdetectionactionExist("citrixadc_videooptimizationdetectionaction.tf_detectionaction", nil),
					resource.TestCheckResourceAttr("citrixadc_videooptimizationdetectionaction.tf_detectionaction", "name", "tf_videooptimizationdetectionaction"),
					resource.TestCheckResourceAttr("citrixadc_videooptimizationdetectionaction.tf_detectionaction", "type", "clear_text_pd"),
				),
			},
			{
				Config: testAccVideooptimizationdetectionaction_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVideooptimizationdetectionactionExist("citrixadc_videooptimizationdetectionaction.tf_detectionaction", nil),
					resource.TestCheckResourceAttr("citrixadc_videooptimizationdetectionaction.tf_detectionaction", "name", "tf_videooptimizationdetectionaction"),
					resource.TestCheckResourceAttr("citrixadc_videooptimizationdetectionaction.tf_detectionaction", "type", "clear_text_abr"),
				),
			},
		},
	})
}

func TestAccVideooptimizationdetectionaction_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_videooptimizationdetectionaction.tf_detectionaction"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVideooptimizationdetectionactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVideooptimizationdetectionaction_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckVideooptimizationdetectionactionExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Videooptimizationdetectionaction.Type(), "tf_videooptimizationdetectionaction"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccVideooptimizationdetectionaction_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckVideooptimizationdetectionactionExist(resAddr, nil)),
			},
		},
	})
}

func TestAccVideooptimizationdetectionaction_import(t *testing.T) {
	const resAddr = "citrixadc_videooptimizationdetectionaction.tf_detectionaction"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVideooptimizationdetectionactionDestroy,
		Steps: []resource.TestStep{
			{Config: testAccVideooptimizationdetectionaction_add},
			{
				Config:                  testAccVideooptimizationdetectionaction_add,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccCheckVideooptimizationdetectionactionExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No videooptimizationdetectionaction name is set")
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
		data, err := client.FindResource("videooptimizationdetectionaction", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("videooptimizationdetectionaction %s not found", n)
		}

		return nil
	}
}

func testAccCheckVideooptimizationdetectionactionDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_videooptimizationdetectionaction" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("videooptimizationdetectionaction", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("videooptimizationdetectionaction %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccVideooptimizationdetectionaction_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckVideooptimizationdetectionactionDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccVideooptimizationdetectionaction_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckVideooptimizationdetectionactionExist("citrixadc_videooptimizationdetectionaction.tf_detectionaction", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccVideooptimizationdetectionaction_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckVideooptimizationdetectionactionExist("citrixadc_videooptimizationdetectionaction.tf_detectionaction", nil)),
			},
		},
	})
}

// comment is the only unset-eligible attribute (Optional, mutable, has a NITRO
// server default of empty string). Step 1 sets it to a non-default value; step 2
// removes it from config so the provider must unset it (revert to empty).
const testAccVideooptimizationdetectionaction_unset_step1 = `
	resource "citrixadc_videooptimizationdetectionaction" "tf_unset" {
		name    = "tf_test_vodetectionaction_unset"
		type    = "clear_text_pd"
		comment = "managed by terraform unset test"
	}
`

const testAccVideooptimizationdetectionaction_unset_step2 = `
	resource "citrixadc_videooptimizationdetectionaction" "tf_unset" {
		name = "tf_test_vodetectionaction_unset"
		type = "clear_text_pd"
		# comment removed from config -> provider must unset it (revert to default "").
	}
`

func TestAccVideooptimizationdetectionaction_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVideooptimizationdetectionactionDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default value is applied and persisted.
				Config: testAccVideooptimizationdetectionaction_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVideooptimizationdetectionactionExist("citrixadc_videooptimizationdetectionaction.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_videooptimizationdetectionaction.tf_unset", "comment", "managed by terraform unset test"),
				),
			},
			{
				// Removing comment must unset it. With Option B (no schema Default),
				// the attr reads back NULL after unset (NITRO omits it from GET), the
				// implicit post-apply plan must be empty, and the appliance value is
				// confirmed independently to be empty.
				Config: testAccVideooptimizationdetectionaction_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVideooptimizationdetectionactionExist("citrixadc_videooptimizationdetectionaction.tf_unset", nil),
					resource.TestCheckNoResourceAttr("citrixadc_videooptimizationdetectionaction.tf_unset", "comment"),
					testAccCheckVideooptimizationdetectionactionADCValue("tf_test_vodetectionaction_unset", "comment", ""),
				),
			},
		},
	})
}

// testAccCheckVideooptimizationdetectionactionADCValue asserts an attribute's
// value directly on the appliance (not just in Terraform state), proving the
// unset actually reverted it.
func testAccCheckVideooptimizationdetectionactionADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Videooptimizationdetectionaction.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("videooptimizationdetectionaction %s not found on appliance", name)
		}
		got := ""
		if v, ok := data[attr]; ok && v != nil {
			got = fmt.Sprintf("%v", v)
		}
		if got != want {
			return fmt.Errorf("videooptimizationdetectionaction %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func TestAccVideooptimizationdetectionactionDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccVideooptimizationdetectionactionDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_videooptimizationdetectionaction.tf_detectionaction", "name", "tf_videooptimizationdetectionaction"),
					resource.TestCheckResourceAttr("data.citrixadc_videooptimizationdetectionaction.tf_detectionaction", "type", "clear_text_pd"),
				),
			},
		},
	})
}
