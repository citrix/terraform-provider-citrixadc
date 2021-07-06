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
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const testAccPolicystringmap_basic_step1 = `

resource "citrixadc_policystringmap" "tf_policystringmap" {
    name = "tf_policystringmap"
    comment = "Some comment"
}
`

const testAccPolicystringmap_basic_step2 = `

resource "citrixadc_policystringmap" "tf_policystringmap" {
    name = "tf_policystringmap"
    comment = "Some other comment"
}
`

func TestAccPolicystringmap_basic(t *testing.T) {
	if adcTestbed != "STANDALONE" {
		t.Skipf("ADC testbed is %s. Expected STANDALONE.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPolicystringmapDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccPolicystringmap_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPolicystringmapExist("citrixadc_policystringmap.tf_policystringmap", nil),
				),
			},
			resource.TestStep{
				Config: testAccPolicystringmap_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPolicystringmapExist("citrixadc_policystringmap.tf_policystringmap", nil),
				),
			},
		},
	})
}

func testAccCheckPolicystringmapExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No policystringmap name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(service.Policystringmap.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("policystringmap %s not found", n)
		}

		return nil
	}
}

func testAccCheckPolicystringmapDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_policystringmap" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Policystringmap.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("policystringmap %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
