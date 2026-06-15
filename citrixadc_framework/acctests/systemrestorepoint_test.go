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

// systemrestorepoint is an action-create + plain-delete + get resource.
// filename is Required and RequiresReplace, so there is no update step.
//
// CAVEATS:
//   - The appliance enforces a MAXIMUM of 3 restore points. If create fails with
//     a cap error, remove stale tf-created restore points (e.g. via NITRO:
//     curl -X DELETE -u nsroot:<pw> http://<ns>/nitro/v1/config/systemrestorepoint/tf_restorepoint_test).
//   - Creating a restore point snapshots the config + a tech-support bundle; it is
//     resource-intensive and may take time, hence the generous timeout in the run command.

const testAccSystemrestorepoint_basic = `

	resource "citrixadc_systemrestorepoint" "tf_systemrestorepoint" {
		filename = "tf_restorepoint_test"
	}
`

func TestAccSystemrestorepoint_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSystemrestorepointDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSystemrestorepoint_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSystemrestorepointExist("citrixadc_systemrestorepoint.tf_systemrestorepoint", nil),
					resource.TestCheckResourceAttr("citrixadc_systemrestorepoint.tf_systemrestorepoint", "filename", "tf_restorepoint_test"),
				),
			},
		},
	})
}

func testAccCheckSystemrestorepointExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No systemrestorepoint filename is set")
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
		data, err := client.FindResource(service.Systemrestorepoint.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("systemrestorepoint %s not found", n)
		}

		return nil
	}
}

func testAccCheckSystemrestorepointDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_systemrestorepoint" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No filename is set")
		}

		_, err := client.FindResource(service.Systemrestorepoint.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("systemrestorepoint %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccSystemrestorepoint_DataSource_basic = `

	resource "citrixadc_systemrestorepoint" "tf_systemrestorepoint" {
		filename = "tf_restorepoint_test"
	}

	data "citrixadc_systemrestorepoint" "tf_systemrestorepoint_data" {
		filename   = citrixadc_systemrestorepoint.tf_systemrestorepoint.filename
		depends_on = [citrixadc_systemrestorepoint.tf_systemrestorepoint]
	}
`

func TestAccSystemrestorepointDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSystemrestorepointDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSystemrestorepoint_DataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_systemrestorepoint.tf_systemrestorepoint_data", "filename", "tf_restorepoint_test"),
				),
			},
		},
	})
}
