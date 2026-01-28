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
	"strings"
	"testing"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testAccDnsnameserver_add = `

	resource "citrixadc_dnsprofile" "dnsprofile" {
		dnsprofilename         = "tf_profile1"
		dnsquerylogging        = "DISABLED"
		dnsanswerseclogging    = "DISABLED"
		dnsextendedlogging     = "DISABLED"
		dnserrorlogging        = "DISABLED"
		cacherecords           = "ENABLED"
		cachenegativeresponses = "ENABLED"
		dropmultiqueryrequest  = "DISABLED"
		cacheecsresponses      = "DISABLED"
	}

	resource "citrixadc_dnsnameserver" "dnsnameserver" {
		ip 				= "192.0.2.0"
		local 			= true
		state 			= "DISABLED"
		type 			= "UDP"
		dnsprofilename 	= citrixadc_dnsprofile.dnsprofile.dnsprofilename
	}
`
const testAccDnsnameserver_update = `

	resource "citrixadc_dnsprofile" "dnsprofile" {
		dnsprofilename         = "tf_profile1"
		dnsquerylogging        = "DISABLED"
		dnsanswerseclogging    = "DISABLED"
		dnsextendedlogging     = "DISABLED"
		dnserrorlogging        = "DISABLED"
		cacherecords           = "ENABLED"
		cachenegativeresponses = "ENABLED"
		dropmultiqueryrequest  = "DISABLED"
		cacheecsresponses      = "DISABLED"
	}
	resource "citrixadc_dnsnameserver" "dnsnameserver" {
		ip 				= "192.0.2.0"
		local 			= false
		state 			= "DISABLED"
		type 			= "UDP"
		dnsprofilename 	= citrixadc_dnsprofile.dnsprofile.dnsprofilename
	}
`

func TestAccDnsnameserver_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckDnsnameserverDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDnsnameserver_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDnsnameserverExist("citrixadc_dnsnameserver.dnsnameserver", nil),
					resource.TestCheckResourceAttr("citrixadc_dnsnameserver.dnsnameserver", "local", "true"),
					resource.TestCheckResourceAttr("citrixadc_dnsnameserver.dnsnameserver", "dnsprofilename", "tf_profile1"),
				),
			},
			{
				Config: testAccDnsnameserver_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDnsnameserverExist("citrixadc_dnsnameserver.dnsnameserver", nil),
					resource.TestCheckResourceAttr("citrixadc_dnsnameserver.dnsnameserver", "local", "false"),
					resource.TestCheckResourceAttr("citrixadc_dnsnameserver.dnsnameserver", "dnsprofilename", "tf_profile1"),
				),
			},
		},
	})
}

func testAccCheckDnsnameserverExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No resource name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		PrimaryId := rs.Primary.ID
		idSlice := strings.SplitN(PrimaryId, ",", 2)
		name := idSlice[0]
		dns_type := idSlice[1]

		// Use the shared utility function to get a configured client
		client, err := testAccGetClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		dataArr, err := client.FindAllResources(service.Dnsnameserver.Type())

		if err != nil {
			return err
		}

		// For UDP_TCP type, verify that both TCP and UDP resources exist
		if dns_type == "UDP_TCP" {
			foundTCP := false
			foundUDP := false
			for _, v := range dataArr {
				if v["ip"] == name || v["dnsvservername"] == name {
					if v["type"] == "TCP" {
						foundTCP = true
					}
					if v["type"] == "UDP" {
						foundUDP = true
					}
				}
			}
			if !foundTCP || !foundUDP {
				return fmt.Errorf("dnsnameserver %s with UDP_TCP type not found (TCP: %v, UDP: %v)", n, foundTCP, foundUDP)
			}
		} else {
			// For single type (TCP or UDP), check that specific type exists
			found := false
			for _, v := range dataArr {
				if v["ip"] == name || v["dnsvservername"] == name {
					found = true
				}
				if found == true {
					if v["type"] != dns_type {
						found = false
					} else {
						break
					}
				}
			}
			if !found {
				return fmt.Errorf("dnsnameserver %s not found", n)
			}
		}

		return nil
	}
}

func testAccCheckDnsnameserverDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_dnsnameserver" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		PrimaryId := rs.Primary.ID
		idSlice := strings.SplitN(PrimaryId, ",", 2)
		name := idSlice[0]
		dns_type := idSlice[1]

		dataArr, err := client.FindAllResources(service.Dnsnameserver.Type())

		if err != nil {
			return err
		}

		// For UDP_TCP type, check that both TCP and UDP are deleted
		if dns_type == "UDP_TCP" {
			foundTCP := false
			foundUDP := false
			for _, v := range dataArr {
				if v["ip"] == name || v["dnsvservername"] == name {
					if v["type"] == "TCP" {
						foundTCP = true
					}
					if v["type"] == "UDP" {
						foundUDP = true
					}
				}
			}
			if foundTCP {
				return fmt.Errorf("TCP dnsnameserver still exists for %s", name)
			}
			if foundUDP {
				return fmt.Errorf("UDP dnsnameserver still exists for %s", name)
			}
		} else {
			// For single type (TCP or UDP), check that specific type is deleted
			found := false
			for _, v := range dataArr {
				if v["ip"] == name || v["dnsvservername"] == name {
					found = true
				}
				if found == true {
					if v["type"] != dns_type {
						found = false
					} else {
						break
					}
				}
			}
			if found {
				return fmt.Errorf("dnsnameserver still exists %s", PrimaryId)
			}
		}
	}

	return nil
}

