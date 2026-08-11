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

const testAccLsnrtspalgprofile_basic = `

	resource "citrixadc_lsnrtspalgprofile" "tf_lsnrtspalgprofile" {
		rtspalgprofilename = "my_lsn_rtspalgprofile"
		rtspportrange      = 4200
		rtspidletimeout    = 150
	}
`
const testAccLsnrtspalgprofile_update = `

	resource "citrixadc_lsnrtspalgprofile" "tf_lsnrtspalgprofile" {
		rtspalgprofilename = "my_lsn_rtspalgprofile"
		rtspportrange      = 4500
		rtspidletimeout    = 100
	}
`

func TestAccLsnrtspalgprofile_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLsnrtspalgprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLsnrtspalgprofile_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsnrtspalgprofileExist("citrixadc_lsnrtspalgprofile.tf_lsnrtspalgprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_lsnrtspalgprofile.tf_lsnrtspalgprofile", "rtspalgprofilename", "my_lsn_rtspalgprofile"),
					resource.TestCheckResourceAttr("citrixadc_lsnrtspalgprofile.tf_lsnrtspalgprofile", "rtspportrange", "4200"),
					resource.TestCheckResourceAttr("citrixadc_lsnrtspalgprofile.tf_lsnrtspalgprofile", "rtspidletimeout", "150"),
				),
			},
			{
				Config: testAccLsnrtspalgprofile_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsnrtspalgprofileExist("citrixadc_lsnrtspalgprofile.tf_lsnrtspalgprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_lsnrtspalgprofile.tf_lsnrtspalgprofile", "rtspalgprofilename", "my_lsn_rtspalgprofile"),
					resource.TestCheckResourceAttr("citrixadc_lsnrtspalgprofile.tf_lsnrtspalgprofile", "rtspportrange", "4500"),
					resource.TestCheckResourceAttr("citrixadc_lsnrtspalgprofile.tf_lsnrtspalgprofile", "rtspidletimeout", "100"),
				),
			},
		},
	})
}

func testAccCheckLsnrtspalgprofileExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lsnrtspalgprofile name is set")
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
		data, err := client.FindResource("lsnrtspalgprofile", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("lsnrtspalgprofile %s not found", n)
		}

		return nil
	}
}

func testAccCheckLsnrtspalgprofileDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_lsnrtspalgprofile" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("lsnrtspalgprofile", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("lsnrtspalgprofile %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccLsnrtspalgprofile_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_lsnrtspalgprofile.tf_lsnrtspalgprofile"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLsnrtspalgprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLsnrtspalgprofile_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckLsnrtspalgprofileExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Lsnrtspalgprofile.Type(), "my_lsn_rtspalgprofile"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccLsnrtspalgprofile_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckLsnrtspalgprofileExist(resAddr, nil)),
			},
		},
	})
}

func TestAccLsnrtspalgprofile_import(t *testing.T) {
	const resAddr = "citrixadc_lsnrtspalgprofile.tf_lsnrtspalgprofile"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLsnrtspalgprofileDestroy,
		Steps: []resource.TestStep{
			{Config: testAccLsnrtspalgprofile_basic},
			{
				Config:                  testAccLsnrtspalgprofile_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

const testAccLsnrtspalgprofileDataSource_basic = `

resource "citrixadc_lsnrtspalgprofile" "tf_lsnrtspalgprofile_ds" {
	rtspalgprofilename = "my_lsn_rtspalgprofile_ds"
	rtspportrange      = "4200"
	rtspidletimeout    = 150
}

data "citrixadc_lsnrtspalgprofile" "tf_lsnrtspalgprofile_ds" {
	rtspalgprofilename = citrixadc_lsnrtspalgprofile.tf_lsnrtspalgprofile_ds.rtspalgprofilename
}
`

// testAccLsnrtspalgprofile_unset_step1 sets the unsettable attributes to
// valid non-default values (rtspportrange is mandatory on add, so it is kept
// in both steps).
const testAccLsnrtspalgprofile_unset_step1 = `
resource "citrixadc_lsnrtspalgprofile" "tf_unset" {
	rtspalgprofilename = "tf_test_lsnrtspalgprofile_unset"
	rtspportrange      = "5000"
	rtspidletimeout    = 100
}
`

// testAccLsnrtspalgprofile_unset_step2 removes the unsettable attribute so the
// provider must unset it (revert to NITRO default: rtspidletimeout=120).
// (rtsptransportprotocol is not exercised: this appliance rejects any value
// other than TCP, so a non-default cannot be applied.)
const testAccLsnrtspalgprofile_unset_step2 = `
resource "citrixadc_lsnrtspalgprofile" "tf_unset" {
	rtspalgprofilename = "tf_test_lsnrtspalgprofile_unset"
	rtspportrange      = "5000"
}
`

func TestAccLsnrtspalgprofile_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLsnrtspalgprofileDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccLsnrtspalgprofile_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsnrtspalgprofileExist("citrixadc_lsnrtspalgprofile.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_lsnrtspalgprofile.tf_unset", "rtspidletimeout", "100"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from
				// the appliance) reverts to the documented NITRO defaults, and the
				// implicit post-apply plan must be empty.
				Config: testAccLsnrtspalgprofile_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsnrtspalgprofileExist("citrixadc_lsnrtspalgprofile.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_lsnrtspalgprofile.tf_unset", "rtspidletimeout", "120"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckLsnrtspalgprofileADCValue("tf_test_lsnrtspalgprofile_unset", "rtspidletimeout", "120"),
				),
			},
		},
	})
}

// testAccCheckLsnrtspalgprofileADCValue asserts an attribute's value directly on
// the appliance (not just in Terraform state), proving the unset actually
// reverted it.
func testAccCheckLsnrtspalgprofileADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Lsnrtspalgprofile.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("lsnrtspalgprofile %s not found on appliance", name)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("lsnrtspalgprofile %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func TestAccLsnrtspalgprofile_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckLsnrtspalgprofileDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccLsnrtspalgprofile_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckLsnrtspalgprofileExist("citrixadc_lsnrtspalgprofile.tf_lsnrtspalgprofile", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccLsnrtspalgprofile_basic,
				Check:                    resource.ComposeTestCheckFunc(testAccCheckLsnrtspalgprofileExist("citrixadc_lsnrtspalgprofile.tf_lsnrtspalgprofile", nil)),
			},
		},
	})
}

func TestAccLsnrtspalgprofileDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccLsnrtspalgprofileDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_lsnrtspalgprofile.tf_lsnrtspalgprofile_ds", "rtspalgprofilename", "my_lsn_rtspalgprofile_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_lsnrtspalgprofile.tf_lsnrtspalgprofile_ds", "rtspportrange", "4200"),
					resource.TestCheckResourceAttr("data.citrixadc_lsnrtspalgprofile.tf_lsnrtspalgprofile_ds", "rtspidletimeout", "150"),
				),
			},
		},
	})
}
