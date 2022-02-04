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

const testAccVpnsessionaction_add = `

	resource "citrixadc_vpnsessionaction" "foo" {
		name 					   = "newsession"
		sesstimeout                = "10"
  		defaultauthorizationaction = "ALLOW"
  		transparentinterception    = "ON"
  		clientidletimeout          = "10"
  		sso                        = "ON"
  		icaproxy                   = "ON"
  		wihome                     = "https://citrix.lab.com"
  		clientlessvpnmode          = "DISABLED"
  		
	}
`
const testAccVpnsessionaction_update = `

	resource "citrixadc_vpnsessionaction" "foo" {
		name 					   = "newsession"
		sesstimeout                = "20"
	 	defaultauthorizationaction = "DENY"
	  	transparentinterception    = "ON"
		clientidletimeout          = "20"
		sso                        = "ON"
		icaproxy                   = "OFF"
		wihome                     = "https://citrix.lab.com"
		clientlessvpnmode          = "DISABLED"
		httpport                   = [8080, 8000, 808]	
	}
`

func TestAccVpnsessionaction_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckVpnsessionactionDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccVpnsessionaction_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnsessionactionExist("citrixadc_vpnsessionaction.foo", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnsessionaction.foo", "name", "newsession"),
					resource.TestCheckResourceAttr("citrixadc_vpnsessionaction.foo", "sesstimeout", "10"),
					resource.TestCheckResourceAttr("citrixadc_vpnsessionaction.foo", "defaultauthorizationaction", "ALLOW"),
				),
			},
			resource.TestStep{
				Config: testAccVpnsessionaction_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnsessionactionExist("citrixadc_vpnsessionaction.foo", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnsessionaction.foo", "name", "newsession"),
					resource.TestCheckResourceAttr("citrixadc_vpnsessionaction.foo", "sesstimeout", "20"),
					resource.TestCheckResourceAttr("citrixadc_vpnsessionaction.foo", "clientidletimeout", "20"),
				),
			},
		},
	})
}

func testAccCheckVpnsessionactionExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No vpnsessionaction name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(service.Vpnsessionaction.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("vpnsessionaction %s not found", n)
		}

		return nil
	}
}

func testAccCheckVpnsessionactionDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_vpnsessionaction" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Vpnsessionaction.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("vpnsessionaction %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
