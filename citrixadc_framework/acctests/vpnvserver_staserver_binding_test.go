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

const testAccVpnvserver_staserver_binding_basic = `
	resource "citrixadc_vpnvserver" "tf_vpnvserver" {
		name        = "tf_vserver"
		servicetype = "SSL"
		ipv46       = "3.3.3.3"
		port        = 443
	}
	resource "citrixadc_vpnvserver_staserver_binding" "tf_binding" {
		name           = citrixadc_vpnvserver.tf_vpnvserver.name
		staserver      = "http://www.example.com/"
		staaddresstype = "IPV6"
	}
`

const testAccVpnvserver_staserver_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion
	resource "citrixadc_vpnvserver" "tf_vpnvserver" {
		name        = "tf_vserver"
		servicetype = "SSL"
		ipv46       = "3.3.3.3"
		port        = 443
	}
`

func TestAccVpnvserver_staserver_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVpnvserver_staserver_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnvserver_staserver_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnvserver_staserver_bindingExist("citrixadc_vpnvserver_staserver_binding.tf_binding", nil),
				),
			},
			{
				Config: testAccVpnvserver_staserver_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnvserver_staserver_bindingNotExist("citrixadc_vpnvserver_staserver_binding.tf_binding", "tf_vserver,http://www.example.com/"),
				),
			},
		},
	})
}

const testAccVpnvserver_staserver_binding_upgrade_basic = `
	resource "citrixadc_vpnvserver" "tf_vpnvserver" {
		name        = "tf_vserver"
		servicetype = "SSL"
		ipv46       = "3.3.3.3"
		port        = 443
	}
	resource "citrixadc_vpnvserver_staserver_binding" "tf_binding" {
		name           = citrixadc_vpnvserver.tf_vpnvserver.name
		staserver      = "http://www.example.com/"
		staaddresstype = "IPV6"
	}
`

func TestAccVpnvserver_staserver_binding_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckVpnvserver_staserver_bindingDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.2.0",
					},
				},
				Config: testAccVpnvserver_staserver_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnvserver_staserver_bindingExist("citrixadc_vpnvserver_staserver_binding.tf_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnvserver_staserver_binding.tf_binding", "id", "tf_vserver,http://www.example.com/"),
				),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccVpnvserver_staserver_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnvserver_staserver_bindingExist("citrixadc_vpnvserver_staserver_binding.tf_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnvserver_staserver_binding.tf_binding", "id", "name:tf_vserver,staserver:http%3A%2F%2Fwww.example.com%2F"),
				),
			},
		},
	})
}

func testAccCheckVpnvserver_staserver_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No vpnvserver_staserver_binding id is set")
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

		idMap, _, err := utils.ParseIdString(bindingId, []string{"name", "staserver"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", bindingId, err)
		}
		name := idMap["name"]
		staserver := idMap["staserver"]

		findParams := service.FindParams{
			ResourceType:             "vpnvserver_staserver_binding",
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
			if v["staserver"].(string) == staserver {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("vpnvserver_staserver_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckVpnvserver_staserver_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		idMap, _, err := utils.ParseIdString(id, []string{"name", "staserver"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", id, err)
		}
		name := idMap["name"]
		staserver := idMap["staserver"]

		findParams := service.FindParams{
			ResourceType:             "vpnvserver_staserver_binding",
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
			if v["staserver"].(string) == staserver {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("vpnvserver_staserver_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckVpnvserver_staserver_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_vpnvserver_staserver_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Vpnvserver_staserver_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("vpnvserver_staserver_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccVpnvserver_staserver_bindingDataSource_basic = `
	resource "citrixadc_vpnvserver" "tf_vpnvserver" {
		name        = "tf_vserver"
		servicetype = "SSL"
		ipv46       = "3.3.3.3"
		port        = 443
	}
	resource "citrixadc_vpnvserver_staserver_binding" "tf_binding" {
		name           = citrixadc_vpnvserver.tf_vpnvserver.name
		staserver      = "http://www.example.com/"
		staaddresstype = "IPV6"
	}
	
	data "citrixadc_vpnvserver_staserver_binding" "tf_binding" {
		name      = citrixadc_vpnvserver_staserver_binding.tf_binding.name
		staserver = citrixadc_vpnvserver_staserver_binding.tf_binding.staserver
	}
`

func TestAccVpnvserver_staserver_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnvserver_staserver_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_vpnvserver_staserver_binding.tf_binding", "name", "tf_vserver"),
					resource.TestCheckResourceAttr("data.citrixadc_vpnvserver_staserver_binding.tf_binding", "staserver", "http://www.example.com/"),
					resource.TestCheckResourceAttr("data.citrixadc_vpnvserver_staserver_binding.tf_binding", "staaddresstype", "IPV6"),
					resource.TestCheckResourceAttrSet("data.citrixadc_vpnvserver_staserver_binding.tf_binding", "id"),
				),
			},
		},
	})
}

func TestAccVpnvserver_staserver_binding_import(t *testing.T) {
	const resAddr = "citrixadc_vpnvserver_staserver_binding.tf_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVpnvserver_staserver_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccVpnvserver_staserver_binding_basic},
			{Config: testAccVpnvserver_staserver_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
		},
	})
}
