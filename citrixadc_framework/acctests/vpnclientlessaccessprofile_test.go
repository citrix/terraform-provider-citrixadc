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

const testAccVpnclientlessaccessprofile_basic = `

	resource "citrixadc_vpnclientlessaccessprofile" "tf_vpnclientlessaccessprofile" {
		profilename = "tf_vpnclientlessaccessprofile"
		requirepersistentcookie = "ON"
	}
`

const testAccVpnclientlessaccessprofile_basic_update = `

	resource "citrixadc_vpnclientlessaccessprofile" "tf_vpnclientlessaccessprofile" {
		profilename = "tf_vpnclientlessaccessprofile"
		requirepersistentcookie = "OFF"
		regexforfindingurlinjavascript = citrixadc_policypatset.tf_patset.name
		regexforfindingurlincss = citrixadc_policypatset.tf_patset.name
		regexforfindingurlinxcomponent = citrixadc_policypatset.tf_patset.name
		regexforfindingurlinxml = citrixadc_policypatset.tf_patset.name
		regexforfindingcustomurls = citrixadc_policypatset.tf_patset.name
	}

	resource "citrixadc_policypatset" "tf_patset" {
		name = "tf_patset"
	}
`

const testAccVpnclientlessaccessprofileDataSource_basic = `

	resource "citrixadc_vpnclientlessaccessprofile" "tf_vpnclientlessaccessprofile" {
		profilename = "tf_vpnclientlessaccessprofile"
		requirepersistentcookie = "ON"
	}

	data "citrixadc_vpnclientlessaccessprofile" "tf_vpnclientlessaccessprofile" {
		profilename = citrixadc_vpnclientlessaccessprofile.tf_vpnclientlessaccessprofile.profilename
	}
`

func TestAccVpnclientlessaccessprofile_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVpnclientlessaccessprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnclientlessaccessprofile_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnclientlessaccessprofileExist("citrixadc_vpnclientlessaccessprofile.tf_vpnclientlessaccessprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnclientlessaccessprofile.tf_vpnclientlessaccessprofile", "profilename", "tf_vpnclientlessaccessprofile"),
					resource.TestCheckResourceAttr("citrixadc_vpnclientlessaccessprofile.tf_vpnclientlessaccessprofile", "requirepersistentcookie", "ON"),
				),
			},
			{
				Config: testAccVpnclientlessaccessprofile_basic_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnclientlessaccessprofileExist("citrixadc_vpnclientlessaccessprofile.tf_vpnclientlessaccessprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnclientlessaccessprofile.tf_vpnclientlessaccessprofile", "requirepersistentcookie", "OFF"),
				),
			},
		},
	})
}

func testAccCheckVpnclientlessaccessprofileExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No vpnclientlessaccessprofile name is set")
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
		data, err := client.FindResource(service.Vpnclientlessaccessprofile.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("vpnclientlessaccessprofile %s not found", n)
		}

		return nil
	}
}

func testAccCheckVpnclientlessaccessprofileDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_vpnclientlessaccessprofile" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Vpnclientlessaccessprofile.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("vpnclientlessaccessprofile %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccVpnclientlessaccessprofile_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_vpnclientlessaccessprofile.tf_vpnclientlessaccessprofile"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVpnclientlessaccessprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnclientlessaccessprofile_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckVpnclientlessaccessprofileExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Vpnclientlessaccessprofile.Type(), "tf_vpnclientlessaccessprofile"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccVpnclientlessaccessprofile_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckVpnclientlessaccessprofileExist(resAddr, nil)),
			},
		},
	})
}

func TestAccVpnclientlessaccessprofile_import(t *testing.T) {
	const resAddr = "citrixadc_vpnclientlessaccessprofile.tf_vpnclientlessaccessprofile"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVpnclientlessaccessprofileDestroy,
		Steps: []resource.TestStep{
			{Config: testAccVpnclientlessaccessprofile_basic},
			{
				Config:                  testAccVpnclientlessaccessprofile_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccVpnclientlessaccessprofile_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckVpnclientlessaccessprofileDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccVpnclientlessaccessprofile_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckVpnclientlessaccessprofileExist("citrixadc_vpnclientlessaccessprofile.tf_vpnclientlessaccessprofile", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccVpnclientlessaccessprofile_basic,
				Check:                    resource.ComposeTestCheckFunc(testAccCheckVpnclientlessaccessprofileExist("citrixadc_vpnclientlessaccessprofile.tf_vpnclientlessaccessprofile", nil)),
			},
		},
	})
}

const testAccVpnclientlessaccessprofile_unset_step1 = `
	resource "citrixadc_vpnclientlessaccessprofile" "tf_unset" {
		profilename             = "tf_vpnclra_unset"
		requirepersistentcookie = "ON"
	}
`

const testAccVpnclientlessaccessprofile_unset_step2 = `
	resource "citrixadc_vpnclientlessaccessprofile" "tf_unset" {
		profilename = "tf_vpnclra_unset"
		# requirepersistentcookie removed -> provider must unset it (revert to "OFF").
	}
`

func TestAccVpnclientlessaccessprofile_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVpnclientlessaccessprofileDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default value is applied and persisted.
				Config: testAccVpnclientlessaccessprofile_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnclientlessaccessprofileExist("citrixadc_vpnclientlessaccessprofile.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnclientlessaccessprofile.tf_unset", "requirepersistentcookie", "ON"),
				),
			},
			{
				// Removing the attribute must unset it: state reverts to the
				// documented NITRO default ("OFF"), the implicit post-apply plan
				// must be empty, and the appliance itself confirms the revert.
				Config: testAccVpnclientlessaccessprofile_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnclientlessaccessprofileExist("citrixadc_vpnclientlessaccessprofile.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnclientlessaccessprofile.tf_unset", "requirepersistentcookie", "OFF"),
					testAccCheckVpnclientlessaccessprofileADCValue("tf_vpnclra_unset", "requirepersistentcookie", "OFF"),
				),
			},
		},
	})
}

// testAccCheckVpnclientlessaccessprofileADCValue asserts an attribute's value
// directly on the appliance (not just in Terraform state), proving the unset
// actually reverted it.
func testAccCheckVpnclientlessaccessprofileADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Vpnclientlessaccessprofile.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("vpnclientlessaccessprofile %s not found on appliance", name)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("vpnclientlessaccessprofile %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func TestAccVpnclientlessaccessprofileDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnclientlessaccessprofileDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_vpnclientlessaccessprofile.tf_vpnclientlessaccessprofile", "profilename", "tf_vpnclientlessaccessprofile"),
					resource.TestCheckResourceAttr("data.citrixadc_vpnclientlessaccessprofile.tf_vpnclientlessaccessprofile", "requirepersistentcookie", "ON"),
				),
			},
		},
	})
}
