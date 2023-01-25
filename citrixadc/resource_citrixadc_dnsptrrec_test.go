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

const testAccDnsptrrec_basic = `

resource "citrixadc_dnsptrrec" "tf_dnsptrrec" {
	reversedomain = "0.2.0.192.in-addr.arpa"
	domain        = "example.com"
	ttl           = 3600
  }
`

func TestAccDnsptrrec_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDnsptrrecDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccDnsptrrec_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDnsptrrecExist("citrixadc_dnsptrrec.tf_dnsptrrec", nil),
					resource.TestCheckResourceAttr("citrixadc_dnsptrrec.tf_dnsptrrec", "reversedomain", "0.2.0.192.in-addr.arpa"),
					resource.TestCheckResourceAttr("citrixadc_dnsptrrec.tf_dnsptrrec", "domain", "example.com"),
					resource.TestCheckResourceAttr("citrixadc_dnsptrrec.tf_dnsptrrec", "ttl", "3600"),
				),
			},
		},
	})
}

func testAccCheckDnsptrrecExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No dnsptrrec name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(service.Dnsptrrec.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("dnsptrrec %s not found", n)
		}

		return nil
	}
}

func testAccCheckDnsptrrecDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_dnsptrrec" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Dnsptrrec.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("dnsptrrec %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
