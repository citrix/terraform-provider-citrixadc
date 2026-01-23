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

const testAccSubscribergxinterface_basic = `

	resource "citrixadc_service" "tf_service" {
		name 		= "pcrf-svc1"
		port 		= 80
		ip 			= "10.202.22.12"
		servicetype = "HTTP"
	}
	resource "citrixadc_subscribergxinterface" "tf_subscribergxinterface" {
		service        = citrixadc_service.tf_service.name
		pcrfrealm      = "myrealm.com"
		healthcheck    = "YES"
		servicepathavp = [26009]
		healthcheckttl = 30
	}
`
const testAccSubscribergxinterface_update = `

	resource "citrixadc_service" "tf_service" {
		name 		= "pcrf-svc1"
		port 		= 80
		ip 			= "10.202.22.12"
		servicetype = "HTTP"
	}

	resource "citrixadc_subscribergxinterface" "tf_subscribergxinterface" {
		service        = citrixadc_service.tf_service.name
		pcrfrealm      = "myrealm2.com"
		healthcheck    = "NO"
		servicepathavp = [26010]
		healthcheckttl = 40
	}
`

func TestAccSubscribergxinterface_basic(t *testing.T) {
	t.Skip("TODO: Need to find a way to test this resource!")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccSubscribergxinterface_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSubscribergxinterfaceExist("citrixadc_subscribergxinterface.tf_subscribergxinterface", nil),
					resource.TestCheckResourceAttr("citrixadc_subscribergxinterface.tf_subscribergxinterface", "service", "pcrf-svc1"),
					resource.TestCheckResourceAttr("citrixadc_subscribergxinterface.tf_subscribergxinterface", "pcrfrealm", "myrealm.com"),
					resource.TestCheckResourceAttr("citrixadc_subscribergxinterface.tf_subscribergxinterface", "healthcheck", "YES"),
					resource.TestCheckResourceAttr("citrixadc_subscribergxinterface.tf_subscribergxinterface", "healthcheckttl", "30"),
				),
			},
			{
				Config: testAccSubscribergxinterface_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSubscribergxinterfaceExist("citrixadc_subscribergxinterface.tf_subscribergxinterface", nil),
					resource.TestCheckResourceAttr("citrixadc_subscribergxinterface.tf_subscribergxinterface", "service", "pcrf-svc1"),
					resource.TestCheckResourceAttr("citrixadc_subscribergxinterface.tf_subscribergxinterface", "pcrfrealm", "myrealm2.com"),
					resource.TestCheckResourceAttr("citrixadc_subscribergxinterface.tf_subscribergxinterface", "healthcheck", "NO"),
					resource.TestCheckResourceAttr("citrixadc_subscribergxinterface.tf_subscribergxinterface", "healthcheckttl", "40"),
				),
			},
		},
	})
}

func testAccCheckSubscribergxinterfaceExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No subscribergxinterface name is set")
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
		data, err := client.FindResource("subscribergxinterface", "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("subscribergxinterface %s not found", n)
		}

		return nil
	}
}
