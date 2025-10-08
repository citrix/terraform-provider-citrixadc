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
	"strconv"
	"strings"

	"testing"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testAccServicegroup_servicegroupmember_binding_ipv4_step1 = `

resource "citrixadc_servicegroup" "tf_servicegroup" {
  servicegroupname = "tf_servicegroup"
  servicetype      = "HTTP"
}

resource "citrixadc_servicegroup_servicegroupmember_binding" "tf_binding" {
    servicegroupname = citrixadc_servicegroup.tf_servicegroup.servicegroupname
    ip = "10.78.22.33"
    port = 80
}
`

const testAccServicegroup_servicegroupmember_binding_ipv4_step2 = `

resource "citrixadc_servicegroup" "tf_servicegroup" {
  servicegroupname = "tf_servicegroup"
  servicetype      = "HTTP"
}
`

func TestAccServicegroup_servicegroupmember_binding_ipv4(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckServicegroup_servicegroupmember_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccServicegroup_servicegroupmember_binding_ipv4_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServicegroup_servicegroupmember_bindingExist("citrixadc_servicegroup_servicegroupmember_binding.tf_binding", nil),
					testAccCheckServicegroup_servicegroupmember_binding_not_exists("tf_servicegroup,10.78.22.33,80", true),
				),
			},
			{
				Config: testAccServicegroup_servicegroupmember_binding_ipv4_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServicegroup_servicegroupmember_binding_not_exists("tf_servicegroup,10.78.22.33,80", false),
				),
			},
		},
	})
}

const testAccServicegroup_servicegroupmember_binding_ipv6_step1 = `

resource "citrixadc_servicegroup" "tf_servicegroup" {
  servicegroupname = "tf_servicegroup"
  servicetype      = "HTTP"
}

resource "citrixadc_servicegroup_servicegroupmember_binding" "tf_binding" {
    servicegroupname = citrixadc_servicegroup.tf_servicegroup.servicegroupname
    ip = "ff::8839"
    port = 80
}
`

const testAccServicegroup_servicegroupmember_binding_ipv6_step2 = `

resource "citrixadc_servicegroup" "tf_servicegroup" {
  servicegroupname = "tf_servicegroup"
  servicetype      = "HTTP"
}
`

func TestAccServicegroup_servicegroupmember_binding_ipv6(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckServicegroup_servicegroupmember_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccServicegroup_servicegroupmember_binding_ipv6_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServicegroup_servicegroupmember_bindingExist("citrixadc_servicegroup_servicegroupmember_binding.tf_binding", nil),
					testAccCheckServicegroup_servicegroupmember_binding_not_exists("tf_servicegroup,ff::8839,80", true),
				),
			},
			{
				Config: testAccServicegroup_servicegroupmember_binding_ipv6_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServicegroup_servicegroupmember_binding_not_exists("tf_servicegroup,ff::8839,80", false),
				),
			},
		},
	})
}

const testAccServicegroup_servicegroupmember_binding_server_no_port_step1 = `

resource "citrixadc_servicegroup" "tf_servicegroup" {
  servicegroupname = "tf_servicegroup"
  servicetype      = "HTTP"
  autoscale = "DNS"
}
resource "citrixadc_server" "tf_server" {
    name = "tf_server"
    domain = "example.com"
    querytype = "SRV"
}

resource "citrixadc_servicegroup_servicegroupmember_binding" "tf_binding" {
    servicegroupname = citrixadc_servicegroup.tf_servicegroup.servicegroupname
    servername = citrixadc_server.tf_server.name
}
`

const testAccServicegroup_servicegroupmember_binding_server_no_port_step2 = `

resource "citrixadc_servicegroup" "tf_servicegroup" {
  servicegroupname = "tf_servicegroup"
  servicetype      = "HTTP"
  autoscale = "DNS"
}
resource "citrixadc_server" "tf_server" {
    name = "tf_server"
    domain = "example.com"
    querytype = "SRV"
}
`

func TestAccServicegroup_servicegroupmember_binding_server_no_port(t *testing.T) {
	t.Skip("TODO: Read error")
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckServicegroup_servicegroupmember_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccServicegroup_servicegroupmember_binding_server_no_port_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServicegroup_servicegroupmember_bindingExist("citrixadc_servicegroup_servicegroupmember_binding.tf_binding", nil),
					testAccCheckServicegroup_servicegroupmember_binding_not_exists("tf_servicegroup,tf_server", true),
				),
			},
			{
				Config: testAccServicegroup_servicegroupmember_binding_server_no_port_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServicegroup_servicegroupmember_binding_not_exists("tf_servicegroup,tf_server", false),
				),
			},
		},
	})
}

const testAccServicegroup_servicegroupmember_binding_server_with_port_step1 = `

resource "citrixadc_servicegroup" "tf_servicegroup" {
  servicegroupname = "tf_servicegroup"
  servicetype      = "HTTP"
  autoscale = "DNS"
}
resource "citrixadc_server" "tf_server" {
    name = "tf_server"
	ipaddress = "10.22.44.33"
}

resource "citrixadc_servicegroup_servicegroupmember_binding" "tf_binding" {
    servicegroupname = citrixadc_servicegroup.tf_servicegroup.servicegroupname
    servername = citrixadc_server.tf_server.name
	port = 80
}
`

