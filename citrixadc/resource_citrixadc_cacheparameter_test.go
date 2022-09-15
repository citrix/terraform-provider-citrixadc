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

const testAccCacheparameter_basic = `

	resource "citrixadc_cacheparameter" "tf_cacheparameter" {
		memlimit    = "3500"
		maxpostlen  = "6000"
		verifyusing = "HOSTNAME"
	}
`
const testAccCacheparameter_update= `

	resource "citrixadc_cacheparameter" "tf_cacheparameter" {
		memlimit    = "3000"
		maxpostlen  = "6500"
		verifyusing = "HOSTNAME_AND_IP"
	}
`
func TestAccCacheparameter_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCacheparameter_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCacheparameterExist("citrixadc_cacheparameter.tf_cacheparameter", nil),
					resource.TestCheckResourceAttr("citrixadc_cacheparameter.tf_cacheparameter", "memlimit", "3500"),
					resource.TestCheckResourceAttr("citrixadc_cacheparameter.tf_cacheparameter", "maxpostlen", "6000"),
					resource.TestCheckResourceAttr("citrixadc_cacheparameter.tf_cacheparameter", "verifyusing", "HOSTNAME"),
				),
			},
			resource.TestStep{
				Config: testAccCacheparameter_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCacheparameterExist("citrixadc_cacheparameter.tf_cacheparameter", nil),
					resource.TestCheckResourceAttr("citrixadc_cacheparameter.tf_cacheparameter", "memlimit", "3000"),
					resource.TestCheckResourceAttr("citrixadc_cacheparameter.tf_cacheparameter", "maxpostlen", "6500"),
					resource.TestCheckResourceAttr("citrixadc_cacheparameter.tf_cacheparameter", "verifyusing", "HOSTNAME_AND_IP"),
				),
			},
		},
	})
}

func testAccCheckCacheparameterExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No cacheparameter name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(service.Cacheparameter.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("cacheparameter %s not found", n)
		}

		return nil
	}
}