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

const testAccLsnsipalgprofile_basic = `

	resource "citrixadc_lsnsipalgprofile" "tf_lsnsipalgprofile" {
		sipalgprofilename      = "my_lsn_sipalgprofile"
		datasessionidletimeout = 150
		sipsessiontimeout      = 150
		registrationtimeout    = 150
		sipsrcportrange        = "4400"
		siptransportprotocol   = "TCP"
	}
  
`
const testAccLsnsipalgprofile_update = `

	resource "citrixadc_lsnsipalgprofile" "tf_lsnsipalgprofile" {
		sipalgprofilename      = "my_lsn_sipalgprofile"
		datasessionidletimeout = 100
		sipsessiontimeout      = 100
		registrationtimeout    = 100
		sipsrcportrange        = "4200"
		siptransportprotocol   = "TCP"
	}
  
`

func TestAccLsnsipalgprofile_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLsnsipalgprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLsnsipalgprofile_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsnsipalgprofileExist("citrixadc_lsnsipalgprofile.tf_lsnsipalgprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_lsnsipalgprofile.tf_lsnsipalgprofile", "sipalgprofilename", "my_lsn_sipalgprofile"),
					resource.TestCheckResourceAttr("citrixadc_lsnsipalgprofile.tf_lsnsipalgprofile", "datasessionidletimeout", "150"),
					resource.TestCheckResourceAttr("citrixadc_lsnsipalgprofile.tf_lsnsipalgprofile", "sipsessiontimeout", "150"),
					resource.TestCheckResourceAttr("citrixadc_lsnsipalgprofile.tf_lsnsipalgprofile", "registrationtimeout", "150"),
					resource.TestCheckResourceAttr("citrixadc_lsnsipalgprofile.tf_lsnsipalgprofile", "sipsrcportrange", "4400"),
					resource.TestCheckResourceAttr("citrixadc_lsnsipalgprofile.tf_lsnsipalgprofile", "siptransportprotocol", "TCP"),
				),
			},
			{
				Config: testAccLsnsipalgprofile_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsnsipalgprofileExist("citrixadc_lsnsipalgprofile.tf_lsnsipalgprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_lsnsipalgprofile.tf_lsnsipalgprofile", "sipalgprofilename", "my_lsn_sipalgprofile"),
					resource.TestCheckResourceAttr("citrixadc_lsnsipalgprofile.tf_lsnsipalgprofile", "datasessionidletimeout", "100"),
					resource.TestCheckResourceAttr("citrixadc_lsnsipalgprofile.tf_lsnsipalgprofile", "sipsessiontimeout", "100"),
					resource.TestCheckResourceAttr("citrixadc_lsnsipalgprofile.tf_lsnsipalgprofile", "registrationtimeout", "100"),
					resource.TestCheckResourceAttr("citrixadc_lsnsipalgprofile.tf_lsnsipalgprofile", "sipsrcportrange", "4200"),
					resource.TestCheckResourceAttr("citrixadc_lsnsipalgprofile.tf_lsnsipalgprofile", "siptransportprotocol", "TCP"),
				),
			},
		},
	})
}

func TestAccLsnsipalgprofile_import(t *testing.T) {
	const resAddr = "citrixadc_lsnsipalgprofile.tf_lsnsipalgprofile"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLsnsipalgprofileDestroy,
		Steps: []resource.TestStep{
			{Config: testAccLsnsipalgprofile_basic},
			{
				Config:                  testAccLsnsipalgprofile_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccCheckLsnsipalgprofileExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lsnsipalgprofile name is set")
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
		data, err := client.FindResource("lsnsipalgprofile", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("lsnsipalgprofile %s not found", n)
		}

		return nil
	}
}

func testAccCheckLsnsipalgprofileDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_lsnsipalgprofile" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("lsnsipalgprofile", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("lsnsipalgprofile %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccLsnsipalgprofile_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_lsnsipalgprofile.tf_lsnsipalgprofile"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLsnsipalgprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLsnsipalgprofile_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckLsnsipalgprofileExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Lsnsipalgprofile.Type(), "my_lsn_sipalgprofile"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccLsnsipalgprofile_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckLsnsipalgprofileExist(resAddr, nil)),
			},
		},
	})
}

// step1 sets every unset-eligible attribute to a valid non-default value;
// step2 removes them (keeping only the key and the mandatory
// siptransportprotocol) so the provider must unset them back to NITRO defaults.
const testAccLsnsipalgprofile_unset_step1 = `
resource "citrixadc_lsnsipalgprofile" "tf_unset" {
	sipalgprofilename      = "tf_test_lsnsipalgprofile_unset"
	siptransportprotocol   = "TCP"
	sipsrcportrange        = "4400"
	datasessionidletimeout = 200
	sipsessiontimeout      = 700
	registrationtimeout    = 90
	opencontactpinhole     = "DISABLED"
	openrecordroutepinhole = "DISABLED"
	openregisterpinhole    = "DISABLED"
	openroutepinhole       = "DISABLED"
	openviapinhole         = "DISABLED"
	rport                  = "DISABLED"
}
`

const testAccLsnsipalgprofile_unset_step2 = `
resource "citrixadc_lsnsipalgprofile" "tf_unset" {
	sipalgprofilename    = "tf_test_lsnsipalgprofile_unset"
	siptransportprotocol = "TCP"
	sipsrcportrange      = "4400"
	# All unset-eligible attributes removed from config -> the provider must
	# unset them (revert to the documented NITRO defaults).
}
`

