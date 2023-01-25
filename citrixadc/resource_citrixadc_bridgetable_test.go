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
	"strings"
)

const testAccBridgetable_basic = `

	resource "citrixadc_vlan" "tf_vlan" {
		vlanid    = 20
		aliasname = "Management VLAN"
	}
	resource "citrixadc_vxlan" "tf_vxlan" {
		vxlanid            = 123
		vlan               = citrixadc_vlan.tf_vlan.vlanid
		port               = 33
		dynamicrouting     = "DISABLED"
		ipv6dynamicrouting = "DISABLED"
		innervlantagging   = "ENABLED"
	}
	resource "citrixadc_bridgetable" "tf_bridgetable" {
		mac       = "00:00:00:00:00:01"
		vxlan     = citrixadc_vxlan.tf_vxlan.vxlanid
		vtep      = "2.34.5.6"
		bridgeage = "250"
	}
`

func TestAccBridgetable_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckBridgetableDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccBridgetable_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBridgetableExist("citrixadc_bridgetable.tf_bridgetable", nil),
				),
			},
		},
	})
}

func testAccCheckBridgetableExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No bridgetable name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

		findParams := service.FindParams{
			ResourceType: service.Bridgetable.Type(),
		}
		dataArray, err := nsClient.FindResourceArrayWithParams(findParams)
		if err != nil {
			return err
		}

		bindingId := rs.Primary.ID

		idSlice := strings.SplitN(bindingId, ",", 3)

		mac := idSlice[0]
		vxlan := idSlice[1]
		vtep := idSlice[2]

		foundIndex := -1
		for i, bridgetable := range dataArray {
			match := true
			if bridgetable["mac"] != mac {
				match = false
			}
			if bridgetable["vxlan"] != vxlan {
				match = false
			}
			if bridgetable["vtep"] != vtep {
				match = false
			}
			if match {
				foundIndex = i
				break
			}
		}

		if foundIndex == -1 {
			return fmt.Errorf("bridgetable %s not found", n)
		}

		return nil
	}
}

func testAccCheckBridgetableDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_bridgetable" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		findParams := service.FindParams{
			ResourceType: service.Bridgetable.Type(),
		}
		dataArray, err := nsClient.FindResourceArrayWithParams(findParams)
		if err != nil {
			return err
		}
		
		bindingId := rs.Primary.ID

		idSlice := strings.SplitN(bindingId, ",", 3)

		mac := idSlice[0]
		vxlan := idSlice[1]
		vtep := idSlice[2]

		foundIndex := -1
		for i, bridgetable := range dataArray {
			match := true
			if bridgetable["mac"] != mac {
				match = false
			}
			if bridgetable["vxlan"] != vxlan {
				match = false
			}
			if bridgetable["vtep"] != vtep {
				match = false
			}
			if match {
				foundIndex = i
				break
			}
		}
		
		if foundIndex != -1 {
			return fmt.Errorf("bridgetable %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
