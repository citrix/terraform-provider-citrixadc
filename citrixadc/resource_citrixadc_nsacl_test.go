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

func TestAccNsacl_basic(t *testing.T) {
	if adcTestbed != "STANDALONE" {
		t.Skipf("ADC testbed is %s. Expected STANDALONE.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNsaclDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNsacl_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsaclExist("citrixadc_nsacl.foo", nil),

					resource.TestCheckResourceAttr(
						"citrixadc_nsacl.foo", "aclaction", "DENY"),
					resource.TestCheckResourceAttr(
						"citrixadc_nsacl.foo", "aclname", "test_acl"),
					resource.TestCheckResourceAttr(
						"citrixadc_nsacl.foo", "destipval", "192.168.1.33"),
					resource.TestCheckResourceAttr(
						"citrixadc_nsacl.foo", "protocol", "TCP"),
					resource.TestCheckResourceAttr(
						"citrixadc_nsacl.foo", "srcportval", "45-1024"),
				),
			},
		},
	})
}

func testAccCheckNsaclExist(n string, id *string) resource.TestCheckFunc {
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
		data, err := nsClient.FindResource(service.Nsacl.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("LB vserver %s not found", n)
		}

		return nil
	}
}

func testAccCheckNsaclDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_nsacl" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Nsacl.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("LB vserver %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccNsacl_basic = `


resource "citrixadc_nsacl" "foo" {

  aclaction = "DENY"
  aclname = "test_acl"
  destipval = "192.168.1.33"
  protocol = "TCP"
  srcportval = "45-1024"
  priority = "100"

}
`

const testAccNsaclEnableDisable_enabled = `
resource "citrixadc_nsacl" "tf_test_acc_nsacl" {
    aclname = "tf_test_acc_nsacl"
    aclaction = "ALLOW"
    priority = "100"
    srcipval = "192.168.10.22"
    destipval = "172.17.0.20"
    state = "ENABLED"
}
`

const testAccNsaclEnableDisable_disabled = `
resource "citrixadc_nsacl" "tf_test_acc_nsacl" {
    aclname = "tf_test_acc_nsacl"
    aclaction = "ALLOW"
    priority = "99"
    srcipval = "192.168.10.22"
    destipval = "172.17.0.20"
    state = "DISABLED"
}
`

func TestAccNsacl_enable_disable(t *testing.T) {
	if adcTestbed != "STANDALONE" {
		t.Skipf("ADC testbed is %s. Expected STANDALONE.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNsaclDestroy,
		Steps: []resource.TestStep{
			// Create enabled
			{
				Config: testAccNsaclEnableDisable_enabled,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsaclExist("citrixadc_nsacl.tf_test_acc_nsacl", nil),
					resource.TestCheckResourceAttr("citrixadc_nsacl.tf_test_acc_nsacl", "state", "ENABLED"),
				),
			},
			// Disable
			{
				Config: testAccNsaclEnableDisable_disabled,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsaclExist("citrixadc_nsacl.tf_test_acc_nsacl", nil),
					resource.TestCheckResourceAttr("citrixadc_nsacl.tf_test_acc_nsacl", "state", "DISABLED"),
				),
			},
			// Re enable
			{
				Config: testAccNsaclEnableDisable_enabled,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsaclExist("citrixadc_nsacl.tf_test_acc_nsacl", nil),
					resource.TestCheckResourceAttr("citrixadc_nsacl.tf_test_acc_nsacl", "state", "ENABLED"),
				),
			},
		},
	})
}
