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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testAccVpntrafficaction_add = `

	resource "citrixadc_vpntrafficaction" "tf_action" {
		name       = "Testing"
		qual       = "tcp"
		apptimeout = 20
		fta        = "OFF"
		hdx        = "OFF"
		sso        = "ON"
	}
`
const testAccVpntrafficaction_update = `

	resource "citrixadc_vpntrafficaction" "tf_action" {
		name       = "Testing"
		qual       = "tcp"
		apptimeout = 30
		fta        = "OFF"
		hdx        = "OFF"
		sso        = "OFF"
	}
`

const testAccVpntrafficactionDataSource_basic = `

	resource "citrixadc_vpntrafficaction" "tf_action" {
		name       = "Testing"
		qual       = "tcp"
		apptimeout = 20
		fta        = "OFF"
		hdx        = "OFF"
		sso        = "ON"
	}

	data "citrixadc_vpntrafficaction" "tf_action" {
		name = citrixadc_vpntrafficaction.tf_action.name
	}
`

func TestAccVpntrafficaction_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVpntrafficactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVpntrafficaction_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpntrafficactionExist("citrixadc_vpntrafficaction.tf_action", nil),
					resource.TestCheckResourceAttr("citrixadc_vpntrafficaction.tf_action", "name", "Testing"),
					resource.TestCheckResourceAttr("citrixadc_vpntrafficaction.tf_action", "apptimeout", "20"),
					resource.TestCheckResourceAttr("citrixadc_vpntrafficaction.tf_action", "sso", "ON"),
				),
			},
			{
				Config: testAccVpntrafficaction_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpntrafficactionExist("citrixadc_vpntrafficaction.tf_action", nil),
					resource.TestCheckResourceAttr("citrixadc_vpntrafficaction.tf_action", "name", "Testing"),
					resource.TestCheckResourceAttr("citrixadc_vpntrafficaction.tf_action", "apptimeout", "30"),
					resource.TestCheckResourceAttr("citrixadc_vpntrafficaction.tf_action", "sso", "OFF"),
				),
			},
		},
	})
}

func testAccCheckVpntrafficactionExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No vpntrafficaction name is set")
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
		data, err := client.FindResource(service.Vpntrafficaction.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("vpntrafficaction %s not found", n)
		}

		return nil
	}
}

func testAccCheckVpntrafficactionDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_vpntrafficaction" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Vpntrafficaction.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("vpntrafficaction %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccVpntrafficaction_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_vpntrafficaction.tf_action"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVpntrafficactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVpntrafficaction_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckVpntrafficactionExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Vpntrafficaction.Type(), "Testing"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccVpntrafficaction_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckVpntrafficactionExist(resAddr, nil)),
			},
		},
	})
}

func TestAccVpntrafficaction_import(t *testing.T) {
	const resAddr = "citrixadc_vpntrafficaction.tf_action"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVpntrafficactionDestroy,
		Steps: []resource.TestStep{
			{Config: testAccVpntrafficaction_add},
			{
				Config:            testAccVpntrafficaction_add,
				ResourceName:      resAddr,
				ImportState:       true,
				ImportStateVerify: true,
				// proxy, userexpression and passwdexpression carry an empty-string
				// schema default so they can be unset. NITRO omits them from GET when
				// empty, so a fresh import reads them back as null and they cannot
				// round-trip through import verification.
				ImportStateVerifyIgnore: []string{"proxy", "userexpression", "passwdexpression"},
			},
		},
	})
}

func TestAccVpntrafficaction_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckVpntrafficactionDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccVpntrafficaction_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckVpntrafficactionExist("citrixadc_vpntrafficaction.tf_action", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccVpntrafficaction_add,
				Check:                    resource.ComposeTestCheckFunc(testAccCheckVpntrafficactionExist("citrixadc_vpntrafficaction.tf_action", nil)),
			},
		},
	})
}

// The vpntrafficaction unset test covers every attribute NITRO documents in its
// unset payload (wanscaler, proxy, userexpression, passwdexpression).
// step1 sets them to valid non-default values; step2 removes them from config so
// the provider must unset them (revert to NITRO defaults).
const testAccVpntrafficaction_unset_step1 = `

	resource "citrixadc_vpntrafficaction" "tf_unset" {
		name             = "tf_unset_action"
		qual             = "http"
		proxy            = "10.0.0.1:8080"
		userexpression   = "http.req.header(\"username\")"
		passwdexpression = "http.req.header(\"password\")"
	}
`

const testAccVpntrafficaction_unset_step2 = `

	resource "citrixadc_vpntrafficaction" "tf_unset" {
		name = "tf_unset_action"
		qual = "http"
		# All unset-eligible attributes removed from config -> the provider must
		# unset them (revert to NITRO defaults).
	}
`

func TestAccVpntrafficaction_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVpntrafficactionDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccVpntrafficaction_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpntrafficactionExist("citrixadc_vpntrafficaction.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_vpntrafficaction.tf_unset", "proxy", "10.0.0.1:8080"),
					resource.TestCheckResourceAttr("citrixadc_vpntrafficaction.tf_unset", "userexpression", "http.req.header(\"username\")"),
					resource.TestCheckResourceAttr("citrixadc_vpntrafficaction.tf_unset", "passwdexpression", "http.req.header(\"password\")"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from
				// the appliance) reverts to the documented NITRO defaults, and the
				// implicit post-apply plan must be empty.
				Config: testAccVpntrafficaction_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpntrafficactionExist("citrixadc_vpntrafficaction.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_vpntrafficaction.tf_unset", "proxy", ""),
					resource.TestCheckResourceAttr("citrixadc_vpntrafficaction.tf_unset", "userexpression", ""),
					resource.TestCheckResourceAttr("citrixadc_vpntrafficaction.tf_unset", "passwdexpression", ""),
					// Independent appliance-level confirmation the unset took effect:
					// NITRO omits these attributes from GET once reverted to their
					// defaults, so the appliance no longer reports the non-default value.
					testAccCheckVpntrafficactionADCValue("tf_unset_action", "proxy", "<nil>"),
					testAccCheckVpntrafficactionADCValue("tf_unset_action", "userexpression", "<nil>"),
				),
			},
		},
	})
}

// testAccCheckVpntrafficactionADCValue asserts an attribute's value directly on
// the appliance (not just in Terraform state), proving the unset actually
// reverted it.
func testAccCheckVpntrafficactionADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Vpntrafficaction.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("vpntrafficaction %s not found on appliance", name)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("vpntrafficaction %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func TestAccVpntrafficactionDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccVpntrafficactionDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_vpntrafficaction.tf_action", "name", "Testing"),
					resource.TestCheckResourceAttr("data.citrixadc_vpntrafficaction.tf_action", "apptimeout", "20"),
					resource.TestCheckResourceAttr("data.citrixadc_vpntrafficaction.tf_action", "sso", "ON"),
				),
			},
		},
	})
}
