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

const testAccCloudallowedngsticketprofile_basic_step1 = `
resource "citrixadc_cloudallowedngsticketprofile" "tf_cloudallowedngsticketprofile" {
  name    = "tf_allowedticket"
  creator = "test_creator"
}
`

const testAccCloudallowedngsticketprofile_basic_step2 = `
resource "citrixadc_cloudallowedngsticketprofile" "tf_cloudallowedngsticketprofile" {
  name    = "tf_allowedticket"
  creator = "test_creator_updated"
}
`

func TestAccCloudallowedngsticketprofile_basic(t *testing.T) {
	t.Skip("TODO: Requires review")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudallowedngsticketprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudallowedngsticketprofile_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudallowedngsticketprofileExist("citrixadc_cloudallowedngsticketprofile.tf_cloudallowedngsticketprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_cloudallowedngsticketprofile.tf_cloudallowedngsticketprofile", "name", "tf_allowedticket"),
					resource.TestCheckResourceAttr("citrixadc_cloudallowedngsticketprofile.tf_cloudallowedngsticketprofile", "creator", "test_creator"),
				),
			},
			{
				Config: testAccCloudallowedngsticketprofile_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudallowedngsticketprofileExist("citrixadc_cloudallowedngsticketprofile.tf_cloudallowedngsticketprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_cloudallowedngsticketprofile.tf_cloudallowedngsticketprofile", "name", "tf_allowedticket"),
					resource.TestCheckResourceAttr("citrixadc_cloudallowedngsticketprofile.tf_cloudallowedngsticketprofile", "creator", "test_creator_updated"),
				),
			},
		},
	})
}

func testAccCheckCloudallowedngsticketprofileExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No cloudallowedngsticketprofile name is set")
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
		data, err := client.FindResource(service.Cloudallowedngsticketprofile.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("cloudallowedngsticketprofile %s not found", n)
		}

		return nil
	}
}

func testAccCheckCloudallowedngsticketprofileDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_cloudallowedngsticketprofile" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Cloudallowedngsticketprofile.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("cloudallowedngsticketprofile %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccCloudallowedngsticketprofileDataSource_basic = `

resource "citrixadc_cloudallowedngsticketprofile" "tf_cloudallowedngsticketprofile" {
  name    = "tf_allowedticket"
  creator = "test_creator"
}

data "citrixadc_cloudallowedngsticketprofile" "tf_cloudallowedngsticketprofile_data" {
  name       = citrixadc_cloudallowedngsticketprofile.tf_cloudallowedngsticketprofile.name
  depends_on = [citrixadc_cloudallowedngsticketprofile.tf_cloudallowedngsticketprofile]
}
`

func TestAccCloudallowedngsticketprofileDataSource_basic(t *testing.T) {
	t.Skip("TODO: Requires review")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudallowedngsticketprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudallowedngsticketprofileDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_cloudallowedngsticketprofile.tf_cloudallowedngsticketprofile_data", "name", "tf_allowedticket"),
					resource.TestCheckResourceAttr("data.citrixadc_cloudallowedngsticketprofile.tf_cloudallowedngsticketprofile_data", "creator", "test_creator"),
				),
			},
		},
	})
}
