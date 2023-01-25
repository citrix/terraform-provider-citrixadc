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
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccRnatClear_basic(t *testing.T) {
	// if adcTestbed != "STANDALONE" {
	// 	t.Skipf("ADC testbed is %s. Expected STANDALONE.", adcTestbed)
	// }
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckRnatClearDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccRnatClear_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRnatClearExist("citrixadc_rnat_clear.foo", nil),
				),
			},
		},
	})
}

func testAccCheckRnatClearExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No rnat name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		//d, err := nsClient.FindFilteredResourceArray(service.RnatClear.Type(), map[string]string{"network": "192.168.96.0", "netmask": "255.255.240.0", "natip": "*"})
		d, err := nsClient.FindFilteredResourceArray(service.Rnat.Type(), map[string]string{"network": "192.168.96.0", "netmask": "255.255.240.0"})

		if err != nil {
			return err
		}

		if len(d) != 1 {
			return fmt.Errorf("RnatClear rule %s not found", n)
		}

		return nil
	}
}

func testAccCheckRnatClearDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_rnat_clear" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindFilteredResourceArray(service.Rnat.Type(), map[string]string{"network": "192.168.96.0", "netmask": "255.255.240.0", "natip": "*"})
		if err == nil {
			return fmt.Errorf("RnatClear rule %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccRnatClear_basic = `


resource "citrixadc_rnat_clear" "foo" {
	rnatsname = "foo"
	rnat {
           network = "192.168.96.0"
           netmask = "255.255.240.0"
         }
}
`
