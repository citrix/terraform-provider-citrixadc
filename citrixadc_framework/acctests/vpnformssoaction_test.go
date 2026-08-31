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

const testAccVpnformssoaction_basic = `

resource "citrixadc_vpnformssoaction" "tf_vpnformssoaction" {
	name = "tf_vpnformssoaction"
	actionurl = "/home"
	userfield = "username"
	passwdfield = "password"
	ssosuccessrule = "true"
	namevaluepair = "name1=value1&name2=value2"
	nvtype = "STATIC"
	responsesize = "150"
	submitmethod = "POST"
}
`

const testAccVpnformssoaction_basic_update_mandatory_attributes = `

	resource "citrixadc_vpnformssoaction" "tf_vpnformssoaction" {
		name = "tf_vpnformssoaction"
		actionurl = "/contact"
		userfield = "username1"
		passwdfield = "password1"
		ssosuccessrule = "false"
	}
`

const testAccVpnformssoaction_basic_update_non_mandatory_attributes = `

	resource "citrixadc_vpnformssoaction" "tf_vpnformssoaction" {
		name = "tf_vpnformssoaction"
		actionurl = "/contact"
		userfield = "username1"
		passwdfield = "password1"
		ssosuccessrule = "false"
		namevaluepair = "name3=value3"
		nvtype = "DYNAMIC"
		responsesize = "151"
		submitmethod = "GET"
	}
`

func TestAccVpnformssoaction_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVpnformssoactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnformssoaction_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnformssoactionExist("citrixadc_vpnformssoaction.tf_vpnformssoaction", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnformssoaction.tf_vpnformssoaction", "name", "tf_vpnformssoaction"),
					resource.TestCheckResourceAttr("citrixadc_vpnformssoaction.tf_vpnformssoaction", "actionurl", "/home"),
					resource.TestCheckResourceAttr("citrixadc_vpnformssoaction.tf_vpnformssoaction", "userfield", "username"),
					resource.TestCheckResourceAttr("citrixadc_vpnformssoaction.tf_vpnformssoaction", "passwdfield", "password"),
					resource.TestCheckResourceAttr("citrixadc_vpnformssoaction.tf_vpnformssoaction", "ssosuccessrule", "true"),
					resource.TestCheckResourceAttr("citrixadc_vpnformssoaction.tf_vpnformssoaction", "namevaluepair", "name1=value1&name2=value2"),
					resource.TestCheckResourceAttr("citrixadc_vpnformssoaction.tf_vpnformssoaction", "nvtype", "STATIC"),
					resource.TestCheckResourceAttr("citrixadc_vpnformssoaction.tf_vpnformssoaction", "responsesize", "150"),
					resource.TestCheckResourceAttr("citrixadc_vpnformssoaction.tf_vpnformssoaction", "submitmethod", "POST"),
				),
			},
			{
				Config: testAccVpnformssoaction_basic_update_mandatory_attributes,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnformssoactionExist("citrixadc_vpnformssoaction.tf_vpnformssoaction", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnformssoaction.tf_vpnformssoaction", "actionurl", "/contact"),
					resource.TestCheckResourceAttr("citrixadc_vpnformssoaction.tf_vpnformssoaction", "userfield", "username1"),
					resource.TestCheckResourceAttr("citrixadc_vpnformssoaction.tf_vpnformssoaction", "passwdfield", "password1"),
					resource.TestCheckResourceAttr("citrixadc_vpnformssoaction.tf_vpnformssoaction", "ssosuccessrule", "false"),
				),
			},
			{
				Config: testAccVpnformssoaction_basic_update_non_mandatory_attributes,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnformssoactionExist("citrixadc_vpnformssoaction.tf_vpnformssoaction", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnformssoaction.tf_vpnformssoaction", "namevaluepair", "name3=value3"),
					resource.TestCheckResourceAttr("citrixadc_vpnformssoaction.tf_vpnformssoaction", "nvtype", "DYNAMIC"),
					resource.TestCheckResourceAttr("citrixadc_vpnformssoaction.tf_vpnformssoaction", "responsesize", "151"),
					resource.TestCheckResourceAttr("citrixadc_vpnformssoaction.tf_vpnformssoaction", "submitmethod", "GET"),
				),
			},
		},
	})
}

func TestAccVpnformssoaction_import(t *testing.T) {
	const resAddr = "citrixadc_vpnformssoaction.tf_vpnformssoaction"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVpnformssoactionDestroy,
		Steps: []resource.TestStep{
			{Config: testAccVpnformssoaction_basic},
			{
				Config:                  testAccVpnformssoaction_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccCheckVpnformssoactionExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No vpnformssoaction name is set")
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
		data, err := client.FindResource(service.Vpnformssoaction.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("vpnformssoaction %s not found", n)
		}

		return nil
	}
}

func testAccCheckVpnformssoactionDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_vpnformssoaction" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Vpnformssoaction.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("vpnformssoaction %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccVpnformssoaction_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_vpnformssoaction.tf_vpnformssoaction"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVpnformssoactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnformssoaction_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckVpnformssoactionExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Vpnformssoaction.Type(), "tf_vpnformssoaction"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccVpnformssoaction_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckVpnformssoactionExist(resAddr, nil)),
			},
		},
	})
}

