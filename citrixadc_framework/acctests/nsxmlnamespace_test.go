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
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"testing"
)

const testAccNsxmlnamespace_add = `
	resource "citrixadc_nsxmlnamespace" "tf_nsxmlnamespace" {
		prefix      = "tf_nsxmlnamespace"
		namespace   = "http://www.w3.org/2001/04/xmlenc#"
		description = "Description"
	}
`
const testAccNsxmlnamespace_update = `
	resource "citrixadc_nsxmlnamespace" "tf_nsxmlnamespace" {
		prefix      = "tf_nsxmlnamespace"
		namespace   = "http://www.w3.org/2001/04/xmlenc#"
		description = "Description_sample"
	}
`

func TestAccNsxmlnamespace_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNsxmlnamespaceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNsxmlnamespace_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsxmlnamespaceExist("citrixadc_nsxmlnamespace.tf_nsxmlnamespace", nil),
					resource.TestCheckResourceAttr("citrixadc_nsxmlnamespace.tf_nsxmlnamespace", "prefix", "tf_nsxmlnamespace"),
					resource.TestCheckResourceAttr("citrixadc_nsxmlnamespace.tf_nsxmlnamespace", "description", "Description"),
				),
			},
			{
				Config: testAccNsxmlnamespace_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsxmlnamespaceExist("citrixadc_nsxmlnamespace.tf_nsxmlnamespace", nil),
					resource.TestCheckResourceAttr("citrixadc_nsxmlnamespace.tf_nsxmlnamespace", "prefix", "tf_nsxmlnamespace"),
					resource.TestCheckResourceAttr("citrixadc_nsxmlnamespace.tf_nsxmlnamespace", "description", "Description_sample"),
				),
			},
		},
	})
}

func testAccCheckNsxmlnamespaceExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No nsxmlnamespace name is set")
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
		data, err := client.FindResource(service.Nsxmlnamespace.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("nsxmlnamespace %s not found", n)
		}

		return nil
	}
}

func testAccCheckNsxmlnamespaceDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_nsxmlnamespace" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Nsxmlnamespace.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("nsxmlnamespace %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
