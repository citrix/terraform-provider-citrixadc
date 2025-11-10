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

const testAccVpnintranetapplication_add = `

	resource "citrixadc_vpnintranetapplication" "tf_vpnintranetapplication" {
		intranetapplication = "tf_vpnintranetapplication"
		protocol            = "TCP"
		destip              = "2.3.6.5"
		interception        = "TRANSPARENT"
	}
`
const testAccVpnintranetapplication_update = `

	resource "citrixadc_vpnintranetapplication" "tf_vpnintranetapplication" {
		intranetapplication = "tf_vpnintranetapplication"
		protocol            = "UDP"
		destip              = "2.3.6.5"
		interception        = "TRANSPARENT"
	}
`

func TestAccVpnintranetapplication_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckVpnintranetapplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnintranetapplication_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnintranetapplicationExist("citrixadc_vpnintranetapplication.tf_vpnintranetapplication", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnintranetapplication.tf_vpnintranetapplication", "intranetapplication", "tf_vpnintranetapplication"),
					resource.TestCheckResourceAttr("citrixadc_vpnintranetapplication.tf_vpnintranetapplication", "protocol", "TCP"),
				),
			},
			{
				Config: testAccVpnintranetapplication_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnintranetapplicationExist("citrixadc_vpnintranetapplication.tf_vpnintranetapplication", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnintranetapplication.tf_vpnintranetapplication", "intranetapplication", "tf_vpnintranetapplication"),
					resource.TestCheckResourceAttr("citrixadc_vpnintranetapplication.tf_vpnintranetapplication", "protocol", "UDP"),
				),
			},
		},
	})
}

func testAccCheckVpnintranetapplicationExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No vpnintranetapplication name is set")
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
		data, err := client.FindResource(service.Vpnintranetapplication.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("vpnintranetapplication %s not found", n)
		}

		return nil
	}
}

func testAccCheckVpnintranetapplicationDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_vpnintranetapplication" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Vpnintranetapplication.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("vpnintranetapplication %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
