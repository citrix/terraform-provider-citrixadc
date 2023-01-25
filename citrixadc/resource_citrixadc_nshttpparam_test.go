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

const testAccNshttpparam_add = `
	resource "citrixadc_nshttpparam" "tf_nshttpparam" {
		dropinvalreqs             = "ON"
		markconnreqinval          = "ON"
		maxreusepool              = 1
		markhttp09inval           = "ON"
		insnssrvrhdr              = "OFF"
		logerrresp                = "OFF"
		conmultiplex              = "DISABLED"
		http2serverside           = "OFF"
		ignoreconnectcodingscheme = "ENABLED"
	}
`
const testAccNshttpparam_update = `
	resource "citrixadc_nshttpparam" "tf_nshttpparam" {
		dropinvalreqs             = "OFF"
		markconnreqinval          = "OFF"
		maxreusepool              = 0
		markhttp09inval           = "OFF"
		insnssrvrhdr              = "ON"
		logerrresp                = "ON"
		conmultiplex              = "ENABLED"
		http2serverside           = "ON"
		ignoreconnectcodingscheme = "DISABLED"
	}
`

func TestAccNshttpparam_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNshttpparam_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNshttpparamExist("citrixadc_nshttpparam.tf_nshttpparam", nil),
					resource.TestCheckResourceAttr("citrixadc_nshttpparam.tf_nshttpparam", "dropinvalreqs", "ON"),
					resource.TestCheckResourceAttr("citrixadc_nshttpparam.tf_nshttpparam", "markconnreqinval", "ON"),
					resource.TestCheckResourceAttr("citrixadc_nshttpparam.tf_nshttpparam", "maxreusepool", "1"),
					resource.TestCheckResourceAttr("citrixadc_nshttpparam.tf_nshttpparam", "markhttp09inval", "ON"),
					resource.TestCheckResourceAttr("citrixadc_nshttpparam.tf_nshttpparam", "conmultiplex", "DISABLED"),
				),
			},
			resource.TestStep{
				Config: testAccNshttpparam_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNshttpparamExist("citrixadc_nshttpparam.tf_nshttpparam", nil),
					resource.TestCheckResourceAttr("citrixadc_nshttpparam.tf_nshttpparam", "dropinvalreqs", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_nshttpparam.tf_nshttpparam", "markconnreqinval", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_nshttpparam.tf_nshttpparam", "maxreusepool", "0"),
					resource.TestCheckResourceAttr("citrixadc_nshttpparam.tf_nshttpparam", "markhttp09inval", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_nshttpparam.tf_nshttpparam", "conmultiplex", "ENABLED"),
				),
			},
		},
	})
}

func testAccCheckNshttpparamExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No nshttpparam name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(service.Nshttpparam.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("nshttpparam %s not found", n)
		}

		return nil
	}
}
