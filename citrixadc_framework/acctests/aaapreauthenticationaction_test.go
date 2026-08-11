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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testAccAaapreauthenticationaction_basic = `

	resource "citrixadc_aaapreauthenticationaction" "tf_aaapreauthenticationaction" {
		name                    = "tf_aaapreauthenticationaction"
		preauthenticationaction = "ALLOW"
		deletefiles             = "/var/tmp/new/hello.txt"
	}
  
`
const testAccAaapreauthenticationaction_update = `

	resource "citrixadc_aaapreauthenticationaction" "tf_aaapreauthenticationaction" {
		name                    = "tf_aaapreauthenticationaction"
		preauthenticationaction = "DENY"
		deletefiles             = "/var/tmp/new/hello.txt"
	}
`

func TestAccAaapreauthenticationaction_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAaapreauthenticationactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAaapreauthenticationaction_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaapreauthenticationactionExist("citrixadc_aaapreauthenticationaction.tf_aaapreauthenticationaction", nil),
					resource.TestCheckResourceAttr("citrixadc_aaapreauthenticationaction.tf_aaapreauthenticationaction", "preauthenticationaction", "ALLOW"),
				),
			},
			{
				Config: testAccAaapreauthenticationaction_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaapreauthenticationactionExist("citrixadc_aaapreauthenticationaction.tf_aaapreauthenticationaction", nil),
					resource.TestCheckResourceAttr("citrixadc_aaapreauthenticationaction.tf_aaapreauthenticationaction", "preauthenticationaction", "DENY"),
				),
			},
		},
	})
}

// The unset test covers the spec-unsettable, mutable string attributes:
// defaultepagroup, deletefiles, killprocess. step1 sets them to non-default
// values; step2 removes them from config -> the provider must issue the NITRO
// ?action=unset so the appliance reverts them (they become absent -> null).
const testAccAaapreauthenticationaction_unset_step1 = `
	resource "citrixadc_aaapreauthenticationaction" "tf_unset" {
		name                    = "tf_test_aaapreauthaction_unset"
		preauthenticationaction = "ALLOW"
		defaultepagroup         = "tf_epa_grp"
		deletefiles             = "/var/tmp/unset/hello.txt"
		killprocess             = "badproc.exe"
	}
`

const testAccAaapreauthenticationaction_unset_step2 = `
	resource "citrixadc_aaapreauthenticationaction" "tf_unset" {
		name                    = "tf_test_aaapreauthaction_unset"
		preauthenticationaction = "ALLOW"
		# defaultepagroup, deletefiles, killprocess removed from config -> unset.
	}
`

func TestAccAaapreauthenticationaction_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAaapreauthenticationactionDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccAaapreauthenticationaction_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaapreauthenticationactionExist("citrixadc_aaapreauthenticationaction.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_aaapreauthenticationaction.tf_unset", "defaultepagroup", "tf_epa_grp"),
					resource.TestCheckResourceAttr("citrixadc_aaapreauthenticationaction.tf_unset", "deletefiles", "/var/tmp/unset/hello.txt"),
					resource.TestCheckResourceAttr("citrixadc_aaapreauthenticationaction.tf_unset", "killprocess", "badproc.exe"),
				),
			},
			{
				// Removing the attributes must unset them: the appliance reverts
				// them (absent), and the implicit post-apply plan must be empty.
				Config: testAccAaapreauthenticationaction_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaapreauthenticationactionExist("citrixadc_aaapreauthenticationaction.tf_unset", nil),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckAaapreauthenticationactionADCUnset("tf_test_aaapreauthaction_unset", "defaultepagroup"),
					testAccCheckAaapreauthenticationactionADCUnset("tf_test_aaapreauthaction_unset", "deletefiles"),
					testAccCheckAaapreauthenticationactionADCUnset("tf_test_aaapreauthaction_unset", "killprocess"),
				),
			},
		},
	})
}

// testAccCheckAaapreauthenticationactionADCUnset asserts an attribute is absent
// (or empty) directly on the appliance, proving the unset actually reverted it.
func testAccCheckAaapreauthenticationactionADCUnset(name, attr string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Aaapreauthenticationaction.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("aaapreauthenticationaction %s not found on appliance", name)
		}
		if v, ok := data[attr]; ok {
			got := strings.TrimSpace(fmt.Sprintf("%v", v))
			if got != "" {
				return fmt.Errorf("aaapreauthenticationaction %s: appliance attr %q = %q, want absent/empty (unset did not revert it)", name, attr, got)
			}
		}
		return nil
	}
}

func testAccCheckAaapreauthenticationactionExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No aaapreauthenticationaction name is set")
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
		data, err := client.FindResource(service.Aaapreauthenticationaction.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("aaapreauthenticationaction %s not found", n)
		}

		return nil
	}
}

func testAccCheckAaapreauthenticationactionDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_aaapreauthenticationaction" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Aaapreauthenticationaction.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("aaapreauthenticationaction %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccAaapreauthenticationaction_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_aaapreauthenticationaction.tf_aaapreauthenticationaction"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAaapreauthenticationactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAaapreauthenticationaction_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAaapreauthenticationactionExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Aaapreauthenticationaction.Type(), "tf_aaapreauthenticationaction"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccAaapreauthenticationaction_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAaapreauthenticationactionExist(resAddr, nil)),
			},
		},
	})
}

func TestAccAaapreauthenticationaction_import(t *testing.T) {
	const resAddr = "citrixadc_aaapreauthenticationaction.tf_aaapreauthenticationaction"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAaapreauthenticationactionDestroy,
		Steps: []resource.TestStep{
			{Config: testAccAaapreauthenticationaction_basic},
			{
				Config:                  testAccAaapreauthenticationaction_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccAaapreauthenticationaction_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckAaapreauthenticationactionDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccAaapreauthenticationaction_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAaapreauthenticationactionExist("citrixadc_aaapreauthenticationaction.tf_aaapreauthenticationaction", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccAaapreauthenticationaction_basic,
				Check:                    resource.ComposeTestCheckFunc(testAccCheckAaapreauthenticationactionExist("citrixadc_aaapreauthenticationaction.tf_aaapreauthenticationaction", nil)),
			},
		},
	})
}

const testAccAaapreauthenticationactionDataSource_basic = `

	resource "citrixadc_aaapreauthenticationaction" "tf_aaapreauthenticationaction" {
		name                    = "tf_aaapreauthenticationaction"
		preauthenticationaction = "ALLOW"
		deletefiles             = "/var/tmp/new/hello.txt"
	}
	
	data "citrixadc_aaapreauthenticationaction" "tf_aaapreauthenticationaction" {
		name = citrixadc_aaapreauthenticationaction.tf_aaapreauthenticationaction.name
	}
`

func TestAccAaapreauthenticationactionDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAaapreauthenticationactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAaapreauthenticationactionDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_aaapreauthenticationaction.tf_aaapreauthenticationaction", "name", "tf_aaapreauthenticationaction"),
					resource.TestCheckResourceAttr("data.citrixadc_aaapreauthenticationaction.tf_aaapreauthenticationaction", "preauthenticationaction", "ALLOW"),
					resource.TestCheckResourceAttr("data.citrixadc_aaapreauthenticationaction.tf_aaapreauthenticationaction", "deletefiles", "/var/tmp/new/hello.txt"),
				),
			},
		},
	})
}
