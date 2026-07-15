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

const testAccNsappflowparam_basic_step1 = `
	resource "citrixadc_nsappflowparam" "tf_nsappflowparam" {
		templaterefresh   = 600
		httpurl           = "ON"
		clienttrafficonly = "NO"
	}
`

const testAccNsappflowparam_basic_step2 = `
	resource "citrixadc_nsappflowparam" "tf_nsappflowparam" {
		templaterefresh   = 700
		httpurl           = "OFF"
		clienttrafficonly = "YES"
	}
`

func TestAccNsappflowparam_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// Singleton resource: never truly deleted from the ADC (state-only removal).
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				Config: testAccNsappflowparam_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsappflowparamExist("citrixadc_nsappflowparam.tf_nsappflowparam", nil),
					resource.TestCheckResourceAttr("citrixadc_nsappflowparam.tf_nsappflowparam", "templaterefresh", "600"),
					resource.TestCheckResourceAttr("citrixadc_nsappflowparam.tf_nsappflowparam", "httpurl", "ON"),
					resource.TestCheckResourceAttr("citrixadc_nsappflowparam.tf_nsappflowparam", "clienttrafficonly", "NO"),
				),
			},
			{
				Config: testAccNsappflowparam_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsappflowparamExist("citrixadc_nsappflowparam.tf_nsappflowparam", nil),
					resource.TestCheckResourceAttr("citrixadc_nsappflowparam.tf_nsappflowparam", "templaterefresh", "700"),
					resource.TestCheckResourceAttr("citrixadc_nsappflowparam.tf_nsappflowparam", "httpurl", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_nsappflowparam.tf_nsappflowparam", "clienttrafficonly", "YES"),
				),
			},
		},
	})
}

func TestAccNsappflowparam_import(t *testing.T) {
	const resAddr = "citrixadc_nsappflowparam.tf_nsappflowparam"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// Singleton resource: never truly deleted from the ADC (state-only removal).
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{Config: testAccNsappflowparam_basic_step1},
			{
				// Import id is the synthetic constant set in Create ("nsappflowparam-config"),
				// resolved by ImportStatePassthroughID from the stored resource id.
				Config:                  testAccNsappflowparam_basic_step1,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccCheckNsappflowparamExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No nsappflowparam name is set")
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
		// Singleton: get-all, no name argument.
		data, err := client.FindResource(service.Nsappflowparam.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("nsappflowparam %s not found", n)
		}

		return nil
	}
}

const testAccNsappflowparamDataSource_basic = `

	resource "citrixadc_nsappflowparam" "tf_nsappflowparam" {
		templaterefresh   = 600
		httpurl           = "OFF"
		clienttrafficonly = "NO"
	}

	data "citrixadc_nsappflowparam" "tf_nsappflowparam" {
		depends_on = [citrixadc_nsappflowparam.tf_nsappflowparam]
	}
`

func TestAccNsappflowparamDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccNsappflowparamDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_nsappflowparam.tf_nsappflowparam", "templaterefresh", "600"),
					resource.TestCheckResourceAttr("data.citrixadc_nsappflowparam.tf_nsappflowparam", "httpurl", "OFF"),
					resource.TestCheckResourceAttr("data.citrixadc_nsappflowparam.tf_nsappflowparam", "clienttrafficonly", "NO"),
				),
			},
		},
	})
}
