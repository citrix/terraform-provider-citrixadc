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

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccSystemextramgmtcpu_basic(t *testing.T) {
	if adcTestbed != "STANDALONE_12CORES" {
		t.Skipf("ADC testbed is %s. Expected STANDALONE_12CORES.", adcTestbed)
	}
	if isCpxRun {
		t.Skip("CPX does not support the feature")
		// TODO actually we need a VPX with 12 cores licensed to test this resource
		// otherwise the systemextramgmtcpu enable action is a noop
	}
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccSystemextramgmtcpu_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSystemextramgmtcpuExist("citrixadc_systemextramgmtcpu.tf_extramgmtcpu", nil),
				),
			},
			resource.TestStep{
				Config: testAccSystemextramgmtcpu_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSystemextramgmtcpuExist("citrixadc_systemextramgmtcpu.tf_extramgmtcpu", nil),
				),
			},
		},
	})
}

func testAccCheckSystemextramgmtcpuExist(n string, id *string) resource.TestCheckFunc {
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
		data, err := nsClient.FindResource("systemextramgmtcpu", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("LB vserver %s not found", n)
		}

		return nil
	}
}

const testAccSystemextramgmtcpu_basic_step1 = `

resource "citrixadc_systemextramgmtcpu" "tf_extramgmtcpu" {
    enabled = true
    reboot = true
}

`

const testAccSystemextramgmtcpu_basic_step2 = `

resource "citrixadc_systemextramgmtcpu" "tf_extramgmtcpu" {
    enabled = false
    reboot = true
}

`
