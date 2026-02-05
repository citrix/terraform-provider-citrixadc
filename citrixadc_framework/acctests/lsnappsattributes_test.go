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

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testAccLsnappsattributes_basic = `


resource "citrixadc_lsnappsattributes" "tf_lsnappsattributes" {
	name              = "my_lsn_appattributes"
	transportprotocol = "TCP"
	port              = 90
	sessiontimeout    = 40
	}
  
`
const testAccLsnappsattributes_update = `


resource "citrixadc_lsnappsattributes" "tf_lsnappsattributes" {
	name              = "my_lsn_appattributes"
	transportprotocol = "TCP"
	port              = 90
	sessiontimeout    = 60
	}
  
`

func TestAccLsnappsattributes_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLsnappsattributesDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLsnappsattributes_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsnappsattributesExist("citrixadc_lsnappsattributes.tf_lsnappsattributes", nil),
					resource.TestCheckResourceAttr("citrixadc_lsnappsattributes.tf_lsnappsattributes", "name", "my_lsn_appattributes"),
					resource.TestCheckResourceAttr("citrixadc_lsnappsattributes.tf_lsnappsattributes", "transportprotocol", "TCP"),
					resource.TestCheckResourceAttr("citrixadc_lsnappsattributes.tf_lsnappsattributes", "port", "90"),
					resource.TestCheckResourceAttr("citrixadc_lsnappsattributes.tf_lsnappsattributes", "sessiontimeout", "40"),
				),
			},
			{
				Config: testAccLsnappsattributes_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsnappsattributesExist("citrixadc_lsnappsattributes.tf_lsnappsattributes", nil),
					resource.TestCheckResourceAttr("citrixadc_lsnappsattributes.tf_lsnappsattributes", "name", "my_lsn_appattributes"),
					resource.TestCheckResourceAttr("citrixadc_lsnappsattributes.tf_lsnappsattributes", "transportprotocol", "TCP"),
					resource.TestCheckResourceAttr("citrixadc_lsnappsattributes.tf_lsnappsattributes", "port", "90"),
					resource.TestCheckResourceAttr("citrixadc_lsnappsattributes.tf_lsnappsattributes", "sessiontimeout", "60"),
				),
			},
		},
	})
}

func testAccCheckLsnappsattributesExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lsnappsattributes name is set")
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
		data, err := client.FindResource("lsnappsattributes", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("lsnappsattributes %s not found", n)
		}

		return nil
	}
}

func testAccCheckLsnappsattributesDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_lsnappsattributes" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("lsnappsattributes", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("lsnappsattributes %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccLsnappsattributesDataSource_basic = `

resource "citrixadc_lsnappsattributes" "tf_lsnappsattributes" {
	name              = "my_lsn_appattributes_ds"
	transportprotocol = "TCP"
	port              = 90
	sessiontimeout    = 40
}

data "citrixadc_lsnappsattributes" "tf_lsnappsattributes_ds" {
	name = citrixadc_lsnappsattributes.tf_lsnappsattributes.name
	depends_on = [citrixadc_lsnappsattributes.tf_lsnappsattributes]
}
`

func TestAccLsnappsattributesDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccLsnappsattributesDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_lsnappsattributes.tf_lsnappsattributes_ds", "name", "my_lsn_appattributes_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_lsnappsattributes.tf_lsnappsattributes_ds", "transportprotocol", "TCP"),
					resource.TestCheckResourceAttr("data.citrixadc_lsnappsattributes.tf_lsnappsattributes_ds", "port", "90"),
					resource.TestCheckResourceAttr("data.citrixadc_lsnappsattributes.tf_lsnappsattributes_ds", "sessiontimeout", "40"),
				),
			},
		},
	})
}
