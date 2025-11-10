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
	"testing"
)

const testAccDnszone_add = `


resource "citrixadc_dnszone" "dnszone" {
	zonename      = "tf_zone1"
	proxymode     = "YES"
	dnssecoffload = "DISABLED"
	nsec          = "DISABLED"
	
	}
`

func TestAccDnszone_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckDnszoneDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDnszone_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDnszoneExist("citrixadc_dnszone.dnszone", nil),
					resource.TestCheckResourceAttr("citrixadc_dnszone.dnszone", "zonename", "tf_zone1"),
					resource.TestCheckResourceAttr("citrixadc_dnszone.dnszone", "proxymode", "YES"),
					resource.TestCheckResourceAttr("citrixadc_dnszone.dnszone", "dnssecoffload", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_dnszone.dnszone", "nsec", "DISABLED"),
				),
			},
		},
	})
}

func testAccCheckDnszoneExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No dnszone name is set")
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
		data, err := client.FindResource(service.Dnszone.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("dnszone %s not found", n)
		}

		return nil
	}
}

func testAccCheckDnszoneDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_dnszone" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Dnszone.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("dnszone %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
