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

const testAccNsdhcpparams_add = `
	resource "citrixadc_nsdhcpparams" "tf_nsdhcpparams" {
		dhcpclient = "ON"
		saveroute  = "ON"
	}
`
const testAccNsdhcpparams_update = `
	resource "citrixadc_nsdhcpparams" "tf_nsdhcpparams" {
		dhcpclient = "OFF"
		saveroute  = "OFF"
	}
`

func TestAccNsdhcpparams_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccNsdhcpparams_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsdhcpparamsExist("citrixadc_nsdhcpparams.tf_nsdhcpparams", nil),
					resource.TestCheckResourceAttr("citrixadc_nsdhcpparams.tf_nsdhcpparams", "dhcpclient", "ON"),
					resource.TestCheckResourceAttr("citrixadc_nsdhcpparams.tf_nsdhcpparams", "saveroute", "ON"),
				),
			},
			{
				Config: testAccNsdhcpparams_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsdhcpparamsExist("citrixadc_nsdhcpparams.tf_nsdhcpparams", nil),
					resource.TestCheckResourceAttr("citrixadc_nsdhcpparams.tf_nsdhcpparams", "dhcpclient", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_nsdhcpparams.tf_nsdhcpparams", "saveroute", "OFF"),
				),
			},
		},
	})
}

func TestAccNsdhcpparams_import(t *testing.T) {
	const resAddr = "citrixadc_nsdhcpparams.tf_nsdhcpparams"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{Config: testAccNsdhcpparams_add},
			{
				Config:                  testAccNsdhcpparams_add,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccCheckNsdhcpparamsExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No nsdhcpparams name is set")
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
		data, err := client.FindResource(service.Nsdhcpparams.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("nsdhcpparams %s not found", n)
		}

		return nil
	}
}

func TestAccNsdhcpparamsDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccNsdhcpparamsDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_nsdhcpparams.tf_nsdhcpparams", "dhcpclient", "ON"),
					resource.TestCheckResourceAttr("data.citrixadc_nsdhcpparams.tf_nsdhcpparams", "saveroute", "ON"),
				),
			},
		},
	})
}

func TestAccNsdhcpparams_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccNsdhcpparams_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckNsdhcpparamsExist("citrixadc_nsdhcpparams.tf_nsdhcpparams", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccNsdhcpparams_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckNsdhcpparamsExist("citrixadc_nsdhcpparams.tf_nsdhcpparams", nil)),
			},
		},
	})
}

const testAccNsdhcpparamsDataSource_basic = `

resource "citrixadc_nsdhcpparams" "tf_nsdhcpparams" {
	dhcpclient = "ON"
	saveroute  = "ON"
}

data "citrixadc_nsdhcpparams" "tf_nsdhcpparams" {
	depends_on = [citrixadc_nsdhcpparams.tf_nsdhcpparams]
}
`

// Unset test: step1 sets the unset-eligible attributes to non-default values
// ("ON"), step2 removes them so the provider must unset them (revert to the
// documented NITRO defaults, "OFF").
const testAccNsdhcpparams_unset_step1 = `
	resource "citrixadc_nsdhcpparams" "tf_unset" {
		dhcpclient = "ON"
		saveroute  = "ON"
	}
`

const testAccNsdhcpparams_unset_step2 = `
	resource "citrixadc_nsdhcpparams" "tf_unset" {
		# All unset-eligible attributes removed from config -> the provider must
		# unset them (revert to NITRO defaults, "OFF").
	}
`

func TestAccNsdhcpparams_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccNsdhcpparams_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsdhcpparamsExist("citrixadc_nsdhcpparams.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_nsdhcpparams.tf_unset", "dhcpclient", "ON"),
					resource.TestCheckResourceAttr("citrixadc_nsdhcpparams.tf_unset", "saveroute", "ON"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from
				// the appliance) reverts to the documented NITRO defaults, and the
				// implicit post-apply plan must be empty.
				Config: testAccNsdhcpparams_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsdhcpparamsExist("citrixadc_nsdhcpparams.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_nsdhcpparams.tf_unset", "dhcpclient", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_nsdhcpparams.tf_unset", "saveroute", "OFF"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckNsdhcpparamsADCValue("dhcpclient", "OFF"),
					testAccCheckNsdhcpparamsADCValue("saveroute", "OFF"),
				),
			},
		},
	})
}

// testAccCheckNsdhcpparamsADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted
// it.
func testAccCheckNsdhcpparamsADCValue(attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Nsdhcpparams.Type(), "")
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("nsdhcpparams not found on appliance")
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("nsdhcpparams: appliance attr %q = %q, want %q (unset did not revert it)", attr, got, want)
		}
		return nil
	}
}
