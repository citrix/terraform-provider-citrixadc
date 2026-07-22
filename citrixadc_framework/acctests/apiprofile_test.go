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

const testAccApiprofile_basic_step1 = `
resource "citrixadc_apiprofile" "tf_apiprofile" {
  name          = "test_apiprofile"
  apivisibility = "ENABLED"
}

`

const testAccApiprofile_basic_step2 = `
resource "citrixadc_apiprofile" "tf_apiprofile" {
  name          = "test_apiprofile"
  apivisibility = "DISABLED"
}

`

func TestAccApiprofile_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckApiprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccApiprofile_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckApiprofileExist("citrixadc_apiprofile.tf_apiprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_apiprofile.tf_apiprofile", "name", "test_apiprofile"),
					resource.TestCheckResourceAttr("citrixadc_apiprofile.tf_apiprofile", "apivisibility", "ENABLED"),
				),
			},
			{
				Config: testAccApiprofile_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckApiprofileExist("citrixadc_apiprofile.tf_apiprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_apiprofile.tf_apiprofile", "name", "test_apiprofile"),
					resource.TestCheckResourceAttr("citrixadc_apiprofile.tf_apiprofile", "apivisibility", "DISABLED"),
				),
			},
		},
	})
}

func TestAccApiprofile_import(t *testing.T) {
	const resAddr = "citrixadc_apiprofile.tf_apiprofile"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckApiprofileDestroy,
		Steps: []resource.TestStep{
			{Config: testAccApiprofile_basic_step1},
			{
				Config:                  testAccApiprofile_basic_step1,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccCheckApiprofileExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No apiprofile name is set")
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
		data, err := client.FindResource(service.Apiprofile.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("apiprofile %s not found", n)
		}

		return nil
	}
}

func testAccCheckApiprofileDestroy(s *terraform.State) error {
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_apiprofile" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Apiprofile.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("apiprofile %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccApiprofileDataSource_basic = `

resource "citrixadc_apiprofile" "tf_apiprofile" {
  name          = "test_apiprofile"
  apivisibility = "ENABLED"
}

data "citrixadc_apiprofile" "tf_apiprofile" {
  name       = citrixadc_apiprofile.tf_apiprofile.name
  depends_on = [citrixadc_apiprofile.tf_apiprofile]
}
`

func TestAccApiprofileDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccApiprofileDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_apiprofile.tf_apiprofile", "name", "test_apiprofile"),
					resource.TestCheckResourceAttr("data.citrixadc_apiprofile.tf_apiprofile", "apivisibility", "ENABLED"),
				),
			},
		},
	})
}
