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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"testing"
)

const testAccVideooptimizationdetectionpolicy_add = `
	resource "citrixadc_videooptimizationdetectionaction" "tf_detectionaction" {
		name = "tf_videooptimizationdetectionaction"
		type = "clear_text_abr"
	}
	
	resource "citrixadc_videooptimizationdetectionpolicy" "tf_detectionpolicy" {
		name   = "tf_videooptimizationdetectionpolicy"
		rule   = "true"
		action = citrixadc_videooptimizationdetectionaction.tf_detectionaction.name
	}
`

const testAccVideooptimizationdetectionpolicy_update = `
	resource "citrixadc_videooptimizationdetectionaction" "tf_detectionaction" {
		name = "tf_videooptimizationdetectionaction"
		type = "clear_text_abr"
	}
	
	resource "citrixadc_videooptimizationdetectionpolicy" "tf_detectionpolicy" {
		name   = "tf_videooptimizationdetectionpolicy"
		rule   = "false"
		action = citrixadc_videooptimizationdetectionaction.tf_detectionaction.name
	}
`

func TestAccVideooptimizationdetectionpolicy_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckVideooptimizationdetectionpolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVideooptimizationdetectionpolicy_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVideooptimizationdetectionpolicyExist("citrixadc_videooptimizationdetectionpolicy.tf_detectionpolicy", nil),
					resource.TestCheckResourceAttr("citrixadc_videooptimizationdetectionpolicy.tf_detectionpolicy", "name", "tf_videooptimizationdetectionpolicy"),
					resource.TestCheckResourceAttr("citrixadc_videooptimizationdetectionpolicy.tf_detectionpolicy", "rule", "true"),
					resource.TestCheckResourceAttr("citrixadc_videooptimizationdetectionpolicy.tf_detectionpolicy", "action", "tf_videooptimizationdetectionaction"),
				),
			},
			{
				Config: testAccVideooptimizationdetectionpolicy_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVideooptimizationdetectionpolicyExist("citrixadc_videooptimizationdetectionpolicy.tf_detectionpolicy", nil),
					resource.TestCheckResourceAttr("citrixadc_videooptimizationdetectionpolicy.tf_detectionpolicy", "name", "tf_videooptimizationdetectionpolicy"),
					resource.TestCheckResourceAttr("citrixadc_videooptimizationdetectionpolicy.tf_detectionpolicy", "rule", "false"),
					resource.TestCheckResourceAttr("citrixadc_videooptimizationdetectionpolicy.tf_detectionpolicy", "action", "tf_videooptimizationdetectionaction"),
				),
			},
		},
	})
}

func testAccCheckVideooptimizationdetectionpolicyExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No videooptimizationdetectionpolicy name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		// Use the shared utility function to get a configured client
		client, err := testAccGetClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource("videooptimizationdetectionpolicy", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("videooptimizationdetectionpolicy %s not found", n)
		}

		return nil
	}
}

func testAccCheckVideooptimizationdetectionpolicyDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_videooptimizationdetectionpolicy" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("videooptimizationdetectionpolicy", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("videooptimizationdetectionpolicy %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
