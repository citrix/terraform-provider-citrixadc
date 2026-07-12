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

const testAccNetprofile_srcportset_binding_basic = `
	resource "citrixadc_netprofile" "tf_netprofile" {
		name                   = "tf_netprofile"
		proxyprotocol          = "ENABLED"
		proxyprotocoltxversion = "V1"
	}
	resource "citrixadc_netprofile_srcportset_binding" "tf_binding" {
		name         = citrixadc_netprofile.tf_netprofile.name
		srcportrange = "2000"
	}
`

const testAccNetprofile_srcportset_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion
	resource "citrixadc_netprofile" "tf_netprofile" {
		name                   = "tf_netprofile"
		proxyprotocol          = "ENABLED"
		proxyprotocoltxversion = "V1"
	}
`

func TestAccNetprofile_srcportset_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNetprofile_srcportset_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNetprofile_srcportset_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetprofile_srcportset_bindingExist("citrixadc_netprofile_srcportset_binding.tf_binding", nil),
				),
			},
			{
				Config: testAccNetprofile_srcportset_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetprofile_srcportset_bindingNotExist("citrixadc_netprofile_srcportset_binding.tf_binding", "tf_netprofile,2000"),
				),
			},
		},
	})
}

func testAccCheckNetprofile_srcportset_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No netprofile_srcportset_binding id is set")
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

		idMap, _, err := utils.ParseIdString(bindingId, []string{"name", "srcportrange"}, nil)
		if err != nil {
			return err
		}
		name := idMap["name"]
		srcportrange := idMap["srcportrange"]

		findParams := service.FindParams{
			ResourceType:             "netprofile_srcportset_binding",
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
			if v["srcportrange"].(string) == srcportrange {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("netprofile_srcportset_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckNetprofile_srcportset_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}
		idMap, _, err := utils.ParseIdString(id, []string{"name", "srcportrange"}, nil)
		if err != nil {
			return err
		}
		name := idMap["name"]
		srcportrange := idMap["srcportrange"]

		findParams := service.FindParams{
			ResourceType:             "netprofile_srcportset_binding",
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
			if v["srcportrange"].(string) == srcportrange {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("netprofile_srcportset_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckNetprofile_srcportset_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_netprofile_srcportset_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Netprofile_srcportset_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("netprofile_srcportset_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccNetprofile_srcportset_bindingDataSource_basic = `
	resource "citrixadc_netprofile" "tf_netprofile" {
		name                   = "tf_netprofile"
		proxyprotocol          = "ENABLED"
		proxyprotocoltxversion = "V1"
	}
	resource "citrixadc_netprofile_srcportset_binding" "tf_binding" {
		name         = citrixadc_netprofile.tf_netprofile.name
		srcportrange = "2000"
	}

	data "citrixadc_netprofile_srcportset_binding" "tf_binding" {
		name         = citrixadc_netprofile_srcportset_binding.tf_binding.name
		srcportrange = citrixadc_netprofile_srcportset_binding.tf_binding.srcportrange
		depends_on   = [citrixadc_netprofile_srcportset_binding.tf_binding]
	}
`

func TestAccNetprofile_srcportset_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNetprofile_srcportset_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_netprofile_srcportset_binding.tf_binding", "name", "tf_netprofile"),
					resource.TestCheckResourceAttr("data.citrixadc_netprofile_srcportset_binding.tf_binding", "srcportrange", "2000"),
				),
			},
		},
	})
}

// testAccNetprofile_srcportset_binding_upgrade_basic reuses the _basic config
// values and is valid under BOTH the last SDK v2 release (2.2.0) schema and the
// current Framework schema (uses only SDK v2 attribute names).
const testAccNetprofile_srcportset_binding_upgrade_basic = `
	resource "citrixadc_netprofile" "tf_netprofile" {
		name                   = "tf_netprofile"
		proxyprotocol          = "ENABLED"
		proxyprotocoltxversion = "V1"
	}
	resource "citrixadc_netprofile_srcportset_binding" "tf_binding" {
		name         = citrixadc_netprofile.tf_netprofile.name
		srcportrange = "2000"
	}
`

// TestAccNetprofile_srcportset_binding_sdkv2StateUpgrade verifies that state
// written by the last SDK v2 release (legacy comma-joined id) is transparently
// upgraded by the current Framework provider. Step 1 creates the binding with
// citrix/citrixadc 2.2.0 (legacy id "name,srcportrange"); step 2 refreshes/plans
// the same config through the current Framework provider, whose Read parses the
// legacy id and recomputes it to the new "name:<v>,srcportrange:<v>" format
// (SetAttrFromGet).
func TestAccNetprofile_srcportset_binding_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckNetprofile_srcportset_bindingDestroy,
		Steps: []resource.TestStep{
			{
				// Step 1: create with the last SDK v2 release, writing the legacy id.
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.2.0",
					},
				},
				Config: testAccNetprofile_srcportset_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetprofile_srcportset_bindingExist("citrixadc_netprofile_srcportset_binding.tf_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_netprofile_srcportset_binding.tf_binding", "id", "tf_netprofile,2000"),
				),
			},
			{
				// Step 2: refresh/apply the same config through the current Framework
				// provider. Read exercises ParseIdString on the legacy id, then
				// recomputes the id to the new key:value canonical format.
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccNetprofile_srcportset_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetprofile_srcportset_bindingExist("citrixadc_netprofile_srcportset_binding.tf_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_netprofile_srcportset_binding.tf_binding", "id", "name:tf_netprofile,srcportrange:2000"),
				),
			},
		},
	})
}

func TestAccNetprofile_srcportset_binding_import(t *testing.T) {
	const resAddr = "citrixadc_netprofile_srcportset_binding.tf_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNetprofile_srcportset_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccNetprofile_srcportset_binding_basic},
			{Config: testAccNetprofile_srcportset_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
		},
	})
}
