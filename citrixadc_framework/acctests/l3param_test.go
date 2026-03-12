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

const testAccL3param_basic = `
	resource "citrixadc_l3param" "tf_l3param" {
		srcnat               = "DISABLED"
		icmpgenratethreshold = 150
		overridernat         = "DISABLED"
		dropdfflag           = "DISABLED"
		implicitpbr		 = "DISABLED"
	}
  
`
const testAccL3param_update = `
	resource "citrixadc_l3param" "tf_l3param" {
		srcnat               = "ENABLED"
		icmpgenratethreshold = 200
		overridernat         = "ENABLED"
		dropdfflag           = "ENABLED"
		implicitpbr		 = "ENABLED"
	}
  
`

func TestAccL3param_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccL3param_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckL3paramExist("citrixadc_l3param.tf_l3param", nil),
					resource.TestCheckResourceAttr("citrixadc_l3param.tf_l3param", "srcnat", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_l3param.tf_l3param", "icmpgenratethreshold", "150"),
					resource.TestCheckResourceAttr("citrixadc_l3param.tf_l3param", "overridernat", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_l3param.tf_l3param", "dropdfflag", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_l3param.tf_l3param", "implicitpbr", "DISABLED"),
				),
			},
			{
				Config: testAccL3param_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckL3paramExist("citrixadc_l3param.tf_l3param", nil),
					resource.TestCheckResourceAttr("citrixadc_l3param.tf_l3param", "srcnat", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_l3param.tf_l3param", "icmpgenratethreshold", "200"),
					resource.TestCheckResourceAttr("citrixadc_l3param.tf_l3param", "overridernat", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_l3param.tf_l3param", "dropdfflag", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_l3param.tf_l3param", "implicitpbr", "ENABLED"),
				),
			},
		},
	})
}

func testAccCheckL3paramExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No l3param name is set")
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
		data, err := client.FindResource(service.L3param.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("l3param %s not found", n)
		}

		return nil
	}
}

const testAccL3paramDataSource_basic = `

	resource "citrixadc_l3param" "tf_l3param" {
		srcnat               = "DISABLED"
		icmpgenratethreshold = 150
		overridernat         = "DISABLED"
		dropdfflag           = "DISABLED"
		implicitpbr          = "DISABLED"
	}

	data "citrixadc_l3param" "tf_l3param" {
		depends_on = [citrixadc_l3param.tf_l3param]
	}
`

func TestAccL3paramDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccL3paramDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_l3param.tf_l3param", "srcnat", "DISABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_l3param.tf_l3param", "icmpgenratethreshold", "150"),
					resource.TestCheckResourceAttr("data.citrixadc_l3param.tf_l3param", "overridernat", "DISABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_l3param.tf_l3param", "dropdfflag", "DISABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_l3param.tf_l3param", "implicitpbr", "DISABLED"),
				),
			},
		},
	})
}
