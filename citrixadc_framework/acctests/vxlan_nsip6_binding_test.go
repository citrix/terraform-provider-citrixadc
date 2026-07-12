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
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testAccVxlan_nsip6_binding_basic = `
	resource "citrixadc_vxlan" "tf_vxlan" {
		vxlanid            = 123
		port               = 33
		dynamicrouting     = "DISABLED"
		ipv6dynamicrouting = "DISABLED"
		innervlantagging   = "ENABLED"
	}
	resource "citrixadc_nsip6" "test_nsip" {
		ipv6address = "2001:db8:100::fb/64"
		type        = "VIP"
		icmp        = "DISABLED"
	}
	resource "citrixadc_vxlan_nsip6_binding" "tf_binding" {
		vxlanid   = citrixadc_vxlan.tf_vxlan.vxlanid
		ipaddress = citrixadc_nsip6.test_nsip.ipv6address
		netmask   = "255.255.255.0"
	}
`

const testAccVxlan_nsip6_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion
	resource "citrixadc_vxlan" "tf_vxlan" {
		vxlanid            = 123
		port               = 33
		dynamicrouting     = "DISABLED"
		ipv6dynamicrouting = "DISABLED"
		innervlantagging   = "ENABLED"
	}
	resource "citrixadc_nsip6" "test_nsip" {
		ipv6address = "2001:db8:100::fb/64"
		type        = "VIP"
		icmp        = "DISABLED"
	}
`

func TestAccVxlan_nsip6_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVxlan_nsip6_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVxlan_nsip6_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVxlan_nsip6_bindingExist("citrixadc_vxlan_nsip6_binding.tf_binding", nil),
				),
			},
			{
				Config: testAccVxlan_nsip6_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVxlan_nsip6_bindingNotExist("citrixadc_vxlan_nsip6_binding.tf_binding", "123,2001:db8:100::fb/64"),
				),
			},
		},
	})
}

func TestAccVxlan_nsip6_binding_import(t *testing.T) {
	const resAddr = "citrixadc_vxlan_nsip6_binding.tf_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVxlan_nsip6_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccVxlan_nsip6_binding_basic},
			{Config: testAccVxlan_nsip6_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{"netmask"}},
		},
	})
}

func testAccCheckVxlan_nsip6_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No vxlan_nsip6_binding id is set")
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

		bindingId := rs.Primary.ID

		idMap, _, err := utils.ParseIdString(bindingId, []string{"vxlanid", "ipaddress"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", bindingId, err)
		}
		vxlanid := idMap["vxlanid"]
		ipaddress := idMap["ipaddress"]

		findParams := service.FindParams{
			ResourceType:             "vxlan_nsip6_binding",
			ResourceName:             vxlanid,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}
		// Resource is missing
		if len(dataArr) == 0 {
			return fmt.Errorf("Cannot find vlan_nsip_binding %s", bindingId)
		}

		// Iterate through results to find the one with the matching secondIdComponent
		foundIndex := -1
		for i, v := range dataArr {
			if v["ipaddress"].(string) == ipaddress {
				foundIndex = i
				break
			}
		}

		// Resource is missing
		if foundIndex == -1 {
			return fmt.Errorf("Resource missing vxlan_nsip6_binding %s", bindingId)
		}

		return nil
	}
}

func testAccCheckVxlan_nsip6_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		idMap, _, err := utils.ParseIdString(id, []string{"vxlanid", "ipaddress"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", id, err)
		}
		vxlanid := idMap["vxlanid"]
		ipaddress := idMap["ipaddress"]

		findParams := service.FindParams{
			ResourceType:             "vxlan_nsip6_binding",
			ResourceName:             vxlanid,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Resource is missing
		if len(dataArr) == 0 {
			return nil
		}

		// Iterate through results to find the one with the matching secondIdComponent
		foundIndex := -1
		for i, v := range dataArr {
			if v["ipaddress"].(string) == ipaddress {
				foundIndex = i
				break
			}
		}

		// Resource is missing
		if foundIndex == -1 {
			return fmt.Errorf("Resource missing vxlan_nsip6_binding %s", id)
		}

		return nil
	}
}

func testAccCheckVxlan_nsip6_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_vxlan_nsip6_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Vxlan_nsip6_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("vxlan_nsip6_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccVxlan_nsip6_bindingDataSource_basic = `
	resource "citrixadc_vxlan" "tf_vxlan" {
		vxlanid            = 123
		port               = 33
		dynamicrouting     = "DISABLED"
		ipv6dynamicrouting = "DISABLED"
		innervlantagging   = "ENABLED"
	}
	resource "citrixadc_nsip6" "test_nsip" {
		ipv6address = "2001:db8:100::fb/64"
		type        = "VIP"
		icmp        = "DISABLED"
	}
	resource "citrixadc_vxlan_nsip6_binding" "tf_binding" {
		vxlanid   = citrixadc_vxlan.tf_vxlan.vxlanid
		ipaddress = citrixadc_nsip6.test_nsip.ipv6address
		netmask   = "255.255.255.0"
	}

	data "citrixadc_vxlan_nsip6_binding" "tf_binding" {
		vxlanid   = citrixadc_vxlan_nsip6_binding.tf_binding.vxlanid
		ipaddress = citrixadc_vxlan_nsip6_binding.tf_binding.ipaddress
		depends_on = [citrixadc_vxlan_nsip6_binding.tf_binding]
	}
