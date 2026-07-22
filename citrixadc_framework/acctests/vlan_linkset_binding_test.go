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

// NOTE: Despite the resource name, vlan_linkset_binding has NO `linkset` attribute.
// It binds an INTERFACE (ifnum) to a VLAN. The user-facing VLAN id is exposed as
// `vlanid` (Int64) to avoid colliding with the synthetic Terraform `id`. The HCL
// below therefore uses `vlanid` and `ifnum` -- there is intentionally no `linkset`.
//
// IMPORTANT: ifnum is TESTBED-SPECIFIC. Binding/tagging an interface to a VLAN can
// disrupt the appliance's networking. Replace the TODO_PLACEHOLDER below with a
// FREE interface on the target testbed (e.g. "1/1") before running these tests.
//
// Participating entity: a citrixadc_vlan (vlanid = 100) is created first and the
// interface is then bound to it. The vlan config is lifted from vlan_test.go.

const testAccVlan_linkset_binding_basic_step1 = `
	resource "citrixadc_vlan" "tf_vlan" {
		vlanid    = 100
		aliasname = "Test alias name"
	}

	resource "citrixadc_vlan_linkset_binding" "tf_vlan_linkset_binding" {
		vlanid     = citrixadc_vlan.tf_vlan.vlanid
		ifnum      = "1/1" # e.g. "1/1" -- a FREE interface on the target testbed
		tagged     = true
		depends_on = [citrixadc_vlan.tf_vlan]
	}

`

const testAccVlan_linkset_binding_basic_step2 = `
	# Keep the parent vlan without the actual binding to verify proper deletion
	resource "citrixadc_vlan" "tf_vlan" {
		vlanid    = 100
		aliasname = "Test alias name"
	}

`

func TestAccVlan_linkset_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVlan_linkset_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVlan_linkset_binding_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlan_linkset_bindingExist("citrixadc_vlan_linkset_binding.tf_vlan_linkset_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_vlan_linkset_binding.tf_vlan_linkset_binding", "vlanid", "100"),
					resource.TestCheckResourceAttr("citrixadc_vlan_linkset_binding.tf_vlan_linkset_binding", "ifnum", "1/1"),
					resource.TestCheckResourceAttr("citrixadc_vlan_linkset_binding.tf_vlan_linkset_binding", "tagged", "true"),
				),
			},
			{
				Config: testAccVlan_linkset_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlan_linkset_bindingNotExist("citrixadc_vlan_linkset_binding.tf_vlan_linkset_binding", "vlanid:100,ifnum:1%2F1"),
				),
			},
		},
	})
}

// vlanLinksetBindingAggregateReadForTest reads the bound interfaces via the
// AGGREGATE parent endpoint (vlan_binding/<vlanid>) and flattens the nested
// "vlan_linkset_binding" arrays. The direct vlan_linkset_binding endpoint can
// return a keyless empty body, so the check helpers must use the aggregate
// parent endpoint to locate the binding (same quirk as lsnpool_lsnip_binding).
// This mirrors vlan_linkset_bindingAggregateRead in the resource implementation.
func vlanLinksetBindingAggregateReadForTest(client *service.NitroClient, vlanid string) ([]map[string]interface{}, error) {
	findParams := service.FindParams{
		ResourceType:             "vlan_binding",
		ResourceName:             vlanid,
		ResourceMissingErrorCode: 258,
	}
	parentArr, err := client.FindResourceArrayWithParams(findParams)
	if err != nil {
		return nil, err
	}

	rows := make([]map[string]interface{}, 0)
	for _, parent := range parentArr {
		// Verified live: an interface bound via the vlan_linkset_binding endpoint is
		// stored by the appliance as a "vlan_interface_binding" row in the aggregate
		// vlan_binding/<id> response (there is no "vlan_linkset_binding" nested array).
		// Mirror the resource's aggregate read.
		nested, ok := parent["vlan_interface_binding"]
		if !ok || nested == nil {
			continue
		}
		nestedArr, ok := nested.([]interface{})
		if !ok {
			continue
		}
		for _, item := range nestedArr {
			if m, ok := item.(map[string]interface{}); ok {
				rows = append(rows, m)
			}
		}
	}
	return rows, nil
}

