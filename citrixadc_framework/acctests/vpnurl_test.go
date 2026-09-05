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

const testAccVpnurl_add = `

resource "citrixadc_vpnurl" "foo" {
	actualurl        = "http://www.citrix.com"
	appjson          = "xyz"
	applicationtype  = "CVPN"
	clientlessaccess = "OFF"
	comment          = "Testing"
	linkname         = "Description"
	ssotype          = "unifiedgateway"
	urlname          = "Firsturl"
	vservername      = "server1"
	}
`
const testAccVpnurl_update = `

resource "citrixadc_vpnurl" "foo" {
	actualurl        = "http://www.citrix1.com"
	appjson          = "xyz"
	applicationtype  = "CVPN"
	clientlessaccess = "OFF"
	comment          = "Testing"
	linkname         = "Description"
	ssotype          = "unifiedgateway"
	urlname          = "Firsturl"
	vservername      = "server1"
	}
`

func TestAccVpnurl_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVpnurlDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnurl_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnurlExist("citrixadc_vpnurl.foo", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnurl.foo", "urlname", "Firsturl"),
					resource.TestCheckResourceAttr("citrixadc_vpnurl.foo", "actualurl", "http://www.citrix.com"),
					resource.TestCheckResourceAttr("citrixadc_vpnurl.foo", "linkname", "Description"),
					resource.TestCheckResourceAttr("citrixadc_vpnurl.foo", "comment", "Testing"),
				),
			},
			{
				Config: testAccVpnurl_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnurlExist("citrixadc_vpnurl.foo", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnurl.foo", "urlname", "Firsturl"),
					resource.TestCheckResourceAttr("citrixadc_vpnurl.foo", "actualurl", "http://www.citrix1.com"),
					resource.TestCheckResourceAttr("citrixadc_vpnurl.foo", "ssotype", "unifiedgateway"),
					resource.TestCheckResourceAttr("citrixadc_vpnurl.foo", "vservername", "server1"),
				),
			},
		},
	})
}

func TestAccVpnurl_import(t *testing.T) {
	const resAddr = "citrixadc_vpnurl.foo"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVpnurlDestroy,
		Steps: []resource.TestStep{
			{Config: testAccVpnurl_add},
			{
				Config:                  testAccVpnurl_add,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccCheckVpnurlExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No vpnurl name is set")
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
		data, err := client.FindResource(service.Vpnurl.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("vpnurl %s not found", n)
		}

		return nil
	}
}

func testAccCheckVpnurlDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_vpnurl" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Vpnurl.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("vpnurl %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccVpnurl_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_vpnurl.foo"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVpnurlDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnurl_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckVpnurlExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Vpnurl.Type(), "Firsturl"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccVpnurl_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckVpnurlExist(resAddr, nil)),
			},
		},
	})
}

// testAccVpnurl_unset_step1 sets the unset-eligible attributes to valid
// non-default values; step2 removes them so the provider must unset them
// (revert to NITRO defaults).
const testAccVpnurl_unset_step1 = `
resource "citrixadc_vpnurl" "tf_unset" {
	urlname          = "tf_test_vpnurl_unset"
	linkname         = "Description"
	actualurl        = "http://www.citrix.com"
	appjson          = "xyz"
	applicationtype  = "CVPN"
	clientlessaccess = "ON"
	comment          = "Testing"
	iconurl          = "http://www.citrix.com/icon.png"
	ssotype          = "unifiedgateway"
	vservername      = "server1"
}
`

const testAccVpnurl_unset_step2 = `
resource "citrixadc_vpnurl" "tf_unset" {
	urlname   = "tf_test_vpnurl_unset"
	linkname  = "Description"
	actualurl = "http://www.citrix.com"
	# All unset-eligible attributes removed from config -> the provider must
	# unset them (revert to NITRO defaults).
}
`

func TestAccVpnurl_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVpnurlDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccVpnurl_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnurlExist("citrixadc_vpnurl.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnurl.tf_unset", "appjson", "xyz"),
					resource.TestCheckResourceAttr("citrixadc_vpnurl.tf_unset", "applicationtype", "CVPN"),
					resource.TestCheckResourceAttr("citrixadc_vpnurl.tf_unset", "clientlessaccess", "ON"),
					resource.TestCheckResourceAttr("citrixadc_vpnurl.tf_unset", "comment", "Testing"),
					resource.TestCheckResourceAttr("citrixadc_vpnurl.tf_unset", "iconurl", "http://www.citrix.com/icon.png"),
					resource.TestCheckResourceAttr("citrixadc_vpnurl.tf_unset", "ssotype", "unifiedgateway"),
					resource.TestCheckResourceAttr("citrixadc_vpnurl.tf_unset", "vservername", "server1"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from
				// the appliance) reverts to the NITRO defaults, and the implicit
				// post-apply plan must be empty.
				Config: testAccVpnurl_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnurlExist("citrixadc_vpnurl.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnurl.tf_unset", "appjson", ""),
					resource.TestCheckResourceAttr("citrixadc_vpnurl.tf_unset", "applicationtype", ""),
					resource.TestCheckResourceAttr("citrixadc_vpnurl.tf_unset", "clientlessaccess", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_vpnurl.tf_unset", "comment", ""),
					resource.TestCheckNoResourceAttr("citrixadc_vpnurl.tf_unset", "iconurl"),
					resource.TestCheckResourceAttr("citrixadc_vpnurl.tf_unset", "ssotype", ""),
					resource.TestCheckResourceAttr("citrixadc_vpnurl.tf_unset", "vservername", ""),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckVpnurlADCValue("tf_test_vpnurl_unset", "clientlessaccess", "OFF"),
				),
			},
		},
	})
}

// testAccCheckVpnurlADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted
// it.
func testAccCheckVpnurlADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Vpnurl.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("vpnurl %s not found on appliance", name)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("vpnurl %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func TestAccVpnurl_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckVpnurlDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.0.0"},
				},
				Config: testAccVpnurl_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckVpnurlExist("citrixadc_vpnurl.foo", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccVpnurl_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckVpnurlExist("citrixadc_vpnurl.foo", nil)),
			},
		},
	})
}

const testAccVpnurlDataSource_basic = `

resource "citrixadc_vpnurl" "foo" {
	actualurl        = "http://www.citrix.com"
	appjson          = "xyz"
	applicationtype  = "CVPN"
	clientlessaccess = "OFF"
	comment          = "Testing"
	linkname         = "Description"
	ssotype          = "unifiedgateway"
	urlname          = "Firsturl"
	vservername      = "server1"
	}

data "citrixadc_vpnurl" "foo" {
	urlname = citrixadc_vpnurl.foo.urlname
}
`

func TestAccVpnurlDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnurlDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_vpnurl.foo", "urlname", "Firsturl"),
					resource.TestCheckResourceAttr("data.citrixadc_vpnurl.foo", "actualurl", "http://www.citrix.com"),
					resource.TestCheckResourceAttr("data.citrixadc_vpnurl.foo", "linkname", "Description"),
					resource.TestCheckResourceAttr("data.citrixadc_vpnurl.foo", "comment", "Testing"),
					resource.TestCheckResourceAttr("data.citrixadc_vpnurl.foo", "applicationtype", "CVPN"),
					resource.TestCheckResourceAttr("data.citrixadc_vpnurl.foo", "clientlessaccess", "OFF"),
					resource.TestCheckResourceAttr("data.citrixadc_vpnurl.foo", "ssotype", "unifiedgateway"),
					resource.TestCheckResourceAttr("data.citrixadc_vpnurl.foo", "vservername", "server1"),
				),
			},
		},
	})
}
