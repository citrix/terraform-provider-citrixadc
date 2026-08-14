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

const testAccQuicbridgeprofile_add = `
	resource citrixadc_quicbridgeprofile tfAcc_quicbridge {
		name             = "tfAcc_quicbridge"
		routingalgorithm = "PLAINTEXT"
		serveridlength   = 4
	}
`

const testAccQuicbridgeprofile_update = `
	resource citrixadc_quicbridgeprofile tfAcc_quicbridge {
		name             = "tfAcc_quicbridge"
		routingalgorithm = "PLAINTEXT"
		serveridlength   = 6
	}
`

func TestAccQuicbridgeprofile_basic(t *testing.T) {
	if isCpxRun {
		t.Skip("No support in CPX")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckQuicbridgeprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccQuicbridgeprofile_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckQuicbridgeprofileExist("citrixadc_quicbridgeprofile.tfAcc_quicbridge", nil),
					resource.TestCheckResourceAttr("citrixadc_quicbridgeprofile.tfAcc_quicbridge", "routingalgorithm", "PLAINTEXT"),
					resource.TestCheckResourceAttr("citrixadc_quicbridgeprofile.tfAcc_quicbridge", "serveridlength", "4"),
				),
			},
			{
				Config: testAccQuicbridgeprofile_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckQuicbridgeprofileExist("citrixadc_quicbridgeprofile.tfAcc_quicbridge", nil),
					resource.TestCheckResourceAttr("citrixadc_quicbridgeprofile.tfAcc_quicbridge", "routingalgorithm", "PLAINTEXT"),
					resource.TestCheckResourceAttr("citrixadc_quicbridgeprofile.tfAcc_quicbridge", "serveridlength", "6"),
				),
			},
		},
	})
}

func TestAccQuicbridgeprofile_selfHealing(t *testing.T) {
	if isCpxRun {
		t.Skip("No support in CPX")
	}
	const resAddr = "citrixadc_quicbridgeprofile.tfAcc_quicbridge"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckQuicbridgeprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccQuicbridgeprofile_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckQuicbridgeprofileExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Quicbridgeprofile.Type(), "tfAcc_quicbridge"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccQuicbridgeprofile_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckQuicbridgeprofileExist(resAddr, nil)),
			},
		},
	})
}

func TestAccQuicbridgeprofile_import(t *testing.T) {
	if isCpxRun {
		t.Skip("No support in CPX")
	}
	const resAddr = "citrixadc_quicbridgeprofile.tfAcc_quicbridge"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckQuicbridgeprofileDestroy,
		Steps: []resource.TestStep{
			{Config: testAccQuicbridgeprofile_add},
			{
				Config:                  testAccQuicbridgeprofile_add,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccCheckQuicbridgeprofileExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No quicbridgeprofile name is set")
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
		data, err := client.FindResource("quicbridgeprofile", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("quicbridgeprofile %s not found", n)
		}

		return nil
	}
}

func testAccCheckQuicbridgeprofileDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_quicbridgeprofile" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("quicbridgeprofile", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("quicbridgeprofile %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccQuicbridgeprofile_sdkv2StateUpgrade(t *testing.T) {
	if isCpxRun {
		t.Skip("No support in CPX")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckQuicbridgeprofileDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccQuicbridgeprofile_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckQuicbridgeprofileExist("citrixadc_quicbridgeprofile.tfAcc_quicbridge", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccQuicbridgeprofile_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckQuicbridgeprofileExist("citrixadc_quicbridgeprofile.tfAcc_quicbridge", nil)),
			},
		},
	})
}

// quicbridgeprofile has one unset-eligible mutable attribute exercisable here:
// serveridlength (NITRO default 4). routingalgorithm is also unset-wired but its
// only possible value is PLAINTEXT (== its default), so it cannot be driven to a
// non-default value and is not asserted here.
const testAccQuicbridgeprofile_unset_step1 = `
resource "citrixadc_quicbridgeprofile" "tf_unset" {
	name             = "tfAcc_quicbridge_unset"
	routingalgorithm = "PLAINTEXT"
	serveridlength   = 6
}
`

const testAccQuicbridgeprofile_unset_step2 = `
resource "citrixadc_quicbridgeprofile" "tf_unset" {
	name = "tfAcc_quicbridge_unset"
	# unset-eligible attributes removed from config -> provider must unset them
	# (revert to NITRO defaults: serveridlength = 4).
}
`

func TestAccQuicbridgeprofile_unset(t *testing.T) {
	if isCpxRun {
		t.Skip("No support in CPX")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckQuicbridgeprofileDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccQuicbridgeprofile_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckQuicbridgeprofileExist("citrixadc_quicbridgeprofile.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_quicbridgeprofile.tf_unset", "serveridlength", "6"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from the
				// appliance) reverts to the documented NITRO defaults, and the
				// implicit post-apply plan must be empty.
				Config: testAccQuicbridgeprofile_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckQuicbridgeprofileExist("citrixadc_quicbridgeprofile.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_quicbridgeprofile.tf_unset", "serveridlength", "4"),
					// routingalgorithm's only valid value is its default PLAINTEXT, so it
					// is never actually unset; the state check above suffices. The raw
					// appliance drops routingalgorithm from GET when serveridlength alone
					// is unset, so a raw-appliance assertion on it is not meaningful.
					resource.TestCheckResourceAttr("citrixadc_quicbridgeprofile.tf_unset", "routingalgorithm", "PLAINTEXT"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckQuicbridgeprofileADCValue("tfAcc_quicbridge_unset", "serveridlength", "4"),
				),
			},
		},
	})
}

// testAccCheckQuicbridgeprofileADCValue asserts an attribute's value directly on
// the appliance (not just in Terraform state), proving the unset actually
// reverted it.
func testAccCheckQuicbridgeprofileADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Quicbridgeprofile.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("quicbridgeprofile %s not found on appliance", name)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("quicbridgeprofile %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func TestAccQuicbridgeprofileDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccQuicbridgeprofileDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_quicbridgeprofile.test", "name", "tf_quicbridgeprofile"),
					resource.TestCheckResourceAttrSet("data.citrixadc_quicbridgeprofile.test", "id"),
					resource.TestCheckResourceAttrSet("data.citrixadc_quicbridgeprofile.test", "routingalgorithm"),
					resource.TestCheckResourceAttrSet("data.citrixadc_quicbridgeprofile.test", "serveridlength"),
				),
			},
		},
	})
}

const testAccQuicbridgeprofileDataSource_basic = `
resource "citrixadc_quicbridgeprofile" "tf_quicbridgeprofile" {
	name              = "tf_quicbridgeprofile"
	routingalgorithm  = "PLAINTEXT"
	serveridlength    = 4
}

data "citrixadc_quicbridgeprofile" "test" {
	name = citrixadc_quicbridgeprofile.tf_quicbridgeprofile.name
}
`
