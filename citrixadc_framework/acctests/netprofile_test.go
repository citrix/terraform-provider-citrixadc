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

func TestAccNetprofile_basic(t *testing.T) {
	if isCpxRun {
		t.Skip("CPX 12.0 is outdated for this resource")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNetprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNetprofile_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetprofileExist("citrixadc_netprofile.tf_netprofile", nil),
				),
			},
			{
				Config: testAccNetprofile_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetprofileExist("citrixadc_netprofile.tf_netprofile", nil),
				),
			},
			{
				Config: testAccNetprofile_basic_step3,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetprofileExist("citrixadc_netprofile.tf_netprofile", nil),
				),
			},
		},
	})
}

const testAccNetprofile_proxyprotocolaftertlshandshake = `

	resource "citrixadc_netprofile" "tf_netprofile_proxyprotocolaftertlshandshake" {
		name                   = "tf_netprofile2"
		proxyprotocol          = "ENABLED"
		proxyprotocoltxversion = "V2"
	}
`

func TestAccNetprofile_proxyprotocolaftertlshandshake(t *testing.T) {
	if isCpxRun {
		t.Skip("CPX 12.0 is outdated for this resource")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNetprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNetprofile_proxyprotocolaftertlshandshake,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetprofileExist("citrixadc_netprofile.tf_netprofile_proxyprotocolaftertlshandshake", nil),
					// resource.TestCheckResourceAttr("citrixadc_netprofile.tf_netprofile_proxyprotocolaftertlshandshake", "proxyprotocolaftertlshandshake", "ENABLED"),
				),
			},
		},
	})
}
func testAccCheckNetprofileExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lb vserver name is set")
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
		data, err := client.FindResource(service.Netprofile.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("LB vserver %s not found", n)
		}

		return nil
	}
}

func testAccCheckNetprofileDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_netprofile" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Netprofile.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("LB vserver %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccNetprofile_basic_step1 = `

resource "citrixadc_netprofile" "tf_netprofile" {
    name 				   = "tf_netprofile"
    proxyprotocol 		   = "ENABLED"
    proxyprotocoltxversion = "V1"
}

`

const testAccNetprofile_basic_step2 = `

resource "citrixadc_netprofile" "tf_netprofile" {
    name 				   = "tf_netprofile"
    proxyprotocol          = "ENABLED"
    proxyprotocoltxversion = "V2"
}

`

const testAccNetprofile_basic_step3 = `

resource "citrixadc_netprofile" "tf_netprofile" {
    name 				   = "tf_netprofile2"
    proxyprotocol 		   = "ENABLED"
    proxyprotocoltxversion = "V2"
}

`

// The netprofile unset test covers the spec-unsettable, mutable attributes that
// have a documented NITRO server default: overridelsn, proxyprotocol,
// proxyprotocolaftertlshandshake, proxyprotocoltxversion, srcippersistency.
// srcip and mbf are excluded (no documented server default); td is create-only.
const testAccNetprofile_unset_step1 = `
resource "citrixadc_netprofile" "tf_unset" {
  name                           = "tf_test_netprofile_unset"
  overridelsn                    = "ENABLED"
  proxyprotocol                  = "ENABLED"
  proxyprotocoltxversion         = "V2"
  srcippersistency               = "ENABLED"
}
`

const testAccNetprofile_unset_step2 = `
resource "citrixadc_netprofile" "tf_unset" {
  name = "tf_test_netprofile_unset"
  # All unset-eligible attributes removed from config -> the provider must
  # unset them (revert to the documented NITRO defaults).
}
`

