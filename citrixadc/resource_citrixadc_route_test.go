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
	"log"
	"net/url"
	"testing"

	"github.com/chiradeep/go-nitro/netscaler"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const testAccRoute_add = `
	resource "citrixadc_route" "foo" {
		depends_on = [citrixadc_nsip.nsip]
		network    = "100.0.100.0"
		netmask    = "255.255.255.0"
		gateway    = "100.0.1.1"
		advertise  = "ENABLED"
	}
	
	resource "citrixadc_nsip" "nsip" {
		ipaddress = "100.0.1.100"
		netmask   = "255.255.255.0"
	}
`

const testAccRoute_update = `
	resource "citrixadc_route" "foo" {
		depends_on = [citrixadc_nsip.nsip]
		network    = "100.0.100.0"
		netmask    = "255.255.255.0"
		gateway    = "100.0.1.1"
		advertise  = "DISABLED"
	}
	
	resource "citrixadc_nsip" "nsip" {
		ipaddress = "100.0.1.100"
		netmask   = "255.255.255.0"
	}
`

func TestAccRoute_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckRouteDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccRoute_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRouteExist("citrixadc_route.foo", nil),
					resource.TestCheckResourceAttr("citrixadc_route.foo", "network", "100.0.100.0"),
					resource.TestCheckResourceAttr("citrixadc_route.foo", "netmask", "255.255.255.0"),
					resource.TestCheckResourceAttr("citrixadc_route.foo", "gateway", "100.0.1.1"),
					resource.TestCheckResourceAttr("citrixadc_route.foo", "advertise", "ENABLED"),
				),
			},
			resource.TestStep{
				Config: testAccRoute_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRouteExist("citrixadc_route.foo", nil),
					resource.TestCheckResourceAttr("citrixadc_route.foo", "network", "100.0.100.0"),
					resource.TestCheckResourceAttr("citrixadc_route.foo", "netmask", "255.255.255.0"),
					resource.TestCheckResourceAttr("citrixadc_route.foo", "gateway", "100.0.1.1"),
					resource.TestCheckResourceAttr("citrixadc_route.foo", "advertise", "DISABLED"),
				),
			},
		},
	})
}

func testAccCheckRouteExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No route is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}
		routeName := rs.Primary.ID
		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		argsMap := make(map[string]string)
		argsMap["network"] = url.QueryEscape(rs.Primary.Attributes["network"])
		argsMap["netmask"] = url.QueryEscape(rs.Primary.Attributes["netmask"])
		argsMap["gateway"] = url.QueryEscape(rs.Primary.Attributes["gateway"])
		findParams := netscaler.FindParams{
			ResourceType: netscaler.Route.Type(),
			ArgsMap:      argsMap,
		}
		dataArray, err := nsClient.FindResourceArrayWithParams(findParams)
		if err != nil {
			log.Printf("[WARN] citrix-provider: acceptance test: Clearing route state %s", routeName)
			return nil
		}
		if len(dataArray) == 0 {
			log.Printf("[WARN] citrix-provider: acceptance test: route does not exist. Clearing state.")
			return nil
		}

		if len(dataArray) > 1 {
			return fmt.Errorf("[ERROR] citrix-provider: acceptance test: multiple entries found for route")
		}

		return nil
	}
}

func testAccCheckRouteDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_route" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}
		argsMap := make(map[string]string)
		argsMap["network"] = url.QueryEscape(rs.Primary.Attributes["network"])
		argsMap["netmask"] = url.QueryEscape(rs.Primary.Attributes["netmask"])
		argsMap["gateway"] = url.QueryEscape(rs.Primary.Attributes["gateway"])
		findParams := netscaler.FindParams{
			ResourceType: netscaler.Route.Type(),
			ArgsMap:      argsMap,
		}
		_, err := nsClient.FindResourceArrayWithParams(findParams)

		if err == nil {
			return fmt.Errorf("Route %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