func testAccCheckVlan_linkset_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No vlan_linkset_binding id is set")
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

		// Composite ID: vlanid:UrlEncode(value),ifnum:UrlEncode(value)
		// Legacy attr order is [vlanid, ifnum].
		idMap, _, err := utils.ParseIdString(rs.Primary.ID, []string{"vlanid", "ifnum"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", rs.Primary.ID, err)
		}

		vlanid := idMap["vlanid"]
		ifnum := idMap["ifnum"]

		dataArr, err := vlanLinksetBindingAggregateReadForTest(client, vlanid)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the matching ifnum (member key)
		found := false
		for _, v := range dataArr {
			if val, ok := v["ifnum"].(string); ok && val == ifnum {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("vlan_linkset_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckVlan_linkset_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		idMap, _, err := utils.ParseIdString(id, []string{"vlanid", "ifnum"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", id, err)
		}

		vlanid := idMap["vlanid"]
		ifnum := idMap["ifnum"]

		dataArr, err := vlanLinksetBindingAggregateReadForTest(client, vlanid)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to hopefully not find the one with the matching ifnum
		found := false
		for _, v := range dataArr {
			if val, ok := v["ifnum"].(string); ok && val == ifnum {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("vlan_linkset_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckVlan_linkset_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_vlan_linkset_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		idMap, _, err := utils.ParseIdString(rs.Primary.ID, []string{"vlanid", "ifnum"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", rs.Primary.ID, err)
		}

		vlanid := idMap["vlanid"]
		ifnum := idMap["ifnum"]

		dataArr, err := vlanLinksetBindingAggregateReadForTest(client, vlanid)

		// If the parent vlan itself is gone, the binding is certainly gone.
		if err != nil {
			continue
		}

		for _, v := range dataArr {
			if val, ok := v["ifnum"].(string); ok && val == ifnum {
				return fmt.Errorf("vlan_linkset_binding %s still exists", rs.Primary.ID)
			}
		}
	}

	return nil
}

const testAccVlan_linkset_binding_DataSource_basic = `
	resource "citrixadc_vlan" "tf_vlan" {
		vlanid    = 100
		aliasname = "Test alias name"
	}

	resource "citrixadc_vlan_linkset_binding" "tf_vlan_linkset_binding" {
		vlanid     = citrixadc_vlan.tf_vlan.vlanid
		ifnum      = "1/1" # e.g. "1/1" -- a FREE interface on the target testbed
		tagged     = true
		depends_on = [citrixadc_vlan.tf_vlan]
	}

	data "citrixadc_vlan_linkset_binding" "tf_vlan_linkset_binding" {
		vlanid     = citrixadc_vlan.tf_vlan.vlanid
		ifnum      = "1/1" # must match the bound interface above
		depends_on = [citrixadc_vlan_linkset_binding.tf_vlan_linkset_binding]
	}

`

func TestAccVlan_linkset_binding_import(t *testing.T) {
	const resAddr = "citrixadc_vlan_linkset_binding.tf_vlan_linkset_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVlan_linkset_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVlan_linkset_binding_basic_step1,
			},
			{
				Config:                  testAccVlan_linkset_binding_basic_step1,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccVlan_linkset_binding_DataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVlan_linkset_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVlan_linkset_binding_DataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_vlan_linkset_binding.tf_vlan_linkset_binding", "vlanid", "100"),
					resource.TestCheckResourceAttr("data.citrixadc_vlan_linkset_binding.tf_vlan_linkset_binding", "ifnum", "1/1"),
				),
			},
		},
	})
}
