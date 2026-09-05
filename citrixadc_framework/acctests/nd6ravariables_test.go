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

const testAccNd6ravariables_basic = `


	resource "citrixadc_nd6ravariables" "tf_nd6ravariables" {
		vlan                     = 1
		ceaserouteradv           = "NO"
		onlyunicastrtadvresponse = "NO"
		srclinklayeraddroption   = "NO"
	}
`
const testAccNd6ravariables_update = `


	resource "citrixadc_nd6ravariables" "tf_nd6ravariables" {
		vlan                     = 1
		ceaserouteradv           = "YES"
		onlyunicastrtadvresponse = "YES"
		srclinklayeraddroption   = "YES"
	}
`

func TestAccNd6ravariables_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccNd6ravariables_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNd6ravariablesExist("citrixadc_nd6ravariables.tf_nd6ravariables", nil),
					resource.TestCheckResourceAttr("citrixadc_nd6ravariables.tf_nd6ravariables", "ceaserouteradv", "NO"),
					resource.TestCheckResourceAttr("citrixadc_nd6ravariables.tf_nd6ravariables", "onlyunicastrtadvresponse", "NO"),
					resource.TestCheckResourceAttr("citrixadc_nd6ravariables.tf_nd6ravariables", "srclinklayeraddroption", "NO"),
				),
			},
			{
				Config: testAccNd6ravariables_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNd6ravariablesExist("citrixadc_nd6ravariables.tf_nd6ravariables", nil),
					resource.TestCheckResourceAttr("citrixadc_nd6ravariables.tf_nd6ravariables", "ceaserouteradv", "YES"),
					resource.TestCheckResourceAttr("citrixadc_nd6ravariables.tf_nd6ravariables", "onlyunicastrtadvresponse", "YES"),
					resource.TestCheckResourceAttr("citrixadc_nd6ravariables.tf_nd6ravariables", "srclinklayeraddroption", "YES"),
				),
			},
		},
	})
}

func TestAccNd6ravariables_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.0.0"},
				},
				Config: testAccNd6ravariables_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckNd6ravariablesExist("citrixadc_nd6ravariables.tf_nd6ravariables", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccNd6ravariables_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckNd6ravariablesExist("citrixadc_nd6ravariables.tf_nd6ravariables", nil)),
			},
		},
	})
}

func testAccCheckNd6ravariablesExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No nd6ravariables name is set")
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
		data, err := client.FindResource(service.Nd6ravariables.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("nd6ravariables %s not found", n)
		}

		return nil
	}
}

