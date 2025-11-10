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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"testing"
)

const testAccContentinspectionprofile_basic = `

	resource "citrixadc_contentinspectionprofile" "tf_contentinspectionprofile" {
		name             = "my_ci_profile"
		type             = "InlineInspection"
		ingressinterface = "LA/2"
		egressinterface  = "LA/3"
	}
`
const testAccContentinspectionprofile_update = `

	resource "citrixadc_contentinspectionprofile" "tf_contentinspectionprofile" {
		name             = "my_ci_profile"
		type             = "InlineInspection"
		ingressinterface = "LA/3"
		egressinterface  = "LA/2"
	}
`

func TestAccContentinspectionprofile_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckContentinspectionprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccContentinspectionprofile_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckContentinspectionprofileExist("citrixadc_contentinspectionprofile.tf_contentinspectionprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_contentinspectionprofile.tf_contentinspectionprofile", "name", "my_ci_profile"),
					resource.TestCheckResourceAttr("citrixadc_contentinspectionprofile.tf_contentinspectionprofile", "type", "InlineInspection"),
					resource.TestCheckResourceAttr("citrixadc_contentinspectionprofile.tf_contentinspectionprofile", "ingressinterface", "LA/2"),
					resource.TestCheckResourceAttr("citrixadc_contentinspectionprofile.tf_contentinspectionprofile", "egressinterface", "LA/3"),
				),
			},
			{
				Config: testAccContentinspectionprofile_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckContentinspectionprofileExist("citrixadc_contentinspectionprofile.tf_contentinspectionprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_contentinspectionprofile.tf_contentinspectionprofile", "name", "my_ci_profile"),
					resource.TestCheckResourceAttr("citrixadc_contentinspectionprofile.tf_contentinspectionprofile", "type", "InlineInspection"),
					resource.TestCheckResourceAttr("citrixadc_contentinspectionprofile.tf_contentinspectionprofile", "ingressinterface", "LA/3"),
					resource.TestCheckResourceAttr("citrixadc_contentinspectionprofile.tf_contentinspectionprofile", "egressinterface", "LA/2"),
				),
			},
		},
	})
}

func testAccCheckContentinspectionprofileExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No contentinspectionprofile name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		// Use the shared utility function to get a configured client
		client, err := testAccGetClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource("contentinspectionprofile", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("contentinspectionprofile %s not found", n)
		}

		return nil
	}
}

func testAccCheckContentinspectionprofileDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_contentinspectionprofile" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("contentinspectionprofile", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("contentinspectionprofile %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
