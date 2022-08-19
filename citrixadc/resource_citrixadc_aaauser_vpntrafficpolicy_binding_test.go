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

const testAccAaauser_vpntrafficpolicy_binding_basic = `

resource "citrixadc_aaauser_vpntrafficpolicy_binding" "tf_aaauser_vpntrafficpolicy_binding" {
	username = "user1"
	policy    = citrixadc_vpntrafficpolicy.tf_vpntrafficpolicy.name
	priority  = 100
  }
  
  resource "citrixadc_vpntrafficaction" "foo" {
	fta        = "ON"
	hdx        = "ON"
	name       = "Testingaction"
	qual       = "tcp"
	sso        = "ON"
  }
  resource "citrixadc_vpntrafficpolicy" "tf_vpntrafficpolicy" {
	name   = "tf_vpntrafficpolicy"
	rule   = "HTTP.REQ.HEADER(\"User-Agent\").CONTAINS(\"CitrixReceiver\").NOT"
	action = citrixadc_vpntrafficaction.foo.name
  }
`

const testAccAaauser_vpntrafficpolicy_binding_basic_step2 = `
resource "citrixadc_vpntrafficaction" "foo" {
	fta        = "ON"
	hdx        = "ON"
	name       = "Testingaction"
	qual       = "tcp"
	sso        = "ON"
  }
  resource "citrixadc_vpntrafficpolicy" "tf_vpntrafficpolicy" {
	name   = "tf_vpntrafficpolicy"
	rule   = "HTTP.REQ.HEADER(\"User-Agent\").CONTAINS(\"CitrixReceiver\").NOT"
	action = citrixadc_vpntrafficaction.foo.name
  }
`

func TestAccAaauser_vpntrafficpolicy_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAaauser_vpntrafficpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccAaauser_vpntrafficpolicy_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaauser_vpntrafficpolicy_bindingExist("citrixadc_aaauser_vpntrafficpolicy_binding.tf_aaauser_vpntrafficpolicy_binding", nil),
				),
			},
			resource.TestStep{
				Config: testAccAaauser_vpntrafficpolicy_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaauser_vpntrafficpolicy_bindingNotExist("citrixadc_aaauser_vpntrafficpolicy_binding.tf_aaauser_vpntrafficpolicy_binding", "user1,tf_vpntrafficpolicy"),
				),
			},
		},
	})
}

func testAccCheckAaauser_vpntrafficpolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No aaauser_vpntrafficpolicy_binding id is set")
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

		username := idSlice[0]
		policy := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "aaauser_vpntrafficpolicy_binding",
			ResourceName:             username,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the matching policy
		found := false
		for _, v := range dataArr {
			if v["policy"].(string) == policy {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("aaauser_vpntrafficpolicy_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckAaauser_vpntrafficpolicy_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*NetScalerNitroClient).client

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}
		idSlice := strings.SplitN(id, ",", 2)

		username := idSlice[0]
		policy := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "aaauser_vpntrafficpolicy_binding",
			ResourceName:             username,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to hopefully not find the one with the matching policy
		found := false
		for _, v := range dataArr {
			if v["policy"].(string) == policy {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("aaauser_vpntrafficpolicy_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckAaauser_vpntrafficpolicy_bindingDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_aaauser_vpntrafficpolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Aaauser_vpntrafficpolicy_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("aaauser_vpntrafficpolicy_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
