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

const testAccDnssoarec_basic_step1 = `
resource "citrixadc_dnssoarec" "tf_dnssoarec" {
	domain =  "test.com"
	originserver  = "10.2.3.5"
	contact =  "other"
	expire = 1800
	refresh = 4000
}
`

const testAccDnssoarec_basic_step2 = `
resource "citrixadc_dnssoarec" "tf_dnssoarec" {
	domain =  "test.com"
	originserver  = "10.2.3.6"
	contact =  "some_other"
	expire = 1600
	refresh = 3600
}
`

func TestAccDnssoarec_basic(t *testing.T) {
	if adcTestbed != "STANDALONE" {
		t.Skipf("ADC testbed is %s. Expected STANDALONE.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDnssoarecDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDnssoarec_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDnssoarecExist("citrixadc_dnssoarec.tf_dnssoarec", nil),
				),
			},
			{
				Config: testAccDnssoarec_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDnssoarecExist("citrixadc_dnssoarec.tf_dnssoarec", nil),
				),
			},
		},
	})
}

func testAccCheckDnssoarecExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No dnssoarec name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(service.Dnssoarec.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("dnssoarec %s not found", n)
		}

		return nil
	}
}

func testAccCheckDnssoarecDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_dnssoarec" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Dnssoarec.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("dnssoarec %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
