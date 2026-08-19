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

// channel_interface_binding binds a physical interface to an LA channel.
//
// IMPORTANT - TESTBED-SPECIFIC INTERFACE: the ifnum value (TODO_PLACEHOLDER below)
// MUST be replaced with a FREE physical interface on the target testbed (for
// example "1/2"). Interface numbers vary per appliance and binding an in-use
// physical interface can disrupt management/data networking. Fill in a known-free
// interface before running these tests.
//
// Parent key attribute is "channelid" (the LA channel name, e.g. "LA/1"); it is
// renamed from the NITRO "id" key to avoid colliding with the synthetic Terraform
// resource "id". The member attribute "ifnum" is a LIST of interface strings.
// Composite resource ID format: id:<channel>,ifnum:<interface> (legacy attr order
// [id, ifnum]). Reads go through the AGGREGATE parent endpoint
// (channel_binding/<id>) which exposes the nested channel_interface_binding[] rows.

import (
	"fmt"
	"net/url"
	"strings"
	"testing"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

// Participating-entity config (citrixadc_channel) is lifted from
// citrixadc_framework/acctests/channel_test.go (testAccChannel_basic). The legacy
// citrixadc_channel resource exposes the channel name as "channel_id".

const testAccChannel_interface_binding_basic_step1 = `
	resource "citrixadc_channel" "tf_channel" {
		channel_id = "LA/1"
		speed      = "1000"
	}

	resource "citrixadc_channel_interface_binding" "tf_channel_interface_binding" {
		channelid = citrixadc_channel.tf_channel.channel_id
		# TODO_PLACEHOLDER: replace with a FREE physical interface on the testbed, e.g. ["1/2"]
		ifnum      = ["1/2"]
		depends_on = [citrixadc_channel.tf_channel]
	}

`

const testAccChannel_interface_binding_basic_step2 = `
	# Keep the parent channel without the binding to verify proper deletion
	resource "citrixadc_channel" "tf_channel" {
		channel_id = "LA/1"
		speed      = "1000"
	}

`

func TestAccChannel_interface_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckChannel_interface_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccChannel_interface_binding_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckChannel_interface_bindingExist("citrixadc_channel_interface_binding.tf_channel_interface_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_channel_interface_binding.tf_channel_interface_binding", "channelid", "LA/1"),
					// TODO_PLACEHOLDER: must match the free interface used in the config above.
					resource.TestCheckResourceAttr("citrixadc_channel_interface_binding.tf_channel_interface_binding", "ifnum.0", "1/2"),
				),
			},
			{
				Config: testAccChannel_interface_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					// TODO_PLACEHOLDER: composite ID must match the free interface used above.
					testAccCheckChannel_interface_bindingNotExist("citrixadc_channel_interface_binding.tf_channel_interface_binding", "id:LA%2F1,ifnum:1%2F2"),
				),
			},
		},
	})
}

func TestAccChannel_interface_binding_import(t *testing.T) {
	const resAddr = "citrixadc_channel_interface_binding.tf_channel_interface_binding"

	// Backward-compat: import via the LEGACY SDK v2 id. Rebuild the legacy positional id from
	// the current canonical key:value id (raw values, only the keys actually set, in legacy
	// order: id,ifnum) so it matches exactly what SDK v2 wrote.
	legacyID := func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resAddr]
		if !ok {
			return "", fmt.Errorf("resource not found in state: %s", resAddr)
		}
		kv := map[string]string{}
		for _, p := range strings.Split(rs.Primary.ID, ",") {
			if i := strings.Index(p, ":"); i >= 0 {
				v, _ := url.QueryUnescape(p[i+1:])
				kv[p[:i]] = v
			}
		}
		ordr := []string{"id", "ifnum"}
		parts := make([]string, 0, len(ordr))
		for _, k := range ordr {
			if v, ok := kv[k]; ok {
				parts = append(parts, v)
			}
		}
		// Fallback: a positional (non key:value) id has no key:value parts to reorder; import it as-is.
		if len(parts) == 0 {
			return rs.Primary.ID, nil
		}
		return strings.Join(parts, ","), nil
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckChannel_interface_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccChannel_interface_binding_basic_step1,
			},
			{
				Config:            testAccChannel_interface_binding_basic_step1,
				ResourceName:      resAddr,
				ImportState:       true,
				ImportStateVerify: true,
				// channelid and ifnum are both recovered from the composite ID by the
				// Read path (readChannelInterfaceBindingFromApi backfills identity attrs
				// from the parsed idMap), so `terraform import` fully round-trips and
				// nothing needs to be ignored.
				ImportStateVerifyIgnore: []string{},
			},
			{Config: testAccChannel_interface_binding_basic_step1, ResourceName: resAddr, ImportState: true, ImportStateIdFunc: legacyID, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
		},
	})
}

