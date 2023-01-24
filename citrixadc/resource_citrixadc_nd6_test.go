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

const testAccNd6_basic = `

	resource "citrixadc_nd6" "tf_nd6" {
		neighbor = "2001::3"
		mac      = "e6:ec:41:50:b1:d1"
		ifnum    = "LO/1"
	}
`
const testAccNd6_update = `

	resource "citrixadc_nd6" "tf_nd6" {
		neighbor = "2001::3"
		mac      = "e6:ec:41:50:b1:d2"
		vxlan    = 1
	}
`

func TestAccNd6_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNd6Destroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNd6_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNd6Exist("citrixadc_nd6.tf_nd6", nil),
					resource.TestCheckResourceAttr("citrixadc_nd6.tf_nd6", "neighbor", "2001::3"),
					resource.TestCheckResourceAttr("citrixadc_nd6.tf_nd6", "mac", "e6:ec:41:50:b1:d1"),
					resource.TestCheckResourceAttr("citrixadc_nd6.tf_nd6", "ifnum", "LO/1"),
				),
			},
			resource.TestStep{
				Config: testAccNd6_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNd6Exist("citrixadc_nd6.tf_nd6", nil),
					resource.TestCheckResourceAttr("citrixadc_nd6.tf_nd6", "neighbor", "2001::3"),
					resource.TestCheckResourceAttr("citrixadc_nd6.tf_nd6", "mac", "e6:ec:41:50:b1:d2"),
					resource.TestCheckResourceAttr("citrixadc_nd6.tf_nd6", "vxlan", "1"),
				),
			},
		},
	})
}

func testAccCheckNd6Exist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No nd6 name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(service.Nd6.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("nd6 %s not found", n)
		}

		return nil
	}
}

func testAccCheckNd6Destroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_nd6" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Nd6.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("nd6 %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
