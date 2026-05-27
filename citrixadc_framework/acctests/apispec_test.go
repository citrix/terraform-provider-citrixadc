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

// TODO_PLACEHOLDER: Replace "TODO_PLACEHOLDER_SPEC_FILE" with the actual name of an API spec file
// that exists on the ADC appliance (e.g., "/nsconfig/apispec/my_spec.yaml").
// TODO_PLACEHOLDER: Replace "TODO_PLACEHOLDER_SPEC_FILE_V2" with another valid spec file for update testing.

const testAccApispec_basic_step1 = `
resource "citrixadc_apispecfile" "tf_apispecfile" {
  name      = "test_apispecfile"
  src       = "local://sample_apispec.yaml"
  overwrite = true
}

resource "citrixadc_apispec" "tf_apispec" {
  name = "tf_apispec"
  file = citrixadc_apispecfile.tf_apispecfile.name
  type = "OAS"
}

`

const testAccApispec_basic_step2 = `
resource "citrixadc_apispecfile" "tf_apispecfile" {
  name      = "test_apispecfile2"
  src       = "local://sample_apispec2.yaml"
  overwrite = true
}
resource "citrixadc_apispec" "tf_apispec" {
  name = "tf_apispec"
  file = citrixadc_apispecfile.tf_apispecfile.name
  type = "OAS"
}

`

func TestAccApispec_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doApiSpecPreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckApispecDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccApispec_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckApispecExist("citrixadc_apispec.tf_apispec", nil),
					resource.TestCheckResourceAttr("citrixadc_apispec.tf_apispec", "name", "tf_apispec"),
					resource.TestCheckResourceAttr("citrixadc_apispec.tf_apispec", "file", "test_apispecfile"),
					resource.TestCheckResourceAttr("citrixadc_apispec.tf_apispec", "type", "OAS"),
				),
			},
			{
				Config: testAccApispec_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckApispecExist("citrixadc_apispec.tf_apispec", nil),
					resource.TestCheckResourceAttr("citrixadc_apispec.tf_apispec", "name", "tf_apispec"),
					resource.TestCheckResourceAttr("citrixadc_apispec.tf_apispec", "file", "test_apispecfile2"),
					resource.TestCheckResourceAttr("citrixadc_apispec.tf_apispec", "type", "OAS"),
				),
			},
		},
	})
}

func testAccCheckApispecExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No apispec name is set")
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
		data, err := client.FindResource(service.Apispec.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("apispec %s not found", n)
		}

		return nil
	}
}

func testAccCheckApispecDestroy(s *terraform.State) error {
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_apispec" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No apispec name is set")
		}

		_, err := client.FindResource(service.Apispec.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("apispec %s still exists", rs.Primary.ID)
		}
	}

	return nil
}

const testAccApispecDataSource_basic = `
resource "citrixadc_apispecfile" "tf_apispecfile" {
  name      = "test_apispecfile"
  src       = "local://sample_apispec.yaml"
  overwrite = true
}

resource "citrixadc_apispec" "tf_apispec" {
  name = "tf_apispec"
  file = citrixadc_apispecfile.tf_apispecfile.name
  type = "OAS"
}

data "citrixadc_apispec" "tf_apispec" {
  name       = citrixadc_apispec.tf_apispec.name
  depends_on = [citrixadc_apispec.tf_apispec]
}
`

func TestAccApispecDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doApiSpecPreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccApispecDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_apispec.tf_apispec", "name", "tf_apispec"),
					resource.TestCheckResourceAttr("data.citrixadc_apispec.tf_apispec", "file", "test_apispecfile"),
					resource.TestCheckResourceAttr("data.citrixadc_apispec.tf_apispec", "type", "OAS"),
				),
			},
		},
	})
}
