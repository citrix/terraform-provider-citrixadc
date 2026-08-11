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
	"testing"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testAccVpnurlaction_add = `

resource "citrixadc_vpnurlaction" "foo" {
	name             = "tf_vpnurlaction"
	actualurl        = "https://www.citrix.com"
	linkname         = "new_link"
	applicationtype  = "CVPN"
	clientlessaccess = "ON"
	comment          = "Testing"
	ssotype          = "unifiedgateway"
	vservername      = "vserver1"
	}
  
`
const testAccVpnurlaction_update = `

resource "citrixadc_vpnurlaction" "foo" {
	name             = "tf_vpnurlaction"
	actualurl        = "https://www.citrix.com/products/citrix-adc/"
	linkname         = "new_link"
	applicationtype  = "CVPN"
	clientlessaccess = "OFF"
	comment          = "Testing"
	ssotype          = "unifiedgateway"
	vservername      = "vserver1"
	}
  
`

func TestAccVpnurlaction_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVpnurlactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnurlaction_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnurlactionExist("citrixadc_vpnurlaction.foo", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnurlaction.foo", "name", "tf_vpnurlaction"),
					resource.TestCheckResourceAttr("citrixadc_vpnurlaction.foo", "actualurl", "https://www.citrix.com"),
					resource.TestCheckResourceAttr("citrixadc_vpnurlaction.foo", "clientlessaccess", "ON"),
				),
			},
			{
				Config: testAccVpnurlaction_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnurlactionExist("citrixadc_vpnurlaction.foo", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnurlaction.foo", "name", "tf_vpnurlaction"),
					resource.TestCheckResourceAttr("citrixadc_vpnurlaction.foo", "actualurl", "https://www.citrix.com/products/citrix-adc/"),
					resource.TestCheckResourceAttr("citrixadc_vpnurlaction.foo", "clientlessaccess", "OFF"),
				),
			},
		},
	})
}

func TestAccVpnurlaction_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_vpnurlaction.foo"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVpnurlactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnurlaction_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckVpnurlactionExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Vpnurlaction.Type(), "tf_vpnurlaction"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccVpnurlaction_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckVpnurlactionExist(resAddr, nil)),
			},
		},
	})
}

func TestAccVpnurlaction_import(t *testing.T) {
	const resAddr = "citrixadc_vpnurlaction.foo"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVpnurlactionDestroy,
		Steps: []resource.TestStep{
			{Config: testAccVpnurlaction_add},
			{
				Config:                  testAccVpnurlaction_add,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccCheckVpnurlactionExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No vpnurlaction name is set")
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
		data, err := client.FindResource("vpnurlaction", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("vpnurlaction %s not found", n)
		}

		return nil
	}
}

func testAccCheckVpnurlactionDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_vpnurlaction" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("vpnurlaction", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("vpnurlaction %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccVpnurlactionDataSource_basic = `

resource "citrixadc_vpnurlaction" "foo" {
	name             = "tf_vpnurlaction"
	actualurl        = "https://www.citrix.com"
	linkname         = "new_link"
	applicationtype  = "CVPN"
	clientlessaccess = "ON"
	comment          = "Testing"
	ssotype          = "unifiedgateway"
	vservername      = "vserver1"
	}

data "citrixadc_vpnurlaction" "foo" {
	name = citrixadc_vpnurlaction.foo.name
}
`

func TestAccVpnurlaction_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckVpnurlactionDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccVpnurlaction_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckVpnurlactionExist("citrixadc_vpnurlaction.foo", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccVpnurlaction_add,
				Check:                    resource.ComposeTestCheckFunc(testAccCheckVpnurlactionExist("citrixadc_vpnurlaction.foo", nil)),
			},
		},
	})
}

// The vpnurlaction unset test covers the spec-unsettable, mutable attributes:
// applicationtype, clientlessaccess, comment, ssotype and vservername. Step 1
// applies non-default values; step 2 removes them from config, so the provider
// must unset them (revert to NITRO defaults -- "OFF" for clientlessaccess, empty
// for the free-form/enum attributes with no documented default).
const testAccVpnurlaction_unset_step1 = `
resource "citrixadc_vpnurlaction" "tf_unset" {
	name             = "tf_test_vpnurlaction_unset"
	actualurl        = "https://www.citrix.com"
	linkname         = "new_link"
	applicationtype  = "CVPN"
	clientlessaccess = "ON"
	comment          = "Testing"
	ssotype          = "unifiedgateway"
	vservername      = "vserver1"
}
`

const testAccVpnurlaction_unset_step2 = `
resource "citrixadc_vpnurlaction" "tf_unset" {
	name      = "tf_test_vpnurlaction_unset"
	actualurl = "https://www.citrix.com"
	linkname  = "new_link"
	# All unset-eligible attributes removed from config -> the provider must
	# unset them (revert to NITRO defaults).
}
`

func TestAccVpnurlaction_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVpnurlactionDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccVpnurlaction_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnurlactionExist("citrixadc_vpnurlaction.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnurlaction.tf_unset", "applicationtype", "CVPN"),
					resource.TestCheckResourceAttr("citrixadc_vpnurlaction.tf_unset", "clientlessaccess", "ON"),
					resource.TestCheckResourceAttr("citrixadc_vpnurlaction.tf_unset", "comment", "Testing"),
					resource.TestCheckResourceAttr("citrixadc_vpnurlaction.tf_unset", "ssotype", "unifiedgateway"),
					resource.TestCheckResourceAttr("citrixadc_vpnurlaction.tf_unset", "vservername", "vserver1"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from
				// the appliance) reverts to the NITRO defaults, and the implicit
				// post-apply plan must be empty.
				Config: testAccVpnurlaction_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnurlactionExist("citrixadc_vpnurlaction.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnurlaction.tf_unset", "applicationtype", ""),
					resource.TestCheckResourceAttr("citrixadc_vpnurlaction.tf_unset", "clientlessaccess", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_vpnurlaction.tf_unset", "comment", ""),
					resource.TestCheckResourceAttr("citrixadc_vpnurlaction.tf_unset", "ssotype", ""),
					resource.TestCheckResourceAttr("citrixadc_vpnurlaction.tf_unset", "vservername", ""),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckVpnurlactionADCValue("tf_test_vpnurlaction_unset", "clientlessaccess", "OFF"),
				),
			},
		},
	})
}

// testAccCheckVpnurlactionADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted it.
func testAccCheckVpnurlactionADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Vpnurlaction.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("vpnurlaction %s not found on appliance", name)
		}
		got := fmt.Sprintf("%v", data[attr])
		if got != want {
			return fmt.Errorf("vpnurlaction %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func TestAccVpnurlactionDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnurlactionDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_vpnurlaction.foo", "name", "tf_vpnurlaction"),
					resource.TestCheckResourceAttr("data.citrixadc_vpnurlaction.foo", "actualurl", "https://www.citrix.com"),
					resource.TestCheckResourceAttr("data.citrixadc_vpnurlaction.foo", "linkname", "new_link"),
					resource.TestCheckResourceAttr("data.citrixadc_vpnurlaction.foo", "applicationtype", "CVPN"),
					resource.TestCheckResourceAttr("data.citrixadc_vpnurlaction.foo", "clientlessaccess", "ON"),
					resource.TestCheckResourceAttr("data.citrixadc_vpnurlaction.foo", "comment", "Testing"),
					resource.TestCheckResourceAttr("data.citrixadc_vpnurlaction.foo", "ssotype", "unifiedgateway"),
					resource.TestCheckResourceAttr("data.citrixadc_vpnurlaction.foo", "vservername", "vserver1"),
				),
			},
		},
	})
}
