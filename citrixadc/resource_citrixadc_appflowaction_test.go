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
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"testing"
)

const testAccAppflowaction_basic = `	

	resource "citrixadc_appflowaction" "tf_appflowaction" {
		name            = "test_action"
		collectors      = ["tf_collector", "tf2_collector" ]
		securityinsight = "ENABLED"
		botinsight      = "ENABLED"
		videoanalytics  = "ENABLED"
		}
		
# -------------------- ADC CLI ----------------------------
#add appflow collector tf2_collector -IPAddress 192.168.2.3
#add appflow collector tf_collector -IPAddress 192.168.2.2

# ----------------- NOT YET IMPLEMENTED -----------------------
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

const testAccAppflowaction_update = `	

resource "citrixadc_appflowaction" "tf_appflowaction" {
	name            = "test_action"
	collectors      = ["tf_collector", "tf2_collector" ]
	securityinsight = "DISABLED"
	botinsight      = "DISABLED"
	videoanalytics  = "DISABLED"
}

# -------------------- ADC CLI ----------------------------
#add appflow collector tf2_collector -IPAddress 192.168.2.3
#add appflow collector tf_collector -IPAddress 192.168.2.2

# ----------------- NOT YET IMPLEMENTED -----------------------
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

func TestAccAppflowaction_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAppflowactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAppflowaction_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppflowactionExist("citrixadc_appflowaction.tf_appflowaction", nil),
					resource.TestCheckResourceAttr("citrixadc_appflowaction.tf_appflowaction", "name", "test_action"),
					resource.TestCheckResourceAttr("citrixadc_appflowaction.tf_appflowaction", "securityinsight", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowaction.tf_appflowaction", "videoanalytics", "ENABLED"),
				),
			},
			{
				Config: testAccAppflowaction_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppflowactionExist("citrixadc_appflowaction.tf_appflowaction", nil),
					resource.TestCheckResourceAttr("citrixadc_appflowaction.tf_appflowaction", "name", "test_action"),
					resource.TestCheckResourceAttr("citrixadc_appflowaction.tf_appflowaction", "securityinsight", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowaction.tf_appflowaction", "videoanalytics", "DISABLED"),
				),
			},
		},
	})
}

func testAccCheckAppflowactionExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No appflowaction name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(service.Appflowaction.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("appflowaction %s not found", n)
		}

		return nil
	}
}

func testAccCheckAppflowactionDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_appflowaction" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Appflowaction.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("appflowaction %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
