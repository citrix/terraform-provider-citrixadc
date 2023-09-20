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

const testAccAaaradiusparams_basic = `


	resource "citrixadc_aaaradiusparams" "tf_aaaradiusparams" {
		radkey             = "sslvpn"
		radnasip           = "ENABLED"
		serverip           = "10.222.74.158"
		authtimeout        = 8
	}
`
const testAccAaaradiusparams_update = `


	resource "citrixadc_aaaradiusparams" "tf_aaaradiusparams" {
		radkey             = "sslvpn2"
		radnasip           = "DISABLED"
		serverip           = "10.222.74.159"
		authtimeout        = 10
	}
`

func TestAccAaaradiusparams_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				Config: testAccAaaradiusparams_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaaradiusparamsExist("citrixadc_aaaradiusparams.tf_aaaradiusparams", nil),
					resource.TestCheckResourceAttr("citrixadc_aaaradiusparams.tf_aaaradiusparams", "radkey", "sslvpn"),
					resource.TestCheckResourceAttr("citrixadc_aaaradiusparams.tf_aaaradiusparams", "radnasip", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_aaaradiusparams.tf_aaaradiusparams", "serverip", "10.222.74.158"),
					resource.TestCheckResourceAttr("citrixadc_aaaradiusparams.tf_aaaradiusparams", "authtimeout", "8"),
				),
			},
			{
				Config: testAccAaaradiusparams_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaaradiusparamsExist("citrixadc_aaaradiusparams.tf_aaaradiusparams", nil),
					resource.TestCheckResourceAttr("citrixadc_aaaradiusparams.tf_aaaradiusparams", "radkey", "sslvpn2"),
					resource.TestCheckResourceAttr("citrixadc_aaaradiusparams.tf_aaaradiusparams", "radnasip", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_aaaradiusparams.tf_aaaradiusparams", "serverip", "10.222.74.159"),
					resource.TestCheckResourceAttr("citrixadc_aaaradiusparams.tf_aaaradiusparams", "authtimeout", "10"),
				),
			},
		},
	})
}

func testAccCheckAaaradiusparamsExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No aaaradiusparams name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(service.Aaaradiusparams.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("aaaradiusparams %s not found", n)
		}

		return nil
	}
}
