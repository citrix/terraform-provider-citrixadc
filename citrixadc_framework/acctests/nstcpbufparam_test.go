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

const testAccNstcpbufparam_add = `
	resource "citrixadc_nstcpbufparam" "tf_nstcpbufparam" {
		size     = 32
		memlimit = 8
	}
`
const testAccNstcpbufparam_update = `
	resource "citrixadc_nstcpbufparam" "tf_nstcpbufparam" {
		size     = 64
		memlimit = 16
	}
`

func TestAccNstcpbufparam_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccNstcpbufparam_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNstcpbufparamExist("citrixadc_nstcpbufparam.tf_nstcpbufparam", nil),
					resource.TestCheckResourceAttr("citrixadc_nstcpbufparam.tf_nstcpbufparam", "size", "32"),
					resource.TestCheckResourceAttr("citrixadc_nstcpbufparam.tf_nstcpbufparam", "memlimit", "8"),
				),
			},
			{
				Config: testAccNstcpbufparam_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNstcpbufparamExist("citrixadc_nstcpbufparam.tf_nstcpbufparam", nil),
					resource.TestCheckResourceAttr("citrixadc_nstcpbufparam.tf_nstcpbufparam", "size", "64"),
					resource.TestCheckResourceAttr("citrixadc_nstcpbufparam.tf_nstcpbufparam", "memlimit", "16"),
				),
			},
		},
	})
}

func testAccCheckNstcpbufparamExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No nstcpbufparam name is set")
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
		data, err := client.FindResource(service.Nstcpbufparam.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("nstcpbufparam %s not found", n)
		}

		return nil
	}
}
