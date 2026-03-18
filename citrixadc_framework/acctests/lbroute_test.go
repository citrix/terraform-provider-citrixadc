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

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testAccLbroute_basic = `
	resource "citrixadc_lbroute" "tf_lbroute" {
		network = "55.0.0.0"
		netmask = "255.0.0.0"
		gatewayname = citrixadc_lbvserver.tf_lbvserver.name

		depends_on = [citrixadc_lbvserver_service_binding.tf_lbvserver_service_binding, citrixadc_nsip.nsip]
	}

	resource "citrixadc_nsip" "nsip" {
		ipaddress = "22.2.2.1"
		netmask   = "255.255.255.0"
	}

	resource "citrixadc_lbvserver_service_binding" "tf_lbvserver_service_binding" {
		name = citrixadc_lbvserver.tf_lbvserver.name
		servicename = citrixadc_service.tf_service.name
	}

	resource "citrixadc_service" "tf_service" {
		name = "tf_service"
		port = 65535
		ip = "22.2.2.2"
		servicetype = "ANY"
	}
	
	resource "citrixadc_lbvserver" "tf_lbvserver" {
		name = "tf_lbvserver"
		ipv46 = "0.0.0.0"
		servicetype = "ANY"
		lbmethod = "ROUNDROBIN"
		persistencetype = "NONE"
		clttimeout = 120
		port = 0
	}
`

const testAccLbrouteDataSource_basic = `
	resource "citrixadc_lbroute" "tf_lbroute" {
		network = "55.0.0.0"
		netmask = "255.0.0.0"
		gatewayname = citrixadc_lbvserver.tf_lbvserver.name

		depends_on = [citrixadc_lbvserver_service_binding.tf_lbvserver_service_binding, citrixadc_nsip.nsip]
	}

	resource "citrixadc_nsip" "nsip" {
		ipaddress = "22.2.2.1"
		netmask   = "255.255.255.0"
	}

	resource "citrixadc_lbvserver_service_binding" "tf_lbvserver_service_binding" {
		name = citrixadc_lbvserver.tf_lbvserver.name
		servicename = citrixadc_service.tf_service.name
	}

	resource "citrixadc_service" "tf_service" {
		name = "tf_service"
		port = 65535
		ip = "22.2.2.2"
		servicetype = "ANY"
	}
	
	resource "citrixadc_lbvserver" "tf_lbvserver" {
		name = "tf_lbvserver"
		ipv46 = "0.0.0.0"
		servicetype = "ANY"
		lbmethod = "ROUNDROBIN"
		persistencetype = "NONE"
		clttimeout = 120
		port = 0
	}

	data "citrixadc_lbroute" "tf_lbroute" {
		network = citrixadc_lbroute.tf_lbroute.network
		netmask = citrixadc_lbroute.tf_lbroute.netmask
		td      = citrixadc_lbroute.tf_lbroute.td
	}
`

func TestAccLbroute_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLbrouteDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLbroute_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbrouteExist("citrixadc_lbroute.tf_lbroute", nil),
				),
			},
		},
	})
}

func testAccCheckLbrouteExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lbroute name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}
		lbrouteName := rs.Primary.ID
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		argsMap := make(map[string]string)
		argsMap["network"] = url.QueryEscape(rs.Primary.Attributes["network"])
		argsMap["netmask"] = url.QueryEscape(rs.Primary.Attributes["netmask"])
		findParams := service.FindParams{
			ResourceType: service.Lbroute.Type(),
			ArgsMap:      argsMap,
		}
		dataArray, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			log.Printf("[WARN] citrix-provider: acceptance test: Clearing lb route state %s", lbrouteName)
			return nil
		}
		if len(dataArray) == 0 {
			log.Printf("[WARN] citrix-provider: acceptance test: lb route does not exist. Clearing state.")
			return nil
		}

		if len(dataArray) > 1 {
			return fmt.Errorf("[ERROR] citrix-provider: acceptance test: multiple entries found for lb route")
		}

		return nil
	}
}

func testAccCheckLbrouteDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_lbroute" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}
		argsMap := make(map[string]string)
		argsMap["network"] = url.QueryEscape(rs.Primary.Attributes["network"])
		argsMap["netmask"] = url.QueryEscape(rs.Primary.Attributes["netmask"])
		findParams := service.FindParams{
			ResourceType: service.Lbroute.Type(),
			ArgsMap:      argsMap,
		}
		_, err := client.FindResourceArrayWithParams(findParams)

		if err == nil {
			return fmt.Errorf("Lbroute %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccLbrouteDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccLbrouteDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_lbroute.tf_lbroute", "network", "55.0.0.0"),
					resource.TestCheckResourceAttr("data.citrixadc_lbroute.tf_lbroute", "netmask", "255.0.0.0"),
					resource.TestCheckResourceAttr("data.citrixadc_lbroute.tf_lbroute", "td", "0"),
					resource.TestCheckResourceAttrSet("data.citrixadc_lbroute.tf_lbroute", "gatewayname"),
				),
			},
		},
	})
}
