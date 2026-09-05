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

const testAccLsnappsprofile_basic = `

	resource "citrixadc_lsnappsprofile" "tf_lsnappsprofile" {
		appsprofilename   = "my_lsn_appsprofile"
		transportprotocol = "TCP"
		mapping           = "ENDPOINT-INDEPENDENT"
	}
`
const testAccLsnappsprofile_update = `

	resource "citrixadc_lsnappsprofile" "tf_lsnappsprofile" {
		appsprofilename   = "my_lsn_appsprofile"
		transportprotocol = "TCP"
		mapping           = "ADDRESS-DEPENDENT"
	}
`

func TestAccLsnappsprofile_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLsnappsprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLsnappsprofile_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsnappsprofileExist("citrixadc_lsnappsprofile.tf_lsnappsprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_lsnappsprofile.tf_lsnappsprofile", "appsprofilename", "my_lsn_appsprofile"),
					resource.TestCheckResourceAttr("citrixadc_lsnappsprofile.tf_lsnappsprofile", "transportprotocol", "TCP"),
					resource.TestCheckResourceAttr("citrixadc_lsnappsprofile.tf_lsnappsprofile", "mapping", "ENDPOINT-INDEPENDENT"),
				),
			},
			{
				Config: testAccLsnappsprofile_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsnappsprofileExist("citrixadc_lsnappsprofile.tf_lsnappsprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_lsnappsprofile.tf_lsnappsprofile", "appsprofilename", "my_lsn_appsprofile"),
					resource.TestCheckResourceAttr("citrixadc_lsnappsprofile.tf_lsnappsprofile", "transportprotocol", "TCP"),
					resource.TestCheckResourceAttr("citrixadc_lsnappsprofile.tf_lsnappsprofile", "mapping", "ADDRESS-DEPENDENT"),
				),
			},
		},
	})
}

func testAccCheckLsnappsprofileExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lsnappsprofile name is set")
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
		data, err := client.FindResource("lsnappsprofile", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("lsnappsprofile %s not found", n)
		}

		return nil
	}
}

func testAccCheckLsnappsprofileDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_lsnappsprofile" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("lsnappsprofile", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("lsnappsprofile %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccLsnappsprofile_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_lsnappsprofile.tf_lsnappsprofile"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLsnappsprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLsnappsprofile_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckLsnappsprofileExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Lsnappsprofile.Type(), "my_lsn_appsprofile"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccLsnappsprofile_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckLsnappsprofileExist(resAddr, nil)),
			},
		},
	})
}

func TestAccLsnappsprofile_import(t *testing.T) {
	const resAddr = "citrixadc_lsnappsprofile.tf_lsnappsprofile"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLsnappsprofileDestroy,
		Steps: []resource.TestStep{
			{Config: testAccLsnappsprofile_basic},
			{
				Config:                  testAccLsnappsprofile_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

const testAccLsnappsprofileDataSource_basic = `

resource "citrixadc_lsnappsprofile" "tf_lsnappsprofile_ds" {
	appsprofilename   = "my_lsn_appsprofile_ds"
	transportprotocol = "TCP"
	mapping           = "ENDPOINT-INDEPENDENT"
	filtering         = "ENDPOINT-INDEPENDENT"
	ippooling         = "RANDOM"
}

data "citrixadc_lsnappsprofile" "tf_lsnappsprofile_ds" {
	appsprofilename = citrixadc_lsnappsprofile.tf_lsnappsprofile_ds.appsprofilename
}
`

func TestAccLsnappsprofile_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckLsnappsprofileDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.0.0"},
				},
				Config: testAccLsnappsprofile_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckLsnappsprofileExist("citrixadc_lsnappsprofile.tf_lsnappsprofile", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccLsnappsprofile_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckLsnappsprofileExist("citrixadc_lsnappsprofile.tf_lsnappsprofile", nil)),
			},
		},
	})
}

