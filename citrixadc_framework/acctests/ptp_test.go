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

const testAccPtp_add = `
	resource "citrixadc_ptp" "tf_ptp" {
		state = "DISABLE"
	}
`
const testAccPtp_update = `
	resource "citrixadc_ptp" "tf_ptp" {
		state = "ENABLE"
	}
`

func TestAccPtp_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccPtp_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPtpExist("citrixadc_ptp.tf_ptp", nil),
					resource.TestCheckResourceAttr("citrixadc_ptp.tf_ptp", "state", "DISABLE"),
				),
			},
			{
				Config: testAccPtp_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPtpExist("citrixadc_ptp.tf_ptp", nil),
					resource.TestCheckResourceAttr("citrixadc_ptp.tf_ptp", "state", "ENABLE"),
				),
			},
		},
	})
}

func testAccCheckPtpExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ptp name is set")
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
		data, err := client.FindResource(service.Ptp.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("ptp %s not found", n)
		}

		return nil
	}
}

const testAccPtpDataSource_basic = `
	resource "citrixadc_ptp" "tf_ptp_ds" {
		state = "ENABLE"
	}

	data "citrixadc_ptp" "tf_ptp_ds" {
		depends_on = [citrixadc_ptp.tf_ptp_ds]
	}
`

func TestAccPtpDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccPtpDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.citrixadc_ptp.tf_ptp_ds", "id"),
					resource.TestCheckResourceAttr("data.citrixadc_ptp.tf_ptp_ds", "state", "ENABLE"),
				),
			},
		},
	})
}
