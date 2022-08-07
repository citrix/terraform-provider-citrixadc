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

const testAccClusterinstance_basic = `

resource "citrixadc_clusterinstance" "tf_clusterinstance" {
	clid          = 1
	deadinterval  = 5
	hellointerval = 600
  }
  
`
const testAccClusterinstance_update = `

resource "citrixadc_clusterinstance" "tf_clusterinstance" {
	clid          = 1
	deadinterval  = 8
	hellointerval = 400
  }
  
`
func TestAccClusterinstance_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckClusterinstanceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccClusterinstance_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClusterinstanceExist("citrixadc_clusterinstance.tf_clusterinstance", nil),
					resource.TestCheckResourceAttr("citrixadc_clusterinstance.tf_clusterinstance", "deadinterval", "5"),
					resource.TestCheckResourceAttr("citrixadc_clusterinstance.tf_clusterinstance", "hellointerval", "600"),
				),
			},
			resource.TestStep{
				Config: testAccClusterinstance_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClusterinstanceExist("citrixadc_clusterinstance.tf_clusterinstance", nil),
					resource.TestCheckResourceAttr("citrixadc_clusterinstance.tf_clusterinstance", "deadinterval", "8"),
					resource.TestCheckResourceAttr("citrixadc_clusterinstance.tf_clusterinstance", "hellointerval", "400"),
				),
			},
		},
	})
}

func testAccCheckClusterinstanceExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No clusterinstance name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(service.Clusterinstance.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("clusterinstance %s not found", n)
		}

		return nil
	}
}

func testAccCheckClusterinstanceDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_clusterinstance" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Clusterinstance.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("clusterinstance %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
