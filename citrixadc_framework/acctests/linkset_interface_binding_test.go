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

// linkset_interface_binding binds a physical interface to a linkset.
//
// IMPORTANT - TESTBED-SPECIFIC INTERFACE: the ifnum value (TODO_PLACEHOLDER below)
// MUST be replaced with a FREE physical interface on the target testbed (for
// example "1/2", or cluster notation "2/1/1"). Interface numbers vary per appliance
// and binding an in-use physical interface can disrupt management/data networking.
// Fill in a known-free interface before running these tests.
//
// Parent key attribute is "linksetid" (the linkset name, e.g. "LS/1"); it is
// renamed from the NITRO "id" key to avoid colliding with the synthetic Terraform
// resource "id". The member attribute "ifnum" is a scalar string. Composite
// resource ID format: id:<linkset>,ifnum:<interface> (legacy attr order
// [id, ifnum]). Reads go through the AGGREGATE parent endpoint
// (linkset_binding/<id>) which exposes the nested linkset_interface_binding[] rows.

import (
	"fmt"
	"testing"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// Participating-entity config (citrixadc_linkset) is lifted from
// citrixadc_framework/acctests/linkset_test.go. The legacy citrixadc_linkset
// resource exposes the linkset name as "linkset_id".

const testAccLinkset_interface_binding_basic_step1 = `
	resource "citrixadc_linkset" "tf_linkset" {
		linkset_id = "LS/1"
	}

	resource "citrixadc_linkset_interface_binding" "tf_linkset_interface_binding" {
		linksetid = citrixadc_linkset.tf_linkset.linkset_id
		# TODO_PLACEHOLDER: replace with a FREE physical interface on the testbed, e.g. "1/2"
		ifnum      = "1/2"
		depends_on = [citrixadc_linkset.tf_linkset]
	}

`

const testAccLinkset_interface_binding_basic_step2 = `
	# Keep the parent linkset without the binding to verify proper deletion
	resource "citrixadc_linkset" "tf_linkset" {
		linkset_id = "LS/1"
	}

`

func TestAccLinkset_interface_binding_basic(t *testing.T) {
	t.Skip("TODO: Requires review")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLinkset_interface_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLinkset_interface_binding_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLinkset_interface_bindingExist("citrixadc_linkset_interface_binding.tf_linkset_interface_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_linkset_interface_binding.tf_linkset_interface_binding", "linksetid", "LS/1"),
					// TODO_PLACEHOLDER: must match the free interface used in the config above.
					resource.TestCheckResourceAttr("citrixadc_linkset_interface_binding.tf_linkset_interface_binding", "ifnum", "1/2"),
				),
			},
			{
				Config: testAccLinkset_interface_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					// TODO_PLACEHOLDER: composite ID must match the free interface used above.
					testAccCheckLinkset_interface_bindingNotExist("citrixadc_linkset_interface_binding.tf_linkset_interface_binding", "id:LS%2F1,ifnum:1%2F2"),
				),
			},
		},
	})
}

