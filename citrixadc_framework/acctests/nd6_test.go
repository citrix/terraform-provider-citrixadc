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
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNd6Destroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNd6_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNd6Exist("citrixadc_nd6.tf_nd6", nil),
					resource.TestCheckResourceAttr("citrixadc_nd6.tf_nd6", "neighbor", "2001::3"),
					resource.TestCheckResourceAttr("citrixadc_nd6.tf_nd6", "mac", "e6:ec:41:50:b1:d1"),
					resource.TestCheckResourceAttr("citrixadc_nd6.tf_nd6", "ifnum", "LO/1"),
				),
			},
			{
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

		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Nd6.Type(), rs.Primary.ID)

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
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_nd6" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Nd6.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("nd6 %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccNd6DataSource_basic = `

	resource "citrixadc_nd6" "tf_nd6_ds" {
		neighbor = "2001::5"
		mac      = "e6:ec:41:50:b1:d3"
		ifnum    = "LO/1"
	}

	data "citrixadc_nd6" "tf_nd6_ds_data" {
		neighbor = citrixadc_nd6.tf_nd6_ds.neighbor
	}
`

func TestAccNd6DataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNd6DataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_nd6.tf_nd6_ds_data", "neighbor", "2001::5"),
					resource.TestCheckResourceAttr("data.citrixadc_nd6.tf_nd6_ds_data", "mac", "e6:ec:41:50:b1:d3"),
					resource.TestCheckResourceAttr("data.citrixadc_nd6.tf_nd6_ds_data", "ifnum", "LO/1"),
				),
			},
		},
	})
}