// channelInterfaceBindingAggregateReadForTest reads the bound interfaces via the
// AGGREGATE parent endpoint (channel_binding/<id>) and flattens the nested
// "channel_interface_binding" arrays. The direct channel_interface_binding endpoint
// returns a keyless empty body on this firmware, so the check helpers must use the
// aggregate parent endpoint to locate the binding (mirrors the resource's
// channel_interface_bindingAggregateRead).
func channelInterfaceBindingAggregateReadForTest(client *service.NitroClient, channelid string) ([]map[string]interface{}, error) {
	findParams := service.FindParams{
		ResourceType: "channel_binding",
		// channelid contains a '/' (e.g. "LA/1"). A path-embedded slash never reaches
		// NITRO (front-end web server returns an HTML 404), so read with the query-arg
		// form (channel_binding?args=id:LA%2F1). Mirrors the resource's
		// channel_interface_bindingAggregateRead.
		ArgsMap:                  map[string]string{"id": utils.UrlEncode(channelid)},
		ResourceMissingErrorCode: 258,
	}
	parentArr, err := client.FindResourceArrayWithParams(findParams)
	if err != nil {
		return nil, err
	}

	rows := make([]map[string]interface{}, 0)
	for _, parent := range parentArr {
		nested, ok := parent["channel_interface_binding"]
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

// channelInterfaceRowHasIfnumForTest reports whether an aggregate
// channel_interface_binding row's "ifnum" contains want. The NITRO aggregate
// response represents ifnum as a JSON array (e.g. ["1/2"]); the scalar form is
// tolerated defensively. Mirrors the resource's channelInterfaceRowHasIfnum.
func channelInterfaceRowHasIfnumForTest(raw interface{}, want string) bool {
	switch v := raw.(type) {
	case string:
		return v == want
	case []interface{}:
		for _, item := range v {
			if s, ok := item.(string); ok && s == want {
				return true
			}
		}
	}
	return false
}

func testAccCheckChannel_interface_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No channel_interface_binding id is set")
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

		// Composite ID: id:UrlEncode(channel),ifnum:UrlEncode(interface)
		// Legacy attr order is [id, ifnum] (the NITRO parent key is "id").
		idMap, _, err := utils.ParseIdString(rs.Primary.ID, []string{"id", "ifnum"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", rs.Primary.ID, err)
		}

		channelid := idMap["id"]
		ifnum := idMap["ifnum"]

		dataArr, err := channelInterfaceBindingAggregateReadForTest(client, channelid)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the matching ifnum
		found := false
		for _, v := range dataArr {
			if channelInterfaceRowHasIfnumForTest(v["ifnum"], ifnum) {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("channel_interface_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckChannel_interface_bindingNotExist(n string, id string) resource.TestCheckFunc {
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

		channelid := idMap["id"]
		ifnum := idMap["ifnum"]

		dataArr, err := channelInterfaceBindingAggregateReadForTest(client, channelid)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to hopefully not find the one with the matching ifnum
		found := false
		for _, v := range dataArr {
			if channelInterfaceRowHasIfnumForTest(v["ifnum"], ifnum) {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("channel_interface_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckChannel_interface_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_channel_interface_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		idMap, _, err := utils.ParseIdString(rs.Primary.ID, []string{"id", "ifnum"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", rs.Primary.ID, err)
		}

		channelid := idMap["id"]
		ifnum := idMap["ifnum"]

		dataArr, err := channelInterfaceBindingAggregateReadForTest(client, channelid)

		// If the parent channel itself is gone, the binding is certainly gone.
		if err != nil {
			continue
		}

		for _, v := range dataArr {
			if channelInterfaceRowHasIfnumForTest(v["ifnum"], ifnum) {
				return fmt.Errorf("channel_interface_binding %s still exists", rs.Primary.ID)
			}
		}
	}

	return nil
}

const testAccChannel_interface_binding_DataSource_basic = `
	resource "citrixadc_channel" "tf_channel" {
		channel_id = "LA/1"
		speed      = "1000"
	}

	resource "citrixadc_channel_interface_binding" "tf_channel_interface_binding" {
		channelid = citrixadc_channel.tf_channel.channel_id
		# TODO_PLACEHOLDER: replace with a FREE physical interface on the testbed, e.g. ["1/2"]
		ifnum      = ["1/2"]
		depends_on = [citrixadc_channel.tf_channel]
	}

	data "citrixadc_channel_interface_binding" "tf_channel_interface_binding" {
		channelid = citrixadc_channel.tf_channel.channel_id
		# TODO_PLACEHOLDER: must match the free interface used above.
		ifnum      = ["1/2"]
		depends_on = [citrixadc_channel_interface_binding.tf_channel_interface_binding]
	}

`

func TestAccChannel_interface_binding_DataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckChannel_interface_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccChannel_interface_binding_DataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_channel_interface_binding.tf_channel_interface_binding", "channelid", "LA/1"),
					// TODO_PLACEHOLDER: must match the free interface used above.
					resource.TestCheckResourceAttr("data.citrixadc_channel_interface_binding.tf_channel_interface_binding", "ifnum.0", "1/2"),
				),
			},
		},
	})
}

// TestAccChannel_interface_binding_selfHealing verifies drift recovery:
// after the binding is deleted out-of-band, the next apply of the same config recreates it.
// Reuses the _basic (step1) config; the out-of-band delete mirrors the resource Delete
// (raw channel id "LA/1" as the URL name, ifnum passed URL-encoded as the only arg).
func TestAccChannel_interface_binding_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_channel_interface_binding.tf_channel_interface_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckChannel_interface_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccChannel_interface_binding_basic_step1,
				Check:  resource.ComposeTestCheckFunc(testAccCheckChannel_interface_bindingExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResourceWithArgs(service.Channel_interface_binding.Type(), "LA/1", []string{"ifnum:" + utils.UrlEncode("1/2")}); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccChannel_interface_binding_basic_step1,
				Check:  resource.ComposeTestCheckFunc(testAccCheckChannel_interface_bindingExist(resAddr, nil)),
			},
		},
	})
}
