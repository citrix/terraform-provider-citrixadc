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

const testAccVrid6_add = `
	resource "citrixadc_vrid6" "tf_vrid6" {
		vrid6_id             = 3
		priority             = 30
		preemption           = "DISABLED"
		sharing              = "DISABLED"
		tracking             = "NONE"
		trackifnumpriority   = 0
		preemptiondelaytimer = 0
	}
`
const testAccVrid6_update = `
	resource "citrixadc_vrid6" "tf_vrid6" {
		vrid6_id             = 3
		priority             = 50
		preemption           = "ENABLED"
		sharing              = "DISABLED"
		tracking             = "NONE"
		trackifnumpriority   = 0
		preemptiondelaytimer = 0
	}
`

func TestAccVrid6_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckVrid6Destroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVrid6_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVrid6Exist("citrixadc_vrid6.tf_vrid6", nil),
					resource.TestCheckResourceAttr("citrixadc_vrid6.tf_vrid6", "vrid6_id", "3"),
					resource.TestCheckResourceAttr("citrixadc_vrid6.tf_vrid6", "priority", "30"),
					resource.TestCheckResourceAttr("citrixadc_vrid6.tf_vrid6", "preemption", "DISABLED"),
				),
			},
			{
				Config: testAccVrid6_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVrid6Exist("citrixadc_vrid6.tf_vrid6", nil),
					resource.TestCheckResourceAttr("citrixadc_vrid6.tf_vrid6", "vrid6_id", "3"),
					resource.TestCheckResourceAttr("citrixadc_vrid6.tf_vrid6", "priority", "50"),
					resource.TestCheckResourceAttr("citrixadc_vrid6.tf_vrid6", "preemption", "ENABLED"),
				),
			},
		},
	})
}

func testAccCheckVrid6Exist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No vrid6 name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(service.Vrid6.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("vrid6 %s not found", n)
		}

		return nil
	}
}

func testAccCheckVrid6Destroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_vrid6" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Vrid6.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("vrid6 %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
