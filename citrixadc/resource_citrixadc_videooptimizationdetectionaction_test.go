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
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"testing"
)

const testAccVideooptimizationdetectionaction_add = `

	resource "citrixadc_videooptimizationdetectionaction" "tf_detectionaction" {
		name = "tf_videooptimizationdetectionaction"
		type = "clear_text_pd"
	}
  
`
const testAccVideooptimizationdetectionaction_update = `

	resource "citrixadc_videooptimizationdetectionaction" "tf_detectionaction" {
		name = "tf_videooptimizationdetectionaction"
		type = "clear_text_abr"
	}
  
`

func TestAccVideooptimizationdetectionaction_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckVideooptimizationdetectionactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVideooptimizationdetectionaction_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVideooptimizationdetectionactionExist("citrixadc_videooptimizationdetectionaction.tf_detectionaction", nil),
					resource.TestCheckResourceAttr("citrixadc_videooptimizationdetectionaction.tf_detectionaction", "name", "tf_videooptimizationdetectionaction"),
					resource.TestCheckResourceAttr("citrixadc_videooptimizationdetectionaction.tf_detectionaction", "type", "clear_text_pd"),
				),
			},
			{
				Config: testAccVideooptimizationdetectionaction_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVideooptimizationdetectionactionExist("citrixadc_videooptimizationdetectionaction.tf_detectionaction", nil),
					resource.TestCheckResourceAttr("citrixadc_videooptimizationdetectionaction.tf_detectionaction", "name", "tf_videooptimizationdetectionaction"),
					resource.TestCheckResourceAttr("citrixadc_videooptimizationdetectionaction.tf_detectionaction", "type", "clear_text_abr"),
				),
			},
		},
	})
}

func testAccCheckVideooptimizationdetectionactionExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No videooptimizationdetectionaction name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource("videooptimizationdetectionaction", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("videooptimizationdetectionaction %s not found", n)
		}

		return nil
	}
}

func testAccCheckVideooptimizationdetectionactionDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_videooptimizationdetectionaction" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource("videooptimizationdetectionaction", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("videooptimizationdetectionaction %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
