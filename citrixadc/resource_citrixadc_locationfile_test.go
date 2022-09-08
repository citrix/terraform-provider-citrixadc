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
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"testing"
)

const testAccLocationfile_basic = `

	resource "citrixadc_locationfile" "tf_locationfile" {
		locationfile = "/var/netscaler/inbuilt_db/Citrix_Netscaler_InBuilt_GeoIP_DB_IPv4"
		format       = "netscaler"
	}
`


func TestAccLocationfile_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckLocationfileDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccLocationfile_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLocationfileExist("citrixadc_locationfile.tf_locationfile", nil),
					resource.TestCheckResourceAttr("citrixadc_locationfile.tf_locationfile", "locationfile", "/var/netscaler/inbuilt_db/Citrix_Netscaler_InBuilt_GeoIP_DB_IPv4"),
					resource.TestCheckResourceAttr("citrixadc_locationfile.tf_locationfile", "format", "netscaler"),
				),
			},
		},
	})
}

func testAccCheckLocationfileExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No locationfile name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(service.Locationfile.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("locationfile %s not found", n)
		}

		return nil
	}
}

func testAccCheckLocationfileDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_locationfile" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		data, _ := nsClient.FindResource(service.Locationfile.Type(), "")
		// if err == nil {
		// 	return fmt.Errorf("locationfile %s still exists", rs.Primary.ID)
		// }
		if data["locationfile"] == rs.Primary.Attributes["locationfile"] {
			return fmt.Errorf("locationfile %s still exists", rs.Primary.ID)
		}
	}

	return nil
}
