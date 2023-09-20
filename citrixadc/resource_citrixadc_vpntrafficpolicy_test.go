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
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"testing"
)

const testAccVpntrafficpolicy_add = `
	resource "citrixadc_vpntrafficaction" "tf_vpntrafficaction" {
		name       = "Testingaction"
		fta        = "ON"
		hdx        = "ON"
		qual       = "http"
		sso        = "ON"
	
	}
	resource "citrixadc_vpntrafficpolicy" "tf_vpntrafficpolicy" {
		name   = "tf_vpntrafficpolicy"
		rule   = "HTTP.REQ.HEADER(\"User-Agent\").CONTAINS(\"CitrixReceiver\").NOT"
		action = citrixadc_vpntrafficaction.tf_vpntrafficaction.name	
	}
`
const testAccVpntrafficpolicy_update = `
	resource "citrixadc_vpntrafficaction" "tf_vpntrafficaction" {
		name       = "Testingaction"
		fta        = "ON"
		hdx        = "ON"
		qual       = "http"
		sso        = "ON"
	}
	resource "citrixadc_vpntrafficpolicy" "tf_vpntrafficpolicy" {
		name   = "tf_vpntrafficpolicy"
		rule   = "HTTP.REQ.HEADER(\"User-Agent\").CONTAINS(\"CitrixReceiver\")"
		action = citrixadc_vpntrafficaction.tf_vpntrafficaction.name	
	}
`

func TestAccVpntrafficpolicy_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckVpntrafficpolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVpntrafficpolicy_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpntrafficpolicyExist("citrixadc_vpntrafficpolicy.tf_vpntrafficpolicy", nil),
					resource.TestCheckResourceAttr("citrixadc_vpntrafficpolicy.tf_vpntrafficpolicy", "name", "tf_vpntrafficpolicy"),
					resource.TestCheckResourceAttr("citrixadc_vpntrafficpolicy.tf_vpntrafficpolicy", "rule", "HTTP.REQ.HEADER(\"User-Agent\").CONTAINS(\"CitrixReceiver\").NOT"),
					resource.TestCheckResourceAttr("citrixadc_vpntrafficpolicy.tf_vpntrafficpolicy", "action", "Testingaction"),
				),
			},
			{
				Config: testAccVpntrafficpolicy_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpntrafficpolicyExist("citrixadc_vpntrafficpolicy.tf_vpntrafficpolicy", nil),
					resource.TestCheckResourceAttr("citrixadc_vpntrafficpolicy.tf_vpntrafficpolicy", "name", "tf_vpntrafficpolicy"),
					resource.TestCheckResourceAttr("citrixadc_vpntrafficpolicy.tf_vpntrafficpolicy", "rule", "HTTP.REQ.HEADER(\"User-Agent\").CONTAINS(\"CitrixReceiver\")"),
					resource.TestCheckResourceAttr("citrixadc_vpntrafficpolicy.tf_vpntrafficpolicy", "action", "Testingaction"),
				),
			},
		},
	})
}

func testAccCheckVpntrafficpolicyExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No vpntrafficpolicy name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(service.Vpntrafficpolicy.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("vpntrafficpolicy %s not found", n)
		}

		return nil
	}
}

func testAccCheckVpntrafficpolicyDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_vpntrafficpolicy" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Vpntrafficpolicy.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("vpntrafficpolicy %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