`

func TestAccVxlan_nsip6_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccVxlan_nsip6_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_vxlan_nsip6_binding.tf_binding", "vxlanid", "123"),
					resource.TestCheckResourceAttr("data.citrixadc_vxlan_nsip6_binding.tf_binding", "ipaddress", "2001:db8:100::fb/64"),
				),
			},
		},
	})
}

const testAccVxlan_nsip6_binding_upgrade_basic = `
	resource "citrixadc_vxlan" "tf_vxlan" {
		vxlanid            = 123
		port               = 33
		dynamicrouting     = "DISABLED"
		ipv6dynamicrouting = "DISABLED"
		innervlantagging   = "ENABLED"
	}
	resource "citrixadc_nsip6" "test_nsip" {
		ipv6address = "2001:db8:100::fb/64"
		type        = "VIP"
		icmp        = "DISABLED"
	}
	resource "citrixadc_vxlan_nsip6_binding" "tf_binding" {
		vxlanid   = citrixadc_vxlan.tf_vxlan.vxlanid
		ipaddress = citrixadc_nsip6.test_nsip.ipv6address
		netmask   = "255.255.255.0"
	}
`

func TestAccVxlan_nsip6_binding_sdkv2StateUpgrade(t *testing.T) {
	resourceAddr := "citrixadc_vxlan_nsip6_binding.tf_binding"
	legacyId := "123,2001:db8:100::fb/64"
	newId := "vxlanid:123,ipaddress:2001%3Adb8%3A100%3A%3Afb%2F64"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckVxlan_nsip6_bindingDestroy,
		Steps: []resource.TestStep{
			// Step 1: create with the last SDK v2 release, writing legacy-id state.
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.2.0",
					},
				},
				Config: testAccVxlan_nsip6_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVxlan_nsip6_bindingExist(resourceAddr, nil),
					resource.TestCheckResourceAttr(resourceAddr, "id", legacyId),
				),
			},
			// Step 2: refresh/plan/apply the legacy-id state through the current framework
			// provider. Read exercises ParseIdString on the legacy id and recomputes the
			// canonical new-format id.
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccVxlan_nsip6_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVxlan_nsip6_bindingExist(resourceAddr, nil),
					resource.TestCheckResourceAttr(resourceAddr, "id", newId),
				),
			},
		},
	})
}
