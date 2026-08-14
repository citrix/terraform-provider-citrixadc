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

const testAccNsratecontrol_add = `
	resource "citrixadc_nsratecontrol" "tf_nsratecontrol" {
		tcpthreshold    = 0
		udpthreshold    = 0
		icmpthreshold   = 100
		tcprstthreshold = 100
	}
`
const testAccNsratecontrol_update = `
	resource "citrixadc_nsratecontrol" "tf_nsratecontrol" {
		tcpthreshold    = 10
		udpthreshold    = 10
		icmpthreshold   = 100
		tcprstthreshold = 100
	}
`

func TestAccNsratecontrol_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccNsratecontrol_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsratecontrolExist("citrixadc_nsratecontrol.tf_nsratecontrol", nil),
					resource.TestCheckResourceAttr("citrixadc_nsratecontrol.tf_nsratecontrol", "tcpthreshold", "0"),
					resource.TestCheckResourceAttr("citrixadc_nsratecontrol.tf_nsratecontrol", "udpthreshold", "0"),
				),
			},
			{
				Config: testAccNsratecontrol_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsratecontrolExist("citrixadc_nsratecontrol.tf_nsratecontrol", nil),
					resource.TestCheckResourceAttr("citrixadc_nsratecontrol.tf_nsratecontrol", "tcpthreshold", "10"),
					resource.TestCheckResourceAttr("citrixadc_nsratecontrol.tf_nsratecontrol", "udpthreshold", "10"),
				),
			},
		},
	})
}

func testAccCheckNsratecontrolExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No nsratecontrol name is set")
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
		data, err := client.FindResource(service.Nsratecontrol.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("nsratecontrol %s not found", n)
		}

		return nil
	}
}

func TestAccNsratecontrol_import(t *testing.T) {
	const resAddr = "citrixadc_nsratecontrol.tf_nsratecontrol"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{Config: testAccNsratecontrol_add},
			{
				Config:                  testAccNsratecontrol_add,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

const testAccNsratecontrolDataSource_basic = `
	resource "citrixadc_nsratecontrol" "tf_nsratecontrol" {
		tcpthreshold    = 15
		udpthreshold    = 20
		icmpthreshold   = 105
		tcprstthreshold = 110
	}

	data "citrixadc_nsratecontrol" "tf_nsratecontrol" {
		depends_on = [citrixadc_nsratecontrol.tf_nsratecontrol]
	}
`

func TestAccNsratecontrol_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccNsratecontrol_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckNsratecontrolExist("citrixadc_nsratecontrol.tf_nsratecontrol", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccNsratecontrol_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckNsratecontrolExist("citrixadc_nsratecontrol.tf_nsratecontrol", nil)),
			},
		},
	})
}

// nsratecontrol unset test: step1 sets all four thresholds to non-default
// values; step2 removes them from config so the provider must unset them,
// reverting to the documented NITRO defaults (icmp/tcprst=100, tcp/udp=0).
const testAccNsratecontrol_unset_step1 = `
	resource "citrixadc_nsratecontrol" "tf_unset" {
		tcpthreshold    = 10
		udpthreshold    = 20
		icmpthreshold   = 105
		tcprstthreshold = 110
	}
`

const testAccNsratecontrol_unset_step2 = `
	resource "citrixadc_nsratecontrol" "tf_unset" {
		# All unset-eligible attributes removed from config -> the provider must
		# unset them (revert to NITRO defaults).
	}
`

func TestAccNsratecontrol_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccNsratecontrol_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsratecontrolExist("citrixadc_nsratecontrol.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_nsratecontrol.tf_unset", "tcpthreshold", "10"),
					resource.TestCheckResourceAttr("citrixadc_nsratecontrol.tf_unset", "udpthreshold", "20"),
					resource.TestCheckResourceAttr("citrixadc_nsratecontrol.tf_unset", "icmpthreshold", "105"),
					resource.TestCheckResourceAttr("citrixadc_nsratecontrol.tf_unset", "tcprstthreshold", "110"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from
				// the appliance) reverts to the documented NITRO defaults, and the
				// implicit post-apply plan must be empty.
				Config: testAccNsratecontrol_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsratecontrolExist("citrixadc_nsratecontrol.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_nsratecontrol.tf_unset", "tcpthreshold", "0"),
					resource.TestCheckResourceAttr("citrixadc_nsratecontrol.tf_unset", "udpthreshold", "0"),
					resource.TestCheckResourceAttr("citrixadc_nsratecontrol.tf_unset", "icmpthreshold", "100"),
					resource.TestCheckResourceAttr("citrixadc_nsratecontrol.tf_unset", "tcprstthreshold", "100"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckNsratecontrolADCValue("icmpthreshold", "100"),
					testAccCheckNsratecontrolADCValue("tcprstthreshold", "100"),
				),
			},
		},
	})
}

// testAccCheckNsratecontrolADCValue asserts an attribute's value directly on
// the appliance (not just in Terraform state), proving the unset actually
// reverted it.
func testAccCheckNsratecontrolADCValue(attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Nsratecontrol.Type(), "")
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("nsratecontrol not found on appliance")
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("nsratecontrol: appliance attr %q = %q, want %q (unset did not revert it)", attr, got, want)
		}
		return nil
	}
}

func TestAccNsratecontrolDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNsratecontrolDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_nsratecontrol.tf_nsratecontrol", "tcpthreshold", "15"),
					resource.TestCheckResourceAttr("data.citrixadc_nsratecontrol.tf_nsratecontrol", "udpthreshold", "20"),
					resource.TestCheckResourceAttr("data.citrixadc_nsratecontrol.tf_nsratecontrol", "icmpthreshold", "105"),
					resource.TestCheckResourceAttr("data.citrixadc_nsratecontrol.tf_nsratecontrol", "tcprstthreshold", "110"),
				),
			},
		},
	})
}
