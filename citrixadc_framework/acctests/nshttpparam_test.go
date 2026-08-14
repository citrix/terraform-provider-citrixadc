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

const testAccNshttpparam_add = `
	resource "citrixadc_nshttpparam" "tf_nshttpparam" {
		dropinvalreqs             = "ON"
		markconnreqinval          = "ON"
		maxreusepool              = 1
		markhttp09inval           = "ON"
		insnssrvrhdr              = "OFF"
		logerrresp                = "OFF"
		conmultiplex              = "DISABLED"
		http2serverside           = "OFF"
		ignoreconnectcodingscheme = "ENABLED"
	}
`
const testAccNshttpparam_update = `
	resource "citrixadc_nshttpparam" "tf_nshttpparam" {
		dropinvalreqs             = "OFF"
		markconnreqinval          = "OFF"
		maxreusepool              = 0
		markhttp09inval           = "OFF"
		insnssrvrhdr              = "ON"
		logerrresp                = "ON"
		conmultiplex              = "ENABLED"
		http2serverside           = "ON"
		ignoreconnectcodingscheme = "DISABLED"
	}
`

func TestAccNshttpparam_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccNshttpparam_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNshttpparamExist("citrixadc_nshttpparam.tf_nshttpparam", nil),
					resource.TestCheckResourceAttr("citrixadc_nshttpparam.tf_nshttpparam", "dropinvalreqs", "ON"),
					resource.TestCheckResourceAttr("citrixadc_nshttpparam.tf_nshttpparam", "markconnreqinval", "ON"),
					resource.TestCheckResourceAttr("citrixadc_nshttpparam.tf_nshttpparam", "maxreusepool", "1"),
					resource.TestCheckResourceAttr("citrixadc_nshttpparam.tf_nshttpparam", "markhttp09inval", "ON"),
					resource.TestCheckResourceAttr("citrixadc_nshttpparam.tf_nshttpparam", "conmultiplex", "DISABLED"),
				),
			},
			{
				Config: testAccNshttpparam_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNshttpparamExist("citrixadc_nshttpparam.tf_nshttpparam", nil),
					resource.TestCheckResourceAttr("citrixadc_nshttpparam.tf_nshttpparam", "dropinvalreqs", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_nshttpparam.tf_nshttpparam", "markconnreqinval", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_nshttpparam.tf_nshttpparam", "maxreusepool", "0"),
					resource.TestCheckResourceAttr("citrixadc_nshttpparam.tf_nshttpparam", "markhttp09inval", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_nshttpparam.tf_nshttpparam", "conmultiplex", "ENABLED"),
				),
			},
		},
	})
}

func TestAccNshttpparam_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccNshttpparam_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNshttpparamExist("citrixadc_nshttpparam.tf_nshttpparam", nil),
				),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccNshttpparam_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNshttpparamExist("citrixadc_nshttpparam.tf_nshttpparam", nil),
				),
			},
		},
	})
}

// nshttpparam is a singleton config resource. Step 1 applies non-default
// values for every unset-eligible attribute; step 2 removes them all from
// config, so the provider must unset them (revert to the documented NITRO
// defaults) and the post-apply plan must be empty.
const testAccNshttpparam_unset_step1 = `
	resource "citrixadc_nshttpparam" "tf_unset" {
		conmultiplex              = "DISABLED"
		dropinvalreqs             = "ON"
		http2serverside           = "ON"
		ignoreconnectcodingscheme = "ENABLED"
		insnssrvrhdr              = "ON"
		logerrresp                = "OFF"
		markconnreqinval          = "ON"
		markhttp09inval           = "ON"
		maxreusepool              = 5
	}
`

const testAccNshttpparam_unset_step2 = `
	resource "citrixadc_nshttpparam" "tf_unset" {
	}
`

