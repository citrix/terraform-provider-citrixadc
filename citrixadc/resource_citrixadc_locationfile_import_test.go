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

const testAccLocationfileImport_basic = `

resource "citrixadc_locationfile_import" "tf_locationfile_import" {
	locationfile = "location_file"
	src          = "local://my_location_file"
	}
  
`

func TestAccLocationfileImport_basic(t *testing.T) {
	t.Skip("TODO: Need to find a way to test this resource!")
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      nil,
		Steps: []resource.TestStep{
			{
				Config: testAccLocationfileImport_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLocationfileImportExist("citrixadc_locationfile_import.tf_locationfile_import", nil),
					resource.TestCheckResourceAttr("citrixadc_locationfile_import.tf_locationfile_import", "locationfile", "location_file"),
					resource.TestCheckResourceAttr("citrixadc_locationfile_import.tf_locationfile_import", "src", "local://my_location_file"),
				),
			},
		},
	})
}

func testAccCheckLocationfileImportExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No locationfile name is set")
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
		data, err := client.FindResource(service.Locationfile.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("locationfile %s not found", n)
		}

		return nil
	}
}
