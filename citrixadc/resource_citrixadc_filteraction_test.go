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
)

const testAccFilteraction_add = `

	resource "citrixadc_filteraction" "tf_filteraction" {
		name  = "tf_filteraction"
		qual  = "corrupt"
		value = "X-Forwarded-For"
	}
`
const testAccFilteraction_update = `

	resource "citrixadc_filteraction" "tf_filteraction" {
		name  = "tf_filteraction"
		qual  = "corrupt"
		value = "X-Forwarded"
	}
`

func TestAccFilteraction_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckFilteractionDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccFilteraction_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFilteractionExist("citrixadc_filteraction.tf_filteraction", nil),
					resource.TestCheckResourceAttr("citrixadc_filteraction.tf_filteraction", "name", "tf_filteraction"),
					resource.TestCheckResourceAttr("citrixadc_filteraction.tf_filteraction", "value", "X-Forwarded-For"),
				),
			},
			resource.TestStep{
				Config: testAccFilteraction_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFilteractionExist("citrixadc_filteraction.tf_filteraction", nil),
					resource.TestCheckResourceAttr("citrixadc_filteraction.tf_filteraction", "name", "tf_filteraction"),
					resource.TestCheckResourceAttr("citrixadc_filteraction.tf_filteraction", "value", "X-Forwarded"),
				),
			},
		},
	})
}

func testAccCheckFilteractionExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No filteraction name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(service.Filteraction.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("filteraction %s not found", n)
		}

		return nil
	}
}

func testAccCheckFilteractionDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_filteraction" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Filteraction.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("filteraction %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
