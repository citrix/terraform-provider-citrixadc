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

const testAccAppflowaction_analyticsprofile_binding_basic = `
	
	# Since the analyticsprofile resource is not yet available on Terraform,
	# the analyticsprofile policy must be created by hand in order for the script to run correctly.
	# You can do that by using the following Citrix ADC cli command:
	# add analyticsprofile <name> -type <type>
	
	resource "citrixadc_appflowaction_analyticsprofile_binding" "tf_appflowaction_analyticsprofile_binding" {
	name      = "test_action"
	analyticsprofile = "ns_analytics_global_profile"
	}




	# -------------------- ADC CLI ----------------------------
	#add appflow collector tf_collector -IPAddress 192.168.2.2
	#add appflowaction test_action -collectors tf_collector

	#---------------------NOT YET IMPLEMENTED ------------------------

	# resource "citrixadc_appflowaction" "tf_appflowaction" {
	#   name = "test_action"
	#   collectors = [citrixadc_appflowcollector.tf_appflowcollector.name,
	#                 citrixadc_appflowcollector.tf_appflowcollector2.name,]
	#   securityinsight = "ENABLED"
	#   botinsight      = "ENABLED"
	#   videoanalytics  = "ENABLED"
	# }
	# resource "citrixadc_appflowcollector" "tf_appflowcollector" {
	#   name      = "tf_collector"
	#   ipaddress = "192.168.2.2"
	#   port      = 80
	# }
	# resource "citrixadc_appflowcollector" "tf_appflowcollector2" {
	#   name      = "tf2_collector"
	#   ipaddress = "192.168.2.3"
	#   port      = 80
	# }
  
`

const testAccAppflowaction_analyticsprofile_binding_basic_step2 = `
  
# Since the analyticsprofile resource is not yet available on Terraform,
# the analyticsprofile policy must be created by hand in order for the script to run correctly.
# You can do that by using the following Citrix ADC cli command:
# add analyticsprofile <name> -type <type>




# -------------------- ADC CLI ----------------------------
#add appflow collector tf_collector -IPAddress 192.168.2.2
#add appflowaction test_action -collectors tf_collector

#---------------------NOT YET IMPLEMENTED ------------------------

# resource "citrixadc_appflowaction" "tf_appflowaction" {
#   name = "test_action"
#   collectors = [citrixadc_appflowcollector.tf_appflowcollector.name,
#                 citrixadc_appflowcollector.tf_appflowcollector2.name,]
#   securityinsight = "ENABLED"
#   botinsight      = "ENABLED"
#   videoanalytics  = "ENABLED"
# }
# resource "citrixadc_appflowcollector" "tf_appflowcollector" {
#   name      = "tf_collector"
#   ipaddress = "192.168.2.2"
#   port      = 80
# }
# resource "citrixadc_appflowcollector" "tf_appflowcollector2" {
#   name      = "tf2_collector"
#   ipaddress = "192.168.2.3"
#   port      = 80
# }
  
`

func TestAccAppflowaction_analyticsprofile_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAppflowaction_analyticsprofile_bindingDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccAppflowaction_analyticsprofile_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppflowaction_analyticsprofile_bindingExist("citrixadc_appflowaction_analyticsprofile_binding.tf_appflowaction_analyticsprofile_binding", nil),
				),
			},
			resource.TestStep{
				Config: testAccAppflowaction_analyticsprofile_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppflowaction_analyticsprofile_bindingNotExist("citrixadc_appflowaction_analyticsprofile_binding.tf_appflowaction_analyticsprofile_binding", "test_action,ns_analytics_global_profile"),
				),
			},
		},
	})
}

func testAccCheckAppflowaction_analyticsprofile_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No appflowaction_analyticsprofile_binding id is set")
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

		name := idSlice[0]
		analyticsprofile := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "appflowaction_analyticsprofile_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the matching analyticsprofile
		found := false
		for _, v := range dataArr {
			if v["analyticsprofile"].(string) == analyticsprofile {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("appflowaction_analyticsprofile_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckAppflowaction_analyticsprofile_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*NetScalerNitroClient).client

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}
		idSlice := strings.SplitN(id, ",", 2)

		name := idSlice[0]
		analyticsprofile := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "appflowaction_analyticsprofile_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to hopefully not find the one with the matching analyticsprofile
		found := false
		for _, v := range dataArr {
			if v["analyticsprofile"].(string) == analyticsprofile {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("appflowaction_analyticsprofile_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckAppflowaction_analyticsprofile_bindingDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_appflowaction_analyticsprofile_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource("appflowaction_analyticsprofile_binding", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("appflowaction_analyticsprofile_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
