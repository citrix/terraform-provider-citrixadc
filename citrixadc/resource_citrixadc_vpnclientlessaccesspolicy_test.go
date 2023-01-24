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
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"testing"
)

const testAccVpnclientlessaccesspolicy_basic = `

	resource "citrixadc_vpnclientlessaccesspolicy" "tf_vpnclientlessaccesspolicy" {
		name = "tf_vpnclientlessaccesspolicy"
		profilename = "ns_cvpn_default_profile"
		rule = "true"
	}
`

const testAccVpnclientlessaccesspolicy_basic_update = `

	resource "citrixadc_vpnclientlessaccesspolicy" "tf_vpnclientlessaccesspolicy" {
		name = "tf_vpnclientlessaccesspolicy"
		profilename = "ns_cvpn_v2_default_profile"
		rule = "false"
	}
`

func TestAccVpnclientlessaccesspolicy_basic(t *testing.T) {
	if adcTestbed != "STANDALONE" {
		t.Skipf("ADC testbed is %s. Expected STANDALONE.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckVpnclientlessaccesspolicyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccVpnclientlessaccesspolicy_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnclientlessaccesspolicyExist("citrixadc_vpnclientlessaccesspolicy.tf_vpnclientlessaccesspolicy", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnclientlessaccesspolicy.tf_vpnclientlessaccesspolicy", "name", "tf_vpnclientlessaccesspolicy"),
					resource.TestCheckResourceAttr("citrixadc_vpnclientlessaccesspolicy.tf_vpnclientlessaccesspolicy", "profilename", "ns_cvpn_default_profile"),
					resource.TestCheckResourceAttr("citrixadc_vpnclientlessaccesspolicy.tf_vpnclientlessaccesspolicy", "rule", "true"),
				),
			},
			resource.TestStep{
				Config: testAccVpnclientlessaccesspolicy_basic_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnclientlessaccesspolicyExist("citrixadc_vpnclientlessaccesspolicy.tf_vpnclientlessaccesspolicy", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnclientlessaccesspolicy.tf_vpnclientlessaccesspolicy", "profilename", "ns_cvpn_v2_default_profile"),
					resource.TestCheckResourceAttr("citrixadc_vpnclientlessaccesspolicy.tf_vpnclientlessaccesspolicy", "rule", "false"),
				),
			},
		},
	})
}

func testAccCheckVpnclientlessaccesspolicyExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No vpnclientlessaccesspolicy name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(service.Vpnclientlessaccesspolicy.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("vpnclientlessaccesspolicy %s not found", n)
		}

		return nil
	}
}

func testAccCheckVpnclientlessaccesspolicyDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_vpnclientlessaccesspolicy" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Vpnclientlessaccesspolicy.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("vpnclientlessaccesspolicy %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
