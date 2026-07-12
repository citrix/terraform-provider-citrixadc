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

const testAccNetbridge_iptunnel_binding_basic = `
	resource "citrixadc_vxlanvlanmap" "tf_vxlanvlanmp" {
		name = "tf_vxlanvlanmp"
	}
	resource "citrixadc_nsip" "nsip" {
		ipaddress = "2.2.2.1"
		type      = "VIP"
		netmask   = "255.255.255.0"
	}
	resource "citrixadc_netbridge" "tf_netbridge" {
		name         = "tf_netbridge"
		vxlanvlanmap = citrixadc_vxlanvlanmap.tf_vxlanvlanmp.name
	}
	resource "citrixadc_iptunnel" "tf_iptunnel" {
		name             = "tf_iptunnel"
		remote           = "66.0.0.11"
		remotesubnetmask = "255.255.255.255"
		local            = citrixadc_nsip.nsip.ipaddress
		protocol         = "GRE"
	}
	resource "citrixadc_netbridge_iptunnel_binding" "tf_binding" {
		name   = citrixadc_netbridge.tf_netbridge.name
		tunnel = citrixadc_iptunnel.tf_iptunnel.name
	}
`

const testAccNetbridge_iptunnel_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion
	resource "citrixadc_vxlanvlanmap" "tf_vxlanvlanmp" {
		name = "tf_vxlanvlanmp"
	}
	resource "citrixadc_nsip" "nsip" {
		ipaddress = "2.2.2.1"
		type      = "VIP"
		netmask   = "255.255.255.0"
	}
	resource "citrixadc_netbridge" "tf_netbridge" {
		name         = "tf_netbridge"
		vxlanvlanmap = citrixadc_vxlanvlanmap.tf_vxlanvlanmp.name
	}
	resource "citrixadc_iptunnel" "tf_iptunnel" {
		name             = "tf_iptunnel"
		remote           = "66.0.0.11"
		remotesubnetmask = "255.255.255.255"
		local            = citrixadc_nsip.nsip.ipaddress
		protocol         = "GRE"
	}
`

func TestAccNetbridge_iptunnel_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNetbridge_iptunnel_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNetbridge_iptunnel_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetbridge_iptunnel_bindingExist("citrixadc_netbridge_iptunnel_binding.tf_binding", nil),
				),
			},
			{
				Config: testAccNetbridge_iptunnel_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetbridge_iptunnel_bindingNotExist("citrixadc_netbridge_iptunnel_binding.tf_binding", "tf_netbridge,tf_iptunnel"),
				),
			},
		},
	})
}

func TestAccNetbridge_iptunnel_binding_import(t *testing.T) {
	const resAddr = "citrixadc_netbridge_iptunnel_binding.tf_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNetbridge_iptunnel_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccNetbridge_iptunnel_binding_basic},
			{Config: testAccNetbridge_iptunnel_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
		},
	})
}

func testAccCheckNetbridge_iptunnel_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No netbridge_iptunnel_binding id is set")
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

		idMap, _, err := utils.ParseIdString(bindingId, []string{"name", "tunnel"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", bindingId, err)
		}
		name := idMap["name"]
		tunnel := idMap["tunnel"]

		findParams := service.FindParams{
			ResourceType:             "netbridge_iptunnel_binding",
			ResourceName:             name,
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
			if v["tunnel"].(string) == tunnel {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("netbridge_iptunnel_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckNetbridge_iptunnel_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}
		idSlice := strings.SplitN(id, ",", 2)

		name := idSlice[0]
		tunnel := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "netbridge_iptunnel_binding",
			ResourceName:             name,
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
			if v["tunnel"].(string) == tunnel {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("netbridge_iptunnel_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckNetbridge_iptunnel_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_netbridge_iptunnel_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Netbridge_iptunnel_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("netbridge_iptunnel_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccNetbridge_iptunnel_bindingDataSource_basic = `
	resource "citrixadc_vxlanvlanmap" "tf_vxlanvlanmp" {
		name = "tf_vxlanvlanmp"
	}
	resource "citrixadc_nsip" "nsip" {
		ipaddress = "2.2.2.1"
		type      = "VIP"
		netmask   = "255.255.255.0"
	}
	resource "citrixadc_netbridge" "tf_netbridge" {
		name         = "tf_netbridge"
		vxlanvlanmap = citrixadc_vxlanvlanmap.tf_vxlanvlanmp.name
	}
	resource "citrixadc_iptunnel" "tf_iptunnel" {
		name             = "tf_iptunnel"
		remote           = "66.0.0.11"
		remotesubnetmask = "255.255.255.255"
		local            = citrixadc_nsip.nsip.ipaddress
		protocol         = "GRE"
	}
	resource "citrixadc_netbridge_iptunnel_binding" "tf_binding" {
		name   = citrixadc_netbridge.tf_netbridge.name
		tunnel = citrixadc_iptunnel.tf_iptunnel.name
	}

	data "citrixadc_netbridge_iptunnel_binding" "tf_binding" {
		name   = citrixadc_netbridge.tf_netbridge.name
		tunnel = citrixadc_iptunnel.tf_iptunnel.name
		depends_on = [citrixadc_netbridge_iptunnel_binding.tf_binding]
	}
