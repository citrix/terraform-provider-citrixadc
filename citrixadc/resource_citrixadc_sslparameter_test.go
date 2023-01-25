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

const testAccSslparameter_basic = `
	resource "citrixadc_sslparameter" "default" {
		denysslreneg   = "NONSECURE"
		defaultprofile = "ENABLED"
	}
`
const testAccSslparameter_basic_update = `
	resource "citrixadc_sslparameter" "default" {
		denysslreneg   = "ALL"
		defaultprofile = "ENABLED"
	}
`

func TestAccSslparameter_basic(t *testing.T) {
	if adcTestbed != "STANDALONE" {
		t.Skipf("ADC testbed is %s. Expected STANDALONE.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		// sslparameter resource do not have DELETE operation
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccSslparameter_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslparameterExist("citrixadc_sslparameter.default", nil),
					resource.TestCheckResourceAttr("citrixadc_sslparameter.default", "denysslreneg", "NONSECURE"),
					resource.TestCheckResourceAttr("citrixadc_sslparameter.default", "defaultprofile", "ENABLED"),
				),
			},
			resource.TestStep{
				Config: testAccSslparameter_basic_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslparameterExist("citrixadc_sslparameter.default", nil),
					resource.TestCheckResourceAttr("citrixadc_sslparameter.default", "denysslreneg", "ALL"),
					resource.TestCheckResourceAttr("citrixadc_sslparameter.default", "defaultprofile", "ENABLED"),
				),
			},
		},
	})
}

func testAccCheckSslparameterExist(n string, id *string) resource.TestCheckFunc {
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
		data, err := nsClient.FindResource(service.Sslparameter.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("SSL Parameter %s not found", n)
		}

		return nil
	}
}
