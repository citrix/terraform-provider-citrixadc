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
	"testing"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testAccDnsaddrec_basic = `


resource "citrixadc_dnsaddrec" "dnsaddrec" {
	hostname  = "ab.root-servers.net"
	ipaddress = "65.200.211.129"
	ttl       = 3600
	}
`

func TestAccDnsaddrec_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDnsaddrecDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDnsaddrec_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDnsaddrecExist("citrixadc_dnsaddrec.dnsaddrec", nil),
					resource.TestCheckResourceAttr("citrixadc_dnsaddrec.dnsaddrec", "hostname", "ab.root-servers.net"),
					resource.TestCheckResourceAttr("citrixadc_dnsaddrec.dnsaddrec", "ipaddress", "65.200.211.129"),
					resource.TestCheckResourceAttr("citrixadc_dnsaddrec.dnsaddrec", "ttl", "3600"),
				),
			},
		},
	})
}

func testAccCheckDnsaddrecExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No dnsaddrec name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		dataArr, err := client.FindAllResources(service.Dnsaddrec.Type())

		if err != nil {
			return err
		}

		found := false
		for _, v := range dataArr {
			if v["hostname"].(string) == rs.Primary.Attributes["hostname"] && v["ipaddress"].(string) == rs.Primary.Attributes["ipaddress"] {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("Dnsaddrec %s not found", n)
		}

		return nil
	}
}

func testAccCheckDnsaddrecDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_dnsaddrec" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}
		dataArr, err := client.FindAllResources(service.Dnsaddrec.Type())

		if err != nil {
			return err
		}

		found := false
		for _, v := range dataArr {
			if v["hostname"].(string) == rs.Primary.Attributes["hostname"] && v["ipaddress"].(string) == rs.Primary.Attributes["ipaddress"] {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("Dnsaddrec %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
