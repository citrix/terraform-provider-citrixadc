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

const testAccNetprofile_natrule_binding_basic = `
	resource "citrixadc_netprofile" "tf_netprofile" {
		name                   = "tf_netprofile"
		proxyprotocol          = "ENABLED"
		proxyprotocoltxversion = "V1"
	}
	resource "citrixadc_netprofile_natrule_binding" "tf_binding" {
		name      = citrixadc_netprofile.tf_netprofile.name
		natrule   = "10.10.10.10"
		netmask   = "255.255.255.255"
		rewriteip = "3.3.3.3"
	}
`

const testAccNetprofile_natrule_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion
	resource "citrixadc_netprofile" "tf_netprofile" {
		name                   = "tf_netprofile"
		proxyprotocol          = "ENABLED"
		proxyprotocoltxversion = "V1"
	}
`

func TestAccNetprofile_natrule_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNetprofile_natrule_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNetprofile_natrule_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetprofile_natrule_bindingExist("citrixadc_netprofile_natrule_binding.tf_binding", nil),
				),
			},
			{
				Config: testAccNetprofile_natrule_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetprofile_natrule_bindingNotExist("citrixadc_netprofile_natrule_binding.tf_binding", "tf_netprofile,10.10.10.10"),
				),
			},
		},
	})
}

func testAccCheckNetprofile_natrule_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No netprofile_natrule_binding id is set")
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

		idMap, _, err := utils.ParseIdString(bindingId, []string{"name", "natrule"}, nil)
		if err != nil {
			return err
		}
		name := idMap["name"]
		natrule := idMap["natrule"]

		findParams := service.FindParams{
			ResourceType:             "netprofile_natrule_binding",
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
			if v["natrule"].(string) == natrule {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("netprofile_natrule_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckNetprofile_natrule_bindingNotExist(n string, id string) resource.TestCheckFunc {
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
		natrule := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "netprofile_natrule_binding",
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
			if v["natrule"].(string) == natrule {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("netprofile_natrule_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckNetprofile_natrule_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_netprofile_natrule_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Netprofile_natrule_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("netprofile_natrule_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccNetprofile_natrule_bindingDataSource_basic = `
	resource "citrixadc_netprofile" "tf_netprofile" {
		name                   = "tf_netprofile"
		proxyprotocol          = "ENABLED"
		proxyprotocoltxversion = "V1"
	}
	resource "citrixadc_netprofile_natrule_binding" "tf_binding" {
		name      = citrixadc_netprofile.tf_netprofile.name
		natrule   = "10.10.10.10"
		netmask   = "255.255.255.255"
		rewriteip = "3.3.3.3"
	}

	data "citrixadc_netprofile_natrule_binding" "tf_binding" {
		name    = citrixadc_netprofile_natrule_binding.tf_binding.name
		natrule = citrixadc_netprofile_natrule_binding.tf_binding.natrule
	}
`

func TestAccNetprofile_natrule_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNetprofile_natrule_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_netprofile_natrule_binding.tf_binding", "name", "tf_netprofile"),
					resource.TestCheckResourceAttr("data.citrixadc_netprofile_natrule_binding.tf_binding", "natrule", "10.10.10.10"),
					resource.TestCheckResourceAttr("data.citrixadc_netprofile_natrule_binding.tf_binding", "netmask", "255.255.255.255"),
					resource.TestCheckResourceAttr("data.citrixadc_netprofile_natrule_binding.tf_binding", "rewriteip", "3.3.3.3"),
				),
			},
		},
	})
}

const testAccNetprofile_natrule_binding_upgrade_basic = `
	resource "citrixadc_netprofile" "tf_netprofile" {
		name                   = "tf_netprofile"
		proxyprotocol          = "ENABLED"
		proxyprotocoltxversion = "V1"
	}
	resource "citrixadc_netprofile_natrule_binding" "tf_binding" {
		name      = citrixadc_netprofile.tf_netprofile.name
		natrule   = "10.10.10.10"
		netmask   = "255.255.255.255"
		rewriteip = "3.3.3.3"
	}
`

func TestAccNetprofile_natrule_binding_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckNetprofile_natrule_bindingDestroy,
		Steps: []resource.TestStep{
			// Step 1: Create the resource with the last SDK v2 release (writes state with the legacy comma ID).
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.2.0",
					},
				},
				Config: testAccNetprofile_natrule_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetprofile_natrule_bindingExist("citrixadc_netprofile_natrule_binding.tf_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_netprofile_natrule_binding.tf_binding", "id", "tf_netprofile,10.10.10.10"),
				),
			},
			// Step 2: Refresh the legacy-id state through the current (framework) provider.
			// Read exercises ParseIdString on the legacy id and recomputes the canonical new-format id.
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccNetprofile_natrule_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetprofile_natrule_bindingExist("citrixadc_netprofile_natrule_binding.tf_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_netprofile_natrule_binding.tf_binding", "id", "name:tf_netprofile,natrule:10.10.10.10"),
				),
			},
		},
	})
}

func TestAccNetprofile_natrule_binding_import(t *testing.T) {
	const resAddr = "citrixadc_netprofile_natrule_binding.tf_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNetprofile_natrule_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccNetprofile_natrule_binding_basic},
			{Config: testAccNetprofile_natrule_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
		},
	})
}
