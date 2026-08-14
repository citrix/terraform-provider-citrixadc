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

const testAccLsntransportprofile_basic = `

	resource "citrixadc_lsntransportprofile" "tf_lsntransportprofile" {
		transportprofilename = "my_lsn_transportprofile"
		transportprotocol    = "TCP"
		portquota            = 10
		sessionquota         = 10
		groupsessionlimit    = 100
	}
  
`

const testAccLsntransportprofile_update = `

	resource "citrixadc_lsntransportprofile" "tf_lsntransportprofile" {
		transportprofilename = "my_lsn_transportprofile"
		transportprotocol    = "TCP"
		portquota            = 20
		sessionquota         = 20
		groupsessionlimit    = 1000
	}
  
`

func TestAccLsntransportprofile_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLsntransportprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLsntransportprofile_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsntransportprofileExist("citrixadc_lsntransportprofile.tf_lsntransportprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_lsntransportprofile.tf_lsntransportprofile", "transportprofilename", "my_lsn_transportprofile"),
					resource.TestCheckResourceAttr("citrixadc_lsntransportprofile.tf_lsntransportprofile", "transportprotocol", "TCP"),
					resource.TestCheckResourceAttr("citrixadc_lsntransportprofile.tf_lsntransportprofile", "portquota", "10"),
					resource.TestCheckResourceAttr("citrixadc_lsntransportprofile.tf_lsntransportprofile", "sessionquota", "10"),
					resource.TestCheckResourceAttr("citrixadc_lsntransportprofile.tf_lsntransportprofile", "groupsessionlimit", "100"),
				),
			},
			{
				Config: testAccLsntransportprofile_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsntransportprofileExist("citrixadc_lsntransportprofile.tf_lsntransportprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_lsntransportprofile.tf_lsntransportprofile", "transportprofilename", "my_lsn_transportprofile"),
					resource.TestCheckResourceAttr("citrixadc_lsntransportprofile.tf_lsntransportprofile", "transportprotocol", "TCP"),
					resource.TestCheckResourceAttr("citrixadc_lsntransportprofile.tf_lsntransportprofile", "portquota", "20"),
					resource.TestCheckResourceAttr("citrixadc_lsntransportprofile.tf_lsntransportprofile", "sessionquota", "20"),
					resource.TestCheckResourceAttr("citrixadc_lsntransportprofile.tf_lsntransportprofile", "groupsessionlimit", "1000"),
				),
			},
		},
	})
}

func testAccCheckLsntransportprofileExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lsntransportprofile name is set")
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
		data, err := client.FindResource("lsntransportprofile", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("lsntransportprofile %s not found", n)
		}

		return nil
	}
}

func testAccCheckLsntransportprofileDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_lsntransportprofile" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("lsntransportprofile", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("lsntransportprofile %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccLsntransportprofile_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_lsntransportprofile.tf_lsntransportprofile"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLsntransportprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLsntransportprofile_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckLsntransportprofileExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Lsntransportprofile.Type(), "my_lsn_transportprofile"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccLsntransportprofile_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckLsntransportprofileExist(resAddr, nil)),
			},
		},
	})
}

func TestAccLsntransportprofile_import(t *testing.T) {
	const resAddr = "citrixadc_lsntransportprofile.tf_lsntransportprofile"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLsntransportprofileDestroy,
		Steps: []resource.TestStep{
			{Config: testAccLsntransportprofile_basic},
			{
				Config:                  testAccLsntransportprofile_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

const testAccLsntransportprofileDataSource_basic = `

	resource "citrixadc_lsntransportprofile" "tf_lsntransportprofile_ds" {
		transportprofilename = "my_lsn_transportprofile_ds"
		transportprotocol    = "TCP"
		portquota            = 10
		sessionquota         = 10
		groupsessionlimit    = 100
	}

	data "citrixadc_lsntransportprofile" "tf_lsntransportprofile_ds" {
		transportprofilename = citrixadc_lsntransportprofile.tf_lsntransportprofile_ds.transportprofilename
	}
`

func TestAccLsntransportprofile_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckLsntransportprofileDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccLsntransportprofile_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckLsntransportprofileExist("citrixadc_lsntransportprofile.tf_lsntransportprofile", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccLsntransportprofile_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckLsntransportprofileExist("citrixadc_lsntransportprofile.tf_lsntransportprofile", nil)),
			},
		},
	})
}

