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

const testAccAnalyticsprofile_basic = `

	resource "citrixadc_analyticsprofile" "tf_analyticsprofile" {
		name             = "my_analyticsprofile"
		type             = "webinsight"
		httppagetracking = "DISABLED"
		httpurl          = "DISABLED"
	}
`
const testAccAnalyticsprofile_update = `

	resource "citrixadc_analyticsprofile" "tf_analyticsprofile" {
		name             = "my_analyticsprofile"
		type             = "webinsight"
		httppagetracking = "ENABLED"
		httpurl          = "ENABLED"
	}
`

func TestAccAnalyticsprofile_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAnalyticsprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAnalyticsprofile_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAnalyticsprofileExist("citrixadc_analyticsprofile.tf_analyticsprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_analyticsprofile.tf_analyticsprofile", "name", "my_analyticsprofile"),
					resource.TestCheckResourceAttr("citrixadc_analyticsprofile.tf_analyticsprofile", "type", "webinsight"),
					resource.TestCheckResourceAttr("citrixadc_analyticsprofile.tf_analyticsprofile", "httppagetracking", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_analyticsprofile.tf_analyticsprofile", "httpurl", "DISABLED"),
				),
			},
			{
				Config: testAccAnalyticsprofile_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAnalyticsprofileExist("citrixadc_analyticsprofile.tf_analyticsprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_analyticsprofile.tf_analyticsprofile", "name", "my_analyticsprofile"),
					resource.TestCheckResourceAttr("citrixadc_analyticsprofile.tf_analyticsprofile", "type", "webinsight"),
					resource.TestCheckResourceAttr("citrixadc_analyticsprofile.tf_analyticsprofile", "httppagetracking", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_analyticsprofile.tf_analyticsprofile", "httpurl", "ENABLED"),
				),
			},
		},
	})
}

func testAccCheckAnalyticsprofileExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No analyticsprofile name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource("analyticsprofile", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("analyticsprofile %s not found", n)
		}

		return nil
	}
}

func testAccCheckAnalyticsprofileDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_analyticsprofile" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource("analyticsprofile", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("analyticsprofile %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
