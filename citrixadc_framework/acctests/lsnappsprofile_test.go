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

const testAccLsnappsprofile_basic = `

	resource "citrixadc_lsnappsprofile" "tf_lsnappsprofile" {
		appsprofilename   = "my_lsn_appsprofile"
		transportprotocol = "TCP"
		mapping           = "ENDPOINT-INDEPENDENT"
	}
`
const testAccLsnappsprofile_update = `

	resource "citrixadc_lsnappsprofile" "tf_lsnappsprofile" {
		appsprofilename   = "my_lsn_appsprofile"
		transportprotocol = "TCP"
		mapping           = "ADDRESS-DEPENDENT"
	}
`

func TestAccLsnappsprofile_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLsnappsprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLsnappsprofile_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsnappsprofileExist("citrixadc_lsnappsprofile.tf_lsnappsprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_lsnappsprofile.tf_lsnappsprofile", "appsprofilename", "my_lsn_appsprofile"),
					resource.TestCheckResourceAttr("citrixadc_lsnappsprofile.tf_lsnappsprofile", "transportprotocol", "TCP"),
					resource.TestCheckResourceAttr("citrixadc_lsnappsprofile.tf_lsnappsprofile", "mapping", "ENDPOINT-INDEPENDENT"),
				),
			},
			{
				Config: testAccLsnappsprofile_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsnappsprofileExist("citrixadc_lsnappsprofile.tf_lsnappsprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_lsnappsprofile.tf_lsnappsprofile", "appsprofilename", "my_lsn_appsprofile"),
					resource.TestCheckResourceAttr("citrixadc_lsnappsprofile.tf_lsnappsprofile", "transportprotocol", "TCP"),
					resource.TestCheckResourceAttr("citrixadc_lsnappsprofile.tf_lsnappsprofile", "mapping", "ADDRESS-DEPENDENT"),
				),
			},
		},
	})
}

func testAccCheckLsnappsprofileExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lsnappsprofile name is set")
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
		data, err := client.FindResource("lsnappsprofile", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("lsnappsprofile %s not found", n)
		}

		return nil
	}
}

func testAccCheckLsnappsprofileDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_lsnappsprofile" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("lsnappsprofile", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("lsnappsprofile %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccLsnappsprofileDataSource_basic = `

resource "citrixadc_lsnappsprofile" "tf_lsnappsprofile_ds" {
	appsprofilename   = "my_lsn_appsprofile_ds"
	transportprotocol = "TCP"
	mapping           = "ENDPOINT-INDEPENDENT"
	filtering         = "ENDPOINT-INDEPENDENT"
	ippooling         = "RANDOM"
}

data "citrixadc_lsnappsprofile" "tf_lsnappsprofile_ds" {
	appsprofilename = citrixadc_lsnappsprofile.tf_lsnappsprofile_ds.appsprofilename
}
`

func TestAccLsnappsprofileDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccLsnappsprofileDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_lsnappsprofile.tf_lsnappsprofile_ds", "appsprofilename", "my_lsn_appsprofile_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_lsnappsprofile.tf_lsnappsprofile_ds", "transportprotocol", "TCP"),
					resource.TestCheckResourceAttr("data.citrixadc_lsnappsprofile.tf_lsnappsprofile_ds", "mapping", "ENDPOINT-INDEPENDENT"),
					resource.TestCheckResourceAttr("data.citrixadc_lsnappsprofile.tf_lsnappsprofile_ds", "filtering", "ENDPOINT-INDEPENDENT"),
					resource.TestCheckResourceAttr("data.citrixadc_lsnappsprofile.tf_lsnappsprofile_ds", "ippooling", "RANDOM"),
				),
			},
		},
	})
}
