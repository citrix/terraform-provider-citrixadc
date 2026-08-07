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

const testAccVlan_basic_step1 = `
resource "citrixadc_vlan" "tf_vlan" {
    vlanid = 40
    aliasname = "Test alias name"
}
`

const testAccVlan_basic_step2 = `
resource "citrixadc_vlan" "tf_vlan" {
    vlanid = 40
    aliasname = "Test alias name 2"
}
`

func TestAccVlan_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVlanDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVlan_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanExist("citrixadc_vlan.tf_vlan", nil),
				),
			},
			{
				Config: testAccVlan_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlanExist("citrixadc_vlan.tf_vlan", nil),
				),
			},
		},
	})
}

func testAccCheckVlanExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No vlan name is set")
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
		data, err := client.FindResource(service.Vlan.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("vlan %s not found", n)
		}

		return nil
	}
}

func testAccCheckVlanDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_vlan" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Vlan.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("vlan %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccVlan_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_vlan.tf_vlan"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVlanDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVlan_basic_step1,
				Check:  resource.ComposeTestCheckFunc(testAccCheckVlanExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Vlan.Type(), "40"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccVlan_basic_step1,
				Check:  resource.ComposeTestCheckFunc(testAccCheckVlanExist(resAddr, nil)),
			},
		},
	})
}

func TestAccVlan_import(t *testing.T) {
	const resAddr = "citrixadc_vlan.tf_vlan"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVlanDestroy,
		Steps: []resource.TestStep{
			{Config: testAccVlan_basic_step1},
			{
				Config:                  testAccVlan_basic_step1,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

const testAccVlanDataSource_basic = `
resource "citrixadc_vlan" "tf_vlan" {
    vlanid = 40
    aliasname = "Test alias name"
}

data "citrixadc_vlan" "tf_vlan" {
    vlanid = citrixadc_vlan.tf_vlan.vlanid
}
`

func TestAccVlan_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckVlanDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccVlan_basic_step1,
				Check:  resource.ComposeTestCheckFunc(testAccCheckVlanExist("citrixadc_vlan.tf_vlan", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccVlan_basic_step1,
				Check:                    resource.ComposeTestCheckFunc(testAccCheckVlanExist("citrixadc_vlan.tf_vlan", nil)),
			},
		},
	})
}

func TestAccVlanDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccVlanDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_vlan.tf_vlan", "vlanid", "40"),
					resource.TestCheckResourceAttr("data.citrixadc_vlan.tf_vlan", "aliasname", "Test alias name"),
				),
			},
		},
	})
}
