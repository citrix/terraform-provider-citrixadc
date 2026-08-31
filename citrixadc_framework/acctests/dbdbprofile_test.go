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

const testAccDbdbprofile_basic = `
	resource "citrixadc_dbdbprofile" "tf_dbdbprofile" {
		name           = "my_dbprofile"
		stickiness     = "YES"
		conmultiplex   = "ENABLED"
		interpretquery = "YES"
	}
`
const testAccDbdbprofile_update = `
	resource "citrixadc_dbdbprofile" "tf_dbdbprofile" {
		name           = "my_dbprofile"
		stickiness     = "NO"
		conmultiplex   = "DISABLED"
		interpretquery = "NO"
	}
`

const testAccDbdbprofileDataSource_basic = `
	resource "citrixadc_dbdbprofile" "tf_dbdbprofile" {
		name           = "my_dbprofile"
		stickiness     = "YES"
		conmultiplex   = "ENABLED"
		interpretquery = "YES"
	}

	data "citrixadc_dbdbprofile" "tf_dbdbprofile_datasource" {
		name = citrixadc_dbdbprofile.tf_dbdbprofile.name
	}
`

func TestAccDbdbprofile_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDbdbprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDbdbprofile_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDbdbprofileExist("citrixadc_dbdbprofile.tf_dbdbprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_dbdbprofile.tf_dbdbprofile", "name", "my_dbprofile"),
					resource.TestCheckResourceAttr("citrixadc_dbdbprofile.tf_dbdbprofile", "stickiness", "YES"),
					resource.TestCheckResourceAttr("citrixadc_dbdbprofile.tf_dbdbprofile", "conmultiplex", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_dbdbprofile.tf_dbdbprofile", "interpretquery", "YES"),
				),
			},
			{
				Config: testAccDbdbprofile_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDbdbprofileExist("citrixadc_dbdbprofile.tf_dbdbprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_dbdbprofile.tf_dbdbprofile", "name", "my_dbprofile"),
					resource.TestCheckResourceAttr("citrixadc_dbdbprofile.tf_dbdbprofile", "stickiness", "NO"),
					resource.TestCheckResourceAttr("citrixadc_dbdbprofile.tf_dbdbprofile", "conmultiplex", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_dbdbprofile.tf_dbdbprofile", "interpretquery", "NO"),
				),
			},
		},
	})
}

func TestAccDbdbprofile_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_dbdbprofile.tf_dbdbprofile"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDbdbprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDbdbprofile_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckDbdbprofileExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Dbdbprofile.Type(), "my_dbprofile"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccDbdbprofile_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckDbdbprofileExist(resAddr, nil)),
			},
		},
	})
}

func testAccCheckDbdbprofileExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No dbdbprofile name is set")
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
		data, err := client.FindResource(service.Dbdbprofile.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("dbdbprofile %s not found", n)
		}

		return nil
	}
}

func testAccCheckDbdbprofileDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_dbdbprofile" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Dbdbprofile.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("dbdbprofile %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccDbdbprofile_import(t *testing.T) {
	const resAddr = "citrixadc_dbdbprofile.tf_dbdbprofile"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDbdbprofileDestroy,
		Steps: []resource.TestStep{
			{Config: testAccDbdbprofile_basic},
			{
				Config:                  testAccDbdbprofile_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccDbdbprofile_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckDbdbprofileDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.0.0"},
				},
				Config: testAccDbdbprofile_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckDbdbprofileExist("citrixadc_dbdbprofile.tf_dbdbprofile", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccDbdbprofile_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckDbdbprofileExist("citrixadc_dbdbprofile.tf_dbdbprofile", nil)),
			},
		},
	})
}

// The dbdbprofile unset test covers all spec-unsettable, mutable attributes
// that carry a documented NITRO default: conmultiplex, enablecachingconmuxoff,
// interpretquery and stickiness. Step1 sets them to non-default values; step2
// removes them from config so the provider must unset them (revert to NITRO
// defaults).
const testAccDbdbprofile_unset_step1 = `
resource "citrixadc_dbdbprofile" "tf_unset" {
  name                   = "tf_test_dbdbprofile_unset"
  conmultiplex           = "DISABLED"
  enablecachingconmuxoff = "ENABLED"
  interpretquery         = "NO"
  stickiness             = "YES"
}
`

const testAccDbdbprofile_unset_step2 = `
resource "citrixadc_dbdbprofile" "tf_unset" {
  name = "tf_test_dbdbprofile_unset"
  # All unset-eligible attributes removed from config -> the provider must
  # unset them (revert to NITRO defaults).
}
`

func TestAccDbdbprofile_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDbdbprofileDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccDbdbprofile_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDbdbprofileExist("citrixadc_dbdbprofile.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_dbdbprofile.tf_unset", "conmultiplex", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_dbdbprofile.tf_unset", "enablecachingconmuxoff", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_dbdbprofile.tf_unset", "interpretquery", "NO"),
					resource.TestCheckResourceAttr("citrixadc_dbdbprofile.tf_unset", "stickiness", "YES"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from
				// the appliance) reverts to the documented NITRO defaults, and the
				// implicit post-apply plan must be empty.
				Config: testAccDbdbprofile_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDbdbprofileExist("citrixadc_dbdbprofile.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_dbdbprofile.tf_unset", "conmultiplex", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_dbdbprofile.tf_unset", "enablecachingconmuxoff", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_dbdbprofile.tf_unset", "interpretquery", "YES"),
					resource.TestCheckResourceAttr("citrixadc_dbdbprofile.tf_unset", "stickiness", "NO"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckDbdbprofileADCValue("tf_test_dbdbprofile_unset", "conmultiplex", "ENABLED"),
					testAccCheckDbdbprofileADCValue("tf_test_dbdbprofile_unset", "interpretquery", "YES"),
					testAccCheckDbdbprofileADCValue("tf_test_dbdbprofile_unset", "stickiness", "NO"),
				),
			},
		},
	})
}

// testAccCheckDbdbprofileADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted
// it.
func testAccCheckDbdbprofileADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Dbdbprofile.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("dbdbprofile %s not found on appliance", name)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("dbdbprofile %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func TestAccDbdbprofileDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDbdbprofileDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_dbdbprofile.tf_dbdbprofile_datasource", "name", "my_dbprofile"),
					resource.TestCheckResourceAttr("data.citrixadc_dbdbprofile.tf_dbdbprofile_datasource", "stickiness", "YES"),
					resource.TestCheckResourceAttr("data.citrixadc_dbdbprofile.tf_dbdbprofile_datasource", "conmultiplex", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_dbdbprofile.tf_dbdbprofile_datasource", "interpretquery", "YES"),
				),
			},
		},
	})
}