func TestAccLsnsipalgprofile_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLsnsipalgprofileDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccLsnsipalgprofile_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsnsipalgprofileExist("citrixadc_lsnsipalgprofile.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_lsnsipalgprofile.tf_unset", "datasessionidletimeout", "200"),
					resource.TestCheckResourceAttr("citrixadc_lsnsipalgprofile.tf_unset", "sipsessiontimeout", "700"),
					resource.TestCheckResourceAttr("citrixadc_lsnsipalgprofile.tf_unset", "registrationtimeout", "90"),
					resource.TestCheckResourceAttr("citrixadc_lsnsipalgprofile.tf_unset", "opencontactpinhole", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_lsnsipalgprofile.tf_unset", "openrecordroutepinhole", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_lsnsipalgprofile.tf_unset", "openregisterpinhole", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_lsnsipalgprofile.tf_unset", "openroutepinhole", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_lsnsipalgprofile.tf_unset", "openviapinhole", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_lsnsipalgprofile.tf_unset", "rport", "DISABLED"),
				),
			},
			{
				// Removing the attributes must unset them: state reverts to the
				// documented NITRO defaults and the implicit post-apply plan is empty.
				Config: testAccLsnsipalgprofile_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsnsipalgprofileExist("citrixadc_lsnsipalgprofile.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_lsnsipalgprofile.tf_unset", "datasessionidletimeout", "120"),
					resource.TestCheckResourceAttr("citrixadc_lsnsipalgprofile.tf_unset", "sipsessiontimeout", "600"),
					resource.TestCheckResourceAttr("citrixadc_lsnsipalgprofile.tf_unset", "registrationtimeout", "60"),
					resource.TestCheckResourceAttr("citrixadc_lsnsipalgprofile.tf_unset", "opencontactpinhole", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_lsnsipalgprofile.tf_unset", "openrecordroutepinhole", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_lsnsipalgprofile.tf_unset", "openregisterpinhole", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_lsnsipalgprofile.tf_unset", "openroutepinhole", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_lsnsipalgprofile.tf_unset", "openviapinhole", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_lsnsipalgprofile.tf_unset", "rport", "ENABLED"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckLsnsipalgprofileADCValue("tf_test_lsnsipalgprofile_unset", "opencontactpinhole", "ENABLED"),
					testAccCheckLsnsipalgprofileADCValue("tf_test_lsnsipalgprofile_unset", "rport", "ENABLED"),
					testAccCheckLsnsipalgprofileADCValue("tf_test_lsnsipalgprofile_unset", "sipsessiontimeout", "600"),
				),
			},
		},
	})
}

// testAccCheckLsnsipalgprofileADCValue asserts an attribute's value directly on
// the appliance (not just in Terraform state), proving the unset reverted it.
func testAccCheckLsnsipalgprofileADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Lsnsipalgprofile.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("lsnsipalgprofile %s not found on appliance", name)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("lsnsipalgprofile %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func TestAccLsnsipalgprofile_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckLsnsipalgprofileDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccLsnsipalgprofile_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckLsnsipalgprofileExist("citrixadc_lsnsipalgprofile.tf_lsnsipalgprofile", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccLsnsipalgprofile_basic,
				Check:                    resource.ComposeTestCheckFunc(testAccCheckLsnsipalgprofileExist("citrixadc_lsnsipalgprofile.tf_lsnsipalgprofile", nil)),
			},
		},
	})
}

const testAccLsnsipalgprofileDataSource_basic = `

resource "citrixadc_lsnsipalgprofile" "tf_lsnsipalgprofile_ds" {
	sipalgprofilename      = "my_lsn_sipalgprofile_ds"
	datasessionidletimeout = 150
	sipsessiontimeout      = 150
	registrationtimeout    = 150
	sipsrcportrange        = "4400"
	siptransportprotocol   = "TCP"
}

data "citrixadc_lsnsipalgprofile" "tf_lsnsipalgprofile_ds" {
	sipalgprofilename = citrixadc_lsnsipalgprofile.tf_lsnsipalgprofile_ds.sipalgprofilename
}
`

func TestAccLsnsipalgprofileDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccLsnsipalgprofileDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_lsnsipalgprofile.tf_lsnsipalgprofile_ds", "sipalgprofilename", "my_lsn_sipalgprofile_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_lsnsipalgprofile.tf_lsnsipalgprofile_ds", "datasessionidletimeout", "150"),
					resource.TestCheckResourceAttr("data.citrixadc_lsnsipalgprofile.tf_lsnsipalgprofile_ds", "sipsessiontimeout", "150"),
					resource.TestCheckResourceAttr("data.citrixadc_lsnsipalgprofile.tf_lsnsipalgprofile_ds", "registrationtimeout", "150"),
					resource.TestCheckResourceAttr("data.citrixadc_lsnsipalgprofile.tf_lsnsipalgprofile_ds", "sipsrcportrange", "4400"),
					resource.TestCheckResourceAttr("data.citrixadc_lsnsipalgprofile.tf_lsnsipalgprofile_ds", "siptransportprotocol", "TCP"),
				),
			},
		},
	})
}
