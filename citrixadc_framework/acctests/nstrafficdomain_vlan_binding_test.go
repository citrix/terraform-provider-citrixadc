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

const testAccNstrafficdomain_vlan_binding_basic = `
	resource "citrixadc_nstrafficdomain" "tf_trafficdomain" {
		td        = 2
		aliasname = "tf_trafficdomain"
		vmac      = "DISABLED"
	}
	resource "citrixadc_vlan" "tf_vlan" {
		vlanid    = 20
		aliasname = "Management VLAN"
	}
	resource "citrixadc_nstrafficdomain_vlan_binding" "tf_binding" {
		td   = citrixadc_nstrafficdomain.tf_trafficdomain.td
		vlan = citrixadc_vlan.tf_vlan.vlanid
	}
`

const testAccNstrafficdomain_vlan_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion
	resource "citrixadc_nstrafficdomain" "tf_trafficdomain" {
		td        = 2
		aliasname = "tf_trafficdomain"
		vmac      = "DISABLED"
	}
	resource "citrixadc_vlan" "tf_vlan" {
		vlanid    = 20
		aliasname = "Management VLAN"
	}
`

const testAccNstrafficdomain_vlan_bindingDataSource_basic = `
	resource "citrixadc_nstrafficdomain" "tf_trafficdomain" {
		td        = 2
		aliasname = "tf_trafficdomain"
		vmac      = "DISABLED"
	}
	resource "citrixadc_vlan" "tf_vlan" {
		vlanid    = 20
		aliasname = "Management VLAN"
	}
	resource "citrixadc_nstrafficdomain_vlan_binding" "tf_binding" {
		td   = citrixadc_nstrafficdomain.tf_trafficdomain.td
		vlan = citrixadc_vlan.tf_vlan.vlanid
	}

	data "citrixadc_nstrafficdomain_vlan_binding" "tf_binding" {
		td   = citrixadc_nstrafficdomain_vlan_binding.tf_binding.td
		vlan = citrixadc_nstrafficdomain_vlan_binding.tf_binding.vlan
	}
`

func TestAccNstrafficdomain_vlan_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNstrafficdomain_vlan_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNstrafficdomain_vlan_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNstrafficdomain_vlan_bindingExist("citrixadc_nstrafficdomain_vlan_binding.tf_binding", nil),
				),
			},
			{
				Config: testAccNstrafficdomain_vlan_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNstrafficdomain_vlan_bindingNotExist("citrixadc_nstrafficdomain_vlan_binding.tf_binding", "2,20"),
				),
			},
		},
	})
}

func TestAccNstrafficdomain_vlan_binding_import(t *testing.T) {
	const resAddr = "citrixadc_nstrafficdomain_vlan_binding.tf_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNstrafficdomain_vlan_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccNstrafficdomain_vlan_binding_basic},
			{Config: testAccNstrafficdomain_vlan_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
		},
	})
}

func testAccCheckNstrafficdomain_vlan_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No nstrafficdomain_vlan_binding id is set")
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

		idMap, _, err := utils.ParseIdString(bindingId, []string{"td", "vlan"}, nil)
		if err != nil {
			return err
		}
		td := idMap["td"]
		vlan := idMap["vlan"]

		findParams := service.FindParams{
			ResourceType:             "nstrafficdomain_vlan_binding",
			ResourceName:             td,
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
			if fmt.Sprintf("%v", v["vlan"]) == vlan {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("nstrafficdomain_vlan_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckNstrafficdomain_vlan_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}
		idMap, _, err := utils.ParseIdString(id, []string{"td", "vlan"}, nil)
		if err != nil {
			return err
		}
		td := idMap["td"]
		vlan := idMap["vlan"]

		findParams := service.FindParams{
			ResourceType:             "nstrafficdomain_vlan_binding",
			ResourceName:             td,
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
			if fmt.Sprintf("%v", v["vlan"]) == vlan {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("nstrafficdomain_vlan_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckNstrafficdomain_vlan_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_nstrafficdomain_vlan_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Nstrafficdomain_vlan_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("nstrafficdomain_vlan_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccNstrafficdomain_vlan_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccNstrafficdomain_vlan_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_nstrafficdomain_vlan_binding.tf_binding", "td", "2"),
					resource.TestCheckResourceAttr("data.citrixadc_nstrafficdomain_vlan_binding.tf_binding", "vlan", "20"),
				),
			},
		},
	})
}

const testAccNstrafficdomain_vlan_binding_upgrade_basic = `
	resource "citrixadc_nstrafficdomain" "tf_trafficdomain" {
		td        = 2
		aliasname = "tf_trafficdomain"
		vmac      = "DISABLED"
	}
	resource "citrixadc_vlan" "tf_vlan" {
		vlanid    = 20
		aliasname = "Management VLAN"
	}
	resource "citrixadc_nstrafficdomain_vlan_binding" "tf_binding" {
		td   = citrixadc_nstrafficdomain.tf_trafficdomain.td
		vlan = citrixadc_vlan.tf_vlan.vlanid
	}
`

// TestAccNstrafficdomain_vlan_binding_sdkv2StateUpgrade verifies that state written
// by the last SDK v2 release (legacy comma-separated ID) is correctly upgraded when
// the same config is subsequently managed by the current Framework provider. Step 1
// creates the binding with citrix/citrixadc 2.2.0 (writes the legacy id "2,20").
// Step 2 refreshes/plans/applies the same config through the Framework provider,
// exercising ParseIdString on the legacy id; because the Framework recomputes the id
// on Read (SetAttrFromGet re-derives data.Id), the id upgrades to the new
// "key:value" form "td:2,vlan:20".
func TestAccNstrafficdomain_vlan_binding_sdkv2StateUpgrade(t *testing.T) {
	resourceAddr := "citrixadc_nstrafficdomain_vlan_binding.tf_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckNstrafficdomain_vlan_bindingDestroy,
		Steps: []resource.TestStep{
			// Step 1: create with the last SDK v2 release -> state carries the legacy id.
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.2.0",
					},
				},
				Config: testAccNstrafficdomain_vlan_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNstrafficdomain_vlan_bindingExist(resourceAddr, nil),
					resource.TestCheckResourceAttr(resourceAddr, "id", "2,20"),
				),
			},
			// Step 2: refresh/plan/apply the SAME config through the current Framework
			// provider. The legacy-id state is read via ParseIdString and the id is
			// recomputed to the new key:value format.
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccNstrafficdomain_vlan_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNstrafficdomain_vlan_bindingExist(resourceAddr, nil),
					resource.TestCheckResourceAttr(resourceAddr, "id", "td:2,vlan:20"),
				),
			},
		},
	})
}
