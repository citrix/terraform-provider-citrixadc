/*
Copyright 2024 Citrix Systems, Inc

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
	"strconv"
	"testing"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testAccGslbservicegroup_gslbservicegroupmember_binding_basic = `

	resource "citrixadc_gslbservicegroup" "tf_gslbservicegroup" {
		servicegroupname = "test_gslbvservicegroup"
		servicetype      = "HTTP"
		cip              = "DISABLED"
		healthmonitor    = "NO"
		sitename         = citrixadc_gslbsite.site_local.sitename
	}
	resource "citrixadc_gslbsite" "site_local" {
		sitename        = "Site-Local"
		siteipaddress   = "172.31.96.234"
		sessionexchange = "DISABLED"
		sitepassword = "password123"
	}
	resource "citrixadc_server" "tf_server" {
		name = "tf_server"
		ipaddress = "192.168.11.13"
	}
	
	resource "citrixadc_gslbservicegroup_gslbservicegroupmember_binding" "tf_binding" {
		servicegroupname = citrixadc_gslbservicegroup.tf_gslbservicegroup.servicegroupname
		servername       = citrixadc_server.tf_server.name
		port             = 60
	}
	
`

const testAccGslbservicegroup_gslbservicegroupmember_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion

	resource "citrixadc_gslbservicegroup" "tf_gslbservicegroup" {
		servicegroupname = "test_gslbvservicegroup"
		servicetype      = "HTTP"
		cip              = "DISABLED"
		healthmonitor    = "NO"
		sitename         = citrixadc_gslbsite.site_local.sitename
	}
	resource "citrixadc_gslbsite" "site_local" {
		sitename        = "Site-Local"
		siteipaddress   = "172.31.96.234"
		sessionexchange = "DISABLED"
		sitepassword = "password123"
	}
	resource "citrixadc_server" "tf_server" {
		name = "tf_server"
		ipaddress = "192.168.11.13"
	}
`

func TestAccGslbservicegroup_gslbservicegroupmember_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckGslbservicegroup_gslbservicegroupmember_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGslbservicegroup_gslbservicegroupmember_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGslbservicegroup_gslbservicegroupmember_bindingExist("citrixadc_gslbservicegroup_gslbservicegroupmember_binding.tf_binding", nil),
				),
			},
			{
				Config: testAccGslbservicegroup_gslbservicegroupmember_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGslbservicegroup_gslbservicegroupmember_bindingNotExist("citrixadc_gslbservicegroup_gslbservicegroupmember_binding.tf_binding", "test_gslbvservicegroup,10.10.10.10,60"),
				),
			},
		},
	})
}

func testAccCheckGslbservicegroup_gslbservicegroupmember_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No gslbservicegroup_gslbservicegroupmember_binding id is set")
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

		bindingId := rs.Primary.ID
		idMap, _, err := utils.ParseIdString(bindingId, []string{"servicegroupname", "servername", "port"}, []string{"servername", "port"})
		if err != nil {
			return fmt.Errorf("Error parsing ID %v: %v", bindingId, err)
		}
		servicegroupname := idMap["servicegroupname"]
		servername := idMap["servername"]

		port := 0
		if p, ok := idMap["port"]; ok && p != "" {
			if port, err = strconv.Atoi(p); err != nil {
				return err
			}
		}

		findParams := service.FindParams{
			ResourceType:             "gslbservicegroup_gslbservicegroupmember_binding",
			ResourceName:             servicegroupname,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the right policy name
		foundIndex := -1
		for i, v := range dataArr {
			if port != 0 {
				portEqual := int(v["port"].(float64)) == port
				servernameEqual := v["servername"] == servername
				if servernameEqual && portEqual {
					foundIndex = i
					break
				}
			} else {
				log.Printf("[DEBUG] teh val sis  %v, %v", v["servername"].(string), servername)
				if v["servername"].(string) == servername {
					foundIndex = i
					break
				}
			}
			log.Printf("[DEBUG] teh val sis  %v, %v", v["servername"].(string), servername)
		}

		if foundIndex == -1 {
			return fmt.Errorf("gslbservicegroup_gslbservicegroupmember_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckGslbservicegroup_gslbservicegroupmember_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		idMap, _, err := utils.ParseIdString(id, []string{"servicegroupname", "servername", "port"}, []string{"servername", "port"})
		if err != nil {
			return fmt.Errorf("Error parsing ID %v: %v", id, err)
		}
		servicegroupname := idMap["servicegroupname"]
		servername := idMap["servername"]

		port := 0
		if p, ok := idMap["port"]; ok && p != "" {
			if port, err = strconv.Atoi(p); err != nil {
				return err
			}
		}

		findParams := service.FindParams{
			ResourceType:             "gslbservicegroup_gslbservicegroupmember_binding",
			ResourceName:             servicegroupname,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to hopefully not find the one with the matching secondIdComponent
		// Iterate through results to find the one with the right policy name
		foundIndex := -1
		for i, v := range dataArr {
			if port != 0 {
				portEqual := int(v["port"].(float64)) == port
				servernameEqual := v["servername"] == servername
				if servernameEqual && portEqual {
					foundIndex = i
					break
				}
			}
		}

		if foundIndex != -1 {
			return fmt.Errorf("servicegroup_servicegroupmember_binding %s found. Should have been deleted", id)
		}

		return nil
	}
}

func testAccCheckGslbservicegroup_gslbservicegroupmember_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_gslbservicegroup_gslbservicegroupmember_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("gslbservicegroup_gslbservicegroupmember_binding", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("gslbservicegroup_gslbservicegroupmember_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccGslbservicegroup_gslbservicegroupmember_bindingDataSource_basic = `

	resource "citrixadc_gslbservicegroup" "tf_gslbservicegroup" {
		servicegroupname = "test_gslbvservicegroup"
		servicetype      = "HTTP"
		cip              = "DISABLED"
		healthmonitor    = "NO"
		sitename         = citrixadc_gslbsite.site_local.sitename
	}
	resource "citrixadc_gslbsite" "site_local" {
		sitename        = "Site-Local"
		siteipaddress   = "172.31.96.234"
		sessionexchange = "DISABLED"
		sitepassword = "password123"
	}
	resource "citrixadc_server" "tf_server" {
		name = "tf_server"
		ipaddress = "192.168.11.13"
	}
	
	resource "citrixadc_gslbservicegroup_gslbservicegroupmember_binding" "tf_binding" {
		servicegroupname = citrixadc_gslbservicegroup.tf_gslbservicegroup.servicegroupname
		servername       = citrixadc_server.tf_server.name
		port             = 60
	}

	data "citrixadc_gslbservicegroup_gslbservicegroupmember_binding" "tf_binding" {
		servicegroupname = citrixadc_gslbservicegroup_gslbservicegroupmember_binding.tf_binding.servicegroupname
		servername       = citrixadc_gslbservicegroup_gslbservicegroupmember_binding.tf_binding.servername
		port             = citrixadc_gslbservicegroup_gslbservicegroupmember_binding.tf_binding.port
	}
`

const testAccGslbservicegroup_gslbservicegroupmember_binding_upgrade_basic = `

	resource "citrixadc_gslbservicegroup" "tf_gslbservicegroup" {
		servicegroupname = "test_gslbvservicegroup"
		servicetype      = "HTTP"
		cip              = "DISABLED"
		healthmonitor    = "NO"
		sitename         = citrixadc_gslbsite.site_local.sitename
	}
	resource "citrixadc_gslbsite" "site_local" {
		sitename        = "Site-Local"
		siteipaddress   = "172.31.96.234"
		sessionexchange = "DISABLED"
		sitepassword = "password123"
	}
	resource "citrixadc_server" "tf_server" {
		name = "tf_server"
		ipaddress = "192.168.11.13"
	}

	resource "citrixadc_gslbservicegroup_gslbservicegroupmember_binding" "tf_binding" {
		servicegroupname = citrixadc_gslbservicegroup.tf_gslbservicegroup.servicegroupname
		servername       = citrixadc_server.tf_server.name
		port             = 60
	}

`

// TestAccGslbservicegroup_gslbservicegroupmember_binding_sdkv2StateUpgrade verifies that a
// resource created with the last SDK v2 release (2.2.0, legacy comma-separated ID) is
// correctly refreshed/planned/applied by the current (framework-muxed) provider, and that
// the ID is upgraded to the new key:value format on Read.
func TestAccGslbservicegroup_gslbservicegroupmember_binding_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckGslbservicegroup_gslbservicegroupmember_bindingDestroy,
		Steps: []resource.TestStep{
			{
				// Step 1: create with the last SDK v2 release -> legacy comma-separated ID.
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.2.0",
					},
				},
				Config: testAccGslbservicegroup_gslbservicegroupmember_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGslbservicegroup_gslbservicegroupmember_bindingExist("citrixadc_gslbservicegroup_gslbservicegroupmember_binding.tf_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_gslbservicegroup_gslbservicegroupmember_binding.tf_binding", "id", "test_gslbvservicegroup,tf_server,60"),
				),
			},
			{
				// Step 2: same config through the current (framework) provider. Read exercises
				// ParseIdString on the legacy ID and recomputes the ID to the new key:value format.
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccGslbservicegroup_gslbservicegroupmember_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGslbservicegroup_gslbservicegroupmember_bindingExist("citrixadc_gslbservicegroup_gslbservicegroupmember_binding.tf_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_gslbservicegroup_gslbservicegroupmember_binding.tf_binding", "id", "servicegroupname:test_gslbvservicegroup,servername:tf_server,port:60"),
				),
			},
		},
	})
}

func TestAccGslbservicegroup_gslbservicegroupmember_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccGslbservicegroup_gslbservicegroupmember_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_gslbservicegroup_gslbservicegroupmember_binding.tf_binding", "servicegroupname", "test_gslbvservicegroup"),
					resource.TestCheckResourceAttr("data.citrixadc_gslbservicegroup_gslbservicegroupmember_binding.tf_binding", "servername", "tf_server"),
					resource.TestCheckResourceAttr("data.citrixadc_gslbservicegroup_gslbservicegroupmember_binding.tf_binding", "port", "60"),
					resource.TestCheckResourceAttr("data.citrixadc_gslbservicegroup_gslbservicegroupmember_binding.tf_binding", "ip", "192.168.11.13"),
				),
			},
		},
	})
}

func TestAccGslbservicegroup_gslbservicegroupmember_binding_import(t *testing.T) {
	const resAddr = "citrixadc_gslbservicegroup_gslbservicegroupmember_binding.tf_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckGslbservicegroup_gslbservicegroupmember_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccGslbservicegroup_gslbservicegroupmember_binding_basic},
			{Config: testAccGslbservicegroup_gslbservicegroupmember_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
		},
	})
}
