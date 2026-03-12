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
	"strings"
	"testing"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testAccHanode_routemonitor6_binding_basic = `

resource "citrixadc_hanode_routemonitor6_binding" "tf_hanode_routemonitor6_binding" {
	hanode_id    = 0
	routemonitor = "fd7f:6bd8:cea9:f32d::/64"
	}
`

const testAccHanode_routemonitor6_binding_basic_step2 = `
`

const testAccHanode_routemonitor6_bindingDataSource_basic = `

resource "citrixadc_hanode_routemonitor6_binding" "tf_hanode_routemonitor6_binding" {
	hanode_id    = 0
	routemonitor = "fd7f:6bd8:ceb9:f32d::/64"
}

data "citrixadc_hanode_routemonitor6_binding" "tf_hanode_routemonitor6_binding" {
	hanode_id    = citrixadc_hanode_routemonitor6_binding.tf_hanode_routemonitor6_binding.hanode_id
	routemonitor = citrixadc_hanode_routemonitor6_binding.tf_hanode_routemonitor6_binding.routemonitor
	depends_on   = [citrixadc_hanode_routemonitor6_binding.tf_hanode_routemonitor6_binding]
}
`

func TestAccHanode_routemonitor6_binding_basic(t *testing.T) {
	if adcTestbed != "HA_PAIR" {
		t.Skipf("ADC testbed is %s. Expected HA.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckHanode_routemonitor6_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccHanode_routemonitor6_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHanode_routemonitor6_bindingExist("citrixadc_hanode_routemonitor6_binding.tf_hanode_routemonitor6_binding", nil),
				),
			},
			{
				Config: testAccHanode_routemonitor6_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHanode_routemonitor6_bindingNotExist("citrixadc_hanode_routemonitor6_binding.tf_hanode_routemonitor6_binding", "0,fd7f:6bd8:cea9:f32d::/64"),
				),
			},
		},
	})
}

func testAccCheckHanode_routemonitor6_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No hanode_routemonitor6_binding id is set")
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

		id := idSlice[0]
		routemonitor := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "hanode_routemonitor6_binding",
			ResourceName:             id,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the matching routemonitor
		found := false
		for _, v := range dataArr {
			if v["routemonitor"].(string) == routemonitor {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("hanode_routemonitor6_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckHanode_routemonitor6_bindingNotExist(n string, id string) resource.TestCheckFunc {
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

		id := idSlice[0]
		routemonitor := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "hanode_routemonitor6_binding",
			ResourceName:             id,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to hopefully not find the one with the matching routemonitor
		found := false
		for _, v := range dataArr {
			if v["routemonitor"].(string) == routemonitor {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("hanode_routemonitor6_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckHanode_routemonitor6_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_hanode_routemonitor6_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Hanode_routemonitor6_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("hanode_routemonitor6_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccHanode_routemonitor6_bindingDataSource_basic(t *testing.T) {
	if adcTestbed != "HA_PAIR" {
		t.Skipf("ADC testbed is %s. Expected HA.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccHanode_routemonitor6_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_hanode_routemonitor6_binding.tf_hanode_routemonitor6_binding", "hanode_id", "0"),
					resource.TestCheckResourceAttr("data.citrixadc_hanode_routemonitor6_binding.tf_hanode_routemonitor6_binding", "routemonitor", "fd7f:6bd8:ceb9:f32d::/64"),
				),
			},
		},
	})
}
