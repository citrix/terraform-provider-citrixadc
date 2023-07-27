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

const testAccNsicapprofile_add = `
	resource "citrixadc_nsicapprofile" "tf_nsicapprofile" {
		name             = "tf_nsicapprofile"
		uri              = "/example"
		mode             = "REQMOD"
		reqtimeout       = 4
		reqtimeoutaction = "RESET"
		preview          = "ENABLED"
		previewlength    = 4096
	}
`
const testAccNsicapprofile_update = `
	resource "citrixadc_nsicapprofile" "tf_nsicapprofile" {
		name             = "tf_nsicapprofile"
		uri              = "/hello"
		mode             = "REQMOD"
		reqtimeout       = 4
		reqtimeoutaction = "RESET"
		preview          = "DISABLED"
		previewlength    = 4096
	}
`

func TestAccNsicapprofile_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNsicapprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNsicapprofile_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsicapprofileExist("citrixadc_nsicapprofile.tf_nsicapprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_nsicapprofile.tf_nsicapprofile", "name", "tf_nsicapprofile"),
					resource.TestCheckResourceAttr("citrixadc_nsicapprofile.tf_nsicapprofile", "uri", "/example"),
					resource.TestCheckResourceAttr("citrixadc_nsicapprofile.tf_nsicapprofile", "preview", "ENABLED"),
				),
			},
			{
				Config: testAccNsicapprofile_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsicapprofileExist("citrixadc_nsicapprofile.tf_nsicapprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_nsicapprofile.tf_nsicapprofile", "name", "tf_nsicapprofile"),
					resource.TestCheckResourceAttr("citrixadc_nsicapprofile.tf_nsicapprofile", "uri", "/hello"),
					resource.TestCheckResourceAttr("citrixadc_nsicapprofile.tf_nsicapprofile", "preview", "DISABLED"),
				),
			},
		},
	})
}

func testAccCheckNsicapprofileExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No nsicapprofile name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource("nsicapprofile", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("nsicapprofile %s not found", n)
		}

		return nil
	}
}

func testAccCheckNsicapprofileDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_nsicapprofile" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource("nsicapprofile", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("nsicapprofile %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
