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

// callhome is a SINGLETON set-get parameter resource:
//   - Create/Update = UpdateUnnamedResource (set), Read = find-unnamed, Delete = no-op.
//   - Fixed ID "callhome"; the object always exists on the ADC and is never deleted.
//   - ipaddress and proxyauthservice are mutually exclusive, so they are never
//     configured together in these tests.

const testAccCallhome_basic_step1 = `
resource "citrixadc_callhome" "tf_callhome" {
  mode             = "Default"
  hbcustominterval = 10
  emailaddress     = "test@example.com"
  proxymode        = "NO"
}

`

func TestAccCallhome_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// Singleton: object always exists on the ADC, so the destroy check only
		// verifies the object is still present (it is never actually deleted).
		CheckDestroy: testAccCheckCallhomeDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCallhome_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCallhomeExist("citrixadc_callhome.tf_callhome", nil),
					resource.TestCheckResourceAttr("citrixadc_callhome.tf_callhome", "mode", "Default"),
					resource.TestCheckResourceAttr("citrixadc_callhome.tf_callhome", "hbcustominterval", "10"),
					resource.TestCheckResourceAttr("citrixadc_callhome.tf_callhome", "emailaddress", "test@example.com"),
					resource.TestCheckResourceAttr("citrixadc_callhome.tf_callhome", "proxymode", "NO"),
				),
			},
		},
	})
}

func TestAccCallhome_import(t *testing.T) {
	const resAddr = "citrixadc_callhome.tf_callhome"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCallhomeDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCallhome_basic_step1,
			},
			{
				Config:                  testAccCallhome_basic_step1,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccCheckCallhomeExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No callhome ID is set")
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
		// Singleton set-get resource: read via find-unnamed (empty name).
		data, err := client.FindResource(service.Callhome.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("callhome %s not found", n)
		}

		return nil
	}
}

// Singleton destroy check: callhome is never deleted on the ADC, so this only
// confirms the object still exists after the config is destroyed (removed from
// Terraform state). It does NOT assert absence.
func testAccCheckCallhomeDestroy(s *terraform.State) error {
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_callhome" {
			continue
		}
		// Singleton resource always exists on the ADC; a successful read is expected.
		_, err := client.FindResource(service.Callhome.Type(), "")
		if err != nil {
			return fmt.Errorf("callhome singleton unexpectedly missing after destroy: %v", err)
		}
	}
	return nil
}

// Unset test: eligible attributes wired for ?action=unset in resource_callhome.go
// are hbcustominterval (default 7), mode (default "Default") and proxymode
// (default "NO"). Step 1 sets the applicable ones to non-defaults; step 2 removes
// them so the provider issues ?action=unset, reverting each to its ADC default with
// an empty post-apply plan.
//
// NOTE on `mode`: its only non-default enum value "CSP" is rejected on a standalone
// appliance with NITRO errorcode 257 "Operation not permitted" (a testbed capability
// limit, not an invalid value), so it cannot be driven to a non-default here. It is
// left at its default "Default" and asserted to remain there (no perpetual diff);
// the true non-default->unset transition for `mode` is only exercisable on an
// appliance where CSP mode is permitted.
const testAccCallhome_unset_step1 = `
resource "citrixadc_callhome" "tf_unset" {
  hbcustominterval = 10
  proxymode        = "YES"
  # mode omitted -> stays at its default "Default" (CSP is "Operation not permitted"
  # on this standalone testbed).
}
`

const testAccCallhome_unset_step2 = `
resource "citrixadc_callhome" "tf_unset" {
  # all eligible attributes removed from config -> provider must unset them
  # (hbcustominterval -> 7, proxymode -> NO); mode remains at its default "Default".
}
`

func TestAccCallhome_unset(t *testing.T) {
	// callhome's other tests (TestAccCallhome_basic/_import/DataSource) have no skip
	// guards; they run on the default standalone testbed, so no guard is added here.
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCallhomeDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values apply and persist.
				Config: testAccCallhome_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCallhomeExist("citrixadc_callhome.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_callhome.tf_unset", "hbcustominterval", "10"),
					resource.TestCheckResourceAttr("citrixadc_callhome.tf_unset", "proxymode", "YES"),
					resource.TestCheckResourceAttr("citrixadc_callhome.tf_unset", "mode", "Default"),
				),
			},
			{
				// Removing them must unset -> state reverts to NITRO defaults, and the
				// implicit post-apply plan must be empty (no perpetual diff).
				Config: testAccCallhome_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCallhomeExist("citrixadc_callhome.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_callhome.tf_unset", "hbcustominterval", "7"),
					resource.TestCheckResourceAttr("citrixadc_callhome.tf_unset", "proxymode", "NO"),
					resource.TestCheckResourceAttr("citrixadc_callhome.tf_unset", "mode", "Default"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckCallhomeADCValue("hbcustominterval", "7"),
					testAccCheckCallhomeADCValue("proxymode", "NO"),
					testAccCheckCallhomeADCValue("mode", "Default"),
				),
			},
		},
	})
}

// testAccCheckCallhomeADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted it.
// callhome is a singleton: read via find-unnamed (empty name).
func testAccCheckCallhomeADCValue(attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Callhome.Type(), "")
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("callhome not found on appliance")
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("callhome: appliance attr %q = %q, want %q (unset did not revert it)", attr, got, want)
		}
		return nil
	}
}

const testAccCallhomeDataSource_basic = `

resource "citrixadc_callhome" "tf_callhome" {
  mode             = "Default"
  hbcustominterval = 10
  emailaddress     = "test@example.com"
  proxymode        = "NO"
}

data "citrixadc_callhome" "tf_callhome" {
  depends_on = [citrixadc_callhome.tf_callhome]
}
`

func TestAccCallhomeDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCallhomeDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_callhome.tf_callhome", "mode", "Default"),
					resource.TestCheckResourceAttr("data.citrixadc_callhome.tf_callhome", "hbcustominterval", "10"),
					resource.TestCheckResourceAttr("data.citrixadc_callhome.tf_callhome", "emailaddress", "test@example.com"),
					resource.TestCheckResourceAttr("data.citrixadc_callhome.tf_callhome", "proxymode", "NO"),
				),
			},
		},
	})
}
