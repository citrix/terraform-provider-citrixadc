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
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"testing"
)

const testAccDbdbprofile_basic = `
	resource "citrixadc_dbdbprofile" "tf_dbdbprofile" {
		name           = "my_dbprofile"
		stickiness     = "YES"
		conmultiplex   = "ENABLED"
		interpretquery = "YES"
	}
`
const testAccDbdbprofile_update = `
	resource "citrixadc_dbdbprofile" "tf_dbdbprofile" {
		name           = "my_dbprofile"
		stickiness     = "NO"
		conmultiplex   = "DISABLED"
		interpretquery = "NO"
	}
`

func TestAccDbdbprofile_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDbdbprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDbdbprofile_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDbdbprofileExist("citrixadc_dbdbprofile.tf_dbdbprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_dbdbprofile.tf_dbdbprofile", "name", "my_dbprofile"),
					resource.TestCheckResourceAttr("citrixadc_dbdbprofile.tf_dbdbprofile", "stickiness", "YES"),
					resource.TestCheckResourceAttr("citrixadc_dbdbprofile.tf_dbdbprofile", "conmultiplex", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_dbdbprofile.tf_dbdbprofile", "interpretquery", "YES"),
				),
			},
			{
				Config: testAccDbdbprofile_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDbdbprofileExist("citrixadc_dbdbprofile.tf_dbdbprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_dbdbprofile.tf_dbdbprofile", "name", "my_dbprofile"),
					resource.TestCheckResourceAttr("citrixadc_dbdbprofile.tf_dbdbprofile", "stickiness", "NO"),
					resource.TestCheckResourceAttr("citrixadc_dbdbprofile.tf_dbdbprofile", "conmultiplex", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_dbdbprofile.tf_dbdbprofile", "interpretquery", "NO"),
				),
			},
		},
	})
}

func testAccCheckDbdbprofileExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No dbdbprofile name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(service.Dbdbprofile.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("dbdbprofile %s not found", n)
		}

		return nil
	}
}

func testAccCheckDbdbprofileDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_dbdbprofile" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Dbdbprofile.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("dbdbprofile %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
