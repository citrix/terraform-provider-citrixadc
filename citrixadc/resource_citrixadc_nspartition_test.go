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

const testAccNspartition_add = `
	resource "citrixadc_nspartition" "tf_nspartition" {
		partitionname = "tf_nspartition"
		maxbandwidth  = 1024
		minbandwidth  = 512
		maxconn       = 512
		maxmemlimit   = 11
	}
`
const testAccNspartition_update = `
	resource "citrixadc_nspartition" "tf_nspartition" {
		partitionname = "tf_nspartition"
		maxbandwidth  = 10240
		minbandwidth  = 512
		maxconn       = 512
		maxmemlimit   = 10
	}
`

func TestAccNspartition_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNspartitionDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNspartition_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNspartitionExist("citrixadc_nspartition.tf_nspartition", nil),
					resource.TestCheckResourceAttr("citrixadc_nspartition.tf_nspartition", "partitionname", "tf_nspartition"),
					resource.TestCheckResourceAttr("citrixadc_nspartition.tf_nspartition", "maxbandwidth", "1024"),
					resource.TestCheckResourceAttr("citrixadc_nspartition.tf_nspartition", "maxmemlimit", "11"),
				),
			},
			resource.TestStep{
				Config: testAccNspartition_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNspartitionExist("citrixadc_nspartition.tf_nspartition", nil),
					resource.TestCheckResourceAttr("citrixadc_nspartition.tf_nspartition", "partitionname", "tf_nspartition"),
					resource.TestCheckResourceAttr("citrixadc_nspartition.tf_nspartition", "maxbandwidth", "10240"),
					resource.TestCheckResourceAttr("citrixadc_nspartition.tf_nspartition", "maxmemlimit", "10"),
				),
			},
		},
	})
}

func testAccCheckNspartitionExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No nspartition name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(service.Nspartition.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("nspartition %s not found", n)
		}

		return nil
	}
}

func testAccCheckNspartitionDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_nspartition" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Nspartition.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("nspartition %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
