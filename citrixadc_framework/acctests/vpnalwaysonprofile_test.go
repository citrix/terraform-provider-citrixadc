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

const testAccVpnalwaysonprofile_basic = `

	resource "citrixadc_vpnalwaysonprofile" "tf_vpnalwaysonprofile" {
		name = "tf_vpnalwaysonprofile"
		clientcontrol = "DENY"
		locationbasedvpn = "Remote"
		networkaccessonvpnfailure = "onlyToGateway"
	}
`

const testAccVpnalwaysonprofile_basic_update = `

	resource "citrixadc_vpnalwaysonprofile" "tf_vpnalwaysonprofile" {
		name = "tf_vpnalwaysonprofile"
		clientcontrol = "ALLOW"
		locationbasedvpn = "Everywhere"
		networkaccessonvpnfailure = "fullAccess"
	}
`

const testAccVpnalwaysonprofileDataSource_basic = `

	resource "citrixadc_vpnalwaysonprofile" "tf_vpnalwaysonprofile" {
		name = "tf_vpnalwaysonprofile"
		clientcontrol = "DENY"
		locationbasedvpn = "Remote"
		networkaccessonvpnfailure = "onlyToGateway"
	}

	data "citrixadc_vpnalwaysonprofile" "tf_vpnalwaysonprofile" {
		name = citrixadc_vpnalwaysonprofile.tf_vpnalwaysonprofile.name
	}
`

func TestAccVpnalwaysonprofile_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVpnalwaysonprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnalwaysonprofile_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnalwaysonprofileExist("citrixadc_vpnalwaysonprofile.tf_vpnalwaysonprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnalwaysonprofile.tf_vpnalwaysonprofile", "name", "tf_vpnalwaysonprofile"),
					resource.TestCheckResourceAttr("citrixadc_vpnalwaysonprofile.tf_vpnalwaysonprofile", "clientcontrol", "DENY"),
					resource.TestCheckResourceAttr("citrixadc_vpnalwaysonprofile.tf_vpnalwaysonprofile", "locationbasedvpn", "Remote"),
					resource.TestCheckResourceAttr("citrixadc_vpnalwaysonprofile.tf_vpnalwaysonprofile", "networkaccessonvpnfailure", "onlyToGateway"),
				),
			},
			{
				Config: testAccVpnalwaysonprofile_basic_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnalwaysonprofileExist("citrixadc_vpnalwaysonprofile.tf_vpnalwaysonprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnalwaysonprofile.tf_vpnalwaysonprofile", "clientcontrol", "ALLOW"),
					resource.TestCheckResourceAttr("citrixadc_vpnalwaysonprofile.tf_vpnalwaysonprofile", "locationbasedvpn", "Everywhere"),
					resource.TestCheckResourceAttr("citrixadc_vpnalwaysonprofile.tf_vpnalwaysonprofile", "networkaccessonvpnfailure", "fullAccess"),
				),
			},
		},
	})
}

func testAccCheckVpnalwaysonprofileExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No vpnalwaysonprofile name is set")
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
		data, err := client.FindResource("vpnalwaysonprofile", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("vpnalwaysonprofile %s not found", n)
		}

		return nil
	}
}

func testAccCheckVpnalwaysonprofileDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_vpnalwaysonprofile" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("vpnalwaysonprofile", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("vpnalwaysonprofile %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccVpnalwaysonprofile_unset_step1 = `
resource "citrixadc_vpnalwaysonprofile" "tf_unset" {
  name                      = "tf_vpnalwaysonprofile_unset"
  clientcontrol             = "ALLOW"
  locationbasedvpn          = "Everywhere"
  networkaccessonvpnfailure = "onlyToGateway"
}
`

const testAccVpnalwaysonprofile_unset_step2 = `
resource "citrixadc_vpnalwaysonprofile" "tf_unset" {
  name = "tf_vpnalwaysonprofile_unset"
  # All unset-eligible attributes removed from config -> the provider must
  # unset them (revert to NITRO defaults).
}
`

