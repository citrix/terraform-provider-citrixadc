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

const testAccNsvariable_add = `
	resource "citrixadc_nsvariable" "tf_nsvariable" {
		name          = "tf_nsvariable"
		type          = "text(20)"
		scope         = "global"
		iffull        = "lru"
		ifvaluetoobig = "undef"
		ifnovalue     = "undef"
		comment       = "Testing"
	}
`
const testAccNsvariable_update = `
	resource "citrixadc_nsvariable" "tf_nsvariable" {
		name          = "tf_nsvariable"
		type          = "text(20)"
		scope         = "global"
		iffull        = "lru"
		ifvaluetoobig = "truncate"
		ifnovalue     = "init"
		comment       = "Testing"
	}
`

func TestAccNsvariable_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNsvariableDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNsvariable_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsvariableExist("citrixadc_nsvariable.tf_nsvariable", nil),
					resource.TestCheckResourceAttr("citrixadc_nsvariable.tf_nsvariable", "name", "tf_nsvariable"),
					resource.TestCheckResourceAttr("citrixadc_nsvariable.tf_nsvariable", "iffull", "lru"),
					resource.TestCheckResourceAttr("citrixadc_nsvariable.tf_nsvariable", "ifvaluetoobig", "undef"),
					resource.TestCheckResourceAttr("citrixadc_nsvariable.tf_nsvariable", "ifnovalue", "undef"),
				),
			},
			{
				Config: testAccNsvariable_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsvariableExist("citrixadc_nsvariable.tf_nsvariable", nil),
					resource.TestCheckResourceAttr("citrixadc_nsvariable.tf_nsvariable", "name", "tf_nsvariable"),
					resource.TestCheckResourceAttr("citrixadc_nsvariable.tf_nsvariable", "iffull", "lru"),
					resource.TestCheckResourceAttr("citrixadc_nsvariable.tf_nsvariable", "ifvaluetoobig", "truncate"),
					resource.TestCheckResourceAttr("citrixadc_nsvariable.tf_nsvariable", "ifnovalue", "init"),
				),
			},
		},
	})
}

func TestAccNsvariable_import(t *testing.T) {
	const resAddr = "citrixadc_nsvariable.tf_nsvariable"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNsvariableDestroy,
		Steps: []resource.TestStep{
			{Config: testAccNsvariable_add},
			{
				Config:                  testAccNsvariable_add,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccNsvariable_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_nsvariable.tf_nsvariable"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNsvariableDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNsvariable_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckNsvariableExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Nsvariable.Type(), "tf_nsvariable"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccNsvariable_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckNsvariableExist(resAddr, nil)),
			},
		},
	})
}

func TestAccNsvariable_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckNsvariableDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccNsvariable_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckNsvariableExist("citrixadc_nsvariable.tf_nsvariable", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccNsvariable_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckNsvariableExist("citrixadc_nsvariable.tf_nsvariable", nil)),
			},
		},
	})
}

// testAccNsvariable_unset_step1 sets the unset-eligible attributes to valid
// non-default values on a text singleton variable.
const testAccNsvariable_unset_step1 = `
	resource "citrixadc_nsvariable" "tf_unset" {
		name          = "tf_nsvariable_unset"
		type          = "text(20)"
		iffull        = "undef"
		ifvaluetoobig = "undef"
		ifnovalue     = "undef"
		expires       = 3600
		comment       = "Testing unset"
	}
`

// testAccNsvariable_unset_step2 removes every unset-eligible attribute (keeping
// only name + required type) so the provider must unset them, reverting the
// appliance to the documented NITRO defaults.
const testAccNsvariable_unset_step2 = `
	resource "citrixadc_nsvariable" "tf_unset" {
		name = "tf_nsvariable_unset"
		type = "text(20)"
	}
`

func TestAccNsvariable_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNsvariableDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccNsvariable_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsvariableExist("citrixadc_nsvariable.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_nsvariable.tf_unset", "iffull", "undef"),
					resource.TestCheckResourceAttr("citrixadc_nsvariable.tf_unset", "ifvaluetoobig", "undef"),
					resource.TestCheckResourceAttr("citrixadc_nsvariable.tf_unset", "ifnovalue", "undef"),
					resource.TestCheckResourceAttr("citrixadc_nsvariable.tf_unset", "expires", "3600"),
					resource.TestCheckResourceAttr("citrixadc_nsvariable.tf_unset", "comment", "Testing unset"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from
				// the appliance) reverts to NITRO defaults and the implicit
				// post-apply plan must be empty.
				Config: testAccNsvariable_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsvariableExist("citrixadc_nsvariable.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_nsvariable.tf_unset", "iffull", "lru"),
					resource.TestCheckResourceAttr("citrixadc_nsvariable.tf_unset", "ifvaluetoobig", "truncate"),
					resource.TestCheckResourceAttr("citrixadc_nsvariable.tf_unset", "ifnovalue", "init"),
					resource.TestCheckResourceAttr("citrixadc_nsvariable.tf_unset", "expires", "0"),
					resource.TestCheckResourceAttr("citrixadc_nsvariable.tf_unset", "comment", ""),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckNsvariableADCValue("tf_nsvariable_unset", "iffull", "lru"),
					testAccCheckNsvariableADCValue("tf_nsvariable_unset", "ifvaluetoobig", "truncate"),
					testAccCheckNsvariableADCValue("tf_nsvariable_unset", "ifnovalue", "init"),
					testAccCheckNsvariableADCValue("tf_nsvariable_unset", "expires", "0"),
				),
			},
		},
	})
}

// testAccCheckNsvariableADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted it.
func testAccCheckNsvariableADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Nsvariable.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("nsvariable %s not found on appliance", name)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("nsvariable %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func testAccCheckNsvariableExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No nsvariable name is set")
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
		data, err := client.FindResource(service.Nsvariable.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("nsvariable %s not found", n)
		}

		return nil
	}
}

func testAccCheckNsvariableDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_nsvariable" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Nsvariable.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("nsvariable %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
func TestAccNsvariableDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNsvariableDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_nsvariable.test", "name", "tf_nsvariable_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_nsvariable.test", "type", "text(20)"),
					resource.TestCheckResourceAttr("data.citrixadc_nsvariable.test", "scope", "global"),
					resource.TestCheckResourceAttr("data.citrixadc_nsvariable.test", "comment", "Testing datasource"),
				),
			},
		},
	})
}

const testAccNsvariableDataSource_basic = `
resource "citrixadc_nsvariable" "tf_nsvariable_ds" {
	name          = "tf_nsvariable_ds"
	type          = "text(20)"
	scope         = "global"
	comment       = "Testing datasource"
}

data "citrixadc_nsvariable" "test" {
	name = citrixadc_nsvariable.tf_nsvariable_ds.name
}
`
