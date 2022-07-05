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

const testAccDnspolicy_add = `


resource "citrixadc_dnspolicy" "dnspolicy" {
	name = "policy1"
	rule = "dns.req.question.type.ne(aaaa)"
	drop = "YES"  
 }
`
const testAccDnspolicy_update = `


resource "citrixadc_dnspolicy" "dnspolicy" {
	name = "policy1"
	rule = "dns.req.question.type.ne(aaaa)"
    drop = "NO"
 }
`
func TestAccDnspolicy_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDnspolicyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccDnspolicy_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDnspolicyExist("citrixadc_dnspolicy.dnspolicy", nil),
					resource.TestCheckResourceAttr("citrixadc_dnspolicy.dnspolicy", "name", "policy1"),
					resource.TestCheckResourceAttr("citrixadc_dnspolicy.dnspolicy", "rule", "dns.req.question.type.ne(aaaa)"),
					resource.TestCheckResourceAttr("citrixadc_dnspolicy.dnspolicy", "drop", "YES"),

				),
			},
			resource.TestStep{
				Config: testAccDnspolicy_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDnspolicyExist("citrixadc_dnspolicy.dnspolicy", nil),
					resource.TestCheckResourceAttr("citrixadc_dnspolicy.dnspolicy", "name", "policy1"),
					resource.TestCheckResourceAttr("citrixadc_dnspolicy.dnspolicy", "rule", "dns.req.question.type.ne(aaaa)"),
					resource.TestCheckResourceAttr("citrixadc_dnspolicy.dnspolicy", "drop", "NO"),
				),
			},
		},
	})
}

func testAccCheckDnspolicyExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No dnspolicy name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(service.Dnspolicy.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("dnspolicy %s not found", n)
		}

		return nil
	}
}

func testAccCheckDnspolicyDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_dnspolicy" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Dnspolicy.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("dnspolicy %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
