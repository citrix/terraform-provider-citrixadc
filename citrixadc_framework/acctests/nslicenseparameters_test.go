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

const testAccNslicenseparameters_add = `

	resource "citrixadc_nslicenseparameters" "tf_nslicenseparameters" {
		alert1gracetimeout = 8
		alert2gracetimeout = 200
		licenseexpiryalerttime = 30
		inventoryrefreshinterval = 50
		heartbeatinterval = 45
	}
`
const testAccNslicenseparameters_update = `

	resource "citrixadc_nslicenseparameters" "tf_nslicenseparameters" {
		alert1gracetimeout = 6
		alert2gracetimeout = 240
		licenseexpiryalerttime = 40
		inventoryrefreshinterval = 60
		heartbeatinterval = 55
	}
`

func TestAccNslicenseparameters_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccNslicenseparameters_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNslicenseparametersExist("citrixadc_nslicenseparameters.tf_nslicenseparameters", nil),
					resource.TestCheckResourceAttr("citrixadc_nslicenseparameters.tf_nslicenseparameters", "alert1gracetimeout", "8"),
					resource.TestCheckResourceAttr("citrixadc_nslicenseparameters.tf_nslicenseparameters", "alert2gracetimeout", "200"),
					resource.TestCheckResourceAttr("citrixadc_nslicenseparameters.tf_nslicenseparameters", "licenseexpiryalerttime", "30"),
					resource.TestCheckResourceAttr("citrixadc_nslicenseparameters.tf_nslicenseparameters", "inventoryrefreshinterval", "50"),
					resource.TestCheckResourceAttr("citrixadc_nslicenseparameters.tf_nslicenseparameters", "heartbeatinterval", "45"),
				),
			},
			{
				Config: testAccNslicenseparameters_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNslicenseparametersExist("citrixadc_nslicenseparameters.tf_nslicenseparameters", nil),
					resource.TestCheckResourceAttr("citrixadc_nslicenseparameters.tf_nslicenseparameters", "alert1gracetimeout", "6"),
					resource.TestCheckResourceAttr("citrixadc_nslicenseparameters.tf_nslicenseparameters", "alert2gracetimeout", "240"),
					resource.TestCheckResourceAttr("citrixadc_nslicenseparameters.tf_nslicenseparameters", "licenseexpiryalerttime", "40"),
					resource.TestCheckResourceAttr("citrixadc_nslicenseparameters.tf_nslicenseparameters", "inventoryrefreshinterval", "60"),
					resource.TestCheckResourceAttr("citrixadc_nslicenseparameters.tf_nslicenseparameters", "heartbeatinterval", "55"),
				),
			},
		},
	})
}

func testAccCheckNslicenseparametersExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No nslicenseparameters name is set")
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
		data, err := client.FindResource("nslicenseparameters", "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("nslicenseparameters %s not found", n)
		}

		return nil
	}
}

func TestAccNslicenseparameters_import(t *testing.T) {
	const resAddr = "citrixadc_nslicenseparameters.tf_nslicenseparameters"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{Config: testAccNslicenseparameters_add},
			{
				Config:                  testAccNslicenseparameters_add,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccNslicenseparameters_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccNslicenseparameters_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckNslicenseparametersExist("citrixadc_nslicenseparameters.tf_nslicenseparameters", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccNslicenseparameters_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckNslicenseparametersExist("citrixadc_nslicenseparameters.tf_nslicenseparameters", nil)),
			},
		},
	})
}

const testAccNslicenseparametersDataSource_basic = `

	resource "citrixadc_nslicenseparameters" "tf_nslicenseparameters" {
		alert1gracetimeout       = 10
		alert2gracetimeout       = 250
		heartbeatinterval        = 270
		inventoryrefreshinterval = 400
		licenseexpiryalerttime   = 40
	}
	
	data "citrixadc_nslicenseparameters" "tf_nslicenseparameters" {
		depends_on = [citrixadc_nslicenseparameters.tf_nslicenseparameters]
	}
`

