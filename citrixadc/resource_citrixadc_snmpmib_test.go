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

const testAccSnmpmib_basic = `

resource "citrixadc_snmpmib" "tf_snmpmib" {
	contact  = "phone_number"
	name     = "my_name"
	location = "LOCATION"
	customid = "CUSTOMER_ID"
  }
  
`
const testAccSnmpmib_update = `

resource "citrixadc_snmpmib" "tf_snmpmib" {
	contact  = "phone_number2"
	name     = "my_name2"
	location = "LOCATION2"
	customid = "CUSTOMER_ID2"
  }
  
`

func TestAccSnmpmib_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccSnmpmib_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSnmpmibExist("citrixadc_snmpmib.tf_snmpmib", nil),
					resource.TestCheckResourceAttr("citrixadc_snmpmib.tf_snmpmib","contact", "phone_number"),
					resource.TestCheckResourceAttr("citrixadc_snmpmib.tf_snmpmib","name", "my_name"),
					resource.TestCheckResourceAttr("citrixadc_snmpmib.tf_snmpmib","location", "LOCATION"),
					resource.TestCheckResourceAttr("citrixadc_snmpmib.tf_snmpmib","customid", "CUSTOMER_ID"),
				),
			},
			resource.TestStep{
				Config: testAccSnmpmib_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSnmpmibExist("citrixadc_snmpmib.tf_snmpmib", nil),
					resource.TestCheckResourceAttr("citrixadc_snmpmib.tf_snmpmib","contact", "phone_number2"),
					resource.TestCheckResourceAttr("citrixadc_snmpmib.tf_snmpmib","name", "my_name2"),
					resource.TestCheckResourceAttr("citrixadc_snmpmib.tf_snmpmib","location", "LOCATION2"),
					resource.TestCheckResourceAttr("citrixadc_snmpmib.tf_snmpmib","customid", "CUSTOMER_ID2"),
				),
			},
		},
	})
}

func testAccCheckSnmpmibExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No snmpmib name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(service.Snmpmib.Type(),"")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("snmpmib %s not found", n)
		}

		return nil
	}
}
