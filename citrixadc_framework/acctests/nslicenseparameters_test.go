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

const testAccNslicenseparameters_add = `

	resource "citrixadc_nslicenseparameters" "tf_nslicenseparameters" {
		alert1gracetimeout = 8
		alert2gracetimeout = 200
		licenseexpiryalerttime = 30
		inventoryrefreshinterval = 50
		heartbeatinterval = 45
	}
`
const testAccNslicenseparameters_update = `

	resource "citrixadc_nslicenseparameters" "tf_nslicenseparameters" {
		alert1gracetimeout = 6
		alert2gracetimeout = 240
		licenseexpiryalerttime = 40
		inventoryrefreshinterval = 60
		heartbeatinterval = 55
	}
`

func TestAccNslicenseparameters_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccNslicenseparameters_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNslicenseparametersExist("citrixadc_nslicenseparameters.tf_nslicenseparameters", nil),
					resource.TestCheckResourceAttr("citrixadc_nslicenseparameters.tf_nslicenseparameters", "alert1gracetimeout", "8"),
					resource.TestCheckResourceAttr("citrixadc_nslicenseparameters.tf_nslicenseparameters", "alert2gracetimeout", "200"),
					resource.TestCheckResourceAttr("citrixadc_nslicenseparameters.tf_nslicenseparameters", "licenseexpiryalerttime", "30"),
					resource.TestCheckResourceAttr("citrixadc_nslicenseparameters.tf_nslicenseparameters", "inventoryrefreshinterval", "50"),
					resource.TestCheckResourceAttr("citrixadc_nslicenseparameters.tf_nslicenseparameters", "heartbeatinterval", "45"),
				),
			},
			{
				Config: testAccNslicenseparameters_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNslicenseparametersExist("citrixadc_nslicenseparameters.tf_nslicenseparameters", nil),
					resource.TestCheckResourceAttr("citrixadc_nslicenseparameters.tf_nslicenseparameters", "alert1gracetimeout", "6"),
					resource.TestCheckResourceAttr("citrixadc_nslicenseparameters.tf_nslicenseparameters", "alert2gracetimeout", "240"),
					resource.TestCheckResourceAttr("citrixadc_nslicenseparameters.tf_nslicenseparameters", "licenseexpiryalerttime", "40"),
					resource.TestCheckResourceAttr("citrixadc_nslicenseparameters.tf_nslicenseparameters", "inventoryrefreshinterval", "60"),
					resource.TestCheckResourceAttr("citrixadc_nslicenseparameters.tf_nslicenseparameters", "heartbeatinterval", "55"),
				),
			},
		},
	})
}

func testAccCheckNslicenseparametersExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No nslicenseparameters name is set")
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
		data, err := client.FindResource("nslicenseparameters", "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("nslicenseparameters %s not found", n)
		}

		return nil
	}
}
