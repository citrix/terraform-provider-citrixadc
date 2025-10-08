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

const testAccSystemparameter_basic = `

resource "citrixadc_systemparameter" "tf_systemparameter" {
    rbaonresponse = "ENABLED"
    natpcbforceflushlimit = 3000
    natpcbrstontimeout = "DISABLED"
    timeout = 500
    doppler = "ENABLED"
}
`

const testAccSystemparameter_update = `

resource "citrixadc_systemparameter" "tf_systemparameter" {
    rbaonresponse = "DISABLED"
    natpcbforceflushlimit = 2000
    natpcbrstontimeout = "ENABLED"
    timeout = 600
    doppler = "DISABLED"
}
`

func TestAccSystemparameter_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      nil,
		Steps: []resource.TestStep{
			{
				Config: testAccSystemparameter_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSystemparameterExist("citrixadc_systemparameter.tf_systemparameter", nil),
					resource.TestCheckResourceAttr("citrixadc_systemparameter.tf_systemparameter", "rbaonresponse", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_systemparameter.tf_systemparameter", "natpcbforceflushlimit", "3000"),
					resource.TestCheckResourceAttr("citrixadc_systemparameter.tf_systemparameter", "natpcbrstontimeout", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_systemparameter.tf_systemparameter", "timeout", "500"),
					resource.TestCheckResourceAttr("citrixadc_systemparameter.tf_systemparameter", "doppler", "ENABLED"),
				),
			},
			{
				Config: testAccSystemparameter_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSystemparameterExist("citrixadc_systemparameter.tf_systemparameter", nil),
					resource.TestCheckResourceAttr("citrixadc_systemparameter.tf_systemparameter", "rbaonresponse", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_systemparameter.tf_systemparameter", "natpcbforceflushlimit", "2000"),
					resource.TestCheckResourceAttr("citrixadc_systemparameter.tf_systemparameter", "natpcbrstontimeout", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_systemparameter.tf_systemparameter", "timeout", "600"),
					resource.TestCheckResourceAttr("citrixadc_systemparameter.tf_systemparameter", "doppler", "DISABLED"),
				),
			},
		},
	})
}

func testAccCheckSystemparameterExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No systemparameter name is set")
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
		data, err := client.FindResource(service.Systemparameter.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("systemparameter %s not found", n)
		}

		return nil
	}
}
