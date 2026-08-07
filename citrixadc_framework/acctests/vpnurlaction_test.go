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
