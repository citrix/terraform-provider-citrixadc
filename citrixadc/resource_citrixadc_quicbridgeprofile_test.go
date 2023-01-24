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

const testAccQuicbridgeprofile_add = `
	resource citrixadc_quicbridgeprofile tfAcc_quicbridge {
		name             = "tfAcc_quicbridge"
		routingalgorithm = "PLAINTEXT"
		serveridlength   = 4
	}
`

const testAccQuicbridgeprofile_update = `
	resource citrixadc_quicbridgeprofile tfAcc_quicbridge {
		name             = "tfAcc_quicbridge"
		routingalgorithm = "PLAINTEXT"
		serveridlength   = 6
	}
`

func TestAccQuicbridgeprofile_basic(t *testing.T) {
	if adcTestbed != "STANDALONE" {
		t.Skipf("ADC testbed is %s. Expected STANDALONE.", adcTestbed)
	}
	if isCpxRun {
		t.Skip("No support in CPX")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckQuicbridgeprofileDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccQuicbridgeprofile_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckQuicbridgeprofileExist("citrixadc_quicbridgeprofile.tfAcc_quicbridge", nil),
					resource.TestCheckResourceAttr("citrixadc_quicbridgeprofile.tfAcc_quicbridge", "routingalgorithm", "PLAINTEXT"),
					resource.TestCheckResourceAttr("citrixadc_quicbridgeprofile.tfAcc_quicbridge", "serveridlength", "4"),
				),
			},
			resource.TestStep{
				Config: testAccQuicbridgeprofile_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckQuicbridgeprofileExist("citrixadc_quicbridgeprofile.tfAcc_quicbridge", nil),
					resource.TestCheckResourceAttr("citrixadc_quicbridgeprofile.tfAcc_quicbridge", "routingalgorithm", "PLAINTEXT"),
					resource.TestCheckResourceAttr("citrixadc_quicbridgeprofile.tfAcc_quicbridge", "serveridlength", "6"),
				),
			},
		},
	})
}

func testAccCheckQuicbridgeprofileExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No quicbridgeprofile name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource("quicbridgeprofile", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("quicbridgeprofile %s not found", n)
		}

		return nil
	}
}

func testAccCheckQuicbridgeprofileDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_quicbridgeprofile" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource("quicbridgeprofile", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("quicbridgeprofile %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
