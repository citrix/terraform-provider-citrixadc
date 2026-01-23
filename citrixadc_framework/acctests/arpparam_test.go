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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"testing"
)

const testAccArpparam_add = `
	resource "citrixadc_arpparam" "tf_arpparam" {
		timeout         = 1000
		spoofvalidation = "ENABLED"
	}
`
const testAccArpparam_update = `
	resource "citrixadc_arpparam" "tf_arpparam" {
		timeout         = 1200
		spoofvalidation = "DISABLED"
	}
`

func TestAccArpparam_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccArpparam_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckArpparamExist("citrixadc_arpparam.tf_arpparam", nil),
					resource.TestCheckResourceAttr("citrixadc_arpparam.tf_arpparam", "timeout", "1000"),
					resource.TestCheckResourceAttr("citrixadc_arpparam.tf_arpparam", "spoofvalidation", "ENABLED"),
				),
			},
			{
				Config: testAccArpparam_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckArpparamExist("citrixadc_arpparam.tf_arpparam", nil),
					resource.TestCheckResourceAttr("citrixadc_arpparam.tf_arpparam", "timeout", "1200"),
					resource.TestCheckResourceAttr("citrixadc_arpparam.tf_arpparam", "spoofvalidation", "DISABLED"),
				),
			},
		},
	})
}

func testAccCheckArpparamExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No arpparam name is set")
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
		data, err := client.FindResource(service.Arpparam.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("arpparam %s not found", n)
		}

		return nil
	}
}
