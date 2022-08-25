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
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"testing"
)

const testAccNtpparam_basic = `
	resource "citrixadc_ntpparam" "tf_ntpparam" {
		authentication = "YES"
		trustedkey     = [123, 456]
		autokeylogsec  = 15
		revokelogsec   = 20
	}
`
const testAccNtpparam_update = `
	resource "citrixadc_ntpparam" "tf_ntpparam" {
		authentication = "NO"
		trustedkey     = [1234, 4567]
		autokeylogsec  = 10
		revokelogsec   = 12
	}
`

func TestAccNtpparam_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNtpparam_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNtpparamExist("citrixadc_ntpparam.tf_ntpparam", nil),
					resource.TestCheckResourceAttr("citrixadc_ntpparam.tf_ntpparam", "authentication", "YES"),
					//resource.TestCheckResourceAttr("citrixadc_ntpparam.tf_ntpparam", "trustedkey", "[123, 456]"),
					resource.TestCheckResourceAttr("citrixadc_ntpparam.tf_ntpparam", "autokeylogsec", "15"),
					resource.TestCheckResourceAttr("citrixadc_ntpparam.tf_ntpparam", "revokelogsec", "20"),
				),
			},
			resource.TestStep{
				Config: testAccNtpparam_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNtpparamExist("citrixadc_ntpparam.tf_ntpparam", nil),
					resource.TestCheckResourceAttr("citrixadc_ntpparam.tf_ntpparam", "authentication", "NO"),
					//resource.TestCheckResourceAttr("citrixadc_ntpparam.tf_ntpparam", "trustedkey", "[1234, 4567]"),
					resource.TestCheckResourceAttr("citrixadc_ntpparam.tf_ntpparam", "autokeylogsec", "10"),
					resource.TestCheckResourceAttr("citrixadc_ntpparam.tf_ntpparam", "revokelogsec", "12"),
				),
			},
		},
	})
}

func testAccCheckNtpparamExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ntpparam name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(service.Ntpparam.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("ntpparam %s not found", n)
		}

		return nil
	}
}