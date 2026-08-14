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

const testAccVpnpcoipprofile_add = `

	resource "citrixadc_vpnpcoipprofile" "tf_vpnpcoipprofile" {
		name               = "tf_vpnpcoipprofile"
		conserverurl       = "http://www.example.com"
		sessionidletimeout = 80
	}
`
const testAccVpnpcoipprofile_update = `

	resource "citrixadc_vpnpcoipprofile" "tf_vpnpcoipprofile" {
		name               = "tf_vpnpcoipprofile"
		conserverurl       = "http://www.example.com"
		sessionidletimeout = 90
	}
`

const testAccVpnpcoipprofileDataSource_basic = `

	resource "citrixadc_vpnpcoipprofile" "tf_vpnpcoipprofile" {
		name               = "tf_vpnpcoipprofile"
		conserverurl       = "http://www.example.com"
		sessionidletimeout = 80
	}

	data "citrixadc_vpnpcoipprofile" "tf_vpnpcoipprofile" {
		name = citrixadc_vpnpcoipprofile.tf_vpnpcoipprofile.name
	}
`

func TestAccVpnpcoipprofile_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVpnpcoipprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnpcoipprofile_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnpcoipprofileExist("citrixadc_vpnpcoipprofile.tf_vpnpcoipprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnpcoipprofile.tf_vpnpcoipprofile", "name", "tf_vpnpcoipprofile"),
					resource.TestCheckResourceAttr("citrixadc_vpnpcoipprofile.tf_vpnpcoipprofile", "sessionidletimeout", "80"),
				),
			},
			{
				Config: testAccVpnpcoipprofile_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnpcoipprofileExist("citrixadc_vpnpcoipprofile.tf_vpnpcoipprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnpcoipprofile.tf_vpnpcoipprofile", "name", "tf_vpnpcoipprofile"),
					resource.TestCheckResourceAttr("citrixadc_vpnpcoipprofile.tf_vpnpcoipprofile", "sessionidletimeout", "90"),
				),
			},
		},
	})
}

func testAccCheckVpnpcoipprofileExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No vpnpcoipprofile name is set")
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
		data, err := client.FindResource("vpnpcoipprofile", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("vpnpcoipprofile %s not found", n)
		}

		return nil
	}
}

func testAccCheckVpnpcoipprofileDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_vpnpcoipprofile" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("vpnpcoipprofile", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("vpnpcoipprofile %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccVpnpcoipprofile_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_vpnpcoipprofile.tf_vpnpcoipprofile"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVpnpcoipprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnpcoipprofile_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckVpnpcoipprofileExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Vpnpcoipprofile.Type(), "tf_vpnpcoipprofile"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccVpnpcoipprofile_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckVpnpcoipprofileExist(resAddr, nil)),
			},
		},
	})
}

func TestAccVpnpcoipprofile_import(t *testing.T) {
	const resAddr = "citrixadc_vpnpcoipprofile.tf_vpnpcoipprofile"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVpnpcoipprofileDestroy,
		Steps: []resource.TestStep{
			{Config: testAccVpnpcoipprofile_add},
			{
				Config:                  testAccVpnpcoipprofile_add,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccVpnpcoipprofile_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckVpnpcoipprofileDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccVpnpcoipprofile_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckVpnpcoipprofileExist("citrixadc_vpnpcoipprofile.tf_vpnpcoipprofile", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccVpnpcoipprofile_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckVpnpcoipprofileExist("citrixadc_vpnpcoipprofile.tf_vpnpcoipprofile", nil)),
			},
		},
	})
}

const testAccVpnpcoipprofile_unset_step1 = `
	resource "citrixadc_vpnpcoipprofile" "tf_unset" {
		name               = "tf_test_vpnpcoipprofile_unset"
		conserverurl       = "http://www.example.com"
		icvverification    = "ENABLED"
		sessionidletimeout = 120
	}
`

const testAccVpnpcoipprofile_unset_step2 = `
	resource "citrixadc_vpnpcoipprofile" "tf_unset" {
		name         = "tf_test_vpnpcoipprofile_unset"
		conserverurl = "http://www.example.com"
		# icvverification and sessionidletimeout removed -> provider must unset
		# them (revert to NITRO defaults: DISABLED / 180).
	}
`

func TestAccVpnpcoipprofile_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVpnpcoipprofileDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccVpnpcoipprofile_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnpcoipprofileExist("citrixadc_vpnpcoipprofile.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnpcoipprofile.tf_unset", "icvverification", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_vpnpcoipprofile.tf_unset", "sessionidletimeout", "120"),
				),
			},
			{
				// Removing the attributes must unset them: state reverts to the
				// documented NITRO defaults and the implicit post-apply plan is empty.
				Config: testAccVpnpcoipprofile_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnpcoipprofileExist("citrixadc_vpnpcoipprofile.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnpcoipprofile.tf_unset", "icvverification", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_vpnpcoipprofile.tf_unset", "sessionidletimeout", "180"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckVpnpcoipprofileADCValue("tf_test_vpnpcoipprofile_unset", "icvverification", "DISABLED"),
					testAccCheckVpnpcoipprofileADCValue("tf_test_vpnpcoipprofile_unset", "sessionidletimeout", "180"),
				),
			},
		},
	})
}

// testAccCheckVpnpcoipprofileADCValue asserts an attribute's value directly on
// the appliance (not just in Terraform state), proving the unset reverted it.
func testAccCheckVpnpcoipprofileADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Vpnpcoipprofile.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("vpnpcoipprofile %s not found on appliance", name)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("vpnpcoipprofile %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func TestAccVpnpcoipprofileDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnpcoipprofileDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_vpnpcoipprofile.tf_vpnpcoipprofile", "name", "tf_vpnpcoipprofile"),
					resource.TestCheckResourceAttr("data.citrixadc_vpnpcoipprofile.tf_vpnpcoipprofile", "conserverurl", "http://www.example.com"),
					resource.TestCheckResourceAttr("data.citrixadc_vpnpcoipprofile.tf_vpnpcoipprofile", "sessionidletimeout", "80"),
				),
			},
		},
	})
}
