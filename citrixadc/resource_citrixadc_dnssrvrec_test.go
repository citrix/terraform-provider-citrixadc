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
	"log"
	"net/url"
	"testing"
)

const testAccDnssrvrec_add = `


resource "citrixadc_dnssrvrec" "dnssrvrec" {
	domain   = "example.com"
	target   = "_sip._udp.example.com"
	priority = 1
	weight   = 1
	port     = 22
	ttl      = 3600
  }
  
`

const testAccDnssrvrec_update = `


resource "citrixadc_dnssrvrec" "dnssrvrec" {
	domain   = "example.com"
	target   = "_sip._udp.example.com"
	priority = 2
	weight   = 4
	port     = 21
	ttl      = 3604
  }
  
`

func TestAccDnssrvrec_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDnssrvrecDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccDnssrvrec_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDnssrvrecExist("citrixadc_dnssrvrec.dnssrvrec", nil),
					resource.TestCheckResourceAttr("citrixadc_dnssrvrec.dnssrvrec", "domain", "example.com"),
					resource.TestCheckResourceAttr("citrixadc_dnssrvrec.dnssrvrec", "target", "_sip._udp.example.com"),
					resource.TestCheckResourceAttr("citrixadc_dnssrvrec.dnssrvrec", "priority", "1"),
					resource.TestCheckResourceAttr("citrixadc_dnssrvrec.dnssrvrec", "weight", "1"),
					resource.TestCheckResourceAttr("citrixadc_dnssrvrec.dnssrvrec", "port", "22"),
					resource.TestCheckResourceAttr("citrixadc_dnssrvrec.dnssrvrec", "ttl", "3600"),
				),
			},

			resource.TestStep{
				Config: testAccDnssrvrec_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDnssrvrecExist("citrixadc_dnssrvrec.dnssrvrec", nil),
					resource.TestCheckResourceAttr("citrixadc_dnssrvrec.dnssrvrec", "domain", "example.com"),
					resource.TestCheckResourceAttr("citrixadc_dnssrvrec.dnssrvrec", "target", "_sip._udp.example.com"),
					resource.TestCheckResourceAttr("citrixadc_dnssrvrec.dnssrvrec", "priority", "2"),
					resource.TestCheckResourceAttr("citrixadc_dnssrvrec.dnssrvrec", "weight", "4"),
					resource.TestCheckResourceAttr("citrixadc_dnssrvrec.dnssrvrec", "port", "21"),
					resource.TestCheckResourceAttr("citrixadc_dnssrvrec.dnssrvrec", "ttl", "3604"),
				),
			},
		},
	})
}

func testAccCheckDnssrvrecExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No dnssrvrec name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		dnssrvrecName := rs.Primary.ID
		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		argsMap := make(map[string]string)
		argsMap["target"] = url.QueryEscape(rs.Primary.Attributes["target"])
		argsMap["ecssubnet"] = url.QueryEscape(rs.Primary.Attributes["ecssubnet"])
		findParams := service.FindParams{
			ResourceType: service.Dnssrvrec.Type(),
			ArgsMap:      argsMap,
		}
		dataArray, err := nsClient.FindResourceArrayWithParams(findParams)

		if err != nil {
			log.Printf("[WARN] citrix-provider: acceptance test: Clearing lb route state %s", dnssrvrecName)
			return nil
		}
		if len(dataArray) == 0 {
			log.Printf("[WARN] citrix-provider: acceptance test: Dnssrvrec does not exist. Clearing state.")
			return nil
		}

		if len(dataArray) > 1 {
			return fmt.Errorf("[ERROR] citrix-provider: acceptance test: multiple entries found for Dnssrvrec")
		}

		return nil
	}
}

func testAccCheckDnssrvrecDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_dnssrvrec" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		argsMap := make(map[string]string)
		argsMap["target"] = url.QueryEscape(rs.Primary.Attributes["target"])
		argsMap["ecssubnet"] = url.QueryEscape(rs.Primary.Attributes["ecssubnet"])
		findParams := service.FindParams{
			ResourceType: service.Dnssrvrec.Type(),
			ArgsMap:      argsMap,
		}
		_, err := nsClient.FindResourceArrayWithParams(findParams)

		if err == nil {
			return fmt.Errorf("dnssrvrec %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
