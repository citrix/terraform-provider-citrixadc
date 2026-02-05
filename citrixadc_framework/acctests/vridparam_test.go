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

const testAccVridparam_add = `

	resource "citrixadc_vridparam" "tf_vridparam" {
		sendtomaster  = "ENABLED"
		hellointerval = 400
		deadinterval  = 4
	}
`
const testAccVridparam_update = `

	resource "citrixadc_vridparam" "tf_vridparam" {
		sendtomaster  = "DISABLED"
		hellointerval = 1000
		deadinterval  = 3
	}
`

func TestAccVridparam_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccVridparam_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVridparamExist("citrixadc_vridparam.tf_vridparam", nil),
					resource.TestCheckResourceAttr("citrixadc_vridparam.tf_vridparam", "sendtomaster", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_vridparam.tf_vridparam", "hellointerval", "400"),
					resource.TestCheckResourceAttr("citrixadc_vridparam.tf_vridparam", "deadinterval", "4"),
				),
			},
			{
				Config: testAccVridparam_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVridparamExist("citrixadc_vridparam.tf_vridparam", nil),
					resource.TestCheckResourceAttr("citrixadc_vridparam.tf_vridparam", "sendtomaster", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_vridparam.tf_vridparam", "hellointerval", "1000"),
					resource.TestCheckResourceAttr("citrixadc_vridparam.tf_vridparam", "deadinterval", "3"),
				),
			},
		},
	})
}

func testAccCheckVridparamExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No vridparam name is set")
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
		data, err := client.FindResource(service.Vridparam.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("vridparam %s not found", n)
		}

		return nil
	}
}

const testAccVridparamDataSource_basic = `

	resource "citrixadc_vridparam" "tf_vridparam" {
		sendtomaster  = "ENABLED"
		hellointerval = 400
		deadinterval  = 4
	}

data "citrixadc_vridparam" "tf_vridparam" {
	depends_on = [citrixadc_vridparam.tf_vridparam]
}
`

func TestAccVridparamDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccVridparamDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_vridparam.tf_vridparam", "sendtomaster", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_vridparam.tf_vridparam", "hellointerval", "400"),
					resource.TestCheckResourceAttr("data.citrixadc_vridparam.tf_vridparam", "deadinterval", "4"),
				),
			},
		},
	})
}
