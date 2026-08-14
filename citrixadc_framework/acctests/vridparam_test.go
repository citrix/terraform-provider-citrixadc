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

const testAccVridparam_add = `

	resource "citrixadc_vridparam" "tf_vridparam" {
		sendtomaster  = "ENABLED"
		hellointerval = 400
		deadinterval  = 4
	}
`
const testAccVridparam_update = `

	resource "citrixadc_vridparam" "tf_vridparam" {
		sendtomaster  = "DISABLED"
		hellointerval = 1000
		deadinterval  = 3
	}
`

func TestAccVridparam_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccVridparam_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVridparamExist("citrixadc_vridparam.tf_vridparam", nil),
					resource.TestCheckResourceAttr("citrixadc_vridparam.tf_vridparam", "sendtomaster", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_vridparam.tf_vridparam", "hellointerval", "400"),
					resource.TestCheckResourceAttr("citrixadc_vridparam.tf_vridparam", "deadinterval", "4"),
				),
			},
			{
				Config: testAccVridparam_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVridparamExist("citrixadc_vridparam.tf_vridparam", nil),
					resource.TestCheckResourceAttr("citrixadc_vridparam.tf_vridparam", "sendtomaster", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_vridparam.tf_vridparam", "hellointerval", "1000"),
					resource.TestCheckResourceAttr("citrixadc_vridparam.tf_vridparam", "deadinterval", "3"),
				),
			},
		},
	})
}

func TestAccVridparam_import(t *testing.T) {
	const resAddr = "citrixadc_vridparam.tf_vridparam"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{Config: testAccVridparam_add},
			{
				Config:                  testAccVridparam_add,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccCheckVridparamExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No vridparam name is set")
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
		data, err := client.FindResource(service.Vridparam.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("vridparam %s not found", n)
		}

		return nil
	}
}

func TestAccVridparam_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccVridparam_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckVridparamExist("citrixadc_vridparam.tf_vridparam", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccVridparam_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckVridparamExist("citrixadc_vridparam.tf_vridparam", nil)),
			},
		},
	})
}

// vridparam is a singleton. Step 1 sets all three unset-eligible attributes to
// non-default values; step 2 removes them so the provider unsets them and the
// appliance reverts to the documented NITRO defaults
// (sendtomaster=DISABLED, hellointerval=1000, deadinterval=3).
const testAccVridparam_unset_step1 = `

	resource "citrixadc_vridparam" "tf_unset" {
		sendtomaster  = "ENABLED"
		hellointerval = 400
		deadinterval  = 4
	}
`

const testAccVridparam_unset_step2 = `

	resource "citrixadc_vridparam" "tf_unset" {
		# All unset-eligible attributes removed from config -> the provider must
		# unset them (revert to NITRO defaults).
	}
`

func TestAccVridparam_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccVridparam_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVridparamExist("citrixadc_vridparam.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_vridparam.tf_unset", "sendtomaster", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_vridparam.tf_unset", "hellointerval", "400"),
					resource.TestCheckResourceAttr("citrixadc_vridparam.tf_unset", "deadinterval", "4"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from
				// the appliance) reverts to the documented NITRO defaults, and the
				// implicit post-apply plan must be empty.
				Config: testAccVridparam_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVridparamExist("citrixadc_vridparam.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_vridparam.tf_unset", "sendtomaster", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_vridparam.tf_unset", "hellointerval", "1000"),
					resource.TestCheckResourceAttr("citrixadc_vridparam.tf_unset", "deadinterval", "3"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckVridparamADCValue("sendtomaster", "DISABLED"),
					testAccCheckVridparamADCValue("hellointerval", "1000"),
					testAccCheckVridparamADCValue("deadinterval", "3"),
				),
			},
		},
	})
}

// testAccCheckVridparamADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted
// it. vridparam is a singleton, so the resource is fetched with an empty name.
func testAccCheckVridparamADCValue(attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Vridparam.Type(), "")
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("vridparam not found on appliance")
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("vridparam: appliance attr %q = %q, want %q (unset did not revert it)", attr, got, want)
		}
		return nil
	}
}

const testAccVridparamDataSource_basic = `

	resource "citrixadc_vridparam" "tf_vridparam" {
		sendtomaster  = "ENABLED"
		hellointerval = 400
		deadinterval  = 4
	}

data "citrixadc_vridparam" "tf_vridparam" {
	depends_on = [citrixadc_vridparam.tf_vridparam]
}
`

func TestAccVridparamDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccVridparamDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_vridparam.tf_vridparam", "sendtomaster", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_vridparam.tf_vridparam", "hellointerval", "400"),
					resource.TestCheckResourceAttr("data.citrixadc_vridparam.tf_vridparam", "deadinterval", "4"),
				),
			},
		},
	})
}
