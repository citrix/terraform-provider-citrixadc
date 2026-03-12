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
	"testing"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testAccLocationfile6_basic = `

resource "citrixadc_locationfile6" "tf_locationfile6" {
	locationfile = "/var/netscaler/inbuilt_db/Citrix_Netscaler_InBuilt_GeoIP_DB_IPv6"
	format       = "netscaler6"
	}
  
`

func TestAccLocationfile6_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLocationfile6Destroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLocationfile6_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLocationfile6Exist("citrixadc_locationfile6.tf_locationfile6", nil),
					resource.TestCheckResourceAttr("citrixadc_locationfile6.tf_locationfile6", "locationfile", "/var/netscaler/inbuilt_db/Citrix_Netscaler_InBuilt_GeoIP_DB_IPv6"),
					resource.TestCheckResourceAttr("citrixadc_locationfile6.tf_locationfile6", "format", "netscaler6"),
				),
			},
		},
	})
}

func testAccCheckLocationfile6Exist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No locationfile name is set")
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
		data, err := client.FindResource(service.Locationfile6.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("locationfile %s not found", n)
		}

		return nil
	}
}

func testAccCheckLocationfile6Destroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_locationfile" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Locationfile6.Type(), "")
		if err == nil {
			return fmt.Errorf("locationfile %s still exists", rs.Primary.ID)
		}
	}

	return nil
}

const testAccLocationfile6DataSource_basic = `

	resource "citrixadc_locationfile6" "tf_locationfile6" {
		locationfile = "/var/netscaler/inbuilt_db/Citrix_Netscaler_InBuilt_GeoIP_DB_IPv6"
		format       = "netscaler6"
	}

	data "citrixadc_locationfile6" "tf_locationfile6" {
		depends_on = [citrixadc_locationfile6.tf_locationfile6]
	}
`

func TestAccLocationfile6DataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccLocationfile6DataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_locationfile6.tf_locationfile6", "locationfile", "/var/netscaler/inbuilt_db/Citrix_Netscaler_InBuilt_GeoIP_DB_IPv6"),
					resource.TestCheckResourceAttr("data.citrixadc_locationfile6.tf_locationfile6", "format", "netscaler6"),
				),
			},
		},
	})
}
