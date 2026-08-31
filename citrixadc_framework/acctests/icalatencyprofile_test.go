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

const testAccIcalatencyprofile_basic = `


resource "citrixadc_icalatencyprofile" "tf_icalatencyprofile" {
	name                     = "my_ica_latencyprofile"
	l7latencymonitoring      = "ENABLED"
	l7latencythresholdfactor = 120
	l7latencywaittime        = 100
	}
  
`
const testAccIcalatencyprofile_update = `


resource "citrixadc_icalatencyprofile" "tf_icalatencyprofile" {
	name                     = "my_ica_latencyprofile"
	l7latencymonitoring      = "DISABLED"
	l7latencythresholdfactor = 100
	l7latencywaittime        = 80
	}
  
`

const testAccIcalatencyprofileDataSource_basic = `

resource "citrixadc_icalatencyprofile" "tf_icalatencyprofile" {
	name                     = "my_ica_latencyprofile"
	l7latencymonitoring      = "ENABLED"
	l7latencythresholdfactor = 120
	l7latencywaittime        = 100
}

data "citrixadc_icalatencyprofile" "tf_icalatencyprofile_ds" {
	name = citrixadc_icalatencyprofile.tf_icalatencyprofile.name
}

`

func TestAccIcalatencyprofile_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckIcalatencyprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccIcalatencyprofile_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIcalatencyprofileExist("citrixadc_icalatencyprofile.tf_icalatencyprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_icalatencyprofile.tf_icalatencyprofile", "name", "my_ica_latencyprofile"),
					resource.TestCheckResourceAttr("citrixadc_icalatencyprofile.tf_icalatencyprofile", "l7latencymonitoring", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_icalatencyprofile.tf_icalatencyprofile", "l7latencythresholdfactor", "120"),
					resource.TestCheckResourceAttr("citrixadc_icalatencyprofile.tf_icalatencyprofile", "l7latencywaittime", "100"),
				),
			},
			{
				Config: testAccIcalatencyprofile_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIcalatencyprofileExist("citrixadc_icalatencyprofile.tf_icalatencyprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_icalatencyprofile.tf_icalatencyprofile", "name", "my_ica_latencyprofile"),
					resource.TestCheckResourceAttr("citrixadc_icalatencyprofile.tf_icalatencyprofile", "l7latencymonitoring", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_icalatencyprofile.tf_icalatencyprofile", "l7latencythresholdfactor", "100"),
					resource.TestCheckResourceAttr("citrixadc_icalatencyprofile.tf_icalatencyprofile", "l7latencywaittime", "80"),
				),
			},
		},
	})
}

func TestAccIcalatencyprofile_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_icalatencyprofile.tf_icalatencyprofile"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckIcalatencyprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccIcalatencyprofile_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckIcalatencyprofileExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Icalatencyprofile.Type(), "my_ica_latencyprofile"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccIcalatencyprofile_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckIcalatencyprofileExist(resAddr, nil)),
			},
		},
	})
}

func TestAccIcalatencyprofile_import(t *testing.T) {
	const resAddr = "citrixadc_icalatencyprofile.tf_icalatencyprofile"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckIcalatencyprofileDestroy,
		Steps: []resource.TestStep{
			{Config: testAccIcalatencyprofile_basic},
			{
				Config:                  testAccIcalatencyprofile_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccCheckIcalatencyprofileExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No icalatencyprofile name is set")
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
		data, err := client.FindResource("icalatencyprofile", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("icalatencyprofile %s not found", n)
		}

		return nil
	}
}

func testAccCheckIcalatencyprofileDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_icalatencyprofile" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("icalatencyprofile", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("icalatencyprofile %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccIcalatencyprofile_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckIcalatencyprofileDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.0.0"},
				},
				Config: testAccIcalatencyprofile_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckIcalatencyprofileExist("citrixadc_icalatencyprofile.tf_icalatencyprofile", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccIcalatencyprofile_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckIcalatencyprofileExist("citrixadc_icalatencyprofile.tf_icalatencyprofile", nil)),
			},
		},
	})
}

