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

const testAccDnspolicy64_add = `

resource "citrixadc_dnspolicy64" "dnspolicy64" {
	name  = "policy_1"
	rule = "dns.req.question.type.ne(aaaa)"
	action = "default_DNS64_action"
  }
`

const testAccDnspolicy64_update = `

resource "citrixadc_dnspolicy64" "dnspolicy64" {
	name  = "policy_1"
	rule = "client.ip.src.in_subnet(23.43.0.0/16)"
	action = "default_DNS64_action"
  }
`
func TestAccDnspolicy64_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDnspolicy64Destroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccDnspolicy64_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDnspolicy64Exist("citrixadc_dnspolicy64.dnspolicy64", nil),
					resource.TestCheckResourceAttr("citrixadc_dnspolicy64.dnspolicy64", "name", "policy_1"),
					resource.TestCheckResourceAttr("citrixadc_dnspolicy64.dnspolicy64", "rule", "dns.req.question.type.ne(aaaa)"),
					resource.TestCheckResourceAttr("citrixadc_dnspolicy64.dnspolicy64", "action", "default_DNS64_action"),
				),
			},
			resource.TestStep{
				Config: testAccDnspolicy64_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDnspolicy64Exist("citrixadc_dnspolicy64.dnspolicy64", nil),
					resource.TestCheckResourceAttr("citrixadc_dnspolicy64.dnspolicy64", "name", "policy_1"),
					resource.TestCheckResourceAttr("citrixadc_dnspolicy64.dnspolicy64", "rule", "client.ip.src.in_subnet(23.43.0.0/16)"),
					resource.TestCheckResourceAttr("citrixadc_dnspolicy64.dnspolicy64", "action", "default_DNS64_action"),
				),
			},
		},
	})
}

func testAccCheckDnspolicy64Exist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No dnspolicy64 name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(service.Dnspolicy64.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("dnspolicy64 %s not found", n)
		}

		return nil
	}
}

func testAccCheckDnspolicy64Destroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_dnspolicy64" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Dnspolicy64.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("dnspolicy64 %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
