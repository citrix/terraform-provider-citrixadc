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

const testAccNspartition_bridgegroup_binding_basic = `
	resource "citrixadc_nspartition" "tf_nspartition" {
		partitionname = "tf_nspartition"
		maxbandwidth  = 10240
		minbandwidth  = 512
		maxconn       = 512
		maxmemlimit   = 11
	}
	resource "citrixadc_bridgegroup" "tf_bridgegroup" {
		bridgegroup_id     = 2
		dynamicrouting     = "DISABLED"
		ipv6dynamicrouting = "DISABLED"
	}
	resource "citrixadc_nspartition_bridgegroup_binding" "tf_binding" {
		partitionname = citrixadc_nspartition.tf_nspartition.partitionname
		bridgegroup   = citrixadc_bridgegroup.tf_bridgegroup.bridgegroup_id
	}
`

const testAccNspartition_bridgegroup_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion
	resource "citrixadc_nspartition" "tf_nspartition" {
		partitionname = "tf_nspartition"
		maxbandwidth  = 10240
		minbandwidth  = 512
		maxconn       = 512
		maxmemlimit   = 11
	}
	resource "citrixadc_bridgegroup" "tf_bridgegroup" {
		bridgegroup_id     = 2
		dynamicrouting     = "DISABLED"
		ipv6dynamicrouting = "DISABLED"
	}
`

func TestAccNspartition_bridgegroup_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNspartition_bridgegroup_bindingDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNspartition_bridgegroup_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNspartition_bridgegroup_bindingExist("citrixadc_nspartition_bridgegroup_binding.tf_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_nspartition_bridgegroup_binding.tf_binding", "partitionname", "tf_nspartition"),
					resource.TestCheckResourceAttr("citrixadc_nspartition_bridgegroup_binding.tf_binding", "bridgegroup", "2"),
				),
			},
			resource.TestStep{
				Config: testAccNspartition_bridgegroup_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNspartition_bridgegroup_bindingNotExist("citrixadc_nspartition_bridgegroup_binding.tf_binding", "tf_nspartition,2"),
				),
			},
		},
	})
}

func testAccCheckNspartition_bridgegroup_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No nspartition_bridgegroup_binding id is set")
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

		partitionname := idSlice[0]
		bridgegroup := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "nspartition_bridgegroup_binding",
			ResourceName:             partitionname,
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
			return fmt.Errorf("nspartition_bridgegroup_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckNspartition_bridgegroup_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*NetScalerNitroClient).client

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}
		idSlice := strings.SplitN(id, ",", 2)

		partitionname := idSlice[0]
		bridgegroup := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "nspartition_bridgegroup_binding",
			ResourceName:             partitionname,
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
			return fmt.Errorf("nspartition_bridgegroup_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckNspartition_bridgegroup_bindingDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_nspartition_bridgegroup_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Nspartition_bridgegroup_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("nspartition_bridgegroup_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
