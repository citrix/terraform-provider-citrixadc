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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"net/url"
	"testing"
)

const testAccDnsaaaarec_basic = `


resource "citrixadc_dnsaaaarec" "dnsaaaarec" {
	hostname = "www.adfihrwpi.com"
    ipv6address = "2001:db8:85a3::8a2e:370:7334"
    ttl = 3600
}

`

func TestAccDnsaaaarec_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckDnsaaaarecDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDnsaaaarec_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDnsaaaarecExist("citrixadc_dnsaaaarec.dnsaaaarec", nil),
					resource.TestCheckResourceAttr("citrixadc_dnsaaaarec.dnsaaaarec", "hostname", "www.adfihrwpi.com"),
					resource.TestCheckResourceAttr("citrixadc_dnsaaaarec.dnsaaaarec", "ipv6address", "2001:db8:85a3::8a2e:370:7334"),
					resource.TestCheckResourceAttr("citrixadc_dnsaaaarec.dnsaaaarec", "ttl", "3600"),
				),
			},
		},
	})
}

func testAccCheckDnsaaaarecExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No dnsaaaarec name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		// Use the shared utility function to get a configured client
		client, err := testAccGetClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Dnsaaaarec.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("dnsaaaarec %s not found", n)
		}

		return nil
	}
}

func testAccCheckDnsaaaarecDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_dnsaaaarec" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		argsMap := make(map[string]string)
		argsMap["ecssubnet"] = url.QueryEscape(rs.Primary.Attributes["ecssubnet"])
		argsMap["ipv6address"] = url.QueryEscape(rs.Primary.Attributes["ipv6address"])
		findParams := service.FindParams{
			ResourceType: service.Dnsaaaarec.Type(),
			ArgsMap:      argsMap,
		}
		_, err := client.FindResourceArrayWithParams(findParams)

		if err == nil {
			return fmt.Errorf("appfwconfidfield %s still exists", rs.Primary.ID)
		}
	}
	return nil
}
