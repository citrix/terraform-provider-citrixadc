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

const testAccAuthenticationsmartaccessprofile_basic_step1 = `
resource "citrixadc_authenticationsmartaccessprofile" "tf_authenticationsmartaccessprofile" {
  name    = "tf_authenticationsmartaccessprofile"
  tags    = "test_tag"
  comment = "test_comment"
}

`

const testAccAuthenticationsmartaccessprofile_basic_step2 = `
resource "citrixadc_authenticationsmartaccessprofile" "tf_authenticationsmartaccessprofile" {
  name    = "tf_authenticationsmartaccessprofile"
  tags    = "test_tag_updated"
  comment = "test_comment_updated"
}

`

func TestAccAuthenticationsmartaccessprofile_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationsmartaccessprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationsmartaccessprofile_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationsmartaccessprofileExist("citrixadc_authenticationsmartaccessprofile.tf_authenticationsmartaccessprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationsmartaccessprofile.tf_authenticationsmartaccessprofile", "name", "tf_authenticationsmartaccessprofile"),
					resource.TestCheckResourceAttr("citrixadc_authenticationsmartaccessprofile.tf_authenticationsmartaccessprofile", "tags", "test_tag"),
					resource.TestCheckResourceAttr("citrixadc_authenticationsmartaccessprofile.tf_authenticationsmartaccessprofile", "comment", "test_comment"),
				),
			},
			{
				Config: testAccAuthenticationsmartaccessprofile_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationsmartaccessprofileExist("citrixadc_authenticationsmartaccessprofile.tf_authenticationsmartaccessprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationsmartaccessprofile.tf_authenticationsmartaccessprofile", "name", "tf_authenticationsmartaccessprofile"),
					resource.TestCheckResourceAttr("citrixadc_authenticationsmartaccessprofile.tf_authenticationsmartaccessprofile", "tags", "test_tag_updated"),
					resource.TestCheckResourceAttr("citrixadc_authenticationsmartaccessprofile.tf_authenticationsmartaccessprofile", "comment", "test_comment_updated"),
				),
			},
		},
	})
}

func testAccCheckAuthenticationsmartaccessprofileExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No authenticationsmartaccessprofile name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Authenticationsmartaccessprofile.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("authenticationsmartaccessprofile %s not found", n)
		}

		return nil
	}
}

func testAccCheckAuthenticationsmartaccessprofileDestroy(s *terraform.State) error {
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_authenticationsmartaccessprofile" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No authenticationsmartaccessprofile name is set")
		}

		_, err := client.FindResource(service.Authenticationsmartaccessprofile.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("authenticationsmartaccessprofile %s still exists", rs.Primary.ID)
		}
	}

	return nil
}

const testAccAuthenticationsmartaccessprofileDataSource_basic = `

resource "citrixadc_authenticationsmartaccessprofile" "tf_authenticationsmartaccessprofile" {
  name    = "tf_authenticationsmartaccessprofile"
  tags    = "test_tag"
  comment = "test_comment"
}

data "citrixadc_authenticationsmartaccessprofile" "tf_authenticationsmartaccessprofile" {
  name       = citrixadc_authenticationsmartaccessprofile.tf_authenticationsmartaccessprofile.name
  depends_on = [citrixadc_authenticationsmartaccessprofile.tf_authenticationsmartaccessprofile]
}
`

func TestAccAuthenticationsmartaccessprofileDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationsmartaccessprofileDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_authenticationsmartaccessprofile.tf_authenticationsmartaccessprofile", "name", "tf_authenticationsmartaccessprofile"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationsmartaccessprofile.tf_authenticationsmartaccessprofile", "tags", "test_tag"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationsmartaccessprofile.tf_authenticationsmartaccessprofile", "comment", "test_comment"),
				),
			},
		},
	})
}
