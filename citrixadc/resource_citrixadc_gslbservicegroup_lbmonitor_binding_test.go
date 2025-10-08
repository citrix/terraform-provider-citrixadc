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
	"strings"
	"testing"
)

const testAccGslbservicegroup_lbmonitor_binding_basic = `

resource "citrixadc_gslbservicegroup_lbmonitor_binding" "tf_gslbservicegroup_lbmonitor_binding" {
	weight           = 20
	servicegroupname = citrixadc_gslbservicegroup.tf_gslbservicegroup.servicegroupname
	monitor_name      = citrixadc_lbmonitor.tfmonitor1.monitorname
  
	}
  
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
	sitepassword    = "password123"
	}
  
  resource "citrixadc_lbmonitor" "tfmonitor1" {
	monitorname = "tf_monitor"
	type        = "HTTP"
	}
`

const testAccGslbservicegroup_lbmonitor_binding_basic_step2 = `

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
	sitepassword    = "password123"
	}
  
  resource "citrixadc_lbmonitor" "tfmonitor1" {
	monitorname = "tf_monitor"
	type        = "HTTP"
	}
`

func TestAccGslbservicegroup_lbmonitor_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckGslbservicegroup_lbmonitor_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGslbservicegroup_lbmonitor_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGslbservicegroup_lbmonitor_bindingExist("citrixadc_gslbservicegroup_lbmonitor_binding.tf_gslbservicegroup_lbmonitor_binding", nil),
				),
			},
			{
				Config: testAccGslbservicegroup_lbmonitor_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGslbservicegroup_lbmonitor_bindingNotExist("citrixadc_gslbservicegroup_lbmonitor_binding.tf_gslbservicegroup_lbmonitor_binding", "test_gslbvservicegroup,tf_monitor"),
				),
			},
		},
	})
}

func testAccCheckGslbservicegroup_lbmonitor_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No gslbservicegroup_lbmonitor_binding id is set")
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

		idSlice := strings.SplitN(bindingId, ",", 2)

		servicegroupname := idSlice[0]
		monitor_name := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "gslbservicegroup_lbmonitor_binding",
			ResourceName:             servicegroupname,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the matching monitor_name
		found := false
		for _, v := range dataArr {
			if v["monitor_name"].(string) == monitor_name {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("gslbservicegroup_lbmonitor_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckGslbservicegroup_lbmonitor_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}
		idSlice := strings.SplitN(id, ",", 2)

		servicegroupname := idSlice[0]
		monitor_name := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "gslbservicegroup_lbmonitor_binding",
			ResourceName:             servicegroupname,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to hopefully not find the one with the matching monitor_name
		found := false
		for _, v := range dataArr {
			if v["monitor_name"].(string) == monitor_name {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("gslbservicegroup_lbmonitor_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckGslbservicegroup_lbmonitor_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_gslbservicegroup_lbmonitor_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("gslbservicegroup_lbmonitor_binding", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("gslbservicegroup_lbmonitor_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
