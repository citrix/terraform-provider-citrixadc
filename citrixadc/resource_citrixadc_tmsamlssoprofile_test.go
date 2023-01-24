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

const testAccTmsamlssoprofile_basic = `

	resource "citrixadc_tmsamlssoprofile" "tf_tmsamlssoprofile" {
		name                        = "my_tmsamlssoprofile"
		assertionconsumerserviceurl = "https://service.example.com"
		sendpassword                = "OFF"
		relaystaterule              = "true"
	}
  
`
const testAccTmsamlssoprofile_update = `

	resource "citrixadc_tmsamlssoprofile" "tf_tmsamlssoprofile" {
		name                        = "my_tmsamlssoprofile"
		assertionconsumerserviceurl = "https://service.example2.com"
		sendpassword                = "ON"
		relaystaterule              = "false"
	}
  
`

func TestAccTmsamlssoprofile_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckTmsamlssoprofileDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccTmsamlssoprofile_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTmsamlssoprofileExist("citrixadc_tmsamlssoprofile.tf_tmsamlssoprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_tmsamlssoprofile.tf_tmsamlssoprofile", "name", "my_tmsamlssoprofile"),
					resource.TestCheckResourceAttr("citrixadc_tmsamlssoprofile.tf_tmsamlssoprofile", "assertionconsumerserviceurl", "https://service.example.com"),
					resource.TestCheckResourceAttr("citrixadc_tmsamlssoprofile.tf_tmsamlssoprofile", "sendpassword", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_tmsamlssoprofile.tf_tmsamlssoprofile", "relaystaterule", "true"),
				),
			},
			resource.TestStep{
				Config: testAccTmsamlssoprofile_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTmsamlssoprofileExist("citrixadc_tmsamlssoprofile.tf_tmsamlssoprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_tmsamlssoprofile.tf_tmsamlssoprofile", "name", "my_tmsamlssoprofile"),
					resource.TestCheckResourceAttr("citrixadc_tmsamlssoprofile.tf_tmsamlssoprofile", "assertionconsumerserviceurl", "https://service.example2.com"),
					resource.TestCheckResourceAttr("citrixadc_tmsamlssoprofile.tf_tmsamlssoprofile", "sendpassword", "ON"),
					resource.TestCheckResourceAttr("citrixadc_tmsamlssoprofile.tf_tmsamlssoprofile", "relaystaterule", "false"),
				),
			},
		},
	})
}

func testAccCheckTmsamlssoprofileExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No tmsamlssoprofile name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(service.Tmsamlssoprofile.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("tmsamlssoprofile %s not found", n)
		}

		return nil
	}
}

func testAccCheckTmsamlssoprofileDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_tmsamlssoprofile" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Tmsamlssoprofile.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("tmsamlssoprofile %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
