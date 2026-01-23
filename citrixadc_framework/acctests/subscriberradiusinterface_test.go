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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"testing"
)

const testAccSubscriberradiusinterface_basic = `


resource "citrixadc_subscriberradiusinterface" "tf_subscriberradiusinterface" {
	listeningservice     = citrixadc_service.tf_service.name
	radiusinterimasstart = "ENABLED"
	}
  
  resource "citrixadc_service" "tf_service" {
	name        = "srad1"
	port        = 1813
	ip          = "192.0.0.206"
	servicetype = "RADIUSListener"
	}
  
`
const testAccSubscriberradiusinterface_update = `


resource "citrixadc_subscriberradiusinterface" "tf_subscriberradiusinterface" {
	listeningservice     = citrixadc_service.tf_service.name
	radiusinterimasstart = "DISABLED"
	}
  
  resource "citrixadc_service" "tf_service" {
	name        = "srad1"
	port        = 1813
	ip          = "192.0.0.206"
	servicetype = "RADIUSListener"
	}
  
  `

func TestAccSubscriberradiusinterface_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccSubscriberradiusinterface_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSubscriberradiusinterfaceExist("citrixadc_subscriberradiusinterface.tf_subscriberradiusinterface", nil),
					resource.TestCheckResourceAttr("citrixadc_subscriberradiusinterface.tf_subscriberradiusinterface", "listeningservice", "srad1"),
					resource.TestCheckResourceAttr("citrixadc_subscriberradiusinterface.tf_subscriberradiusinterface", "radiusinterimasstart", "ENABLED"),
				),
			},
			{
				Config: testAccSubscriberradiusinterface_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSubscriberradiusinterfaceExist("citrixadc_subscriberradiusinterface.tf_subscriberradiusinterface", nil),
					resource.TestCheckResourceAttr("citrixadc_subscriberradiusinterface.tf_subscriberradiusinterface", "listeningservice", "srad1"),
					resource.TestCheckResourceAttr("citrixadc_subscriberradiusinterface.tf_subscriberradiusinterface", "radiusinterimasstart", "DISABLED"),
				),
			},
		},
	})
}

func testAccCheckSubscriberradiusinterfaceExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No subscriberradiusinterface name is set")
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
		data, err := client.FindResource("subscriberradiusinterface", "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("subscriberradiusinterface %s not found", n)
		}

		return nil
	}
}
