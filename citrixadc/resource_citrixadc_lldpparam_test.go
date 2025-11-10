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

const testAccLldpparam_basic = `
	resource "citrixadc_lldpparam" "tf_lldpparam" {
		holdtimetxmult = 3
		mode           = "TRANSMITTER"
		timer          = 40
	}
`
const testAccLldpparam_update = `
	resource "citrixadc_lldpparam" "tf_lldpparam" {
		holdtimetxmult = 10
		mode           = "RECEIVER"
		timer          = 60
	}
`

func TestAccLldpparam_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      nil,
		Steps: []resource.TestStep{
			{
				Config: testAccLldpparam_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLldpparamExist("citrixadc_lldpparam.tf_lldpparam", nil),
					resource.TestCheckResourceAttr("citrixadc_lldpparam.tf_lldpparam", "holdtimetxmult", "3"),
					resource.TestCheckResourceAttr("citrixadc_lldpparam.tf_lldpparam", "mode", "TRANSMITTER"),
					resource.TestCheckResourceAttr("citrixadc_lldpparam.tf_lldpparam", "timer", "40"),
				),
			},
			{
				Config: testAccLldpparam_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLldpparamExist("citrixadc_lldpparam.tf_lldpparam", nil),
					resource.TestCheckResourceAttr("citrixadc_lldpparam.tf_lldpparam", "holdtimetxmult", "10"),
					resource.TestCheckResourceAttr("citrixadc_lldpparam.tf_lldpparam", "mode", "RECEIVER"),
					resource.TestCheckResourceAttr("citrixadc_lldpparam.tf_lldpparam", "timer", "60"),
				),
			},
		},
	})
}

func testAccCheckLldpparamExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lldpparam name is set")
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
		data, err := client.FindResource("lldpparam", "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("lldpparam %s not found", n)
		}

		return nil
	}
}
