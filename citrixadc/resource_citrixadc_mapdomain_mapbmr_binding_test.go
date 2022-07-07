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

const testAccMapdomain_mapbmr_binding_basic = `
	resource "citrixadc_mapbmr" "tf_mapbmr" {
		name           = "tf_mapbmr"
		ruleipv6prefix = "2001:db8:abcd:12::/64"
		psidoffset     = 6
		eabitlength    = 16
		psidlength     = 8
	}
	resource "citrixadc_mapdmr" "tf_mapdmr" {
		name         = "tf_mapdmr"
		bripv6prefix = "2002:db8::/64"
	}
	resource "citrixadc_mapdomain" "tf_mapdomain" {
		name       = "tf_mapdomain"
		mapdmrname = citrixadc_mapdmr.tf_mapdmr.name
	}
	resource "citrixadc_mapdomain_mapbmr_binding" "tf_binding" {
		name       = citrixadc_mapdomain.tf_mapdomain.name
		mapbmrname = citrixadc_mapbmr.tf_mapbmr.name
	}
`

const testAccMapdomain_mapbmr_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion
	resource "citrixadc_mapbmr" "tf_mapbmr" {
		name           = "tf_mapbmr"
		ruleipv6prefix = "2001:db8:abcd:12::/64"
		psidoffset     = 6
		eabitlength    = 16
		psidlength     = 8
	}
	resource "citrixadc_mapdmr" "tf_mapdmr" {
		name         = "tf_mapdmr"
		bripv6prefix = "2002:db8::/64"
	}
	resource "citrixadc_mapdomain" "tf_mapdomain" {
		name       = "tf_mapdomain"
		mapdmrname = citrixadc_mapdmr.tf_mapdmr.name
	}
`

func TestAccMapdomain_mapbmr_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckMapdomain_mapbmr_bindingDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccMapdomain_mapbmr_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMapdomain_mapbmr_bindingExist("citrixadc_mapdomain_mapbmr_binding.tf_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_mapdomain_mapbmr_binding.tf_binding", "name", "tf_mapdomain"),
				),
			},
			resource.TestStep{
				Config: testAccMapdomain_mapbmr_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMapdomain_mapbmr_bindingNotExist("citrixadc_mapdomain_mapbmr_binding.tf_binding", "tf_mapdomain,tf_mapbmr"),
				),
			},
		},
	})
}

func testAccCheckMapdomain_mapbmr_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No mapdomain_mapbmr_binding id is set")
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
		mapbmrname := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "mapdomain_mapbmr_binding",
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
			if v["mapbmrname"].(string) == mapbmrname {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("mapdomain_mapbmr_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckMapdomain_mapbmr_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*NetScalerNitroClient).client

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}
		idSlice := strings.SplitN(id, ",", 2)

		name := idSlice[0]
		mapbmrname := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "mapdomain_mapbmr_binding",
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
			if v["mapbmrname"].(string) == mapbmrname {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("mapdomain_mapbmr_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckMapdomain_mapbmr_bindingDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_mapdomain_mapbmr_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource("mapdomain_mapbmr_binding", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("mapdomain_mapbmr_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
