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

const testAccAuthenticationpolicylabel_add = `

	resource "citrixadc_authenticationpolicylabel" "tf_authenticationpolicylabel" {
		labelname = "tf_authenticationpolicylabel"
		type 	  = "AAATM_REQ"
		comment   = "Testing"
	}
`
const testAccAuthenticationpolicylabel_update = `

	resource "citrixadc_authenticationpolicylabel" "tf_authenticationpolicylabel" {
		labelname = "tf_authenticationpolicylabel"
		type 	  = "RBA_REQ"
		comment   = "Testing1"
	}
`
func TestAccAuthenticationpolicylabel_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAuthenticationpolicylabelDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccAuthenticationpolicylabel_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationpolicylabelExist("citrixadc_authenticationpolicylabel.tf_authenticationpolicylabel", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationpolicylabel.tf_authenticationpolicylabel", "labelname", "tf_authenticationpolicylabel"),
					resource.TestCheckResourceAttr("citrixadc_authenticationpolicylabel.tf_authenticationpolicylabel", "type", "AAATM_REQ"),
					resource.TestCheckResourceAttr("citrixadc_authenticationpolicylabel.tf_authenticationpolicylabel", "comment", "Testing"),
				),
			},
			resource.TestStep{
				Config: testAccAuthenticationpolicylabel_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationpolicylabelExist("citrixadc_authenticationpolicylabel.tf_authenticationpolicylabel", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationpolicylabel.tf_authenticationpolicylabel", "labelname", "tf_authenticationpolicylabel"),
					resource.TestCheckResourceAttr("citrixadc_authenticationpolicylabel.tf_authenticationpolicylabel", "type", "RBA_REQ"),
					resource.TestCheckResourceAttr("citrixadc_authenticationpolicylabel.tf_authenticationpolicylabel", "comment", "Testing1"),
				),
			},
		},
	})
}

func testAccCheckAuthenticationpolicylabelExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No authenticationpolicylabel name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(service.Authenticationpolicylabel.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("authenticationpolicylabel %s not found", n)
		}

		return nil
	}
}

func testAccCheckAuthenticationpolicylabelDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_authenticationpolicylabel" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Authenticationpolicylabel.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("authenticationpolicylabel %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