const testAccLsntransportprofile_unset_step1 = `
	resource "citrixadc_lsntransportprofile" "tf_unset" {
		transportprofilename = "tf_test_lsntransportprofile_unset"
		transportprotocol    = "TCP"
		finrsttimeout        = 45
		groupsessionlimit    = 500
		portpreserveparity   = "ENABLED"
		portpreserverange    = "ENABLED"
		portquota            = 500
		sessionquota         = 500
		sessiontimeout       = 300
		stuntimeout          = 900
		syncheck             = "DISABLED"
		synidletimeout       = 90
	}
`

const testAccLsntransportprofile_unset_step2 = `
	resource "citrixadc_lsntransportprofile" "tf_unset" {
		transportprofilename = "tf_test_lsntransportprofile_unset"
		transportprotocol    = "TCP"
		# All unset-eligible attributes removed from config -> the provider must
		# unset them (revert to NITRO defaults).
	}
`

func TestAccLsntransportprofile_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLsntransportprofileDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccLsntransportprofile_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsntransportprofileExist("citrixadc_lsntransportprofile.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_lsntransportprofile.tf_unset", "finrsttimeout", "45"),
					resource.TestCheckResourceAttr("citrixadc_lsntransportprofile.tf_unset", "groupsessionlimit", "500"),
					resource.TestCheckResourceAttr("citrixadc_lsntransportprofile.tf_unset", "portpreserveparity", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_lsntransportprofile.tf_unset", "portpreserverange", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_lsntransportprofile.tf_unset", "portquota", "500"),
					resource.TestCheckResourceAttr("citrixadc_lsntransportprofile.tf_unset", "sessionquota", "500"),
					resource.TestCheckResourceAttr("citrixadc_lsntransportprofile.tf_unset", "sessiontimeout", "300"),
					resource.TestCheckResourceAttr("citrixadc_lsntransportprofile.tf_unset", "stuntimeout", "900"),
					resource.TestCheckResourceAttr("citrixadc_lsntransportprofile.tf_unset", "syncheck", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_lsntransportprofile.tf_unset", "synidletimeout", "90"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from
				// the appliance) reverts to the documented NITRO defaults, and the
				// implicit post-apply plan must be empty.
				Config: testAccLsntransportprofile_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsntransportprofileExist("citrixadc_lsntransportprofile.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_lsntransportprofile.tf_unset", "finrsttimeout", "30"),
					resource.TestCheckResourceAttr("citrixadc_lsntransportprofile.tf_unset", "groupsessionlimit", "0"),
					resource.TestCheckResourceAttr("citrixadc_lsntransportprofile.tf_unset", "portpreserveparity", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_lsntransportprofile.tf_unset", "portpreserverange", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_lsntransportprofile.tf_unset", "portquota", "0"),
					resource.TestCheckResourceAttr("citrixadc_lsntransportprofile.tf_unset", "sessionquota", "0"),
					resource.TestCheckResourceAttr("citrixadc_lsntransportprofile.tf_unset", "sessiontimeout", "120"),
					resource.TestCheckNoResourceAttr("citrixadc_lsntransportprofile.tf_unset", "stuntimeout"),
					resource.TestCheckResourceAttr("citrixadc_lsntransportprofile.tf_unset", "syncheck", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_lsntransportprofile.tf_unset", "synidletimeout", "60"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckLsntransportprofileADCValue("tf_test_lsntransportprofile_unset", "syncheck", "ENABLED"),
					testAccCheckLsntransportprofileADCValue("tf_test_lsntransportprofile_unset", "sessiontimeout", "120"),
					testAccCheckLsntransportprofileADCValue("tf_test_lsntransportprofile_unset", "portquota", "0"),
				),
			},
		},
	})
}

// testAccCheckLsntransportprofileADCValue asserts an attribute's value directly
// on the appliance (not just in Terraform state), proving the unset actually
// reverted it.
func testAccCheckLsntransportprofileADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Lsntransportprofile.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("lsntransportprofile %s not found on appliance", name)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("lsntransportprofile %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func TestAccLsntransportprofileDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccLsntransportprofileDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_lsntransportprofile.tf_lsntransportprofile_ds", "transportprofilename", "my_lsn_transportprofile_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_lsntransportprofile.tf_lsntransportprofile_ds", "transportprotocol", "TCP"),
					resource.TestCheckResourceAttr("data.citrixadc_lsntransportprofile.tf_lsntransportprofile_ds", "portquota", "10"),
					resource.TestCheckResourceAttr("data.citrixadc_lsntransportprofile.tf_lsntransportprofile_ds", "sessionquota", "10"),
					resource.TestCheckResourceAttr("data.citrixadc_lsntransportprofile.tf_lsntransportprofile_ds", "groupsessionlimit", "100"),
				),
			},
		},
	})
}
