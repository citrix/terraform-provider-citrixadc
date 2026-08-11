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

const testAccContentinspectionprofile_basic = `

	resource "citrixadc_contentinspectionprofile" "tf_contentinspectionprofile" {
		name             = "my_ci_profile"
		type             = "InlineInspection"
		ingressinterface = "LA/2"
		egressinterface  = "LA/3"
	}
`
const testAccContentinspectionprofile_update = `

	resource "citrixadc_contentinspectionprofile" "tf_contentinspectionprofile" {
		name             = "my_ci_profile"
		type             = "InlineInspection"
		ingressinterface = "LA/3"
		egressinterface  = "LA/2"
	}
`

func TestAccContentinspectionprofile_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckContentinspectionprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccContentinspectionprofile_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckContentinspectionprofileExist("citrixadc_contentinspectionprofile.tf_contentinspectionprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_contentinspectionprofile.tf_contentinspectionprofile", "name", "my_ci_profile"),
					resource.TestCheckResourceAttr("citrixadc_contentinspectionprofile.tf_contentinspectionprofile", "type", "InlineInspection"),
					resource.TestCheckResourceAttr("citrixadc_contentinspectionprofile.tf_contentinspectionprofile", "ingressinterface", "LA/2"),
					resource.TestCheckResourceAttr("citrixadc_contentinspectionprofile.tf_contentinspectionprofile", "egressinterface", "LA/3"),
				),
			},
			{
				Config: testAccContentinspectionprofile_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckContentinspectionprofileExist("citrixadc_contentinspectionprofile.tf_contentinspectionprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_contentinspectionprofile.tf_contentinspectionprofile", "name", "my_ci_profile"),
					resource.TestCheckResourceAttr("citrixadc_contentinspectionprofile.tf_contentinspectionprofile", "type", "InlineInspection"),
					resource.TestCheckResourceAttr("citrixadc_contentinspectionprofile.tf_contentinspectionprofile", "ingressinterface", "LA/3"),
					resource.TestCheckResourceAttr("citrixadc_contentinspectionprofile.tf_contentinspectionprofile", "egressinterface", "LA/2"),
				),
			},
		},
	})
}

func TestAccContentinspectionprofile_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckContentinspectionprofileDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccContentinspectionprofile_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckContentinspectionprofileExist("citrixadc_contentinspectionprofile.tf_contentinspectionprofile", nil),
				),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccContentinspectionprofile_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckContentinspectionprofileExist("citrixadc_contentinspectionprofile.tf_contentinspectionprofile", nil),
				),
			},
		},
	})
}

func TestAccContentinspectionprofile_import(t *testing.T) {
	const resAddr = "citrixadc_contentinspectionprofile.tf_contentinspectionprofile"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckContentinspectionprofileDestroy,
		Steps: []resource.TestStep{
			{Config: testAccContentinspectionprofile_basic},
			{
				Config:                  testAccContentinspectionprofile_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccCheckContentinspectionprofileExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No contentinspectionprofile name is set")
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
		data, err := client.FindResource("contentinspectionprofile", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("contentinspectionprofile %s not found", n)
		}

		return nil
	}
}

func testAccCheckContentinspectionprofileDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_contentinspectionprofile" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("contentinspectionprofile", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("contentinspectionprofile %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccContentinspectionprofile_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_contentinspectionprofile.tf_contentinspectionprofile"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckContentinspectionprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccContentinspectionprofile_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckContentinspectionprofileExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Contentinspectionprofile.Type(), "my_ci_profile"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccContentinspectionprofile_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckContentinspectionprofileExist(resAddr, nil)),
			},
		},
	})
}

// Unset test: egressvlan and ingressvlan are the only unset-eligible attributes
// per the NITRO spec (the unset payload lists only these two). VLAN 1 is the
// default appliance VLAN, used here as a valid non-default value (the NITRO
// default is 0). Removing the attributes from config must unset them, reverting
// the appliance to 0.
const testAccContentinspectionprofile_unset_step1 = `
	resource "citrixadc_contentinspectionprofile" "tf_unset" {
		name             = "tf_ci_profile_unset"
		type             = "InlineInspection"
		ingressinterface = "LA/2"
		egressinterface  = "LA/3"
		ingressvlan      = 1
		egressvlan       = 1
	}
`

const testAccContentinspectionprofile_unset_step2 = `
	resource "citrixadc_contentinspectionprofile" "tf_unset" {
		name             = "tf_ci_profile_unset"
		type             = "InlineInspection"
		ingressinterface = "LA/2"
		egressinterface  = "LA/3"
		# ingressvlan/egressvlan removed -> provider must unset them (revert to 0).
	}
`

func TestAccContentinspectionprofile_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckContentinspectionprofileDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default vlan values are applied and persisted.
				Config: testAccContentinspectionprofile_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckContentinspectionprofileExist("citrixadc_contentinspectionprofile.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_contentinspectionprofile.tf_unset", "ingressvlan", "1"),
					resource.TestCheckResourceAttr("citrixadc_contentinspectionprofile.tf_unset", "egressvlan", "1"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from
				// the appliance) reverts to the NITRO default (0), and the
				// implicit post-apply plan must be empty.
				Config: testAccContentinspectionprofile_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckContentinspectionprofileExist("citrixadc_contentinspectionprofile.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_contentinspectionprofile.tf_unset", "ingressvlan", "0"),
					resource.TestCheckResourceAttr("citrixadc_contentinspectionprofile.tf_unset", "egressvlan", "0"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckContentinspectionprofileADCValue("tf_ci_profile_unset", "ingressvlan", "0"),
					testAccCheckContentinspectionprofileADCValue("tf_ci_profile_unset", "egressvlan", "0"),
				),
			},
		},
	})
}

// testAccCheckContentinspectionprofileADCValue asserts an attribute's value
// directly on the appliance (not just in Terraform state), proving the unset
// actually reverted it.
func testAccCheckContentinspectionprofileADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Contentinspectionprofile.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("contentinspectionprofile %s not found on appliance", name)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("contentinspectionprofile %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

const testAccContentinspectionprofileDataSource_basic = `

	resource "citrixadc_contentinspectionprofile" "tf_contentinspectionprofile" {
		name             = "my_ci_profile_ds"
		type             = "InlineInspection"
		ingressinterface = "LA/2"
		egressinterface  = "LA/3"
	}

	data "citrixadc_contentinspectionprofile" "tf_contentinspectionprofile_ds" {
		name = citrixadc_contentinspectionprofile.tf_contentinspectionprofile.name
	}
`

func TestAccContentinspectionprofileDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccContentinspectionprofileDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_contentinspectionprofile.tf_contentinspectionprofile_ds", "name", "my_ci_profile_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_contentinspectionprofile.tf_contentinspectionprofile_ds", "type", "InlineInspection"),
					resource.TestCheckResourceAttr("data.citrixadc_contentinspectionprofile.tf_contentinspectionprofile_ds", "ingressinterface", "LA/2"),
					resource.TestCheckResourceAttr("data.citrixadc_contentinspectionprofile.tf_contentinspectionprofile_ds", "egressinterface", "LA/3"),
				),
			},
		},
	})
}
