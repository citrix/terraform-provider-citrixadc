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

const testAccLsnhttphdrlogprofile_basic = `


resource "citrixadc_lsnhttphdrlogprofile" "tf_lsnhttphdrlogprofile" {
	httphdrlogprofilename = "my_lsn_httphdrlogprofile"
	logurl                = "DISABLED"
	logversion            = "DISABLED"
	loghost               = "DISABLED"
	}
  
`
const testAccLsnhttphdrlogprofile_update = `


resource "citrixadc_lsnhttphdrlogprofile" "tf_lsnhttphdrlogprofile" {
	httphdrlogprofilename = "my_lsn_httphdrlogprofile"
	logurl                = "ENABLED"
	logversion            = "ENABLED"
	loghost               = "ENABLED"
	}
  
`

func TestAccLsnhttphdrlogprofile_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLsnhttphdrlogprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLsnhttphdrlogprofile_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsnhttphdrlogprofileExist("citrixadc_lsnhttphdrlogprofile.tf_lsnhttphdrlogprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_lsnhttphdrlogprofile.tf_lsnhttphdrlogprofile", "httphdrlogprofilename", "my_lsn_httphdrlogprofile"),
					resource.TestCheckResourceAttr("citrixadc_lsnhttphdrlogprofile.tf_lsnhttphdrlogprofile", "logurl", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_lsnhttphdrlogprofile.tf_lsnhttphdrlogprofile", "logversion", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_lsnhttphdrlogprofile.tf_lsnhttphdrlogprofile", "loghost", "DISABLED"),
				),
			},
			{
				Config: testAccLsnhttphdrlogprofile_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsnhttphdrlogprofileExist("citrixadc_lsnhttphdrlogprofile.tf_lsnhttphdrlogprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_lsnhttphdrlogprofile.tf_lsnhttphdrlogprofile", "httphdrlogprofilename", "my_lsn_httphdrlogprofile"),
					resource.TestCheckResourceAttr("citrixadc_lsnhttphdrlogprofile.tf_lsnhttphdrlogprofile", "logurl", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_lsnhttphdrlogprofile.tf_lsnhttphdrlogprofile", "logversion", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_lsnhttphdrlogprofile.tf_lsnhttphdrlogprofile", "loghost", "ENABLED"),
				),
			},
		},
	})
}

func TestAccLsnhttphdrlogprofile_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_lsnhttphdrlogprofile.tf_lsnhttphdrlogprofile"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLsnhttphdrlogprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLsnhttphdrlogprofile_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckLsnhttphdrlogprofileExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Lsnhttphdrlogprofile.Type(), "my_lsn_httphdrlogprofile"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccLsnhttphdrlogprofile_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckLsnhttphdrlogprofileExist(resAddr, nil)),
			},
		},
	})
}

func TestAccLsnhttphdrlogprofile_import(t *testing.T) {
	const resAddr = "citrixadc_lsnhttphdrlogprofile.tf_lsnhttphdrlogprofile"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLsnhttphdrlogprofileDestroy,
		Steps: []resource.TestStep{
			{Config: testAccLsnhttphdrlogprofile_basic},
			{
				Config:                  testAccLsnhttphdrlogprofile_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccCheckLsnhttphdrlogprofileExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lsnhttphdrlogprofile name is set")
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
		data, err := client.FindResource("lsnhttphdrlogprofile", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("lsnhttphdrlogprofile %s not found", n)
		}

		return nil
	}
}

func testAccCheckLsnhttphdrlogprofileDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_lsnhttphdrlogprofile" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("lsnhttphdrlogprofile", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("lsnhttphdrlogprofile %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccLsnhttphdrlogprofileDataSource_basic = `

resource "citrixadc_lsnhttphdrlogprofile" "tf_lsnhttphdrlogprofile_ds" {
	httphdrlogprofilename = "my_lsn_httphdrlogprofile_ds"
	logurl                = "DISABLED"
	logversion            = "DISABLED"
	loghost               = "DISABLED"
}

data "citrixadc_lsnhttphdrlogprofile" "tf_lsnhttphdrlogprofile_ds" {
	httphdrlogprofilename = citrixadc_lsnhttphdrlogprofile.tf_lsnhttphdrlogprofile_ds.httphdrlogprofilename
}
`

func TestAccLsnhttphdrlogprofile_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckLsnhttphdrlogprofileDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.0.0"},
				},
				Config: testAccLsnhttphdrlogprofile_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckLsnhttphdrlogprofileExist("citrixadc_lsnhttphdrlogprofile.tf_lsnhttphdrlogprofile", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccLsnhttphdrlogprofile_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckLsnhttphdrlogprofileExist("citrixadc_lsnhttphdrlogprofile.tf_lsnhttphdrlogprofile", nil)),
			},
		},
	})
}