// linksetInterfaceBindingAggregateReadForTest reads the bound interfaces via the
// AGGREGATE parent endpoint (linkset_binding/<id>) and flattens the nested
// "linkset_interface_binding" arrays. The direct linkset_interface_binding endpoint
// returns a keyless empty body on this firmware, so the check helpers must use the
// aggregate parent endpoint to locate the binding (mirrors the resource's
// linkset_interface_bindingAggregateRead).
func linksetInterfaceBindingAggregateReadForTest(client *service.NitroClient, linksetid string) ([]map[string]interface{}, error) {
	findParams := service.FindParams{
		ResourceType: "linkset_binding",
		// linksetid contains a '/' (e.g. "LS/1"); a path-embedded slash never reaches
		// NITRO, so read with the query-arg form. Mirrors the resource read.
		ArgsMap:                  map[string]string{"id": utils.UrlEncode(linksetid)},
		ResourceMissingErrorCode: 258,
	}
	parentArr, err := client.FindResourceArrayWithParams(findParams)
	if err != nil {
		return nil, err
	}

	rows := make([]map[string]interface{}, 0)
	for _, parent := range parentArr {
		nested, ok := parent["linkset_interface_binding"]
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

func testAccCheckLinkset_interface_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No linkset_interface_binding id is set")
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

		// Composite ID: id:UrlEncode(linkset),ifnum:UrlEncode(interface)
		// Legacy attr order is [id, ifnum] (the NITRO parent key is "id").
		idMap, _, err := utils.ParseIdString(rs.Primary.ID, []string{"id", "ifnum"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", rs.Primary.ID, err)
		}

		linksetid := idMap["id"]
		ifnum := idMap["ifnum"]

		dataArr, err := linksetInterfaceBindingAggregateReadForTest(client, linksetid)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the matching ifnum
		found := false
		for _, v := range dataArr {
			if val, ok := v["ifnum"].(string); ok && val == ifnum {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("linkset_interface_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckLinkset_interface_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		idMap, _, err := utils.ParseIdString(id, []string{"id", "ifnum"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", id, err)
		}

		linksetid := idMap["id"]
		ifnum := idMap["ifnum"]

		dataArr, err := linksetInterfaceBindingAggregateReadForTest(client, linksetid)

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
			return fmt.Errorf("linkset_interface_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckLinkset_interface_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_linkset_interface_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		idMap, _, err := utils.ParseIdString(rs.Primary.ID, []string{"id", "ifnum"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", rs.Primary.ID, err)
		}

		linksetid := idMap["id"]
		ifnum := idMap["ifnum"]

		dataArr, err := linksetInterfaceBindingAggregateReadForTest(client, linksetid)

		// If the parent linkset itself is gone, the binding is certainly gone.
		if err != nil {
			continue
		}

		for _, v := range dataArr {
			if val, ok := v["ifnum"].(string); ok && val == ifnum {
				return fmt.Errorf("linkset_interface_binding %s still exists", rs.Primary.ID)
			}
		}
	}

	return nil
}

const testAccLinkset_interface_binding_DataSource_basic = `
	resource "citrixadc_linkset" "tf_linkset" {
		linkset_id = "LS/1"
	}

	resource "citrixadc_linkset_interface_binding" "tf_linkset_interface_binding" {
		linksetid = citrixadc_linkset.tf_linkset.linkset_id
		# TODO_PLACEHOLDER: replace with a FREE physical interface on the testbed, e.g. "1/2"
		ifnum      = "1/2"
		depends_on = [citrixadc_linkset.tf_linkset]
	}

	data "citrixadc_linkset_interface_binding" "tf_linkset_interface_binding" {
		linksetid = citrixadc_linkset.tf_linkset.linkset_id
		# TODO_PLACEHOLDER: must match the free interface used above.
		ifnum      = "1/2"
		depends_on = [citrixadc_linkset_interface_binding.tf_linkset_interface_binding]
	}

`

func TestAccLinkset_interface_binding_import(t *testing.T) {
	t.Skip("TODO: Requires review")
	const resAddr = "citrixadc_linkset_interface_binding.tf_linkset_interface_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLinkset_interface_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLinkset_interface_binding_basic_step1,
			},
			{
				Config:                  testAccLinkset_interface_binding_basic_step1,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccLinkset_interface_binding_DataSource_basic(t *testing.T) {
	t.Skip("TODO: Requires review")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLinkset_interface_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLinkset_interface_binding_DataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_linkset_interface_binding.tf_linkset_interface_binding", "linksetid", "LS/1"),
					// TODO_PLACEHOLDER: must match the free interface used above.
					resource.TestCheckResourceAttr("data.citrixadc_linkset_interface_binding.tf_linkset_interface_binding", "ifnum", "1/2"),
				),
			},
		},
	})
}