func TestAccVpnalwaysonprofile_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVpnalwaysonprofileDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccVpnalwaysonprofile_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnalwaysonprofileExist("citrixadc_vpnalwaysonprofile.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnalwaysonprofile.tf_unset", "clientcontrol", "ALLOW"),
					resource.TestCheckResourceAttr("citrixadc_vpnalwaysonprofile.tf_unset", "locationbasedvpn", "Everywhere"),
					resource.TestCheckResourceAttr("citrixadc_vpnalwaysonprofile.tf_unset", "networkaccessonvpnfailure", "onlyToGateway"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from
				// the appliance) reverts to the documented NITRO defaults, and the
				// implicit post-apply plan must be empty.
				Config: testAccVpnalwaysonprofile_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnalwaysonprofileExist("citrixadc_vpnalwaysonprofile.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnalwaysonprofile.tf_unset", "clientcontrol", "DENY"),
					resource.TestCheckResourceAttr("citrixadc_vpnalwaysonprofile.tf_unset", "locationbasedvpn", "Remote"),
					resource.TestCheckResourceAttr("citrixadc_vpnalwaysonprofile.tf_unset", "networkaccessonvpnfailure", "fullAccess"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckVpnalwaysonprofileADCValue("tf_vpnalwaysonprofile_unset", "clientcontrol", "DENY"),
					testAccCheckVpnalwaysonprofileADCValue("tf_vpnalwaysonprofile_unset", "locationbasedvpn", "Remote"),
					testAccCheckVpnalwaysonprofileADCValue("tf_vpnalwaysonprofile_unset", "networkaccessonvpnfailure", "fullAccess"),
				),
			},
		},
	})
}

// testAccCheckVpnalwaysonprofileADCValue asserts an attribute's value directly
// on the appliance (not just in Terraform state), proving the unset actually
// reverted it.
func testAccCheckVpnalwaysonprofileADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Vpnalwaysonprofile.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("vpnalwaysonprofile %s not found on appliance", name)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("vpnalwaysonprofile %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func TestAccVpnalwaysonprofile_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_vpnalwaysonprofile.tf_vpnalwaysonprofile"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVpnalwaysonprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnalwaysonprofile_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckVpnalwaysonprofileExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Vpnalwaysonprofile.Type(), "tf_vpnalwaysonprofile"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccVpnalwaysonprofile_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckVpnalwaysonprofileExist(resAddr, nil)),
			},
		},
	})
}

func TestAccVpnalwaysonprofile_import(t *testing.T) {
	const resAddr = "citrixadc_vpnalwaysonprofile.tf_vpnalwaysonprofile"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVpnalwaysonprofileDestroy,
		Steps: []resource.TestStep{
			{Config: testAccVpnalwaysonprofile_basic},
			{
				Config:                  testAccVpnalwaysonprofile_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccVpnalwaysonprofile_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckVpnalwaysonprofileDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccVpnalwaysonprofile_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckVpnalwaysonprofileExist("citrixadc_vpnalwaysonprofile.tf_vpnalwaysonprofile", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccVpnalwaysonprofile_basic,
				Check:                    resource.ComposeTestCheckFunc(testAccCheckVpnalwaysonprofileExist("citrixadc_vpnalwaysonprofile.tf_vpnalwaysonprofile", nil)),
			},
		},
	})
}

func TestAccVpnalwaysonprofileDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnalwaysonprofileDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_vpnalwaysonprofile.tf_vpnalwaysonprofile", "name", "tf_vpnalwaysonprofile"),
					resource.TestCheckResourceAttr("data.citrixadc_vpnalwaysonprofile.tf_vpnalwaysonprofile", "clientcontrol", "DENY"),
					resource.TestCheckResourceAttr("data.citrixadc_vpnalwaysonprofile.tf_vpnalwaysonprofile", "locationbasedvpn", "Remote"),
					resource.TestCheckResourceAttr("data.citrixadc_vpnalwaysonprofile.tf_vpnalwaysonprofile", "networkaccessonvpnfailure", "onlyToGateway"),
				),
			},
		},
	})
}
