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

const testAccSubscriberparam_basic = `

resource "citrixadc_subscriberparam" "tf_subscriberparam" {
	keytype       = "IP"
	interfacetype = "None"
	idlettl       = 40
	idleaction    = "ccrTerminate"
	ipv6prefixlookuplist = [64]
	}
  
`
const testAccSubscriberparam_update = `

resource "citrixadc_subscriberparam" "tf_subscriberparam" {
	keytype       = "IP"
	interfacetype = "RadiusOnly"
	idlettl       = 50
	idleaction    = "ccrTerminate"
	ipv6prefixlookuplist = [64]
	}
  
`

const testAccSubscriberparamDataSource_basic = `

resource "citrixadc_subscriberparam" "tf_subscriberparam" {
	keytype       = "IP"
	interfacetype = "None"
	idlettl       = 40
	idleaction    = "ccrTerminate"
	ipv6prefixlookuplist = [64]
}

data "citrixadc_subscriberparam" "tf_subscriberparam" {
	depends_on = [citrixadc_subscriberparam.tf_subscriberparam]
}
`

func TestAccSubscriberparam_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccSubscriberparam_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSubscriberparamExist("citrixadc_subscriberparam.tf_subscriberparam", nil),
					resource.TestCheckResourceAttr("citrixadc_subscriberparam.tf_subscriberparam", "keytype", "IP"),
					resource.TestCheckResourceAttr("citrixadc_subscriberparam.tf_subscriberparam", "interfacetype", "None"),
					resource.TestCheckResourceAttr("citrixadc_subscriberparam.tf_subscriberparam", "idlettl", "40"),
					resource.TestCheckResourceAttr("citrixadc_subscriberparam.tf_subscriberparam", "idleaction", "ccrTerminate"),
				),
			},
			{
				Config: testAccSubscriberparam_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSubscriberparamExist("citrixadc_subscriberparam.tf_subscriberparam", nil),
					resource.TestCheckResourceAttr("citrixadc_subscriberparam.tf_subscriberparam", "keytype", "IP"),
					resource.TestCheckResourceAttr("citrixadc_subscriberparam.tf_subscriberparam", "interfacetype", "RadiusOnly"),
					resource.TestCheckResourceAttr("citrixadc_subscriberparam.tf_subscriberparam", "idlettl", "50"),
					resource.TestCheckResourceAttr("citrixadc_subscriberparam.tf_subscriberparam", "idleaction", "ccrTerminate"),
				),
			},
		},
	})
}

func testAccCheckSubscriberparamExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No subscriberparam name is set")
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
		data, err := client.FindResource("subscriberparam", "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("subscriberparam %s not found", n)
		}

		return nil
	}
}

func TestAccSubscriberparam_import(t *testing.T) {
	const resAddr = "citrixadc_subscriberparam.tf_subscriberparam"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{Config: testAccSubscriberparam_basic},
			{
				Config:                  testAccSubscriberparam_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccSubscriberparam_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.0.0"},
				},
				Config: testAccSubscriberparam_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckSubscriberparamExist("citrixadc_subscriberparam.tf_subscriberparam", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccSubscriberparam_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckSubscriberparamExist("citrixadc_subscriberparam.tf_subscriberparam", nil)),
			},
		},
	})
}

// testAccSubscriberparam_unset covers the spec-unsettable, mutable attributes of
// subscriberparam (idleaction, idlettl, interfacetype, keytype). keytype =
// IPANDVLAN is only valid when interfacetype = GxOnly, so step1 sets that
// combination; step2 removes every unset-eligible attribute so the provider
// unsets them (revert to NITRO defaults: keytype=IP, interfacetype=None,
// idlettl=0, idleaction=ccrTerminate). ipv6prefixlookuplist is NOT unsettable
// per the NITRO spec (absent from the unset payload) and is left out.
const testAccSubscriberparam_unset_step1 = `
resource "citrixadc_subscriberparam" "tf_unset" {
	keytype       = "IPANDVLAN"
	interfacetype = "GxOnly"
	idlettl       = 50
	idleaction    = "delete"
}
`

const testAccSubscriberparam_unset_step2 = `
resource "citrixadc_subscriberparam" "tf_unset" {
	# All unset-eligible attributes removed from config -> the provider must
	# unset them (revert to NITRO defaults).
}
`

func TestAccSubscriberparam_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccSubscriberparam_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSubscriberparamExist("citrixadc_subscriberparam.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_subscriberparam.tf_unset", "keytype", "IPANDVLAN"),
					resource.TestCheckResourceAttr("citrixadc_subscriberparam.tf_unset", "interfacetype", "GxOnly"),
					resource.TestCheckResourceAttr("citrixadc_subscriberparam.tf_unset", "idlettl", "50"),
					resource.TestCheckResourceAttr("citrixadc_subscriberparam.tf_unset", "idleaction", "delete"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from
				// the appliance) reverts to the documented NITRO defaults, and the
				// implicit post-apply plan must be empty.
				Config: testAccSubscriberparam_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSubscriberparamExist("citrixadc_subscriberparam.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_subscriberparam.tf_unset", "keytype", "IP"),
					resource.TestCheckResourceAttr("citrixadc_subscriberparam.tf_unset", "interfacetype", "None"),
					resource.TestCheckResourceAttr("citrixadc_subscriberparam.tf_unset", "idlettl", "0"),
					resource.TestCheckResourceAttr("citrixadc_subscriberparam.tf_unset", "idleaction", "ccrTerminate"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckSubscriberparamADCValue("keytype", "IP"),
					testAccCheckSubscriberparamADCValue("interfacetype", "None"),
					testAccCheckSubscriberparamADCValue("idleaction", "ccrTerminate"),
				),
			},
		},
	})
}

// testAccCheckSubscriberparamADCValue asserts an attribute's value directly on
// the appliance (not just in Terraform state), proving the unset actually
// reverted it.
func testAccCheckSubscriberparamADCValue(attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Subscriberparam.Type(), "")
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("subscriberparam not found on appliance")
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("subscriberparam: appliance attr %q = %q, want %q (unset did not revert it)", attr, got, want)
		}
		return nil
	}
}

func TestAccSubscriberparamDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccSubscriberparamDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_subscriberparam.tf_subscriberparam", "keytype", "IP"),
					resource.TestCheckResourceAttr("data.citrixadc_subscriberparam.tf_subscriberparam", "interfacetype", "None"),
					resource.TestCheckResourceAttr("data.citrixadc_subscriberparam.tf_subscriberparam", "idlettl", "40"),
					resource.TestCheckResourceAttr("data.citrixadc_subscriberparam.tf_subscriberparam", "idleaction", "ccrTerminate"),
				),
			},
		},
	})
}
