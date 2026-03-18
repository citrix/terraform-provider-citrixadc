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

const testAccNsicapprofileDataSource_basic = `
	resource "citrixadc_nsicapprofile" "tf_nsicapprofile" {
		name             = "tf_nsicapprofile_ds"
		uri              = "/avscan"
		mode             = "REQMOD"
		reqtimeout       = 30
		reqtimeoutaction = "RESET"
		preview          = "ENABLED"
		previewlength    = 2048
		allow204         = "ENABLED"
		connectionkeepalive = "ENABLED"
	}

	data "citrixadc_nsicapprofile" "nsicapprofile_data" {
		name = citrixadc_nsicapprofile.tf_nsicapprofile.name
	}
`

func TestAccNsicapprofile_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNsicapprofileDestroy,
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

func TestAccNsicapprofileDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNsicapprofileDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_nsicapprofile.nsicapprofile_data", "name", "tf_nsicapprofile_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_nsicapprofile.nsicapprofile_data", "uri", "/avscan"),
					resource.TestCheckResourceAttr("data.citrixadc_nsicapprofile.nsicapprofile_data", "mode", "REQMOD"),
					resource.TestCheckResourceAttr("data.citrixadc_nsicapprofile.nsicapprofile_data", "preview", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_nsicapprofile.nsicapprofile_data", "previewlength", "2048"),
					resource.TestCheckResourceAttr("data.citrixadc_nsicapprofile.nsicapprofile_data", "allow204", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_nsicapprofile.nsicapprofile_data", "connectionkeepalive", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_nsicapprofile.nsicapprofile_data", "reqtimeout", "30"),
					resource.TestCheckResourceAttr("data.citrixadc_nsicapprofile.nsicapprofile_data", "reqtimeoutaction", "RESET"),
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

		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource("nsicapprofile", rs.Primary.ID)

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
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_nsicapprofile" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("nsicapprofile", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("nsicapprofile %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
