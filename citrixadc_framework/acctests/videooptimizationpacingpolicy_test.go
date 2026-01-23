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

const testAccVideooptimizationpacingpolicy_add = `

	resource "citrixadc_videooptimizationpacingaction" "tf_action" {
		name = "tf_action"
		rate = 10
	}
	
	resource "citrixadc_videooptimizationpacingpolicy" "tf_policy" {
		name   = "tf_policy"
		rule   = "true"
		action = citrixadc_videooptimizationpacingaction.tf_action.name
	}
`

const testAccVideooptimizationpacingpolicy_update = `

	resource "citrixadc_videooptimizationpacingaction" "tf_action" {
		name = "tf_action"
		rate = 10
	}
	
	resource "citrixadc_videooptimizationpacingpolicy" "tf_policy" {
		name   = "tf_policy"
		rule   = "false"
		action = citrixadc_videooptimizationpacingaction.tf_action.name
	}
`

func TestAccVideooptimizationpacingpolicy_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVideooptimizationpacingpolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVideooptimizationpacingpolicy_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVideooptimizationpacingpolicyExist("citrixadc_videooptimizationpacingpolicy.tf_policy", nil),
					resource.TestCheckResourceAttr("citrixadc_videooptimizationpacingpolicy.tf_policy", "name", "tf_policy"),
					resource.TestCheckResourceAttr("citrixadc_videooptimizationpacingpolicy.tf_policy", "rule", "true"),
				),
			},
			{
				Config: testAccVideooptimizationpacingpolicy_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVideooptimizationpacingpolicyExist("citrixadc_videooptimizationpacingpolicy.tf_policy", nil),
					resource.TestCheckResourceAttr("citrixadc_videooptimizationpacingpolicy.tf_policy", "name", "tf_policy"),
					resource.TestCheckResourceAttr("citrixadc_videooptimizationpacingpolicy.tf_policy", "rule", "false"),
				),
			},
		},
	})
}

func testAccCheckVideooptimizationpacingpolicyExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No videooptimizationpacingpolicy name is set")
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
		data, err := client.FindResource("videooptimizationpacingpolicy", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("videooptimizationpacingpolicy %s not found", n)
		}

		return nil
	}
}

func testAccCheckVideooptimizationpacingpolicyDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_videooptimizationpacingpolicy" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("videooptimizationpacingpolicy", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("videooptimizationpacingpolicy %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
