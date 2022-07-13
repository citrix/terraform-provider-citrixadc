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
	"strings"
	"testing"
)

const testAccGslbservice_lbmonitor_binding_basic = `

resource "citrixadc_gslbservice_lbmonitor_binding" "tf_gslbservice_lbmonitor_binding" {
	monitor_name = citrixadc_lbmonitor.tfmonitor1.monitorname
	monstate    = "DISABLED"
	servicename = citrixadc_gslbservice.tf_gslbservice.servicename
	weight      = "20" 
  }
  
  resource "citrixadc_gslbservice" "tf_gslbservice" {
	ip          = "172.16.1.200"
	port        = "80"
	servicename = "tf_gslb1vservice"
	servicetype = "HTTP"
	sitename    = citrixadc_gslbsite.tf_gslbsite.sitename
  }
  
  resource "citrixadc_gslbsite" "tf_gslbsite" {
	sitename      = "tf_sitename"
	siteipaddress = "10.222.70.210"
  }
  
  resource "citrixadc_lbmonitor" "tfmonitor1" {
	monitorname = "tf_monitor"
	type        = "HTTP"
  }
  
`

const testAccGslbservice_lbmonitor_binding_basic_step2 = `
resource "citrixadc_gslbservice" "tf_gslbservice" {
	ip          = "172.16.1.200"
	port        = "80"
	servicename = "tf_gslb1vservice"
	servicetype = "HTTP"
	sitename    = citrixadc_gslbsite.tf_gslbsite.sitename
  }
  
  resource "citrixadc_gslbsite" "tf_gslbsite" {
	sitename      = "tf_sitename"
	siteipaddress = "10.222.70.210"
  }
  
  resource "citrixadc_lbmonitor" "tfmonitor1" {
	monitorname = "tf_monitor"
	type        = "HTTP"
  }
`

func TestAccGslbservice_lbmonitor_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
                CheckDestroy: testAccCheckGslbservice_lbmonitor_bindingDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
                                Config: testAccGslbservice_lbmonitor_binding_basic,
				Check: resource.ComposeTestCheckFunc(
                                        testAccCheckGslbservice_lbmonitor_bindingExist("citrixadc_gslbservice_lbmonitor_binding.tf_gslbservice_lbmonitor_binding", nil),								
                                        
				),
			},
			resource.TestStep{
                                Config: testAccGslbservice_lbmonitor_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
                                        testAccCheckGslbservice_lbmonitor_bindingNotExist("citrixadc_gslbservice_lbmonitor_binding.tf_gslbservice_lbmonitor_binding", "tf_gslb1vservice,tf_monitor"),
                                        
				),
			},
		},
	})
}

func testAccCheckGslbservice_lbmonitor_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No gslbservice_lbmonitor_binding id is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		client := testAccProvider.Meta().(*NetScalerNitroClient).client

		bindingId := rs.Primary.ID

		idSlice := strings.SplitN(bindingId, ",", 2)

		servicename := idSlice[0]
		monitor_name := idSlice[1]

		findParams := service.FindParams{
            ResourceType:             "gslbservice_lbmonitor_binding",
            ResourceName:             servicename,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the matching secondIdComponent
		found := false
		for _, v := range dataArr {
			if v["monitor_name"].(string) == monitor_name{
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("gslbservice_lbmonitor_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckGslbservice_lbmonitor_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*NetScalerNitroClient).client

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}
		idSlice := strings.SplitN(id, ",", 2)

		servicename := idSlice[0]
		monitor_name := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "gslbservice_lbmonitor_binding",
			ResourceName:             servicename,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to hopefully not find the one with the matching secondIdComponent
		found := false
		for _, v := range dataArr {
			if v["monitor_name"].(string) == monitor_name {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("gslbservice_lbmonitor_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckGslbservice_lbmonitor_bindingDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
                if rs.Type != "citrixadc_gslbservice_lbmonitor_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

                _, err := nsClient.FindResource(service.Gslbservice_lbmonitor_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("gslbservice_lbmonitor_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}