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

const testAccForwardingsession_add = `
	resource "citrixadc_forwardingsession" "tf_forwarding" {
		name             = "tf_forwarding"
		network          = "10.102.105.90"
		netmask          = "255.255.255.255"
		connfailover     = "ENABLED"
		sourceroutecache = "ENABLED"
		processlocal     = "DISABLED"
	}
`
const testAccForwardingsession_update = `
	resource "citrixadc_forwardingsession" "tf_forwarding" {
		name             = "tf_forwarding"
		network          = "10.102.105.90"
		netmask          = "255.255.255.255"
		connfailover     = "DISABLED"
		sourceroutecache = "DISABLED"
		processlocal     = "DISABLED"
	}
`

func TestAccForwardingsession_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckForwardingsessionDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccForwardingsession_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckForwardingsessionExist("citrixadc_forwardingsession.tf_forwarding", nil),
					resource.TestCheckResourceAttr("citrixadc_forwardingsession.tf_forwarding", "name", "tf_forwarding"),
					resource.TestCheckResourceAttr("citrixadc_forwardingsession.tf_forwarding", "network", "10.102.105.90"),
					resource.TestCheckResourceAttr("citrixadc_forwardingsession.tf_forwarding", "connfailover", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_forwardingsession.tf_forwarding", "sourceroutecache", "ENABLED"),
				),
			},
			resource.TestStep{
				Config: testAccForwardingsession_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckForwardingsessionExist("citrixadc_forwardingsession.tf_forwarding", nil),
					resource.TestCheckResourceAttr("citrixadc_forwardingsession.tf_forwarding", "name", "tf_forwarding"),
					resource.TestCheckResourceAttr("citrixadc_forwardingsession.tf_forwarding", "network", "10.102.105.90"),
					resource.TestCheckResourceAttr("citrixadc_forwardingsession.tf_forwarding", "connfailover", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_forwardingsession.tf_forwarding", "sourceroutecache", "DISABLED"),
				),
			},
		},
	})
}

func testAccCheckForwardingsessionExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No forwardingsession name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(service.Forwardingsession.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("forwardingsession %s not found", n)
		}

		return nil
	}
}

func testAccCheckForwardingsessionDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_forwardingsession" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Forwardingsession.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("forwardingsession %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