// The icalatencyprofile unset test exercises every unset-eligible attribute:
// step1 sets all mutable attributes to valid non-default values, step2 removes
// them so the provider must unset them (revert to the documented NITRO
// defaults).
const testAccIcalatencyprofile_unset_step1 = `
resource "citrixadc_icalatencyprofile" "tf_unset" {
	name                     = "tf_test_icalatencyprofile_unset"
	l7latencymonitoring      = "ENABLED"
	l7latencythresholdfactor = 120
	l7latencywaittime        = 100
	l7latencynotifyinterval  = 50
	l7latencymaxnotifycount  = 30
}
`

const testAccIcalatencyprofile_unset_step2 = `
resource "citrixadc_icalatencyprofile" "tf_unset" {
	name = "tf_test_icalatencyprofile_unset"
	# All unset-eligible attributes removed from config -> the provider must
	# unset them (revert to NITRO defaults).
}
`

func TestAccIcalatencyprofile_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckIcalatencyprofileDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccIcalatencyprofile_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIcalatencyprofileExist("citrixadc_icalatencyprofile.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_icalatencyprofile.tf_unset", "l7latencymonitoring", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_icalatencyprofile.tf_unset", "l7latencythresholdfactor", "120"),
					resource.TestCheckResourceAttr("citrixadc_icalatencyprofile.tf_unset", "l7latencywaittime", "100"),
					resource.TestCheckResourceAttr("citrixadc_icalatencyprofile.tf_unset", "l7latencynotifyinterval", "50"),
					resource.TestCheckResourceAttr("citrixadc_icalatencyprofile.tf_unset", "l7latencymaxnotifycount", "30"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from
				// the appliance) reverts to the documented NITRO defaults, and the
				// implicit post-apply plan must be empty.
				Config: testAccIcalatencyprofile_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIcalatencyprofileExist("citrixadc_icalatencyprofile.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_icalatencyprofile.tf_unset", "l7latencymonitoring", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_icalatencyprofile.tf_unset", "l7latencythresholdfactor", "4"),
					resource.TestCheckResourceAttr("citrixadc_icalatencyprofile.tf_unset", "l7latencywaittime", "20"),
					resource.TestCheckResourceAttr("citrixadc_icalatencyprofile.tf_unset", "l7latencynotifyinterval", "20"),
					resource.TestCheckResourceAttr("citrixadc_icalatencyprofile.tf_unset", "l7latencymaxnotifycount", "5"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckIcalatencyprofileADCValue("tf_test_icalatencyprofile_unset", "l7latencymonitoring", "DISABLED"),
					testAccCheckIcalatencyprofileADCValue("tf_test_icalatencyprofile_unset", "l7latencythresholdfactor", "4"),
				),
			},
		},
	})
}

// testAccCheckIcalatencyprofileADCValue asserts an attribute's value directly on
// the appliance (not just in Terraform state), proving the unset actually
// reverted it.
func testAccCheckIcalatencyprofileADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Icalatencyprofile.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("icalatencyprofile %s not found on appliance", name)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("icalatencyprofile %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func TestAccIcalatencyprofileDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIcalatencyprofileDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_icalatencyprofile.tf_icalatencyprofile_ds", "name", "my_ica_latencyprofile"),
					resource.TestCheckResourceAttr("data.citrixadc_icalatencyprofile.tf_icalatencyprofile_ds", "l7latencymonitoring", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_icalatencyprofile.tf_icalatencyprofile_ds", "l7latencythresholdfactor", "120"),
					resource.TestCheckResourceAttr("data.citrixadc_icalatencyprofile.tf_icalatencyprofile_ds", "l7latencywaittime", "100"),
				),
			},
		},
	})
}
