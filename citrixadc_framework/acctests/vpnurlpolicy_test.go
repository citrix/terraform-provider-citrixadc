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

const testAccVpnurlpolicy_add = `

	resource "citrixadc_vpnurlaction" "tf_vpnurlaction" {
		name             = "tf_vpnurlaction"
		linkname         = "new_link"
		actualurl        = "http://www.citrix.com"
		applicationtype  = "CVPN"
		clientlessaccess = "OFF"
		comment          = "Testing"
		ssotype          = "unifiedgateway"
		vservername      = "vserver1"
	}
	resource "citrixadc_vpnurlpolicy" "tf_vpnurlpolicy" {
		name = "new_policy"
		rule = "true"
		action = citrixadc_vpnurlaction.tf_vpnurlaction.name
	}
`
const testAccVpnurlpolicy_update = `

	resource "citrixadc_vpnurlaction" "tf_vpnurlaction" {
		name             = "tf_vpnurlaction"
		linkname         = "new_link"
		actualurl        = "http://www.citrix.com"
		applicationtype  = "CVPN"
		clientlessaccess = "OFF"
		comment          = "Testing"
		ssotype          = "unifiedgateway"
		vservername      = "vserver1"
	}
	resource "citrixadc_vpnurlpolicy" "tf_vpnurlpolicy" {
		name = "new_policy"
		rule = "false"
		action = citrixadc_vpnurlaction.tf_vpnurlaction.name
	}
`

func TestAccVpnurlpolicy_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVpnurlpolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnurlpolicy_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnurlpolicyExist("citrixadc_vpnurlpolicy.tf_vpnurlpolicy", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnurlpolicy.tf_vpnurlpolicy", "name", "new_policy"),
					resource.TestCheckResourceAttr("citrixadc_vpnurlpolicy.tf_vpnurlpolicy", "rule", "true"),
					resource.TestCheckResourceAttr("citrixadc_vpnurlpolicy.tf_vpnurlpolicy", "action", "tf_vpnurlaction"),
				),
			},
			{
				Config: testAccVpnurlpolicy_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnurlpolicyExist("citrixadc_vpnurlpolicy.tf_vpnurlpolicy", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnurlpolicy.tf_vpnurlpolicy", "name", "new_policy"),
					resource.TestCheckResourceAttr("citrixadc_vpnurlpolicy.tf_vpnurlpolicy", "rule", "false"),
					resource.TestCheckResourceAttr("citrixadc_vpnurlpolicy.tf_vpnurlpolicy", "action", "tf_vpnurlaction"),
				),
			},
		},
	})
}

func TestAccVpnurlpolicy_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_vpnurlpolicy.tf_vpnurlpolicy"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVpnurlpolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnurlpolicy_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckVpnurlpolicyExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Vpnurlpolicy.Type(), "new_policy"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccVpnurlpolicy_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckVpnurlpolicyExist(resAddr, nil)),
			},
		},
	})
}

func TestAccVpnurlpolicy_import(t *testing.T) {
	const resAddr = "citrixadc_vpnurlpolicy.tf_vpnurlpolicy"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVpnurlpolicyDestroy,
		Steps: []resource.TestStep{
			{Config: testAccVpnurlpolicy_add},
			{
				Config:                  testAccVpnurlpolicy_add,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccCheckVpnurlpolicyExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No vpnurlpolicy name is set")
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
		data, err := client.FindResource("vpnurlpolicy", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("vpnurlpolicy %s not found", n)
		}

		return nil
	}
}

func testAccCheckVpnurlpolicyDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_vpnurlpolicy" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("vpnurlpolicy", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("vpnurlpolicy %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccVpnurlpolicy_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckVpnurlpolicyDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccVpnurlpolicy_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckVpnurlpolicyExist("citrixadc_vpnurlpolicy.tf_vpnurlpolicy", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccVpnurlpolicy_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckVpnurlpolicyExist("citrixadc_vpnurlpolicy.tf_vpnurlpolicy", nil)),
			},
		},
	})
}

