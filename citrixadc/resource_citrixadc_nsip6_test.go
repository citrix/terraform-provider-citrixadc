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
	"errors"
	"fmt"
	"testing"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccNsip6_basic(t *testing.T) {
	if adcTestbed != "STANDALONE" {
		t.Skipf("ADC testbed is %s. Expected STANDALONE.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNsip6Destroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNsip6_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsip6Exist("citrixadc_nsip6.tf_nsip6", nil, "2002:db8:100::ff/64"),
				),
			},
			resource.TestStep{
				Config: testAccNsip6_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsip6Exist("citrixadc_nsip6.tf_nsip6", nil, "2002:db8:100::ff/64"),
				),
			},
			resource.TestStep{
				Config: testAccNsip6_basic_step3,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsip6Exist("citrixadc_nsip6.tf_nsip6", nil, "2002:db8:100::ff/64"),
				),
			},
		},
	})
}

const testAccNsip6_mptcpadvertise = `
	resource "citrixadc_nsip6" "tf_test_nsip6_mptcpadvertise" {
		ipv6address = "2002:db8:100::ff/64"
		type = "VIP"
		icmp = "ENABLED"
		mptcpadvertise = "YES"
	}
`

func TestAccNsip6_mptcpadvertise(t *testing.T) {
	if adcTestbed != "STANDALONE" {
		t.Skipf("ADC testbed is %s. Expected STANDALONE.", adcTestbed)
	}
	if isCpxRun {
		t.Skip("No support in CPX")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNsip6Destroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNsip6_mptcpadvertise,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsip6Exist("citrixadc_nsip6.tf_test_nsip6_mptcpadvertise", nil, "2002:db8:100::ff/64"),
					resource.TestCheckResourceAttr("citrixadc_nsip6.tf_test_nsip6_mptcpadvertise", "mptcpadvertise", "YES"),
				),
			},
		},
	})
}

func testAccCheckNsip6Exist(n string, id *string, ipv6address string) resource.TestCheckFunc {
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

		array, _ := nsClient.FindAllResources(service.Nsip6.Type())

		foundAddress := false
		for _, item := range array {
			if item["ipv6address"] == ipv6address {
				foundAddress = true
				break
			}
		}
		if !foundAddress {
			return errors.New("Cannot find resource nsip6 with ipv6address %v")
		}

		return nil
	}
}

func testAccCheckNsip6Destroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_nsip6" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Nsip6.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("LB vserver %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccNsip6_basic_step1 = `

resource "citrixadc_nsip6" "tf_nsip6" {
    ipv6address = "2002:db8:100::ff/64"
    type = "VIP"
    icmp = "DISABLED"
}
`

const testAccNsip6_basic_step2 = `

resource "citrixadc_nsip6" "tf_nsip6" {
    ipv6address = "2002:db8:100::ff/64"
    type = "VIP"
    icmp = "ENABLED"
}
`

const testAccNsip6_basic_step3 = `

resource "citrixadc_nsip6" "tf_nsip6" {
    ipv6address = "2002:db8:100::ff/64"
    type = "SNIP"
    icmp = "ENABLED"
}
`
