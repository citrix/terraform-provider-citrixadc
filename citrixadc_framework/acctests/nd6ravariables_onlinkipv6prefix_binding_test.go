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
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"strings"
	"testing"
)

const testAccNd6ravariables_onlinkipv6prefix_binding_basic = `

	resource "citrixadc_vlan" "tf_vlan" {
		vlanid 		= 40
		aliasname 	= "Management VLAN"
	}
	resource "citrixadc_onlinkipv6prefix" "tf_onlinkipv6prefix" {
		ipv6prefix      = "2003::/64"
		onlinkprefix    = "YES"
		autonomusprefix = "NO"
	}

	resource "citrixadc_nd6ravariables_onlinkipv6prefix_binding" "tf_nd6ravariables_onlinkipv6prefix_binding" {
		vlan      = citrixadc_vlan.tf_vlan.vlanid
		ipv6prefix = citrixadc_onlinkipv6prefix.tf_onlinkipv6prefix.ipv6prefix
	}
`

const testAccNd6ravariables_onlinkipv6prefix_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion

	resource "citrixadc_vlan" "tf_vlan" {
		vlanid 		= 40
		aliasname 	= "Management VLAN"
	}
	resource "citrixadc_onlinkipv6prefix" "tf_onlinkipv6prefix" {
		ipv6prefix      = "2003::/64"
		onlinkprefix    = "YES"
		autonomusprefix = "NO"
	}
`

func TestAccNd6ravariables_onlinkipv6prefix_binding_basic(t *testing.T) {
	t.Skip("TODO: Need to find a way to test this resource!")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNd6ravariables_onlinkipv6prefix_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNd6ravariables_onlinkipv6prefix_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNd6ravariables_onlinkipv6prefix_bindingExist("citrixadc_nd6ravariables_onlinkipv6prefix_binding.tf_nd6ravariables_onlinkipv6prefix_binding", nil),
				),
			},
			{
				Config: testAccNd6ravariables_onlinkipv6prefix_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNd6ravariables_onlinkipv6prefix_bindingNotExist("citrixadc_nd6ravariables_onlinkipv6prefix_binding.tf_nd6ravariables_onlinkipv6prefix_binding", "40,2003::/64"),
				),
			},
		},
	})
}

func testAccCheckNd6ravariables_onlinkipv6prefix_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No nd6ravariables_onlinkipv6prefix_binding id is set")
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

		vlan := idSlice[0]
		ipv6prefix := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "nd6ravariables_onlinkipv6prefix_binding",
			ResourceName:             vlan,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the matching ipv6prefix
		found := false
		for _, v := range dataArr {
			if v["ipv6prefix"].(string) == ipv6prefix {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("nd6ravariables_onlinkipv6prefix_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckNd6ravariables_onlinkipv6prefix_bindingNotExist(n string, id string) resource.TestCheckFunc {
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

		vlan := idSlice[0]
		ipv6prefix := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "nd6ravariables_onlinkipv6prefix_binding",
			ResourceName:             vlan,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to hopefully not find the one with the matching ipv6prefix
		found := false
		for _, v := range dataArr {
			if v["ipv6prefix"].(string) == ipv6prefix {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("nd6ravariables_onlinkipv6prefix_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckNd6ravariables_onlinkipv6prefix_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_nd6ravariables_onlinkipv6prefix_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Nd6ravariables_onlinkipv6prefix_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("nd6ravariables_onlinkipv6prefix_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
