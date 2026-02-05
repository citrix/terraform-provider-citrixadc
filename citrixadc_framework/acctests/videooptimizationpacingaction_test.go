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

const testAccVideooptimizationpacingaction_add = `

	resource "citrixadc_videooptimizationpacingaction" "tf_pacingaction" {
		name 	= "tf_pacingaction"
		rate 	= 20
		comment = "Some Comment"
	}
`

const testAccVideooptimizationpacingaction_update = `

	resource "citrixadc_videooptimizationpacingaction" "tf_pacingaction" {
		name 	= "tf_pacingaction"
		rate 	= 10
		comment = "Some Comment"
	}
`

func TestAccVideooptimizationpacingaction_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVideooptimizationpacingactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVideooptimizationpacingaction_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVideooptimizationpacingactionExist("citrixadc_videooptimizationpacingaction.tf_pacingaction", nil),
					resource.TestCheckResourceAttr("citrixadc_videooptimizationpacingaction.tf_pacingaction", "name", "tf_pacingaction"),
					resource.TestCheckResourceAttr("citrixadc_videooptimizationpacingaction.tf_pacingaction", "rate", "20"),
				),
			},
			{
				Config: testAccVideooptimizationpacingaction_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVideooptimizationpacingactionExist("citrixadc_videooptimizationpacingaction.tf_pacingaction", nil),
					resource.TestCheckResourceAttr("citrixadc_videooptimizationpacingaction.tf_pacingaction", "name", "tf_pacingaction"),
					resource.TestCheckResourceAttr("citrixadc_videooptimizationpacingaction.tf_pacingaction", "rate", "10"),
				),
			},
		},
	})
}

func testAccCheckVideooptimizationpacingactionExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No videooptimizationpacingaction name is set")
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
		data, err := client.FindResource("videooptimizationpacingaction", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("videooptimizationpacingaction %s not found", n)
		}

		return nil
	}
}

func testAccCheckVideooptimizationpacingactionDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_videooptimizationpacingaction" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("videooptimizationpacingaction", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("videooptimizationpacingaction %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccVideooptimizationpacingactionDataSource_basic = `

	resource "citrixadc_videooptimizationpacingaction" "tf_pacingaction" {
		name 	= "tf_pacingaction"
		rate 	= 20
		comment = "Some Comment"
	}

	data "citrixadc_videooptimizationpacingaction" "tf_pacingaction" {
		name = citrixadc_videooptimizationpacingaction.tf_pacingaction.name
	}
`

func TestAccVideooptimizationpacingactionDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccVideooptimizationpacingactionDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_videooptimizationpacingaction.tf_pacingaction", "name", "tf_pacingaction"),
					resource.TestCheckResourceAttr("data.citrixadc_videooptimizationpacingaction.tf_pacingaction", "rate", "20"),
					resource.TestCheckResourceAttr("data.citrixadc_videooptimizationpacingaction.tf_pacingaction", "comment", "Some Comment"),
				),
			},
		},
	})
}
