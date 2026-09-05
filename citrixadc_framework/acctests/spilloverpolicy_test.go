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

const testAccSpilloverpolicy_add = `
	resource "citrixadc_spilloveraction" "tf_spilloveraction" {
		name   = "my_spilloveraction"
		action = "SPILLOVER"
	}
	resource "citrixadc_spilloverpolicy" "tf_spilloverpolicy" {
		name    = "tf_spilloverpolicy"
		rule    = "true"
		action  = citrixadc_spilloveraction.tf_spilloveraction.name
		comment = "This is example of spilloverpolicy"
	}
`
const testAccSpilloverpolicy_update = `
	resource "citrixadc_spilloveraction" "tf_spilloveraction" {
		name   = "my_spilloveraction"
		action = "SPILLOVER"
	}
	resource "citrixadc_spilloverpolicy" "tf_spilloverpolicy" {
		name    = "tf_spilloverpolicy"
		rule    = "false"
		action  = citrixadc_spilloveraction.tf_spilloveraction.name
		comment = "This is example of spilloverpolicy"
	}
`

const testAccSpilloverpolicyDataSource_basic = `
resource "citrixadc_spilloveraction" "tf_spilloveraction" {
	name   = "my_spilloveraction_ds"
	action = "SPILLOVER"
}
resource "citrixadc_spilloverpolicy" "tf_spilloverpolicy" {
	name    = "tf_spilloverpolicy_ds"
	rule    = "true"
	action  = citrixadc_spilloveraction.tf_spilloveraction.name
	comment = "This is example of spilloverpolicy"
}

data "citrixadc_spilloverpolicy" "tf_spilloverpolicy" {
	name = citrixadc_spilloverpolicy.tf_spilloverpolicy.name
}
`

func TestAccSpilloverpolicy_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSpilloverpolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSpilloverpolicy_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSpilloverpolicyExist("citrixadc_spilloverpolicy.tf_spilloverpolicy", nil),
					resource.TestCheckResourceAttr("citrixadc_spilloverpolicy.tf_spilloverpolicy", "name", "tf_spilloverpolicy"),
					resource.TestCheckResourceAttr("citrixadc_spilloverpolicy.tf_spilloverpolicy", "rule", "true"),
				),
			},
			{
				Config: testAccSpilloverpolicy_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSpilloverpolicyExist("citrixadc_spilloverpolicy.tf_spilloverpolicy", nil),
					resource.TestCheckResourceAttr("citrixadc_spilloverpolicy.tf_spilloverpolicy", "name", "tf_spilloverpolicy"),
					resource.TestCheckResourceAttr("citrixadc_spilloverpolicy.tf_spilloverpolicy", "rule", "false"),
				),
			},
		},
	})
}

func TestAccSpilloverpolicyDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSpilloverpolicyDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_spilloverpolicy.tf_spilloverpolicy", "name", "tf_spilloverpolicy_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_spilloverpolicy.tf_spilloverpolicy", "rule", "true"),
					resource.TestCheckResourceAttr("data.citrixadc_spilloverpolicy.tf_spilloverpolicy", "comment", "This is example of spilloverpolicy"),
					// Universal runtime-binding proof.
					resource.TestCheckResourceAttrSet("data.citrixadc_spilloverpolicy.tf_spilloverpolicy", "id"),
					// Read-only counter metadata exposed only by the data source; hits is
					// always populated (0 for a freshly-created policy).
					resource.TestCheckResourceAttrSet("data.citrixadc_spilloverpolicy.tf_spilloverpolicy", "hits"),
				),
			},
		},
	})
}

func TestAccSpilloverpolicy_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_spilloverpolicy.tf_spilloverpolicy"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSpilloverpolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSpilloverpolicy_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckSpilloverpolicyExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Spilloverpolicy.Type(), "tf_spilloverpolicy"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccSpilloverpolicy_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckSpilloverpolicyExist(resAddr, nil)),
			},
		},
	})
}

func TestAccSpilloverpolicy_import(t *testing.T) {
	const resAddr = "citrixadc_spilloverpolicy.tf_spilloverpolicy"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSpilloverpolicyDestroy,
		Steps: []resource.TestStep{
			{Config: testAccSpilloverpolicy_add},
			{
				Config:                  testAccSpilloverpolicy_add,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccSpilloverpolicy_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckSpilloverpolicyDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.0.0"},
				},
				Config: testAccSpilloverpolicy_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckSpilloverpolicyExist("citrixadc_spilloverpolicy.tf_spilloverpolicy", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccSpilloverpolicy_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckSpilloverpolicyExist("citrixadc_spilloverpolicy.tf_spilloverpolicy", nil)),
			},
		},
	})
}

// spilloverpolicy has exactly one spec-unsettable, mutable attribute: comment
// (name/rule/action are mandatory; newname is rename-only). Step 1 sets comment
// to a non-default value; step 2 removes it so the provider must unset it,
// reverting the appliance to the NITRO default (empty).
const testAccSpilloverpolicy_unset_step1 = `
	resource "citrixadc_spilloveraction" "tf_spilloveraction_unset" {
		name   = "tf_spilloveraction_unset"
		action = "SPILLOVER"
	}
	resource "citrixadc_spilloverpolicy" "tf_unset" {
		name    = "tf_spilloverpolicy_unset"
		rule    = "true"
		action  = citrixadc_spilloveraction.tf_spilloveraction_unset.name
		comment = "unset-test-comment"
	}
`

const testAccSpilloverpolicy_unset_step2 = `
	resource "citrixadc_spilloveraction" "tf_spilloveraction_unset" {
		name   = "tf_spilloveraction_unset"
		action = "SPILLOVER"
	}
	resource "citrixadc_spilloverpolicy" "tf_unset" {
		name   = "tf_spilloverpolicy_unset"
		rule   = "true"
		action = citrixadc_spilloveraction.tf_spilloveraction_unset.name
		# comment removed from config -> the provider must unset it (revert to "").
	}
`

func TestAccSpilloverpolicy_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSpilloverpolicyDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default value is applied and persisted.
				Config: testAccSpilloverpolicy_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSpilloverpolicyExist("citrixadc_spilloverpolicy.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_spilloverpolicy.tf_unset", "comment", "unset-test-comment"),
				),
			},
			{
				// Removing comment must unset it: state reverts to the NITRO default
				// (empty) and the implicit post-apply plan must be empty.
				Config: testAccSpilloverpolicy_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSpilloverpolicyExist("citrixadc_spilloverpolicy.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_spilloverpolicy.tf_unset", "comment", ""),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckSpilloverpolicyADCValue("tf_spilloverpolicy_unset", "comment", ""),
				),
			},
		},
	})
}

// testAccCheckSpilloverpolicyADCValue asserts an attribute's value directly on
// the appliance (not just in Terraform state), proving the unset actually
// reverted it. A missing/nil attribute is treated as empty.
func testAccCheckSpilloverpolicyADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Spilloverpolicy.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("spilloverpolicy %s not found on appliance", name)
		}
		got := ""
		if v, ok := data[attr]; ok && v != nil {
			got = strings.TrimSpace(fmt.Sprintf("%v", v))
		}
		if got != want {
			return fmt.Errorf("spilloverpolicy %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func testAccCheckSpilloverpolicyExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No spilloverpolicy name is set")
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
		data, err := client.FindResource(service.Spilloverpolicy.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("spilloverpolicy %s not found", n)
		}

		return nil
	}
}

func testAccCheckSpilloverpolicyDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_spilloverpolicy" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Spilloverpolicy.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("spilloverpolicy %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
