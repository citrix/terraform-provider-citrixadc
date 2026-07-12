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
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testAccBridgegroup_nsip_binding_basic = `
	resource "citrixadc_bridgegroup" "tf_bridgegroup" {
		bridgegroup_id     = 2
		dynamicrouting     = "DISABLED"
		ipv6dynamicrouting = "DISABLED"
	}
	resource "citrixadc_nsip" "nsip" {
		ipaddress = "2.2.2.3"
		type      = "VIP"
		netmask   = "255.255.255.0"
	}
	resource "citrixadc_bridgegroup_nsip_binding" "tf_binding" {
		bridgegroup_id = citrixadc_bridgegroup.tf_bridgegroup.bridgegroup_id
		ipaddress      = citrixadc_nsip.nsip.ipaddress
		netmask        = citrixadc_nsip.nsip.netmask
	}
`

const testAccBridgegroup_nsip_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion
	resource "citrixadc_bridgegroup" "tf_bridgegroup" {
		bridgegroup_id     = 2
		dynamicrouting     = "DISABLED"
		ipv6dynamicrouting = "DISABLED"
	}
	resource "citrixadc_nsip" "nsip" {
		ipaddress = "2.2.2.3"
		type      = "VIP"
		netmask   = "255.255.255.0"
	}
`

func TestAccBridgegroup_nsip_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckBridgegroup_nsip_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccBridgegroup_nsip_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBridgegroup_nsip_bindingExist("citrixadc_bridgegroup_nsip_binding.tf_binding", nil),
				),
			},
			{
				Config: testAccBridgegroup_nsip_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBridgegroup_nsip_bindingNotExist("citrixadc_bridgegroup_nsip_binding.tf_binding", "2,2.2.2.3"),
				),
			},
		},
	})
}

func testAccCheckBridgegroup_nsip_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No bridgegroup_nsip_binding id is set")
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

		idMap, _, err := utils.ParseIdString(bindingId, []string{"bridgegroup_id", "ipaddress"}, nil)
		if err != nil {
			return err
		}
		bridgegroup_id := idMap["bridgegroup_id"]
		ipaddress := idMap["ipaddress"]

		findParams := service.FindParams{
			ResourceType:             "bridgegroup_nsip_binding",
			ResourceName:             bridgegroup_id,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the matching secondIdComponent
		found := false
		for _, v := range dataArr {
			if v["ipaddress"].(string) == ipaddress {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("bridgegroup_nsip_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckBridgegroup_nsip_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}
		idMap, _, err := utils.ParseIdString(id, []string{"bridgegroup_id", "ipaddress"}, nil)
		if err != nil {
			return err
		}
		bridgegroup_id := idMap["bridgegroup_id"]
		ipaddress := idMap["ipaddress"]

		findParams := service.FindParams{
			ResourceType:             "bridgegroup_nsip_binding",
			ResourceName:             bridgegroup_id,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to hopefully not find the one with the matching secondIdComponent
		found := false
		for _, v := range dataArr {
			if v["ipaddress"].(string) == ipaddress {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("bridgegroup_nsip_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckBridgegroup_nsip_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_bridgegroup_nsip_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Bridgegroup_nsip_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("bridgegroup_nsip_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccBridgegroup_nsip_bindingDataSource_basic = `
	resource "citrixadc_bridgegroup" "tf_bridgegroup" {
		bridgegroup_id     = 2
		dynamicrouting     = "DISABLED"
		ipv6dynamicrouting = "DISABLED"
	}
	resource "citrixadc_nsip" "nsip" {
		ipaddress = "2.2.2.3"
		type      = "VIP"
		netmask   = "255.255.255.0"
	}
	resource "citrixadc_bridgegroup_nsip_binding" "tf_binding" {
		bridgegroup_id = citrixadc_bridgegroup.tf_bridgegroup.bridgegroup_id
		ipaddress      = citrixadc_nsip.nsip.ipaddress
		netmask        = citrixadc_nsip.nsip.netmask
	}

	data "citrixadc_bridgegroup_nsip_binding" "tf_binding" {
		bridgegroup_id = citrixadc_bridgegroup_nsip_binding.tf_binding.bridgegroup_id
		ipaddress      = citrixadc_bridgegroup_nsip_binding.tf_binding.ipaddress
		depends_on     = [citrixadc_bridgegroup_nsip_binding.tf_binding]
	}
`

func TestAccbridgegroup_nsip_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccBridgegroup_nsip_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_bridgegroup_nsip_binding.tf_binding", "bridgegroup_id", "2"),
					resource.TestCheckResourceAttr("data.citrixadc_bridgegroup_nsip_binding.tf_binding", "ipaddress", "2.2.2.3"),
					resource.TestCheckResourceAttr("data.citrixadc_bridgegroup_nsip_binding.tf_binding", "netmask", "255.255.255.0"),
				),
			},
		},
	})
}

// testAccBridgegroup_nsip_binding_upgrade_basic is valid under BOTH the last
// SDK v2 release (2.2.0) and the current Framework schema. It reuses the same
// resource labels and values as testAccBridgegroup_nsip_binding_basic so the
// existing Exist/Destroy helpers and addresses match.
const testAccBridgegroup_nsip_binding_upgrade_basic = `
	resource "citrixadc_bridgegroup" "tf_bridgegroup" {
		bridgegroup_id     = 2
		dynamicrouting     = "DISABLED"
		ipv6dynamicrouting = "DISABLED"
	}
	resource "citrixadc_nsip" "nsip" {
		ipaddress = "2.2.2.3"
		type      = "VIP"
		netmask   = "255.255.255.0"
	}
	resource "citrixadc_bridgegroup_nsip_binding" "tf_binding" {
		bridgegroup_id = citrixadc_bridgegroup.tf_bridgegroup.bridgegroup_id
		ipaddress      = citrixadc_nsip.nsip.ipaddress
		netmask        = citrixadc_nsip.nsip.netmask
	}
`

// TestAccBridgegroup_nsip_binding_sdkv2StateUpgrade verifies that state written
// by the last SDK v2 release (legacy comma-joined id) is upgraded correctly by
// the current Framework provider. Step 1 creates the binding with citrix/citrixadc
// 2.2.0 (legacy id "2,2.2.2.3"). Step 2 re-plans/applies the SAME config through
// the current (Framework) provider, whose Read (SetAttrFromGet) recomputes the id
// into the new key:value format ("bridgegroup_id:2,ipaddress:2.2.2.3").
func TestAccBridgegroup_nsip_binding_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckBridgegroup_nsip_bindingDestroy,
		Steps: []resource.TestStep{
			{
				// Create with the last SDK v2 release from the registry.
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.2.0",
					},
				},
				Config: testAccBridgegroup_nsip_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBridgegroup_nsip_bindingExist("citrixadc_bridgegroup_nsip_binding.tf_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_bridgegroup_nsip_binding.tf_binding", "id", "2,2.2.2.3"),
				),
			},
			{
				// Refresh/plan/apply the legacy-id state through the current provider.
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccBridgegroup_nsip_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBridgegroup_nsip_bindingExist("citrixadc_bridgegroup_nsip_binding.tf_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_bridgegroup_nsip_binding.tf_binding", "id", "bridgegroup_id:2,ipaddress:2.2.2.3"),
				),
			},
		},
	})
}

func TestAccBridgegroup_nsip_binding_import(t *testing.T) {
	const resAddr = "citrixadc_bridgegroup_nsip_binding.tf_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckBridgegroup_nsip_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccBridgegroup_nsip_binding_basic},
			{Config: testAccBridgegroup_nsip_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
		},
	})
}
