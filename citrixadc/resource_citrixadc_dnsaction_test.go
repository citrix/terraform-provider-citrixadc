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

const testAccDnsaction_add = `


resource "citrixadc_dnsaction" "dnsaction" {
	actionname       = "tf_action1"
	actiontype       = "Rewrite_Response"
	ipaddress        = ["192.0.2.20","192.0.2.56","198.51.130.10"]
	dnsprofilename   = "tf_profile1"
  }
`

func TestAccDnsaction_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDnsactionDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccDnsaction_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDnsactionExist("citrixadc_dnsaction.dnsaction", nil),
					resource.TestCheckResourceAttr("citrixadc_dnsaction.dnsaction", "actionname", "tf_action1"),
					resource.TestCheckResourceAttr("citrixadc_dnsaction.dnsaction","actiontype", "Rewrite_Response"),
					resource.TestCheckResourceAttr("citrixadc_dnsaction.dnsaction","dnsprofilename", "tf_profile1"),
				),
			},
		},
	})
}

func testAccCheckDnsactionExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No dnsaction name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(service.Dnsaction.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("dnsaction %s not found", n)
		}

		return nil
	}
}

func testAccCheckDnsactionDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_dnsaction" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Dnsaction.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("dnsaction %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
