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

func TestAccNsparam_basic(t *testing.T) {
	if adcTestbed != "STANDALONE" {
		t.Skipf("ADC testbed is %s. Expected STANDALONE.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNsparam_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsparamExist("citrixadc_nsparam.tf_nsparam", nil, map[string]interface{}{"maxconn": "10", "useproxyport": "DISABLED"}),
				),
			},
			resource.TestStep{
				Config: testAccNsparam_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsparamExist("citrixadc_nsparam.tf_nsparam", nil, map[string]interface{}{"maxconn": "0", "useproxyport": "ENABLED"}),
				),
			},
		},
	})
}

func testAccCheckNsparamExist(n string, id *string, expectedValues map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lb vserver name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(service.Nsparam.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("NS parameters %s not found", n)
		}

		if data["useproxyport"] != expectedValues["useproxyport"] {
			return fmt.Errorf("Expected value for \"useproxyport\" differs. Expected: \"%v\", Retrieved \"%v\"", expectedValues["proxyprotocol"], data["proxyprotocol"])
		}

		if data["maxconn"] != expectedValues["maxconn"] {
			return fmt.Errorf("Expected value for \"maxconn\" differs. Expected: \"%v\", Retrieved \"%v\"", expectedValues["maxconn"], data["maxconn"])
		}

		return nil
	}
}

const testAccNsparam_basic_step1 = `

resource "citrixadc_nsparam" "tf_nsparam" {
  maxconn = 10
  useproxyport = "DISABLED"
}
`

const testAccNsparam_basic_step2 = `

resource "citrixadc_nsparam" "tf_nsparam" {
  maxconn = 0
  useproxyport = "ENABLED"
}
`
