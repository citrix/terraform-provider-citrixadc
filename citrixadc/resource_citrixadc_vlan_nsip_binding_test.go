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
	"log"
	"strings"
	"testing"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testAccVlan_nsip_binding_basic_step1 = `
resource "citrixadc_vlan" "tf_vlan" {
    vlanid 	  = 40
    aliasname = "Management VLAN"
}

resource "citrixadc_nsip" "tf_snip" {
    ipaddress = "10.222.74.145"
    type 	  = "SNIP"
    netmask   = "255.255.255.0"
    icmp 	  = "ENABLED"
    state 	  = "ENABLED"
}

resource "citrixadc_vlan_nsip_binding" "tf_bind" {
    vlanid 	  = citrixadc_vlan.tf_vlan.vlanid
    ipaddress = citrixadc_nsip.tf_snip.ipaddress
    netmask   = citrixadc_nsip.tf_snip.netmask
    td 		  = 0
}
`

const testAccVlan_nsip_binding_basic_step2 = `
resource "citrixadc_vlan" "tf_vlan" {
    vlanid    = 40
    aliasname = "Management VLAN"
}

resource "citrixadc_nsip" "tf_snip" {
    ipaddress = "10.222.74.145"
    type 	  = "SNIP"
    netmask   = "255.255.255.0"
    icmp      = "ENABLED"
    state     = "ENABLED"
}
`

func TestAccVlan_nsip_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckVlan_nsip_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVlan_nsip_binding_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlan_nsip_bindingExist("citrixadc_vlan_nsip_binding.tf_bind", nil),
				),
			},
			{
				Config: testAccVlan_nsip_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVlan_nsip_bindingNotExist("40,10.222.74.145"),
				),
			},
		},
	})
}

func testAccCheckVlan_nsip_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No vlan_nsip_binding name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		// Use the shared utility function to get a configured client
		client, err := testAccGetClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		bindingId := rs.Primary.ID

		idSlice := strings.SplitN(bindingId, ",", 2)

		vlanid := idSlice[0]
		ipaddress := idSlice[1]

		log.Printf("[DEBUG] citrixadc-provider: Reading vlan_nsip_bindingName state %s", bindingId)
		findParams := service.FindParams{
			ResourceType:             "vlan_nsip_binding",
			ResourceName:             vlanid,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Resource is missing
		if len(dataArr) == 0 {
			return fmt.Errorf("Cannot find vlan_nsip_binding %s", bindingId)
		}

		// Iterate through results to find the one with the right monitor name
		foundIndex := -1
		for i, v := range dataArr {
			if v["ipaddress"].(string) == ipaddress {
				foundIndex = i
				break
			}
		}

		// Resource is missing
		if foundIndex == -1 {
			return fmt.Errorf("Resource missing vlan_nsip_binding %s", bindingId)
		}

		return nil
	}
}

func testAccCheckVlan_nsip_bindingNotExist(bindingId string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		// Use the shared utility function to get a configured client
		client, err := testAccGetClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		idSlice := strings.SplitN(bindingId, ",", 2)

		vlanid := idSlice[0]
		ipaddress := idSlice[1]

		log.Printf("[DEBUG] citrixadc-provider: Reading vlan_nsip_bindingName state %s", bindingId)
		findParams := service.FindParams{
			ResourceType:             "vlan_nsip_binding",
			ResourceName:             vlanid,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Resource is missing
		if len(dataArr) == 0 {
			return nil
		}

		// Iterate through results to find the one with the right monitor name
		foundIndex := -1
		for i, v := range dataArr {
			if v["ipaddress"].(string) == ipaddress {
				foundIndex = i
				break
			}
		}

		// Resource is missing
		if foundIndex == -1 {
			return nil
		}

		return fmt.Errorf("Resource still exists vlan_nsip_binding %s", bindingId)
	}
}

func testAccCheckVlan_nsip_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_vlan_nsip_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Vlan_nsip_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("vlan_nsip_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
