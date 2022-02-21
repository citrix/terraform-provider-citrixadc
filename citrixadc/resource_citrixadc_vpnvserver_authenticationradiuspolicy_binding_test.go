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
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"strings"
	"testing"
)

const testAccVpnvserver_authenticationradiuspolicy_binding_basic = `
# Since the authenticationradiuspolicy resource is not yet available on Terraform,
# the tf_radiuspolicy policy must be created by hand(manually) in order for the script to run correctly.
# You can do that by using the following Citrix ADC cli commands:
# add authenticationradiusaction tf_radiusaction -serverIP 2.3.2.1 -radKey WareTheLorax
# add authenticationradiuspolicy tf_radiuspolicy ns_true tf_radiusaction

# TODO authenticationradiusaction and authenticationradiuspolicy resource 

	resource "citrixadc_vpnvserver" "tf_vpnvserver" {
		name        = "tf_vpnvserver"
		servicetype = "SSL"
	}
	resource "citrixadc_vpnvserver_authenticationradiuspolicy_binding" "tf_bind" {
		name      = citrixadc_vpnvserver.tf_vpnvserver.name
		policy    = "tf_radiuspolicy"
		priority  = 20
	}
	
`

const testAccVpnvserver_authenticationradiuspolicy_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion
	resource "citrixadc_vpnvserver" "tf_vpnvserver" {
		name        = "tf_vpnvserver"
		servicetype = "SSL"
	}
`

func TestAccVpnvserver_authenticationradiuspolicy_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckVpnvserver_authenticationradiuspolicy_bindingDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccVpnvserver_authenticationradiuspolicy_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnvserver_authenticationradiuspolicy_bindingExist("citrixadc_vpnvserver_authenticationradiuspolicy_binding.tf_bind", nil),
				),
			},
			resource.TestStep{
				Config: testAccVpnvserver_authenticationradiuspolicy_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnvserver_authenticationradiuspolicy_bindingNotExist("citrixadc_vpnvserver_authenticationradiuspolicy_binding.tf_bind", "tf_vpnvserver,tf_radiuspolicy"),
				),
			},
		},
	})
}

func testAccCheckVpnvserver_authenticationradiuspolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No vpnvserver_authenticationradiuspolicy_binding id is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		client := testAccProvider.Meta().(*NetScalerNitroClient).client

		bindingId := rs.Primary.ID

		idSlice := strings.SplitN(bindingId, ",", 2)

		name := idSlice[0]
		policy := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "vpnvserver_authenticationradiuspolicy_binding",
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
			if v["policy"].(string) == policy {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("vpnvserver_authenticationradiuspolicy_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckVpnvserver_authenticationradiuspolicy_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*NetScalerNitroClient).client

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}
		idSlice := strings.SplitN(id, ",", 2)

		name := idSlice[0]
		policy := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "vpnvserver_authenticationradiuspolicy_binding",
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
			if v["policy"].(string) == policy {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("vpnvserver_authenticationradiuspolicy_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckVpnvserver_authenticationradiuspolicy_bindingDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_vpnvserver_authenticationradiuspolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Vpnvserver_authenticationradiuspolicy_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("vpnvserver_authenticationradiuspolicy_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
