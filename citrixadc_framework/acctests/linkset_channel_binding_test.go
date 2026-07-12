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
	"net/url"
	"strings"
	"testing"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testAccLinkset_channel_binding_basic = `

resource "citrixadc_linkset_channel_binding" "tf_linkset_channel_binding" {
	linkset_id = citrixadc_linkset.tf_linkset.linkset_id
	ifnum      = citrixadc_channel.tf_channel.channel_id
	}
  
  
  resource "citrixadc_linkset" "tf_linkset"{
	  linkset_id = "LS/3"
	}
  
  resource "citrixadc_channel" "tf_channel"{
	  channel_id = "LA/3"
	}
  
`

const testAccLinkset_channel_binding_basic_step2 = `
resource "citrixadc_linkset" "tf_linkset"{
	linkset_id = "LS/3"
}

resource "citrixadc_channel" "tf_channel"{
	channel_id = "LA/3"
}
`

func TestAccLinkset_channel_binding_basic(t *testing.T) {
	t.Skip("TODO: Need to find a way to test this resource!")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLinkset_channel_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLinkset_channel_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLinkset_channel_bindingExist("citrixadc_linkset_channel_binding.tf_linkset_channel_binding", nil),
				),
			},
			{
				Config: testAccLinkset_channel_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLinkset_channel_bindingNotExist("citrixadc_linkset_channel_binding.tf_linkset_channel_binding", "LS/3,LA/3"),
				),
			},
		},
	})
}

func testAccCheckLinkset_channel_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No linkset_channel_binding id is set")
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

		idSlice := strings.SplitN(bindingId, ",", 2)

		id := idSlice[0]
		ifnum := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "linkset_channel_binding",
			ResourceName:             url.QueryEscape(url.QueryEscape(id)),
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the matching ifnum
		found := false
		for _, v := range dataArr {
			if v["ifnum"].(string) == ifnum {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("linkset_channel_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckLinkset_channel_bindingNotExist(n string, id string) resource.TestCheckFunc {
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

		id := idSlice[0]
		ifnum := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "linkset_channel_binding",
			ResourceName:             url.QueryEscape(url.QueryEscape(id)),
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to hopefully not find the one with the matching ifnum
		found := false
		for _, v := range dataArr {
			if v["ifnum"].(string) == ifnum {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("linkset_channel_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckLinkset_channel_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_linkset_channel_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Linkset_channel_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("linkset_channel_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccLinkset_channel_bindingDataSource_basic = `

resource "citrixadc_linkset_channel_binding" "tf_linkset_channel_binding" {
	linkset_id = citrixadc_linkset.tf_linkset.linkset_id
	ifnum      = citrixadc_channel.tf_channel.channel_id
	}
  
  
  resource "citrixadc_linkset" "tf_linkset"{
	  linkset_id = "LS/3"
	}
  
  resource "citrixadc_channel" "tf_channel"{
	  channel_id = "LA/3"
	}

	data "citrixadc_linkset_channel_binding" "tf_linkset_channel_binding" {
		linkset_id = "LS/3"
		ifnum      = "LA/3"
		depends_on = [citrixadc_linkset_channel_binding.tf_linkset_channel_binding]
	}
`

func TestAccLinkset_channel_bindingDataSource_basic(t *testing.T) {
	t.Skip("TODO: Need to find a way to test this resource!")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccLinkset_channel_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_linkset_channel_binding.tf_linkset_channel_binding", "linkset_id", "LS/3"),
					resource.TestCheckResourceAttr("data.citrixadc_linkset_channel_binding.tf_linkset_channel_binding", "ifnum", "LA/3"),
				),
			},
		},
	})
}

func TestAccLinkset_channel_binding_import(t *testing.T) {
	t.Skip("TODO: Need to find a way to test this resource!")
	const resAddr = "citrixadc_linkset_channel_binding.tf_linkset_channel_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLinkset_channel_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccLinkset_channel_binding_basic},
			{Config: testAccLinkset_channel_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
		},
	})
}

// testAccLinkset_channel_binding_upgrade_basic mirrors the _basic config (a
// linkset + a channel bound together). It uses only the SDK v2 attribute names
// (linkset_id, ifnum, channel_id) that the migration restored, so it is valid
// under BOTH the SDK v2 2.2.0 schema and the current framework schema. This lets
// it be applied with the old provider in step 1 and re-planned with the new
// provider in step 2 of the state-upgrade test below.
const testAccLinkset_channel_binding_upgrade_basic = `

resource "citrixadc_linkset" "tf_linkset" {
	linkset_id = "LS/3"
}

resource "citrixadc_channel" "tf_channel" {
	channel_id = "LA/3"
}

resource "citrixadc_linkset_channel_binding" "tf_linkset_channel_binding" {
	linkset_id = citrixadc_linkset.tf_linkset.linkset_id
	ifnum      = citrixadc_channel.tf_channel.channel_id
}
`

// TestAccLinkset_channel_binding_sdkv2StateUpgrade verifies that a resource
// created by the LAST SDK v2 release (2.2.0) — which writes the legacy
// comma-joined id "linkset_id,ifnum" (e.g. "LS/3,LA/3") — is refreshed and
// re-applied correctly by the CURRENT framework provider. Step 2 exercises
// ParseIdString on the legacy id during the framework Read.
//
// On this branch the framework recomputes data.Id into the new key:value form
// during Read (linkset_channel_bindingSetAttrFromGet in resource_schema.go), so
// after the step-2 refresh the id becomes the canonical
// "linkset_id:LS%2F3,ifnum:LA%2F3".
//
// Skipped for the same reason as the other tests of this resource (creating the
// underlying linkset/channel binding is not reliably testable on the shared ADC).
func TestAccLinkset_channel_binding_sdkv2StateUpgrade(t *testing.T) {
	t.Skip("TODO: Need to find a way to test this resource!")
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckLinkset_channel_bindingDestroy,
		Steps: []resource.TestStep{
			// Step 1: create with the last SDK v2 release from the registry. This
			// writes state carrying the LEGACY comma-joined id.
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.2.0",
					},
				},
				Config: testAccLinkset_channel_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLinkset_channel_bindingExist("citrixadc_linkset_channel_binding.tf_linkset_channel_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_linkset_channel_binding.tf_linkset_channel_binding", "id", "LS/3,LA/3"),
				),
			},
			// Step 2: same config through the CURRENT framework provider. Terraform
			// refreshes the legacy-id state through the framework Read (exercising
			// ParseIdString on the legacy id) then plans/applies. The framework Read
			// recomputes the id into the new key:value format.
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccLinkset_channel_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLinkset_channel_bindingExist("citrixadc_linkset_channel_binding.tf_linkset_channel_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_linkset_channel_binding.tf_linkset_channel_binding", "id", "linkset_id:LS%2F3,ifnum:LA%2F3"),
				),
			},
		},
	})
}
