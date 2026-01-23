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

const testAccDnskey_add = `

resource "citrixadc_dnskey" "dnskey" {
	keyname            = "adckey_1"
	publickey          = "/nsconfig/dns/demo.key"
	privatekey         = "/nsconfig/dns/demo.private"
	expires            = 120
	units1             = "DAYS"
	notificationperiod = 7
	units2             = "DAYS"
	ttl                = 3600
	rollovermethod	 = "PrePublication"
	autorollover	 = "ENABLED"
	}
`
const testAccDnskey_update = `

resource "citrixadc_dnskey" "dnskey" {
	keyname            = "adckey_1"
	publickey          = "/nsconfig/dns/demo.key"
	privatekey         = "/nsconfig/dns/demo.private"
	expires            = 121
	units1             = "HOURS"
	notificationperiod = 12
	units2             = "HOURS"
	ttl                = 3601
	rollovermethod	 = "DoubleSignature"
	autorollover	 = "DISABLED"
	}
`

func TestAccDnskey_basic(t *testing.T) {
	t.Skip("TODO: Need to find a way to test this resource!")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDnskeyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDnskey_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDnskeyExist("citrixadc_dnskey.dnskey", nil),
					resource.TestCheckResourceAttr("citrixadc_dnskey.dnskey", "keyname", "adckey_1"),
					resource.TestCheckResourceAttr("citrixadc_dnskey.dnskey", "expires", "120"),
					resource.TestCheckResourceAttr("citrixadc_dnskey.dnskey", "units1", "DAYS"),
					resource.TestCheckResourceAttr("citrixadc_dnskey.dnskey", "notificationperiod", "7"),
					resource.TestCheckResourceAttr("citrixadc_dnskey.dnskey", "units2", "DAYS"),
					resource.TestCheckResourceAttr("citrixadc_dnskey.dnskey", "ttl", "3600"),
					resource.TestCheckResourceAttr("citrixadc_dnskey.dnskey", "rollovermethod", "PrePublication"),
					resource.TestCheckResourceAttr("citrixadc_dnskey.dnskey", "autorollover", "ENABLED"),
				),
			},
			{
				Config: testAccDnskey_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDnskeyExist("citrixadc_dnskey.dnskey", nil),
					resource.TestCheckResourceAttr("citrixadc_dnskey.dnskey", "keyname", "adckey_1"),
					resource.TestCheckResourceAttr("citrixadc_dnskey.dnskey", "expires", "121"),
					resource.TestCheckResourceAttr("citrixadc_dnskey.dnskey", "units1", "HOURS"),
					resource.TestCheckResourceAttr("citrixadc_dnskey.dnskey", "notificationperiod", "12"),
					resource.TestCheckResourceAttr("citrixadc_dnskey.dnskey", "units2", "HOURS"),
					resource.TestCheckResourceAttr("citrixadc_dnskey.dnskey", "ttl", "3601"),
					resource.TestCheckResourceAttr("citrixadc_dnskey.dnskey", "rollovermethod", "DoubleSignature"),
					resource.TestCheckResourceAttr("citrixadc_dnskey.dnskey", "autorollover", "DISABLED"),
				),
			},
		},
	})
}

func testAccCheckDnskeyExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No dnskey name is set")
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
		data, err := client.FindResource(service.Dnskey.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("dnskey %s not found", n)
		}

		return nil
	}
}

func testAccCheckDnskeyDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_dnskey" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Dnskey.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("dnskey %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
