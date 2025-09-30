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

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

const testAccLacp_basic = `

resource "citrixadc_lacp" "tf_lacp" {
	syspriority = 30
  }
  
`
const testAccLacp_update = `

resource "citrixadc_lacp" "tf_lacp" {
	syspriority = 50
  }
  
`

func TestAccLacp_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				Config: testAccLacp_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLacpExist("citrixadc_lacp.tf_lacp", nil),
					resource.TestCheckResourceAttr("citrixadc_lacp.tf_lacp", "syspriority", "30"),
					resource.TestCheckResourceAttr("citrixadc_lacp.tf_lacp", "ownernode", "255"),
				),
			},
			{
				Config: testAccLacp_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLacpExist("citrixadc_lacp.tf_lacp", nil),
					resource.TestCheckResourceAttr("citrixadc_lacp.tf_lacp", "syspriority", "50"),
					resource.TestCheckResourceAttr("citrixadc_lacp.tf_lacp", "ownernode", "255"),
				),
			},
		},
	})
}

func testAccCheckLacpExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lacp name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(service.Lacp.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("lacp %s not found", n)
		}

		return nil
	}
}
