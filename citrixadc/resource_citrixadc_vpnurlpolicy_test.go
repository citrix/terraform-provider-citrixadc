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

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testAccVpnurlpolicy_add = `

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
	resource "citrixadc_vpnurlpolicy" "tf_vpnurlpolicy" {
		name = "new_policy"
		rule = "true"
		action = citrixadc_vpnurlaction.tf_vpnurlaction.name
	}
`
const testAccVpnurlpolicy_update = `

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
	resource "citrixadc_vpnurlpolicy" "tf_vpnurlpolicy" {
		name = "new_policy"
		rule = "false"
		action = citrixadc_vpnurlaction.tf_vpnurlaction.name
	}
`

func TestAccVpnurlpolicy_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckVpnurlpolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnurlpolicy_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnurlpolicyExist("citrixadc_vpnurlpolicy.tf_vpnurlpolicy", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnurlpolicy.tf_vpnurlpolicy", "name", "new_policy"),
					resource.TestCheckResourceAttr("citrixadc_vpnurlpolicy.tf_vpnurlpolicy", "rule", "true"),
					resource.TestCheckResourceAttr("citrixadc_vpnurlpolicy.tf_vpnurlpolicy", "action", "tf_vpnurlaction"),
				),
			},
			{
				Config: testAccVpnurlpolicy_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnurlpolicyExist("citrixadc_vpnurlpolicy.tf_vpnurlpolicy", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnurlpolicy.tf_vpnurlpolicy", "name", "new_policy"),
					resource.TestCheckResourceAttr("citrixadc_vpnurlpolicy.tf_vpnurlpolicy", "rule", "false"),
					resource.TestCheckResourceAttr("citrixadc_vpnurlpolicy.tf_vpnurlpolicy", "action", "tf_vpnurlaction"),
				),
			},
		},
	})
}

func testAccCheckVpnurlpolicyExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No vpnurlpolicy name is set")
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
		data, err := client.FindResource("vpnurlpolicy", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("vpnurlpolicy %s not found", n)
		}

		return nil
	}
}

func testAccCheckVpnurlpolicyDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_vpnurlpolicy" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("vpnurlpolicy", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("vpnurlpolicy %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
