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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testAccSnmpengineid_basic = `

	resource "citrixadc_snmpengineid" "tf_snmpengineid" {
		engineid = "123456789012345"
	}
`
const testAccSnmpengineid_update = `

	resource "citrixadc_snmpengineid" "tf_snmpengineid" {
		engineid = "1234567890123456"
	}
`

func TestAccSnmpengineid_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccSnmpengineid_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSnmpengineidExist("citrixadc_snmpengineid.tf_snmpengineid", nil),
					resource.TestCheckResourceAttr("citrixadc_snmpengineid.tf_snmpengineid", "engineid", "123456789012345"),
				),
			},
			{
				Config: testAccSnmpengineid_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSnmpengineidExist("citrixadc_snmpengineid.tf_snmpengineid", nil),
					resource.TestCheckResourceAttr("citrixadc_snmpengineid.tf_snmpengineid", "engineid", "1234567890123456"),
				),
			},
		},
	})
}

func testAccCheckSnmpengineidExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No snmpengineid name is set")
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
		data, err := client.FindResource(service.Snmpengineid.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("snmpengineid %s not found", n)
		}

		return nil
	}
}

func TestAccSnmpengineidDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccSnmpengineidDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_snmpengineid.tf_snmpengineid_ds", "engineid", "1234567890abcdef"),
				),
			},
		},
	})
}

const testAccSnmpengineidDataSource_basic = `

resource "citrixadc_snmpengineid" "tf_snmpengineid_ds" {
	engineid  = "1234567890abcdef"
	ownernode = -1
}

data "citrixadc_snmpengineid" "tf_snmpengineid_ds" {
	ownernode = -1
	depends_on = [citrixadc_snmpengineid.tf_snmpengineid_ds]
}
`