// The vpnformssoaction unset test covers the spec-unsettable mutable
// attributes (responsesize, namevaluepair, nvtype, submitmethod). Step 1 sets
// them to non-default values; step 2 removes them so the provider must unset
// them and NITRO reverts to its documented defaults.
const testAccVpnformssoaction_unset_step1 = `
resource "citrixadc_vpnformssoaction" "tf_unset" {
	name           = "tf_test_vpnformssoaction_unset"
	actionurl      = "/home"
	userfield      = "username"
	passwdfield    = "password"
	ssosuccessrule = "true"
	namevaluepair  = "name1=value1&name2=value2"
	nvtype         = "STATIC"
	responsesize   = "150"
	submitmethod   = "POST"
}
`

const testAccVpnformssoaction_unset_step2 = `
resource "citrixadc_vpnformssoaction" "tf_unset" {
	name           = "tf_test_vpnformssoaction_unset"
	actionurl      = "/home"
	userfield      = "username"
	passwdfield    = "password"
	ssosuccessrule = "true"
	# All unset-eligible attributes removed from config -> the provider must
	# unset them (revert to NITRO defaults).
}
`

func TestAccVpnformssoaction_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVpnformssoactionDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccVpnformssoaction_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnformssoactionExist("citrixadc_vpnformssoaction.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnformssoaction.tf_unset", "namevaluepair", "name1=value1&name2=value2"),
					resource.TestCheckResourceAttr("citrixadc_vpnformssoaction.tf_unset", "nvtype", "STATIC"),
					resource.TestCheckResourceAttr("citrixadc_vpnformssoaction.tf_unset", "responsesize", "150"),
					resource.TestCheckResourceAttr("citrixadc_vpnformssoaction.tf_unset", "submitmethod", "POST"),
				),
			},
			{
				// Removing the attributes must unset them: state reverts to the
				// documented NITRO defaults, and the implicit post-apply plan
				// must be empty.
				Config: testAccVpnformssoaction_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnformssoactionExist("citrixadc_vpnformssoaction.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnformssoaction.tf_unset", "nvtype", "DYNAMIC"),
					resource.TestCheckResourceAttr("citrixadc_vpnformssoaction.tf_unset", "responsesize", "8096"),
					resource.TestCheckResourceAttr("citrixadc_vpnformssoaction.tf_unset", "submitmethod", "GET"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckVpnformssoactionADCValue("tf_test_vpnformssoaction_unset", "nvtype", "DYNAMIC"),
					testAccCheckVpnformssoactionADCValue("tf_test_vpnformssoaction_unset", "responsesize", "8096"),
					testAccCheckVpnformssoactionADCValue("tf_test_vpnformssoaction_unset", "submitmethod", "GET"),
				),
			},
		},
	})
}

// testAccCheckVpnformssoactionADCValue asserts an attribute's value directly on
// the appliance (not just in Terraform state), proving the unset actually
// reverted it.
func testAccCheckVpnformssoactionADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Vpnformssoaction.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("vpnformssoaction %s not found on appliance", name)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("vpnformssoaction %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

const testAccVpnformssoactionDataSource_basic = `

resource "citrixadc_vpnformssoaction" "tf_vpnformssoaction" {
	name = "tf_vpnformssoaction"
	actionurl = "/home"
	userfield = "username"
	passwdfield = "password"
	ssosuccessrule = "true"
	namevaluepair = "name1=value1&name2=value2"
	nvtype = "STATIC"
	responsesize = "150"
	submitmethod = "POST"
}

data "citrixadc_vpnformssoaction" "tf_vpnformssoaction" {
	name = citrixadc_vpnformssoaction.tf_vpnformssoaction.name
}
`

func TestAccVpnformssoaction_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckVpnformssoactionDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.0.0"},
				},
				Config: testAccVpnformssoaction_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckVpnformssoactionExist("citrixadc_vpnformssoaction.tf_vpnformssoaction", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccVpnformssoaction_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckVpnformssoactionExist("citrixadc_vpnformssoaction.tf_vpnformssoaction", nil)),
			},
		},
	})
}

func TestAccVpnformssoactionDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVpnformssoactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnformssoactionDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_vpnformssoaction.tf_vpnformssoaction", "name", "tf_vpnformssoaction"),
					resource.TestCheckResourceAttr("data.citrixadc_vpnformssoaction.tf_vpnformssoaction", "actionurl", "/home"),
					resource.TestCheckResourceAttr("data.citrixadc_vpnformssoaction.tf_vpnformssoaction", "userfield", "username"),
					resource.TestCheckResourceAttr("data.citrixadc_vpnformssoaction.tf_vpnformssoaction", "passwdfield", "password"),
					resource.TestCheckResourceAttr("data.citrixadc_vpnformssoaction.tf_vpnformssoaction", "ssosuccessrule", "true"),
					resource.TestCheckResourceAttr("data.citrixadc_vpnformssoaction.tf_vpnformssoaction", "namevaluepair", "name1=value1&name2=value2"),
					resource.TestCheckResourceAttr("data.citrixadc_vpnformssoaction.tf_vpnformssoaction", "nvtype", "STATIC"),
					resource.TestCheckResourceAttr("data.citrixadc_vpnformssoaction.tf_vpnformssoaction", "responsesize", "150"),
					resource.TestCheckResourceAttr("data.citrixadc_vpnformssoaction.tf_vpnformssoaction", "submitmethod", "POST"),
				),
			},
		},
	})
}
