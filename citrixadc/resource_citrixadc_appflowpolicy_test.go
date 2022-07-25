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
	"testing"
)

const testAccAppflowpolicy_basic = `
resource "citrixadc_appflowpolicy" "tf_appflowpolicy" {
	name      = "test_policy"
	action    = "test_action"
	rule      = "client.TCP.DSTPORT.EQ(22)"
  }
  
  # -------------------- ADC CLI ----------------------------
  #add appflow collector tf_collector -IPAddress 192.168.2.2
  #add appflowaction test_action -collectors tf_collector
  
  # ---------------- NOT YET IMPLEMENTED -------------------
  # resource "citrixadc_appflowaction" "tf_appflowaction" {
  #   name = "test_action"
  #   collectors     = [citrixadc_appflowcollector.tf_appflowcollector.name,
  #                    citrixadc_appflowcollector.tf_appflowcollector2.name, ]
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

const testAccAppflowpolicy_update = `
resource "citrixadc_appflowpolicy" "tf_appflowpolicy" {
	name      = "test_policy"
	action    = "test_action"
	rule      = "client.TCP.DSTPORT.EQ(25)"
  }
  
  # -------------------- ADC CLI ----------------------------
  #add appflow collector tf_collector -IPAddress 192.168.2.2
  #add appflowaction test_action -collectors tf_collector
  
  # ---------------- NOT YET IMPLEMENTED -------------------
  # resource "citrixadc_appflowaction" "tf_appflowaction" {
  #   name = "test_action"
  #   collectors     = [citrixadc_appflowcollector.tf_appflowcollector.name,
  #                    citrixadc_appflowcollector.tf_appflowcollector2.name, ]
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

func TestAccAppflowpolicy_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAppflowpolicyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccAppflowpolicy_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppflowpolicyExist("citrixadc_appflowpolicy.tf_appflowpolicy", nil),
					resource.TestCheckResourceAttr("citrixadc_appflowpolicy.tf_appflowpolicy", "name", "test_policy"),
					resource.TestCheckResourceAttr("citrixadc_appflowpolicy.tf_appflowpolicy", "action", "test_action"),
					resource.TestCheckResourceAttr("citrixadc_appflowpolicy.tf_appflowpolicy", "rule", "client.TCP.DSTPORT.EQ(22)"),
				),
			},
			resource.TestStep{
				Config: testAccAppflowpolicy_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppflowpolicyExist("citrixadc_appflowpolicy.tf_appflowpolicy", nil),
					resource.TestCheckResourceAttr("citrixadc_appflowpolicy.tf_appflowpolicy", "name", "test_policy"),
					resource.TestCheckResourceAttr("citrixadc_appflowpolicy.tf_appflowpolicy", "action", "test_action"),
					resource.TestCheckResourceAttr("citrixadc_appflowpolicy.tf_appflowpolicy", "rule", "client.TCP.DSTPORT.EQ(25)"),
				),
			},
		},
	})
}

func testAccCheckAppflowpolicyExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No appflowpolicy name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(service.Appflowpolicy.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("appflowpolicy %s not found", n)
		}

		return nil
	}
}

func testAccCheckAppflowpolicyDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_appflowpolicy" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Appflowpolicy.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("appflowpolicy %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
