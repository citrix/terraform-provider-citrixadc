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

// The vpnvserver participating-entity HCL is reused from acctests/vpnvserver_test.go
// (name + servicetype=SSL + ipv46/port). The secureprivateaccessurl is a literal URL
// value supplied directly (there is no separate entity resource for it).

const testAccVpnvserver_secureprivateaccessurl_binding_basic = `
	resource "citrixadc_vpnvserver" "tf_vpnvserver" {
		name        = "tf.citrix.example.com"
		servicetype = "SSL"
		ipv46       = "3.3.3.3"
		port        = 443
	}

	resource "citrixadc_vpnvserver_secureprivateaccessurl_binding" "tf_binding" {
		name                   = citrixadc_vpnvserver.tf_vpnvserver.name
		secureprivateaccessurl = "https://app.example.com/"
		depends_on             = [citrixadc_vpnvserver.tf_vpnvserver]
	}
`

const testAccVpnvserver_secureprivateaccessurl_binding_basic_step2 = `
	# Keep the participating entity without the actual binding to confirm deletion
	resource "citrixadc_vpnvserver" "tf_vpnvserver" {
		name        = "tf.citrix.example.com"
		servicetype = "SSL"
		ipv46       = "3.3.3.3"
		port        = 443
	}
`

func TestAccVpnvserver_secureprivateaccessurl_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVpnvserver_secureprivateaccessurl_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnvserver_secureprivateaccessurl_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnvserver_secureprivateaccessurl_bindingExist("citrixadc_vpnvserver_secureprivateaccessurl_binding.tf_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnvserver_secureprivateaccessurl_binding.tf_binding", "name", "tf.citrix.example.com"),
					resource.TestCheckResourceAttr("citrixadc_vpnvserver_secureprivateaccessurl_binding.tf_binding", "secureprivateaccessurl", "https://app.example.com/"),
				),
			},
			{
				// Binding dropped from config; confirm it was deleted from the ADC.
				Config: testAccVpnvserver_secureprivateaccessurl_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnvserver_secureprivateaccessurl_bindingNotExist("citrixadc_vpnvserver_secureprivateaccessurl_binding.tf_binding", "tf.citrix.example.com", "https://app.example.com/"),
				),
			},
		},
	})
}

func testAccCheckVpnvserver_secureprivateaccessurl_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No vpnvserver_secureprivateaccessurl_binding id is set")
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

		// ID is key:UrlEncode(value) pairs - parse with ParseIdString
		idMap, _, err := utils.ParseIdString(rs.Primary.ID, nil, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", rs.Primary.ID, err)
		}

		name := idMap["name"]
		secureprivateaccessurl := idMap["secureprivateaccessurl"]

		findParams := service.FindParams{
			ResourceType:             service.Vpnvserver_secureprivateaccessurl_binding.Type(),
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the matching secureprivateaccessurl
		found := false
		for _, v := range dataArr {
			if val, ok := v["secureprivateaccessurl"].(string); ok && val == secureprivateaccessurl {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("vpnvserver_secureprivateaccessurl_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckVpnvserver_secureprivateaccessurl_bindingNotExist(n string, name string, secureprivateaccessurl string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		findParams := service.FindParams{
			ResourceType:             service.Vpnvserver_secureprivateaccessurl_binding.Type(),
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			return err
		}

		// Iterate through results to hopefully not find the matching secureprivateaccessurl
		found := false
		for _, v := range dataArr {
			if val, ok := v["secureprivateaccessurl"].(string); ok && val == secureprivateaccessurl {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("vpnvserver_secureprivateaccessurl_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckVpnvserver_secureprivateaccessurl_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_vpnvserver_secureprivateaccessurl_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		idMap, _, err := utils.ParseIdString(rs.Primary.ID, nil, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", rs.Primary.ID, err)
		}

		name := idMap["name"]
		secureprivateaccessurl := idMap["secureprivateaccessurl"]

		findParams := service.FindParams{
			ResourceType:             service.Vpnvserver_secureprivateaccessurl_binding.Type(),
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)
		// If the parent vserver itself is gone, the binding cannot exist
		if err != nil {
			continue
		}

		for _, v := range dataArr {
			if val, ok := v["secureprivateaccessurl"].(string); ok && val == secureprivateaccessurl {
				return fmt.Errorf("vpnvserver_secureprivateaccessurl_binding %s still exists", rs.Primary.ID)
			}
		}
	}

	return nil
}

const testAccVpnvserver_secureprivateaccessurl_binding_DataSource_basic = `
	resource "citrixadc_vpnvserver" "tf_vpnvserver" {
		name        = "tf.citrix.example.com"
		servicetype = "SSL"
		ipv46       = "3.3.3.3"
		port        = 443
	}

	resource "citrixadc_vpnvserver_secureprivateaccessurl_binding" "tf_binding" {
		name                   = citrixadc_vpnvserver.tf_vpnvserver.name
		secureprivateaccessurl = "https://app.example.com/"
		depends_on             = [citrixadc_vpnvserver.tf_vpnvserver]
	}

	data "citrixadc_vpnvserver_secureprivateaccessurl_binding" "tf_binding" {
		name                   = citrixadc_vpnvserver.tf_vpnvserver.name
		secureprivateaccessurl = "https://app.example.com/"
		depends_on             = [citrixadc_vpnvserver_secureprivateaccessurl_binding.tf_binding]
	}
`

func TestAccVpnvserver_secureprivateaccessurl_binding_DataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVpnvserver_secureprivateaccessurl_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnvserver_secureprivateaccessurl_binding_DataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_vpnvserver_secureprivateaccessurl_binding.tf_binding", "name", "tf.citrix.example.com"),
					resource.TestCheckResourceAttr("data.citrixadc_vpnvserver_secureprivateaccessurl_binding.tf_binding", "secureprivateaccessurl", "https://app.example.com/"),
				),
			},
		},
	})
}
