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

func TestAccNsip_basic(t *testing.T) {
	if adcTestbed != "STANDALONE" {
		t.Skipf("ADC testbed is %s. Expected STANDALONE.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNsipDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNsip_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsipExist("citrixadc_nsip.tf_test_nsip", nil),
				),
			},
			resource.TestStep{
				Config: testAccNsip_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsipExist("citrixadc_nsip.tf_test_nsip", nil),
				),
			},
			resource.TestStep{
				Config: testAccNsip_basic_step3,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsipExist("citrixadc_nsip.tf_test_nsip", nil),
				),
			},
			resource.TestStep{
				Config: testAccNsip_basic_step4,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsipExist("citrixadc_nsip.tf_test_nsip", nil),
				),
			},
		},
	})
}

const testAccNsip_mptcpadvertise = `
	resource "citrixadc_nsip" "tf_test_nsip_mptcpadvertise" {
		ipaddress = "192.168.1.55"
		type = "VIP"
		netmask = "255.255.255.0"
		icmp = "ENABLED"
		mptcpadvertise = "YES"
	}
`

func TestAccNsip_mptcpadvertise(t *testing.T) {
	if adcTestbed != "STANDALONE" {
		t.Skipf("ADC testbed is %s. Expected STANDALONE.", adcTestbed)
	}
	if isCpxRun {
		t.Skip("No support in CPX")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNsipDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNsip_mptcpadvertise,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsipExist("citrixadc_nsip.tf_test_nsip_mptcpadvertise", nil),
					resource.TestCheckResourceAttr("citrixadc_nsip.tf_test_nsip_mptcpadvertise", "mptcpadvertise", "YES"),
				),
			},
		},
	})
}

func testAccCheckNsipExist(n string, id *string) resource.TestCheckFunc {
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
		data, err := nsClient.FindResource(service.Nsip.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("LB vserver %s not found", n)
		}

		return nil
	}
}

func testAccCheckNsipDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_nsip" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Nsip.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("LB vserver %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccNsip_basic_step1 = `

resource "citrixadc_nsip" "tf_test_nsip" {
    ipaddress = "192.168.2.55"
    type = "VIP"
    netmask = "255.255.255.0"
    icmp = "ENABLED"
}
`

const testAccNsip_basic_step2 = `

resource "citrixadc_nsip" "tf_test_nsip" {
    ipaddress = "192.168.2.55"
    type = "VIP"
    netmask = "255.255.254.0"
    icmp = "ENABLED"
}
`

const testAccNsip_basic_step3 = `

resource "citrixadc_nsip" "tf_test_nsip" {
    ipaddress = "192.168.2.55"
    type = "VIP"
    netmask = "255.255.254.0"
    icmp = "DISABLED"
}
`

const testAccNsip_basic_step4 = `

resource "citrixadc_nsip" "tf_test_nsip" {
    ipaddress = "192.168.2.55"
    type = "VIP"
    netmask = "255.255.254.0"
    icmp = "DISABLED"
	state = "DISABLED"
}
`