func TestAccNetprofile_unset(t *testing.T) {
	if isCpxRun {
		t.Skip("CPX 12.0 is outdated for this resource")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNetprofileDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccNetprofile_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetprofileExist("citrixadc_netprofile.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_netprofile.tf_unset", "overridelsn", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_netprofile.tf_unset", "proxyprotocol", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_netprofile.tf_unset", "proxyprotocoltxversion", "V2"),
					resource.TestCheckResourceAttr("citrixadc_netprofile.tf_unset", "srcippersistency", "ENABLED"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from
				// the appliance) reverts to the documented NITRO defaults, and the
				// implicit post-apply plan must be empty.
				Config: testAccNetprofile_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetprofileExist("citrixadc_netprofile.tf_unset", nil),
					// overridelsn is omit-on-default: after unset the appliance omits it from
					// GET, so it reads back as null/absent (driven by the unsetOnRemove
					// modifier, no schema Default — matches the import round-trip).
					resource.TestCheckNoResourceAttr("citrixadc_netprofile.tf_unset", "overridelsn"),
					resource.TestCheckResourceAttr("citrixadc_netprofile.tf_unset", "proxyprotocol", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_netprofile.tf_unset", "proxyprotocolaftertlshandshake", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_netprofile.tf_unset", "proxyprotocoltxversion", "V1"),
					resource.TestCheckResourceAttr("citrixadc_netprofile.tf_unset", "srcippersistency", "DISABLED"),
					// Independent appliance-level confirmation the unset took effect.
					// (overridelsn is omit-on-default -> absent from GET, so it is asserted
					// only at the state level above, not here.)
					testAccCheckNetprofileADCValue("tf_test_netprofile_unset", "proxyprotocol", "DISABLED"),
					testAccCheckNetprofileADCValue("tf_test_netprofile_unset", "srcippersistency", "DISABLED"),
				),
			},
		},
	})
}

// testAccCheckNetprofileADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted it.
func testAccCheckNetprofileADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Netprofile.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("netprofile %s not found on appliance", name)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("netprofile %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func TestAccNetprofile_selfHealing(t *testing.T) {
	if isCpxRun {
		t.Skip("CPX 12.0 is outdated for this resource")
	}
	const resAddr = "citrixadc_netprofile.tf_netprofile"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNetprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNetprofile_basic_step1,
				Check:  resource.ComposeTestCheckFunc(testAccCheckNetprofileExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Netprofile.Type(), "tf_netprofile"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccNetprofile_basic_step1,
				Check:  resource.ComposeTestCheckFunc(testAccCheckNetprofileExist(resAddr, nil)),
			},
		},
	})
}

func TestAccNetprofile_import(t *testing.T) {
	if isCpxRun {
		t.Skip("CPX 12.0 is outdated for this resource")
	}
	const resAddr = "citrixadc_netprofile.tf_netprofile"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNetprofileDestroy,
		Steps: []resource.TestStep{
			{Config: testAccNetprofile_basic_step1},
			{
				Config:                  testAccNetprofile_basic_step1,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

const testAccNetprofileDataSource_basic = `

	resource "citrixadc_netprofile" "tf_netprofile" {
		name                   = "tf_netprofile_ds"
		proxyprotocol          = "ENABLED"
		proxyprotocoltxversion = "V1"
		srcippersistency       = "ENABLED"
	}

	data "citrixadc_netprofile" "tf_netprofile_ds" {
		name = citrixadc_netprofile.tf_netprofile.name
	}
`

func TestAccNetprofile_sdkv2StateUpgrade(t *testing.T) {
	if isCpxRun {
		t.Skip("CPX 12.0 is outdated for this resource")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckNetprofileDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.0.0"},
				},
				Config: testAccNetprofile_basic_step1,
				Check:  resource.ComposeTestCheckFunc(testAccCheckNetprofileExist("citrixadc_netprofile.tf_netprofile", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccNetprofile_basic_step1,
				Check:  resource.ComposeTestCheckFunc(testAccCheckNetprofileExist("citrixadc_netprofile.tf_netprofile", nil)),
			},
		},
	})
}

func TestAccNetprofileDataSource_basic(t *testing.T) {
	if isCpxRun {
		t.Skip("CPX 12.0 is outdated for this resource")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccNetprofileDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_netprofile.tf_netprofile_ds", "name", "tf_netprofile_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_netprofile.tf_netprofile_ds", "proxyprotocol", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_netprofile.tf_netprofile_ds", "proxyprotocoltxversion", "V1"),
					resource.TestCheckResourceAttr("data.citrixadc_netprofile.tf_netprofile_ds", "srcippersistency", "ENABLED"),
					// Universal runtime-binding proof.
					resource.TestCheckResourceAttrSet("data.citrixadc_netprofile.tf_netprofile_ds", "id"),
				),
			},
		},
	})
}