// Unset test: all four log* attributes are Optional+Computed booleans-as-strings
// with a documented NITRO default of "ENABLED". Step 1 sets them to the
// non-default "DISABLED"; step 2 removes them, so the provider must unset them
// back to "ENABLED".
const testAccLsnhttphdrlogprofile_unset_step1 = `
resource "citrixadc_lsnhttphdrlogprofile" "tf_unset" {
	httphdrlogprofilename = "tf_test_lsnhttphdrlogprofile_unset"
	loghost               = "DISABLED"
	logmethod             = "DISABLED"
	logurl                = "DISABLED"
	logversion            = "DISABLED"
}
`

const testAccLsnhttphdrlogprofile_unset_step2 = `
resource "citrixadc_lsnhttphdrlogprofile" "tf_unset" {
	httphdrlogprofilename = "tf_test_lsnhttphdrlogprofile_unset"
	# All unset-eligible attributes removed from config -> the provider must
	# unset them (revert to NITRO defaults, "ENABLED").
}
`

func TestAccLsnhttphdrlogprofile_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLsnhttphdrlogprofileDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccLsnhttphdrlogprofile_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsnhttphdrlogprofileExist("citrixadc_lsnhttphdrlogprofile.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_lsnhttphdrlogprofile.tf_unset", "loghost", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_lsnhttphdrlogprofile.tf_unset", "logmethod", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_lsnhttphdrlogprofile.tf_unset", "logurl", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_lsnhttphdrlogprofile.tf_unset", "logversion", "DISABLED"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from the
				// appliance) reverts to the documented NITRO defaults, and the
				// implicit post-apply plan must be empty.
				Config: testAccLsnhttphdrlogprofile_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsnhttphdrlogprofileExist("citrixadc_lsnhttphdrlogprofile.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_lsnhttphdrlogprofile.tf_unset", "loghost", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_lsnhttphdrlogprofile.tf_unset", "logmethod", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_lsnhttphdrlogprofile.tf_unset", "logurl", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_lsnhttphdrlogprofile.tf_unset", "logversion", "ENABLED"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckLsnhttphdrlogprofileADCValue("tf_test_lsnhttphdrlogprofile_unset", "loghost", "ENABLED"),
					testAccCheckLsnhttphdrlogprofileADCValue("tf_test_lsnhttphdrlogprofile_unset", "logurl", "ENABLED"),
				),
			},
		},
	})
}

// testAccCheckLsnhttphdrlogprofileADCValue asserts an attribute's value directly
// on the appliance (not just in Terraform state), proving the unset actually
// reverted it.
func testAccCheckLsnhttphdrlogprofileADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Lsnhttphdrlogprofile.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("lsnhttphdrlogprofile %s not found on appliance", name)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("lsnhttphdrlogprofile %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func TestAccLsnhttphdrlogprofileDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccLsnhttphdrlogprofileDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_lsnhttphdrlogprofile.tf_lsnhttphdrlogprofile_ds", "httphdrlogprofilename", "my_lsn_httphdrlogprofile_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_lsnhttphdrlogprofile.tf_lsnhttphdrlogprofile_ds", "logurl", "DISABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_lsnhttphdrlogprofile.tf_lsnhttphdrlogprofile_ds", "logversion", "DISABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_lsnhttphdrlogprofile.tf_lsnhttphdrlogprofile_ds", "loghost", "DISABLED"),
				),
			},
		},
	})
}
