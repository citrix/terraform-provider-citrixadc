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

const testAccVlan_nsip6_binding_basic = `

	resource "citrixadc_vlan" "tf_vlan" {
		vlanid = 2
		aliasname = "VLAN"
	}
	resource "citrixadc_nsip6" "tf_nsip6" {
		ipv6address = "2001::a/96"
		type = "VIP"
	}

	resource "citrixadc_vlan_nsip6_binding" "tf_vlan_nsip6_binding" {
		vlanid    = citrixadc_vlan.tf_vlan.vlanid
		ipaddress = citrixadc_nsip6.tf_nsip6.ipv6address
	}
`

const testAccVlan_nsip6_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion

	resource "citrixadc_vlan" "tf_vlan" {
		vlanid = 2
		aliasname = "VLAN"
	}
	resource "citrixadc_nsip6" "tf_nsip6" {
		ipv6address = "2001::a/96"
		type = "VIP"
	}
`

func TestAccVlan_nsip6_binding_import(t *testing.T) {
	const resAddr = "citrixadc_vlan_nsip6_binding.tf_vlan_nsip6_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVlan_nsip6_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccVlan_nsip6_binding_basic},
			{Config: testAccVlan_nsip6_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
		},
	})
}

func TestAccVlan_nsip6_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVlan_nsip6_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVlan_nsip6_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlan_nsip6_bindingExist("citrixadc_vlan_nsip6_binding.tf_vlan_nsip6_binding", nil),
				),
			},
			{
				Config: testAccVlan_nsip6_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlan_nsip6_bindingNotExist("citrixadc_vlan_nsip6_binding.tf_vlan_nsip6_binding", "2,2001::a/96"),
				),
			},
		},
	})
}

func testAccCheckVlan_nsip6_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No vlan_nsip6_binding id is set")
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

		idMap, _, err := utils.ParseIdString(bindingId, []string{"vlanid", "ipaddress"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", bindingId, err)
		}
		vlanid := idMap["vlanid"]
		ipaddress := idMap["ipaddress"]

		findParams := service.FindParams{
			ResourceType:             "vlan_nsip6_binding",
			ResourceName:             vlanid,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the matching ipaddress
		found := false
		for _, v := range dataArr {
			if v["ipaddress"].(string) == ipaddress {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("vlan_nsip6_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckVlan_nsip6_bindingNotExist(n string, id string) resource.TestCheckFunc {
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

		vlanid := idSlice[0]
		ipaddress := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "vlan_nsip6_binding",
			ResourceName:             vlanid,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to hopefully not find the one with the matching ipaddress
		found := false
		for _, v := range dataArr {
			if v["ipaddress"].(string) == ipaddress {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("vlan_nsip6_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckVlan_nsip6_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_vlan_nsip6_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Vlan_nsip6_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("vlan_nsip6_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccVlan_nsip6_bindingDataSource_basic = `
	resource "citrixadc_vlan" "tf_vlan" {
		vlanid = 2
		aliasname = "VLAN"
	}
	resource "citrixadc_nsip6" "tf_nsip6" {
		ipv6address = "2001::a/96"
		type = "VIP"
	}

	resource "citrixadc_vlan_nsip6_binding" "tf_vlan_nsip6_binding" {
		vlanid    = citrixadc_vlan.tf_vlan.vlanid
		ipaddress = citrixadc_nsip6.tf_nsip6.ipv6address
	}

	data "citrixadc_vlan_nsip6_binding" "tf_vlan_nsip6_binding" {
		vlanid      = citrixadc_vlan.tf_vlan.vlanid
		ipaddress   = citrixadc_nsip6.tf_nsip6.ipv6address
		depends_on = [citrixadc_vlan_nsip6_binding.tf_vlan_nsip6_binding]
	}
`

// Config for the SDK v2 -> Framework state-upgrade test. Reuses the _basic
// config values (same terraform resource labels) so the Exist/Destroy helpers
// and resource addresses match. It is valid under both the SDK v2 2.2.0 schema
// and the current Framework schema.
const testAccVlan_nsip6_binding_upgrade_basic = `
	resource "citrixadc_vlan" "tf_vlan" {
		vlanid = 2
		aliasname = "VLAN"
	}
	resource "citrixadc_nsip6" "tf_nsip6" {
		ipv6address = "2001::a/96"
		type = "VIP"
	}

	resource "citrixadc_vlan_nsip6_binding" "tf_vlan_nsip6_binding" {
		vlanid    = citrixadc_vlan.tf_vlan.vlanid
		ipaddress = citrixadc_nsip6.tf_nsip6.ipv6address
	}
`

// TestAccVlan_nsip6_binding_sdkv2StateUpgrade verifies that state written by the
// last SDK v2 release (legacy comma-joined id "vlanid,ipaddress") is transparently
// upgraded by the current Framework provider. Step 1 creates the binding with
// citrix/citrixadc 2.2.0 (legacy id "2,2001::a/96"); step 2 refreshes/plans the
// same config through the current Framework provider, whose Read parses the legacy
// id and recomputes it to the new "vlanid:<v>,ipaddress:<v>" canonical format
// (SetAttrFromGet).
func TestAccVlan_nsip6_binding_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckVlan_nsip6_bindingDestroy,
		Steps: []resource.TestStep{
			{
				// Step 1: create with the last SDK v2 release, writing the legacy id.
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.2.0",
					},
				},
				Config: testAccVlan_nsip6_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlan_nsip6_bindingExist("citrixadc_vlan_nsip6_binding.tf_vlan_nsip6_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_vlan_nsip6_binding.tf_vlan_nsip6_binding", "id", "2,2001::a/96"),
				),
			},
			{
				// Step 2: refresh/apply the same config through the current Framework
				// provider. Read exercises ParseIdString on the legacy id, then
				// recomputes the id to the new key:value canonical format.
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccVlan_nsip6_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlan_nsip6_bindingExist("citrixadc_vlan_nsip6_binding.tf_vlan_nsip6_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_vlan_nsip6_binding.tf_vlan_nsip6_binding", "id", "vlanid:2,ipaddress:2001%3A%3Aa%2F96"),
				),
			},
		},
	})
}

func TestAccVlan_nsip6_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccVlan_nsip6_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_vlan_nsip6_binding.tf_vlan_nsip6_binding", "vlanid", "2"),
					resource.TestCheckResourceAttr("data.citrixadc_vlan_nsip6_binding.tf_vlan_nsip6_binding", "ipaddress", "2001::a/96"),
				),
			},
		},
	})
}