// nslicenseparameters is a singleton config object. All five read/write
// attributes support NITRO unset and revert to their documented defaults:
// alert1gracetimeout=6, alert2gracetimeout=240, heartbeatinterval=280,
// inventoryrefreshinterval=360, licenseexpiryalerttime=30.
const testAccNslicenseparameters_unset_step1 = `
	resource "citrixadc_nslicenseparameters" "tf_unset" {
		alert1gracetimeout       = 12
		alert2gracetimeout       = 300
		heartbeatinterval        = 200
		inventoryrefreshinterval = 500
		licenseexpiryalerttime   = 60
	}
`

const testAccNslicenseparameters_unset_step2 = `
	resource "citrixadc_nslicenseparameters" "tf_unset" {
		# All unset-eligible attributes removed from config -> the provider must
		# unset them (revert to NITRO defaults).
	}
`

func TestAccNslicenseparameters_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccNslicenseparameters_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNslicenseparametersExist("citrixadc_nslicenseparameters.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_nslicenseparameters.tf_unset", "alert1gracetimeout", "12"),
					resource.TestCheckResourceAttr("citrixadc_nslicenseparameters.tf_unset", "alert2gracetimeout", "300"),
					resource.TestCheckResourceAttr("citrixadc_nslicenseparameters.tf_unset", "heartbeatinterval", "200"),
					resource.TestCheckResourceAttr("citrixadc_nslicenseparameters.tf_unset", "inventoryrefreshinterval", "500"),
					resource.TestCheckResourceAttr("citrixadc_nslicenseparameters.tf_unset", "licenseexpiryalerttime", "60"),
				),
			},
			{
				// Removing the attributes must unset them: state reverts to the
				// documented NITRO defaults and the implicit post-apply plan is empty.
				Config: testAccNslicenseparameters_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNslicenseparametersExist("citrixadc_nslicenseparameters.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_nslicenseparameters.tf_unset", "alert1gracetimeout", "6"),
					resource.TestCheckResourceAttr("citrixadc_nslicenseparameters.tf_unset", "alert2gracetimeout", "240"),
					resource.TestCheckResourceAttr("citrixadc_nslicenseparameters.tf_unset", "heartbeatinterval", "280"),
					resource.TestCheckResourceAttr("citrixadc_nslicenseparameters.tf_unset", "inventoryrefreshinterval", "360"),
					resource.TestCheckResourceAttr("citrixadc_nslicenseparameters.tf_unset", "licenseexpiryalerttime", "30"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckNslicenseparametersADCValue("alert1gracetimeout", "6"),
					testAccCheckNslicenseparametersADCValue("inventoryrefreshinterval", "360"),
				),
			},
		},
	})
}

// testAccCheckNslicenseparametersADCValue asserts an attribute's value directly
// on the appliance (not just in Terraform state), proving the unset reverted it.
func testAccCheckNslicenseparametersADCValue(attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Nslicenseparameters.Type(), "")
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("nslicenseparameters not found on appliance")
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("nslicenseparameters: appliance attr %q = %q, want %q (unset did not revert it)", attr, got, want)
		}
		return nil
	}
}

func TestAccNslicenseparametersDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccNslicenseparametersDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_nslicenseparameters.tf_nslicenseparameters", "alert1gracetimeout", "10"),
					resource.TestCheckResourceAttr("data.citrixadc_nslicenseparameters.tf_nslicenseparameters", "alert2gracetimeout", "250"),
					resource.TestCheckResourceAttr("data.citrixadc_nslicenseparameters.tf_nslicenseparameters", "heartbeatinterval", "270"),
					resource.TestCheckResourceAttr("data.citrixadc_nslicenseparameters.tf_nslicenseparameters", "inventoryrefreshinterval", "400"),
					resource.TestCheckResourceAttr("data.citrixadc_nslicenseparameters.tf_nslicenseparameters", "licenseexpiryalerttime", "40"),
				),
			},
		},
	})
}
