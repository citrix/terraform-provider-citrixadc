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

const testAccContentinspectioncallout_basic = `

resource "citrixadc_nsicapprofile" "tf_nsicapprofile" {
	name             = "new-profile"
	uri              = "/example"
	mode             = "REQMOD"
	reqtimeout       = 4
	reqtimeoutaction = "RESET"
	preview          = "ENABLED"
	previewlength    = 4096
}
resource "citrixadc_contentinspectioncallout" "tf_contentinspectioncalloout" {
	name        = "my_ci_callout"
	type        = "ICAP"
	profilename = citrixadc_nsicapprofile.tf_nsicapprofile.name
	serverip    = "2.2.2.2"
	returntype  = "TEXT"
	resultexpr  = "true"
	}
`

const testAccContentinspectioncallout_update = `

resource "citrixadc_nsicapprofile" "tf_nsicapprofile" {
	name             = "new-profile"
	uri              = "/example"
	mode             = "REQMOD"
	reqtimeout       = 4
	reqtimeoutaction = "RESET"
	preview          = "ENABLED"
	previewlength    = 4096
}
resource "citrixadc_contentinspectioncallout" "tf_contentinspectioncalloout" {
	name        = "my_ci_callout"
	type        = "ICAP"
	profilename = citrixadc_nsicapprofile.tf_nsicapprofile.name
	serverip    = "2.2.2.2"
	returntype  = "TEXT"
	resultexpr  = "icap.res.header(\"ISTag\")"
	}
`

func TestAccContentinspectioncallout_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckContentinspectioncalloutDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccContentinspectioncallout_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckContentinspectioncalloutExist("citrixadc_contentinspectioncallout.tf_contentinspectioncalloout", nil),
					resource.TestCheckResourceAttr("citrixadc_contentinspectioncallout.tf_contentinspectioncalloout", "name", "my_ci_callout"),
					resource.TestCheckResourceAttr("citrixadc_contentinspectioncallout.tf_contentinspectioncalloout", "type", "ICAP"),
					resource.TestCheckResourceAttr("citrixadc_contentinspectioncallout.tf_contentinspectioncalloout", "profilename", "new-profile"),
					resource.TestCheckResourceAttr("citrixadc_contentinspectioncallout.tf_contentinspectioncalloout", "serverip", "2.2.2.2"),
					resource.TestCheckResourceAttr("citrixadc_contentinspectioncallout.tf_contentinspectioncalloout", "returntype", "TEXT"),
					resource.TestCheckResourceAttr("citrixadc_contentinspectioncallout.tf_contentinspectioncalloout", "resultexpr", "true"),
				),
			},
			{
				Config: testAccContentinspectioncallout_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckContentinspectioncalloutExist("citrixadc_contentinspectioncallout.tf_contentinspectioncalloout", nil),
					resource.TestCheckResourceAttr("citrixadc_contentinspectioncallout.tf_contentinspectioncalloout", "name", "my_ci_callout"),
					resource.TestCheckResourceAttr("citrixadc_contentinspectioncallout.tf_contentinspectioncalloout", "type", "ICAP"),
					resource.TestCheckResourceAttr("citrixadc_contentinspectioncallout.tf_contentinspectioncalloout", "profilename", "new-profile"),
					resource.TestCheckResourceAttr("citrixadc_contentinspectioncallout.tf_contentinspectioncalloout", "serverip", "2.2.2.2"),
					resource.TestCheckResourceAttr("citrixadc_contentinspectioncallout.tf_contentinspectioncalloout", "returntype", "TEXT"),
					resource.TestCheckResourceAttr("citrixadc_contentinspectioncallout.tf_contentinspectioncalloout", "resultexpr", "icap.res.header(\"ISTag\")"),
				),
			},
		},
	})
}

func testAccCheckContentinspectioncalloutExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No contentinspectioncallout name is set")
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
		data, err := client.FindResource("contentinspectioncallout", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("contentinspectioncallout %s not found", n)
		}

		return nil
	}
}

func testAccCheckContentinspectioncalloutDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_contentinspectioncallout" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("contentinspectioncallout", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("contentinspectioncallout %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccContentinspectioncalloutDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccContentinspectioncalloutDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_contentinspectioncallout.tf_contentinspectioncallout_ds", "name", "my_ci_callout_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_contentinspectioncallout.tf_contentinspectioncallout_ds", "type", "ICAP"),
					resource.TestCheckResourceAttr("data.citrixadc_contentinspectioncallout.tf_contentinspectioncallout_ds", "profilename", "new-profile-ds"),
					resource.TestCheckResourceAttr("data.citrixadc_contentinspectioncallout.tf_contentinspectioncallout_ds", "serverip", "2.2.2.2"),
					resource.TestCheckResourceAttr("data.citrixadc_contentinspectioncallout.tf_contentinspectioncallout_ds", "returntype", "TEXT"),
					resource.TestCheckResourceAttr("data.citrixadc_contentinspectioncallout.tf_contentinspectioncallout_ds", "resultexpr", "true"),
				),
			},
		},
	})
}

const testAccContentinspectioncalloutDataSource_basic = `

resource "citrixadc_nsicapprofile" "tf_nsicapprofile_ds" {
	name             = "new-profile-ds"
	uri              = "/example"
	mode             = "REQMOD"
	reqtimeout       = 4
	reqtimeoutaction = "RESET"
	preview          = "ENABLED"
	previewlength    = 4096
}

resource "citrixadc_contentinspectioncallout" "tf_contentinspectioncallout_ds" {
	name        = "my_ci_callout_ds"
	type        = "ICAP"
	profilename = citrixadc_nsicapprofile.tf_nsicapprofile_ds.name
	serverip    = "2.2.2.2"
	returntype  = "TEXT"
	resultexpr  = "true"
}

data "citrixadc_contentinspectioncallout" "tf_contentinspectioncallout_ds" {
	name = citrixadc_contentinspectioncallout.tf_contentinspectioncallout_ds.name
	depends_on = [citrixadc_contentinspectioncallout.tf_contentinspectioncallout_ds]
}

`
