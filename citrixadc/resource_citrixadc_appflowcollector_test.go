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
	"testing"
)

const testAccAppflowcollector_basic = `
resource "citrixadc_appflowcollector" "tf_appflowcollector" {
	name      = "tf_collector3"
	ipaddress = "192.168.2.3"
	transport = "logstream"
	port      =  80
	}
`
const testAccAppflowcollector_update = `
resource "citrixadc_appflowcollector" "tf_appflowcollector" {
	name      = "tf_collector3"
	ipaddress = "192.168.2.4"
	transport = "rest"
	port      = 90
	}
`

func TestAccAppflowcollector_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckAppflowcollectorDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAppflowcollector_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppflowcollectorExist("citrixadc_appflowcollector.tf_appflowcollector", nil),
					resource.TestCheckResourceAttr("citrixadc_appflowcollector.tf_appflowcollector", "name", "tf_collector3"),
					resource.TestCheckResourceAttr("citrixadc_appflowcollector.tf_appflowcollector", "transport", "logstream"),
					resource.TestCheckResourceAttr("citrixadc_appflowcollector.tf_appflowcollector", "ipaddress", "192.168.2.3"),
					resource.TestCheckResourceAttr("citrixadc_appflowcollector.tf_appflowcollector", "port", "80"),
				),
			},
			{
				Config: testAccAppflowcollector_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppflowcollectorExist("citrixadc_appflowcollector.tf_appflowcollector", nil),
					resource.TestCheckResourceAttr("citrixadc_appflowcollector.tf_appflowcollector", "name", "tf_collector3"),
					resource.TestCheckResourceAttr("citrixadc_appflowcollector.tf_appflowcollector", "transport", "rest"),
					resource.TestCheckResourceAttr("citrixadc_appflowcollector.tf_appflowcollector", "ipaddress", "192.168.2.4"),
					resource.TestCheckResourceAttr("citrixadc_appflowcollector.tf_appflowcollector", "port", "90"),
				),
			},
		},
	})
}

func testAccCheckAppflowcollectorExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No appflowcollector name is set")
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
		data, err := client.FindResource(service.Appflowcollector.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("appflowcollector %s not found", n)
		}

		return nil
	}
}

func testAccCheckAppflowcollectorDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_appflowcollector" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Appflowcollector.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("appflowcollector %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