func TestAccNshttpparam_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccNshttpparam_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNshttpparamExist("citrixadc_nshttpparam.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_nshttpparam.tf_unset", "conmultiplex", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_nshttpparam.tf_unset", "dropinvalreqs", "ON"),
					resource.TestCheckResourceAttr("citrixadc_nshttpparam.tf_unset", "http2serverside", "ON"),
					resource.TestCheckResourceAttr("citrixadc_nshttpparam.tf_unset", "ignoreconnectcodingscheme", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_nshttpparam.tf_unset", "insnssrvrhdr", "ON"),
					resource.TestCheckResourceAttr("citrixadc_nshttpparam.tf_unset", "logerrresp", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_nshttpparam.tf_unset", "markconnreqinval", "ON"),
					resource.TestCheckResourceAttr("citrixadc_nshttpparam.tf_unset", "markhttp09inval", "ON"),
					resource.TestCheckResourceAttr("citrixadc_nshttpparam.tf_unset", "maxreusepool", "5"),
				),
			},
			{
				// Removing the attributes must unset them: state reverts to the
				// documented NITRO defaults and the implicit post-apply plan is empty.
				Config: testAccNshttpparam_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNshttpparamExist("citrixadc_nshttpparam.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_nshttpparam.tf_unset", "conmultiplex", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_nshttpparam.tf_unset", "dropinvalreqs", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_nshttpparam.tf_unset", "http2serverside", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_nshttpparam.tf_unset", "ignoreconnectcodingscheme", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_nshttpparam.tf_unset", "insnssrvrhdr", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_nshttpparam.tf_unset", "logerrresp", "ON"),
					resource.TestCheckResourceAttr("citrixadc_nshttpparam.tf_unset", "markconnreqinval", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_nshttpparam.tf_unset", "markhttp09inval", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_nshttpparam.tf_unset", "maxreusepool", "0"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckNshttpparamADCValue("conmultiplex", "ENABLED"),
					testAccCheckNshttpparamADCValue("dropinvalreqs", "OFF"),
					testAccCheckNshttpparamADCValue("logerrresp", "ON"),
					testAccCheckNshttpparamADCValue("maxreusepool", "0"),
				),
			},
		},
	})
}

// testAccCheckNshttpparamADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted it.
func testAccCheckNshttpparamADCValue(attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Nshttpparam.Type(), "")
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("nshttpparam not found on appliance")
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("nshttpparam: appliance attr %q = %q, want %q (unset did not revert it)", attr, got, want)
		}
		return nil
	}
}

func testAccCheckNshttpparamExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No nshttpparam name is set")
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
		data, err := client.FindResource(service.Nshttpparam.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("nshttpparam %s not found", n)
		}

		return nil
	}
}

func TestAccNshttpparam_import(t *testing.T) {
	const resAddr = "citrixadc_nshttpparam.tf_nshttpparam"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{Config: testAccNshttpparam_add},
			{
				Config:                  testAccNshttpparam_add,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

const testAccNshttpparamDataSource_basic = `
	resource "citrixadc_nshttpparam" "tf_nshttpparam_ds" {
		dropinvalreqs             = "ON"
		markconnreqinval          = "ON"
		maxreusepool              = 2
		markhttp09inval           = "ON"
		insnssrvrhdr              = "OFF"
		logerrresp                = "OFF"
		conmultiplex              = "DISABLED"
		http2serverside           = "OFF"
		ignoreconnectcodingscheme = "ENABLED"
	}

	data "citrixadc_nshttpparam" "tf_nshttpparam_ds" {
		depends_on = [citrixadc_nshttpparam.tf_nshttpparam_ds]
	}
`

func TestAccNshttpparamDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNshttpparamDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_nshttpparam.tf_nshttpparam_ds", "dropinvalreqs", "ON"),
					resource.TestCheckResourceAttr("data.citrixadc_nshttpparam.tf_nshttpparam_ds", "markconnreqinval", "ON"),
					resource.TestCheckResourceAttr("data.citrixadc_nshttpparam.tf_nshttpparam_ds", "maxreusepool", "2"),
					resource.TestCheckResourceAttr("data.citrixadc_nshttpparam.tf_nshttpparam_ds", "markhttp09inval", "ON"),
					resource.TestCheckResourceAttr("data.citrixadc_nshttpparam.tf_nshttpparam_ds", "conmultiplex", "DISABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_nshttpparam.tf_nshttpparam_ds", "http2serverside", "OFF"),
					resource.TestCheckResourceAttr("data.citrixadc_nshttpparam.tf_nshttpparam_ds", "ignoreconnectcodingscheme", "ENABLED"),
				),
			},
		},
	})
}
