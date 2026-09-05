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

const testAccIpsecalgprofile_basic = `
	resource "citrixadc_ipsecalgprofile" "tf_ipsecalgprofile" {
		name              = "my_ipsecalgprofile"
		ikesessiontimeout = 50
		espsessiontimeout = 20
		connfailover      = "DISABLED"
	}
  
`

const testAccIpsecalgprofile_update = `
	resource "citrixadc_ipsecalgprofile" "tf_ipsecalgprofile" {
		name              = "my_ipsecalgprofile"
		ikesessiontimeout = 40
		espsessiontimeout = 30
		connfailover      = "ENABLED"
	}
  
`

func TestAccIpsecalgprofile_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckIpsecalgprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccIpsecalgprofile_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpsecalgprofileExist("citrixadc_ipsecalgprofile.tf_ipsecalgprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_ipsecalgprofile.tf_ipsecalgprofile", "name", "my_ipsecalgprofile"),
					resource.TestCheckResourceAttr("citrixadc_ipsecalgprofile.tf_ipsecalgprofile", "ikesessiontimeout", "50"),
					resource.TestCheckResourceAttr("citrixadc_ipsecalgprofile.tf_ipsecalgprofile", "espsessiontimeout", "20"),
					resource.TestCheckResourceAttr("citrixadc_ipsecalgprofile.tf_ipsecalgprofile", "connfailover", "DISABLED"),
				),
			},
			{
				Config: testAccIpsecalgprofile_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpsecalgprofileExist("citrixadc_ipsecalgprofile.tf_ipsecalgprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_ipsecalgprofile.tf_ipsecalgprofile", "name", "my_ipsecalgprofile"),
					resource.TestCheckResourceAttr("citrixadc_ipsecalgprofile.tf_ipsecalgprofile", "ikesessiontimeout", "40"),
					resource.TestCheckResourceAttr("citrixadc_ipsecalgprofile.tf_ipsecalgprofile", "espsessiontimeout", "30"),
					resource.TestCheckResourceAttr("citrixadc_ipsecalgprofile.tf_ipsecalgprofile", "connfailover", "ENABLED"),
				),
			},
		},
	})
}

func testAccCheckIpsecalgprofileExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ipsecalgprofile name is set")
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
		data, err := client.FindResource("ipsecalgprofile", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("ipsecalgprofile %s not found", n)
		}

		return nil
	}
}

func testAccCheckIpsecalgprofileDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_ipsecalgprofile" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("ipsecalgprofile", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("ipsecalgprofile %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccIpsecalgprofile_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_ipsecalgprofile.tf_ipsecalgprofile"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckIpsecalgprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccIpsecalgprofile_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckIpsecalgprofileExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Ipsecalgprofile.Type(), "my_ipsecalgprofile"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccIpsecalgprofile_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckIpsecalgprofileExist(resAddr, nil)),
			},
		},
	})
}

func TestAccIpsecalgprofile_import(t *testing.T) {
	const resAddr = "citrixadc_ipsecalgprofile.tf_ipsecalgprofile"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckIpsecalgprofileDestroy,
		Steps: []resource.TestStep{
			{Config: testAccIpsecalgprofile_basic},
			{
				Config:                  testAccIpsecalgprofile_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

const testAccIpsecalgprofileDataSource_basic = `
	resource "citrixadc_ipsecalgprofile" "tf_ipsecalgprofile_ds" {
		name              = "my_ipsecalgprofile_ds"
		ikesessiontimeout = 50
		espsessiontimeout = 20
		connfailover      = "DISABLED"
	}

	data "citrixadc_ipsecalgprofile" "tf_ipsecalgprofile_ds" {
		name = citrixadc_ipsecalgprofile.tf_ipsecalgprofile_ds.name
	}
`

func TestAccIpsecalgprofile_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckIpsecalgprofileDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.0.0"},
				},
				Config: testAccIpsecalgprofile_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckIpsecalgprofileExist("citrixadc_ipsecalgprofile.tf_ipsecalgprofile", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccIpsecalgprofile_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckIpsecalgprofileExist("citrixadc_ipsecalgprofile.tf_ipsecalgprofile", nil)),
			},
		},
	})
}

// TestAccIpsecalgprofile_unset verifies that removing unset-eligible attributes
// from config reverts them to their documented NITRO defaults (connfailover
// ENABLED, espgatetimeout 30, espsessiontimeout 60, ikesessiontimeout 60).
const testAccIpsecalgprofile_unset_step1 = `
resource "citrixadc_ipsecalgprofile" "tf_unset" {
  name              = "tf_test_ipsecalgprofile_unset"
  connfailover      = "DISABLED"
  espgatetimeout    = 100
  espsessiontimeout = 120
  ikesessiontimeout = 90
}
`

const testAccIpsecalgprofile_unset_step2 = `
resource "citrixadc_ipsecalgprofile" "tf_unset" {
  name = "tf_test_ipsecalgprofile_unset"
  # All unset-eligible attributes removed from config -> the provider must
  # unset them (revert to NITRO defaults).
}
`

func TestAccIpsecalgprofile_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckIpsecalgprofileDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccIpsecalgprofile_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpsecalgprofileExist("citrixadc_ipsecalgprofile.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_ipsecalgprofile.tf_unset", "connfailover", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_ipsecalgprofile.tf_unset", "espgatetimeout", "100"),
					resource.TestCheckResourceAttr("citrixadc_ipsecalgprofile.tf_unset", "espsessiontimeout", "120"),
					resource.TestCheckResourceAttr("citrixadc_ipsecalgprofile.tf_unset", "ikesessiontimeout", "90"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from
				// the appliance) reverts to the documented NITRO defaults, and the
				// implicit post-apply plan must be empty.
				Config: testAccIpsecalgprofile_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpsecalgprofileExist("citrixadc_ipsecalgprofile.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_ipsecalgprofile.tf_unset", "connfailover", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_ipsecalgprofile.tf_unset", "espgatetimeout", "30"),
					resource.TestCheckResourceAttr("citrixadc_ipsecalgprofile.tf_unset", "espsessiontimeout", "60"),
					resource.TestCheckResourceAttr("citrixadc_ipsecalgprofile.tf_unset", "ikesessiontimeout", "60"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckIpsecalgprofileADCValue("tf_test_ipsecalgprofile_unset", "connfailover", "ENABLED"),
					testAccCheckIpsecalgprofileADCValue("tf_test_ipsecalgprofile_unset", "espgatetimeout", "30"),
				),
			},
		},
	})
}

// testAccCheckIpsecalgprofileADCValue asserts an attribute's value directly on
// the appliance (not just in Terraform state), proving the unset actually
// reverted it.
func testAccCheckIpsecalgprofileADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Ipsecalgprofile.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("ipsecalgprofile %s not found on appliance", name)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("ipsecalgprofile %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func TestAccIpsecalgprofileDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIpsecalgprofileDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_ipsecalgprofile.tf_ipsecalgprofile_ds", "name", "my_ipsecalgprofile_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_ipsecalgprofile.tf_ipsecalgprofile_ds", "ikesessiontimeout", "50"),
					resource.TestCheckResourceAttr("data.citrixadc_ipsecalgprofile.tf_ipsecalgprofile_ds", "espsessiontimeout", "20"),
					resource.TestCheckResourceAttr("data.citrixadc_ipsecalgprofile.tf_ipsecalgprofile_ds", "connfailover", "DISABLED"),
				),
			},
		},
	})
}
