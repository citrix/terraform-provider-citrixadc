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

const testAccVpnpcoipvserverprofile_add = `
	resource "citrixadc_vpnpcoipvserverprofile" "tf_vpnpcoipvserverprofile" {
		name        = "tf_vpnpcoipvserverprofile"
		logindomain = "domainname"
		udpport     = "802"
	}
`
const testAccVpnpcoipvserverprofile_update = `
	resource "citrixadc_vpnpcoipvserverprofile" "tf_vpnpcoipvserverprofile" {
		name        = "tf_vpnpcoipvserverprofile"
		logindomain = "domainname"
		udpport     = "200"
	}
`

const testAccVpnpcoipvserverprofileDataSource_basic = `
	resource "citrixadc_vpnpcoipvserverprofile" "tf_vpnpcoipvserverprofile" {
		name        = "tf_vpnpcoipvserverprofile"
		logindomain = "domainname"
		udpport     = "802"
	}

	data "citrixadc_vpnpcoipvserverprofile" "tf_vpnpcoipvserverprofile" {
		name = citrixadc_vpnpcoipvserverprofile.tf_vpnpcoipvserverprofile.name
	}
`

func TestAccVpnpcoipvserverprofile_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVpnpcoipvserverprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnpcoipvserverprofile_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnpcoipvserverprofileExist("citrixadc_vpnpcoipvserverprofile.tf_vpnpcoipvserverprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnpcoipvserverprofile.tf_vpnpcoipvserverprofile", "name", "tf_vpnpcoipvserverprofile"),
					resource.TestCheckResourceAttr("citrixadc_vpnpcoipvserverprofile.tf_vpnpcoipvserverprofile", "udpport", "802"),
				),
			},
			{
				Config: testAccVpnpcoipvserverprofile_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnpcoipvserverprofileExist("citrixadc_vpnpcoipvserverprofile.tf_vpnpcoipvserverprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnpcoipvserverprofile.tf_vpnpcoipvserverprofile", "name", "tf_vpnpcoipvserverprofile"),
					resource.TestCheckResourceAttr("citrixadc_vpnpcoipvserverprofile.tf_vpnpcoipvserverprofile", "udpport", "200"),
				),
			},
		},
	})
}

func testAccCheckVpnpcoipvserverprofileExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No vpnpcoipvserverprofile name is set")
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
		data, err := client.FindResource("vpnpcoipvserverprofile", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("vpnpcoipvserverprofile %s not found", n)
		}

		return nil
	}
}

func testAccCheckVpnpcoipvserverprofileDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_vpnpcoipvserverprofile" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("vpnpcoipvserverprofile", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("vpnpcoipvserverprofile %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccVpnpcoipvserverprofile_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_vpnpcoipvserverprofile.tf_vpnpcoipvserverprofile"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVpnpcoipvserverprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnpcoipvserverprofile_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckVpnpcoipvserverprofileExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Vpnpcoipvserverprofile.Type(), "tf_vpnpcoipvserverprofile"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccVpnpcoipvserverprofile_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckVpnpcoipvserverprofileExist(resAddr, nil)),
			},
		},
	})
}

func TestAccVpnpcoipvserverprofile_import(t *testing.T) {
	const resAddr = "citrixadc_vpnpcoipvserverprofile.tf_vpnpcoipvserverprofile"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVpnpcoipvserverprofileDestroy,
		Steps: []resource.TestStep{
			{Config: testAccVpnpcoipvserverprofile_add},
			{
				Config:                  testAccVpnpcoipvserverprofile_add,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccVpnpcoipvserverprofile_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckVpnpcoipvserverprofileDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccVpnpcoipvserverprofile_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckVpnpcoipvserverprofileExist("citrixadc_vpnpcoipvserverprofile.tf_vpnpcoipvserverprofile", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccVpnpcoipvserverprofile_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckVpnpcoipvserverprofileExist("citrixadc_vpnpcoipvserverprofile.tf_vpnpcoipvserverprofile", nil)),
			},
		},
	})
}

const testAccVpnpcoipvserverprofile_unset_step1 = `
	resource "citrixadc_vpnpcoipvserverprofile" "tf_unset" {
		name        = "tf_pcoip_unset"
		logindomain = "domainname"
		udpport     = 802
	}
`

const testAccVpnpcoipvserverprofile_unset_step2 = `
	resource "citrixadc_vpnpcoipvserverprofile" "tf_unset" {
		name        = "tf_pcoip_unset"
		logindomain = "domainname"
		# udpport removed from config -> the provider must unset it (revert to
		# NITRO default 4172).
	}
`

func TestAccVpnpcoipvserverprofile_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVpnpcoipvserverprofileDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default value is applied and persisted.
				Config: testAccVpnpcoipvserverprofile_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnpcoipvserverprofileExist("citrixadc_vpnpcoipvserverprofile.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnpcoipvserverprofile.tf_unset", "udpport", "802"),
				),
			},
			{
				// Removing udpport must unset it: state (read back from the
				// appliance) reverts to the documented NITRO default, and the
				// implicit post-apply plan must be empty.
				Config: testAccVpnpcoipvserverprofile_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnpcoipvserverprofileExist("citrixadc_vpnpcoipvserverprofile.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnpcoipvserverprofile.tf_unset", "udpport", "4172"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckVpnpcoipvserverprofileADCValue("tf_pcoip_unset", "udpport", "4172"),
				),
			},
		},
	})
}

// testAccCheckVpnpcoipvserverprofileADCValue asserts an attribute's value
// directly on the appliance (not just in Terraform state), proving the unset
// actually reverted it.
func testAccCheckVpnpcoipvserverprofileADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Vpnpcoipvserverprofile.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("vpnpcoipvserverprofile %s not found on appliance", name)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("vpnpcoipvserverprofile %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func TestAccVpnpcoipvserverprofileDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnpcoipvserverprofileDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_vpnpcoipvserverprofile.tf_vpnpcoipvserverprofile", "name", "tf_vpnpcoipvserverprofile"),
					resource.TestCheckResourceAttr("data.citrixadc_vpnpcoipvserverprofile.tf_vpnpcoipvserverprofile", "logindomain", "domainname"),
					resource.TestCheckResourceAttr("data.citrixadc_vpnpcoipvserverprofile.tf_vpnpcoipvserverprofile", "udpport", "802"),
				),
			},
		},
	})
}