`

func TestAccNetbridge_iptunnel_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNetbridge_iptunnel_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_netbridge_iptunnel_binding.tf_binding", "name", "tf_netbridge"),
					resource.TestCheckResourceAttr("data.citrixadc_netbridge_iptunnel_binding.tf_binding", "tunnel", "tf_iptunnel"),
					resource.TestCheckResourceAttrSet("data.citrixadc_netbridge_iptunnel_binding.tf_binding", "id"),
				),
			},
		},
	})
}

// Config for the SDK v2 -> Framework state-upgrade test. Reuses the _basic
// values and is valid under BOTH the last SDK v2 release (2.2.0) schema and the
// current Framework schema (uses only SDK v2 attribute names).
const testAccNetbridge_iptunnel_binding_upgrade_basic = `
	resource "citrixadc_vxlanvlanmap" "tf_vxlanvlanmp" {
		name = "tf_vxlanvlanmp"
	}
	resource "citrixadc_nsip" "nsip" {
		ipaddress = "2.2.2.1"
		type      = "VIP"
		netmask   = "255.255.255.0"
	}
	resource "citrixadc_netbridge" "tf_netbridge" {
		name         = "tf_netbridge"
		vxlanvlanmap = citrixadc_vxlanvlanmap.tf_vxlanvlanmp.name
	}
	resource "citrixadc_iptunnel" "tf_iptunnel" {
		name             = "tf_iptunnel"
		remote           = "66.0.0.11"
		remotesubnetmask = "255.255.255.255"
		local            = citrixadc_nsip.nsip.ipaddress
		protocol         = "GRE"
	}
	resource "citrixadc_netbridge_iptunnel_binding" "tf_binding" {
		name   = citrixadc_netbridge.tf_netbridge.name
		tunnel = citrixadc_iptunnel.tf_iptunnel.name
	}
`

// TestAccNetbridge_iptunnel_binding_sdkv2StateUpgrade verifies that state written
// by the last SDK v2 release (legacy comma-joined id) is transparently upgraded by
// the current Framework provider. Step 1 creates the binding with citrix/citrixadc
// 2.2.0 (legacy id "name,tunnel"); step 2 refreshes/plans the same config through
// the current Framework provider, whose Read parses the legacy id and recomputes
// it to the new "name:<v>,tunnel:<v>" format (SetAttrFromGet).
func TestAccNetbridge_iptunnel_binding_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckNetbridge_iptunnel_bindingDestroy,
		Steps: []resource.TestStep{
			{
				// Step 1: create with the last SDK v2 release, writing the legacy id.
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.2.0",
					},
				},
				Config: testAccNetbridge_iptunnel_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetbridge_iptunnel_bindingExist("citrixadc_netbridge_iptunnel_binding.tf_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_netbridge_iptunnel_binding.tf_binding", "id", "tf_netbridge,tf_iptunnel"),
				),
			},
			{
				// Step 2: refresh/apply the same config through the current Framework
				// provider. Read exercises ParseIdString on the legacy id, then
				// recomputes the id to the new key:value canonical format.
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccNetbridge_iptunnel_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetbridge_iptunnel_bindingExist("citrixadc_netbridge_iptunnel_binding.tf_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_netbridge_iptunnel_binding.tf_binding", "id", "name:tf_netbridge,tunnel:tf_iptunnel"),
				),
			},
		},
	})
}