// The unset test exercises the unset-eligible attributes of lsnappsprofile that
// can be set together on a single TCP profile: filtering, ippooling, mapping,
// tcpproxy. l2info is excluded because NITRO rejects (errorcode 257) enabling it
// together with tcpproxy=ENABLED. td is excluded because setting a non-default
// value requires a pre-existing traffic domain on the appliance. Filtering must
// be at least as strict as mapping, so both are set to ENDPOINT-INDEPENDENT.
const testAccLsnappsprofile_unset_step1 = `
resource "citrixadc_lsnappsprofile" "tf_unset" {
  appsprofilename   = "tf_test_lsnappsprofile_unset"
  transportprotocol = "TCP"
  filtering         = "ENDPOINT-INDEPENDENT"
  ippooling         = "PAIRED"
  mapping           = "ENDPOINT-INDEPENDENT"
  tcpproxy          = "ENABLED"
}
`

const testAccLsnappsprofile_unset_step2 = `
resource "citrixadc_lsnappsprofile" "tf_unset" {
  appsprofilename   = "tf_test_lsnappsprofile_unset"
  transportprotocol = "TCP"
  # All unset-eligible attributes removed from config -> the provider must
  # unset them (revert to NITRO defaults).
}
`

func TestAccLsnappsprofile_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLsnappsprofileDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccLsnappsprofile_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsnappsprofileExist("citrixadc_lsnappsprofile.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_lsnappsprofile.tf_unset", "filtering", "ENDPOINT-INDEPENDENT"),
					resource.TestCheckResourceAttr("citrixadc_lsnappsprofile.tf_unset", "ippooling", "PAIRED"),
					resource.TestCheckResourceAttr("citrixadc_lsnappsprofile.tf_unset", "mapping", "ENDPOINT-INDEPENDENT"),
					resource.TestCheckResourceAttr("citrixadc_lsnappsprofile.tf_unset", "tcpproxy", "ENABLED"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from
				// the appliance) reverts to the documented NITRO defaults, and the
				// implicit post-apply plan must be empty.
				Config: testAccLsnappsprofile_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsnappsprofileExist("citrixadc_lsnappsprofile.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_lsnappsprofile.tf_unset", "filtering", "ADDRESS-PORT-DEPENDENT"),
					resource.TestCheckResourceAttr("citrixadc_lsnappsprofile.tf_unset", "ippooling", "RANDOM"),
					resource.TestCheckResourceAttr("citrixadc_lsnappsprofile.tf_unset", "mapping", "ADDRESS-PORT-DEPENDENT"),
					resource.TestCheckResourceAttr("citrixadc_lsnappsprofile.tf_unset", "tcpproxy", "DISABLED"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckLsnappsprofileADCValue("tf_test_lsnappsprofile_unset", "mapping", "ADDRESS-PORT-DEPENDENT"),
					testAccCheckLsnappsprofileADCValue("tf_test_lsnappsprofile_unset", "tcpproxy", "DISABLED"),
				),
			},
		},
	})
}

// testAccCheckLsnappsprofileADCValue asserts an attribute's value directly on
// the appliance (not just in Terraform state), proving the unset actually
// reverted it.
func testAccCheckLsnappsprofileADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Lsnappsprofile.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("lsnappsprofile %s not found on appliance", name)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("lsnappsprofile %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func TestAccLsnappsprofileDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccLsnappsprofileDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_lsnappsprofile.tf_lsnappsprofile_ds", "appsprofilename", "my_lsn_appsprofile_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_lsnappsprofile.tf_lsnappsprofile_ds", "transportprotocol", "TCP"),
					resource.TestCheckResourceAttr("data.citrixadc_lsnappsprofile.tf_lsnappsprofile_ds", "mapping", "ENDPOINT-INDEPENDENT"),
					resource.TestCheckResourceAttr("data.citrixadc_lsnappsprofile.tf_lsnappsprofile_ds", "filtering", "ENDPOINT-INDEPENDENT"),
					resource.TestCheckResourceAttr("data.citrixadc_lsnappsprofile.tf_lsnappsprofile_ds", "ippooling", "RANDOM"),
				),
			},
		},
	})
}
