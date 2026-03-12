/*
Copyright 2024 Citrix Systems, Inc

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

const testAccLbaction_basic = `

	resource "citrixadc_lbaction" "tf_act" {
		name  = "tf_act"
		type  = "SELECTIONORDER"
		value = [1]
	}
  
`

func TestAccLbaction_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLbactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLbaction_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbactionExist("citrixadc_lbaction.tf_act", nil),
				),
			},
		},
	})
}

func testAccCheckLbactionExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lbaction name is set")
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
		data, err := client.FindResource("lbaction", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("lbaction %s not found", n)
		}

		return nil
	}
}

func testAccCheckLbactionDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_lbaction" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("lbaction", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("lbaction %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccLbactionDataSource_basic = `

	resource "citrixadc_lbaction" "tf_act_ds" {
		name  = "tf_act_ds"
		type  = "SELECTIONORDER"
		value = [1, 2]
	}

	data "citrixadc_lbaction" "tf_lbaction_ds" {
		name = citrixadc_lbaction.tf_act_ds.name
	}
`

func TestAccLbactionDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccLbactionDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_lbaction.tf_lbaction_ds", "name", "tf_act_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_lbaction.tf_lbaction_ds", "type", "SELECTIONORDER"),
				),
			},
		},
	})
}
