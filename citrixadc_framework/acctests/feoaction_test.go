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

const testAccFeoaction_basic = `

	resource "citrixadc_feoaction" "tf_feoaction" {
		name              = "my_feoaction"
		cachemaxage       = 50
		imgshrinktoattrib = "false"
		imggiftopng       = "false"
	}
`
const testAccFeoaction_update = `

	resource "citrixadc_feoaction" "tf_feoaction" {
		name              = "my_feoaction"
		cachemaxage       = 40
		imgshrinktoattrib = "true"
		imggiftopng       = "true"
	}
`

func TestAccFeoaction_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckFeoactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccFeoaction_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFeoactionExist("citrixadc_feoaction.tf_feoaction", nil),
					resource.TestCheckResourceAttr("citrixadc_feoaction.tf_feoaction", "name", "my_feoaction"),
					resource.TestCheckResourceAttr("citrixadc_feoaction.tf_feoaction", "cachemaxage", "50"),
					resource.TestCheckResourceAttr("citrixadc_feoaction.tf_feoaction", "imgshrinktoattrib", "false"),
					resource.TestCheckResourceAttr("citrixadc_feoaction.tf_feoaction", "imggiftopng", "false"),
				),
			},
			{
				Config: testAccFeoaction_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFeoactionExist("citrixadc_feoaction.tf_feoaction", nil),
					resource.TestCheckResourceAttr("citrixadc_feoaction.tf_feoaction", "name", "my_feoaction"),
					resource.TestCheckResourceAttr("citrixadc_feoaction.tf_feoaction", "cachemaxage", "40"),
					resource.TestCheckResourceAttr("citrixadc_feoaction.tf_feoaction", "imgshrinktoattrib", "true"),
					resource.TestCheckResourceAttr("citrixadc_feoaction.tf_feoaction", "imggiftopng", "true"),
				),
			},
		},
	})
}

func testAccCheckFeoactionExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No feoaction name is set")
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
		data, err := client.FindResource("feoaction", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("feoaction %s not found", n)
		}

		return nil
	}
}

func testAccCheckFeoactionDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_feoaction" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("feoaction", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("feoaction %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccFeoactionDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccFeoactionDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_feoaction.tf_feoaction_ds", "name", "tf_feoaction_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_feoaction.tf_feoaction_ds", "cachemaxage", "60"),
					resource.TestCheckResourceAttr("data.citrixadc_feoaction.tf_feoaction_ds", "imgshrinktoattrib", "true"),
					resource.TestCheckResourceAttr("data.citrixadc_feoaction.tf_feoaction_ds", "imggiftopng", "true"),
					resource.TestCheckResourceAttr("data.citrixadc_feoaction.tf_feoaction_ds", "cssminify", "true"),
					resource.TestCheckResourceAttr("data.citrixadc_feoaction.tf_feoaction_ds", "jsminify", "true"),
					resource.TestCheckResourceAttr("data.citrixadc_feoaction.tf_feoaction_ds", "htmlminify", "true"),
				),
			},
		},
	})
}

const testAccFeoactionDataSource_basic = `

resource "citrixadc_feoaction" "tf_feoaction_ds" {
    name              = "tf_feoaction_ds"
    cachemaxage       = 60
    imgshrinktoattrib = "true"
    imggiftopng       = "true"
    cssminify         = "true"
    jsminify          = "true"
    htmlminify        = "true"
}

data "citrixadc_feoaction" "tf_feoaction_ds" {
    name = citrixadc_feoaction.tf_feoaction_ds.name
    depends_on = [citrixadc_feoaction.tf_feoaction_ds]
}

`
