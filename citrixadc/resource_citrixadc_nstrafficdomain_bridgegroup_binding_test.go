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
	"strings"
	"testing"
)

const testAccNstrafficdomain_bridgegroup_binding_basic = `

	resource "citrixadc_nstrafficdomain" "tf_trafficdomain" {
		td        = 2
		aliasname = "tf_trafficdomain"
		vmac      = "DISABLED"
	}
	resource "citrixadc_bridgegroup" "tf_bridgegroup" {
		bridgegroup_id     = 2
		dynamicrouting     = "DISABLED"
		ipv6dynamicrouting = "DISABLED"
	}
	resource "citrixadc_nstrafficdomain_bridgegroup_binding" "tf_binding" {
		td          = citrixadc_nstrafficdomain.tf_trafficdomain.td
		bridgegroup = citrixadc_bridgegroup.tf_bridgegroup.bridgegroup_id
	}
`

const testAccNstrafficdomain_bridgegroup_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion
	
	resource "citrixadc_nstrafficdomain" "tf_trafficdomain" {
		td        = 2
		aliasname = "tf_trafficdomain"
		vmac      = "DISABLED"
	}
	resource "citrixadc_bridgegroup" "tf_bridgegroup" {
		bridgegroup_id     = 2
		dynamicrouting     = "DISABLED"
		ipv6dynamicrouting = "DISABLED"
	}
`

func TestAccNstrafficdomain_bridgegroup_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNstrafficdomain_bridgegroup_bindingDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNstrafficdomain_bridgegroup_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNstrafficdomain_bridgegroup_bindingExist("citrixadc_nstrafficdomain_bridgegroup_binding.tf_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_nstrafficdomain_bridgegroup_binding.tf_binding", "td", "2"),
					resource.TestCheckResourceAttr("citrixadc_nstrafficdomain_bridgegroup_binding.tf_binding", "bridgegroup", "2"),
				),
			},
			resource.TestStep{
				Config: testAccNstrafficdomain_bridgegroup_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNstrafficdomain_bridgegroup_bindingNotExist("citrixadc_nstrafficdomain_bridgegroup_binding.tf_binding", "2,2"),
				),
			},
		},
	})
}

func testAccCheckNstrafficdomain_bridgegroup_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No nstrafficdomain_bridgegroup_binding id is set")
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

		td := idSlice[0]
		bridgegroup := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "nstrafficdomain_bridgegroup_binding",
			ResourceName:             td,
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
			if v["bridgegroup"].(string) == bridgegroup {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("nstrafficdomain_bridgegroup_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckNstrafficdomain_bridgegroup_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*NetScalerNitroClient).client

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}
		idSlice := strings.SplitN(id, ",", 2)

		td := idSlice[0]
		bridgegroup := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "nstrafficdomain_bridgegroup_binding",
			ResourceName:             td,
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
			if v["bridgegroup"].(string) == bridgegroup {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("nstrafficdomain_bridgegroup_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckNstrafficdomain_bridgegroup_bindingDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_nstrafficdomain_bridgegroup_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Nstrafficdomain_bridgegroup_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("nstrafficdomain_bridgegroup_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
