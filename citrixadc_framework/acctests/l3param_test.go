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

const testAccL3param_basic = `
	resource "citrixadc_l3param" "tf_l3param" {
		srcnat               = "DISABLED"
		icmpgenratethreshold = 150
		overridernat         = "DISABLED"
		dropdfflag           = "DISABLED"
		implicitpbr		 = "DISABLED"
	}
  
`
const testAccL3param_update = `
	resource "citrixadc_l3param" "tf_l3param" {
		srcnat               = "ENABLED"
		icmpgenratethreshold = 200
		overridernat         = "ENABLED"
		dropdfflag           = "ENABLED"
		implicitpbr		 = "ENABLED"
	}
  
`

func TestAccL3param_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccL3param_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckL3paramExist("citrixadc_l3param.tf_l3param", nil),
					resource.TestCheckResourceAttr("citrixadc_l3param.tf_l3param", "srcnat", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_l3param.tf_l3param", "icmpgenratethreshold", "150"),
					resource.TestCheckResourceAttr("citrixadc_l3param.tf_l3param", "overridernat", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_l3param.tf_l3param", "dropdfflag", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_l3param.tf_l3param", "implicitpbr", "DISABLED"),
				),
			},
			{
				Config: testAccL3param_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckL3paramExist("citrixadc_l3param.tf_l3param", nil),
					resource.TestCheckResourceAttr("citrixadc_l3param.tf_l3param", "srcnat", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_l3param.tf_l3param", "icmpgenratethreshold", "200"),
					resource.TestCheckResourceAttr("citrixadc_l3param.tf_l3param", "overridernat", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_l3param.tf_l3param", "dropdfflag", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_l3param.tf_l3param", "implicitpbr", "ENABLED"),
				),
			},
		},
	})
}

func TestAccL3param_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.0.0"},
				},
				Config: testAccL3param_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckL3paramExist("citrixadc_l3param.tf_l3param", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccL3param_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckL3paramExist("citrixadc_l3param.tf_l3param", nil)),
			},
		},
	})
}

func testAccCheckL3paramExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No l3param name is set")
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
		data, err := client.FindResource(service.L3param.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("l3param %s not found", n)
		}

		return nil
	}
}

func TestAccL3param_import(t *testing.T) {
	const resAddr = "citrixadc_l3param.tf_l3param"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{Config: testAccL3param_basic},
			{
				Config:                  testAccL3param_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

const testAccL3paramDataSource_basic = `

	resource "citrixadc_l3param" "tf_l3param" {
		srcnat               = "DISABLED"
		icmpgenratethreshold = 150
		overridernat         = "DISABLED"
		dropdfflag           = "DISABLED"
		implicitpbr          = "DISABLED"
	}

	data "citrixadc_l3param" "tf_l3param" {
		depends_on = [citrixadc_l3param.tf_l3param]
	}
`

// l3param is a singleton global config. step1 sets every unset-eligible
// attribute to a valid NON-default value; step2 removes them all from config so
// the provider must unset them, reverting the appliance to the documented NITRO
// defaults.
const testAccL3param_unset_step1 = `
	resource "citrixadc_l3param" "tf_unset" {
		acllogtime           = 6000
		allowclasseipv4      = "ENABLED"
		dropdfflag           = "ENABLED"
		dropipfragments      = "ENABLED"
		externalloopback     = "ENABLED"
		forwardicmpfragments = "ENABLED"
		icmpgenratethreshold = 200
		implicitaclallow     = "DISABLED"
		implicitpbr          = "ENABLED"
		miproundrobin        = "DISABLED"
		overridernat         = "ENABLED"
		srcnat               = "DISABLED"
		tnlpmtuwoconn        = "DISABLED"
		usipserverstraypkt   = "ENABLED"
	}
`

const testAccL3param_unset_step2 = `
	resource "citrixadc_l3param" "tf_unset" {
		# All unset-eligible attributes removed from config -> the provider must
		# unset them (revert to NITRO defaults).
	}
`

func TestAccL3param_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccL3param_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckL3paramExist("citrixadc_l3param.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_l3param.tf_unset", "acllogtime", "6000"),
					resource.TestCheckResourceAttr("citrixadc_l3param.tf_unset", "allowclasseipv4", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_l3param.tf_unset", "dropdfflag", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_l3param.tf_unset", "dropipfragments", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_l3param.tf_unset", "externalloopback", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_l3param.tf_unset", "forwardicmpfragments", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_l3param.tf_unset", "icmpgenratethreshold", "200"),
					resource.TestCheckResourceAttr("citrixadc_l3param.tf_unset", "implicitaclallow", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_l3param.tf_unset", "implicitpbr", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_l3param.tf_unset", "miproundrobin", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_l3param.tf_unset", "overridernat", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_l3param.tf_unset", "srcnat", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_l3param.tf_unset", "tnlpmtuwoconn", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_l3param.tf_unset", "usipserverstraypkt", "ENABLED"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from the
				// appliance) reverts to the documented NITRO defaults, and the implicit
				// post-apply plan must be empty.
				Config: testAccL3param_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckL3paramExist("citrixadc_l3param.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_l3param.tf_unset", "acllogtime", "5000"),
					resource.TestCheckResourceAttr("citrixadc_l3param.tf_unset", "allowclasseipv4", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_l3param.tf_unset", "dropdfflag", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_l3param.tf_unset", "dropipfragments", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_l3param.tf_unset", "externalloopback", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_l3param.tf_unset", "forwardicmpfragments", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_l3param.tf_unset", "icmpgenratethreshold", "100"),
					resource.TestCheckResourceAttr("citrixadc_l3param.tf_unset", "implicitaclallow", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_l3param.tf_unset", "implicitpbr", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_l3param.tf_unset", "miproundrobin", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_l3param.tf_unset", "overridernat", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_l3param.tf_unset", "srcnat", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_l3param.tf_unset", "tnlpmtuwoconn", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_l3param.tf_unset", "usipserverstraypkt", "DISABLED"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckL3paramADCValue("srcnat", "ENABLED"),
					testAccCheckL3paramADCValue("acllogtime", "5000"),
					testAccCheckL3paramADCValue("miproundrobin", "ENABLED"),
					testAccCheckL3paramADCValue("implicitaclallow", "ENABLED"),
				),
			},
		},
	})
}

// testAccCheckL3paramADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted
// it. l3param is a singleton, so no name key is needed.
func testAccCheckL3paramADCValue(attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.L3param.Type(), "")
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("l3param not found on appliance")
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("l3param: appliance attr %q = %q, want %q (unset did not revert it)", attr, got, want)
		}
		return nil
	}
}

func TestAccL3paramDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccL3paramDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_l3param.tf_l3param", "srcnat", "DISABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_l3param.tf_l3param", "icmpgenratethreshold", "150"),
					resource.TestCheckResourceAttr("data.citrixadc_l3param.tf_l3param", "overridernat", "DISABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_l3param.tf_l3param", "dropdfflag", "DISABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_l3param.tf_l3param", "implicitpbr", "DISABLED"),
				),
			},
		},
	})
}