const testAccServicegroup_servicegroupmember_binding_server_with_port_step2 = `

resource "citrixadc_servicegroup" "tf_servicegroup" {
  servicegroupname = "tf_servicegroup"
  servicetype      = "HTTP"
  autoscale = "DNS"
}
resource "citrixadc_server" "tf_server" {
    name = "tf_server"
	ipaddress = "10.22.44.33"
}

`

func TestAccServicegroup_servicegroupmember_binding_server_with_port(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckServicegroup_servicegroupmember_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccServicegroup_servicegroupmember_binding_server_with_port_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServicegroup_servicegroupmember_bindingExist("citrixadc_servicegroup_servicegroupmember_binding.tf_binding", nil),
					testAccCheckServicegroup_servicegroupmember_binding_not_exists("tf_servicegroup,tf_server,80", true),
				),
			},
			{
				Config: testAccServicegroup_servicegroupmember_binding_server_with_port_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServicegroup_servicegroupmember_binding_not_exists("tf_servicegroup,tf_server,80", false),
				),
			},
		},
	})
}

const testAccServicegroup_servicegroupmember_binding_mixed_bindings_step1 = `
resource "citrixadc_server" "tf_server" {
    name = "tf_server"
    domain = "example.com"
    querytype = "SRV"
}

resource "citrixadc_servicegroup" "tf_servicegroup" {
  servicegroupname = "tf_servicegroup"
  servicetype      = "HTTP"
  autoscale = "DNS"
}

resource "citrixadc_servicegroup_servicegroupmember_binding" "tf_binding" {
    servicegroupname = citrixadc_servicegroup.tf_servicegroup.servicegroupname
    servername = citrixadc_server.tf_server.name
	disable_read = true
}

resource "citrixadc_servicegroup_servicegroupmember_binding" "tf_binding2" {
    servicegroupname = citrixadc_servicegroup.tf_servicegroup.servicegroupname
    ip = "10.78.22.33"
    port = 80
}
`
const testAccServicegroup_servicegroupmember_binding_mixed_bindings_step2 = `
resource "citrixadc_server" "tf_server" {
    name = "tf_server"
    domain = "example.com"
    querytype = "SRV"
}

resource "citrixadc_servicegroup" "tf_servicegroup" {
  servicegroupname = "tf_servicegroup"
  servicetype      = "HTTP"
  autoscale = "DNS"
}

`

func TestAccServicegroup_servicegroupmember_binding_mixed_bindings(t *testing.T) {
	t.Skip("TODO:")
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckServicegroup_servicegroupmember_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccServicegroup_servicegroupmember_binding_mixed_bindings_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServicegroup_servicegroupmember_bindingExist("citrixadc_servicegroup_servicegroupmember_binding.tf_binding", nil),
					testAccCheckServicegroup_servicegroupmember_bindingExist("citrixadc_servicegroup_servicegroupmember_binding.tf_binding2", nil),
					testAccCheckServicegroup_servicegroupmember_binding_not_exists("tf_servicegroup,tf_server", true),
					testAccCheckServicegroup_servicegroupmember_binding_not_exists("tf_servicegroup,10.78.22.33,80", true),
				),
			},
			{
				Config: testAccServicegroup_servicegroupmember_binding_mixed_bindings_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServicegroup_servicegroupmember_binding_not_exists("tf_servicegroup,tf_server", false),
					testAccCheckServicegroup_servicegroupmember_binding_not_exists("tf_servicegroup,10.78.22.33,80", false),
				),
			},
		},
	})
}

func testAccCheckServicegroup_servicegroupmember_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No servicegroup_servicegroupmember_binding name is set")
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

		bindingId := rs.Primary.ID
		idSlice := strings.SplitN(bindingId, ",", 3)
		servicegroupname := idSlice[0]

		servername := idSlice[1]

		port := 0
		if len(idSlice) == 3 {
			if port, err = strconv.Atoi(idSlice[2]); err != nil {
				return err
			}
		}

		findParams := service.FindParams{
			ResourceType:             "servicegroup_servicegroupmember_binding",
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
			return fmt.Errorf("servicegroup_servicegroupmember_binding %s not found 123 %v and %v ", bindingId, servername, port)
		}

		return nil
	}
}

func testAccCheckServicegroup_servicegroupmember_binding_not_exists(bindingId string, invert bool) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		// Use the shared utility function to get a configured client
		client, err := testAccGetClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		idSlice := strings.SplitN(bindingId, ",", 3)
		servicegroupname := idSlice[0]

		servername := idSlice[1]

		port := 0
		if len(idSlice) == 3 {
			if port, err = strconv.Atoi(idSlice[2]); err != nil {
				return err
			}
		}

		findParams := service.FindParams{
			ResourceType:             "servicegroup_servicegroupmember_binding",
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
				if v["servername"].(string) == servername {
					foundIndex = i
					break
				}
			}
		}
		if !invert {
			if foundIndex != -1 {
				return fmt.Errorf("servicegroup_servicegroupmember_binding %s found. Should have been deleted", bindingId)
			}
		} else {
			if foundIndex == -1 {
				return fmt.Errorf("servicegroup_servicegroupmember_binding %s not found.", bindingId)
			}
		}

		return nil
	}
}

func testAccCheckServicegroup_servicegroupmember_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_servicegroup_servicegroupmember_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Servicegroup_servicegroupmember_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("servicegroup_servicegroupmember_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