// The vpnurlpolicy unset test covers the two spec-unsettable, mutable
// attributes: comment and logaction. Both are Optional+Computed with no server
// default; after a NITRO ?action=unset the appliance drops them (absent on GET).
const testAccVpnurlpolicy_unset_step1 = `
	resource "citrixadc_vpnurlaction" "tf_unset_action" {
		name             = "tf_unset_urlaction"
		linkname         = "new_link"
		actualurl        = "http://www.citrix.com"
		applicationtype  = "CVPN"
		clientlessaccess = "OFF"
		ssotype          = "unifiedgateway"
		vservername      = "vserver1"
	}
	resource "citrixadc_auditmessageaction" "tf_unset_msg" {
		name              = "tf_unset_msgaction"
		loglevel          = "NOTICE"
		stringbuilderexpr = "\"hello\""
	}
	resource "citrixadc_vpnurlpolicy" "tf_unset" {
		name      = "tf_unset_policy"
		rule      = "true"
		action    = citrixadc_vpnurlaction.tf_unset_action.name
		comment   = "unset me"
		logaction = citrixadc_auditmessageaction.tf_unset_msg.name
	}
`

const testAccVpnurlpolicy_unset_step2 = `
	resource "citrixadc_vpnurlaction" "tf_unset_action" {
		name             = "tf_unset_urlaction"
		linkname         = "new_link"
		actualurl        = "http://www.citrix.com"
		applicationtype  = "CVPN"
		clientlessaccess = "OFF"
		ssotype          = "unifiedgateway"
		vservername      = "vserver1"
	}
	resource "citrixadc_auditmessageaction" "tf_unset_msg" {
		name              = "tf_unset_msgaction"
		loglevel          = "NOTICE"
		stringbuilderexpr = "\"hello\""
	}
	resource "citrixadc_vpnurlpolicy" "tf_unset" {
		name   = "tf_unset_policy"
		rule   = "true"
		action = citrixadc_vpnurlaction.tf_unset_action.name
		# comment and logaction removed from config -> provider must unset them
		# (appliance reverts to defaults: absent/empty).
	}
`

func TestAccVpnurlpolicy_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVpnurlpolicyDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccVpnurlpolicy_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnurlpolicyExist("citrixadc_vpnurlpolicy.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnurlpolicy.tf_unset", "comment", "unset me"),
					resource.TestCheckResourceAttr("citrixadc_vpnurlpolicy.tf_unset", "logaction", "tf_unset_msgaction"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from the
				// appliance) reverts to the NITRO defaults, and the implicit post-apply
				// plan must be empty.
				Config: testAccVpnurlpolicy_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnurlpolicyExist("citrixadc_vpnurlpolicy.tf_unset", nil),
					// After unset the attributes are absent (null) in state.
					resource.TestCheckNoResourceAttr("citrixadc_vpnurlpolicy.tf_unset", "comment"),
					resource.TestCheckNoResourceAttr("citrixadc_vpnurlpolicy.tf_unset", "logaction"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckVpnurlpolicyADCValue("tf_unset_policy", "comment", ""),
					testAccCheckVpnurlpolicyADCValue("tf_unset_policy", "logaction", ""),
				),
			},
		},
	})
}

// testAccCheckVpnurlpolicyADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted
// it. An unset attribute is absent from the GET response, which reads back as "".
func testAccCheckVpnurlpolicyADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Vpnurlpolicy.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("vpnurlpolicy %s not found on appliance", name)
		}
		got := ""
		if v, ok := data[attr]; ok && v != nil {
			got = strings.TrimSpace(fmt.Sprintf("%v", v))
		}
		if got != want {
			return fmt.Errorf("vpnurlpolicy %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

const testAccVpnurlpolicyDataSource_basic = `

	resource "citrixadc_vpnurlaction" "tf_vpnurlaction" {
		name             = "tf_vpnurlaction"
		linkname         = "new_link"
		actualurl        = "http://www.citrix.com"
		applicationtype  = "CVPN"
		clientlessaccess = "OFF"
		comment          = "Testing"
		ssotype          = "unifiedgateway"
		vservername      = "vserver1"
	}
	resource "citrixadc_vpnurlpolicy" "tf_vpnurlpolicy" {
		name = "new_policy"
		rule = "true"
		action = citrixadc_vpnurlaction.tf_vpnurlaction.name
	}

data "citrixadc_vpnurlpolicy" "tf_vpnurlpolicy" {
	name = citrixadc_vpnurlpolicy.tf_vpnurlpolicy.name
}
`

func TestAccVpnurlpolicyDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnurlpolicyDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_vpnurlpolicy.tf_vpnurlpolicy", "name", "new_policy"),
					resource.TestCheckResourceAttr("data.citrixadc_vpnurlpolicy.tf_vpnurlpolicy", "rule", "true"),
					resource.TestCheckResourceAttr("data.citrixadc_vpnurlpolicy.tf_vpnurlpolicy", "action", "tf_vpnurlaction"),
				),
			},
		},
	})
}
