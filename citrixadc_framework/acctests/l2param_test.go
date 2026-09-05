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

const testAccL2param_basic = `


	resource "citrixadc_l2param" "tf_l2param" {
		mbfpeermacupdate   = 20
		maxbridgecollision = 30
		bdggrpproxyarp     = "DISABLED"
	}
`

const testAccL2param_update = `


	resource "citrixadc_l2param" "tf_l2param" {
		mbfpeermacupdate   = 30
		maxbridgecollision = 40
		bdggrpproxyarp     = "ENABLED"
	}
`

func TestAccL2param_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccL2param_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckL2paramExist("citrixadc_l2param.tf_l2param", nil),
					resource.TestCheckResourceAttr("citrixadc_l2param.tf_l2param", "mbfpeermacupdate", "20"),
					resource.TestCheckResourceAttr("citrixadc_l2param.tf_l2param", "maxbridgecollision", "30"),
					resource.TestCheckResourceAttr("citrixadc_l2param.tf_l2param", "bdggrpproxyarp", "DISABLED"),
				),
			},
			{
				Config: testAccL2param_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckL2paramExist("citrixadc_l2param.tf_l2param", nil),
					resource.TestCheckResourceAttr("citrixadc_l2param.tf_l2param", "mbfpeermacupdate", "30"),
					resource.TestCheckResourceAttr("citrixadc_l2param.tf_l2param", "maxbridgecollision", "40"),
					resource.TestCheckResourceAttr("citrixadc_l2param.tf_l2param", "bdggrpproxyarp", "ENABLED"),
				),
			},
		},
	})
}

func TestAccL2param_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.0.0"},
				},
				Config: testAccL2param_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckL2paramExist("citrixadc_l2param.tf_l2param", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccL2param_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckL2paramExist("citrixadc_l2param.tf_l2param", nil)),
			},
		},
	})
}

func testAccCheckL2paramExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No l2param name is set")
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
		data, err := client.FindResource(service.L2param.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("l2param %s not found", n)
		}

		return nil
	}
}

func TestAccL2param_import(t *testing.T) {
	const resAddr = "citrixadc_l2param.tf_l2param"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{Config: testAccL2param_basic},
			{
				Config:                  testAccL2param_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

// l2param is a singleton global-config resource. Step 1 sets every
// unset-eligible attribute to a valid non-default value; step 2 removes them
// all from config so the provider must unset them (revert to NITRO defaults).
const testAccL2param_unset_step1 = `
	resource "citrixadc_l2param" "tf_unset" {
		bdggrpproxyarp         = "DISABLED"
		bdgsetting             = "ENABLED"
		bridgeagetimeout       = 200
		garponvridintf         = "DISABLED"
		garpreply              = "ENABLED"
		macmodefwdmypkt        = "ENABLED"
		maxbridgecollision     = 30
		mbfinstlearning        = "ENABLED"
		mbfpeermacupdate       = 20
		proxyarp               = "DISABLED"
		returntoethernetsender = "ENABLED"
		rstintfonhafo          = "ENABLED"
		skipproxyingbsdtraffic = "ENABLED"
		stopmacmoveupdate      = "ENABLED"
		usemymac               = "ENABLED"
		usenetprofilebsdtraffic = "ENABLED"
	}
`

const testAccL2param_unset_step2 = `
	resource "citrixadc_l2param" "tf_unset" {
		# All unset-eligible attributes removed from config -> the provider must
		# unset them (revert to NITRO defaults).
	}
`

func TestAccL2param_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccL2param_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckL2paramExist("citrixadc_l2param.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_l2param.tf_unset", "bdggrpproxyarp", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_l2param.tf_unset", "bdgsetting", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_l2param.tf_unset", "bridgeagetimeout", "200"),
					resource.TestCheckResourceAttr("citrixadc_l2param.tf_unset", "garponvridintf", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_l2param.tf_unset", "garpreply", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_l2param.tf_unset", "macmodefwdmypkt", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_l2param.tf_unset", "maxbridgecollision", "30"),
					resource.TestCheckResourceAttr("citrixadc_l2param.tf_unset", "mbfinstlearning", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_l2param.tf_unset", "mbfpeermacupdate", "20"),
					resource.TestCheckResourceAttr("citrixadc_l2param.tf_unset", "proxyarp", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_l2param.tf_unset", "returntoethernetsender", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_l2param.tf_unset", "rstintfonhafo", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_l2param.tf_unset", "skipproxyingbsdtraffic", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_l2param.tf_unset", "stopmacmoveupdate", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_l2param.tf_unset", "usemymac", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_l2param.tf_unset", "usenetprofilebsdtraffic", "ENABLED"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from
				// the appliance) reverts to the documented NITRO defaults, and the
				// implicit post-apply plan must be empty.
				Config: testAccL2param_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckL2paramExist("citrixadc_l2param.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_l2param.tf_unset", "bdggrpproxyarp", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_l2param.tf_unset", "bdgsetting", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_l2param.tf_unset", "bridgeagetimeout", "300"),
					resource.TestCheckResourceAttr("citrixadc_l2param.tf_unset", "garponvridintf", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_l2param.tf_unset", "garpreply", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_l2param.tf_unset", "macmodefwdmypkt", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_l2param.tf_unset", "maxbridgecollision", "20"),
					resource.TestCheckResourceAttr("citrixadc_l2param.tf_unset", "mbfinstlearning", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_l2param.tf_unset", "mbfpeermacupdate", "10"),
					resource.TestCheckResourceAttr("citrixadc_l2param.tf_unset", "proxyarp", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_l2param.tf_unset", "returntoethernetsender", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_l2param.tf_unset", "rstintfonhafo", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_l2param.tf_unset", "skipproxyingbsdtraffic", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_l2param.tf_unset", "stopmacmoveupdate", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_l2param.tf_unset", "usemymac", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_l2param.tf_unset", "usenetprofilebsdtraffic", "DISABLED"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckL2paramADCValue("bdggrpproxyarp", "ENABLED"),
					testAccCheckL2paramADCValue("bridgeagetimeout", "300"),
					testAccCheckL2paramADCValue("usemymac", "DISABLED"),
				),
			},
		},
	})
}

// testAccCheckL2paramADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted
// it.
func testAccCheckL2paramADCValue(attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.L2param.Type(), "")
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("l2param not found on appliance")
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("l2param: appliance attr %q = %q, want %q (unset did not revert it)", attr, got, want)
		}
		return nil
	}
}

const testAccL2paramDataSource_basic = `
	resource "citrixadc_l2param" "tf_l2param" {
		mbfpeermacupdate   = 20
		maxbridgecollision = 30
		bdggrpproxyarp     = "DISABLED"
	}

	data "citrixadc_l2param" "l2param" {
		depends_on = [citrixadc_l2param.tf_l2param]
	}
`

func TestAccL2paramDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccL2paramDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_l2param.l2param", "mbfpeermacupdate", "20"),
					resource.TestCheckResourceAttr("data.citrixadc_l2param.l2param", "maxbridgecollision", "30"),
					resource.TestCheckResourceAttr("data.citrixadc_l2param.l2param", "bdggrpproxyarp", "DISABLED"),
				),
			},
		},
	})
}
