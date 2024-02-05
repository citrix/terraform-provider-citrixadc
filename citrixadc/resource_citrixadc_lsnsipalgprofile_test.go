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
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"testing"
)

const testAccLsnsipalgprofile_basic = `

	resource "citrixadc_lsnsipalgprofile" "tf_lsnsipalgprofile" {
		sipalgprofilename      = "my_lsn_sipalgprofile"
		datasessionidletimeout = 150
		sipsessiontimeout      = 150
		registrationtimeout    = 150
		sipsrcportrange        = "4400"
		siptransportprotocol   = "TCP"
	}
  
`
const testAccLsnsipalgprofile_update = `

	resource "citrixadc_lsnsipalgprofile" "tf_lsnsipalgprofile" {
		sipalgprofilename      = "my_lsn_sipalgprofile"
		datasessionidletimeout = 100
		sipsessiontimeout      = 100
		registrationtimeout    = 100
		sipsrcportrange        = "4200"
		siptransportprotocol   = "TCP"
	}
  
`

func TestAccLsnsipalgprofile_basic(t *testing.T) {
	t.Skip("TODO: Need to find a way to test this LSN resource!")
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckLsnsipalgprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLsnsipalgprofile_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsnsipalgprofileExist("citrixadc_lsnsipalgprofile.tf_lsnsipalgprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_lsnsipalgprofile.tf_lsnsipalgprofile", "sipalgprofilename", "my_lsn_sipalgprofile"),
					resource.TestCheckResourceAttr("citrixadc_lsnsipalgprofile.tf_lsnsipalgprofile", "datasessionidletimeout", "150"),
					resource.TestCheckResourceAttr("citrixadc_lsnsipalgprofile.tf_lsnsipalgprofile", "sipsessiontimeout", "150"),
					resource.TestCheckResourceAttr("citrixadc_lsnsipalgprofile.tf_lsnsipalgprofile", "registrationtimeout", "150"),
					resource.TestCheckResourceAttr("citrixadc_lsnsipalgprofile.tf_lsnsipalgprofile", "sipsrcportrange", "4400"),
					resource.TestCheckResourceAttr("citrixadc_lsnsipalgprofile.tf_lsnsipalgprofile", "siptransportprotocol", "TCP"),
				),
			},
			{
				Config: testAccLsnsipalgprofile_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsnsipalgprofileExist("citrixadc_lsnsipalgprofile.tf_lsnsipalgprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_lsnsipalgprofile.tf_lsnsipalgprofile", "sipalgprofilename", "my_lsn_sipalgprofile"),
					resource.TestCheckResourceAttr("citrixadc_lsnsipalgprofile.tf_lsnsipalgprofile", "datasessionidletimeout", "100"),
					resource.TestCheckResourceAttr("citrixadc_lsnsipalgprofile.tf_lsnsipalgprofile", "sipsessiontimeout", "100"),
					resource.TestCheckResourceAttr("citrixadc_lsnsipalgprofile.tf_lsnsipalgprofile", "registrationtimeout", "100"),
					resource.TestCheckResourceAttr("citrixadc_lsnsipalgprofile.tf_lsnsipalgprofile", "sipsrcportrange", "4200"),
					resource.TestCheckResourceAttr("citrixadc_lsnsipalgprofile.tf_lsnsipalgprofile", "siptransportprotocol", "TCP"),
				),
			},
		},
	})
}

func testAccCheckLsnsipalgprofileExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lsnsipalgprofile name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource("lsnsipalgprofile", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("lsnsipalgprofile %s not found", n)
		}

		return nil
	}
}

func testAccCheckLsnsipalgprofileDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_lsnsipalgprofile" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource("lsnsipalgprofile", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("lsnsipalgprofile %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