func testAccCheckDnsnameserverUdpTcpExist(n string, expectedIP string, expectedState string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No resource name is set")
		}

		// Use the shared utility function to get a configured client
		client, err := testAccGetClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		dataArr, err := client.FindAllResources(service.Dnsnameserver.Type())
		if err != nil {
			return err
		}

		// For UDP_TCP type, ADC creates two resources: one TCP and one UDP
		foundTCP := false
		foundUDP := false

		for _, v := range dataArr {
			if v["ip"] == expectedIP {
				if v["type"] == "TCP" {
					foundTCP = true
					if v["state"] != expectedState {
						return fmt.Errorf("TCP dnsnameserver %s has incorrect state: expected %s, got %s", expectedIP, expectedState, v["state"])
					}
				}
				if v["type"] == "UDP" {
					foundUDP = true
					if v["state"] != expectedState {
						return fmt.Errorf("UDP dnsnameserver %s has incorrect state: expected %s, got %s", expectedIP, expectedState, v["state"])
					}
				}
			}
		}

		if !foundTCP {
			return fmt.Errorf("TCP dnsnameserver with IP %s not found on ADC", expectedIP)
		}
		if !foundUDP {
			return fmt.Errorf("UDP dnsnameserver with IP %s not found on ADC", expectedIP)
		}

		return nil
	}
}

const testAccDnsnameserver_tcp = `
	resource "citrixadc_dnsnameserver" "dnsnameserver_tcp" {
		ip    = "192.0.2.10"
		local = false
		state = "ENABLED"
		type  = "TCP"
	}
`

func TestAccDnsnameserver_tcp(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckDnsnameserverDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDnsnameserver_tcp,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDnsnameserverExist("citrixadc_dnsnameserver.dnsnameserver_tcp", nil),
					resource.TestCheckResourceAttr("citrixadc_dnsnameserver.dnsnameserver_tcp", "ip", "192.0.2.10"),
					resource.TestCheckResourceAttr("citrixadc_dnsnameserver.dnsnameserver_tcp", "type", "TCP"),
					resource.TestCheckResourceAttr("citrixadc_dnsnameserver.dnsnameserver_tcp", "state", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_dnsnameserver.dnsnameserver_tcp", "local", "false"),
				),
			},
		},
	})
}

const testAccDnsnameserver_udp_tcp_step1 = `
	resource "citrixadc_dnsnameserver" "dnsnameserver_udp_tcp" {
		ip    = "192.0.2.25"
		state = "ENABLED"
		type  = "UDP_TCP"
	}
`

const testAccDnsnameserver_udp_tcp_step2 = `
	resource "citrixadc_dnsnameserver" "dnsnameserver_udp_tcp" {
		ip    = "192.0.2.25"
		state = "DISABLED"
		type  = "UDP_TCP"
	}
`

func TestAccDnsnameserver_udp_tcp(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckDnsnameserverDestroy,
		Steps: []resource.TestStep{
			// Create with UDP_TCP type - ADC creates both TCP and UDP resources
			{
				Config: testAccDnsnameserver_udp_tcp_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDnsnameserverExist("citrixadc_dnsnameserver.dnsnameserver_udp_tcp", nil),
					// Verify terraform state
					resource.TestCheckResourceAttr("citrixadc_dnsnameserver.dnsnameserver_udp_tcp", "ip", "192.0.2.25"),
					resource.TestCheckResourceAttr("citrixadc_dnsnameserver.dnsnameserver_udp_tcp", "state", "ENABLED"),
					// Verify ADC has both TCP and UDP resources
					testAccCheckDnsnameserverUdpTcpExist("citrixadc_dnsnameserver.dnsnameserver_udp_tcp", "192.0.2.25", "ENABLED"),
				),
			},
			// Update the dnsnameserver (change state)
			{
				Config: testAccDnsnameserver_udp_tcp_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDnsnameserverExist("citrixadc_dnsnameserver.dnsnameserver_udp_tcp", nil),
					// Verify terraform state
					resource.TestCheckResourceAttr("citrixadc_dnsnameserver.dnsnameserver_udp_tcp", "ip", "192.0.2.25"),
					resource.TestCheckResourceAttr("citrixadc_dnsnameserver.dnsnameserver_udp_tcp", "state", "DISABLED"),
					// Verify ADC has both TCP and UDP resources with updated state
					testAccCheckDnsnameserverUdpTcpExist("citrixadc_dnsnameserver.dnsnameserver_udp_tcp", "192.0.2.25", "DISABLED"),
				),
			},
			// Deletion is verified automatically by CheckDestroy after test completion
		},
	})
}
