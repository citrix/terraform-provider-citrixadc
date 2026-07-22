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

const testAccQuicparam_add = `
	resource "citrixadc_quicparam" "tf_quicparam" {
		quicsecrettimeout = 3600
	}
`
const testAccQuicparam_update = `
	resource "citrixadc_quicparam" "tf_quicparam" {
		quicsecrettimeout = 7200
	}
`

func TestAccQuicparam_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccQuicparam_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckQuicparamExist("citrixadc_quicparam.tf_quicparam", nil),
					resource.TestCheckResourceAttr("citrixadc_quicparam.tf_quicparam", "quicsecrettimeout", "3600"),
				),
			},
			{
				Config: testAccQuicparam_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckQuicparamExist("citrixadc_quicparam.tf_quicparam", nil),
					resource.TestCheckResourceAttr("citrixadc_quicparam.tf_quicparam", "quicsecrettimeout", "7200"),
				),
			},
		},
	})
}

func TestAccQuicparam_import(t *testing.T) {
	const resAddr = "citrixadc_quicparam.tf_quicparam"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{Config: testAccQuicparam_add},
			{
				Config:                  testAccQuicparam_add,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccCheckQuicparamExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No quicparam name is set")
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
		data, err := client.FindResource(service.Quicparam.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("quicparam %s not found", n)
		}

		return nil
	}
}

const testAccQuicparamDataSource_basic = `

	resource "citrixadc_quicparam" "tf_quicparam" {
		quicsecrettimeout = 3600
	}

	data "citrixadc_quicparam" "tf_quicparam" {
		depends_on = [citrixadc_quicparam.tf_quicparam]
	}
`

func TestAccQuicparamDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccQuicparamDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_quicparam.tf_quicparam", "quicsecrettimeout", "3600"),
				),
			},
		},
	})
}
