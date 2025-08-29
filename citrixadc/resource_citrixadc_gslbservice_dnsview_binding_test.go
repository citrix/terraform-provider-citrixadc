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
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

const testAccGslbservice_dnsview_binding_basic = `

resource "citrixadc_gslbservice_dnsview_binding" "tf_gslbservice_dnsview_binding" {
	servicename = citrixadc_gslbservice.gslb_svc1.servicename
	viewname    = citrixadc_dnsview.tf_dnsview.viewname
	viewip      = "192.168.2.1"
  }
  
  resource "citrixadc_gslbsite" "site_remote" {
	sitename        = "Site-Remote"
	siteipaddress   = "172.31.48.18"
	sessionexchange = "ENABLED"
	sitepassword = "password123"
  }
  
  resource "citrixadc_gslbservice" "gslb_svc1" {
	ip          = "172.16.1.121"
	port        = "80"
	servicename = "gslb1vservice"
	servicetype = "HTTP"
	sitename    = citrixadc_gslbsite.site_remote.sitename
  }
  
  resource "citrixadc_dnsview" "tf_dnsview" {
	viewname = "view4"
  } 
`

const testAccGslbservice_dnsview_binding_basic_step2 = `
resource "citrixadc_gslbsite" "site_remote" {
	sitename        = "Site-Remote"
	siteipaddress   = "172.31.48.18"
	sessionexchange = "ENABLED"
	sitepassword = "password123"
  }
  
  resource "citrixadc_gslbservice" "gslb_svc1" {
	ip          = "172.16.1.121"
	port        = "80"
	servicename = "gslb1vservice"
	servicetype = "HTTP"
	sitename    = citrixadc_gslbsite.site_remote.sitename
  }
  resource "citrixadc_dnsview" "tf_dnsview" {
	viewname = "view4"
  } 
`

func TestAccGslbservice_dnsview_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckGslbservice_dnsview_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGslbservice_dnsview_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGslbservice_dnsview_bindingExist("citrixadc_gslbservice_dnsview_binding.tf_gslbservice_dnsview_binding", nil),
				),
			},
			{
				Config: testAccGslbservice_dnsview_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGslbservice_dnsview_bindingNotExist("citrixadc_gslbservice_dnsview_binding.tf_gslbservice_dnsview_binding", "gslb1vservice,view4"),
				),
			},
		},
	})
}

func testAccCheckGslbservice_dnsview_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No gslbservice_dnsview_binding id is set")
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
		viewname := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "gslbservice_dnsview_binding",
			ResourceName:             servicename,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the matching dnsview
		found := false
		for _, v := range dataArr {
			if v["viewname"].(string) == viewname {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("gslbservice_dnsview_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckGslbservice_dnsview_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*NetScalerNitroClient).client

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}
		idSlice := strings.SplitN(id, ",", 2)

		servicename := idSlice[0]
		viewname := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "gslbservice_dnsview_binding",
			ResourceName:             servicename,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to hopefully not find the one with the matching secondview
		found := false
		for _, v := range dataArr {
			if v["viewname"].(string) == viewname {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("gslbservice_dnsview_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckGslbservice_dnsview_bindingDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_gslbservice_dnsview_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Gslbservice_dnsview_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("gslbservice_dnsview_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
