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
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDnsnameserverDestroy,
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
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		dataArr, err := client.FindAllResources(service.Dnsnameserver.Type())

		if err != nil {
			return err
		}

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

		return nil
	}
}

func testAccCheckDnsnameserverDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
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
			return fmt.Errorf("dnsnameserver Still exists %s not found", PrimaryId)
		}

	}

	return nil
}
