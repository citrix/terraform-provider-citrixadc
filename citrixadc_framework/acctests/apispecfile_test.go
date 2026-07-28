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

// TODO_PLACEHOLDER: Replace src with a real URL accessible by the ADC appliance.
const testAccApispecfile_basic_step1 = `
resource "citrixadc_apispecfile" "tf_apispecfile" {
  name      = "test_apispecfile"
  src       = "local://sample_apispec.yaml"
  overwrite = true
}

`

// Step 2 changes overwrite — triggers resource replacement due to RequiresReplace on all attrs.
// TODO_PLACEHOLDER: Replace src with a real URL accessible by the ADC appliance.
const testAccApispecfile_basic_step2 = `
resource "citrixadc_apispecfile" "tf_apispecfile" {
  name      = "test_apispecfile"
  src       = "local://sample_apispec.yaml"
  overwrite = false
}

`

func TestAccApispecfile_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doApiSpecPreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckApispecfileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccApispecfile_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckApispecfileExist("citrixadc_apispecfile.tf_apispecfile", nil),
					resource.TestCheckResourceAttr("citrixadc_apispecfile.tf_apispecfile", "name", "test_apispecfile"),
					resource.TestCheckResourceAttr("citrixadc_apispecfile.tf_apispecfile", "overwrite", "true"),
				),
			},
			{
				Config: testAccApispecfile_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckApispecfileExist("citrixadc_apispecfile.tf_apispecfile", nil),
					resource.TestCheckResourceAttr("citrixadc_apispecfile.tf_apispecfile", "name", "test_apispecfile"),
					resource.TestCheckResourceAttr("citrixadc_apispecfile.tf_apispecfile", "overwrite", "false"),
				),
			},
		},
	})
}

func TestAccApispecfile_import(t *testing.T) {
	const resAddr = "citrixadc_apispecfile.tf_apispecfile"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doApiSpecPreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckApispecfileDestroy,
		Steps: []resource.TestStep{
			{Config: testAccApispecfile_basic_step1},
			{
				Config:            testAccApispecfile_basic_step1,
				ResourceName:      resAddr,
				ImportState:       true,
				ImportStateVerify: true,
				// `src` and `overwrite` are write-only Import-action inputs that
				// NITRO does not echo back on GET, so they cannot round-trip
				// through import.
				ImportStateVerifyIgnore: []string{"overwrite", "src"},
			},
		},
	})
}

func testAccCheckApispecfileExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No apispecfile name is set")
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
		data, err := client.FindResource(service.Apispecfile.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("apispecfile %s not found", n)
		}

		return nil
	}
}

func testAccCheckApispecfileDestroy(s *terraform.State) error {
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_apispecfile" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No apispecfile name is set")
		}

		_, err := client.FindResource(service.Apispecfile.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("apispecfile %s still exists", rs.Primary.ID)
		}
	}

	return nil
}

// TODO_PLACEHOLDER: Replace src with a real URL accessible by the ADC appliance.
const testAccApispecfileDataSource_basic = `

resource "citrixadc_apispecfile" "tf_apispecfile" {
  name      = "test_apispecfile"
  src       = "local://sample_apispec.yaml"
  overwrite = true
}

data "citrixadc_apispecfile" "tf_apispecfile" {
  name       = citrixadc_apispecfile.tf_apispecfile.name
  depends_on = [citrixadc_apispecfile.tf_apispecfile]
}
`

func TestAccApispecfileDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doApiSpecPreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccApispecfileDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_apispecfile.tf_apispecfile", "name", "test_apispecfile"),
				),
			},
		},
	})
}
