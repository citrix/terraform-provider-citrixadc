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

const testAccLsnlogprofile_basic = `

	resource "citrixadc_lsnlogprofile" "tf_lsnlogprofile" {
		logprofilename = "my_lsn_logprofile"
		logsubscrinfo   = "ENABLED"
		logcompact      = "ENABLED"
		logipfix        = "ENABLED"
	}
	
`
const testAccLsnlogprofile_update = `

	resource "citrixadc_lsnlogprofile" "tf_lsnlogprofile" {
		logprofilename = "my_lsn_logprofile"
		logsubscrinfo   = "DISABLED"
		logcompact      = "DISABLED"
		logipfix        = "DISABLED"
	}
	
`

func TestAccLsnlogprofile_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLsnlogprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLsnlogprofile_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsnlogprofileExist("citrixadc_lsnlogprofile.tf_lsnlogprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_lsnlogprofile.tf_lsnlogprofile", "logprofilename", "my_lsn_logprofile"),
					resource.TestCheckResourceAttr("citrixadc_lsnlogprofile.tf_lsnlogprofile", "logsubscrinfo", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_lsnlogprofile.tf_lsnlogprofile", "logcompact", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_lsnlogprofile.tf_lsnlogprofile", "logipfix", "ENABLED"),
				),
			},
			{
				Config: testAccLsnlogprofile_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsnlogprofileExist("citrixadc_lsnlogprofile.tf_lsnlogprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_lsnlogprofile.tf_lsnlogprofile", "logprofilename", "my_lsn_logprofile"),
					resource.TestCheckResourceAttr("citrixadc_lsnlogprofile.tf_lsnlogprofile", "logsubscrinfo", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_lsnlogprofile.tf_lsnlogprofile", "logcompact", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_lsnlogprofile.tf_lsnlogprofile", "logipfix", "DISABLED"),
				),
			},
		},
	})
}

func TestAccLsnlogprofile_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckLsnlogprofileDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccLsnlogprofile_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckLsnlogprofileExist("citrixadc_lsnlogprofile.tf_lsnlogprofile", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccLsnlogprofile_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckLsnlogprofileExist("citrixadc_lsnlogprofile.tf_lsnlogprofile", nil)),
			},
		},
	})
}

func testAccCheckLsnlogprofileExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lsnlogprofile name is set")
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
		data, err := client.FindResource("lsnlogprofile", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("lsnlogprofile %s not found", n)
		}

		return nil
	}
}

func testAccCheckLsnlogprofileDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_lsnlogprofile" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("lsnlogprofile", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("lsnlogprofile %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccLsnlogprofile_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_lsnlogprofile.tf_lsnlogprofile"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLsnlogprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLsnlogprofile_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckLsnlogprofileExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Lsnlogprofile.Type(), "my_lsn_logprofile"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccLsnlogprofile_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckLsnlogprofileExist(resAddr, nil)),
			},
		},
	})
}

// Unset test: step1 sets the unset-eligible attributes to non-default values;
// step2 removes them so the provider must unset them (revert to NITRO defaults).
const testAccLsnlogprofile_unset_step1 = `
resource "citrixadc_lsnlogprofile" "tf_unset" {
  logprofilename  = "tf_test_lsnlogprofile_unset"
  logsubscrinfo   = "ENABLED"
  logcompact      = "DISABLED"
  logipfix        = "ENABLED"
  logsessdeletion = "DISABLED"
}
`

const testAccLsnlogprofile_unset_step2 = `
resource "citrixadc_lsnlogprofile" "tf_unset" {
  logprofilename = "tf_test_lsnlogprofile_unset"
  # All unset-eligible attributes removed from config -> the provider must
  # unset them (revert to NITRO defaults).
}
`

func TestAccLsnlogprofile_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLsnlogprofileDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccLsnlogprofile_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsnlogprofileExist("citrixadc_lsnlogprofile.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_lsnlogprofile.tf_unset", "logsubscrinfo", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_lsnlogprofile.tf_unset", "logcompact", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_lsnlogprofile.tf_unset", "logipfix", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_lsnlogprofile.tf_unset", "logsessdeletion", "DISABLED"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from
				// the appliance) reverts to the documented NITRO defaults, and the
				// implicit post-apply plan must be empty.
				Config: testAccLsnlogprofile_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsnlogprofileExist("citrixadc_lsnlogprofile.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_lsnlogprofile.tf_unset", "logsubscrinfo", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_lsnlogprofile.tf_unset", "logcompact", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_lsnlogprofile.tf_unset", "logipfix", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_lsnlogprofile.tf_unset", "logsessdeletion", "ENABLED"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckLsnlogprofileADCValue("tf_test_lsnlogprofile_unset", "logsubscrinfo", "DISABLED"),
					testAccCheckLsnlogprofileADCValue("tf_test_lsnlogprofile_unset", "logcompact", "ENABLED"),
				),
			},
		},
	})
}

// testAccCheckLsnlogprofileADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted it.
func testAccCheckLsnlogprofileADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Lsnlogprofile.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("lsnlogprofile %s not found on appliance", name)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("lsnlogprofile %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func TestAccLsnlogprofile_import(t *testing.T) {
	const resAddr = "citrixadc_lsnlogprofile.tf_lsnlogprofile"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLsnlogprofileDestroy,
		Steps: []resource.TestStep{
			{Config: testAccLsnlogprofile_basic},
			{
				Config:                  testAccLsnlogprofile_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

const testAccLsnlogprofileDataSource_basic = `

resource "citrixadc_lsnlogprofile" "tf_lsnlogprofile_ds" {
	logprofilename = "my_lsn_logprofile_ds"
	logsubscrinfo  = "ENABLED"
	logcompact     = "ENABLED"
	logipfix       = "ENABLED"
}

data "citrixadc_lsnlogprofile" "tf_lsnlogprofile_ds" {
	logprofilename = citrixadc_lsnlogprofile.tf_lsnlogprofile_ds.logprofilename
}
`

func TestAccLsnlogprofileDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccLsnlogprofileDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_lsnlogprofile.tf_lsnlogprofile_ds", "logprofilename", "my_lsn_logprofile_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_lsnlogprofile.tf_lsnlogprofile_ds", "logsubscrinfo", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_lsnlogprofile.tf_lsnlogprofile_ds", "logcompact", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_lsnlogprofile.tf_lsnlogprofile_ds", "logipfix", "ENABLED"),
				),
			},
		},
	})
}
