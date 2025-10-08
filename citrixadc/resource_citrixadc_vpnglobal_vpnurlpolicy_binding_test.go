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
	"testing"
)

const testAccVpnglobal_vpnurlpolicy_binding_basic = `

	resource "citrixadc_vpnurlaction" "tf_vpnurlaction" {
		name             = "tf_vpnurlaction"
		linkname         = "new_link"
		actualurl        = "http://www.citrix.com"
		applicationtype  = "CVPN"
		clientlessaccess = "OFF"
		comment          = "Testing"
		ssotype          = "unifiedgateway"
		vservername      = "vserver1"
	}
	resource "citrixadc_vpnurlpolicy" "citrixadc_vpnurlpolicy" {
		name   = "new_policy"
		rule   = "true"
		action = citrixadc_vpnurlaction.tf_vpnurlaction.name
	}
	resource "citrixadc_vpnglobal_vpnurlpolicy_binding" "tf_bind" {
		policyname = citrixadc_vpnurlpolicy.citrixadc_vpnurlpolicy.name
		priority = "20"
	}
`

const testAccVpnglobal_vpnurlpolicy_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion

	resource "citrixadc_vpnurlaction" "tf_vpnurlaction" {
		name             = "tf_vpnurlaction"
		linkname         = "new_link"
		actualurl        = "http://www.citrix.com"
		applicationtype  = "CVPN"
		clientlessaccess = "OFF"
		comment          = "Testing"
		ssotype          = "unifiedgateway"
		vservername      = "vserver1"
	}
	resource "citrixadc_vpnurlpolicy" "citrixadc_vpnurlpolicy" {
		name   = "new_policy"
		rule   = "true"
		action = citrixadc_vpnurlaction.tf_vpnurlaction.name
	}
`

func TestAccVpnglobal_vpnurlpolicy_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckVpnglobal_vpnurlpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnglobal_vpnurlpolicy_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnglobal_vpnurlpolicy_bindingExist("citrixadc_vpnglobal_vpnurlpolicy_binding.tf_bind", nil),
				),
			},
			{
				Config: testAccVpnglobal_vpnurlpolicy_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnglobal_vpnurlpolicy_bindingNotExist("citrixadc_vpnglobal_vpnurlpolicy_binding.tf_bind", "policyname"),
				),
			},
		},
	})
}

func testAccCheckVpnglobal_vpnurlpolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No vpnglobal_vpnurlpolicy_binding id is set")
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

		policyname := rs.Primary.ID

		findParams := service.FindParams{
			ResourceType:             "vpnglobal_vpnurlpolicy_binding",
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
			if v["policyname"].(string) == policyname {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("vpnglobal_vpnurlpolicy_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckVpnglobal_vpnurlpolicy_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		policyname := id
		findParams := service.FindParams{
			ResourceType:             "vpnglobal_vpnurlpolicy_binding",
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
			if v["policyname"].(string) == policyname {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("vpnglobal_vpnurlpolicy_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckVpnglobal_vpnurlpolicy_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_vpnglobal_vpnurlpolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("vpnglobal_vpnurlpolicy_binding", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("vpnglobal_vpnurlpolicy_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
