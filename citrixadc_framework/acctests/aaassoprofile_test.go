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

const testAccAaassoprofile_basic = `

	resource "citrixadc_aaassoprofile" "tf_aaassoprofile" {
		name = "myssoprofile"
		username = "john"
		password = "my_password"
	}
`
const testAccAaassoprofile_update = `

	resource "citrixadc_aaassoprofile" "tf_aaassoprofile" {
		name = "myssoprofile"
		username = "maria"
		password = "my_password2"
	}
`

func TestAccAaassoprofile_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAaassoprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAaassoprofile_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaassoprofileExist("citrixadc_aaassoprofile.tf_aaassoprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_aaassoprofile.tf_aaassoprofile", "name", "myssoprofile"),
					resource.TestCheckResourceAttr("citrixadc_aaassoprofile.tf_aaassoprofile", "username", "john"),
					resource.TestCheckResourceAttr("citrixadc_aaassoprofile.tf_aaassoprofile", "password", "my_password"),
				),
			},
			{
				Config: testAccAaassoprofile_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaassoprofileExist("citrixadc_aaassoprofile.tf_aaassoprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_aaassoprofile.tf_aaassoprofile", "name", "myssoprofile"),
					resource.TestCheckResourceAttr("citrixadc_aaassoprofile.tf_aaassoprofile", "username", "maria"),
					resource.TestCheckResourceAttr("citrixadc_aaassoprofile.tf_aaassoprofile", "password", "my_password2"),
				),
			},
		},
	})
}

func testAccCheckAaassoprofileExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No aaassoprofile name is set")
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
		data, err := client.FindResource("aaassoprofile", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("aaassoprofile %s not found", n)
		}

		return nil
	}
}

func testAccCheckAaassoprofileDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_aaassoprofile" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("aaassoprofile", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("aaassoprofile %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccAaassoprofileDataSource_basic = `

	resource "citrixadc_aaassoprofile" "tf_aaassoprofile" {
		name = "myssoprofile"
		username = "john"
		password = "my_password"
	}

	data "citrixadc_aaassoprofile" "tf_aaassoprofile" {
		name = citrixadc_aaassoprofile.tf_aaassoprofile.name
	}
`

func TestAccAaassoprofileDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccAaassoprofileDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_aaassoprofile.tf_aaassoprofile", "name", "myssoprofile"),
					resource.TestCheckResourceAttr("data.citrixadc_aaassoprofile.tf_aaassoprofile", "username", "john"),
					// password is not checked as it may be returned encrypted by the API
				),
			},
		},
	})
}