func TestAccNd6ravariables_import(t *testing.T) {
	const resAddr = "citrixadc_nd6ravariables.tf_nd6ravariables"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{Config: testAccNd6ravariables_basic},
			{
				Config:                  testAccNd6ravariables_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

const testAccNd6ravariablesDataSource_basic = `

	resource "citrixadc_nd6ravariables" "tf_nd6ravariables" {
		vlan                     = 1
		ceaserouteradv           = "NO"
		onlyunicastrtadvresponse = "NO"
		srclinklayeraddroption   = "NO"
	}

	data "citrixadc_nd6ravariables" "tf_nd6ravariables_data" {
		vlan = citrixadc_nd6ravariables.tf_nd6ravariables.vlan
	}
`

// step1 sets every unset-eligible attribute to a valid non-default value
// (respecting the RA interval constraints: minrtadvinterval <= 0.75 *
// maxrtadvinterval and defaultlifetime >= maxrtadvinterval).
const testAccNd6ravariables_unset_step1 = `
	resource "citrixadc_nd6ravariables" "tf_unset" {
		vlan                     = 1
		ceaserouteradv           = "YES"
		srclinklayeraddroption   = "NO"
		onlyunicastrtadvresponse = "YES"
		managedaddrconfig        = "YES"
		otheraddrconfig          = "YES"
		currhoplimit             = 32
		maxrtadvinterval         = 700
		minrtadvinterval         = 300
		linkmtu                  = 1400
		reachabletime            = 1000
		retranstime              = 500
		defaultlifetime          = 2000
	}
`

// step2 removes all unset-eligible attributes (key only) -> the provider must
// unset them (revert to NITRO defaults).
const testAccNd6ravariables_unset_step2 = `
	resource "citrixadc_nd6ravariables" "tf_unset" {
		vlan = 1
	}
`

func TestAccNd6ravariables_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccNd6ravariables_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNd6ravariablesExist("citrixadc_nd6ravariables.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_nd6ravariables.tf_unset", "ceaserouteradv", "YES"),
					resource.TestCheckResourceAttr("citrixadc_nd6ravariables.tf_unset", "srclinklayeraddroption", "NO"),
					resource.TestCheckResourceAttr("citrixadc_nd6ravariables.tf_unset", "onlyunicastrtadvresponse", "YES"),
					resource.TestCheckResourceAttr("citrixadc_nd6ravariables.tf_unset", "managedaddrconfig", "YES"),
					resource.TestCheckResourceAttr("citrixadc_nd6ravariables.tf_unset", "otheraddrconfig", "YES"),
					resource.TestCheckResourceAttr("citrixadc_nd6ravariables.tf_unset", "currhoplimit", "32"),
					resource.TestCheckResourceAttr("citrixadc_nd6ravariables.tf_unset", "maxrtadvinterval", "700"),
					resource.TestCheckResourceAttr("citrixadc_nd6ravariables.tf_unset", "minrtadvinterval", "300"),
					resource.TestCheckResourceAttr("citrixadc_nd6ravariables.tf_unset", "linkmtu", "1400"),
					resource.TestCheckResourceAttr("citrixadc_nd6ravariables.tf_unset", "reachabletime", "1000"),
					resource.TestCheckResourceAttr("citrixadc_nd6ravariables.tf_unset", "retranstime", "500"),
					resource.TestCheckResourceAttr("citrixadc_nd6ravariables.tf_unset", "defaultlifetime", "2000"),
				),
			},
			{
				Config: testAccNd6ravariables_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNd6ravariablesExist("citrixadc_nd6ravariables.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_nd6ravariables.tf_unset", "ceaserouteradv", "NO"),
					resource.TestCheckResourceAttr("citrixadc_nd6ravariables.tf_unset", "srclinklayeraddroption", "YES"),
					resource.TestCheckResourceAttr("citrixadc_nd6ravariables.tf_unset", "onlyunicastrtadvresponse", "NO"),
					resource.TestCheckResourceAttr("citrixadc_nd6ravariables.tf_unset", "managedaddrconfig", "NO"),
					resource.TestCheckResourceAttr("citrixadc_nd6ravariables.tf_unset", "otheraddrconfig", "NO"),
					resource.TestCheckResourceAttr("citrixadc_nd6ravariables.tf_unset", "currhoplimit", "64"),
					resource.TestCheckResourceAttr("citrixadc_nd6ravariables.tf_unset", "maxrtadvinterval", "600"),
					resource.TestCheckResourceAttr("citrixadc_nd6ravariables.tf_unset", "minrtadvinterval", "198"),
					resource.TestCheckResourceAttr("citrixadc_nd6ravariables.tf_unset", "linkmtu", "0"),
					resource.TestCheckResourceAttr("citrixadc_nd6ravariables.tf_unset", "reachabletime", "0"),
					resource.TestCheckResourceAttr("citrixadc_nd6ravariables.tf_unset", "retranstime", "0"),
					resource.TestCheckResourceAttr("citrixadc_nd6ravariables.tf_unset", "defaultlifetime", "1800"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckNd6ravariablesADCValue("1", "ceaserouteradv", "NO"),
					testAccCheckNd6ravariablesADCValue("1", "currhoplimit", "64"),
					testAccCheckNd6ravariablesADCValue("1", "srclinklayeraddroption", "YES"),
				),
			},
		},
	})
}

// testAccCheckNd6ravariablesADCValue asserts an attribute's value directly on
// the appliance (not just in Terraform state), proving the unset actually
// reverted it.
func testAccCheckNd6ravariablesADCValue(vlan, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Nd6ravariables.Type(), vlan)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("nd6ravariables vlan %s not found on appliance", vlan)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("nd6ravariables vlan %s: appliance attr %q = %q, want %q (unset did not revert it)", vlan, attr, got, want)
		}
		return nil
	}
}

func TestAccNd6ravariablesDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccNd6ravariablesDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_nd6ravariables.tf_nd6ravariables_data", "vlan", "1"),
					resource.TestCheckResourceAttr("data.citrixadc_nd6ravariables.tf_nd6ravariables_data", "ceaserouteradv", "NO"),
					resource.TestCheckResourceAttr("data.citrixadc_nd6ravariables.tf_nd6ravariables_data", "onlyunicastrtadvresponse", "NO"),
					resource.TestCheckResourceAttr("data.citrixadc_nd6ravariables.tf_nd6ravariables_data", "srclinklayeraddroption", "NO"),
					// Universal runtime-binding proof.
					resource.TestCheckResourceAttrSet("data.citrixadc_nd6ravariables.tf_nd6ravariables_data", "id"),
				),
			},
		},
	})
}
